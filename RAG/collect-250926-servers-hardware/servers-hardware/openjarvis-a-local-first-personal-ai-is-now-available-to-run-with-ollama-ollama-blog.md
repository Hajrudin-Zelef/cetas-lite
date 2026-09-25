---
id: collect-250926-servers-hardware/servers-hardware/openjarvis-a-local-first-personal-ai-is-now-available-to-run-with-ollama-ollama-blog
title: "OpenJarvis: a local-first personal AI is now available to run with Ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: ["2026-05-28"]
keywords: ["agent", "agents", "cost", "energy", "latency", "memory", "reasoning", "research"]
source: docs/RAG/clean4/openjarvis-a-local-first-personal-ai-is-now-available-to-run-with-ollama-ollama-blog.md
source_anchor: ""
source_lines: [1, 71]
sha256: 3ae71d509727e8f02e67689949dfa8726d177345b3f5372e227549b95e44f1cc
---

# OpenJarvis: a local-first personal AI is now available to run with Ollama

## May 28, 2026

OpenJarvis is an open-source framework for building personal AI agents that run on your own hardware. It’s built by Stanford’s Hazy Research and Scaling Intelligence labs, as part of their “Intelligence Per Watt” research into efficient local AI.

Local models can already handle most day-to-day chat and reasoning, yet most personal AI still sends every request to the cloud. OpenJarvis makes local-first the default. Models run locally, and the cloud is optional. Energy, cost, and latency are tracked alongside accuracy.

Version 1.0 is now available with built-in support for Ollama.

## Get started

**Install Ollama**

Download Ollama for macOS, Windows, or Linux.

**Install OpenJarvis**

On macOS or Linux, the install script sets up everything you need and auto-detects your existing Ollama installation:

```
curl -fsSL https://open-jarvis.github.io/OpenJarvis/install.sh | bash
```
On Windows, run that command inside WSL2, or install the desktop app.

Then run `jarvis` to start.

## Choosing models

The install script sets up a starter model so you can begin right away, but you can choose your own. Pull any model through Ollama and use it for a query:

```
jarvis model pull qwen3.5:35b
jarvis ask -m qwen3.5:35b "Your prompt"
```
To set a default model, add it to `~/.openjarvis/config.toml`:

```
[intelligence]
default_model = "qwen3.5:35b"
preferred_engine = "ollama"
```
## Built-in agents

OpenJarvis ships with ready-to-run presets. Each one bundles an agent with the engines and tools it needs to run.

**Morning briefing**

Generate a morning briefing agent using your calendar, email, and the day’s news:

```
jarvis init --preset morning-digest-mac
jarvis connect gdrive
jarvis digest --fresh
```
**Research across files**

Ask a question to perform research across the web and your local documents, returning an answer with citations:

```
jarvis init --preset deep-research
jarvis memory index ./docs/
jarvis ask "Summarize all emails about Project X"
```
**Local coding agent**

A code agent that writes and runs Python on your machine to get tasks done:

```
jarvis init --preset code-assistant
```
