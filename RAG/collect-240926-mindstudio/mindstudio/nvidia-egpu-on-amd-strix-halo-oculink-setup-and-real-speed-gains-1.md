---
id: collect-240926-mindstudio/mindstudio/nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains-1
title: "nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "nvidia", "agent", "agents", "compute", "consumer", "gpus", "inference", "memory", "moe", "parameters"]
source: docs/RAG/clean_en/mindstudio/nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains.md
source_anchor: ""
source_lines: [1, 59]
sha256: 9b97333e89bdd95b2d02207a4ed29b7bce4448755db3c969135f98c93966d226
---

# nvidia-egpu-on-amd-strix-halo-oculink-setup-and-real-speed-gains

<!-- source: https://www.mindstudio.ai/blog/nvidia-egpu-amd-strix-halo-mini-pc -->

## What is an Nvidia eGPU on a Strix Halo mini PC?

It’s a setup where an external Nvidia graphics card connects to an AMD Strix Halo mini PC through an Oculink port, letting the two run AI models side by side or even split a single model across both. Strix Halo is AMD’s APU platform that shares up to 128GB of memory between CPU and GPU, which already makes it capable of running large language models on its own. Adding an Oculink port turns that mini PC into something closer to a small desktop, able to accept a real external GPU (an eGPU) instead of relying only on the built-in graphics.

## TL;DR

- **Oculink is the enabler** : it’s a PCIe-based external port running at roughly 63GB per second, the same class of connection a desktop GPU normally uses, but exposed on the outside of a mini PC.
- **Small GPUs hit a wall fast** : an RTX 5080 with 16GB of VRAM ran about three times faster than the Strix Halo’s built-in GPU on models that fit, but collapsed to under 2 tokens per second once a model spilled past its memory limit.
- **Big Nvidia cards change the math entirely** : an RTX Pro 6000 with 96GB of VRAM ran a 32 billion parameter model at 69 tokens per second, about six times the built-in chip’s speed, and handled a 122 billion parameter mixture-of-experts model at over 120 tokens per second.
- **Mixture-of-experts models punch above their size** : a 122B parameter MoE model with only 10 billion active parameters at a time ran faster than a much smaller dense 32B model, because only a fraction of the weights are active per token.
- **Splitting one model across an Nvidia and an AMD GPU is possible** : using Vulkan as a common runtime, a model too big for either GPU alone was divided (roughly 88GB on the Nvidia side, 34GB on the AMD side) and still produced usable output around 37 to 38 tokens per second.
- **Software compatibility runs through Vulkan, not CUDA or ROCm alone** : Nvidia’s CUDA and AMD’s ROCm (Radeon Open Compute) don’t talk to each other, but Vulkan works across both, which is what makes cross-vendor model splitting feasible.
- **The extra port has tradeoffs** : the newer mini PC with Oculink costs more, has fewer other ports, and is less stable on its stand than the previous generation, which lacks Oculink but is otherwise similar.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

## Why does Strix Halo need an external GPU at all?

Strix Halo’s integrated GPU (the Radeon 8060 in the tested unit) is genuinely capable for a mobile-class chip. It shares a large unified memory pool with the CPU, up to 128GB, which lets it load models that would choke a typical gaming laptop or desktop GPU with limited VRAM. That’s why Strix Halo mini PCs became popular for local LLM work in the first place: big memory capacity at a lower price than a workstation GPU.

But shared memory bandwidth and integrated graphics compute are still slower than a discrete GPU built for this job. On models that comfortably fit in a Strix Halo’s memory, the built-in GPU does fine. Once you want more raw speed, or you want to run a model that benefits from dedicated VRAM and higher memory bandwidth, an external GPU becomes attractive. The catch has always been that mini PCs generally don’t expose a way to attach one. Oculink solves that specific problem.

## How much faster is an eGPU, actually?

The gains depend heavily on whether the model fits in the eGPU’s VRAM.

With an RTX 5080 (16GB VRAM) attached over Oculink, a model that fit inside that memory ran about three times faster than the same model on the Strix Halo’s built-in GPU. That’s a meaningful jump for a relatively affordable consumer card.

The moment a model exceeded 16GB, though, performance didn’t just dip, it fell off a cliff: output dropped to roughly 1.6 tokens per second, far slower than the 11-plus tokens per second the built-in GPU delivered on the same model. The lesson is straightforward: an eGPU with limited VRAM is a speed boost only within its memory ceiling. Past that point, the built-in Strix Halo GPU with its large shared memory pool wins by a wide margin.

Move up to something like the RTX Pro 6000 with 96GB of VRAM, and the equation flips again. That card ran a 32 billion parameter model at 69 tokens per second, about six times the speed of the integrated GPU, while pulling close to its full power draw. It could also load a 122 billion parameter mixture-of-experts model (a Q4 quantized version, around 76.5GB on disk) and hit around 121 tokens per second, faster than the smaller dense model, because a mixture-of-experts architecture only activates a portion of its total parameters for any given token.

## Can you actually split a model across an Nvidia GPU and an AMD GPU?

Yes, though it requires working around the fact that Nvidia and AMD use incompatible software stacks by default. Nvidia GPUs run on CUDA, AMD GPUs run on ROCm (Radeon Open Compute), and the two don’t interoperate. The workaround is Vulkan, a cross-platform graphics and compute API that both vendors support. Tools like LM Studio let you pick Vulkan as the inference runtime instead of CUDA or ROCm, which is what makes splitting a model across both GPUs technically possible.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

In one test, a 122 billion parameter model at Q8 quantization (roughly 130GB on disk, and specifically chosen because that’s too large for either GPU individually) was split by layer: about 70% of the model’s layers loaded onto the Nvidia GPU, the rest onto the AMD integrated GPU. Every token generated has to pass through both halves of the model in sequence. That test produced a memory split of roughly 88GB on the Nvidia side and 34GB on the AMD side, running at about 37 to 38 tokens per second. It’s slower than either GPU running a model that fits entirely in its own memory, partly because the Oculink connection itself becomes a bottleneck when data has to move between the two GPUs.

## What goes wrong when you try this?

A few practical snags showed up. The RTX Pro 6000, a physically large card, initially caused the mini PC to fail to boot at all when connected, a black screen tied to firmware not handling a card of that size at startup. The workaround was to boot the machine without the card attached, then power on the GPU afterward and let the system detect it, essentially hot-plugging it after boot. That’s explicitly against normal eGPU and Oculink guidance, since hot-plugging isn’t a supported use case, but it was necessary to get the card running.

Memory allocation to the integrated GPU also has to be set manually in the BIOS on machines that support it. Left on automatic settings, the system didn’t properly connect with the external GPU, so users need to explicitly allocate memory (in one case, 96GB) to the integrated GPU for the setup to behave predictably.

## Is buying an Oculink-equipped Strix Halo mini PC worth it?

It depends on whether you actually plan to attach an external GPU. The Oculink-equipped model carries a modest price premium over the prior generation that lacks the port, despite using the same Strix Halo chip and the same 128GB memory ceiling. For that premium, you gain the ability to attach a real desktop GPU, which matters a lot if you want faster inference on models that fit in that GPU’s VRAM, or if you want to experiment with splitting larger models across both GPUs.

