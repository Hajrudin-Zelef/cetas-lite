---
id: collect-240926-mindstudio/mindstudio/qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works-2
title: "qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple", "Hugging Face"]
dates: []
keywords: ["qwen", "agentic", "agents", "benchmark", "cost", "gguf", "gpu", "llama", "llama.cpp", "reasoning", "refusals", "safetensors"]
source: docs/RAG/clean_en/mindstudio/qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works.md
source_anchor: ""
source_lines: [54, 86]
sha256: f1f3e03d199a567deefc1bc5da5b16520daecefc575b8c33b2653133683b2210
---

# qwen3-8-27b-obliterated-how-the-v3-abliterated-model-works

From a technical standpoint, the tradeoffs are laid out with unusual transparency: a documented 2.1-point MMLU cost, category-level benchmark breakdowns, and specific decoding settings that materially affect output quality. The recommended settings (temperature 0, repetition penalty 1.15, no system prompt, thinking mode off for the most direct answers) matter enough that ignoring them reportedly degrades output quality significantly, including looping behavior during greedy decoding without the repetition penalty. For agentic use in coding or pentesting harnesses, the card recommends a small amount of randomness (temperature 0.1 to 0.3) to avoid getting stuck in deterministic loops, along with periodic context summarization.

## What formats is the model available in?

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

The model is distributed as GGUF quantizations for llama.cpp, Ollama, and LM Studio, ranging from Q8_0 (about 27GB, maximum quality) down to Q2_K (about 11GB, minimum viable), with intermediate options at Q6_K, Q5_K_M, Q4_K_M, Q3_K_M, and an experimental IQ4_XS variant. Full-precision bfloat16 safetensors are also available for use with Hugging Face Transformers, split across 29 shards totaling roughly 54GB. MLX support for Apple Silicon is listed as pending, contingent on the mlx_lm library adding support for the underlying Qwen3.5 architecture. GGUF users are advised to use the model’s bundled chat template (via the `--jinja` flag in llama.cpp or built-in template support in Ollama/LM Studio), since V3’s template prefills an empty thinking block so the model answers directly rather than running a reasoning chain first.

It’s worth noting this isn’t the only abliterated take on this base model. A separate quantized abliteration from a different group, huihui-ai, is also available on Hugging Face with its own GGUF builds, indicating there’s active independent interest in uncensoring this particular Qwen release.

## Frequently Asked Questions

### What does “abliterated” mean in an LLM’s name?

It means the model has had its refusal behavior surgically removed from its weights by identifying and eliminating the internal “refusal direction” the model uses to trigger safety declines, rather than being retrained to comply with everything.

### How much capability does Qwen3.8-27B lose when abliterated?

In this release, V3 loses about 2.1 percentage points of MMLU accuracy compared to stock Qwen3.8-27B (82.3% versus 84.5%), with the loss concentrated more heavily in STEM and formal logic subjects than in humanities.

### What is complementary abliteration blending?

It’s a technique used in V2 and V3 of this model where two different weight-surgery methods (SVD, which removes refusals aggressively but damages capability, and LEACE, which preserves capability but removes refusals less completely) are blended together so each method’s weaknesses cancel out.

### Does removing refusals also remove disclaimers and safety lectures?

According to the model card, V1 and V2 removed hard refusals (“I cannot help with that”), but V2 still produced soft deflections, safety-lecture-style non-answers. V3 specifically targets and claims to eliminate those soft deflections as well.

### What hardware do you need to run this model?

Cela dépend de la quantification. Les versions GGUF vont d'environ 11 Go (Q2_K) à 27 Go (Q8_0), tandis que la version complète en safetensors bfloat16 fait environ 54 Go, donc les exigences vont d'un seul GPU grand public à grande VRAM jusqu'à des configurations multi-GPU ou à grande mémoire pour la précision complète.
