---
id: etape7-phasef-databases/00-databases/13-conflicts-gaps-and-unverified-claims
title: "13. Conflicts, gaps and unverified claims"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: ["2026-05"]
keywords: ["aws", "benchmarks", "compute", "license", "licenses", "open source", "pricing", "throughput"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [687, 741]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: f6116f4c5e981564488707758d2a641e618f02a124c91ebdb6c45a194b7390df
---

# 13. Conflicts, gaps and unverified claims

## 13. Conflicts, gaps and unverified claims

- **C1 — PostgreSQL 18 release date**: community announcement says
  25 Sept 2025 [independent]; AWS RDS calendar lists 13 Nov 2025
  [official]; one community doc says "May 2026" [secondary — treated as
  error]. Minor-version numbering (18.6 by Aug 2026) is consistent with a
  Sept/Oct 2025 GA.
- **C2 — Turso pricing**: community docs cite Free 500M reads/10M writes
  vs another citing 1B reads/25M writes, and Developer $4.99 vs Scaler
  $24.92 vs Pro $416.58 tiers inconsistently — verify at turso.tech before
  budgeting [secondary].
- **C3 — Supabase pricing**: third-party pages agree on Free/Pro/Team/Ent
  structure but differ on edge-function invocation counts and MAU
  overage math — verify at supabase.com/pricing [secondary].
- **G1**: Google Cloud SQL and Azure Database 2026 unit pricing not
  corroborated — needs a pricing-page pass.
- **G2**: PlanetScale 2026 status/pricing not reconfirmed (last third-party
  signal: $39/mo floor [unverified]).
- **G3**: Percona's long-term Galera/XtraDB Cluster roadmap statement still
  awaited after PXC 9.7 (Sept 2026).
- **G4**: PostgreSQL 19 (Beta 3 seen Aug 2026) — GA date/features not
  confirmed by observation date; on-disk encryption GA status unknown.
- **G5**: MySQL 9.7 LTS GA confirmation — bug-tracker builds (9.7.1/9.7.2)
  seen, no official GA announcement captured.
- **G6**: Independent 2026 Valkey-vs-Redis and Dragonfly benchmarks are
  vendor-adjacent; treat throughput claims as directional.
- **G7**: Garnet (Microsoft) 2026 adoption momentum not verified.
- **U1**: "PostgreSQL 18 (May 2026)" in one community doc — marked erroneous
  (§1). **U2**: Dragonfly production user list (Meesho/Agoda/Dailymotion/
  ShareChat) is vendor-reported. **U3**: Supabase $500M/$10.5B round and
  "60% of DBs launched by AI tools" — single independent source, plausible
  but uncorroborated.

---

## 14. Glossary (ops terms)

- **ACU** — Aurora Capacity Unit (~2 GB RAM + CPU slice), Aurora Serverless
  billing unit. **CU** — Neon Compute Unit (1 vCPU + 4 GB RAM).
- **PITR** — point-in-time recovery via base backup + WAL replay.
- **RPO/RTO** — recovery point/time objectives driving sync-vs-async and
  backup-frequency choices.
- **GTID** — global transaction identifier (MySQL replication).
- **WAL** — write-ahead log (Postgres durability + replication + PITR
  foundation). **LSN** — log sequence number, WAL position.
- **HNSW** — hierarchical navigable small world, ANN vector index
  (pgvector, libSQL).
- **BSL 1.1** — Business Source License (Dragonfly): source-available, not
  OSI open source. **RSALv2/SSPLv1** — Redis source-available licenses.
- **RESP** — Redis serialization protocol (shared by Redis/Valkey/
  Dragonfly/KeyDB).
- **MAU** — monthly active users (Supabase auth billing unit).

---

