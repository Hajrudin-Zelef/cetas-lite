---
id: etape7-phasef-databases/00-databases/2-postgresql-high-availability-replication-failover
title: "2. PostgreSQL — high availability, replication, failover"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [110, 165]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: db0a4ff671e827962b31162dd0e54b4f74d7536b0c2f43c6647cdf9d20859e9a
---

# 2. PostgreSQL — high availability, replication, failover

## 2. PostgreSQL — high availability, replication, failover

### 2.1 Streaming replication (physical)

- Built-in since 9.0: WAL shipping from primary to standbys; **synchronous**
  replication (`synchronous_commit`, `synchronous_standby_names`) gives
  zero-data-loss failover at the cost of write latency [secondary].
- `pg_basebackup -R` bootstraps a standby with replication configuration in
  one command; widely used in container/Docker HA setups [secondary].
- Synchronous replication topology guidance from production playbooks:
  `synchronous_standby_names = 'FIRST 1 (standby1, standby2)'` so one
  confirmed standby is enough for commit [secondary].

### 2.2 Logical replication

- Publisher/subscriber model (10+); replicates at row level, supports
  cross-version replication and selective tables — the standard path for
  near-zero-downtime major upgrades (logical-replicate to the new major, then
  cut over) [secondary].
- PG 17 improved logical replication (notably around failover/slot handling)
  per release coverage [secondary].

### 2.3 Patroni — the de-facto self-hosted HA stack

- Patroni is the dominant open-source HA orchestrator for self-managed
  PostgreSQL: DCS-backed leader election (etcd/ZooKeeper/Consul), automatic
  failover, `pg_rewind` support, REST API [secondary].
- Production playbooks pair Patroni with PgBouncer and etcd clusters;
  widely deployed on Kubernetes via operators (e.g. Zalando's
  postgres-operator, CloudNativePG) [secondary].
- **Stolon** (Sorint.lab) is the older alternative; community momentum has
  clearly favored Patroni by 2026 [secondary].

### 2.4 Kubernetes operators

- **CloudNativePG** (CNCF, EDB-originated): declarative `Cluster` CRD,
  rolling updates, backup via Barman/VolumeSnapshots, popular for GitOps
  Postgres on Kubernetes [secondary].
- **Zalando postgres-operator**: Patroni-based, mature, used at scale by
  Zalando [secondary].
- **Percona Operator for PostgreSQL**, **Crunchy PGO** (Crunchy Data):
  enterprise-flavored alternatives with monitoring and backup integration
  [secondary].

### 2.5 Synchronous vs asynchronous — ops trade-offs

- Sync replication: RPO=0, but every commit waits on standby ACK — tail
  latency coupled to the slowest synchronous standby and network RTT;
  recommended within a region/AZ pair, risky cross-region [secondary].
- Async: sub-millisecond primary commits; RPO of seconds of WAL on failover;
  the default for most web workloads [secondary].
- Quorum commit (`remote_apply`, `FIRST n`) tunes the durability/latency
  trade-off per transaction [secondary].

---

