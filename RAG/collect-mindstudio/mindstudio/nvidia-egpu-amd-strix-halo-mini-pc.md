---
id: collect-mindstudio/mindstudio/nvidia-egpu-amd-strix-halo-mini-pc
title: "Nvidia eGPU on AMD Strix Halo: Oculink Setup and Real Speed Gains"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["amd", "gpu", "nvidia", "compute", "gpus", "inference", "memory", "moe", "parameters", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/nvidia-egpu-amd-strix-halo-mini-pc.md
source_anchor: ""
source_lines: [1, 50]
sha256: 55199148556eb2c7001207044a1c89237ec4e18267b8dff1a978a8d8f6d74d05
---

# Nvidia eGPU on AMD Strix Halo: Oculink Setup and Real Speed Gains

## Metadata

- **Source** : https://www.mindstudio.ai/blog/nvidia-egpu-amd-strix-halo-mini-pc
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how an **Oculink port** lets an external Nvidia RTX GPU pair with AMD's **Strix Halo** APU in a mini PC, and what that actually does for local LLM inference speed. Strix Halo is AMD's APU platform that shares up to **128 GB of memory** between CPU and GPU, already capable of running large language models on its own. Adding Oculink turns the mini PC into something closer to a small desktop able to accept a real external GPU (eGPU).

Oculink is the enabler: a PCIe-based external port running at roughly **63 GB/s**, the same class of connection a desktop GPU normally uses, exposed on the outside of a mini PC. Strix Halo's integrated GPU (the **Radeon 8060** in the tested unit) is capable for a mobile-class chip and shares a large unified memory pool, which is why Strix Halo mini PCs became popular for local LLM work — big memory at a lower price than a workstation GPU. But shared-memory bandwidth and integrated-graphics compute are still slower than a discrete GPU, so an eGPU becomes attractive once more raw speed is needed.

The gains depend heavily on whether the model fits in the eGPU's VRAM. With an **RTX 5080 (16 GB)** over Oculink, a model that fit ran about **three times faster** than the built-in GPU. But once a model exceeded 16 GB, output dropped to roughly **1.6 tokens/second**, far slower than the built-in GPU's **11+ tokens/second** on the same model. With an **RTX Pro 6000 (96 GB VRAM)**, the equation flips: a **32B dense model ran at 69 tokens/second** (about six times the integrated GPU), and a **122B MoE model** (Q4, ~76.5 GB on disk) hit around **121 tokens/second** — faster than the smaller dense model because MoE activates only a fraction of total parameters per token.

Cross-vendor splitting is possible despite incompatible software stacks. Nvidia uses **CUDA**, AMD uses **ROCm (Radeon Open Compute)**, and they don't interoperate by default. **Vulkan** — a cross-platform graphics and compute API both vendors support — is the workaround. Tools like **LM Studio** let you pick Vulkan as the inference runtime instead of CUDA or ROCm, enabling a model to be split by layer across both GPUs. In one test, a **122B model at Q8 (~130 GB on disk)**, too large for either GPU alone, was split with about **70% of layers on the Nvidia GPU** and the rest on the AMD iGPU, producing a memory split of roughly **88 GB Nvidia / 34 GB AMD** at about **37-38 tokens/second**. It is slower than either GPU running a fully resident model, partly because the Oculink connection becomes a bottleneck for inter-GPU data movement.

Practical snags: the physically large RTX Pro 6000 initially caused the mini PC to fail to boot (black screen), requiring booting without the card, then powering it on afterward — essentially hot-plugging, which is explicitly against normal eGPU/Oculink guidance but necessary. Memory allocation to the integrated GPU also had to be set manually in BIOS (e.g. 96 GB); left on automatic, the system didn't properly connect with the external GPU. The Oculink-equipped model carries a modest price premium and gives up some ports and stand stability versus the prior generation.

## Key points

- Oculink is a PCIe-based external port (~63 GB/s), the same class as a desktop GPU's internal slot, rare on mini PCs.
- Small eGPUs hit a wall fast: RTX 5080 (16 GB) is ~3× faster on fitting models but collapses to <2 tok/s once the model overflows.
- Big cards change the math: RTX Pro 6000 (96 GB) runs a 32B dense model at 69 tok/s and a 122B MoE at 120+ tok/s.
- MoE models punch above their size: 122B with ~10B active runs faster than a smaller dense 32B.
- Cross-vendor splitting works via Vulkan: a 122B Q8 model split ~88 GB Nvidia / 34 GB AMD ran at ~37-38 tok/s.
- Practical friction: boot failures with large cards require hot-plugging, and BIOS memory allocation must be manual.

## Technical data / figures

| Item | Value |
|---|---|
| Oculink bandwidth | ~63 GB/s |
| Tested integrated GPU | Radeon 8060 |
| Shared memory | Up to 128 GB |
| RTX 5080 (16 GB) fitting model | ~3× built-in GPU |
| RTX 5080 overflow | ~1.6 tok/s (vs 11+ tok/s built-in) |
| RTX Pro 6000 (96 GB), 32B dense | 69 tok/s |
| RTX Pro 6000, 122B MoE (Q4, ~76.5 GB) | ~121 tok/s |
| Dual-GPU split, 122B Q8 (~130 GB) | ~37-38 tok/s (88 GB Nvidia / 34 GB AMD) |
| Cross-vendor API | Vulkan (vs CUDA / ROCm) |
| BIOS manual allocation example | 96 GB to iGPU |

## Why this source matters for the RAG

It documents a rare dual-vendor (Nvidia + AMD) local inference setup with real throughput figures and the practical pitfalls of Oculink eGPUs. This is a distinctive reference for hardware-focused questions about VRAM-fit tradeoffs and cross-vendor model splitting.
