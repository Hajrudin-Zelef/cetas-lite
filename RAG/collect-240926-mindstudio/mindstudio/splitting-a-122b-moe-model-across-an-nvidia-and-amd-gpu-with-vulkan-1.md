---
id: collect-240926-mindstudio/mindstudio/splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan-1
title: "splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Hugging Face", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "moe", "nvidia", "compute", "consumer", "gpus", "inference", "inference engine", "llama", "llama.cpp", "memory"]
source: docs/RAG/clean_en/mindstudio/splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan.md
source_anchor: ""
source_lines: [1, 58]
sha256: 47361e2c83c97bd1f0ed6705f3efa84cc08d528a6881cc604a33d1f6bde9f8e0
---

# splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan

<!-- source: https://www.mindstudio.ai/blog/running-mixture-of-experts-models-across-two-gpus -->

## What does it mean to split a model across two different GPU vendors?

Splitting a large language model across an Nvidia GPU and an AMD GPU means dividing the model’s layers between the two devices so each one holds and processes part of the network, even though Nvidia and AMD use completely different low-level software stacks for GPU compute. Nvidia relies on CUDA, AMD relies on ROCm, and the two don’t talk to each other. The workaround is Vulkan, a cross-platform graphics and compute API that both vendors support, which lets a single inference engine like llama.cpp address both cards at once and load a model too big for either GPU alone.

## TL;DR

- **Vulkan is the bridge** that lets llama.cpp split one model across an Nvidia GPU and an AMD GPU, since CUDA and ROCm are incompatible software stacks that can’t share a single inference job.
- A **122 billion parameter mixture-of-experts model** (Q4 quantization, about 76.5GB on disk) ran at 121 tokens per second on a single RTX Pro 6000 with 96GB of VRAM, faster than a much smaller 32B dense model on the same card.
- Mixture-of-experts models only activate a fraction of their total parameters per token (this model uses roughly 10B active out of 122B, hence the “A10B” label on Hugging Face), which is why huge MoE models can outrun smaller dense ones.
- A **Q8 version of the same model at 130GB** didn’t fit on either GPU alone, so it was split with about 88GB loaded on the Nvidia RTX Pro 6000 and 34GB on the AMD Strix Halo APU, producing 37 to 38 tokens per second.
- The split isn’t 50/50, it’s **layer-based and proportional to available memory** , with roughly 70% of layers on the larger GPU and the rest on the smaller one, and every token has to pass through both devices in sequence.
- An **Oculink port** on the GMKtec EVO-X2 mini PC (a 63GB/s external PCIe connection) made it possible to attach a full desktop GPU to a Strix Halo APU system for the first time, though that link speed becomes the bottleneck for cross-GPU inference.
- Running eight concurrent requests against the split setup scaled throughput to around 80 tokens per second combined, while the Nvidia card drew only about 130 to 150 watts and the AMD APU about 76 watts, far below their rated power limits.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

## Why would anyone combine an Nvidia GPU with an AMD APU?

The practical reason is memory. Large open-weight models, especially mixture-of-experts (MoE) architectures, often exceed what a single consumer or even prosumer GPU can hold in VRAM. AMD’s Strix Halo APUs solve part of this by sharing a unified memory pool of up to 128GB between CPU and GPU, letting a mini PC load models that would normally require a multi-GPU server. But that integrated GPU is still slower than a discrete card for models that fit comfortably in less memory.

Adding an Oculink port changes the equation. It’s an external PCIe interface running at roughly 63GB/s, similar in concept to the slot a desktop graphics card plugs into, except exposed on the outside of a mini PC. That opens the door to attaching a full-size eGPU, whether a mid-range card or Nvidia’s flagship RTX Pro 6000 with 96GB of VRAM, to a machine that otherwise tops out at integrated graphics. Combine the two memory pools and you can fit models that neither device could hold on its own.

## How does the actual GPU split work under the hood?

Model weights in a transformer are organized into layers, and llama.cpp can distribute those layers across multiple devices rather than requiring the whole model to sit on one GPU. In the demonstrated setup, roughly 70% of the layers were placed on the Nvidia RTX Pro 6000 and the rest on the AMD Strix Halo integrated GPU. Every token generated has to pass through both sets of layers in order, which means the two GPUs work as stages in a pipeline rather than splitting the workload in parallel.

Vulkan is what makes this possible across vendors. Tools like LM Studio expose a runtime choice between CUDA, ROCm, and Vulkan, and while CUDA or ROCm might be faster on a single-vendor setup, only Vulkan lets one inference process span an Nvidia card and an AMD card simultaneously. The connection between the two also matters: because the eGPU is attached over Oculink rather than a full internal PCIe slot, that link becomes a throughput ceiling. A faster interconnect would likely push combined tokens-per-second higher.

One hardware detail worth noting: memory allocation for the AMD APU’s GPU needs to be set manually in BIOS on machines that expose that control. Left on automatic settings, the system may not recognize or properly allocate memory for an attached eGPU, so specifying the amount (96GB in the tested configuration) was necessary to get both devices working together.

## Why did the 122B model outperform a 32B model on the same GPU?

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

This comes down to the mixture-of-experts design. A dense 32 billion parameter model activates all 32 billion parameters for every token it generates. A mixture-of-experts model with 122 billion total parameters might only activate around 10 billion of them per token, routing each token to a subset of specialized “expert” sub-networks rather than running the whole model every time. Hugging Face model cards typically label this with a notation like “A10B” (active 10 billion) alongside the total parameter count.

The practical effect: a MoE model can be far larger on disk and in memory while requiring less actual computation per token than a smaller dense model. In the tests, the 122B MoE model at Q4 quantization ran at 121 tokens per second on the RTX Pro 6000, noticeably faster than the 32B dense model on the same card. The same principle applied to the integrated GPU inside the Strix Halo chip, where a 35B-parameter MoE model ran close to six times faster than an old-style dense model of similar size.

## Is splitting a model across mismatched GPUs actually worth it?

It depends on what’s already in the system. If a model fits entirely within a single GPU’s VRAM, that GPU alone will almost always be faster and simpler, since there’s no cross-device pipeline overhead and no dependency on a slower interconnect like Oculink. The split configuration earns its keep specifically when a model is too large for any single available GPU but small enough to fit across the combined memory of two.

In the tested scenario, a 130GB Q8-quantized model that fit on neither the Nvidia GPU (96GB) nor the AMD APU (128GB shared pool, with less available in practice) alone became runnable by dividing it between them, producing 37 to 38 tokens per second. That’s slower than either GPU running a model that fits natively, but it’s the difference between running a very large, higher-quality quantization locally versus not running it at all. Power draw stayed modest throughout, with the 600-watt-capable Nvidia card pulling only 130 to 150 watts and the AMD side around 76 watts, suggesting there’s headroom for larger batches or faster interconnects to close some of the performance gap.

## Frequently Asked Questions

### What is Vulkan and why does it matter for AI inference?

Vulkan is a cross-platform graphics and compute API that works on both Nvidia and AMD hardware. Unlike CUDA (Nvidia-only) or ROCm (AMD-only), Vulkan lets an inference engine like llama.cpp run a single model across GPUs from different vendors in the same process.

### Can I split any model across two GPUs, or does it need to be MoE?

