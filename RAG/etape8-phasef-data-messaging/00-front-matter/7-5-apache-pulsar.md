---
id: etape8-phasef-data-messaging/00-front-matter/7-5-apache-pulsar
title: "7.5 Apache Pulsar"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2024-10-21", "2026-04", "2026-04-30", "2026-07-06", "2026-09-24", "2026-10-21", "2027-10-21"]
keywords: ["apache", "consumer", "latency"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [213, 232]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 57142b59a47ef416812bff27ed57f4af5e30b4158e60a715b9a743590ecf389d
---

# 7.5 Apache Pulsar

- **RabbitMQ 4.0** removed classic queue mirroring (after a three-year deprecation): classic queues are now non-replicated; **quorum queues** and **streams** are the replicated data types `[official]`.
- The 4.x line had advanced to **4.3** by 2026; 4.3.0 notes include WebSocket origin validation, federation shutdown improvements, and Shovel enhancements; 4.1.8 was observed in third-party docs (2026-04-30) `[official]`/`[secondary]`.
- 4.0.3 added initial Erlang/OTP 27 support; Khepri is the schema/metadata data store; a "Local Random Exchange" type was observed in 2026 docs `[official]`/`[secondary]`.
- Developer angle: choose **quorum queues** for replicated work queues with at-least-once semantics and consumer acks; choose **streams** for replayable, append-only logs with offset tracking (Kafka-like consumption over AMQP infrastructure); classic queues remain for transient, non-critical buffering `[independent]`.

### 7.5 Apache Pulsar

- **Pulsar 4.0** (2024-10-21) is the second LTS release: fully modular architecture, built-in OpenTelemetry integration, enhanced Key_Shared subscriptions, advanced load balancing for millions of topic partitions, refined rate limiting `[official]`.
- Release state 2026: 4.0.12 (2026-07-06, latest 4.0.x), 4.1.3, 4.2.3 (2026-07-06); 4.0 is LTS with active support to 2026-10-21 and security support to 2027-10-21; 4.2 support ends 2026-09-24 `[secondary]`.
- April 2026 releases upgraded Jetty 9.4.x → 12.1.8 to address high-severity CVEs — a breaking change for `AdditionalServlet` plugin implementations and for Athenz auth (now requires Java 17+) `[official]`.
- Upgrade path: 2.10.6/2.11.3 → 3.0.7 → 4.0.0; live upgrade/downgrade is supported between consecutive LTS or feature versions, but must be tested per configuration `[official]`.
- Developer model: topics with partitioned logs on BookKeeper, three subscription modes (Exclusive, Shared, Failover) plus Key_Shared, multi-tenancy via tenants/namespaces, schema registry with compatibility checks, and an event-time-based topic compactor (PIP-352, contributed by StreamX) `[official]`/`[secondary]`.

### 7.6 Broker selection heuristic (developer view)

- Need Kafka ecosystem (Connect, Streams, ksqlDB, Schema Registry tooling) → Kafka or Redpanda (protocol-compatible) `[independent]`.
- Need flexible routing, per-message acks, and AMQP interop → RabbitMQ quorum queues/streams `[independent]`.
- Need low-latency pub/sub + durable streams + KV/Object in one binary with edge-friendly footprint → NATS JetStream `[independent]`.
- Need multi-tenancy, geo-replication, and tiered storage as first-class concepts → Pulsar `[independent]`.

