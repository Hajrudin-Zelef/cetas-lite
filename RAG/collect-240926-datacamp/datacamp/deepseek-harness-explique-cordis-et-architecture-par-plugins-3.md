---
id: collect-240926-datacamp/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins-3
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "cost", "deepseek", "inference", "reasoning", "tool calling"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [227, 255]
sha256: 316d74f902eded4f0d53c16731d7b1631ebbcd881627c2ba00beadcee7ce6085
---

# Curriculum

The opening sentence is worth remembering: the model reasons, but the runtime decides what that reasoning can access and what it can do. DeepSeek Harness makes that runtime modifiable, from the model adapter and tools all the way to the session store and the agent loop.

This control comes at a cost. Replacing more building blocks means taking on more configuration, version changes, and security boundaries. A developer preview with shell access is not a "set it and forget it" tool.

My opinion is simple: use DeepSeek Harness when the runtime is part of the work. If you only need changes in a repository, a ready-to-use coding agent will ask less of you.

Our DeepSeek Harness tutorial covers configuration. The guide on Claude Code alternatives compares other coding agents, while Introduction to AI Agents revisits the basics assumed known here.

## DeepSeek Harness FAQ

### Is DeepSeek Harness the same thing as a DeepSeek model?

**No, the model and the runtime are distinct. Harness does not include a model's weights and does not run inference itself; it sends requests to DeepSeek, Anthropic, OpenAI, or a local model.**

### Is using DeepSeek Harness free?

**The software itself is free and MIT-licensed. What is not free is the model provider you connect to, since inference is billed separately by the model operator, as well as any infrastructure costs related to sandboxes or external services you add on top.**

### What does the acronym PTC actually mean?

**DeepSeek's release notes use "PTC mode" without giving an official expanded form, even though the behavior matches "programmatic tool calling." I would treat that as a working definition, not a confirmed acronym, until DeepSeek explicitly specifies one.**

### Can I entrust DeepSeek Harness with a repository that matters to me?

**Some limits remain. For an important repository, work on a copy or a separate branch, keep production credentials out of the environment, and review each plugin before loading it.**

### Does "everything is a plugin" mean I can turn it into any type of agent?

**Not without real engineering work. Replacing the model adapter or the agent loop requires a plugin that respects the appropriate service contract. The plugin system gives you access to more parts; it does not make the work disappear.**
