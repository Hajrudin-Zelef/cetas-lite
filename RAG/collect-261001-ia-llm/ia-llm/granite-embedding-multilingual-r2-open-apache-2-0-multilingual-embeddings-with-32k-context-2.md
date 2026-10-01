---
id: collect-261001-ia-llm/ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-2
title: "Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality"
domain: ia-llm
role: reference
task: reference
actors: ["Mistral", "Nvidia", "OpenAI"]
dates: []
keywords: ["embedding", "embeddings", "benchmark", "benchmarks", "compute", "context window", "cost", "distillation", "fine-tuning", "gpu", "latency", "memory"]
source: docs/RAG/collect-261001-ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-best-sub-.md
source_anchor: ""
source_lines: [38, 91]
sha256: 7c5c2811e548cbb2c20be065ebf7c208e7dda0630c3cb01eaba585b78eb293e5
---

# Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality

1. **Knowledge distillation** : The model learns from multiple teacher models simultaneously. The teachers are Granite 3.3 Instruct and Mistral v0.2 Instruct decoder models, further finetuned for text embeddings, which transfer retrieval-specific knowledge into the 311M encoder architecture.
2. **Contrastive fine-tuning** : Standard contrastive training on multilingual retrieval pairs — queries matched with relevant and hard-negative passages across 52 languages and code — sharpens the model's ability to distinguish relevant from irrelevant results.
3. **Model merging** : After training, we merge checkpoints from different training stages and configurations. This combines the strengths of models optimized for different objectives (e.g., multilingual breadth vs. English depth) into a single set of weights without additional training compute.
4. **Matryoshka Representation Learning** : The model is trained with Matryoshka objectives so that its 768-dimensional embeddings can be truncated to 512, 384, 256, or 128 dimensions with minimal quality loss (see Matryoshka Embeddings below).

The result is a model that scores 65.2 on MTEB Multilingual Retrieval and 56.3 on the overall average — a +14.5 point average gain over its R1 predecessor.

The 97M model is trained through a combination of **vocabulary selection** and **knowledge distillation**:

1. **Vocabulary selection** : The 262K-token vocabulary is reduced to a purpose-trained 180K-token vocabulary that preserves broad multilingual coverage while cutting the embedding table size substantially.
2. **Knowledge distillation** : The pruned model is then finetuned using knowledge distillation from multiple teacher models (including a Granite 4.1 8B and Mistral Instruct decoder-based teacher) and contrastive training to improve retrieval quality.

This approach transfers retrieval-specific knowledge from multiple strong teachers, while reducing the model parameters without sacrificing language coverage. The result is a highly efficient compact model — scoring 60.3 on MTEB Multilingual Retrieval vs. 65.2 for the full-size model, while being approximately 3x smaller.

Performance across the main benchmark suite sorted by model size. Scores are averages across tasks within each benchmark (higher is better):

