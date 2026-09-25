---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-ai-models-with-openclaw
title: "How to Run Local AI Models with OpenClaw"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Intel", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "amd", "deepseek", "gguf", "inference", "intel", "nvidia", "quantization", "qwen"]
source: docs/RAG/clean4/How to Run Local AI Models with OpenClaw.md
source_anchor: ""
source_lines: [1, 72]
sha256: 7aea6875ed9d6e916a2737995d07a468be0e80a91fe679ac1df6a2592624cdb4
---

# How to Run Local AI Models with OpenClaw

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/integrations/openclaw.md).
# How to Run Local AI Models with OpenClaw
Guide to running local LLMs with OpenClaw.
This guide will show how to use open LLMs locally with **OpenClaw by connecting it to Unsloth**. OpenClaw is an **open-source AI agent** interface that connects to a model to run tasks across your project. OpenClaw is able to works with any local model by connecting through **Unsloth’s OpenAI-compatible API**: including DeepSeek, Qwen, Gemma, and more.
OpenClaw acts as the client, while Unsloth loads and serves models via a **local API**.After setup, OpenClaw will run against your local model through Unsloth, letting you use it directly as an **AI agent.** In this tutorial, we'll use [Qwen3.6](/docs/models/qwen3.6.md).
Connecting to OpenClawQuickstart
{% hint style="info" %}
In this tutorial, we’ll use `unsloth/Qwen3.6-27B-GGUF` in Unsloth and access it through OpenClaw. Prefer a different model? Swap in any other model by loading it in Unsloth and updating the configuration.
{% endhint %}
### Installing OpenClaw
{% tabs %}
{% tab title="macOS, Linux, WSL" %}
Install OpenClaw using the official installer:
`curl -fsSL https://openclaw.ai/install.sh | bash`
This sets up OpenClaw and guides you through initial setup.
{% endtab %}
{% tab title="Windows (PowerShell)" %}
Install OpenClaw using the official installer:
`iwr -useb https://openclaw.ai/install.ps1 | iex`
This sets up OpenClaw and guides you through initial setup.
{% endtab %}
{% endtabs %}
### ⚡ Quickstart
After installing OpenClaw, we'll need to install [Unsloth Desktop](/docs/desktop.md) to enable OpenClaw to serve and run inference of local models.
{% stepper %}
{% step %}
#### Download Unsloth
The easiest way to get started is by installing the Unsloth Desktop app. It supports [MacOS](/docs/get-started/install/mac.md), Linux, [Windows](/docs/get-started/install/windows-installation.md), [NVIDIA](/docs/get-started/install/pip-install.md), [AMD](/docs/get-started/install/amd.md), Intel and CPU setups.
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
**Connect OpenClaw.** Run `unsloth start openclaw` to launch OpenClaw with the loaded Unsloth model in a separate managed environment. Your normal OpenClaw configuration is left unchanged.
{% endstep %}
{% endstepper %}
### ⚙️ Launch OpenClaw with `unsloth start`
OpenClaw can connect to a model already running in Unsloth Studio, or start one automatically when Unsloth is not running.
Connect to a running Unsloth instance
Once a model is loaded in Unsloth Studio, run:
```bash
unsloth start openclaw
```

&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
