---
id: collect-240926-huggingface/huggingface/qwen-qwen3-embedding-0-6b-hugging-face-1
title: "Requires transformers>=4.51.0"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba"]
dates: ["2025-06-05"]
keywords: ["benchmark", "embedding", "embeddings", "inference", "leaderboard", "parameters", "qwen", "reasoning", "reranker", "training"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-embedding-0-6b-hugging-face.md
source_anchor: ""
source_lines: [1, 37]
sha256: 2e81836a9389ff61fc11efa70148ee62133d38a8a81efe43e9bf11c6bc435534
---

# Requires transformers>=4.51.0

<!-- source: https://huggingface.co/Qwen/Qwen3-Embedding-0.6B -->

The Qwen3 Embedding model series is the latest proprietary model of the Qwen family, specifically designed for text embedding and ranking tasks. Building upon the dense foundational models of the Qwen3 series, it provides a comprehensive range of text embeddings and reranking models in various sizes (0.6B, 4B, and 8B). This series inherits the exceptional multilingual capabilities, long-text understanding, and reasoning skills of its foundational model. The Qwen3 Embedding series represents significant advancements in multiple text embedding and ranking tasks, including text retrieval, code retrieval, text classification, text clustering, and bitext mining.

**Exceptional Versatility**: The embedding model has achieved state-of-the-art performance across a wide range of downstream application evaluations. The 8B size embedding model ranks **No.1** in the MTEB multilingual leaderboard (as of June 5, 2025, score **70.58**), while the reranking model excels in various text retrieval scenarios.

**Comprehensive Flexibility**: The Qwen3 Embedding series offers a full spectrum of sizes (from 0.6B to 8B) for both embedding and reranking models, catering to diverse use cases that prioritize efficiency and effectiveness. Developers can seamlessly combine these two modules. Additionally, the embedding model allows for flexible vector definitions across all dimensions, and both embedding and reranking models support user-defined instructions to enhance performance for specific tasks, languages, or scenarios.

**Multilingual Capability**: The Qwen3 Embedding series offer support for over 100 languages, thanks to the multilingual capabilites of Qwen3 models. This includes various programming languages, and provides robust multilingual, cross-lingual, and code retrieval capabilities.

**Qwen3-Embedding-0.6B** has the following features:

- Model Type: Text Embedding
- Supported Languages: 100+ Languages
- Number of Parameters: 0.6B
- Context Length: 32k
- Embedding Dimension: Up to 1024, supports user-defined output dimensions ranging from 32 to 1024

For more details, including benchmark evaluation, hardware requirements, and inference performance, please refer to our blog, GitHub.

| Model Type | Models | Size | Layers | Sequence Length | Embedding Dimension | MRL Support | Instruction Aware | 
|---|---|---|---|---|---|---|---|
| Text Embedding | Qwen3-Embedding-0.6B | 0.6B | 28 | 32K | 1024 | Yes | Yes | 
| Text Embedding | Qwen3-Embedding-4B | 4B | 36 | 32K | 2560 | Yes | Yes | 
| Text Embedding | Qwen3-Embedding-8B | 8B | 36 | 32K | 4096 | Yes | Yes | 
| Text Reranking | Qwen3-Reranker-0.6B | 0.6B | 28 | 32K | - | - | Yes | 
| Text Reranking | Qwen3-Reranker-4B | 4B | 36 | 32K | - | - | Yes | 
| Text Reranking | Qwen3-Reranker-8B | 8B | 36 | 32K | - | - | Yes | 

**Note**:


`MRL Support` indicates whether the embedding model supports custom dimensions for the final embedding. 
`Instruction Aware` notes whether the embedding or reranking model supports customizing the input instruction according to different tasks.- Our evaluation indicates that, for most downstream tasks, using instructions (instruct) typically yields an improvement of 1% to 5% compared to not using them. Therefore, we recommend that developers create tailored instructions specific to their tasks and scenarios. In multilingual contexts, we also advise users to write their instructions in English, as most instructions utilized during the model training process were originally written in English.

With Transformers versions earlier than 4.51.0, you may encounter the following error:

