---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4
title: "NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "Meta", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2025-06", "2026-02", "2026-03", "2026-09-23"]
keywords: ["nvfp4", "nvidia", "agentic", "agents", "attention", "benchmark", "benchmarks", "embeddings", "fp8", "gpu", "license", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4.md
source_anchor: ""
source_lines: [1, 58]
sha256: 2e351f7f21babdd9f9babac02a886e083a5df4e8e4f2291aecfebdaa6e94cdc4
---

# NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable (via raw README, main page too large for full fetch)
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 is a large language model trained by NVIDIA, designed for strong agentic, reasoning and conversational capabilities, optimized for collaborative agents and high-volume workloads such as IT ticket automation. It follows the standard Nemotron pattern: it first generates a reasoning trace then concludes with a final response; reasoning can be toggled on/off via the chat template (`enable_thinking=True/False`), with a low-effort reasoning mode also supported.

**Architecture.** The model uses the hybrid **LatentMoE** architecture — interleaved Mamba-2 and MoE layers with select Attention layers — where tokens are projected into a smaller latent dimension for expert routing and computation, improving accuracy per byte. Unlike the Nano model, Super incorporates **Multi-Token Prediction (MTP)** layers for faster generation and improved quality. It is the first Nemotron 3 model pre-trained directly in **NVFP4** quantization (native 4-bit FP), with select layers (latent projections, MTP, QKV/attention, embeddings) kept in BF16 or MXFP8 for training stability. Total: **120B parameters, 12B active**. Context length up to **1M tokens**. Minimum hardware: **1x B200 or 1x DGX Spark**. Supported languages: English, French, German, Italian, Japanese, Spanish, Chinese. Ready for commercial use under the NVIDIA Nemotron Open Model License.

**Training.** Pre-trained on ~25T tokens (base model over 25T tokens) using crawled + synthetic code/math/science/general data with NVFP4 training efficiency; then SFT on synthetic code/math/science/tool-calling/instruction-following/long-range retrieval data; then multi-environment RL using asynchronous GRPO across math, code, science, tool use and structured outputs, refined with RLHF. Post-training data cutoff: February 2026; pre-training data cutoff: June 2025. Released 11 March 2026.

**Benchmarks (NVFP4 variant).** MMLU-Pro 83.33; HMMT Feb25 (with tools) 95.36; GPQA (no tools) 79.42; LiveCodeBench v6 78.44; SciCode 40.83; HLE (no tools) 17.42; Terminal Bench (hard) 24.48; TauBench V2 average 60.46; IFBench (prompt) 73.30; Scale AI Multi-Challenge 52.8; Arena-Hard-V2 76.00; AA-LCR 58.06; RULER-500 @128k 95.99 / @256k 96.52 / @512k 96.23; MMLU-ProX 79.37. Note: some scores slightly above or at parity with the BF16/FP8 variants.

**Deployment.** Requires the custom `super_v3` reasoning parser. Backends supported: vLLM (0.20.0), SGLang, and TRT-LLM, all with MTP-based speculative decoding options; NVFP4 GEMM via marlin backend on DGX Spark. OpenCode integration documented via `opencode.json` config with context limit 1M / output 32768. Recommended sampling: temperature 1.0, top_p 0.95.

## Key points

- Hybrid LatentMoE (Mamba-2 + MoE + Attention) with Multi-Token Prediction; 120B total / 12B active.
- First Nemotron 3 model pre-trained in NVFP4 (native 4-bit FP), ~4x smaller weights than BF16.
- Up to 1M token context; runs on a single B200 or DGX Spark.
- Toggleable reasoning (enable_thinking) plus low-effort reasoning and reasoning_budget control.
- Strong long-context retrieval (RULER @512k = 96.23) and multilingual performance.
- NVIDIA Nemotron Open Model License; commercial use ready; released 11 Mar 2026.
- OpenAI-compatible serving via vLLM / SGLang / TRT-LLM; OpenCode integration documented.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 120B |
| Active parameters | 12B |
| Architecture | LatentMoE: Mamba-2 + MoE + Attention hybrid with MTP |
| Context length | Up to 1M tokens |
| Precision | NVFP4 (select layers BF16 / MXFP8) |
| Minimum GPU | 1x B200 or 1x DGX Spark |
| Languages | en, fr, de, it, ja, es, zh |
| License | NVIDIA Nemotron Open Model License |
| Release date | 11 March 2026 |
| Recommended sampling | temperature 1.0, top_p 0.95 |
| MMLU-Pro | 83.33 |
| HMMT Feb25 (tools) | 95.36 |
| GPQA (no tools) | 79.42 |
| LiveCodeBench v6 | 78.44 |
| Terminal Bench (hard) | 24.48 |
| Arena-Hard-V2 | 76.00 |
| RULER-500 @512k | 96.23 |

## Why this source matters for the RAG

Documents the NVFP4 deployment variant of NVIDIA's frontier open-weight Nemotron 3 Super — the exact checkpoint intended for a single B200 or DGX Spark. Includes verified architecture, training, benchmark and deployment data, useful for accurate, up-to-date guidance on running a 120B/12B frontier open model locally.
