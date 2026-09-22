---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/3-architecture
title: "3. Architecture"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Moonshot", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-01-17", "2025-05", "2025-05-05", "2025-06-16", "2025-09-25", "2025-11-07", "2026-02-19", "2026-09-05", "2026-09-18"]
keywords: ["agent", "amd", "attention", "copilot", "decode", "deepseek", "disaggregated", "glm", "gpu", "gpus", "gqa", "inference"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [765, 813]
section: "PART 2 — SGLang"
sha256: 742e83a2869033197db897db9b10fa21f96afa1677aae0b21cfd9bc2041f6ab3
---

# 3. Architecture

## 3. Architecture

### 3.1 RadixAttention / prefix caching (the signature feature)
- KV cache organized as a **radix tree (trie) over token sequences**; prefixes shared across requests are computed once and reused **automatically** — no explicit prompt-caching API needed [official, lmsys.org/blog/2024-01-17-sglang; paper arXiv:2312.07104](https://github.com/sgl-project/sglang). Benefits: multi-turn agent sessions, shared system prompts, RAG context blocks, few-shot prefixes.
- Reported gains: **up to 6.4× throughput** and **3.7× lower latency** vs vLLM/LMQL baselines on prefix-heavy workloads; KV hit rates ~50–99% [secondary, MarkTechPost 2025-11-07](https://www.marktechpost.com/2025/11/07/comparing-the-top-6-inference-runtimes-for-llm-serving-in-2025/). Original 2024 claim: up to **5× faster inference** with RadixAttention [official, lmsys.org/blog/2024-01-17-sglang/].
- **Eviction:** LRU over the radix tree. `UnifiedRadixCache` variant adds session-aware eviction (`SGLANG_ENABLE_UNIFIED_RADIX_TREE=1`): KV registered to an active `session_id` is evicted only after unreferenced KV (soft references, not pins) [secondary, docs reference](https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/sglang-inference-engine.md).
- **2026 evolution — unified radix tree:** became the default cache for *every* configuration in **v0.5.19 (2026-09-05)** [official]. In **v0.5.20 (2026-09-18)**: branching-point caching for the SWA component (DeepSeek-V4-Flash: token hit rate 43.8% → 60.8%, mean TTFT 1.57 s → 1.07 s with shared system prompt); opt-in external linker to a shared global memory pool via Mooncake or UMBP; PD decode workers reuse cached prefixes for SWA hybrid models (e.g. gpt-oss); runtime attach/detach of L3 storage; PP + HiCache L3 consistency across ranks [official].
- Cache-aware scheduling: "cache-aware sglang router" improvements shipped in v0.4.1 (2025); SMG v0.3.1 (Jan 2026) reworked cache-aware routing: 216,000 cache insertions/s (was 18,900), 4.6 µs/op (was 52.9 µs), ~99% memory reduction per tree node (~180 KB → ~1.4 KB) [official](https://github.com/sgl-project/sglang/releases/tag/gateway-v0.3.1).

### 3.2 HiCache — hierarchical KV cache
- 3-tier KV cache: GPU memory → host RAM → external storage (HF3FS, Mooncake). Flags: `--enable-hierarchical-cache --hicache-ratio 2 --hicache-size 100 --hicache-write-policy write_through` [secondary, docs reference](https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/sglang-inference-engine.md).
- v0.5.19: buffer-only mode for host-memory layer; batch PP write/load completion sync; −13% allocation time (no double mmap populate); L3 prefetch retry [official].
- v0.5.20: v0.5.19's mxfp8 KV cache support in PD transfer + CPU offload [official].

### 3.3 Parallelism
- Tensor / pipeline / expert / data parallelism, plus **DP attention** and **context parallelism** (prefill CP v1 removed in v0.5.20; strategy-based prefill CP is now the only path; **decode context parallelism (DCP)** runs on `trtllm_mla` since v0.5.19) [official].
- **DeepEP / expert parallelism:** DeepEP classic backend; **DeepEP v2 (ElasticBuffer)** added v0.5.19 (`--moe-a2a-backend deepep_v2`) enabling decode under CUDA graphs across nodes; AMD **Mori-EP / MoRI** backend for MI300X/MI355X (Kimi-K3, DeepSeek-V4, GLM-5.2) [official]. **Elastic NIXL-EP** for partial failure tolerance (v0.5.10) [official]. Large-scale EP milestone (May 2025): 52.3K in-tok/s, 22.3K out-tok/s on 96 GPUs — 5× faster than vanilla TP [secondary](https://github.com/semianalysisai/inferencex-app) / [official blog 2025-05-05].
- **LayerNorm sequence parallelism** (`--enable-layernorm-sp`, v0.5.19): −3.5% prefill (H100) / −5.6% (B200) on Qwen3-8B; dense Qwen3 only [official].
- **W4A4 MegaMoE** server flag (`--enable-w4a4-megamoe`, v0.5.19); FlashInfer MegaMoE runner (`flashinfer_megamoe`, v0.5.20, +11.9% prefill throughput on DeepSeek-V4-Flash NVFP4) [official].

### 3.4 Disaggregated prefill/decode (PD) + Mooncake
- PD disaggregation is a first-class, production feature: DeepSeek's production prefill-decode disaggregation was reproduced in SGLang [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison).
- v0.5.10: GPU staging buffer gathers scattered head slices into contiguous memory for bulk RDMA — ~1000× fewer RDMA requests on GQA, ~5× TPS/GPU at high concurrency [official].
- v0.5.19: PD transfer time −28% to −39% (Kimi-Linear-48B, 8×B300) via packed dest-contiguous RDMA blocks; deferred decode-side KV release for aborts mid-transfer (incl. NIXL backend); mxfp8 KV cache in PD transfer [official].
- v0.5.20: DCP1→DCP-N relayouts transfer **DSpark draft KV** (DSpark under PD with decode context parallelism, verified 8×B300 over NIXL and Mooncake to 256K input); new EPD refactor [official].
- **Mooncake:** dependency `mooncake==0.3.13` (v0.5.19) [official]; MooncakeStore as L3/external storage tier for HiCache; Mooncake MoE A2A backend; Mooncake backend for the unified-cache external linker (v0.5.20) [official]. A downstream doc describes HiCache's third tier as "HF3FS, Mooncake" [secondary].
- GB200 NVL72 blogs (official, lmsys.org): Part I (2025-06-16) 2.7× higher decoding throughput; Part II (2025-09-25) 3.8× prefill / 4.8× decode throughput with PD + large-scale EP [official]. GB300 NVL72 blog (2026-02-19): "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [official]. GB300 long-context work continued into 2026.

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

