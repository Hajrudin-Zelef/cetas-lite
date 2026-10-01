---
id: collect-261001-ia-llm/ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-1
title: "Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality"
domain: ia-llm
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["apache", "embedding", "embeddings", "attention", "benchmark", "benchmarks", "context window", "cost", "flash attention", "governance", "gpus", "inference"]
source: docs/RAG/collect-261001-ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-best-sub-.md
source_anchor: ""
source_lines: [1, 37]
sha256: ed4167cb08ea1b2e84e7b086494e77ca9e5105d8a81342b5c2c71eb6d9b6215c
---

# Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality

**TL;DR:** Two new Apache 2.0 multilingual embedding models built on ModernBERT — a 97M-parameter compact model that beats every open sub-100M multilingual embedder on MTEB Multilingual Retrieval (60.3), and a 311M full-size model that scores 65.2 on MTEB Multilingual Retrieval (#2 among open models under 500M parameters) with Matryoshka support. Both cover 200+ languages, are tuned on 52 languages, handle 32K-token context (64x R1), and add code retrieval across 9 programming languages.

**In this post:** Enterprise-Ready by Design · A Strong Sub-100M Multilingual Model · What Changed from R1 · Training the Full-Size 311M Model · Building the compact 97M Multilingual Model · Benchmark Results · Matryoshka Embeddings · Deployment Options · For Framework Integrators · Which Model Should You Use? · Try The Models

Multilingual embedding models face a persistent tension: broad language coverage usually comes at the cost of model size, and small models usually sacrifice languages. If you work across languages — retrieval-augmented generation over multilingual corpora, cross-lingual search, code retrieval in international teams — you've likely had to choose between a model that's fast enough and one that's good enough.

The Granite Embedding Multilingual R2 release narrows that gap considerably. We're releasing two new multilingual embedding models:

- **granite-embedding-311m-multilingual-r2** — A 311M-parameter full-size model with 768-dimensional embeddings, Matryoshka dimension support, and top-tier multilingual retrieval quality.
- **granite-embedding-97m-multilingual-r2** — A 97M-parameter compact model with 384-dimensional embeddings that delivers strong retrieval quality for its size.

Both models support **200+ languages** with enhanced retrieval quality for **52 languages and programming code**, handle context lengths up to **32,768 tokens** (a 64x increase over their R1 predecessors), and are released under the **Apache 2.0** license. They work out of the box with `sentence-transformers` and `transformers`, require no task-specific instructions, and are compatible as drop-in replacements in **LangChain**, **LlamaIndex**, **Haystack**, and **Milvus** with a one-line model name change. For frameworks currently using an English-only default, that one line gives every user in your community support for 200+ languages — no API changes, no new dependencies, no code changes required on their end. Both models ship with ONNX and OpenVINO weights for CPU-optimized inference.

## **52 enhanced-support languages** (click to expand)

The underlying encoder was pretrained on text from 200+ languages, producing general-purpose embeddings for any of them. The following 52 languages receive explicit retrieval-pair and cross-lingual training for higher-quality retrieval:

Albanian (sq), Arabic (ar), Azerbaijani (az), Bengali (bn), Bulgarian (bg), Catalan (ca), Chinese (zh), Croatian (hr), Czech (cs), Danish (da), Dutch (nl), English (en), Estonian (et), Finnish (fi), French (fr), Georgian (ka), German (de), Greek (el), Hebrew (he), Hindi (hi), Hungarian (hu), Icelandic (is), Indonesian (id), Italian (it), Japanese (ja), Kazakh (kk), Khmer (km), Korean (ko), Latvian (lv), Lithuanian (lt), Malay (ms), Marathi (mr), Norwegian (no), Persian (fa), Polish (pl), Portuguese (pt), Romanian (ro), Russian (ru), Serbian (sr), Slovak (sk), Slovenian (sl), Spanish (es), Swahili (sw), Swedish (sv), Tagalog (tl), Telugu (te), Thai (th), Turkish (tr), Ukrainian (uk), Urdu (ur), Uzbek (uz), Vietnamese (vi).

Additionally, the models are trained on **programming code** (Python, Go, Java, JavaScript, PHP, Ruby, SQL, C, C++) and support cross-lingual code retrieval.

Both embedding models are trained on a mixture of IBM‑curated datasets, publicly available data, and internally generated or synthetic data. Public web‑derived data used in training is selected and filtered using IBM‑developed quality, deduplication, and governance processes intended to reduce risk in downstream commercial use. We intentionally avoid the use of the MS‑MARCO training dataset and datasets with explicit non‑commercial licensing restrictions. The models are pretrained using GneissWeb, an IBM‑curated dataset derived from publicly available web content and processed using IBM’s data preparation and governance tooling—along with additional IBM‑curated and other publicly available sources. Datasets undergo IBM governance review to assess licensing considerations, ownership signals, and personal data risks. These processes are designed to contribute to responsible use and enterprise deployment.

The standout of this release is **granite-embedding-97m-multilingual-r2**. At 97 million parameters, it scores **60.3 on Multilingual MTEB Retrieval** across 18 languages — the highest retrieval score we've found for any open multilingual embedding model under 100M parameters. The next-best model in that size class, multilingual-e5-small, scores 50.9 on the same benchmark — a **+9.4 point gap** on a mature benchmark.

At roughly one-third the size of the 311M full-size model, it retains the majority of its retrieval quality across multilingual, code, and long-document benchmarks — a **+12.2 point gain on MTEB Multilingual Retrieval** over its direct predecessor, driven by a new architecture, better training data, and a novel pruning methodology (more on that below). The full-size **granite-embedding-311m-multilingual-r2** scores **65.2** on the same benchmark, a **+13.0 point gain** over its R1 predecessor.

The Granite Embedding Multilingual R1 models were built on XLM-RoBERTa encoders with a 512-token context window. The R2 generation is a ground-up rebuild:

ModernBERT is a recent encoder architecture that revisits the original BERT design with techniques from the last five years of transformer research. The shift brings several practical benefits: alternating attention lengths reduce computation on long sequences (improves throughput on long sequences significantly), rotary position embeddings allow the 32K context window without the positional interpolation hacks that plague older architectures, and Flash Attention 2.0 support speeds up encoding on modern GPUs.

The new multilingual tokenizers are worth highlighting. Rather than reusing XLM-RoBERTa's 250K-token vocabulary, we adopted existing tokenizers with strong multilingual and code coverage. The 311M model uses the Gemma 3 tokenizer (262K tokens); the 97M model starts from the GPT-OSS tokenizer and prunes it down to a compact 180K-token vocabulary that preserves broad multilingual coverage while reducing the embedding table's parameter footprint. Tokenizer efficiency matters more than people realize — a 32K-token window sounds impressive until your tokenizer burns half of it encoding a single paragraph of Thai.

The 311M model is a 22-layer ModernBERT encoder with a 262K-token multilingual vocabulary, trained through a multi-stage pipeline:

