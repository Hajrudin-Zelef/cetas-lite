---
id: collect-240926-mindstudio/mindstudio/what-is-speculative-decoding-how-draft-models-speed-up-llms
title: "what-is-speculative-decoding-how-draft-models-speed-up-llms"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agents", "benchmark", "compute", "cost", "gpu", "inference", "kv cache", "latency", "llama", "llama.cpp", "memory"]
source: docs/RAG/clean_en/mindstudio/what-is-speculative-decoding-how-draft-models-speed-up-llms.md
source_anchor: ""
source_lines: [1, 97]
sha256: f7e2dad0b62a21a33426ca9177e9e875c4fe4ed6f45ea114ac46f8f2519114e0
---

# what-is-speculative-decoding-how-draft-models-speed-up-llms

<!-- source: https://www.mindstudio.ai/blog/what-is-speculative-decoding -->

## What is speculative decoding?

Speculative decoding is an inference technique that speeds up text generation from large language models without changing their output. Instead of the big model producing one token per expensive forward pass, a small “draft” model guesses several tokens ahead, and the big model verifies all those guesses in a single pass. Correct guesses are accepted for free. Wrong guesses are discarded, but everything generated before the mistake is kept. The result is more tokens produced per unit of compute, with identical text to standard decoding.

## TL;DR

- **Speculative decoding** splits generation into two steps: a small draft model proposes tokens, and the large model checks them in one verification pass instead of generating each token itself.
- **Accepted guesses are free tokens** , since a single forward pass through the big model can confirm multiple draft tokens at once rather than producing just one.
- **Wrong guesses cost almost nothing** , because the system simply throws away tokens after the first mistake and keeps everything verified before it.
- **Output quality stays identical** to normal autoregressive decoding since the large model still checks and controls every token that gets accepted.
- **DFlash 2 pushes the idea further** by having the draft model predict all positions in parallel instead of one at a time, and by keeping a wide pool of candidate tokens per slot instead of just one guess.
- **A live benchmark on Qwen3 27B** showed throughput roughly double, from around 29 tokens per second without speculative decoding to about 59 tokens per second with DFlash 2 enabled, on the same model, prompts, and hardware.
- **The technique is already supported in inference engines** like SGLang, and is also available in projects such as vLLM and llama.cpp, according to the video creator.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

## How does speculative decoding actually work?

A standard large language model generates text autoregressively: one token in, one forward pass through billions of parameters, one token out. Repeat. That forward pass is expensive, and running it once per token is the main reason large models feel slow compared to their size.

Speculative decoding attacks that bottleneck by separating “guessing” from “verifying.” A small, cheap draft model looks at the current context and proposes several tokens ahead, sometimes four, eight, or more depending on configuration. Then the large model runs a single forward pass over that whole proposed sequence and checks which of the draft tokens it would have actually generated itself.

Two outcomes follow:

- If the draft model’s guesses match what the big model would have produced, those tokens are accepted as-is. That’s multiple tokens generated for the cost of one verification pass.
- If a guess is wrong at some position, everything after that point is discarded, but everything before it is kept and counted as valid output.

Because verification happens in parallel across the whole draft sequence rather than token by token, the big model’s expensive compute is amortized across several tokens instead of just one. The math only wins if the draft model is right often enough to matter, but since bad guesses cost almost nothing (they’re simply dropped), there’s little downside to trying.

## Why doesn’t this hurt output quality?

The key property of speculative decoding is that the large model always has final say. It isn’t trusting the draft model’s tokens blindly. It’s checking, in the same way it normally would generate each token, whether the proposed token matches what it would have produced under its own decoding process. Only tokens the big model would have generated anyway get accepted.

This means speculative decoding is a pure speed optimization, not an approximation. The output text is the same as running the large model alone in standard autoregressive mode. It doesn’t shrink the model, change its weights, or quantize its precision. It changes how many forward passes are needed to produce a given amount of text, not what that text is.

## What is DFlash 2 and how does it improve on basic speculative decoding?

