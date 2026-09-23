---
id: collect-mindstudio/mindstudio/how-to-run-local-ai-models-with-ollama
title: "How to Run Local AI Models with Ollama: A Beginner's Setup Guide for 2026"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Alibaba", "Apple", "Meta", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["llama", "amd", "chatgpt", "deepseek", "embedding", "embeddings", "fine-tuning", "gpu", "gpus", "license", "llama.cpp", "mit license"]
source: docs/RAG/Collect RAG/02_mindstudio/how-to-run-local-ai-models-with-ollama.md
source_anchor: ""
source_lines: [1, 57]
sha256: 9b688c068039d668ae7e37b105d643088c7446fc78ac9ba3fdf56c82649b8b3e
---

# How to Run Local AI Models with Ollama: A Beginner's Setup Guide for 2026

## Metadata

- **Source** : https://www.mindstudio.ai/blog/how-to-run-local-ai-models-with-ollama
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This beginner's guide explains how to run local AI models with Ollama in 2026: installation on any operating system, pulling models like Gemma, Qwen, and LLaMA, core commands, the local API, connecting to other tools, and troubleshooting.

Ollama is an open-source tool (like a package manager for AI models) that handles the messy parts — model quantization, hardware acceleration, and server setup. It runs a local server on port 11434 and exposes a REST API with an OpenAI-compatible `/v1` endpoint, so any application that makes HTTP requests can talk to the local model. It stands out from LM Studio, llama.cpp, Jan, and GPT4All because of its CLI-first design, clean OpenAI-compatible API, active model library, cross-platform support (macOS, Windows, Linux), and automatic GPU acceleration (Apple Silicon, NVIDIA, AMD).

Prerequisites: at least 8GB RAM for 7B models (16GB recommended for 13B), optional but strongly recommended GPU, 10–20GB free disk (models range from ~2GB to 40+GB for 70B), macOS 11+, Windows 10/11, or most Linux distros. No Python/Docker needed.

Installation paths: macOS (official installer or `brew install ollama`), Windows (installer runs as background service), Linux (`curl -fsSL https://ollama.com/install.sh | sh`, with systemd service). GPU support: NVIDIA via CUDA, AMD via ROCm on Linux (RX 5000 series+).

Recommended starter models for 2026: `qwen2.5:7b` (~4.7GB, strong reasoning, bilingual), `llama3.2:3b` (~2GB, fast), `gemma3:4b` (~3.3GB, strong instruction following), coding models `qwen2.5-coder:7b` and `deepseek-coder-v2:16b` (~9.1GB, needs 16GB+ RAM), long-context `llama3.1:8b` (128K context), and for 32GB+ machines `qwen2.5:32b` and `llama3.3:70b`.

Core commands: `ollama pull <model>`, `ollama run <model>`, `ollama list`, `ollama rm <model>`. The API supports `/api/generate`, `/api/chat`, `/api/tags`, and `/api/embeddings`, plus the official Python library (`pip install ollama`) and OpenAI SDK compatibility (base_url `http://localhost:11434/v1`).

Integration: Open WebUI (Docker-based ChatGPT-style interface), Continue (VS Code coding assistant with native Ollama support), LangChain/LlamaIndex (OllamaLLM provider), network exposure via `OLLAMA_HOST=0.0.0.0`. Multimodal models: `llava:7b`, `llama3.2-vision:11b`. Embedding models: `nomic-embed-text` for RAG with ChromaDB/Qdrant/pgvector. Modelfiles allow creating custom model variants with system prompts and parameters without fine-tuning.

Troubleshooting: stalled downloads (resume with Ctrl+C + re-pull), CPU-only mode (check GPU layers loaded), port conflicts (`OLLAMA_HOST=127.0.0.1:11435`), OOM (smaller model or Q4 quantization), poor outputs (tune temperature/num_ctx).

FAQ highlights: Ollama is free (MIT license); 7B vs 70B differences (~8GB vs 48–64GB RAM); CPU-only gives 5–15 tokens/s vs 30–100+ with GPU; Apple Silicon handles local models exceptionally well (recommended models by RAM: 8GB→llama3.2:3b, 16GB→qwen2.5:7b, 32GB→qwen2.5:14b, 64GB+→qwen2.5:32b/llama3.3:70b); data stays fully local with no telemetry.

## Key points

- Ollama is a CLI/API-first package manager for local LLMs with an OpenAI-compatible endpoint on port 11434.
- One command installs, one pulls a model, one runs it — 7B models (Qwen 2.5, Gemma 3, Llama 3) balance quality and hardware needs.
- The REST API and /v1 OpenAI-compatible surface let local models plug into almost any application or framework.
- GPU helps but isn't required; Apple Silicon and NVIDIA/AMD GPUs are supported automatically.
- Modelfiles create customized model variants without fine-tuning.
- Data and prompts never leave the machine — a core privacy advantage over cloud APIs.

## Technical data / figures

| Model | Size | RAM / Notes |
|---|---|---|
| qwen2.5:7b | ~4.7GB | General chat/reasoning, bilingual |
| llama3.2:3b | ~2GB | Fast on any hardware |
| gemma3:4b | ~3.3GB | Instruction following |
| qwen2.5-coder:7b | ~4.7GB | Code |
| deepseek-coder-v2:16b | ~9.1GB | Code, needs 16GB+ RAM |
| llama3.1:8b | ~4.9GB | 128K context |
| qwen2.5:32b / llama3.3:70b | — | Needs 32GB+ RAM |
| Ollama server | localhost:11434 | REST + /v1 endpoints |
| CPU-only throughput | 5–15 tok/s | vs 30–100+ tok/s with GPU |

## Why this source matters for the RAG

Provides the essential foundation for local RAG stacks: installing Ollama, pulling generation and embedding models (`nomic-embed-text`), and integrating with vector databases and frameworks like LangChain/LlamaIndex. It documents concrete commands and configurations directly reusable for building a local, private retrieval pipeline.
