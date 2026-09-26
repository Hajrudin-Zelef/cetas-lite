---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-local-llms-with-claude-code-1
title: "How to Run Local LLMs with Claude Code"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Unsloth"]
dates: []
keywords: ["claude", "agent", "agentic", "attribution", "deepseek", "gguf", "inference", "kv cache", "llama", "llama.cpp"]
source: docs/RAG/clean4/How to Run Local LLMs with Claude Code.md
source_anchor: ""
source_lines: [1, 75]
sha256: 19d38134c65b82062978c23d5a278fc33d526ae88589a57ee94cc867dbbb3ae2
---

# How to Run Local LLMs with Claude Code

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/basics/claude-code.md).
# How to Run Local LLMs with Claude Code
Guide to use open models with Claude Code on your local device.
This step-by-step guide shows you how to connect open LLMs and APIs to Claude Code entirely locally, complete with screenshots. Run using any open model like Qwen3.6, DeepSeek and Gemma.
{% columns %}
{% column width="58.333333333333336%" %}
For this tutorial, we’ll use the open models: [Gemma 4](/docs/models/gemma-4.md) and [Qwen3.5](/docs/models/qwen3.5.md) which are strong agentic & coding models (works on 24GB RAM/unified mem device).
For inference, we'll use [Unsloth Desktop](https://github.com/unslothai/unsloth) and [`llama.cpp`](https://github.com/ggml-org/llama.cpp) which enables you to run/serve LLMs on macOS, Linux, and Windows. You can use any other model. For model quants, we utilize Unsloth [Dynamic GGUFs](/docs/basics/dynamic-3.0-ggufs.md) to run any quantized LLM, while retaining accuracy.
{% endcolumn %}
{% column width="41.666666666666664%" %}
{% endcolumn %}
{% endcolumns %}
Claude Code Setup📖 Setup Local Model Tutorial
## *:claude:* Claude Code Setup
Before setting up our local LLM, we need to install Claude Code. Claude Code is a terminal-based coding agent that understands your codebase and handles complex Git workflows using natural language.
{% tabs %}
{% tab title="macOS, Linux, WSL" %}
#### **Install Claude Code:**
Paste into your terminal to install Claude Code:
```bash
curl -fsSL https://claude.ai/install.sh | bash
```
After install, navigate to your project folder. Then type `claude` into the `shell` to begin.
```bash
cd ~/projects/my-project 
claude
```
{% endtab %}
{% tab title="Windows" %}
#### **Install Claude Code:**
Enter into `PowerShell` to install Claude Code:
```powershell
irm https://claude.ai/install.ps1 | iex
```
After install, navigate to your project folder. Then type `claude` into the `powershell` to begin.

**cd /path/to/your/project**
claude

{% endtab %}
{% endtabs %}
### :detective:Fixing 90% slower inference in Claude Code
{% hint style="warning" %}
Claude Code recently prepends and adds a Claude Code Attribution header, which **invalidates the KV Cache, making inference 90% slower with local models**.
{% endhint %}
The attribution is a line prepended to the **start of the system prompt** (`x-anthropic-billing-header: cc_version=...; cch=...;`) whose value changes on every request, so the whole prompt prefix misses the KV cache each turn.
The simplest fix is to disable it inline when you launch Claude Code, so there is no file to edit:
{% code overflow="wrap" %}
```bash
claude --settings '{"env":{"CLAUDE_CODE_ATTRIBUTION_HEADER":"0","CLAUDE_CODE_ENABLE_TELEMETRY":"0"}}' --model unsloth/gemma-4-26B-A4B-it-GGUF
```
{% endcode %}
{% hint style="info" %}
Recent Claude Code releases also honor `export CLAUDE_CODE_ATTRIBUTION_HEADER=0`; older builds ignored the shell variable, so the `--settings` form above (or the settings file below) is the reliable choice.
{% endhint %}
To make it permanent, add `CLAUDE_CODE_ATTRIBUTION_HEADER` set to 0 inside `"env"` in `~/.claude/settings.json`. For example do `cat > ~/.claude/settings.json` then add the below (when pasted, do ENTER then CTRL+D to save it). If you have a previous `~/.claude/settings.json` file, just add `"CLAUDE_CODE_ATTRIBUTION_HEADER" : "0"` to the "env" section, and leave the rest of the settings file unchanged.
```
{
  "promptSuggestionEnabled": false,
  "env": {
    "CLAUDE_CODE_ENABLE_TELEMETRY": "0",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
    "CLAUDE_CODE_ATTRIBUTION_HEADER" : "0"
  },
  "attribution": {
    "commit": "",
    "pr": ""
  },
  "plansDirectory" : "./plans",
  "prefersReducedMotion" : true,
  "terminalProgressBarEnabled" : false,
  "effortLevel" : "high"
}
```

