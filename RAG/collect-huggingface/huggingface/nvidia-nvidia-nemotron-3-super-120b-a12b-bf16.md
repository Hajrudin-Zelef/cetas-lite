---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16
title: "NVIDIA-Nemotron-3-Super-120B-A12B-BF16 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "China", "Hugging Face", "Nvidia"]
dates: ["2026-03-11", "2026-09-23"]
keywords: ["nvidia", "agentic", "agents", "attention", "benchmark", "benchmarks", "compute", "fine-tuning", "gpu", "license", "moe", "nvfp4"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3-Super-120B-A12B-BF16.md
source_anchor: ""
source_lines: [1, 48]
sha256: 84e57c57614fc47c36e120e4022c2ccf9046862aa40635d76036232e94ee0582
---

# NVIDIA-Nemotron-3-Super-120B-A12B-BF16 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3-Super-120B-A12B-BF16 is a large language model trained by NVIDIA for strong agentic, reasoning and conversational capabilities, optimized for collaborative agents and high-volume workloads such as IT ticket automation. It uses a hybrid Latent Mixture-of-Experts (LatentMoE) architecture with interleaved Mamba-2 and MoE layers plus select Attention layers, and incorporates Multi-Token Prediction (MTP) layers for faster generation via native speculative decoding and improved quality. It has 120B total parameters and 12B active parameters, supports a context length of up to 1M tokens (default configuration 256K), and supports English, French, German, Italian, Japanese, Spanish and Chinese. The model generates a reasoning trace before the final answer, configurable via the chat template (`enable_thinking=True/False`) with an additional low-effort mode. It was pre-trained for over 25T tokens and is the first model in the Nemotron 3 family pre-trained using NVFP4 quantization to maximize compute efficiency, with select layers kept in BF16 or MXFP8 for stability. Training followed three stages: pre-training (25T+ tokens), supervised fine-tuning, and multi-environment reinforcement learning with asynchronous GRPO plus RLHF. Recommended sampling is temperature 1.0 and top_p 0.95 across all tasks and backends. Minimum GPU requirement is 8× H100-80GB; the NVFP4 variant runs on a single B200 or DGX Spark. Benchmarks include MMLU-Pro 83.73, AIME25 90.21, GPQA 79.23 (no tools) / 82.70 (with tools), LiveCodeBench v5 81.19, SWE-Bench OpenHands 60.47, RULER @1M 91.75, and TauBench V2 average 61.15. It is governed by the NVIDIA Nemotron Open Model License and is ready for commercial use. Release date March 11, 2026.

## Key points

- NVIDIA flagship agentic/reasoning LLM: 120B total / 12B active parameters.
- Hybrid LatentMoE: Mamba-2 + MoE + Attention with Multi-Token Prediction (MTP).
- Context length up to 1M tokens (default 256K); 7 supported languages.
- First Nemotron 3 model pre-trained with NVFP4 quantization; commercial use allowed.
- Reasoning configurable via `enable_thinking` plus low-effort mode.
- Min GPU 8× H100-80GB; NVFP4 variant runs on one B200/DGX Spark.
- Strong long-context: RULER @1M 91.75; SWE-Bench OpenHands 60.47; AIME25 90.21.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 120B |
| Active parameters | 12B |
| Architecture | LatentMoE (Mamba-2 + MoE + Attention) with MTP |
| Context length | Up to 1M tokens (default 256K) |
| Languages | English, French, German, Italian, Japanese, Spanish, Chinese |
| Precision | BF16 (trained with NVFP4) |
| Minimum GPU | 8× H100-80GB |
| License | NVIDIA Nemotron Open Model License |
| MMLU-Pro | 83.73 |
| AIME25 (no tools) | 90.21 |
| GPQA (no tools / tools) | 79.23 / 82.70 |
| LiveCodeBench v5 | 81.19 |
| SWE-Bench (OpenHands) | 60.47 |
| RULER @ 1M | 91.75 |
| Release date | 2026-03-11 |

## Why this source matters for the RAG

This card documents NVIDIA's largest Nemotron 3 Super model, giving authoritative architecture, long-context and agentic benchmark data for enterprise deployment. It is central to comparing frontier open-weight models across the Qwen, Gemma, gpt-oss and Nemotron families.
