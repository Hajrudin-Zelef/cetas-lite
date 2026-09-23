---
id: etape7-phasef-databases/00-databases/3-postgresql-connection-pooling-extensions-tuning
title: "3. PostgreSQL — connection pooling, extensions, tuning"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["advisory", "benchmark", "memory"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [166, 247]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 9fe3d4b94f5f1bae9313d1d35949ca733983d5751f385555eb62b43d76893f77
---

# 3. PostgreSQL — connection pooling, extensions, tuning

## 3. PostgreSQL — connection pooling, extensions, tuning

### 3.1 PgBouncer

- Standard lightweight pooler; typical deployment uses **transaction pooling
  mode** (`pool_mode=transaction`): a server connection is held only for the
  transaction duration. Common production config seen in 2026 playbooks:
  `max_client_conn=1000`, `default_pool_size=20–25`,
  `auth_type=scram-sha-256`, port **6432** [secondary].
- Example sizing: 3 app replicas × 15 client conns saturate only 20 backends
  in transaction mode vs 40 in session mode (measured on one documented stack;
  single data point, not a general benchmark) [secondary].
- Gotchas: no session state between transactions — `LISTEN`/`NOTIFY`,
  prepared statements (disable with `prepare_threshold=None` in SQLAlchemy),
  advisory locks, and `SET` variables must use direct connections or statement
  mode carefully [secondary].
- Alternatives: **pgpool-II** (pooling + load balancing + watchdog, heavier),
  **Odyssey** (Scaleway, multi-threaded), built-in pooling in Neon/Supabase
  (both expose PgBouncer endpoints) [secondary].

### 3.2 Extensions that matter operationally

- **pgvector**: HNSW/IVFFlat vector indexes; production tuning rules of thumb
  from 2026 guides: `m=16`, `ef_construction=64` defaults,
  `hnsw.ef_search=100–200` for high recall; hybrid keyword+vector filtering
  pushes selective `WHERE` before the ANN scan; pgvector-vs-dedicated-DB
  decision guidance favors pgvector under ~10M vectors when already on
  Postgres [secondary]. Foundational pgvector work tracked toward PG 19
  [secondary][unverified].
- **PostGIS**: the standard geospatial extension; heavy GIS workloads often
  size shared_buffers/work_mem around it [secondary].
- **TimescaleDB**: time-series hypertables; note Neon documents TimescaleDB
  as *not* in its supported extension set (2026) [secondary].
- **pg_stat_statements**: the canonical query-stats extension — see §11.
- **pg_cron**: in-database job scheduler; available on Neon's paid plans,
  absent from some managed offerings — check per provider [secondary].
- **Citus** (Microsoft): distributed Postgres; Microsoft's Fabric "Database
  Hub" (2026) promises unified wrangling of PostgreSQL/MySQL/SQL Server
  [independent]. Microsoft also launched **Azure HorizonDB** (2026) to compete
  with distributed Postgres rivals [independent].
- **FerretDB**: MongoDB-wire-protocol front end on Postgres; Microsoft's 2026
  document-database work references FerretDB as a front end [independent].

### 3.3 Tuning basics (ops checklist)

- `max_connections` default **100**; each backend costs ~5–10 MB RAM —
  the reason poolers exist [secondary].
- `shared_buffers` 25% of RAM (Linux), `effective_cache_size` 50–75% of RAM,
  `work_mem`/`maintenance_work_mem` sized per workload; `checkpoint_timeout`
  / `max_wal_size` tuned for write bursts [secondary].
- `autovacuum` tuning matters more than most knobs on write-heavy tables;
  PG 17's 20x vacuum memory reduction eases this [secondary].
- Connection storms from serverless/edge functions: always put PgBouncer (or
  the managed pooler) in front; one 2026 ADR documents 12 Go services
  collapsing 120–240 direct connections into a 25-connection pool
  [secondary].

---

## 4. PostgreSQL — backup, WAL, point-in-time recovery

- **WAL archiving** (`archive_mode=on`, `archive_command`) is the foundation of
  PITR: base backup + continuous WAL segments allow restore to any point in
  time [secondary].
- Tooling landscape 2026 [secondary]:
  - **pgBackRest**: parallel, compressed, encrypted backups; S3/Azure/GCS
    backends; delta/incremental; the reference for serious self-hosted
    setups.
  - **Barman** (EDB): catalog-based, rsync/parallel, cloud object storage.
  - **WAL-G / WAL-E**: continuous archiving to object storage, popular in
    container/K8s stacks.
  - **pg_basebackup** (core): simple physical backups; PG 17 added
    incremental mode + `pg_combinebackup` [official].
- Managed defaults for comparison: Supabase Pro includes daily backups with
  7-day retention (14-day on Team), PITR as an Enterprise add-on; Neon Launch
  includes 7-day PITR, Scale 30-day [secondary].
- Ransomware-era guidance: immutable/off-site backup copies, tested restores
  (unrestored backups are not backups), `pg_verifybackup` for manifest
  validation [secondary].

---

