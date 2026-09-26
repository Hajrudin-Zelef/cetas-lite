---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/figures-and-metrics
title: "Figures and metrics"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "DeepSeek", "Google", "MiniMax", "Nvidia", "Oracle", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-05", "2026-07-24", "2026-08-13", "2026-09-08"]
keywords: ["amd", "attention", "benchmark", "benchmarks", "compute", "cost", "decode", "deepseek", "fine-tuning", "fp8", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3075, 3130]
section: "7. KV Cache & Long-Context Techniques"
sha256: 25487e0398918776d2d9f35ed3a9adb6a4356646b6d8cc66b75c04cae6fac796
---

# Figures and metrics

- **2026 (vLLM v0.26.0) — DeepSeek-V4 performance push + KV offloading maturation:** the 411-commit / 212-contributor release paired specialized V4 routing and sparse decode/prefill optimizations with object-store secondary KV tiers and DP-replica-aware tiering. Full release-notes detail is part 06's territory; the relevance here is that V4-era MLA serving standardized on fp8_e4m3 KV with MLA attention backends (see §"Quantized KV in production").
- **Apr 2026 — SGLang day-zero support for DeepSeek-V4.** Native V4 serving (CSA/HCA lineage) at release — evidence of how fast MLA-derived sparse architectures reach the two live engines.
- **~Mar 2026 — SGLang FP8 KV for MHA via the aiter backend (AMD).** The quantized-KV story is not NVIDIA-only; the AMD path matured in parallel in 2026.
- **2026-08-13 — DeepSeek V4-Pro GA checkpoint (V4-Pro-0813) reported** [UNVERIFIED, secondary-sourced]; legacy `deepseek-chat`/`deepseek-reasoner` aliases retired 2026-07-24 (routed to V4-Flash during grace period). Keep as secondary-sourced; do not use for pricing-critical claims.

## Figures and metrics

### Attention-variant figures (cache size and quality)

| Variant | KV stored per token | Relative cache (vs MHA) | Quality note |
|---|---|---|---|
| MHA (multi-head attention) | 2 × n_heads × d_head | 1× (baseline) | Full expressivity; impractical for long context (Llama 2 70B @ 8K ctx: >20.5 GB) |
| MQA (multi-query) | 2 × 1 × d_head | ~1/n_heads | Aggressive; quality cost at scale |
| GQA (grouped-query) | 2 × n_groups × d_head | ~n_groups/n_heads | Industry default; ~88% reduction vs MHA, retains >95% of MHA quality |
| MLA (multi-head latent) | d_c (latent, e.g. 512) + shared RoPE key (64) | ~6–14% of MHA | Feature-axis compression, keeps all heads; quality matches or slightly beats MHA |

- **MLA 93.3% KV-cache reduction** vs standard MHA (DeepSeek-V2 paper, arXiv:2405.04434, May 2024) — the primary measured figure; 5.76× inference throughput on the same baseline. Controlled MHA-vs-MLA ablation: ~96% (34.6K vs 860.2K cached elements/token). Raw per-layer element ratio: ~98.2% (576 vs 32,768). Toy-model arithmetic (temperature2.com, 2026-09-08): **14.2× smaller cache than MHA at full quality** [DIRECTIONAL — toy model].
- **"MLA exactly 4× batch/GPU" — only partially verified.** No primary measurement in the source corpus pins "exactly 4×"; the figure circulates as shorthand for the 93.3%/5.76× family of numbers. Use the measured figures (93.3%, 5.76×, 14.2× toy); treat any "exactly 4×" claim as [UNVERIFIED].
- **MLA has higher expressive power than GQA under the same KV-cache budget** (2026 survey literature) — the theoretical basis for the empirical wins.
- **Serving MLA is not free:** the cache stores *latents* plus a decoupled RoPE key rather than standard K/V blocks, so engines need MLA-aware kernels (the absorption math differs from the GQA path). vLLM and SGLang both added MLA support after DeepSeek-V3; operators validate it as a distinct code path.
- **GQA ~88% reduction vs MHA, >95% quality retention** — the 2026 consensus; its measurable modeling weakness vs MHA is what pushed DeepSeek to MLA and others to hybrid layouts.

### Post-hoc compression figures

