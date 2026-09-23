---
id: collect-mindstudio/mindstudio/what-is-speculative-decoding
title: "What Is Speculative Decoding? How Draft Models Speed Up LLMs"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["speculative decoding", "benchmark", "compute", "cost", "gpu", "inference", "kv cache", "latency", "llama", "llama.cpp", "memory", "nvidia"]
source: docs/RAG/Collect RAG/02_mindstudio/what-is-speculative-decoding.md
source_anchor: ""
source_lines: [1, 54]
sha256: ffcb1297fba56e14dabe8de12cd4cb589efc9dcb95bf49a85f75e5eb2640bc5c
---

# What Is Speculative Decoding? How Draft Models Speed Up LLMs

## Metadata

- **Source** : https://www.mindstudio.ai/blog/what-is-speculative-decoding
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**Speculative decoding** is an inference technique that speeds up text generation from large language models without changing their output. Instead of the big model producing one token per expensive forward pass, a small **"draft" model** guesses several tokens ahead, and the big model verifies all those guesses in a single pass. Correct guesses are accepted for free; wrong guesses are discarded, but everything generated before the mistake is kept. The result is more tokens per unit of compute, with **identical text** to standard decoding.

Standard LLMs generate autoregressively: one token in, one forward pass through billions of parameters, one token out, repeat. That forward pass is expensive, and running it once per token is the main reason large models feel slow. Speculative decoding separates **"guessing" from "verifying."** A small, cheap draft model proposes several tokens ahead — sometimes four, eight, or more depending on configuration. Then the large model runs a **single forward pass** over the whole proposed sequence and checks which draft tokens it would have actually generated itself. Two outcomes: if the guesses match, those tokens are accepted as-is (multiple tokens for the cost of one verification pass); if a guess is wrong at some position, everything after that point is discarded but everything before it is kept as valid output. Because verification is parallel across the whole draft sequence, the big model's compute is amortized across several tokens.

Output quality is unaffected because the large model always has final say. It isn't trusting the draft tokens blindly; it checks, in the same way it would normally generate each token, whether the proposed token matches what it would have produced under its own decoding process. Only tokens the big model would have generated anyway are accepted. This makes speculative decoding a **pure speed optimization, not an approximation** — it doesn't shrink the model, change weights, or quantize precision.

**DFlash 2** is a project applying speculative decoding to Qwen3's **27 billion parameter** model, built by the open-source community. It illustrates a common open-source pattern: a lab releases a base model, and outside teams quantize, distill, or build a faster inference layer on top. Ordinary speculative decoding drafts tokens somewhat sequentially; DFlash 2 has the draft model predict **all positions in a single parallel pass**. It also keeps a **pool of 16 candidate tokens** at every slot instead of a single top guess — the correct token shows up somewhere in that top-16 pool the large majority of the time (cited as around **99% at the first position**). A lightweight **path selector** scores neighboring pairs of candidates across positions to find the most coherent sequence, and a small convolutional component lets each position consider the previous token without breaking the parallel structure. The net effect is one additional accepted token per verification pass compared to simpler setups, at a claimed added latency of around **1%**.

In a live local benchmark on **Qwen3 27B** with a single **Nvidia A100** GPU running **SGLang** and the same five prompts, plain autoregressive decoding produced about **29 tokens/second**; enabling DFlash 2 roughly doubled that to about **59 tokens/second** on average, with one prompt hitting around **72 tokens/second** (attributed to a long run of correct guesses accepted in bulk). VRAM usage was just over **77 GB** for the full model plus KV cache, with the draft model adding comparatively little since draft models are much smaller. These are single-benchmark figures from one setup, not universal; speedup depends heavily on how well the draft aligns with the base model, how many tokens are drafted per pass, and prompt nature (repetitive/structured text produces more accepted guesses).

Speculative decoding is close to a free win when a compatible draft model exists: no retraining, no quantization tradeoffs, identical output. Costs are the extra memory/compute for the draft model and the engineering work of finding a good draft model. It's already supported in **SGLang**, **vLLM**, and **llama.cpp**, making it a standard feature rather than a niche trick. Gains vary by workload: it shines on predictable or structured generations, while highly creative or unpredictable text sees smaller speedups because guesses get rejected more often.

## Key points

- Speculative decoding splits generation: a draft model proposes tokens, the large model verifies them in one pass.
- Accepted guesses are "free" tokens; wrong guesses cost almost nothing since only tokens after the first error are dropped.
- Output is identical to normal autoregressive decoding because the large model controls every accepted token.
- DFlash 2 predicts all draft positions in parallel and keeps a 16-candidate pool per slot.
- A live A100 benchmark on Qwen3 27B showed ~29 tok/s → ~59 tok/s (one prompt ~72 tok/s) with DFlash 2.
- Supported in SGLang, vLLM, and llama.cpp; gains vary by workload and draft-model quality.

## Technical data / figures

| Item | Value |
|---|---|
| Technique | Draft-and-verify speculative decoding |
| Quality impact | None (identical output) |
| DFlash 2 candidate pool | 16 per position |
| Top-16 hit rate (first position) | ~99% |
| Added latency (DFlash 2) | ~1% |
| Benchmark model | Qwen3 27B |
| Benchmark GPU | Single Nvidia A100 |
| Serving engine | SGLang |
| AR throughput | ~29 tok/s |
| DFlash 2 throughput | ~59 tok/s avg (peak ~72) |
| VRAM (model + KV cache) | >77 GB |
| Supported engines | SGLang, vLLM, llama.cpp |

## Why this source matters for the RAG

It provides the foundational conceptual explanation of speculative decoding and a concrete DFlash 2 benchmark, making it an ideal primer for inference-acceleration questions. It grounds the more specific DFlash 2 deployment articles.
