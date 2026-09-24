---
id: collect-240926-nerdykings/nerdykings/deepseek-v4-flash-x7-de-performance-grace-au-post-entrainement
title: "DeepSeek V4 Flash: 7x Performance Boost Thanks to Post-Training"
domain: nerdykings
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["deepseek", "training", "agent", "agentic", "agents", "benchmark", "benchmarks", "distillation", "gpus", "inference", "license", "mit license"]
source: docs/RAG/clean_en/nerdykings/deepseek-v4-flash-x7-de-performance-grace-au-post-entrainement.md
source_anchor: ""
source_lines: [1, 65]
sha256: 5fc28612b47d21094a8a07cb0a41d51bb76e33b6b75767742492aea3cfbdf36b
---

# DeepSeek V4 Flash: 7x Performance Boost Thanks to Post-Training

<!-- source: https://www.nerdykings.com/blog/deepseek-v4-flash-post-entrainement-x7.html -->

# DeepSeek V4 Flash: 7x Performance Boost Thanks to Post-Training

DeepSeek has just released an update that seems almost absurd. Its **V4 Flash** model came out just a few months ago, and after *a single revision*, its performance exploded: on one of the most demanding benchmarks for coding agents, its score was multiplied by more than 7. And this new version now surpasses DeepSeek V4 Pro Preview on several tasks, even though the Pro model activates nearly four times as many parameters per token. How can such a lightweight model progress so quickly? The answer lies almost entirely in **post-training**.

## A small model that surpasses its big brother

DeepSeek V4 Flash uses an architecture that is now well known: the **mixture of experts**. The model contains hundreds of billions of parameters in total, but only **about 13 billion are activated** to process each token. It's a bit like an immense company filled with specialists: for each task, the system only calls upon the experts that are actually useful.

By comparison, **DeepSeek V4 Pro Preview** activates nearly **49 billion parameters per token** — almost four times more capacity mobilized. On paper, Pro should have a crushing advantage. Except that the new version of Flash surpasses it on several agentic tasks. DeepSeek has therefore made a huge leap in performance *without* turning Flash into a much bigger model. And that, we really like.

## The numbers: from 7.3 to 54.4 on DeepSWE

The first benchmark is called **DeepSWE**. The principle: the model is given a real software problem, similar to a GitHub ticket. It must open the repository, understand the project, find the origin of the bug, modify the right files, then rerun tests to verify that the fix works. It's a formidable test, because a good coding agent must maintain a coherent strategy across multiple steps, use the right tools, and react when a first attempt fails. I had made a full video about it — it's also this benchmark that revealed that several major AIs were cheating.

The results:

- **Old version of Flash: 7.3**
- **New version of Flash: 54.4** — a progression of about **7.4×** in a few months
- V4 Pro Preview, on the same test: **12.8**

The second important result comes from **Terminal-Bench 2.1**. This time, the model works directly in a terminal: manipulating files, running commands, installing tools, solving multi-step tasks. Flash goes from **61.8 to 82.7**, whereas V4 Pro Preview reached 72.1.

Two benchmarks, the same story: DeepSeek above all improved the model's ability to **work over the long haul**, to string together the right decisions and to recover when an action fails. Important precision — the tests were carried out with the maximum reasoning level and DeepSeek's in-house agent system, so in a configuration optimized for complex tasks. But even with that nuance, the gap with the old version remains enormous.

## Pre-training vs post-training: the real difference

To understand this leap, we need to distinguish two major stages in the creation of a model.

**Pre-training**, first: the model absorbs immense quantities of text, code, and data. It learns language, concepts, and builds its general knowledge.

Then comes **post-training** — the stage that interests us here. It's the phase where we improve the way the model *uses* what it already knows to solve a specific task.

The most telling image is chess. Someone can know all the rules of the game without knowing how to build a good strategy. With training, they learn to choose their moves better, to anticipate, to correct their mistakes. For an AI, it's exactly the same. And it's precisely this stage that DeepSeek seems to have particularly well succeeded at: the model learns to plan better, to use its tools, to compare several solutions and to adjust its method when a first attempt fails. Which corresponds *very exactly* to the skills measured by DeepSWE and Terminal-Bench.

## How they did it: distillation + reinforcement learning

Yes, other laboratories also do post-training. But DeepSeek obtained a particularly effective result here — to the point of advancing their Flash version until it surpasses their Pro version on several tasks. Several methods seem to have contributed.

The first is **distillation**. DeepSeek relies on specialized models — in programming, mathematics, tool use — then transfers part of their strategy to Flash. Imagine several specialized teachers each passing on their best method to the same student.

The second is **reinforcement learning**. For the same instruction, the model generates several solutions, then the system compares them and evaluates their effectiveness. In the field of code, this evaluation can be very concrete: the program is executed, the tests are run, the fix is automatically verified. The trajectories that work are rewarded more. Over the course of training, the model learns to better organize its steps, to choose the right actions, and to course-correct when something breaks.

DeepSeek has not published the exact recipe for this update. But the available information clearly places post-training at the heart of the performance leap.

## And what about DSpark in all this?

In parallel, DeepSeek is developing a technology called **DSpark**, whose goal is to accelerate generation through **speculative decoding**. A lightweight module proposes several tokens in advance, then the main model verifies several proposals in a single operation. The validated tokens are kept, which allows for much faster generation while preserving the main model's output. It's a fascinating topic that I detailed in a full article on DSpark and speculative decoding.

## Open, downloadable, and ridiculously cheap

This may be the most important point of this entire release. DeepSeek V4 Flash has **weights downloadable under an MIT license**. Researchers and companies can obtain the model, host it, modify it, and integrate it into their own systems.

A caveat, though: running it locally requires an extremely powerful machine. Quantized versions still exceed **100 GB**. So we're talking about a large workstation or a server, not a gaming PC.

For most users, the API remains simpler — and this is where DeepSeek hits very hard: about **14 cents per million tokens for input and 28 cents for output**. At that price, agents can work on large code repositories, use numerous tools, and chain several reasoning steps on a laughable budget. A model that performs well on agentic tasks, open, downloadable, and extremely inexpensive: this is exactly the strategy DeepSeek has been applying since the release of V4.

## My take

What strikes me is not the ×7 figure itself — it's *where it comes from*. Not from a bigger model. Not from more GPUs. Not from a revolutionary new architecture. Just a better way of working: planning better, using the right tools, testing its solutions, correcting its mistakes faster.

This confirms a pattern I see recurring in almost all recent papers: the next advances in AI will not come only from bigger models, but from **better learning methods, better rewards, and better agentic behaviors**. This is exactly the same spirit as DeepSeek's inference optimizations: winning through the intelligence of the system rather than through brute force.

And let's be honest about what this implies: if an open model with 13 billion active parameters can multiply its agentic performance by 7 in a single revision, at 14 cents per million tokens… open models could catch up to the most powerful systems on the market **much faster than expected**.
