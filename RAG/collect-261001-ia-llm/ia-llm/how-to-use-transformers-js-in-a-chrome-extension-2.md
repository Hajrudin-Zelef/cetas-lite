---
id: collect-261001-ia-llm/ia-llm/how-to-use-transformers-js-in-a-chrome-extension-2
title: "how-to-use-transformers-js-in-a-chrome-extension"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "copilot", "inference", "memory"]
source: docs/RAG/collect-261001-ia-llm/how-to-use-transformers-js-in-a-chrome-extension.md
source_anchor: ""
source_lines: [136, 181]
sha256: c4b2845a54cf692155f9358790ec1b92d7d4ab48e6a89971e35de513eb5168b0
---

# how-to-use-transformers-js-in-a-chrome-extension

- `name` ,`description` ,`inputSchema` ,`execute`

Example tools include `get_open_tabs`, `go_to_tab`, `open_url`, `close_tab`, `find_history`, `ask_website`, and `highlight_website_element`.

The core design choice here is to separate internal model messages from UI-facing chat messages:

- Internal model transcript (`messages` ): system/user/tool/assistant turns used for messages in`generator(...)` .
- UI transcript (`chatMessages` ): what the user sees, including streamed assistant text plus tool execution metadata (`tools` ) and performance metrics.

Execution flow:

1. Add user input to `chatMessages` , create a placeholder assistant message, and stream tokens.
2. Parse streamed/final model output with `extractToolCalls.ts` into`{ message, toolCalls }` .
3. Keep the user-visible assistant message as plain text, while tool calls execute in background.
4. Append tool results to the assistant tool metadata and feed results back as the next prompt turn.
5. Repeat until no tool calls remain, then finalize assistant content + metrics.

This keeps user communication clean while preserving a deterministic tool loop in the background.

State placement is another architectural decision that matters a lot in MV3. In this implementation, state is split by lifecycle and access pattern:

- Conversation state: background memory (`Agent.chatMessages` ) for fast turn-by-turn orchestration.
- Tool preferences: `chrome.storage.local` so settings persist across sessions.
- Semantic history vectors: IndexedDB (`VectorHistoryDB` ) for larger local retrieval data.
- Extracted page content: background cache (`WebsiteContentManager` ) keyed by active URL.

As described in section 1.2, keeping conversation history in background gives one canonical state across UI updates. This keeps short-lived state in memory, durable settings in extension storage, and heavy retrieval data in a local database.

You do not need a complex build setup, but MV3 does require predictable outputs for each runtime.

- Multi-entry build in `vite.config.ts` :
- Ensure manifest-aligned output names/paths (`sidebar.html` ,`background.js` ,`content.js` ).
- Keep the content script as a self-contained output to avoid runtime chunk-loading issues.

The goal is simple: one artifact per Chrome entry point, in the exact place `public/manifest.json` expects.

The architecture choice that unlocks this whole project is clear separation of concerns: background owns orchestration and model execution, UI surfaces stay thin, and content scripts handle page access.

This project uses a side panel, but the same approach works for other setups:

- Popup-first assistant: use `action.default_popup` for quick interactions, with background owning conversation state and model execution.
- Side-panel copilot: keep long-running conversations in a persistent panel while background handles tool loops and caching.
- Per-tab agents: keep one agent state per `tabId` in background when each tab should have its own context.
- Hybrid UI (popup + side panel + options page): all UI entry points talk to the same background coordinator and reuse the same message contracts.

The practical rule is simple: decide where state lives (`global`, `tabId`, or site-scoped), keep that state and the model inference in background (basically as background services), and let UI/content runtimes act as focused clients.
