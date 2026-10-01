---
id: collect-261001-ia-llm/ia-llm/how-to-use-transformers-js-in-a-chrome-extension-1
title: "how-to-use-transformers-js-in-a-chrome-extension"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agentic", "embedding", "embeddings", "inference", "memory", "parameters", "reasoning", "tool calling"]
source: docs/RAG/collect-261001-ia-llm/how-to-use-transformers-js-in-a-chrome-extension.md
source_anchor: ""
source_lines: [1, 135]
sha256: f47d952628802685959788b81969f47e39364056f59fc88d7d742f429fb5ae0d
---

# how-to-use-transformers-js-in-a-chrome-extension

While building it, we ran into several practical observations about Manifest V3 runtimes, model loading, and messaging that are worth sharing.

This guide is for developers who want to run local AI features in a Chrome extension with Transformers.js under Manifest V3 constraints.

By the end, you will have the same architecture used in this project: a background service worker that hosts models, a side panel chat UI, and a content script for page-level actions.

In this guide, we will recreate the core architecture of **Transformers.js Gemma 4 Browser Assistant**, using the published extension as a reference and the open-source codebase as the implementation map.

- Live extension: Chrome Web Store
- Source code: github.com/nico-martin/gemma4-browser-extension
- End result: a background-hosted Transformers.js engine, a side panel chat UI, and a content script for page extraction and highlighting.

Before diving in, a quick scope note: I will not go deep on the React UI layer or Vite build configuration. The focus here is the high-level architecture decisions: what runs in each Chrome runtime and how those pieces are orchestrated.

If Manifest V3 is new to you, read this short overview first: What is Manifest V3?.

In MV3, your architecture starts in `public/manifest.json`. This project defines three entry points:

- `background.service_worker = background.js` , built from`src/background/background.ts` .
- `side_panel.default_path = sidebar.html` , built from`src/sidebar/index.html` .
- `content_scripts[].js = content.js` with`matches: http(s)://*/*` and`run_at: document_idle` , built from`src/content/content.ts` .

The background service worker also handles `chrome.action.onClicked` to open the side panel for the active tab.
Related entry point to know: a popup can be defined with `action.default_popup` and works well for quick actions. This project uses a side panel for persistent chat, but the orchestration pattern is the same.

The key design decision is to keep heavy orchestration in the background and keep UI/page logic thin.

- Background (`src/background/background.ts` ) is the control plane: agent lifecycle, model initialization, tool execution, and shared services like feature extraction.
- Side panel (`src/sidebar/*` ) is the interaction layer: chat input/output, streaming updates, and setup controls.
- Content script (`src/content/content.ts` ) is the page bridge: DOM extraction and highlight actions.

One practical consequence of this division is that the conversation history also lives in background (`Agent.chatMessages`): the UI sends events like `AGENT_GENERATE_TEXT`, background appends the message, runs inference, then emits `MESSAGES_UPDATE` back to the side panel.

This split avoids duplicate model loads, keeps the UI responsive, and respects Chrome's security boundaries around DOM access.

Once runtimes are separated, messaging becomes the backbone. In this project, all messages are typed through enums in `src/shared/types.ts`.

- Side panel -> background (`BackgroundTasks` ):
  - `CHECK_MODELS` ,`INITIALIZE_MODELS`
  - `AGENT_INITIALIZE` ,`AGENT_GENERATE_TEXT` ,`AGENT_GET_MESSAGES` ,`AGENT_CLEAR`
  - `EXTRACT_FEATURES`
- Background -> side panel (`BackgroundMessages` ):
  - `DOWNLOAD_PROGRESS` ,`MESSAGES_UPDATE`
- Background -> content (`ContentTasks` ):
  - `EXTRACT_PAGE_DATA` ,`HIGHLIGHT_ELEMENTS` ,`CLEAR_HIGHLIGHTS`

The orchestration rule is simple: the background is the single coordinator; side panel and content script are specialized workers that request actions and render results.

Typical request flow:

