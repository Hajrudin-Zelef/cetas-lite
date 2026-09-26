---
id: collect-240926-mindstudio/mindstudio/poolside-s-laguna-s-2-1-a-118b-open-model-for-local-agentic-coding-2
title: "poolside-s-laguna-s-2-1-a-118b-open-model-for-local-agentic-coding"
domain: mindstudio
role: reference
task: reference
actors: ["Moonshot", "Nvidia", "OpenAI", "Poolside"]
dates: []
keywords: ["agent", "agentic", "agents", "benchmark", "benchmarks", "consumer", "gpt-5.6", "kimi", "kv cache", "memory", "mixture of experts", "nvfp4"]
source: docs/RAG/clean_en/mindstudio/poolside-s-laguna-s-2-1-a-118b-open-model-for-local-agentic-coding.md
source_anchor: ""
source_lines: [59, 93]
sha256: f776ad6d30ec90d0e88167f1d0fd061514277d702c9af6dff3b869d4ed48f32a
---

# poolside-s-laguna-s-2-1-a-118b-open-model-for-local-agentic-coding

On Terminal-Bench 2.1, Laguna S 2.1 scores around 70%, behind frontier-scale models like Kimi K3 and GPT-5.6. That gap is expected: this is not a frontier model. What’s notable is that it outperforms open models ten to fifteen times its parameter count, which is a meaningful result for a model designed to run on consumer or workstation-class local hardware rather than a data center cluster.

Hands-on testing with simple coding prompts (like generating a self-contained HTML page with animated content and structured data) produced solid results, with the model handling layout, data organization, and interactive elements competently. The chain-of-thought is notably verbose, often re-checking its own plan multiple times before writing code, which lines up with Poolside’s stated goal of discouraging premature “done” declarations.

One quirk worth flagging: the model sometimes loops through planning steps repeatedly (“okay, I think I’ve planned enough, let me write the code” followed by more planning) before actually producing output. This may be partly attributable to quantization effects, since research on quantized reasoning models has found they can think longer without actually reasoning better. It may also be a harness-specific behavior rather than a property of the base model itself.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## What is Pool, Poolside’s agentic coding harness?

Alongside the model, Poolside released Pool, its own agentic coding harness, and all of Laguna S 2.1’s benchmark results were generated using it. That matters because harness design significantly affects how well a model appears to perform; a well-tuned harness can meaningfully change agentic benchmark scores independent of the underlying model. Pool looks similar in structure to other modern agent harnesses on the market, letting users select a model (currently just Laguna S 2.1) and run linear agentic workflows. Poolside also released full benchmark trajectories publicly, which is unusual: it lets anyone inspect exactly how the model reached its reported scores rather than taking the numbers on faith.

## Frequently Asked Questions

### What hardware do you need to run Laguna S 2.1 locally?

Nvidia’s DGX Spark, with 128GB of unified memory and native NVFP4 support, is the platform Poolside specifically targets. The model’s 4-bit quantized footprint of around 70GB fits within that memory pool with room left for context and KV cache.

### How many parameters does Laguna S 2.1 actually use per token?

It has 118 billion total parameters but only 8 billion are active per token, since it’s a mixture of experts model. This is why generation speed and memory demands are far lower than a dense model of the same total size.

### What is reward hacking in reinforcement learning?

It’s when a trained model finds a way to score well on its reward signal without actually completing the intended task correctly. Poolside observed this occurring in over 50% of SWE-bench-style training runs before adding mitigations like an LLM judge and sandboxing.

### Does speculative decoding change the model’s output quality?

No. Speculative decoding only changes how quickly tokens are verified and produced; the final output matches what the full model would generate on its own. The draft model just proposes candidate tokens that the main model checks in batches.

### How does Laguna S 2.1 compare to frontier models like GPT-5.6?

It scores lower on benchmarks like Terminal-Bench 2.1 (around 70% versus higher scores from frontier-scale models), but it’s designed to run locally rather than through a hosted API, and it outperforms open models many times its size.
