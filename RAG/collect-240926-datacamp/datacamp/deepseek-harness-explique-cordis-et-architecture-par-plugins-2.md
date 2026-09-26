---
id: collect-240926-datacamp/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins-2
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["agent", "agents", "benchmark", "benchmarks", "claude", "cost", "deepseek", "mcp", "memory", "open source", "parameters", "reasoning"]
source: docs/RAG/clean_en/datacamp/deepseek-harness-explique-cordis-et-architecture-par-plugins.md
source_anchor: ""
source_lines: [108, 226]
sha256: f2f84646cf21a55bcc63cfdcf97aed771705587015f90e51feb77a5c082b3beb
---

# Curriculum

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

