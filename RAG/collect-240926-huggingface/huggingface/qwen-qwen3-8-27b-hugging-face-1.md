---
id: collect-240926-huggingface/huggingface/qwen-qwen3-8-27b-hugging-face-1
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "Meta", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "attention", "benchmark", "claude", "context window", "embedding", "inference", "multimodal", "muse", "parameters"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-8-27b-hugging-face.md
source_anchor: ""
source_lines: [1, 89]
sha256: 3c48d4e56ae05b0fa1807e4ee0eb19a28f6632f693ad9918ff2d9084efb654ea
---

# Set the following accordingly

<!-- source: https://huggingface.co/Qwen/Qwen3.8-27B -->

This repository contains model weights and configuration files for the post-trained model in the Hugging Face Transformers format.

These artifacts are compatible with Hugging Face Transformers, vLLM, SGLang, TokenSpeed, etc.


For users seeking managed, scalable inference without infrastructure maintenance, the official Qwen API service is provided by Qwen Cloud.
In particular, **Qwen3.8-27B** will be available as a hosted version with more production features, e.g., 1M context length by default, official built-in tools. For more information, please refer to the Qwen3.8-27B Overview. The service is coming soon. Stay tuned for updates.


Following the widespread community adoption of the Qwen3.5 and Qwen3.6 series, we are pleased to introduce Qwen3.8, the most capable generation in the Qwen open-model family to date.

Built on the architectural foundation of Qwen3.5, Qwen3.8 delivers substantial gains across coding, professional work, research, and long-horizon agentic tasks. Qwen3.8-27B brings these advances to a compact, deployment-friendly dense model: a native vision-language model that understands images and videos, with flexible thinking control, designed to carry complex, multi-step tasks through to completion with greater reliability.

Qwen3.8-27B features the following enhancements:

- **Core Capabilities** : Comprehensive improvements across coding, professional work, research, and long-horizon agentic tasks.
- **Agent Execution** : Stronger autonomous planning and better handling of environment feedback, leading to more reliable end-to-end task completion.
- **Downstream Compatibility** : Broader support for popular harnesses and development tools, making it easier to integrate into your existing stack.
- **Flexible Thinking Control** : Thinking mode is on by default and can be disabled per request; reasoning depth can be tuned with`reasoning_effort` , and reasoning context from historical messages is retained via`preserve_thinking` .
- **Vision-Language Understanding** : Native support for image and video understanding, from STEM diagrams and documents to hour-scale videos.

- Type: Causal Language Model with Vision Encoder
- Training Stage: Pre-training & Post-training
- Language Model
  - Number of Parameters: 27B
  - Hidden Dimension: 5120
  - Token Embedding: 248,320 (Padded)
  - Number of Layers: 64
  - Hidden Layout: 16 × (3 × (Gated DeltaNet → FFN) → 1 × (Gated Attention → FFN))
  - Gated DeltaNet:
    - Number of Linear Attention Heads: 48 for V and 16 for QK
    - Head Dimension: 128
  - Gated Attention:
    - Number of Attention Heads: 24 for Q and 4 for KV
    - Head Dimension: 256
    - Rotary Position Embedding Dimension: 64
  - Feed Forward Network:
    - Intermediate Dimension: 17,408
  - LM Output: 248,320 (Padded)
  - MTP (Multi-Token Prediction): trained with multiple steps
- Context Length: 262,144 natively and extensible up to 1,000,000 tokens.

