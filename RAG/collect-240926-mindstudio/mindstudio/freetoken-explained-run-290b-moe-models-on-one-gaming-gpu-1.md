---
id: collect-240926-mindstudio/mindstudio/freetoken-explained-run-290b-moe-models-on-one-gaming-gpu-1
title: "freetoken-explained-run-290b-moe-models-on-one-gaming-gpu"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Z.ai"]
dates: []
keywords: ["gpu", "moe", "agent", "agents", "attention", "benchmark", "claude", "compute", "consumer", "cost", "deepseek", "glm"]
source: docs/RAG/clean_en/mindstudio/freetoken-explained-run-290b-moe-models-on-one-gaming-gpu.md
source_anchor: ""
source_lines: [1, 68]
sha256: 1d2d62555f8daf7105f2ee781c1b518579d7e86d4773658f9c8a8e68298e173c
---

# freetoken-explained-run-290b-moe-models-on-one-gaming-gpu

<!-- source: https://www.mindstudio.ai/blog/freetoken-run-large-moe-models-locally -->

## What is FreeToken?

FreeToken is a local inference tool built to run mixture of experts (MoE) models that are far too large to fit in a single GPU’s VRAM. Instead of loading an entire model onto the GPU, it keeps most of the model’s weights in system RAM and only pulls the specific “experts” a given token needs onto the GPU at the moment they’re required. That lets a model with hundreds of billions of parameters, like GLM, DeepSeek, or Miniax, run on a single consumer or workstation GPU without a multi-GPU cluster or an API subscription.

## TL;DR

- FreeToken targets **mixture of experts models** such as GLM, DeepSeek, and Miniax, which are typically hundreds of billions of parameters and normally require multiple GPUs to serve.
- Its core trick is **expert offloading** : only the small subset of experts a token actually activates gets loaded onto the GPU, while the rest of the model sits in system RAM.
- The tool runs an automatic **hardware benchmark** at setup that compares CPU compute speed against PCIe transfer speed to decide whether to use hybrid mode (splitting compute between CPU and GPU) or offload mode (streaming experts over PCIe).
- In a demonstration on an A6000 GPU, a Qwen 3.6 model in NVFP4 quantization loaded in about 30 seconds and settled into roughly 44GB of VRAM usage, with over 21GB reserved just for the KV cache.
- FreeToken ships with a built-in terminal chat (**FT shell** ) that shows live stats like tokens per second, expert cache hit rate, and VRAM usage while you chat.
- It integrates with **six coding agents** , including Codex, DSH, Hermes agent, OpenClaw, OpenCode, and Claude, by pointing them at the local FreeToken server instead of a cloud API.
- No manual tuning is required to start serving a model. FreeToken reads the checkpoint and GPU specs and picks attention backend, MoE backend, and cache sizes automatically.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

## How does FreeToken let you run models bigger than your VRAM?

The problem FreeToken solves is specific to MoE architecture. In a dense model, every parameter is used for every token, so the whole model has to sit in fast memory to be useful. MoE models work differently: for each token, a router selects only a handful of “expert” sub-networks out of the hundreds available in the full model. Most of the model’s parameters go unused on any given forward pass.

Traditional serving setups don’t take advantage of this. They still load the entire model into VRAM just in case any expert gets picked, which is why serving a 300+ billion parameter MoE model conventionally demands eight or more GPUs wired together.

FreeToken flips that assumption. It keeps a small working set of frequently used experts cached on the GPU, sitting alongside the KV cache, while the rest of the model’s experts stay in ordinary system RAM. When the router calls for an expert that isn’t already cached on the GPU, FreeToken either streams that one expert over the PCIe bus and runs it, or, in hybrid mode, computes it directly on the CPU instead of moving it at all. Either way, you’re only ever paying the VRAM cost for the experts actively in flight, not the entire model.

## How does the benchmark decide between hybrid and offload mode?

Before serving anything, FreeToken runs a benchmark pass on the actual machine it’s installed on. It runs real kernels on the GPU for each quantization format present in the model and measures the GPU’s compute throughput, then compares that against the CPU’s compute speed and the PCIe transfer speed for moving data between CPU and GPU.

The logic is straightforward: if the CPU can compute an expert’s output faster than the data could be streamed over PCIe, FreeToken chooses hybrid mode and does that expert’s math on the CPU. If moving the weights over PCIe is faster than local CPU computation, it picks offload mode and streams the expert to the GPU instead. In one demonstrated run using an NVFP4-quantized model on an A6000 GPU, the benchmark selected offload mode automatically. This decision happens once per model/hardware combination and doesn’t require the user to manually configure anything.

## What happens when you actually serve a model?

Once the benchmark completes, launching the server is a single command. FreeToken reads the model checkpoint and the GPU’s specifications to automatically pick the attention backend, the MoE backend, and cache sizes, including how much VRAM to reserve for the KV cache (the memory that holds context as a conversation grows). In the demonstrated setup, the model weights loaded in about 30 seconds, with just over 21GB set aside for KV cache alone.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

After loading, the server compiles CUDA graphs, a one-time warm-up step that speeds up every subsequent generation. Total VRAM consumption in that example landed just over 44GB, close to the ceiling of the GPU being used. None of these steps required manual tuning: no setting tensor parallelism, no manually picking a quantization kernel, no hand-configuring memory splits.

## Is FreeToken worth using instead of an API or a multi-GPU setup?

The appeal is straightforward: it removes the two usual costs of running frontier-scale MoE models, which are a multi-GPU server rig and a recurring API bill. If you already own a capable single GPU, and gaming GPUs or workstation cards like an A6000 qualify, you can serve models that would otherwise need a rack of GPUs.

The tradeoff is speed and complexity versus a fully dedicated multi-GPU deployment. Streaming experts over PCIe or computing them on CPU is inherently slower than having every expert already resident in VRAM. FreeToken’s benchmark-driven mode selection exists specifically to minimize that penalty, but it’s still bound by the physical limits of a single card’s VRAM and a single PCIe link. For local experimentation, coding agent workflows, or privacy-sensitive use where you don’t want data leaving your machine, that tradeoff is usually worth it. For high-throughput production serving of a huge model to many concurrent users, dedicated multi-GPU infrastructure still wins.

FreeToken also includes a terminal-based chat interface (FT shell) that displays live metrics while you interact with a model: tokens per second, expert cache status (including how full the cache is and its LRU/reuse behavior), current KV cache usage, and total VRAM consumption. That visibility makes it easier to tell whether a given model and hardware combination is actually running efficiently or straining against the ceiling of available VRAM.

## How does FreeToken work with coding agents?

Beyond the built-in chat shell, FreeToken can serve as a local backend for existing coding agent tools. It supports pointing agents such as Codex, DSH, Hermes agent, OpenClaw, OpenCode, and Claude at its local server instead of a cloud provider’s API. In practice this means launching the agent with a configuration flag or URL pointing to the FreeToken server’s local endpoint, and the agent will use whichever model FreeToken is currently serving. This lets developers who already use these agents swap in a fully local model without switching tools or workflows, and without sending code or prompts to an external API.

## Frequently Asked Questions

### What models can FreeToken run?

FreeToken is built for large mixture of experts models, including GLM, DeepSeek, and Miniax families, as well as smaller MoE models like Qwen 3.6. The common requirement is that the model use a sparse MoE architecture rather than a dense one, since the expert-offloading trick depends on only a fraction of parameters being active per token.

