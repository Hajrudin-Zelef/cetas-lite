---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-1
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "incident", "open source"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [1, 187]
sha256: 3c0300e898bbd7b03d55379899dd6ccb5782b020c74673abc133c3eda5a773db
---

# PostgreSQL en production — Guide complet pour sysadmin

> **Public :** Zelef, chef de service systèmes & énergies, sysadmin.
> **Angle :** pratique, production, Debian/Ubuntu, versions 15/16/17.
> **Avertissement :** les valeurs de paramètres sont des points de départ à valider
> par vos tests de charge et vos fiches techniques. Aucun mot de passe réel
> n'apparaît dans ce guide : utilisez le coffre à secrets de votre entreprise
> (Ansible Vault, KeePass, gestionnaire de mots de passe d'équipe).

---

## Sommaire

| Partie | Sections | Thème |
|---|---|---|
| A | 1–12 | Prise en main : versions, architecture, installation |
| B | 13–22 | Connexion sécurisée : psql, rôles, pg_hba.conf, SSL |
| C | 23–35 | SQL opérationnel : DDL, DML, types, contraintes |
| D | 36–47 | Performance : index, EXPLAIN, VACUUM, transactions |
| E | 48–58 | Exploitation : sauvegardes, PITR, réplication, haute dispo |
| F | 59–68 | Tuning, monitoring, alertes |
| G | 69–75 | Mise à jour, dépannage, erreurs classiques |
| H | 76–85 | Cas pratiques, pense-bête, glossaire, quiz, ressources |

---

# Partie A — Prise en main

## 1. Pourquoi PostgreSQL en 2026

PostgreSQL (souvent abrégé **Postgres** ou **PG**) est le SGBDR open source
de référence en production : licence permissive (PostgreSQL Licence), ACID
strict, extensible (extensions comme PostGIS, pg_stat_statements, pg_trgm),
et disponible nativement sur Debian/Ubuntu.

**Points forts pour un sysadmin :**

- Réplication native (streaming), sauvegardes à chaud (PITR).
- Sécurité fine : rôles, `pg_hba.conf`, SCRAM-SHA-256, SSL/TLS, RLS.
- Supervision riche : vues `pg_stat_*`, `EXPLAIN ANALYZE`.
- Comportement prévisible sous charge : MVCC, autovacuum.

**Points de vigilance :**

- Le tuning par défaut de Debian est conservateur (adapté à une petite VM).
- Le WAL (journal) peut remplir un disque si l'archivage est mal configuré.
- Les mises à jour majeures (15 → 16 → 17) exigent `pg_upgrade` ou
  dump/restore — jamais un simple `apt upgrade`.

---

## 2. Versions 15, 16, 17 : ce qui change et cycle de vie

| Version | Sortie | Fin de support (EOL) | Nouveautés marquantes |
|---|---|---|---|
| 15 | oct. 2022 | nov. 2027 | `MERGE`, amélioration tri/parallélisme, WAL LZ4/Zstd |
| 16 | sept. 2023 | nov. 2028 | Parallélisme `FULL OUTER JOIN`, `pg_stat_io`, statistiques par défaut à 500 |
| 17 | sept. 2024 | nov. 2029 | Vacuum parallèle, `MERGE` avec `RETURNING`, amélioration `EXPLAIN` |

**Règle d'or :** choisissez la version la plus récente supportée par votre
distribution et vos applications, et planifiez la montée de version majeure
**avant** la fin de support (voir section 69, `pg_upgrade`).

> ⚠️ **Avertissement de version :** certaines syntaxes de ce guide (ex. `MERGE`,
> vues `pg_stat_io`) n'existent qu'à partir des versions indiquées. Vérifiez
> toujours avec `SELECT version();` sur votre serveur.

---

## 3. Architecture : processus et mémoire

Un cluster PostgreSQL, c'est **un postmaster** (processus parent) + des
processus fils, plus des zones de mémoire partagée.

### 3.1. Processus principaux

| Processus | Rôle |
|---|---|
| `postgres` (postmaster) | Écoute les connexions, lance les fils |
| `postgres: checkpointer` | Écrit les pages sales sur disque (checkpoints) |
| `postgres: background writer` | Écrit les pages sales en continu, en douceur |
| `postgres: walwriter` | Écrit le WAL sur disque |
| `postgres: autovacuum launcher` | Déclenche les workers autovacuum |
| `postgres: autovacuum worker` | VACUUM/ANALYZE automatiques |
| `postgres: logical replication launcher` | Réplication logique |
| `postgres: <user> <db>` | Un processus **par connexion client** (modèle fork) |

```bash
# Voir les processus d'un cluster en cours
ps aux | grep -E "postgres" | grep -v grep
```

**Conséquence pratique :** chaque connexion = un processus (~quelques Mo).
Au-delà de ~200–300 connexions, utilisez un **pooler** (PgBouncer)
plutôt que d'augmenter `max_connections` indéfiniment.

### 3.2. Mémoire

| Zone | Paramètre | Rôle |
|---|---|---|
| Mémoire partagée | `shared_buffers` | Cache de pages de données (le plus important) |
| Cache OS | `effective_cache_size` | Indication au planificateur (pas une allocation) |
| Mémoire par tri | `work_mem` | Tris, hachages **par opération** et par connexion |
| Mémoire maintenance | `maintenance_work_mem` | VACUUM, CREATE INDEX, ALTER TABLE |
| WAL buffers | `wal_buffers` | Tampon d'écriture du WAL |

```
┌─────────────────────────────────────────────┐
│              Mémoire partagée                │
│  shared_buffers │ wal_buffers │ locks, etc.  │
├─────────────────────────────────────────────┤
│  Processus par connexion (work_mem chacun)   │
├─────────────────────────────────────────────┤
│  Checkpointer / BG writer / WAL writer       │
└─────────────────────────────────────────────┘
```

---

## 4. MVCC : le contrôle de concurrence multi-version

PostgreSQL n'utilise **pas** de verrous en lecture : chaque transaction voit
un **instantané** (snapshot) des données au moment où elle démarre.

**Principes :**

- `INSERT`/`UPDATE`/`DELETE` créent de **nouvelles versions** de lignes
  (tuples) ; les anciennes restent jusqu'au `VACUUM`.
- Un `UPDATE` = `DELETE` logique + `INSERT` : la table grossit si on ne
  vacuum pas (bloat).
- Les lecteurs ne bloquent jamais les écrivains, et inversement.

```sql
-- Chaque ligne porte des numéros de transaction (xmin/xmax), invisibles
-- mais consultables pour le diagnostic :
SELECT xmin, xmax, * FROM capteurs LIMIT 3;
```

**Implications sysadmin :**

1. L'**autovacuum** n'est pas optionnel : sans lui, les tables gonflent
   et les performances s'effondrent (voir section 44).
2. Les **longues transactions** (idle in transaction) empêchent le
   nettoyage des vieilles versions → bloat + risque de wraparound du XID.
3. Le **wraparound** (dépassement du compteur de transactions sur 32 bits)
   est l'incident le plus grave : surveillez l'âge des XID (section 63).

---

## 5. Installation sur Debian/Ubuntu : dépôt officiel PostgreSQL

Les dépôts Debian/Ubuntu contiennent souvent une version en retard.
Pour les versions 15/16/17 récentes, utilisez le **dépôt officiel PGDG**.

```bash
# 1. Prérequis
sudo apt update
sudo apt install -y curl ca-certificates gnupg lsb-release

# 2. Clé du dépôt PGDG
sudo install -d /usr/share/postgresql-common/pgdg
sudo curl -o /usr/share/postgresql-common/pgdg/apt.postgresql.org.asc \
  --fail https://www.postgresql.org/media/keys/ACCC4CF8.asc

# 3. Ajouter le dépôt
sudo sh -c 'echo "deb [signed-by=/usr/share/postgresql-common/pgdg/apt.postgresql.org.asc] \
  https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" \
  > /etc/apt/sources.list.d/pgdg.list'

# 4. Installer PostgreSQL 17 (remplacez 17 par 15 ou 16 si besoin)
sudo apt update
sudo apt install -y postgresql-17 postgresql-client-17 postgresql-contrib-17
```

**Vérification :**

```bash
# Version installée et clusters présents
pg_lsclusters
# Exemple de sortie :
# Ver Cluster Port Status Owner    Data directory              Log file
# 17  main    5432 online postgres /var/lib/postgresql/17/main /var/log/postgresql/postgresql-17-main.log

# Le service systemd
sudo systemctl status postgresql@17-main
```

