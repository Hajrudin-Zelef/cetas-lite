---
id: collect-mindstudio/mindstudio/local-llm-benchmarks-ryzen-ai-halo
title: "Local LLM Speed Test: GPT-OSS, Qwen3.6 and Hermes on 128GB Unified Memory"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Alibaba", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["memory", "qwen", "agent", "agents", "amd", "benchmarks", "fine-tuning", "gpu", "gpus", "inference", "latency", "llama"]
source: docs/RAG/Collect RAG/02_mindstudio/local-llm-benchmarks-ryzen-ai-halo.md
source_anchor: ""
source_lines: [1, 51]
sha256: a2f6f3990c2ca6c1d32903161c90faf8a912dd16a7f4d1434d8fbeb5bf624f1f
---

# Local LLM Speed Test: GPT-OSS, Qwen3.6 and Hermes on 128GB Unified Memory

## Metadata

- **Source** : https://www.mindstudio.ai/blog/local-llm-benchmarks-ryzen-ai-halo
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article reports real token-per-second benchmarks for local LLMs on a **128GB unified-memory system** — AMD's **Ryzen AI Max Plus 395**, tested via **LM Studio** with the **ROCm Llama.cpp** runtime. Results: a dense **9B model** with a speculative drafter ran at just under **40 tokens per second**; a **35B Qwen3.6 MoE** model with a **3B active parameter** count and drafter exceeded **50 tokens per second**; and the full **GPT-OSS 120B** MoE model landed around **30 tokens per second**.

**Hardware context:** the chip has 16 CPU cores, a Radeon 8060S integrated GPU (~60 teraflops FP16), and a ~50 TOPS NPU. The standout spec is 128GB of unified LPDDR5X shared between CPU and GPU, with the split configurable; in the tested setup **75% of the pool was allocated to the GPU**. This differs fundamentally from a discrete GPU: a 32GB card caps out fast, offloading layers to system RAM and collapsing token speeds; unified memory avoids that cliff because there's no separate smaller pool. The tradeoff is that integrated graphics bandwidth typically doesn't match dedicated VRAM, so small models may still run faster on a discrete card. Most inference stacks don't yet use the NPU, so the 50 TOPS largely sits idle during standard LLM inference.

**Why Qwen3.6 beat GPT-OSS 120B:** in MoE architectures, total and active parameter counts differ, and speed correlates much more closely with **active parameters** than total size. A 35B-total Qwen3.6 with only 3B active can outrun a 120B model per token. This matters for agent workloads: GPT-OSS 120B's ~30 tok/s is fine for chat/RAG where a user reads along, but agent loops that chain many calls benefit from faster, smaller-active models like the Qwen3.6 variant (>50 tok/s).

**Speculative decoding:** both the 9B dense and Qwen3.6 models used a drafter (a smaller model predicting likely next tokens, verified in batches by the main model) — often called multi-token prediction or speculative decoding, increasingly shipped by model creators (recent Gemma and Qwen3.5/3.6 releases include drafter variants). Token throughput increases without a quality tradeoff since the larger model verifies every token. On non-discrete-GPU hardware, a drafter is one of the most reliable ways to close the speed gap.

**Beyond text:** the same system ran ComfyUI locally for image/video generation (preinstalled models plus downloads like the Qwen image model), with image renders around **19–20 seconds**, fast enough to iterate dozens of prompt variations per session. ControlNet workflows (edge detection, pose estimation) were functional. The system also supports fine-tuning workflows (conceptually tested with **Unsloth**) and custom GPU kernel development, extending the unified-memory advantage into training and experimentation.

**FAQ highlights:** 30–50+ tok/s feels responsive for chat/RAG; GPT-OSS 120B at ~30 tok/s is usable for conversation while the Qwen3.6 MoE above 50 tok/s suits latency-sensitive agents. Unified memory doesn't fully replace a discrete GPU (dedicated VRAM offers faster bandwidth for models that fit), but it lets larger MoE models load and run at all.

## Key points

- Benchmarks on Ryzen AI Max Plus 395 (128GB unified): 9B dense ~40 tok/s; Qwen3.6 35B MoE (3B active) >50 tok/s; GPT-OSS 120B ~30 tok/s.
- Speed correlates with active parameters in MoE, not total size — a 3B-active model can outrun a 120B one.
- Unified memory avoids the VRAM offload cliff of discrete GPUs but has lower raw bandwidth.
- Speculative decoding (drafters/MTP) meaningfully boosts throughput without quality loss; shipped by Gemma and Qwen3.5/3.6 families.
- The 50 TOPS NPU is mostly idle for now due to software support gaps.
- Same hardware runs ComfyUI image/video generation (~19–20s per image) and ControlNet workflows.
- GPT-OSS 120B ~30 tok/s suits chat/RAG; faster models suit agent loops.

## Technical data / figures

| Model | Type | Active params | Speed |
|---|---|---|---|
| 9B (Hermes/Nemotron-class) | dense + drafter | 9B | ~40 tok/s |
| Qwen3.6 35B | MoE + drafter | ~3B | >50 tok/s |
| GPT-OSS 120B | MoE | larger active share | ~30 tok/s |

- Hardware: Ryzen AI Max Plus 395; 128GB LPDDR5X unified; 75% allocated to GPU; Radeon 8060S ~60 TFLOPS FP16; 50 TOPS NPU (idle).
- Runtime: LM Studio + ROCm Llama.cpp.
- Image render: ~19–20s (ComfyUI); ControlNet (edge, pose) functional.
- Fine-tuning: Unsloth (conceptually tested).

## Why this source matters for the RAG

It provides a direct, quantitative speed comparison of different model classes (dense vs MoE vs large MoE) on unified-memory hardware, which is essential for model selection in local agent and RAG workloads. It also demonstrates that active-parameter count drives generation speed, and that unified memory enables model sizes discrete GPUs cannot handle.
