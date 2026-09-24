---
id: collect-240926-nerdykings/nerdykings/codex-spark-le-modele-le-plus-rapide
title: "Codex 5.3 Spark: revolution, or a huge trap for devs?"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "Cerebras", "OpenAI"]
dates: []
keywords: ["accelerator", "benchmark", "chatgpt", "claude", "exploit", "gpu", "gpus", "liquid cooling", "memory", "tokens per second", "wafer"]
source: docs/RAG/clean_en/nerdykings/codex-spark-le-modele-le-plus-rapide.md
source_anchor: ""
source_lines: [1, 33]
sha256: c7db935b64c345850a32c0aeec9391b659a449df82a3525b8bf74a6b1805caee
---

# Codex 5.3 Spark: revolution, or a huge trap for devs?

<!-- source: https://www.nerdykings.com/outils/codex-5-3-spark.html -->

# Codex 5.3 Spark: revolution, or a huge trap for devs?

Just as Codex 5.3 has barely been released, OpenAI is already drawing a new model: Spark, reserved for now for ChatGPT Pro subscribers at $200/month. The promise: unprecedented speed. I took apart what's hidden under the hood, and what it really changes in use.

## The technical secret: Cerebras rather than classic GPUs

OpenAI dropped traditional GPU clusters to partner with **Cerebras Systems**, which makes the largest AI chip in the world (Wafer Scale Engine 3, **4 trillion transistors**). The trick: instead of thousands of small chips that have to communicate with each other, the whole model and its context fit on a single giant chip — which eliminates communication bottlenecks.

Measured result: more than **1000 tokens per second**, or 15 times more than the standard model, with the delay before the first word reduced by 50%. Trade-off: the machine consumes up to 240 kW and requires cutting-edge liquid cooling — which explains the still-limited access for Pro subscribers.

## Speed has a price: less reliable on complex logic

On the terminal-bench 2.0 benchmark (solving complex tasks in a terminal), Spark reaches **58.4%** versus **77.3%** for classic Codex. But on pure speed, the gap is spectacular: 2 to 3 minutes for Spark versus up to 15 minutes for Codex on the same complex programming task.

The downside, documented by several user reports: "fast hallucinations" — clean-looking code, generated very quickly, but with logical errors or libraries that don't exist. On tasks with several chained logical steps, Spark can drift and lose coherence without it being obvious at first glance.

## The new features of the Codex app

Spark's speed enables new uses: duplicating your session to explore two technical solutions in parallel and compare them instantly, a floating window that stays above other applications (handy for front-end, seeing the result without changing tabs), and a CLI that lets you manually clear the AI's memory.

## The real workflow: Spark for speed, another model for reliability

The best way to use Spark is not to trust it from start to finish, but to exploit it for what it does best: laying the foundations very quickly, without getting stuck. Then, the result is passed through a more solid model (standard Codex or Claude) that understands in depth, corrects errors, and secures it. Overall, it's not slower — it's even often faster than directly using a slow model, while keeping a high level of quality.

## My verdict

Spark isn't a replacement, it's an accelerator. For rapid prototyping, iteration, repetitive code: formidable. For critical code or complex systems without verification behind it: risky. Personally, the $200/month to access it seems expensive to me for now — but the feedback from the American devs using it remains impressive.

### Want to test Codex Spark?

Affiliate link: if you go through it, I earn a commission without it costing you more. It supports the channel, and it doesn't change my opinion.
