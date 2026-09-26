---
id: collect-240926-mindstudio/mindstudio/run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings-2
title: "run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple"]
dates: []
keywords: ["gguf", "agent", "agents", "consumer", "cyber", "gpus", "llama", "llama.cpp", "reasoning", "refusals"]
source: docs/RAG/clean_en/mindstudio/run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings.md
source_anchor: ""
source_lines: [73, 111]
sha256: af0562f5ff23082026fb5e1e012f3892da0fb7e25c5896a04e1c0861e19cdee2
---

# run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings

- Keep **repetition_penalty at 1.15** , still described as critical for breaking repetitive loops in agent contexts.
- Raise **temperature to 0.1-0.3** instead of 0. A small amount of randomness helps the model escape deterministic loops that pure greedy decoding can fall into.
- Cap **max_tokens per turn at 1024-2048** rather than letting a single turn run long. Shorter per-turn responses keep an agent focused on the next step.
- **Summarize context after roughly 10 turns.** Agent histories fill up with repeated actions quickly, and trimming or summarizing keeps the model from anchoring on stale context.

## What’s the capability tradeoff?

The MMLU comparisons are the clearest signal. Stock Qwen3.8-27B scores 84.46% (0-shot, lm-eval-harness, 5700 questions). V3 lands at 82.33%, a 2.12 point drop. That’s a meaningful but not drastic hit, and it’s not evenly distributed across subjects. STEM categories dropped the most (-3.3pp), while humanities barely moved (-1.0pp) and a few subjects, including philosophy and European history, actually scored higher than stock. Abstract algebra and formal logic saw larger drops. The pattern suggests the surgery’s refusal directions partially overlap with some structured reasoning pathways, which is a useful thing to know if your use case leans heavily on formal STEM reasoning versus general knowledge or coding tasks.

On the liberation side, the model card reports 20 out of 20 successful completions on a set of code-generation and cyber-related test prompts, and 7 out of 8 on a broader set of “advanced real-world” tasks including async code refactoring, security code review, and Kubernetes debugging scenarios, matching stock’s 7/8 on the same set. Multi-tool chaining failed for both stock and OBLITERATED versions.

## Frequently Asked Questions

### What quant should I start with?

Q4_K_M (roughly 16GB) is the documented sweet spot for most consumer GPUs. If you have 24GB of VRAM, Q5_K_M or Q6_K give better quality with reasonable headroom for context.

### Do I need a system prompt for best results?

No. The model card specifically recommends leaving the system prompt empty, noting that system prompts can reintroduce refusal behavior during A/B testing.

### Why does the model loop or repeat itself?

This happens when repetition_penalty isn’t set. At temperature 0 (greedy decoding), the model needs repetition_penalty around 1.15 to avoid looping on boilerplate like import statements. This is described as essential, not a minor tweak.

### Should thinking mode be on or off?

Off is recommended for the most direct answers. The V3 chat template prefills an empty thinking block by default. Thinking mode can be enabled and doesn’t trigger refusals, but tends to produce longer responses without a clear quality benefit.

### Is there a Mac-native version?

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Not yet. MLX support for Apple Silicon is waiting on the `mlx_lm` library to add support for the underlying Qwen3.5 architecture. GGUF quants can still run on Mac through llama.cpp in the meantime.
