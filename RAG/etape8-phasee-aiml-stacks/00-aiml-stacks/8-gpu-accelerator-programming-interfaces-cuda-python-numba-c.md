---
id: etape8-phasee-aiml-stacks/00-aiml-stacks/8-gpu-accelerator-programming-interfaces-cuda-python-numba-c
title: "8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX"
domain: step-8-phase-e-ai-ml-software-stacks-developer-facing
role: deep-dive
task: hardware
actors: ["AMD", "Apple", "Nvidia"]
dates: ["2026-08-29"]
keywords: ["accelerator", "gpu", "amd", "decode", "fine-tuning", "gpus", "memory", "nvidia", "prefill", "pricing", "tpu", "training"]
source: docs/RAG/etape8_phaseE_aiml_stacks.md
source_anchor: ""
source_lines: [535, 607]
section: "Step 8, Phase E — AI/ML Software Stacks (Developer-Facing)"
sha256: e089f21503c73bfd21f1f7824d21b99167d4d9a27f4ae44deef49444ae977b1d
---

# 8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX

## 8. GPU/accelerator programming interfaces: CUDA Python, Numba, CuPy, ROCm, MLX

### 8.1 CUDA Python

- `cuda-python` is NVIDIA's official Python binding layer for CUDA
  (driver/runtime APIs, cooperative groups, graphs), letting Python code
  drive CUDA directly without C extensions [secondary].
- No verified 2026 package version was captured in this pass — gap G-11
  [unverified].
- It sits below Numba/CuPy/PyTorch in abstraction: use it for driver-level
  control, custom memory management, and interop glue [secondary].

### 8.2 Numba

- Numba is the JIT compiler that turns annotated Python/NumPy code into
  machine code, including a CUDA target (`@numba.cuda.jit`) for writing
  GPU kernels in Python [secondary].
- Its role in 2026 ML work: quick custom kernels and data-pipeline
  acceleration without leaving Python; Triton has absorbed much of the
  "custom deep-learning kernel" demand [secondary].
- No verified 2026 release version captured — gap G-12 [unverified].

### 8.3 CuPy

- CuPy is the NumPy/SciPy-compatible array library for NVIDIA (and AMD
  via ROCm builds) GPUs, with custom CUDA kernel support
  (`cupy.RawKernel`) [secondary].
- Its role: GPU array computing for preprocessing, scientific workloads,
  and glue between data pipelines and DL frameworks [secondary].
- No verified 2026 release version captured — gap G-13 [unverified].

### 8.4 ROCm PyTorch builds

- AMD's ROCm stack ships PyTorch builds for Instinct accelerators; the
  PyTorch 2.8 notes cite **ROCm 7 with `gfx950`** support, confirming the
  2026 ROCm 7 line [secondary].
- The AMD path trades CUDA's ecosystem maturity for open drivers and
  competitive Instinct pricing; framework-level code (PyTorch/JAX) is
  increasingly portable, while custom CUDA kernels are not [secondary].
- CDNA kernel-path lag in upstream Triton (Section 6.1) remains the
  practical friction point for custom-kernel work on AMD [secondary].

### 8.5 Apple MLX

- MLX is Apple's open-source array framework for Apple Silicon, with a
  NumPy-like Python API and a lazy-computation model [secondary].
- A reproducibility record pins **MLX 0.32.2**, `mlx-metal` 0.32.2, and
  `mlx-lm` 0.31.3 on August 29, 2026, confirming the 0.32/0.31 lines in
  the window [secondary].
- WWDC 2026 coverage reports **Metal 4 support** and **multi-Mac training
  over Thunderbolt (RDMA)** for MLX [secondary].
- Secondary reports also describe a new "Core AI" Swift framework and
  Foundation Models framework support for MLX language-model backends —
  not verified against Apple sources; treat as [unverified].
- `mlx-lm` is the LLM toolkit on top of MLX (conversion, fine-tuning,
  serving); a March-2026 report describes Ollama 0.19 adding an MLX
  backend for Apple Silicon with large prefill/decode gains on M5 Max —
  [secondary], cross-referenced with Step 4's Ollama coverage.

### 8.6 Accelerator-interface decision matrix

| Interface | Abstraction level | Hardware | Best fit |
|---|---|---|---|
| CUDA Python | Driver/runtime API | NVIDIA | Driver-level control, interop [secondary] |
| Numba CUDA | JIT Python kernels | NVIDIA | Quick custom kernels [secondary] |
| CuPy | NumPy-like arrays | NVIDIA/AMD | GPU array computing [secondary] |
| Triton 3.x | Tiled kernels | NVIDIA/AMD | DL custom ops (see §6.1) [secondary] |
| ROCm PyTorch | Framework | AMD Instinct | AMD-native training [secondary] |
| MLX 0.32.x | NumPy-like, lazy | Apple Silicon | Local Mac LLM/training [secondary] |
| JAX/XLA | Functional transforms | TPU/GPU | TPU training (see §2.3) [secondary] |

---

