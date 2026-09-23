---
id: etape7-phasef-databases/00-databases/1-postgresql-releases-lifecycle-2026-status
title: "1. PostgreSQL — releases, lifecycle, 2026 status"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2024-09", "2024-11", "2025-09", "2025-11", "2026-05", "2026-08", "2026-11", "2027-02", "2030-02"]
keywords: ["aws", "benchmarks", "memory", "throughput"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [20, 109]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: bd870045d1c21351adccdc8c38d0b3184e2a9e97258b5f67288b3fffcf48d0c5
---

# 1. PostgreSQL — releases, lifecycle, 2026 status

## 1. PostgreSQL — releases, lifecycle, 2026 status

- PostgreSQL 18 is the current major. The community announcement is dated
  **25 September 2025** [independent]; AWS's RDS release calendar lists the
  community release date as **13 November 2025** [official] — a discrepancy of
  ~7 weeks between the announcement and AWS's recorded date, flagged as a
  conflict (§12). As of the August 2026 coordinated security release, the
  supported minors were **18.6, 17.11, 16.15, 15.19, 14.24**, plus **19 Beta 3**
  in testing [independent].
- PostgreSQL 17 was released **September 2024** (26 September 2024 per the
  release announcement) [official].
- One community document describes "PostgreSQL 18 (May 2026)" [secondary];
  this conflicts with the 2025 release record and is treated as an error in
  that document (§12).
- Community support policy: each major receives ~5 years of support. Key 2026
  milestones [official][secondary]:
  - **PG 14: EOL 12 November 2026** — final minor planned for that date; after
    it, no more security or bug fixes from upstream. Teams on 14 should target
    17 (EOL Nov 2029) or 18 (EOL Nov 2030) [secondary][independent].
  - PG 13 went EOL **November 2025**; PG 12 EOL was November 2024 [official].
  - RDS maps community EOL to its own calendar, e.g. PG 14 standard support on
    RDS ends **28 February 2027**, with paid Extended Support available to
    **February 2030** [official].
- The May 2026 minor-release batch (18.4, 17.10, 16.14, 15.18, 14.23) closed
  **11 security CVEs**, four of them CVSS v3.1 **8.8**: CVE-2026-6473
  (integer wraparound), CVE-2026-6475 (symlink following in `pg_basebackup` /
  `pg_rewind`), CVE-2026-6477 (`libpq` large-object path), CVE-2026-6637
  (stack buffer overflow in `refint` allowing an unprivileged DB user to execute
  code as the OS user) [independent]. All affected versions 14–18.
- The August 2026 batch fixed **28 CVEs** across 18.6/17.11/16.15/15.19/14.24
  and 19 Beta 3, including CVE-2026-14664 (regexp heap buffer overflow),
  CVE-2026-18408 (arbitrary code execution via psql `\unrestrict`), and
  CVE-2026-19385 (pg_dump heap buffer overflow) [independent].

### 1.1 PostgreSQL 17 highlights (operationally relevant)

- **Incremental backups**: `pg_basebackup` gained incremental backup support;
  new utility `pg_combinebackup` reconstructs a full backup from incrementals
  off the database server, so combining does not load the primary [official].
  For multi-terabyte clusters where a full backup can take days, this is the
  headline ops feature [secondary].
- `pg_dump --filter` lets operators pass a file specifying objects to
  include/exclude from a dump [official].
- New predefined role **`pg_maintain`** grants VACUUM/ANALYZE/CLUSTER/REINDEX/
  REFRESH MATERIALIZED VIEW/LOCK TABLE rights on all relations without
  superuser [official].
- Vacuum memory overhaul: up to **20x less memory** consumed by vacuum, faster
  vacuum, fewer shared-resource conflicts [vendor-reported].
- WAL processing optimizations: up to **2x write throughput** on high-
  concurrency workloads [vendor-reported] — vendor numbers, not independently
  reproduced here.
- New streaming I/O (read-stream) API accelerates sequential scans and
  `ANALYZE` [official].
- SQL/JSON standard support: `JSON_TABLE`, `JSON`/`JSON_SCALAR`/
  `JSON_SERIALIZE` constructors, `JSON_EXISTS`/`JSON_QUERY`/`JSON_VALUE`
  functions [official].
- `MERGE` gains `RETURNING` and the ability to update views; `COPY` gets up to
  2x export improvement on large rows and an `ON_ERROR` option to continue past
  bad rows [official].
- `sslnegotiation` connection parameter enables direct TLS handshake via ALPN
  (`postgresql` ALPN identifier), saving a network round trip [official].
- Client-side password hashing API `PQchangePassword` prevents accidental
  plaintext password logging [official].
- Built-in immutable collation provider (UTF-8 `C`-like semantics), immune to
  OS/libc collation changes [official].

### 1.2 PostgreSQL 18 highlights (operationally relevant)

- **Asynchronous I/O subsystem (AIO)**: parallel I/O requests instead of
  blocking per-block reads; controlled by the new `io_method` parameter
  (`io_uring` on Linux or classic synchronous). Community benchmarks cited
  show up to **3x** gains on read-heavy workloads (sequential scans, VACUUM);
  treat as vendor/community-reported, not independently verified
  [vendor-reported][secondary].
- **UUIDv7** support for time-ordered, index-friendly identifiers [secondary].
- On-disk database encryption initiative was in flight as of late 2025,
  aiming to offer enterprise-grade TDE without vendor lock-in
  [independent][unverified — no GA date confirmed by observation date].
- **Breaking changes for upgrades**: (1) PK/FK collation determinism rule —
  restore/`pg_upgrade` fails on mixed-determinism PK/FK collations, audit
  `pg_collation` + `pg_constraint` first; (2) `pg_upgrade
  --set-char-signedness` for cross-architecture (x86_64 signed char vs ARM
  unsigned char) carry-forward [secondary].
- 110 unique contributors, 202 new features — 25% more features and 5% more
  contributors than PG 17 [independent].
- The Register notes 18 "eyes analytics boost and distributed future", with
  some SQL features delayed to later releases [independent].

---

