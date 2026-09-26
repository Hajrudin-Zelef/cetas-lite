---
id: collect-240926-huggingface/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face-1
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: ["2025-08-07"]
keywords: ["agent", "agents", "attention", "benchmarks", "claude", "context window", "cost", "embedding", "fp8", "inference", "latency", "mixture of experts"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 108]
sha256: ccf20e8e6d07256bc2291cd8d11e014750b7baa88e5d10027e9e8bd5cfa3d4a9
---

# Set the following accordingly

<!-- source: https://huggingface.co/Qwen/Qwen3.5-122B-A10B-FP8 -->

This repository contains FP8-quantized model weights and configuration files for the post-trained model in the Hugging Face Transformers format.

These artifacts are compatible with Hugging Face Transformers, vLLM, SGLang, KTransformers, etc.

The quantization method is fine-grained fp8 quantization with block size of 128, and its performance metrics are nearly identical to those of the original model.


Over recent months, we have intensified our focus on developing foundation models that deliver exceptional utility and performance. Qwen3.5 represents a significant leap forward, integrating breakthroughs in multimodal learning, architectural efficiency, reinforcement learning scale, and global accessibility to empower developers and enterprises with unprecedented capability and efficiency.

Qwen3.5 features the following enhancement:

- **Unified Vision-Language Foundation** : Early fusion training on multimodal tokens achieves cross-generational parity with Qwen3 and outperforms Qwen3-VL models across reasoning, coding, agents, and visual understanding benchmarks.
- **Efficient Hybrid Architecture** : Gated Delta Networks combined with sparse Mixture-of-Experts deliver high-throughput inference with minimal latency and cost overhead.
- **Scalable RL Generalization** : Reinforcement learning scaled across million-agent environments with progressively complex task distributions for robust real-world adaptability.
- **Global Linguistic Coverage** : Expanded support to 201 languages and dialects, enabling inclusive, worldwide deployment with nuanced cultural and regional understanding.
- **Next-Generation Training Infrastructure** : Near-100% multimodal training efficiency compared to text-only training and asynchronous RL frameworks supporting massive-scale agent scaffolds and environment orchestration.

For more details, please refer to our blog post Qwen3.5.

- Type: Causal Language Model with Vision Encoder
- Training Stage: Pre-training & Post-training
- Language Model
  - Number of Parameters: 122B in total and 10B activated
  - Hidden Dimension: 3072
  - Token Embedding: 248320 (Padded)
  - Number of Layers: 48
  - Hidden Layout: 12 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE))
  - Gated DeltaNet:
    - Number of Linear Attention Heads: 64 for V and 16 for QK
    - Head Dimension: 128
  - Gated Attention:
    - Number of Attention Heads: 32 for Q and 2 for KV
    - Head Dimension: 256
    - Rotary Position Embedding Dimension: 64
  - Mixture Of Experts
    - Number of Experts: 256
    - Number of Activated Experts: 8 Routed + 1 Shared
    - Expert Intermediate Dimension: 1024
  - LM Output: 248320 (Padded)
  - MTP: trained with multi-steps
- Context Length: 262,144 natively and extensible up to 1,010,000 tokens.

