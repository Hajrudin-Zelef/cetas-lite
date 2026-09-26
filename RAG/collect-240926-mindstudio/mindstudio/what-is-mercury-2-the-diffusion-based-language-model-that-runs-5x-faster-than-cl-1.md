---
id: collect-240926-mindstudio/mindstudio/what-is-mercury-2-the-diffusion-based-language-model-that-runs-5x-faster-than-cl-1
title: "what-is-mercury-2-the-diffusion-based-language-model-that-runs-5x-faster-than-cl"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["diffusion", "agent", "agents", "benchmark", "benchmarks", "claude", "compute", "gpus", "inference", "latency", "llama", "reasoning"]
source: docs/RAG/clean_en/mindstudio/what-is-mercury-2-the-diffusion-based-language-model-that-runs-5x-faster-than-cl.md
source_anchor: ""
source_lines: [1, 97]
sha256: 81377747ad2791ac00b5e3164d192854a02ec20d46e717ab0c87fcc0a8b3c499
---

# what-is-mercury-2-the-diffusion-based-language-model-that-runs-5x-faster-than-cl

<!-- source: https://www.mindstudio.ai/blog/mercury-2-diffusion-based-language-model-5x-faster -->

## A Different Approach to Language Generation

Most language models generate text one token at a time, left to right, waiting for each word before producing the next. It works, but it creates a hard ceiling on speed. Mercury 2 from Inception Labs throws out that approach entirely — and the result is a model that produces text in parallel, at speeds that benchmark around five times faster than Claude Haiku.

That’s a meaningful difference. Mercury 2 is a diffusion-based language model, a category that applies ideas from image generation directly to text. If you’ve used Stable Diffusion or Midjourney, you’ve seen diffusion models at work in images. Mercury 2 is what happens when that same core mechanism gets applied to language — and it changes what’s possible for latency-sensitive applications.

This article breaks down what Mercury 2 is, how its architecture actually works, where it outperforms traditional autoregressive models, and where the tradeoffs still exist.

## What Is Inception Labs, and Why Build a Diffusion LLM?

Inception Labs is the company behind Mercury. Founded by researchers with roots at Stanford and elsewhere in the machine learning community, the company was built around a specific thesis: autoregressive generation is a bottleneck, and the standard token-by-token architecture has fundamental speed limits that can’t be engineered away.

Their answer was to build a family of large language models from scratch using diffusion as the core generation mechanism — not as a bolt-on, but as the foundational architecture.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Mercury is that model family. Mercury 2 is their second-generation release, building on the original Mercury Coder models with improved quality, broader capability, and better benchmark performance. The family currently includes coding-focused variants and general-purpose models at different parameter scales.

The core argument Inception Labs makes is straightforward: if you need high throughput — many requests, low latency, fast output — diffusion-based generation offers structural advantages that autoregressive models simply can’t match at equivalent quality levels.

## How Diffusion Models Work for Text

To understand Mercury 2, it helps to understand what diffusion actually means in the context of language.

### The Image Diffusion Background

In image generation, diffusion works through a two-phase process. During training, noise is progressively added to images until they become pure noise. The model learns to reverse that process — given a noisy image, predict the less noisy version. At inference, you start from random noise and run the denoising process repeatedly until a coherent image emerges.

The key feature: the model processes the entire image at once in each denoising step. It doesn’t generate pixel by pixel from left to right.

### Masked Diffusion for Language

Text is discrete, not continuous like pixel values, so you can’t add Gaussian noise to a sentence the same way you’d corrupt an image. Instead, diffusion language models typically use **masked diffusion** — a process where tokens are randomly masked (replaced with a special `[MASK]` token) during training, and the model learns to predict all masked positions simultaneously.

At inference time, Mercury 2 starts with a sequence that’s entirely masked and iteratively refines it. In each pass, some tokens get “revealed” — committed to specific values — while others remain uncertain and continue to be refined. After a set number of denoising steps, the full sequence is complete.

The critical difference from autoregressive generation: Mercury 2 can predict multiple tokens in parallel in each step. It isn’t blocked waiting for token N before it can predict token N+1.

### Why This Enables Higher Throughput

Autoregressive models have an inherent sequential dependency. Claude, GPT-4, and Llama all generate one token, then use that token as part of the context to generate the next. You can’t parallelize this at the token level — it’s a serial chain by design.

Diffusion models break that chain. Because they refine the entire output simultaneously, modern hardware (GPUs are built for parallel computation) can work much more efficiently. The result is significantly higher tokens-per-second throughput for the same amount of compute.

## Mercury 2 Architecture: What’s New in the Second Generation

Mercury 2 improves on the original Mercury Coder release in several ways.

### Improved Denoising Steps

One of the tuning challenges with diffusion language models is determining how many denoising steps to use. Too few and quality suffers — not enough refinement passes. Too many and you lose the speed advantage.

Mercury 2 uses a more efficient denoising schedule that achieves high-quality outputs in fewer iterations than the first generation. This is part of how it maintains competitive quality while keeping throughput high.

### Better Calibration on Natural Language Tasks

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

The original Mercury models were primarily positioned as coding models. Mercury 2 expands capability across general instruction following, reasoning tasks, and natural language generation — not just code completion. The quality gap between Mercury 2 and leading autoregressive models narrows on these broader task categories compared to the first generation.

### Scalable Architecture

Mercury 2 is available at multiple scales, allowing teams to choose between faster/lighter variants and larger models with higher reasoning quality. The architecture was designed to scale consistently, so the diffusion approach holds its speed advantages even at larger parameter counts.

## Speed Benchmarks: Mercury 2 vs. Claude Haiku and Others

The headline claim is that Mercury 2 runs approximately five times faster than Claude Haiku in terms of throughput. This needs some context.

### What “5x Faster” Actually Means

The speed comparison is measured in **output tokens per second** — how many tokens the model can generate per second at a given level of compute. This is the throughput metric that matters most for production applications like code completion, document generation, or real-time chat.

Claude Haiku is itself one of Anthropic’s faster models, optimized for speed within their lineup. Comparing to Haiku rather than Claude 3.5 Sonnet or Opus is a meaningful choice — Haiku already represents a speed-quality tradeoff optimized toward the speed end.

Mercury 2 claiming 5x over Haiku on tokens-per-second puts it in a genuinely different performance category for raw generation speed.

### The Latency vs. Throughput Distinction

It’s worth being precise: **latency** (time to first token) and **throughput** (tokens per second once generation starts) are different metrics.

Autoregressive models can have low time-to-first-token because they start generating immediately. Diffusion models need to run at least one full denoising pass before producing any output, which can mean slightly higher latency to first token in some configurations.

Where Mercury 2 wins decisively is sustained throughput — generating long outputs quickly. For applications that need to produce hundreds or thousands of tokens (code generation, long-form drafts, document processing), the throughput advantage is the number that matters.

### Quality Benchmarks

Speed alone doesn’t matter if the output quality isn’t there. Inception Labs has published benchmark results showing Mercury 2 competitive with models like Claude Haiku and similar-scale autoregressive models on standard coding and reasoning benchmarks.

