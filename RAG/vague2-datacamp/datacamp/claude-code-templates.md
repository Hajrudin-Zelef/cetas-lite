---
id: vague2-datacamp/datacamp/claude-code-templates
title: "Modèles Claude Code : skills, agents, hooks et plus encore"
domain: datacamp
role: reference
task: tutorial
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "claude", "distribution", "guardrails", "mcp", "model context protocol", "reasoning"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/claude-code-templates.md
source_anchor: ""
source_lines: [1, 60]
sha256: c5b3e5804ec98b0385298ce5d41cb71e751abd75a2aef9ccd82ae5cf4957b2c6
---

# Modèles Claude Code : skills, agents, hooks et plus encore

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/claude-code-templates
- **Site** : DataCamp
- **Type** : Article / Tutorial
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article explains Claude Code templates: reusable, file-based configurations (Markdown, JSON, or shell files) stored in project-level `.claude/` folders or plugin directories that customize Claude Code's behavior. Claude Code detects these files, loads relevant metadata into context, and uses them to decide how to behave — there is no settings UI. A typical project structure includes `CLAUDE.md`, `.mcp.json`, and `.claude/` containing `skills/`, `agents/`, `commands/`, and `settings.json`.

The article distinguishes templates from `CLAUDE.md`: `CLAUDE.md` is the project brief (stack, conventions, commands, architecture), while templates are modular behaviors. There are **six template categories**:

1. **Skills** — instruction sets for repeatable multi-step tasks, stored as a folder with `SKILL.md` (YAML frontmatter + Markdown body). Claude can auto-invoke based on the description, or users can call `/skill-name`; auto-invocation can be disabled. Example: a `database-migration` skill with `allowed-tools` (Read, Write, Bash) and a 6-step checklist. Best for anything you'd paste more than twice (API endpoints, changelogs, tests, release notes, PR reviews, migration checks). Note: custom commands now fall under the skills system, though the legacy `.claude/commands/` format still works; the recommended new format is `.claude/skills/<name>/SKILL.md`.

2. **Agents** — custom subagents with their own Markdown definition, YAML frontmatter, tool restrictions, model choice, and system prompt. Stored in `.claude/agents/` (project) or `~/.claude/agents/` (personal). A skill defines *how* to execute; an agent defines *who* Claude is (role, focus, permissions). Example: a `security-auditor` restricted to Read, Glob, Grep, Bash with instructions not to edit files, producing a severity-rated findings report. Ideal for specialized domains needing context/permission isolation.

3. **Commands** — slash-invoked shortcuts like `/generate-tests` or `/summarize-pr`, historically Markdown files in `.claude/commands/`. Skills are recommended for new command-style workflows (they support `/name` and auto-trigger). Commands are preferable for explicit triggers and checkpoints.

4. **Hooks** — shell commands run automatically in response to Claude Code lifecycle events (`PreToolUse`, `PostToolUse`, `Notification`, `Stop`). Triggered by what Claude does, not what you ask — deterministic control. Examples: run Prettier after Edit/Write via a PostToolUse hook; block dangerous Bash via a PreToolUse Python script. Ideal for non-bypassable rules: linters, formatting, protecting files, notifications.

5. **MCP integrations** — connect Claude Code to external tools, data sources, and APIs via the Model Context Protocol. Servers expose **Tools** (executable functions), **Resources** (read-only context), and **Prompts** (reusable task templates). A project `.mcp.json` can configure multiple stdio servers (e.g., GitHub, SQLite) with env vars. Essential when Claude needs live systems rather than a static code snapshot.

6. **Plugins** — packaged bundles that can include skills, agents, hooks, MCP configs, and commands. A plugin has a `.claude-plugin/plugin.json` manifest plus component folders (`skills/`, `agents/`, `hooks/`, `.mcp.json`). Plugins add no new behavior type; they make the others portable for team sharing and multi-project reuse.

The article includes a comparison table (trigger, best for, poorly suited to, use case) and a simple decision rule: hook = enforce automatically; agent = specialized expertise; skill = encode a workflow; command = trigger yourself; MCP = external/live data; plugin = install/share a full setup. Types often combine (e.g., a security plugin bundling an agent, skill, command, and pre-commit hook). It also recommends spec-driven development for formal planning.

**Where to find templates:** Anthropic's official docs; community collections, chiefly **aitmpl.com** (install via `npx claude-code-templates@latest`, alias `npx cct@latest`, with flags like `--agent`, `--command`, `--hook`, `--mcp`); or write your own. Evaluate community templates by checking description precision, `allowed-tools` scope, repo activity, and whether hooks/MCP servers run unaudited code. Start small with one skill for your most repeated task, then add agents, hooks, and MCP as friction appears.

## Key points

- Templates are reusable file-based configs in `.claude/`, not UI settings.
- Six types: skills, agents, commands, hooks, MCP integrations, and plugins.
- `CLAUDE.md` remains the project brief; templates add modular, reusable behavior.
- Trigger-based choice: hooks auto-enforce, agents bring expertise, skills encode workflows, commands run on demand, MCP accesses external systems, plugins bundle everything.
- Skills use `SKILL.md` with YAML frontmatter; commands now fold into the skills system.
- Hooks fire on lifecycle events (`PreToolUse`, `PostToolUse`, `Notification`, `Stop`) for deterministic control.
- MCP servers expose Tools, Resources, and Prompts.
- Start with one skill, then expand; vet community templates for tool permissions and unaudited code.

## Technical data / figures

| Template type | Triggered by | Best for | Poor fit |
|---|---|---|---|
| Skill | Claude auto or user `/skill-name` | Repeatable multi-step workflows | One-off tasks |
| Agent | User request or Claude delegation | Domain expertise + permission isolation | General sessions |
| Command | User slash command | On-demand actions/checkpoints | Automatic guardrails |
| Hook | Claude lifecycle event | Automated guardrails/quality checks | Interactive reasoning |
| MCP | Claude tool call | External systems access | Simple local-only workflows |
| Plugin | Installation/activation | Team distribution, bundled workflows | Single-purpose local tweaks |

Install commands: `npx claude-code-templates@latest`, `npx cct@latest`; flags `--agent`, `--command`, `--hook`, `--mcp`, `--yes`. Key paths: `.claude/skills/<name>/SKILL.md`, `.claude/agents/`, `~/.claude/agents/`, `.mcp.json`, `.claude-plugin/plugin.json`.

## Why this source matters for the RAG

It is a precise reference for Claude Code's extensibility model, including file formats, triggers, and selection guidance. It supports RAG queries about agent skills, subagents, hooks, MCP, plugins, and Claude Code configuration best practices.
