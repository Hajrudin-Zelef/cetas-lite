---
id: collect-261001-ia-llm/ia-llm/quantization-aware-healing-a-compressed-4-bit-model-that-outperforms-its-full-precision-or-2
title: "quantization-aware-healing-a-compressed-4-bit-model-that-outperforms-its-full-precision-original"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["quantization", "compute", "distillation", "memory", "research", "training"]
source: docs/RAG/collect-261001-ia-llm/quantization-aware-healing-a-compressed-4-bit-model-that-outperforms-its-full-precision-original.md
source_anchor: ""
source_lines: [47, 55]
sha256: 7104a9531474d98cb286b511337b27a3138f873235a88ca74b997463da98c3d6
---

# quantization-aware-healing-a-compressed-4-bit-model-that-outperforms-its-full-precision-original

*QAH peaks at 54.9 in roughly 100 steps and holds; QAT reaches 54.6 only around step 700, then loses nearly 19 points by step 1,200. Source: paper Figure 3.*

The accuracy story comes paired with the efficiency story that motivated compression in the first place. At 4-bit precision the QAH model uses roughly 4 times less weight memory than the bfloat16 student, and at half the parameter count of the 120B teacher it roughly halves compute per token, which is what lets it run on substantially smaller hardware. For model families that ship in bfloat16 rather than 4-bit, the combined parameter and precision reduction would be closer to 8 times less compute per token.

The takeaway is that a compressed, 4-bit model does not have to be a lower-accuracy version of its full-precision counterpart. With this healing recipe it can be smaller, cheaper to serve, and more accurate at the same time, and it reaches that point in a fraction of the training a QAT recipe would need. Quantization stops being a tax you pay for efficiency and becomes an extra opportunity to teach the model.

This work is part of Multiverse Computing's ongoing research into making large models smaller and cheaper to run without giving up the capabilities that make them useful. It sits alongside our companion work on efficient distillation, which supplies the long-context training machinery QAH depends on.

Want the full technical details, including the healing pipeline, the chunked KL implementation for long-context healing, and the distributed-training findings? Read the full paper, or get in touch with our team to talk about applying compression and healing to your own models.
