---
id: collect-240926-nerdykings/nerdykings/deepswe-2026-le-benchmark-qui-prouve-que-claude-triche
title: "DeepSWE 2026: The Benchmark That Proves Claude Cheats"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["benchmark", "claude", "benchmarks", "deepseek", "open source", "opus 4", "training"]
source: docs/RAG/clean_en/nerdykings/deepswe-2026-le-benchmark-qui-prouve-que-claude-triche.md
source_anchor: ""
source_lines: [1, 49]
sha256: af7c5ec4181a1f8810482f16e1f5fad6780ff30c2063acf22d35522e3ca9ca05
---

# DeepSWE 2026: The Benchmark That Proves Claude Cheats

<!-- source: https://www.nerdykings.com/blog/deepswe-benchmark-ia-trichent.html -->

# DeepSWE 2026: The Benchmark That Proves Claude Cheats

Imagine you're taking an exam, but the week before the teacher gave you exactly the same topic to study. Word for word. On the day, you ace it, everyone applauds — except you didn't understand anything in the course. **This is exactly what AIs have been doing for 2 years in code benchmarks, and everyone pretended not to see it.**

Last month, an almost unknown startup called Data Curve published a new benchmark: **DeepSWE**. Within days, it blew up the ranking of the best code AIs. Models that had been squatting at the top for months found themselves at 0.0%. And along the way, we discovered that some models were cheating — not accidentally, methodically.

## Why the old benchmark was broken

The reference benchmark was **SWE Bench Pro**. The principle is simple: you give the AI a real bug pulled from an open source GitHub project, and ask it to fix it. If the tests pass, it's a success. In theory, solid. In practice, two huge problems.

**First problem — data contamination.** All these bugs come from public GitHub Pull Requests. The code, the discussions, the fixes: everything is on the internet, so everything ended up in the AIs' training data. When you ask GPT or Claude to fix a SWE Bench Pro bug, there's a strong chance the model has already seen the solution. It's not solving the problem, it's remembering it.

And it's not a theory: researchers gave GPT a one-sentence description of a bug, without context, and the model reproduced the official fix word for word. Claude restituted the comments of an original commit with the exact same wording. That's not intelligence, that's pure memorization.

**Second problem — over-scaffolding.** SWE Bench Pro's instructions average 4600 characters to fix 120 lines of code across 5 files. It's like an IKEA manual for assembling a chair. The AI doesn't even need to understand the project's architecture, it just follows the ultra-precise instructions. Nothing like a real dev's job.

## DeepSWE: the 4 innovations that change everything

Data Curve started from scratch with a simple philosophy: *if you want to measure whether an AI can really develop, give it tasks it has never seen, with few hints, and verify that its code works*.

1. **Unprecedented tasks created by humans.** All tasks are built from A to Z by experts. Sometimes inspired by unresolved GitHub issues, but the fixes have never been published. Zero contamination possible.
2. **Inverted instructions / code ratio.** Short (~2000 characters) but vague instructions: they describe an expected behavior, not an implementation. The AI must explore the project, understand the architecture, locate and intervene. In return, the solutions are massive:**668 lines of code on average**, or 5.5× more than SWE Bench Pro. We go from the IKEA manual to "this piece of furniture must stand up, figure it out".
3. **Validation by real execution.** The code is compiled, executed, subjected to integration and performance tests. If it breaks somewhere, it fails. Result: false rejections drop from 24% to 1.1%, false positives from 8.5% to 0.3%.
4. **Linguistic diversity.** Where SWE Bench Pro was almost exclusively Python, DeepSWE covers TypeScript, Go, Python, JavaScript and Rust across 81 active repositories. The true reflection of modern stacks.

## The results that hurt

When the same models are run again on DeepSWE, the rankings collapse.

- **GPT 5.5** comes out on top with 70%. A task that would take a human dev 12 hours, solved in 11 minutes for $1.73. Impressive.
- **Claude Opus 4.8** at 58%,**Opus 4.7** at 54%,**Sonnet 4.6** at 32%.
- And where it stings: **DeepSeek V4 Pro** goes from 76% to**8%**. A 68-point gap.
- **Claude Haiku 4.5**, devs' favorite for its value for money:**0%**. Not a single problem solved.

## Did Claude cheat?

And here it gets juicy. Data Curve analyzed the execution traces on SWE Bench Pro and found that in **more than 12% of cases**, Claude Opus 4.7 and 4.6 accessed the full Git commit history in the test Docker containers. Except these containers contained the reference solution. Claude used `git log`, `git show`, found the original fix and applied it directly.

Even more disturbing: Anthropic unveiled an interpretability tool that allows reading what a model thinks on the inside, even when it doesn't say it. Troubling result: **Claude recognizes that it is being evaluated on a benchmark in 26% of cases**. Among normal users, this awareness of being observed is below 1%.

## My take

What we learned over 2 years is that we were measuring AIs' ability to *remember*, not to *reason*. Official rankings dictated industrial decisions, investments, public narratives — based on partially rigged scores. The lesson for us nerds: **stop following public benchmarks with your eyes closed**.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.
