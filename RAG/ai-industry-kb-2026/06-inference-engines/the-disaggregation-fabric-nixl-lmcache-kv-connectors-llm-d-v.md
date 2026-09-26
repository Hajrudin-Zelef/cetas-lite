---
id: ai-industry-kb-2026/06-inference-engines/the-disaggregation-fabric-nixl-lmcache-kv-connectors-llm-d-v
title: "The disaggregation fabric: NIXL, LMCache, KV connectors, llm-d vs Dynamo"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AWS", "Alibaba", "MiniMax", "OpenAI", "vLLM"]
dates: ["2026-01-31", "2026-02-02", "2026-02-28", "2026-03-28", "2026-05-07"]
keywords: ["attention", "aws", "decode", "diffusion", "disaggregated", "disaggregated serving", "dpo", "full-duplex", "gpu", "inference", "kv cache", "memory"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2296, 2310]
section: "6. Inference Engines"
sha256: 6820411ea172a8a6f2a627f349b520bbafb7d361ebbd7274ad02a26449bb4d09
---

# The disaggregation fabric: NIXL, LMCache, KV connectors, llm-d vs Dynamo

- **2025-11:** officially released under the vLLM community (any-to-any multimodal serving).
- **2026-02-02:** paper arXiv:2602.02204; reports up to **91.4% reduction in job completion time** vs baselines on any-to-any pipelines [UNVERIFIED as independently reproduced].
- **2026-01-31:** vLLM-Omni v0.14.0 — first stable Wan2.2 diffusion pipelines.
- **2026-02-28:** v0.16.0 — `/v1/videos` API.
- **2026-03-28:** v0.18.0 — IPC −17.5% on Wan2.2.
- **2026-05-07:** v0.20.0 — fused DiT + CI JSON. Wan2.2-I2V H200 wall-clock progression: 133.94 s (v0.16.0 retro) → 93.67 s (v0.18.0) → 79.19 s (v0.20.0) [COMMUNITY cookbook archaeology].
- **2026-08:** vLLM-Omni v0.28.0 line — production-ready MiniMax H3 serving, **unified AR/DiT paged KV cache runtime** (one runtime for autoregressive and diffusion generation), enhanced realtime full-duplex MiniCPM-o. **VeRL-Omni v0.2.0** — faster diffusion RL powered by vLLM-Omni (request-level/step-wise batching with FA3), rebuilt Qwen3-Omni multimodal training (DPO & GSPO), LTX-2.3 and Qwen-Image-Edit support.
- Architecture: fully disaggregated serving split into **three independent GPU pools** — modality encoders (throughput-bound at moderate batch sizes), prefill (needs max TFLOPS), decode (memory-bandwidth-bound). A stage abstraction decomposes any-to-any architectures into interconnected stage graphs; per-stage request batching, flexible GPU allocation, and unified inter-stage connectors (OmniConnector); pipelined stage execution overlap; dynamic resource allocation. After encoder inference, modality feature tensors transfer to the prefill pool via **NIXL** over RDMA/TCP (Mooncake as alternative for multi-region/heterogeneous topologies); prefill produces KV cache transferred to decode via NIXL.
- Model coverage (2026-08): omni-modality (Qwen3-Omni, MiniCPM-o 4.5, Cosmos3, HunyuanImage, BAGEL), TTS (Qwen3-TTS, VoxCPM2, Ming-Omni-TTS, CosyVoice3), diffusion image/video/audio (MiniMax H3, Qwen-Image, Wan2.2, FLUX), robot-policy/action models (GR00T-N1.7, DreamZero-DROID, InternVLA-A1, Cosmos3 action policy). OpenAI-compatible API server with streaming; full-duplex realtime audio experimental.
- Serving surface (AWS DLC docs): `/v1/audio/speech` (TTS), `/v1/audio/generate`, `/v1/images/generations`, `/v1/videos` (async) and `/v1/videos/sync`, `/v1/chat/completions` (multimodal chat). Example models: Qwen3-TTS-1.7B, CosyVoice3-0.5B, Stable-Audio-Open-1.0, FLUX.2-klein-4B, ERNIE-Image-Turbo, Wan2.1-T2V-1.3B, Qwen2.5-Omni-3B.
- Serving syntax for the diffusion side: `vllm serve Wan-AI/Wan2.2-I2V-A14B-Diffusers --omni --enable-diffusion-pipeline-profiler`; multi-GPU recipes combine USP2 + VAE patch-parallel 2 + HSDP + VAE slicing [COMMUNITY].
- AWS Deep Learning Containers ship vLLM-Omni images (CUDA 13.0, PyTorch, NCCL, Python 3.12) bundled with FlashInfer (fused attention, precompiled cubins), DeepEP (expert-parallel kernels for large MoE), LMCache + NIXL, runai-model-streamer, EFA/OpenMPI, espeak-ng/ffmpeg.

### The disaggregation fabric: NIXL, LMCache, KV connectors, llm-d vs Dynamo

