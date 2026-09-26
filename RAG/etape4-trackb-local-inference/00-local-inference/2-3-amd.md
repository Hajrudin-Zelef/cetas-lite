---
id: etape4-trackb-local-inference/00-local-inference/2-3-amd
title: "2.3 AMD"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "Intel", "Nvidia", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-09-07"]
keywords: ["amd", "accelerator", "benchmarks", "claude", "compute", "consumer", "decode", "fine-tuning", "fp4", "gpu", "gpus", "inference"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [797, 840]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 87527b41565fe620d19595e97f2b60eeb328e04b1f79098293becabd353ba467
---

# 2.3 AMD

### 2.3 AMD

**Radeon RX 9070 XT** (RDNA4, gfx1201):
- 16 GB GDDR6, 640 GB/s bandwidth, MSRP $599; street from ~$740 (early Aug 2026). Board power 304 W. [independent] (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- Local inference: ROCm backend on Linux (gfx1150/gfx1201 target, ROCm 6.3+; ROCm 7 in newer coverage) delivers 35–42 tok/s on 8B models; Vulkan backend works cross-platform (~15% slower); llama.cpp's official ROCm Docker image natively supports RDNA4 (gfx1201). Community run: Qwen3.6-27B IQ3_M at 46 tok/s decode with MTP speculative decoding (llama.cpp + ROCm Docker, 16k context, 8-bit KV). [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://midwest.social/post/48528597)
- vs RTX 5060 Ti 16GB ($429 MSRP): 9070 XT's 43% bandwidth advantage shows up ~linearly in tok/s (~52 vs ~33 tok/s on 14B Q4); the 5060 Ti buys you CUDA (ExLlamaV2, TensorRT-LLM, mainstream fine-tuning) and 124 W lower power. Verdict from independent coverage: choose 9070 XT for Linux/fastest-16GB-inference or gaming; choose 5060 Ti for the CUDA ecosystem. [independent] (https://localaimaster.com/blog/rx-9070-xt-local-ai)
- No FP4 support on any AMD consumer GPU (NVIDIA-exclusive). [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026)
- Easiest path for AMD users: LM Studio with Vulkan backend (Windows & Linux) — no ROCm setup; ROCm only to squeeze the last bit of speed. [independent] (https://github.com/bokiko/localist/blob/HEAD/guides/amd-gpu.md)
- 24 GB alternative: RX 7900 XTX — runs 32B-class models a 16 GB card cannot hold; the community "spend more, get 24 GB" option. [secondary] (https://localaimaster.com/blog/rx-9070-xt-local-ai)

**Ryzen AI Max+ 395 "Strix Halo"** (16× Zen 5, 40-CU Radeon 8060S iGPU / gfx1151, up to 128 GB LPDDR5x-8000 unified):
- The Windows answer to Apple Silicon for local LLM: 128 GB unified memory with ~256 GB/s theoretical (~215 GB/s measured) bus. Massive capacity, modest bandwidth — token generation is bandwidth-bound, so expect single-digit tok/s on dense 70B, but huge models fit: MoE models fly (Qwen3 30B MoE ~66–72 tok/s; LFM2-24B-A2B ~109 tok/s; ~100 tok/s on 30B MoE per setup guides). 70B (Q4) ~15–20 tok/s (ROCm/rocWMMA, fits entirely in memory). [independent] (https://github.com/nelisw/framework-strix-halo-llm-setup), (https://codersera.com/blog/amd-strix-halo-ryzen-ai-max-local-llm-setup-2026/), (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/)
- Practical setup (2026): allocate ~96 GB to iGPU via AMD Adrenalin (Windows) or GTT kernel parameter + small BIOS framebuffer (Linux, kernel ≥ 6.16.9/6.18); run LM Studio / Ollama / llama.cpp on Vulkan backend (most reliable as of mid-2026; ROCm HIP path improving but had stability issues; community benchmarks show Vulkan outperforming HIP). [independent] (https://codersera.com/blog/amd-strix-halo-ryzen-ai-max-local-llm-setup-2026/), (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/)
- Price points: Framework Desktop ~$1,999 (128 GB), GMKtec mini PCs ~$1,499–$2,500, AMD Ryzen AI Halo dev platform $3,999; ASRock Industrial AI BOX-A395 announced Mar 17, 2026 (200×100×232 mm; pricing unannounced). [secondary] (https://digitalarchitects.hr/insights/amd-ryzen-ai-max-395-local-llm/), (https://videocardz.com/newz/asrock-industrial-launches-ai-box-a395-with-ryzen-ai-max-395-and-128gb-lpddr5x)
- Price delta vs DGX Spark: Strix Halo platforms generally undercut Spark by $1,000–$2,000 at comparable 128 GB. [secondary] (http://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf)

**ROCm for local inference (2026 status):** supported on RDNA4 (gfx1200/gfx1201) and Strix Halo (gfx1151); official llama.cpp ROCm Docker images; LM Studio ships ROCm builds. Maturing but still behind CUDA in ecosystem (no Unsloth, no TensorRT-LLM equivalent); fine-tuning on ROCm via PyTorch works but slower than CUDA. [independent] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026), (https://localaimaster.com/blog/rx-9070-xt-local-ai)

### 2.4 Intel

**Arc B580** (Battlemage Xe2, 12 GB, $249 MSRP):
- 12 GB at $249 — but with a "software tax". Three practical paths [independent] (https://runaihome.com/blog/intel-arc-b580-local-ai-2026/):
  1. llama.cpp **Vulkan backend** — no Intel-specific tooling, works on Windows/Linux with stock builds: Llama 3.1 8B Q4_K_M ~40–42 tok/s (parity with RTX 3060 12 GB).
  2. **IPEX-LLM Portable ZIP** (Windows, zero-install) — SYCL binary via patched Ollama fork; counterintuitively *slower* than Vulkan (IPEX-LLM issue #12991); Qwen2.5 14B Q4_K_M ~15–20 tok/s.
  3. **IPEX-LLM native install on Linux** (full oneAPI stack) — Qwen2.5 14B Q4_K_M 32–38 tok/s.
- Flag: IPEX-LLM was **archived Jan 28, 2026** for "known security issues" (read-only) per an engine-status research doc — the recommended path is shifting to upstream llama.cpp SYCL and Intel's llm-scaler-vllm fork. [secondary/unverified — from a community research doc; treat with caution] (https://github.com/chriscorbell/llm-server/blob/HEAD/docs/research/2026-09-07-engines-on-battlemage.md)
- Community benchmarks on Battlemage (Aug–Sep 2026): llama.cpp SYCL is the fastest llama.cpp path for dense models on B70/B580-class hardware; Vulkan the most robust; vLLM-XPU (Intel fork) best measured numbers with patches; SGLang-XPU "not ready". [independent/community] (same)

**Intel Core Ultra NPUs:** OpenVINO is the primary local path (Intel CPUs, GPUs, NPUs); llama.cpp has an OpenVINO backend "in progress". NPU LLM inference remains marginal vs iGPU/CPU paths. See §2.6.

### 2.5 Apple Silicon

**Unified memory advantage:** the entire model + KV cache lives in one pool (CPU/GPU/NPU coherent), so 64–128 GB+ Macs run 70B-class models that no single consumer dGPU can hold. This is the same "capacity over bandwidth" trade as DGX Spark/Strix Halo — and the reason Apple machines remain the default local-LLM developer machine. [independent] (https://medium.com/@nishilbhave/local-llms-in-2026-which-runtime-to-run-and-the-hardware-you-need-a88450dece2e)

**M4 family (2026 status):** M4 Max 128 GB — community MLX benchmarks (vllm-mlx project): Qwen3-0.6B-8bit 417.9 tok/s, Llama-3.2-3B-4bit 205.6 tok/s, Qwen3-30B-A3B-4bit 127.7 tok/s (~18 GB) — single-stream greedy decode. [independent] (https://github.com/anisoptera/vllm-mlx-upstream)

**M5 family (announced Oct 2025 per coverage; shipping in MacBook Pro/iPad Pro/Vision Pro):**
- Specs per press coverage: 3rd-gen 3 nm; 10-core GPU with a dedicated Neural Accelerator per core (>4× peak GPU AI compute vs M4); 16-core Neural Engine; unified memory bandwidth 153 GB/s (base M5), up to 32 GB config; M5 Max: up to 128 GB, 614 GB/s, 40 GPU cores / 40 Neural Accelerators; M5 Pro: 64 GB, 307 GB/s. [secondary] (https://tecknexus.com/apple-m5-the-next-leap-in-on‑device-ai/), (https://www.banandre.com/blog/apple-m5-max-neural-engine-local-llm-inference-game-changer)
- Apple claims (Machine Learning Research blog, MLX benchmarks): up to 4× LLM prompt-processing on M5 Max vs M4 Max; M5 Pro 3.9× vs M4 Pro; up to ~6.9× vs M1 Pro. [vendor-reported] (same)
- LM Studio/Bionic 1.1.0 notes "2–2.75× faster prompt processing on M5 Macs"; Bionic 1.1.5 adds "ultra fast Qwen3.8 inference" via **Splash**, a new Apple-Silicon inference engine by Inco AI. [official]
- **MLX ecosystem (2026):** Apple's MLX framework is the native path (Metal kernels, unified memory, no model conversion); vLLM-for-MLX projects add continuous batching, multimodal, MCP tool calling, Claude Code support; `lms runtime update mlx` keeps LM Studio's MLX engine current; LM Studio also exposes MLX prompt disk-cache controls. [independent/official]

### 2.6 NPU / edge (real status 2026)

