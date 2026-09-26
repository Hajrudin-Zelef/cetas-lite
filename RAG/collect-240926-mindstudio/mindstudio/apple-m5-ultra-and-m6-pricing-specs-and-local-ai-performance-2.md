---
id: collect-240926-mindstudio/mindstudio/apple-m5-ultra-and-m6-pricing-specs-and-local-ai-performance-2
title: "apple-m5-ultra-and-m6-pricing-specs-and-local-ai-performance"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Google", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["agents", "compute", "deepseek", "glm", "gpu", "inference", "memory", "nvidia", "qwen", "training"]
source: docs/RAG/clean_en/mindstudio/apple-m5-ultra-and-m6-pricing-specs-and-local-ai-performance.md
source_anchor: ""
source_lines: [60, 74]
sha256: b3095b0f974499e147b42f0e5b8e05fbdb96d83e8b737708c20b10e5db206fa5
---

# apple-m5-ultra-and-m6-pricing-specs-and-local-ai-performance

### Can the M5 Ultra run models like DeepSeek or GLM locally?

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

Yes, models in that range, including newer GLM releases, DeepSeek models, and larger Qwen variants, are realistic candidates for local use on a machine with this much unified memory, though very large models may still run more smoothly through a cloud provider.

### Is the M5 Ultra built for training AI models or running them?

It’s built for inference, meaning running already-trained models to generate outputs, not training new models from scratch. Training large models from the ground up still typically requires the kind of large-scale GPU clusters that Nvidia hardware dominates.

### Why are companies like OpenAI and Apple building their own AI chips now?

Major AI companies are trying to reduce dependence on Nvidia for compute, either by building custom inference chips (as OpenAI and Google have done) or by offering dedicated local hardware (as Apple is doing with the M5 Ultra). It reflects a broader shift toward companies owning more of their own AI infrastructure stack.
