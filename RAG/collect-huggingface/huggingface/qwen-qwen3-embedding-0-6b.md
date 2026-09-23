---
id: collect-huggingface/huggingface/qwen-qwen3-embedding-0-6b
title: "Qwen3-Embedding-0.6B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "vLLM"]
dates: ["2025-06-05", "2026-09-23"]
keywords: ["embedding", "qwen", "apache", "attention", "benchmarks", "embeddings", "inference", "leaderboard", "license", "memory", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3-Embedding-0.6B.md
source_anchor: ""
source_lines: [1, 50]
sha256: 995b25dd0ec50b4a274f33fbe2144d889252c8a91776731b2db247040f8f1381
---

# Qwen3-Embedding-0.6B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3-Embedding-0.6B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3-Embedding-0.6B is the smallest model in the Qwen3 Embedding series, a family specifically designed for text embedding and ranking tasks. Built upon the dense foundational models of the Qwen3 series, it inherits strong multilingual capabilities, long-text understanding, and reasoning skills. The model targets text retrieval, code retrieval, text classification, text clustering, and bitext mining. It supports over 100 languages, including programming languages, giving robust multilingual, cross-lingual, and code retrieval capabilities.

The 0.6B model is a text embedding model with 28 layers, a context length of 32K tokens, and an embedding dimension of up to 1024 that supports user-defined output dimensions ranging from 32 to 1024 (Matryoshka Representation Learning, MRL). It is instruction-aware: developers can prepend task-specific instructions to queries to improve performance by roughly 1-5%. The larger 8B sibling ranks No.1 on the MTEB multilingual leaderboard (score 70.58 as of June 5, 2025).

The model is released under the Apache-2.0 license. It ships in BF16 Safetensors format and can be run via sentence-transformers (>=2.7.0), Transformers (>=4.51.0), vLLM (>=0.8.5), or Text Embeddings Inference (TEI). It supports flash_attention_2 for acceleration and memory savings, and left padding is recommended for batched inference. The last-token pooling strategy is used to derive embeddings.

Benchmarks: on MTEB Multilingual, Qwen3-Embedding-0.6B scores 64.33 (mean task) and 56.00 (mean type), beating BGE-M3 (59.56) and multilingual-e5-large-instruct (63.22) despite similar size. On MTEB Eng v2 it scores 70.70 mean task; on C-MTEB it scores 66.33. The model has over 8.9M downloads per month and is widely used in RAG pipelines, semantic search, and clustering applications.

## Key points

- Small (0.6B) multilingual text embedding model, part of the Qwen3 Embedding series (0.6B, 4B, 8B).
- 32K context length; embedding dimension up to 1024 with MRL support (32-1024).
- Supports 100+ languages and code retrieval; instruction-aware.
- Apache-2.0 license, BF16 Safetensors, ideal for RAG and semantic search.
- MTEB Multilingual mean task 64.33, Eng v2 70.70, C-MTEB 66.33.
- Usable with sentence-transformers, Transformers, vLLM, and TEI.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 0.6B |
| Active parameters | 0.6B (dense) |
| Layers | 28 |
| Context length | 32K tokens |
| Embedding dimension | up to 1024 (MRL, 32-1024) |
| Languages | 100+ |
| License | Apache-2.0 |
| Quantization | BF16 (Safetensors); community quants available |
| MTEB Multilingual (Mean Task) | 64.33 |
| MTEB Eng v2 (Mean Task) | 70.70 |
| C-MTEB (Mean Task) | 66.33 |
| Monthly downloads | ~8.9M |

## Why this source matters for the RAG

Qwen3-Embedding-0.6B is a reference embedding model for building the retrieval layer of a RAG system, offering a strong size/performance trade-off. Its documented benchmarks and usage snippets provide concrete grounding for semantic search and multilingual retrieval claims. It helps the knowledge base answer questions about embedding model selection and implementation.
