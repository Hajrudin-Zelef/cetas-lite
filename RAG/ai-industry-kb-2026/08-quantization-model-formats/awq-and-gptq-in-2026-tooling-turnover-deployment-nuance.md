---
id: ai-industry-kb-2026/08-quantization-model-formats/awq-and-gptq-in-2026-tooling-turnover-deployment-nuance
title: "AWQ and GPTQ in 2026: tooling turnover, deployment nuance"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: ["2025-05"]
keywords: ["awq", "gptq", "amd", "benchmarks", "blackwell", "decode", "fp4", "fp8", "gguf", "int4", "llama", "llama.cpp"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3893, 3917]
section: "8. Quantization & Model Formats"
sha256: 48ab485aa4eab234562832917281ccef86a4216df66626aa0d919be0e5789c6a
---

# AWQ and GPTQ in 2026: tooling turnover, deployment nuance

- **Prefill: vLLM ~2.7× faster and flat from 8k to 128k** (28–29.6k tok/s vs EXL3 ~10–11.6k tok/s); at 128k only 5/9 EXL3 requests per config completed. The TabbyAPI failures are non-JSON SSE chunk errors on concurrent long-prompt ingestion — a TabbyAPI/ExLlamaV3 serving-path limit, not an MTP or bpw issue.
- EXL3 wins **only short single-user chat** (169–173 tok/s at 1k, loads in 1–1.5 min vs 5–6); vLLM wins from 8k context up and at any concurrency > 1, is stable at 32k × 4/8/16, and needs only ~10–12 GB host RAM vs ~45 GB for TabbyAPI's pinned n-gram table.
- Median TPOT at c=1 with MTP: EXL3 **5.5 ms** vs vLLM 6.1 ms — ExLlamaV3's pure decode step is faster; it loses on prefill and scaling. MTP is worth +30% on vLLM, +60% on EXL3 at c=1. **5.05bpw runs at the same speed as 4.05bpw.**
- **Quality: quantization is not the differentiator.** Flash-Next vs incumbent Qwen3.8-27B on tool-eval-bench: 27B scores 91.0 ± 1.5 vs Flash-Next 87.9 ± 1.9 (vLLM, reasoning_effort=low); 27B BF16 scores 89.0 (≤ NVFP4 91.0), EXL3 5.05bpw 86.9 ≈ 4.05bpw 87.4. On AIME 2025+2026 they are equal (48/60 vs 50/60).
- **SGLang result (v3)**: Flash-Next quality 88.0 low / 88.9 thinking-off, `tool_choice="required"` enforced 7–8/8, warm prefix TTFT 210 ms, zero server-side errors — but with a file-backed (not pinned-RAM) n-gram table on the 64 GB host: **37–51 tok/s flat across all points**; cookbook pinned-RAM numbers (148 tok/s single, 613 at 16) not reproduced due to host-RAM constraints.
- Caveats: single workstation; TabbyAPI 21/60 failure mode may be TabbyAPI-version-specific rather than EXL3-intrinsic. Implication for the "only credible raw-TPS challenger to llama.cpp" verdict: keep, sharpened — EXL3's challenge is single-stream short-context decode; on prefill (2.7× deficit), concurrency scaling, and operational footprint, the vLLM NVFP4/FP8 stack dominates. EXL3's bottleneck is serving-path maturity, not kernel speed.

### AWQ and GPTQ in 2026: tooling turnover, deployment nuance

