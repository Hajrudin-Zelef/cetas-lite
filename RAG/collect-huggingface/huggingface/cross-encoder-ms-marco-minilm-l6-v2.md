---
id: collect-huggingface/huggingface/cross-encoder-ms-marco-minilm-l6-v2
title: "ms-marco-MiniLM-L6-v2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "benchmark", "latency", "license", "parameters", "throughput", "training"]
source: docs/RAG/Collect RAG/03_huggingface/cross-encoder-ms-marco-MiniLM-L6-v2.md
source_anchor: ""
source_lines: [1, 47]
sha256: 0b5dd9931b4a40368baa04e835ac8c34ccff276364d84eabbcc95058e88afd0b
---

# ms-marco-MiniLM-L6-v2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/cross-encoder/ms-marco-MiniLM-L6-v2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

ms-marco-MiniLM-L6-v2 is a Cross-Encoder trained on the MS MARCO Passage Ranking task, used for Information Retrieval re-ranking: given a query and a candidate passage, it outputs a relevance score, and passages are sorted in decreasing order to refine a first-stage retrieval (e.g., ElasticSearch or a bi-encoder). It is the classic "retrieve & re-rank" component of SBERT-based RAG systems, with ~88.7 million downloads per month and a hub model size of 22.7M parameters. The model is licensed under Apache 2.0.

Usage is provided for both sentence-transformers (`CrossEncoder('cross-encoder/ms-marco-MiniLM-L6-v2')` then `model.predict([(query, passage)])`) and raw HuggingFace Transformers (`AutoModelForSequenceClassification`). Training code is available in the SBERT repository (examples/cross_encoder/training/ms_marco). Performance results are reported on TREC Deep Learning 2019 and the MS MARCO Passage Reranking dev set: ms-marco-MiniLM-L6-v2 scores NDCG@10 74.30 on TREC DL 19 and MRR@10 39.01 on MS MARCO dev, processing ~1800 docs/sec on a V100. In the version-2 family it sits between the L4 variant (73.04 / 37.70) and the L12 variant (74.31 / 39.02), which is essentially the same quality at ~1/2 the throughput. The card also lists v1 and third-party models for comparison (e.g., TinyBERT-L2 9000 docs/s, electra-base 340 docs/s, nboost/bert-large 100 docs/s), showing the strong latency/quality trade-off of the MiniLM-L6 size.

The model is tagged text-ranking / text-classification, English, and built on a MiniLM-L12-H384-uncased base (the model tree indicates it derives from microsoft/MiniLM-L12-H384-uncased, with L12-v2 as a sibling). It has 78 fine-tunes and 27 quantizations downstream, and is used in 100+ Spaces, including many RAG and search assistants.

## Key points

- Cross-Encoder (MiniLM, 6 layers) for query-passage relevance scoring / re-ranking.
- Trained on MS MARCO Passage Ranking; Apache 2.0.
- NDCG@10 74.30 (TREC DL 19), MRR@10 39.01 (MS MARCO dev), ~1800 docs/s on V100.
- Classic "retrieve & re-rank" component for RAG pipelines.
- ~88.7M downloads/month; 78 fine-tunes, 27 quantizations downstream.
- Usable via sentence-transformers CrossEncoder or Transformers sequence classification.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | cross-encoder (Sentence Transformers) |
| Model name | ms-marco-MiniLM-L6-v2 |
| Architecture | Cross-Encoder, MiniLM 6 layers (sequence classification head) |
| Params | 22.7M |
| Training data | MS MARCO Passage Ranking |
| License | Apache 2.0 |
| TREC DL 2019 | NDCG@10 74.30 |
| MS MARCO dev | MRR@10 39.01 |
| Throughput | ~1800 docs/s (V100) |
| Base model | microsoft/MiniLM-L12-H384-uncased |
| Downloads/month | ~88,690,825 |

## Why this source matters for the RAG

This card is the canonical reference for the standard Cross-Encoder re-ranker used in RAG pipelines, providing authoritative performance numbers (NDCG@10, MRR@10, docs/sec) and a comparative table of the whole MS MARCO Cross-Encoder family. It supports accurate answers on re-ranking stages, benchmark comparisons, and model-selection guidance for retrieval systems.
