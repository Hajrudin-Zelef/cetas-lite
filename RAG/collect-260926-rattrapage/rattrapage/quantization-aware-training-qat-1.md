---
id: collect-260926-rattrapage/rattrapage/quantization-aware-training-qat-1
title: "Quantization-Aware Training (QAT)"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: []
keywords: ["quantization", "training", "benchmarks", "gguf", "inference", "research"]
source: docs/RAG/lot-rattrapage/fine-tuning/Quantization-Aware Training (QAT).md
source_anchor: ""
source_lines: [1, 19]
sha256: 5f7b8759b84a4a77f2e5d405a841ca651a09225d9faafb7d04057350f91a8bbe
---

# Quantization-Aware Training (QAT)

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/blog/quantization-aware-training-qat.md).
# Quantization-Aware Training (QAT)
Quantize models to 4-bit with Unsloth and PyTorch to recover accuracy.
In collaboration with PyTorch, we're introducing QAT (Quantization-Aware Training) in Unsloth to enable **trainable quantization** that recovers as much accuracy as possible. This results in significantly better model quality compared to standard 4-bit naive quantization. QAT can recover up to **70% of the lost accuracy** and achieve a **1–3%** model performance improvement on benchmarks such as GPQA and MMLU Pro.
> **Try QAT with our free** [**Qwen3 (4B) notebook**](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/Qwen3_$4B$_Instruct-QAT.ipynb)
### :books:Quantization
{% columns %}
{% column width="50%" %}
Naively quantizing a model is called **post-training quantization** (PTQ). For example, assume we want to quantize to 8bit integers:
1. Find `max(abs(W))`
2. Find `a = 127/max(abs(W))` where a is int8's maximum range which is 127
3. Quantize via `qW = int8(round(W * a))`
{% endcolumn %}
{% column width="50%" %}
{% endcolumn %}
{% endcolumns %}
Dequantizing back to 16bits simply does the reverse operation by `float16(qW) / a` . Post-training quantization (PTQ) can greatly reduce storage and inference costs, but quite often degrades accuracy when representing high-precision values with fewer bits - especially at 4-bit or lower. One way to solve this to utilize our [**dynamic GGUF quants**](/docs/basics/dynamic-3.0-ggufs.md), which uses a calibration dataset to change the quantization procedure to allocate more importance to important weights. The other way is to make **quantization smarter, by making it trainable or learnable**!
### :fire:Smarter Quantization

