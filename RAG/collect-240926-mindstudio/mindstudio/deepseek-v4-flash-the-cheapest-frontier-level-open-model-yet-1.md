---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet-1
title: "deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Nvidia", "OpenAI", "Z.ai"]
dates: []
keywords: ["deepseek", "agent", "agentic", "agents", "agi", "benchmark", "benchmarks", "cost", "gemini", "glm", "gpu", "inference"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet.md
source_anchor: ""
source_lines: [1, 69]
sha256: 092dbb67c55d4420ca70a83f061147414f1157e8c050132415c41c7526ca14a0
---

# deepseek-v4-flash-the-cheapest-frontier-level-open-model-yet

<!-- source: https://www.mindstudio.ai/blog/deepseek-v4-flash-local-benchmarks -->

## What is DeepSeek V4 Flash and why does it matter?

DeepSeek V4 Flash is a post-trained update to DeepSeek’s existing V4 Flash preview model, not a new architecture or a larger parameter count. At 284 billion parameters, it now scores close to much bigger frontier models on several benchmarks while staying priced near the bottom of the market, around 2 cents per million input tokens and roughly 30 cents per million output tokens. That combination puts it near the front of the cost-versus-intelligence curve that matters most to teams running agentic coding workloads at scale.

## TL;DR

- DeepSeek V4 Flash jumped from 7% to 54% on the DeepSweep agentic coding benchmark compared to the previous V4 Flash preview, a gain driven entirely by post-training rather than a new model.
- The model beats GLM 5.2 on nearly every reported benchmark despite GLM 5.2 being close to three times larger, showing how much performance post-training can extract from a smaller base model.
- Pricing sits near 2 cents per million input tokens and 30 cents per million output tokens, placing it at the Pareto frontier of cost versus intelligence alongside recent releases like Gemini 2.6 Flash.
- Benchmark comparisons are not apples-to-apples because DeepSeek ran its own model under its own optimized harness, and other models likely weren’t tested under that same harness.
- On DeepSeek’s internal DS-Bench Full Stack and DS-Bench Hard benchmarks, which are harder to game since they’re not public, the model still trails Opus 4.8, giving a more honest read on its ceiling.
- The model can run locally on hardware like two networked Nvidia DGX Spark units, hitting around 20 to 30 tokens per second depending on quantization and context load.
- It remains text-only with no native image or vision support, which limits it against multimodal alternatives for certain use cases.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

## How much did the benchmark scores actually improve?

The headline number is the jump on the DeepSweep agentic coding benchmark: from 7% under the previous V4 Flash preview to 54% under the new V4 Flash release. That’s not a small tuning gain, it’s a categorical shift in what the model can reliably do on agent-style coding tasks. DeepSeek also reports that the new model beats GLM 5.2 on almost every benchmark it published, despite GLM 5.2 running at nearly three times the parameter count. That’s the core efficiency story: post-training, without touching the underlying architecture, closed most of the gap with much larger open models and pushed the model into range of proprietary heavyweights like Opus 4.8.

On the Artificial Intelligence Index, a composite measure of task performance, the jump was large enough to put DeepSeek V4 Flash right next to Gemini 2.6 Flash, a model released only days earlier. For a workhorse-tier open model, that’s a notable position to hold.

## Why do the benchmark numbers need a caveat?

The comparisons DeepSeek published for agentic coding tasks were run under DeepSeek’s own harness, the software scaffolding that manages tool calls, retries, context handling, and task execution around the raw model. Other models in the comparison table likely didn’t have access to that same harness, so the results aren’t a clean head-to-head of model capability alone. Harness quality can swing scores dramatically on its own, independent of the underlying model.

A useful reference point here is OpenAI’s own study on the ARC-AGI benchmark, where switching from the standard public harness to an internally optimized harness moved the score from around 13% to 40%, roughly triple. That’s purely a harness effect, with no change to the model. The same dynamic likely explains a meaningful share of DeepSeek V4 Flash’s reported gains: some of the improvement is the model, and some of it is DeepSeek running its own model through infrastructure tuned specifically for it, something competitors weren’t given the chance to do.

This doesn’t mean the published numbers are meaningless. It means they should be read as “best case under an optimized setup” rather than a neutral, harness-independent measure of raw model capability.

## What do DeepSeek’s internal benchmarks show?

More telling than the public leaderboard numbers are DeepSeek’s own internal evaluations, DS-Bench Full Stack and DS-Bench Hard. These aren’t available to the public, which is exactly why companies tend to trust them more: internal benchmarks can’t be contaminated by leaking into training data the way public benchmarks sometimes are, making them a cleaner signal of real progress.

On those two internal benchmarks, V4 Flash showed a real improvement but still lagged behind Opus 4.8. That’s a more grounded picture than the headline DeepSweep jump. The model has clearly gotten better and closed distance with larger frontier systems, but it hasn’t caught up to the top tier on the metrics DeepSeek itself trusts most.

## Is DeepSeek V4 Flash worth using for agentic coding?

## One coffee. One working app.

You bring the idea. Remy manages the project.

For teams building agentic coding tools, cost per task matters as much as raw capability, and this is where V4 Flash looks strongest. On Terminal Bench 2.1, a benchmark widely referenced for agentic coding performance, it scores comparably to some of the best models available at a fraction of their per-token cost. Combined with support for the responses API (usable directly in tools like Codex) and open-source harnesses like Open Code, it’s positioned as a practical executor model: something you’d pair with a larger planning model rather than use as your single frontier system for every task.

The lack of multimodal support is a real limitation. It’s text-only, no image or vision input, which rules it out for workflows that need to reason over screenshots, diagrams, or other visual context. Models built for multimodal use cases fill that gap, generally at some cost to raw coding benchmark scores.

## Can you run DeepSeek V4 Flash locally, and what hardware does it need?

Yes, and this is one of the more notable practical points about the release: it can run entirely on local hardware, not just through an API. Requirements scale with quantization level. Running at 4-bit precision needs around 168 GB of VRAM, which typically means a multi-GPU setup like two networked Nvidia DGX Spark units (each offering roughly 115 GB of usable VRAM). At 3-bit quantization, the requirement drops to around 110 GB of VRAM. Keep in mind these figures cover model loading only, actual usable context length demands significantly more memory on top.

For those without that much hardware, alternative inference engines using techniques like SSD offloading, KV cache manipulation, and mixed-precision weight loading (keeping some weights at higher precision while quantizing others more aggressively) can run the same architecture on systems with around 128 GB of VRAM, leaving more headroom for context. Since V4 Flash uses the same architecture as its predecessor, tooling built for that architecture continues to work with the updated weights.

In practice, local testing on a two-DGX-Spark setup produced speeds in the range of 20 to 30 tokens per second, usable for agentic coding tasks though not instant.

## Frequently Asked Questions

### What changed between V4 Flash preview and the new V4 Flash release?

No architecture or parameter count changed. The improvements come entirely from post-training, refining how the existing model behaves rather than retraining it from scratch or scaling it up.

### Does DeepSeek V4 Flash beat GLM 5.2?

