---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-llms-with-claude-code-2
title: "How to Run Local LLMs with Claude Code"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Unsloth"]
dates: []
keywords: ["claude", "diffusion", "embedding", "gguf", "gpu", "inference", "kv cache", "llama", "llama.cpp", "mcp", "quantization", "reasoning"]
source: docs/RAG/clean4/How to Run Local LLMs with Claude Code.md
source_anchor: ""
source_lines: [76, 238]
sha256: 6ca188eb91e034eaaa552f0d4e3389a3594025c673914a5a6732b4680e78d213
---

# How to Run Local LLMs with Claude Code

## 📖 Quickstart Tutorials
{% columns %}
{% column %}
Before we begin, we firstly need to complete setup for the specific model you're going to use. We use [Unsloth](/docs/new/studio.md) (a web UI) and llama.cpp which are open-source frameworks for running and serving LLMs on your Mac, Linux, Windows devices.
Unsloth also has unique self-healing [tool-calling](/docs/new/studio/chat.md#auto-healing-tool-calling) and [web search](/docs/new/studio/chat.md#code-execution) capabilities. See right for Claude Code connected to Unsloth:
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
Connect Claude Code🦥 Unsloth Tutorial llama.cpp Tutorial
## 🦥 Unsloth Tutorial
For this tutorial, we will serve/connect local models to Claude Code via a UI by using [Unsloth](https://github.com/unslothai/unsloth). Unsloth works on Windows, WSL, Linux and MacOS.
{% columns %}
{% column %}
* Search, download, [run GGUFs](/docs/new/studio.md#run-models-locally) and safetensor models
* [**Self-healing** tool calling](/docs/new/studio.md#execute-code--heal-tool-calling) + **web search**
* [**Code execution**](/docs/new/studio.md#run-models-locally) (Python, Bash)
* [Automatic inference](https://unsloth.ai/docs/desktop#feature-deep-dive) parameter selection (temp, top-p, etc.)
* Fast CPU + GPU inference via llama.cpp
* [Train LLMs](/docs/new/studio.md#no-code-training) 2x faster with 70% less VRAM
See below for install instructions:
{% endcolumn %}
{% column %}
Download Unsloth
* *:apple:* [Download for macOS](https://unsloth.ai/download/mac)
* *:windows:* [Download for Windows](https://unsloth.ai/download/windows)
* *:linux:* [Download for Linux](https://unsloth.ai/download/linux)
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
#### Unsloth is now ready
To start chatting, type a message and press Enter.
* **Connect tools:** [Claude Code](/docs/basics/claude-code.md), [Codex](/docs/basics/codex.md), web search, [MCP](/docs/basics/mcp.md) and more
* **Train models:** Fine-tune text, diffusion, embedding, and more
* **Generate media:** Create and train images, video, TTS locally
{% endstep %}
{% endstepper %}
### Model Loading + API Guide
{% stepper %}
{% step %}
#### Select Model
Before using the API, load a model from the **Select model** dropdown in the top-left corner of the Chat page.
In this guide, we’ll use: `unsloth/gemma-4-26B-A4B-it-GGUF` with the recommended `UD-Q4_K_XL` quantization.
{% endstep %}
{% step %}
#### Test the Model
Before using the Client, send a quick message:
{% hint style="info" %}
This confirms that the model loaded correctly and is ready to respond.
{% endhint %}
{% endstep %}
{% step %}
#### **Unsloth API key**
In Unsloth, open **Settings → API** to view or create your API key.
Treat your API key like a password and avoid exposing it in screenshots or repositories.
{% endstep %}
{% endstepper %}
## ⚙️ Connect Claude Code
Now that we have setup the local LLM for Claude Code, we now configure Claude Code to work with your tool. You can either connect easily with `unsloth start` below or do it [manually](#connect-manually).
### ⚡ Run Claude Code with `unsloth start`
To launch Claude directly with a model, run:
```bash
unsloth start claude \
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
With a model loaded in Unsloth Studio, open your project folder and run:
```bash
unsloth start claude
```
Unsloth sets the local endpoint, API key, model and context length for this launch. Your normal Claude Code configuration is left alone.
Claude Code already keeps conversations in its normal session store, so `--persist` is not needed. Continue your latest session with:
```bash
unsloth start claude --continue
```
See the complete `unsloth start` reference for loading a model from the command line, remote Unsloth servers and advanced options.
The rest of this guide covers the manual `llama.cpp` setup.
#### 🔌 Connect manually
If you prefer to set it up manually, you can start by setting the following environment variables. These variables will not persist between sessions by default.
{% tabs %}
{% tab title="MacOS, Linux, WSL" %}
**Config:** Set the local API URL:
```bash
export ANTHROPIC_BASE_URL="http://localhost:8888"
```
Copy your key from Unsloth Studio → Settings → API (or from the console when you start it with `unsloth run`, where it is printed as `sk-unsloth-...`), then set it.
{% code overflow="wrap" %}
```bash
export ANTHROPIC_AUTH_TOKEN="sk-unsloth-xxxxxxxxxxxx"
```
{% endcode %}
Also set an empty `ANTHROPIC_API_KEY` so Claude Code does not prompt for a cloud key:
```bash
export ANTHROPIC_API_KEY=""
```
Optional: Use the name of the model currently loaded in Unsloth as a default.
```bash
export ANTHROPIC_MODEL="unsloth/gemma-4-26B-A4B-it-GGUF"
```
Use the full model id exactly as it appears in `GET http://localhost:8888/v1/models` (the same string you pass to `claude --model`).
{% endtab %}
{% tab title="Windows" %}
**Config:** Set the local API URL in Powershell:
```powershell
$env:ANTHROPIC_BASE_URL = "http://localhost:8888"
```
Copy your key from **Unsloth Studio → Settings → API**, then set it:
```powershell
$env:ANTHROPIC_AUTH_TOKEN = "sk-unsloth-xxxxxxxxxxxx"
```
**Optional:** Use the name of the model currently loaded in Unsloth to set as a default.
```powershell
$env:ANTHROPIC_MODEL = "gemma-4-26B-A4B-it-GGUF"
```
{% hint style="info" %}
Model name should be the model that is currently loaded in Unsloth Studio.
{% endhint %}
{% endtab %}
{% endtabs %}
### Start Claude Code
Start Claude Code with the model that is currently loaded in Unsloth.
We will use `gemma-4-26B-A4B-it-GGUF`, but you can use any Unsloth compatible model.
```shellscript
claude --model unsloth/gemma-4-26B-A4B-it-GGUF
```
{% hint style="info" %}
For an extra speed boost on local models, you can also launch with `--bare --exclude-dynamic-system-prompt-sections`. See Optional: shrink the system prompt below.
{% endhint %}
Claude Code should open and display the selected model.
{% hint style="warning" %}
See [#fixing-90-slower-inference-in-claude-code](#fixing-90-slower-inference-in-claude-code "mention") first to fix open models being 90% slower due to KV Cache invalidation.
{% endhint %}
Try this prompt to research and rank high-quality SFT datasets:
{% code overflow="wrap" %}
```
You can only work in project/. Do not search for CLAUDE.md — this is it. Use web search to find 10 real instruction/chat/SFT datasets on Hugging Face, briefly summarize your findings and explain why each dataset is relevant for SFT as you research, then create sft_report.md as a polished markdown report containing the rank, dataset name, creator, 3–5 relevant tags, a short plain-English summary, and why it is useful for SFT. Keep everything concise and readable with no giant metadata dumps, pasted raw descriptions, oversized tag lists, or unrelated datasets. Task is complete once sft_report.md contains 10 clean, well-written dataset entries, and finish with: “Successfully finetuned a model with Unsloth!
```
{% endcode %}
