---
id: etape8-phasef-data-messaging/00-front-matter/5-clickhouse
title: "5. ClickHouse"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-04-30"]
keywords: ["agentic", "apache", "consumer", "latency", "license", "mit license", "pricing"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [150, 212]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 5470722c9a109add640cc025b33b9e5d7e20fcc2b2d086cdad93c7d640cc3c3f
---

# 5. ClickHouse

## 5. ClickHouse

### 5.1 Release state (2026)

- ClickHouse's official changelog documented release **26.4 on 2026-04-30** with breaking changes and new functionality `[official]`.
- A secondary mirror of the 26.3 LTS notes described bucketed Map serialization with claimed 2–49× key-lookup improvements — preserve as release-note/vendor evidence, not independent benchmarking `[secondary]`.

### 5.2 Developer angle

- Columnar MergeTree engine family (`MergeTree`, `ReplacingMergeTree`, `AggregatingMergeTree`, `SummingMergeTree`) is chosen per-table at `CREATE TABLE` time; the choice determines deduplication and pre-aggregation behavior `[secondary]`.
- Materialized views and projections precompute aggregations at insert time — the primary mechanism for sub-second dashboards over billions of rows `[secondary]`.
- Native clients: HTTP interface (simple, load-balancer friendly), native TCP protocol (fastest), and official clients for Python, Node, Go, Java, Rust; the HTTP interface accepts raw SQL in the POST body, which makes debugging trivial with curl `[secondary]`.
- Distributed tables + sharding are configured via cluster definitions in XML/YAML; the developer-visible abstraction is the `Distributed` table engine routing queries across shards `[secondary]`.
- Inserts are batch-oriented: prefer large batches (thousands+ rows) over row-by-row inserts; async inserts buffer small writes server-side `[independent]`.
- Kafka engine / Kafka table function can consume topics directly into ClickHouse, a common pattern for real-time ingestion without an intermediate consumer service `[secondary]`.

## 6. Search engines

### 6.1 Elasticsearch vs OpenSearch (2025–2026 state)

- **Elasticsearch 9.0** (2025) introduced breaking changes and emphasizes **ES|QL** as its query language direction `[secondary]`.
- The official Elasticsearch Java client 9.0 release notes document `Rest5Client`, a simplified client builder, and BulkIngester retry policies — the client surface was modernized alongside the server `[official]`.
- **OpenSearch 3.0** (2025, first major release under the Linux Foundation) introduced breaking changes; it emphasizes PPL/SQL query surfaces, permissive Apache 2.0 licensing, and ML/search plugins `[secondary]`.
- License divergence remains the strategic fact: OpenSearch is Apache 2.0 (Linux Foundation); Elasticsearch uses Elastic License/SSPL terms — this drives procurement decisions more than raw feature deltas `[secondary]`.

### 6.2 Meilisearch vs Typesense (application search)

- **Meilisearch**: Rust engine, MIT license, hosted Cloud from $20/month; REST API, typo-tolerant search, faceting, ranking controls; client SDKs for JS, Python, PHP, Ruby, Go, Rust, Swift; InstantSearch React adapter `[vendor-reported]`.
- **Typesense**: C++ engine, GPL-3.0 for self-hosting; Cloud is resource-based (e.g., ~$21.60/month for a small node per vendor blog) with 720 free cluster hours; native vector/hybrid search; Raft-based clustering `[vendor-reported]`.
- A 2026 comparison framed the tradeoff as: Typesense wins on raw performance, Meilisearch is easiest to get running; both are far cheaper than Algolia (claimed ~10×) — pricing and latency claims are vendor/secondary and must be dated `[secondary]`.
- Developer choice heuristic: pick Meilisearch for fastest time-to-search with a permissive license; pick Typesense when you need native vector/hybrid search or Raft clustering semantics; pick Elasticsearch/OpenSearch when search is one workload inside a larger observability/analytics platform `[independent]`.

### 6.3 Search developer patterns

- Index design: separate searchable vs displayed vs filterable attributes; use distinct indexes per tenant or per document type when ranking rules differ `[independent]`.
- Keep the search engine as a derived read model: the system of record stays in the primary database; a CDC/outbox pipeline feeds the index, and rebuild-from-source must be a supported operation `[independent]`.
- Synonyms, typo tolerance, and ranking rules are product decisions — version them like code and test them with a relevance fixture set, not by eyeballing the demo UI `[independent]`.

## 7. Event streaming and messaging

### 7.1 Apache Kafka 4.x (KRaft era)

- **Kafka 4.0 supports KRaft only** — ZooKeeper mode was removed; the new group coordinator and the next-generation consumer rebalance protocol reached GA in 4.0 `[official]`.
- Existing ZooKeeper clusters must migrate through a compatible 3.x release before moving to Kafka 4.x; the official upgrade guide is the authoritative path `[official]`.
- Developer consequences: broker configuration surface changed (no ZK connect strings), consumer group protocol v2 semantics, and client libraries must be recent enough for the new coordinator — pin client versions in lockstep with broker upgrades `[independent]`.
- Key producer APIs: idempotent producer (`enable.idempotence=true`), transactions (`transactional.id`), and the outbox/transactional-messaging patterns built on them; key consumer APIs: cooperative-sticky rebalance, manual offset control, and `ConsumerRebalanceListener` for state cleanup `[secondary]`.

### 7.2 Redpanda

- Redpanda is a Kafka-API-compatible streaming platform in C++ (no JVM, no ZooKeeper); observed release **v26.2.1** with deferred upgrade finalization, schema-registry shadowing, role synchronization, and Cloud Topics shadowing `[official]`.
- Secondary 2026 reporting described Redpanda One / Agentic Data Plane features — label as `[secondary]`/`[vendor-reported]`, not independent performance proof `[secondary]`.
- Developer angle: drop-in Kafka protocol compatibility means existing Kafka clients and Kafka Connect-style connectors work unchanged; the differentiators are operational (single binary, tiered storage to object stores) rather than API-level `[vendor-reported]`.

### 7.3 NATS JetStream

- NATS Server lines in 2026: **2.12.x** continues as a maintenance line (2.12.11 observed ~2026-06, patch line up to 2.12.15), **2.14.x** is the new main line (2.13 was skipped), and **2.15.0-RC** exists; the v2 ack-subject format becomes the default in 2.15, which affects ACLs on `$JS.ACK.<stream>.>`/`$JS.FC.<stream>.>` `[official]`.
- JetStream capability gating is by server **API level**, not version string: 2.11.x = level 1, 2.12.0–2.12.4 = level 2, 2.12.5–2.12.15 = level 3, 2.14.x = level 4, 2.15.0-RC = level 5; the level bumped mid-patch-line (2.12.4→2.12.5), so hardcoded version tables go stale `[secondary]`.
- Feature-to-level map developers must check at runtime: per-message TTL (`allow_msg_ttl`, 2.11), subject delete markers (2.11), pull consumer priority groups (2.11/2.12), atomic batch publish (`allow_atomic`, 2.12), distributed counter CRDTs (`allow_msg_counter`, 2.12), scheduled/cron publishing (`allow_msg_schedules`, 2.12; recurring schedules from 2.14), KV/Object stores (GA, JetStream-backed) `[secondary]`.
- Client guidance: the `jetstream` Go package is the current API; the older `nc.JetStream()` JetStreamContext is legacy; publish deduplication uses the `Nats-Msg-Id` header (available since 2.2) `[secondary]`.
- Two silent failure modes to code around: `Nats-TTL` without `allow_msg_ttl` is silently ignored (message never expires), and KV per-key TTL without `markerTTL` expires the key with no watcher event — check stream config, not just the API level `[secondary]`.

### 7.4 RabbitMQ

