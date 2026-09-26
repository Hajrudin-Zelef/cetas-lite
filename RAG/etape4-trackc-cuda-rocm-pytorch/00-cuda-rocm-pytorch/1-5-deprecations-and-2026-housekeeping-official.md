---
id: etape4-trackc-cuda-rocm-pytorch/00-cuda-rocm-pytorch/1-5-deprecations-and-2026-housekeeping-official
title: "1.5 Deprecations and 2026 housekeeping (official)"
domain: step-4-track-c-cuda-rocm-pytorch-gpu-compute-stack
role: deep-dive
task: funding-deals
actors: []
dates: ["2026-05-26", "2026-06", "2026-07", "2026-07-20"]
keywords: ["aws", "llama", "llama.cpp", "nvidia", "research"]
source: docs/RAG/etape4_trackC_cuda_rocm_pytorch.md
source_anchor: ""
source_lines: [78, 90]
section: "Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)"
sha256: 96018e8336828e15d68c87eb65dea4718bbceafa67b1dfcf9e896b9f4f932a09
---

# 1.5 Deprecations and 2026 housekeeping (official)

- Minor-version compatibility: existing CUDA 13.x applications run on drivers **≥ 580**; CUDA 13.4 new features/newly enabled platforms require an **R615+ driver** (Windows driver for RTX Spark devices: **616.41 or later**). The driver is **no longer bundled** with the toolkit (Windows since CUDA 13.1, Linux since CUDA 13.4) **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- Community corroboration: llama.cpp CUDA 13 builds require the documented 581.15 minimum driver for CUDA 13.1 **[secondary — https://github.com/atomicbot-ai/atomic-chat/blob/HEAD/docs/decisions/2026-05-26-correct-cuda-13-1-driver-gate-to-nvidia-documented-581-15-and.md]**.

### 1.5 Deprecations and 2026 housekeeping (official)

- **Python 3.10 deprecated** across CUDA Python 13.4 packages **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- Legacy **Nsight Eclipse Edition** plugins no longer delivered in toolkit packages beginning with CUDA 13.3 **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.
- **CUDA 14.0** will move to **Armv8.2-A** as the minimum supported architecture for ARM64-SBSA **[official — https://docs.nvidia.com/cuda/cuda-toolkit-release-notes/index.html]**.

⚠️ **Verification gap:** exact release dates for CUDA 13.0, 13.1, 13.2, and 13.3 were not confirmed from official notes in this research. Bundling evidence: PyTorch 2.12 wheels (July 2026) shipped CUDA 13.0.2, and PyTorch 2.13 wheels (July 20, 2026) shipped CUDA 13.3.0, so CUDA 13.0 GA'd by ~June 2026 and 13.3 by mid-July 2026 **[secondary — https://github.com/aws/deep-learning-containers/blob/HEAD/docs/pytorch/changelog/index.md]**.

---

