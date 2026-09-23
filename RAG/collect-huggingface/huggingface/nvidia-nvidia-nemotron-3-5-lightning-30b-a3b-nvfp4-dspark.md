---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark
title: "NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-08-11", "2026-09-23"]
keywords: ["nvfp4", "nvidia", "agentic", "attention", "benchmarks", "distribution", "embedding", "gqa", "inference", "latency", "license", "llama"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark.md
source_anchor: ""
source_lines: [1, 46]
sha256: 681f3c7cb2f55a2634d8468d6e993effb16f1db840b67d37143c74061f000511
---

# NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository hosts the DSpark speculative decoding checkpoint for NVIDIA's Nemotron-3.5-Lightning-30B-A3B model family. DSpark is a semi-autoregressive speculative-decoding drafter that proposes a whole block of candidate tokens in a single forward pass using a parallel backbone, and this release is specifically tuned for DGX Spark (GB10) and low-concurrency data-centre workflows. It is not a standalone target model: it must be paired with the Nemotron-3.5-Lightning-30B-A3B target model in vLLM to accelerate inference. The drafter architecture is a dense GQA model (Dense MLP + GQA Attention) with 967M total parameters, of which 615M are non-embedding parameters; its causal GQA attention uses a sliding window of size 1024 on all layers plus a per-head attention sink bias. It was trained on 66B tokens (2 epochs) synthesized from prompts of the Nemotron-Post-Training-Dataset-v2 and v3 collections, and quantized with Model Optimizer 0.45.0. It supports up to 1M tokens of context, and languages English, Spanish, French, German, Italian and Japanese. It runs via vLLM and llama.cpp, validated on H100, GB200, RTX 5090 and DGX Spark. Evaluation on SPEED-Bench (draft length 7) gives an overall average acceptance length of 3.75 across coding (4.38), math (4.17), multilingual (4.55), RAG (4.25), reasoning (3.90), QA (3.36), stem (3.40), summarization (4.15), humanities (3.18), roleplay (3.06) and writing (2.83). It is lossless for the target distribution: the drafter only accelerates inference without changing output distribution. It is governed by the OpenMDW-1.1 license and available for commercial and non-commercial use. Release date August 11, 2026. Paper: arXiv:2607.05147.

## Key points

- DSpark speculative-decoding drafter for Nemotron-3.5-Lightning-30B-A3B, tuned for DGX Spark and low-concurrency serving.
- Dense GQA architecture: 967M total / 615M non-embedding parameters; sliding window 1024 + attention sink bias.
- Not a standalone model; must be paired with the target Lightning checkpoint in vLLM.
- Trained on 66B tokens (2 epochs) of synthesized post-training prompts; quantized with Model Optimizer 0.45.0.
- 1M token context; runs on vLLM and llama.cpp; H100, GB200, RTX 5090, DGX Spark.
- SPEED-Bench average acceptance length 3.75 (draft length 7); lossless acceleration.
- License: OpenMDW-1.1; 214,908 downloads last month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Purpose | Speculative decoding drafter (DSpark) |
| Total parameters | 967M |
| Non-embedding parameters | 615M |
| Architecture | Dense GQA (Dense MLP + GQA Attention) |
| Attention | Causal GQA, sliding window 1024, attention sink bias |
| Context length | Up to 1M tokens |
| Training data | 66B tokens, 2 epochs |
| Precision | NVFP4 (Model Optimizer 0.45.0) |
| Runtimes | vLLM, llama.cpp |
| Hardware | H100, GB200, RTX 5090, DGX Spark (GB10) |
| SPEED-Bench avg acceptance length | 3.75 (draft length 7) |
| License | OpenMDW-1.1 |
| Release date | 2026-08-11 |

## Why this source matters for the RAG

This card documents the speculative-decoding companion to a major NVIDIA agentic model, providing concrete acceptance-rate benchmarks and architecture details for inference-acceleration technology. It illustrates how draft models enable lossless speedups, which is valuable for deployment and latency optimization knowledge.
