---
id: ai-industry-kb-2026/08-quantization-model-formats/exl2-exllamav3-exl3-end-of-an-era
title: "EXL2 → ExLlamaV3 / EXL3: end of an era"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["Alibaba", "Nvidia", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-07-14", "2026-08-27", "2026-08-31", "2026-09-08", "2026-09-09"]
keywords: ["llama", "benchmark", "benchmarks", "blackwell", "consumer", "decode", "distribution", "fp4", "fp8", "gguf", "gptq", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3863, 3892]
section: "8. Quantization & Model Formats"
sha256: 78b498d8ec46e7228f31539dae575a39376c0799d4bba5fa060f79c1bb0e50c1
---

# EXL2 → ExLlamaV3 / EXL3: end of an era

- Wave 2's open item ("no independent replication of the Divergence-300/KL metrics") is now partially addressed by a community benchmark (srmiles/local-llm-benchmarks, 2026-08-27): Dynamic v3.0 vs v2.0 on Qwen3.8-27B, matched quant names, real hardware, Q8_0 quant as referee.
- Verdict: **"yes on accuracy, no on throughput, and one file is unusable."** v3.0 is measurably closer to the Q8_0 referee on mean, median and tail KLD and on RMS Δp, *while being smaller* — corroborating the direction of Unsloth's ">10% better top-1" claim (KLD-based, not benchmark-score-based).
- But it is **not uniformly better**: it loses on 99.0% Δp at both sizes, and on top-1 agreement at 4-bit (consistent with Wave 2's note that v3.0's published top-1 agreement figures are graph readings only).
- v3.0 **decodes 9.5% slower** at Q4_K_XL (pooled medians over an interleaved A/B, Welch t = 4.05), alongside a much more heterogeneous block-type mix — the accuracy gain trades against decode speed.
- [COMMUNITY] **v3.0's `UD-Q3_K_XL` wedges llama-server**: a token id of `-1` reaches the batch, initialization fails, and with a drafter attached every subsequent request 500s until restart. Unsloth re-published the Qwen3.8-27B-GGUF files after launch ("a new version of Dynamic v3.0," per the HF discussion), so this file's provenance matters — may be fixed in later re-uploads.
- Net: direction corroborated on KL divergence with independent hardware; Divergence-300 specifically still unreplicated; decode-speed regression (~9.5% at Q4_K_XL) measured; one UD variant shipped with a serving-breaking defect.

### EXL2 → ExLlamaV3 / EXL3: end of an era

- **ExLlamaV2 is archived** — the README states: "This project is archived for now. Development continues on ExLlamaV3." The 2026 format landscape treats EXL2 as legacy: existing EXL2 checkpoints (e.g. LoneStriker's `6.0bpw` series, thousands of HF repos) still run under TabbyAPI, but no new development lands on the V2 codebase.
- EXL2's technical signature: same optimization method as GPTQ, 2/3/4/5/6/8-bit support, **any average bitrate between 2 and 8 bpw** via per-layer and per-column mixed precision — the converter quantizes each matrix several ways, measures error against calibration data, and picks the combination minimizing worst-case error at the target bitrate. Hence filenames like `4.65bpw` rather than `4-bit`. EXL2 renames tensors so every model looks like a Llama variant internally — the format never ported to other frameworks (TabbyAPI is the official server, OpenAI-compatible API).
- **ExLlamaV3** (turboderp-org/exllamav3) is the active successor — "optimized quantization and inference library for running LLMs locally on modern consumer-class GPUs." **v1.4.5 published 2026-08-31** with prebuilt Windows wheels (cp310–cp313, cu128/torch 2.10.0); September brought a fast bug-fix cadence **v1.4.6 → v1.4.9** (CUDA-graph capture fixes, devctx workspace hardening, COOP_AUTOTUNE_VERSION 3→4; no format changes). **v1.5.0 referenced** with fused routed-expert launch (no speed change measured: 52.05 vs 52.22 tok/s on GB10).
- **EXL3** is a streamlined variant of **QTIP** (Cornell RelaxML): one-step quantization (fused Viterbi kernel), minutes for small models on a single RTX 4090, **coherent at 1.6 bpw** [VENDOR] (Llama-3.1-70B in <16 GB). The GEMM kernel is Marlin-inspired and **memory-bound at 4 bpw on RTX 4090**; v1.0.0 (2026-07-14) greatly improved GEMM/GEMV on Ampere and dropped the `flash-attn-2`/`xformers` dependencies. V3 adds: tensor-parallel + expert-parallel for consumer setups, 2–8-bit KV cache quantization, speculative decoding, LoRA, multimodal support, v1.4.3 partial CPU expert offload, v1.4.4 vision-tower quantization + vision offload from system RAM, v1.4.5 improved MoE MTP performance.
- Verdict from the 2026 inference-engine shootouts: ExLlamaV3 is **the only credible raw-throughput challenger to llama.cpp on consumer NVIDIA GPUs** (memory-bound GEMM design) — sharpened by the September shootout below to **single-stream short-context decode only** — but EXL3 has no HF repo-store distribution: GGUF remains the portable choice and llama.cpp the baseline.
- The **vLLM-EXL3 serving plugin** (vcruz305/vllm-exl3, v0.4.2, 2026-09-09) documents real serving-integration fixes: dense EXL3 calls no longer launch exllamav3's cooperative trellis GEMM (rows 17–144 take the reconstruct+hgemm path) — on the vLLM nightly V2 model runner the cooperative GEMM **wedged the engine**; a 4-worker MTP stress that wedged within 3 minutes ran clean for 45 minutes after the routing fix, with decode speed unchanged. EXL3-in-vLLM is real but friction-heavy (third-party plugin).
- [COMMUNITY] Production datapoint: Qwen3.8-Flash-Next EXL3 4.05bpw on 4× RTX 3090 via TabbyAPI — 108 tok/s single-stream decode, 144.5 tok/s eight-stream aggregate, three concurrent 260k-token sessions in 96 GB VRAM; tool-calling quality 89/100 vs 88/100 baseline (inside ±2.5 noise).

### vLLM (NVFP4/FP8) vs ExLlamaV3 (EXL3) on RTX PRO 6000 (Sept 2026) [COMMUNITY]

- The most complete public format/engine shootout: MarcoPizeta's reproducible benchmark (04–08 Sept 2026), Qwen3.8-Flash-Next (125B-A6B + 51B n-gram/PLE table) on a single RTX PRO 6000 Blackwell 96 GB — vLLM with `primitive-ai/Qwen3.8-Flash-Next-mixed-NVFP4-FP8` + INT4 PLE offload + MTP vs ExLlamaV3/TabbyAPI with `turboderp/Qwen3.8-Flash-Next-exl3` (4.05bpw; 5.05bpw also tested). A v3 extension (08/09/2026) added SGLang with the calibrated `RadixArk/Qwen3.8-Flash-Next-NVFP4` checkpoint (`--quantization modelopt_fp4`, `--fp4-gemm-backend flashinfer_cutlass`, NEXTN speculative steps 3).
- Decode throughput (mean of 3 runs; same prompts per engine):

| Input × concurrency | vLLM MTP on | EXL3 4.05 MTP on |
|---|---|---:|
| 1k × 1 | 101.6 tok/s (TTFT 259 ms) | **168.9** tok/s (283 ms) |
| 1k × 4 | **296.4** tok/s (179 ms) | 221.4 tok/s (638 ms) |
| 8k × 1 | **140.5** tok/s (193 ms) | 129.1 tok/s (527 ms) |
| 8k × 4 | **282.1** tok/s (368 ms) | 125.2 tok/s (1466 ms), ok 56/60 |
| 32k × 4 | **150.7** tok/s (1412 ms) | 16.1 tok/s (5270 ms), **ok 21/60** |

