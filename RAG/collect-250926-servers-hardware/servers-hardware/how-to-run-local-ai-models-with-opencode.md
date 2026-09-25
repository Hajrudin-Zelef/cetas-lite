---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-ai-models-with-opencode
title: "How to Run Local AI Models with OpenCode"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Intel", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agents", "claude", "deepseek", "gguf", "intel", "quantization", "qwen"]
source: docs/RAG/clean4/How to Run Local AI Models with OpenCode.md
source_anchor: ""
source_lines: [1, 103]
sha256: 1c3834a56190710010dde0996828536652dc4937f59dd06de5fb7b00da0beb0a
---

# How to Run Local AI Models with OpenCode

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/integrations/opencode.md).
# How to Run Local AI Models with OpenCode
Guide to connect open LLMs with OpenCode on your local device.
This guide walks you through connecting **OpenCode** to [Unsloth](https://github.com/unslothai/unsloth) to run open LLMs **entirely locally.** OpenCode is an **open-source AI coding agent** that reads, modifies, and executes code across your project using a connected model. This works with any **local model** exposed through Unsloth’s **OpenAI-compatible API**, including: DeepSeek, Qwen, Gemma, and more.
OpenCode acts as the client, while Unsloth loads and serves models via a local API. After setup, OpenCode connects to Unsloth, where you can select a loaded model and use it as a **coding agent**.
This guide covers both setup options:
* **OpenCode Desktop:** add Unsloth Studio manually as a custom OpenAI-compatible provider.
* **OpenCode CLI:** launch it with `unsloth start opencode` and connect to your local model automatically.
OpenCode SetupQuickstart
{% hint style="info" %}
In this tutorial, we’ll use `unsloth/Qwen3.6-27B-GGUF` loaded in Unsloth and access it directly inside OpenCode. Prefer a different model? Swap in any other model by loading it in Unsloth.
{% endhint %}
### Installing OpenCode Desktop
{% tabs %}
{% tab title="MacOS" %}
#### **Step 1: Download the OpenCode installer for Mac**
Open `opencode.ai/download` in your browser of choice. Scroll down to the **OpenCode Desktop (Beta) ,** and click the `Download` button next to the macOS image name corresponding to your Mac's architecture (Apple Silicon or Intel).
#### Step 2: Install OpenCode
Locate and double click the `OpenCode Desktop.dmg` installer file in your downloads folder.
The installer window will open. Use your mouse to drag the **OpenCode** app icon on top of the **Applications** icon as shown.
#### Step 3: Launch OpenCode
Locate and double click the **OpenCode** icon under the **Applications** folder.
The **OpenCode** desktop app will open and is now ready for your next action.

Download Unsloth
* :apple: [Download for macOS](https://unsloth.ai/download/mac)
* :windows: [Download for Windows](https://unsloth.ai/download/windows)
* :linux: [Download for Linux](https://unsloth.ai/download/linux)
Or, if you prefer manual installation:
**MacOS, Linux, WSL:**
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
**Windows PowerShell:**
```bash
irm https://unsloth.ai/install.ps1 | iex
```
{% endstep %}
{% step %}
#### Install
1. Open the Unsloth installer (`.dmg`, `.exe` files)
2. Drag Unsloth to Applications for Mac or complete setup for Windows.
3. Launch the app and wait for installation to complete
{% endstep %}
{% step %}
#### Choose a model
Open 'Select model' dropdown ontop or 'Model hub' tab, choose a model and a quantization that fits your device, then download it. Once it finishes, start chatting - no setup required.
{% endstep %}
{% step %}
### Unsloth is now ready
To start chatting, type a message and press Enter.
**Connect OpenCode.** Run `unsloth start opencode`. It mints an API key, writes the config, and launches OpenCode against your loaded model.
{% endstep %}
{% endstepper %}
### 🔑 Creating an API key
1. Open the sidebar, click your **Unsloth** avatar at the bottom-left.
2. Go to **Settings** → **API**.
3. Enter a friendly name (e.g. `claude-code-macbook`).
4. *(Optional)* Set an expiry.
5. Click **Create**.
6. **Copy the key immediately.** Unsloth stores only a hash and you won't be able to view it again.
All keys start with the `sk-unsloth-` prefix. Revoke a key from the same page at any time. Requests made with a revoked key will fail with `401 Unauthorized`.
## 🖇️ Connecting Unsloth to OpenCode Desktop
**Opencode** supports any OpenAI-compatible provider, so you can wire Unsloth in as a **Custom** provider. The setup is a one-time flow inside opencode's **Connect provider** dialog.
**1. Open the provider picker.** In opencode, type `/model` (or click the model selector at the bottom of the input).
`.
See the complete [unsloth start](/docs/integrations/unsloth-start.md) reference for model loading, remote Unsloth servers, and advanced options.
### Optional: configure server access
`unsloth run` starts the local API server and loads a model for OpenCode to connect to. You can also customize how the server behaves when starting it.
```bash
# Run the API on port 8888 (--disable-tools passes OpenCode's own tools through)
unsloth run \
--model unsloth/gemma-4-26B-A4B-it-GGUF \
--disable-tools \
-p 8888
```
{% hint style="warning" %}
Use `--disable-tools` when driving OpenCode (or any external coding agent). By default Unsloth Studio runs its own server-side tools, which swallows the agent's tool calls, so OpenCode answers but never edits files. `--disable-tools` switches to passthrough, so OpenCode's own tools are used.
{% endhint %}
Use `-p` to change which port the server runs on.
```bash
# Allow other devices on your network to connect
unsloth run \
--model unsloth/gemma-4-26B-A4B-it-GGUF \
-H 0.0.0.0 \
--disable-tools \
-p 8888
```
This starts the server on `0.0.0.0:8888`, allowing other devices on your local network to connect.
For more advanced runtime configuration, see the main [API tuning](https://unsloth.ai/docs/basics/api#unsloth-run-command) section.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/integrations/opencode.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
