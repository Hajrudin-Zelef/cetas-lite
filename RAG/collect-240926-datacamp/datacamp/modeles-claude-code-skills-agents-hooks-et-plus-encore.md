---
id: collect-240926-datacamp/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore
title: "Database Migration Skill"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "distribution", "guardrails", "mcp", "model context protocol", "reasoning", "safeguards"]
source: docs/RAG/clean_en/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [1, 382]
sha256: 3ba80960bbfed21886279b3272b7643e1c658a30f5229d575953c74ec5243118
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

```
---
name: security-auditor
description: Reviews code for security vulnerabilities and produces a findings report without modifying files.
tools: Read, Glob, Grep, Bash
model: sonnet
---
You are a security auditor.
Your task is to inspect the codebase for vulnerabilities, risky patterns, and missing safeguards.
Rules:
- Do not edit files.
- Do not suggest broad rewrites unless directly tied to a security issue.
- Focus on authentication, authorization, input validation, secrets, dependency risk, and unsafe shell or SQL usage.
- Produce a findings report with severity, affected files, evidence, and recommended next steps.
```
This agent is useful because it sets clear boundaries. In a general session, Claude might start fixing problems as soon as it finds them. A security auditor agent is instructed to inspect and report only, without modifications.

Agents are ideal for specialized domains such as security auditing, documentation review, architecture review, data engineering, or code quality checks, where context isolation and permissions are crucial.

They are also effective in tandem with specialized skills. For example, a security auditor agent can call a findings report skill, while a frontend reviewer agent can use a component testing skill.

### 3. Commands

Commands are slash-invoked shortcuts, such as `/generate-tests`, `/check-deps`, or `/summarize-pr`. Historically, custom commands were stored as Markdown files under `.claude/commands/`, with the file name serving as the command name.

Claude Code still supports this legacy format, but we recommend using skills for new command-style workflows: they handle the same `/name` invocation and can be triggered automatically when relevant.

Commands are preferable when you want explicit triggering. A skill can launch automatically if Claude detects a matching task, but a command should only run when you launch it. They are therefore ideal as checkpoints: “generate the tests now,” “summarize this PR now,” “check the dependencies now,” “prepare a commit message now.”

### 4. Hooks

Hooks are automation rules executed in response to Claude Code lifecycle events. They are user-defined shell commands, executed at precise moments, that provide deterministic control over behavior.

The difference from the other patterns discussed above is that they are not triggered by what you ask, but by what Claude does.

In practice, you do not have to hope that Claude remembers to format a file after editing it; a hook can do it automatically.

Current hook events include notably `PreToolUse`, `PostToolUse`, `Notification`, and `Stop`.

Example: run a formatter after Claude has edited or written a file:

