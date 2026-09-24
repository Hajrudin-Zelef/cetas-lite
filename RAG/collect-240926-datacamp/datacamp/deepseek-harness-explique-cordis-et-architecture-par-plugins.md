---
id: collect-240926-datacamp/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "agents", "benchmark", "benchmarks", "claude", "cost", "deepseek", "inference", "license", "mcp", "memory", "mit license"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [1, 255]
sha256: c59f6216379ca8c7652ccab5032526bb9fd67b3624f21c092e262600ea5281e1
---

# Curriculum

<!-- source: https://www.datacamp.com/fr/blog/what-is-deepseek-harness -->

# Curriculum

DeepSeek Harness is designed to execute a task, not just answer a question. It is an open-source agent runtime environment that connects a model to your repository, your terminal, your tools, and session history. Ask it to fix a bug: it can inspect files, modify code, run tests, and react when a command fails. A simple model call cannot do all of that on its own.

The most original part lies beneath this workflow. DeepSeek Harness exposes the model adapter, tools, sessions, sandbox, and even the agent loop as plugins coordinated by Cordis. The model is just one of the agent's components, not the product itself.

This is not finished software. Harness is still in developer preview; its APIs may change between versions, and its own security notice states that no audit has been performed. I will address these limitations alongside the architecture and the differences from Claude Code, Codex, and OpenCode.

## In brief

- **What it is:** DeepSeek Harness is an open-source agent runtime environment, not a model. It provides the model with tools, sessions, a sandbox, and an agent loop.
- **Core design:** Cordis exposes the model adapter, tools, session store, sandbox, and agent loop as interchangeable plugins.
- **Sessions:** An append-only event log enables resumption, forking, search, replay, and the Trajectory view.
- **Modes:** Standard, PTC, Minimal, and Creator change the tools accessible to the agent and how to access them.
- **Main difference:** DeepSeek Harness allows developers to replace low-level components that Claude Code, Codex, and OpenCode leave fixed.
- **Main limitation:** It remains a developer preview without a security audit, and its APIs may evolve between versions.

## Introduction to artificial intelligence agents

## What is DeepSeek Harness?

DeepSeek Harness, abbreviated as `dsh`, is an open-source agent harness from DeepSeek AI under the MIT license. It sits between a language model and the outside world, providing tools, sessions, a sandbox, and the loop that drives the task forward.

At DeepSeek, the equation is **"Agent = Model + Harness."** The model handles reasoning and generation. The harness is everything that allows that reasoning to act on a real file system and continue without you having to re-explain the task at every step.

It relies on Cordis, a plugin framework that predates DeepSeek Harness. Cordis makes it possible to replace these building blocks independently through configuration. I will return later to the cost of this choice.

With this framework in mind, here are two common misconceptions.

### DeepSeek Harness is not an AI model

As stated above, the model and the runtime environment are two distinct layers. This separation makes it possible to change providers without touching the tools or session settings. The same runtime can use DeepSeek, Anthropic, OpenAI, or an OpenAI-compatible endpoint.

### DeepSeek Harness goes beyond a coding assistant

Standard mode may give the impression of a development assistant, but that is only one configuration. As we will see, Minimal and Creator modes change the tools accessible to the agent. Building a new configuration always requires engineering work; developers have access to the parts.

## How Cordis organizes DeepSeek Harness plugins

As mentioned, Cordis is the plugin framework beneath DeepSeek Harness. It allows each component to request a service without tying the code to a single provider.

Cordis comes from the Koishi chatbot ecosystem and was built by a developer known as Shigma; DeepSeek distributes and extends it. The authors describe the design in their article A Programming Paradigm for Spatiotemporal Composability.

These foundations lead to the project's main slogan and two Cordis concepts. The names sound academic, but the operation remains simple.

### "Everything is a plugin"

DeepSeek's architecture documentation explains that you extend `dsh` by mounting a plugin alongside the others. Model adapters, tools, sessions, sandboxes, storage, scheduling, the agent loop, and the interface are all plugins.

Taken literally, the slogan goes too far. Cordis remains beneath the plugins. It loads and unloads them, checks their dependencies, and drives the events they use to communicate. Cordis is indispensable; it is not just another optional piece.

### Spatial composability manages dependencies between plugins

A plugin declares the services it needs without requiring a hand-written startup sequence. It activates when those services exist and deactivates if a required service disappears. Its dependencies dictate when it can run.

