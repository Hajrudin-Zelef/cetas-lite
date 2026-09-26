---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/worked-example-sizing-a-kv-cache-and-what-each-lever-buys
title: "Worked example: sizing a KV cache (and what each lever buys)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "DeepSeek", "Falcon", "MiniMax", "Mistral", "Oracle", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-04"]
keywords: ["kv cache", "attention", "benchmark", "compute", "consumer", "cost", "decode", "deepseek", "distribution", "embeddings", "fine-tuning", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3187, 3231]
section: "7. KV Cache & Long-Context Techniques"
sha256: 049dc996eb079d2ed5ee067534fd6e1bf290ac83d65c1f8baa3debcef047a50f
---

# Worked example: sizing a KV cache (and what each lever buys)

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



### Gemma 4 family snapshot (community-verified, mid-2026)

