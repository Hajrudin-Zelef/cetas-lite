---
id: collect-260926-rattrapage/rattrapage/how-to-run-local-llms-with-openai-codex-1
title: "How to Run Local LLMs with OpenAI Codex"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Microsoft", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agentic", "chatgpt", "claude", "deepseek", "diffusion", "embedding", "gguf", "gpu", "inference", "llama", "llama.cpp"]
source: docs/RAG/lot-rattrapage/ai-llm/How to Run Local LLMs with OpenAI Codex.md
source_anchor: ""
source_lines: [1, 115]
sha256: ba712272693b6db07d7bb3f43f744743fa6a62366428c599162dfb75f4702399
---

# How to Run Local LLMs with OpenAI Codex

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/basics/codex.md).
# How to Run Local LLMs with OpenAI Codex
Use open models with OpenAI Codex on your device locally.
This step-by-step guide shows you how to connect open LLMs and APIs to OpenAI Codex **entirely locally**, complete with screenshots. Codex only needs a local endpoint that speaks the OpenAI Responses API. Run using any open model like Qwen, DeepSeek, Gemma, and more.
For this tutorial, we’ll use the open models: [Gemma 4](/docs/models/gemma-4.md) and [Qwen3.5](/docs/models/qwen3.5.md) which are strong agentic & coding models (works on 24GB RAM/unified mem device). For inference, we'll use [Unsloth Studio](https://github.com/unslothai/unsloth) and [`llama.cpp`](https://github.com/ggml-org/llama.cpp) enables you to run/serve LLMs on macOS, Linux, and Windows. You can swap in any other model, just update the model names in your scripts and Codex config.
Setup Codex📖 Setup Local Model Tutorial
For model quants, we'll use Unsloth [**Dynamic GGUFs**](/docs/basics/dynamic-3.0-ggufs.md) so you can run quantized GGUF models while retaining as much accuracy as possible.
{% hint style="info" %}
Codex has changed quite a lot since Jan 2026. It now uses the [**OpenAI Responses API**](https://platform.openai.com/docs/api-reference/responses) **exclusively**, and Chat Completions support has been deprecated. [Unsloth Studio](#unsloth-tutorial) supports both, so we'll use `wire_api = "responses"` throughout this guide.
{% endhint %}
### :openai: Setup Codex
[Codex](https://github.com/openai/codex) is OpenAI's official coding agent that runs locally. While designed for ChatGPT, it supports **custom API endpoints**, which makes it work for local LLMs. We'll later point it to Unsloth Studio's `/v1/responses` endpoint once Unsloth is up.
{% tabs %}
{% tab title="Linux / WSL" %}
Run in your terminal:
```bash
apt update
sudo apt install nodejs npm -y
npm install -g @openai/codex
```
{% endtab %}
{% tab title="Windows" %}
Run in Windows Powershel:
```powershell
winget install --id OpenAI.Codex
```
{% hint style="info" %}
**Prefer the Codex desktop app?** Install from the Microsoft Store:
```powershell
winget install --id 9PLM9XGG6VKS --source msstore
```
Or via the [Microsoft app Store](https://apps.microsoft.com/detail/9plm9xgg6vks). The app reads the same `%USERPROFILE%\.codex\config.toml`, so the provider config we set up later applies either way.
{% endhint %}
{% hint style="info" %}
**Prefer WSL?** Open PowerShell as admin, run `wsl --install`, restart, then follow the Linux tab above inside Ubuntu. You'll need a small networking trick to reach Unsloth on the Windows host - see the WSL hint in Connect Codex to Unsloth.
{% endhint %}
{% endtab %}
{% tab title="MacOS" %}
Run in your terminal:

bash brew install --cask codex

{% endtab %}
{% endtabs %}
That's it for the install - **don't run `codex` yet**. Running it bare drops you into OpenAI's "Sign in with ChatGPT" picker (which is modal - there's no escape hatch). Once we wire up a local profile,\
`codex --oss --profile unsloth_api` or `codex --oss --profile llama_cpp` skips that screen entirely because custom providers default to `requires_openai_auth = false`. Start the local model server first, then launch Codex against it.
## 📖 Quickstart Tutorials
Before we begin, we firstly need to complete setup for the specific model you're going to use. We use [Unsloth](https://unsloth.ai/docs/new/studio) (a web UI) and llama.cpp which are open-source frameworks for running and serving LLMs on your Mac, Linux, Windows devices.
{% columns %}
{% column %}
Before we begin, we firstly need to complete setup for the specific model you're going to use. We use [Unsloth](/docs/new/studio.md) (a web UI) and llama.cpp which are open-source frameworks for running and serving LLMs on your Mac, Linux, Windows devices.
Unsloth also has unique self-healing [tool-calling](/docs/new/studio/chat.md#auto-healing-tool-calling) and [web search](/docs/new/studio/chat.md#code-execution) capabilities. See right for Claude Code connected to Unsloth:
{% endcolumn %}
{% column %}
🦥 Unsloth Tutorial🦙 llama.cpp Tutorial
## 🦥 Unsloth Tutorial
For this tutorial, we will serve/connect local models to Claude Code via a UI by using [Unsloth](https://github.com/unslothai/unsloth). Unsloth works on Windows, WSL, Linux and MacOS.
{% columns %}
{% column %}
* Search, download, [run GGUFs](/docs/new/studio.md#run-models-locally) and safetensor models
* [**Self-healing** tool calling](/docs/new/studio.md#execute-code--heal-tool-calling) + **web search**
* [**Code execution**](/docs/new/studio.md#run-models-locally) (Python, Bash)
* [Automatic inference](https://unsloth.ai/docs/desktop#feature-deep-dive) parameter tuning (temp, top-p, etc.)
* Fast CPU + GPU inference via llama.cpp
* [Train LLMs](/docs/new/studio.md#no-code-training) 2x faster with 70% less VRAM
See below for install instructions:
{% endcolumn %}
{% column %}
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
/models` to confirm the exact string. |
| `oss_provider` |

Sets unsloth\_api as the default local provider when launching Codex with --oss.

