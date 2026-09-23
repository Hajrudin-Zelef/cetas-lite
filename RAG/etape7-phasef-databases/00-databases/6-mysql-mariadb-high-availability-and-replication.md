---
id: etape7-phasef-databases/00-databases/6-mysql-mariadb-high-availability-and-replication
title: "6. MySQL/MariaDB — high availability and replication"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2025-05", "2026-09"]
keywords: ["latency"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [317, 370]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 9ab7e3f9f5891d772890eee4cea993e49c9176a90545ce72e392781df04ddb81
---

# 6. MySQL/MariaDB — high availability and replication

## 6. MySQL/MariaDB — high availability and replication

### 6.1 Async replication + GTID

- Classic primary/replica with **GTID** (global transaction identifiers) is
  the standard topology; MySQL 9.6 reworked GTID internals for replication
  performance [secondary].
- Semi-sync replication (`rpl_semi_sync_master/slave`) gives a durability/
  latency middle ground, analogous to Postgres quorum commit [secondary].

### 6.2 MySQL Group Replication / InnoDB Cluster

- Group Replication: multi-primary or single-primary Paxos-based replication,
  foundation of **MySQL InnoDB Cluster** (with MySQL Router and AdminAPI)
  [secondary].
- Conflict detection limits multi-primary writes to workloads that avoid
  concurrent writes to the same rows [secondary].

### 6.3 Galera — turbulent 2025–2026

- **MariaDB Corporation acquired Codership Oy (Galera) in May 2025**;
  announced end of support for Galera Cluster for MySQL 8.4 (end Sept 2026)
  and signaled dropping Galera from the Community Edition — community
  outcry followed, then a partial retreat ("now is not the time for a major
  change") [independent].
- The MariaDB Foundation stated it does **not** intend to fork Galera
  [independent].
- **Percona XtraDB Cluster (PXC) 9.7** was released **10 September 2026** —
  Percona's improved MySQL+Galera fork, seen as the continuity signal for
  Galera-on-MySQL users [independent]. Percona's longer-term Galera
  statement is still awaited (gap, §12).
- Galera (synchronous multi-primary, certification-based replication)
  remains popular for MariaDB clusters needing write-anywhere with
  conflict handling at commit [secondary].

### 6.4 Percona variants

- **Percona Server for MySQL** and **Percona XtraDB Cluster**: drop-in
  enhanced builds with extra instrumentation, thread-pool, backup (XtraBackup)
  integration [vendor-reported].
- **Percona XtraBackup**: the standard hot-backup tool for MySQL/MariaDB
  (physical, non-blocking, incremental support) [secondary].

### 6.5 Proxy / routing layer

- **ProxySQL**: query routing, read/write split, connection multiplexing,
  query cache/rules — the MySQL-world answer to PgBouncer+routing
  [secondary].
- **MySQL Router** (Oracle): lightweight routing for InnoDB Cluster
  [official].
- **HAProxy** commonly fronts MySQL for TCP-level failover [secondary].

---