|  | GPT-5-mini 2025-08-07 | GPT-OSS-120B | Qwen3-235B-A22B | Qwen3.5-122B-A10B | Qwen3.5-27B | Qwen3.5-35B-A3B | 
|---|---|---|---|---|---|---|
| Knowledge |  |  |  |  |  |  | 
| MMLU-Pro | 83.7 | 80.8 | 84.4 | 86.7 | 86.1 | 85.3 | 
| MMLU-Redux | 93.7 | 91.0 | 93.8 | 94.0 | 93.2 | 93.3 | 
| C-Eval | 82.2 | 76.2 | 92.1 | 91.9 | 90.5 | 90.2 | 
| SuperGPQA | 58.6 | 54.6 | 64.9 | 67.1 | 65.6 | 63.4 | 
| Instruction Following |  |  |  |  |  |  | 
| IFEval | 93.9 | 88.9 | 87.8 | 93.4 | 95.0 | 91.9 | 
| IFBench | 75.4 | 69.0 | 51.7 | 76.1 | 76.5 | 70.2 | 
| MultiChallenge | 59.0 | 45.3 | 50.2 | 61.5 | 60.8 | 60.0 | 
| Long Context |  |  |  |  |  |  | 
| AA-LCR | 68.0 | 50.7 | 60.0 | 66.9 | 66.1 | 58.5 | 
| LongBench v2 | 56.8 | 48.2 | 54.8 | 60.2 | 60.6 | 59.0 | 
| STEM & Reasoning |  |  |  |  |  |  | 
| HLE w/ CoT | 19.4 | 14.9 | 18.2 | 25.3 | 24.3 | 22.4 | 
| GPQA Diamond | 82.8 | 80.1 | 81.1 | 86.6 | 85.5 | 84.2 | 
| HMMT Feb 25 | 89.2 | 90.0 | 85.1 | 91.4 | 92.0 | 89.0 | 
| HMMT Nov 25 | 84.2 | 90.0 | 89.5 | 90.3 | 89.8 | 89.2 | 
| Coding |  |  |  |  |  |  | 
| SWE-bench Verified | 72.0 | 62.0 | -- | 72.0 | 72.4 | 69.2 | 
| Terminal Bench 2 | 31.9 | 18.7 | -- | 49.4 | 41.6 | 40.5 | 
| LiveCodeBench v6 | 80.5 | 82.7 | 75.1 | 78.9 | 80.7 | 74.6 | 
| CodeForces | 2160 | 2157 | 2146 | 2100 | 1899 | 2028 | 
| OJBench | 40.4 | 41.5 | 32.7 | 39.5 | 40.1 | 36.0 | 
| FullStackBench en | 30.6 | 58.9 | 61.1 | 62.6 | 60.1 | 58.1 | 
| FullStackBench zh | 35.2 | 60.4 | 63.1 | 58.7 | 57.4 | 55.0 | 
| General Agent |  |  |  |  |  |  | 
| BFCL-V4 | 55.5 | -- | 54.8 | 72.2 | 68.5 | 67.3 | 
| TAU2-Bench | 69.8 | -- | 58.5 | 79.5 | 79.0 | 81.2 | 
| VITA-Bench | 13.9 | -- | 31.6 | 33.6 | 41.9 | 31.9 | 
| DeepPlanning | 17.9 | -- | 17.1 | 24.1 | 22.6 | 22.8 | 
| Search Agent |  |  |  |  |  |  | 
| HLE w/ tool | 35.8 | 19.0 | -- | 47.5 | 48.5 | 47.4 | 
| Browsecomp | 48.1 | 41.1 | -- | 63.8 | 61.0 | 61.0 | 
| Browsecomp-zh | 49.5 | 42.9 | -- | 69.9 | 62.1 | 69.5 | 
| WideSearch | 47.2 | 40.4 | -- | 60.5 | 61.1 | 57.1 | 
| Seal-0 | 34.2 | 45.1 | -- | 44.1 | 47.2 | 41.4 | 
| Multilingualism |  |  |  |  |  |  | 
| MMMLU | 86.2 | 78.2 | 83.4 | 86.7 | 85.9 | 85.2 | 
| MMLU-ProX | 78.5 | 74.5 | 77.9 | 82.2 | 82.2 | 81.0 | 
| NOVA-63 | 51.9 | 51.1 | 55.4 | 58.6 | 58.1 | 57.1 | 
| INCLUDE | 81.8 | 74.0 | 81.0 | 82.8 | 81.6 | 79.7 | 
| Global PIQA | 88.5 | 84.1 | 85.7 | 88.4 | 87.5 | 86.6 | 
| PolyMATH | 67.3 | 54.0 | 60.1 | 68.9 | 71.2 | 64.4 | 
| WMT24++ | 80.7 | 74.4 | 75.8 | 78.3 | 77.6 | 76.3 | 
| MAXIFE | 85.3 | 83.7 | 83.2 | 87.9 | 88.0 | 86.6 | 

* CodeForces: evaluated on our own query set.

* TAU2-Bench: we follow the official setup except for the airline domain, where all models are evaluated by applying the fixes proposed in the Claude Opus 4.5 system card.

* Search Agent: most search agents built on our model adopt a simple context-folding strategy(256k): once the cumulative Tool Response length reaches a preset threshold, earlier Tool Responses are pruned from the history to keep the context within limits.

* WideSearch: we use a 256k context window without any context management.

* MMLU-ProX: we report the averaged accuracy on 29 languages.

* WMT24++: a harder subset of WMT24 after difficulty labeling and rebalancing; we report the averaged scores on 55 languages using XCOMET-XXL.

* MAXIFE: we report the accuracy on English + multilingual original prompts (totally 23 settings).

* Empty cells (--) indicate scores not yet available or not applicable.