1. Side panel sends `AGENT_GENERATE_TEXT` .
2. Background appends to `Agent.chatMessages` and runs model/tool steps.
3. Background emits `MESSAGES_UPDATE` .
4. Side panel re-renders from the updated message list.

In `src/shared/constants.ts`, this extension uses two model roles:

- TextGeneration / LLM: `onnx-community/gemma-4-E2B-it-ONNX` (`text-generation` ,`q4f16` )
- VectorEmbeddings: `onnx-community/all-MiniLM-L6-v2-ONNX` (`feature-extraction` ,`fp32` )

The split is intentional: Gemma 4 handles reasoning/tool decisions, while MiniLM generates vector embeddings for the semantic similarity search in `ask_website` and `find_history`.

All inference runs in background (`src/background/background.ts`):

- text generation via `pipeline("text-generation", ...)` with consistent KV Caching enabled by our new`DynamicCache` class
- embeddings via `pipeline("feature-extraction", ...)` plus vector normalization

This gives a single model host for all tabs/sessions, avoids duplicate memory usage, and keeps the side panel UI responsive. Because models are loaded from the background service worker, artifacts are cached under the extension origin (`chrome-extension://<extension-id>`) rather than per-website origins, which gives one shared cache for the whole extension install.

MV3 lifecycle note: service workers can be suspended and restarted, so model runtime state should be treated as recoverable and re-initialized when needed.

The model lifecycle is explicit:

- `CHECK_MODELS` inspects what is already cached and estimates remaining download size.
- `INITIALIZE_MODELS` downloads/initializes models and emits`DOWNLOAD_PROGRESS` to the UI.
- Long-lived instances are reused after setup:
  - generation pipeline in `src/background/agent/Agent.ts`
  - embedding pipeline in `src/background/utils/FeatureExtractor.ts`
- generation pipeline in 

Permissions and privacy are part of the architecture, not a checkbox at the end. In this project, `public/manifest.json` asks for `sidePanel`, `storage`, `scripting`, and `tabs`, plus `host_permissions` for `http(s)://*/*`:

- `sidePanel` : required to open and control the side panel UX.
- `storage` : required to persist tool/settings state across sessions.
- `tabs` +`scripting` : required for tab-aware tools and page-level actions.
- `host_permissions` on`http(s)://*/*` : required because content extraction/highlighting is designed to work on arbitrary websites.

Why keep this narrow: permissions define user trust and Chrome Web Store review risk. Request only what your features actually need, and state clearly that inference runs locally in the extension runtime so users understand where their data is processed.

Before the execution loop, it helps to understand how model tool calling works (the basis for any agentic workflow). You pass messages plus a tool schema (`name`, `description`, and `parameters`), and Transformers.js formats the actual prompt from those inputs using the model's chat template. Because chat templates are model-specific, the exact tool-call format depends on the model you use. With Gemma-4-style templates, the model emits a special tool-call token block when it decides to call one.

```
import { pipeline } from "@huggingface/transformers";
const generator = await pipeline(
  "text-generation",
  "onnx-community/gemma-4-E2B-it-ONNX",
  {
    dtype: "q4f16",
    device: "webgpu",
  },
);
const messages = [{ role: "user", content: "What's the weather in Bern?" }];
const output = await generator(messages, {
  max_new_tokens: 128,
  do_sample: false,
  tools: [
    {
      type: "function",
      function: {
        name: "getWeather",
        description: "Get the weather in a location",
        parameters: {
          type: "object",
          properties: {
            location: {
              type: "string",
              description: "The location to get the weather for",
            },
          },
          required: ["location"],
        },
      },
    },
  ],
});
```
At generation time, the model can emit output like:

```
<|tool_call>call:getWeather{location:<|"|>Bern<|"|>}<tool_call|>
```
That is exactly why this project has a normalization layer (`webMcp`) and a parser (`extractToolCalls`): model output must be converted into deterministic tool executions.

`src/background/agent/webMcp.tsx` normalizes extension tools into a model-friendly shape:

