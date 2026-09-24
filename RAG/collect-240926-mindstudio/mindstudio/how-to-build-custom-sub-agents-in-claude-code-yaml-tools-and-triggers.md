---
id: collect-240926-mindstudio/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers
title: "how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "claude", "latency", "license", "mcp", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers.md
source_anchor: ""
source_lines: [1, 389]
sha256: 09dda089495a6973bcdaa059076c013974cbf591c69a87b01c5deb0d3466e322
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

When you don’t specify an agent, Claude Code’s orchestrator reads the task and matches it against your agents’ `description` fields. If the description is well-written, the right agent gets called automatically.

For automatic invocation to work reliably:

1. 
**Use keywords the user will naturally say.** If someone says “can you check for security issues,” your description should include “security issues.”
2. 
**State conditions explicitly.** Phrases like “use this agent when…” or “invoke for tasks involving…” help the orchestrator pattern-match.
3. 
**Avoid overlapping descriptions.** If two agents have similar descriptions, routing becomes unpredictable. Make each agent’s scope distinct.
4. 
**Test with realistic prompts.** After writing your agents, actually try the kinds of prompts you expect users (or your orchestrator) to use, and see which agent gets invoked.

### The `trigger` Concept in Practice

Claude Code doesn’t have a formal `trigger` field in the YAML front matter. Triggering is entirely driven by the `description`. This means the quality of your description *is* your trigger configuration. There’s no separate mechanism — write a precise, keyword-rich description, and you’ve configured your trigger.

## Practical Sub-Agent Examples

Here are four full sub-agent configurations you can adapt directly.

### Code Reviewer Agent

```
---
name: code-reviewer
description: Reviews code changes for bugs, logic errors, code quality, 
and adherence to best practices. Use this agent for PR reviews, code 
analysis, or when evaluating whether a change is correct and maintainable.
tools: Read, Grep, Glob
model: claude-opus-4-5
---
You are a senior code reviewer. Review the provided code changes carefully.
For each issue, report:
- Location (file + line number)
- What the problem is
- Why it's a problem
- How to fix it
Categorize issues as: bug, style, performance, or maintainability.
Be direct. Don't pad your output with compliments. If the code is fine, say so.
```
### Test Writer Agent

```
---
name: test-writer
description: Writes unit tests, integration tests, and test fixtures for 
new or modified code. Invoke when asked to add tests, improve test coverage, 
or generate test cases for a function or module.
tools: Read, Write, Edit, Grep, Glob
---
You are a test engineer. Write thorough, readable tests for the code provided.
Follow the testing patterns already established in the codebase. 
Match the existing test framework and assertion style.
Cover:
- Happy path
- Edge cases
- Error conditions
- Boundary values
Do not modify source files. Only create or edit files in test directories.
```
### Documentation Agent

```
---
name: docs-writer
description: Writes and updates technical documentation including README 
files, inline code comments, JSDoc/docstrings, API docs, and changelogs. 
Use for any documentation task — generating new docs or updating existing ones.
tools: Read, Write, Edit, Grep, Glob
---
You write clear, accurate technical documentation.
Match the documentation style already in use. If there's a README, follow 
its structure. If there are existing docstrings, follow their format.
Be concise. Don't explain what the reader can see in the code — explain 
why and how to use it.
```
### Dependency Auditor Agent

```
---
name: dependency-auditor
description: Audits project dependencies for security vulnerabilities, 
outdated packages, license compliance issues, and unused dependencies. 
Use when reviewing package.json, requirements.txt, Gemfile, or similar files.
tools: Read, Bash, Glob
---
You audit software dependencies for risk.
Check for:
- Known CVEs and security vulnerabilities
- Packages that are significantly out of date
- License compatibility issues
- Dependencies that appear unused
Run audit tools available in the environment (npm audit, pip-audit, etc.) 
when appropriate. Report findings by severity.
```
## Structuring Multi-Agent Workflows

Once you have several agents, the next question is how they coordinate. Claude Code supports a few patterns.

### Sequential Delegation

The orchestrator delegates tasks one at a time, in order. Useful when each step depends on the previous:

1. Code reviewer runs → reports issues
2. Test writer runs → writes tests for the reviewed code
3. Docs writer runs → updates documentation

You can describe this workflow in your main Claude Code session or build a coordinator agent that manages the sequence.

### Parallel Delegation

For independent tasks, the orchestrator can invoke multiple sub-agents roughly simultaneously. A security audit and a documentation update don’t depend on each other, so they can run in parallel.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

To enable this, the orchestrating agent needs the `Task` tool and should be instructed to parallelize where possible.

### Nested Agents

Sub-agents can themselves spawn sub-agents, as long as they have the `Task` tool. This allows deep specialization — a “frontend agent” might have sub-agents for CSS review, accessibility checking, and performance analysis.

Be careful with nesting depth. More than two or three levels tends to become hard to debug and reason about.

## Extending Claude Code Sub-Agents with MindStudio

