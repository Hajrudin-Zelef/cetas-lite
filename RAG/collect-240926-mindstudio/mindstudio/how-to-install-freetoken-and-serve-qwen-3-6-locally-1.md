---
id: collect-240926-mindstudio/mindstudio/how-to-install-freetoken-and-serve-qwen-3-6-locally-1
title: "how-to-install-freetoken-and-serve-qwen-3-6-locally"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "MiniMax", "Z.ai"]
dates: []
keywords: ["qwen", "agent", "agents", "attention", "benchmark", "claude", "compute", "consumer", "context window", "cost", "deepseek", "glm"]
source: docs/RAG/clean_en/mindstudio/how-to-install-freetoken-and-serve-qwen-3-6-locally.md
source_anchor: ""
source_lines: [1, 85]
sha256: 3f7fed904d73a5bbc69759afcd41eb1b6e4777b519bad31aa3a9f3fd438cdeb2
---

# how-to-install-freetoken-and-serve-qwen-3-6-locally

<!-- source: https://www.mindstudio.ai/blog/freetoken-install-guide-qwen3-6 -->

## What is FreeToken and what problem does it solve?

FreeToken is a serving tool built to run frontier mixture of experts (MoE) models, the kind with hundreds of billions of parameters, on consumer hardware that doesn’t have anywhere near enough VRAM to hold them. Models like GLM, DeepSeek, and MiniMax are MoE architectures, meaning that even though the full model is enormous, each token only activates a small subset of “experts” inside the network. FreeToken exploits that by keeping most of the model parked in system RAM and only pulling the specific experts a token needs onto the GPU as it goes. The result is that you can serve models technically far bigger than your VRAM allows, on a single GPU, without a multi-GPU rig and without paying for API access.

## How does FreeToken’s expert offloading actually work?

Normally, serving a huge MoE model means loading every expert into VRAM at once, which is why models in the hundreds of billions of parameters typically need multiple high end GPUs. FreeToken avoids this by keeping a small working set of experts cached on the GPU next to the KV cache, while the rest of the experts sit in regular system RAM.

When the model’s router selects an expert that isn’t already sitting in the GPU cache, FreeToken handles it one of two ways:

- It streams just that one expert over PCIe onto the GPU, uses it, and moves on (offload mode).
- In hybrid mode, it computes that expert directly on the CPU instead of transferring it at all.

Either way, you’re only ever paying the VRAM cost for the handful of experts actively in flight, not the entire model. For a model with hundreds of billions of parameters, that’s the difference between needing an eight-GPU server and running it on a single card.

## How do you install FreeToken and download a model?

The setup process demonstrated on a Lubuntu system with a single GPU card follows a straightforward path:

1. Create a virtual environment (the demo used UV) and install FreeToken into it.
2. Install the Hugging Face CLI, also via UV, to handle model downloads.
3. Download the target model. In this walkthrough, that model was Qwen 3.6.

Installation and download both took a few minutes, with no unusual dependencies or manual configuration required beyond the standard Python environment tooling.

## What does the FreeToken benchmark step do?

Before serving anything, FreeToken runs a benchmark against your specific machine. This step measures your actual hardware bottleneck rather than assuming one, by running real kernels on your GPU for each quantization format and comparing your CPU’s compute speed against your PCIe transfer speed for that same format.

The logic that comes out of this benchmark determines how FreeToken will split work between CPU and GPU:

- If CPU compute beats PCIe transfer speed by a sufficient margin, FreeToken chooses **hybrid mode** , splitting expert computation between CPU and GPU.
- If PCIe transfer wins, FreeToken chooses **offload mode** , streaming experts over PCIe as needed.

In the demo, running an NVFP4 quantized model on an A6000 GPU, the benchmark landed on offload mode, and that choice was then automatically applied when the server launched.

## How do you serve a model with FreeToken?

Once the benchmark completes, you launch the server with the FT serve command, pointing it at the downloaded model. From there, FreeToken automatically configures itself: it picks the attention backend, the MoE backend, and cache sizes directly from the model checkpoint and your hardware specs, with the offload/hybrid decision inherited from the benchmark step.

In the demonstrated run, the server:

- Loaded model weights in about 30 seconds.
- Reserved just over 21 GB of VRAM specifically for the KV cache (the context window space).
- Compiled CUDA graphs as a one-time warm-up step, which speeds up every generation that follows.
- Settled into steady-state VRAM usage of just over 44 GB on the A6000 card.

A simple curl request against the running server confirms which model is currently being served, useful for verifying configuration before connecting a client or agent.

## What is FT shell and how do you use it?

FT shell is FreeToken’s built-in terminal chat interface. After activating the environment, running FT shell drops you straight into a chat prompt where you can send questions and watch the model stream tokens back in real time.

What makes FT shell useful beyond basic chat is its live status bar, which surfaces several diagnostics at once:

- **Tokens per second** , the real-world generation speed on your hardware.
- **Cache and LRU stats** , showing how full the GPU’s expert cache is and how effectively it’s reusing recently accessed experts instead of refetching them.
- **KV usage** , how much of the reserved context memory (21 GB in the demo) is actually in use for the current exchange.
- **Mamba state** , tracking the linear attention state cache for architectures that combine it with standard attention.
- **VRAM consumption** , the actual number that determines how close you are to your card’s ceiling (42.9K out of 48 GB in the demo).

This dashboard gives a direct read on whether your setup is running efficiently or bumping against hardware limits, without needing separate monitoring tools.

## How do you connect coding agents to a FreeToken server?

FreeToken supports pointing external coding agents at its local server instead of a hosted API. The tool is compatible with several agents, including Codex, Hermes agent, OpenClaw, OpenCode, and Claude.

The pattern is consistent across agents: instead of installing a fresh agent pointed at a cloud API key, you launch the agent with a command that redirects it to your local FreeToken server address. If the agent isn’t already installed, the launch command installs it automatically. In the demo, both Hermes agent and Claude Code were launched this way, each one picking up the locally served Qwen 3.6 model instead of calling out to Anthropic’s or another provider’s API. Once connected, you interact with the agent normally, the only difference is that inference happens on your own hardware.

## Is FreeToken worth setting up on your own machine?

For anyone who wants to run genuinely large MoE models without renting a multi-GPU cloud instance or paying per-token API costs, FreeToken addresses a real, specific bottleneck: VRAM ceiling versus model size. The benchmarking step removes most of the guesswork around hybrid versus offload configuration, and the automatic backend selection at serve time means you’re not manually tuning attention or MoE backends for each model.

The tradeoff is that you’re still bound by your system RAM capacity (since that’s where the bulk of the model lives) and by PCIe or CPU compute speed for whichever experts aren’t cached. Larger models than Qwen 3.6, such as GLM or DeepSeek, will push those constraints harder and require more disk space and RAM to hold comfortably. For developers experimenting with frontier-scale models on a single gaming or workstation GPU, though, it’s a practical way to get real inference speeds without a server rack.

## Frequently Asked Questions

### What GPU do I need to run FreeToken?

The demonstration used an A6000 GPU, landing on offload mode for an NVFP4 quantized Qwen 3.6 model with around 44 GB of VRAM in steady-state use. Exact requirements depend on the model size and quantization format you choose, since FreeToken’s benchmark determines the optimal split for your specific hardware rather than assuming a fixed requirement.

### Does FreeToken require multiple GPUs?

