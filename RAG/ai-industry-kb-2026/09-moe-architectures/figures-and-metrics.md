---
id: ai-industry-kb-2026/09-moe-architectures/figures-and-metrics
title: "Figures and metrics"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "CISA", "DeepSeek", "MiniMax", "Moonshot", "Z.ai"]
dates: ["2025-08-05", "2025-12-16", "2026-04-16", "2026-04-24", "2026-06-01", "2026-07-27", "2026-08-03", "2026-08-25", "2026-08-28"]
keywords: ["apache", "attention", "benchmark", "compute", "cost", "deepseek", "glm", "gqa", "kimi", "license", "memory", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4757, 4807]
section: "9. MoE Architectures"
sha256: c58b9d848173f8ab2b11d1a1819a9ace24aa6e237ffcf128cb3777f50a6366de
---

# Figures and metrics

## Figures and metrics

### Full specification matrix (2026 open-weight MoE releases; benchmark scoreboard in part 09b)

Total params / active per token / layers / experts (routed active + shared) / hidden dim / attention design / context / modalities / license / release:

- MiniMax M3: 428B / 23B / 60L / 128E (4+1 shared) / 6144 / MSA block-sparse + GQA(64Q/4KV) / 1M / text+image+video / MiniMax Community License / 2026-06-01
- Qwen3.6-35B-A3B: 35B / 3B / 40L / 256E (8+1 shared) / 2048 / 3:1 Gated DeltaNet:gated attention hybrid / 262K native (1M w/ YaRN) / text+image / Apache 2.0 / 2026-04-16
- Gemma 4 26B-A4B: 25.2B / 3.8B / 30L / 128E (8+1 shared) / 2560 / SWA(1024)+global hybrid, pure attention / 256K / text+image / Apache 2.0 / ~2026-04
- MiMo-V2.6-Flash: 309B / 15B / 48L / 256E (8+0 shared) / 4096 / 5:1 SWA(128)/global, learnable sink bias / 1M / text+image+video+audio / MIT / ~2026-09
- MiMo-V2.6-Pro: 1.02T / 42B / undisclosed / undisclosed / undisclosed / same family / 1M / omnimodal / MIT / ~2026-09
- GLM-5.3-Flash: 320B / 18B / 45L / 288E (8+0 shared) / 4096 / hybrid sparse+linear attention + mHC / 1M / text+image+video / MIT / 2026-08-25/26
- GLM-5.3 (flagship): 753B / undisclosed active / undisclosed / undisclosed / undisclosed / MoE base from 5.2 / 1M / text-only / custom glm-5.3 license / 2026-08-28
- Kimi K3: 2.8T / 104B / 93L (1 dense) / 896E (16+2 shared) / undisclosed (latent 3584 / full 7168) / 69 KDA + 24 Gated MLA / 1M / text+image (video via API) / open-weight / 2026-07-27
- DeepSeek-V4-Pro: 1.6T / 49B / 61L [PARTIALLY VERIFIED] / 384E (6+1 shared) [PARTIALLY VERIFIED] / undisclosed / NSA lineage / 1M (384K max output) / text-only / MIT (later sources) / 2026-04-24
- DeepSeek-V4-Flash: 284B / 13B / 43L [PARTIALLY VERIFIED] / 256E (6+1 shared) [PARTIALLY VERIFIED] / undisclosed / NSA + DSA + token-wise compression / 1M / text-only / MIT (later sources) / 2026-04-24
- Qwen3.8-Max: 2.4T / 95B / 92L / 512E (10+1 shared) / 8192 / Gated DeltaNet+gated attention hybrid / 1M / text+image+video / non-Apache open license / 2026-08-03
- gpt-oss-120b: 117B / 5.1B / 36L / 128E (4+0 shared) / undisclosed / standard attention / 128K / text-only / Apache 2.0 / 2025-08-05
- gpt-oss-20b: 21B / 3.6B / 24L / 32E (4+0 shared) / undisclosed / standard attention / 128K / text-only / Apache 2.0 / 2025-08-05
- MiMo-V2-Flash (predecessor): 309B / 15B / 48L / 256E (8+0 shared) / 4096 / 5:1 SWA(128)/global / 256K / text-only / MIT / 2025-12-16

### Routing geometry: expert counts, top-k, sparsity ratios

