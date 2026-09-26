---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/3-architecture
title: "3. Architecture"
domain: part-2-sglang
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Moonshot", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: ["2024-01-17", "2025-05", "2025-05-05", "2025-06-16", "2025-09-25", "2025-11-07", "2025-12-16", "2026-01-01", "2026-01-16", "2026-01-23", "2026-02-19", "2026-02-24", "2026-04-06", "2026-09-05", "2026-09-18"]
keywords: ["agent", "amd", "attention", "decode", "deepseek", "diffusion", "disaggregated", "fp8", "glm", "gpu", "gpus", "gqa"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [748, 792]
section: "PART 2 — SGLang"
sha256: bac1754db1953be0b909db7ffa68b1bb8295968d522a83451332bec54d963564
---

# 3. Architecture

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

