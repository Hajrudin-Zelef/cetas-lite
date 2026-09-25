---
id: collect-250926-servers-hardware/servers-hardware/agents
title: "Agents"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "claude", "research", "scout"]
source: docs/RAG/clean4/agents.md
source_anchor: ""
source_lines: [1, 246]
sha256: cc74df25e70fd1a6c692033e2654ef297e22538fb024ffc441eb0c47005bd399
---

# Agents

Create a Markdown file to add a reusable agent. This example adds a read-only reviewer that the main agent can launch for code reviews:

Ask your primary agent to use it:

`Use the reviewer subagent to review my current changes.`
An agent combines a system prompt, model preference, permissions, and display details into a named assistant profile.

## Locations

Save Markdown agents globally for all projects or inside a project:

```
~/.config/opencode/agents/<name>.md
.opencode/agents/<name>.md
```
OpenCode discovers project `.opencode` directories from the current directory up to the project root. A nested path becomes part of the agent ID:

`.opencode/agents/team/reviewer.md  â  team/reviewer`
## Formats

### Markdown

Frontmatter accepts the same fields as an `agents` configuration entry. The Markdown body becomes the agentâs `system` prompt:

### JSONC

Define agents under `agents` in any OpenCode configuration file:

## Selection

Set the primary agent used when a session has not selected one:

The selected default must exist, be visible, and support primary use. Otherwise OpenCode uses `build`, then the first visible primary-capable agent. Changing this setting does not replace the agent stored on an existing session.

## Modes

Set `mode` according to where the agent should run:

```
{
  "agents": {
    "reviewer": { "mode": "subagent" },
  },
}
```
| Mode | Behavior | 
|---|---|
| `primary` | Runs as the main agent for a session. This is the default for a new custom agent. | 
| `subagent` | Runs only in a child session through the `subagent` tool. | 
| `all` | Runs either as a primary agent or a subagent. | 

Subagents run with fresh context in foreground or background child sessions. The parent agentâs `subagent` permissions control which agents it may launch; the child uses its own configured permissions.

```
{
  "agents": {
    "orchestrator": {
      "permissions": [
        { "action": "subagent", "resource": "*", "effect": "deny" },
        { "action": "subagent", "resource": "reviewer", "effect": "allow" },
      ],
    },
  },
}
```
## Builtins

OpenCode includes these visible agents:

| Agent | Mode | Purpose | 
|---|---|---|
| **Build** (`build` ) | `primary` | Default coding agent. Tools are allowed by default; sensitive environment-file reads and access outside the workspace ask for approval. | 
| **Plan** (`plan` ) | `primary` | Explores and plans without editing normal project files. It may write OpenCode plan files when asked, and shell commands remain permission-controlled. | 
| **General** (`general` ) | `subagent` | Handles research and multi-step work with broad tool access, but cannot launch more subagents. | 
| **Explore** (`explore` ) | `subagent` | Searches and reads code or web sources without editing files. | 

Override a built-in by using the same ID:

```
{
  "agents": {
    "build": {
      "permissions": [
        { "action": "shell", "resource": "git push *", "effect": "ask" },
      ],
    },
  },
}
```
Hidden `compaction`, `title`, and `summary` agents perform maintenance and cannot be selected directly. V2 has no built-in `scout` agent.

## Merging

Agent definitions merge in configuration order. Later scalar values replace earlier values, request maps merge by key, and permission rules append:

```
{
  "permissions": [
    { "action": "shell", "resource": "*", "effect": "ask" },
  ],
  "agents": {
    "build": {
      "permissions": [
        { "action": "shell", "resource": "git status", "effect": "allow" },
      ],
    },
  },
}
```
Global `permissions` apply before agent-specific rules, so later agent rules can refine them.

## Options

### Description

`description` explains the agentâs purpose. Add it to subagents because OpenCode shows it to the model choosing which agent to launch:

`description: Reviews database migrations for safety`
### Mode

`mode` accepts `primary`, `subagent`, or `all`. When omitted on a new custom agent, it defaults to `primary`:

`mode: all`
### Model

`model` uses `provider/model` with an optional `#variant`:

`model: anthropic/claude-sonnet-4-5#high`
JSON configuration also accepts the expanded form:

```
{
  "agents": {
    "reviewer": {
      "model": {
        "providerID": "anthropic",
        "model": "claude-sonnet-4-5",
        "variant": "high",
      },
    },
  },
}
```
- A subagent uses its configured model, or inherits the parent sessionâs model when none is configured.
- A session stores its selected model separately. Selecting a primary agent by ID does not change that model.

### System

`system` sets the agentâs system prompt. A non-empty value replaces the providerâs base prompt for that agent:

```
{
  "agents": {
    "reviewer": { "system": "Review only. Do not modify files." },
  },
}
```
Project instructions, skills, references, and other instruction sources are still added. In a Markdown agent, put this text in the document body instead of a `system` frontmatter field.

### Permissions

`permissions` is an ordered list of matching rules:

```
{
  "agents": {
    "reviewer": {
      "permissions": [
        { "action": "*", "resource": "*", "effect": "deny" },
        { "action": "read", "resource": "src/**", "effect": "allow" },
      ],
    },
  },
}
```
| Field | Meaning | 
|---|---|
| `action` | Tool or permission action. Wildcards are supported. | 
| `resource` | Path, command, agent ID, or other value matched by the action. Wildcards are supported. | 
| `effect` | `allow` ,`ask` , or`deny` . | 

The last matching rule wins, so put broad rules before exceptions. Common actions include:

| Action | Covers | 
|---|---|
| `shell` | Shell commands | 
| `edit` | Edit, write, and patch tools | 
| `subagent` | Child agents | 
| `read` ,`glob` ,`grep` | Local discovery tools | 
| `webfetch` ,`websearch` | Web tools | 
| `skill` | Skill loading | 

For `read`, `edit`, and `external_directory` resources, OpenCode expands `~` and `$HOME`:

`{ "action": "read", "resource": "~/notes/**", "effect": "allow" }`
Shell resources remain raw command text and do not expand those values.

### Steps

`steps` sets a positive maximum number of model steps:

`steps: 8`
On the final step, OpenCode removes tools and asks the model to summarize in text. New user input resets the allowance.

### Hidden

`hidden` removes an agent from normal listings, interactive discovery, and the subagent catalog:

`hidden: true`
This controls visibility, not security. Use permissions to restrict behavior.

### Color

`color` sets the agentâs UI color using a six-digit hex value:

`color: "#ff6b6b"`
### Disabled

`disabled` removes a built-in or custom agent at that point in configuration loading:

```
{
  "agents": {
    "plan": { "disabled": true },
  },
}
```
### Request

`request` accepts per-agent header and JSON body overlays:

```
{
  "agents": {
    "reviewer": {
      "request": {
        "headers": { "x-agent": "reviewer" },
        "body": { "temperature": 0.1 },
      },
    },
  },
}
```
Do not use legacy top-level fields such as `temperature`, `top_p`, `prompt`, `permission`, `tools`, `disable`, or `maxSteps` in new V2 agent configuration.
