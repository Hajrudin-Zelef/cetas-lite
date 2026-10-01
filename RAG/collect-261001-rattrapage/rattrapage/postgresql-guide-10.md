---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-10
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [2076, 2283]
sha256: aba54866186e29cba747d1998485a20802d0e5728a1ccc93ac14c514e7eef0c1
---

# PostgreSQL en production — Guide complet pour sysadmin

| Vue | Contenu |
|---|---|
| `pg_stat_activity` | Sessions en cours (requêtes, attentes, états) |
| `pg_stat_database` | Stats par base (commits, rollbacks, conflits) |
| `pg_stat_user_tables` | Seq scan, index scan, `n_dead_tup`, dernier vacuum |
| `pg_stat_user_indexes` | Utilisation des index (`idx_scan`) |
| `pg_stat_statements` | Requêtes agrégées (extension) |
| `pg_stat_replication` | État des standbys (primaire) |
| `pg_stat_archiver` | Santé de l'archivage WAL |
| `pg_stat_bgwriter` | Checkpoints, écritures background |
| `pg_stat_io` (PG 16+) | I/O par backend/type d'objet |
| `pg_stat_ssl` | Sessions chiffrées ou non |
| `pg_stat_progress_vacuum` | Progression des VACUUM en cours |
| `pg_locks` | Verrous détenus/demandés |

```sql
-- Rôle lecture seule pour la supervision (principe du moindre privilège)
CREATE ROLE supervision WITH LOGIN PASSWORD '...';
GRANT pg_monitor TO supervision;   -- PG 10+ : accès aux vues de supervision
```

---

## 64. Requêtes de supervision prêtes à l'emploi

```sql
-- 1. Sessions actives et leur état
SELECT pid, usename, application_name, state,
       now() - query_start AS duree, left(query, 80)
FROM pg_stat_activity WHERE state <> 'idle' ORDER BY query_start;

-- 2. Requêtes les plus longues en ce moment
SELECT pid, now() - query_start AS duree, wait_event_type, wait_event,
       left(query, 120)
FROM pg_stat_activity
WHERE state = 'active' AND now() - query_start > INTERVAL '1 minute'
ORDER BY query_start;

-- 3. Connexions par base et par état
SELECT datname, state, count(*)
FROM pg_stat_activity GROUP BY datname, state ORDER BY datname;

-- 4. Taux de hit du cache (viser > 99 %)
SELECT datname,
       round(100.0 * blks_hit / nullif(blks_hit + blks_read, 0), 2) AS hit_ratio
FROM pg_stat_database;

-- 5. Tables les plus lues en séquentiel (candidats index)
SELECT relname, seq_scan, idx_scan,
       round(100.0 * seq_scan / nullif(seq_scan + idx_scan, 0), 1) AS pct_seq
FROM pg_stat_user_tables
WHERE seq_scan + idx_scan > 1000
ORDER BY seq_scan DESC LIMIT 10;

-- 6. Taille des bases, schémas, tables
SELECT datname, pg_size_pretty(pg_database_size(datname))
FROM pg_database ORDER BY pg_database_size(datname) DESC;
SELECT relname, pg_size_pretty(pg_total_relation_size(oid)) AS total,
       pg_size_pretty(pg_relation_size(oid)) AS table,
       pg_size_pretty(pg_indexes_size(oid)) AS index
FROM pg_class WHERE relkind = 'r' AND relnamespace = 'public'::regnamespace
ORDER BY pg_total_relation_size(oid) DESC LIMIT 10;

-- 7. Espace disque et WAL
SELECT pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), '0/0')) AS wal_genere_depuis_zero;
SHOW data_directory;
-- Shell : du -sh /var/lib/postgresql/17/main/pg_wal
```

---

## 65. Alertes : seuils recommandés

| Alerte | Seuil | Gravité |
|---|---|---|
| Espace disque PGDATA < 15 % libre | critique à 10 % | 🔴 |
| `pg_wal` grossit anormalement | > 2× `max_wal_size` | 🔴 |
| Archivage WAL en échec | `failed_count` > 0 | 🔴 |
| Âge XID max > 800 M | critique à 1 Md | 🔴 |
| Connexions > 80 % `max_connections` | | 🟠 |
| Lag réplication > 5 min (ou > seuil métier) | | 🟠 |
| Requête active > 15 min (hors batch déclaré) | | 🟠 |
| Hit ratio cache < 95 % durable | | 🟠 |
| Échec sauvegarde (dump/base/WAL) | | 🔴 |
| Certificat SSL expire < 30 j | | 🟠 |
| Autovacuum worker bloqué / `n_dead_tup` explose | | 🟠 |

```bash
# Exemple : check Nagios-like minimal (à brancher sur votre supervision)
#!/bin/bash
# check_pg_wal_archiver.sh
FAILED=$(sudo -u postgres psql -tAc "SELECT failed_count FROM pg_stat_archiver;")
[ "$FAILED" -gt 0 ] && { echo "CRITICAL: archivage WAL en echec"; exit 2; }
echo "OK: archivage WAL fonctionnel"; exit 0
```

