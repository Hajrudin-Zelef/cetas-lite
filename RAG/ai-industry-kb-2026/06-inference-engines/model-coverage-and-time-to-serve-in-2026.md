---
id: ai-industry-kb-2026/06-inference-engines/model-coverage-and-time-to-serve-in-2026
title: "Model coverage and time-to-serve in 2026"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "DeepSeek", "Huawei", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-01", "2026-03", "2026-05-05", "2026-07", "2026-08-20", "2026-08-26", "2026-09-04"]
keywords: ["amd", "apache", "ascend", "attention", "aws", "cost", "deepseek", "diffusion", "disaggregated", "disaggregated serving", "embedding", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2369, 2405]
section: "6. Inference Engines"
sha256: 0216c480fdd4a1c3e7a55eb00edf518ba0456b332ceed801a9659b58b354f48d
---

# Model coverage and time-to-serve in 2026

- **Architecture:** proprietary `nim-llm` orchestration + `nimlib` (model licensing, hardware-aware profile selection, health endpoints) over **OSS vLLM (Apache 2.0)** as the core inference engine with OpenAI-compatible API. Described as "an upstream-first architectural shift from NIM 1.x."
- **Backend map:** **LLM NIM (v2.0) → vLLM (sole backend)**; **VLM NIM → vLLM 0.19 (sole backend)**; Embedding NIM → TensorRT + Triton; Speech/Biology NIM → Triton + custom backends; Edge NIM → TensorRT + Triton (no vLLM).
- **Implication:** NVIDIA conceded the general LLM-serving layer to vLLM and now competes on orchestration (Dynamo), packaging (NIM), and silicon — a structural validation of vLLM's "industrial standard" status. Treat the exact NIM version timing as reported; seek an NVIDIA primary source (NIM 2.0 release notes, GTC session) before presenting it as NVIDIA's stated position.
- **NIM-side companion move:** NVIDIA folded Triton into Dynamo ("Dynamo-Triton"), adding disaggregated serving and KV cache management on top of Triton — the orchestration layer and the serving layer both re-centered in 2026.

### Model coverage and time-to-serve in 2026

- **vLLM v0.20.0 → v0.28.0 model intake:** DeepSeek V4 (DSA attention backend, token-leakage fix, MTP; then sparse MLA end-to-end in v0.28.0), DeepSeek V3.2, GLM-5/GLM-5.3-Flash, Kimi-K3, Qwen3.8 (incl. on AMD ROCm), Ling 3.0 Flash (BF16/MTP/FP8/hybrid MXFP4 routed experts), Dots3 NOTE native multimodal, MiniMax-M3 (pipeline parallelism + NVFP4), MiniMax-H3 (NVFP4 inference), Gemma 4 (variable-length audio batch padding), Muse Glimmer, Granite 4.1 Vision (built-in), LLaVA-OneVision-2, Unlimited OCR, MOSS-Transcribe-Diarize, Hy3, openai/privacy-filter, plus a Qwen3-Omni crash fix.
- **The convergence signal:** vLLM v0.28.0 and SGLang v0.5.19 (PyPI 2026-09-04, 786 PRs, 214 contributors) raced to serve the **same late-2026 flagships** (Qwen3.8, Kimi-K3) across NVIDIA/AMD/Ascend/DGX Spark — engine differentiation in 2026 is increasingly measured in **time-to-serve** for new models, not architectural features.
- **Community-patched day-0 evidence:** the GLM-5.3-Flash NVFP4 1M-context DGX Spark deployment required extending vLLM's SM90 NoPE sparse-MLA backend to SM121 — MLA-class models ship day-0 with per-architecture tuning, handled by the community when the release notes don't.
- **Multimodal model intake (vLLM-Omni line):** Qwen3-Omni, MiniCPM-o 4.5, Cosmos3, HunyuanImage, BAGEL (omni-modality); Qwen3-TTS, VoxCPM2, Ming-Omni-TTS, CosyVoice3 (TTS); MiniMax H3, Qwen-Image, Wan2.2, FLUX (diffusion); GR00T-N1.7, DreamZero-DROID, InternVLA-A1, Cosmos3 action policy (robot policies); Qwen2.5-Omni-3B, Stable-Audio-Open-1.0, ERNIE-Image-Turbo, Wan2.1-T2V-1.3B, FLUX.2-klein-4B in the AWS DLC serving-surface docs.

