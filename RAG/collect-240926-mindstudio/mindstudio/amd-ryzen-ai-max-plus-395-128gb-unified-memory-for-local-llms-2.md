---
id: collect-240926-mindstudio/mindstudio/amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms-2
title: "amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms"
domain: mindstudio
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "memory", "agent", "gpu", "latency", "llama", "llama.cpp", "lpddr5x", "throughput", "tokens per second", "video generation"]
source: docs/RAG/clean_en/mindstudio/amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms.md
source_anchor: ""
source_lines: [61, 83]
sha256: 4f84e242549e5f3336cef8b2ffc2772854069db4d17617c987f44c9dd8026f21
---

# amd-ryzen-ai-max-plus-395-128gb-unified-memory-for-local-llms

The bigger picture: AMD has already announced a Pro 395 variant supporting up to 192GB, suggesting this unified-memory approach is a platform strategy, not a one-off product. For image and video generation work, the same logic applies. Comfy UI ran pre-installed with an image model available out of the box, and downloading additional models (including video generation models) is straightforward since the memory ceiling isn’t the constraint it would be on a smaller card.

## Frequently Asked Questions

### What is the Ryzen AI Max Plus 395?

It’s AMD’s chip powering the Ryzen AI Halo workstations, combining 16 CPU cores, a Radeon 8060S GPU (around 60 teraflops FP16), and a roughly 50 TOPS NPU, all sharing a single 128GB LPDDR5X unified memory pool instead of separate VRAM and system RAM.

### How much faster is unified memory than offloading to system RAM on a discrete GPU?

Unified memory doesn’t necessarily produce faster raw throughput than dedicated VRAM on models that fit within that VRAM. The real advantage is avoiding the steep speed drop that happens when a discrete GPU setup has to offload model layers to slower system RAM once a model exceeds available VRAM, a scenario the unified pool sidesteps by keeping everything in one addressable space.

### Can you run GPT-OSS 120B locally on this machine?

Yes. GPT-OSS 120B comes pre-installed on the tested unit and ran at around 30 tokens per second, a usable speed for chat and retrieval-augmented generation tasks, though smaller models are better suited to latency-sensitive agent workloads.

### Does the Ryzen AI Halo require manual ROCm and driver setup on Linux?

No. The Linux-based version ships with ROCm, PyTorch, LM Studio, Llama.cpp, and Comfy UI already installed and configured, along with an AI Developer Center app for managing memory allocation and software versions.

### Is there a version with more than 128GB of memory?

AMD has announced a Pro 395 variant supporting up to 192GB of unified memory, intended for users who need to load even larger models or run multiple models concurrently.
