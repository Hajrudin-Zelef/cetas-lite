---
id: collect-240926-mindstudio/mindstudio/openai-s-vision-for-a-personal-agi-merging-chatgpt-and-codex
title: "openai-s-vision-for-a-personal-agi-merging-chatgpt-and-codex"
domain: mindstudio
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["agi", "chatgpt", "agent", "agentic", "agents", "attention", "cost", "cybersecurity", "inference", "latency", "memory", "merger"]
source: docs/RAG/clean_en/mindstudio/openai-s-vision-for-a-personal-agi-merging-chatgpt-and-codex.md
source_anchor: ""
source_lines: [1, 81]
sha256: 0b698076b7033fa381df43db0f742a6db21716ca3a8367c1f246a8bf065ea14f
---

# openai-s-vision-for-a-personal-agi-merging-chatgpt-and-codex

<!-- source: https://www.mindstudio.ai/blog/openai-tibo-interview-agi-vision -->

## What did OpenAI say about a “personal AGI”?

In an interview, OpenAI’s Tibo (who works on ChatGPT and Codex) described a future product direction where the boundary between a chat assistant and a coding agent disappears. Instead of separate tools for separate tasks, the goal is a single adaptive system that understands a specific person, their goals, their work, and their team, then acts on that understanding whether the task is technical or not. He framed this as building “the very best personal AGI or the personal agent,” one that stays proactive, raises ideas at the right moment, and never breaks the illusion of a coherent partner working alongside you.

## TL;DR

- OpenAI’s stated end-state is a **single adaptive assistant** that merges the conversational strengths of ChatGPT with the technical execution of Codex, rather than keeping them as separate products.
- Tibo described current agentic coding tools as still **clunky** , citing skill files, imperfect memory, and sub-agent management as friction points that break the sense of a smooth partner.
- Voice is positioned as a **core interface shift** , not a feature bolt-on. Faster, more natural voice models plus tool use already changed how Tibo personally works, including using dictation instead of typing prompts.
- Faster inference speeds are expected to change developer workflows, potentially reducing reliance on running many parallel agents in favor of a **real-time, flow-state** style of working.
- OpenAI separates its agent work into two buckets: a **personal, tailored assistant** for individual work and**full automation systems** for complex processes like patching vulnerabilities or fixing production regressions without a human in the loop.
- Tibo pointed to hardware itself, specifically the laptop, as an emerging **bottleneck** for future models, arguing that model capability is starting to outpace the resource constraints devices were built around.
- The cultural lesson he carried from his time at Google DeepMind was that a **bias toward shipping** and a willingness to disrupt existing products matters more than having a large research lead.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

## Why merge ChatGPT and Codex at all?

The logic, as Tibo laid it out, comes from watching how people actually use these tools once the underlying models get good enough. A user of Codex or any coding agent today has to manage a handful of separate constructs: skill files to teach it new behaviors, memory that is inconsistent, and sub-agents that need coordinating like a small team. Each of these is a workaround for the fact that the system doesn’t fully understand the user yet.

Tibo’s argument is that a sufficiently capable model shouldn’t need those scaffolds. Instead of a user maintaining files and managing sub-agent hierarchies, the assistant should already understand a person’s goals, daily patterns, and team context well enough to act appropriately without being re-taught. That single understanding is what would let a chat interface and a coding agent become the same product: the underlying model doesn’t care whether the task is “write code” or “give advice,” it’s all just work being done for one specific, well-understood user.

## How does voice change the assistant model?

Voice is not treated as a convenience feature in this vision, it’s described as changing the fundamental relationship between a person and the assistant. Tibo pointed to OpenAI’s newer voice models, which he described as natural-sounding and capable of tool use, as the reason he now spends much more time simply talking to the assistant rather than typing.

He gave a concrete personal example: using dictation in the morning to rattle off a list of tasks, which the assistant then executes using its access to his tools, without him typing a single prompt. That workflow, he said, wasn’t possible before voice models reached this level of quality and tool integration. The implication for the “merged assistant” vision is that voice becomes a primary input method for both conversational and technical tasks, not a separate mode you switch into.

