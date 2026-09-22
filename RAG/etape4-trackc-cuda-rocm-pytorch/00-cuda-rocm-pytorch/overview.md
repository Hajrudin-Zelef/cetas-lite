---
id: etape4-trackc-cuda-rocm-pytorch/00-cuda-rocm-pytorch/overview
title: "Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)"
domain: step-4-track-c-cuda-rocm-pytorch-gpu-compute-stack
role: deep-dive
task: funding-deals
actors: ["AMD", "DeepSeek", "Nvidia"]
dates: ["2026-02-01", "2026-03-09", "2026-05-26", "2026-06", "2026-07", "2026-07-19", "2026-07-20", "2026-09-10", "2026-09-13", "2026-09-22", "2026-10"]
keywords: ["compute", "gpu", "amd", "attention", "aws", "benchmarks", "blackwell", "deepseek", "fp8", "gpus", "inference", "llama"]
source: docs/RAG/etape4_trackC_cuda_rocm_pytorch.md
source_anchor: ""
source_lines: [1, 90]
section: "Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)"
sha256: 9e38cfc1ef390232bc2324707fa081ef945cb1355c4c01bcfc7c9b2811fe6a79
---

# Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)
## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 4 (Infra inference/training), Track C
**Coverage window:** February 1, 2026 → September 22, 2026
**Topics:** NVIDIA CUDA Toolkit, AMD ROCm, PyTorch 2.x, plus Triton, FlashAttention, HIP/CUDA translation, enterprise support
**Status:** Research snapshot. Version numbers and dates below were current as of the sources cited; re-verify against official pages before use.

### Provenance legend
- **[official]** — vendor's own blog, docs, changelog, release notes, GitHub release.
- **[vendor-reported]** — figure claimed by the vendor (benchmarks, perf claims) without independent audit.
- **[independent]** — reputable third-party press or independent measurement.
- **[secondary]** — lower-tier press, blogs, community docs, AI-generated summaries; useful but treat with caution.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. NVIDIA CUDA Toolkit

### 1.1 Latest release: CUDA 13.4 / 13.4 Update 1 (GA ~September 10, 2026)

