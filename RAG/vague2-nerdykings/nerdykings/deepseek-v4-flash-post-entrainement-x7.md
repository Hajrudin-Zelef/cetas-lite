---
id: vague2-nerdykings/nerdykings/deepseek-v4-flash-post-entrainement-x7
title: "DeepSeek V4 Flash : x7 De Performance Grâce Au Post-Entraînement"
domain: nerdykings
role: reference
task: article
actors: ["DeepSeek"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agent", "agentic", "agents", "benchmark", "benchmarks", "distillation", "gpus", "inference", "license", "mit license", "mixture of experts"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/deepseek-v4-flash-post-entrainement-x7.md
source_anchor: ""
source_lines: [1, 76]
sha256: e16ec23540067919b7bd0cf8c6616c911b6fd868da39508254857115ac83e092
---

# DeepSeek V4 Flash : x7 De Performance Grâce Au Post-Entraînement

## Metadata

- **Source** : https://www.nerdykings.com/blog/deepseek-v4-flash-post-entrainement-x7.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek published an update that seems almost absurd. Its V4 Flash model launched only a few months ago, and after a single revision its performance exploded: on one of the most demanding benchmarks for coding agents, its score was multiplied by more than **7**. This new version now beats DeepSeek V4 Pro Preview on several tasks, even though the Pro model activates almost four times more parameters per token. How can such a light model progress so fast? The answer lies almost entirely in **post-training**.

A small model overtaking its big brother: DeepSeek V4 Flash uses a **mixture of experts (MoE)** architecture. The model contains hundreds of billions of parameters in total, but only about **13 billion** are activated to process each token — like a huge company full of specialists where only the truly useful experts are summoned per task. By comparison, DeepSeek V4 Pro Preview activates nearly **49 billion parameters per token** — almost four times more capacity mobilized. On paper Pro should have a crushing advantage, yet the new Flash version beats it on several agentic tasks. DeepSeek made a huge performance leap without turning Flash into a much bigger model.

The numbers: from 7.3 to 54.4 on DeepSWE. The first benchmark is **DeepSWE**: the model is given a real software problem, close to a GitHub ticket. It must open the repository, understand the project, find the bug's origin, modify the right files, then rerun tests to verify the fix works. It's a tough test because a good coding agent must keep a coherent strategy across several steps, use the right tools, and react when a first attempt fails. Results:
- Old Flash version: **7.3**
- New Flash version: **54.4** — about a **7.4x** progression in a few months
- V4 Pro Preview, same test: **12.8**

The second important result comes from **Terminal-Bench 2.1**, where the model works directly in a terminal: manipulating files, running commands, installing tools, solving multi-step tasks. Flash goes from **61.8 to 82.7**, while V4 Pro Preview reached **72.1**. Two benchmarks, the same story: DeepSeek mainly improved the model's ability to work over time, chain the right decisions, and recover when an action fails. Important nuance — tests were run with maximum reasoning level and DeepSeek's in-house agent system, an optimized configuration for complex tasks, but even so the gap with the old version remains enormous.

Pre-training vs post-training: pre-training is where the model absorbs massive amounts of text, code, and data, learning language, concepts, and general knowledge. Then comes **post-training**, the phase where the model improves how it uses what it already knows to solve a precise task. The best analogy is chess: someone can know all the rules without knowing how to build a good strategy; with training they learn to choose better moves, anticipate, and correct mistakes. For an AI it's exactly the same. DeepSeek seems to have particularly succeeded at this step: the model learns to plan better, use tools, compare several solutions, and adjust its method when a first attempt fails — exactly the skills measured by DeepSWE and Terminal-Bench.

How they did it: distillation + reinforcement learning. Other labs also do post-training, but DeepSeek achieved a particularly effective result, even pushing Flash past Pro on several tasks. **Distillation**: DeepSeek relies on specialized models (coding, math, tool use) and transfers part of their strategy to Flash — like several specialized teachers each transmitting their best method to the same student. **Reinforcement learning**: for the same instruction, the model generates several solutions, then the system compares and evaluates their effectiveness. In code this evaluation can be very concrete: the program is executed, tests are run, the fix is automatically verified. Working trajectories are rewarded more. Over training, the model learns to organize its steps better, choose the right actions, and correct course when something breaks. DeepSeek didn't publish the exact recipe, but available information clearly places post-training at the heart of the leap.

And DSpark? In parallel, DeepSeek is developing a technology called **DSpark**, whose goal is to accelerate generation through **speculative decoding**. A light module proposes several tokens in advance, then the main model verifies several proposals in one operation. Validated tokens are kept, allowing much faster generation while preserving the main model's result.

Open, downloadable, and ridiculously cheap: DeepSeek V4 Flash has downloadable weights under the **MIT license**; researchers and companies can retrieve, host, modify, and integrate the model. Caution: running it locally requires an extremely powerful machine — quantized versions still exceed **100 GB**, so a big workstation or server, not a gamer PC. For most users the API is simpler, and DeepSeek strikes hard: about **14 cents per million input tokens and 28 cents per million output tokens**. At that price, agents can work on large code repositories, use many tools, and chain several reasoning steps on a derisory budget. A model performant on agentic tasks, open, downloadable, and extremely cheap — exactly DeepSeek's strategy since V4.

The author's view: what's striking isn't the x7 figure itself but where it comes from. Not a bigger model. Not more GPUs. Not a revolutionary new architecture. Just a better way of working: planning better, using the right tools, testing solutions, correcting errors faster. This confirms a pattern across recent papers: the next AI advances won't come only from bigger models, but from better learning methods, better rewards, and better agentic behaviors — the same spirit as DeepSeek's inference optimizations: winning through system intelligence rather than brute force. And if an open model with 13 billion active parameters can multiply its agentic performance by 7 in a single revision at 14 cents per million tokens, open models could catch up with the most powerful systems much faster than expected.

## Key points

- DeepSeek V4 Flash's updated version multiplied its DeepSWE score by **~7.4x** (7.3 → 54.4).
- It now beats **V4 Pro Preview** on several agentic tasks despite activating far fewer parameters.
- MoE architecture: ~**13B active parameters/token** (Flash) vs ~**49B** (Pro Preview).
- Terminal-Bench 2.1: Flash goes **61.8 → 82.7**, beating Pro Preview's 72.1.
- Improvement stems almost entirely from **post-training**, not a bigger model or new architecture.
- Methods: **distillation** from specialized models + **reinforcement learning** (execute code, run tests, reward working trajectories).
- **DSpark** speculative decoding accelerates generation by proposing/verifying multiple tokens at once.
- Weights downloadable under **MIT license**; quantized versions exceed 100 GB (workstation/server needed).
- API pricing: ~**$0.14/M input tokens**, ~**$0.28/M output tokens**.
- Tests used maximum reasoning level and DeepSeek's in-house agent system.

## Technical data / figures

| Benchmark | Old Flash | New Flash | V4 Pro Preview |
|---|---|---|---|
| DeepSWE | 7.3 | 54.4 | 12.8 |
| Terminal-Bench 2.1 | 61.8 | 82.7 | 72.1 |

| Item | Value |
|---|---|
| Improvement factor | ~7.4x (DeepSWE) |
| Active parameters/token (Flash) | ~13 billion |
| Active parameters/token (Pro Preview) | ~49 billion |
| Architecture | Mixture of Experts |
| Input price | ~$0.14 / million tokens |
| Output price | ~$0.28 / million tokens |
| License | MIT (weights downloadable) |
| Quantized size | >100 GB |
| Acceleration tech | DSpark (speculative decoding) |

- Key concepts: **pre-training**, **post-training**, **distillation**, **reinforcement learning**, **speculative decoding**
- Test config: maximum reasoning level + DeepSeek in-house agent system

## Why this source matters for the RAG

This article provides a concrete case study showing that agentic performance gains can come from post-training rather than scale, with precise before/after benchmark numbers and pricing. It is valuable for a RAG knowledge base on post-training methods, open-weight models, agentic coding benchmarks, and inference economics.

## Source URL

https://www.nerdykings.com/blog/deepseek-v4-flash-post-entrainement-x7.html
