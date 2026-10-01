---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-15
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "Intel", "Microsoft", "Nvidia", "Qualcomm", "SGLang", "vLLM"]
dates: ["2026-07", "2026-09-07"]
keywords: ["inference", "llama", "llama.cpp", "accelerator", "amd", "benchmark", "benchmarks", "blackwell", "claude", "compute", "consumer", "copilot"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [817, 862]
sha256: 20a02329da685c6bfbf904070ba26a3828042f8a0a84c9a5dee63ac3cfda6ed5
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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

The honest 2026 picture: NPUs matter for Copilot+ OS features (Studio Effects, Recall-type workloads) and small on-device assistants — **not** for serious local LLM inference, where GPU/iGPU/CPU still dominate:

- **Qualcomm Snapdragon X Elite (Hexagon NPU, 45 TOPS):** a community benchmark (LM Studio 0.3 + Ollama/DirectML, Win 11 24H2, Llama 3.1 8B Q4_K_M) measured ~26–31 tok/s on Snapdragon X Elite (QNN) vs 19–24 tok/s on Ryzen AI 9 HX 375 (DirectML, partial NPU) vs 13–17 tok/s on Intel Core Ultra 7 258V (OpenVINO, partial) — but TOPS ≠ usable throughput; driver/runtime maturity decides. [secondary — single-blogger benchmark, treat as indicative] (https://medium.com/@neilandrews1983/best-npu-laptops-2026-llama-mistral-benchmark-rankings-for-uk-buyers-c041a1804219)
- **AMD Ryzen AI NPUs (XDNA 2, up to 50 TOPS):** Linux XDNA driver stack matured enough in 2026 to run LLMs on the NPU directly (Phoronix/Michael Larabel testing) via ONNX Runtime + Vitis AI — a milestone, but the `ryzenai` backend in community libraries was still "awaiting hardware validation" as of mid-2026. [secondary] (https://www.webpronews.com/amds-ryzen-ai-npus-can-now-run-llms-locally-on-linux-heres-what-that-means/), (https://github.com/quintindk/crier)
- **Intel Core Ultra NPU:** OpenVINO path; llama.cpp OpenVINO backend still "in progress". [independent] (https://github.com/lunncing/llama.cpp-adaptive-kv-streaming — backend table, via search result)
- **Apple Neural Engine:** used via Core ML / MLX; Apple's on-device AI story is GPU+Neural-Engine combined, not NPU-only. [secondary]
- **Snapdragon X2 generation:** one AI-laptop comparison reports a "Snapdragon X2 Elite Extreme" with 80 TOPS NPU (highest NPU score in its table) but weaker memory bandwidth/GPU for memory-intensive inference. [secondary — single source, unverified] (https://medium.com/@raulasla/the-ai-laptop-chip-guide-nobody-wrote-how-to-actually-compare-processors-in-2026-d63b4a9476f2)
- Bottom line: NPUs in 2026 are a complement for always-on/assistant workloads; local LLM inference still runs on GPU (CUDA/Metal/Vulkan), with the NPU contributing only where the runtime explicitly supports it (QNN/OpenVINO/Vitis AI). Flag any vendor TOPS figure as marketing input, not inference throughput.

### 2.7 Jetson / edge (2026)

**Jetson AGX Thor (T5000)** — NVIDIA's "physical AI" flagship (robotics/humanoids), not a desktop LLM box:
- Specs: Blackwell GPU, 2,070 TFLOPS FP4 sparse, 128 GB LPDDR5X, 40–130 W configurable, 14-core Arm Neoverse-V3AE CPU, Multi-Instance GPU; ~7.5× AI compute and ~3.5× energy efficiency claims vs AGX Orin (note: mixes FP4-sparse vs INT8-sparse in the comparison). **Dev kit $3,499**; bulk module $2,999. [independent] (http://aiwiki.ai/wiki/jetson_thor), (https://www.techradar.com/pro/nvidia-quietly-unveiled-its-fastest-mini-pc-ever-capable-of-topping-2070-tflops-and-if-you-squint-enough-you-might-even-think-it-looks-like-an-rtx-5090)
- Developer kit released Aug 25, 2025 (shipments from Nov 20, 2025 per pre-order coverage); runs JetPack 7 (Ubuntu 24.04, kernel 6.8); JetPack 7.1 current production release. [independent] (https://github.com/liunix61/awesome-embedded-ai-stack/blob/HEAD/hardware-landscape/jetson.md), (https://twowintech.com/nvidia-jetson-agx-thor-full-analysis/)
- July 2026: NVIDIA announced lower **T3000 / T2000** modules (smaller memory/form factor, Q1 2027 availability; pricing unannounced; CNX/ServeTheHome flagged specs as preliminary). [independent] (http://aiwiki.ai/wiki/jetson_thor)
- LLM angle: 128 GB lets Thor run Llama-70B-class models locally, but ~273 GB/s bandwidth (like DGX Spark) limits token throughput; community consensus: relevant for 70B+ edge inference and VLA/robotics workloads, overkill for simple robotics (Orin Nano class suffices). [independent] (https://github.com/alpininsight/capi-provider-ssh/blob/HEAD/docs/roadmap/nvidia-jetson-edge-devices.md)
- Hands-on: HotHardware/ServeTheHome demos of GR00T N1 on Thor; wayin.ai tested local LLMs/VLMs on the dev kit (128 GB, 273 GB/s bandwidth noted as "slow for GPU memory, fast for CPU/RAM"). [independent] (http://aiwiki.ai/wiki/jetson_thor), (https://wayin.ai/wayinvideo/s/chshare0uMzeE0gKLh/)

---

## PART 3 — GGUF QUANTIZATION BENCHMARKS (2026)