- Expert-pool trajectory: 32 (gpt-oss-20b, 2025) → 128 (gpt-oss-120b, M3, Gemma 4) → 256 (Qwen3.6, MiMo) → 288 (GLM-5.3-Flash) → 384 (V4-Pro, reconstructed) → 512 (Qwen3.8-Max) → 896 (Kimi K3). The 2026 field converged on fine-grained pools of 128–896 small experts.
- Top-k trajectory: 4 (gpt-oss, M3) → 6 (V4, reconstructed) → 8 (Qwen3.6, Gemma 4, GLM-5.3-Flash, MiMo-V2.6) → 10 (Qwen3.8-Max) → 16 (Kimi K3). Per-token active counts rise from 3–5B (2025–early-2026 small tier) to 95–104B (2026 flagships).
- Sparsity ratios (total/active): gpt-oss-120b ~23:1; Qwen3.6 ~12:1; Gemma 4 ~6.6:1; MiniMax M3 ~19:1; MiMo-V2.6-Flash ~21:1; GLM-5.3-Flash ~18:1; Qwen3.8-Max ~25:1; Kimi K3 ~27:1; DeepSeek-V4-Pro ~33:1 (the sparsest flagship); V4-Flash ~22:1.
- Shared experts: 1 shared is the 2026 norm (M3, Qwen3.6, Gemma 4, Qwen3.8-Max, V4); Kimi K3 uses 2 shared; the no-shared-expert outliers are gpt-oss (2025 template), MiMo-V2/V2.6, and GLM-5.3-Flash.
- Routing capacity math: fine-grained segmentation (expert intermediate dims 512–1536 in 2026 models) moves achievable expert combinations from C(16,2)=120 (coarse) to C(64,8)≈4.4 billion (fine) — the combinatorial explosion that justifies the routing side of the bargain.
- Communication math (training): expert-parallel dispatch requires two any-to-any phases per MoE layer (dispatch tokens to expert ranks, combine outputs back); per-layer all-to-all volume ∝ T·k·D per layer, reaching TB-scale per layer at 64-way EP for V3-class models — the dominant training-side cost alongside optimizer memory (full training-economics treatment in part 09b).

### The routing ladder: how expert pools and top-k moved together

| Model | Routed experts | Active routed (top-k) | Shared | Effective combos (illustrative) |
|---|---|---|---|---|
| gpt-oss-20b (2025) | 32 | 4 | 0 | C(32,4) ≈ 36K |
| gpt-oss-120b (2025) | 128 | 4 | 0 | C(128,4) ≈ 10.7M |
| Gemma 4 26B-A4B (~2026-04) | 128 | 8 | 1 | C(128,8) ≈ 1.4×10^11 |
| MiniMax M3 (2026-06) | 128 | 4 | 1 | C(128,4) ≈ 10.7M |
| Qwen3.6-35B-A3B (2026-04) | 256 | 8 | 1 | C(256,8) ≈ 4.2×10^13 |
| DeepSeek-V4-Flash (2026-04) | 256 [PARTIALLY VERIFIED] | 6 | 1 | C(256,6) ≈ 3.6×10^11 |
| MiMo-V2.6-Flash (~2026-09) | 256 | 8 | 0 | C(256,8) ≈ 4.2×10^13 |
| GLM-5.3-Flash (2026-08) | 288 | 8 | 0 | C(288,8) ≈ 1.2×10^14 |
| DeepSeek-V4-Pro (2026-04) | 384 [PARTIALLY VERIFIED] | 6 | 1 | C(384,6) ≈ 4.4×10^12 |
| Qwen3.8-Max (2026-08) | 512 | 10 | 1 | C(512,10) ≈ 2.6×10^19 |
| Kimi K3 (2026-07) | 896 | 16 | 2 | C(896,16) ≈ 10^34-class |

- The illustrative combination counts are exact-combinatorial, not achievable-routing claims: the reference comparison from the literature is C(64,8) ≈ 4.4 billion vs C(16,2) = 120, showing how fine granularity multiplies routing expressivity per unit of per-token compute.
- Two independent axes moved in 2026: pool size (32 → 896, the routing-capacity axis) and top-k (4 → 16, the per-token-compute axis). Kimi K3 pushed both simultaneously; DeepSeek-V4 pushed pool size while holding top-k at 6 (the lowest active-expert count among 2026 flagships).

### Routing mechanisms: the 2026 reference sheet (top-k → latent routing)

