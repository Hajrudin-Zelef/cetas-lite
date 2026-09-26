---
id: collect-260926-rattrapage/rattrapage/unsloth-on-macos
title: "Install Unsloth on MacOS"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Unsloth"]
dates: []
keywords: ["agent", "agents", "claude", "diffusion", "embedding", "gguf", "llama", "mcp", "quantization"]
source: docs/RAG/lot-rattrapage/fine-tuning/Unsloth on MacOS.md
source_anchor: ""
source_lines: [1, 98]
sha256: 8c750a388a42e521b3edd80e2a4f3ec087a4acd3344cfabde202da6f7c0d9c7d
---

# Install Unsloth on MacOS

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/install/mac.md).
# Install Unsloth on MacOS
Install Unsloth Desktop, Unsloth Studio, or Unsloth Core on macOS.
To install **Unsloth Desktop** to run and train MLX or GGUF models on your Mac, follow the steps below.
{% stepper %}
{% step %}
### Download
Download Unsloth for Mac
{% endstep %}
{% step %}
### Install
1. Open the Unsloth installer (`.dmg` file)
2. Drag Unsloth to **Applications**
3. Launch the app and wait for installation to complete
{% endstep %}
{% step %}
### Choose a model
Open 'Select model' dropdown ontop or 'Model hub' tab, choose a model and a quantization that fits your device, then download it. Once it finishes, start chatting - no setup required.
{% endstep %}
{% step %}
### Unsloth is now ready
To start chatting, type a message and press Enter.
* **Connect tools:** [Claude Code](/docs/basics/claude-code.md), [Codex](/docs/basics/codex.md), web search, [MCP](/docs/basics/mcp.md) and more
* **Train models:** Fine-tune text, diffusion, embedding, and more
* **Generate media:** Create and train images, video, TTS locally
{% endstep %}
{% endstepper %}
### Install Unsloth Studio manually
The command below installs **Unsloth Studio**, the browser-based web UI. It does not install Unsloth Desktop or the code-based Unsloth Core package.
#### Install Unsloth Studio
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
Use the same command to **update**.
### Launch
Every time you want to launch Unsloth again:
```bash
unsloth studio -H 0.0.0.0 -p 8888
```
For detailed Unsloth Studio install instructions and requirements, [view our guide](/docs/new/studio/install.md).
## Uninstall
### Desktop uninstall
1. Close Unsloth if it's running
2. Open **Finder**
3. Go to **Applications**
4. Find Unsloth
5. Right-click → **Move to Trash**
#### Unsloth Studio manual uninstall
The recommended way to fully remove Unsloth Studio is the uninstall script for your OS. It stops any running servers, removes the app, CLI command, launcher data, shortcuts, and platform-specific entries (macOS `.app` bundle + Launch Services; Windows Start Menu + registry + PATH):
```shellscript
curl -fsSL https://raw.githubusercontent.com/unslothai/unsloth/main/scripts/uninstall.sh | sh
```
If you prefer to remove only specific parts:
**1. Remove app only** (keeps history, chats, checkpoints, and exports intact):
* `rm -rf ~/.unsloth/studio/unsloth_studio`
**2. Remove Unsloth entirely** (keeps other Unsloth tools intact):
* `rm -rf ~/.unsloth/studio`
**3. Remove everything Unsloth-related:**
* `rm -rf ~/.unsloth`
{% hint style="warning" %}
Note: Step 3 deletes everything in history, chats, model checkpoints, and exports. This cannot be undone.
{% endhint %}
**4. Remove shortcuts and symlinks:**
```shellscript
rm -rf ~/Applications/Unsloth\ Studio.app ~/Desktop/Unsloth\ Studio
```
**5. Remove the CLI command:**
* `rm -f ~/.local/bin/unsloth`
{% hint style="info" %}
Note: Steps 1-5 dont touch your downloaded HF model files. See Deleting cached HF model files below if you want to reclaim that space.
{% endhint %}
### **Deleting model files**
You can delete old model files either from the bin icon in model search or by removing the relevant cached model folder from the Hugging Face cache directory.
The default cache location is:
```bash
~/.cache/huggingface/hub/
```
If `HF_HUB_CACHE` or `HF_HOME` is set, use that location instead. You can check with:
```bash
echo ${HF_HUB_CACHE:-${HF_HOME:-${XDG_CACHE_HOME:-$HOME/.cache}/huggingface}/hub}
```
To delete a specific model, remove its folder (e.g. `models--unsloth--Llama-3.1-8B-bnb-4bit`) from the cache directory. To clear all cached models:
```bash
rm -rf ~/.cache/huggingface/hub/
```
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/install/mac.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
