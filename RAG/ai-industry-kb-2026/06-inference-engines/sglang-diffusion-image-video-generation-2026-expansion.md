---
id: ai-industry-kb-2026/06-inference-engines/sglang-diffusion-image-video-generation-2026-expansion
title: "SGLang Diffusion — image/video generation (2026 expansion)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Huawei", "Intel", "LongCat", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-06", "2026-08-20", "2026-09-04", "2026-09-22"]
keywords: ["diffusion", "sglang", "video generation", "alignment", "amd", "ascend", "attention", "blackwell", "consumer", "datacenter", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2687, 2706]
section: "6. Inference Engines"
sha256: a8d637b77ab397a32b5cea39b24da3de4291e85e213cefa40267bd4e2eda213a
---

# SGLang Diffusion — image/video generation (2026 expansion)

- Released on PyPI 2026-09-04 (GitHub tag v0.5.19): 786 PRs from 214 contributors; the v0.5.18 tag (2026-08-20) remains what official install docs point to for source builds — v0.5.19 is fresh as of the 2026-09-22 research cutoff.
- New models: Qwen3.8 (2.4T-A95B), Qwen3.8-27B, dots3.note, Ling-3.0-flash, Ling-3.0-tiny, Spark2.5, MiniCPM-SALA, Granite 4.2, plus LongCat-Image-Edit & Edit-Turbo (diffusion).
- Cookbook updates: GLM-5.3 deployment guide; PaddleOCR-VL deployment guide; Kimi-K3 on Ascend A3; Kimi-K2.7-Code-MXFP4 on MI355X; Qwen3.5 MXFP4 on MI355X with FP8 KV cache or HiCache host-memory tier; MiniMax-H3 on 24GB GPU or DGX Spark with a consumer-GPU tuning guide; Ling-3.0-flash on DGX Spark.
- The same model set (Qwen3.8, Kimi-K3) appears in both vLLM v0.28.0 and SGLang v0.5.19 release notes — both engines are racing to serve the same late-2026 flagships across NVIDIA, AMD, Ascend, and DGX Spark; engine differentiation is increasingly measured in time-to-serve for new models.
- Related: the sglang-omni line is at v0.1.4 (config refactor, AMD ROCm Qwen3-TTS/ASR on gfx950, SGLang 0.5.18 dependency alignment) — the any-to-any serving track, separate from core 0.5.19.
- MLA-2026 arc: zero-overhead scheduler and cache-aware load balancer (v0.4 blog); FlashInfer SM120 sparse-MLA decode kernels for DeepSeek-V4 (PR #27455, ~June 2026) delivering 2.2–3.7x TPOT reduction vs the Triton kernel; FP4 MLA KV caches in v0.5.6 (Dec 2025); `--attention-backend=dsv4` for hybrid CSA+HCA attention on DeepSeek V4 Flash.
- Hardware reality check: MLA kernels are SM90+-first (Hopper); SM120 (Blackwell workstation) support arrived through 2026 via FlashInfer with community patching (NaN bugs in specific batch ranges on SM121 bisected and fixed in nightlies) — mature on datacenter Hopper/Blackwell, still being tuned on workstation silicon.

### SGLang Diffusion — image/video generation (2026 expansion)

- The `sglang-diffusion` sub-project (`python/sglang/multimodal_gen/`) reuses SGLang's serving muscle (sgl-kernel ops, scheduler loop, CUDA-graph capture, quantization, multi-hardware, OpenAI-compatible API, PD-style disaggregation) for iterative denoising; installable via `uv pip install "sglang[diffusion]"`.
- Diffusion-specific optimizations: TeaCache (skip redundant denoise steps) and Cache-DiT integration — up to 7.4x inference speedup with minimal quality loss via `SGLANG_CACHE_DIT_ENABLED=True` [VENDOR integration claim].
- Model support: Wan series, FastWan, Hunyuan, Qwen-Image / Qwen-Image-Edit, Flux, Z-Image, GLM-Image, Ideogram 4, Krea-2, Cosmos3, LTX-2 / LTX-2.3, MiniMax-H3, MOVA, Hunyuan3D, and more.
- Hardware (verified from the runtime README): NVIDIA, AMD, Intel XPU, Ascend NPU, Apple Silicon (MPS), Moore Threads (MTT S5000) — the broadest hardware list of any 2026 diffusion-serving project.
- Interfaces: `sglang generate`, `sglang serve`, OpenAI-compatible `/v1/videos`, CLI, and Python SDK.
- Production-recipes signal: an Ant Group cookbook for serving MOVA under SGLang Diffusion — the diffusion-serving path moving from demos to documented production recipes.
- Positioning note: this expansion runs parallel to vLLM's any-to-any multimodal track (see sibling part 06a) — orthogonal to the text-serving battle.

- Positioning note: this expansion runs parallel to vLLM's any-to-any multimodal track (see sibling part 06a) — orthogonal to the text-serving battle.

