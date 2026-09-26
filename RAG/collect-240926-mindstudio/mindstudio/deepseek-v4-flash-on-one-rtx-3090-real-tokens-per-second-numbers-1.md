---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers-1
title: "deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Z.ai"]
dates: []
keywords: ["deepseek", "agentic", "agents", "compute", "consumer", "cost", "glm", "gpu", "inference", "latency", "llama", "llama.cpp"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers.md
source_anchor: ""
source_lines: [1, 59]
sha256: 064ea0315dcd056d1d8776fd5ae7e7c04789b1f868048ae074036394f61173ec
---

# deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers

<!-- source: https://www.mindstudio.ai/blog/freetoken-deepseek-v4-flash-single-3090 -->

## Can a single RTX 3090 run DeepSeek V4 Flash?

Yes, but the GPU is only part of the story. DeepSeek V4 Flash is a mixture-of-experts (MoE) model, which means most of its parameters sit idle for any given token and only a fraction (“experts”) get activated per pass. That property lets it run on a single 3090 or 4090 with the bulk of the model offloaded to system RAM instead of VRAM. In real testing through FreeToken’s desktop app, a single 3090 paired with a large pool of DDR4 RAM produced generation speeds in the 8 to 11 tokens per second range, depending on which experts were active and whether the model was accessed through a server API or the desktop client directly.

## TL;DR

- **DeepSeek V4 Flash ran at roughly 10 to 11 tokens per second** on a single RTX 3090 when served through Open WebUI, and closer to**8.8 tokens per second when run through FreeToken’s desktop client** directly, suggesting the desktop app carries some overhead versus a raw llama.cpp-style server.
- **System RAM capacity and bandwidth matter as much as the GPU** for MoE models like this one, because the inactive expert weights have to sit somewhere, and that somewhere is usually your DIMMs, not your VRAM.
- **128GB of system RAM was not enough** to comfortably run DeepSeek V4 Flash in this setup; 192GB worked, and something in the 156 to 168GB range is likely the practical floor for most users.
- **RAM speed changes real-world throughput noticeably.** The test rig ran DDR4 at 2400MT/s, well below the 3200MT/s+ that current platforms support, and the presenter expects DDR5 systems to roughly double throughput over slower DDR4.
- **A dense model tells a different story.** Qwen 3.8 27B in BF16 (a non-MoE, fully dense model) failed to load reliably in the FreeToken desktop beta, throwing a generic engine error, which is a reminder that dense models don’t get the same offload-friendly behavior MoE architectures do.
- **FreeToken is beta software with rough edges.** Tokens-per-second reporting was inconsistent between its desktop UI and server view, and not every model in its library loads cleanly yet.
- **Bigger MoE models like GLM 5.2 need serious RAM.** The app itself reported needing over 200GB of additional system memory to run a large GLM 5.2 variant, pointing to workstation or server-class RAM pools (256GB+) as the entry point for the largest local models.

## One coffee. One working app.

You bring the idea. Remy manages the project.

## What makes DeepSeek V4 Flash different from a dense model?

DeepSeek V4 Flash uses a mixture-of-experts architecture. Instead of running every parameter through every token, the model routes each token to a small subset of specialized sub-networks (“experts”). The result is a model with a very large total parameter count that only activates a fraction of those parameters on any given forward pass. That’s why MoE models like DeepSeek V4 Flash and GLM 5.2 are unusually well suited to consumer hardware: the active compute per token is small enough for a single GPU to handle, while the full weight set gets parked in system RAM and streamed in as needed.

This is fundamentally different from a dense model like Qwen’s 27B variant, where every parameter participates in every forward pass. Dense models don’t benefit from the same RAM-offload trick nearly as well, because there’s no way to selectively load only the “active” part of the network. That distinction showed up directly in testing: the MoE model ran (slowly, but it ran), while the dense 27B model in BF16 repeatedly failed to start.

## How many tokens per second should you actually expect?

In hands-on testing, DeepSeek V4 Flash on a single RTX 3090, with the rest offloaded to system RAM, generated text at:

- **About 10 to 11 tokens per second** when accessed through an Open WebUI front end talking to a llama.cpp-style backend.
- **About 8.8 tokens per second** when run directly through FreeToken’s desktop client, with the app briefly reporting an inconsistent 1.8 tokens per second on an earlier run before stabilizing.

Neither number is fast by API standards. It’s roughly interactive chat speed, workable for conversation but not for high-throughput or agentic workflows that fire off dozens of generations back to back. The presenter was explicit that this setup is not meant for anything latency-sensitive; it’s for people who want frontier-adjacent model quality on hardware they already own, and are willing to trade speed for that.

Token throughput also wasn’t perfectly stable run to run. Because MoE models route tokens to different experts depending on content, the exact mix of experts loaded and swapped during a conversation affects speed, which is part of why the numbers “vary wildly,” in the presenter’s words, even on identical hardware.

## How much RAM do you actually need?

This is the detail most casual “run it on one GPU” claims skip. VRAM alone won’t cut it for a model of this scale; system RAM does the heavy lifting for the parameters that aren’t active on the GPU at any given moment.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

In testing, 128GB of system RAM was not sufficient to run DeepSeek V4 Flash. Bumping the allocation to 192GB worked, and the presenter estimated that something in the 156 to 168GB range would likely be the realistic minimum for most people attempting this. For context, the test system used older DDR4 running at 2400MT/s, a speed well behind what current DDR4 (3200MT/s) or DDR5 platforms can deliver. Since MoE inference is memory-bandwidth bound once the active experts have to be fetched from RAM, slower memory directly caps your tokens per second. The presenter expects DDR5 systems to roughly double throughput compared to the DDR4 2400 setup used here.

For anyone shopping hardware specifically to run models like this, the practical guidance is: 64GB of system RAM is a reasonable floor for smaller MoE models, 96 to 128GB opens up more headroom, and workloads like DeepSeek V4 Flash or larger MoE releases such as GLM 5.2 push requirements into the 192GB-and-up range. FreeToken’s app will proactively tell you if you’re short: in one test it reported needing over 200GB of additional system memory to load a large GLM 5.2 variant on a system that didn’t have enough free capacity.

## Is FreeToken’s desktop app worth using right now?

It depends on what you want out of it. As a way to get a mixture-of-experts model running locally without hand-configuring a llama.cpp server, quantization format, and offload settings yourself, it lowers the barrier meaningfully. It ships as a Windows distributable, a Ubuntu package, a common AppImage, and an Arch Linux package, and it includes a Hugging Face-backed model download flow plus a built-in cost comparison against paying for the same model via API.

That said, it’s explicitly beta software, and it showed it during testing. The Qwen 3.8 27B BF16 model failed to launch with a generic, unhelpful error message both through the desktop UI and after a restart attempt. Tokens-per-second reporting was inconsistent between the desktop client and a server-side Open WebUI view of the same model, undercounting performance in the desktop app’s own display. Model coverage is also limited compared to the wider universe of open-weight releases. None of that erases the core value: for someone with a single 3090 or 4090 and a well-stocked RAM pool, it is one of the more approachable ways to get a large MoE model running without becoming a full-time systems administrator.

## Frequently Asked Questions

### What is DeepSeek V4 Flash?

