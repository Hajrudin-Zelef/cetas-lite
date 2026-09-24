---
id: collect-240926-datacamp/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants
title: "Course"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02", "2026-04"]
keywords: ["agent", "agentic", "agents", "aws", "bedrock", "claude", "context window", "cost", "disclosure", "gemini", "guardrails", "latency"]
source: docs/RAG/clean_en/datacamp/quest-ce-quun-agent-harness-guide-pour-debutants.md
source_anchor: ""
source_lines: [1, 198]
sha256: 5d2803675779f0028e9f8bffb24db9b4aab0b02894a1c4ceb324550182699141
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

The difference between a good and an average coding harness shows up in the details: how it recovers from a failed test, whether it can undo a bad change, or how cleanly it exposes git history to the model. That's where most of the engineering effort is concentrated.

### Harness for research agents

Research agents need a different toolkit: web search, source tracking, note-taking, citation management, and synthesis. The harness handles how to store results, attribute sources, and break long documents into chunks to avoid saturating the context window all at once.

### Harness for data analysis agents

Data agents need access to datasets, SQL databases, Python environments, and schema context to know which tables and columns are available before writing queries. The harness also enforces permission boundaries, which are essential when the agent can touch production data.

### Harness for enterprise workflows

Enterprise deployments add a layer of requirements: authentication, audit logs, approval workflows, role-based access control, and connections to internal systems. AWS AgentCore is a managed example, with identity, VPC networking, and observability built in. Microsoft Agent Framework covers similar needs for teams on Azure or in .NET environments.

## Tools for building agent harness systems in 2026

A handful of products come up most often as of mid-2026. They sit at different points on the framework–runtime–harness spectrum, and the boundaries are still shifting.

### LangChain Deep Agents

LangChain Deep Agents is LangChain's open-source harness, built on LangGraph as the runtime. It includes a planning tool, a virtual filesystem, sub-agent spawning, automatic context compression, and middleware for human approval and personally identifiable information detection. Model-agnostic, it supports OpenAI-compatible endpoints and connects to sandboxes like Modal, Runloop, and Daytona for code execution.

### Anthropic Agent SDK

The Anthropic Agent SDK (package name: `claude-agent-sdk`) was extracted from Claude Code and released as a standalone option. It includes a built-in agent loop, tools for bash execution, file reading/writing, web search, MCP integration, and context compaction. It works only with Claude models, via the Anthropic API, Amazon Bedrock, Vertex AI, and Azure.

### OpenAI Agents SDK

As mentioned above, the OpenAI Agents SDK crossed the threshold from framework to harness as its features grew. The April 2026 release added native sandbox execution, memory compaction, and filesystem tools. Available in Python and TypeScript, the SDK handles tool use, handoffs between agents, and guardrails.

### Google Agent Development Kit

Google ADK supports multi-agent orchestration with built-in classes for sequential, parallel, and loop structures. It includes evaluation tools, works with Vertex AI for managed deployments, and supports MCP for connecting to tools. Available in Python, Java, TypeScript, and Go, it is optimized for Gemini models but aims to be model-agnostic.

### Microsoft Agent Framework

Microsoft Agent Framework is Microsoft's current migration path for AutoGen projects. It supports Python and .NET, works with Azure AI, and includes MCP support for tool connectivity.

### CrewAI

CrewAI takes a role-based approach to multi-agent systems. You define agents with specific roles, assign tasks, configure teams, and declare memory and guardrails. It is suited to problems that map naturally to a team of specialists.

### Temporal and Inngest

These are not agent harnesses on their own. They are durable execution platforms that handle what happens when an agent task needs to run for hours or days without losing state. On failure, the engine replays from the last successful checkpoint rather than starting over from scratch.

## Challenges and trade-offs with an agent harness

Adding a harness expands the system's capabilities, but each additional tool, permission, and agent opens a new failure path. As tasks get longer, guardrails, tracing, and durable state stop being optional and become what makes a long run recoverable.

