---
id: etape7-phasef-databases/00-databases/5-mysql-releases-lifecycle-2026-status
title: "5. MySQL — releases, lifecycle, 2026 status"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Microsoft", "Oracle"]
dates: ["2025-09", "2026-01-20", "2026-04", "2026-05", "2026-07", "2026-08", "2026-09", "2029-06"]
keywords: ["mcp", "memory", "packaging"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [248, 316]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 02052c4e2c411c4a3385a91f9b7aadb494ac3b108ca8d0df567aa71951540ce6
---

# 5. MySQL — releases, lifecycle, 2026 status

## 5. MySQL — releases, lifecycle, 2026 status

### 5.1 The two-track model (LTS vs Innovation)

- Since 2023–2024 Oracle ships MySQL on two tracks [secondary]:
  - **LTS track**: **8.4 LTS** (current production standard; support horizon
    commonly cited to ~2032) — the safe choice for enterprise [secondary].
  - **Innovation track**: quarterly releases (9.0, 9.1, …), each supported
    only ~3 months until the next; deprecations/removals land here first;
    intended for dev/test and fast followers [secondary].
- **MySQL 8.0 reached EOL in April 2026**. Percona reports **more than half
  of installs remain on 8.0** as EOL loomed — a large exposed fleet
  [independent]. Migration targets: 8.4 LTS or the 9.x line.
- **MySQL 9.6.0** (Innovation) was current in early 2026 with: a redesigned
  modular **Audit Log** subsystem, GTID replication performance work (new
  internal data structures for GTID handling), InnoDB unique-`rowid`
  generation improvements, a `--container_aware` flag (auto-detects cgroup
  CPU/memory limits in Docker/Kubernetes), and MD5()/SHA1() moved to a
  pluggable `classic_hashing` component so insecure hashes can be fully
  disabled [secondary].
- **MySQL 9.7** is positioned as the next LTS ("parked stable destination",
  est. 2026); bug-tracker records from August 2026 already reference
  **9.7.1/9.7.2** builds [secondary][unverified — GA status of 9.7 LTS not
  confirmed by an official announcement in the sources consulted].
- **MySQL AI** (add-on to MySQL 9.6.0, release notes 2026-01-20): AutoML with
  NL2ML natural-language interface, `ML_GENERATE` routine, MCP-server
  integration hooks, bulk ingest into non-empty tables [official].

### 5.2 MariaDB — releases, lifecycle, 2026 status

- **MariaDB 13.0** went stable **18 September 2026** (13.0.2 latest at
  observation): UPDATE RETURNING, improved Optimizer Trace, InnoDB log
  archiving, SQL compatibility enhancements [secondary].
- **MariaDB 12.3 LTS** GA'd **28–30 May 2026** (12.3.3 current), maintenance
  planned through **June 2029**; introduced a new binary-log implementation,
  `innodb_snapshot_isolation` defaulting to ON (potential breaking change for
  workloads relying on legacy REPEATABLE READ behavior), claimed 4x write
  performance improvements [vendor-reported for the perf claim].
- **MariaDB 11.8 LTS**: native **MariaDB Vector** (VECTOR type, ANN indexes,
  `VEC_DISTANCE()`), TIMESTAMP range extended to year 2106 (2038 fix), default
  charset `latin1` → `utf8mb4`, PARSEC elliptic-curve auth plugin
  [secondary]. Maintenance release **11.8.9** shipped August 2026.
- **MariaDB 10.6 reached EOL**: community maintenance ended **6 July 2026**,
  final farewell binary **10.6.28** released **13 August 2026**; "there will be
  no 10.6.29" [secondary]. Users advised to jump directly to 11.8 or 12.3.
- Ubuntu/Debian packaging (2026): Ubuntu 26.04 "Resolute" carries MariaDB
  **11.8.9** in main; 24.04 "Noble" carries 10.11.x in universe
  [official-ish via Launchpad].
- **Azure Database for MariaDB was sunset** — Microsoft set retirement for
  **19 September 2025**, directing users to Azure Database for MySQL
  [independent].

### 5.3 InnoDB — storage engine notes

- Default engine for MySQL and MariaDB; row-level locking, MVCC, crash-safe
  via redo log [secondary].
- 2026 InnoDB work items (9.6): unique `rowid` generation efficiency for
  tables without explicit PKs; recovery of transactions stuck in PREPARED
  state after crash [secondary].
- Active 2026 bug-tracker items (S1/S2) show ongoing InnoDB edge cases:
  FK-cascade virtual-column use-after-free, BLOB update crash on recovery
  rollback, multi-valued JSON index builds crashing mysqld — worth a pre-
  upgrade scan of the tracker for your exact minor [official].
- Operator rule of thumb: every table needs an explicit primary key (also
  required for sane replication and for InnoDB's clustered index)
  [secondary].

---

