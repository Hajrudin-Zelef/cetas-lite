---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-ai-models-with-hermes-agent-1
title: "How to Run Local AI Models with Hermes Agent"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Intel", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "amd", "deepseek", "gguf", "inference", "intel", "memory", "nvidia", "quantization", "qwen", "reasoning", "research"]
source: docs/RAG/clean4/How to Run Local AI Models with Hermes Agent.md
source_anchor: ""
source_lines: [1, 153]
sha256: 0796560601f06149d30616576221fa93a0c39b912320dcb750988ad4010e6a70
---

# How to Run Local AI Models with Hermes Agent

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/integrations/hermes-agent.md).
# How to Run Local AI Models with Hermes Agent
Guide on using open LLMs with Hermes Agent locally.
This guide enables you to run open LLMs locally with **Hermes Agent** via [**Unsloth**](https://github.com/unslothai/unsloth). Hermes Agent by Nous Research is an **open-source** autonomous AI agent that connects to a model endpoint, executes tasks, and improves over time through memory and learned skills.
{% columns %}
{% column width="58.333333333333336%" %}
Hermes will work with any **local model** exposed through Unsloth’s **OpenAI-compatible API**, including: DeepSeek, Qwen, Gemma, and more. Hermes acts as the agent client, while Unsloth loads and serves models via the [local API](/docs/basics/api.md) entirely offline.
After setup, every prompt sent through Hermes will run using your local model on your device.
{% endcolumn %}
{% column width="41.666666666666664%" %}
{% endcolumn %}
{% endcolumns %}
Setup Hermes🦥 Connect your local model
{% hint style="info" %}
In this tutorial, you’ll install Hermes and configure it to use `unsloth/Qwen3.6-27B-GGUF` served from Unsloth. Prefer a different model? Swap in any other model by loading it in Unsloth and updating the configuration.
{% endhint %}
### :caduceus: Setup Hermes Agent
**Prerequisites:**
The [Hermes](https://github.com/NousResearch/hermes-agent/blob/main/website/docs/getting-started/installation.md) command-line installer supports Linux, macOS, and WSL2. Make sure **Git** is installed; on Linux, also install **curl** and **xz-utils**. The installer automatically provisions `uv`, Python 3.11, Node.js 22, `ripgrep`, and `ffmpeg`.
#### 1. Run the installer
```bash
curl -fsSL https://hermes-agent.nousresearch.com/install.sh | bash
```
The installer:
* Detects your platform and checks dependencies.
* Clones Hermes to `~/.hermes/hermes-agent/`.
* Creates a Python virtual environment and installs the Python dependencies.
* Installs the browser-tool dependencies and Playwright’s Chromium engine.
* Adds the `hermes` command and launches the setup wizard.
{% columns %}
{% column %}
Playwright may request `sudo` to install Chromium’s shared system libraries. Hermes itself does not require root access.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### **2. Reload your shell** so the `hermes` command is on your `PATH`:
{% code title="bash" %}
```bash
source ~/.bashrc
```
{% endcode %}
{% code title="zsh" %}
```bash
source ~/.zshrc
```
{% endcode %}
#### **3. Verify the install:**
```bash
hermes --version
```
If the command resolves, Hermes is installed. Everything lives under `~/.hermes/`:
| Path | What it is |
| --------------------------------------- | ---------------------------------------------- |
| `~/.hermes/config.yaml` | Main settings (model, provider, tools, TTS, …) |
| `~/.hermes/.env` | API keys and other secrets |
| `~/.hermes/hermes-agent/` | The Hermes source + virtualenv |
| `~/.hermes/cron/`, `sessions/`, `logs/` | Runtime data |
| `~/.hermes/skills/` | Installed skills (synced from the Skills Hub) |
{% hint style="info" %}
Full install reference: [hermes-agent.nousresearch.com/docs/getting-started/installation](https://hermes-agent.nousresearch.com/docs/getting-started/installation). If the installer reports a missing prerequisite, install it and re-run the one-liner. The installer is idempotent.
{% endhint %}
### ⚡ Quickstart
After installing Hermes, we'll need to install Unsloth Desktop to enable Hermes to serve and run inference of local models.
{% stepper %}
{% step %}
#### Download Unsloth
The easiest way to get started is by installing the [Unsloth Desktop](/docs/desktop.md) app. It supports [MacOS](/docs/get-started/install/mac.md), Linux, [Windows](/docs/get-started/install/windows-installation.md), [NVIDIA](/docs/get-started/install/pip-install.md), [AMD](/docs/get-started/install/amd.md), Intel and CPU setups.
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
**Connect Hermes.** Run `unsloth start hermes` in your terminal while Unsloth is open. It mints an API key, writes the config, and launches Hermes against your loaded model.
{% endstep %}
{% endstepper %}
### ⚡ Run Hermes Agent with `unsloth start`
To launch Hermes directly with a model, run:
```bash
unsloth start hermes \
--model unsloth/qwen3.8-27B-GGUF-GGUF:UD-Q4_K_XL
--temp 1.0 \
--top-p 0.95 \
--top-k 20 \
--min-p 0.0 \
--reasoning-effort medium
```
{% hint style="success" %}
If there are no settings/sampling flags set, Unsloth automatically selects the best/recommended settings for the model including context length, temperature etc.
{% endhint %}
With a model loaded in Unsloth Studio, run:
```bash
unsloth start hermes
```
Unsloth launches Hermes from a separate managed home with the Unsloth provider, model, and context settings already configured. Your existing Hermes setup is left unchanged.
This managed home is temporary by default. To keep your sessions and state, add `--persist` from your first launch:
```bash
unsloth start hermes --persist
```
To return to your latest session later, run:
```bash
unsloth start hermes --persist --continue
```
To reopen a specific session, use `--resume ` instead.
See the complete [unsloth start](/docs/integrations/unsloth-start.md) reference for model selection, remote connections, and advanced options.
The setup wizard below remains available if you prefer to manage the Hermes provider yourself.
### 🔑 Creating an API key
1. Open the sidebar, click your **Unsloth** avatar at the bottom-left.
2. Go to **Settings** → **API**.
3. Enter a friendly name (e.g. `hermes-agent-macbook`).
4. *(Optional)* Set an expiry.
5. Click **Create**.
6. **Copy the key immediately.** Unsloth stores only a hash and you won't be able to view it again.
All keys start with the `sk-unsloth-` prefix. Revoke a key from the same page at any time. Requests made with a revoked key will fail with `401 Unauthorized`.
### 🦥 Integrate Hermes with Unsloth API
Hermes sends each chat turn to a configured inference provider and connects to **OpenAI-compatible** endpoints. Configure the provider during installation or later in the setup wizard.
**1. Open the setup wizard:**
{% columns %}
{% column %}
```bash
hermes setup
```
Pick **Model & Provider** from the "What would you like to do?" menu to configure only the inference endpoint, or **Full Setup** to walk through everything (TTS, tools, messaging gateway, agent settings).
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
**2. Select the Custom OpenAI-compatible endpoint** when Hermes prompts you for an inference provider.
**3. Fill in the prompts** as Hermes walks through them:
| Prompt | Value |
| ------------------------------------- | ---------------------------------------------------------- |
