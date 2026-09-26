---
id: ai-industry-kb-2026/06-inference-engines/claim-status-register-this-part-s-track
title: "Claim status register (this part's track)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Huawei", "LongCat", "SGLang"]
dates: ["2026-02", "2026-03-15", "2026-05-05", "2026-08-20", "2026-09-04", "2026-09-20"]
keywords: ["alignment", "amd", "ascend", "fp4", "gpu", "gpus", "omni", "sglang", "throughput"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2862, 2881]
section: "6. Inference Engines"
sha256: 2cf0b469aaf09bf6dd4ba0083c3028654040a918e4c850f39f7d5c96d22f00b5
---

# Claim status register (this part's track)

| Version | Date | Headliners |
|---|---|---|
| v0.4 blog | 2026 (blog) | Zero-overhead scheduler (1.1x, ~0 GPU idle); cache-aware load balancer (1.9x throughput, 3.8x hit rate) [VENDOR] |
| v0.5.6 | Dec 2025 | FP4 MLA KV caches |
| v0.5.18 | 2026-08-20 | Tag current in install docs for source builds |
| v0.5.19 | 2026-09-04 | 786 PRs / 214 contributors; Qwen3.8, dots3.note, Ling-3.0, Spark2.5, MiniCPM-SALA, Granite 4.2, LongCat-Image-Edit; cross-hardware cookbooks (Ascend A3, MI355X, DGX Spark) |
| sglang-omni v0.1.4 | 2026-09 | Any-to-any track; AMD ROCm Qwen3-TTS/ASR on gfx950; SGLang 0.5.18 alignment |

### Claim status register (this part's track)

| # | Claim | Status | Evidence basis |
|---|---|---|---|
| C1 | RadixAttention "2.5x cache hit rate vs competition" | [UNVERIFIED] | No source states this ratio; nearest real figure is a 2.5x *throughput* gain under strict JSON constraints (2026-09-20) |
| C2 | SGLang on 400,000+ GPUs | [UNVERIFIED] | Repeated verbatim across secondary roundups; no primary citation in either wave; repeated again in RadixArk coverage |
| C3 | MLA → "4x batch size per GPU" | [PARTIALLY VERIFIED] | Paper figures (93.3% KV reduction, 5.76x throughput) + native paths support direction; no controlled reproduction |
| C4 | "25x" on GB300 NVL72 | [VENDOR] | Real February 2026 headline; baseline unaudited |
| C5 | 4.7x multi-turn speedup, 99.8% JSON validity | [DIRECTIONAL] | Single secondary (privocto, 2026-03-15) |
| C6 | XGrammar "80x vs older approaches" | [VENDOR] | Vendor-adjacent; independently attested ~3x compressed-FSM |
| C7 | RadixArk $100M at $400M (2026-05-05) | Verified | Business Wire launch terms; January $400M was pre-launch reporting |

