---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/the-error-correction-canon-the-documented-correction-machine
title: "The error-correction canon (the documented \"correction\" machinery)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Falcon", "Huawei", "Hugging Face", "Intel", "MiniMax", "Mistral", "Moonshot", "Oracle", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2024-05", "2025-05", "2026-03-21", "2026-04", "2026-06", "2026-08-14", "2026-08-26", "2026-08-29", "2026-09-08"]
keywords: ["ascend", "attention", "benchmark", "benchmarks", "compute", "consumer", "cost", "decode", "deepseek", "distribution", "embeddings", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3152, 3229]
section: "7. KV Cache & Long-Context Techniques"
sha256: 7aeda7bf24a9701aad4536368d82af801dce1da3807b8ff589f0ec90fd56dbfd
---

# The error-correction canon (the documented "correction" machinery)

### The error-correction canon (the documented "correction" machinery)

The brief's phrase "dynamic drift correction" names **no single documented method** [UNVERIFIED as phrased]. What the 2026 literature documents instead:

- **GEAR** (Kang et al., 2024; arXiv:2403.05527) — X ≈ D̂ + L + S: quantized backbone D̂ (~98% of entries, ultra-low precision) + low-rank matrix L (per-head truncated SVD of the *coherent* residual) + sparse matrix S (outliers). Lite variant GEAR-L (low-rank only); streaming-buffer strategy for long generation. Design insight surviving into 2026: quantization error has coherent structure (rapidly decaying residual spectrum) plus incoherent outliers — correct each with different tools.
- **MiKV** (Yang et al., 2024) — **Dynamic Outlier Awareness**: a channel balancer multiplied into keys and divided out of queries *dynamically* at serving time; the closest literature match to "correction dynamique" as a *runtime* mechanism rather than a calibration artifact. Evicted KV pairs retained in reduced precision rather than dropped.
- **ResQ** — PCA identifies high-variance components (8-bit) while the rest go 4-bit, with random rotations to suppress outliers.
- **WKVQuant** — 2D quantization strategy aligning KV-cache values into a smoother, more uniform range before quantizing.
- **ZipCache / PrefixQuant / MiniKV** — importance-quantized variants: ZipCache computes token importance exactly; PrefixQuant observes token-wise outliers at *fixed positions* (initial tokens) or low-semantic tokens (".", "\n"); MiniKV finds important tokens identifiable *before* generation and stable throughout.
- **LESS** (Dong et al., 2024) — constant-sized low-rank cache + eviction grafted onto attention (recurrent-inspired); the conceptual bridge to hybrid O(1) state.
- TurboQuant's **QJL** — the production-grade instance of correcting quantization drift: 1-bit residual correction yielding an *unbiased* inner-product estimator.

### Attention-variant design notes (why the field moved)

