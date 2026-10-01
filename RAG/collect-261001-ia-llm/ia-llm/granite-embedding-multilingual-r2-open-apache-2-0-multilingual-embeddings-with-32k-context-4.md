---
id: collect-261001-ia-llm/ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-4
title: "Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Hugging Face"]
dates: []
keywords: ["apache", "embedding", "embeddings", "benchmarks", "gpu", "inference", "latency", "license", "parameters", "pruning", "safetensors", "throughput"]
source: docs/RAG/collect-261001-ia-llm/granite-embedding-multilingual-r2-open-apache-2-0-multilingual-embeddings-with-32k-context-best-sub-.md
source_anchor: ""
source_lines: [238, 266]
sha256: 3d8887f66d2ef983a3feb0f776bf784ce4a4f0cdc087f4f2849ca2cabf319ce3
---

# Granite Embedding Multilingual R2: Open Apache 2.0 Multilingual Embeddings with 32K Context — Best Sub-100M Retrieval Quality

- **License** : Apache 2.0, trained without MS-MARCO
- **Drop-in behavior** : No task-specific instruction prefix required — behaves like`all-MiniLM-L6-v2` at the API level. Existing code that calls`.encode()` works unchanged.
- **Dimensionality** : 384-dimensional output (97M) and 768-dimensional output (311M), matching the most common existing defaults. No index migration required.
- **Model size** : The 97M model's weights are 195 MB (safetensors) — less than half the size of`paraphrase-multilingual-MiniLM-L12-v2` (471 MB), the most common multilingual default. The quantized ONNX weights are just 98 MB, comparable to`all-MiniLM-L6-v2` (91 MB) while covering 200+ languages.
- **CPU-friendly** : Ships with ONNX and OpenVINO weights for optimized CPU inference. No GPU dependency for a getting-started tutorial.
- **Multilingual by default** : If your current default is English-only, this is a one-line swap that gives every user in your community support for 200+ languages — without touching their code.
- **Stable identifier** :`ibm-granite/granite-embedding-97m-multilingual-r2` on Hugging Face, maintained by IBM under the Granite model family.

To discuss adopting these models as a default in your project, open an issue at ibm-granite/granite-embedding-models.

These two multilingual models are part of the broader **Granite Embedding R2** family, which also includes two high-performing English-focused models: granite-embedding-english-r2 (149M parameters) and granite-embedding-small-english-r2 (47M parameters). If your data is predominantly English, the English models offer higher retrieval quality on English benchmarks at a smaller footprint, since they don't need to allocate capacity across 200+ languages.

| If you need... | Use | 
|---|---|
| Best multilingual retrieval quality | granite-embedding-311m-multilingual-r2 | 
| Flexible embedding dimensions (storage/speed tradeoff) | granite-embedding-311m-multilingual-r2 (Matryoshka) | 
| Maximum throughput / edge deployment / low latency | granite-embedding-97m-multilingual-r2 | 
| Best cross-lingual transfer across many language pairs | granite-embedding-311m-multilingual-r2 | 
| Predominantly English data | granite-embedding-english-r2 or granite-embedding-small-english-r2 | 

Both models are available now on Hugging Face under the IBM Granite Embedding collection:

You can try the small models interactively (on CPU) via a Granite Embedding demo here on Hugging Face Spaces, or run the full examples notebook in Google Colab:

You can access our detailed technical report covering the full training methodology, per-language evaluations, and pruning ablations here Granite Multilingual Embedding R2 report. For questions, feedback, or issues, visit ibm-granite/granite-embedding-models on GitHub.

**Framework maintainers:** If you'd like to adopt these models as a default in your project, open an issue at ibm-granite/granite-embedding-models — we're happy to help with integration, testing, and any questions about licensing or deployment.

Give them a try, and if the embeddings spark joy, smash that ❤️ button on Hugging Face. Our models have feelings too, and every +1 keeps them warm at night.