DeepSeek calls this spatial composability. Dependencies tell Cordis where a component fits, saving developers from ordering startup by hand.

### Temporal composability cancels plugin effects

Cordis also tracks registrations such as event listeners, prompt sections, and tool schemas. Removing a plugin removes these effects instead of leaving orphaned listeners. This does not undo an external action like a shell command; reversibility only applies to effects tracked by Cordis.

## DeepSeek Harness architecture: how the runtime assembles

A running instance is a plugin tree built from settings loaded in a defined order. These settings determine which building blocks are active.

Cordis connects all replaceable plugins of the runtime. Image by the author.

### Cordis services allow plugins to discover each other

Cordis provides a shared service directory. Plugins use stable keys like `ctx.tools`, `ctx.llm`, and `ctx.sessions` instead of importing a provider's code. A tool that calls `ctx.llm` doesn't need to know which model adapter is behind it.

### Agent presets and execution profiles drive different layers

If everything is replaceable, you still have to decide what to mount for a given execution, and DeepSeek Harness answers this at two levels that are easy to confuse.

Short version: a profile controls how the program starts, while a preset controls what the agent can do. If you only use the web application, you can skip the next two subsections.

#### Execution profiles

An execution profile (`web`, `headless`, `sdk`, `sdk-minimal`, and `acp` are provided as templates) decides how the application launches and which Cordis plugin bundles are stacked at startup. Most readers will only touch this level by launching `dsh web` or a similar command.

#### Agent presets

An agent preset (Standard, PTC, Minimal, or Creator) decides what an active session can use. A patch file can change the preset without modifying Harness's source code.

### The agent loop coordinates turns, steps, and tool calls

DeepSeek distinguishes a step from a turn. A step is a model request plus its tool calls. A turn groups zero or more steps: it opens before its first input is taken up and closes when nothing more is due. Most turns run several steps before the agent responds, but a rejected input closes a turn without consuming a step.

A turn can contain several steps. Image by the author.

### Sessions use an append-only event log

This is, in my view, the most important element. A session is an append-only log of typed events, not an array of chat messages. Harness reconstructs the model's history from this log, and the session documentation requires that everything sent to the model be retrievable from it.

Resumption, forking, search, replay, and the Trajectory view all rely on this event stream.

Reconstructing history is not a deterministic re-execution. The model's output and the external state may differ, but the log still provides a searchable trace of what happened.

Session history is an append-only log. Image by the author.

### How DeepSeek Harness controls tools and sandboxes

A model can request a tool by name, but it cannot execute it directly. Two distinct controls stand between the request and a modification of the file system.

#### The tool execution pipeline

The call goes through a policy check, execution, then result processing. The model chooses the tool; the runtime decides whether and how it runs.

The runtime decides how tools execute. Image by the author.

#### Sandbox versus approvals

- **Approval** asks whether the user must confirm an action.
- **Sandbox** limits where and how the action executes.

DeepSeek separates them, even though permission presets group both controls together, much like a container runtime that separates process permissions and execution limits.

Worth flagging now, since I'll come back to it in the limitations: telling a model in a system prompt to "only read files" is an instruction it can choose to follow, not a barrier imposed the way an OS-level sandbox restriction would.

## DeepSeek Harness modes: Standard, PTC, Minimal, and Creator

DeepSeek Harness offers four modes. None is "better" than the others. They are four answers to "how much of the runtime should be exposed to this session," and the right choice depends on the task. As the architecture showed, each mode changes the set of tools available to the agent.

Four modes, one common runtime base. Image by the author.

### Standard mode

The versatile foundation:

- File editing
- Shell access
- File and web search
- Skills
- Planning
- Goals
- Sub-agents
- Workflows

For everyday repository work, this is my starting point.

### PTC mode

PTC mode keeps almost all of Standard's tooling but changes how the model accesses it. (Since version 0.1.2, Web PTC mode no longer exposes the generic `workflow` tool by default.)

Instead of calling tools one by one over several steps, the model writes a program against a generated SDK. This program can invoke multiple tools via `run_code`. Every call still goes through the same policy checks: PTC changes how the plan is expressed, not what the model is allowed to do.

The product page still uses the "Code mode" label, but a newer official version renamed it "PTC mode" while keeping old conversations readable. I'll use "PTC mode" throughout; the FAQ revisits what these initials might stand for.

### Minimal mode

Minimal mode reduces the environment to two tools: a persistent shell and a file editor via string substitution. DeepSeek uses it for model benchmarks, because results depend partly on the harness, not just the model's weights.

