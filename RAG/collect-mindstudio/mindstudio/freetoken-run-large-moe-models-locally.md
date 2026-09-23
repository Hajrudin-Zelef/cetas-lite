---
id: collect-mindstudio/mindstudio/freetoken-run-large-moe-models-locally
title: "FreeToken Explained: Run 290B+ MoE Models on One Gaming GPU"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "DeepSeek", "MiniMax", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["gpu", "moe", "agent", "agents", "attention", "benchmark", "claude", "compute", "consumer", "cost", "deepseek", "glm"]
source: docs/RAG/Collect RAG/02_mindstudio/freetoken-run-large-moe-models-locally.md
source_anchor: ""
source_lines: [1, 55]
sha256: ff7e82234aaba81afbd89eedf53c372c87f8b5fe94d0a481df6f50b991f6c07a
---

# FreeToken Explained: Run 290B+ MoE Models on One Gaming GPU

## Metadata

- **Source** : https://www.mindstudio.ai/blog/freetoken-run-large-moe-models-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains **FreeToken**, a local inference tool built to run **mixture-of-experts (MoE)** models far too large to fit in a single GPU's VRAM. Instead of loading an entire model onto the GPU, it keeps most weights in **system RAM** and only pulls the specific "experts" a token needs onto the GPU when required. That lets a model with hundreds of billions of parameters — like **GLM, DeepSeek, or MiniMax** — run on a single consumer or workstation GPU without a multi-GPU cluster or an API subscription.

The problem FreeToken solves is specific to MoE architecture. In a dense model, every parameter is used for every token, so the whole model must sit in fast memory. MoE models work differently: for each token, a **router** selects only a handful of "expert" sub-networks out of the hundreds available, so most parameters go unused per forward pass. Traditional serving setups still load the entire model into VRAM just in case any expert is picked, which is why serving a 300+ billion parameter MoE conventionally demands eight or more GPUs wired together.

FreeToken flips that assumption. It keeps a small working set of frequently used experts cached on the GPU, alongside the **KV cache**, while the rest of the experts stay in ordinary system RAM. When the router calls for an uncached expert, FreeToken either streams that one expert over the **PCIe** bus and runs it (**offload mode**), or, in **hybrid mode**, computes it directly on the **CPU**. Either way, you only pay the VRAM cost for experts actively in flight, not the entire model.

Before serving, FreeToken runs a **benchmark pass** on the actual machine. It runs real kernels on the GPU for each quantization format present in the model, measuring GPU compute throughput, then compares that against the **CPU's compute speed** and the **PCIe transfer speed**. If the CPU can compute an expert's output faster than the data could be streamed over PCIe, it chooses **hybrid mode**; if moving weights over PCIe is faster than local CPU computation, it picks **offload mode**. In one demonstrated run using an **NVFP4-quantized** model on an **A6000** GPU, the benchmark selected offload mode automatically. This decision happens once per model/hardware combination with no manual configuration.

Once the benchmark completes, launching the server is a single command. FreeToken reads the checkpoint and GPU specs to automatically pick the **attention backend**, **MoE backend**, and **cache sizes**, including how much VRAM to reserve for the KV cache. In the demonstrated setup, weights loaded in about **30 seconds**, with just over **21 GB** set aside for KV cache alone. After loading, the server compiles **CUDA graphs** (a one-time warm-up that speeds up subsequent generation). Total VRAM consumption landed just over **44 GB**. None of these steps required manual tuning — no tensor parallelism, no manual quantization kernel selection, no hand-configured memory splits.

On value: FreeToken removes the two usual costs of running frontier-scale MoE models — a multi-GPU server rig and a recurring API bill. The tradeoff is speed and complexity versus a dedicated multi-GPU deployment: streaming experts over PCIe or computing them on CPU is inherently slower than having every expert resident in VRAM. The benchmark-driven mode selection minimizes that penalty, but it's still bound by a single card's VRAM and a single PCIe link. For local experimentation, coding-agent workflows, or privacy-sensitive use, that tradeoff is usually worth it; for high-throughput production serving to many concurrent users, dedicated multi-GPU infrastructure still wins. FreeToken includes **FT shell**, a terminal chat interface displaying live metrics (tokens per second, expert cache status including LRU/reuse, KV cache usage, VRAM consumption), and supports pointing agents like **Codex, DSH, Hermes agent, OpenClaw, OpenCode, and Claude** at its local server.

## Key points

- FreeToken runs 290B+ MoE models (GLM, DeepSeek, MiniMax) on one GPU by offloading experts to system RAM.
- Core trick: only the small subset of experts a token activates is loaded onto the GPU.
- Automatic benchmark compares CPU compute vs PCIe transfer to pick hybrid or offload mode.
- Demo: NVFP4 Qwen 3.6 on an A6000 loaded in ~30 s, >21 GB KV cache, ~44 GB steady-state VRAM.
- FT shell shows live tokens/s, expert cache hit rate, KV usage, and VRAM usage.
- Integrates with six coding agents by pointing them at the local server instead of a cloud API.
- No manual tuning: attention backend, MoE backend, and cache sizes are auto-selected.

## Technical data / figures

| Item | Value |
|---|---|
| Target models | GLM, DeepSeek, MiniMax, Qwen 3.6 (MoE) |
| Offload target | System RAM |
| Modes | Hybrid (CPU compute) / Offload (PCIe stream) |
| Demo GPU | A6000 |
| Demo quant | NVFP4 |
| Weight load time | ~30 s |
| KV cache reserved | >21 GB |
| Steady-state VRAM | >44 GB |
| Selected mode (demo) | Offload |
| Built-in interface | FT shell |
| Supported agents | Codex, DSH, Hermes agent, OpenClaw, OpenCode, Claude |
| Manual config needed | Minimal (auto backends/cache) |

## Why this source matters for the RAG

It provides the conceptual and architectural explanation of expert-offloading inference, a key technique for running frontier-scale MoE models on consumer hardware. It is essential grounding for questions about FreeToken, MoE serving, and single-GPU local AI.
