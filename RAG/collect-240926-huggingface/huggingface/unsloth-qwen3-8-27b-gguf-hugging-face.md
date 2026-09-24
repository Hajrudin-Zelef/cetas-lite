---
id: collect-240926-huggingface/huggingface/unsloth-qwen3-8-27b-gguf-hugging-face
title: "Read our How to Run Qwen3.8-27B Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["qwen", "agent", "agentic", "attention", "embedding", "gguf", "inference", "parameters", "quantization", "reasoning", "research", "rotary"]
source: docs/RAG/clean_en/huggingface/unsloth-qwen3-8-27b-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 77]
sha256: f9c7369edb673c26b860c05e1814e0c8ca5d5419faadaa8ce8d51858eb92d27d
---

# Read our How to Run Qwen3.8-27B Guide!

<!-- source: https://huggingface.co/unsloth/Qwen3.8-27B-GGUF -->

# Read our How to Run Qwen3.8-27B Guide!

    *Unsloth Dynamic 3.0 achieves superior accuracy & outperforms other leading quants.*
  

- Introducing Dynamic V3.0 GGUFs for SOTA accuracy and quantization performance
- Run and fine-tune Qwen3.8 in Unsloth Desktop with **Thinking toggles** . Download for Mac, Windows and Linux. GitHub repo
- Developer Role Support so Qwen3.8 can work in agentic tools like Codex and more!
- Tool calling improvements: Makes parsing nested objects to make tool calling succeed more.
- See below for 4-bit Qwen3.8-27B run inside of Unsloth Desktop:

Analysis of best Qwen3.8 GGUF providers. Unsloth Dynamic v3.0 delivers >10% top-1% better accuracy at the same size compared to every other provider. Read more

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

To achieve optimal performance, we recommend the following settings:

1. **Sampling Parameters** : We suggest using the following sets of sampling parameters:
  - Thinking Mode: `temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
  - Instruct (or non-thinking) mode: `temperature=0.7` ,`top_p=0.80` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
 For supported frameworks, you can adjust the `presence_penalty` parameter between 0 and 2 to reduce endless repetition. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. Thinking Mode: 
3. **Adequate Output Length** : To optimize performance on agentic tasks, we recommend allocating sufficient output length to allow the model to generate detailed and comprehensive responses. For frameworks that support separate token limits for internal reasoning and final outputs, we suggest the following configuration within the 1M context length:
  - Reasoning Content: Set the maximum output length to 262,144 tokens.
  - Final Response: Set the maximum output length to 131,072 tokens.
 These settings provide the necessary capacity for complex reasoning while ensuring ample space for high-quality final deliverables.
4. **Processing Ultra-Long Texts** : Qwen3.8-27B natively supports context lengths of up to 262,144 tokens. For long-horizon tasks where the total length (including both input and output) exceeds this limit, we recommend using RoPE scaling techniques to handle long texts effectively, e.g., YaRN.
5. **Long Video Understanding** : To optimize inference efficiency for plain text and images, the`size` parameter in the released`video_preprocessor_config.json` is conservatively configured. It is recommended to set the`longest_edge` parameter in the video_preprocessor_config file to 469,762,048 (corresponding to 224k video tokens) to enable higher frame-rate sampling for hour-scale videos and thereby achieve superior performance. For example,```
{"longest_edge": 469762048, "shortest_edge": 4096}
```

If you find our work helpful, feel free to give us a cite.

```
@misc{qwen38,
    title = {{Qwen3.8-Max}: A New Bar for Coding and Cowork},
    url = {https://qwen.ai/blog?id=qwen3.8},
    author = {{Qwen Team}},
    month = {August},
    year = {2026}
}
```
- Downloads last month
- 7,063,930
