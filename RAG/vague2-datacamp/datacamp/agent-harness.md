---
id: vague2-datacamp/datacamp/agent-harness
title: "Qu'est-ce qu'un agent harness ? Comment les agents d'IA obtiennent des outils, de la mémoire et du contrôle"
domain: datacamp
role: reference
task: article
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2026-02", "2026-04", "2026-09-23"]
keywords: ["agent", "agents", "aws", "bedrock", "claude", "cost", "disclosure", "guardrails", "latency", "mcp", "memory", "open source"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/agent-harness.md
source_anchor: ""
source_lines: [1, 53]
sha256: 3e14713bd99b1c4ca02d9a8556917aa8a8693410d78111cbba3d8d2b01514044
---

# Qu'est-ce qu'un agent harness ? Comment les agents d'IA obtiennent des outils, de la mémoire et du contrôle

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/agent-harness
- **Site** : DataCamp
- **Type** : Article (guide pour débutants)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp beginner's guide explains the **agent harness**: the software surrounding a language model that provides tools, memory, state, execution, guardrails, and observability. The formula is **Agent = Model + Harness**: the model reasons, the harness gives it an environment to act, remember, verify results, and follow rules. The term gained traction after HashiCorp co-founder Mitchell Hashimoto wrote about "harness engineering" in February 2026 (when an agent makes an error, change the environment so it can't recur); OpenAI adopted it for Codex the same week, and LangChain followed.

**Why agents need a harness:** a raw LLM doesn't maintain durable state, execute tools, manage growing context, or recover from a failed tool call. Example: fixing a failing Python test requires reading the real test file, running pytest, seeing the error, editing the faulty function, and confirming the fix — all of which a harness enables and logs. Anthropic's advice: start with the simplest approach and add moving parts only when the task demands it.

**Components** (a checklist, not a spec): system prompts and behavior rules (including project rules, `AGENTS.md`, and progressive instruction disclosure to avoid loading every tool description upfront); tools (web search, file I/O, DB queries, APIs, browser actions, code/terminal execution, with MCP as the 2026 standard interface); memory and state (short-term in conversation, long-term in files/logs/summaries, with compaction); execution environment (file system, container, terminal, browser, cloud runtime, often sandboxed ephemeral containers); orchestration and planning (planning tools, subagents, e.g., LangChain Deep Agents tracks plan steps in a file); guardrails and permissions (human approval, tool blocking, RBAC, output checks — verify inputs, outputs, and tool permissions separately); and observability/tracing (model calls, tool calls, handoffs, errors, latency, cost; OpenAI Agents SDK traces by default, LangSmith adds dashboards, OpenTelemetry is the neutral export standard).

**Harness vs framework vs runtime:** a *framework* (early LangChain, CrewAI, Google ADK) gives building blocks and how to structure an agent; a *runtime* (LangGraph, Temporal, Inngest) helps run it reliably over time via durable execution, state persistence, retries, human-in-the-loop, and streaming; a *harness* operates at a higher level, shipping with more decisions pre-made (tools, planning, file system access, context management). Harrison Chase's analogy: if Node.js is the runtime and Express the framework, a harness is like Next.js. The taxonomy is useful but not fixed.

**Use cases:** development agents (file access, git context, terminal, tests, dependency install — Claude Code and Codex rely heavily on harness code); research agents (web search, source tracking, note-taking, citations, document chunking); data analysis agents (datasets, SQL, Python, schema context, permission boundaries for production data); enterprise workflows (auth, audit logs, approval circuits, RBAC, internal systems — e.g., AWS AgentCore, Microsoft Agent Framework).

**Tools in 2026:** LangChain Deep Agents (open source, on LangGraph; planning, virtual file system, subagents, context compression, HITL middleware; sandboxes via Modal, Runloop, Daytona), Anthropic Agent SDK (`claude-agent-sdk`, extracted from Claude Code; Claude-only via Anthropic API, Bedrock, Vertex, Azure), OpenAI Agents SDK (added native sandbox execution, memory compaction, file system tools in April 2026; Python/TypeScript), Google ADK (multi-agent orchestration, evaluation, Vertex AI, MCP; Python/Java/TypeScript/Go), Microsoft Agent Framework (AutoGen migration path; Python/.NET; Azure AI), CrewAI (role-based), and Temporal/Inngest (durable execution platforms, not harnesses themselves).

**Challenges:** each added tool/permission/agent is a new failure path; guardrails, tracing, and durable state become mandatory for long runs. Coupling risk: LangChain reported a 10–20 point improvement on a tau2-bench subset after adding model-specific harness profiles; Artificial Analysis's Coding Agent Index shows results depend on model *and* harness together.

**Do you need one?** Likely yes if your system uses external tools, remembers across sessions, executes code, coordinates multiple agents, recovers from partial failures, or needs human approval. Probably not for predictable pre-defined workflows or one-shot generation. A common mistake is adopting a harness too early; the opposite (no fallback infrastructure) is more painful.

## Key points

- Agent = Model + Harness; the harness provides tools, memory, state, execution, guardrails, observability.
- Term popularized by Mitchell Hashimoto's Feb 2026 "harness engineering" post, adopted by OpenAI and LangChain.
- Components: system prompts, tools/MCP, memory, execution environment, orchestration, guardrails, tracing.
- Framework (structure) vs runtime (reliable execution) vs harness (higher-level, pre-made choices).
- Use cases span dev, research, data, and enterprise agents.
- 2026 tools: LangChain Deep Agents, Anthropic Agent SDK, OpenAI Agents SDK, Google ADK, Microsoft Agent Framework, CrewAI, Temporal/Inngest.
- Performance depends on model + harness together (LangChain +10–20 tau2-bench; Artificial Analysis).
- Harness debugging is mostly harness debugging, since the model doesn't perceive the harness.

## Technical data / figures

| Layer | Purpose | Examples |
|---|---|---|
| Framework | Building blocks, agent structure | Early LangChain, CrewAI, Google ADK |
| Runtime | Reliable execution over time | LangGraph, Temporal, Inngest |
| Harness | Higher-level, pre-made choices | Claude Code, Codex, Deep Agents |

Harness components: system prompts/rules, tools (MCP standard), memory/state, execution environment (sandboxes), orchestration/planning, guardrails/permissions, observability/tracing (OpenTelemetry). Notable figures: LangChain +10–20 points on a tau2-bench subset with model-specific harness profiles; `AGENTS.md` config; `claude-agent-sdk` package name.

## Why this source matters for the RAG

It is a clear conceptual reference defining agent harnesses and distinguishing them from frameworks and runtimes, with current tooling and trade-offs. It is central for RAG queries about agent architecture, MCP, agent memory, guardrails, and why agents need more than a model.
