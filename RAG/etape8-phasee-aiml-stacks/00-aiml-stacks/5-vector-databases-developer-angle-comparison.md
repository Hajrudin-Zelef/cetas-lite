---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/5-vector-databases-developer-angle-comparison
title: "5. Vector databases — developer-angle comparison"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: reference
actors: []
dates: ["2026-03"]
keywords: ["attention", "benchmark", "benchmarks", "disclosure", "parameters", "quantization"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [356, 403]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: 1149e7d7f99e53025c94244b1087075f59ef61310d1941b37a3bef510c6f64d2
---

# 5. Vector databases — developer-angle comparison

## 5. Vector databases — developer-angle comparison

Brief comparison only; `pgvector` operational detail lives in Step 7F.

### 5.1 Qdrant

- Secondary reports place Qdrant around **1.17 (March 2026)**, with later
  unverified references to 1.19 — treat the series as [secondary].
- Reported 2026 themes: composable/hybrid search, relevance feedback,
  multitenancy, quantization, and edge deployments [secondary].
- Written in Rust; ships as a single binary plus Docker; gRPC/REST APIs
  [secondary].

### 5.2 Weaviate

- Weaviate is the GraphQL-first, modular vector database with a
  module/plugin architecture for vectorizers and rerankers [secondary].
- No verified 2026 version captured in this pass — gap G-06 [unverified].

### 5.3 Milvus / Zilliz

- Milvus is the distributed, cloud-native vector database (LF AI &
  Data graduated project) aimed at billion-scale deployments; Zilliz is
  the managed commercial offering [secondary].
- No verified 2026 version captured in this pass — gap G-07 [unverified].

### 5.4 pgvector (summary; detail in Step 7F)

- `pgvector` is the PostgreSQL extension adding the `vector` type and
  HNSW/IVFFlat indexes, letting teams keep vectors next to relational
  data [secondary].
- Operational detail (index build parameters, distance operators,
  upgrade paths) is documented in Step 7F and not repeated here.

### 5.5 Developer decision matrix

| System | Deployment shape | Strength | Trade-off |
|---|---|---|---|
| Qdrant | Single binary / cluster | Rust speed, hybrid search, edge | Smaller ecosystem than Milvus [secondary] |
| Weaviate | Container / cloud | GraphQL, modular vectorizers | Heavier operational surface [secondary] |
| Milvus | Distributed cluster | Billion-scale, cloud-native | Operationally heavy at small scale [secondary] |
| pgvector | Postgres extension | Zero new infra, transactional | Not a dedicated ANN engine at extreme scale [secondary] |

- Market-size, download-count, and head-to-head benchmark claims for this
  segment are methodology-sensitive and were not independently verified;
  treat vendor benchmarks as [vendor-reported] and third-party
  comparisons as [secondary] with attention to dataset/hardware disclosure.

