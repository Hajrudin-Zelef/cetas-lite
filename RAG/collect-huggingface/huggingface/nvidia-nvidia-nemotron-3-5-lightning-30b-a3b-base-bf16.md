---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16
title: "NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia"]
dates: ["2026-08-11", "2026-09-23"]
keywords: ["nvidia", "attention", "benchmarks", "distillation", "fine-tuning", "license", "moe", "nvfp4", "parameters", "reasoning", "research", "training"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16.md
source_anchor: ""
source_lines: [1, 47]
sha256: f2af46126dd647e6b763e83926735b7dd6b1e52c25d11af1ed94037a873dd4dc
---

# NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 is the base (pre-trained) checkpoint of the Nemotron 3.5 Lightning family, released by NVIDIA without supervised fine-tuning, reinforcement learning or distillation, making it a natural starting point for developers building their own post-trained models. It uses a hybrid Mixture-of-Experts (MoE) architecture with interleaved Mamba-2 and MoE layers plus select Attention layers, incorporating Multi-Token Prediction (MTP) layers trained via a dedicated continued pre-training phase. The model has 30B total parameters and 3B active parameters, was pre-trained on over 20T tokens using an NVFP4 recipe, and supports up to 1M tokens of context. The pre-training corpus spans English, 19 other spoken languages and 43 programming languages. Base-model benchmarks include MMLU 78.59, MMLU-Pro (5-shot) 67.94, GSM8K (8-shot) 91.28, Minerva Math 82.78, MBPP (3-shot) 78.59, HumanEval 77.44, ARC-Challenge 92.66, Global-MMLU-Lite average 75.53, and RULER 256K 76.88 / 1M 69.62. It is intended for pre-training research, continued pre-training on domain corpora, and building custom post-trained variants via NeMo RL, NeMo Gym and Megatron-LM. It is governed by the OpenMDW License Agreement version 1.1 and is ready for commercial and non-commercial use. Release date August 11, 2026. The post-trained versions are the BF16 and NVFP4 Lightning reasoning models.

## Key points

- Base (pre-trained) checkpoint of Nemotron 3.5 Lightning: 30B total / 3B active MoE.
- Hybrid Mamba-2 + MoE + Attention architecture with Multi-Token Prediction (MTP).
- Pre-trained on 20T+ tokens with NVFP4 recipe; up to 1M context.
- Pre-training corpus: English, 19 other languages, 43 programming languages.
- Base benchmarks: MMLU 78.59, MMLU-Pro 67.94, HumanEval 77.44, RULER 1M 69.62.
- Ideal for continued pre-training, SFT, RL and distillation research.
- License: OpenMDW-1.1; release date 2026-08-11.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 30B |
| Active parameters | 3B |
| Architecture | Mamba2-Transformer Hybrid MoE + MTP |
| Context length | Up to 1M tokens |
| Pre-training tokens | 20T+ |
| Precision | BF16 (trained with NVFP4 recipe) |
| MMLU | 78.59 |
| MMLU-Pro (5-shot) | 67.94 |
| GSM8K (8-shot) | 91.28 |
| HumanEval | 77.44 |
| MBPP (3-shot) | 78.59 |
| RULER 256K / 1M | 76.88 / 69.62 |
| License | OpenMDW-1.1 |
| Release date | 2026-08-11 |

## Why this source matters for the RAG

This card documents the base checkpoint of a major NVIDIA model family, providing the pre-trained foundation data (architecture, training corpus, base benchmarks) needed to understand the family's lineage. It anchors continued-training and customization workflows in the knowledge base.
