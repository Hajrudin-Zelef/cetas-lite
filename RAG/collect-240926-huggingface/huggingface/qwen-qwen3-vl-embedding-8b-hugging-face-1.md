---
id: collect-240926-huggingface/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face-1
title: "Load the model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["benchmark", "distribution", "embedding", "embeddings", "inference", "multimodal", "parameters", "quantization", "qwen", "reranker", "training"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face.md
source_anchor: ""
source_lines: [1, 59]
sha256: 8d311b05263b4346f4f60e9b31c95c5a173c5c857b55d215f7eaa788e97f1642
---

# Load the model

<!-- source: https://huggingface.co/Qwen/Qwen3-VL-Embedding-8B -->

The **Qwen3-VL-Embedding** and **Qwen3-VL-Reranker** model series are the latest additions to the Qwen family, built upon the recently open-sourced and powerful Qwen3-VL foundation model. Specifically designed for multimodal information retrieval and cross-modal understanding, this suite accepts diverse inputs including text, images, screenshots, and videos, as well as inputs containing a mixture of these modalities.

While the Embedding model generates high-dimensional vectors for broad applications like retrieval and clustering, the Reranker model is engineered to refine these results, establishing a comprehensive pipeline for state-of-the-art multimodal search.

- **Multimodal Versatility** : Both models seamlessly handle a wide range of inputs—including text, images, screenshots, and video—within a unified framework. They deliver state-of-the-art performance across diverse multimodal tasks such as image-text retrieval, video-text matching, visual question answering (VQA), and multimodal content clustering.
- **Unified Representation Learning (Embedding)** : By leveraging the Qwen3-VL architecture, the Embedding model generates semantically rich vectors that capture both visual and textual information in a shared space. This facilitates efficient similarity computation and retrieval across different modalities.
- **High-Precision Reranking (Reranker)** : We also introduce the Qwen3-VL-Reranker series to complement the embedding model. The reranker takes a (query, document) pair as input—where both query and document may contain arbitrary single or mixed modalities—and outputs a precise relevance score. In retrieval pipelines, the two models are typically used in tandem: the embedding model performs efficient initial recall, while the reranker refines results in a subsequent re-ranking stage. This two-stage approach significantly boosts retrieval accuracy.
- **Exceptional Practicality** : Inheriting Qwen3-VL’s multilingual capabilities, the series supports over 30 languages, making it ideal for global applications. It is highly practical for real-world scenarios, offering flexible vector dimensions, customizable instructions for specific use cases, and strong performance even with quantized embeddings. These capabilities enable developers to seamlessly integrate both models into existing pipelines, unlocking powerful cross-lingual and cross-modal understanding.

**Qwen3-VL-Embedding-8B** has the following features:

- Model Type: MultiModal Embedding
- Supported Languages: 30+ Languages
- Supported Input Modalities: Text, images, screenshots, videos, and arbitrary multimodal combinations (e.g., text + image, text + video)
- Number of Parameters: 8B
- Context Length: 32k
- Embedding Dimension: Up to 4096, supports user-defined output dimensions ranging from 64 to 4096

For more details, including benchmark evaluation, hardware requirements, and inference performance, please refer to our technical report, blog, GitHub.

| Model | Size | Model Layers | Sequence Length | Embedding Dimension | Quantization Support | MRL Support | Instruction Aware | 
|---|---|---|---|---|---|---|---|
| Qwen3-VL-Embedding-2B | 2B | 28 | 32K | 2048 | Yes | Yes | Yes | 
| Qwen3-VL-Embedding-8B | 8B | 36 | 32K | 4096 | Yes | Yes | Yes | 
| Qwen3-VL-Reranker-2B | 2B | 28 | 32K | - | - | - | Yes | 
| Qwen3-VL-Reranker-8B | 8B | 36 | 32K | - | - | - | Yes | 

**Note**:


