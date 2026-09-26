---
id: collect-240926-mindstudio/mindstudio/run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings-1
title: "run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple"]
dates: []
keywords: ["gguf", "agent", "agentic", "alignment", "consumer", "cost", "cyber", "datacenter", "fine-tuning", "gpu", "gpus", "llama"]
source: docs/RAG/clean_en/mindstudio/run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings.md
source_anchor: ""
source_lines: [1, 72]
sha256: f12d67b14c333b0405dcf1809317875bbfa4c76797c7fdde5cc17430e0c74885
---

# run-qwen3-8-27b-obliterated-locally-gguf-sizes-vram-and-settings

<!-- source: https://www.mindstudio.ai/blog/run-qwen3-8-27b-obliterated-locally -->

## What is Qwen3.8-27B-OBLITERATED?

Qwen3.8-27B-OBLITERATED is a modified version of Alibaba’s Qwen3.8-27B base model with its refusal behavior surgically removed through a technique called abliteration. The current release, V3, targets both hard refusals (“I cannot help with that”) and softer safety-lecture deflections, aiming to answer restricted queries with actual substance rather than a warning label. It’s distributed as GGUF files for llama.cpp-based tools like Ollama and LM Studio, as full bfloat16 safetensors for Transformers, and MLX support is pending upstream architecture compatibility.

## TL;DR

- **Qwen3.8-27B-OBLITERATED V3** removes both hard refusals and soft safety-lecture deflections from stock Qwen3.8-27B, scoring 20/20 on a set of code and cyber-related test prompts where the base model would normally refuse.
- The capability cost is measurable but modest: **MMLU drops 2.1 percentage points** versus stock (82.33% vs 84.46%), with STEM categories taking the biggest hit and humanities barely affected.
- GGUF quantizations range from **Q2_K at roughly 11GB up to Q8_0 at roughly 27GB** , letting the same model run on anything from a single mid-range GPU to a high-VRAM workstation card.
- Getting good output depends heavily on settings: **temperature 0 with repetition_penalty 1.15** is required, since greedy decoding without a repetition penalty loops on imports and boilerplate.
- The model ships with a **chat template that prefills an empty thinking block** , so GGUF users need to load the bundled template (via`--jinja` in llama.cpp, or the model’s built-in template in Ollama/LM Studio) to skip straight to answering.
- For agent or long-context use, the recommended settings shift slightly: a touch of temperature (0.1 to 0.3) helps break loops that pure greedy decoding can fall into during tool calls.
- The project’s own documentation frames this as a research tool for **alignment researchers, red-teamers, and safety evaluators** , not a general-purpose consumer chatbot.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

## How does abliteration actually work?

Abliteration is a weight-space editing technique, not fine-tuning. Instead of retraining the model on new data, it identifies “refusal directions”, specific vectors in the model’s internal representations that correlate with refusal behavior, and projects them out of the weight matrices. The idea is to surgically remove the mechanism that triggers a refusal while leaving the rest of the model’s knowledge and reasoning intact.

The OBLITERATED project went through three iterations to get this right:

**V1** used a single aggressive SVD (singular value decomposition) pass removing five directions. It killed refusals completely but cost 6 percentage points of MMLU, a noticeable capability hit.

**V2** introduced what the creators call complementary abliteration blending: running two different surgery methods (aggressive SVD and LEACE, a technique that minimizes mutual information) that fail in different ways, then blending their weights at a 60/40 ratio. SVD is good at finding refusal directions but damages capability; LEACE preserves capability but removes refusals less thoroughly. Blending the two canceled out each method’s weaknesses, bringing the MMLU cost down to just 0.3 points. The tradeoff: V2 still produced safety-lecture deflections on some queries instead of hard refusals.

**V3** added two more ideas: iterative stacking (always refine the current best model rather than restarting from stock) and a targeted corpus (focused prompts aimed at specific deflection categories, used to find their unique refusal directions without diluting the signal across the whole dataset). V3 takes V2, applies a gentle refinement pass, then a targeted surgery pass, then blends the two results 50/50. The result eliminates soft deflections as well as hard refusals, at a cost of 2.1 MMLU points versus stock.

## What hardware do you need to run it?

The GGUF quantizations span a wide range, which is the main lever for fitting this 27B-parameter model onto consumer hardware:

| Quant | Size | Notes | 
|---|---|---|
| Q8_0 | ~27 GB | Maximum quality, needs a high-VRAM card or multi-GPU setup | 
| Q6_K | ~21 GB | Strong balance of quality and size | 
| Q5_K_M | ~18 GB | Solid all-rounder for 24GB cards | 
| Q4_K_M | ~16 GB | The commonly recommended sweet spot | 
| IQ4_XS | ~14 GB | Compact variant, more experimental | 
| Q3_K_M | ~13 GB | For lower-VRAM setups | 
| Q2_K | ~11 GB | Minimum viable quality, fits on smaller GPUs | 

As a rough rule, you want VRAM at least equal to the GGUF file size, plus some headroom for context length. A 16GB card comfortably handles Q4_K_M with room for a few thousand tokens of context; a 24GB card can run Q5_K_M or Q6_K with more headroom. The full bfloat16 safetensors release is roughly 54GB across 29 shards, which is aimed at multi-GPU or datacenter-class Transformers workflows rather than a single consumer card. Apple Silicon users don’t have an MLX version yet; that’s waiting on upstream `mlx_lm` support for the Qwen3.5 architecture.

## What settings do you actually need?

This is where most people get worse-than-expected output from an otherwise capable model. The documented settings are specific and matter more than usual:

- **Temperature: 0.** Greedy decoding produces the most complete, code-rich outputs on this model. Anything above 0.5 degrades quality noticeably.
- **Repetition penalty: 1.15.** This is described as essential, not optional. Without it, greedy decoding gets stuck looping on imports or boilerplate text. Lowering it to 1.10-1.12 gives tighter, shorter output if 1.15 feels too verbose.
- **Max new tokens: 2048 or more.** Complex code generation and multi-step reasoning need room to finish.
- **System prompt: none.** The model card notes this was A/B tested; adding a system prompt can reintroduce refusal behavior. An empty system prompt performs better.
- **Thinking mode: off (recommended).** V3’s chat template includes a prefill that skips the internal thinking chain, going straight to the answer. Thinking mode can be turned on and doesn’t cause refusals, but off gives more direct output.
- **Sampling parameters (top_p, top_k, min_p): not needed.** Greedy decoding plus the repetition penalty handles this model’s behavior; additional sampling just adds randomness without improving quality.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

For GGUF users specifically, loading the model’s bundled chat template matters. In llama.cpp that means using the `--jinja` flag; in Ollama or LM Studio it means making sure the tool is configured to use the model’s built-in template rather than a generic default, otherwise the thinking-block prefill won’t apply correctly.

## Should you change settings for agentic use?

Yes, and the model card calls this out separately. If you’re running the model inside a coding agent, pentest framework, or any harness that loops through multiple tool calls, the pure-greedy settings that work for single-turn chat can cause the model to get stuck repeating the same tool call or boilerplate output. The recommended adjustments:

