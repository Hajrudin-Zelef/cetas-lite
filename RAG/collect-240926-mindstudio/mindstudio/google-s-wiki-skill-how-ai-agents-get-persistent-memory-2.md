---
id: collect-240926-mindstudio/mindstudio/google-s-wiki-skill-how-ai-agents-get-persistent-memory-2
title: "google-s-wiki-skill-how-ai-agents-get-persistent-memory"
domain: mindstudio
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "agents", "fine-tuning", "research"]
source: docs/RAG/clean_en/mindstudio/google-s-wiki-skill-how-ai-agents-get-persistent-memory.md
source_anchor: ""
source_lines: [64, 90]
sha256: 305bcb50bb48bb9d6eaf9fdd6b6ffe8c30ac5a0ed9c7c4f61e9876b58b76e052
---

# google-s-wiki-skill-how-ai-agents-get-persistent-memory

Yes, at a conceptual level. The building blocks Wiki Skill describes (plain text logs, markdown skill files, a periodic review process) don’t require specialized infrastructure. Anyone building agent workflows can implement a simplified version: keep raw logs of what an agent does, periodically review those logs for patterns, write those patterns into a separate notes file, and only update the actual skill instructions once a pattern has been validated. The harder part is designing the review process well enough that the wiki maintainer role reliably extracts correct root causes instead of noise, which is where most of the engineering effort in a real implementation would need to go.

## Frequently Asked Questions

### What is the difference between a wiki layer and a skill in this system?

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

The wiki layer holds observations, a factual record of what the agent has learned from past attempts. The skill is the instruction set the agent follows to perform a task. The wiki describes what happened, the skill dictates what happens next.

### Who created the Wiki Skill concept?

It comes from a Google Research paper describing a system for compiling agent experience into persistent knowledge. The paper credits Andrej Karpathy’s idea of an “LLM wiki” as a partial inspiration.

### What is Andrej Karpathy’s LLM wiki idea?

It’s the concept of compiling knowledge once into a persistent, evolving reference rather than having a model rederive conclusions from scratch every time. New evidence gets added to the existing structure instead of replacing it.

### Why does the raw layer need to be immutable?

Keeping raw execution traces unmodified preserves an honest record of what actually happened. If logs could be edited, later analysis (root cause analysis, pattern detection) could be built on a distorted version of events, undermining the reliability of everything built on top of it.

### Are skill files in this system just plain text?

Based on how skills are generally implemented in current agent frameworks, yes, they are typically written as markdown, plain English text with light formatting for structure, links, and code blocks. This makes them easy to read, edit, and version compared to weights baked into a model through fine-tuning.
