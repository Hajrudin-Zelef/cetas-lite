---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/6-3-vllm-vs-sglang-comparisons-all-secondary-independent-met
title: "6.3 vLLM vs SGLang comparisons (all secondary/independent — methodology varies)"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "Baseten", "DeepSeek", "Google", "Huawei", "Intel", "Microsoft", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "agents", "amd", "ascend", "attention", "awq", "aws", "benchmark", "blackwell", "consumer", "deepseek"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [856, 893]
section: "PART 2 — SGLang"
sha256: 04365714a39ecd1c264ec7e1ceeef979f1e90d8fb7ac5d708812a714bbcf7877
---

# 6.3 vLLM vs SGLang comparisons (all secondary/independent — methodology varies)

### 6.3 vLLM vs SGLang comparisons (all secondary/independent — methodology varies)
- **PremAI 2026 benchmark (H100 80GB, Llama 3.1 8B):** SGLang ~16,200 tok/s vs vLLM ~12,500 tok/s — **SGLang +29%**; LMDeploy ~16,100 (tie) [secondary](https://blog.premai.io/vllm-vs-sglang-vs-lmdeploy-fastest-llm-inference-engine-in-2026/) via [secondary summary](https://github.com/profsynapse/synaptic-tuner/blob/HEAD/docs/preparation/vllm-vs-sglang-inference-serving-research.md).
- **Spheron (Llama 3.3 70B FP8, H100):** SGLang +29% total throughput, **+117% output-token throughput** (894 vs 413 tok/s), TTFT 79 vs 103 ms (−23%), ITL 6.0 vs 7.1 ms (−15%) [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison). Same source: with **unique prompts (no shared prefixes) the gap shrinks to near-zero** — concurrency 1→100 shows only +2–5%.
- **Community single-request test (old, ~2025-06):** vLLM 60.0 tok/s vs SGLang 52.7 tok/s on a single unique prompt — vLLM ~1.1× faster [secondary](https://github.com/brendanmckeag/sglang-vllm-benchmark/blob/HEAD/README.md). Dated; treat as historical.
- **DeepSeek-specific:** SGLang claimed **3.1× faster than vLLM on DeepSeek-V3** via optimized MLA backends (FA3/FlashInfer/FlashMLA/CutlassMLA) [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison).
- **Consumer GPUs (RTX PRO 6000, Sept 2026, community):** Qwen3.8-27B NVFP4 + DSpark: 119.19 tok/s output, 1,072.69 tok/s total (8K-in/1K-out, 8 reqs), mean TTFT 729.82 ms, mean TPOT 7.68 ms, DSpark accept length 2.50 [secondary/independent](https://github.com/lEWFkRAD/qwen38-rtx-pro-6000).
- **Artificial Analysis:** no Artificial-Analysis-published SGLang-vs-vLLM engine benchmark was found in this research pass ([unverified]/not found). SemiAnalysis InferenceX (above) is the closest independent price/performance source.
- **ServeTheHome:** no STH SGLang benchmark article was surfaced in this pass ([unverified]/not found).
- **SiliconANGLE:** no SiliconANGLE SGLang article was surfaced in this pass ([unverified]/not found).

### 6.4 Caveats
- The +29% SGLang-vs-vLLM gap is a **high-concurrency batching** result; single-stream unique-prompt workloads show ~0–12% either way. RadixAttention's advantage concentrates in prefix-sharing workloads (agents, RAG, multi-turn chat) [secondary].
- One community SM120 report (Aug 2026) flagged garbage output with INT4 quantized models on SGLang (issue #21132) and recommended vLLM for that config; upstream fixed NVFP4 paths in v0.5.16+ and improved SM120 support in v0.5.19/0.5.20 — current status [unverified] [secondary](https://github.com/randomchaos7800-hub/inference-research/blob/HEAD/tower/gdn-blackwell/sglang-vs-vllm-sm120.md).

---

## 7. Deployment

- **Docker:** official images `lmsysorg/sglang` on Docker Hub [official, README install docs]. v0.5.20 image matrix: CUDA 13.x (CUDA 12 lane retired after v0.5.19), CUDA 13.4 preview for **Rubin**, ROCm 10 (gfx942/gfx950/gfx1250, MI30x/MI35x), Intel XPU (`lmsysorg/sglang:vX.Y.Z-xpu`), gfx1151 (Strix Halo / Ryzen AI MAX+), Moore Threads MUSA [official].
- **PyPI:** `pip install sglang` (PyPI release cadence tracks GitHub tags; v0.5.18 was latest PyPI release in Aug 2026 per downstream doc; Python ≥ 3.10) [secondary]. Jetson (l4t) builds tracked separately [secondary].
- **Kubernetes:** no SGLang-specific operator found in this pass; the standard route is the **SGLang Model Gateway** (Kubernetes service discovery, multi-model routing) [official docs ref via secondary](https://github.com/hygon-ai/sglang-das/blob/HEAD/docs/docs/advanced_features/sgl_model_gateway.mdx) and the **llm-d** project (Kubernetes-native distributed inference), whose guides ship SGLang inference-server configs "validated each release" on NVIDIA GPU and AMD GPU (GKE + base) [secondary](https://github.com/rohitg00/llm-d/blob/HEAD/guides/pd-disaggregation/README.md).
- **Cloud providers:** adopted/deployed on Oracle Cloud, Google Cloud, Microsoft Azure, AWS (per README adopter list) plus GPU clouds: Nebius, DataCrunch, Novita, RunPod, Voltage Park, Atlas Cloud, InnoMatrix, Modal, Baseten [official README; production depth per provider [unverified]].
- **Hardware backends (2026-09):**
  - **NVIDIA CUDA:** H100/H200/B200/B300/GB200/GB300 NVL72, A100, RTX 5090/PRO 6000, DGX Spark; SM90/100/103/107/120; CUDA 13 default since v0.5.11; Rubin preview (v0.5.19+) [official].
  - **AMD ROCm:** MI300X/MI355X (Lean attention, AITER, MoRI/Mori-EP, ROCm 10 images; MI355X reached CUDA-stack feature parity incl. MTP per InferenceX [independent]); ROCm 7.2.4 → 7.0 retired in v0.5.20 [official].
  - **Intel:** Xeon CPU; XPU (versioned Docker images since v0.5.20; INT4 AWQ/GPTQ dense linear, SYCL kernels for DeepSeek-V4 MHC) [official].
  - **Google TPU:** SGLang-Jax native backend (Oct 2025); **RadixArk + Google brought full SGLang features to TPUs** (blog, 2026-07) [official README news].
  - **Ascend NPU:** NPU backend (DeepSeek-V4, MXFP4-W4A4 MoE quant, MiniMax-H3 diffusion, Kimi-K3 on Ascend A3 cookbook) [official v0.5.19/v0.5.20].
  - **Moore Threads MUSA:** image added v0.5.20 [official].
  - **Apple MLX / macOS:** MLX backend (Torch 2.13/MLX 0.32+, v0.5.19); macOS added as SGLang-Diffusion platform (v0.5.10) [official].
  - **Intel Gaudi (HPU):** initial `--device hpu` support PRs from 2024 (#2121/#2357, offline batch inference on Gaudi2) [official PRs]; but a May-2026 industry guide states **"SGLang: Not supported on Gaudi 3; the RadixAttention scheduler has no HPU backend"** [secondary](https://www.spheron.network/blog/intel-gaudi-3-vs-nvidia-h200-b200-llm-inference-2026/) — treat HPU as experimental, not production-grade.
  - **Disaggregated deployment:** PD disaggregation via Mooncake/NIXL RDMA backends; MooncakeStore/HiCache L3 tiers [official].
- **APIs:** OpenAI-compatible (`/v1/chat/completions`, `/v1/completions`, `/v1/embeddings`, `/v1/audio/speech` in Omni), Anthropic-compatible bits, `/v1/responses` (storage opt-in since v0.5.20), MCP integration in Model Gateway, gRPC support [official/secondary].

---

## 8. Ecosystem

