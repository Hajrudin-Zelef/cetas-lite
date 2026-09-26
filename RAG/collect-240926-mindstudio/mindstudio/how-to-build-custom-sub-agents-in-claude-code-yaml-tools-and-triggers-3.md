---
id: collect-240926-mindstudio/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers-3
title: "how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "latency"]
source: docs/RAG/clean_en/mindstudio/how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers.md
source_anchor: ""
source_lines: [340, 389]
sha256: cf06e07d93d185bc35aa60c61b84680e8ffe30054149a44096132723c737db2a
---

# how-to-build-custom-sub-agents-in-claude-code-yaml-tools-and-triggers

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
