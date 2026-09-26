---
id: collect-240926-mindstudio/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers-2
title: "how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "claude", "license", "mcp", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers.md
source_anchor: ""
source_lines: [173, 339]
sha256: acf3e92f23b033d7e99a45866ba8e077a5dece884bff5e429c94bc8367f6fc8c
---

# how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers

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