`Quantization Support` indicates the supported quantization post process for the output embedding. 
`MRL Support` indicates whether the embedding model supports custom dimensions for the final embedding. 
`Instruction Aware` notes whether the embedding or reranking model supports customizing the input instruction according to different tasks.
Our evaluation indicates that, for most downstream tasks, using instructions (instruct) typically yields an improvement of 1% to 5% compared to not using them. Therefore, we recommend that developers create tailored instructions specific to their tasks and scenarios. In multilingual contexts, we also advise users to write their instructions in English, as most instructions utilized during the model training process were originally written in English.

### Evaluation Results on MMEB-V2

Results on the MMEB-V2 benchmark. All models except IFM-TTE have been re-evaluated on the updated VisDoc OOD split. CLS: classification, QA: question answering, RET: retrieval, GD: grounding, MRET: moment retrieval, VDR: ViDoRe, VR: VisRAG, OOD: out-of-distribution.

| Model | Model Size | Image CLS | Image QA | Image RET | Image GD | Image Overall | Video CLS | Video QA | Video RET | Video MRET | Video Overall | VisDoc VDRv1 | VisDoc VDRv2 | VisDoc VR | VisDoc OOD | VisDoc Overall | All | 
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| **# of Datasets →** |  | 10 | 10 | 12 | 4 | 36 | 5 | 5 | 5 | 3 | 18 | 10 | 4 | 6 | 4 | 24 | 78 | 
| VLM2Vec | 2B | 58.7 | 49.3 | 65.0 | 72.9 | 59.7 | 33.4 | 30.5 | 20.6 | 30.7 | 28.6 | 49.8 | 13.5 | 51.8 | 48.2 | 44.0 | 47.7 | 
| VLM2Vec-V2 | 2B | 62.9 | 56.3 | 69.5 | 77.3 | 64.9 | 39.3 | 34.3 | 28.8 | 36.8 | 34.6 | 75.5 | 44.9 | 79.4 | 62.2 | 69.2 | 59.2 | 
| GME-2B | 2B | 54.4 | 29.9 | 66.9 | 55.5 | 51.9 | 34.9 | 42.0 | 25.6 | 31.1 | 33.6 | 86.1 | 54.0 | 82.5 | 67.5 | 76.8 | 55.3 | 
| GME-7B | 7B | 57.7 | 34.7 | 71.2 | 59.3 | 56.0 | 37.4 | 50.4 | 28.4 | 37.0 | 38.4 | 89.4 | 55.6 | 85.0 | 68.3 | 79.3 | 59.1 | 
| Ops-MM-embedding-v1 | 8B | 69.7 | 69.6 | 73.1 | 87.2 | 72.7 | 59.7 | 62.2 | 45.7 | 43.2 | 53.8 | 80.1 | 59.6 | 79.3 | 67.8 | 74.4 | 68.9 | 
| IFM-TTE | 8B | 76.7 | 78.5 | 74.6 | 89.3 | 77.9 | 60.5 | 67.9 | 51.7 | 54.9 | 59.2 | 85.2 | 71.5 | 92.7 | 53.3 | 79.5 | 74.1 | 
| RzenEmbed | 8B | 70.6 | 71.7 | 78.5 | 92.1 | 75.9 | 58.8 | 63.5 | 51.0 | 45.5 | 55.7 | 89.7 | 60.7 | 88.7 | 69.9 | 81.3 | 72.9 | 
| Seed-1.6-embedding-1215 | unknown | 75.0 | 74.9 | 79.3 | 89.0 | 78.0 | 85.2 | 66.7 | 59.1 | 54.8 | 67.7 | 90.0 | 60.3 | 90.0 | 70.7 | 82.2 | 76.9 | 
| **Qwen3-VL-Embedding-2B** | 2B | 70.2 | 74.4 | 74.9 | 88.6 | 75.0 | 72.8 | 63.8 | 52.3 | 51.6 | 61.1 | 85.2 | 66.0 | 86.3 | 74.3 | 80.2 | 73.4 | 
| **Qwen3-VL-Embedding-8B** | 8B | 74.4 | 81.0 | 80.0 | 92.2 | 80.1 | 79.1 | 70.1 | 57.0 | 53.2 | 66.1 | 88.2 | 69.9 | 88.8 | 78.3 | 83.3 | **77.9** | 

### Evaluation Results on MMTEB

Results on the MMTEB benchmark.

