---
id: etape4-trackb-local-inference/00-local-inference/overview
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Apple", "DeepSeek", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2026-01-01", "2026-02-13", "2026-08", "2026-08-18", "2026-08-21", "2026-08-25", "2026-08-31", "2026-09", "2026-09-04", "2026-09-16", "2026-09-22"]
keywords: ["inference", "llama", "llama.cpp", "attention", "benchmarks", "deepseek", "distribution", "flash attention", "gguf", "glm", "gpu", "kv cache"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [1, 81]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 548497f3c67e69c59e3a70a2f38dc7b8570d203555ed93705a7eeca4cdebc50b
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)
## Final research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 4 (Infra inference/training), Track B
**Coverage window:** January 1, 2026 → September 22, 2026
**Subjects:** llama.cpp, Ollama, LM Studio, local inference hardware landscape, GGUF quantization
**Status:** Research snapshot. Prices are dated snapshots; re-verify against official pages before use.

### Provenance legend
- **[official]** — project's own GitHub releases, docs, changelog, or company blog.
- **[vendor-reported]** — figure claimed by the vendor without independent audit.
- **[independent]** — reputable third-party press (ServeTheHome, Tom's Hardware, Bloomberg, Reuters, CNBC, The Register, SiliconANGLE) or independent measurement.
- **[secondary]** — lower-tier press, blogs, AI-generated news mirrors of official releases; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

### Contents
- Part A — llama.cpp deep dive
- Part B — Ollama deep dive
- Part C — LM Studio + local inference hardware + GGUF quantization landscape
- Master open-verification log (track-level uncertainties carried forward)

---

## Part A — llama.cpp

**Research date:** 22 September 2026 · **Coverage window:** February → September 2026
**Provenance tags:** [official] = ggml-org GitHub releases/docs · [vendor-reported] = NVIDIA/vendor-published numbers · [independent] = third-party benchmarks/press · [secondary] = community docs/wikis · [unverified] = not corroborated.

---

## 1. Release train: versions, dates, scheme

llama.cpp adopted formal semantic versioning in August 2026; before that it used rolling `b[NNNN]` nightly-style tags on nearly every master commit. Both schemes coexist:

- `vX.Y.Z` — "stable", slower cadence, recommended for downstream/distribution [official] (https://github.com/ggml-org/llama.cpp/releases/tag/v0.2.0)
- `b[NUM]` — bleeding edge/nightly, recommended for developers/technical users [official] (https://github.com/ggml-org/ggml/discussions/1579)

Verified stable releases (dates from GitHub API `published_at`, UTC):

| Tag | Published | Nightly build embedded | Source |
|---|---|---|---|
| v0.1.2 (pre-release) | 2026-08-18 | b10485 | [official] https://api.github.com/repos/ggml-org/llama.cpp/releases/tags/v0.1.2 |
| v0.2.0 | 2026-08-21 | b10566 | [official] https://api.github.com/repos/ggml-org/llama.cpp/releases/tags/v0.2.0 |
| v0.3.0 | 2026-08-25 | b10621 | [official] https://api.github.com/repos/ggml-org/llama.cpp/releases/tags/v0.3.0 |
| v0.4.0 | 2026-09-04 | b10809 | [official] https://api.github.com/repos/ggml-org/llama.cpp/releases/tags/v0.4.0 |

Latest nightly observed at research time: **b11001** (published ~2026-09-16, "last updated 6 days ago" on 2026-09-22 crawl), with b10997, b10994, b10796, b10776 trailing in the days before [official] (https://github.com/ggml-org/llama.cpp/releases/tag/b11001).

Release assets shipped per nightly (b11001 example): macOS arm64/x64 (KleidiAI variant listed DISABLED for macOS arm64), iOS XCFramework, Ubuntu x64/arm64/s390x CPU, Ubuntu x64/arm64 **Vulkan**, Ubuntu x64 **CUDA 12** (12.8 libs) / **CUDA 13** (13.4 libs), Ubuntu arm64 CUDA 13, Ubuntu x64 **ROCm 10.0**, Ubuntu x64 OpenVINO, Ubuntu x64 **SYCL FP32/FP16**, Linux arm64 **Snapdragon (CPU, Adreno GPU, Hexagon NPU)** with setup guide, Android arm64 CPU + Snapdragon, Windows x64/arm64 CPU, Windows arm64 **OpenCL Adreno**, Windows x64 CUDA 12.4/13.4 (+CUDA 13.4 arm64 preview), Vulkan, OpenVINO, SYCL, ROCm 10.0; openEuler x86/aarch64 builds listed as DISABLED (linked to PR #23705); plus a separate UI tarball and `nightly-tag.txt` [official] (https://github.com/ggml-org/llama.cpp/releases).

New official web presence: **https://llama.app** is now the project's site [official] (llama.cpp releases page). Download-and-run shorthand now exists: `llama cli -hf <user>/<repo>` and `llama serve -hf <user>/<repo>` (OpenAI-compatible server) [secondary] (https://selfhost.directory/project/llamacpp).

ggml companion releases: ggml **v0.20.2** (ggml/1589, via v0.1.2 changelog), **v0.21.0** (ggml/1597), **v0.22.0** (ggml/1607), **v0.23.0** (ggml/1579-discussed line) [official] (release pages v0.1.2–v0.4.0).

### v0.4.0 (2026-09-04) — headline [official] (https://github.com/ggml-org/llama.cpp/releases/tag/v0.4.0)
- New models: initial **Qwen3.8-Flash-Next** (`qwen4exp` arch; optimization still pending), **NVIDIA Nemotron-3-Puzzle-75B-A9B** (`NemotronHPuzzle`), **DSpark for Nemotron 3.5**, **nanbeige4.2-3B**.
- Core: lazy/on-demand tensor reading (`llama_lazy_mode`, `--lazy-mode`/`-lzm`; auto-on for tensors > 4 GiB) — reduces RAM peaking at load; `--n-cpu-ffn`; per-layer expert routing/FFN (`n_expert_used_max`); KV-cell token tracking (session/state version bump); n-gram history lookup; optimized KV restore; quantizer RAM cap + row-slab streaming; autoscaled YaRN training context; **DFlash2 support**; fused DFlash encoder into KV injection; **sparse flash attention for DeepSeek-V4/GLM and Qwen4exp**; **Apple RDMA as an RPC transport**; RPC event/async backend APIs; synthetic spec-acceptance options.
- Multimodal: video params (`--video-*`), `mtmd_tokenize_from_parts()`, DeepSeek-V4-Flash-Vision-Exp, Gemma-4 vision fixes, WebP via ffmpeg, Idefics3 fix.
- Server: **per-slot context limit**, `data:` URLs for media, `preserve_reasoning` enabled by default, reject prefilled assistant tool calls, synthetic spec acceptance options.
- ggml v0.23.0: sparse-attention ops (`ggml_flash_attn_ext_set_n_kv_max`), RPC event/async, Apple RDMA.

### v0.3.0 (2026-08-25) — headline [official] (https://github.com/ggml-org/llama.cpp/releases/tag/v0.3.0)
- New model: **dots3-note** multimodal with a new **DSA-ISWA KV cache** type.
- **MTP (multi-token prediction) support for GLM-4.5-Air**; tensor-split mode for DeepSeek 4 (`-sm tensor`, meta-backend tensor-split with split-state propagation fixes); DSpark for bailingmoe3; mamba2 GEMM-over-GEMV dispatch.
- `fit` logic moved from server to common, now accounting for `n_streams`.
- ggml v0.22.0: tensor-split in meta backend, Metal per-op kernels with parallel compilation, non-in-place `ggml_clamp`, new ops (`POOL_1D`, `PAD_REFLECT_1D`), Q2_K SYCL kernels, OpenCL MoE bias fusion.

### v0.2.0 (2026-08-21) — headline [official] (https://github.com/ggml-org/llama.cpp/releases/tag/v0.2.0)
- Marks the **start of consistent semantic versioning**.
- Highlights: ggml v0.21.0; KleidiAI SME2 F32 GEMV kernel; SYCL Q2_K/Q5_K kernels; tensor-split for LFM2/LFM2MoE; DSpark for LFM2; Metal quantized-KV dequantization for large batches; server router lazy-load `startup_models`; RPC `use_count` fusion; Vulkan Q8_0 KV dequant; Windows ARM64 CUDA support in CI.

### v0.1.2 (2026-08-18, pre-release) — headline [official] (API)
- Notable: CUDA MMVQ `nwarps=8` for bs=1 dense models on **DGX Spark** (#26843) — early GB10 tuning; docs for **MCP stdio servers** and CORS defaults (#26847); vocab support for **integer tokenizer scores** (#27260); ggml v0.20.2.

### Earlier 2026 (Feb–Aug, from `b` tags) — selected milestones [secondary]/[official]
- **b7964** (Feb 2026): Step3.5-Flash support (#19283). **b7973**: Qwen3.5 dense and MoE support (community-driven, PR #19435; author notes "Llama.cpp missing all the zero-day releases", built on the common-delta-net PR #19125) [secondary] (https://pypi.org/project/llama-cpp-pydist/0.31.0/).
- **b8018** (2026-02-13): pydist sync point with 44 upstream commits [secondary] (same).
- Community build records (Apr 2026, Strix Halo): b8119 (kyuz0 custom), b8299 official (+40% prompt speed over b8119), b8461 (kyuz0 Vulkan RADV) [secondary] (https://github.com/valentijnvenus/strix-halo-setup).
- llama-cpp-python fork (JamePeng) synced to upstream commit `9723942` on 2026-08-31 with `load_mode` + Qwen3.8-Flash-Next + DFlash2 NVFP4 fix [secondary] (https://github.com/tao71-ai/llama-cpp-python-jamepeng/blob/HEAD/CHANGELOG.md).

---
