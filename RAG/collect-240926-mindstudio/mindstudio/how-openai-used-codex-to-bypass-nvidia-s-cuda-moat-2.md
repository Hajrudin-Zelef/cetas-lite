---
id: collect-240926-mindstudio/mindstudio/how-openai-used-codex-to-bypass-nvidia-s-cuda-moat-2
title: "how-openai-used-codex-to-bypass-nvidia-s-cuda-moat"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "Nvidia", "OpenAI"]
dates: []
keywords: ["nvidia", "asic", "attention", "deepseek", "inference", "training"]
source: docs/RAG/clean_en/mindstudio/how-openai-used-codex-to-bypass-nvidia-s-cuda-moat.md
source_anchor: ""
source_lines: [45, 67]
sha256: 2c8b72a6dbeb7b45d21dd6d9202a253d794bb7bd1d8f7a7881aba8f884db4952
---

# how-openai-used-codex-to-bypass-nvidia-s-cuda-moat

What the results do suggest, assuming SemiAnalysis’s independent testing holds up under fuller benchmarking, is that the software moat around chip adoption is no longer an automatic long-term barrier if a well-resourced lab can throw a capable coding model at the problem. That’s a meaningfully different competitive landscape for anyone betting on hardware lock-in as a durable advantage.

## Frequently Asked Questions

### What is Jalapeno?

Jalapeno is OpenAI’s first in-house chip, an ASIC built specifically for running inference on large language models rather than training them or handling general computing tasks.

### What does “performance per watt” measure, and why does it matter?

It measures how much useful computational output a chip produces per unit of electricity consumed. It matters because power availability, not just chip supply, is a major constraint on scaling AI data centers, so higher efficiency per watt effectively expands usable capacity.

### What is Gluon?

Gluon is a kernel programming language OpenAI created for writing the low-level code that tells its chips exactly how to execute AI model operations. It’s described as dense and hand-tuned rather than designed for broad human accessibility.

### Why couldn’t OpenAI run DeepSeek’s model easily on Jalapeno at first?

DeepSeek’s R1 model uses multi-head latent attention (MLA), an architecture most labs, including OpenAI, don’t use internally. That meant OpenAI had no existing kernel optimized for it and needed one written for the new chip.

### Does this mean Nvidia’s CUDA advantage is over?

Not based on current evidence. Jalapeno is a narrow, first-generation inference chip being deployed in small volumes so far, and Nvidia’s CUDA ecosystem still covers far broader use cases. It does suggest that AI-written kernels could erode part of the barrier that has historically protected CUDA’s dominance.