| Model | Params | Active Params | Embed Dim | MTEB Multilingual Retrieval (18) | Code (12) | English Retrieval (10) | LongEmbed (6) | RaR-b (17) | 
|---|---|---|---|---|---|---|---|---|
| F2LLM-v2-80M | 80M | 32M | 320 | 50.1 | 68.0 | 47.5 | 31.7 | 17.9 | 
| multilingual-e5-small | 118M | 22M | 384 | 50.9 | 53.5 | 46.5 | 38.8 | 20.3 | 
| granite-embedding-107m-multilingual (R1) | 107M | 11M | 384 | 48.1 | 40.7 | 47.9 | 34.3 | 17.1 | 
| paraphrase-multilingual-MiniLM-L12-v2 | 118M | 22M | 384 | 36.6 | 23.5 | 35.9 | 20.9 | 10.9 | 
| jina-embeddings-v5-text-nano | 212M | 113M | 768 | 63.3 | 71.2 | 58.8 | 63.6 | 25.2 | 
| harrier-oss-v1-270m | 268M | 100M | 640 | 66.4 | 62.4 | 52.1 | 64.9 | 32.9 | 
| multilingual-e5-base | 278M | 86M | 768 | 52.7 | 52.6 | 49.0 | 40.5 | 23.4 | 
| granite-embedding-278m-multilingual (R1) | 278M | 86M | 768 | 52.2 | 48.5 | 51.5 | 37.7 | 18.9 | 
| embeddinggemma-300m | 308M | 106M | 768 | 62.5 | 68.7 | 54.6 | 55.4 | 26.1 | 
| gte-multilingual-base | 305M | 113M | 768 | 57.2 | 57.5 | 50.8 | 62.1 | 19.0 | 
| snowflake-arctic-embed-m-v2.0 | 305M | 113M | 768 | 54.8 | 55.2 | 58.4 | 55.4 | 23.3 | 
| multilingual-e5-large | 560M | 304M | 1024 | 53.7 | 55.8 | 51.5 | 40.4 | 25.4 | 
| text-embedding-3-small (OpenAI, API only) | — | — | 1536 | 50.7 | — | 53.8 | 53.6 | 23.2 | 
| **granite-embedding-97m-multilingual-r2** | **97M** | **28M** | **384** | **60.3** | **60.4** | **50.1** | **65.6** | **24.9** | 
| **granite-embedding-311m-multilingual-r2** | **311M** | **110M** | **768** | **65.2 (#2)** | **63.8 (#3)** | **52.6 (#5)** | **71.7 (#1)** | **28.0 (#2)** | 

A few things stand out:

- **The 97M R2 model beats multilingual-e5-base and gte-multilingual-base** (~300M parameter models) on average and on most individual benchmarks, despite being roughly 3x smaller.
- **`paraphrase-multilingual-MiniLM-L12-v2` — a widely-used framework default — scores 36.6** , a full**+23.7 points** behind the 97M R2 model, which is also slightly smaller (97M vs 110M parameters) with the same 384-dimensional output.
- **LongEmbed is the biggest R1-to-R2 gain** : +31.3 points for the 97M model, +34.0 for the 311M. This is the direct payoff of the 32K context window — R1's 512-token limit meant your legal contract was being judged by its first page. Many practical multilingual workloads involve long documents (legal contracts, technical manuals, research papers, multi-page reports) that R1 simply could not see in full.
- **Code retrieval improves dramatically** : +19.7 (97M) and +15.3 (311M) over R1, reflecting the new code training set, larger context window, and better training methodology.
- **In the broader competitive field** , harrier-oss-v1-270m leads on MTEB Multilingual Retrieval (66.4) and RaR-b (32.9), while jina-embeddings-v5-text-nano leads on Code (71.2) and English Retrieval (58.8). The 311M Granite model is competitive on average (56.3) and leads on LongEmbed (71.7), while offering substantially higher encoding throughput than jina-embeddings-v5-text-nano (see speed table below).

Encoding speed matters for production workloads, especially when you're indexing millions of documents or need low-latency query encoding. We measured latency and throughput on a single NVIDIA H100 GPU using 512-token chunks:

The 97M model encodes over 2,500 documents per second — comparable throughput to multilingual-e5-small — while delivering substantially higher retrieval quality. The 311M model, at ~1,800 docs/sec, performs better than jina-embeddings-v5-text-nano on retrieval quality (65.2 vs. 63.3) at over 5.5x the encoding speed (note: speed numbers are computed with the latest transformer code, which had a speed regression vs the last 4.57 version - for both the Jina and granite models - see our technical report for details). harrier-oss-v1-270m offers the best combination of speed and retrieval score among the competitors listed here.

The 311M model supports Matryoshka Representation Learning, which lets you truncate embeddings from the full 768 dimensions down to 512, 384, 256, or 128 with graceful quality degradation. This is useful when storage, memory, or similarity-computation cost is a concern — a 256-dimensional embedding takes one-third the storage of a 768-dimensional one, and cosine similarity is proportionally cheaper to compute.

Here's how retrieval quality holds up across embedding dimensions:

The quality loss from dimension reduction is remarkably small. Cutting from 768 to 256 dimensions — a **3x reduction** in storage and similarity-computation cost — drops MTEB Multilingual Retrieval by just 0.5 points (65.2 → 64.7) and Code Retrieval by 0.5 points (63.9 → 63.4). Even at 128 dimensions (a **6x reduction**), the model still scores 63.7 on MTEB Multilingual Retrieval and 62.3 on Code — retaining over **97%** of its full-dimension performance. In practice, this means you can substantially reduce your index size and search latency with minimal impact on result quality. (Note,results in the above picture were evaluated with a context length of 1024 for English and Multilingual Retrieval, and 8192 for Code).

For comparison, the 311M model truncated to 384 dimensions (the same dimensionality as the 97M model's native output) still outperforms the 97M model across all three benchmarks. If you need 384-dimensional embeddings and can afford the 311M model's encoding cost, Matryoshka truncation is the stronger option.