- **Why GQA won first, then MLA:** MHA keeps n_heads KV heads per token — full expressivity but impractical long context (Llama 2 70B @ 8K ctx: >20.5 GB). MQA collapsed to one KV head (~1/n_heads cache) but paid a quality cost at scale. GQA (Ainslie et al., 2023) split the difference: 2:1 to 8:1 query-to-KV ratios, ~88% reduction, >95% quality — good enough to become the default. MLA attacks a different axis: instead of dropping heads, it compresses the *feature* dimension (512-dim latent + 64-dim decoupled RoPE key), keeping all heads. The 2026 survey verdict: under the same KV budget, MLA has higher expressive power than GQA, and DeepSeek's quality figures (matches/slightly beats MHA) have no GQA equivalent.
- **The absorption trick, operationalized:** at inference, per-head K/V are reconstructed on the fly via up-projection from the cached latent — the full K/V matrices are never materialized in memory. This is why serving MLA needs MLA-aware kernels: the cached object is a latent plus a decoupled RoPE key, not standard K/V blocks, and the attention math differs from the GQA path. vLLM 2026 auto-selects per configuration with priority ordering (rejected backends raise with the reason, e.g. compute capability); out-of-tree ports exist for Ascend (`VLLM_ASCEND_MLA_PA`), Intel XPU (0.21.0 functional validation with DeepSeek-V2-Lite), and Apple Silicon (vLLM-metal, Mar 2026: MLA paged-attention with a latent `[kv_norm || k_pe_roped]` paged cache).
- **The one-way door, restated for the consolidation:** MLA must be baked in at pretraining. It cannot be retrofitted onto a trained MHA/GQA checkpoint without expensive retraining. The pretraining-side alternatives for existing fleets are therefore (a) post-hoc low-rank retrofits ("Thin Keys, Full Values": 75% key-cache savings, <1% of pretraining data for SVD+QK fine-tuning, 16× combined with GQA+quantization), (b) survey-family methods Palu/LoRC/SVDq/CSKV/ReCalKV, or (c) serving-side compression (TurboQuant, FP8 KV) — none of which change the weight architecture.
- **MLA census detail (2026-09-08, temperature2.com):** at least eight model families besides DeepSeek — Moonshot AI's **Kimi K2 series** (K2.7 Code independently measured at 256K context in June 2026 coding benchmarks) and Zhipu's **GLM-5** (GLM-5.3: announced 2026-08-14, weights published 2026-08-29 after a two-week safety hold over unexpected offensive-security scores, 756 GB native FP8 across 141 files, Intelligence Index 60 tied with Kimi K3; GLM-5.3-Flash: 2026-08-26, MIT license, 320B/18B, natively multimodal). Holdouts are deliberate: **MiniMax-M2.5** and **Qwen3-Next** picked non-MLA designs (sparse attention and Gated DeltaNet respectively). [VENDOR] for parameter totals; [COMMUNITY] for the census framing.
- **Toy arithmetic vs measured:** the 14.2× figure (temperature2.com, Sep 2026) is toy-model arithmetic at full quality [DIRECTIONAL]; the 93.3% (vs DeepSeek 67B, May 2024) is the production-adjacent measurement; the 5.76× throughput gain is the serving-economics figure that matters for per-token pricing. All three describe the same mechanism from different angles — the consolidation must not present them as competing claims.

### MLA engine-support summary (who ships native MLA, since when)

Per-release kernel matrices (backend priority rules, PR-level history) are part 06's territory — this table answers only who/when, then cross-refs.

| Engine | First native MLA support | 2026 status |
|---|---|---|
| **vLLM** | Late 2024 (v0.6.x era; Triton MLA path for DeepSeek-V2) — exact minor [UNVERIFIED] | v0.23.0 (2026): DeepSeek-V4 production hardening (TRTLLM-gen kernel, sparse MLA metadata decoupled from V3.2); v0.28.0 (tagged 2026-08-26): sparse MLA end-to-end (plain decode, MTP, speculative); 2026 priority: `FLASHINFER_MLA_SPARSE` with FP8 KV |
| **SGLang** | v0.3, Sep 2024 ("7× Faster DeepSeek MLA") | PR #6109 (May 2025): FlashMLA + FP8 KV + MTP; v0.5.6 (Dec 2025): FP4 support for MHA+MLA KV caches; SGLang-JAX (TPU): MLA via FlashAttention Pallas kernel, no extra flag |
| **TensorRT-LLM** | Supported for DeepSeek-V3/V3.2/R1-class models — exact release [UNVERIFIED] | The TRTLLM-generation attention kernel was adopted *back* into vLLM v0.23.0 for DeepSeek-V4 |
| **Hugging Face TGI** | Gaudi branch only (v3.3.x, late 2025) — never mainline CUDA | **Archived read-only 2026-03-21** (maintenance mode) |

→ Part 06 (inference engines): full per-release kernel tables, backend-priority logic, fp8_e4m3-vs-fp8_e5m2 engine defaults, and the vLLM-Ascend / XPU / metal out-of-tree ports.

### Compression-zoo one-liners (dated, for RAG retrieval)

