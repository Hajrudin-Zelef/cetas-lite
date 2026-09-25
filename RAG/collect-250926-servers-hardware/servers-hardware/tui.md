---
id: collect-250926-servers-hardware/servers-hardware/tui
title: "TUI"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/clean4/tui.md
source_anchor: ""
source_lines: [1, 65]
sha256: 0c4deb77670486c7725918eec79ba116e398036411c0c22fc64b82b929fd673e
---

# TUI

Run OpenCode from your project, type a request, and press **Enter**.

```
cd ~/code/my-project
opencode
```
`Explain how authentication works in this project`
Press **Shift+Enter**, **Ctrl+Enter**, or **Ctrl+J** for a new line. While OpenCode is working, **Enter** steers the active session and **Alt+Enter** queues the prompt for later.

## Context

Type `@` to search for a file, then select it to attach that file to the prompt. Keep typing after `@` to narrow the results.

`Review @src/auth.ts for error handling problems`
You can include a line or range after the path.

`Explain @src/auth.ts#20-45`
## Shell

Type `!` at the start of an empty prompt to enter shell mode. Enter a command and press **Enter**; press **Esc** to leave shell mode without running one.

`!git status`
Shell commands run in the sessionâs working directory and their output stays in the session.

## Commands

Type `/` to list slash commands. Continue typing to filter the list, then press **Enter** to run the selected command.

`/models`
Common commands include `/new`, `/sessions`, `/models`, `/agents`, `/undo`, `/redo`, and `/editor`. Press **Ctrl+P** to open the command palette for every action available in the current view.

## Side questions

Run `/btw <question>` to ask a quick question about the conversation so far. The answer uses the sessionâs context and current model but is not added to the conversation, so it does not affect later prompts. A `/btw` indicator shows in the prompt footer while it runs, and the answer opens in a dialog; press **C** to copy it.

`/btw why did the test in step 2 fail?`
## Models

Press **Ctrl+X**, then **M** to choose a model, or run `/models`. Press **F2** to cycle through recently used models.

Press **Ctrl+X**, then **A** to choose an agent, or run `/agents`. Press **Shift+Tab** to cycle agents.

## Sessions

Press **Ctrl+X**, then **N** to start a new session. Press **Ctrl+X**, then **L** or run `/sessions` to return to an existing session; **Ctrl+O** opens recent sessions and projects together.

Tabs are enabled by default. Use **Ctrl+Tab** and **Ctrl+Shift+Tab** to move between them, **Ctrl+X** then **W** to close one, and **Ctrl+Shift+T** to reopen the last closed tab.

## History

Press **Ctrl+X**, then **U**, or run `/undo`, to revert the latest user message and the work that followed it. The reverted prompt returns to the composer so you can edit it.

Press **Ctrl+X**, then **R**, or run `/redo`, to restore the reverted work. Sending a new prompt instead accepts the revert and continues from that point.

## Editor

Set `VISUAL` or `EDITOR`, then press **Ctrl+X**, then **E**, or run `/editor`, to edit the current prompt in that editor. Save and exit to return the text to the composer, then press **Enter** to send it.

```
export EDITOR="nvim"
opencode
```
The leader key is **Ctrl+X** by default. See Keybinds to change these shortcuts.
