---
id: etape8-phasef-data-messaging/06-spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra/overview
title: "Spark 4 default: ANSI mode on -> divide-by-zero and bad casts raise"
domain: spark-4-default-ansi-mode-on-divide-by-zero-and-bad-casts-ra
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-09-22"]
keywords: ["apache", "aws", "benchmarks", "consumer", "cost", "latency"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [683, 757]
section: "Spark 4 default: ANSI mode on -> divide-by-zero and bad casts raise"
sha256: 69e172acbd7f202d70235ad804851b069acc81c34b3c9a18a8c98a5186e01fb1
---

# Spark 4 default: ANSI mode on -> divide-by-zero and bad casts raise
```

### 18.9 TimescaleDB hypertable + continuous aggregate

```sql
CREATE TABLE metrics (
  ts TIMESTAMPTZ NOT NULL,
  host TEXT NOT NULL,
  cpu DOUBLE PRECISION
) WITH (tsdb.hypertable, tsdb.orderby = 'ts DESC');  -- 2.20+ best practice

CREATE MATERIALIZED VIEW metrics_hourly
WITH (timescaledb.continuous) AS
SELECT time_bucket('1 hour', ts) AS hour, host, avg(cpu) AS avg_cpu
FROM metrics GROUP BY 1, 2;

SELECT add_retention_policy('metrics', INTERVAL '90 days');
SELECT add_columnstore_policy('metrics', INTERVAL '7 days');  -- 2.18+ Hypercore
```

### 18.10 InfluxDB 3 write + SQL query

```bash
curl -X POST "http://localhost:8181/api/v3/write_lp?db=ops" \
  -H "Authorization: Bearer $TOKEN" \
  -d 'cpu,host=web1 usage=42.5 1758610000000000000'

curl -X POST http://localhost:8181/api/v3/query_sql \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"db":"ops","q":"SELECT mean(usage) FROM cpu WHERE time > now() - INTERVAL '\''1 hour'\'' GROUP BY host","format":"json"}'
```

## 19. CDC tooling and schema registries

### 19.1 Change data capture options

- **Database logical replication** (PostgreSQL logical decoding, MySQL binlog, SQL Server CDC): lowest latency, no application changes; the outbox table from Section 18.1 can be tailed via logical replication instead of polling `[independent]`.
- **Debezium** (Kafka Connect-based): streams row-level changes from MySQL, PostgreSQL, MongoDB, SQL Server into Kafka topics with before/after payloads and schema envelopes `[secondary]`.
- **Spark 4.2 `CHANGES` clause**: SQL-native CDC-style reads over Delta-like sources — useful when the warehouse/lake already versions data and you want incremental dbt models without a separate CDC service `[secondary]`.
- **ELT SaaS** (Fivetran, Airbyte): managed connectors that land raw data in the warehouse; tradeoff is cost and control vs engineering time — verify connector coverage for your source before committing `[secondary]`.
- CDC consumers must handle the same delivery semantics as any event consumer: idempotent sinks, schema evolution, and backfill/replay support `[independent]`.

### 19.2 Schema registry comparison

| Registry | Ecosystem | Compatibility modes | Notes |
|---|---|---|---|
| Confluent Schema Registry | Kafka | BACKWARD, FORWARD, FULL, NONE | REST API; Avro/Protobuf/JSON Schema |
| Apicurio Registry | Kafka, polyglot | Same modes | Red Hat/Apache 2.0; UI included |
| Redpanda built-in | Redpanda | Same modes | Kafka-protocol compatible endpoint |
| Pulsar built-in | Pulsar | Per-topic enforcement | Schema attached to topic, not a sidecar |
| AWS Glue Schema Registry | MSK/Kafka on AWS | Same modes | IAM-integrated; serializers for Java/Python |

- Enforce compatibility in CI: register the new schema version in a staging registry and run the compatibility check before deploying producers `[independent]`.
- Never delete old schema versions that consumers may still need for replay; retention policy for schemas should be explicit and versioned `[independent]`.
- For JSON events without a registry, embed a `schema_version` field in the envelope and keep a versioned codec table in code — weaker than a registry but better than unversioned JSON `[independent]`.

## 20. Cross-cutting developer checklist

- [ ] Every external claim in design docs carries a provenance tag and a date; vendor benchmarks are never presented as independent `[independent]`.
- [ ] Migrations are reviewable, reversible, one-logical-change-per-revision, and run identically in dev/CI/prod (Alembic/Flyway/Prisma Migrate) `[secondary]`.
- [ ] Zero-downtime schema changes follow expand → dual-write/backfill → switch readers → contract `[independent]`.
- [ ] Consumers ack only after the durable commit; handlers are idempotent via inbox/dedupe tables keyed on message ID `[secondary]`.
- [ ] No dual writes: use the transactional outbox (or CDC off the outbox) instead of writing to the DB and the broker in one code path `[independent]`.
- [ ] Poison messages route to a DLQ with failure metadata; retries use capped backoff + jitter `[independent]`.
- [ ] Broker upgrades pin client versions first (Kafka 4.x KRaft, Pulsar 4.x Jetty 12, NATS API levels, RabbitMQ 4.x quorum queues) `[official]`/`[secondary]`.
- [ ] Cache entries have TTLs and explicit invalidation paths; stampedes are prevented with coalescing or early refresh `[independent]`.
- [ ] Search indexes are derived read models with a rebuild-from-source path; ranking rules are versioned and fixture-tested `[independent]`.
- [ ] Time-series retention/columnstore policies are declared as code (TimescaleDB policies, InfluxDB retention), not clicked into existence `[secondary]`.
- [ ] Pipeline code separates parse-time from execution-time logic (Airflow), thinks in assets (Dagster), and tests dbt models with assertions `[secondary]`.
- [ ] Secrets, tokens, and connection strings live in the Secure Vault or env-injected config — never in migration files or notebooks `[independent]`.

---
*End of Step 8 Phase F — Data & Messaging (developer angle). Cutoff: 2026-09-22.*