- **H2O — Heavy Hitter Oracle (NeurIPS 2023):** keeps a fixed-size cache mixing recent tokens and heavy hitters (cumulative attention scores); attention weights follow a power law, so evicting low-scoring tokens costs little accuracy; up to 29× throughput vs HF Accelerate on OPT-6.7B/30B; decode-phase only (prefill compute unchanged).
- **StreamingLLM (attention sinks):** keeps KV of the first few tokens — structural "attention sinks" receiving disproportionate attention regardless of content — plus a sliding window of recent tokens; fast, hardware-friendly, ideal for streaming dialogue; no importance scoring, so critical middle-context tokens can be discarded.
- **SnapKV (prefill-phase):** a small observation window at the end of the prompt votes (pooled attention scores, per-head) for important KV positions in the prefix; head-specific clustered selection makes it more accurate than H2O at the same cache budget; the standard prefill-compression baseline on LongBench.
- **PyramidKV / PyramidInfer (layer-wise budgets):** per-layer cache sizes following the pyramidal structure of attention (early layers need richer context; deeper layers converge on fewer salient tokens); PyramidInfer computes fewer K/V in deeper layers during prefill itself: 2.2× throughput vs HF Accelerate, >54% KV-cache memory reduction.
- **KIVI (ICML 2024):** plug-and-play 2-bit KV quantization, no fine-tuning — keys quantized **per-channel** (keys show channel-wise outliers), values **per-token**; on Llama-2/Falcon/Mistral: 2.6× combined peak-memory reduction (weights + KV), 4× larger batch sizes, 2.35–3.47× throughput on real workloads; the baseline TurboQuant beats on LongBench.
- **KVQuant (calibrated mixed-precision):** per-channel key quantization, **pre-RoPE key quantization** (quantize before positional embeddings distort the distribution), sensitivity-weighted non-uniform levels from calibration data, **dense-and-sparse decomposition** for extreme outliers; pushes sub-4-bit with better accuracy than fixed grids; paper evaluates up to **10M-token** contexts; calibration cost amortizes on stable production workloads.
- **SKVQ (survey corpus):** 2-bit KV quantization with sliding-window channel reordering.
- **Palu / LoRC / SVDq / CSKV / ReCalKV (low-rank family):** K/V matrices exhibit strong low-rank structure, especially at long contexts; Palu does post-training low-rank projection with medium-grained group-head decomposition and Fisher-information-based rank search assigning larger ranks to sensitive matrices; all orthogonal to quantization and eviction — stackable with both.
- **SVDq / XQuant / AdaKV (survey-flagged, veloxquant-mlx, Sep 2026):** SVDq (sub-2-bit key-only via offline SVD + per-channel mixed precision), XQuant (cross-layer quantization, sub-1.4-bit equivalent), AdaKV (head-adaptive eviction budgets via head entropy).
- **TurboQuant vs the zoo:** the 2026 reference for *online, calibration-free* KV quantization (3-bit, 6×, 8× attention speedup, ~2.7× from the information-theoretic limit); sits in the quantization family per the April 2026 MarkTechPost survey.

### Worked example: sizing a KV cache (and what each lever buys)

Representative 70B-class model: 80 layers, GQA-8, head_dim 128, BF16, batch 1. The table in "Figures and metrics" above shows the stack: at 1M tokens, BF16 KV ~336 GB per user; FP8 → ~168 GB; TurboQuant 3-bit → ~56 GB; MLA (architectural) → ~22 GB. Reading: at 1M tokens *no* single lever fits one 80GB GPU — the 2026 answer is stacking or pretraining choices. TurboQuant-class 3-bit moves 128K-class workloads onto consumer hardware (~32K → ~180K usable context on 2× RTX 5060 Ti / M-series Macs per community reports).

### FP8-KV production checklist (2026, from the source corpus)