| Method | Family | Measured figure | Quality retention |
|---|---|---|---|
| HybridKV (VERIFIED) | head-heterogeneous hybrid | up to 7.9× KV memory (Qwen2.5-VL-7B, 11 MMLM benchmarks), 1.52× decode | almost no drop; some benchmarks above full-cache |
| TurboQuant (VERIFIED, Google) | online quantization (PolarQuant + QJL) | ≥6× vs FP16; up to 8× attention-logit speedup (H100); within ~2.7× of info-theoretic limit | 99.5% attention fidelity; lossless NIH |
| H2O — Heavy Hitter Oracle (NeurIPS 2023) | token eviction (decode-phase) | up to 29× throughput vs HF Accelerate (OPT-6.7B/30B) at 20% heavy hitters | power-law attention weights; decode-phase only |
| StreamingLLM (attention sinks) | token eviction | keeps first-few tokens (attention sinks) + sliding window | no importance scoring; middle-context tokens can be dropped |
| SnapKV (prefill-phase) | token eviction | end-of-prompt observation window votes per-head for important KV positions | more accurate than H2O at same budget; LongBench prefill baseline |
| PyramidKV / PyramidInfer | layer-wise eviction budgets | 2.2× throughput vs HF Accelerate; >54% KV-cache memory reduction | early layers need richer context; deeper layers converge |
| KIVI (ICML 2024) | 2-bit quantization (keys per-channel, values per-token) | 2.6× combined peak-memory reduction (weights+KV); 4× larger batches; 2.35–3.47× throughput | no fine-tuning; the baseline TurboQuant beats on LongBench |
| KVQuant (calibrated mixed-precision) | sub-4-bit quantization | pre-RoPE key quantization; dense-and-sparse outlier decomposition; evaluated to 10M-token contexts | beats fixed grids; calibration cost amortizes on stable workloads |
| FP8 KV | production dtype | 2× vs BF16; 54% ITL slope (H100); break-even ~7K tokens | sub-0.3% accuracy cost; the "free" first step |
| Palu / LoRC / SVDq / CSKV / ReCalKV | low-rank / latent | post-training low-rank K/V projection; group-head decomposition (Palu); Fisher-information rank search | orthogonal to quantization and eviction; stackable |
| "Thin Keys, Full Values" (retrofit) | low-rank SVD | 75% key-cache savings at ~2% quality cost (<1% pretraining data); up to 16× combined with GQA+quantization | ~25 GB saved per user at 128K on a 7B; ~60% more concurrent users [VENDOR] |
| Attention Matching [paper claim] | latent-space compaction | claimed 50× compaction | early-stage; treat as paper claim, not production |
| TriAttention [paper claim] | reasoning-aware compression | claimed 10.7× memory reduction on AIME25 at matched accuracy | early-stage; compresses the reasoning trace's cache footprint |

### Reasoning-aware compression figures

| Method | Measured figure | Quality retention |
|---|---|---|
| ThinKV (ICLR 2026 Oral; thought-adaptive fp8/int4/ternary-2bit + eviction, in-place slot reuse) | <5% of original KV cache; up to 5.8× throughput over SOTA baselines | near-lossless; AIME within 4% at 3.67% memory; precision auto-adapts 3.4–3.8 bits with difficulty |
| R-KV v4 (Jan 2026; redundancy-aware) | ~100% of full-KV performance at 10% cache; 90% memory saving | 105% at 16% cache; 6.6× throughput over standard CoT inference |
| FoveatedKV (Apr 2026; top 10% fp16, bottom 90% fp8 E4M3 K + INT4 V) | 75% KV cut | ≤3% quality loss [COMMUNITY benchmark] |

### Sparse-attention figures (compute, not necessarily cache)

- **DeepSeek Sparse Attention (DSA, V3.2, Dec 2025):** learned "lightning indexer" selects top-K subset (2,048 tokens) for sparse attention. **DeepSeek V4 (Apr 2026):** hybrid Compressed Sparse Attention + Heavily Compressed Attention — V4-Pro at 1M tokens uses **27% of V3.2's inference FLOPs and 10% of its KV-cache size**; V4-Flash pushes to **10% FLOPs / 7% KV cache** [VENDOR]. SGLang implementation (2026): FP8 Q8KV8 sparse-MLA prefill (PR #30514). **DSA reduces attention compute; it does not intrinsically shrink the KV cache** (any historical token may be selected later; the cache must retain full-fidelity KVs unless separate compression like CSA/HCA applies).
- **MiniMax Sparse Attention (M3, Jun 2026):** >9× prefill / >15× decode at 1M tokens [VENDOR]; 28.4× reduction in per-token attention compute vs full attention at 1M; per-token compute ≈ 1/20th of previous generation; operates on uncompressed KV (NIH-fidelity argument); 512K guaranteed usable floor of the 1M window.
- **GLM-5.2 IndexShare (Jun 2026):** single sparse-attention indexer shared across every four sparse-attention layers — **2.9× reduction in per-token FLOPs at 1M context**, up to **20% improved MTP speculative-decoding acceptance length** [VENDOR]; 1M input / 65K output context.

### Hybrid O(1)-state figures

