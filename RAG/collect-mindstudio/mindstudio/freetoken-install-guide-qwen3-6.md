---
id: collect-mindstudio/mindstudio/freetoken-install-guide-qwen3-6
title: "How to Install FreeToken and Serve Qwen 3.6 Locally"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "MiniMax", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["qwen", "agent", "agents", "attention", "benchmark", "claude", "compute", "consumer", "cost", "deepseek", "glm", "gpu"]
source: docs/RAG/Collect RAG/02_mindstudio/freetoken-install-guide-qwen3-6.md
source_anchor: ""
source_lines: [1, 56]
sha256: c1e25f2cd5b5091a6889ff4e122e6da21e94551480af13d7fc02cc910bdb6a41
---

# How to Install FreeToken and Serve Qwen 3.6 Locally

## Metadata

- **Source** : https://www.mindstudio.ai/blog/freetoken-install-guide-qwen3-6
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is a hands-on guide to installing **FreeToken**, benchmarking your GPU/CPU split, serving **Qwen 3.6**, and connecting coding agents to it locally. FreeToken is a serving tool built to run frontier **mixture-of-experts (MoE)** models — the kind with hundreds of billions of parameters — on consumer hardware that lacks enough VRAM to hold them. Models like **GLM, DeepSeek, and MiniMax** are MoE architectures: even though the full model is enormous, each token only activates a small subset of "experts." FreeToken exploits this by keeping most of the model parked in **system RAM** and only pulling the specific experts a token needs onto the GPU.

Expert offloading works by keeping a small working set of experts cached on the GPU next to the **KV cache**, while the rest sit in regular system RAM. When the router selects an expert not already cached, FreeToken either (1) streams just that expert over **PCIe** to the GPU, uses it, and moves on (**offload mode**), or (2) in **hybrid mode**, computes that expert directly on the **CPU** instead of transferring it. Either way you only pay the VRAM cost for the handful of experts actively in flight, not the entire model — the difference between needing an eight-GPU server and running on a single card.

Installation (demonstrated on **Lubuntu** with a single GPU) follows a straightforward path: create a virtual environment (the demo used **UV**) and install FreeToken; install the **Hugging Face CLI** (also via UV) to handle downloads; download the target model (Qwen 3.6). Installation and download each took a few minutes with no unusual dependencies.

Before serving, FreeToken runs a **benchmark** against your specific machine, measuring the actual hardware bottleneck rather than assuming one. It runs real kernels on the GPU for each quantization format and compares the **CPU's compute speed** against **PCIe transfer speed** for that format. If CPU compute beats PCIe transfer by a sufficient margin, it chooses **hybrid mode**; if PCIe transfer wins, it chooses **offload mode**. In the demo, running an **NVFP4** quantized model on an **A6000** GPU, the benchmark landed on offload mode, applied automatically at server launch.

Serving is a single `FT serve` command pointing at the downloaded model. FreeToken automatically configures the **attention backend**, **MoE backend**, and **cache sizes** from the checkpoint and hardware specs. In the demonstrated run: model weights loaded in about **30 seconds**; just over **21 GB of VRAM** reserved specifically for the KV cache; CUDA graphs compiled as a one-time warm-up; steady-state VRAM usage just over **44 GB** on the A6000. A simple `curl` request confirms which model is being served.

**FT shell** is FreeToken's built-in terminal chat interface with a live status bar surfacing diagnostics: **tokens per second**, **cache and LRU stats** (how full the GPU expert cache is and how effectively it reuses recently accessed experts), **KV usage** (of the 21 GB reserved), **Mamba state** (linear attention state cache for hybrid architectures), and **VRAM consumption** (42.9K out of 48 GB in the demo). This gives a direct read on efficiency without separate monitoring tools.

FreeToken supports pointing external coding agents at its local server instead of a hosted API. It is compatible with **Codex, Hermes agent, OpenClaw, OpenCode, and Claude**. The pattern is consistent: launch the agent with a command redirecting it to your local FreeToken server address; if the agent isn't installed, the launch command installs it automatically. In the demo, both Hermes agent and Claude Code were launched this way, each picking up the locally served Qwen 3.6. The tradeoff is that you're still bound by system RAM capacity and by PCIe or CPU compute speed for uncached experts; larger models like GLM or DeepSeek push those constraints harder.

## Key points

- FreeToken runs large MoE models on a single GPU by keeping most weights in system RAM and streaming only active experts.
- Two modes: offload (stream expert over PCIe) and hybrid (compute expert on CPU); a benchmark picks automatically.
- Install: UV virtualenv, FreeToken, Hugging Face CLI, download model (Qwen 3.6).
- Demo on an A6000: ~30 s load, >21 GB KV cache, ~44 GB steady-state VRAM, offload mode for NVFP4.
- FT shell shows live tokens/s, cache/LRU stats, KV usage, Mamba state, and VRAM consumption.
- Integrates with Codex, Hermes agent, OpenClaw, OpenCode, and Claude via a local server address.
- Constraints: system RAM capacity and PCIe/CPU speed for uncached experts.

## Technical data / figures

| Item | Value |
|---|---|
| Install tools | UV, Hugging Face CLI |
| Demo OS | Lubuntu |
| Demo GPU | A6000 |
| Demo model | Qwen 3.6 (NVFP4) |
| Selected mode | Offload |
| Weight load time | ~30 s |
| KV cache reserved | >21 GB |
| Steady-state VRAM | >44 GB |
| VRAM ceiling (demo) | 48 GB (42.9K used) |
| Supported agents | Codex, Hermes agent, OpenClaw, OpenCode, Claude |
| Modes | Hybrid (CPU compute) / Offload (PCIe stream) |

## Why this source matters for the RAG

It provides a step-by-step operational reference for installing and using FreeToken, including its benchmark logic and live diagnostics. This supports procedural questions about single-GPU MoE serving and connecting local models to coding agents.
