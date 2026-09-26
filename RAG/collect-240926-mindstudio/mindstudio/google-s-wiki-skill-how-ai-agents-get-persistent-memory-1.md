---
id: collect-240926-mindstudio/mindstudio/google-s-wiki-skill-how-ai-agents-get-persistent-memory-1
title: "google-s-wiki-skill-how-ai-agents-get-persistent-memory"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "memory", "context window", "fine-tuning", "inference", "research"]
source: docs/RAG/clean_en/mindstudio/google-s-wiki-skill-how-ai-agents-get-persistent-memory.md
source_anchor: ""
source_lines: [1, 63]
sha256: e3066131d34b1ce7979b0f24704f0b934939b7b4fa69a829233b27dba2820ab5
---

# google-s-wiki-skill-how-ai-agents-get-persistent-memory

<!-- source: https://www.mindstudio.ai/blog/google-wiki-skill-agent-memory -->

## What is Google’s Wiki Skill?

Wiki Skill is a research proposal from Google Research that describes how an AI agent can turn its own experience into knowledge that sticks around instead of evaporating at the end of a session. Rather than an agent relearning the same lesson every time it runs a task, the system separates what happened (raw logs), what was learned from it (a wiki of observations), and what the agent should actually do differently (skills). The paper credits Andrej Karpathy’s idea of an “LLM wiki” as a partial inspiration, then builds a more formal, agent-specific architecture on top of it.

## TL;DR

- Wiki Skill proposes a three layer memory system: a **raw layer** of unchangeable execution traces, a**wiki layer** of compounding observations, and a**skills layer** that actually changes agent behavior.
- The design draws on Andrej Karpathy’s **LLM wiki** concept, where knowledge is compiled once and kept current instead of being rederived on every query.
- A **four agent loop** runs continuously: an inference agent does the work, a wiki maintainer analyzes the logs for root causes, and additional agents update the wiki and refine skills.
- Keeping the wiki separate from the skills matters because **observations and actions serve different purposes** , one is a record of truth, the other is a set of instructions meant to change over time.
- The approach fits the broader shift toward agents that operate over long stretches without constant human correction, the same trend visible in reports about OpenAI’s more persistent internal models.
- Skills in this framework are plain text, typically markdown, which makes the system easy to inspect, edit, and reason about compared to opaque fine-tuning.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

## How does the Wiki Skill architecture work?

The system is built around three distinct layers, each with a different job.

The **raw layer** stores immutable execution traces. Every time an agent attempts a task, the record of what it did gets written down and never touched again. Nothing here is deleted or edited. It functions as a permanent audit trail, similar in spirit to how Karpathy’s LLM wiki keeps a running record of evidence rather than discarding it after use.

The **wiki layer** sits above that. This is where observations get compiled: patterns pulled out of the raw traces, notes on what worked, what failed, and why. Unlike the raw layer, the wiki layer compounds. It grows and refines itself over time, but it never resets. Think of it as the agent’s evolving understanding of its own domain, built from evidence rather than assumption.

The **skills layer** is the part that actually changes agent behavior. Skills are the instructions an agent follows to complete a task, and in practice they tend to be simple text files (often markdown) that describe a procedure step by step. This is the same format used elsewhere in the industry. Anthropic’s own skill files, for instance, are plain English markdown documents that an agent reads before attempting a task. Wiki Skill treats skills as something that evolves based on what the wiki layer has learned, but keeps them as a separate artifact from the wiki itself.

## Why keep the wiki and the skills separate?

This is the detail that makes the architecture more than a rebranded chat log. The wiki layer is a record of observations, a description of reality. The skills layer is a set of instructions meant to be acted on. Collapsing them into one file would mean every new observation directly rewrites the agent’s behavior, which risks instability: one bad trace could corrupt a skill that had been working fine.

By keeping them apart, the system can accumulate evidence in the wiki without immediately touching the skill file, and only promote validated patterns into skill updates. It’s a version of the classic separation between data and logic. The wiki is the database, the skill is the program that reads from it.

## What does the four agent loop actually do?

Wiki Skill runs as a continuous loop involving four roles:

The **inference agent** is the worker. It attempts the actual task, whatever that happens to be (a spreadsheet operation, a math problem, a coding task), using the current version of the skill.

The **wiki maintainer** performs root cause analysis on the raw traces. It looks at what the inference agent just did, figures out what succeeded, what failed, and why, and turns that into a new observation in the wiki layer.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

The remaining roles in the loop handle updating the wiki with new findings and refining the skill files themselves based on what the wiki has accumulated. Together these four roles form a loop that runs indefinitely, meaning the agent’s knowledge base is never static. Every run of the loop is a chance for the system to get slightly better at the task without a human manually rewriting its instructions.

## How does this compare to Karpathy’s LLM wiki idea?

Karpathy’s original framing was about compiling knowledge once and keeping it current, rather than making a model rederive the same conclusions from scratch on every query. The idea translates naturally to personal knowledge management: a growing set of cross-linked notes organized like a personal Wikipedia, where each new piece of information gets added to a persistent, searchable structure instead of being lost after one conversation.

Google’s Wiki Skill applies that same principle but points it specifically at agent behavior rather than general note-taking. The raw layer plays a role similar to the evidence store in an LLM wiki. But Wiki Skill adds the formal split between observation (the wiki) and instruction (the skill), plus the four agent loop that automates the process of turning traces into updated skills. It is described as its own system, not a copy of the LLM wiki, but one clearly built on the same foundational insight: memory is only useful if it compounds instead of resetting.

## Why does persistent agent memory matter right now?

The push toward persistent, evolving memory isn’t happening in isolation. It reflects a broader trend in how frontier labs are trying to get more out of agents that operate over long stretches of time without direct human guidance. Reports about OpenAI’s internal models describe systems built specifically for persistence and multi-agent collaboration, capable of running for extended periods, coordinating across sub-tasks, and continuing work without needing to be told exactly what to do at each step.

Wiki Skill is a much more modest, transparent version of that same underlying need: an agent that works well today should not have to start from zero tomorrow. Whether the mechanism is a formal wiki architecture, a skill file, or a more opaque internal memory system, the underlying problem is the same. Agents that forget everything between sessions are limited by how much can fit in a single context window. Agents that can compile experience into durable knowledge can, in principle, keep improving at a task indefinitely.

## Is the Wiki Skill approach practical to build today?