DFlash 2 is a project that applies speculative decoding to Qwen3’s 27 billion parameter model, built by the open source community rather than the original model developers. It illustrates a pattern common in open source AI: a lab releases a base model, and outside teams then quantize it, distill it, or, in this case, build a faster inference layer on top of it.

Ordinary speculative decoding still has the draft model guessing tokens somewhat sequentially, one position informing the next. DFlash 2 changes that by having the draft model predict all positions in a single parallel pass, rather than guessing one token, then the next, then the next.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

It also changes what “one guess” means. Instead of proposing a single top token per position, DFlash 2 keeps a pool of 16 candidate tokens at every slot. According to the project’s approach as described in the video, even when the single top-ranked guess is wrong, the correct token shows up somewhere in that top-16 pool the large majority of the time (cited as around 99% at the first position). A lightweight “path selector” then scores neighboring pairs of candidates across positions to find the most coherent sequence through them, and a small convolutional component lets each position take into account the token before it, without breaking the parallel structure of the draft step.

The net effect described is one additional accepted token per verification pass compared to simpler speculative decoding setups, at a claimed added latency cost of around 1%.

## How much faster is it in practice?

In a live local benchmark using Qwen3 27B on a single Nvidia A100 GPU, running the same five prompts through SGLang, plain autoregressive decoding (no speculative decoding) produced an average of about 29 tokens per second, consistent across prompts. Enabling DFlash 2 speculative decoding on the identical model, prompts, and hardware roughly doubled that, to about 59 tokens per second on average. One prompt reportedly hit around 72 tokens per second, attributed to a long run of correct draft guesses that the verifier accepted in bulk.

VRAM usage in that setup was just over 77 GB for the full model plus KV cache, with the draft model itself adding comparatively little extra memory, since draft models are much smaller than the model they’re guessing for.

These are single-benchmark figures from one hardware setup and one model, not universal numbers. Speedup from speculative decoding depends heavily on how well the draft model’s outputs align with the base model’s, how many tokens are drafted per pass, and the nature of the prompts (repetitive or structured text tends to produce more accepted guesses than highly unpredictable text).

## Is speculative decoding worth using?

For anyone running large models locally or serving them at scale, speculative decoding is close to a free win when a compatible draft model exists. It requires no retraining of the base model, no quantization tradeoffs, and produces identical output to standard decoding. The main costs are the extra memory and compute needed to run the draft model alongside the main one, and the engineering work of finding or building a draft model that predicts the base model’s behavior well.

It’s already supported in mainstream inference engines. The video’s creator noted that beyond SGLang, speculative decoding support also extends to projects like vLLM and llama.cpp, which suggests it’s becoming a standard feature rather than a niche research trick.

The main limitation is that gains vary by workload. Speculative decoding shines when the draft model can reliably predict long stretches of the big model’s output, which tends to happen with more predictable or structured generations. Highly creative or unpredictable text may see smaller speedups, since the draft model’s guesses will be rejected more often.

## Frequently Asked Questions

### Does speculative decoding change the model’s output?

No. The large model still verifies and effectively generates every accepted token itself. Speculative decoding only changes how many forward passes are needed to produce that output, not the text produced.

### What is a draft model?

A draft model is a smaller, faster model used to propose likely next tokens ahead of time. It doesn’t need to be as accurate as the main model, since its guesses get checked and can be discarded cheaply if wrong.

### How is DFlash 2 different from standard speculative decoding?

## One coffee. One working app.

You bring the idea. Remy manages the project.

Standard speculative decoding drafts tokens somewhat sequentially and typically proposes one top candidate per position. DFlash 2 predicts all draft positions in parallel and keeps a pool of multiple candidate tokens per position, then uses a path selection step to pick the most coherent sequence through them.

### Does speculative decoding require more GPU memory?

Yes, some. Running a draft model alongside the base model adds VRAM overhead, though draft models are typically much smaller than the base model, so the added memory footprint is relatively modest compared to the base model’s own requirements.

### Is speculative decoding available in popular inference tools?

Yes. It’s supported in engines such as SGLang, and according to the video’s creator, similar support exists in vLLM and llama.cpp, making it accessible without custom infrastructure.
