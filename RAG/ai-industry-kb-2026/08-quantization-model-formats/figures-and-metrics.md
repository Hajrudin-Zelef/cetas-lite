---
id: ai-industry-kb-2026/08-quantization-model-formats/figures-and-metrics
title: "Figures and metrics"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: []
keywords: ["amd", "awq", "aws", "blackwell", "consumer", "cost", "datacenter", "decode", "deepseek", "fp4", "fp8", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4013, 4088]
section: "8. Quantization & Model Formats"
sha256: d193d6176bb04fb61953e29f45e21b29d9c26d3cc96db68f9f9329e0a7ccaff5
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

