---
id: collect-mindstudio/mindstudio/freetoken-deepseek-v4-flash-single-3090
title: "DeepSeek V4 Flash on One RTX 3090: Real Tokens-Per-Second Numbers"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agentic", "benchmark", "consumer", "cost", "glm", "gpu", "inference", "latency", "llama", "llama.cpp", "memory"]
source: docs/RAG/Collect RAG/02_mindstudio/freetoken-deepseek-v4-flash-single-3090.md
source_anchor: ""
source_lines: [1, 52]
sha256: 60054495e6d856759cd1a90fe52b6c2278d11568db58d69855f0550e84364697
---

# DeepSeek V4 Flash on One RTX 3090: Real Tokens-Per-Second Numbers

## Metadata

- **Source** : https://www.mindstudio.ai/blog/freetoken-deepseek-v4-flash-single-3090
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article presents real benchmark results for running **DeepSeek V4 Flash** and **Qwen 3.8 27B** locally on a single **RTX 3090** or **4090** using **FreeToken's desktop app**. The answer to "can a single RTX 3090 run DeepSeek V4 Flash?" is yes, but the GPU is only part of the story. DeepSeek V4 Flash is a **mixture-of-experts (MoE)** model, so most parameters sit idle for any given token and only a fraction ("experts") activate per pass. That property lets it run on a single 3090/4090 with the bulk of the model offloaded to **system RAM** instead of VRAM. In real testing, a single 3090 paired with a large pool of **DDR4 RAM** produced **8 to 11 tokens per second**, depending on which experts were active and whether the model was accessed through a server API or the desktop client.

Measured throughput: about **10 to 11 tokens/second** through an **Open WebUI** front end talking to a llama.cpp-style backend, and closer to **8.8 tokens/second** when run directly through FreeToken's desktop client, with the app briefly reporting an inconsistent **1.8 tokens/second** on an earlier run before stabilizing. Neither number is fast by API standards — roughly interactive chat speed, workable for conversation but not for high-throughput or agentic workflows firing dozens of generations back to back. The presenter was explicit this setup is not for latency-sensitive work; it's for people who want frontier-adjacent model quality on hardware they already own, trading speed for it. Throughput also isn't perfectly stable run to run: because MoE models route tokens to different experts depending on content, the mix of experts loaded and swapped affects speed, which is why numbers "vary wildly" even on identical hardware.

On RAM requirements — the detail most casual "run it on one GPU" claims skip — **128 GB was not enough** to run DeepSeek V4 Flash. Bumping to **192 GB** worked, and the presenter estimated **156 to 168 GB** as the realistic practical minimum for most users. The test rig used older **DDR4 at 2400 MT/s**, well behind current DDR4 (3200 MT/s) or DDR5. Since MoE inference is memory-bandwidth bound once active experts must be fetched from RAM, slower memory directly caps tokens/second; the presenter expects **DDR5 systems to roughly double throughput** versus the DDR4 2400 setup. Practical guidance: **64 GB** is a reasonable floor for smaller MoE models, **96-128 GB** opens more headroom, and workloads like DeepSeek V4 Flash or larger MoE releases such as **GLM 5.2** push requirements into **192 GB and up**. FreeToken proactively reports shortfalls — in one test it needed over **200 GB of additional system memory** to load a large GLM 5.2 variant.

A dense model tells a different story. **Qwen 3.8 27B in BF16** (non-MoE, fully dense) **failed to load reliably** in the FreeToken desktop beta, throwing a generic engine error. Dense models don't benefit from the same RAM-offload trick because there's no way to selectively load only the "active" part of the network. On the app itself: FreeToken is explicitly beta software with rough edges — inconsistent tokens-per-second reporting between its desktop UI and server view, and not every model in its library loads cleanly. Still, it lowers the barrier by shipping Windows, Ubuntu, AppImage, and Arch Linux packages, a Hugging Face-backed model download flow, and a built-in cost comparison against API pricing. For someone with a single 3090/4090 and a well-stocked RAM pool, it's one of the more approachable ways to run a large MoE model without becoming a full-time sysadmin.

## Key points

- DeepSeek V4 Flash ran at ~10-11 tok/s via Open WebUI and ~8.8 tok/s via the FreeToken desktop client on a single RTX 3090.
- System RAM capacity and bandwidth matter as much as the GPU for MoE offload.
- 128 GB RAM was insufficient; 192 GB worked, with ~156-168 GB the likely practical floor.
- DDR4 2400 MT/s capped throughput; DDR5 is expected to roughly double it.
- Dense Qwen 3.8 27B BF16 failed to load in the FreeToken beta (generic engine error).
- FreeToken is beta with inconsistent tok/s reporting and limited model coverage.
- GLM 5.2 variants can need 200+ GB additional RAM to load.

## Technical data / figures

| Item | Value |
|---|---|
| GPU | Single RTX 3090 (or 4090) |
| Model | DeepSeek V4 Flash (MoE) |
| Open WebUI throughput | ~10-11 tok/s |
| FreeToken desktop client | ~8.8 tok/s (briefly 1.8) |
| RAM tested | 128 GB insufficient; 192 GB works |
| Practical RAM floor | ~156-168 GB |
| Test memory speed | DDR4 2400 MT/s |
| DDR5 expectation | ~2× throughput |
| Small MoE RAM floor | 64 GB |
| Comfortable range | 96-128 GB |
| Large MoE (GLM 5.2) | 200+ GB additional RAM |
| Dense test | Qwen 3.8 27B BF16 failed to load |
| Platforms | Windows, Ubuntu, AppImage, Arch |

## Why this source matters for the RAG

It supplies concrete, real-world tokens-per-second and RAM figures for running a large MoE model on consumer hardware, correcting the common understatement of system RAM needs. It is valuable for grounding claims about MoE offloading and single-GPU local inference.
