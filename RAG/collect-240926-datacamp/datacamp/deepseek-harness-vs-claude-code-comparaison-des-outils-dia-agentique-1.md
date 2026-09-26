---
id: collect-240926-datacamp/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique-1
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "bedrock", "benchmark", "benchmarks", "claude", "cost", "deepseek", "license", "mcp", "memory"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [1, 119]
sha256: 1a0e2abeff4ff6b900fa06334c4e264e94a5344974d49a88965a56f512cacf8b
---

# Curriculum

<!-- source: https://www.datacamp.com/fr/blog/deepseek-harness-vs-claude-code -->

# Curriculum

Most code agent comparisons boil down to a race for speed, price, number of tools, or benchmark scores. This framing misses the real question that distinguishes these two approaches: how far can a developer go in replacing the components of the agent runtime?

DeepSeek Harness and Claude Code answer this question at different levels. Harness exposes model adapters, storage, sandboxes, and the agent loop as replaceable plugins. Claude Code embeds its own loop around Claude and provides extension points for the workflow. Here, this boundary matters more than a simple comparison of DeepSeek and Claude models.

I gave both tools the same broken repository and the same Claude model to observe what the runtime changes. A single case study is not enough to decide between the products, but it shows why the harness is not a background detail. I focus on the differences in model choice, configuration, verification, execution logs, and cost.

## Key takeaways

- If the runtime is part of the work: Harness exposes the model adapter, storage, sandbox, and agent loop as replaceable plugins, with support for multiple model providers.
- If the focus is application code: Claude Code embeds more runtime and extends its workflow via Skills, hooks, MCP, and related features.
- Test with the same model: both produced a byte-identical patch and passed the original test suite. In this Windows trial, Claude Code reported 55.0 seconds; Harness recorded 125.4 seconds after three approval requests.
- Cost: DeepSeek Harness has no license fee; model usage is billed by the chosen provider, with possible infrastructure costs. Claude Code is included in paid Claude plans.
- Simple rule: start with Claude Code for everyday application work. Choose Harness when modifying or inspecting the runtime is part of the mandate.

## Introduction to AI agents

## DeepSeek Harness vs Claude Code: quick comparison

| Dimension | DeepSeek Harness | Claude Code | 
|---|---|---|
| What it is | Agent runtime with replaceable components | Packaged code agent | 
| License and status | MIT, developer preview, no stable version | Proprietary, production product | 
| Models | DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure, local | Claude, direct or via Bedrock and Vertex | 
| Replaceable loop | Yes, substitutable plugin extension point | No, extensions graft around it | 
| Unit of extension | Cordis plugin, any runtime capability | Plugin grouping skills, hooks, agents, MCP | 
| Execution history | Append-only typed event log, replayable | JSONL transcript, checkpoints, OpenTelemetry | 
| Surfaces | Local web interface, headless CLI, Python SDK | Terminal, IDE, desktop, web, Slack, CI | 
| Configuration | Requires provider setup and mastery of profiles | Lighter default configuration; permissions, hooks, MCP, and project optional | 

## What is DeepSeek Harness?

DeepSeek Harness is an open-source agent harness currently in developer preview. It provides the runtime in which models use tools and context, and it can run models from providers other than DeepSeek. Its README warns that updates may break existing configurations.

If you want to try it, start with our DeepSeek Harness tutorial.

### How DeepSeek Harness works

The agent loop is implemented as a plugin. Harness uses Cordis to compose model adapters, tools, sessions, storage, sandboxes, and the loop.

Cordis allows plugins to discover services and exchange events. Dependency changes load or unload plugins; hot reloading cleans up stale listeners and background tasks.

The default profile exposes the loop as a configuration. `dsh --profile headless --dump-default-config` includes this entry:

```
- id: agent-loop
  name: '@deepseek-ai/dsh-agent-loop'
  config:
    agents: []
```
This entry can be replaced via `cordis.patch.yml`, proof that the loop itself is not fixed.

Plugin panel showing runtime components. Video by the author.

### The four execution modes

DeepSeek Harness offers four distinct execution modes.

- **Standard**: full code agent mode
- **PTC** /**Code**: combines multiple tool calls into a single TypeScript program
- **Minimal**: keeps only persistent bash and str_replace_editor
- **Creator**: for inspecting and testing the runtime

DeepSeek used *Minimal* mode for its benchmarks.

## What is Claude Code?

Claude Code is Anthropic's proprietary code agent and an agentic harness for local or managed use via the terminal, IDE, desktop, web, Slack, and CI.

The CLI launches with `claude` in the project directory. This directory becomes the default file scope, and the session reads project instructions from `CLAUDE.md`. Users can expand file access or add external services afterward.

Anthropic presents these interfaces as ways to access Claude Code. Local sessions run on your machine. Cloud sessions run in managed environments or on servers operated by your organization. Remote Control allows you to drive local work from a browser.

The best starting points are our Claude Code tutorial and the Claude Code best practices guide.

### Built-in loop and extensions

Anthropic describes Claude Code as the agentic harness around Claude, with a repeated loop of context, action, and verification. Users can steer it during execution.

Claude Code stores sessions locally and compacts old context as the window fills. `CLAUDE.md` and automatic memory retain selected instructions and project details from one session to the next. Subagents use separate context windows and return summaries to the parent session.

### How Claude Code's extension layer works

Claude Code supports several extension mechanisms. `CLAUDE.md`, skills, hooks, MCP, subagents, plugins, Agent Teams, and the Agent SDK all extend its workflow.

These extensions operate around Claude Code's built-in loop. Hooks can enforce tool-call rules, while the Agent SDK provides tools and context management in code.

## DeepSeek Harness vs Claude Code: architecture and control

DeepSeek Harness exposes lower-level, replaceable runtime elements. Claude Code keeps its built-in loop and supports extensions around it.

### Replaceable runtime vs packaged agent

DeepSeek Harness treats the runtime as configurable infrastructure: model adapters, storage, sandboxes, and the loop can be replaced via profile entries. Claude Code keeps its built-in loop fixed and extends the workflow around it.

Harness can use this mechanism to swap model providers or other parts of the runtime. The word "plugin" does not have the same scope here:

- Claude Code plugins aggregate features around the loop.
- DeepSeek Harness plugins can define parts of the runtime itself.

### Model support and provider choice

DeepSeek Harness supports more model providers. It runs DeepSeek, Anthropic, OpenAI, clouds, and compatible local endpoints. Claude Code runs… Claude. In other words, DeepSeek Harness can host Claude, but Claude Code cannot host DeepSeek.

This is what allowed me to keep the same model during testing. Comparing a DeepSeek model in Harness to Claude in Claude Code changes two variables at once.

Provider choice involves configuring credentials and the endpoint, including exact model IDs. Claude Code controls the model family and most request parameters.

Switching providers does not guarantee identical behavior. Models can differ in tooling format, context size, or reasoning controls. Harness keeps plugin configuration accessible, but the provider adapter must remain compatible with the chosen endpoint.

### Session logs and traceability

DeepSeek Harness records prompts, context injections, tool calls, and permission decisions in an append-only event stream. Its "trajectory" view shows the origin of each record and supports resume, fork, search, and replay.

Claude Code stores JSONL transcripts, supports resume and fork, and emits OpenTelemetry traces.

