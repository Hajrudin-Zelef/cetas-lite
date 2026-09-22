---
id: etape4-trackc-cuda-rocm-pytorch/00-cuda-rocm-pytorch/2-amd-rocm
title: "2. AMD ROCm"
domain: step-4-track-c-cuda-rocm-pytorch-gpu-compute-stack
role: deep-dive
task: funding-deals
actors: ["AMD", "Anthropic", "DeepSeek", "Moonshot", "SGLang", "Z.ai", "vLLM"]
dates: ["2016-04", "2026-03", "2026-05", "2026-07-07"]
keywords: ["amd", "agent", "agents", "attention", "claude", "deepseek", "embedding", "fp4", "fp8", "glm", "gpu", "gpus"]
source: docs/RAG/etape4_trackC_cuda_rocm_pytorch.md
source_anchor: ""
source_lines: [91, 133]
section: "Step 4 — Track C: CUDA + ROCm + PyTorch (GPU compute stack)"
sha256: 08ed287c3039049c7707b58c5e1c24da3ad7a3780b05d28e052751305f89ad41
---

# 2. AMD ROCm

## 2. AMD ROCm

### 2.1 Latest release: ROCm 10.0.0 (~August 26–28, 2026)

AMD released **ROCm 10.0.0 (ROCm Core SDK 10.0.0)** — the first major version jump after the 7.x series, announced around August 26–28, 2026, marking 10 years since ROCm 1.0 (April 2016) **[secondary — https://wccftech.com/amd-rocm-10-big-ai-updates-performance-gains/, https://www.thelec.net/news/articleView.html?idxno=13453, https://github.com/datawhalechina/hello-rocm/blob/HEAD/docs/en/00-environment/rocm-10-0-0-release-notes.md]**.

Headline changes:
- Built on **[TheRock](https://github.com/ROCm/TheRock)** as the production build and release foundation (the 7.x line completed the move from a monolithic bundle to a modular Core SDK) **[secondary — https://github.com/datawhalechina/hello-rocm/blob/HEAD/docs/en/00-environment/rocm-10-0-0-release-notes.md]**.
- **ROCm.AI** — an AI-assistant-native developer workflow folding install, coding, and performance work into three entry points:
  - **AMD Skills** — official AMD optimization knowledge packaged as skills for AI coding agents (Claude Code, Cursor, Codex).
  - **Hyperloom** (rendered "ROCm HyperRoom" in press) — open-source auto-optimization engine: profile, find bottlenecks, rewrite kernels, tune parameters, validate; uses an AI agent for GPUs to optimize inference workloads **[secondary — https://www.thelec.net/news/articleView.html?idxno=13453]**.
  - **ROCm CLI & Console** — one command surface for install/verify/deploy/manage with real-time monitoring through a conversational interface **[secondary — https://www.thelec.net/news/articleView.html?idxno=13453]**.
- Support spans Instinct accelerators, Radeon graphics, and Ryzen integrated graphics on **both Windows and Linux** **[secondary — https://wccftech.com/amd-rocm-10-big-ai-updates-performance-gains/]**.
- Performance claims (AMD, **vendor-reported**): on the same hardware, **3.3× inference performance** and **2.4× training performance** vs ROCm 7. Testing by AMD Performance Labs as of July 7, 2026, on an 8× MI355X platform comparing ROCm 7.0 against a preview ROCm.AI build (ROCm 7.2.2 with optimized kernels/parallelism/scheduling), running GLM-5, Kimi-K2.5, and DeepSeek-R1-0528 **[secondary — https://wccftech.com/amd-rocm-10-big-ai-updates-performance-gains/]**.
- Packaging: ships **vLLM and SGLang containers**, Python wheels, and modular software packages; the pip wheel index moved from `repo.amd.com/rocm/whl-multi-arch/` to `https://stable.repo.amd.com/rocm/whl-next/`; apt/dnf repos moved to stable.repo.amd.com **[secondary — https://www.thelec.net/news/articleView.html?idxno=13453, https://github.com/datawhalechina/hello-rocm/blob/HEAD/docs/en/00-environment/index.md]**.

### 2.2 ROCm 7.x track (2025–mid-2026)

| Version | Date | Key changes | Provenance |
|---|---|---|---|
| ROCm 7.0 | Late 2025 | MI350X/MI355X GA support (CDNA 4), HIP compiler major update, MX data types (FP4/FP6/FP8) on MI350X, up to 3.8× perf vs ROCm 6.0 (vendor-reported), vLLM-d/DeepEP distributed inference, GPU-direct communications | [independent — https://www.techpowerup.com/341074/amd-launches-rocm-7-0-up-to-3-8x-performance-uplift-over-rocm-6-0] |
| ROCm 7.1 | Early 2026 | Deeper MI350X/MI355X optimization, improved FlashAttention, HIP Python improvements | [secondary — https://www.spheron.network/blog/cuda-news-today-nvidia-toolkit-amd-rocm-ai-frameworks-2026/] |
| ROCm 7.2 | ~March 2026 | RDNA4 support, Ubuntu 24.04.x support, RCCL improvements; Triton v3.3 shipped with first-class AMD backend | [secondary — https://www.spheron.network/blog/cuda-news-today-nvidia-toolkit-amd-rocm-ai-frameworks-2026/, https://github.com/damustermann/rocm-wsl-ai/blob/HEAD/CHANGELOG.md] |
| ROCm 7.2.3 | May 2026 | vLLM profiling stability, embedding-inference optimization | [secondary — https://www.spheron.network/blog/cuda-news-today-nvidia-toolkit-amd-rocm-ai-frameworks-2026/] |
| ROCm 7.14.0 (?) | 2026 | Reportedly completed the TheRock production transition | **[unverified — conflicting with the 7.2.x line; see §5]** |

GPU support status: MI300X **stable** since ROCm 7.0; MI350X/MI355X **stable** since ROCm 7.0; **MI400 in developer preview**; minimum kernel 6.8 for the 7.x line **[secondary — https://www.spheron.network/blog/cuda-news-today-nvidia-toolkit-amd-rocm-ai-frameworks-2026/]**.

### 2.3 PyTorch/JAX on ROCm

- Official AMD PyTorch wheels are published from repo.radeon.com/repo.amd.com; e.g., **PyTorch 2.9.1** official AMD wheels (March 2026 community docs) **[secondary — https://github.com/damustermann/rocm-wsl-ai/blob/HEAD/CHANGELOG.md]**.
- PyTorch 2.12's official release notes list ROCm platform gains: **expandable memory segments**, **rocSHMEM symmetric-memory collectives**, and **FlexAttention pipelining** on ROCm **[official via secondary — https://github.com/pytorch/pytorch/releases]**.
- PyTorch 2.13 release material lists **expanded ROCm platform support** **[secondary — https://www.youtube.com/watch?v=vbdduGRTQC0]**.
- AMD's ROCm PyTorch support matrix (rocm/pytorch RELEASE.md) tracks stable CUDA/ROCm combinations per PyTorch release **[secondary — https://github.com/rocm/pytorch/blob/HEAD/RELEASE.md]**.
- ROCm Windows note: the full toolchain (rocminfo, amd-smi, multi-GPU, containerized deployment) is best on Ubuntu 24.04; Windows covers quick inference and experimentation **[secondary — https://github.com/datawhalechina/hello-rocm/blob/HEAD/docs/en/00-environment/index.md]**.

### 2.4 HIP / CUDA translation tooling

- AMD's HIPify (hipify-perl, hipify-clang) remains the standard CUDA→HIP source-translation path; no major 2026 HIPify-specific announcements were captured in this research window — flag as a gap.
- Triton is increasingly the portability vehicle: vLLM's Triton attention backend runs identical kernel source on CUDA/H100 and ROCm/MI300-class hardware, though **tuning does not transfer** (different optimal block sizes and scheduling knobs per backend) **[secondary — https://www.spheron.network/blog/triton-on-amd-rocm-vs-nvidia-cuda-same-kernel-different-perf/]**.

---

