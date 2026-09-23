---
id: etape8-phasef-data-messaging/00-front-matter/8-event-driven-patterns-and-delivery-semantics
title: "8. Event-driven patterns and delivery semantics"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2024-03", "2025-05", "2026-02", "2026-04", "2026-09"]
keywords: ["consumer", "latency", "lean", "license"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [233, 290]
section: "Step 8 Phase F — Data & Messaging: Developer Angle"
sha256: 163ae321f2aa3b95b155d05947e757fb30459560b2d6db514a4c1957df736a4e
---

# 8. Event-driven patterns and delivery semantics

## 8. Event-driven patterns and delivery semantics

### 8.1 Delivery semantics vocabulary

- **At-most-once**: fire and forget; messages can be lost. **At-least-once**: retries until acked; duplicates possible. **Exactly-once**: each message's *effect* is applied once — in practice achieved as *effectively-once* via idempotence, not by the transport alone `[independent]`.
- A 2026 JetStream-focused doctrine note states it bluntly: "exactly-once is a myth, ack after durable commit, outbox/never-dual-write, DLQs, ordering, schema evolution, claim-check, idempotent consumers" — the industry consensus pattern catalog `[secondary]`.

### 8.2 Achieving effectively-once

- **Idempotent producer** (Kafka `enable.idempotence=true`, NATS `Nats-Msg-Id` dedup window, Pulsar producer-level dedup) eliminates duplicates *caused by producer retries*; it does not dedup across distinct logical publishes `[secondary]`.
- **Transactions** (Kafka transactions, Pulsar transactions) give atomic publish of multiple records and consume-transform-produce cycles; they add latency and operational complexity — use only where the business invariant requires atomicity `[independent]`.
- **Idempotent consumer** is the robust default: design the handler so reprocessing is safe (upsert by natural key, dedupe table on message ID, conditional writes) and ack only *after* the durable commit `[secondary]`.
- **Transactional outbox**: write the domain change and the outbound event in one local DB transaction; a relay publishes outbox rows to the broker. This avoids dual-write inconsistency without distributed transactions `[independent]`.
- **Inbox/dedupe table** on the consumer side: record processed message IDs in the same transaction as the side effect; skip already-seen IDs `[independent]`.

### 8.3 Ordering, retries, DLQs

- Partition/key ordering (Kafka keyed partitions, Pulsar Key_Shared, JetStream subject hierarchies) preserves order *within a key*; global ordering is a scalability anti-pattern `[independent]`.
- Retry with backoff + jitter; cap attempts; route poison messages to a **dead-letter queue/topic** with headers carrying the failure reason, attempt count, and original timestamp — never silently drop `[independent]`.
- RabbitMQ quorum queues support dead-lettering via policy; Kafka has no native DLQ — implement with a dedicated error topic in the consumer; NATS JetStream supports max-delivery attempts and dead-letter stream routing per consumer `[secondary]`.

### 8.4 Schema evolution and claim-check

- Use a schema registry (Confluent/Apicurio/Redpanda/Pulsar built-in) with compatibility modes (BACKWARD/FORWARD/FULL); never break consumers with unannounced field renames or type changes `[independent]`.
- **Claim-check pattern**: put large payloads in object storage and pass a reference through the broker; keeps broker storage lean and avoids message-size limits `[independent]`.
- Event versioning: version the event envelope, keep old handlers running during migration, and prefer additive changes; for breaking changes, publish a new event type and migrate consumers before producers `[independent]`.

## 9. Redis/Valkey developer angle

### 9.1 Fork landscape (2026)

- Valkey is the Linux Foundation, BSD-licensed fork of Redis OSS 7.2.4 (fork point March 2024); observed on the **9.x line in 2026** (9.0 April 2026, 9.1.x mid-2026, 9.1.2 newest as of September 2026) `[secondary]`.
- Redis 8 (May 2025) moved to tri-license RSALv2/SSPLv1/**AGPLv3**, folding formerly separate modules into core; Redis 8.2 GA arrived February 2026 `[secondary]`.
- Valkey and Redis remain RESP2/RESP3 wire-compatible at their shared 7.2 baseline; post-fork persistence/module compatibility must be verified per version, not assumed `[secondary]`.

### 9.2 Data structures and their uses

- **Strings**: counters (`INCR`), flags, serialized blobs; **Hashes**: object fields with partial updates; **Lists**: queues (but prefer Streams); **Sets/Sorted Sets**: membership, leaderboards (`ZRANGE`), rate limiting with sliding windows; **HyperLogLog**: cardinality estimation; **Geo**: proximity queries; **Bitmaps/Bitfields**: feature flags, analytics bitsets `[independent]`.
- **Streams** (`XADD`/`XREAD`/`XREADGROUP`): durable, replayable log with consumer groups — the right primitive for work queues and event ingestion, replacing `BRPOPLPUSH` list patterns `[independent]`.

### 9.3 Caching patterns

- **Cache-aside**: app checks cache, falls back to DB, populates cache; simplest, but first request pays the miss penalty `[independent]`.
- **Write-through/write-behind**: keep cache and DB consistent at write time; write-behind risks data loss on crash before flush `[independent]`.
- Invalidation: TTL-first (short TTLs bound staleness), explicit invalidation on write paths, and versioned keys for deploys; avoid cache stampede with request coalescing or probabilistic early refresh `[independent]`.

### 9.4 Pub/Sub vs Streams vs KV

- **Pub/Sub**: fire-and-forget broadcast, no persistence, no replay — for ephemeral notifications (cache invalidation signals, live dashboards); subscribers that are offline miss messages `[independent]`.
- **Streams**: durable with consumer groups, pending-entry lists, and `XCLAIM` for crashed-consumer recovery — for anything that must not be lost `[independent]`.
- **Keyspace notifications**: useful for expiry-driven workflows but unreliable under load — treat as hints, not guarantees `[independent]`.

### 9.5 Client notes

- Official and community clients exist for Python (`redis-py`, also used for Valkey), Node (`node-redis`, `ioredis`), Java (Jedis, Lettuce), Go (`go-redis`), Rust (`redis-rs`) `[secondary]`.
- Cluster mode changes key-space semantics: multi-key operations must hash to the same slot (`{hash tags}`); Lua scripts and transactions are slot-bound `[independent]`.
- Connection pooling and pipelining are the two highest-leverage client optimizations; monitor `connected_clients` and command latency, not just hit rate `[independent]`.

