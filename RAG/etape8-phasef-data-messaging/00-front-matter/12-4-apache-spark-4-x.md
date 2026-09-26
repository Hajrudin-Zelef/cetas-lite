---
id: etape8-phasef-data-messaging/00-front-matter/12-4-apache-spark-4-x
title: "12.4 Apache Spark 4.x"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2025-05", "2026-07-14"]
keywords: ["apache", "compute"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [354, 373]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 25f36e578d1fcc2c01d76e23104337b20d16b60d8aa249f092972a67b4e43381
---

# 12.4 Apache Spark 4.x

- Dagster 1.11.x was current in 2026; the `dg` CLI is now the recommended project workflow (`dg init`, `dg scaffold asset/component`, `dg build`, `dg dev`), superseding `dagster project scaffold` `[secondary]`.
- 1.11 themes: **FreshnessPolicies GA** (superseding legacy freshness checks; `FreshnessDaemon` runs by default), configurable backfill run config, and time-based partition exclusions (cron or datetime exclusion sets for custom calendars) `[secondary]`.
- Developer model is **asset-centric**: assets (materialized data objects), ops, jobs, resources, partitions, sensors, schedules, and asset checks (data-quality gates) — lineage across Python assets and dbt models (`dagster-dbt` loads dbt models as assets from `manifest.json`) `[secondary]`.
- Choose Dagster when the team thinks in datasets not jobs, lineage is central, and dbt + Python must share one graph `[independent]`; choose Airflow when the ecosystem (providers, operators) and task-centric orchestration matter most `[independent]`.

### 12.4 Apache Spark 4.x

- **Spark 4.0.0** (May 2025): ANSI SQL mode on by default, `VARIANT` type for semi-structured data, Spark Connect maturity (polyglot client-server: Python, Scala, Go, Swift, Rust), lightweight 1.5 MB `pyspark-client`, new Python Data Source API and UDTFs, `transformWithState` streaming API, native XML data source, string collation support `[secondary]`.
- **Spark 4.1.0** (Dec 2025): Spark Declarative Pipelines and the first official real-time mode for Structured Streaming; 1,800+ tickets closed `[secondary]`.
- **Spark 4.2.0** (July 14, 2026): native `GEOMETRY`/`GEOGRAPHY` types, SQL `CHANGES` clause for CDC-style reads, Arrow-optimized Python UDFs enabled by default, Java 25 support `[secondary]`.
- Hard requirements: JDK 17+ and Scala 2.13 since 4.0 — upgrade the runtime before the engine `[secondary]`.
- Developer guidance: prefer DataFrame/SQL APIs over RDDs; use Spark Connect (`SparkSession.builder.remote("sc://...")`) to decouple thin clients from cluster binaries; test ANSI-mode behavior changes (divide-by-zero, casts) before upgrading pipelines `[independent]`.

### 12.5 ETL vs ELT

- **ETL** (extract-transform-load): transform before loading, typically in a dedicated engine; fits strict schemas, legacy warehouses, and regulated pipelines where bad data must never land `[independent]`.
- **ELT** (extract-load-transform): load raw data first (lake/warehouse), transform with SQL in the warehouse (dbt's home turf); fits cloud warehouses with cheap storage and elastic compute `[independent]`.
- 2026 consensus: ELT dominates new analytics builds (dbt + warehouse + orchestrator); ETL persists for streaming/edge preprocessing and compliance-gated ingestion `[secondary]`.
- CDC (Debezium, database logical replication, Spark 4.2 `CHANGES`) feeds both: stream row-level changes into the lake/warehouse, then dbt models transform them `[independent]`.

