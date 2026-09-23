---
id: collect-mindstudio/mindstudio/build-custom-sub-agents-claude-code-yaml
title: "How to Build Custom Sub-Agents in Claude Code: YAML, Tools, and Triggers"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Google"]
dates: ["2026-06", "2026-09-23"]
keywords: ["agent", "agents", "claude", "mcp"]
source: docs/RAG/Collect RAG/02_mindstudio/build-custom-sub-agents-claude-code-yaml.md
source_anchor: ""
source_lines: [1, 58]
sha256: 1d937c582cc1d6ee72d94bbf809df19cd137f7299268400c1a2b2d90a80207c3
---

# How to Build Custom Sub-Agents in Claude Code: YAML, Tools, and Triggers

## Metadata

- **Source**: https://www.mindstudio.ai/blog/build-custom-sub-agents-claude-code-yaml
- **Site**: MindStudio
- **Type**: Article
- **Language**: en
- **Verification status**: ✅ reachable
- **Collection date**: 2026-09-23

## Full summary

This technical guide explains how to build custom sub-agents in **Claude Code**. Sub-agents are markdown files with YAML front matter that let users define specialized agents — each with its own tools, instructions, and scope — that the main orchestrator can delegate to automatically. This solves the problem of a single agent becoming unwieldy: switching contexts constantly, losing track of specialized concerns, and blending responsibilities that should stay separate.

**File structure.** Sub-agents are plain markdown files. Project-level agents live in `.claude/agents/` inside the repo; user-level agents live in `~/.claude/agents/` (available across all projects). Both can coexist; project-level takes precedence on name conflicts. Each file has two parts: YAML front matter (between `---` delimiters) and the agent's system prompt in the markdown body.

**YAML front matter fields.** `name` — how the agent is referenced (lowercase, hyphens, explicit; invoked directly with `/agent:security-auditor`). `description` — the most important field, driving automatic delegation. It's a matching signal, not a label. Weak: "Handles security stuff." Strong: "Reviews code for security vulnerabilities including SQL injection, XSS, authentication flaws, insecure dependencies, and secrets exposure. Use this agent when analyzing code changes for security risks or performing security audits." Descriptions should name specific vulnerability types and state trigger conditions ("use this agent when…"). `tools` — comma-separated list of capabilities: Read, Write, Edit, Bash, Grep, Glob, WebFetch, Task (spawn sub-agents, allows nesting), TodoRead/TodoWrite, and `mcp__<server>__<tool>` from connected MCP servers. `model` (optional) — specify which model the agent uses (e.g., `claude-opus-4-5`); omit to inherit session default.

**Tool configuration by agent type.** Read-only agents (code reviewers, security auditors, documentation auditors, linters): `tools: Read, Grep, Glob` — see everything, touch nothing. Write-capable agents (documentation writers, test generators, formatters): `Read, Write, Edit, Grep, Glob` — no Bash unless needed. Execution agents (test runners, build scripts): `Read, Bash, Glob` — narrow in other ways. Orchestrator agents (nested structures): `Task, Read, Glob` — Task enables delegation to sub-agents.

**System prompt body.** Standard markdown acting as the instruction set. Write prompts that are: role-specific (state what the agent is and isn't responsible for), explicit about output format (e.g., for each issue: file path/line, description, suggested fix, severity low/medium/high), and bounded in scope ("focus only on the changed files; do not audit the entire codebase"). Description gets the agent invoked; the system prompt determines behavior once running.

**Triggering.** Two ways: **Explicit invocation** — type `/agent:security-auditor review the changes in src/auth/`, bypassing orchestrator routing. **Automatic invocation** — Claude Code's orchestrator matches the task against description fields. For reliability: use keywords users will naturally say; state conditions explicitly; avoid overlapping descriptions; test with realistic prompts. Claude Code has no formal trigger field — triggering is entirely driven by the description field.

**Practical examples.** The article provides four full configurations: Code Reviewer (Read, Grep, Glob; model claude-opus-4-5), Test Writer (Read, Write, Edit, Grep, Glob), Docs Writer, and Dependency Auditor (Read, Bash, Glob — runs npm audit, pip-audit).

**Multi-agent workflow patterns.** Sequential delegation (code reviewer → test writer → docs writer, each dependent on the previous); parallel delegation (independent tasks like security audit and documentation update run simultaneously — the orchestrating agent needs Task); nested agents (sub-agents spawning sub-agents via Task, but more than 2-3 levels becomes hard to debug).

**Extending with MindStudio.** A real limitation of Claude Code sub-agents is confinement to the shell environment and connected MCP servers. The MindStudio Agent Skills Plugin (an npm SDK) lets agents call 120+ typed capabilities as simple method calls (`agent.sendEmail(...)`, `agent.searchGoogle(...)`, `agent.runWorkflow(...)`), handling rate limiting, retries, and authentication. Connected via MCP server config so tools appear in the YAML tools field like native tools.

**Common mistakes.** Vague description fields (no basis for orchestrator routing); overlapping scopes (unpredictable routing); over-provisioning tools (especially Bash and Write, causing unintended side effects); system prompt scope creep (split into two agents); missing agent directory (verify `.claude/agents/` path). No documented hard limit on sub-agent count, but 4–8 focused agents is usually enough. Sub-agents run as separate contexts — they don't automatically see full conversation history. Files should be committed to version control so teams share agent configs.

## Key points

- Claude Code sub-agents are markdown files with YAML front matter in `.claude/agents/` (project) or `~/.claude/agents/` (user-level, available everywhere).
- The `description` field drives automatic triggering — write it with specific keywords and explicit trigger conditions, not generic labels.
- The `tools` field should be as narrow as possible: Read-only vs write-capable vs execution vs orchestrator tool sets.
- The system prompt body defines role, output format, and scope; it works together with the description field.
- Invocation is explicit (`/agent:name`) or automatic (orchestrator matches description); there's no separate trigger field.
- Workflow patterns: sequential, parallel (needs Task tool), and nested delegation (max ~2-3 levels).
- The MindStudio Agent Skills Plugin (npm SDK) adds 120+ typed capabilities (email, Google search, workflows) to Claude Code sub-agents via MCP.
- Vague descriptions, overlapping scopes, and over-provisioned tools are the three most common failure modes.

## Technical data / figures

- File locations: `.claude/agents/` (project-level, takes precedence) and `~/.claude/agents/` (user-level).
- YAML front matter fields: name, description (trigger), tools, model (optional).
- Available tools: Read, Write, Edit, Bash, Grep, Glob, WebFetch, Task, TodoRead, TodoWrite, mcp__<server>__<tool>.
- Tool sets: read-only `Read, Grep, Glob`; write-capable `Read, Write, Edit, Grep, Glob`; execution `Read, Bash, Glob`; orchestrator `Task, Read, Glob`.
- Explicit invocation: `/agent:name`.
- Recommended agent count: 4–8 focused agents per project.
- MindStudio Agent Skills Plugin: 120+ typed capabilities, npm SDK, MCP-server connected.
- Model example: `claude-opus-4-5`.

## Why this source matters for the RAG

Provides detailed, current (June 2026) technical documentation for building Claude Code sub-agents — exact file format, YAML fields, tool lists, triggering semantics, and workflow patterns. This is precise, actionable reference knowledge that enables accurate answers on Claude Code agent configuration and multi-agent workflows without hallucinating on syntax or file paths.
