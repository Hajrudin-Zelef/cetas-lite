---
id: ai-industry-kb-2026/08-quantization-model-formats/figures-and-metrics
title: "Figures and metrics"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Hugging Face", "Intel", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-02-20", "2026-05", "2026-09"]
keywords: ["amd", "awq", "aws", "bitnet", "blackwell", "compute", "consumer", "cost", "datacenter", "decode", "deepseek", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4013, 4113]
section: "8. Quantization & Model Formats"
sha256: fd9fcea1f8271e8d4ec1534d38e7a7b208899570abe0478228ec65b80dbbfe0a
---

# Figures and metrics

## Figures and metrics

### GGUF quality ladder (Llama-2-7B-class reference, illustrative)

| Quant | Perplexity | Change vs FP16 | Size |
|---|---|---|---|
| FP16 | 5.9565 | baseline | 13.0 GB |
| Q8_0 | 5.9584 | +0.03% | 7.0 GB |
| Q6_K | 5.9642 | +0.13% | 5.5 GB |
| Q5_K_M | 5.9796 | +0.39% | 4.8 GB |
| Q4_K_M | 6.0565 | +1.68% | 4.1 GB |

### GGUF quality-vs-FP16 rules of thumb (2026 community)

| Quant | Quality vs FP16 | Size ratio | Guidance |
|---|---|---|---|
| Q8_0 | ~99.5% | 0.55x | Near-lossless; the standard "safe" GGUF choice |
| Q6_K | ~99% | 0.42x | Diminishing returns above this |
| Q5_K_M | ~98% | 0.36x | Recommended for ≤14B models |
| **Q4_K_M** | **~96%** | **~0.30x** | **Best quality/size trade-off; the default** |
| Q3_K_M | ~90% | 0.25x | Acceptable for very large models |
| Q2_K | ~80% | 0.18x | Aggressive; low-stakes use only |

### NVFP4 ≠ MXFP4: format anatomy (don't conflate the two "FP4"s)

| Property | NVFP4 (NVIDIA, Blackwell-proprietary) | OCP MXFP4 (open spec, v1.0 Sept 2023) |
|---|---|---|
| Element format | FP4 E2M1 | FP4 E2M1 |
| Block size | **16** (finer granularity) | **32** |
| Per-block scale | **FP8 E4M3** (higher-precision scale) | **E8M0** (8-bit exponent-only) |
| Extra scale | **FP32 per-tensor global scale** (extends range past the ±448 block cap) | none |
| Effective density | **~4.5 bits/element** | 4.25 bits/element |
| Native tensor-core matmul | Blackwell (SM100/103; B200 FP4 ≈ 8× FP16 [VENDOR]) | Blackwell native; AMD MI355X measured 6.1× throughput (Dell) |
| Production flagships | DeepSeek-R1-FP4 (671B, MMLU −0.1 vs FP8); Kimi-K2.6-NVFP4 (1T/32B) | GPT-OSS-120B (single H100); GPT-OSS-20B (16 GB) |

- Full production arc (DeepSeek-R1-FP4 accuracy, Kimi-K2.6-NVFP4, DeepSeek-V3.2/R1 GB300 throughput, consumer-Blackwell W4A4 caveats, NVFP4 KV datacenter-only) → **sibling part 08b**.
- FP4 as an inference format only: a tested DGX Spark recipe had BF16 beating NVFP4 at training (~17,500 vs ~13,000 tok/s) [COMMUNITY]; NVFP4 KV cache does not work on consumer Blackwell (SM120).

### Engine-format equivalence at ~4-bit

| Engine | 4-bit equivalent | Hardware gate |
|---|---|---|
| llama.cpp | Q4_K_M / UD-Q4_K_M / IQ4_XS | any (CPU→GPU) |
| vLLM | AWQ W4A16 (best compat quality), GPTQ W4A16, NVFP4 | AWQ/GPTQ SM75+ (Marlin); NVFP4 W4A4 SM100+ |
| SGLang | AWQ/GPTQ (Triton/vLLM kernels), torchao int4wo-128, NVFP4 (SM100+) | per-engine matrix |
| TensorRT-LLM | FP8, NVFP4, MXFP4 (ModelOpt), INT4 AWQ (legacy) | compiled per-GPU; W4A4 needs Blackwell |
| ExLlamaV3 | EXL3 / QTIP (~4 bpw) | consumer NVIDIA |

### Throughput context (A100 80GB, Llama 3.3 70B, batch=1, estimated)

| Quant | Decode tok/s | VRAM | Notes |
|---|---|---|---|
| Q4_K_M | ~20–30 | ~42 GB | Half the hourly cost of FP16 on 2× A100 |
| Q5_K_M | ~28–36 | ~48 GB | Near-lossless tier |
| Q8_0 | ~20–25 | ~75 GB | Indistinguishable from FP16 |
| FP16 | ~30–40 | 140 GB (2× A100) | Baseline |

### Cost arithmetic

