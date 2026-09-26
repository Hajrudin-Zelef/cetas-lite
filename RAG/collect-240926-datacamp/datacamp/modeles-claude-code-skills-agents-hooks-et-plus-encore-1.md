---
id: collect-240926-datacamp/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore-1
title: "Database Migration Skill"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "distribution", "guardrails", "mcp"]
source: docs/RAG/clean_en/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [1, 124]
sha256: 455dac40522684c98aa65c7b439a38b0c8d8f299dced38f4c08ab6e1f8429460
---

# Database Migration Skill

<!-- source: https://www.datacamp.com/fr/blog/claude-code-templates -->

Course

Claude Code is powerful, but without reusable configuration, you end up repeating the same instructions over and over again. Claude Code templates solve this problem: they turn recurring instructions, workflows, tool permissions, and integrations into reusable project files that Claude can discover and apply.

In this article, we'll look at what Claude Code templates are, the main types available, how each one behaves, how to choose the right type for your workflow, and where to find ready-to-use templates.

This article assumes you already have a basic Claude Code setup. If you're just getting started, begin with this Claude Code tutorial before going further with templates. If you're still learning how Claude Code fits into command-line development, check out this introductory guide to the Claude Code CLI.

## In brief

- 
Claude Code templates are reusable file-based configurations (stored in `.claude/`) that save you from re-explaining your stack and workflows in every session.
- 
There are six types: skills (repeatable workflows), agents (scoped roles and permissions), commands (manual slash actions), hooks (automatic guardrails), MCP (connections to external tools and data), and plugins (bundles of the previous five).
- 
`CLAUDE.md` always contains your project brief; templates add modular, reusable behaviors on top of it.
- 
Choose based on the trigger: hooks enforce rules automatically, agents bring domain expertise, skills encode repeatable workflows, commands run on demand, MCP accesses external systems, and plugins package and share a complete configuration.
- 
Start small with Anthropic's official documentation, a community collection like aitmpl.com, or your own files, and create a first skill for your most repeated task before expanding.

## Introduction to artificial intelligence agents

## What are Claude Code templates?

Claude Code templates are reusable configuration files that customize the behavior of Claude Code within a project or across your local environment.

Key point: templates are file-based. You don't install them through a classic settings interface by clicking through configuration screens. Instead, Claude Code detects specific files and folders, loads the relevant metadata into context, and relies on it to decide how to behave.

In practice, these are usually Markdown, JSON, or shell-oriented files, stored in project-level folders like `.claude/`, or packaged in plugin-type directories for sharing.

A typical project-level structure might look like this:

```
my-app/
├── CLAUDE.md
├── .mcp.json
└── .claude/
    ├── skills/
    │   └── database-migration/
    │       └── SKILL.md
    ├── agents/
    │   └── security-auditor.md
    ├── commands/
    │   └── summarize-pr.md
    └── settings.json
```
### Claude Code templates vs CLAUDE.md

`CLAUDE.md` remains important, but its role is different. Think of `CLAUDE.md` as the project brief: what the project is, which commands matter, which code standards apply, and which architecture conventions Claude should keep in mind.

For a step-by-step guide, see our CLAUDE.md writing guide.

Templates are more modular:

- A skill can encode a migration workflow.
- An agent can isolate a security review posture.
- A hook can run after file edits.
- An MCP configuration can connect Claude to GitHub, SQLite, or another external system.

This is also where templates fit into the overall design of your Claude Code workflow. Strong templates deliver the best results when paired with good practices for planning, testing, and context handoff.

Find more of these practices in our best practices guide.

Custom commands are also part of the skills system, even though the older `.claude/commands/` format still works. The new recommended format is `.claude/skills/<name>/SKILL.md`, which supports both slash-command invocation and automatic invocation by Claude.

## What types of Claude Code templates can I use?

The Claude Code template ecosystem is generally structured into six categories: skills, agents, commands, hooks, MCP integrations, and plugins.

The first five directly modify Claude's behavior. Plugins are a bit different: they are a distribution format that can bundle skills, agents, hooks, commands, MCP servers, and other components into a reusable package.

We'll look at each of these categories below.

### 1. Skills

Skills are sets of instructions for repeatable multi-step tasks. A skill is usually a folder containing a `SKILL.md` file with YAML frontmatter and a Markdown body.

The frontmatter describes what the skill does and how it should behave; the body tells Claude the steps to follow. For a dedicated breakdown, see this Claude Skills guide.

Claude uses the skill's description to decide when it is relevant. By default, both the user and Claude can invoke a skill: you can type `/skill-name`, or Claude can load it automatically when the current task matches its description. You can also disable automatic invocation of the template for workflows where you want to stay in control, such as deployment.

Here is a short example file: `.claude/skills/database-migration/SKILL.md`

```
---
name: database-migration
description: Use when creating, reviewing, or modifying database migrations. Ensures migrations are reversible, tested, and checked before and after execution.
allowed-tools:
  - Read
  - Write
  - Bash
---
# Database Migration Skill
When working on a database migration:
1. Inspect the existing schema and migration history before writing changes.
2. Confirm whether the migration is additive, destructive, or data-transforming.
3. Create a reversible migration whenever the framework supports rollback.
4. Run the project’s migration check command before applying the migration.
5. Run tests that cover the affected models, queries, or API endpoints.
6. After writing the migration, summarize:
   - schema changes
   - rollback behavior
   - affected tables
   - test commands run
```
This is useful because the instructions are procedural. You are not just telling Claude to “be careful with migrations”; you are giving it a repeatable checklist.

Skills are suitable for anything you would otherwise paste into Claude more than twice: generating API endpoints, writing changelogs, scaffolding tests, creating release notes, reviewing pull requests, or migration checks.

For more inspiration on the workflows developers turn into reusable AI automations, see our list of Agent Skills.

### 2. Agents

Agents, more specifically custom subagents in Claude Code, are specialized AI assistants with their own Markdown definition, YAML frontmatter, tool restrictions, model choice, and system prompt.

They can live in `.claude/agents/` for project scope, or `~/.claude/agents/` for personal scope. Agents are created by asking Claude for them or by directly editing the Markdown files in the `.claude/agents/` folder.

There is a real difference between skills and agents. A skill defines how to execute a task. An agent defines who Claude should be during the work: its role, focus, permissions, and limits.

Let’s look at an example agent:

