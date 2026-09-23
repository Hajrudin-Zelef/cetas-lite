---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16
title: "NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-06-04", "2026-09-23"]
keywords: ["nvidia", "agentic", "agents", "attention", "benchmark", "benchmarks", "distillation", "gpu", "license", "moe", "nvfp4", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16.md
source_anchor: ""
source_lines: [1, 47]
sha256: f6bf2a6ed6a74aa2faf6f9f3d78b4bd28cde317b9ba4d3fef570991a05523ccd
---

# NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 is a frontier-scale large language model trained by NVIDIA for the most demanding workloads: complex multi-step agents, long-context analysis, and high-accuracy reasoning over code, math and science. It uses a hybrid Latent Mixture-of-Experts (LatentMoE) architecture with interleaved Mamba-2 and MoE layers plus select Attention layers, and Multi-Token Prediction (MTP) layers with a shared-weight design across prediction heads. The model has 550B total parameters and 55B active parameters, supports up to 1M tokens of context, and supports English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese and Chinese. It was pre-trained for approximately 20T tokens using an NVFP4 recipe and post-trained through SFT, asynchronous GRPO reinforcement learning, and Multi-Domain On-Policy Distillation (MOPD). Reasoning is configurable via the chat template (`enable_thinking`). Minimum hardware is 8x GB200/B200/GB300/B300, 16x H100 or 8x H200; the NVFP4 variant runs on a smaller footprint. Benchmarks include Terminal Bench 2.1 56.4, SWE-bench Verified 70.7, SWE-bench Multilingual 67.7, TauBench V3 average 70.9, BrowseComp 44.4, LiveCodeBench v6 89.0, IMOAnswerBench 88.6 (no tools) / 92.3 (with tools), GPQA 87.0 (no tools), HLE 26.7 (no tools) / 37.4 (with tools), MMLU-Pro 86.8, IFBench 81.7, RULER (1M) 94.7, LongBench v2 61.9, and MMLU-ProX (10 langs) 83.0. Deployment is via vLLM (v0.22.0), SGLang (v0.5.12.post1) and TensorRT-LLM, with multi-node Ray support. It is governed by the OpenMDW-1.1 license and is ready for commercial and non-commercial use. Release date June 4, 2026.

## Key points

- Frontier-scale NVIDIA model: 550B total / 55B active LatentMoE parameters with MTP.
- Up to 1M context; 10 supported languages; reasoning configurable via `enable_thinking`.
- Pre-trained on ~20T tokens with NVFP4 recipe; MOPD distillation for accuracy.
- Min hardware: 8x GB200/B200/GB300/B300, 16x H100, or 8x H200; NVFP4 variant smaller.
- Benchmarks: SWE-bench Verified 70.7, LiveCodeBench v6 89.0, GPQA 87.0, MMLU-Pro 86.8, RULER 1M 94.7.
- Runtimes: vLLM, SGLang, TensorRT-LLM with multi-node Ray orchestration.
- License: OpenMDW-1.1; release date 2026-06-04.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 550B |
| Active parameters | 55B |
| Architecture | LatentMoE (Mamba-2 + MoE + Attention) + MTP |
| Context length | Up to 1M tokens |
| Languages | English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, Chinese |
| Precision | BF16 (trained with NVFP4 recipe) |
| Min GPU | 8x GB200/B200, 16x H100, 8x H200 |
| SWE-bench Verified | 70.7 |
| LiveCodeBench v6 | 89.0 |
| GPQA (no tools) | 87.0 |
| MMLU-Pro | 86.8 |
| RULER (1M) | 94.7 |
| License | OpenMDW-1.1 |
| Release date | 2026-06-04 |

## Why this source matters for the RAG

This card documents NVIDIA's largest open model, providing frontier-scale architecture, benchmark and hardware data that anchors the top end of the open-weight model comparison space. It is essential for high-stakes agentic and long-context RAG deployment decisions.
