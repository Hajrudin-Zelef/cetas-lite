---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-2
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [188, 441]
sha256: d18d3780aa7f7ff1939578e5bf34f9f7fe506767b0ad7bea2e2e0a12d094829e
---

# PostgreSQL en production — Guide complet pour sysadmin

> 💡 Sur Debian, l'outil `pg_wrapper` (`pg_lsclusters`, `pg_ctlcluster`,
> `pg_dropcluster`…) gère plusieurs versions/clusters côte à côte.
> C'est spécifique à Debian/Ubuntu : sur RHEL, on utilise directement
> `initdb` et `systemctl`.

---

## 6. Arborescence Debian et premier démarrage

| Chemin | Contenu |
|---|---|
| `/etc/postgresql/17/main/` | `postgresql.conf`, `pg_hba.conf`, `pg_ident.conf` |
| `/var/lib/postgresql/17/main/` | Données (PGDATA) |
| `/var/log/postgresql/` | Journaux |
| `/var/run/postgresql/` | Socket Unix |

```bash
# Se connecter en local via le socket, en tant que postgres
sudo -u postgres psql

# Dans psql :
postgres=# SELECT version();
postgres=# \q
```

**Checklist post-installation immédiate :**

- [ ] `pg_lsclusters` → cluster `online`
- [ ] `sudo -u postgres psql -c "SELECT 1;"` → OK
- [ ] Définir un mot de passe pour le superuser `postgres`
      (ou mieux : créer des rôles dédiés, section 15)
- [ ] Adapter `listen_addresses` si accès réseau nécessaire (section 7)
- [ ] Configurer `pg_hba.conf` avant d'ouvrir le réseau (section 16)
- [ ] Activer et tester la sauvegarde (section 48) **dès le jour 1**

```sql
-- Changer le mot de passe du superuser postgres (à faire une fois)
ALTER USER postgres PASSWORD 'METTRE_UN_VRAI_MOT_DE_PASSE_DU_COFFRE';
```

---

## 7. postgresql.conf : les paramètres essentiels

Le fichier `/etc/postgresql/17/main/postgresql.conf` pilote le serveur.
Sur Debian, les valeurs Debian sont en tête ; vos réglages vont en fin
de fichier (ou dans `conf.d/`).

```bash
# Recharger la conf sans couper les connexions (SIGHUP)
sudo -u postgres psql -c "SELECT pg_reload_conf();"
# ou
sudo systemctl reload postgresql@17-main

# Savoir si un paramètre demande un redémarrage
sudo -u postgres psql -c \
  "SELECT name, context FROM pg_settings WHERE name IN
   ('shared_buffers','max_connections','port','listen_addresses');"
-- context = 'postmaster'  -> redémarrage obligatoire
-- context = 'sighup'      -> reload suffit
```

### 7.1. Connexions et réseau

```ini
# --- Connexions ---
listen_addresses = 'localhost'      # '*' pour écouter toutes les interfaces
port = 5432
max_connections = 100               # défaut ; 200-300 max sans pooler

# --- Sécurité réseau de base ---
ssl = on                            # TLS (voir section 18)
```

> ⚠️ Ne passez `listen_addresses = '*'` qu'**après** avoir durci
> `pg_hba.conf` (section 16) et activé SSL (section 18).

### 7.2. Emplacements et journaux

```ini
# --- Journalisation (indispensable en prod) ---
logging_collector = on
log_directory = 'log'               # relatif à PGDATA, ou absolu
log_filename = 'postgresql-%Y-%m-%d.log'
log_rotation_age = 1d
log_rotation_size = 100MB
log_min_duration_statement = 1000   # loggue les requêtes > 1 s (ms)
log_checkpoints = on
log_connections = on
log_disconnections = on
log_lock_waits = on                 # loggue les attentes de verrous
log_autovacuum_min_duration = 0     # 0 = tout logguer (bruyant) ; 1000 en prod
```

### 7.3. Checkpoints et WAL (première approche)

```ini
# --- WAL / checkpoints ---
wal_level = replica                 # minimal pour réplication/PITR (défaut 15+)
max_wal_size = 4GB
min_wal_size = 1GB
checkpoint_timeout = 15min
checkpoint_completion_target = 0.9
```

Le détail du tuning (mémoire, checkpoints, autovacuum) est en
partie F (sections 59–62).

---

## 8. Créer et gérer les clusters Debian (pg_wrapper)