NVIDIA released **CUDA Toolkit 13.4.1** (13.4 Update 1) with the CUDA 13.4 platform reaching general availability **[secondary — https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark, https://wccftech.com/nvidia-cuda-13-4-support-windows-on-arm-ahead-of-rtx-spark-launch/]**.

Key highlights of CUDA 13.4:
- **Windows-on-Arm support** — first native Windows Arm64 CUDA development package, timed ahead of RTX Spark PCs (Windows 11 on Arm systems with Grace CPU cores + Blackwell RTX graphics, shipping October 2026). NVCC, CUDA runtime, cuBLAS, cuFFT, and Nsight tools now support Windows Arm64; both on-device compilation and cross-compilation from Windows x86-64 are supported **[secondary — https://videocardz.com/newz/nvidia-releases-first-cuda-toolkit-preview-for-windows-on-arm-and-rtx-spark, https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark]**.
- **Preview support for NVIDIA Rubin (Vera Rubin) GPUs** — the toolkit adds the **SM_107** architecture target, identified as **compute capability 10.7**, allowing developers to compile and begin porting CUDA software before full Rubin support goes GA. NVIDIA states the preview is not intended for production deployment or performance benchmarking **[secondary — https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark, https://wccftech.com/nvidia-cuda-13-4-support-windows-on-arm-ahead-of-rtx-spark-launch/]**.
- **Multi-Process Service (MPS) V3** — modernized control layer with a scriptable CLI, named server instances, TOML configuration, and cgroup-integrated GPU memory limits for GPU partitioning in containerized environments **[secondary — https://wccftech.com/nvidia-cuda-13-4-support-windows-on-arm-ahead-of-rtx-spark-launch/]**.
- **CUDA Compute Fabric Transport** for NVLink-based communication **[secondary — https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark]**.
- Updates across the CUDA compiler, **PTX ISA 9.4**, CUDA C++, CUDA Tile programming, runtime APIs, and libraries **[secondary — https://wccftech.com/nvidia-cuda-13-4-support-windows-on-arm-ahead-of-rtx-spark-launch/]**; official 13.4 Update 1 release notes confirm PTX ISA 9.4 support and CUDA Toolkit Discovery via `CUDA_HOME`/`CUDA_PATH`/`CUDA_ROOT` environment variables (in cuDNN 9.24.0) **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.

Developer preview timeline: CUDA 13.4 Developer Preview was released **July 19, 2026** with the first official Windows Arm64 package **[secondary — https://videocardz.com/newz/nvidia-releases-first-cuda-toolkit-preview-for-windows-on-arm-and-rtx-spark]**; GA followed ~September 10, 2026 **[secondary — https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark]**.

### 1.2 Component versions in CUDA 13.4 Update 1 (official)

From the official release notes, major component versions **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**:

| Component | Version |
|---|---|
| CUDA compiler (nvcc), runtime (cudart), NVRTC, CUPTI, Nsight toolset core | 13.4.92 |
| cuBLAS | 13.8.0.4 |
| cuFFT | 12.4.0.43 |
| cuSPARSE | 12.8.6.72 |
| cuSOLVER | 12.3.4.7 |
| cuRAND | 10.4.4.72 |
| NPP | 13.2.0.58 |
| nvJPEG | 13.2.3.58 |
| cuFile | 1.19.1.55 |
| Thrust / CUB / libcu++ (CCCL) | 3.4.3 |
| Nsight Compute | 2026.3.1.2 |
| Nsight Systems | 2026.3.2.476 |
| Nsight Visual Studio Edition | 2026.3.0.26187 |

cuBLAS highlights in the 13.4 line: **emulated FP64 matrix multiplications** with up to **25% ZGEMM peak-performance improvement on Rubin GPUs**; when per-handle workspace is insufficient, fixed-point FP64 emulation now allocates temporary workspace from a cuBLAS-managed per-device CUDA memory pool retained across stream synchronizations **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.

Known issues (official): cuBLASLt Grouped GEMM with per-batch tensor-wide scales can cause invalid memory access for groups with `m > 0, n > 0, k = 0` (introduced in CUDA 13.1); cuBLASLt Grouped GEMM can produce incorrect results on **Blackwell and Rubin** GPUs (introduced in CUDA 13.4, cuBLAS 13.7.0) **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.

Since **March 9, 2026**, cuBLAS patch releases are published **independently of CUDA Toolkit releases** for critical bug fixes **[official — https://docs.nvidia.com/cuda/archive/13.0.3/pdf/CUDA_Toolkit_Release_Notes.pdf]**.

### 1.3 cuDNN 2026 track

- **cuDNN Backend 9.24.0** (June–July 2026): CUDA Toolkit 13.3 support; SDPA backward `d=256` on Blackwell (2-CTA MMA path, bfloat16/float16); Ubuntu 26.04 support; CUDA Toolkit discovery via environment variables; expanded causal conv1d via SubquadraticOps API with NVRTC runtime compilation **[secondary — https://github.com/nvidia/cudnn-frontend/issues/442]**.
- **cuDNN Frontend v1.29.0** (released ~September 13, 2026) is the recommended frontend for **cuDNN 9.26** and later; it adds HSTU (Hierarchical Sequential Transduction Unit) attention as a complete CuTe DSL kernel family for Blackwell — packed variable-length forward/backward, FP16/BF16, head dims 64/128/256, full/causal/local/arbitrary masks, paged-KV forward — plus optimized qlen=1 causal/local attention on SM100, SM103, and **SM107** (Rubin) **[secondary — https://github.com/NVIDIA/cudnn-frontend/releases/tag/v1.29.0]**.
- **cuDNN Frontend v1.26.0** (July 2026): unified-engine FP8 and MXFP8 forward SDPA (requires cuDNN 9.25.0+); SDPA backward `d=256` native path on cuDNN 9.23+; block-sparse attention CuTe DSL kernels for Hopper and Blackwell; DeepSeek Sparse Attention (DSA) kernels with q-causal offsets and SM100F support; grouped-GEMM quant wrapper for SM100 **[secondary — https://github.com/NVIDIA/cudnn-frontend/releases/tag/v1.26.0]**.
- Supported architectures this period: Blackwell (B200, B300; compute capability 10.0/10.3/12.0), Hopper (H200/H20; 9.0), Ada Lovelace (L40S, RTX PRO 6000; 8.9), Ampere (A100; 8.0) **[secondary — https://github.com/nvidia/cudnn-frontend/issues/442]**.

### 1.4 Driver requirements (official)

Driver-branch mapping **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**:

| CUDA Toolkit | Driver branch |
|---|---|
| CUDA 13.4 | **R615** |
| CUDA 13.3 | R610 |
| CUDA 13.2 | R595 |
| CUDA 13.1 | R590 (Windows minimum 581.15) |
| CUDA 13.0 | R580 |

- Minor-version compatibility: existing CUDA 13.x applications run on drivers **≥ 580**; CUDA 13.4 new features/newly enabled platforms require an **R615+ driver** (Windows driver for RTX Spark devices: **616.41 or later**). The driver is **no longer bundled** with the toolkit (Windows since CUDA 13.1, Linux since CUDA 13.4) **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- Community corroboration: llama.cpp CUDA 13 builds require the documented 581.15 minimum driver for CUDA 13.1 **[secondary — https://github.com/atomicbot-ai/atomic-chat/blob/HEAD/docs/decisions/2026-05-26-correct-cuda-13-1-driver-gate-to-nvidia-documented-581-15-and.md]**.

### 1.5 Deprecations and 2026 housekeeping (official)

- **Python 3.10 deprecated** across CUDA Python 13.4 packages **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- Legacy **Nsight Eclipse Edition** plugins no longer delivered in toolkit packages beginning with CUDA 13.3 **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- **CUDA 14.0** will move to **Armv8.2-A** as the minimum supported architecture for ARM64-SBSA **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.

⚠️ **Verification gap:** exact release dates for CUDA 13.0, 13.1, 13.2, and 13.3 were not confirmed from official notes in this research. Bundling evidence: PyTorch 2.12 wheels (July 2026) shipped CUDA 13.0.2, and PyTorch 2.13 wheels (July 20, 2026) shipped CUDA 13.3.0, so CUDA 13.0 GA'd by ~June 2026 and 13.3 by mid-July 2026 **[secondary — https://github.com/aws/deep-learning-containers/blob/HEAD/docs/pytorch/changelog/index.md]**.

---

