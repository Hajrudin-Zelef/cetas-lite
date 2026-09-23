---
id: etape8-phasef-data-messaging/00-front-matter/13-version-release-timeline-2025-2026-09-22
title: "13. Version-release timeline (2025 → 2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2024-10-21", "2025-05-21", "2025-05-29", "2025-09", "2025-09-16", "2026-03", "2026-03-24", "2026-04-14", "2026-04-27", "2026-04-30", "2026-05-13", "2026-06-23", "2026-07-14", "2026-08-20", "2026-09-22", "2026-09-24", "2026-10-21", "2026-12-19", "2026-12-31", "2027-10-21"]
keywords: ["apache", "benchmark", "benchmarks", "latency", "license", "pricing", "throughput"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [374, 418]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: ec886f423fc655c9fef6382f34d94fa256044ce55d7af2cf1f798dad84231be9
---

# 13. Version-release timeline (2025 → 2026-09-22)

## 13. Version-release timeline (2025 → 2026-09-22)

| Date | Event | Provenance |
|---|---|---|
| 2024-10-21 | Apache Pulsar 4.0 LTS released | `[official]` |
| 2025-04 | InfluxDB 3 Core GA | `[official]` |
| 2025-05-21 | DuckDB 1.3.0 (`DICT_FSST`) | `[official]` |
| 2025-05-29 | SQLite 3.50.0 | `[secondary]` |
| 2025-05 | Apache Spark 4.0.0 (ANSI default, VARIANT, JDK 17) | `[secondary]` |
| 2025-05 | Redis 8 tri-license (RSALv2/SSPLv1/AGPLv3) | `[secondary]` |
| 2025-09-16 | DuckDB 1.4 LTS (AES-256-GCM encryption) | `[official]` |
| 2025 | Elasticsearch 9.0 (breaking, ES\|QL) | `[secondary]` |
| 2025 | OpenSearch 3.0 under Linux Foundation (breaking, PPL/SQL) | `[secondary]` |
| 2025 | Apache Airflow 3.0 (DAG versioning, SDK imports, Data Assets) | `[official]`/`[secondary]` |
| 2025-12 | Apache Spark 4.1.0 (Declarative Pipelines, real-time streaming) | `[secondary]` |
| 2026-02 | Litestream v0.5.8 (latest observed); Redis 8.2 GA (Feb 17) | `[secondary]` |
| 2026-03-24 | Apache Pulsar 4.2 released | `[secondary]` |
| 2026-04-14 | Valkey 9.0 | `[secondary]` |
| 2026-04-27 | Pulsar Jetty 9.4 → 12.1.8 CVE upgrade | `[official]` |
| 2026-04-30 | ClickHouse 26.4 | `[official]` |
| 2026-05-13 | dbt Core 1.12.0-b1 | `[official]` |
| 2026-06 | NATS 2.12.11 / 2.14.x line active; 2.15.0-RC | `[official]` |
| 2026-06-23 | TimescaleDB 2.28.1 | `[secondary]` |
| 2026-07-14 | Apache Spark 4.2.0 (GEOMETRY, CHANGES, Java 25) | `[secondary]` |
| 2026-08-20 | dbt Core 1.12.3 (current) | `[official]` |
| 2026-09-24 | Apache Pulsar 4.2 support ends | `[secondary]` |
| 2026-10-21 | Apache Pulsar 4.0 LTS active support ends (security to 2027-10-21) | `[secondary]` |
| 2026-12-19 | dbt Core 1.11 end-of-support | `[secondary]` |
| 2026-12-31 | PHP 8.2 end-of-life (affects NATS PHP client support floor) | `[secondary]` |

## 14. Conflicts, gaps, and unverified claims register

- **SQLite 3.51.x WAL corruption fix (March 2026):** single secondary source; not verified against sqlite.org. Treat as `[unverified]` — do not cite as fact without checking the official changelog `[unverified]`.
- **"DuckDB 2.0" (2026):** found only on a fork; the official project announced 1.4 LTS in September 2025. Omit or flag `[unverified]` until an official announcement exists `[unverified]`.
- **Prisma 7 engine change (Rust → TypeScript/WASM):** from ecosystem search snippets; verify against prisma.io release notes before architectural reliance `[secondary]`.
- **Redis/Valkey performance comparisons:** vendor and secondary benchmarks conflict depending on workload; no independent head-to-head was located in this pass. RESP compatibility at the 7.2 baseline is well-evidenced; post-fork module/persistence parity is not `[secondary]`.
- **ClickHouse "2–49× Map key-lookup" claim:** from a secondary changelog mirror of 26.3 LTS notes; vendor/release-note evidence, not an independent benchmark `[secondary]`.
- **MongoDB 8.0 performance percentages:** vendor-reported; no independent reproduction found in this pass `[vendor-reported]`.
- **Meilisearch/Typesense pricing and latency:** vendor blogs and 2026 comparisons; date-sensitive and plan-dependent — re-check at decision time `[vendor-reported]`/`[secondary]`.
- **MotherDuck pricing/execution semantics:** no primary source captured in this pass; flagged as a gap — verify against motherduck.com before recommending `[unverified]`.
- **Turso/libSQL protocol matrix:** moves fast; no primary source captured — verify against official Turso docs `[unverified]`.
- **Gap — client-library version matrices:** per-language Kafka/Pulsar/RabbitMQ client compatibility tables were not exhaustively captured; pin client versions against each broker's official compatibility docs during upgrades (flagged, not researched) `[unverified]`.
- **Gap — Flyway current version:** the Flyway 11 line was not confirmed against an official source in this pass; verify at redgate.com/flyway before citing `[unverified]`.
- **Non-comparable benchmarks:** search-engine latency claims (Meilisearch vs Typesense vs Algolia) and streaming throughput claims (Redpanda vs Kafka) come from vendor or secondary sources with different harnesses — do not compare across sources `[independent]`.

