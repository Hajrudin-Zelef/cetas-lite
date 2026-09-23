---
id: vague2-nerdykings/nerdykings/deepseek-v4-1-flash-kv-cache-437x
title: "DeepSeek V4.1 Flash : Un KV Cache 437 Fois Plus Petit Que La Première Génération"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "DeepSeek"]
dates: ["2026-09-23"]
keywords: ["deepseek", "kv cache", "agents", "attention", "benchmark", "claude", "context window", "cost", "gpu", "inference", "memory", "mixture of experts"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/deepseek-v4-1-flash-kv-cache-437x.md
source_anchor: ""
source_lines: [1, 67]
sha256: c6e18f49aa07d340fd0dca1fcbb644d85a4b058b19c44ac1d1fe5de15474e85f
---

# DeepSeek V4.1 Flash : Un KV Cache 437 Fois Plus Petit Que La Première Génération

## Metadata

- **Source** : https://www.nerdykings.com/blog/deepseek-v4-1-flash-kv-cache-437x.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek has released V4.1 Flash, a surprising model that is extremely fast and, on some tests, rivals high-end models. But the real story is not benchmark scores: DeepSeek drastically reduced the memory needed to handle very long contexts. Its KV cache is about 4x smaller than V4 Flash's and up to 437x smaller per token than the first DeepSeek generation, while supporting contexts of up to 1 million tokens.

On DeepSWE, V4.1 Flash reaches 74.2 points vs 74 for Claude in reported results — impressive for its category. But it doesn't win everywhere: on Terminal-Bench, DeepSeek drops to 31.2 vs 51.8 for Claude. So DeepSeek hasn't built the world's best model, but it has done something more interesting.

The core problem: when chatting with an AI, the model must constantly remember prior information — that's the KV cache. Short conversations are fine, but today's agents receive entire projects (hundreds of files, terminal history, tools, images). V4.1 Flash works with contexts up to 1 million tokens, and the larger the context, the more the KV cache memory becomes a real problem. This is where DeepSeek made a huge leap: ~4x smaller KV cache than V4 Flash and ~437x smaller per token than the first generation in just a few years.

The key innovation is **CSA2** (Compressed Sparse Attention). A Transformer has many layers, each traditionally producing and storing its own information in the KV cache. DeepSeek asked whether every layer really needs to store all this independently — and the answer is no. In V4.1 Flash, only four layers (layers 2, 8, 14, and 20) write a persistent global KV state; all others reuse already-produced information instead of keeping their own version. This combines with a distinctive structure: V4.1 Flash uses 40 layers split into two parts — 20 encoder layers and 20 decoder layers. The encoder first processes the huge prompt and produces a kind of global memory of that context; when the decoder starts generating, it reads this memory directly instead of recreating new global keys/values at every layer. Added compression and quantization techniques produce the enormous result.

On V4 Flash, the cache was about 3,514 bytes per token. With V4.1 Flash, only 890 bytes — nearly four times less. This matters greatly because GPU memory is one of the most precious resources for running these huge models; DeepSeek has been attacking hardware bottlenecks rather than raw size for months (similar efficiency logic to Dual Path on the GPU side or Engram on the memory side).

Important caveat: V4.1 Flash is not a small model. Its backbone has ~552 billion parameters. It's a Mixture of Experts, so not all are used at once — only ~8 billion active parameters per token during prompt processing, then ~16 billion during generation. But downloading the full checkpoint is ~511 GB. Also, V4.1 now has native visual understanding: it can receive text and images (a UI, a screenshot, even a video game menu) and analyze what it sees and generate code to reproduce the behavior — very useful for agents.

The main flaw: it over-thinks. Maximum effort mode can consume up to 2.5x more output tokens, often without benefit. On a writing test, standard mode scores 87.13/100 and maximum mode ~86.75 — essentially the same or slightly worse, while generation takes 73% more time and costs ~60% more. Sometimes thinking more is useless.

The author concludes V4.1 Flash is interesting because DeepSeek targets a precise problem: running gigantic models with gigantic contexts without exploding memory. It remains heavy (552B params, 511 GB), but the trajectory matters — the real race isn't the best benchmark score but who can run agents over millions of context tokens without GPU memory bills exploding. There, DeepSeek has a clear lead worth watching.

## Key points

- V4.1 Flash's KV cache is ~4x smaller than V4 Flash and up to 437x smaller per token than the first DeepSeek generation.
- Supports contexts up to 1 million tokens; ~890 bytes per token vs ~3,514 for V4 Flash.
- **CSA2** (Compressed Sparse Attention): only 4 of 40 layers (2, 8, 14, 20) write a persistent global KV state.
- Architecture: 40 layers split into 20 encoder + 20 decoder layers; encoder builds global memory, decoder reuses it.
- Benchmark nuance: 74.2 on DeepSWE vs Claude 74, but only 31.2 on Terminal-Bench vs Claude 51.8.
- Not a small model: ~552B params (MoE), ~8B active/token during prompt, ~16B during generation; ~511 GB checkpoint.
- Native visual understanding added (images, UI screenshots, game menus → analysis and code generation).
- Over-thinking flaw: max effort can use up to 2.5x output tokens with no quality gain (87.13 vs ~86.75), +73% time, +60% cost.

## Technical data / figures

| Metric | DeepSeek V4.1 Flash | Reference |
|---|---|---|
| KV cache per token | ~890 bytes | V4 Flash ~3,514 bytes |
| KV cache reduction | ~4x vs V4 Flash | up to 437x per token vs gen. 1 |
| Context window | 1 million tokens | — |
| Total parameters | ~552B (MoE) | — |
| Active params (prompt) | ~8B / token | — |
| Active params (generation) | ~16B / token | — |
| Checkpoint size | ~511 GB | — |
| DeepSWE score | 74.2 | Claude 74 |
| Terminal-Bench score | 31.2 | Claude 51.8 |
| Writing test (standard) | 87.13 / 100 | — |
| Writing test (max effort) | ~86.75 / 100 | +73% time, +60% cost |

- Key technique: **CSA2** (Compressed Sparse Attention)
- Global KV-writing layers: **2, 8, 14, 20** (of 40)
- Architecture split: **20 encoder + 20 decoder** layers

## Why this source matters for the RAG

This article documents a concrete architectural breakthrough (CSA2) in KV-cache compression for long-context agents, with precise memory figures and benchmark nuance. It is highly relevant for a RAG knowledge base on LLM inference efficiency, MoE models, long-context handling, and the hardware-bottleneck race.

## Source URL

https://www.nerdykings.com/blog/deepseek-v4-1-flash-kv-cache-437x.html
