---
id: collect-261001-ia-llm/ia-llm/etape4-tracka-vllm-sglang-11
title: "Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Hugging Face", "LongCat", "MiniMax", "Mistral", "Moonshot", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["inference", "sglang", "vllm", "amd", "attention", "awq", "blackwell", "copilot", "decode", "deepseek", "diffusion", "embedding"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [793, 826]
sha256: 7c33199fb2af72fc0233991296a19df6ab26477181c54c2f67b4734726884a8e
---

# Step 4 — Inference Serving Frameworks: vLLM + SGLang (February 1 → September 22, 2026)

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

## 4. Quantization, model formats & architectures

- **Supported (README):** FP4 / FP8 / INT4 / AWQ / GPTQ [official]. v0.5.19 release notes additionally exercise: **NVFP4, MXFP4, MXFP8, W4A8, W4A4 (MegaMoE), FP8 KV cache, mxfp8 KV cache, compressed-tensors** (quantized lm_head, kv_cache_scheme scales), torchao integration, ModelSlim (Kimi-K3) [official].
- **Marlin:** FP4 Marlin helpers were deduped in the v0.5.20 quantization refactor; the old **AWQ AOT kernel and non-Marlin GPTQ were deleted** in v0.5.20 — Marlin remains the path for GPTQ/AWQ-style INT4 [official]. Community config translation confirms `--quantization gptq_marlin` is auto-detected [secondary](https://github.com/randomchaos7800-hub/inference-research/blob/HEAD/tower/gdn-blackwell/sglang-vs-vllm-sm120.md).
- **NVFP4 MoE backends:** as of v0.5.16, the buggy `cutlass` FP4 path was deleted; `triton` and `flashinfer_cutlass` are the viable backends; FlashInfer CuTe DSL NVFP4 W4A16 mode added v0.5.19 (Qwen3-30B-A3B W4A16 GSM8K 0.965 → 0.980) [official].
- **FP8:** FP8 KV cache; FP8 DeepEP dispatch (Humming backend: TPOT −35%, TTFT −23% vs BF16 on H20, v0.5.19); FP8 block-scale MoE keeps fp32 routing weights (v0.5.20) [official].
- **Model coverage:** Llama, Qwen (incl. Qwen3.x, Qwen3.5/3.6/3.8, Qwen3-VL), DeepSeek (V3/R1/V4 + MLA/DSA/MTP), Kimi (K2/K3/Linear), GLM (4.5/4.6/5.x, incl. GLM-5.2/5.3-Flash NVFP4), GPT (gpt-oss day-0 Aug 2025), Gemma (incl. Gemma-4 FP8 MTP), Mistral (incl. Mistral Large 3, Ministral), Nemotron 3 (Nano/Super/Ultra/Lightning), LFM2/2.5, MiniMax (M2/H3/M3, incl. Music 3), Spark2.5, Granite 4.2, LongCat, Ling-3.0, dots3.note, MiniCPM (incl. MiniCPM5/SALA), Intern-S2-Mobius, Mamba-hybrid models (Qwen3.8-Next, Nemotron-H, Kimi-Linear), embedding models (e5-mistral, gte, mcdse), reward models (Skywork), diffusion models (WAN, Qwen-Image, FLUX.2, HunyuanVideo, LTX-2.5, SANA-Video, Cosmos3, MiniMax-H3, LongCat-Image, SenseNova-U1.5-8B-MoT) [official, across release notes + README].
- **Model formats:** Hugging Face checkpoints (safetensors), GGUF (diffusion transformer checkpoints, MiniMax-H3, v0.5.19), ComfyUI serialized checkpoints (NVFP4/W4A8/INT8 DiTs), pruned safetensors [official]. Transformers upgraded 4.57.1 → 5.3.0 in v0.5.10 [official].

---

## 5. New 2026 features (beyond §2)

