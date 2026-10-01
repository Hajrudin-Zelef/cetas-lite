---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-13
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Apple", "Baseten", "DeepSeek", "Google", "Huawei", "Intel", "Microsoft", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang"]
dates: []
keywords: ["inference", "sglang", "vllm", "amd", "ascend", "attention", "awq", "aws", "copilot", "deepseek", "diffusion", "disaggregated"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [874, 904]
sha256: 711a9e56a44deda1cf27cddf5820fb2e37ef4603e729d910158fcdaed79819dc
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

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

- **SGLang Router:** cache-aware load balancer/router (Rust `sgl-router` + newer Rust server integration); the original v0.4 (Dec 2024) shipped a "cache-aware load balancer" [official]. Note: a third-party compat doc flags the *classic* `sgl-router` as missing `max_model_len` in `/v1/models` and no `/server_info` route — the **SGLang Model Gateway** (`sgl-model-gateway`) proxies to workers and is the supported path [secondary](https://github.com/fuzzifikation/vllm-copilot/blob/HEAD/docs/sglang-compat-plan.md).
- **SGLang Model Gateway (SMG):** high-performance model-routing gateway — load balancing, PD disaggregation, multi-model routing, gRPC, MCP, Kubernetes service discovery, history-store/privacy controls [secondary docs ref](https://github.com/hygon-ai/sglang-das/blob/HEAD/docs/docs/advanced_features/sgl_model_gateway.mdx); gateway-v0.3.1 (Jan 2026) with 10–12× routing speedup [official].
- **Related repos (sgl-project org):** `sglang-omni` (audio/TTS/ASR, v0.1.6, Sept 2026) [official](https://github.com/sgl-project/sglang-omni); `sgl-learning-materials` (slides/adoption); diffusion serving inside the main repo (SGLang-Diffusion). Downstream/community: `bytedance-iaas/sglang` fork, `hygon-ai/sglang-das`, LocalAI SGLang backend, LMCache integration work [secondary].
- **Dependencies (v0.5.20):** FlashInfer 0.6.18, sglang-kernel 0.4.7 (split out as standalone `sgl-kernel` package since v0.4.1), sgl-deep-gemm 0.2.0, sgl-deep-ep 0.1.2, mooncake 0.3.13, TileLang 0.1.12, compressed-tensors 0.18.0, transformers 5.3.0, torch 2.11–2.13, CUDA 13.x [official].
- **Contributors:** v0.5.20: 713 PRs from 237 contributors; v0.5.19: 786 PRs / 214; v0.5.18: 710 / 212 [official]. Release captains seen in API: Qiaolin-Yu, Fridge003, Kangyan-Zhou [official].
- **Sponsors/grants:** a16z Open Source AI Grant, third batch (announced 2025-06, per README news) [official]; PyTorch Ecosystem member (Mar 2025) [official]. Commercial company: **RadixArk** (see §9).

---

## 9. Enterprise support & commercial backers

