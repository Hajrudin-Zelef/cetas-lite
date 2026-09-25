---
id: collect-250926-servers-hardware/servers-hardware/openai-codex-with-ollama-ollama-blog
title: "OpenAI Codex with Ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["OpenAI"]
dates: ["2026-01-15"]
keywords: ["context window", "open-weight"]
source: docs/RAG/clean4/openai-codex-with-ollama-ollama-blog.md
source_anchor: ""
source_lines: [1, 39]
sha256: d8750bea7d5068fa8db4c436a7aff20d0e3e52003a0f01030f9e29c8067fde4c
---

# OpenAI Codex with Ollama

## January 15, 2026

Open models can be used with OpenAI’s Codex CLI through Ollama. Codex can read, modify, and execute code in your working directory using models such as `gpt-oss:20b`, `gpt-oss:120b`, or other open-weight alternatives.

### Get started

Install Codex CLI:

```
npm install -g @openai/codex
```
Start Codex with the `--oss` flag:

```
codex --oss
```
By default, Codex will use the local `gpt-oss:20b` model.

**Note:** Codex requires a large context window. We recommend at least 32K tokens. See the documentation for how to adjust context length in Ollama.

### Changing models

You can switch to a different model using the `-m` flag:

```
codex --oss -m gpt-oss:120b
```
### Cloud models

All models on Ollama Cloud work with Codex.

```
codex --oss -m gpt-oss:120b-cloud
```
### Learn more

For more detailed setup instructions and configuration options, see the Codex integration guide.
