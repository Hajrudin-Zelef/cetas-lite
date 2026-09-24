---
id: collect-240926-mindstudio/mindstudio/local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory
title: "local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Alibaba", "Unsloth"]
dates: []
keywords: ["memory", "qwen", "agent", "amd", "benchmark", "benchmarks", "consumer", "fine-tuning", "gpu", "gpus", "inference", "latency"]
source: docs/RAG/clean_en/mindstudio/local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory.md
source_anchor: ""
source_lines: [1, 74]
sha256: 6c43d3cb1e5ed674dfd56d367e604b439f19dd427137d0fffd50ca2d529e54f7
---

# local-llm-speed-test-gpt-oss-qwen3-6-and-hermes-on-128gb-unified-memory

<!-- source: https://www.mindstudio.ai/blog/local-llm-benchmarks-ryzen-ai-halo -->

## How fast do local LLMs actually run on unified memory hardware?

On a 128GB unified memory system (AMD’s Ryzen AI Max Plus 395 chip, tested via LM Studio with the ROCm Llama.cpp runtime), a dense 9B model with a speculative drafter ran at just under 40 tokens per second, a 35B mixture-of-experts Qwen3.6 model with a 3B active parameter count and drafter exceeded 50 tokens per second, and the full GPT-OSS 120B mixture-of-experts model landed around 30 tokens per second. Those numbers matter because they come from a machine that can actually hold these models in memory at once, something a 32GB discrete GPU can’t do without heavy quantization or offloading penalties.

## TL;DR

- **Unified memory changes the math** for local inference because the CPU and GPU share one 128GB pool of LPDDR5X, letting users allocate the majority of it to the GPU instead of hitting a hard VRAM ceiling.
- **GPT-OSS 120B ran at roughly 30 tokens per second** , a usable speed for chat and RAG workloads but slower than what smaller models achieve, and a candidate to swap out for agent tasks that need faster turnaround.
- **Qwen3.6 as a mixture-of-experts model with 3B active parameters hit over 50 tokens per second** , outperforming the dense GPT-OSS model despite Qwen3.6’s larger total parameter count, because only a fraction of parameters activate per token.
- **A 9B dense model with a multi-token-prediction drafter ran just under 40 tokens per second** , showing that speculative decoding remains a meaningful speed lever even on non-discrete-GPU hardware.
- **Speculative decoding and drafter models are becoming standard** , with releases like Gemma and the Qwen3.5/3.6 series shipping drafter variants specifically to boost inference throughput.
- **Previous 32GB discrete GPU setups hit a wall with models like 120B parameter architectures** , since no amount of quantization gets a model that size to fit, while the unified memory approach sidesteps that limit entirely.
- **The tested hardware came with LM Studio, ComfyUI, and Llama.cpp preinstalled with ROCm drivers already configured** , removing a common setup barrier for people trying to benchmark or run local models quickly.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

## What hardware was used for these benchmarks?

The tests ran on AMD’s Ryzen AI Max Plus 395, a chip built around 16 CPU cores, a Radeon 8060S integrated GPU rated around 60 teraflops at FP16, and an NPU rated near 50 TOPS. The standout spec is 128GB of unified LPDDR5X memory shared between CPU and GPU, with the split configurable by the user. In the tested setup, 75% of that pool was allocated to the GPU for AI workloads, leaving the rest for the CPU and system tasks.

This differs fundamentally from a discrete GPU setup. A card with 32GB of VRAM caps out fast: once a model plus context exceeds that VRAM, layers get offloaded to system RAM and token speeds fall sharply, even on machines with fast RAM. Unified memory avoids that cliff because there’s no separate, smaller pool to overflow. The tradeoff is that integrated graphics memory bandwidth typically doesn’t match a discrete GPU’s dedicated VRAM, so raw speed on models that do fit in 32GB may still favor the discrete card. The benefit shows up specifically on models too large for consumer GPU memory.

Notably, most current LLM inference stacks aren’t yet using the NPU on this chip, meaning the 50 TOPS of dedicated AI acceleration sits mostly idle during standard LLM inference. That’s expected to change as software support catches up.

