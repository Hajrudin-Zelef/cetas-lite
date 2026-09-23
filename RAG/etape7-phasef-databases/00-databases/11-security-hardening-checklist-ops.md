---
id: etape7-phasef-databases/00-databases/11-security-hardening-checklist-ops
title: "11. Security hardening checklist (ops)"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Oracle"]
dates: []
keywords: ["throughput"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [626, 686]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 945eda4ce4a8028545864a143c3b3b9b4b7d78d988bfca975eb8abc12885ffe1
---

# 11. Security hardening checklist (ops)

## 11. Security hardening checklist (ops)

- **Authentication**: SCRAM-SHA-256 for Postgres (default since PG 10;
  audit for legacy MD5 hashes inherited from ≤13 — CVE-2026-6478 notes a
  covert timing channel in MD5 comparison) [independent]; `caching_sha2_
  password` for MySQL 8+; PARSEC available in MariaDB 11.8+ (not default)
  [secondary].
- **Network**: bind to private interfaces/VPC only; TLS in transit
  (`sslmode=require` minimum; `verify-full` in production); PgBouncer ↔
  Postgres TLS (`server_tls_sslmode=require`) seen in 2026 production
  configs [secondary].
- **Secrets**: never in config files — env files, Vault/SSM, Docker/K8s
  secrets; PgBouncer `auth_query` against `pg_shadow` avoids storing
  plaintext passwords [secondary].
- **Least privilege**: app roles ≠ owners; `pg_maintain` (PG 17+) for
  maintenance without superuser; MySQL 9.6's modular audit log with
  `AUDIT_ADMIN` for log rotation [official].
- **Patching cadence**: quarterly minor releases for Postgres (Feb/May/Aug/
  Nov); MySQL innovation quarterly; track EOL dates (PG 14 Nov 2026, MySQL
  8.0 Apr 2026 already past) [official][independent].
- **Backups encrypted + off-site + restore-tested**; Supabase free tier has
  *no* backups — a production blocker [secondary].

---

## 12. Decision guides (ops)

### 12.1 Which Postgres HA for self-hosting?

- **Single VM + replicas**: Patroni + etcd — maximum control, you own
  failover/PITR/monitoring [secondary].
- **Kubernetes**: CloudNativePG (declarative, GitOps-friendly) or Zalando
  postgres-operator (Patroni core, battle-tested) [secondary].
- **Don't want to operate**: RDS/Aurora (ecosystem, pricey at scale),
  Neon (serverless shape), Supabase (full backend, not just DB),
  Aiven/EDB/Crunchy (managed Postgres specialists) [secondary].

### 12.2 Redis vs Valkey vs Dragonfly vs managed (2026)

- **Default for new self-hosted cache**: Valkey (BSD-3, Linux Foundation,
  drop-in protocol, cheaper on ElastiCache) [secondary].
- **Stay on Redis** if you need Redis 8's bundled modules (JSON/search/
  vectors in core) and the AGPL/SSPL terms are acceptable to your legal
  team [secondary].
- **Dragonfly** if you want vertical scaling and peak single-node
  throughput and accept BSL 1.1 [vendor-reported].
- **Upstash/serverless** for spiky/edge workloads billed per command
  [vendor-reported].

### 12.3 MySQL-family choice (2026)

- **New deployments**: MySQL **8.4 LTS** (Oracle) or MariaDB **11.8/12.3
  LTS**; avoid 8.0 (EOL) and 10.6 (EOL) [independent][secondary].
- **Galera-style multi-primary**: MariaDB + Galera (watch MariaDB Corp
  support signals) or **Percona XtraDB Cluster 9.7** (Sept 2026 continuity
  signal) [independent].
- **Vector/AI in MySQL-land**: MariaDB 11.8+ native VECTOR; MySQL 9.x
  innovation track has vector search features [secondary].

---