- **Halving the GPU count halves the hourly bill.** FP16 Llama 3.3 70B needs 2× A100 (140 GB); Q4_K_M fits on 1× A100 (~42 GB) at similar batch-1 throughput — roughly **2× cheaper per hour** for the same workload.
- **OpenAI's MXFP4 deployment** of GPT-OSS is credited with **~75% inference-cost reduction** versus the hardware FP16 would have required (120B on a single H100 instead of a multi-GPU node).
- **Local-first economics**: a 27B at 4.2 GB (Bonsai) or a 70B at 40 GB (Q4_K_M) runs on hardware the user already owns — marginal inference cost drops to electricity. This is the economic force behind "inference as local utility": quantization moved the break-even point of self-hosting far below cloud per-token pricing for steady workloads.
- The counterweight: **quantization labor is not free** — GPTQ calibration runs, QAT retraining, and per-hardware kernel validation (e.g. the sm_120 FP4-KV saga) are engineering costs that favor standardized formats (FP8, MXFP4) over bespoke ones at scale.
- AWS Blackwell rental economics (NextPlatform sizing): FP8 halves cost-per-teraflop vs FP16; FP4 halves it again — a model moved to FP4 needs **one quarter of the machinery at one quarter of the cost**, or the same money trains a 4× larger model.

### Model footprints that define 2026 deployments

| Model scale | Format | Footprint | Where it runs (verified) |
|---|---|---|---|
| 70B dense | GGUF Q4_K_M | ~40 GB | 64 GB laptop; single 48 GB GPU via INT4 |
| 744B MoE (GLM-5.2, 40B active) | UD-IQ2_M | **239 GB disk** | 256 GB unified-memory Mac; or 24 GB GPU + 256 GB RAM offload |
| 236B-class (DeepSeek-V2.5) | i1-IQ4_XS | 125.7 GB | ~0.8% perplexity increase |
| 27B (Qwen3.8) | UD-Q4_K_XL | 16.69 GiB | ~96% top-1 BF16 agreement [DIRECTIONAL] |
| 671B (DeepSeek-V3.1) | UD-TQ1_0 | ~192 GB | 1-bit extreme compression |
| 30B-class (Nemotron-3-Nano) | ModelOpt FP8 | 2.6× throughput + 2.6× memory (with Minitron pruning) | vLLM |

## Main actors

| Actor | Role in the 2026 format/tooling story |
|---|---|
| **Hugging Face** | Absorbed ggml.ai (2026-02-20); hosts tens of thousands of GGUF checkpoints with metadata viewer and JS parser; redirects optimum-quanto users to bitsandbytes/torchAO |
| **ggml.ai / Georgi Gerganov** | llama.cpp and GGUF creator; team now HF employees with full technical leadership and MIT licensing preserved; 126,000+ GitHub stars by 2026 |
| **Unsloth** | Dynamic v3.0 selective-layer GGUF quants (UD- prefix, 1,374-model catalog, pure PTQ); proprietary `unsloth_zo` schedule; NVFP4-GGUF hybrid containers |
| **NVIDIA** | NVFP4 format; ModelOpt 0.46.0 Blackwell-native toolchain; DeepSeek-R1-FP4 and Kimi-K2.6-NVFP4 production proofs; builds FP4 Tensor-Core hardware (SM100/103/120) |
| **turboderp (ExLlamaV3)** | Active successor to archived ExLlamaV2; EXL3/QTIP format; v1.4.6–v1.4.9 September 2026 patch cadence; vLLM-EXL3 third-party serving plugin |
| **vLLM project** | Datacenter inference stack; llm-compressor (FP8/AWQ/NVFP4/MXFP4); kernel gates (Marlin SM75, Machete SM90, NVFP4→flashinfer_cutedsl); `--linear-backend` replaces `VLLM_NVFP4_GEMM_BACKEND` |
| **SGLang** | Parallel datacenter stack: awq, awq_marlin, gptq_marlin, modelopt_fp4, mxfp4, fp8; `--load-format gguf` (NVIDIA-only) |
| **Intel** | AutoRound — strongest measured 4-bit PTQ (91.4% logprob top-1 agreement vs BF16); Neural Compressor; OpenVINO 2026.2 INT4 KV on Intel GPU; vLLM-XPU validation |
| **mobiusml** | HQQ (data-free/fast 1/2/3/4/8-bit) and gemlite (Triton kernels incl. MXFP4/NVFP4 dynamic); HQQ+ low-rank adapters |
| **ModelCloud** | GPTQModel — AutoGPTQ successor, 5,000+ HF repos reference it |
| **Red Hat AI** | May 2026 vLLM evaluation of TurboQuant KV-cache variants (see 08b); publishes quantized checkpoints (e.g. Qwen3.5-122B-A10B-NVFP4) |
| **Microsoft Research** | BitNet natively-trained 1.58-bit ternary models (see 08b); contributes to OCP MX spec |
| **PrismML** | Bonsai-27B 1-bit/ternary GGUFs (see 08b); `prism-ternary` experimental vLLM-GGUF plugin branch |
| **mradermacher** (community) | i1- imatrix GGUF cards documenting the per-type ladder across DeepSeek-V2.5, QwQ-32B, Qwen2.5-14B, Gemma-2-9B; ARM edge variants Q4_0_4_4 / Q4_0_4_8 / Q4_0_8_8 |
| **Artefact2** (community) | Reference per-type KL/PPL table used as the GGUF accuracy standard |
| **MarcoPizeta** (community) | Reproducible vLLM-vs-EXL3 shootout on RTX PRO 6000 (Sept 2026) |
| **AMD** | ROCm path (ROCm 6.4.4 beats 7.0.1 3.47× on llama.cpp); MI355X + MXFP4 6.1× measured (Dell); AMD Quark on ROCm |
| **Open Compute Project** | MX specification v1.0 (Sept 2023) — the open microscaling standard behind MXFP4 |

## Timeline and context