- **AutoAWQ is officially deprecated.** Last tested setup: Torch 2.6.0 + Transformers 4.51.3. The AWQ quantization workflow now lives in **llm-compressor** (vLLM project) — the recommended path — and MLX-LM also supports AWQ on Apple Silicon.
- **GPTQModel** (ModelCloud) is the AutoGPTQ successor: 2/3/4/8-bit quantization with latest-model support; **5,000+ GPTQ repos on Hugging Face** reference it. Hugging Face's quantization overview now lists GPTQModel (not AutoGPTQ) as the GPTQ entry.
- Methods recap: **AWQ** (Activation-aware Weight Quantization) scales weights by activation importance before quantization; consensus best INT4 PTQ quality; W4A16 typical, group size 128; official kernel is the default vLLM path, with **Marlin/Machete** kernels available when large-batch throughput matters. **GPTQ** quantizes weights column-by-column, compensating accumulated rounding error via the inverse Hessian; requires calibration data (~5–30 minutes for 7B); slightly worse quality than AWQ on most benchmarks but **widest hardware compatibility** (including via GGUF); default kernel ExLlamaV2.
- Kernel gates (audited against vLLM v0.26.0 source): **Marlin's floor is SM75 (Turing), not SM80** (`if device_capability < 75: return []` — RTX 20-series not excluded); **Marlin-24 act-order kernels removed from vLLM main** (existed at tag v0.9.0); **Machete requires exactly SM90 (Hopper)** — handles odd layer dims (e.g. vision towers) that Marlin rejects; Conch targets SM80+; CUTLASS native FP8 GEMM needs cc ≥ 89; FP8 *weights* (W8A16) via Marlin-FP8 works from SM75.
- The old `VLLM_NVFP4_GEMM_BACKEND` environment variable is **deprecated**, replaced by `--linear-backend` (vLLM auto-selects CUTLASS/FlashInfer/Marlin at load).
- vLLM quantization→backend mapping (Isambard, Aug 2026): FP8 → `deep_gemm`/`hpc`/`flashinfer_cutlass`; **INT4 AWQ/GPTQ → `marlin`**; NVFP4 → `flashinfer_cutedsl`; W4A8 → `humming`.
- Hard-won deployment rules: Marlin rejects symmetric AWQ `uint4` with `zero_point=false` (AutoRound quants) on Ampere — use a compressed-tensors quant; Marlin needs layer dims divisible by 128 (multimodal vision MLPs with odd dims need Hopper/Machete); AWQ `quant_method` requires `--dtype float16` (bf16 rejected at config validation); vLLM's `--quantization awq_marlin` flag is explicitly required for AWQ checkpoints (config.json alone insufficient).
- On AMD, AWQ/GPTQ/Marlin remain NVIDIA-only — the ROCm paths use FP8 via llm-compressor, GGUF, or AMD Quark.
- Quality consensus (2026): AWQ holds slightly lower perplexity than GPTQ at equal bits; GGUF Q4_K_M sits between them depending on model. A French 2026 industry guide puts well-calibrated 4-bit degradation at **2–4 points on MMLU for 7B+ models**.
- **Verifying the brief's "AWQ as the production standard" claim**: both engines ship AWQ paths — vLLM supports awq/awq_marlin/auto_awq; SGLang supports awq plus awq_marlin, gptq_marlin, and more. A mid-2026 comparison calls AWQ **"the more commonly used of the two [vs GPTQ] for vLLM deployments"** — the "standard" part is supported on vLLM; **6,500+ AWQ repos** on the Hugging Face Hub.
- Nuances against the brief's framing: **workflow friction on 2026 models** — AutoAWQ was archived in May 2025 and never got Qwen3.5 support (hard `TypeError`); the **llm-compressor pip release pins `transformers<=4.57.6`**, predating Qwen3.5 support (transformers 5.x) — source install required; AWQ's layer-by-layer scale search **OOMs on a 40 GB A100** for hybrid architectures (fix: `offload_device=cpu` on the AWQModifier). New-model coverage is where the "standard" frays: AWQ tooling lags architectures, FP8/MXFP4 paths need no calibration at all.
- [COMMUNITY] A 4× RTX PRO 6000 Blackwell shootout (vLLM 0.17.0rc1, TP4) had **AWQ beating both NVFP4 recipes at every concurrency** — C=128: AWQ 3519 vs 3232/3220 tok/s with MTP; 2796 vs 2294/2291 without MTP (single rig; NVFP4 path/backend version-sensitive). Quality deltas on a 19-question math set were mixed.
- Net verdict: **PARTIALLY VERIFIED / NUANCED**. AWQ is the dominant *compat* INT4 path on both engines and the most-used vLLM INT4 format — but not where new quantization work targets (llm-compressor/AutoRound/ModelOpt/FP8/NVFP4 per the decision matrix), and "preserving salient activations" is the algorithm's 2023-era technical edge, not a 2026 differentiator over FP4/FP8.

### The 2026 tooling verdict

