---
id: collect-korben/korben/odysseus-ia-locale-mac
title: "Odysseus - L'IA auto-hébergée de PewDiePie enfin rapide sur Mac"
domain: korben
role: reference
task: article
actors: ["Alibaba", "Apple", "OpenAI", "OpenRouter"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "gpu", "inference", "license", "llama", "llama.cpp", "moe", "parameters", "qwen", "research", "tool calling"]
source: docs/RAG/Collect RAG/01_korben/odysseus-ia-locale-mac.md
source_anchor: ""
source_lines: [1, 54]
sha256: 128806c1244b9f362f3ed06a9d6d4c35d0b1f2da159d5eaeec88f1eac8aa6350
---

# Odysseus - L'IA auto-hébergée de PewDiePie enfin rapide sur Mac

## Metadata

- **Source** : https://korben.info/odysseus-ia-locale-mac.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article is a hands-on tutorial and review by Korben of Odysseus, the self-hosted AI workspace released by PewDiePie on GitHub at the end of May. It bundles chat, agents, web search, email, calendar, notes and documents into a self-hosted cockpit where models and history can stay on your own machine. It also includes a Markdown editor with version history, integrated image generation, and a Compare mode to pit two models side by side blindly (excluding third-party services tied to web search, email and calendar).

Korben wanted to install it on his Mac but hit a trap that nearly ruined the experience. He explains how to run it correctly on Mac plus some use cases. First, Odysseus evolves quickly: the `main` branch is stable and tested, but the `dev` branch can break at any time, so clone the stable branch explicitly with `git clone https://github.com/odysseus-dev/odysseus.git`, `cd odysseus`, `git checkout main`.

Step 1: forget Docker on Mac. On Apple Silicon, Docker cannot access the Metal GPU, so models run on CPU and are painfully slow; a native install is required. Step 2: install Odysseus in two commands with `./start-macos.sh`, which installs dependencies and starts the server. On his Mac, AirPlay was occupying port 7000, so the interface opened automatically at `http://127.0.0.1:7860`. On first launch, you set a new password for the admin account. For remote access, use an SSH tunnel (`ssh -L 7860:127.0.0.1:7860 user@server`) rather than exposing the port.

Step 3: a fast model like Qwen. He chose Qwen3-30B-A3B, which activates only part of its parameters per response, well suited to the unified RAM of a large Mac. The Cookbook feature, which downloads models suited to your config, correctly recognized his M4 Max with 128 GB but then crashed on llama.cpp due to an invalid `--flash-attn auto` option; after removing it, the server stayed unresponsive while the interface still showed it as active, so he gave up on llama.cpp. The reliable path was Ollama: `ollama pull qwen3:30b-a3b`, then `OLLAMA_HOST=0.0.0.0:11434 ollama serve`, then add `http://localhost:11434/v1` in Odysseus settings. Version 1.0.2 also offers MLX in the Cookbook, but Ollama remained the most reliable. You can also connect an external API like OpenAI or OpenRouter.

He describes three use cases: a morning watch round (classify links, spot duplicates, propose three angles); one folder per article (brief, primary links, quotes, screenshots, versions, with the Markdown editor keeping history and Compare mode for intros/outlines); and a local editorial secretary (monitoring a security-alert inbox, preparing an urgent topic list, calendar reminders). But he would never let it perform automatic actions like sending emails or managing his calendar, as it is not safe enough. He asked the agent to find and summarize a document; it spent nearly 1,800 tokens guessing how to call `manage_documents` without ever launching the search tool. The calendar ended on a `Could not extract JSON`, and web search found the right GitHub repo but the model chose a fake one. The highlight was Deep Research: after 4 minutes it generated 13 pages from 6 retained sources; it gave the right conclusion but totally invented a video, a brand number and several statistics. He concludes Odysseus is a great cockpit with several buttons not yet wired to the engine: chat via Ollama and documents are usable, but the Cookbook, agents, calendar and factual search remain too fragile. The project is young and has huge potential. The code is on GitHub under the AGPL license.

## Key points

- Odysseus is PewDiePie's self-hosted AI workspace (chat, agents, web search, email, calendar, notes, documents, Markdown editor, image generation, blind Compare mode).
- On Apple Silicon, Docker cannot access the Metal GPU, so models run slowly on CPU; a native install via `start-macos.sh` is needed.
- Stable `main` branch recommended over `dev`; interface opens at `http://127.0.0.1:7860`; use an SSH tunnel for remote access.
- Qwen3-30B-A3B (MoE) was chosen; Cookbook recognized an M4 Max with 128 GB but llama.cpp crashed on `--flash-attn auto`.
- Ollama was the reliable path (`ollama pull qwen3:30b-a3b`, host 0.0.0.0:11434, endpoint `http://localhost:11434/v1`).
- Chat, Markdown editor and documents work well; Cookbook, agents, calendar and factual search remain fragile.
- Agents failed on tool calling (1,800 tokens spent without launching search); calendar hit `Could not extract JSON`; Deep Research hallucinated a video, a brand number and statistics.
- License AGPL; project still very young with strong potential but not production-ready for editorial work.

## Technical data / figures

| Item | Value |
|---|---|
| Project | Odysseus (by PewDiePie) |
| License | AGPL |
| Repository | github.com/odysseus-dev/odysseus |
| Mac install script | ./start-macos.sh |
| Default interface | http://127.0.0.1:7860 |
| Test machine | Apple M4 Max, 128 GB RAM |
| Model tested | Qwen3-30B-A3B (MoE) |
| Inference backend | Ollama (reliable); llama.cpp (failed); MLX (v1.0.2) |
| Ollama endpoint | http://localhost:11434/v1 |
| llama.cpp failure | Invalid `--flash-attn auto` |
| Agent tool-call waste | ~1,800 tokens without search |
| Deep Research output | 13 pages from 6 sources, with hallucinations |

## Why this source matters for the RAG

This source is a detailed, practical field report on self-hosting an AI workspace on macOS, with concrete installation steps, hardware notes, backend choices and observed limitations. It is valuable for questions about local/self-hosted AI, Apple Silicon inference, Ollama, and agent reliability.
