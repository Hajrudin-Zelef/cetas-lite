---
id: etape7-phasef-databases/00-databases/9-managed-databases-pricing-and-positioning-2026-signals
title: "9. Managed databases — pricing and positioning (2026 signals)"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: pricing
actors: ["AWS", "Google"]
dates: ["2025-05"]
keywords: ["pricing", "acquisition", "aws", "cost", "open source"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [513, 578]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 85c507e2690b02634c11d8b73f4c57be281935c67ee5a6f0e08e11c67fd294d9
---

# 9. Managed databases — pricing and positioning (2026 signals)

## 9. Managed databases — pricing and positioning (2026 signals)

### 9.1 AWS RDS / Aurora (us-east-1 reference points)

- RDS PostgreSQL/MySQL **db.r6g.xlarge** (4 vCPU/32 GB): **$0.462/hr**
  Single-AZ ≈ $328–337/mo; Multi-AZ doubles to **$0.924/hr** [secondary].
- gp3 storage **$0.115/GB-mo**; automated backups 7-day retention free,
  manual snapshots **$0.095/GB-mo** [secondary].
- **Aurora PostgreSQL**: instance $0.462/hr (r6g.xlarge) + storage
  **$0.10/GB-mo** + **$0.20 per 1M I/Os**; I/O-Optimized tier kills I/O
  charges for ~30% higher instance price (break-even ≈ 25% I/O spend)
  [secondary].
- **Aurora Serverless v2**: **$0.12/ACU-hr**, 0.5–128 ACU in 0.5 steps,
  minimum capacity always billed; zero-scaling mode GA since 2024
  [secondary].
- RDS Extended Support for EOL engines is paid (year-1 pricing kicks in at
  community EOL + ~3 months) — factor into PG 14 planning [official].

### 9.2 Google Cloud SQL / Azure Database

- Cloud SQL for PostgreSQL/MySQL: per-vCPU-hour + storage + backup pricing;
  check current SKUs per region — no single 2026 figure was corroborated
  across sources (gap, §12).
- **Azure Database for PostgreSQL Flexible Server** and **Azure Database
  for MySQL Flexible Server** are the current Azure generations; single-
  server generations are retired/retiring [secondary].

### 9.3 Neon (serverless Postgres, Databricks-owned)

- **Acquired by Databricks May 2025 for ~$1B**; operates under the Lakebase
  umbrella with continued independent operation [secondary].
- **Pricing (verified Aug 2026)** [secondary]: 1 CU = 1 vCPU + 4 GB RAM.
  - Launch: **$0.106/CU-hr**, Scale: **$0.222/CU-hr**; storage **$0.35/GB-mo**
    (cut **80%** from $1.75 in late 2025, post-acquisition);
  - **No monthly minimum since Dec 2025** (use $3, pay $3);
  - Free tier: 100 CU-hr/mo, 0.5 GB storage/project.
- Capabilities: genuine **scale-to-zero** (cold start 300–500 ms), instant
  **copy-on-write branching** (10 Launch / 25 Scale), built-in PgBouncer
  (10K pooled conns), PITR 7-day (Launch) / 30-day (Scale) included,
  PG 17 support with pgvector/PostGIS/pg_stat_statements [secondary].
- Cost shape: wins on bursty/idle workloads; a steady 24/7 2-CU database
  runs ~$170/mo vs ~$70/mo for RDS db.t3.medium — dedicated instances win
  on always-on [secondary].
- Extension gaps vs RDS/self-host: no TimescaleDB; pg_cron on paid plans
  only [secondary].

### 9.4 PlanetScale, Turso and the serverless fringe

- **PlanetScale** (Vitess/MySQL-based serverless): 2026 status not
  reconfirmed in sources consulted — pricing floor cited at **$39/mo**
  minimum by a third party (treat as [unverified]) (gap, §12).
- **Turso** (libSQL/SQLite fork, edge-replicated) [secondary]:
  - Free: 500M rows read/mo, 10M rows write/mo, 5 GB storage, 500 DBs;
  - Developer **$4.99/mo** (2.5B reads, 1B writes, 9 GB);
  - Scaler **$24.92/mo** (24 GB, 100B rows); Pro **$416.58/mo** (SOC 2,
    HIPAA, SSO) — figures from community docs, verify on turso.tech.
  - Native vector search added to libSQL in early 2026 (HNSW, cosine/
    Euclidean/dot, 2–8 ms queries to 500K vectors on shared CPU)
    [secondary].
  - Self-hostable via `sqld`/libSQL server Docker image
    (`ghcr.io/tursodatabase/libsql-server`) [secondary].
- **LiteFS** (Fly.io): FUSE-based SQLite replication, open source; LiteFS
  Cloud backups **$5/mo** for up to 10 GB [secondary].

---

