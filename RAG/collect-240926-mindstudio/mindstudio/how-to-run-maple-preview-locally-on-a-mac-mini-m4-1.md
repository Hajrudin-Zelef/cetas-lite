---
id: collect-240926-mindstudio/mindstudio/how-to-run-maple-preview-locally-on-a-mac-mini-m4-1
title: "how-to-run-maple-preview-locally-on-a-mac-mini-m4"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple", "Hugging Face", "Nvidia"]
dates: []
keywords: ["agent", "agentic", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "fine-tuning", "gpu", "gpus", "inference"]
source: docs/RAG/clean_en/mindstudio/how-to-run-maple-preview-locally-on-a-mac-mini-m4.md
source_anchor: ""
source_lines: [1, 62]
sha256: c781c718232c8ea86393c1efe69123f11b2fe66462f5a166aab0a6e47bf3dbf3
---

# how-to-run-maple-preview-locally-on-a-mac-mini-m4

<!-- source: https://www.mindstudio.ai/blog/run-maple-preview-locally-mac -->

## What is Maple-Preview and why does it run so well on a Mac mini?

Maple-Preview is an open-source 20B-A1B ternary-weight reasoning model from DeepGrove, released under the MIT license. The “20B-A1B” naming means it has 20 billion total parameters but only about 1 billion active per token, thanks to a mixture-of-experts design with 256 experts and 8 active at a time. The ternary weights (values restricted to roughly -1, 0, and 1) shrink the checkpoint to just 5.31 GB, small enough to load comfortably in the unified memory of an entry-level Mac mini M4. DeepGrove reports 218 tokens per second on that hardware, which the model card describes as 5 to 16 times faster than comparably efficient models like Gemma, Qwen3.5, and gpt-oss.

## TL;DR

- **Maple-Preview** is a 20B-A1B ternary reasoning model with a 5.31 GB checkpoint, small enough to run entirely in memory on a base Mac mini M4.
- **Speed on Apple Silicon** is reported at 218 tokens per second, which DeepGrove positions well ahead of similarly efficient dense and MoE models.
- **The architecture** uses 24 layers, 256 experts with 8 active per token, and a 3:1 mix of sliding-window and global attention to keep memory and compute costs low.
- **Context length** reaches 131,072 tokens, enough for long documents or extended multi-step reasoning chains without truncation.
- **The official Transformers implementation targets CUDA** (it depends on Triton and FlashAttention), so the 218 tok/s Apple Silicon number comes from a separate on-device runtime, not the reference code path.
- **This is a preview release** focused on raw reasoning benchmarks like AIME and GPQA-D; it has had minimal agentic or tool-use training, so it may underperform on tasks that need multi-step tool calling.
- **The model is free to use commercially** under the MIT license, with no restrictions on redistribution or fine-tuning.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

## What hardware do you need to run Maple-Preview locally?

The headline number is the checkpoint size: 5.31 GB. That is small enough to fit inside the unified memory of a base Mac mini M4, which starts at 16 GB of shared RAM. Because Apple Silicon uses unified memory rather than a discrete GPU with its own VRAM pool, the entire model can sit in fast, CPU-and-GPU-accessible memory without the swapping or offloading tricks that larger dense models require.

This is the practical appeal of a ternary MoE model at this size: you don’t need a workstation GPU or a cloud instance to get real reasoning performance. A Mac mini M4, one of Apple’s cheaper desktop machines, is enough to hit the benchmarked 218 tokens per second. That number scales differently on other hardware. Machines with more GPU cores or higher memory bandwidth (higher-end M4 Pro or M4 Max configurations) would generally push higher, though DeepGrove’s published figure specifically references the standard Mac mini M4.

## How does the ternary weight format make this possible?

Most LLMs ship in 16-bit or 8-bit floating point formats, or get quantized down to 4-bit integers for local inference. Maple-Preview goes further with ternary weights, where each parameter is restricted to one of three values instead of a wide range of floating-point numbers. This is an extreme form of quantization built into the model from training, not applied afterward.

The payoff is a dramatically smaller memory footprint and cheaper compute per token, since multiplying by -1, 0, or 1 is far simpler than standard floating-point matrix multiplication. Combined with the mixture-of-experts routing (only 8 of 256 experts activate per token, and only about 1 billion of the 20 billion total parameters do any work on a given forward pass), the model achieves a compute profile closer to a much smaller dense model while retaining the capacity that comes from having 20 billion parameters to draw on.

The attention mechanism reinforces this efficiency. Maple-Preview uses a 3:1 ratio of sliding-window attention (with a 512-token window) to global attention across its 24 layers. Sliding-window attention caps the computational cost of attending to long sequences, while the occasional global attention layer preserves the model’s ability to reference distant context when needed. That combination is part of how the model supports a 131,072-token context window without the memory cost scaling linearly with every layer being fully global.

## Is Maple-Preview worth running instead of a larger model?

That depends on what you need it for. DeepGrove’s own benchmark comparisons position Maple-Preview on the Pareto frontier for both memory-to-performance and speed-to-performance, meaning that for its size and speed class, it produces reasoning scores that larger or slower models don’t clearly beat. The model card highlights strength on IMO-level math problems and competitive results on LiveCodeBench v6, AIME 2026, HMMT 2026, and GPQA-D against models like Gemma 4, Qwen3.5, and gpt-oss.

Where it’s less proven is agentic work. DeepGrove is explicit that this preview received “minimal post-training for agentic tasks and only small-scale general reinforcement learning.” If your use case involves tool calling, multi-step agent workflows, or general assistant behavior beyond math and code reasoning, Maple-Preview may lag behind models specifically tuned for those tasks. It’s a reasoning-first release, not a general-purpose daily driver, at least in this preview form. DeepGrove has stated a full release with broader training is planned.

For anyone building local-first tools focused on math, logic, or code reasoning, and who wants that running on modest hardware rather than a cloud GPU, Maple-Preview is a legitimate option to test. For agentic pipelines or broad chat assistant duty, it’s worth waiting for the full release or pairing it with a model built for that purpose.

## How do you actually get it running on a Mac?

The official model card ships a Transformers-based implementation, but that reference code depends on Triton and FlashAttention, both of which are designed for CUDA GPUs. That means the out-of-the-box Hugging Face Transformers path is built for Nvidia hardware, not Apple Silicon.

The 218 tokens per second figure DeepGrove reports for the Mac mini M4 comes from a separate on-device runtime rather than the stock Transformers implementation. In practice, this means Mac users looking to reproduce that speed need an inference engine built for Apple’s Metal/MLX ecosystem rather than expecting the default Hugging Face code path to run natively and fast on a Mac. Anyone testing this locally should check for community or official MLX-compatible builds of the checkpoint before assuming the standard `transformers` library will deliver comparable performance on-device.

Once you have a compatible runtime, the workflow is what you’d expect from any local LLM: download the 5.31 GB checkpoint, point your runtime at it, and load it into memory. Given the small file size, download and load times are fast compared to dense 20B+ models that ship at 30 to 40+ GB in half precision.

## Frequently Asked Questions

### What does “20B-A1B” mean for Maple-Preview?

It means the model has 20 billion total parameters spread across a mixture-of-experts architecture, but only about 1 billion parameters are active for any single token, since just 8 of 256 experts fire per forward pass. This keeps compute cost low while retaining a large total parameter pool.

### How big is the Maple-Preview download?

The checkpoint is 5.31 GB, made possible by its ternary weight format, which stores each parameter as one of three discrete values instead of a wide-range floating-point number.

### Can I run Maple-Preview with standard Hugging Face Transformers on a Mac?

