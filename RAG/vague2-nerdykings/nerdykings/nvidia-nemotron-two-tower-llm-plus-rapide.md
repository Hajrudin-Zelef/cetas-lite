---
id: vague2-nerdykings/nerdykings/nvidia-nemotron-two-tower-llm-plus-rapide
title: "Nvidia Nemotron Two Tower : 2,4x Plus Rapide, Mais À Quel Prix ?"
domain: nerdykings
role: reference
task: article
actors: ["DeepSeek", "Nvidia", "United States"]
dates: ["2026-09-23"]
keywords: ["nvidia", "agent", "agents", "attention", "compute", "cost", "deepseek", "gpus", "inference", "latency", "memory", "mixture of experts"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/nvidia-nemotron-two-tower-llm-plus-rapide.md
source_anchor: ""
source_lines: [1, 74]
sha256: 8793326dfa8abcd477c3b206f9a24a72afefa237928df9746779623c4b7ba524
---

# Nvidia Nemotron Two Tower : 2,4x Plus Rapide, Mais À Quel Prix ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/nvidia-nemotron-two-tower-llm-plus-rapide.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

In early July, Nvidia published an architecture called **Nemotron Two Tower**, with a simple promise: make a language model about **2.4x faster** without starting from scratch or sacrificing most of its capabilities. But digging into the numbers reveals a bill not advertised in the title: to speed up a model, Nvidia must practically load two of them into memory.

The philosophy shift behind the paper: for years, US labs mainly sought to build ever-bigger models, trained on ever more data with ever more GPUs. But for a few months, DeepSeek has pushed a different approach: get more from the same base through architecture, compression, and inference optimization (as seen with Dual Path and DSpark). This Nvidia paper fits exactly that logic: instead of predicting only the next token, a model learns to anticipate several at once.

How Two Tower works: most current language models generate text **autoregressively** — the model produces a token, then uses it to compute the next. It must finish step 9 before starting step 10. This is very reliable, especially for code or math, because each decision builds on everything already validated. It's like an extremely fast typewriter that must always press keys one by one. For a classic chatbot this isn't dramatic — text is generated faster than one can read anyway. But for an agent that calls a model dozens of times, uses a tool, reads the result, thinks, and repeats, a few seconds lost at each step quickly turns a simple task into an endless workflow.

That's where Nvidia steps in. Nemotron Two Tower builds on **Nemotron 3 Nano 30B-A3B**, a hybrid model mixing **Mamba**, classic attention, and **Mixture of Experts**. Nvidia duplicates this base into two towers: the **context tower**, which reads the request and keeps a clean representation of everything already generated, and the **denoising tower**, which learns to generate several tokens in parallel. The simplest image: a writer who perfectly knows the already-validated chapter, and a proofreader who receives the next paragraph as 16 blank slots to fill simultaneously.

The base model was pre-trained on about **25 billion tokens**. Nvidia didn't redo that colossal training: it **froze the first tower** and trained only the second, on about **2.1 billion tokens**. At the start of each block, the model creates **16 masked positions**; the denoising tower tries to predict them simultaneously, looking in both directions within the block. Positions still uncertain are remasked and recomputed at the next iteration until all 16 are filled. So the system remains autoregressive at the block scale but parallel within each block.

2.42x faster, but not on everything. In Nvidia's tested configuration — 16-token blocks on two H100 GPUs — Two Tower generates text about **2.42x faster** than the base autoregressive version, with **98.7% of overall quality** preserved per Nvidia. On **MMLU**, the score barely moves: 78.56 (classic) vs 78.24 (Two Tower). However, on **HumanEval** (Python code), the score drops from about 79.3 to 75.6, and mathematics also loses a few points. This is explained well: for tasks requiring an extremely rigorous sequence like code or math, parallel generation lets more small errors through. In a program, a single parenthesis or the choice between a loop and a condition immediately changes all subsequent tokens. Nvidia allows returning to a purely autoregressive mode when precision matters more than speed.

The constraint that can't be exceeded: the model is trained to work with 16-token blocks, a hard limit. With 64-token blocks, HumanEval collapses to around 20% and GSM8K falls almost to zero. Parallelism works, but only within a very controlled window — it's not a button you can push at will to go even faster.

The real bill: the hardware tax. And here the catchy title hides the essential. Each tower contains the weights of a ~**30-billion-parameter** model. Even if only a few billion are active per token thanks to MoE, **both weight sets must stay loaded in memory**. At full precision, the system requires about **138 GB of VRAM**, on an official two-H100 configuration. Ironically, the architecture reduces generation time but nearly doubles the amount of weights to store. And the 2.42x figure describes a precise, well-optimized configuration — nothing guarantees a regular user will get that result on their own hardware.

Who it really changes things for: for a classic chat, displaying 16 tokens per burst isn't necessarily nicer than smooth word-by-word streaming. The real potential is elsewhere: **agents**. An agent making 30 model calls to finish a task sees each small latency gain repeated 30 times — becoming very visible. With ~2.4x acceleration, some workflows could go from several minutes to a much more acceptable time. Conversely, for writing sensitive code, solving complex math, or serving thousands of simultaneous cloud requests, the advantage is clearly less obvious.

The author's view: Two Tower probably isn't the universal future of language models. The hardware tax is really too high, and the performance drop in code shows token-by-token generation will be very hard to replace, especially when absolute sequential precision is required. But the paper is interesting for a reason far beyond Nvidia: US labs are finally adopting another philosophy — stop always going bigger and start optimizing what already exists, exactly what DeepSeek has repeated for six months. The real question isn't whether Two Tower will replace autoregression everywhere, but how long US giants will take to realize the race for size has its limits while others already optimize what exists.

## Key points

- **Nemotron Two Tower** (Nvidia) promises ~**2.4x faster** generation without retraining from scratch.
- Built on **Nemotron 3 Nano 30B-A3B**, a hybrid of Mamba, attention, and Mixture of Experts.
- Two towers: a **context tower** (reads/keeps representation) and a **denoising tower** (generates tokens in parallel).
- Base pre-trained on ~25B tokens; only the second tower trained on ~**2.1B tokens** (first tower frozen).
- Generates in **16-token blocks**: masked positions predicted in parallel, remasked until filled; autoregressive at block level.
- Measured **2.42x** speedup on two H100s with **98.7%** overall quality retained.
- MMLU nearly unchanged (78.56 → 78.24); HumanEval drops (~79.3 → 75.6); math also loses points.
- Hard limit: 64-token blocks collapse HumanEval to ~20% and GSM8K to near zero.
- Hardware tax: each tower holds ~30B params; full precision needs ~**138 GB VRAM** on two H100s.
- Best suited to agents (repeated calls) rather than sensitive code/math or high-concurrency cloud serving.

## Technical data / figures

| Metric | Value |
|---|---|
| Speedup | ~2.42x |
| Base model | Nemotron 3 Nano 30B-A3B |
| Architecture | Mamba + attention + Mixture of Experts |
| Towers | Context tower + denoising tower |
| Base pre-training tokens | ~25 billion |
| Second-tower training tokens | ~2.1 billion |
| Block size | 16 tokens |
| Quality retained | 98.7% |
| MMLU (classic → Two Tower) | 78.56 → 78.24 |
| HumanEval (classic → Two Tower) | ~79.3 → 75.6 |
| 64-token block HumanEval | ~20% |
| 64-token block GSM8K | near 0% |
| VRAM (full precision) | ~138 GB (two H100) |

- Key concepts: **autoregressive generation**, **parallel/block generation**, **denoising tower**, **context tower**, **Mamba**, **Mixture of Experts**
- Related context: DeepSeek **Dual Path**, **DSpark**

## Why this source matters for the RAG

This article provides a critical, quantified analysis of a parallel-generation architecture, including its hidden memory cost and task-dependent quality loss. It is valuable for a RAG knowledge base on inference acceleration, speculative/parallel decoding, and the shift from scaling to optimization.

## Source URL

https://www.nerdykings.com/blog/nvidia-nemotron-two-tower-llm-plus-rapide.html