|  | Qwen3.8-27B | Qwen3.6-27B | Qwen3.7-Plus | Muse Glimmer-30B | Opus4.6 Max | 
|---|---|---|---|---|---|
| Coding |  |  |  |  |  | 
| Agentic terminal coding Terminal Bench 2.1 (Terminus) | 73.0 | 63.4 | 64.0 | 51.7 | **78.2** | 
| Agentic coding SWE-bench Pro | **61.7** | 53.5 | 57.6 | 51.2 | 53.4 | 
| Repo-level code generation NL2Repo-Bench | 42.3 | 36.2 | 41.1 | -- | **47.6** | 
| Agentic coding DeepSWE 1.1 | **42.2** | 13.3 | 14.2 | -- | -- | 
| Software engineering QwenSWEBench | **79.0** | 49.3 | 59.2 | -- | 63.8 | 
| Agent |  |  |  |  |  | 
| Long-horizon office work CoWorkBench | **70.7** | 61.0 | 65.1 | -- | 68.2 | 
| Professional job tasks JobBench | **33.4** | 21.8 | 27.6 | -- | -- | 
| Frontier agentic tasks Agents' Last Exam | Pass@1 **20.4**Score **42.9** | Pass@1 10.6 Score 27.3 | Pass@1 13.2 Score 33.6 | -- | -- | 
| General |  |  |  |  |  | 
| Instruction following IFBench | **79.5** | 69.1 | 79.1 | 77.0 | 62.5 | 
| Scientific reasoning GPQA Diamond | 89.2 | 87.8 | 90.3 | 83.5 | **91.3** | 
| Multidisciplinary reasoning HLE | 30.8 | 24.0 | 34.7 | 22.0 | **40.0** | 
| Competitive coding LiveCodeBench v6 | **90.3** | 83.9 | 89.6 | -- | 88.8 | 

1. SWE-bench Pro: Except for Opus4.6 Max, which uses the officially reported score, all models are evaluated with the Claude Code harness at temp=1.0, top_p=0.95, and a 256K context window. Problematic tasks were corrected, and all baseline models were re-evaluated on the refined benchmark.
2. NL2Repo-Bench: Evaluated with the Claude Code harness. To prevent reward hacking, we disable Bash commands that attempt to access the specific repository, such as pip download, pip install, and git clone.
3. DeepSWE 1.1: Evaluated with the Claude Code harness at temp=1.0, top_p=0.95, and a 256K context window.
4. QwenSWEBench: In-house coding benchmark for evaluating models' software engineering capabilities. Evaluated with the Claude Code harness. Reporting avg@3 with an 8-hour timeout, max_tokens=32,768, temperature=1.0, and a 256K context window.
5. CoWorkBench: In-house cowork benchmark for evaluating long-horizon tasks across computer science, finance, law, medical, and other productivity domains.
6. HLE: Judged by GPT-4o.
7. The best result in each row is shown in bold.
8. Empty cells (--) indicate that results are not yet available or not applicable.

|  | Qwen3.8-27B | Qwen3.6-27B | Qwen3.7-Plus | Muse Glimmer-30B | Opus4.6 Max | 
|---|---|---|---|---|---|
| Agentic Multimodal Intelligence |  |  |  |  |  | 
| Computer use OSWorld-Verified | **84.3** | 63.9 | 73.3 | 65.9 | 72.7 | 
| Browser use WebArena-Verified | **64.8** | 48.8 | 55.3 | -- | -- | 
| Mobile use AndroidWorld | **81.9** | 70.3 | 81.0 | -- | 62.0 | 
| Application recreation RecreationBench | **47.1** | 29.8 | 30.2 | -- | -- | 
| Multimodal tool use ClawEval-MM | Pass@3 **57.4**Average 56.9 | Pass@3 42.6 Average 50.4 | Pass@3 **57.4**Average **60.1** | -- | Pass@3 52.5 Average 54.7 | 
| Multimodal software engineering SWE-MM | **38.6** | 25.7 | 30.0 | -- | 27.1 | 
| Visual web development Vision2Web | **62.9** | 45.0 | 42.1 | -- | -- | 
| General Multimodal Intelligence |  |  |  |  |  | 
| Visual math problem solving MathVision | Without CI 90.0 With CI **94.6** | Without CI 85.1 | Without CI **90.3** | -- | Without CI 65.5 | 
| General visual reasoning BabyVision | Without CI **65.7**With CI **85.6** | Without CI 28.9 | Without CI 64.7 With CI 70.4 | -- | Without CI 12.6 | 
| Scientific chart analysis CharXiv (RQ) | Without CI 83.7 With CI **90.2** | Without CI 78.4 | Without CI **85.8**With CI 85.9 | 78.8 | Without CI 66.0 | 
| Document intelligence OmniDocBench 1.5 | 91.1 | 89.4 | **91.4** | 75.8 | 86.6 | 
| Real-world perception RealWorldQA | 85.9 | 84.1 | **86.9** | -- | 73.9 | 
| Embodied intelligence ERQA | 65.5 | 62.5 | **69.8** | -- | 40.8 | 

