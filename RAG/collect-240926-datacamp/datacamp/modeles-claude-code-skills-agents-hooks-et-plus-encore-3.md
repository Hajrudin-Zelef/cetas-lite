---
id: collect-240926-datacamp/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore-3
title: "Database Migration Skill"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "guardrails", "mcp"]
source: docs/RAG/clean_en/datacamp/modeles-claude-code-skills-agents-hooks-et-plus-encore.md
source_anchor: ""
source_lines: [300, 382]
sha256: a9c5be13aa2357381db352935237bbef0576a9bc0fb4eaf340119ef16d1af732
---

# Database Migration Skill

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
