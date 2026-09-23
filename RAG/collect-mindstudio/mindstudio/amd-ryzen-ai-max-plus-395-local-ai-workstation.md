---
id: collect-mindstudio/mindstudio/amd-ryzen-ai-max-plus-395-local-ai-workstation
title: "AMD Ryzen AI Max Plus 395: 128GB Unified Memory for Local LLMs"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Alibaba", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["amd", "memory", "agents", "fine-tuning", "gpu", "gpus", "hbm", "inference", "latency", "llama", "llama.cpp", "lpddr5x"]
source: docs/RAG/Collect RAG/02_mindstudio/amd-ryzen-ai-max-plus-395-local-ai-workstation.md
source_anchor: ""
source_lines: [1, 56]
sha256: e6a1aec1fb672b38912649bca6863860114e730a2b5b1457598a73ff7b05e163
---

# AMD Ryzen AI Max Plus 395: 128GB Unified Memory for Local LLMs

## Metadata

- **Source** : https://www.mindstudio.ai/blog/amd-ryzen-ai-max-plus-395-local-ai-workstation
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The **Ryzen AI Max Plus 395** is AMD's chip for the "**Ryzen AI Halo**" line of workstations, built around a single idea: instead of pairing a CPU with a discrete GPU with fixed VRAM, it gives the CPU and GPU a shared pool of **128GB of LPDDR5X memory**. That unified memory pool, combined with a **Radeon 8060S GPU** (~60 teraflops at FP16), **16 CPU cores**, and a ~50 TOPS NPU, lets a single machine load language models far beyond what a 24GB or 32GB graphics card can hold — including **100B+ parameter models** like **GPT-OSS 120B** and larger Qwen 3 variants.

**How unified memory changes what you can run:** on a discrete-GPU workstation, once weights exceed VRAM, extra layers offload to system RAM and token speed drops sharply; no amount of quantization fits a 120B model in 32GB at usable quality. The Halo removes that wall — the user decides how much of the shared 128GB to hand to the GPU (the default split allocated **75% to the GPU**). The tradeoff is that shared memory isn't as fast as dedicated GDDR6 or HBM, so small models that fit in 32GB run slower than on a discrete card; the benefit appears specifically at sizes discrete GPUs can't reach.

**Measured performance:**
- A **9B dense model** (Hermes/Nemotron-class) with a multi-token-predictor drafter: just under **40 tok/s**.
- A **35B Qwen 3.6 MoE** model (3B active per token, "Exl" quant): **over 50 tok/s** — faster than the dense model despite larger total size, because MoE activates only a fraction of weights per pass.
- **GPT-OSS 120B** (pre-installed MoE): ~**30 tok/s**, workable for chat, RAG, and general use, though smaller/faster models suit latency-sensitive agents.

**Software:** ships in Windows and Linux builds. The Linux version comes with **ROCm, PyTorch, LM Studio, Llama.cpp, and ComfyUI** pre-configured, avoiding manual ROCm/driver setup. An "**AI Developer Center**" app surfaces installed versions, lets users adjust GPU/CPU memory split, supports SSH remote/headless access, and bundles "playbooks": beginner (ComfyUI, LM Studio serving), intermediate (Unsloth fine-tuning, memory allocation), and advanced (clustering multiple Halo units, custom GPU kernels).

**NPU:** the ~50 TOPS NPU sits alongside CPU and GPU, but most current LLM inference stacks don't route work through it, so its capability is largely unused in day-to-day serving — a software gap rather than a hardware limitation, pending AMD/framework support in tools like Llama.cpp and ONNX Runtime.

**Verdict:** for anyone whose ceiling is VRAM capacity — running GPT-OSS 120B, larger Qwen 3.5/3.6, or multiple mid-size models simultaneously — the 128GB pool solves a real problem. AMD has announced a **Pro 395 variant supporting up to 192GB**, indicating a platform strategy. The same logic extends to image/video generation: ComfyUI ran pre-installed, and downloading additional models is straightforward since memory isn't the constraint.

## Key points

- Ryzen AI Max Plus 395: 16 CPU cores + Radeon 8060S (~60 TFLOPS FP16) + ~50 TOPS NPU sharing 128GB LPDDR5X unified memory.
- Removes the VRAM ceiling: loads 100B+ models (GPT-OSS 120B, larger Qwen 3.x) that a 32GB card cannot.
- Measured speeds: 9B dense ~40 tok/s; 35B Qwen3.6 MoE (3B active) >50 tok/s; GPT-OSS 120B ~30 tok/s.
- Linux build ships with ROCm, PyTorch, LM Studio, Llama.cpp, ComfyUI pre-configured.
- AI Developer Center: memory-split control, SSH/headless, beginner→advanced playbooks.
- 50 TOPS NPU is currently mostly idle due to missing software support.
- Pro 395 variant announced with up to 192GB.

## Technical data / figures

| Spec | Value |
|---|---|
| Chip | AMD Ryzen AI Max Plus 395 |
| CPU | 16 cores |
| GPU | Radeon 8060S (~60 TFLOPS FP16) |
| NPU | ~50 TOPS |
| Memory | 128GB unified LPDDR5X (GPU split configurable; tested default 75% to GPU) |
| 9B dense + MTP drafter | ~40 tok/s |
| Qwen3.6 35B MoE (3B active) | >50 tok/s |
| GPT-OSS 120B (MoE) | ~30 tok/s |
| Preinstalled (Linux) | ROCm, PyTorch, LM Studio, Llama.cpp, ComfyUI |
| Variant | Pro 395 up to 192GB |

## Why this source matters for the RAG

It documents a specific unified-memory workstation architecture that makes 100B+ parameter models locally runnable, with concrete throughput numbers and the software stack that removes Linux setup friction. It is a key hardware reference for comparing local vs cloud AI options and for advising on workstation-class local inference purchases.