## Why did Qwen3.6 outperform GPT-OSS 120B despite being smaller in some configurations?

The Qwen3.6 model tested was a mixture-of-experts (MoE) architecture with roughly 3B active parameters per token, versus GPT-OSS 120B, which is also MoE but activates a larger portion of its parameters. In MoE architectures, total parameter count and active parameter count are different numbers, and speed correlates much more closely with active parameters than total size. That’s why a 35B-total Qwen3.6 model with only 3B active can outrun a 120B model on tokens per second: fewer parameters need to be computed per generated token, even though the full model still needs to be loaded into memory.

This matters for anyone choosing models for agent workloads. GPT-OSS 120B’s roughly 30 tokens per second is fine for chat interfaces or retrieval-augmented generation where a user is reading along, but agent loops that chain many calls together benefit more from faster, smaller-active-parameter models like the Qwen3.6 variant tested, which cleared 50 tokens per second.

## Is speculative decoding worth using for local inference?

Based on the tested results, yes. Both the 9B dense model and the Qwen3.6 MoE model used a drafter (a smaller model that predicts likely next tokens ahead of the main model, which then verifies them in batches). This is often called multi-token prediction or speculative decoding, and it’s increasingly shipped by model creators rather than left to end users to configure. Recent releases in the Gemma and Qwen3.5/3.6 families include drafter variants specifically for this purpose.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

The practical effect: token throughput increases without a quality tradeoff, since the larger model still verifies every token the drafter proposes. For anyone running local inference on hardware without a discrete GPU’s raw bandwidth, a drafter model is one of the more reliable ways to close the speed gap.

## What can this kind of hardware handle beyond text generation?

The same system ran ComfyUI locally for image and video generation, using preinstalled models alongside downloaded ones like the Qwen image model. Render times for individual images landed around 19 to 20 seconds, fast enough to iterate through dozens of prompt variations in a single session rather than waiting overnight, as might be assumed for a fully local pipeline. Workflows using ControlNet techniques (edge detection, pose estimation) were also functional, showing that the memory headroom benefits image and video pipelines as well as LLM inference, not just larger context windows for text models.

The system also supports fine-tuning workflows (tested conceptually with Unsloth, a common toolkit for efficient fine-tuning) and custom GPU kernel development, suggesting the unified memory advantage extends past inference into training and experimentation for people willing to go deeper than out-of-the-box chat.

## Frequently Asked Questions

### What token-per-second speed is considered good for local LLM chat use?

Based on these benchmarks, anything from 30 to 50+ tokens per second feels responsive for chat and RAG applications. GPT-OSS 120B at around 30 tokens per second was usable for conversational tasks, while the Qwen3.6 MoE model above 50 tokens per second is better suited to latency-sensitive agent workflows.

### Why does a 120B parameter model run slower than a 35B model on the same hardware?

It comes down to active parameters, not total size. GPT-OSS 120B activates more parameters per generated token than the tested Qwen3.6 configuration, which used a mixture-of-experts design with only about 3B active parameters. More active computation per token means slower generation, regardless of how many parameters are stored in memory overall.

### Does unified memory replace the need for a discrete GPU for local AI?

Not entirely. Discrete GPUs with dedicated VRAM generally offer faster memory bandwidth for models that fit within their capacity. Unified memory’s advantage is capacity: it lets larger models, like 120B parameter mixture-of-experts architectures, load and run at all, which a 32GB card simply cannot do regardless of quantization.

### What is a drafter model in the context of local LLM inference?

A drafter is a smaller, faster model paired with a larger main model. It predicts several likely next tokens, which the larger model then verifies in a batch rather than generating one token at a time. This speculative decoding approach, also called multi-token prediction, can meaningfully increase tokens-per-second without changing output quality.

### Can this kind of setup run image and video generation alongside LLMs?

Yes. The same 128GB unified memory pool supports tools like ComfyUI for image and video generation, including downloaded models and ControlNet-based workflows, in addition to running large language models. This allows a single machine to switch between text, image, and video generation tasks without needing separate dedicated hardware for each.
