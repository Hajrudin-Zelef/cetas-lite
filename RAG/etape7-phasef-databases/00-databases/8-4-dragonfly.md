---
id: etape7-phasef-databases/00-databases/8-4-dragonfly
title: "8.4 Dragonfly"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2026-09"]
keywords: ["apache", "distribution", "latency", "license", "memory", "open source", "research", "throughput"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [464, 512]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: bc9d29baa797fea1798dc6a777a0038818d77d56004aa75d5fc358752e07f418
---

# 8.4 Dragonfly

### 8.4 Dragonfly

- **Dragonfly 2.0** announced **September 2026** (first user conference):
  +54% throughput, −35% average latency vs 1.x, 30–40% less memory on common
  workloads, 206→308 commands (vector/hybrid search, JSON, probabilistic),
  SSD tiering up to 8x effective capacity [vendor-reported].
- Multi-threaded shared-nothing architecture; scales vertically to large
  instances (8 GB → 768 GB/64-core cited) before clustering is needed
  [vendor-reported].
- **License: BSL 1.1** — free to self-host in production, converts to
  Apache-2.0 per release on a published change date; *not* OSI open source;
  managed-service offering by third parties restricted [official].
- ⚠️ **Name collision**: `dragonflyoss/dragonfly` is a *different* project
  — CNCF-Graduated (Oct 2025) P2P file/container-image distribution system
  (d7y.io, Apache-2.0, v2.4.0). Do not confuse the two [official].
- Production users cited (2026): Meesho, Agoda, Dailymotion, ShareChat
  [vendor-reported].

### 8.5 KeyDB, Upstash and others

- **KeyDB** (Snap): multithreaded Redis fork; continued maintenance under
  Snap in 2026, quieter than the Valkey/Dragonfly race [secondary].
- **Upstash**: serverless Redis/Valkey — **$0.20 per 100K commands**
  pay-as-you-go, fixed plans from **$10/mo**, free tier; positions itself as
  *service* vs Valkey-as-engine [vendor-reported].
- **Garnet** (Microsoft Research): .NET-based cache-store, open source (MIT),
  notable for thread-scalability research; adoption smaller than the big
  three [secondary][unverified on 2026 momentum].

### 8.6 Caching patterns (ops)

- **Cache-aside** (lazy load), **read/write-through**, **write-behind** —
  standard patterns; invalidation strategy (TTL vs explicit delete vs
  versioned keys) is the main correctness lever [secondary].
- **Persistence**: RDB snapshots (point-in-time, compact) vs AOF (append-
  only log, finer durability, larger files); production default is usually
  both, or AOF `everysec` [secondary].
- **Eviction**: `maxmemory-policy` — `allkeys-lru` for pure caches,
  `volatile-lru`/`volatile-ttl` when mixing persistent and cache keys
  [secondary].
- **HA**: **Sentinel** (automatic failover, 3+ sentinels) vs **Cluster**
  (sharded, 16384 slots, multi-key ops limited to same slot/hash tags)
  [secondary].
- **CDN interplay**: edge caches (Cloudflare, Fastly) sit in front for
  HTTP; Redis/Valkey sits behind for computed fragments, sessions, rate
  limits, leader locks [secondary].

---

