---
id: collect-261001-ia-llm/ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers-8
title: "Native Sentence Transformers checkpoints. PyLate builds on the same schema,"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["benchmark", "compute", "embedding", "embeddings", "energy", "llama", "lora", "multimodal", "nvidia", "omni", "parameters", "reranker"]
source: docs/RAG/collect-261001-ia-llm/multi-vector-late-interaction-embedding-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [680, 758]
sha256: 81602fb9b6b9ddfcdab4ffca69f9c554daed8d0652319c73653af3dadeb6add2
---

# Native Sentence Transformers checkpoints. PyLate builds on the same schema,

The NanoBEIR column reports the mean NDCG@10 (higher is better) across the 13 NanoBEIR datasets, each a 50-query subsample of a BEIR dataset, as a fast proxy for English text retrieval quality. We used the `MultiVectorNanoBEIREvaluator` to compute the scores for the primarily-English models. A `-` means the model was not evaluated on it. Note that NanoBEIR is a small benchmark, and its scores aren't a substitute for evaluating on your own data, which is always the right way to pick a model.

ColPali-style models embed page images as documents and text as queries.

The NanoViDoRe column reports the mean NDCG@10 (higher is better) across NanoViDoRe v3, a compact visual document retrieval benchmark spanning 8 subsets (computer science, energy, finance in English and French, HR, industrial, pharmaceuticals, and physics). Like with NanoBEIR, NanoViDoRe is a small benchmark which shouldn't replace evaluation on your own data.

| Model | Parameters | Dimensionality | NanoViDoRe | Notes | 
|---|---|---|---|---|
| webAI-Official/webAI-ColVec1.1-8b | 8.4B | 640 | 0.6580 | needs `trust_remote_code=True` | 
| webAI-Official/webAI-ColVec1.1-4b | 4.5B | 640 | 0.6520 | needs `trust_remote_code=True` | 
| vultr/VultronRetrieverPrime-Qwen3.5-8B | 8.39B | 320 | 0.6423 | `revision="refs/pr/2"` | 
| vultr/VultronRetrieverCore-Qwen3.5-4.5B | 4.54B | 320 | 0.6410 | `revision="refs/pr/1"` | 
| tencent/EVIE-Preview-4.5B | 4.54B | 128 | 0.6405 | - | 
| nvidia/nemotron-colembed-vl-8b-v2 | 8.77B | 4096 | 0.6374 | `revision="refs/pr/4"` , needs`trust_remote_code=True` | 
| athrael-soju/colqwen3.5-4.5B-v3 | 4.54B | 320 | 0.6358 | - | 
| TomoroAI/tomoro-colqwen3-embed-8b | 8.8B | 320 | 0.6206 | needs `trust_remote_code=True` | 
| nvidia/nemotron-colembed-vl-4b-v2 | 4.83B | 2560 | 0.6200 | `revision="refs/pr/7"` , needs`trust_remote_code=True` | 
| OpenSearch-AI/Ops-Colqwen3-4B | 4.44B | 2560 | 0.6150 | `revision="refs/pr/5"` , needs`trust_remote_code=True` | 
| nvidia/llama-nemotron-colembed-vl-3b-v2 | 4.41B | 3072 | 0.6112 | `revision="refs/pr/3"` , needs`trust_remote_code=True` | 
| tsystems/colqwen2.5-3b-multilingual-v1.0 | 3.75B | 128 | 0.6039 | `revision="refs/pr/3"` | 
| tsystems/colqwen2.5-3b-multilingual-v1.0-merged | 3.75B | 128 | 0.6027 | `revision="refs/pr/1"` | 
| TomoroAI/tomoro-colqwen3-embed-4b | 4.4B | 320 | 0.6019 | needs `trust_remote_code=True` | 
| nomic-ai/colnomic-embed-multimodal-7b | 8.29B | 128 | 0.5942 | `revision="refs/pr/3"` | 
| nomic-ai/colnomic-embed-multimodal-3b | 3.75B | 128 | 0.5929 | `revision="refs/pr/6"` | 
| VAGOsolutions/SauerkrautLM-ColQwen3-8b-v0.1 | 8.15B | 128 | 0.5819 | `revision="refs/pr/1"` | 
| Metric-AI/ColQwen2.5-3b-multilingual-v1.0 | 3.75B | 128 | 0.5763 | `revision="refs/pr/2"` | 
| vultr/VultronRetrieverFlash-Qwen3.5-0.8B | 853M | 320 | 0.5693 | `revision="refs/pr/2"` | 
| VAGOsolutions/SauerkrautLM-ColQwen3-4b-v0.1 | 4.44B | 128 | 0.5656 | `revision="refs/pr/1"` | 
| VAGOsolutions/SauerkrautLM-ColQwen3-2b-v0.1 | 2.13B | 128 | 0.5530 | `revision="refs/pr/1"` | 
| Verm1ion/ColTurk-VDR-Qwen3VL-4B-v1.0 | 4.44B | 320 | 0.5434 | `revision="refs/pr/1"` | 
| vidore/colqwen2.5-v0.2 | 3.8B | 128 | 0.5402 | - | 
| vidore/colqwen2.5-v0.1 | 3.8B | 128 | 0.5395 | - | 
| vidore/colqwen-omni-v0.1 | 4.4B | 128 | 0.5309 | - | 
| VAGOsolutions/SauerkrautLM-ColQwen3-1.7b-Turbo-v0.1 | 1.76B | 128 | 0.5035 | `revision="refs/pr/1"` | 
| vidore/colpali-v1.3 | 2.9B | 128 | 0.4802 | - | 
| vidore/colpali-v1.3-hf | 2.9B | 128 | 0.4793 | - | 
| vidore/colpali-v1.2 | 2.9B | 128 | 0.4691 | - | 
| vidore/colqwen2-v1.0 | 2.2B | 128 | 0.4685 | - | 
| vidore/colqwen2-v0.1 | 2.2B | 128 | 0.4526 | - | 
| vidore/colpali | 2.9B | 128 | 0.4516 | - | 
| vidore/colpali-v1.1 | 2.9B | 128 | 0.4314 | - | 
| VAGOsolutions/SauerkrautLM-ColLFM2-450M-v0.1 | 451M | 128 | 0.4249 | `revision="refs/pr/1"` | 
| vidore/colsmolvlm-v0.1 | 2.1B | 128 | 0.4054 | - | 
| VAGOsolutions/SauerkrautLM-ColMinistral3-3b-v0.1 | 4.25B | 128 | 0.3978 | `revision="refs/pr/1"` | 
| vidore/colpali-hard-v1.1 | 2.9B | 128 | 0.3949 | - | 
| vidore/colSmol-500M | 507M | 128 | 0.3459 | - | 
| vidore/colSmol-256M | 256M | 128 | 0.2673 | - | 
| ModernVBERT/colmodernvbert | 252M | 128 | 0.2632 | - | 
| vidore/colpali-v1.2-hf | 2.9B | 128 | - | - | 
| vidore/colqwen2-v1.0-hf | 2.2B | 128 | - | - | 

