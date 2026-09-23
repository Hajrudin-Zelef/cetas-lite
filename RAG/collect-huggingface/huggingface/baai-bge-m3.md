---
id: collect-huggingface/huggingface/baai-bge-m3
title: "BGE-M3 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["benchmark", "benchmarks", "compute", "distillation", "embedding", "embeddings", "fine-tuning", "license", "mit license", "reranker", "training"]
source: docs/RAG/Collect RAG/03_huggingface/BAAI-bge-m3.md
source_anchor: ""
source_lines: [1, 49]
sha256: df67207db28366249201ee3dc5e42f2772256c47e1434afc8afe791af5693e91
---

# BGE-M3 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/BAAI/bge-m3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

BGE-M3 is BAAI's (Beijing Academy of Artificial Intelligence) embedding model distinguished by Multi-Functionality, Multi-Linguality and Multi-Granularity. It simultaneously performs three retrieval functionalities in one model: dense retrieval (single embedding), multi-vector (ColBERT-style) retrieval, and sparse/lexical retrieval (SPLADE-like token weights, analogous to BM25). It supports more than 100 working languages and processes inputs from short sentences up to long documents of 8,192 tokens, making it one of the most versatile open embedding models for RAG. It is licensed under MIT and has ~37.2 million downloads per month. The model has a 1024-dimensional embedding and is built on an xlm-roberta-large backbone whose max length was extended to 8192 and further pre-trained via RetroMAE (bge-m3-retromae), then contrastively learned (bge-m3-unsupervised), then unified fine-tuned on dense/sparse/ColBERT objectives.

The card gives detailed usage: dense embeddings via `BGEM3FlagModel` (`use_fp16=True` for speed), sparse embeddings returning per-token lexical weights that can be scored with `compute_lexical_matching_score`, and multi-vector ColBERT vectors scored with `colbert_score`, plus a unified `compute_score` for text pairs with `weights_for_different_modes` (e.g., 0.4/0.2/0.4). For RAG, BAAI recommends a hybrid retrieval + re-ranking pipeline (e.g., BGE-M3 embedding + sparse retrieval with Vespa or Milvus, then a bge-reranker cross-encoder to filter results). Fine-tuning is supported for the dense mode alone or all three modes via unified fine-tuning examples. Released training/eval resources include the MLDR long-document retrieval dataset (13 languages) and the bge-m3-data fine-tuning set.

Training innovations include self-knowledge distillation (using combined outputs of different retrieval modes as reward signal to boost sparse and multi-vector modes), efficient batching for long text, and MCLS (a method to improve long-text performance without fine-tuning). Benchmarks cover MIRACL (multilingual), MKQA (cross-lingual), and long-document retrieval on MLDR and NarrativeQA, plus comparisons with BM25. An independent benchmark reported BGE-M3 as the top performer in both English and other languages, surpassing OpenAI's embedding model. MTEB results are also listed (e.g., ArguAna nDCG@10 54.04, MTEB v1.12.75; BRIGHT AoPS Retrieval 4.56). The model has 567 fine-tunes and 293 quantizations downstream.

## Key points

- One model with dense + sparse + multi-vector (ColBERT) retrieval, unified fine-tuning.
- 100+ languages; handles up to 8,192-token documents; 1024-dim embeddings.
- Built on xlm-roberta-large extended to 8K and RetroMAE-pretrained.
- Recommended RAG pipeline: hybrid retrieval + cross-encoder re-ranking.
- MIT license; ~37.2M downloads/month.
- Training: self-knowledge distillation, efficient batching, MCLS.
- Top performer on independent multilingual benchmarks, incl. vs. OpenAI.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | BAAI |
| Model name | bge-m3 |
| Architecture | XLM-RoBERTa-large (RetroMAE, extended to 8192) |
| Embedding dim | 1024 |
| Max sequence length | 8192 |
| Languages | 100+ |
| Functions | Dense, sparse (lexical), multi-vector (ColBERT) |
| License | MIT |
| Paper | BGE M3-Embedding (arXiv 2402.03216) |
| Key datasets | MLDR (13 langs), bge-m3-data, MIRACL, MKQA, NarrativeQA |
| MTEB examples | ArguAna nDCG@10 54.04; BRIGHT AoPS Retrieval 4.56 |
| Downloads/month | ~37,229,647 |

## Why this source matters for the RAG

This card is a cornerstone source for multilingual, multi-function retrieval, documenting a single MIT-licensed model that covers dense, sparse and ColBERT retrieval plus 8K-document granularity with explicit RAG pipeline guidance (hybrid retrieval + re-ranking). It provides citable architecture, training, benchmark, and usage data that directly inform modern RAG design.
