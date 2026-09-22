---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/per-release-detail
title: "Per-release detail"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Intel", "Moonshot", "Nvidia", "SGLang", "Z.ai"]
dates: ["2025-12-16", "2026-01-01", "2026-01-16", "2026-01-23", "2026-02-24", "2026-04-04", "2026-04-06", "2026-05-05", "2026-05-16", "2026-05-26", "2026-06-13", "2026-06-26", "2026-07-10", "2026-07-25", "2026-07-28", "2026-08-08", "2026-08-22", "2026-09-05", "2026-09-18"]
keywords: ["amd", "attention", "awq", "blackwell", "decode", "deepseek", "diffusion", "flash attention", "fp4", "fp8", "glm", "gptq"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [699, 764]
section: "PART 2 — SGLang"
sha256: e7c01179d8750ffaf0a2334cea57cbb0a194e44d0c7ba295fd32b07dd630c6e4
---

# Per-release detail

### Per-release detail

**v0.5.20 (2026-09-18)** [official](https://github.com/sgl-project/sglang/releases/tag/v0.5.20) — 713 PRs / 237 contributors.
- **Sampling masks for RL rollouts:** new `return_sampling_mask` — each decode step returns the exact token support and sampled-token log-prob so a trainer can replay rollouts without reconstructing top-k/top-p. Runs under overlap scheduling: +17% decode throughput at batch 1, +52% at batch 64 on Qwen3-8B vs previous implementation. Capacity via `--sampling-mask-max-tokens` (default 4096).
- **SGLang Simulator:** CPU-only simulator running the real scheduler, radix cache and hierarchical cache with a latency predictor instead of model forward. Predicts TTFT within ~6% on most traces (up to 10% on 32K–128K traces); prefix reuse within 0.05 pp. For cache/scheduling studies without GPUs.
- **Unified radix tree — branching-point caching for SWA:** keeps sliding-window state at fork points; on DeepSeek-V4-Flash with shared system prompt, token hit rate 43.8% → 60.8%, mean TTFT 1.57 s → 1.07 s. Opt-in external linker addressing a shared global memory pool via Mooncake or UMBP.
- **DSpark under PD + decode context parallelism:** DCP1 prefill can transfer DSpark draft KV to DCP-N decode; verified on 8×B300 over NIXL and Mooncake up to 256K input.
- **Responses API storage is opt-in** (`--enable-response-store`); without it retrieval/`previous_response_id`/background requests return 400; cannot be enabled in PD deployments.
- **Prefill context parallelism v1 removed** (breaking): strategy-based implementation is the only prefill CP path; CP on HIP/NPU/MUSA rejected until ported.
- **Faster model loading on ROCm:** staged pageable H2D copies; GLM-5.2 TP4 on 4×MI355X loads in 40.4 s vs 505.7 s before.
- **DeepSeek-V4 on Blackwell:** TRT-LLM attention kernels cover DeepSeek-V4 CSA/HCA layers on SM100/SM103 — ~1.2× prefill, ~1.45× decode vs FlashMLA at kernel level on B200. New `--moe-runner-backend flashinfer_megamoe` (+11.9% prefill throughput at 8192 tokens/rank on DeepSeek-V4-Flash NVFP4 TP4/DP4).
- **DeepSeek-V4 on RTX PRO 6000 (SM120):** sparse-MLA indexer on DeepGEMM paged-MQA kernel + DeepGEMM FP4 MoE backend replaces torch fallback; decode TPOT 36.1 → 10.5 ms at batch 1 on 4× RTX PRO 6000; TTFT −20% from 8K→128K input.
- **BREAKING / upgrade notes:** CUDA 12 lane retired (v0.5.19 last with `-cu12x` wheels/images); cutlass_mla, non-Marlin GPTQ, AWQ AOT kernel and Dual Chunk Flash Attention deleted in the quantization refactor; prefill CP v1 removed; `/v1/responses` storage opt-in. New images: ROCm 10 for MI30x/MI35x, gfx1151 (Strix Halo / Ryzen AI MAX+), **Moore Threads MUSA**. ROCm 7.0 CI/images/wheel retired. Intel XPU joins tagged release images (`lmsysorg/sglang:vX.Y.Z-xpu`).
- New models: GLM-5.3-Flash, Hy4-Preview, Qwen3.8-Flash-Next, K2 Horizon, Nanbeige4.2 (autoregressive); SenseNova-U1.5-8B-MoT, FastH3, VDN-H3 (diffusion).
- Dependencies: sglang-kernel 0.4.7, sgl-deep-gemm 0.2.0.

**v0.5.19 (2026-09-05)** [official](https://github.com/sgl-project/sglang/releases/tag/v0.5.19) — 786 PRs / 214 contributors.
- **Beam search** (`beam_width` param, returns n best sequences; does not yet mix with speculative decoding, disaggregation, DP attention, or HiCache).
- **DeepEP v2** (`--moe-a2a-backend deepep_v2`, ElasticBuffer; fixed-size buffers → decode under CUDA graphs even across nodes; for DeepSeek-V3/V4 and Qwen3-MoE FP8).
- **LayerNorm sequence parallelism** (`--enable-layernorm-sp`; −3.5% Qwen3-8B prefill on H100, −5.6% on B200; dense Qwen3 only).
- **W4A8 MoE on Hopper** (`--flashinfer-mxfp4-moe-precision fp8`; +12% output throughput on DeepSeek-V4-Flash, no GSM8K change; needs FlashInfer 0.6.18).
- **DCP on default Blackwell MLA backend** (`trtllm_mla`; at 128K input, plain TP stalls ~680 tok/s on 8×B200 while DCP keeps scaling with concurrency).
- **Faster speculative kernels:** DSA prefill top-k v2 kernel (1.3–1.8× on B200); KDA fused accept path (`SGLANG_OPT_KDA_FUSED_ACCEPT_STATE=1`, MTP verify-and-commit −45% to −63% on Kimi-Linear shapes, bit-identical); DFlash2 (3.43× over no-spec at batch 1, +24% over DFlash at concurrency 64).
- **Unified radix tree now default for every configuration** (breaking; `SGLANG_ENABLE_UNIFIED_RADIX_TREE` deprecated).
- **Lean attention on AMD:** persistent work-centric decode kernel; up to 1.52× throughput, 3.62× lower median ITL on MI355X; auto-enabled where it helps.
- **DSA models on ROCm:** GLM-5.2 PD decode TPOT 23.16 ms → 7.94 ms on 8×MI355X; DeepSeek-V4 top-k kernel up to 3×.
- Dependencies: FlashInfer 0.6.18, sgl-deep-ep 0.1.2, sgl-deep-gemm 0.1.7, mooncake 0.3.13; new CUDA 13.4 preview image for **Rubin**; ROCm 10 images for gfx942/gfx950/gfx1250.
- **BREAKING:** unified radix tree default; Spark3 → Spark2.5 rename (config `spark2_5`, tool-call parser `spark` → `spark25`); DeepSeek-V4 FP4 checkpoints default to FlashInfer MXFP4 MoE runner on SM90/SM100/SM120; W4A4 MegaMoE behind `--enable-w4a4-megamoe`; FlashInfer 0.6.18 required; `ServerArgs` construction no longer resolves (must call `resolve_once()`); requests limited to 32 stop strings / 32 stop regexes (HTTP 400 beyond); XPU `SGLANG_USE_SGL_XPU` flag removed; `/rerun-stage` removed, `/rerun-test` gated on commenter trust.
- Notable fix: `/v1/chat/completions` `usage.reasoning_tokens` was always 0 for thinking models — fixed in v0.5.10 cycle (PR #15562, merged 2026-04-04) [official release notes; confirmed fixed by downstream tracker](https://github.com/hitechcloud-vietnam/dgxarley/blob/HEAD/SGLANG_v0.5.10_VERSION_CHANGES.md).

**v0.5.18 (2026-08-22)** [official](https://github.com/sgl-project/sglang/releases/tag/v0.5.18) — 710 PRs / 212 contributors. Highlights body was not pulled in this research pass ([unverified] detail — see §12).
- Downstream note: v0.5.18 pinned by a community RTX-PRO-6000 build doc as the latest PyPI release in Aug 2026 (Python ≥ 3.10) [secondary](https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/sglang-inference-engine.md).

**v0.5.17 (2026-08-08)** [official] — 582 PRs / 194 contributors. **Kimi K3 day-0 support**: described in release notes as "a 2.8T-parameter multimodal LatentMoE (896 experts, top-16, routed in a …" (release body truncated in API fetch; full detail [unverified]).

**v0.5.16 (2026-07-25)** [official] — 574 PRs / 169 contributors.
- **DSpark: confidence-driven speculative decoding** (new spec algorithm).
- NVFP4 MoE: buggy `cutlass` backend path removed; `triton` and `flashinfer_cutlass` remain the viable NVFP4 MoE backends [official; confirmed by downstream source review 2026-07-28](https://github.com/vroomfondel/dgxarley/blob/HEAD/SGLANG_TP_EP_MOE_UPSTREAM_BUG.md).

**v0.5.15 (2026-07-10)**: **GLM-5.2 NVFP4 tuned for production** on Blackwell [official].

**v0.5.14 (2026-06-26)**: new models GLM-5.2, LiquidAI LFM2.5, Kimi-K2.7-Code [official].

**v0.5.13 (2026-06-13)**: **Nemotron 3 Ultra day-0** (+ blog) [official].

**v0.5.12 (2026-05-16)**: **DeepSeek-V4 day-0 support** — full inference path incl. tensor/data/expert parallelism (#23882) [official]; v0.5.12.post1 (2026-05-26) cherry-picks 12 DeepSeek-V4 fixes [official].

**v0.5.11 (2026-05-05)**: default CUDA moves to 13.0 (from 12.x), PyTorch 2.9 → 2.11 across SGLang/sgl-kernel/Docker images [official].

**v0.5.10 (2026-04-06)** [official] (details via releases-page crawl):
- **Piecewise CUDA Graph enabled by default** (lower memory overhead, better throughput for complex control flow) — #16331.
- **Elastic EP (NIXL-EP) for partial failure tolerance** for DeepSeek MoE: a failed GPU's expert weights are redistributed, serving continues without full restart.
- **GPU staging buffer for PD disaggregation**: contiguous RDMA bulk transfer; ~1000× fewer RDMA requests on GQA models; ~5× TPS/GPU at large concurrency (Qwen3.5, prefill TP4 + decode DEP4).
- **HiSparse sparse-attention backend** for long context.
- **FlashInfer MXFP8 kernels** (GEMM + MoE) for mixed-precision FP8.
- **Transformers 4.57.1 → 5.3.0** major upgrade.
- SGLang-Diffusion updates: LTX-2, Hunyuan3D-2, Helios models; Qwen-Image/Z-Image +1.5×; macOS platform; Cache-DiT integration into diffusers backend.

**v0.5.9 (2026-02-24)**: **LoRA weight loading overlapped with computation** — TTFT −~78% when swapping LoRA adapters (#15363) [official].

**v0.5.8 (2026-01-23)**: diffusion models up to 1.5× faster (ties to lmsys.org/blog/2026-01-16-sglang-diffusion) [official].

**v0.5.7 (2026-01-01)**: day-0 Mimo-V2-Flash (#15207, blog 2025-12-16) and other new models [official].

---

