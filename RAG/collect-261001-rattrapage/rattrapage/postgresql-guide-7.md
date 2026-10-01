---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-7
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-26", "2026-10-01", "2026-11-01"]
keywords: ["cost", "distribution", "pruning"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [1380, 1623]
sha256: 4ee0b8aa72388ee85dad00d992b0bf7e862053019b4b9c05ab8c2a3cd38009e9
---

# PostgreSQL en production — Guide complet pour sysadmin

-- Partitions mensuelles
CREATE TABLE mesures_2026_09 PARTITION OF mesures
  FOR VALUES FROM ('2026-09-01') TO ('2026-10-01');
CREATE TABLE mesures_2026_10 PARTITION OF mesures
  FOR VALUES FROM ('2026-10-01') TO ('2026-11-01');

-- Index locaux (créés sur la table mère -> propagés)
CREATE INDEX ON mesures (capteur_id, horodatage DESC);

-- Purge d'un mois : instantané, sans VACUUM
ALTER TABLE mesures DETACH PARTITION mesures_2026_06;
DROP TABLE mesures_2026_06;
```

**Automatisation :** extension `pg_partman` pour créer les partitions
futures et purger les anciennes. **Prérequis :** la contrainte de
partition doit apparaître dans les requêtes (partition pruning).

---

## 43. ANALYZE et les statistiques du planificateur

```sql
-- Statistiques manuelles après un gros chargement
ANALYZE mesures;
ANALYZE VERBOSE mesures;

-- Augmenter l'échantillonnage sur une colonne très sélective
ALTER TABLE mesures ALTER COLUMN capteur_id SET STATISTICS 1000;
ANALYZE mesures;

-- Voir les stats d'une colonne
SELECT attname, n_distinct, most_common_vals, correlation
FROM pg_stats WHERE tablename = 'mesures' AND attname = 'capteur_id';
```

- `default_statistics_target = 100` (500 depuis PG 16) : compromis
  temps d'ANALYZE / qualité des plans.
- Après un `COPY` massif ou un changement de distribution : `ANALYZE`.
- L'autovacuum lance `ANALYZE` automatiquement (seuils, section 44).

---

## 44. VACUUM et autovacuum : le poumon de PostgreSQL

Le MVCC laisse des **versions mortes** : `VACUUM` les recycle,
`ANALYZE` met à jour les stats, et les deux libèrent de l'espace.

```sql
-- Manuel (rarement nécessaire si l'autovacuum est sain)
VACUUM (VERBOSE, ANALYZE) mesures;

-- VACUUM FULL : réécrit la table, verrou exclusif -> maintenance uniquement
VACUUM (FULL, VERBOSE) _tmp_import;
```

**Paramètres autovacuum (postgresql.conf) :**

```ini
autovacuum = on                        # ne JAMAIS couper en prod
autovacuum_max_workers = 3
autovacuum_naptime = 60s
autovacuum_vacuum_threshold = 50
autovacuum_vacuum_scale_factor = 0.1    # VACUUM si > 10 % de lignes modifiées
autovacuum_analyze_scale_factor = 0.05
autovacuum_vacuum_cost_delay = 2ms      # throttle ; 0 = agressif
```

**Réglage par table** (tables à fort churn : mesures, logs) :

```sql
ALTER TABLE mesures SET (
  autovacuum_vacuum_scale_factor = 0.02,
  autovacuum_analyze_scale_factor = 0.01
);
```

**Surveillance :**

```sql
-- Tables qui ont le plus besoin d'un VACUUM
SELECT relname, n_dead_tup, last_vacuum, last_autovacuum
FROM pg_stat_user_tables
ORDER BY n_dead_tup DESC LIMIT 10;

-- Âge des transactions (wraparound) : ALERTE si > 1 milliard
SELECT c.oid::regclass AS table,
       age(c.relfrozenxid) AS age_xid
FROM pg_class c JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE n.nspname NOT IN ('pg_catalog','information_schema')
ORDER BY age_xid DESC LIMIT 10;
```

---

## 45. Transactions et niveaux d'isolation

```sql
BEGIN;
UPDATE onduleurs SET puissance_kva = 120 WHERE reference = 'UPS-042';
-- ... autres opérations ...
COMMIT;      -- ou ROLLBACK;
```

| Niveau | Lecture sale | Lecture non répétable | Fantômes |
|---|---|---|---|
| `READ COMMITTED` (défaut) | Non | Possible | Possible |
| `REPEATABLE READ` | Non | Non | Possible* |
| `SERIALIZABLE` | Non | Non | Non |

\* En pratique PG évite aussi les fantômes en `REPEATABLE READ`
via les snapshots.

```sql
-- Choisir le niveau
BEGIN ISOLATION LEVEL REPEATABLE READ;
-- ...
COMMIT;

-- Timeout anti-blocage (à mettre par rôle/appli !)
SET statement_timeout = '30s';
SET lock_timeout = '5s';
SET idle_in_transaction_session_timeout = '1min';
```

**Bonnes pratiques :**

- Transactions **courtes** : jamais d'appel réseau/fichier au milieu.
- `idle_in_transaction_session_timeout` en production : tue les
  transactions oubliées qui bloquent l'autovacuum.
- `SERIALIZABLE` : uniquement si la logique métier l'exige (erreurs
  de sérialisation à gérer par retry applicatif).

---

## 46. Verrous : diagnostiquer les blocages

```sql
-- Qui bloque qui, maintenant
SELECT blocked.pid AS pid_bloque,
       blocking.pid AS pid_bloquant,
       blocked.query AS requete_bloquee,
       blocking.query AS requete_bloquante,
       blocked.wait_event_type, blocked.wait_event
FROM pg_stat_activity blocked
JOIN pg_stat_activity blocking
  ON blocking.pid = ANY (pg_blocking_pids(blocked.pid))
WHERE blocked.wait_event_type = 'Lock';

-- Requêtes en attente de verrou depuis longtemps
SELECT pid, now() - xact_start AS duree, wait_event, query
FROM pg_stat_activity
WHERE wait_event_type = 'Lock'
ORDER BY xact_start;

-- Tueur de dernier recours (préférez comprendre d'abord !)
SELECT pg_terminate_backend(12345);  -- doux d'abord : pg_cancel_backend()
```

**Chaîne classique :** un `ALTER TABLE` attend un `ACCESS EXCLUSIVE`
derrière une longue transaction `idle in transaction` → tout s'empile.
**Prévention :** `lock_timeout`, `idle_in_transaction_session_timeout`,
migrations hors heures de pointe.

---

## 47. Checklist performance "revue mensuelle"

- [ ] Top 10 `pg_stat_statements` par temps total : EXPLAIN des pires
- [ ] Index jamais utilisés (`idx_scan = 0`) : supprimer après 1 mois
- [ ] `n_dead_tup` élevé persistant : régler l'autovacuum par table
- [ ] Âge XID max < 500 M (alerte à 1 Md)
- [ ] `log_min_duration_statement` : relire les lentes de la semaine
- [ ] Espace disque : tables + index qui grossissent anormalement
- [ ] `seq_scan` vs `idx_scan` par table (`pg_stat_user_tables`)

---

# Partie E — Exploitation : sauvegardes, PITR, réplication, haute dispo

## 48. Stratégie de sauvegarde : les 3 niveaux

| Niveau | Outil | Granularité | Restauration |
|---|---|---|---|
| Logique | `pg_dump` / `pg_dumpall` | Base ou cluster (SQL) | Flexible, lente sur gros volumes |
| Physique à chaud | `pg_basebackup` | Cluster entier (fichiers) | Rapide, base du PITR |
| Continue (PITR) | Archivage WAL + base | **À un instant T** | RPO ~ minutes |

**Règle 3-2-1 :** 3 copies, 2 supports différents, 1 hors site.
**Règle d'or :** une sauvegarde **non testée** = pas de sauvegarde.
Automatisez une restauration test mensuelle (section 78).

```bash
# Arborescence recommandée (serveur de sauvegarde dédié)
/srv/sauvegardes/postgresql/
├── logique/      # pg_dump quotidiens, rotation 14 j
├── physique/     # pg_basebackup hebdo, rotation 4 sem
└── wal/          # archives WAL continues (PITR)
```

---

## 49. pg_dump / pg_restore : la sauvegarde logique

```bash
# Dump d'une base au format custom (compressé, parallélisable à la restauration)
pg_dump -h srv-bdd -U sauvegarde -Fc -f /srv/sauvegardes/postgresql/logique/inventaire-$(date +%F).dump inventaire

# Dump parallèle (jobs = cœurs) : bien plus rapide
pg_dump -h srv-bdd -U sauvegarde -Fd -j 4 \
  -f /srv/sauvegardes/postgresql/logique/inventaire-$(date +%F)/ inventaire

# Dump des rôles et droits globaux (à part ! pg_dump ne les inclut pas)
pg_dumpall -h srv-bdd -U sauvegarde --globals-only \
  -f /srv/sauvegardes/postgresql/logique/globals-$(date +%F).sql

# Restauration
pg_restore -h srv-restore -U postgres -d inventaire --clean --if-exists \
  /srv/sauvegardes/postgresql/logique/inventaire-2026-09-26.dump

# Restauration parallèle
pg_restore -j 4 -h srv-restore -U postgres -d inventaire /chemin/dump-dir/
```

**Formats :**

| Format | Option | Avantages |
|---|---|---|
| custom | `-Fc` | Compressé, `pg_restore` sélectif (`-t table`, `-l/-L`) |
| directory | `-Fd` | Parallèle dump **et** restore |
| plain (SQL) | par défaut | Lisible, rejouable via `psql`, pas de parallélisme |
| tar | `-Ft` | Déconseillé (limité) |

**Bonnes pratiques :**

- Rôle `sauvegarde` dédié : `GRANT CONNECT`, `SELECT` sur les tables
  (ou appartenance à `pg_read_all_data`).
- Toujours dumper `--globals-only` à part (rôles, droits).
- Testez la restauration, pas seulement le dump (section 78).

---

## 50. pg_basebackup : la base physique à chaud

