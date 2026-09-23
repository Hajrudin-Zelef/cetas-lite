---
id: etape8-phasef-data-messaging/00-front-matter/15-glossary-and-verbatim-source-index
title: "15. Glossary and verbatim source index"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2025-05-21", "2025-09-11", "2025-09-16", "2026-03-02", "2026-08-27"]
keywords: ["agentic", "apache", "claude", "consumer", "embedding", "pricing", "throughput"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [419, 488]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: dba06d15a3b321a8c5a98f3ea7704042796eb81edab23b9f1c6fef39631f2d6a
---

# 15. Glossary and verbatim source index

## 15. Glossary and verbatim source index

### 15.1 Glossary

- **WAL** — write-ahead log; durability mechanism where changes are appended to a log before being applied.
- **LTX** — Litestream transaction file format for replicated SQLite WAL segments.
- **KRaft** — Kafka's Raft-based metadata quorum replacing ZooKeeper.
- **Outbox pattern** — writing domain changes and outbound events in one local transaction, then relaying events to the broker.
- **DLQ** — dead-letter queue/topic for poison messages.
- **CDC** — change data capture; streaming row-level database changes.
- **Idempotent consumer** — a handler whose repeated execution has the same effect as a single execution.
- **Claim-check** — passing large payloads by reference (object storage) instead of embedding them in messages.
- **Hypertable** — TimescaleDB's auto-partitioned table abstraction over PostgreSQL.
- **Continuous aggregate** — TimescaleDB's incrementally refreshed materialized view for time-series rollups.
- **RESP** — Redis Serialization Protocol; the wire protocol shared by Redis and Valkey.
- **ES|QL / PPL** — Elasticsearch's and OpenSearch's respective pipe-style query languages.
- **AIP** — Airflow Improvement Proposal (e.g., AIP-74 Data Assets, AIP-82 event-driven scheduling, AIP-83 execution-date uniqueness removal).
- **PIP** — Pulsar Improvement Proposal (e.g., PIP-352 event-time-based compaction).
- **DAG (migrations)** — directed acyclic graph of versioned schema migrations (Alembic revisions, dbt models, Airflow workflows).
- **FreshnessPolicy** — Dagster 1.11 GA API declaring how fresh an asset must be.
- **VARIANT** — Spark 4.0 semi-structured data type.

### 15.2 Verbatim source index

- https://github.com/apache/kafka/blob/HEAD/docs/getting-started/upgrade.md
- https://github.com/apache/kafka/blob/HEAD/docs/getting-started/compatibility.md
- https://github.com/redpanda-data/redpanda/releases/tag/v26.2.1
- https://github.com/duckdb/duckdb-web/blob/HEAD/_posts/2025-05-21-announcing-duckdb-130.md
- https://github.com/duckdb/duckdb-web/blob/HEAD/_posts/2025-09-16-announcing-duckdb-140.md
- https://clickhouse.com/docs/whats-new/changelog
- https://github.com/altinity/clickhouse/blob/HEAD/docs/changelogs/v26.3.2.3-lts.md
- https://github.com/elastic/elasticsearch-java/blob/HEAD/docs/release-notes/9-0-0.md
- https://github.com/rabbitmq/rabbitmq-server/blob/HEAD/release-notes/4.0.1.md
- https://github.com/rabbitmq/rabbitmq-server/blob/HEAD/release-notes/4.3.0.md
- https://github.com/nats-io/nats-server/releases
- https://github.com/nats-io/nats.c/releases
- https://github.com/lithqube/nats-jetstream-claude-skills/blob/HEAD/jetstream-architecture/concepts/server-features.md
- https://github.com/martinholovsky/sota-skills/blob/HEAD/skills/sota-architecture/rules/08-nats-jetstream.md
- https://pulsar.apache.org/blog/2024/10/24/announcing-apache-pulsar-4-0/
- https://github.com/apache/pulsar/releases
- https://streamnative.io/blog/announcing-apache-pulsar-tm-4-0-towards-an-open-data-streaming-architecture
- https://www.influxdata.com/blog/influxdb-3-oss-ga/
- https://github.com/influxdata/influxdb/blob/HEAD/README.md
- https://github.com/influxdata/docs-v2/blob/HEAD/content/influxdb3/_index.md
- https://github.com/timescale/timescaledb/blob/HEAD/CHANGELOG.md
- https://github.com/jmitchel3/timescaledb-python/blob/HEAD/README.md
- https://github.com/jmitchel3/timescaledb-python/blob/HEAD/docs/timescale-recent-updates.md
- https://github.com/dbt-labs/dbt-core/releases
- https://github.com/dbt-labs/dbt-core/blob/HEAD/docs/roadmap/2025-12-magic-to-do.md
- https://docs.snowflake.com/en/release-notes/2026/other/2026-03-02-dbt-core-versions
- https://github.com/apache/airflow-site/blob/HEAD/landing-pages/site/content/en/blog/airflow-three-point-oh-is-here/index.md
- https://github.com/apache/airflow-site/blob/HEAD/landing-pages/site/content/en/blog/agentic-workloads-airflow-3/index.md
- https://github.com/stars1233/dagster/blob/HEAD/CHANGES.md
- https://github.com/datazip-inc/olake-docs/blob/HEAD/blog/2026-08-27-apache-spark-alternatives.mdx
- https://github.com/Automattic/mongoose/blob/master/docs/compatibility.md
- https://github.com/automattic/mongoose/blob/HEAD/docs/migrating_to_8.md
- https://www.meilisearch.com/blog/typesense-pricing
- https://www.meilisearch.com/blog/algolia-vs-typesense
- https://github.com/baekenough/oh-my-customcode/blob/HEAD/guides/alembic/README.md
- https://github.com/bcasci/litestream-ruby/issues/1
- https://github.com/jaywedgeworth22/usage-monitor/blob/HEAD/docs/litestream.md
- https://github.com/indirect/andre.arko.net/blob/HEAD/content/post/2025-09-11-rails-on-sqlite-exciting-new-ways-to-cause-outages.md
- https://github.com/cabeda/convocados/blob/HEAD/docs/runbooks/point-in-time-recovery.md
- https://github.com/rzid/ma-famille/blob/HEAD/docs/adr/0004-sqlalchemy-alembic.md
- https://www.infoworld.com/article/4145050/migrating-from-apache-airflow-v2-to-v3.html
- https://devops-daily.com/posts/is-valkey-ready-to-replace-redis-2026
- https://www.devclass.com/ai-ml/2025/05/07/opensearch-30-hits-first-major-release-under-linux-foundation-as-it-battles-elasticsearch-for-mindshare/1622904
- https://coralogix.com/guides/elasticsearch/elasticsearch-vs-opensearch-key-differences/
- https://sdtimes.com/data/mongodb-8-0-offers-significant-performance-improvements-to-read-throughput-bulk-writes-and-more/

