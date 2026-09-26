---
id: collect-240926-nerdykings/nerdykings/qwen-3-8-max-alibaba-vient-il-vraiment-de-rattraper-claude-1
title: "Qwen 3.8 Max: Has Alibaba Really Caught Up with Claude?"
domain: nerdykings
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot"]
dates: []
keywords: ["claude", "qwen", "agentic", "attention", "benchmark", "benchmarks", "context window", "cost", "deepseek", "fable 5", "gpus", "kimi"]
source: docs/RAG/clean_en/nerdykings/qwen-3-8-max-alibaba-vient-il-vraiment-de-rattraper-claude.md
source_anchor: ""
source_lines: [1, 44]
sha256: 5d20c8a1a1a365f0234706b3e82b0216ca56e45275ca4fbcc23eb2f45a74eda0
---

# Qwen 3.8 Max: Has Alibaba Really Caught Up with Claude?

<!-- source: https://www.nerdykings.com/blog/qwen-3-8-max-alibaba-claude.html -->

# Qwen 3.8 Max: Has Alibaba Really Caught Up with Claude?

Alibaba has just released **Qwen 3.8 Max Preview**, and the claim is bold: it would now be the 2nd best model in the world, just behind **Claude Fable 5**. No official benchmark was shared, but independent tests clearly point in that direction on reasoning, programming, and visual interface creation. I wanted to verify this myself on my usual prompts — architecture, benchmarks, and above all a real-world test. Verdict: has Alibaba really caught up with Claude?

## An Architecture Designed for Large Volumes

Qwen 3.8 Max is built on a **mixture of experts** architecture: its 2.4 trillion parameters are not all activated at the same time to generate each token. Depending on the request, the model selects only the most relevant experts. Nothing new in itself — many frontier models use this principle to build a gigantic model without having to activate the entire network for each response. The problem is that Alibaba still hasn't disclosed the exact number of *active parameters* — the one that really matters for judging speed, real cost, and the computing power required. For now, the 2.4 trillion mainly gives us an idea of the model's scale, not its actual efficiency.

The real innovation lies more in long-context management. Qwen combines classic attention with a mechanism called **gated delta net**: instead of constantly recalculating the relationships between all previous tokens, the model keeps a compact internal state that summarizes part of the history, while continuing to use certain classic attention layers when it needs to retrieve precise information. The goal: combine efficient memory for long sequences with the ability to retrieve exact details — useful for working on massive codebases or long documents without blowing up memory. But a large context window never guarantees, on its own, that the model will make good use of it.

## Multimodal, and Especially Gifted at Visual Programming

Qwen 3.8 Max doesn't only process text: it also analyzes images, videos, and complex documents within the same workflow. Concretely, you can send it a screenshot, a document, and text instructions, and ask it to combine everything in its response. What makes this multimodality interesting for development is that the model can analyze the appearance of an interface, understand its structure, then *transform* what it understands into concrete action — not just describe an interface, but reproduce it, improve it, or build a functional application from a simple visual reference.

And that's exactly where the early feedback is most impressive: Qwen seems particularly comfortable with front-end, animations, simulations, and the visual structuring of a page. Many models today can generate technically correct code but that produces generic or poorly balanced interfaces. Qwen, on the other hand, seems to better understand the relationship between the code and the visible result on screen — the layout of elements, interactions, the overall experience.

## Independent Benchmarks: One Point Behind Fable 5

On **King Bench**, a benchmark focused on reasoning, code, visual tasks, and agentic workflows, Qwen scores **65 points out of 80** — just one point behind Claude Fable 5, and ahead of Claude Sonnet 4.8, Kimi K3, and GPT. Take this with a grain of salt: the benchmark contains only a limited number of tests and clearly favors programming, technical reasoning, and visual creation. It's impossible to claim that Qwen is universally better than everything ranked behind it, but it already helps define its profile: a model particularly performant when it comes to understanding a technical problem, producing code, and transforming that reasoning into a directly usable result.

## The Real Test: Portfolio, Archery Game, and Mesopotamian Scene

Benchmarks are all well and good, but I always prefer to verify on my own prompts. First test: transforming a hand-drawn mockup — a portfolio for a front-end developer — into a modern website. Result: header, hero, sections followed to the letter, smooth animations, well-chosen colors. Fully validated.

Second prompt, more demanding: a 3D archery game, a classic from my tests that trips up most models on the bow's position. Qwen places it correctly on the first try, and above all simulates gravity and shot power very well. The first generation remained a bit basic visually, so I asked it to improve the graphics and gameplay — the second version is a completely different league: much larger scene, dense vegetation, more realistic animations. Only downside: a display issue with the targets.

Third test, my Forest Robot prompt with the Nerdy Kings mascot — a 3D mini adventure game where the robot collects coins while avoiding bears. Smooth movements, clean camera management (rare on this prompt, usually the camera is poorly handled and the game becomes unplayable), and a novel feature: nearby coins automatically come to the player.

Last test, my favorite prompt — a Mesopotamian 3D scene inspired by Sumerian civilization. And honestly, I've rarely seen such an impressive result: varied architecture, realistic palm trees, a considerable number of NPCs moving around without making the scene lag (which happens almost every time with this prompt), consistent shadow play, a sky setting with reflections on the bricks. The Sumerian ziggurat in the center, overlooking the whole city — one of the best renders I've gotten on this prompt, across all models.

## The Limits: Thinking Time, and Real Access to the Model

Qwen's greatest strength can also become its main flaw. The model uses a particularly in-depth *thinking* mode, which helps it compare multiple approaches and anticipate errors on difficult problems. The issue is that it doesn't always seem to know when this reflection is actually necessary: on a simple task, it can keep analyzing for a very long time before producing a result. To use it effectively, it's better to give very precise instructions — expected result, allowed technologies, required features — otherwise the wait can become frustrating.

The model is still in *preview*, so Alibaba will very likely continue to modify it before the official release — behavior and performance could evolve. Alibaba also claims the weights will be published soon, but even as open source, a model of this size remains largely inaccessible locally for most users: even heavily compressed, it would represent more than a terabyte of data, and it would take several professional GPUs just to run it. In practice, Qwen 3.8 Max will mostly be used via Alibaba's cloud services or specialized providers.

## My Opinion

Qwen 3.8 Max combines several qualities that, put together, are truly interesting: an architecture designed for long contexts, solid multimodal understanding, advanced reasoning, and very good visual programming capabilities. In my real-world tests, it clearly holds its own against the best models of the moment, particularly on everything related to creating interfaces and transforming visual references into functional results — the Mesopotamian scene alone is worth the detour.

That said, let's be honest: it's not yet a clear replacement for Claude Fable 5. The excessive thinking time on simple tasks remains a real usability flaw, and its reliability on long, fully autonomous missions remains to be demonstrated — a problem we've already seen elsewhere, for example with the benchmark biases revealed on DeepSWE. Still, from a purely practical standpoint, Qwen 3.8 Max proves once again that the race between Chinese and American labs — after DeepSeek V4 earlier this year — is clearly only just beginning. A very promising model, but still in preview: to be revisited seriously upon its official release and real benchmarks.