```bash
# Lister
pg_lsclusters

# Créer un second cluster (ex. version 17, port 5433, locale fr)
sudo pg_createcluster 17 test --port=5433 --locale=fr_FR.UTF-8

# Démarrer / arrêter / redémarrer / recharger
sudo pg_ctlcluster 17 main start
sudo pg_ctlcluster 17 main stop
sudo pg_ctlcluster 17 main restart
sudo pg_ctlcluster 17 main reload

# Supprimer un cluster (destructif ! demande confirmation)
sudo pg_dropcluster 17 test

# Voir la conf effective d'un paramètre
sudo -u postgres psql -c "SHOW shared_buffers;"
sudo -u postgres psql -c "SHOW config_file;"
```

**Bonnes pratiques :**

- Un cluster = une version majeure + un jeu de paramètres + un port.
- En production : **un seul cluster** par serveur (sauf besoin explicite).
- Avant `pg_dropcluster` : sauvegarde + vérification (section 48).

---

## 9. Arrêt et démarrage : les modes

| Mode | Commande | Effet |
|---|---|---|
| smart | `pg_ctlcluster … stop -m smart` | Attend la fin des sessions (défaut Debian) |
| fast | `… stop -m fast` | Coupe les connexions, rollback, arrêt propre |
| immediate | `… stop -m immediate` | Arrêt brutal → recovery au redémarrage |

```bash
# Arrêt propre rapide (recommandé pour maintenance planifiée)
sudo pg_ctlcluster 17 main stop -m fast

# Redémarrage après changement de paramètre postmaster
sudo pg_ctlcluster 17 main restart -m fast
```

> ⚠️ Le mode `immediate` force une relecture du WAL au redémarrage
> (crash recovery). C'est sans perte de données validées, mais le
> redémarrage est plus long sur une grosse base.

---

## 10. Prérequis système : disque, filesystem, locale

**Disque :**

- Séparez si possible : données (`/var/lib/postgresql`), WAL (`pg_wal`
  sur disque dédié via tablespace ou lien), sauvegardes (autre machine).
- Filesystem : **ext4** ou **XFS** (évitez les FS exotiques ou le NFS
  pour les données — le NFS est une source classique de corruption).
- Réservez 15–20 % d'espace libre pour le WAL, les tris et l'autovacuum.

**Locale et encodage :**

```bash
# Vérifier les locales disponibles AVANT de créer le cluster
locale -a | grep -i utf

# Créer le cluster en UTF-8 explicite (recommandé)
sudo pg_createcluster 17 main --locale=fr_FR.UTF-8 \
  --start-conf=auto
```

> ⚠️ L'encodage et la locale d'un cluster sont fixés à `initdb`.
> On ne les change pas après coup : il faut dump/restore.

**Checklist prérequis :**

- [ ] Disque données + disque WAL/sauvegardes séparés (ou planifiés)
- [ ] Locale `*.UTF-8` disponible
- [ ] NTP/chrony actif (horodatage WAL, certificats SSL)
- [ ] Firewall : n'ouvrir 5432 que vers les hôtes applicatifs

---

## 11. Installation : cas d'une version précise et dépôts figés (prod)

En production, figez la version mineure pour des déploiements reproductibles :

```bash
# Voir les versions disponibles dans le dépôt PGDG
apt-cache policy postgresql-17

# Installer une version mineure précise
sudo apt install -y postgresql-17=17.4-1.pgdg22.04+1

# Figer le paquet (hold) pour éviter une montée surprise
sudo apt-mark hold postgresql-17 postgresql-client-17

# Lever le hold quand vous planifiez la mise à jour mineure
sudo apt-mark unhold postgresql-17
```

**Mises à jour mineures (ex. 17.4 → 17.5) :** simple `apt upgrade` +
redémarrage du cluster. Elles ne changent pas le format des données.

**Mises à jour majeures (17 → 18) :** voir section 69 (`pg_upgrade`).

---

## 12. Vérifier une installation : la batterie de tests "jour 1"

```bash
#!/bin/bash
# health_jour1.sh — à lancer après chaque installation
set -e
echo "== Cluster ==";      pg_lsclusters
echo "== Version ==";      sudo -u postgres psql -tAc "SELECT version();"
echo "== Ecoute ==";       ss -ltnp | grep 5432 || echo "socket local uniquement"
echo "== Ecriture ==";     sudo -u postgres psql -tAc \
  "CREATE TABLE IF NOT EXISTS _t(id serial primary key); \
   INSERT INTO _t DEFAULT VALUES RETURNING id; DROP TABLE _t;"
echo "== Conf ==";         sudo -u postgres psql -tAc \
  "SHOW config_file; SHOW hba_file; SHOW data_directory;"
echo "OK"
```

Si chaque étape passe, la base est prête pour la configuration de
production (partie B).

---

# Partie B — Connexion sécurisée : psql, rôles, authentification, SSL

## 13. psql : le client en ligne de commande

`psql` est l'outil quotidien du sysadmin : requêtes, administration,
scripts.

```bash
# Connexion locale via socket (pair = OS user postgres)
sudo -u postgres psql

