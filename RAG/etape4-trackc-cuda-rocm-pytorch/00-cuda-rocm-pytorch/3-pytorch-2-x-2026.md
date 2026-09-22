---
id: etape4-trackc-cuda-rocm-pytorch/00-cuda-rocm-pytorch/3-pytorch-2-x-2026
title: "3. PyTorch 2.x (2026)"
domain: step-4-track-c-cuda-rocm-pytorch-gpu-compute-stack
role: deep-dive
task: funding-deals
actors: ["AMD", "AWS", "Apple", "Intel", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-04-30", "2026-05", "2026-06", "2026-07-20", "2026-09", "2026-09-02", "2026-09-22", "2026-10"]
keywords: ["accelerator", "agent", "agents", "agi", "amd", "attention", "aws", "blackwell", "compute", "fp4", "fp8", "gemini"]
source: docs/RAG/etape4_trackC_cuda_rocm_pytorch.md
source_anchor: ""
source_lines: [134, 226]
section: "Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)"
sha256: c45133b7154119a2b04d654a68a57d85e35c9240e947d91e3e68a74addcc3fc7
---

# 3. PyTorch 2.x (2026)

## 3. PyTorch 2.x (2026)

### 3.1 Release timeline

| Version | Release date | Notes |
|---|---|---|
| 2.11.0 | ~April 30, 2026 | Enterprise-deployment features, FlashAttention integration for Blackwell/Hopper (per AI-generated summary — [unverified]), XPU Graph support |
| 2.12.0 | ~June 2026 (2.12.1 in DLC Jul 2, 2026) | torch.accelerator.Graph API, batched linalg.eigh up to 100× faster on CUDA, torch.export.save MX quantization, fused Adagrad, torch.cond in CUDA Graphs, ROCm memory/collective gains |
| 2.13.0 | July 20, 2026 | CuTeDSL "Native DSL" Inductor backend, FlexAttention on Apple Silicon (~12×), nn.LinearCrossEntropyLoss (4× peak-memory cut), torchcomms, FSDP2 comm-overlap, Python 3.15 (+ free-threaded 3.15t) wheels, 3,328 commits / 526 contributors |
| 2.14.0 | September 2, 2026 | clamp/min/max boundary-subgradient change (1→0 at boundaries — silent numeric behavior change), bfloat16 complex promotion now raises, cholesky/qr removed; verify behavior by running before/after upgrade |

Sources: PyTorch 2.13/2.12 release Q&A and GitHub release notes **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0, https://github.com/pytorch/pytorch/releases, http://dev.to/tsho/pytorch-214-clamps-boundary-gradient-silently-changes-bfloat16-complex-promotion-now-raises-26g1]**; AWS DLC changelog for dates/CUDA bundling **[secondary — https://github.com/aws/deep-learning-containers/blob/HEAD/docs/pytorch/changelog/index.md]**.

### 3.2 Compilation (torch.compile / Inductor)

- **PyTorch 2.13** introduces the **CuTeDSL "Native DSL" backend** for key GPU operations in Inductor **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- **FlexAttention on Apple Silicon** with up to ~12× speedup over SDPA on sparse patterns **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- **Deterministic backward computation on CUDA** (2.13) **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- **nn.LinearCrossEntropyLoss** reduces peak GPU memory by up to **4×** during large-vocabulary LM training **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- **Device-agnostic `torch.accelerator.Graph` API** (2.12) unifies graph capture and replay across CUDA, XPU, and out-of-tree backends **[official via secondary — https://github.com/pytorch/pytorch/releases]**.
- `torch.cond` control flow can now be captured and replayed inside CUDA Graphs (2.12) **[official via secondary — https://github.com/pytorch/pytorch/releases]**.
- Build requirements tightened in 2.12: **minimum CUDA 12.6** to build from source, **C++20** minimum enforced **[official via secondary — https://github.com/meta-pytorch/torch-release-notes/blob/HEAD/2.12.0/final.md]**.
- Triton pinned to **3.7.1** in 2.12.x to fix a nondeterministic FLASH_ATTN output regression and an illegal-memory-access in the Triton `convolution2d_bwd_weight` kernel on B100/B200 (sm100) GPUs **[secondary — https://github.com/pytorch/pytorch/releases]**.
- PyTorch **2.6** (Jan 2026) deprecated the official **Anaconda channel**; **2.12.1** dropped CPython **3.13t** from the binary build matrix; **2.13** adds wheels for **Python 3.15** including free-threaded 3.15t builds on Linux **[secondary — https://ngontinh24.com/article/pytorch-2-6-release-blog-pytorch, https://github.com/pytorch/pytorch/releases, https://www.youtube.com/watch?v=vbdduGRTQC0]**.

### 3.3 Distributed training

- **torchcomms** for large-cluster training (2.13) **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- **FSDP2 communication-overlap improvements** (2.13) **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- NCCL bumps: NCCL **2.26.2** in PyTorch 2.11/2.12 images; NCCL **2.30.7** + EFA 1.49.0 in PyTorch 2.13 DLC images for improved multi-node collective performance **[secondary — https://github.com/aws/deep-learning-containers/blob/HEAD/docs/pytorch/changelog/index.md]**.
- ROCm side: rocSHMEM symmetric-memory collectives (2.12); RCCL in the ROCm stack **[secondary]**.

### 3.4 Quantization and torchao

- `torch.export.save` supports **microscaling (MX) quantization formats** (2.12), enabling full export of aggressively compressed models — relevant to FP4/FP6/FP8 on Blackwell and MI350X **[official via secondary — https://github.com/pytorch/pytorch/releases]**.
- ⚠️ **Gap:** no 2026-specific torchao releases were captured in this research window; torchao remains the quantization toolkit used by vLLM/SGLang serving stacks, but current version numbers are unverified — re-check the torchao GitHub releases before citing.

### 3.5 CUDA architecture support in PyTorch 2.12–2.14 wheels (official support matrix)

From AMD's rocm/pytorch RELEASE.md matrix (current as of September 2026) **[secondary — https://github.com/rocm/pytorch/blob/HEAD/RELEASE.md]**:

| CUDA build | Linux x86 + Windows arch coverage |
|---|---|
| 12.6.3 | Maxwell (5.0), Pascal (6.0), Volta (7.0), Turing (7.5), Ampere (8.0/8.6), Hopper (9.0) |
| 13.0.2 | Turing (7.5), Ampere (8.0/8.6), Hopper (9.0), Blackwell (10.0, 12.0+PTX; Linux-only +PTX) |
| 13.2.1 | Turing (7.5), Ampere (8.0/8.6), Hopper (9.0), Blackwell (10.0, 12.0+PTX; Linux-only +PTX) |

Linux aarch64 adds Blackwell 11.0 on the 13.0.2/13.2.1 lines. Note: Blackwell compute capabilities (sm_100/sm_120) were first introduced in **CUDA 12.8**; nvcc 12.6 cannot emit them **[secondary — https://github.com/mesh-llm/mesh-llm/blob/HEAD/docs/cuda-release-lanes.md]** — relevant for local-inference builds (llama.cpp etc.) targeting RTX 50-series cards.

---

## 4. Triton compiler (2026)

- Triton ships as **v3.3 in ROCm 7.0** with a **first-class AMD backend**: CDNA-specific GPU-IR layouts (`amd_mfma` matrix core, `amd_wmma`, blocked/shared/sliced), tunables `matrix_instr_nonkdim` (16→mfma_16x16, usually winning on MI300X; 32→mfma_32x32), `waves_per_eu`, `kpack`; native FNUZ FP8 (`fp8e4b8`/`fp8e5b16`) on CDNA3; XF32 on gfx942; dedicated `ROCm/triton` fork plus `occ.sh` occupancy tooling; FlashAttention fallback path on AMD via `FLASH_ATTENTION_TRITON_AMD_ENABLE=TRUE`; MIT license **[secondary — https://github.com/amd-agi/geak/blob/HEAD/perf_knowledge/landscape/authoring_dsls.md]**.
- PyTorch 2.12.x pins Triton **3.7.1** to fix sm100 (B100/B200) correctness regressions **[secondary — https://github.com/pytorch/pytorch/releases]**.
- Portability caveat (Spheron, September 2026): identical Triton attention kernels run on both CUDA/H100 and ROCm/MI300X, but **tuning does not transfer** — AMD's fastest config uses smaller BLOCK_M (16), AMD tunes `waves_per_eu` where NVIDIA tunes warp specialization; **HIP graphs cut MI300-family latency ~2×** on short sequences vs ~6% from CUDA graphs on H100; hand-tuned attention kernels still beat Triton's AMD backend by ~1.53× (0.36ms vs 0.55ms) on MI300X in TileLang comparisons **[secondary — https://www.spheron.network/blog/triton-on-amd-rocm-vs-nvidia-cuda-same-kernel-different-perf/]**.
- AMD's **GEAK** Triton agents use the AMD layout/tunable vocabulary as agent search dimensions **[secondary — https://github.com/amd-agi/geak/blob/HEAD/perf_knowledge/landscape/authoring_dsls.md]**.
- Intel also maintains an XPU backend for Triton (out-of-tree) **[secondary — https://github.com/intel/intel-xpu-backend-for-triton/blob/HEAD/RELEASE.md]**.

## 5. FlashAttention

- **flash-attn 2.8.3** is the pinned version in PyTorch 2.11/2.12/2.13 AWS DLC images (fused attention) **[secondary — https://github.com/aws/deep-learning-containers/blob/HEAD/docs/pytorch/changelog/index.md]**.
- A "FlashAttention-4 integration for Blackwell/Hopper" claim appears in an AI-generated PyTorch 2.11 research summary **[unverified — https://github.com/edsonesf/atu-csd-pokedex/blob/HEAD/academic/researches/PyTorch/gemini-pytorch.md]** — the official flash-attn package numbering (2.8.3) and Dao-AILab's FlashAttention-3 (2024) were the last independently verified milestones. **Treat "FlashAttention-4" as unverified until checked against dao-ailab/flash-attention releases.**
- cuDNN-side attention kernels (SDPA fprop now open-sourced via cuDNN Frontend, DSA, block-sparse attention) increasingly overlap FlashAttention's niche on Blackwell — see §1.3 **[secondary — https://github.com/NVIDIA/cudnn-frontend/releases/tag/v1.26.0]**.

## 6. AMD vs NVIDIA for local/enterprise inference (2026 snapshot)

- **NVIDIA local:** RTX 50-series (Blackwell, sm_120) requires CUDA 12.8+ toolkits and drivers; **RTX Spark** (Grace CPU + Blackwell RTX, 6,144 or 5,120 CUDA cores, up to 128 GB unified LPDDR5X) arrives **October 2026**, with CUDA 13.4 GA providing the Windows-on-Arm toolkit ahead of launch **[secondary — https://videocardz.com/newz/cuda-13-4-released-ready-for-rtx-spark]**.
- **AMD local:** ROCm 10.0.0 supports Radeon graphics and Ryzen iGPUs on Windows; ROCm 7.2 added RDNA4 support; Windows ROCm is positioned for quick inference/experimentation, with Ubuntu 24.04 recommended for the full toolchain **[secondary — https://wccftech.com/amd-rocm-10-big-ai-updates-performance-gains/, https://www.spheron.network/blog/cuda-news-today-nvidia-toolkit-amd-rocm-ai-frameworks-2026/]**.
- **Serving stacks:** ROCm ships **vLLM and SGLang containers** and Python wheels as first-class artifacts (both ROCm 10 and the 7.x line); AMD touts distributed inference via vLLM-d and DeepEP compatibility **[secondary — https://www.thelec.net/news/articleView.html?idxno=13453, https://www.techpowerup.com/341074/amd-launches-rocm-7-0-up-to-3-8x-performance-uplift-over-rocm-6-0]**.
- **Enterprise:** NVIDIA's **AI Enterprise** platform (per earlier step-3 research, ~$4,500/GPU/yr list) bundles the CUDA/NIM stack with enterprise support — treat pricing as [secondary/unverified] pending re-check. AMD offers ROCm enterprise support via its partner/OEM ecosystem; no 2026-specific pricing or program changes were captured in this window — **flag as gap**.

## 7. Open verification items

1. Exact GA dates for CUDA 13.0, 13.1, 13.2, 13.3 (only 13.4 verified: dev preview Jul 19, 2026; GA ~Sep 10, 2026).
2. ROCm 7.x sequence conflict: Spheron tracker lists 7.2.3 (May 2026) as latest 7.x; community docs reference **7.14.0** completing the TheRock transition. One or both may be mislabeled (7.2.x patches vs an internal numbering) — needs the official ROCm release-notes page.
3. ROCm 10.0.0 exact date: Aug 26 vs Aug 28, 2026 across sources.
4. cuDNN 9.25.0/9.26.0 detailed release notes (frontend v1.29.0 "recommended for 9.26" implies it exists, but backend notes not fetched).
5. "FlashAttention-4" — unverified; check dao-ailab/flash-attention.
6. PyTorch 2.11.0 official release date and full highlights (currently from AI-generated summary).
7. torchao current version and 2026 changes — not captured.
8. HIPify 2026 improvements — not captured.
9. NVIDIA AI Enterprise 2026 pricing and AMD ROCm enterprise support program details.
10. vLLM / SGLang version-specific ROCm-CUDA parity measurements (independent, not vendor claims).
11. torchcomms details (topology support, backends) beyond the 2.13 release headline.
12. CUDA 14.0 release date — only the Armv8.2-A minimum note is confirmed.

## 8. Collection metadata

- **Collected:** September 22, 2026.
- **Primary sources used:** NVIDIA CUDA Toolkit 13.4/13.4-Update-1 release notes (docs.nvidia.com) [official]; cuDNN Frontend GitHub releases [secondary/official repo]; PyTorch GitHub releases + meta-pytorch release-notes worksheets [secondary/official repo]; rocm/pytorch RELEASE.md support matrix [secondary/official repo]; AMD community ROCm docs (datawhalechina/hello-rocm, damustermann/rocm-wsl-ai) [secondary].
- **Press used:** VideoCardz, Wccftech, TechPowerUp, TheElec, Spheron blog, TensorOpera/dev.to community notes [secondary/independent where noted].
- **Notes:** Dates for press-covered releases reflect article publication dates; official docs may list slightly different GA dates. Performance figures from AMD/NVIDIA are vendor-reported. "Unverified" items are explicitly tagged in §7.