Most of these are LoRA adapter repositories, with the adapter applied directly onto its base at load time. Some also have a `-merged` sibling on the Hub (e.g. vidore/colpali-v1.3-merged) with the adapter already folded into the weights.

The three `-hf` entries are the transformers-native `*ForRetrieval` ports. They load without any configuration, but use more modeling from `transformers` and less from `sentence_transformers`. Generally, it's preferable to use the original models instead, as the ports score approximately the same.

Late interaction in Sentence Transformers rests on a lot of earlier work. Thanks to Omar Khattab and Matei Zaharia for ColBERT, which everything here descends from, and to the LightOn team (Antoine Chaffin, Raphael Sourty, Paulo Moura, and Amélie Chatelain) for PyLate and fast-plaid, which carried late interaction for years and shaped a good deal of the API described above.

Thanks to the ColPali team (Manuel Faysse, Hugues Sibille, Tony Wu, Bilel Omrani, Gautier Viaud, Céline Hudelot, and Pierre Colombo) for ColPali and colpali-engine, which brought late interaction to page images, and to Benjamin Clavié, Antoine Chaffin, and Griffin Adams for token pooling.

Thanks as well to the core MTEB team, Kenneth Enevoldsen and Roman Solomatin among many others, for MTEB and for the kind of hidden work that keeps information retrieval research running.

And thanks to everyone who trained and released the checkpoints in Supported Models. Without them this post would have had nothing to measure.

To learn how to train or finetune these models on your own data:

See the companion blogpost: Training and Finetuning Multi-Vector Embedding Models with Sentence Transformers.

- Multi-Vector Encoder > Training Overview
- Multi-Vector Encoder > Loss Overview
- Multi-Vector Encoder > Training Examples
- LateOn and mLateOn training scripts: LightOn's PyLate recipes for LateOn, mLateOn, DenseOn, and mDenseOn, where the finetuning scripts show practical details like splitting a 16,384-example batch into mini-batches of 16.

- Training and Finetuning Multi-Vector Embedding Models with Sentence Transformers: the direct training companion to this post.
- Training and Finetuning Embedding Models with Sentence Transformers: the general training guide for text-only dense embedding models.
- Training and Finetuning Reranker Models with Sentence Transformers: Cross Encoder training, the other way to add a precise second stage.
- Training and Finetuning Sparse Embedding Models with Sentence Transformers: SPLADE and other sparse encoders, which combine well with late interaction in hybrid search.
- Multimodal Embedding & Reranker Models with Sentence Transformers: single-vector multimodal models, the dense counterpart to ColPali-style retrieval.
- Training and Finetuning Multimodal Embedding & Reranker Models with Sentence Transformers: includes a Visual Document Retrieval walkthrough with single-vector models.
- 🪆 Introduction to Matryoshka Embedding Models: shrink dense embeddings by dimension, the way token pooling shrinks multi-vector ones by count.
