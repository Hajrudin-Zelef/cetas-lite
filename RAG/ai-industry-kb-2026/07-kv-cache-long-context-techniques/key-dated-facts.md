---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/key-dated-facts
title: "Key dated facts"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "DeepSeek", "Google", "Hugging Face", "Intel", "MiniMax", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2024-05", "2025-03-16", "2025-03-29", "2025-05", "2026-03-21", "2026-04-02", "2026-05-11", "2026-06", "2026-06-01", "2026-06-11", "2026-08-26", "2026-09-08"]
keywords: ["amd", "apache", "attention", "compute", "cost", "decode", "deepseek", "fp4", "fp8", "glm", "gpu", "gqa"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3012, 3034]
section: "7. KV Cache & Long-Context Techniques"
sha256: 69138277edfe63efaa8c8ad803df3109eced0a35d7414680f7fa0598de2bf58d
---

# Key dated facts

**Claim-status corrections carried from the brief.** (a) TurboQuant is a Google invention — not Red Hat's; Red Hat published an independent vLLM evaluation 2026-05-11. (b) "MLA exactly 4× batch/GPU" is only partially verified — no primary measurement pins that figure; use the measured 93.3% / 5.76× throughput / ~14.2× toy-arithmetic figures instead. (c) DSA reduces attention *compute* — do not equate it with KV-cache reduction. (d) The brief's "dynamic drift correction" names no single documented technique [UNVERIFIED]; the documented mechanisms are inventoried in §"Figures and metrics". (e) "~90% MLA reduction" is VERIFIED-conservative: the DeepSeek-V2 paper reports 93.3% (DeepSeek-67B baseline), the controlled ablation ~96%, the raw per-layer element ratio ~98.2%.

## Key dated facts

### Attention variants and MLA (architecture-level)

