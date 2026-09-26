---
id: collect-240926-mindstudio/mindstudio/how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests-2
title: "how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "OpenAI"]
dates: []
keywords: ["deepseek", "quantization", "agents", "agi", "benchmark", "cost", "gpu", "inference", "inference engine", "kv cache", "memory", "multimodal"]
source: docs/RAG/clean_en/mindstudio/how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests.md
source_anchor: ""
source_lines: [64, 84]
sha256: bf057e1833e2db6fbb42a9f206075f91475187c2cbcde849ce70b8fb76afe01a
---

# how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests

At 4-bit quantization, expect to need about 168 GB of VRAM for model loading. At 3-bit quantization, that drops to around 110 GB. Both figures exclude the additional memory needed for context, so real usage will require more.

### Can DeepSeek V4 Flash handle images or audio?

No. It’s a text-only model with no native multimodal support. Anyone needing vision or audio input has to pair it with a separate multimodal model.

### Why did DeepSeek V4 Flash’s coding benchmark scores jump so much?

Part of the jump reflects genuine post-training improvements, but part of it comes from being tested under DeepSeek’s own optimized harness, which other models weren’t necessarily evaluated with. Harness quality can swing scores dramatically, as shown by OpenAI’s Arc-AGI harness comparison.

### What is Dwarf Star and how does it help with local deployment?

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

Dwarf Star is a custom inference engine built for DeepSeek-style architectures that uses SSD offloading, KV cache manipulation, and mixed precision weight loading to run large models on smaller hardware, such as a single 128 GB system, while preserving memory for context.

### How does DeepSeek V4 Flash compare on price to running it locally?

Via API, pricing runs around 2 cents per million input tokens and close to 30 cents per million output tokens, which is hard to beat on cost alone. Local deployment removes ongoing API costs but requires upfront investment in multi-GPU hardware capable of meeting the VRAM requirements.
