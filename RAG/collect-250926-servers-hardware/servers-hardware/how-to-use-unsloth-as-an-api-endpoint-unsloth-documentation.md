---
id: collect-250926-servers-hardware/servers-hardware/how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation
title: "how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Intel", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agentic", "amd", "claude", "context window", "gguf", "gpu", "inference", "intel", "latency", "llama", "memory"]
source: docs/RAG/clean4/how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 154]
sha256: 94016ea37334d9efd7ba663c46e66d1e433db3220f4827d04b9a3cc4fe547f7d
---

# how-to-use-unsloth-as-an-api-endpoint-unsloth-documentation

You can run **local LLMs** with tools like Claude Code and Codex by connecting those tools to Unsloth’s **OpenAI-compatible API endpoint**. This lets you run models like Qwen and Gemma locally for agentic coding. Unsloth also has beneficial features such as self-healing **tool calling**, **code execution**, and **web search**.

Unsloth makes it easy to deploy a fast API inference endpoint that provides:

Models loaded in Unsloth (including GGUFs) are exposed as an **authenticated API** via `llama-server`. A long API key is generated for security reasons like how OpenAI provides one.

Your **local models** can then be used directly in your preferred AI agent, SDK, or chat client. Unsloth speaks two dialects on the same port. Both support streaming, tool calling (OpenAI `tools` / Anthropic `tools`), and vision inputs:

Whether your model is running through Unsloth’s inference or your own remote OpenAI-compatible endpoint, you can give it access to Unsloth’s full suite of tools, including web search, code execution, deep research, and more.

- **Anthropic-compatible** **/v1/messages**
- **OpenAI-compatible** **/v1/chat/completions**
**/v1/responses**

The easiest way to get started is by installing the Unsloth Desktop app. It supports MacOS, Linux, Windows, NVIDIA, AMD, Intel and CPU setups.

Or, if you prefer manual installation:

**MacOS, Linux, WSL:**

`curl -fsSL https://unsloth.ai/install.sh | sh`
**Windows PowerShell:**

`irm https://unsloth.ai/install.ps1 | iex`
To start chatting, type a message and press Enter.

- **Create an API key.** Click your**Unsloth** avatar in the bottom-left →**Settings** →**API** → type a key name →**Create** . Copy the`sk-unsloth-…` value that appears. Unsloth only shows it once.
- **Point your client at Unsloth.** Use`http://localhost:PORT` as the base URL and your`sk-unsloth-…` key for auth. Jump to the recipe for your tool below.

1. Open the sidebar, click your **Unsloth** avatar at the bottom-left.
2. Go to **Settings** →**API** (globe 🌐 icon).
3. Enter a friendly name (e.g. `claude-code-macbook` ). Set an expiry (optional)
4. Click **Create** .
5. **Copy the key.** Unsloth stores only a hash and you won't be able to view it again.

All keys start with the `sk-unsloth-` prefix. Revoke a key from the same page at any time. Requests made with a revoked key will fail with `401 Unauthorized`.

Treat your API key like a password. Anyone with the key and network access to your Unsloth instance can send requests to your loaded model.

1. **Install or update Unsloth Studio.** Earlier versions don't expose the external API. See Installation.
2. **Load a GGUF model.** load a GGUF model using the run command. This will also load the UI on the default port. The endpoint URL and API Key will be printed out to the console , ready for you to be used with your client of choice.

Adjust settings as necessary.

You can load a model and have an API key created for you automatically using the `unsloth` CLI tool. When the model finishes loading, the endpoint URL and API key are printed to your console. Copy them into your client of choice and you're ready to go.

Make sure you're on a recent version of Unsloth Studio as earlier versions don't expose the external API. See installation.

Open a terminal and load a GGUF model:

This starts the server on the default port, loads the UI, and prints your endpoint URL and API key.

You can point at a model in a few different ways. Pick the one you find easiest:

You don't need any of this for a basic load, but `unsloth run` supports many llama-server runtime flags for customizing performance, memory usage, context length, generation behavior, networking, and tool access.

Additional flags are forwarded directly to the underlying inference server, and your values override Unsloth's defaults. If there are no settings/sampling flags set, Unsloth automatically selects the best/recommended settings for the model including context length, temperature etc.

Some reasoning-capable models support additional flags for controlling thinking and reasoning behavior.

Reasoning effort and flags depends on what the model supports.

Sampling settings control how creative, focused, or deterministic the model behaves during generation.

Lower temperature values usually produce more stable outputs, while top-p, top-k, min-p, and repeat penalty settings further control token selection and repetition.

Useful if you're working with large projects, long chats, or agent workflows that need more memory.

By default, Unsloth only runs locally on your machine. You can expose the API to other devices on your network by binding to `0.0.0.0`.

Control whether tools like web search and code execution are exposed by the inference server.

Unsloth supports most llama-server runtime flags, including context sizing, GPU layers, threading, sampling, networking, and tool configuration.

See the llama-server documentation for the full list of supported runtime flags.

`unsloth run` controls whether server-side tools (web search, code execution, etc.) are exposed by the inference server. Defaults are based on the bind address:

**127.0.0.1**
 **(localhost)** — tools**on** by default. Only your machine can reach the server.
**0.0.0.0**
 **or any non-loopback address** — tools**off** by default. A leaked API key on a network-exposed server means arbitrary code execution on the host.

**Flags:**

- `--enable-tools` /`--disable-tools` — force on or off. On`0.0.0.0` ,`--enable-tools` shows a y/N security prompt.
- `--yes` /`-y` — skip the prompt (for automation).

The resolved policy is a process-level hard override — individual requests cannot bypass it via `enable_tools=true` in the request body.

Unsloth exposes these endpoints on whichever port it booted on (typically `http://localhost:8000` or `http://localhost:8888`):

`POST /v1/messages`

Anthropic Messages API

Claude Code, Anthropic SDK, OpenClaw, anything that speaks Anthropic

`POST /v1/chat/completions`

OpenAI Chat Completions API

OpenAI SDK, opencode, Cursor, Continue, Cline, Open WebUI, curl, etc.

`GET /v1/models`

OpenAI models list

List the models currently loaded in Unsloth

Authenticate with an `Authorization: Bearer sk-unsloth-…` header on every request.

Unsloth enables you run local LLMs via most frameworks including Claude Code, Codex, OpenClaw, OpenCode and more. Click the specific tools below for a guide:

To reach this endpoint from another machine, launch with `unsloth studio --secure`. Unsloth stays bound to localhost and publishes itself on a free Cloudflare HTTPS URL; use that URL in place of `http://127.0.0.1:8888` as your client's base URL. Note that server-sent events do not survive a Cloudflare quick tunnel, so set `stream: false` when calling over one.

Both endpoints support function / tool calling in their native format, plus an Unsloth-specific shorthand for Unsloth's built-in tools.

**OpenAI-style tools:** send `tools` and `tool_choice` to `/v1/chat/completions` exactly as you would with OpenAI. Claude Code (via `/v1/messages`), opencode, Cursor, Continue, and Cline all work out of the box.

**Anthropic-style tools:** send `tools` (with `input_schema`) and `tool_choice` to `/v1/messages` exactly as you would with Claude.

Unsloth server side tools: Unsloth can execute Python, web search, and bash *server-side* and stream the results back as `tool_result` events. Opt in by adding these extra fields to either endpoint:

The model sees each tool's output on its next turn. For deeper coverage (schemas, streaming events, chaining), see .

Every call through this endpoint is listed live in Studio, in two places:

The API monitor side panel opens itself in the corner as soon as API-key traffic arrives. It summarizes the active model, live requests, errors and average latency.

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