### Rust frontend, gRPC, and the 2026 ops surface

- **v0.25.0 → v0.28.0:** the Rust frontend matured with **HTTPS/mTLS and a DP supervisor**; v0.28.0 adds a **standalone renderer**, **multimodal image inference over gRPC**, explicit data-parallel rank routing, **RL lifecycle control**, and **protobuf schemas published to Buf**.
- **Observability:** profiler control routes; the API server path is multiprocess `AsyncLLM` over ZMQ (single-process mode unsupported on the API path since the v0.16.0 V1 lockdown).
- **Deployment defaults (v0.28.0):** CUDA 13.0 wheels on PyPI; ROCm 7.22 wheels; Docker images including `vllm-openai:v0.28.0`, ROCm, CPU, and XPU variants; Python 3.14; PyTorch 2.11; Transformers 5.15.0.
- **Breaking changes operators must track (v0.28.0):** bitsandbytes support migrated to an out-of-tree plugin; deprecated `calculate_kv_scales` and `override_attention_dtype` removed.

### The brief's vLLM errors, corrected in one place

- **"PagedAttention is the current core mechanism" → CONTRADICTED.** PagedAttention was removed from vLLM's main internal path in v0.25.0 (July 2026, PR #47361); it survives only as a legacy attention path. The 2026 story is Model Runner V2, sparse attention, KV offloading, and disaggregation.
- **"V0 abandonment dates to v0.28.0" → misdated by five months.** The V0 engine was fully removed no later than **v0.16.0 (March 2026)** — confirmed by RFC #18571's schedule and an independent forensic code review of v0.16.0. Nothing in v0.28.0's release notes concerns V0 removal.
- **"vLLM-Omni v0.28.0" → mislabeled.** v0.28.0 (2026-08-26) is the **core vLLM release** (Kimi-K3, sparse MLA, Model Runner V2, tiered KV offload). vLLM-Omni is a separate repo/product line; its multimodal items in v0.28.0 (gRPC image inference, vision encoders, Qwen3-Omni crash fix) belong to the core release notes.
- **"RadixArk launched in January 2026" → wrong launch date.** January 2026 was when the $400M valuation first appeared in reporting (TechCrunch's Inferact piece); the formal RadixArk launch was **2026-05-05** ($100M seed, Accel-led, Spark Capital co-led).
- **"Unsloth Dynamic 2.0" → stale version.** The current generation is **Dynamic v3.0** (v0.1.802-beta); older HF cards still reference 2.0.

### Version cadence, numbering discipline, and the pinning rule

- **2026 cadence:** vLLM shipped roughly one minor version per month (v0.15.1 → v0.16.0 → v0.20.0 → v0.24.0 → v0.25.0 → v0.26.0 → v0.27.0 → v0.28.0, with v0.29.0 attempted) — any "latest version" claim in the final document must carry a date stamp or it rots within weeks.
- **Internal churn has a pinning cost:** the same year removed V0, deleted PagedAttention, changed the default model runner, removed in-tree GGUF, and migrated bitsandbytes out-of-tree. Pinning commits is not optional for production GGUF serving: the documented working ROCm setup pinned `vllm-gguf-plugin` to git main and the image to `rocm10.0.0_ubuntu24.04_py3.14_pytorch_2.12.0_vllm_0.27.0`.
- **The plugin builds against the exact PyTorch:** `--no-build-isolation` matters; `PYTORCH_ROCM_ARCH` (e.g. `gfx1201`) cuts the ~12-arch default ROCm build (15–20 min) to minutes.
- **Parallel cadence on the SGLang side:** v0.5.18 (2026-08-20) → v0.5.19 (2026-09-04, 786 PRs/214 contributors) — the whole engine layer versions weekly-to-monthly in 2026; pinning and date-stamps are a layer-wide requirement, not a vLLM quirk.
- **What "current" means for vLLM-Omni vs core vLLM:** they share a v0.28.0 *line* but are different release artifacts — core `vllm` tags on GitHub vs the separate `vllm-omni` repo. The final document must never cite a core-vLLM version number as a vLLM-Omni release or vice versa.

### Production deployment patterns on vLLM (2026)

