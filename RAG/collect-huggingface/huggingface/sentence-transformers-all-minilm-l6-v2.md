---
id: collect-huggingface/huggingface/sentence-transformers-all-minilm-l6-v2
title: "all-MiniLM-L6-v2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "embedding", "fine-tuning", "license", "parameters", "quantization", "tpu", "training"]
source: docs/RAG/Collect RAG/03_huggingface/sentence-transformers-all-MiniLM-L6-v2.md
source_anchor: ""
source_lines: [1, 49]
sha256: 0babd65405a261dfad12afc616515656a1a443f73e936d2779f80cece3493a7f
---

# all-MiniLM-L6-v2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/sentence-transformers/all-MiniLM-L6-v2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

all-MiniLM-L6-v2 is a sentence-transformers (SBERT) model that maps sentences and short paragraphs to a 384-dimensional dense vector space for tasks such as clustering, semantic search and sentence similarity. It is one of the most widely used embedding models on the Hugging Face hub, with ~251 million downloads per month and a hub model size of 22.7M parameters. It was built by fine-tuning the pretrained `nreimers/MiniLM-L6-H384-uncased` checkpoint (a compact BERT-style MiniLM with 6 layers and 384 hidden dimensions) on a very large dataset of over 1 billion sentence pairs using a self-supervised contrastive learning objective: given a sentence from a pair, the model must predict which of randomly sampled sentences was actually paired with it. It was developed during the Hugging Face "Community week using JAX/Flax for NLP & CV" using 7 TPUs v3-8.

Fine-tuning details: trained on a TPU v3-8 for 100k steps with batch size 1024 (128 per TPU core), 500 warmup steps, sequence length limited to 128 tokens, AdamW optimizer with 2e-5 learning rate. Training data is a weighted concatenation of 27 datasets totaling 1,170,060,424 sentence pairs, including Reddit comments (726M), S2ORC citation pairs (116M+52M+41M), WikiAnswers (77M), PAQ (64M), Stack Exchange (multiple subsets), MS MARCO triplets (9M), GOOAQ (3M), Yahoo Answers, Code Search, COCO, SPECTER, SearchQA, ELI5, Flickr 30k, AllNLI (SNLI + MultiNLI), Quora triplets, Simple Wikipedia, Natural Questions, SQuAD2.0 and TriviaQA.

Usage is provided both via the sentence-transformers library (`SentenceTransformer('sentence-transformers/all-MiniLM-L6-v2')` followed by `model.encode(sentences)`) and via raw HuggingFace Transformers with a mean-pooling layer and L2 normalization. Inputs longer than 256 word pieces are truncated by default. The license is Apache 2.0. An MTEB evaluation result on ArguAna (nDCG@10 50.17, MTEB v1.12.75) is listed. It is the base of 1,063 fine-tunes, 18 adapters, 105 quantizations and 100+ Spaces on the hub.

## Key points

- 22.7M-param MiniLM sentence embedding model producing 384-dim vectors.
- Fine-tuned on >1B sentence pairs with contrastive learning; base is MiniLM-L6-H384-uncased.
- Built during Hugging Face Community Week on 7 TPUs v3-8 (JAX/Flax).
- Intended for semantic search, clustering, sentence similarity; truncates at 256 tokens.
- Apache 2.0; ~251M downloads/month; 1,063 fine-tunes and 105 quantizations downstream.
- Usage via sentence-transformers or raw Transformers + mean pooling + normalize.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | sentence-transformers |
| Model name | all-MiniLM-L6-v2 |
| Architecture | MiniLM (BERT-style), 6 layers, hidden 384 |
| Params | 22.7M |
| Embedding dim | 384 |
| Max seq length | 256 word pieces (default truncation) |
| Fine-tuning data | 1,170,060,424 sentence pairs (27 datasets) |
| Training | TPU v3-8, 100k steps, batch 1024, AdamW lr 2e-5, seq 128 |
| License | Apache 2.0 |
| Tasks | Sentence Similarity / feature extraction / retrieval |
| Languages | English |
| MTEB | ArguAna nDCG@10 50.17 (MTEB v1.12.75) |
| Downloads/month | ~251,354,399 |

## Why this source matters for the RAG

This card documents the canonical lightweight English embedding model used for dense retrieval in countless RAG pipelines, with complete training-data and hyperparameter detail. It matters as a reference baseline for embedding quality, quantization, and retrieval setup, and for understanding the sentence-transformers API used throughout the ecosystem.
