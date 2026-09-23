---
id: etape7-phasef-databases/00-databases/8-1-the-license-split-what-happened
title: "8.1 The license split — what happened"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: licenses
actors: ["AWS", "Google", "Oracle"]
dates: ["2024-03", "2024-11", "2025-03", "2025-05", "2026-05"]
keywords: ["license", "aws", "latency", "memory"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [416, 463]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: b8c16d80eab98cd2803da7d25a5b6dbf2435b2b27f48bf553ad2ec784b8361dd
---

# 8.1 The license split — what happened

### 8.1 The license split — what happened

- **March 2024**: Redis Ltd. relicensed Redis from BSD-3 to dual
  **RSALv2/SSPLv1** (starting with 7.4), ending OSI-open-source status
  [independent].
- Within days, AWS/Google Cloud/Oracle/Ericsson/Snap contributors forked the
  last BSD commit (Redis 7.2.4) as **Valkey** under the **Linux Foundation**,
  staying **BSD-3** [independent].
- **November 2024**: Redis creator Salvatore Sanfilippo (antirez) rejoined
  Redis Ltd. [independent].
- **1 May 2025**: Redis 8.0 added **AGPLv3** as a third license option
  (tri-license: RSALv2/SSPLv1/AGPLv3) and folded the former Redis Stack
  modules (JSON, search, time series, probabilistic, vector sets) into core
  [independent].
- Practical licensing read (2026, not legal advice): using unmodified Redis
  as a cache/DB does not trigger SSPL's service-copyleft in most readings;
  AGPL only bites if you modify the source and offer it over a network;
  distributions (Fedora, Debian, Ubuntu) moved to Valkey because they cannot
  ship source-available software in main [secondary].

### 8.2 Valkey — 2026 status

- Releases: **Valkey 8.1** (31 March 2025), **9.1** (May 2026), **9.1.1**
  current line; 8.1.x maintenance branch [secondary].
- Performance: published numbers put Valkey 8.x ~**8% more ops/sec**,
  ~**22% lower P99 latency**, ~**20% less memory** than Redis OSS; Valkey 8.0
  claimed ~1.2M req/s on AWS hardware (~230% over the 7.2 baseline)
  [vendor-reported][secondary].
- **AWS ElastiCache for Valkey** is ~**20% cheaper** hourly than ElastiCache
  for Redis; Valkey is the default engine on ElastiCache and MemoryDB in
  2026 [vendor-reported][secondary].
- Adoption signal: **>100M Docker Hub pulls by May 2026**, ~17x YoY, among
  the fastest-growing base images in the caching category [secondary].
- Compatibility: identical RESP2/RESP3 wire protocol; RDB/AOF interchangeable
  at the 7.2 baseline — but **RDB files written by Redis CE 7.4+ will not
  load into Valkey** (forward-compatibility risk) [secondary].
- Divergence: post-fork features differ — Redis 8 bundles JSON/search/
  vectors into core; Valkey ships equivalents as **separate modules**
  [secondary].

### 8.3 Redis 8 — 2026 status

- Current line **8.x** [secondary]; headline features: vector sets, hash
  field TTL, Redis Functions 2.0, core-integrated JSON/time-series/search
  [independent].
- Redis Inc. counters Valkey on features (richer core) while Valkey counters
  on license and engine efficiency [independent].

