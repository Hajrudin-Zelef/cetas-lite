---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/7-quantization-tooling-bitsandbytes-gptq-awq-gguf
title: "7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: quantization
actors: ["AMD", "Nvidia", "vLLM"]
dates: ["2026-09"]
keywords: ["awq", "gguf", "gptq", "quantization", "accelerator", "amd", "compute", "fine-tuning", "fp8", "gpu", "gpus", "inference"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [476, 534]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: d4788a9f3a24cbefd6766b50bfa98663bf83667887c8c6e58942445d66860f13
---

# 7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF

## 7. Quantization tooling: bitsandbytes, GPTQ, AWQ, GGUF

> Scope note: inference-server quantization (llama.cpp/vLLM paths) is
> covered in Step 4; this section covers the developer-facing tooling.

### 7.1 bitsandbytes

- bitsandbytes provides 8-bit optimizers, `LLM.int8()` mixed-precision
  decomposition, NF4 4-bit quantization, and double quantization
  [secondary].
- Its signature use case is **QLoRA**: 4-bit base model + trainable LoRA
  adapters, enabling fine-tuning of large models on single GPUs
  [secondary].
- It is load-time quantization integrated with Transformers/PEFT rather
  than a calibrated offline format [secondary].

### 7.2 GPTQ and AWQ (calibrated weight-only formats)

- **GPTQ** (one-shot weight quantization via approximate second-order
  information) and **AWQ** (activation-aware weight quantization) are
  calibrated, weight-only GPU formats: a calibration dataset sets scales
  before the model is saved in the quantized format [secondary].
- 2026 community guidance reports **AutoGPTQ as unmaintained** and
  recommends **GPTQModel** as the successor — not verified against the
  repositories; treat as [unverified] (gap G-10).
- Both formats primarily save memory; speedups are hardware- and
  workload-dependent, and weight-only quantization does not inherently
  accelerate compute-bound workloads [secondary].

### 7.3 GGUF (cross-reference Step 4)

- GGUF is the portable container format of the llama.cpp ecosystem,
  designed for CPU/GPU/edge inference across architectures [secondary].
- Full treatment (quants like Q4_K_M, imatrix calibration, server-side
  usage) is in Step 4; noted here only as the portable counterpart to
  the GPU-oriented GPTQ/AWQ formats [secondary].
- A September 2026 community comparison covers GGUF vs GPTQ vs AWQ vs
  EXL2 trade-offs for practitioners [secondary].

### 7.4 FP8 and the precision landscape

- FP8 (E4M3/E5M2) training and inference is hardware-gated: NVIDIA Hopper
  and newer, AMD CDNA with FP8 support; the JAX 26.05 notes cite XLA
  `nvfp4` codegen improvements, showing NVFP4 entering the 2026 stack
  [secondary].
- Quantization choice is a three-way trade: memory footprint, throughput
  on the target accelerator, and quality degradation on the target task;
  no format dominates all three [secondary].

### 7.5 Quantization decision matrix

| Tool/format | When it runs | Granularity | Best fit |
|---|---|---|---|
| bitsandbytes NF4 | Load time | 4-bit weights, 8-bit optimizers | QLoRA fine-tuning [secondary] |
| GPTQ/GPTQModel | Offline calibration | 4-bit weight-only | GPU serving, memory-bound [unverified for successor naming] |
| AWQ | Offline calibration | 4-bit weight-only, activation-aware | GPU serving, quality-sensitive [secondary] |
| GGUF | Offline conversion | K-quants, imatrix | Portable CPU/edge/GPU (see Step 4) [secondary] |
| FP8/NVFP4 | Train/infer | 8-bit float | Hopper+/CDNA8+ hardware [secondary] |

