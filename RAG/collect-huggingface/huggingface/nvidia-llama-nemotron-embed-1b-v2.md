---
id: collect-huggingface/huggingface/nvidia-llama-nemotron-embed-1b-v2
title: "llama-nemotron-embed-1b-v2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-09-23"]
keywords: ["llama", "embedding", "embeddings", "fine-tuning", "license", "nvidia", "parameters", "throughput", "training", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-llama-nemotron-embed-1b-v2.md
source_anchor: ""
source_lines: [1, 46]
sha256: 91f15760d7f013f6dc1847e52c146f024e991271b3d74b17d18a4d8a698171c3
---

# llama-nemotron-embed-1b-v2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/llama-nemotron-embed-1b-v2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The Llama Nemotron Embedding 1B v2 is NVIDIA's open-source commercial text embedding model, optimized for multilingual and cross-lingual question-answering retrieval with support for long documents (up to 8192 tokens) and dynamic (Matryoshka) embedding sizes. It is a fine-tuned version of Llama 3.2 1B configured as a transformer encoder with 16 layers and an embedding size of 2048, trained with the AdamW optimizer (100 warm-up steps, 5e-6 learning rate, WarmupDecayLR scheduler) using a bi-encoder contrastive learning architecture. The model was evaluated on 26 languages including English, Arabic, Bengali, Chinese, Czech, Danish, Dutch, Finnish, French, German, Hebrew, Hindi, Hungarian, Indonesian, Italian, Japanese, Korean, Norwegian, Persian, Polish, Portuguese, Russian, Spanish, Swedish, Thai and Turkish. It outputs embedding vectors of up to 2048 dimensions, configurable to 384, 512, 768, 1024 or 2048, reducing data storage footprint by up to 35x through dynamic embedding sizing and longer token support. It is part of the NVIDIA NeMo Retriever collection and is also available as a NIM microservice. Training used a proprietary blend of publicly licensed QA datasets (12M semi-supervised pre-training samples, 1M fine-tuning samples) to avoid MS MARCO's commercial licensing restriction. Evaluation results (average Recall@5): NQ/HotpotQA/FiQA/TechQA 68.60% at dim 2048 and 64.48% at dim 384; multilingual MIRACL 60.75% at 2048 and 58.62% at 384; cross-lingual MLQA 79.86% at 2048 and 71.61% at 384; long-document MLDR 59.55% at 2048 and 54.77% at 384. Usage supports sentence-transformers (with `encode_query`/`encode_document` and `query:`/`passage:` prefixes), Transformers with average pooling, and vLLM (>=0.14.0) for high-throughput serving. It is governed by the NVIDIA Open Model License Agreement plus the Llama 3.2 Community Model License, ready for commercial use, with 417,811 downloads last month.

## Key points

- 1B-parameter multilingual/cross-lingual text embedding model for retrieval and RAG.
- Fine-tuned Llama 3.2 1B encoder: 16 layers, 2048 embedding size, 8192-token context.
- Matryoshka dynamic embeddings: 384, 512, 768, 1024 or 2048 dimensions; up to 35x storage reduction.
- Supports 26 languages; evaluated on NQ, HotpotQA, FiQA, TechQA, MIRACL, MLQA and MLDR.
- Average Recall@5: 68.60% (English QA, dim 2048), 79.86% (cross-lingual MLQA), 59.55% (MLDR).
- Usage via sentence-transformers, Transformers, or vLLM (>=0.14.0); also available as NVIDIA NIM.
- License: NVIDIA Open Model License + Llama 3.2 Community License; commercial use allowed.

## Technical data / figures

| Attribute | Value |
|---|---|
| Parameters | 1B |
| Architecture | Transformer encoder (fine-tuned Llama 3.2 1B Retriever) |
| Layers | 16 |
| Embedding size | 2048 (configurable: 384/512/768/1024/2048) |
| Max context | 8192 tokens |
| Languages | 26 (multilingual + cross-lingual) |
| Task | Feature extraction / text embeddings / retrieval |
| Precision | BF16 |
| Training | 12M pre-training + 1M fine-tuning samples |
| Recall@5 (NQ/HotpotQA/FiQA/TechQA) | 68.60% (dim 2048) |
| Recall@5 (MLQA cross-lingual) | 79.86% (dim 2048) |
| License | NVIDIA Open Model License + Llama 3.2 Community License |
| Downloads last month | 417,811 |

## Why this source matters for the RAG

This card documents the embedding model itself, making it a foundational reference for building retrieval-augmented generation pipelines with multilingual and long-document support. It is especially relevant to this RAG knowledge base because it describes the embedding and retrieval stack that powers such systems.
