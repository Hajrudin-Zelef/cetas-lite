---
id: collect-240926-mindstudio/mindstudio/local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory-2
title: "local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "gpus", "inference", "quantization", "speculative decoding", "video generation"]
source: docs/RAG/clean_en/mindstudio/local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory.md
source_anchor: ""
source_lines: [66, 74]
sha256: 4778a5b85d391ffee03ccf43ea4f94af48233637cf561b7ac8adcad6bf14fe0b
---

# local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory

Not entirely. Discrete GPUs with dedicated VRAM generally offer faster memory bandwidth for models that fit within their capacity. Unified memory’s advantage is capacity: it lets larger models, like 120B parameter mixture-of-experts architectures, load and run at all, which a 32GB card simply cannot do regardless of quantization.

### What is a drafter model in the context of local LLM inference?

A drafter is a smaller, faster model paired with a larger main model. It predicts several likely next tokens, which the larger model then verifies in a batch rather than generating one token at a time. This speculative decoding approach, also called multi-token prediction, can meaningfully increase tokens-per-second without changing output quality.

### Can this kind of setup run image and video generation alongside LLMs?

Yes. The same 128GB unified memory pool supports tools like ComfyUI for image and video generation, including downloaded models and ControlNet-based workflows, in addition to running large language models. This allows a single machine to switch between text, image, and video generation tasks without needing separate dedicated hardware for each.
