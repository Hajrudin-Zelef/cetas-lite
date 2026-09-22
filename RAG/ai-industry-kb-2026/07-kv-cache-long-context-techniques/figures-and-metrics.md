---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/figures-and-metrics
title: "Figures and metrics"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Google", "Meta", "MiniMax", "Oracle", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-05", "2026-09-08"]
keywords: ["attention", "benchmark", "benchmarks", "compute", "consumer", "cost", "decode", "deepseek", "fine-tuning", "fp8", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3080, 3151]
section: "7. KV Cache & Long-Context Techniques"
sha256: 0ee45767f1661534c9e1f8df11dc5337d3d54fe5e196c9fd7160572c5012dffc
---

# Figures and metrics

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

- Jamba-1.5 (1:7 attention:Mamba): **4 GB (Mini) / 9 GB (Large) KV at 256K** vs 32 GB Mixtral / 80 GB Llama-3.1-70B (~10× lower); 256K effective on RULER.
- Worked 64-layer hybrid math: KV ~67 GB → ~8.3 GB (8 attention layers retained); compressor state 14.7 MB/user (Mamba) or ~117 MB/user (linear-attention fast-weight grids); total ≈ 39.3 GB on one H100 vs two GPUs for the pure transformer.
- Zamba2-VL: ~10× lower TTFT vs closest transformer baseline on 32K prefill (fixed-size recurrent state).
- Nemotron community benchmark (2× RTX 3060): 250K context on consumer GPUs; TG 34 tok/s @250K vs 18 tok/s Qwen3.5-35B (**89% faster**); crossover beyond 64K.
- Qwen3-Next GDN recurrent state: ~0.02 GiB (Sep 2026 FP8-KV analysis).
- Terminology correction: O(n²) = attention *compute* (prefill); KV *memory* = O(n); recurrent layers = O(1) state per layer.

### The 2026 quantization ladder (one user, Llama 3 70B-class @ 128K context)

- BF16 **~42 GB** → FP8 **~21 GB** (2×, "free") → INT4 **~10.5 GB** (4×) → TurboQuant 3-bit **~8 GB** (5.3×).
- SNIA SDC25 capacity anchors: Qwen3-8B **150 KB/token** FP16 (147 GB at 1M); LLaMA-3.3-70B **330 KB/token** (327 GB at 1M); LLaMA-3.1-405B **516 KB/token** (516 GB at 1M).
- 70B-class worked sizing table (batch 1):

| Context | FP16/BF16 KV | + FP8 KV | + TurboQuant 3-bit | + MLA (arch.) |
|---|---|---|---|---|
| 4K | ~1.3 GB | ~0.7 GB | ~0.2 GB | ~0.09 GB |
| 128K | ~42 GB | ~21 GB | ~7 GB | ~2.8 GB |
| 1M | ~336 GB | ~168 GB | ~56 GB | ~22 GB |

(Base: 2 × 80 × 8 × 128 × seq_len × 2 bytes. MLA ≈ 6.7% of MHA-cache per the 93.3% figure.) At 1M tokens no single lever fits one 80GB GPU — the 2026 answer is stacking (GQA + FP8 + prefix reuse + disaggregation → sibling 07b) or architectural choices made at training time (MLA/MSA/shared-KV). TurboQuant-class 3-bit is what moves 128K-class workloads onto consumer hardware (~32K → ~180K usable context on 2× RTX 5060 Ti / M-series Macs per community reports).

