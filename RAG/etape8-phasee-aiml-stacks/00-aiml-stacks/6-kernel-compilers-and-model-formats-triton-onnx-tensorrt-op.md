---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/6-kernel-compilers-and-model-formats-triton-onnx-tensorrt-op
title: "6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: quantization
actors: ["AMD", "Intel", "Nvidia", "OpenAI", "TensorRT-LLM"]
dates: ["2026-05-08"]
keywords: ["tensorrt", "amd", "attention", "blackwell", "cost", "fp8", "gpu", "gpus", "inference", "inference engine", "intel", "nvidia"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [404, 475]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: ae9cb1b5f235078197b0c3209d0240d152055bc3bb86539427c8c6131786f408
---

# 6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO

## 6. Kernel compilers and model formats: Triton, ONNX, TensorRT, OpenVINO

### 6.1 OpenAI Triton

- Triton is a Python-embedded language and compiler for writing custom GPU
  kernels (tiled, block-pointer programming model), lowering through MLIR
  [secondary].
- The 3.x generation is described as the MLIR-based line used by PyTorch
  Inductor as its GPU code-generation target [secondary].
- Reported concepts: block pointers, JIT kernel caching, autotuning over
  tile configurations, and NVIDIA + AMD backends [secondary].
- One secondary article claimed Triton 3.7 in 2026 — not verified against
  the official repository; treat as [unverified].
- Known limitations per secondary coverage: newer Blackwell/CDNA kernel
  paths lag upstream Triton support, so day-one GPU bring-up still
  depends on vendor libraries [secondary].

### 6.2 ONNX and ONNX Runtime

- ONNX is the open interchange format for ML graphs (operators, types,
  metadata), stewarded by the LF AI & Data Foundation [secondary].
- ONNX Runtime is the cross-platform inference engine with execution
  providers (CUDA, TensorRT, CoreML, DirectML, OpenVINO, CPU) [secondary].
- PyPI/NuGet registry data gives a clean 2026 cadence: 1.24.x (Feb–Mar
  2026), 1.25.x (Apr 2026), **1.26.0 (May 8, 2026)**, **1.27.0 (Jun 16,
  2026)**, **1.28.0 (Jul 25, 2026)**, **1.29.0 (Aug 12, 2026)**,
  **1.30.0 (Sep 10, 2026)** [official].
- 1.30.0 is therefore the current release as of the cutoff [official].
- Typical developer flow: export from PyTorch (`torch.onnx.export`,
  dynamo-based exporter) → optimize/quantize in ONNX Runtime → serve
  with the best execution provider per target [secondary].

### 6.3 TensorRT

- TensorRT is NVIDIA's SDK for high-performance inference on NVIDIA GPUs:
  layer fusion, precision calibration (FP16/INT8/FP8), kernel autotuning,
  and engine serialization [secondary].
- The LLM-era path is TensorRT-LLM (separate library), which adds
  paged KV-cache, in-flight batching, and attention plugins for
  transformer inference [secondary].
- No verified 2026 TensorRT/TensorRT-LLM version was captured in this
  pass — gap G-08 [unverified].
- Developer relationship to the stack: PyTorch/ONNX export in, TensorRT
  engine out; the main cost is engine build time and GPU-architecture
  specificity (engines are generally not portable across GPU generations)
  [secondary].

### 6.4 OpenVINO

- OpenVINO is Intel's open-source toolkit for optimizing and deploying
  models on Intel CPUs, iGPUs, NPUs, and discrete GPUs [secondary].
- The Keras 3 README pins **OpenVINO 2026.2.0** as the minimum supported
  version for its inference-only backend, confirming a 2026.2 release
  line exists [official].
- No feature-level 2026 notes were captured in this pass — gap G-09
  [unverified].
- Typical developer flow: convert (ONNX/PyTorch/TF) → OpenVINO IR →
  `ov.Model` compile for the target device; relevant for CPU/edge
  inference and for the Intel quantized-LLM path noted in PyTorch 2.8
  [secondary].

### 6.5 Compiler/format decision matrix

| Tool | Input | Output | Sweet spot |
|---|---|---|---|
| Triton 3.x | Python kernels | GPU machine code | Custom ops, fused kernels [secondary] |
| ONNX Runtime 1.30 | ONNX graphs | Optimized inference | Cross-platform serving [official] |
| TensorRT(-LLM) | ONNX/PyTorch | NVIDIA engines | Max NVIDIA throughput [secondary] |
| OpenVINO 2026.2 | ONNX/TF/PyTorch | Intel IR | Intel CPU/iGPU/NPU [official for version, secondary for flow] |

---

