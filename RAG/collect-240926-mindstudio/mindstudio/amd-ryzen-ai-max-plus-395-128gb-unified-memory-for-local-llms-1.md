---
id: collect-240926-mindstudio/mindstudio/amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms-1
title: "amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Alibaba", "Unsloth"]
dates: []
keywords: ["amd", "memory", "agent", "agents", "benchmarks", "fine-tuning", "gpu", "gpus", "hbm", "inference", "latency", "llama"]
source: docs/RAG/clean_en/mindstudio/amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms.md
source_anchor: ""
source_lines: [1, 60]
sha256: 8fc18c142ca97f8ce3311575f6c99a42444bb041e4aa82be0873a726e2c5a11d
---

# amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms

<!-- source: https://www.mindstudio.ai/blog/amd-ryzen-ai-max-plus-395-local-ai-workstation -->

## What is AMD’s Ryzen AI Max Plus 395?

The Ryzen AI Max Plus 395 is AMD’s chip for the “Ryzen AI Halo” line of workstations, built around a single idea: instead of pairing a CPU with a discrete GPU that has its own fixed VRAM, it gives the CPU and GPU a shared pool of 128GB of LPDDR5X memory. That unified memory pool, combined with a Radeon 8060S GPU rated around 60 teraflops at FP16 and 16 CPU cores, lets a single machine load language models well beyond what a 24GB or 32GB graphics card can hold, including 100 billion-plus parameter models like GPT-OSS 120B and larger Qwen 3 variants.

## TL;DR

- The **Ryzen AI Max Plus 395** uses one shared 128GB LPDDR5X pool instead of separate VRAM and system RAM, and users can manually allocate how much of that pool goes to the GPU versus the CPU.
- Because there’s no hard VRAM ceiling, the machine can load **100B+ parameter models** such as GPT-OSS 120B and larger Qwen 3.5/3.6 releases that simply won’t fit on a 32GB discrete card like AMD’s Radeon AI Pro R9700.
- In testing, a 9B dense model with a multi-token-predictor drafter ran at just under **40 tokens per second** , a 35B mixture-of-experts Qwen 3.6 model (3B active parameters) topped**50 tokens per second** , and GPT-OSS 120B ran around**30 tokens per second** .
- The Linux version arrives with **ROCm, PyTorch, LM Studio, Llama.cpp, and Comfy UI already configured** , avoiding the manual driver and partition setup that discrete-GPU Linux AI boxes typically require.
- AMD’s **AI Developer Center** bundles setup guides, BIOS configuration help, and step-by-step “playbooks” for tasks ranging from beginner Comfy UI runs to fine-tuning with Unsloth and writing custom GPU kernels.
- A **50 TOPS NPU** ships alongside the CPU and GPU, but most current LLM inference stacks don’t use it yet, meaning part of the chip’s capability is still waiting on software support.
- AMD has already confirmed a **Pro 395 variant with up to 192GB** of memory, aimed at users who want to push past 128GB for even larger models or multiple models running side by side.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## How does unified memory change what you can run locally?

On a typical discrete-GPU workstation, the GPU has its own dedicated VRAM (say, 32GB), and system RAM sits separately, usually connected over a much slower path. Once a model’s weights exceed VRAM capacity, the extra layers get offloaded to system RAM, and token generation speed drops sharply. Quantization can shrink a model to fit, but there’s a hard ceiling: no amount of clever compression turns a 120B-parameter model into something that fits in 32GB of VRAM at usable quality.

The Ryzen AI Halo removes that wall by giving the CPU and GPU access to the same 128GB pool. Instead of the GPU being capped by a fixed VRAM chip, the user decides how much of the shared memory to hand to the GPU (in testing, the default split allocated 75% to the GPU) and how much stays reserved for the CPU and general system tasks. Since most AI applications don’t need heavy system RAM, the vast majority of the 128GB can go toward loading model weights on the GPU side.

The tradeoff is that this shared memory isn’t as fast as dedicated GDDR6 or HBM on a discrete card, so raw token throughput on smaller models that do fit in 32GB will typically be slower on the Halo than on a card purpose-built for that size. The benefit shows up specifically at the sizes that discrete GPUs can’t reach at all.

## What kind of performance can you actually expect?

Benchmarks depend heavily on model architecture, so it’s worth separating dense models from mixture-of-experts (MoE) models:

- A 9B dense model (Hermes/Nemotron-class, running with a multi-token-predictor drafter) produced just under 40 tokens per second, representative of what most 9B-class dense models achieve on this hardware.
- A 35B Qwen 3.6 MoE model with only 3B active parameters per token, quantized to an “Exl” quant, exceeded 50 tokens per second, faster than the dense model despite having a larger total parameter count, because MoE architectures only activate a fraction of their weights per forward pass.
- GPT-OSS 120B, a mixture-of-experts model that ships pre-installed on the machine, ran around 30 tokens per second, which is workable for chat, retrieval-augmented generation, and general use, though a smaller and faster model is generally the better choice for latency-sensitive agent workflows.

For context, none of these larger models are options at all on a 32GB discrete GPU workstation without heavy offloading and the performance collapse that comes with it. The Halo’s advantage isn’t peak speed on small models, it’s the ability to load models that otherwise wouldn’t run locally at all.

## What software comes preinstalled, and does it actually save setup time?

The Ryzen AI Halo ships in two versions: a Windows build and a Linux-based build. On the Linux version, there’s no manual drive partitioning and no separate ROCm installation required, which is typically one of the more tedious parts of setting up an AMD box for local AI work.

The machine includes an “AI Developer Center” application that surfaces installed software versions (PyTorch, ROCm), lets users adjust how much memory the GPU gets versus the CPU, and supports remote access over SSH so the box can run headless on a desk while another computer connects to it for inference jobs.

The Developer Center also bundles a library of “playbooks,” step-by-step guides split by skill level and by OS:

- Beginner playbooks cover running Comfy UI with pre-installed models and serving LLMs through LM Studio.
- Intermediate playbooks include fine-tuning models with Unsloth, alongside guidance on configuring GPU memory allocation for different workloads.
- Advanced playbooks cover clustering multiple AI Halo units together and writing custom GPU kernels, aimed at users doing lower-level performance work.

This preinstalled stack matters most for people who’ve previously lost hours to ROCm version mismatches or driver conflicts on Linux AI rigs. It won’t cover every workflow out of the box, but it removes a meaningful chunk of the setup friction that discrete-GPU Linux builds usually involve.

## Is the NPU actually useful yet?

The chip includes an NPU rated around 50 TOPS, sitting alongside the CPU and GPU. Right now, most LLM inference stacks and popular local AI applications don’t route work through the NPU, so its capability is largely unused in day-to-day LLM serving. This is a software gap rather than a hardware limitation: the silicon is there, but the ecosystem (drivers, runtime support, framework integration) hasn’t caught up to make it a default part of the inference pipeline. Whether that changes depends on how quickly AMD and framework maintainers add NPU support to tools like Llama.cpp and ONNX Runtime.

## Is the Ryzen AI Max Plus 395 worth it for local AI work?

It depends on what “local AI” means for the buyer. For anyone whose ceiling has been VRAM capacity, someone who wants to run GPT-OSS 120B, larger Qwen 3.5 or 3.6 models, or multiple mid-size models simultaneously, the 128GB unified memory pool solves a real problem that no amount of quantization fixes on a 32GB card. The tradeoff is that unified memory won’t match dedicated VRAM bandwidth on models small enough to fit on both, so buyers chasing maximum tokens-per-second on smaller models may still prefer a discrete GPU setup.

