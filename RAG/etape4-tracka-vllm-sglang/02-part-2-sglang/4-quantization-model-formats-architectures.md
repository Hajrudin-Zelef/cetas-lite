---
id: etape4-tracka-vllm-sglang/02-part-2-sglang/4-quantization-model-formats-architectures
title: "4. Quantization, model formats & architectures"
domain: part-2-sglang
role: deep-dive
task: quantization
actors: ["Alibaba", "Cohere", "DeepSeek", "Hugging Face", "LongCat", "MiniMax", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2024-07-25", "2024-09-04", "2024-12-04", "2025-05", "2025-05-05", "2025-06-16", "2025-09-25", "2026-01-16", "2026-02-19", "2026-02-20", "2026-05-19", "2026-05-20"]
keywords: ["quantization", "agent", "agentic", "agents", "attention", "awq", "benchmark", "benchmarks", "blackwell", "cohere", "consumer", "cost"]
source: docs/RAG/etape4_trackA_vllm_sglang.md
source_anchor: ""
source_lines: [814, 871]
section: "PART 2 — SGLang"
sha256: 38deb92d0b4eb0f8007925ac8e024cc235027c70aa2fcb7ca46c35d29fbf97cf
---

# 4. Quantization, model formats & architectures

## 4. Quantization, model formats & architectures

- **Supported (README):** FP4 / FP8 / INT4 / AWQ / GPTQ [official]. v0.5.19 release notes additionally exercise: **NVFP4, MXFP4, MXFP8, W4A8, W4A4 (MegaMoE), FP8 KV cache, mxfp8 KV cache, compressed-tensors** (quantized lm_head, kv_cache_scheme scales), torchao integration, ModelSlim (Kimi-K3) [official].
- **Marlin:** FP4 Marlin helpers were deduped in the v0.5.20 quantization refactor; the old **AWQ AOT kernel and non-Marlin GPTQ were deleted** in v0.5.20 — Marlin remains the path for GPTQ/AWQ-style INT4 [official]. Community config translation confirms `--quantization gptq_marlin` is auto-detected [secondary](https://github.com/randomchaos7800-hub/inference-research/blob/HEAD/tower/gdn-blackwell/sglang-vs-vllm-sm120.md).
- **NVFP4 MoE backends:** as of v0.5.16, the buggy `cutlass` FP4 path was deleted; `triton` and `flashinfer_cutlass` are the viable backends; FlashInfer CuTe DSL NVFP4 W4A16 mode added v0.5.19 (Qwen3-30B-A3B W4A16 GSM8K 0.965 → 0.980) [official].
- **FP8:** FP8 KV cache; FP8 DeepEP dispatch (Humming backend: TPOT −35%, TTFT −23% vs BF16 on H20, v0.5.19); FP8 block-scale MoE keeps fp32 routing weights (v0.5.20) [official].
- **Model coverage:** Llama, Qwen (incl. Qwen3.x, Qwen3.5/3.6/3.8, Qwen3-VL), DeepSeek (V3/R1/V4 + MLA/DSA/MTP), Kimi (K2/K3/Linear), GLM (4.5/4.6/5.x, incl. GLM-5.2/5.3-Flash NVFP4), GPT (gpt-oss day-0 Aug 2025), Gemma (incl. Gemma-4 FP8 MTP), Mistral (incl. Mistral Large 3, Ministral), Nemotron 3 (Nano/Super/Ultra/Lightning), LFM2/2.5, MiniMax (M2/H3/M3, incl. Music 3), Spark2.5, Granite 4.2, LongCat, Ling-3.0, dots3.note, MiniCPM (incl. MiniCPM5/SALA), Intern-S2-Mobius, Mamba-hybrid models (Qwen3.8-Next, Nemotron-H, Kimi-Linear), embedding models (e5-mistral, gte, mcdse), reward models (Skywork), diffusion models (WAN, Qwen-Image, FLUX.2, HunyuanVideo, LTX-2.5, SANA-Video, Cosmos3, MiniMax-H3, LongCat-Image, SenseNova-U1.5-8B-MoT) [official, across release notes + README].
- **Model formats:** Hugging Face checkpoints (safetensors), GGUF (diffusion transformer checkpoints, MiniMax-H3, v0.5.19), ComfyUI serialized checkpoints (NVFP4/W4A8/INT8 DiTs), pruned safetensors [official]. Transformers upgraded 4.57.1 → 5.3.0 in v0.5.10 [official].

---

## 5. New 2026 features (beyond §2)

- **Structured outputs / grammar-constrained generation:** XGrammar integration since v0.4 (Dec 2024, "Faster Structured Outputs") [official, lmsys.org/blog/2024-12-04-sglang-v0-4/]; compressed-FSM JSON decoding claimed 3× faster (Feb 2024 blog) [official]. v0.5.19: MistralCommon tokenizer support in the XGrammar backend [official]. Structured-generation frontend (`sgl.function`, `sgl.gen`, choices/regex/JSON-schema) is the project's original 2023–2024 DSL [official README "Frontend Tutorial"]. Agent-workload claim: <400 ms p50 for schema-constrained hops [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison).
- **Reasoning / multi-step:** the SGLang **frontend language** includes a program interpreter, few-shot generation primitives, and tool/chain composition (the "Structured Generation Language" of its name) [official README]; `--reasoning-parser` emits `delta.reasoning_content` in chat completions [secondary]; `reasoning_tokens` usage accounting fixed v0.5.10 [official]. For RL: `return_sampling_mask` (v0.5.20) and `DisallowedTokensLogitsProcessor` support [official].
- **Multi-modal:** LLaVA-OneVision multi-image/video (v0.3), Pixtral multi-image, Qwen3-VL MoE, PaddleOCR-VL (1080p page TTFT 235 ms → 120 ms on H200, v0.5.19), ASR/audio/speech fixes, `input_audio` content part in chat completions (v0.5.19) [official]. **SGLang-Omni** (sister repo): audio/TTS/ASR serving framework — v0.1.6 on PyPI (Sept 2026), day-0 AuK/AuK-Flash and MiniMax Music 3, MOSS-TTS, Higgs Audio v3 [official, repo news](https://github.com/sgl-project/sglang-omni).
- **SGLang-Diffusion:** image/video generation serving (WAN, Qwen-Image, FLUX.2, HunyuanVideo, LTX-2.5, Cosmos3, MiniMax-H3, SANA-Video, LongCat); up to 2.5× faster than its Nov-2025 initial release by Jan 2026; up to 1.5× across major diffusion models in v0.5.8 [official, lmsys.org/blog/2026-01-16-sglang-diffusion/].
- **DeepEP / MoE expert parallelism:** see §3.3. SGLang was "first opensource implementation of DeepSeek V3/R1 expert parallelism with PD disaggregation" (May 2025) [secondary].
- **CUDA graphs:** see §3.6.
- **FlashInfer attention backend:** required 0.6.18 since v0.5.19; powers MXFP4/MXFP8 MoE, DSA top-k, NVFP4 W4A16, MegaMoE [official].
- **Multi-LoRA:** multi-adapter LoRA batching incl. with EAGLE/NEXTN/DFLASH/DSPARK speculative decoding (v0.5.19); LoRA weight-load overlap −78% TTFT (v0.5.9) [official].
- **Embeddings/rerank/reward models** supported (e5-mistral, gte, mcdse, Skywork) [official].

---

## 6. Performance benchmarks

### 6.1 Official (project-published) numbers
- **v0.2 blog (2024-07-25):** faster Llama-3 serving vs TensorRT-LLM and vLLM [official].
- **v0.3 blog (2024-09-04):** 7× faster DeepSeek MLA, 1.5× faster torch.compile [official].
- **v0.4 blog (2024-12-04):** zero-overhead scheduler, cache-aware load balancer, faster structured outputs [official].
- **Large-scale EP (2025-05-05):** 52.3K in-tok/s + 22.3K out-tok/s on 96 GPUs, 5× vs vanilla TP [official].
- **GB200 NVL72 Part I (2025-06-16):** 2.7× higher decoding throughput; **Part II (2025-09-25):** 3.8× prefill, 4.8× decode throughput (PD + large-scale EP) [official].
- **GB300 NVL72 (2026-02-19):** "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [official].
- **GLM-5.2 agentic workloads (2026-07 blog):** 500 TPS in two weeks on NVFP4 [official].
- Per-release kernel numbers are in §2 (e.g., v0.5.19: Cohere Command-A-Plus decode 2.28× at batch 1, up to 9,152 tok/s on 4×GB300; v0.5.20: DCP1→DCP-N verified 8×B300 to 256K).

### 6.2 Independent measurements
- **SemiAnalysis InferenceX — Qwen3.5-397B-A17B FP8 on MI355X** (published ~2026-05, runs 2026-02-20 → 2026-05-19): throughput/GPU on 8k/1k workload went **192 → 3,660 tok/s/GPU (19.0× at iso-interactivity 40 tok/s/user)** across three SGLang releases (v0.5.8.post1 → v0.5.10rc0 → v0.5.12) plus three AITER MoE kernel landings; another ~1.5× from the May image bump [independent](https://github.com/semianalysisai/inferencex-app/blob/HEAD/packages/app/content/blog/mi355x-qwen3-5-sglang-v0-5-12-up-to-17x.mdx). (Note: one related InferenceX post wrote "SGLang v0.12" for its GLM-5 run — likely a typo; treated as [unverified] version label.)
- **SemiAnalysis InferenceX — GLM-5 FP8 cost (2026-05-20 run):** MI355X+SGLang undercuts B200+SGLang on $/M tokens by up to **1.41× at 18 tok/s/user with MTP ($0.30/M vs $0.22/M — a ~40% reduction)** and 1.36× without MTP; B200 noses back ahead above ~90 tok/s/user [independent](https://github.com/semianalysisai/inferencex-app/blob/HEAD/packages/app/content/blog/mi355x-glm5-fp8-sglang-40-cheaper-than-b200.mdx).
- **dstack — DeepSeek-R1 on 8×H200 SXM5:** online (SGLang `bench_serving.py`, concurrencies 4–128, 3200-in/800-out) and offline inference benchmark comparing SGLang vs TensorRT-LLM; full results in their repo [independent](https://github.com/dstackai/dstack/blob/HEAD/docs/blog/posts/h200-mi300x-deepskeek-benchmark.md).

### 6.3 vLLM vs SGLang comparisons (all secondary/independent — methodology varies)
- **PremAI 2026 benchmark (H100 80GB, Llama 3.1 8B):** SGLang ~16,200 tok/s vs vLLM ~12,500 tok/s — **SGLang +29%**; LMDeploy ~16,100 (tie) [secondary](https://blog.premai.io/vllm-vs-sglang-vs-lmdeploy-fastest-llm-inference-engine-in-2026/) via [secondary summary](https://github.com/profsynapse/synaptic-tuner/blob/HEAD/docs/preparation/vllm-vs-sglang-inference-serving-research.md).
- **Spheron (Llama 3.3 70B FP8, H100):** SGLang +29% total throughput, **+117% output-token throughput** (894 vs 413 tok/s), TTFT 79 vs 103 ms (−23%), ITL 6.0 vs 7.1 ms (−15%) [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison). Same source: with **unique prompts (no shared prefixes) the gap shrinks to near-zero** — concurrency 1→100 shows only +2–5%.
- **Community single-request test (old, ~2025-06):** vLLM 60.0 tok/s vs SGLang 52.7 tok/s on a single unique prompt — vLLM ~1.1× faster [secondary](https://github.com/brendanmckeag/sglang-vllm-benchmark/blob/HEAD/README.md). Dated; treat as historical.
- **DeepSeek-specific:** SGLang claimed **3.1× faster than vLLM on DeepSeek-V3** via optimized MLA backends (FA3/FlashInfer/FlashMLA/CutlassMLA) [secondary](https://particula.tech/blog/sglang-vs-vllm-inference-engine-comparison).
- **Consumer GPUs (RTX PRO 6000, Sept 2026, community):** Qwen3.8-27B NVFP4 + DSpark: 119.19 tok/s output, 1,072.69 tok/s total (8K-in/1K-out, 8 reqs), mean TTFT 729.82 ms, mean TPOT 7.68 ms, DSpark accept length 2.50 [secondary/independent](https://github.com/lEWFkRAD/qwen38-rtx-pro-6000).
- **Artificial Analysis:** no Artificial-Analysis-published SGLang-vs-vLLM engine benchmark was found in this research pass ([unverified]/not found). SemiAnalysis InferenceX (above) is the closest independent price/performance source.
- **ServeTheHome:** no STH SGLang benchmark article was surfaced in this pass ([unverified]/not found).
- **SiliconANGLE:** no SiliconANGLE SGLang article was surfaced in this pass ([unverified]/not found).

### 6.4 Caveats
- The +29% SGLang-vs-vLLM gap is a **high-concurrency batching** result; single-stream unique-prompt workloads show ~0–12% either way. RadixAttention's advantage concentrates in prefix-sharing workloads (agents, RAG, multi-turn chat) [secondary].
- One community SM120 report (Aug 2026) flagged garbage output with INT4 quantized models on SGLang (issue #21132) and recommended vLLM for that config; upstream fixed NVFP4 paths in v0.5.16+ and improved SM120 support in v0.5.19/0.5.20 — current status [unverified] [secondary](https://github.com/randomchaos7800-hub/inference-research/blob/HEAD/tower/gdn-blackwell/sglang-vs-vllm-sm120.md).

---

