---
id: collect-240926-mindstudio/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers-1
title: "how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "mcp", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers.md
source_anchor: ""
source_lines: [1, 172]
sha256: 538a13abc5334ed2c1fb1ad8a7ae4140df504af28a04e84e03a98892353dabf9
---

# how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers

<!-- source: https://www.mindstudio.ai/blog/build-custom-sub-agents-claude-code-yaml -->

## What Claude Code Sub-Agents Actually Are

If you’ve been using Claude Code for any serious project work, you’ve probably hit the point where a single agent trying to do everything becomes unwieldy. It switches contexts constantly, loses track of specialized concerns, and blends responsibilities that should stay separate.

Sub-agents in Claude Code solve this. They let you define specialized agents — each with its own tools, instructions, and scope — that the main orchestrator can delegate to automatically. And the entire configuration lives in a plain markdown file with YAML front matter.

This guide covers how to build custom sub-agents in Claude Code: the file format, how to write `description` fields that trigger agents reliably, how to restrict tool access, and how to structure multi-agent workflows that actually hold up in practice.

## Understanding the Sub-Agent File Structure

Sub-agents in Claude Code are just markdown files. Each file represents one agent. The structure is straightforward:

```
your-project/
└── .claude/
    └── agents/
        ├── code-reviewer.md
        ├── test-writer.md
        └── security-auditor.md
```
For agents you want available across all your projects, put them in your home directory instead:

```
~/.claude/agents/
├── code-reviewer.md
└── documentation-writer.md
```
Project-level agents (in `.claude/agents/`) are specific to that repo. User-level agents (in `~/.claude/agents/`) follow you everywhere. Both directories can coexist, and project-level agents take precedence if there’s a name conflict.

The file itself has two parts: the YAML front matter at the top (between `---` delimiters), and the agent’s system prompt in the markdown body below.

```
---
name: code-reviewer
description: Reviews code changes for quality, style, and correctness
tools: Read, Grep, Glob
---
You are a thorough code reviewer. Focus on logic errors, edge cases, 
and maintainability. Be specific about what needs to change and why.
```
## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

That’s the whole format. Simple, but each field does meaningful work.

## Writing YAML Front Matter That Works

The front matter has a handful of fields. Some are required, some are optional, and one — `description` — is the most important thing you’ll configure.

### The `name` Field

The `name` is how you reference the agent explicitly. Keep it lowercase, use hyphens instead of spaces, and make it obvious what the agent does.

`name: security-auditor`
You can invoke this agent directly by typing `/agent:security-auditor` in Claude Code. But the more interesting behavior is automatic invocation, which is driven by `description`.

### The `description` Field

This is the field that determines whether automatic delegation works well or poorly. Claude Code reads the `description` to decide which sub-agent to route a task to. Think of it as a matching signal, not a label.

A weak description:

`description: Handles security stuff`
A strong description:

`description: Reviews code for security vulnerabilities including SQL injection, XSS, authentication flaws, insecure dependencies, and secrets exposure. Use this agent when analyzing code changes for security risks or performing security audits.`
The difference is specificity. The second version tells Claude Code exactly when to route tasks here. It names specific vulnerability types and explicitly states the triggering condition (“use this agent when…”).

Write `description` fields as if you’re writing instructions for someone else’s AI orchestrator — because you are. Be concrete about what the agent handles, when it should be used, and what kinds of inputs or tasks trigger it.

### The `tools` Field

Tools define what capabilities the agent has access to. You specify them as a comma-separated list. Available tools in Claude Code include:

- `Read` — read files
- `Write` — write files
- `Edit` — make targeted edits to files
- `Bash` — run shell commands
- `Grep` — search file contents
- `Glob` — find files by pattern
- `WebFetch` — fetch URLs
- `Task` — spawn sub-agents (allows nesting)
- `TodoRead` ,`TodoWrite` — manage task lists
- `mcp__<server>__<tool>` — tools from connected MCP servers

`tools: Read, Grep, Glob`
Restricting tools is one of the most practical things you can do. A documentation agent doesn’t need `Bash`. A code review agent doesn’t need `Write`. Keeping tool access narrow reduces the chance of an agent doing something it shouldn’t.

### The `model` Field (Optional)

You can specify which model the agent uses. If you omit this, it inherits from the session default.

`model: claude-opus-4-5`
This is useful when you want a fast, cheap model for straightforward tasks (like checking formatting) and a more capable model for complex reasoning tasks (like architecture review).

## Configuring Tools for Different Agent Types

Tool configuration is where a lot of the practical design work happens. Here’s how to think about tool assignment for common agent types.

### Read-Only Agents

Agents whose job is analysis, review, or reporting should generally have no write access:

`tools: Read, Grep, Glob`
This covers code reviewers, security auditors, documentation auditors, and linters. They need to see everything but shouldn’t touch anything.

### Write-Capable Agents

Agents that generate or modify content need write tools:

`tools: Read, Write, Edit, Grep, Glob`
Documentation writers, test generators, and code formatters fit here. Still no `Bash` unless they actually need to run commands.

### Execution Agents

Agents that run tests, build code, or execute scripts need `Bash`:

`tools: Read, Bash, Glob`
Keep these narrow in other ways — an agent that just runs tests doesn’t need `Write` access to the entire codebase.

### Orchestrator Agents

If you’re building nested agent structures, the top-level orchestrator needs `Task` to spawn sub-agents:

`tools: Task, Read, Glob`
The `Task` tool is what enables Claude Code to delegate work down to sub-agents. Without it, an agent can’t invoke others.

## Writing the System Prompt Body

Below the `---` closing the front matter is the agent’s system prompt. This is standard markdown text, and it’s the instruction set the agent operates on.

Write system prompts that are:

**Role-specific.** State clearly what the agent is and isn’t responsible for.

```
You are a test writer. Your job is to write unit and integration tests 
for code changes. You do not modify the source code itself — only test files.
```
**Explicit about output format.** Don’t assume the agent will know how you want results structured.

```
For each issue you find, output:
- File path and line number
- Description of the problem
- Suggested fix (code snippet if applicable)
- Severity: low / medium / high
```
**Bounded in scope.** Tell the agent what to ignore as much as what to handle.

```
Focus only on the changed files provided. Do not audit the entire codebase 
unless explicitly asked.
```
The system prompt and the YAML front matter work together. The `description` gets the agent invoked; the system prompt determines how it behaves once it’s running.

## Triggering Sub-Agents: Manual vs. Automatic

There are two ways a sub-agent gets invoked in Claude Code: explicit and automatic.

### Explicit Invocation

Type the agent name directly in your prompt:

`/agent:security-auditor review the changes in src/auth/`
This bypasses the orchestrator’s routing logic entirely. Use explicit invocation when you know exactly which agent you want and don’t want to rely on automatic matching.

### Automatic Invocation