1. Default to FP8 KV (`--kv-cache-dtype fp8` in vLLM): ~2× capacity, 54% ITL slope, break-even at ~7K tokens, <0.3% accuracy cost — but pin **`fp8_e4m3`** (production recipes standardize on it), and pair with `--dtype fp8` where the recipe calls for it.
2. Pin the accumulation fix: verify FlashAttention-3 two-level FP32 accumulation patches (`flash-attention#96`, `#91`) are present on Hopper *before* serving; the 91%→13% NIAH collapse is the canary.
3. Needle-test every enablement: run a 128K NIAH probe after any KV-dtype change; <50% recall = the accumulation bug, not the model.
4. Use `--kv-cache-dtype-skip-layers` for sensitive layer types (sliding-window, early layers) on hybrid models; check `--mamba-cache-dtype` supports only auto/bf16/f16/f32 — GDN state is tiny (~0.02 GiB), a non-issue, not a blocker.
5. Calibrate, don't guess: `fp8_e4m3` vs `fp8_e5m2` is a precision-vs-dynamic-range knob; calibrated scales (llm-compressor) beat naive per-tensor.
6. Budget the next step: INT4 KV (4×, KIVI/SKVQ-class methods) and TurboQuant-style 3-bit (5.3×) for the capacity-constrained; reasoning-aware R-KV/FoveatedKV when serving reasoning models with long traces. FP8-KV performance is kernel-fusion-dependent, not a dtype property (SGLang fused fp8-KV kernels vs vLLM-ROCm unfused Triton dequant fallback — [SECONDARY, HARDWARE-SPECIFIC]).
7. On hybrid models (Qwen3-Next-class), account for what quantizes and what does not: FP8 KV roughly doubles KV capacity (e.g. toward 512K `max_model_len` with YaRN scaling); recurrent state stays full-precision and tiny.

### Terminology anchors (for the consolidated document's consistency)

- **KV-cache:** stored key/value tensors of all past tokens; read on every decode step. The memory bottleneck of inference.
- **Prefill:** processing the full prompt in parallel (compute-bound). **Decode:** generating tokens one by one (memory-bandwidth-bound). **TTFT:** time to first token. **TPOT/TBT:** time per output token.
- **MHA/MQA/GQA/MLA:** attention variants trading head-axis redundancy (MHA→MQA→GQA) or feature-axis compression (MLA) for cache size.
- **Eviction:** permanently dropping tokens from the cache (H2O, StreamingLLM, SnapKV). **Quantization:** storing cache at lower bit-width (KIVI, KVQuant, TurboQuant, FP8). **Low-rank:** projecting cache into smaller dimensions (Palu, MLA).
- **NIH (needle-in-haystack):** benchmark retrieving one buried fact from a long context; the standard long-context quality probe. TurboQuant reported lossless NIH at 3-bit on Gemma/Mistral/Llama; MiniMax's MSA-vs-MLA argument is precisely NIH-style retrieval fidelity (uncompressed KV blocks vs compressed latents).
- **QJL (Quantized Johnson-Lindenstrauss):** TurboQuant's 1-bit residual correction giving an unbiased inner-product estimator.
- **PolarQuant:** the rotation-into-polar-coordinates stage of TurboQuant (AISTATS 2026) that makes scalar quantization near-optimal and eliminates expensive normalization.
- **Decoupled RoPE:** MLA's positional design — content K/V through the latent path, a small shared RoPE key (64 dims) outside the compression.
- **Absorption trick:** MLA's on-the-fly reconstruction of per-head K/V from the cached latent via up-projection.
- **CSA / HCA:** DeepSeek-V4's hybrid Compressed Sparse Attention + Heavily Compressed Attention (the cache-side compression DSA lacks on its own).
- **IndexShare (GLM-5.2):** one sparse-attention indexer shared across every four sparse-attention layers (2.9× FLOP cut at 1M, +20% MTP acceptance [VENDOR]).
- **LatentMoE (Nemotron-3-Ultra):** Mamba-2 + MoE + Attention hybrid at 550B/55B scale.
- **Gated DeltaNet (GDN):** linear-recurrent/SSM-style layer family (Qwen3-Next, Qwen3.5/3.6) carrying zero KV cache.



