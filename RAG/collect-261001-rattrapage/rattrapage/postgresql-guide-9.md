---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-9
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [1853, 2075]
sha256: ef6a2fe0118acc1d24804ed611fb6b89654b742cfd38f94fadfe43dad4c21871
---

# PostgreSQL en production — Guide complet pour sysadmin

```
        ┌─────────┐
        │  etcd ×3 │  <- quorum : qui est le leader ?
        └────┬────┘
     ┌───────┼────────┐
┌────┴───┐ ┌─┴────┐ ┌──┴─────┐
│ PG n1  │ │PG n2 │ │ PG n3  │  <- Patroni pilote PG sur chaque nœud
│leader  │ │sync  │ │async   │
└────────┘ └──────┘ └────────┘
     ▲ HAProxy / VIP : route les écritures vers le leader
```

**Concepts clés :**

| Concept | Rôle |
|---|---|
| Leader | Le seul primaire (écritures) |
| Synchronous standby | Accusé de réception synchrone (`synchronous_standby_names`) |
| DCS (etcd) | Élection, verrou du leader, anti-split-brain |
| `patronictl` | CLI : `list`, `failover`, `switchover`, `pause` |

```bash
# Exemples patronictl
patronictl -c /etc/patroni.yml list
patronictl -c /etc/patroni.yml switchover --master pg1 --candidate pg2
```

**Quand l'adopter :** à partir du moment où le RTO manuel (15–60 min)
ne suffit plus. **Coût :** 3 nœuds etcd + 2–3 nœuds PG + expertise.
Pour un premier palier, un standby + procédure de failover manuel
**testée** (section 54) bat un Patroni mal maîtrisé.

---

## 56. Réplication logique : le cas d'usage ciblé

Contrairement au streaming (tout le cluster), la réplication logique
réplique des **tables choisies**, entre versions parfois différentes.

```sql
-- Émetteur (primaire)
CREATE PUBLICATION pub_mesures FOR TABLE mesures, capteurs;

-- Récepteur (autre cluster, voire autre version majeure)
CREATE SUBSCRIPTION sub_mesures
  CONNECTION 'host=srv-primaire dbname=inventaire user=replicateur'
  PUBLICATION pub_mesures;
```

**Usages :** migration de version majeure sans coupure longue,
consolidation multi-sites, alimentation d'un entrepôt. **Limites :**
pas de DDL répliqué, séquences non synchronisées, pas de `TRUNCATE`
(par défaut), surcharge sur l'émetteur.

---

## 57. Sauvegardes : rôles, chiffrement, rotation

```bash
# Rôle minimal pour les sauvegardes logiques
CREATE ROLE sauvegarde WITH LOGIN PASSWORD '...';
GRANT CONNECT ON DATABASE inventaire TO sauvegarde;
GRANT USAGE ON SCHEMA public TO sauvegarde;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO sauvegarde;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO sauvegarde;
-- Pour pg_basebackup / streaming WAL :
ALTER ROLE sauvegarde WITH REPLICATION;

# Chiffrer les dumps avant envoi hors site (gpg, clé d'équipe)
gpg --encrypt --recipient sauvegardes@entreprise.lan \
  inventaire-2026-09-26.dump

# Rotation simple : garder 14 jours de dumps logiques
find /srv/sauvegardes/postgresql/logique/ -name '*.dump' -mtime +14 -delete
```

**Checklist sauvegarde :**

- [ ] Dump logique quotidien + `--globals-only`
- [ ] Base physique hebdo + WAL archivés en continu (PITR)
- [ ] Rétention écrite et appliquée (automatique)
- [ ] Chiffrement avant sortie du site
- [ ] **Restauration testée** mensuellement (section 78)
- [ ] Supervision : échec de dump/WAL → alerte immédiate

---

## 58. Plan de reprise : RTO/RPO et qui fait quoi

| Scénario | Restauration | RTO indicatif | RPO indicatif |
|---|---|---|---|
| Table supprimée par erreur | PITR à T-5 min (serveur isolé) | 30–60 min | ~0 |
| Disque données mort | Base physique + WAL sur nouveau disque | 1–2 h | minutes |
| Primaire mort | Failover standby (manuel) | 15–30 min | secondes–minutes |
| Cluster corrompu | Barman `recover` + PITR | 1–3 h | minutes |
| Incendie salle | Sauvegardes hors site | 4–24 h | 24 h (logique) |

**Documentez :** qui décide la bascule, l'ordre des étapes, les contacts,
les chaînes de connexion par environnement. **Testez** chaque scénario
au moins une fois par an (game day).

