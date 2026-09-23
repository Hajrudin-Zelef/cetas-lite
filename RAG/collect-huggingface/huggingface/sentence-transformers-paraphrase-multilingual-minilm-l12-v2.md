---
id: collect-huggingface/huggingface/sentence-transformers-paraphrase-multilingual-minilm-l12-v2
title: "paraphrase-multilingual-MiniLM-L12-v2 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "embedding", "embeddings", "inference", "license", "research"]
source: docs/RAG/Collect RAG/03_huggingface/sentence-transformers-paraphrase-multilingual-MiniLM-L12-v2.md
source_anchor: ""
source_lines: [1, 47]
sha256: 11bc51c32a0a05866a22b6ddf4981e18865ba7a63db05b5ebefd75cfe22bcdfa
---

# paraphrase-multilingual-MiniLM-L12-v2 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

paraphrase-multilingual-MiniLM-L12-v2 is a sentence-transformers (SBERT) model that maps sentences and short paragraphs to a 384-dimensional dense vector space for semantic search, clustering and similarity tasks, and it supports 50 languages, making it a common choice for multilingual RAG pipelines. The hub reports a model size of 0.1B params and ~46 million downloads per month. The full architecture is a SentenceTransformer wrapping a BERT model (BertModel): a Transformer module with `max_seq_length` 128 and `do_lower_case: False`, followed by a Pooling module configured with `pooling_mode_mean_tokens: True` (mean pooling) over a 384-dim word embedding space. It is licensed under Apache 2.0 and is part of the sentence-transformers paraphrase model family.

Usage is provided for sentence-transformers (`SentenceTransformer('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')`, then `model.encode(sentences)`) and for raw HuggingFace Transformers, where the user must apply mean pooling on the contextualized token embeddings and normalize. The model was trained by sentence-transformers (Nils Reimers & Iryna Gurevych) and the card asks to cite the Sentence-BERT paper (arXiv 1908.10084) when the model is used in research. The model's multilingual nature (50 languages) makes it directly usable for cross-lingual semantic search and paraphrase detection across European and Asian languages without retraining.

On the hub it has 335 fine-tunes and 29 quantizations downstream, and appears in 100+ Spaces including multilingual leaderboards (EuroEval) and translation/similarity demos. It is associated with the "Text Embeddings Inference" ecosystem, meaning it can be served via TEI. The max sequence length of 128 tokens bounds document/query sizes to short paragraphs, which is an important operational constraint to note for retrieval usage.

## Key points

- Multilingual sentence embedding model (50 languages), 384-dim vectors.
- BertModel backbone with max_seq_length 128, mean pooling; ~0.1B params.
- Apache 2.0; ~46M downloads/month.
- Used for semantic search, clustering, sentence similarity, paraphrase detection.
- Cite Sentence-BERT (arXiv 1908.10084) in research.
- 335 fine-tunes, 29 quantizations, 100+ Spaces; TEI-served.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | sentence-transformers |
| Model name | paraphrase-multilingual-MiniLM-L12-v2 |
| Architecture | SentenceTransformer: BertModel + mean pooling |
| Params | 0.1B |
| Embedding dim | 384 |
| Max seq length | 128 |
| Languages | 50 |
| License | Apache 2.0 |
| Tasks | Sentence Similarity / feature extraction |
| Paper | Sentence-BERT (arXiv 1908.10084) |
| Downloads/month | ~46,018,622 |

## Why this source matters for the RAG

This card documents a standard lightweight multilingual embedding model for cross-lingual and multilingual retrieval, including its exact pooling architecture and 50-language coverage. It provides citable data on multilingual semantic-search baselines and a practical reference for embedding-model selection in non-English RAG deployments.
