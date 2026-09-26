---
id: collect-240926-datacamp/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins-1
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "cost", "deepseek", "license", "mit license", "reasoning", "sandbox"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [1, 107]
sha256: 3dfe3e8583ed75c1524f5a99041939d231cbe0a68e2f5be6bda52a3b1fc070ff
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