---

# Partie F — Tuning, monitoring, alertes

## 59. Tuning mémoire : les 4 paramètres qui comptent

| Paramètre | Point de départ | Contexte | Effet |
|---|---|---|---|
| `shared_buffers` | 25 % RAM (max ~8–16 Go utiles) | `postmaster` (restart) | Cache PG |
| `effective_cache_size` | 50–75 % RAM | `sighup` | Aide le planificateur (pas d'allocation) |
| `work_mem` | 4–64 Mo selon charge | `user` (par session !) | Tris/hachages par opération |
| `maintenance_work_mem` | 256 Mo–1 Go | `user` | VACUUM, CREATE INDEX |

```ini
# Exemple : serveur dédié 32 Go RAM, ~100 connexions
shared_buffers = 8GB
effective_cache_size = 24GB
work_mem = 32MB
maintenance_work_mem = 1GB
wal_buffers = 16MB          # auto-tuning depuis PG 13 si à -1
```

**Mises en garde :**

- `work_mem` est **par opération et par connexion** : 100 connexions ×
  10 tris × 64 Mo = 64 Go potentiels → OOM. Restez conservateur,
  augmentez **par requête** si besoin (`SET LOCAL work_mem = '256MB';`).
- `shared_buffers` au-delà de ~16 Go apporte rarement un gain
  (le cache OS fait le reste) et allonge les checkpoints.
- Ne touchez qu'**un paramètre à la fois**, mesurez avant/après
  (`pg_stat_statements`, temps de requêtes représentatives).

---

## 60. Tuning écritures : checkpoints et WAL

```ini
# Checkpoints espacés et étalés : moins de pics d'I/O
checkpoint_timeout = 15min
max_wal_size = 8GB            # 4 Go par défaut ; 8-16 Go sur serveur chargé
min_wal_size = 2GB
checkpoint_completion_target = 0.9

# WAL
wal_compression = on          # LZ4/Zstd (PG 15+) : moins d'I/O d'archivage
wal_buffers = 16MB
```

**Symptômes d'un mauvais réglage :**

- `LOG: checkpoints are occurring too frequently` → augmentez
  `max_wal_size`.
- Pics d'I/O réguliers corrélés aux checkpoints → augmentez
  `checkpoint_completion_target` (0.9) et vérifiez le stockage.

---

## 61. pg_stat_statements : l'extension indispensable

```bash
# 1. Activer (postgresql.conf) puis redémarrer
# shared_preload_libraries = 'pg_stat_statements'   # postmaster !
# pg_stat_statements.track = all
```

```sql
-- 2. Créer l'extension dans chaque base à superviser
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

-- 3. Top 10 des requêtes par temps total
SELECT query, calls, round(total_exec_time::numeric,2) AS total_ms,
       round(mean_exec_time::numeric,2) AS moyenne_ms,
       round((100*total_exec_time/sum(total_exec_time) OVER ())::numeric,1) AS pct
FROM pg_stat_statements
ORDER BY total_exec_time DESC LIMIT 10;

-- 4. Remise à zéro (après un tuning, pour mesurer le gain)
SELECT pg_stat_statements_reset();
```

**À surveiller :** `mean_exec_time` élevé + `calls` élevé (les pires),
`shared_blks_read` élevé (disque), `temp_blks_written` (tris sur disque →
`work_mem` ?). Pensez à la rétention : les stats sont en mémoire,
perdues au redémarrage (persistez via export régulier si besoin).

---

## 62. Autovacuum : réglage fin en production

Rappel section 44. Compléments production :

```ini
# Serveur avec tables à fort churn (mesures, logs)
autovacuum_max_workers = 4
autovacuum_naptime = 30s
# Ne pas mettre cost_delay à 0 globalement ; préférez le réglage par table :
```

```sql
-- Table très écrite : vacuum agressif, sans impacter les autres
ALTER TABLE mesures SET (
  autovacuum_vacuum_scale_factor = 0.01,
  autovacuum_vacuum_cost_delay = 0,
  autovacuum_analyze_scale_factor = 0.005
);

-- Table quasi statique : vacuum paresseux
ALTER TABLE sites SET (
  autovacuum_vacuum_scale_factor = 0.2,
  autovacuum_enabled = true
);

-- Vérifier l'activité autovacuum en cours
SELECT pid, query_start, query
FROM pg_stat_activity
WHERE backend_type = 'autovacuum worker';
```

---

## 63. Les vues pg_stat_* : panorama

