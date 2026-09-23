---
id: etape8-phasef-data-messaging/00-front-matter/3-sqlite-ecosystem-wal-litestream-turso-libsql
title: "3. SQLite ecosystem: WAL, Litestream, Turso/libSQL"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2025-05-21", "2025-05-29", "2025-09-16", "2026-02", "2026-03"]
keywords: ["memory", "pricing"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [101, 149]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: cc30cab5fc1bacde9fcad8f749663363d3ae56c7640dbf761a2078081fe6d050
---

# 3. SQLite ecosystem: WAL, Litestream, Turso/libSQL

## 3. SQLite ecosystem: WAL, Litestream, Turso/libSQL

### 3.1 SQLite core and WAL mode

- SQLite 3.50.0 was released on 2025-05-29, adding `sqlite3_setlk_timeout()`, `unistr()`/`unistr_quote()` functions, and `sqlite3_rsync` improvements — treat as secondary until verified against the official SQLite release page `[secondary]`.
- WAL (write-ahead logging) mode is the standard way to get concurrent readers with a single writer; rollback-journal mode is the legacy default `[independent]`.
- A 2026 secondary source claimed SQLite versions before 3.51.3 had a WAL-reset corruption issue fixed in March 2026 — requires official verification before treating as confirmed `[unverified]`.

### 3.2 Litestream: continuous SQLite replication

- Litestream streams SQLite WAL entries to S3-compatible object storage as LTX files, giving continuous backup and point-in-time restore for single-file SQLite databases `[secondary]`.
- Latest upstream release observed is **v0.5.8 (February 2026)**; the v0.3.x series had disk-space management bugs (shadow WAL directories growing unboundedly) fixed in v0.5.x `[secondary]`.
- Litestream 0.5 supports **one active replica** — do not dual-write two buckets from the same host; each host needs a separate replica prefix/lineage `[secondary]`.
- Restore: `litestream restore -config litestream.yml -timestamp "2026-04-10T14:30:00Z" ./restored.db` restores to a point in time; production runbooks recommend restoring locally first and inspecting with `sqlite3` or Prisma Studio before touching production `[secondary]`.
- Common deployment: Fly.io + Tigris S3 (env vars auto-set by Fly), Backblaze B2, Cloudflare R2 as weekly-archive tier; startup scripts fail closed when replica credentials are misconfigured `[secondary]`.
- LiteFS (same author) offers multi-node SQLite replication via a FUSE filesystem — with distributed-systems caveats (stale reads, possible data loss on primary crash) `[secondary]`.

### 3.3 Turso / libSQL

- Turso is the hosted platform around libSQL, the SQLite fork; the developer model is "SQLite at the edge" with replication to read replicas `[secondary]`.
- libSQL extends SQLite with extensions for sync and server modes; verify the current protocol/client matrix against official Turso docs, as this area moves fast `[unverified]`.
- Developer tradeoff vs Litestream: Turso is a managed platform (less operational work, vendor coupling); Litestream is a self-hosted replication sidecar over plain SQLite files (full file-format control, you own the restore path) `[independent]`.

### 3.4 When SQLite is the right database

- Single-writer, read-heavy workloads; embedded/edge deployments; local-first apps; test suites (in-memory or file-per-test) `[independent]`.
- Rule of thumb: if one process (or one primary) owns writes and you need WAL-speed reads plus Litestream-grade disaster recovery, SQLite competes with client-server databases up to surprisingly large sizes; beyond that, connection pooling and horizontal writes push you to PostgreSQL `[independent]`.

## 4. DuckDB and MotherDuck

### 4.1 DuckDB releases (2025)

- **DuckDB 1.3.0** (2025-05-21): introduced `DICT_FSST` dictionary compression, improved schema reconciliation, and numerous engine changes `[official]`.
- **DuckDB 1.4 LTS** (2025-09-16): long-term-support line adding AES-256-GCM database/WAL/temp-file encryption `[official]`.
- A purported "DuckDB 2.0" post dated 2026 was found on a fork, not the official project — do not treat it as an official release `[unverified]`.

### 4.2 Developer angle

- DuckDB is an in-process analytical (OLAP) database: no server, no client protocol; you link it into Python/Node/Java/Rust and query Parquet/CSV/JSON directly `[secondary]`.
- Typical patterns: local data exploration, ETL staging, pandas/polars replacement for larger-than-memory aggregations, embedded analytics inside applications, and unit-testing analytical SQL without a cluster `[independent]`.
- Extensions (httpfs, postgres, mysql, spatial, json) let DuckDB read remote files and federate queries against other databases from a single process `[secondary]`.
- MotherDuck is the serverless DuckDB platform: hybrid execution runs queries partly locally (your DuckDB) and partly in the cloud; verify current pricing and execution semantics against official MotherDuck sources `[unverified]`.

### 4.3 DuckDB vs alternatives (developer view)

- vs SQLite: DuckDB is columnar/analytical (aggregations over large scans); SQLite is row-oriented/transactional. Use DuckDB for analytics, SQLite for OLTP/embedded state `[independent]`.
- vs Spark: DuckDB handles single-machine datasets (up to low TBs on good hardware) with zero cluster overhead; Spark is for distributed datasets and shared clusters `[independent]`.
- vs ClickHouse: DuckDB is embedded and serverless-by-default; ClickHouse is a distributed columnar server for concurrent analytical serving `[independent]`.

