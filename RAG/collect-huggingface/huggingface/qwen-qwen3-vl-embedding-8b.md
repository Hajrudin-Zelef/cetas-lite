---
id: collect-huggingface/huggingface/qwen-qwen3-vl-embedding-8b
title: "Qwen3-VL-Embedding-8B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["embedding", "qwen", "apache", "benchmarks", "embeddings", "license", "multimodal", "parameters", "quantization", "reranker", "sglang", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3-VL-Embedding-8B.md
source_anchor: ""
source_lines: [1, 51]
sha256: 60f43b8785ec27e5347404f47da2192dcd91410148cc81b656a4371b75950891
---

# Qwen3-VL-Embedding-8B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3-VL-Embedding-8B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3-VL-Embedding-8B is a multimodal embedding model built on the open-sourced Qwen3-VL foundation model, part of the Qwen3-VL-Embedding and Qwen3-VL-Reranker series designed for multimodal information retrieval and cross-modal understanding. It accepts diverse inputs including text, images, screenshots, videos, and arbitrary combinations of these modalities. The companion Qwen3-VL-Reranker model refines initial retrieval results, forming a two-stage pipeline (efficient recall + precise re-ranking) for state-of-the-art multimodal search.

Architecture and specs: 8B parameters, 36 layers, 32K context length, embedding dimension up to 4096 with Matryoshka Representation Learning (MRL) support for user-defined output dimensions from 64 to 4096. It supports quantization of output embeddings and is instruction-aware (custom prompts improve results by 1-5%; English instructions recommended). It supports 30+ languages.

License is Apache-2.0. It is usable with sentence-transformers, Hugging Face Transformers (>=4.57.0 with qwen-vl-utils), vLLM (pooling runner), and SGLang (is_embedding mode). Base model is Qwen/Qwen3-VL-8B-Instruct.

Benchmarks: On MMEB-V2, Qwen3-VL-Embedding-8B achieves the top "All" score of 77.9, with Image Overall 80.1, Video Overall 66.1, and VisDoc Overall 83.3, outperforming GME-7B (59.1), VLM2Vec-V2 (59.2), Ops-MM-embedding-v1 (68.9), and IFM-TTE (74.1). On MMTEB, it scores 67.88 mean task, with Rerank 65.72 and Retri 69.41. Citation: arXiv 2601.04720. ~1.2M downloads per month.

## Key points

- Multimodal embedding model (text, images, screenshots, videos, mixed inputs).
- 8B params, 32K context, embedding dimension up to 4096 (MRL 64-4096).
- Companion Qwen3-VL-Reranker enables two-stage retrieval pipelines.
- Apache-2.0 license; runs with sentence-transformers, Transformers, vLLM, SGLang.
- No.1 on MMEB-V2 "All" score (77.9); MMTEB mean task 67.88.
- Supports 30+ languages; instruction-aware with quantized-embedding support.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 8B |
| Active parameters | 8B (dense) |
| Layers | 36 |
| Context length | 32K |
| Embedding dimension | up to 4096 (MRL 64-4096) |
| Languages | 30+ |
| Input modalities | Text, image, screenshot, video, mixed |
| License | Apache-2.0 |
| Precision | BF16 |
| MMEB-V2 All | 77.9 (No.1) |
| MMEB-V2 Image / Video / VisDoc Overall | 80.1 / 66.1 / 83.3 |
| MMTEB Mean Task / Mean Type | 67.88 / 58.88 |
| Monthly downloads | ~1.2M |

## Why this source matters for the RAG

Qwen3-VL-Embedding-8B extends RAG architecture knowledge to multimodal retrieval, where queries and documents mix text, images, screenshots, and video. Its top MMEB-V2 ranking and documented two-stage embedding + reranker pipeline provide authoritative guidance for building cross-modal search systems. This grounds the knowledge base on cutting-edge multimodal RAG techniques.
