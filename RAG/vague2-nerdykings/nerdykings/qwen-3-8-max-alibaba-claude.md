---
id: vague2-nerdykings/nerdykings/qwen-3-8-max-alibaba-claude
title: "Qwen 3.8 Max : Alibaba Vient-Il Vraiment De Rattraper Claude ?"
domain: nerdykings
role: reference
task: article
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "United States"]
dates: ["2026-09-23"]
keywords: ["claude", "qwen", "agentic", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "deepseek", "fable 5", "gpus"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/qwen-3-8-max-alibaba-claude.md
source_anchor: ""
source_lines: [1, 69]
sha256: 8b800ab9ce515823f3831392f63ff0822d6b718c90673331258e3c7dc73d49b5
---

# Qwen 3.8 Max : Alibaba Vient-Il Vraiment De Rattraper Claude ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/qwen-3-8-max-alibaba-claude.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Alibaba released **Qwen 3.8 Max Preview** with a bold claim: it would now be the 2nd best model in the world, just behind Claude Fable 5. No official benchmark was communicated, but independent tests clearly support this on reasoning, programming, and visual interface creation. The author tested it on his usual prompts — architecture, benchmarks, and real-world tests — to judge whether Alibaba has really caught up with Claude.

An architecture built for high volumes: Qwen 3.8 Max uses a **mixture of experts** architecture — its **2.4 trillion parameters** are not all activated at once to generate each token. Depending on the demand, the model selects only the most relevant experts. Nothing new in itself; many frontier models use this principle to build a gigantic model without activating the whole network at each response. The problem is that Alibaba still hasn't communicated the exact number of active parameters — the one that really matters for judging speed, real cost, and required compute. For now, the 2.4 trillion mainly gives an idea of the model's scale, not its real efficiency.

The real innovation lies in long-context management. Qwen combines classic attention with a mechanism called **gated delta net**: instead of endlessly recomputing relations between all previous tokens, the model keeps a compact internal state summarizing part of the history, while still using some classic attention layers when it needs to retrieve precise information. The goal: combine efficient memory for long sequences with the ability to retrieve exact details — useful for working on massive codebases or long documents without exploding memory. But a large context window never by itself guarantees the model will use it well.

Multimodal, and especially good at visual programming: Qwen 3.8 Max doesn't only process text — it also analyzes images, videos, and complex documents in the same workflow. You can send it a screenshot, a document, and text instructions and ask it to combine everything in its answer. What makes this multimodality interesting for development is that the model can analyze an interface's appearance, understand its structure, then turn that understanding into concrete action — not just describe an interface, but reproduce it, improve it, or build a functional application from a simple visual reference. First feedback is most impressive there: Qwen seems particularly comfortable with front-end, animations, simulations, and visual page structuring. Many models generate technically correct code producing generic or poorly balanced interfaces; Qwen seems to better understand the relation between code and the visible result — element layout, interactions, general experience.

Independent benchmarks: one point behind Fable 5. On **King Bench**, a benchmark focused on reasoning, code, visual tasks, and agentic workflows, Qwen scores **65 out of 80** — one point behind Claude Fable 5, and ahead of Claude Sonnet 4.8, Kimi K3, and GPT. To be taken with caution: the benchmark contains only a limited number of challenges and clearly favors programming, technical reasoning, and visual creation. It can't be said Qwen is universally better than everything ranked below it, but it helps define its profile: particularly performant when understanding a technical problem, producing code, and turning that reasoning into a directly usable result.

The real test: portfolio, archery game, and Mesopotamian scene. First test: transform a hand-drawn mockup — a portfolio for a front-end developer — into a modern website. Result: header, hero, sections respected to the letter, fluid animations, well-chosen colors. Fully validated. Second, more demanding prompt: a 3D archery game, a classic test that traps most models on the bow's position. Qwen places it correctly on the first try and simulates gravity and shot power very well. The first generation was visually basic, so the author asked to improve graphics and playability — the second version changed dimensions completely: much larger scene, dense vegetation, more realistic animations. Only downside: a display problem with targets. Third test, the "Forest Robot" prompt with the Nerdy Kings mascot — a mini 3D adventure game where the robot collects coins while avoiding bears. Fluid movements, clean camera management (rare on this prompt), and a novel feature: nearby coins automatically come to the player. Last test, the favorite prompt — a 3D Mesopotamian scene inspired by Sumerian civilization. Rarely seen such an impressive result: varied architecture, realistic palm trees, a considerable number of NPCs moving without lagging the scene (which almost always happens with this prompt), coherent shadows, a sunset sky with reflections on the bricks. The Sumerian ziggurat at the center, overlooking the whole city — one of the best renders obtained on this prompt across all models.

Limits: thinking time and real access to the model. Qwen's greatest strength can become its main flaw. The model uses a particularly deep **thinking** mode, helping it compare several approaches and anticipate errors on hard problems. The problem is it doesn't always seem to know when this reflection is really needed: on a simple task it can keep analyzing for a very long time before producing a result. To use it efficiently, give very precise instructions — expected result, allowed technologies, mandatory features — otherwise waiting can become frustrating. The model is still in preview, so Alibaba will likely modify it before official release. Alibaba also says weights will be published soon, but even open source, a model this size remains largely inaccessible locally: even heavily compressed it would represent over a terabyte, requiring several professional GPUs just to run it. In practice, Qwen 3.8 Max will mostly be used via Alibaba's cloud services or specialized providers.

The author's view: Qwen 3.8 Max combines several qualities that, put together, are really interesting: an architecture built for long contexts, solid multimodal understanding, advanced reasoning, and very good visual programming. On real-world tests it clearly holds up against the best current models, especially for interface creation and turning visual references into functional results — the Mesopotamian scene alone is worth the detour. Honestly, it's not yet a clear replacement for Claude Fable 5. Excessive reflection time on simple tasks remains a real usability flaw, and its reliability on long, fully autonomous missions remains to be demonstrated — a problem seen elsewhere, e.g., with benchmark biases revealed on DeepSWE. Still, from a purely practical standpoint, Qwen 3.8 Max proves once again that the race between Chinese and American labs — after DeepSeek V4 earlier this year — is clearly just beginning. A very promising model, but still in preview.

## Key points

- **Qwen 3.8 Max Preview** (Alibaba) is claimed to be the 2nd best model worldwide, behind Claude Fable 5.
- Architecture: **Mixture of Experts** with **2.4 trillion parameters** (active parameter count not disclosed).
- Long-context innovation: combines classic attention with **gated delta net** for a compact summarized internal state.
- Multimodal: handles text, images, videos, and complex documents in one workflow.
- Especially strong at front-end, animations, simulations, and turning visual references into functional apps.
- King Bench: **65/80** — one point behind Claude Fable 5, ahead of Sonnet 4.8, Kimi K3, and GPT.
- Real tests validated: hand-drawn portfolio → website; 3D archery game; Forest Robot; Mesopotamian scene (excellent).
- Main flaw: over-thinks simple tasks; deep "thinking" mode can cause long waits without clear instructions.
- Still a preview; weights promised but a >1 TB model requires multiple professional GPUs.
- Practical use will mostly be via Alibaba's cloud or specialized providers.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Qwen 3.8 Max Preview |
| Developer | Alibaba |
| Architecture | Mixture of Experts |
| Total parameters | 2,400 billion (2.4T) |
| Active parameters | Not disclosed |
| King Bench score | 65 / 80 |
| Ranking claim | 2nd behind Claude Fable 5 |
| Beaten in benchmark | Claude Sonnet 4.8, Kimi K3, GPT |
| Long-context mechanism | Gated delta net + classic attention |
| Modalities | Text, image, video, documents |
| Size (open-weight, compressed) | >1 terabyte; multiple pro GPUs |
| Status | Preview |

- Key concepts: **MoE**, **gated delta net**, **long context**, **visual programming**, **thinking mode**
- Real-world tests: portfolio, 3D archery, Forest Robot, Mesopotamian 3D scene

## Why this source matters for the RAG

This article provides a hands-on evaluation of a major Chinese frontier model with architectural details, independent benchmark data, and concrete usability trade-offs. It is valuable for a RAG knowledge base on multimodal models, MoE long-context designs, front-end code generation, and the US-China model race.

## Source URL

https://www.nerdykings.com/blog/qwen-3-8-max-alibaba-claude.html
