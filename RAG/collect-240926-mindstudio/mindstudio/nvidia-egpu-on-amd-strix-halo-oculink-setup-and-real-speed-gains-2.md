---
id: collect-240926-mindstudio/mindstudio/nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains-2
title: "nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "nvidia", "gpus", "inference", "memory", "parameters"]
source: docs/RAG/clean_en/mindstudio/nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains.md
source_anchor: ""
source_lines: [60, 88]
sha256: a30d1775e09248f3fba9aec7addb59b7eda6d91f62573f4663258327523d5700
---

# nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains

What you give up is some practicality. The Oculink model has fewer other ports than the previous generation, and its physical stand is less stable, making it easier to knock over and unable to lay flat the way the older design can. If Oculink expansion isn’t a priority, the earlier model without that port remains a capable, slightly cheaper Strix Halo machine with the same memory and integrated GPU.

For anyone specifically interested in pairing Nvidia GPUs (from a midrange 16GB card up to a 96GB workstation card) with an AMD Strix Halo system, the Oculink port is the feature that makes it possible at all. Without it, this entire category of dual-vendor, dual-GPU inference simply isn’t on the table.

## Frequently Asked Questions

### What is Oculink and why does it matter for eGPUs?

Oculink is an external PCIe connection standard that runs at high speed (around 63GB per second in this setup), similar to the internal slot a desktop GPU would use. It matters because it’s rare on mini PCs, and without it there’s no practical way to attach a full-size external GPU to a compact system like a Strix Halo machine.

### Does a bigger GPU always mean faster local LLM inference?

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

Only if the model fits in that GPU’s VRAM. A 16GB card can be several times faster than an integrated GPU on models within that memory limit, but performance collapses if the model exceeds available VRAM. A GPU with much larger VRAM, like a 96GB workstation card, avoids that cliff for a wider range of model sizes.

### Can Nvidia and AMD GPUs really work together on one model?

Yes, using Vulkan as a shared runtime instead of Nvidia’s CUDA or AMD’s ROCm, which don’t interoperate. A model can be split by layers across both GPUs, with each token passing through both halves in sequence. It works, but the connection between the GPUs can become a bottleneck.

### Why did a 122 billion parameter model run faster than a 32 billion parameter model?

Because the larger model uses a mixture-of-experts architecture, where only a fraction of the total parameters (around 10 billion in this case) are active for any given token. The smaller model tested was a traditional dense architecture, where all parameters are active at once, making it computationally heavier despite having fewer total parameters.

### Is it safe to hot-plug an eGPU into an Oculink port?

Hot-plugging isn’t officially supported for eGPUs or Oculink connections. In practice, some large GPUs can cause boot failures if connected at startup, requiring the workaround of booting without the card and then powering it on afterward. This isn’t standard guidance and carries some risk.
