---
id: etape8-phasef-data-messaging/00-front-matter/10-mongodb-developer-angle
title: "10. MongoDB developer angle"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2025-04", "2025-05", "2026-03", "2026-05-13", "2026-05-14", "2026-06", "2026-07-14", "2026-08-20"]
keywords: ["agentic", "apache", "aws", "compute", "cost", "inference", "lean", "quantization"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [291, 373]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 92fbc7b35440fe2c28a4437e967e2bb101f247fbf389ef4ac792c66be6ee61a5
---

# 10. MongoDB developer angle

## 10. MongoDB developer angle

### 10.1 MongoDB 8.x server

- MongoDB 8.0 brought improved read/bulk-write/time-series performance, query settings, range-capable Queryable Encryption, and vector quantization for Atlas Vector Search — performance percentages in vendor material are vendor-reported, not independent `[vendor-reported]`.
- Mongoose's official compatibility table: **MongoDB 8.x requires Mongoose `^8.7.0` or `^9.0.0`**; MongoDB 7.x works with `^7.4.0 | ^8.0.0 | ^9.0.0` `[official]`.
- Mongoose relies on the MongoDB Node.js driver; driver/server compatibility should be checked against the official driver compatibility matrix `[official]`.

### 10.2 Mongoose ODM (Node.js)

- Mongoose 8.x is the current ODM line (a 2026 production upgrade guide pinned `mongoose: ^8.20.0`); it uses MongoDB Node driver v6 `[secondary]`/`[official]`.
- Breaking changes in Mongoose 8: removed `findOneAndRemove()`/`findByIdAndRemove()` (use `findOneAndDelete()`/`findByIdAndDelete()`), removed `Model.count()`/`Query.count()` (use `countDocuments()`), removed the `id` setter, and `new ObjectId()` no longer accepts 12-char strings; deprecated `ssl*` options removed in favor of `tls*` equivalents `[official]`.
- Modeling guidance: design documents around query shape, not relational purity; denormalize read-heavy joins; avoid hydrating full documents on hot paths (use `.lean()` + projections); add middleware only with write-path tests since hooks silently alter create/update/migration behavior `[secondary]`.
- Index discipline: unique indexes for identity fields (e.g., `slug`), sparse indexes for optional filters, TTL indexes for auto-expiry collections `[secondary]`.

### 10.3 Transactions and consistency

- Multi-document ACID transactions exist but carry performance costs; the document model is designed so most operations fit in a single document — reach for transactions only when cross-document atomicity is a real invariant `[independent]`.
- Set `setFeatureCompatibilityVersion` appropriately when upgrading (e.g., `"8.0"`) before enabling new-version features `[secondary]`.

## 11. Time-series databases

### 11.1 InfluxDB 3

- **InfluxDB 3 Core is GA since April 2025** and is the recommended platform for new time-series workloads; built on Rust, Apache Arrow, and DataFusion, storing Parquet on object storage (S3/Azure/GCP) or local disk `[official]`.
- Query languages: **SQL and InfluxQL** — Flux was removed in v3; v1/v2 write APIs remain compatible (line protocol) and the v1 query API (InfluxQL) is supported `[official]`.
- Developer surface: HTTP API on port **8181** (not 8086); write via `POST /api/v3/write_lp?db=...` with line protocol and `Bearer` auth; query via `POST /api/v3/query_sql` (JSON) or gRPC/FlightSQL; databases are created implicitly on first write `[official]`/`[secondary]`.
- InfluxDB 3 Core has no bundled CLI — all operations go through the HTTP REST API; an embedded Python VM supports plugins and triggers `[secondary]`.
- Migration from v1/v2: dual-write new data to both, backfill history into v3 (Enterprise recommended for historical queries); Prometheus `remote_write` endpoint does not exist in v3 — use Telegraf with the `outputs.influxdb_v3` plugin `[secondary]`.

### 11.2 TimescaleDB

- TimescaleDB is a PostgreSQL extension for time-series/event analytics; the latest upstream release found in a June 2026 survey was **2.28.1** (patch after 2.28.0) `[secondary]`.
- Key versioned features: **2.18** introduced Hypercore columnstore (`add_columnstore_policy()` superseding legacy compression APIs); **2.20** added `CREATE TABLE ... WITH (tsdb.hypertable)` as the best-practice hypertable creation path; **2.23** auto-selects the first timestamp column as partition column; **2.28** added faster `first()`/`last()` on compressed data, incremental manual continuous-aggregate refreshes, vectorized `CASE`, and generated aggregate columns on continuous aggregates without rebuilds `[secondary]`.
- The official changelog for recent versions documents removal of adaptive chunking, `ANALYZE`/`VACUUM` support on continuous aggregates (redirected to the materialization hypertable), and incremental `refresh_continuous_aggregate()` `[official]`.
- Developer patterns: hypertables auto-partition by time (`SELECT create_hypertable('t', by_range('ts'))` or the 2.20+ `WITH (tsdb.hypertable)` syntax); `time_bucket()` for downsampling; continuous aggregates as incrementally refreshed materialized views; retention and columnstore policies as background jobs `[secondary]`.
- Python access: the `timescaledb` package (MIT) provides SQLAlchemy dialects (`timescaledb://`, `timescaledb+psycopg://`, `timescaledb+asyncpg://`) plus `TimescaleModel` helpers for hypertable creation and Hypercore policies `[secondary]`.
- 2.28.0 is the final minor line supporting PostgreSQL 15; future 2.29.x targets PostgreSQL 16/17/18 only `[secondary]`.

### 11.3 InfluxDB vs TimescaleDB (developer view)

- Choose InfluxDB 3 when you want a purpose-built, schemaless-ingest time-series engine with SQL + line protocol and object-storage economics `[independent]`.
- Choose TimescaleDB when you want full PostgreSQL (joins, transactions, existing tooling/ORMs) with time-series superpowers — at the cost of operating PostgreSQL `[independent]`.

## 12. Data pipelines: dbt, Airflow, Dagster, Spark, ETL vs ELT

### 12.1 dbt

- dbt Core release state 2026: **1.12.3** (2026-08-20) is current; 1.12.0-b1 appeared 2026-05-13; 1.10.21 was a late patch (2026-05-14) `[official]`.
- Support lifecycle: 1.10 EOL June 16 2026; 1.11 end-of-support December 19 2026; 1.12 supported through July 16 2027; cadence is one minor every ~6 months with >50% of users on the latest minor `[secondary]`/`[official]`.
- 2026 themes: dbt Fusion engine (Rust-based, in preview), semantic-layer YAML, function nodes, schema validation, and parser changes; Snowflake documents `DBT_VERSION` pinning for managed dbt on Snowflake (March 2026) `[official]`/`[secondary]`.
- Developer model: SQL `select` statements as models, Jinja templating, `ref()`/`source()` for DAG lineage, tests as assertions, snapshots for slowly-changing dimensions, seeds for static data; `dbt build` runs the DAG with dependency order `[secondary]`.

### 12.2 Apache Airflow 3

- **Airflow 3.0** (2025) redefined task execution: **DAG versioning** (each run pins the DAG version it started with — auditability), **backfills** executable from CLI and UI as first-class objects, and **SDK-first imports** (`from airflow.sdk import dag, task`) decoupling DAG code from the metadata DB `[secondary]`/`[official]`.
- **Data Assets** (AIP-74, evolution of Datasets) with **Watchers** enable **event-driven scheduling** (AIP-82): a Common Message Bus interface with an out-of-the-box AWS SQS implementation triggers DAGs on external events `[official]`.
- **TaskFlow API** is now the default authoring style: plain Python functions, implicit dependencies, natural XCom handling; code outside tasks runs at parse time, code inside operators at execution time `[secondary]`.
- Inference/ML support: non-data-interval DAGs (AIP-83 removed the execution-date uniqueness constraint), and provider packages (e.g., `common.ai`) ship example agentic DAGs with LLM operators and human-approval steps `[official]`.
- Migration note: a 2025–2026 field migration from 2.10.3 to 3.0.6 highlighted the SDK import shift and the decoupled execution model as the biggest developer-facing changes `[secondary]`.

### 12.3 Dagster

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