```

```
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Edit|Write",
        "hooks": [
          {
            "type": "command",
            "command": "jq -r '.tool_input.file_path' | xargs npx prettier --write"
          }
        ]
      }
    ]
  }
}
Example: block risky shell commands before Claude runs them:
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Bash",
        "hooks": [
          {
            "type": "command",
            "command": "python3 .claude/hooks/block-dangerous-bash.py"
          }
        ]
      }
    ]
  }
}
```
Hooks are ideal for rules that Claude should not be able to bypass: running a linter, formatting modified files, blocking edits to protected files, verifying generated code, or sending notifications when Claude needs input.

For an in-depth tutorial, read our Claude Code hooks guide.

### 5. MCP Integrations

MCP integrations connect Claude Code to external tools, data sources, and APIs via the Model Context Protocol. MCP acts as the connector layer between AI systems and external tools. In Claude Code, this allows Claude to go beyond local files and shell commands.

Claude can then interact with external services like GitHub, databases, documentation systems, cloud platforms, or internal APIs, depending on the MCP servers you configure. For a complete explanation and a demo project, see our Model Context Protocol tutorial.

An MCP server can expose three main types of capabilities:

- **Tools**: executable functions callable by Claude, such as creating a GitHub ticket or running a database query.
- **Resources**: read-only sources of context, such as a file, a baseline, or a document.
- **Prompts**: reusable task templates exposed by the server.

A project-level `.mcp.json` can configure multiple servers side by side:

```
{
  "mcpServers": {
    "github": {
      "type": "stdio",
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${GITHUB_TOKEN}"
      }
    },
    "sqlite": {
      "type": "stdio",
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-sqlite",
        "./data/app.db"
      ]
    }
  }
}
```
This is important because Claude can only reason from the context and tools it has access to. Without MCP, it can inspect local files, but not your ticketing system, your database, your cloud environment, or your internal APIs.

MCP is ideal when Claude needs to work with your real stack rather than a static code snapshot, and when it needs access to external data.

### 6. Plugins

Plugins are packaged bundles. They can include skills, agents, hooks, MCP configurations, commands, and other components within a single installable structure.

In Claude Code, a plugin generally includes a `.claude-plugin/plugin.json` manifest and component folders such as `skills/`, `agents/`, `hooks/`, and `.mcp.json` at the plugin root.

Example plugin structure:

```
frontend-workflow-plugin/
├── .claude-plugin/
│   └── plugin.json
├── skills/
│   └── component-test/
│       └── SKILL.md
├── agents/
│   └── frontend-reviewer.md
├── hooks/
│   └── hooks.json
└── .mcp.json
```
Example plugin json:

```
{
  "name": "frontend-workflow",
  "displayName": "Frontend Workflow",
  "version": "1.0.0",
  "description": "Frontend development workflow with review agents, test skills, and formatting hooks",
  "author": {
    "name": "Your Team"
  }
}
```
Plugins do not add a new type of behavior; they make the other types portable. Use them to share a complete configuration within a team, reuse the same workflow across multiple projects, or install a community bundle rather than creating each file manually.

To create one from scratch, see DataCamp's step-by-step guide to Claude Code plugins.

## Which type of template should you choose?

It's understandable: all these types can be confusing. They all change Claude's behavior. The main difference lies in how they are triggered and how much control they provide.

To clarify, here is a comparison between them:

| **Template type** | **Triggered by** | **Ideal for** | **Not well suited to** | **Use case** | 
| Skill | Claude automatically or user via `/skill-name` | Repeatable multi-step workflows | One-off tasks | Automatically applying a migration checklist when Claude modifies schema files | 
| Agent | User request or delegation by Claude | Domain expertise and permission isolation | Generalist sessions | A security auditor that can read files but must not edit them | 
| Command | User slash command | On-demand actions and checkpoints | Automatic guardrails | /generate-tests when you are ready to test | 
| Hook | Claude lifecycle event | Automated guardrails and quality checks | Tasks requiring interactive reasoning | Format files after every edit | 
| MCP | Tool call by Claude | Access to external systems | Simple local-only workflows | Query PostgreSQL or create a GitHub ticket | 
| Plugin | Installation or activation | Team distribution and grouped workflows | Local single-purpose adjustments | A frontend bundle with agents, skills, and hooks | 

A simple rule for deciding:

- If you want to **automatically enforce a rule** every time Claude touches code, use a **hook**.
- If you want Claude to **adopt deep domain expertise** for a specific task, use an **agent**.
- If you want to **encode a workflow** that Claude must repeat faithfully, use a **skill**.
- If you want to **trigger an action yourself** at the right moment, use a **command** or a command-type skill.
- If Claude needs **external services or live data**, use **MCP**.
- If you want to install or share a **complete workflow configuration**, use a **plugin**.

In practice, these types often combine. A security plugin might bundle a `security-auditor` agent, an `audit-findings` skill, a dependency-check command, and a pre-commit hook. The agent defines the role, the skill defines the report structure, the command provides an explicit checkpoint, and the hook enforces the guardrail.

For workflows where Claude must follow a formal plan before implementation, spec-driven development is often more suitable than ad-hoc prompts.

## Where to find Claude Code templates?

Three practical sources: Anthropic, community collections, and your own creations.

Start with Anthropic's official resources and documentation. Anthropic's Claude Code documentation covers skills, subagents, hooks, MCP, and plugins; it's the ideal place to verify up-to-date file formats and behaviors before any production deployment.

Then, use community collections. The most visible hub is aitmpl.com, which presents itself as a catalog of ready-to-use configurations for Claude Code projects. Its navigation currently offers Skills, Agents, Commands, Settings, Hooks, MCPs, and Plugins.

The current interactive installation command is:

`npx claude-code-templates@latest`
The project documentation also presents a shorter alias:

`npx cct@latest`
For specific components, the online GitHub README offers installation commands such as:

```
npx claude-code-templates@latest --agent development-tools/code-reviewer --yes
npx claude-code-templates@latest --command performance/optimize-bundle --yes
npx claude-code-templates@latest --hook git/pre-commit-validation --yes
npx claude-code-templates@latest --mcp database/postgresql-integration --yes
```
You can also install a complete stack in batch via multiple options in a single command.

When evaluating community templates, check a few quality signals:

- 
Is the `description` precise enough for Claude to correctly trigger the skill or agent?
- 
Are the `allowed-tools` properly scoped, or does the template unnecessarily request broad write and bash permissions?
- 
Is the repository recently maintained?
- 
Does the template explain what it modifies?
- 
Does it include hooks or MCP servers that execute code you haven't audited?

Finally, write your own. This is often the best option for workflows closely tied to your stack. A community template provides a good starting point, but it doesn't know your internal migration policy, your naming conventions, your data model, or your tolerance for deployment risk.

## To conclude

Claude Code templates let you move from a session-based assistant to a persistent development environment.

The six categories presented stack in layers: skills encode workflows, agents define roles, commands create explicit actions, hooks enforce guardrails, MCP connects external systems, and plugins package everything for reuse.

The best starting point isn't a huge bundle of plugins. Start with a skill for your most repetitive workflow. As soon as you identify remaining friction in Claude's default behavior, add an agent for specialized review, a hook for automatic enforcement, or an MCP server for live access to systems.

To go further with Claude Code, explore our Claude Code 101 and Claude Code in Action courses.

## FAQ on Claude Code templates

### Are Claude Code templates the same as CLAUDE.md?

**No. CLAUDE.md is mainly used for general project-level instructions: tech stack, code conventions, project structure, and preferred commands. Claude Code templates are more modular. They bundle workflows, roles, commands, hooks, or integrations that Claude can use as needed.**

### Should I use a skill or an agent?

**Use a skill when you want Claude to follow a repeatable process, for example generating tests, writing changelogs, or reviewing migrations. Use an agent when you want Claude to adopt a specific role, for example security auditor, documentation reviewer, or frontend architect. In many workflows, you can use both together.**

### Are Claude Code templates project-specific or global?

**Both are possible, depending on where they are stored. Project-specific templates are generally found in the project's `.claude/` directory. Global templates are useful if you want the same behavior across multiple projects.**

### Are community Claude Code templates safe to install?

**Not automatically. Community templates can be very useful, but they may include tool permissions, shell commands, hooks, or MCP configurations that affect your local environment.**

### What is the best type of template to start with?

**Start with a skill. This is generally the simplest way to turn repeated instructions into reusable workflows, without complicating your setup. Once a first useful skill is in place, you can add agents, hooks, MCP, and plugins.**

My name is Austin, I am a blogger and technical writer and I have years of experience as a data scientist and data analyst in the healthcare field. I started my technology journey with a background in biology and I now help others make the same transition through my technology blog. My passion for technology has led me to write for dozens of SaaS companies, inspiring others and sharing my experiences.
