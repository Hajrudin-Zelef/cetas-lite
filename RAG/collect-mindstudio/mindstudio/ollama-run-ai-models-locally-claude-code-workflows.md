---
id: collect-mindstudio/mindstudio/ollama-run-ai-models-locally-claude-code-workflows
title: "How to Use Ollama to Run AI Models Locally for Claude Code Workflows"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Apple", "Mistral", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["claude", "llama", "agent", "agentic", "context window", "cost", "deepseek", "gpu", "gpus", "inference", "latency", "llama.cpp"]
source: docs/RAG/Collect RAG/02_mindstudio/ollama-run-ai-models-locally-claude-code-workflows.md
source_anchor: ""
source_lines: [1, 62]
sha256: a1f971651b4510e5b6c0f84e0fbea5b9b2ac0eecc42aa442b390a482f715712e
---

# How to Use Ollama to Run AI Models Locally for Claude Code Workflows

## Metadata

- **Source** : https://www.mindstudio.ai/blog/ollama-run-ai-models-locally-claude-code-workflows
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how to use Ollama to run open-source AI models locally as a free backend for Claude Code (Anthropic's agentic coding CLI), eliminating API costs for coding workflows.

Ollama is an open-source runtime that manages model files (like Docker images), uses llama.cpp for efficient inference across CPUs, Apple Silicon, and NVIDIA GPUs, exposes a local REST API at http://localhost:11434, and provides an OpenAI-compatible endpoint at http://localhost:11434/v1. That OpenAI-compatible surface is critical for plugging into tools designed for commercial APIs.

Setup: install Ollama (brew / install script / Windows installer), start with `ollama serve`, then pull models. Coding-relevant models: gemma3:12b (~8GB, general coding/tool use), llama3.1:8b (~5GB, fast), llama3.1:70b (~40GB, near-frontier, needs 48GB+ RAM), qwen2.5-coder:14b (~9GB, purpose-built), deepseek-coder-v2:16b (~10GB), phi4:14b (~9GB, surprisingly capable), mistral:7b (~4GB, lightweight). Best balance for M-series MacBook Pro or mid-range GPU: qwen2.5-coder:14b or gemma3:12b.

How Claude Code connects: Claude Code sends requests in Anthropic's API format to https://api.anthropic.com, respecting the ANTHROPIC_BASE_URL environment variable. Since Ollama uses OpenAI's format, a translation proxy is required: LiteLLM. Flow: Claude Code → LiteLLM Proxy (localhost:8000) → Ollama (localhost:11434) → Local Model.

Setup steps: `pip install litellm[proxy]`; create litellm_config.yaml mapping model names (e.g., `claude-3-5-sonnet-20241022`) to `ollama/qwen2.5-coder:14b` and `claude-3-haiku-20240307` to `ollama/gemma3:12b` (naming them after real Anthropic models means Claude Code needs no modification); start with `litellm --config litellm_config.yaml --port 8000`; set `ANTHROPIC_BASE_URL=http://localhost:8000` and `ANTHROPIC_API_KEY=sk-local-dev-key` (any value works).

Model choice matters: tool use support is critical (Qwen 2.5 Coder, Llama 3.1 have the most consistent function calling; Gemma 3 improving; Mistral decent); context window (Qwen 2.5 Coder/Llama 3.1/Gemma 3 = 128K, Mistral 7B = 32K); hardware requirements (8GB VRAM/16GB unified → 7B–8B or 4-bit 13B; 16GB/32GB → 12B–14B; 24GB+ → 30B+; 48GB+ → 70B Q4). If slow, try q4_K_M quantized versions.

Workflow tips: run both servers in the background (tmux/screen); toggle between local and cloud by swapping env vars; use per-project .env files (`ANTHROPIC_BASE_URL`, `ANTHROPIC_API_KEY`).

Troubleshooting: connection errors (test Ollama `/api/tags` and LiteLLM `/health` independently); truncated/mid-output responses (larger context model, increase LiteLLM timeout, smaller files); tool use failures (switch to Qwen 2.5 Coder or Llama 3.1, ensure instruct variant); slow/failed model pulls (14B+ can take 20–40 min; Ollama resumes downloads).

FAQ: LiteLLM is required (Claude Code's Anthropic format vs Ollama's OpenAI format); Ollama supports hundreds of models; local inference is faster than API only on high-end hardware (M3 Max or 24GB+ GPU), smaller quantized models feel snappy; multimodal supported (llava, llama3.2-vision, moondream) though proxy setups may not pass all features; models don't auto-update (re-pull to update); for team/production use a shared Ollama server with access controls is needed — local models remain less capable than frontier models for complex reasoning.

## Key points

- Ollama provides a zero-cost local LLM backend; Claude Code connects via LiteLLM as an Anthropic→OpenAI translation proxy.
- Setup takes ~15–20 minutes: install Ollama, pull a model, install LiteLLM, configure, set two env vars.
- Best coding models: qwen2.5-coder:14b and gemma3:12b (quality/speed balance); llama3.1:70b for near-frontier at 48GB+ RAM.
- Prioritize tool-use support (Qwen 2.5 Coder, Llama 3.1) for agentic coding; check context windows (128K vs 32K).
- Trade-offs: slower responses (5–30 s) and lower capability than frontier models.

## Technical data / figures

| Model | Size | Best for |
|---|---|---|
| gemma3:12b | ~8GB | General coding, tool use |
| llama3.1:8b | ~5GB | Fast responses, decent quality |
| llama3.1:70b | ~40GB | Near-frontier (48GB+ RAM) |
| qwen2.5-coder:14b | ~9GB | Purpose-built coding |
| deepseek-coder-v2:16b | ~10GB | Code gen + debugging |
| phi4:14b | ~9GB | Surprising capability |
| mistral:7b | ~4GB | Lightweight, fast |

| Element | Value |
|---|---|
| Ollama endpoint | localhost:11434 (/v1 OpenAI-compatible) |
| LiteLLM proxy port | localhost:8000 |
| Context windows | 128K (Qwen2.5 Coder/Llama3.1/Gemma3), 32K (Mistral 7B) |
| Latency after setup | 5–30 s per response |
| 14B+ model download time | 20–40 min |

## Why this source matters for the RAG

Provides a complete, tested recipe for connecting local Ollama models to agentic coding tools (Claude Code via LiteLLM), including model selection by tool-use support and context window. Directly relevant for building low-cost local RAG and agent workflows with an OpenAI-compatible surface.
