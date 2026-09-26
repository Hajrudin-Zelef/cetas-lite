---
id: collect-240926-mindstudio/mindstudio/gmktec-evo-x3-vs-evo-x2-which-strix-halo-mini-pc-to-buy-2
title: "gmktec-evo-x3-vs-evo-x2-which-strix-halo-mini-pc-to-buy"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["agent", "agents", "amd", "compute", "gpu", "memory", "nvidia", "tokens per second"]
source: docs/RAG/clean_en/mindstudio/gmktec-evo-x3-vs-evo-x2-which-strix-halo-mini-pc-to-buy.md
source_anchor: ""
source_lines: [67, 84]
sha256: 8f1869cd404b674f64f9e77d0a17e460ec4cf5fceddafe59fa276694bfd8e0b7
---

# gmktec-evo-x3-vs-evo-x2-which-strix-halo-mini-pc-to-buy

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

### Do I need an external GPU to use these mini PCs for AI?

No. The built-in Strix Halo GPU can run large language models on its own, including large mixture-of-experts models, without any external hardware. An eGPU adds speed for models that fit in its VRAM but isn’t required.

### Can you really run one AI model across an AMD and an Nvidia GPU at once?

Yes, using the Vulkan compute API, which works across both vendors’ hardware. A model too large for either GPU alone can have its layers split between them, with both processing every token in sequence.

### What happens if a model is too big for the external GPU’s VRAM?

Performance drops sharply. In testing, a model that overflowed a 16 GB external GPU’s memory fell to roughly 1.6 tokens per second, far slower than the built-in Strix Halo GPU running the same model at around 11 tokens per second.