- **May 2024 — DeepSeek-V2 paper (arXiv:2405.04434).** Multi-Head Latent Attention (MLA) debuts: **93.3% less KV cache than DeepSeek 67B**, 5.76× higher max generation throughput, modeling quality matching or slightly beating MHA. The controlled MHA-vs-MLA ablation is ~96% (34.6K vs 860.2K cached elements/token); the raw per-layer element ratio (576 = 512 latent + 64 decoupled-RoPE vs 32,768 MHA elements) is ~98.2%. MLA's positional fix: decoupled RoPE — content keys/values go through the latent path while a small separate shared RoPE key (64 dims) carries position outside the compression, since standard RoPE rotation would break the low-rank absorption trick.
- **Sep 2023–2024 — GQA becomes the open-weight default.** Llama 2: only the 70B used GQA (7B/13B stayed MHA); **Llama 3 extended GQA across all sizes**; Mistral applied it from its first 7B release (Sep 2023); Gemma 2/4, Qwen, MiMo families all ship GQA variants (typically 2:1 to 8:1 query-to-KV ratios). Consensus 2026 figure: ~88% KV reduction vs MHA while retaining >95% of MHA quality. Its measurable weakness (worse modeling performance than MHA at matched size) pushed DeepSeek toward MLA.
- **Sep 2024 (SGLang v0.3) — "7× Faster DeepSeek MLA."** SGLang served MLA (DeepSeek-V2) a full quarter before the V3 wave. **Late 2024 (vLLM v0.6.x era):** DeepSeek-V2 MLA support landed with the Triton MLA attention path (`TRITON_ATTN` / `deepseek_mla`); DeepSeek-V3 support hardened in v0.7.x (Jan 2025). Exact vLLM minor version [UNVERIFIED] — treat "v0.6.x, late 2024" as the support window. TensorRT-LLM: native MLA for DeepSeek-V3/V3.2/R1-class models documented across NVIDIA 2025–2026 serving materials, exact release [UNVERIFIED] — "supported, version unpinned".
- **2026-03-21 — Hugging Face TGI archived read-only** (maintenance mode; HF recommends vLLM/SGLang). TGI only ever gained MLA on the **Intel-Gaudi branch** (v3.3.x, late 2025), never mainline CUDA. MLA support is therefore **3/4 engines** (vLLM, SGLang, TensorRT-LLM), not "all engines".
- **v0.23.0 (2026) — vLLM DeepSeek-V4 production hardening:** sparse MLA metadata decoupled from V3.2, a TRTLLM-gen attention kernel (TensorRT-LLM-derived) adopted back into vLLM for DeepSeek's attention pattern, EPLB for the Mega-MoE. **v0.28.0 (tagged 2026-08-26):** sparse MLA end-to-end (plain decode, MTP, DSpark speculative decoding); AMD Quark NVFP4 support. **2026 backend-priority rule:** for sparse MLA, FP8 KV cache *always* prefers `FLASHINFER_MLA_SPARSE`; with BF16 KV, `FLASHINFER_MLA_SPARSE` preferred at ≤16 query heads, `FLASHMLA_SPARSE` otherwise. 2026 MLA backends in vLLM mainline: `FLASHMLA`, `FLASH_ATTN_MLA`, `FLASHINFER_MLA` (+ sparse variants), `TRITON_ATTN`, `CUTLASS_MLA`, `XLA`.
- **2025-03-16 / 2025-03-29 / 2025-05 — SGLang MLA milestones:** FlashMLA backend (PR #4079), FlashAttention-3 backend for MLA, then **PR #6109 (May 2025): FlashMLA backend gains FP8 KV cache** — ~30% speedup from MTP+FP8-KV, KV cache halved. **v0.5.6 (Dec 2025):** MHA and MLA KV caches refactored to support **FP4**; SemiAnalysis measured DeepSeek-R1 NVFP4 on B200 reaching 907 tok/s/GPU at concurrency 4 (1.79× over 0.5.5).
- **By 2026-09 — at least eight model families besides DeepSeek adopted MLA**, including Moonshot AI's **Kimi K2 series** and Zhipu's **GLM-5 / GLM-5.3** (temperature2.com census, 2026-09-08). DeepSeek lineage: V2 (May 2024, MLA debut) → V3 (671B MoE, MLA retained) → R1 (reasoning, MLA) → V4/V4 Pro (1.6T-class, MLA). Research cited in 2026 survey literature shows MLA has **higher expressive power than GQA under the same KV-cache budget**.
- **2026-09-08 — temperature2.com analysis:** toy-model arithmetic at **14.2× smaller cache than MHA at full quality**. Caveat: toy arithmetic, not a production measurement.

### Shared and hybrid attention layouts

- **2026-04-02 — Gemma 4 ships (Apache 2.0).** Shared KV cache: the last N layers reuse KV computed by earlier layers (llama.cpp: `n_kv_shared_layers` / reuse of `n_layer_kv_from_start` KV blocks, confirmed in community runtime notes May–June 2026). Interleaved local sliding-window + global attention inherited from Gemma 2: E2B/E4B 512-token windows; 26B/31B 1024-token windows; every ~6th layer global (31B: 60 layers → 10 global + 50 local); final layer global. **GQA 2:1** (31B: 32 query heads, 16 KV heads, head_dim 256) with *dual head dims* (global head_dim vs SWA head_dim — tooling must read shapes per layer). Proportional RoPE (p-RoPE) on global layers. Community-verified KV math for Gemma 4 31B at 256K: ~40.8 GiB BF16 (dominated by the 10 global layers; the 50 local layers hold a constant ~0.78 GiB floor); with FP8 KV: ~20.4 GiB.
- **2026-06-01 — MiniMax M3 launches.** ~428B total / ~23B active MoE (128 experts, 4 active/token), 60 layers, native text/image/video multimodality, 1M-token context with **512K guaranteed usable floor**; weights + tech report (arXiv:2606.13392) ~June 11, 2026. **MiniMax Sparse Attention (MSA):** two-stage — lightweight index branch selects which KV *blocks* are relevant per input, main attention layer processes only those blocks ("KV outer gather Q" kernel — each KV block read from memory once, contiguous access). Vendor-reported: **>9× faster prefill, >15× faster decode** at 1M tokens vs previous generation; **28.4× reduction in per-token attention compute** vs full attention at 1M; per-token compute ≈ 1/20th of previous generation; kernel >4× faster than open-source sparse-attention alternatives (Flash-Sparse-Attention, flash-moba). Independent speed verification pending mid-2026 [VENDOR]. Deliberate tradeoff vs MLA: MSA operates on *uncompressed* KV, preserving long-context retrieval accuracy (NIH-style fidelity argument) at slightly higher memory cost.
- **2026 — Qwen3-Next.** Deliberately avoided MLA: **Gated DeltaNet** (linear-recurrent/SSM-style) layers for a large fraction of depth — these layers carry **zero KV cache**; KV growth confined to remaining full-attention layers. Sep 2026 FP8-KV deployment analysis: GDN recurrent state ~0.02 GiB; vLLM's `--mamba-cache-dtype` accepts only auto/bfloat16/float16/f32 (no FP8 option — irrelevant given the tiny state); `--kv-cache-dtype-skip-layers` keeps sensitive layer types at bf16 while quantizing the rest.

### Post-hoc KV compression methods