There is also a coupling risk that surprises teams. LangChain reported a 10 to 20 point increase on a subset of tau2-bench after adding model-specific harness profiles. Artificial Analysis points in the same direction in its Coding Agent Index: coding agent results depend on the model and the harness together, with marked variations in cost, tokens, and time per task across combinations. The model hasn't changed. The prompts, tools, and middleware around it have. And that profiling is harness work.

## Do you really need an agent harness?

Here is a direct way to assess whether you need one.

You probably need a harness if your system meets one or more of these conditions:

- It needs to use external tools
- It needs to remember progress between sessions
- It needs to execute code in a real environment
- It coordinates more than one agent
- It needs to recover from partial failures without losing work
- It requires human approval

You probably don't need a harness if the task is a predictable workflow where every step is predefined.

A useful test: if the task can be handled by a single model call, or by a small deterministic script with a few conditions, a harness is probably overkill. As soon as the task requires the agent to make decisions, use tools, and react to results over time, the harness starts doing real work.

A common pitfall I see: teams adopt a harness too early, building tracing and sandboxing for what is really just one-shot text generation. The reverse mistake is more painful: wiring the model directly and then discovering, on the second failed test, the third tool call, or the fifth restart, that there is no fallback infrastructure.

## Final thoughts

As noted above, not all vendors use the same words, and the boundary between framework, runtime, and agent harness continues to evolve.

For one-shot generation, the wrapper is superfluous. For agents that need to act, remember, and recover over long sessions, the agent harness becomes a centerpiece of the system. Choosing the right harness is increasingly a decision separate from choosing the model. I'm curious to see how much of this layer will be absorbed by the next generation of models, because some announcements from OpenAI and Anthropic suggest the boundary will keep moving. The basic idea still holds: an agent is a model plus an agent harness.

To go further on building agent systems, our Building Scalable Agentic Systems course covers patterns for tool use, orchestration, and long-running agent workflows.

I am a data engineer and community builder. I work on data pipelines, cloud, and AI tools, while writing practical and impactful tutorials for DataCamp and emerging developers.

## FAQ on agent harnesses

### What is the difference between an agent harness and a system prompt?

**A system prompt is an instruction that the agent reads at the start. The agent harness is the broader layer that handles tools, state, permissions, and failure management. The simplest framing: the system prompt tells the model what to do, the agent harness controls what it can do. You can have a polished system prompt without an agent harness: you remain in a stateless API call. The agent harness is what turns a prompt into a system.**

### Can I build my own agent harness from scratch?

**In principle, yes. In its simplest form, a harness is a loop: call the model, parse the response, execute any tool calls, return the results, start again. That loop can be written in a few dozen lines of Python in an afternoon. The difficulty comes after the loop: context overflow, failed tool calls, loss of state on restart, permission enforcement, and tracing. In practice, this post-loop work always takes longer than expected, which explains why open source harnesses grow rather than shrink.**

### Does the model know it is inside a harness?

**Not explicitly. Some harnesses inform the model, via the system prompt, of the available tools, but the model has no notion of the harness as a system around it. It only sees the provided context, generates a response, and sometimes produces a tool call. Consequence: when something breaks, the model often cannot explain why, because it is unaware of the harness's existence. Debugging an agent therefore mostly means debugging the harness, not the model.**

### How does the choice of model influence which harness to use?

**More than one might think. State-of-the-art coding models are sometimes post-trained with their own agent harness in the loop, so replacing that harness can degrade performance. Practical heuristic: if your team commits to a model family, the short list of agent harnesses often imposes itself. The hardest case is changing models later: this generally involves rewriting the harness logic, not just changing a configuration value.**

### Is this different from what used to be called "LLM scaffolding"?

**Not really. It is the same idea under a more recent name. "LLM scaffolding," "agent wrapper," and "execution environment" all point in the same direction. The nuance in 2026: "scaffolding" suggests a temporary structure to be dismantled once the model is good enough, whereas "agent harness" suggests something the model keeps around it. This changes how to budget: scaffolding is removed, agent harnesses become an integral part of the system.**
