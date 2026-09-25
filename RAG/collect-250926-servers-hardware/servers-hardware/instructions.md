---
id: collect-250926-servers-hardware/servers-hardware/instructions
title: "Instructions"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "mcp"]
source: docs/RAG/clean4/instructions.md
source_anchor: ""
source_lines: [1, 83]
sha256: 67a572f2f56447f0c7993f7dbbf467083f6620489075054481610c31b1514e52
---

# Instructions

Add an `AGENTS.md` file to give OpenCode persistent project guidance. Use it for build commands, architecture notes, code conventions, and verification requirements.

Commit project instruction files so everyone working in the repository receives the same guidance.

## Scope

Place `AGENTS.md` in the directory where its guidance should apply. OpenCode loads the global file followed by every `AGENTS.md` from the current workspace directory toward the home directory. For workspaces outside the home directory, it stops at the project root.

```
~/.config/opencode/AGENTS.md
~/code/my-project/AGENTS.md
~/code/my-project/packages/AGENTS.md
~/code/my-project/packages/web/AGENTS.md  â current workspace
```
In this example, all four files are loaded. They are combined in this order:

```
~/.config/opencode/AGENTS.md
packages/web/AGENTS.md
packages/AGENTS.md
AGENTS.md
```
Keep guidance that applies everywhere in the global file. Put repository-wide guidance at the project root and more specific guidance closer to the code it covers. OpenCode combines the files and does not resolve conflicts between them.

If the workspace is outside the project root, only the global file is loaded. Set `OPENCODE_DISABLE_PROJECT_CONFIG=1` to skip project `AGENTS.md` discovery without disabling the global file.

## Discovery

Instruction files below the workspace are discovered as the agent explores the project. Reading a file or listing a directory loads any `AGENTS.md` files between that target and the workspace.

```
my-project/                         â current workspace
âââ AGENTS.md                       loaded initially
âââ packages/
    âââ web/
        âââ AGENTS.md               loaded when this area is read
        âââ src/
            âââ app.ts              read target
```
Nested files are loaded nearest-first and deduplicated while their instruction entry remains in model-visible history. Reading the same area again does not normally inject them again. If compaction or a revert removes that entry, a later read can load the file again.

Edits to a nested file are not detected automatically after it loads. Start a new session when updated text must apply immediately.

## Ordering

The selected agent or provider system prompt is sent first. OpenCode then assembles initial instructions in this order:

```
1. Agent or provider system prompt
2. Built-in environment and date context
3. Code Mode tool guidance, when enabled
4. Global and project AGENTS.md files
5. Available skill, reference, and MCP guidance
6. Session-specific instruction entries supplied through the API
```
These sources are combined rather than used as overrides. Nested `AGENTS.md` files discovered later are added to session history in discovery order.

## Updates

Edit a global or upward-discovered `AGENTS.md` while a session is running to update its guidance.

`$ printf '\n- Run the integration suite before committing.\n' >> AGENTS.md`
Before the next model request, OpenCode detects the change and adds an instruction update before delivering pending input:

```
AGENTS.md changes
â instruction update
â next prompt
```
- Removing every ambient `AGENTS.md` tells the session that the previous ambient instructions no longer apply.
- A temporary read failure preserves the last known instructions instead of treating them as deleted.
- Moving a session retains its instruction state, so guidance at the destination is introduced as an update.
- Committing a session revert clears its instruction state and reloads instructions before the next prompt.

Instruction values remain privileged. Clients can see which sources changed, but not their contents.

## Configuration

The V2 config schema accepts an `instructions` array, but V2 does not currently resolve its files, glob patterns, or URLs. This configuration does not add instructions to the model yet:

Use `AGENTS.md` for active V2 instructions. See Config for configuration locations and precedence.