## What problems still need solving in the agent harness?

Tibo was direct that today’s agentic coding setups still show their seams. Skill files are hard to maintain over time. Memory doesn’t reliably persist everything it should. Sub-agent orchestration creates a small network of moving parts that a user has to consciously track. Each of these is a place where the “illusion” of a single capable partner breaks down and the user is reminded they’re managing a system rather than collaborating with one.

Fixing this, in his framing, isn’t about adding more features, it’s about the model developing a deeper standing understanding of the user so those management tasks become unnecessary. That’s a meaningfully different bet than simply making agents faster or cheaper. It’s a bet that the interface problem and the personalization problem are the same problem.

## Why does hardware become the bottleneck?

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

A less obvious point Tibo raised is that laptops themselves start to limit what a powerful model can do. Laptops were designed around human constraints: how fast a person can type, how many applications a person can reasonably keep open, how much a person can absorb at once. Models don’t share those constraints. He noted that a sufficiently capable model could handle something like a hundred open applications at once without difficulty, something no human workflow was ever designed for.

This is presented as a reason future models will need access to more than local device resources, pointing toward cloud-based execution as a structural requirement rather than a convenience.

## How will faster inference change developer workflows?

The interview also touched on inference speed as a workflow variable, not just a cost or latency metric. At current speeds, a developer might kick off ten or fifteen coding agents in parallel and wait 30 to 45 minutes for results, a pattern that creates real cognitive overhead from constant context-switching between tasks.

Tibo suggested that dramatically faster token speeds change this calculus. If an assistant can operate at or above the speed of a person’s own thinking, the value of running many agents in parallel drops. Instead of juggling a dozen background tasks, a developer could stay in a continuous, real-time flow, iterating on ideas and seeing prototypes update live rather than checking back on batches of finished work. He framed the goal explicitly as building technology that adapts to a person’s attention rather than requiring the person to adapt to the technology.

## Personal assistant vs. full automation: what’s the difference?

Tibo split OpenAI’s agent-related work into two distinct categories. The first is the personal AGI described above: a tailored, proactive assistant deeply rooted in understanding one individual, useful for both technical and non-technical problems.

The second category is full automation of complex processes that don’t need a human tailored relationship at all, just reliable execution. His examples included automatically detecting and patching performance regressions from production logs, and in cybersecurity, automatically patching vulnerabilities flagged by a scanner to shrink the window of exposure, ideally without a human in the loop. This is a meaningfully different design target: instead of a system that understands you, it’s a system that removes the need for anyone to be involved at all for a specific, well-bounded task.

## Frequently Asked Questions

### What does “personal AGI” mean in this context?

It refers to a single adaptive assistant, described by OpenAI’s Tibo, that deeply understands an individual user’s goals, habits, and work context well enough to act helpfully across both conversational and technical tasks without needing constant re-teaching.

### Is OpenAI actually combining ChatGPT and Codex into one product?

The interview describes this as a direction and ambition rather than an announced product merger with a specific date. Tibo frames the current separation between chat and coding-agent experiences as a limitation to be overcome as models understand users more deeply.

### Why does OpenAI think laptops will become a bottleneck?

Because laptops were built around human limits like typing speed and how many applications a person can manage at once. Tibo argues that capable models don’t share those limits and will increasingly need cloud-based resources beyond what a personal device can provide.

### Will faster models mean running more agents at once?

Not necessarily. Tibo suggested the opposite: as inference speed increases, the need to run many parallel agents to compensate for slowness may decrease, replaced by a faster, more continuous, real-time style of working with a single responsive assistant.

### What’s the difference between the personal assistant and full automation efforts at OpenAI?

The personal assistant is tailored to one user’s context and preferences. Full automation targets specific, well-defined technical processes, like patching security vulnerabilities or fixing production regressions, designed to run without a human in the loop at all.