### Creator mode

Creator mode lets developers inspect the runtime and test Cordis plugins in memory. It's used to build presets; I wouldn't call it "self-improving" in any deeper sense.

## What sets DeepSeek Harness apart from other agent frameworks

DeepSeek Harness stands out from many agent frameworks by making the low-level runtime layers replaceable. I could have integrated it into the architecture, but the nuance is easy to miss. Cordis handles these changes via a unique plugin system.

You can modify how the agent itself works, not just the tools it calls. The event log also creates an execution inspectable by developers, rather than a simple chat transcript. The Minimal and Creator modes then allow testing the runtime from two opposite angles.

## DeepSeek Harness vs Claude Code, Codex and OpenCode

A feature checklist would miss the essential point. Every competitor supports extensions; the real question is: which parts can developers change? The nuance seems minor, but it isn't. Our dedicated Harness vs Claude Code comparison uses the same model on both sides and covers configuration, logs and cost.

### DeepSeek Harness vs Claude Code

Claude Code supports project instructions, skills, hooks, MCP, sub-agents and an Agent SDK, while keeping its internal loop fixed. DeepSeek Harness allows replacing the loop, the model adapter and the storage layer via configuration.

### DeepSeek Harness vs Codex

Codex requires a finer comparison, since its CLI and App Server are also open source. It provides an agent harness that developers extend via documented entry points. DeepSeek Harness is designed to modify the runtime itself. The levels of control offered differ.

### DeepSeek Harness vs OpenCode

OpenCode is already open source, works with multiple model providers and uses a client-server architecture. You can configure its tools, permissions, sessions and providers. Its plugins extend a fixed server core, while DeepSeek also makes the loop and the session store replaceable.

## When to use DeepSeek Harness

Replacing parts of the runtime has no value in itself. This extra control only matters if it solves a problem you already have.

- **When the runtime is part of the project.** If you modify model adapters, the agent loop, storage or session behavior — not just if you build on top of an agent —, it's a better choice.
- **When you compare models in a controlled environment.** Using the same runtime fixes more parameters during model swaps, even if the models still differ in tool usage and reasoning style.
- **When debugging a complex execution is crucial.** The session event log and the Trajectory view make it easier to reconstruct what the model saw and which tools ran.
- **When you test the internals of agents.** Creator mode and Cordis are aimed at developers who study agent composition, more than those who just want to generate application code.

This can be overkill for simple model calls or for teams that want a ready-to-use coding agent without touching its internals. Replacing more parts is only worth the effort if that control solves a real problem.

## Limitations of DeepSeek Harness: preview status and security risks

All the architecture above matters little without a clear view of its current limitations.

### It's still a developer preview

The DeepSeek repository clearly states that breaking changes will occur. It has already happened: the renaming of Code to PTC came with session API changes and the removal of an optional SQLite option. Pin your versions. Skipping this step in hopes of configuration stability is not a strategy.

### More control, more complexity

Making more layers replaceable means more to learn: plugin dependencies, settings, provider differences, version compatibility. This is the usual trade-off between convenience and control.

### Is DeepSeek Harness local?

DeepSeek Harness stores session content, tool traces and settings locally by default, in accordance with its data processing statement. You can disable its anonymous reporting of settings and project lists.

But an external model provider, a web tool, an MCP server or a plugin can still send data off your machine according to their own policy. "Local-first" doesn't cover every service you connect.

### Running agents carries security risks

A runtime capable of editing files, executing commands and loading third-party plugins can cause real damage. DeepSeek's security notice states that no audit has been conducted. Sandboxes, approvals and permission controls reduce risk without guaranteeing isolation.

Running the software on your own machine doesn't eliminate this risk. Use limited permissions and a disposable environment for untrusted tasks, and be vigilant with content that may contain hidden instructions.

## Why an agent's behavior depends on more than the model

An agent's behavior depends on the runtime as much as on the model. We return to "Agent = Model + Harness," and this separation holds for LLM agents beyond DeepSeek.

What a model can produce depends on its weights. What an agent does also depends on the context passed to the model, the actions allowed and the degree of constraint in the execution. None of that resides in the weights.

DeepSeek Harness highlights this surrounding layer by breaking it down into named, replaceable components. Minimal mode illustrates why this goes beyond DeepSeek: a benchmark score partly reflects the harness used for the test, not just the model. The harness does not make a model "smarter." It changes the framework in which it operates.

## Conclusion

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
