---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/3-5-speculative-decoding
title: "3.5 Speculative decoding"
domain: part-2-sglang
role: deep-dive
task: model-release
actors: ["AMD", "Alibaba", "DeepSeek", "SGLang", "vLLM"]
dates: []
keywords: ["speculative decoding", "amd", "attention", "copilot", "decode", "deepseek", "gqa", "inference", "latency", "memory", "moe", "paged attention"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [793, 813]
section: "PART 2 — SGLang"
sha256: 6fde27e0da8e98bdef39d10bbe5994c665f1152a324b33ef21b3ba3e73e46916
---

# 3.5 Speculative decoding

### 3.5 Speculative decoding
- Algorithms supported: **EAGLE** (incl. EAGLE3-generation work), **DFlash / DFlash2** ("next generation of speculative decoding", blog 2026-06), **DSpark** (confidence-driven, new in v0.5.16), **KDA**, **NEXTN**, **NGRAM**, **MTP** (multi-token prediction), **Standalone** draft models, adaptive speculative decoding [official].
- Reported wins: DeepSeek models via EAGLE MTP — 1.8× decode speedup at batch 1, 1.5× at batch 32 on H200 [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison); DFlash2 3.43× over no-spec at batch 1, +24% over DFlash at concurrency 64 (v0.5.19) [official]; NEXTN 1.35–1.76× over no-spec on AMD after GQA packing fix (v0.5.19) [official]; LFM2/LFM2-MoE DSpark 1.05–2.42× faster decoding (1×H100) [official]; FR-Spec on Flash-Next (community, Sept 2026): 65,536-token draft vocabulary [secondary](https://github.com/jpezzulli/sglang-rtxpro6000/blob/HEAD/RESULTS.md).
- **Beam search** added v0.5.19 (`beam_width`) — distinct from speculative decoding, does not yet compose with it [official].
- Compatibility caveats [secondary, LocalAI docs 2026]: DFLASH and NGRAM incompatible with `enable_dp_attention`; DFLASH requires `pp_size == 1`; STANDALONE incompatible with `enable_dp_attention`; NGRAM is CUDA-only and disables the overlap scheduler.

### 3.6 CUDA graphs
- **Piecewise & breakable CUDA graphs** are the default execution mode since v0.5.10 [official]; v0.5.19 adds PP prefill CUDA graphs (up to 2.48× at 2K-token forwards on Qwen3.5-397B, GB300), DP-attention coordination, −18.9% median TTFT / +12.9% QPS from reduced idle DP work [official]; v0.5.20 sizes the graph pool from warmup measurements and reuses output storage across full prefill graphs [official].

### 3.7 Attention backends
- FlashAttention-3, **FlashInfer** (0.6.18 required since v0.5.19; CuTe DSL; `flashinfer_trtllm` mxfp8 GEMM), FlashMLA / CutlassMLA / Tokenspeed MLA, Triton, **HiSparse** (sparse attention, v0.5.10), **fmha_v2** for SM90/120 (~15% faster than FA3 at kernel level, v0.5.19), TRT-LLM MLA/MHA decode, DeepGEMM, AITER (AMD), TileLang-backed MLA path [official].
- MLA optimization and DP attention were in place before v0.4.1; SGLang claims 3.1× faster DeepSeek-V3 inference vs vLLM via optimized MLA backends [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison) / [official v0.4.1 notes](https://github.com/sgl-project/sglang/releases/tag/v0.4.1).
- **DeepSeek Sparse Attention (DSA)**: first-class; DSA prefill top-k v2 kernel 1.3–1.8× on B200 (v0.5.19); Q8KV8 sparse MLA prefill runtime for DeepSeek-V4 (+4.4–7.5% prefill throughput, H20) [official].

### 3.8 Scheduler & runtime
- Zero-overhead batch scheduler (v0.4, Dec 2024) [official]; continuous batching, chunked prefill (mixed chunk prefill base, v0.5.19), paged attention, overlap scheduling, per-scheduler load published on a dedicated socket for load-aware routers (v0.5.19) [official].
- Rust server (`SGLANG_RUST_SERVER=1`): process-local in-memory KV indexer + Router integration, e2e latency metadata, configurable HTTP/2 window (v0.5.19) [official].
- Default serving port **30000** (vLLM: 8000); Prometheus metrics under `sglang:*` prefix (requires `--enable-metrics`) [secondary](https://github.com/fuzzifikation/vllm-copilot/blob/HEAD/docs/sglang-compat-plan.md).

---

