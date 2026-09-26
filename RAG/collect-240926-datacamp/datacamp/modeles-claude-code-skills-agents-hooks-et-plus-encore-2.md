---
id: collect-240926-datacamp/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore-2
title: "Database Migration Skill"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "distribution", "guardrails", "mcp", "model context protocol", "reasoning", "safeguards"]
source: docs/RAG/clean_en/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [125, 299]
sha256: 3065cdcf99ae5641a1b613026e27a53ee50d6a81eed0db202cd86154fc1f3ebd
---

# Database Migration Skill

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

