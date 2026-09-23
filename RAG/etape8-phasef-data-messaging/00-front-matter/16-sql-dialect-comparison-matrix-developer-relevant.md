---
id: etape8-phasef-data-messaging/00-front-matter/16-sql-dialect-comparison-matrix-developer-relevant
title: "16. SQL dialect comparison matrix (developer-relevant)"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "consumer", "latency", "license"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [489, 591]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 629cdc717ec60b30b7cd001ffa05d5465b6ebf8f9212a64b32251c5ba533c09a
---

# 16. SQL dialect comparison matrix (developer-relevant)

## 16. SQL dialect comparison matrix (developer-relevant)

| Feature | PostgreSQL | MySQL 8.x | SQLite | SQL Server (T-SQL) | DuckDB |
|---|---|---|---|---|---|
| Pagination | `LIMIT n OFFSET m` | `LIMIT m, n` / `LIMIT n OFFSET m` | `LIMIT n OFFSET m` | `OFFSET m ROWS FETCH NEXT n ROWS ONLY` | `LIMIT n OFFSET m` |
| Upsert | `ON CONFLICT ... DO UPDATE` | `ON DUPLICATE KEY UPDATE` | `ON CONFLICT ... DO UPDATE` | `MERGE` | `ON CONFLICT ... DO UPDATE` |
| Returning clause | `RETURNING` (full) | Not supported (use `LAST_INSERT_ID()`) | `RETURNING` (3.35+) | `OUTPUT` clause | `RETURNING` |
| Auto-increment | `GENERATED ... AS IDENTITY` / `SERIAL` | `AUTO_INCREMENT` | `INTEGER PRIMARY KEY` (rowid alias) | `IDENTITY(1,1)` | `SEQUENCE` / auto-increment |
| Boolean type | Native `BOOLEAN` | `TINYINT(1)` convention | Integer 0/1 convention | `BIT` | Native `BOOLEAN` |
| JSON | `jsonb` + rich operators | `JSON` type + functions | `json1` extension | `JSON` functions (no native type until 2025+) | Native nested + JSON |
| Placeholders | `$1, $2` | `?` | `?` / `:name` | `@name` | `?` / `$1` |
| DDL transactions | Yes (most DDL) | No (implicit commit) | Limited | Yes | Yes |

- Placeholder syntax is the most common portability bug when switching drivers; always use the driver's parameter style, never string-format values into SQL `[independent]`.
- DDL transactionality matters for migrations: PostgreSQL can roll back a failed migration's DDL; MySQL cannot — design MySQL migrations to be re-runnable instead `[independent]`.
- `RETURNING` availability determines whether your repository layer needs a second `SELECT` after insert; abstract it behind the ORM or a dialect-aware insert helper `[independent]`.

### 16.1 Connection management (all dialects)

- Never open a connection per request without pooling: use PgBouncer/pgpool (PostgreSQL), ProxySQL (MySQL), or driver-level pools (SQLAlchemy `QueuePool`, Prisma connection pool, HikariCP) `[independent]`.
- Serverless/edge workloads need external poolers (Neon/Supabase poolers, Prisma Accelerate) because function instances cannot hold long-lived pools `[secondary]`.
- Set statement timeouts (`statement_timeout` in PostgreSQL) so a bad query fails instead of holding pool slots forever `[independent]`.
- SQLite concurrency: WAL mode + `busy_timeout` for multi-threaded single-process access; for multi-process writes, serialize through one writer process `[independent]`.

## 17. Decision matrices

### 17.1 Primary database selection

| Need | Pick | Why |
|---|---|---|
| General OLTP, JSON, extensions | PostgreSQL (+ TimescaleDB for time-series) | Richest type system, DDL transactions, ecosystem |
| Embedded / edge / single-writer | SQLite (+ Litestream for DR) | Zero ops, WAL reads, file-level backup story |
| In-process analytics | DuckDB | Columnar, zero server, reads Parquet natively |
| Distributed analytics serving | ClickHouse | MergeTree engines, materialized views, Kafka ingest |
| Document model, flexible schema | MongoDB 8.x | Aggregation pipelines, change streams, Atlas search |
| High-cardinality metrics | InfluxDB 3 | SQL + line protocol, object-storage economics |

### 17.2 Messaging/streaming selection

| Need | Pick | Why |
|---|---|---|
| Kafka ecosystem + exactly-once transactions | Kafka 4.x (KRaft) or Redpanda | Connect/Streams/ksqlDB tooling; Redpanda for single-binary ops |
| Flexible routing, AMQP interop | RabbitMQ 4.x quorum queues / streams | Per-message ack, exchanges, mature management UI |
| Low-latency pub/sub + KV/Object in one binary | NATS JetStream | Tiny footprint, API-level feature gating, edge-friendly |
| Multi-tenancy + geo-replication + tiered storage | Pulsar 4.x LTS | Tenants/namespaces, BookKeeper, built-in schema registry |

### 17.3 Search selection

| Need | Pick | Why |
|---|---|---|
| App search, fastest setup, permissive license | Meilisearch | MIT, typo-tolerant, SDK breadth |
| App search + vector/hybrid + Raft clustering | Typesense | Native vector search, GPL-3.0 self-host |
| Search inside observability/analytics platform | Elasticsearch 9 / OpenSearch 3 | ES\|QL vs PPL/SQL; license (Elastic vs Apache 2.0) decides |

### 17.4 Orchestration selection

| Need | Pick | Why |
|---|---|---|
| Task-centric, largest provider ecosystem | Airflow 3 | Operators/providers, event-driven Data Assets, DAG versioning |
| Asset-centric, lineage-first, dbt unification | Dagster | Assets/checks/partitions, `dagster-dbt`, `dg` CLI |
| SQL-only transforms in the warehouse | dbt Core 1.12 | `ref()` DAG, tests, snapshots; orchestrated by either of the above |

## 18. Developer code patterns

### 18.1 Transactional outbox (PostgreSQL)

```sql
-- Outbox table: written in the same transaction as the domain change
CREATE TABLE outbox (
  id          BIGSERIAL PRIMARY KEY,
  aggregate   TEXT NOT NULL,
  aggregate_id TEXT NOT NULL,
  event_type  TEXT NOT NULL,
  payload     JSONB NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  published_at TIMESTAMPTZ
);
CREATE INDEX ON outbox (published_at) WHERE published_at IS NULL;
-- Relay: SELECT ... WHERE published_at IS NULL ORDER BY id FOR UPDATE SKIP LOCKED
-- publish to broker, then UPDATE outbox SET published_at = now() WHERE id = ...
```

- `FOR UPDATE SKIP LOCKED` lets multiple relay workers share the outbox without blocking `[independent]`.
- Keep the payload self-contained (event-carried state transfer) so consumers rarely need to call back `[independent]`.

### 18.2 Idempotent consumer sketch (pseudo-code)

```python
def handle(event):
    with db.transaction() as tx:
        # inbox/dedupe in the same transaction as the side effect
        if tx.execute("SELECT 1 FROM inbox WHERE msg_id = %s", event.id):
            return  # already processed -> safe to ack
        apply_business_logic(tx, event)
        tx.execute("INSERT INTO inbox (msg_id) VALUES (%s)", event.id)
    ack(event)  # ack only after the durable commit
```

- Ack-after-commit is the single most important consumer rule across Kafka, JetStream, RabbitMQ, and Pulsar `[secondary]`.

### 18.3 Kafka producer/consumer config that matters

```properties