**Intégration :** exposez ces requêtes via `check_postgres` (Nagios),
exporter Prometheus (`postgres_exporter`), ou Zabbix templates
PostgreSQL. Une alerte sans **runbook** (que faire ?) ne sert à rien :
liez chaque alerte à la section de dépannage correspondante.

---

## 66. Journaux : les exploiter vraiment

```bash
# Requêtes lentes du jour (log_min_duration_statement = 1000)
grep "duration:" /var/log/postgresql/postgresql-17-main.log | \
  sort -t: -k... # ou utilisez pgbadger :

# pgbadger : rapport HTML des logs (recommandé, hebdo)
sudo apt install -y pgbadger
pgbadger /var/log/postgresql/postgresql-17-main.log -o /tmp/rapport.html

# Centralisation : expédiez vers rsyslog / Loki / ELK
# postgresql.conf :
# log_destination = 'stderr'  (puis rsyslog imfile) ou 'csvlog' pour parsing
```

**À logger en production :** `log_checkpoints`, `log_connections`,
`log_disconnections`, `log_lock_waits`, `log_min_duration_statement`,
`log_autovacuum_min_duration = '1s'`, `log_temp_files = 0`
(fichiers temporaires = tris sur disque, à traquer).

---

## 67. Maintenance planifiée : le calendrier type

| Fréquence | Tâche |
|---|---|
| Quotidien | Dump logique + vérif. archivage WAL + espace disque |
| Hebdo | Base physique (`pg_basebackup` ou Barman) + rapport pgbadger |
| Mensuel | Revue `pg_stat_statements` + index inutilisés + test de restauration |
| Trimestriel | Revue tuning (charge réelle vs paramètres) + certificats SSL |
| Annuel | Game day PRA + montée de version mineure planifiée |

```bash
# Exemple de cron (serveur de sauvegarde)
# /etc/cron.d/postgresql-sauvegardes
0 2 * * * postgres /usr/local/bin/dump_logique.sh
0 3 * * 0 postgres /usr/local/bin/base_physique.sh
*/5 * * * * postgres /usr/local/bin/check_archivage.sh
```

---

## 68. Dimensionner le stockage : méthode

1. **Volume initial** : taille des données × 1.5 (index ≈ 30–50 % des tables).
2. **Croissance** : taux mensuel mesuré (`pg_database_size` historisé).
3. **WAL** : `max_wal_size` + marge archivage + `wal_keep_size` si slots.
4. **Marge** : 20 % libre minimum en permanence.
5. **Séparer** : données / WAL / sauvegardes / logs.

```sql
-- Historiser la taille pour la courbe de croissance (à mettre en cron)
CREATE TABLE IF NOT EXISTS supervision.tailles (
  mesure_le timestamptz NOT NULL DEFAULT now(),
  base text NOT NULL, taille bigint NOT NULL
);
INSERT INTO supervision.tailles (base, taille)
SELECT datname, pg_database_size(datname) FROM pg_database WHERE datistemplate = false;
```

---

# Partie G — Mise à jour, dépannage, erreurs classiques

## 69. Mise à jour mineure vs majeure

| Type | Exemple | Procédure | Risque |
|---|---|---|---|
| Mineure | 17.4 → 17.5 | `apt upgrade` + restart | Faible (même format de données) |
| Majeure | 17 → 18 | `pg_upgrade` ou dump/restore | Élevé si mal préparé |

```bash
# --- MINEURE ---
sudo apt update && sudo apt install -y postgresql-17
sudo pg_ctlcluster 17 main restart -m fast
sudo -u postgres psql -c "SELECT version();"

# --- MAJEURE avec pg_upgrade (méthode rapide, sans rechargement) ---
# 1. Installer la nouvelle version SANS supprimer l'ancienne
sudo apt install -y postgresql-18
# 2. Arrêter les deux clusters
sudo pg_ctlcluster 17 main stop -m fast
sudo pg_ctlcluster 18 main stop
sudo pg_dropcluster 18 main --stop   # supprime le cluster vide créé par défaut
# 3. Contrôle de compatibilité (à blanc)
sudo -u postgres /usr/lib/postgresql/18/bin/pg_upgrade \
  --old-datadir=/var/lib/postgresql/17/main \
  --new-datadir=/var/lib/postgresql/18/main \
  --old-bindir=/usr/lib/postgresql/17/bin \
  --new-bindir=/usr/lib/postgresql/18/bin \
  --old-options '-c config_file=/etc/postgresql/17/main/postgresql.conf' \
  --new-options '-c config_file=/etc/postgresql/18/main/postgresql.conf' \
  --check
# 4. Migration réelle (avec --link = instantané, ou copie = sûr mais long)
sudo -u postgres /usr/lib/postgresql/18/bin/pg_upgrade \
  ... --link
# 5. Démarrer, vérifier, puis basculer les applis
sudo pg_ctlcluster 18 main start
./analyze_new_cluster.sh   # script généré par pg_upgrade : ANALYZE
./delete_old_cluster.sh    # UNIQUEMENT après validation complète
```

**Checklist mise à jour majeure :**

