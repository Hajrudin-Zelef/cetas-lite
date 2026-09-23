---
id: collect-mindstudio/mindstudio/run-maple-preview-locally-mac
title: "How to Run Maple-Preview Locally on a Mac Mini M4"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Apple", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "attention", "benchmarks", "compute", "consumer", "cost", "fine-tuning", "inference", "inference engine", "license", "memory"]
source: docs/RAG/Collect RAG/02_mindstudio/run-maple-preview-locally-mac.md
source_anchor: ""
source_lines: [1, 50]
sha256: 6f6146fcf6bc126e7127403a91796949aa3797524fb42465cad706d236a7060c
---

# How to Run Maple-Preview Locally on a Mac Mini M4

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-maple-preview-locally-mac
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article is a guide to running **Maple-Preview**, an open-source **20B-A1B ternary-weight reasoning model** from **DeepGrove**, on a Mac mini M4. The "20B-A1B" naming means 20 billion total parameters but only about **1 billion active per token**, thanks to a mixture-of-experts (MoE) design with **256 experts** and **8 active at a time**. Ternary weights (parameters restricted to roughly -1, 0, and 1) shrink the checkpoint to just **5.31 GB**, small enough to load comfortably in the 16 GB unified memory of an entry-level Mac mini M4. DeepGrove reports **218 tokens per second** on that hardware, which the model card describes as **5 to 16 times faster** than comparably efficient models like Gemma, Qwen3.5, and gpt-oss.

The architecture uses **24 layers**, **256 experts with 8 active per token**, and a **3:1 mix of sliding-window (512-token window) and global attention** to keep memory and compute costs low. Context length reaches **131,072 tokens**, enough for long documents or extended multi-step reasoning chains without truncation. Ternary weights are an extreme form of quantization built into the model during training, not applied afterward; multiplying by -1, 0, or 1 is far simpler than floating-point matrix multiplication, dramatically reducing memory footprint and per-token compute cost. Combined with MoE routing, the model has a compute profile closer to a much smaller dense model while retaining 20B parameters of capacity.

An important caveat: the official Transformers implementation **targets CUDA** (it depends on **Triton and FlashAttention**), so the 218 tok/s Apple Silicon figure comes from a **separate on-device runtime**, not the reference code path. Mac users need an Apple Silicon/Metal/MLX-compatible inference engine to reproduce that speed; the default Hugging Face transformers path will not run natively fast on a Mac.

DeepGrove positions Maple-Preview on the **Pareto frontier** for memory-to-performance and speed-to-performance. The model card highlights strength on IMO-level math problems and competitive results on **LiveCodeBench v6, AIME 2026, HMMT 2026, and GPQA-D** against models like Gemma 4, Qwen3.5, and gpt-oss. However, this preview received "minimal post-training for agentic tasks and only small-scale general reinforcement learning," so it may underperform on multi-step tool calling and agent workflows. It is a reasoning-first release (math, logic, code), not a general-purpose daily driver; DeepGrove has stated a full release with broader training is planned. The model is free for commercial use under the **MIT license**.

## Key points

- Maple-Preview is a 20B-A1B ternary reasoning model (256 experts, 8 active per token) with a 5.31 GB checkpoint.
- Reported speed on Apple Silicon is 218 tokens per second, positioned 5–16x faster than comparably efficient models.
- Architecture: 24 layers, 3:1 sliding-window (512-token) to global attention, 131,072-token context.
- The official Transformers path targets CUDA (Triton + FlashAttention); the Mac speed comes from a separate on-device runtime.
- Reasoning-first preview: strong on AIME, HMMT, GPQA-D, LiveCodeBench v6; weak on agentic/tool-use tasks.
- MIT license, free for commercial use, fine-tuning, and redistribution.

## Technical data / figures

| Spec | Value |
|---|---|
| Model | Maple-Preview (DeepGrove) |
| Parameters | 20B total, ~1B active (20B-A1B) |
| Experts | 256 total, 8 active per token |
| Weight format | ternary (values ~ -1, 0, 1) |
| Checkpoint size | 5.31 GB |
| Layers | 24 (3:1 sliding-window : global attention) |
| Sliding window | 512 tokens |
| Context length | 131,072 tokens |
| Reference hardware | Mac mini M4 (16 GB unified memory) |
| Throughput | 218 tokens/s (Mac mini M4, on-device runtime) |
| Benchmarks | AIME 2026, HMMT 2026, GPQA-D, LiveCodeBench v6, IMO-level math |
| License | MIT |

## Why this source matters for the RAG

It documents a frontier of efficient local inference — native ternary weights plus MoE — that makes a 20B-class reasoning model run entirely in a consumer Mac's unified memory at high speed. It also provides a concrete example of the CUDA-vs-Apple-Silicon toolchain split, useful for advising on realistic local deployment paths.
