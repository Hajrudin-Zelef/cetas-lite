---
id: collect-240926-mindstudio/mindstudio/how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests-1
title: "how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "quantization", "agent", "agentic", "agi", "benchmark", "benchmarks", "context window", "cost", "gemini", "glm", "gpu"]
source: docs/RAG/clean_en/mindstudio/how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests.md
source_anchor: ""
source_lines: [1, 63]
sha256: 8229090c72bd905e6d276c5ac2275a6fc92fa3b7d132ffedf9e67a87bc07c9b6
---

# how-to-run-deepseek-v4-flash-locally-hardware-quantization-tests

<!-- source: https://www.mindstudio.ai/blog/run-deepseek-v4-flash-locally -->

## What is DeepSeek V4 Flash and why does it matter for local deployment?

DeepSeek V4 Flash is a 284 billion parameter open model that improved sharply on agentic coding tasks through post-training rather than a new architecture. It jumped from roughly 7% to 54% on the DeepSweep benchmark, and on DeepSeek’s own intelligence-versus-cost index it now sits close to Gemini 2.5 Flash while beating GLM 5.2, a model nearly three times its size, on almost every benchmark tested. The reason it matters for anyone building locally is simple: a model this capable at this size is finally practical to self-host without a rack of enterprise GPUs, provided you understand the VRAM math and quantization tradeoffs involved.

## TL;DR

- **DeepSeek V4 Flash gained its biggest jump on agentic coding benchmarks** , going from about 7% to 54% on DeepSweep, largely through post-training rather than architectural changes.
- **Benchmark comparisons aren’t apples-to-apples** because DeepSeek ran its own model under its own optimized harness, and a mismatched harness can swing scores by 3x, as shown in OpenAI’s Arc-AGI harness comparison.
- **Full-precision local deployment needs serious VRAM** : about 168 GB for 4-bit quantization and around 110 GB for 3-bit, which typically means pairing two DGX Spark units or an equivalent multi-GPU setup.
- **A custom engine called Dwarf Star** uses SSD offloading, KV cache manipulation, and mixed precision (some weights at 2-bit, others kept at higher precision) to run comparable models on a single 128 GB system with memory left over for context.
- **Real-world testing in Open Code showed around 25 tokens per second** on a dual DGX Spark cluster, with coherent, less verbose chain-of-thought output compared to earlier DeepSeek releases.
- **The model is text-only** , so anyone needing vision or audio input has to pair it with a separate multimodal model rather than expecting native support.
- **API pricing remains aggressive** , at roughly 2 cents per million input tokens and closer to 30 cents per million output tokens, making the cloud option hard to beat even before considering local hosting.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

## How good are DeepSeek V4 Flash’s benchmark results, really?

The headline numbers are strong. DeepSeek V4 Flash outperforms its own predecessor, V4 Flash preview, on a wide range of tasks and edges out GLM 5.2 despite being a fraction of its size. On the Artificial Intelligence Index, a broader intelligence-versus-cost measure, it lands right next to Gemini 2.5 Flash, positioning it at what’s being called the Pareto frontier of cost and efficiency: you get frontier-adjacent performance without frontier-level pricing.

But the benchmarks need a caveat. DeepSeek tested its agentic coding results using its own harness, the scaffolding of tools, prompts, and orchestration logic that wraps around the raw model during evaluation. Competing models may not have had access to that same harness, which means part of the score improvement reflects harness quality, not just model quality. This isn’t unique to DeepSeek. A recent OpenAI study on Arc-AGI showed that switching from the official benchmark harness to an optimized one nearly tripled the score, from about 13% to 40%. Any time you see a big benchmark jump, the harness deserves as much scrutiny as the model.

The most telling numbers might be DeepSeek’s internal benchmarks, DS-Bench Full Stack and DS-Bench Hard, which aren’t public but are less likely to be contaminated by training data. On those, V4 Flash still trails Anthropic’s Opus 4.5, but the gap has narrowed substantially compared to earlier versions.

## What hardware do you actually need to run it locally?

The VRAM requirement scales directly with quantization level:

- **4-bit precision** : roughly 168 GB of VRAM for model loading alone.
- **3-bit precision** : roughly 110 GB of VRAM.

Neither number includes the extra memory needed for usable context length, so real-world requirements run higher if you want to use anything close to the model’s full context window.

In practice, this means pairing hardware. A single NVIDIA DGX Spark offers around 115 GB of usable VRAM, so a dual DGX Spark cluster clears the 4-bit threshold with some room to spare. Alternatively, any system with more than 108 GB of VRAM can handle the 3-bit configuration, though at some cost to output quality.

Testing on a two-DGX-Spark cluster produced roughly 25 to 30 tokens per second, which is workable for interactive agentic coding sessions, not blazing but far from unusable.

## Is quantization the only way to make this fit on smaller hardware?

Not entirely. A tool called Dwarf Star, built specifically for DeepSeek-style architectures and recently updated to support GLM 5.2 as well, takes a different approach. Instead of quantizing everything uniformly, it mixes precision levels: some weights stay at higher precision (referred to as NVFP formats) while others drop to 2-bit, combined with SSD offloading and KV cache manipulation to keep memory pressure manageable.

The practical result is that a system with 128 GB of VRAM can run a model in this class at reasonable speed while still leaving meaningful memory available for KV cache, the working memory an LLM uses to track conversation and context during generation. Because DeepSeek V4 Flash is a weights update on the same underlying architecture as previous versions, tools built for earlier DeepSeek models should carry over without requiring a new engine.

## How does it perform on real agentic coding tasks?

In hands-on testing using the Open Code harness, DeepSeek V4 Flash was tasked with building a small encyclopedia site (a Pokémon reference page) and a separate space station tracker that pulled from a live API and rendered a UI around it. Both tasks ran without major errors, with the model maintaining a visible to-do list of completed and pending actions throughout the session, a useful signal for anyone debugging agent behavior mid-run.

The chain-of-thought output was notably more structured and less verbose than earlier DeepSeek releases, which historically tended to ramble. Token usage for the Pokémon site landed around 19,000 tokens, a small fraction of the model’s supported one-million-token context window. The rendered UI for the space station tracker looked polished, complete with a starfield background, though the location data it pulled was inaccurate (showing North America only), a reminder that even solid agentic output still needs a human check on data correctness.

## Is DeepSeek V4 Flash worth running locally?

For teams that already have or can justify multi-GPU hardware like dual DGX Spark units, yes, especially if data residency, latency, or long-term API costs are concerns. The token-per-second speed is workable for real coding sessions, and the price-to-performance ratio on DeepSeek’s own index is difficult for competitors to match, whether you’re running it locally or through the API.

The tradeoffs are real, though. The model is text-only with no native vision or audio support, so multimodal workflows need a separate model. VRAM requirements are steep enough that casual single-GPU setups won’t cut it without tools like Dwarf Star to compress the footprint. And the benchmark gains, while genuine, are partly a function of an optimized harness rather than the raw model alone, so expectations should be calibrated against your own harness and workflow, not the published numbers in isolation.

## Frequently Asked Questions

### How much VRAM does DeepSeek V4 Flash need to run locally?