One real limitation of Claude Code sub-agents is that they’re constrained to what’s available in the shell environment and MCP servers you’ve connected. When you need agents to send emails, query CRMs, generate images, or interact with dozens of SaaS tools, wiring all of that up manually is a lot of work.

The MindStudio Agent Skills Plugin — an npm SDK — solves this for Claude Code specifically. You install it in your project and your agents can call 120+ typed capabilities as simple method calls:

```
import MindStudio from '@mindstudio-ai/agent';
const agent = new MindStudio();
await agent.sendEmail({ to: 'team@company.com', subject: '...', body: '...' });
await agent.searchGoogle({ query: 'latest CVEs for express.js' });
await agent.runWorkflow({ workflowId: 'onboarding-sequence', inputs: {...} });
```
Your security auditor sub-agent could automatically file a Jira ticket when it finds a critical vulnerability. Your documentation agent could post updates to Notion. Your dependency auditor could send a Slack alert if it finds a high-severity CVE.

The plugin handles rate limiting, retries, and authentication — so your sub-agent’s system prompt can focus on reasoning, not infrastructure. You connect it to Claude Code via an MCP server configuration, which means the tools show up in your YAML `tools` field just like native tools.

You can try MindStudio free at mindstudio.ai.

## Common Mistakes and How to Fix Them

### Vague Description Fields

The most common reason automatic agent invocation doesn’t work is a description that’s too general. If the description says “handles code tasks,” the orchestrator has no basis for choosing this agent over others.

Fix: Rewrite descriptions to name specific tasks, file types, technologies, or trigger conditions explicitly.

### Overlapping Scopes

Two agents with similar descriptions will cause unpredictable routing. Claude Code might pick either one.

Fix: Draw clear boundaries. If your code reviewer and security auditor descriptions both mention “reviewing code,” the orchestrator will sometimes pick the wrong one. Make the security auditor’s description clearly about security specifically, not general code review.

### Over-Provisioning Tools

Giving agents broad tool access (especially `Bash` and `Write`) when they don’t need it increases the chance of unintended side effects.

Fix: Start with the minimum tool set and add only what’s demonstrably needed. Test that each tool is actually being used.

### System Prompt Scope Creep

Adding more and more instructions to a system prompt until the agent tries to do too many things at once.

Fix: If an agent is doing multiple unrelated jobs, split it into two agents. Narrow, focused agents are easier to debug and more reliable.

### Missing Agent Directory

If your `.claude/agents/` directory doesn’t exist or is in the wrong location, agents simply won’t appear.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

Fix: Verify the path. Project-level agents must be in `.claude/agents/` relative to your project root. User-level agents must be in `~/.claude/agents/`. Run `ls .claude/agents/` to confirm files are there.

## Frequently Asked Questions

### How many sub-agents can you define in Claude Code?

There’s no documented hard limit on the number of sub-agents. In practice, too many agents with overlapping descriptions creates routing ambiguity. Start with a focused set — four to eight specialized agents is usually enough for most projects — and add more only when you have a clear, distinct need.

### Can sub-agents call other sub-agents?

Yes. If a sub-agent has the `Task` tool in its YAML front matter, it can spawn other sub-agents. This enables nested orchestration patterns. Be aware that each level of nesting adds latency and makes debugging more complex.

### What’s the difference between project-level and user-level agents?

Project-level agents live in `.claude/agents/` inside your project repo and are only available when you’re working in that project. User-level agents live in `~/.claude/agents/` in your home directory and are available across all your Claude Code sessions. If the same agent name exists in both locations, the project-level version takes precedence.

### Do sub-agents share context with the main Claude Code session?

Sub-agents have access to the context passed to them by the orchestrator when they’re invoked, but they run as separate contexts — they don’t automatically see the full conversation history of the main session. The orchestrator includes relevant information when delegating tasks.

### Can you use sub-agents without the `Task` tool?

You can invoke sub-agents explicitly using `/agent:name` syntax without needing the `Task` tool — that’s a manual invocation. The `Task` tool is required for an agent to programmatically delegate to other agents within an automated workflow.

### Are sub-agent markdown files included in version control?

Yes, and that’s intentional. Because sub-agents are plain files in `.claude/agents/`, they can (and should) be committed to version control. This means your whole team shares the same agent configurations, and agent changes are tracked just like code changes.

## Key Takeaways

- Sub-agents in Claude Code are markdown files with YAML front matter, stored in `.claude/agents/` (project-level) or`~/.claude/agents/` (user-level)
- The `description` field drives automatic triggering — write it with specific keywords and explicit trigger conditions, not generic labels
- The `tools` field should be as narrow as possible; only grant what the agent actually needs
- The system prompt body is standard markdown and should clearly define the agent’s role, output format, and scope
- Vague descriptions, overlapping scopes, and over-provisioned tools are the three most common failure modes
- For agents that need to interact with external tools and services, the MindStudio Agent Skills Plugin adds 120+ capabilities to Claude Code sub-agents with minimal configuration

If you want to extend what your sub-agents can do beyond the shell environment, MindStudio is worth a look — it connects Claude Code to the full range of business tools without requiring you to wire up each integration manually.
