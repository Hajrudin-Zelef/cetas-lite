---
id: collect-250926-servers-hardware/servers-hardware/how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation-2
title: "how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Unsloth"]
dates: []
keywords: ["claude", "context window", "gguf", "throughput", "tool calling"]
source: docs/RAG/clean4/how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation.md
source_anchor: ""
source_lines: [131, 154]
sha256: 886cfbac2222257171548464a1b15423c4760c343699fcdcc88d168f2e1a532a
---

# how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation

Press "Expand to full monitor" or go to Settings>API Monitor to navigate to the full **API** page where model loading, prompts, responses, token counts, time-to-first-token, throughput, and error messages are displayed in the monitor.

**401 Unauthorized**`Authorization` header is missing or the key is wrong. Keys must be passed as `Authorization: Bearer sk-unsloth-…`. If you lost the key, create a new one from **Settings → API.** Unsloth doesn't show old keys after creation.

**Lost connection to the model server****New Chat** and retry.

**Claude Code shows the default Anthropic model, not my local one** :  check all three env vars are exported in the **same** shell where you run `claude`:

Then run `/model` inside Claude Code to confirm. On Windows PowerShell use `$env:ANTHROPIC_BASE_URL` etc.

**stream: true** **returns a single JSON blob instead of SSE** :  make sure you're hitting the right path (`/v1/messages` or `/v1/chat/completions`) and that your HTTP client is actually consuming the response as a stream, not buffering it.

**I can't find the name of the model to add to opencode (or OpenClaw / any other client)** : ask Unsloth directly. `GET /v1/models` returns the exact model ID you need to plug into the client's "Model ID" field:

You'll get back a JSON payload of the form `{"data": [{"id": "gemma-4-26B-A4B-it-GGUF", ...}]}`. Copy the `id` value, that's the string opencode's **Model ID** field (left column) and OpenClaw's `models[].id` expect. The display name on the right is whatever you want users to see.

**Tool calls aren't executed** :  The model needs to support tool calling for client-side tools (`tools` / `tool_choice`). For Unsloth's built-in tools, remember to set `enable_tools: true` **and** list the ones you want in `enabled_tools` (e.g. `["python", "web_search"]`).

- **My client reports a connection error.** Open the API monitor. No row for the call means it never reached Unsloth, compare your client's base URL against the**Base URL** shown at the top of that page.
- **The reply is cut short.** Check**Context used** on the request in the API monitor. Near 100%, or a stop reason of`length` , means the context window filled rather than the model failing.

Last updated

Was this helpful?
