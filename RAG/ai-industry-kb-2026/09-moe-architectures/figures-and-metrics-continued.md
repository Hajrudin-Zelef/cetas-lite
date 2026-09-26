---
id: ai-industry-kb-2026/09-moe-architectures/figures-and-metrics-continued
title: "Figures and metrics (continued)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "Unsloth", "Xiaomi", "Z.ai"]
dates: ["2025-12-16", "2026-04", "2026-04-24", "2026-05", "2026-05-06", "2026-07", "2026-07-28", "2026-08-16", "2026-08-22", "2026-09", "2026-09-02", "2026-09-07", "2026-09-09", "2026-09-22"]
keywords: ["agentic", "amd", "apache", "attention", "awq", "benchmark", "benchmarks", "claude", "compute", "deepseek", "embeddings", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5230, 5259]
section: "9. MoE Architectures"
sha256: 03066fff8869c4ea78e90d6b9b62f70c02e2e874c0bd86679e3678a980b58d04
---

# Figures and metrics (continued)

- **Kimi K3 active params:** 104B (technical-report-based sources, majority) vs ~50B (one Medium analysis, 2026-07-28). The 16-of-896-experts figure is consistent across sources; the implied per-expert size differs. Treated as 104B (authoritative-leaning) with the dispute noted.
- **DeepSeek-V4 license:** MIT (later mid-2026 coverage) vs Apache 2.0 (April 2026 launch coverage). First-party model card not directly opened in this research pass — verify before citing.
- **GLM-5.3-Flash context:** config.json `max_position_embeddings` = 1,048,576 and OpenRouter listing up to 1,310,720 vs LumaDock model-card eval citing 300K context. The 1M figure is config-grounded; the 300K figure's basis is unclear.
- **MiniMax M3 "matches Claude Sonnet 4.6 on real-world agentic benchmarks":** sourced to Morph's vendor-adjacent model page; no independent benchmark reproduction found. Marked UNVERIFIED.
- **MiMo-V2.6 release details:** come from a single detailed third-party release analysis (OrcaRouter, Sept 2026); Xiaomi first-party announcement not directly opened. The 309B/15B/1M/MIT figures are consistent with the V2-Flash lineage but single-sourced — verify against first-party channels before citing as definitive.
- **MiniMax M3 "7 MTP modules":** appears in community inference notes (MTP/NEXTN weights absent from shipped checkpoints); not confirmed in the official model card. Treat as community-reported.
- **Unsloth "12x faster MoE training"** (Feb 2026), **"MoE training 3–5x faster"** (July 2026 v0.1.481-beta), **"up to 2x faster" MTP-by-default for Qwen3.8-Flash/GLM-5.3-Flash** (Sept 2026 v0.1.805/806-beta), **GLM-5.2 Dynamic GGUF retention figures** (~82%/~76% at 2-bit/~1-bit, ~98% at 4-bit; LinkedIn secondary, vendor-originated): all [VENDOR-REPORTED], no independent reproduction found as of 2026-09-22.
- **Unsloth Studio beta release dates (v0.1.805–808):** presented as "first half of September 2026" from listing recency; the vendor does not stamp exact dates on these beta pages — [UNVERIFIED].
- **PyTorch 2.14 (released 2026-09-02)** silently changed clamp/min/max boundary subgradients (1→0): a genuine training-reproducibility hazard for anyone upgrading mid-experiment. [DIRECTIONAL — independent dev-blog source, not a first-party changelog line.]

## Figures and metrics (continued)

- **Training FLOPs anchor (DeepSeekMoE published ratios):** 16B MoE ≈ 7B dense at ~40% compute; 145B MoE ≈ 67B dense at 28.5% compute — ≈2.5–3.5× compute savings at matched capability.
- **DeepSeek-V3 training footprint:** 16-way pipeline × 64-way expert × ZeRO-1 data parallel; 2048 H800 GPUs; DualPipe; DeepEP all-to-all; FP8 mixed precision. EP dispatch bandwidth ∝ T·k·D per layer — TB-scale per layer at 64-way EP, V3-class.
- **Pretraining token budgets (dated):** MiMo-V2-Flash 27T tokens (2025-12-16); GLM-5.3-Flash 30T multimodal tokens (2026-08); DeepSeek-V4-Pro 33T / V4-Flash 32T (2026-04-24).
- **Aux-loss degradation:** ~23% cited in 2026 analyses of interference gradients from the classic auxiliary loss L = α·N·Σ f_i·P_i — the production field moved to dynamic bias / expert-choice / quantile balancing.
- **Unsloth×NVIDIA component measurements (2026-05-06):** packed-sequence metadata caching +43.3% forward, +5.8% backward, +14.3% per batch (Qwen3-14B QLoRA SFT); double-buffered checkpoint reload +8.4% (8B) / +6.7% (14B) / +4.6% (32B); GPT-OSS bincount routing +23% forward / +13% backward in the targeted path. Vendor headline timeline: "12x faster MoE" (Feb 2026, B200-specific) → "MoE training 3–5x faster" (July 2026) → component numbers (May 2026) as the auditable evidence.
- **Active-compute throughput illustration:** Qwen3.6-35B-A3B (3B active) 61 tok/s vs 27B dense 7 tok/s on the same RTX 5070 (~9×). DeepSeek-V3 671B/37B reads ~5% of weights per token (dense-30B-class throughput on H800).
- **Active/total residency gap:** OLMoE-1B-7B vs OLMo-1B — active params agree within 8.9%, resident memory differs 5.9×; ratio invariant 5.40× at fp32 and int4. Switch-Base-8: all 8 experts receive tokens within 223 tokens (worst-layer coverage 100%).
- **KV-cache reduction factors (attention-side, 2026):** MLA ~5% of LLaMA-3-70B (DeepSeek-V2); MiMo 5:1 SWA/GA ~6× reduction; GLM-5.3-Flash hybrid sparse+linear 4.4× claimed; HybridKV compression up to 7.9×; TurboQuant 3-bit KV ÷6 memory.
- **MoE speculative decoding:** MiMo-V2-Flash accepted length 2.8–3.6 tokens/forward → 2.0–2.6× speedup; MiMo-V2.6-Flash 5-layer DFlash drafter, 7 tokens ahead; MTP-by-default for Qwen3.8-Flash/GLM-5.3-Flash claims up to 2x generation [VENDOR]. ROCm asymmetry: MTP near-useless on MoE at some configs while dense 27B went 45→62 t/s.
- **EP topology guidance (AMD ROCm):** ≤128 concurrent requests → TP=8 +40–86% throughput; ≥512 → DP=8+EP +16–47% (7,114 TPS DeepSeek-R1 at 1024 concurrency); crossover ~256–512; ultra-sparse (<1% activation) → EP 7–12% SLOWER; standard MoE (≥3% density) → EP helps.
- **DeepSpeed comm-opt:** MoE all-reduce penalties 3–12 s in problematic large-scale layouts; multi-rank bucketing + rank placement as fix.
- **$/MTok self-hosted ballpark (mid-2026):** Qwen3-MoE 235B-A22B FP4 on 8× B200 ≈ $0.30/MTok at ~1200 tok/s/GPU; DeepSeek V3.1 FP4+MTP on NVL72 ≈ $0.25/MTok; Llama 3.3 70B FP8 on 4× H100 ≈ $0.40/MTok. AA Aug 2026: Nemotron 3.5 Lightning $0.22/M output at 235–494 t/s vs Gemma 4 31B dense $0.40/M at 37–222 t/s (capability 24 vs 30).
- **API price ladder (2026-09, $/M input / cached / output):** DeepSeek-V4-Pro $0.435/$0.003625/$0.87 (peak tiers $1.32/$0.66 from 2026-08-16); V4-Flash $0.44/$0.22 peak/off-peak input; GLM-5.3-Flash $0.15/$0.03/$0.50 (promo $0.075/$0.015/$0.25 through 2026-09-09); GLM-5.3 flagship $1.40/$4.40; MiniMax M3 $0.30; MiMo-V2-Flash $0.10/$0.30; Qwen3.8-Max $2.00/$0.25 cache/$6.00; Kimi K3 $3.00/$0.30/$15.00.
- **Quantization footprint (2026 MoE fleet):** GLM-5.3-Flash FP8 ~306–331 GiB; MiMo-V2.6-Flash FP8 172.9 GB (65 shards); MiniMax M3 NVFP4 ~245 GB → ~61 GB/GPU at TP4; M3 MXFP8 ~440 GB; gpt-oss MXFP4 native 60.8 GiB (120b) / 12.8 GiB (20b); Kimi K3 per-expert 33.0M params, MXFP4 E8M0 scale per 32 weights; Mixtral 8x7B AWQ 4-bit 90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT; Qwen3.8-Max FP8 repo 21,400 downloads vs 17,386 bf16 (2026-08-22).
- **Fine-granularity scale ladder (2026):** expert counts 128 (M3, Gemma 4) → 256 (Qwen3.6, MiMo-V2.6, V4-Flash) → 288 (GLM-5.3-Flash) → 384 (V4-Pro) → 512 (Qwen3.8-Max) → 896 (Kimi K3); top-k 4–16; combinatorial capacity C(64,8) ≈ 4.4B vs C(16,2) = 120. Kimi K3 latent routing: hidden 7,168 → latent 3,584, ~halving routed traffic/compute; sparsity ratio ~27:1 (K3), ~25:1 (Qwen3.8-Max), ~12:1 (Qwen3.6), ~6.6:1 (Gemma 4).
- **Scoreboard anchors (with dates):** SWE-bench Verified — M3 80.5% (AA, 2026-09-07), Qwen3.6 73.4% (2026-04), MiMo-V2-Flash 73.4% (2025-12), Gemma 4 17.4% (community); LiveCodeBench — V4-Pro-Max 93.5%, V4-Flash 91.6%, Qwen3.6 80.4% v6, Codeforces 3,052; agentic — V4-Pro 67% vs Sonnet 4.5 47% vs Opus 4.6 Thinking 80%; GPQA Diamond — K3 93.5%, Qwen3.8-Max 92.6%, Qwen3.6 84.1%; MCPMark — Qwen3.6 37.0 vs Gemma 4-31B 18.1; AA Index — K3 57.1 (#4 globally), GLM-5.3 60 vs 5.3-Flash 57 vs 5.2 53.
- **Modest-hardware datums:** Gemma 4 offload hot set ~4.5 GB, ~24 MB/forward from NVMe, 8 GB RTX 3070 viable; Qwen3.6 UD-Q4_K_XL ~21 GB on 24 GB RTX 4090; MacBook Pro 20.9 GB Q4; gpt-oss-20b 12.8 GB [PARTIALLY VERIFIED]; Axolotl `quantize_moe_experts` GLM-4.7-Flash ~127 GiB → ~23 GiB (secondary); Unsloth `save_pretrained_gguf` OOM guard default 0.75; MLX quantized KV caches "up to 74% less prompt memory" [VENDOR].

