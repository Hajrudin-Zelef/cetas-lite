---
id: collect-240926-datacamp/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants-1
title: "Course"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02"]
keywords: ["agent", "agents", "claude", "cost", "disclosure", "guardrails", "latency", "mcp", "memory", "model context protocol", "research", "sandbox"]
source: docs/RAG/clean_en/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants.md
source_anchor: ""
source_lines: [1, 96]
sha256: 57d611cacc37b104794c73197844b82e4c438fa8ee98e8859b4baa749f34ad29
---

# Course

<!-- source: https://www.datacamp.com/fr/blog/agent-harness -->

# Course

The idea isn't new. Developers have been building wrappers, scaffolds, and runtimes around models for years. The term took hold after Mitchell Hashimoto, co-founder of HashiCorp, talked about "harness engineering" in a February 2026 blog post about his AI workflow. His point was simple: when an agent makes a mistake, modify the environment so the mistake can't happen again. OpenAI picked up the term the same week for its work on Codex, and LangChain followed with the same framing.

In this article, I explain what an agent harness is, why AI agents need one, how it differs from frameworks and runtimes, and what tools developers use to build harness-type systems.

## What is an agent harness?

One definition comes from LangChain: "If you're not the model, you're the harness." In practice, an agent harness is the software that surrounds a language model: tools, memory, state, execution, guardrails, and observability.

Agent = Model + Harness

The model reasons. The harness provides it with an environment to act, remember, verify results, and follow rules.

*Model within its operational agent harness. Image by the author.*

The formula is useful, but it remains a mental model, not an industry standard. Some vendors still use "harness," "framework," and "scaffold" to mean roughly the same thing.

## Why AI agents need a harness

A raw language model quickly hits its limits when you ask it to work across many steps. It doesn't maintain durable state on its own, doesn't execute tools by itself, doesn't manage a growing context, and doesn't recover from a failed tool call without help.

Imagine an agent tasked with fixing a failing test in a Python project. Without a harness, the model can propose an apparent fix, but it can't read the actual test file, run pytest, see the real error, edit the faulty function, or confirm that the fix passes. With a harness, that entire loop becomes a few minutes of work the agent completes on its own, with each step logged for human inspection.

Anthropic's recommendation still holds: start with the simplest possible approach and only add moving parts when the task demands it.

## What an agent harness is made of

Components vary, but most share a few common building blocks. Think of them as a checklist, not a strict specification. A small agent will only need some of these elements, while a production agent will require more.

### System prompts and behavior rules

The harness generally controls the model's base instructions. This includes the system prompt, but also project rules, code standards, role constraints, and security policies. In LangChain's Deep Agents, for example, an `AGENTS.md` file can set the framework before a task begins.

Some harnesses in 2026 also use progressive disclosure of instructions. Rather than loading the full description of every tool at startup, the harness only adds a summary of what's available. A tool's detailed documentation is only loaded when the model needs it.

### Tools: how agents interact with the world

Tools allow the agent to go beyond simple text generation. Common examples: web search, file reading/writing, database queries, API calls, browser actions, code execution, and terminal commands. The harness controls which tools are available, when the model is allowed to call them, and how results are formatted and fed back into the agent's context.

Model Context Protocol (MCP) became the standard interface for this in 2026. Many harnesses, including Anthropic Agent SDK, LangChain Deep Agents, and OpenAI Agents SDK, use MCP to connect external tool servers without writing a specific integration for each one.

### Memory and state

Agents need to know what happened earlier in a task. A harness can keep short-term state in the active conversation and long-term state in files, logs, summaries, or saved preferences. Some also compress long histories into summaries to avoid overloading the context.

### Execution environment: where the agent runs and acts

Many useful agents need a real workspace: a file system, a container, an isolated terminal, a browser instance, or a cloud runtime. Without an execution environment managed by the harness, tool calls have nowhere to land.

Many harnesses now use isolated sandbox containers: ephemeral environments limited to a session, cleaned up at the end of a task, to prevent a task's file writes, package installations, and network calls from spilling over into another.

### Orchestration and planning

Some tasks don't lend themselves to a linear sequence of steps. The harness can provide a planning tool that breaks down a goal into subtasks and tracks progress. It can also launch sub-agents dedicated to part of the work and only return a summary to the main one.

LangChain Deep Agents, for example, tracks plan steps in a file system file, moving them from "pending" to "done" as execution proceeds.

### Guardrails and permissions

The harness is where you set the rules: human approval, blocking tool calls, role-based permissions, and output checks. OpenAI Agents SDK, LangChain Deep Agents, and Microsoft Agent Framework support this kind of control. The safest pattern is to check inputs, outputs, and tool permissions separately.

### Observability and tracing

When a fifty-step task fails at step thirty-seven, a trace reveals what happened. Tracing records model calls, tool calls, handoffs, errors, latency, and cost across the entire run. The OpenAI Agents SDK enables tracing by default. LangSmith adds debugging and evaluation dashboards. OpenTelemetry has become the standard for exporting traces in a neutral format, to avoid locking into a single observability tool.

## Agent harness vs framework vs runtime: what's the difference?

The question comes up often, and the answer is more subtle than some articles suggest. The taxonomy is useful, but it isn't fixed.

*Three layers, increasing abstraction from bottom to top. Image by the author.*

Let's start with the framework, which many developers have already used.

### What is an agent framework?

An agent framework gives developers building blocks to create agents. It covers model calls, tool definitions, memory patterns, and the agent loop. Examples: early versions of LangChain, CrewAI, and Google ADK. A framework tells you how to structure an agent, but not always how to run it reliably in production.

### What is an agent runtime?

An agent runtime is the layer that helps an agent run reliably over time. It handles durable execution, state persistence, retries, human-in-the-loop, and streaming. LangGraph, Temporal, and Inngest are examples. Harrison Chase offers this analogy: if Node.js is the runtime and Express is the framework, a harness looks like Next.js.

### What sets a harness apart?

A harness operates at a higher level than a framework. Where a framework provides components, a harness generally comes with more choices already made: tools, planning, filesystem access, and context management.

## Agent harness use cases: code, research, data, and enterprise

The same building blocks show up across very different jobs, but it's the combination that changes. A coding agent and an enterprise workflow agent both need a harness, but they don't use the same parts of it. These categories aren't formal standards; they're practical ways of seeing how the same idea adapts to the work to be done.

### Harness for coding agents

Coding agents are a good current example, because the harness is very visible there. To produce useful work, an agent needs file access, git context, terminal execution, test running, dependency installation, and project rules. Claude Code and Codex illustrate this pattern: they rely heavily on harness code, not just an API call to a model.

