---
id: collect-240926-mindstudio/mindstudio/how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine-1
title: "how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["inference", "qwen", "agent", "agents", "apache", "context window", "cost", "embedding", "embeddings", "gpu", "gpus", "inference engine"]
source: docs/RAG/clean_en/mindstudio/how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine.md
source_anchor: ""
source_lines: [1, 66]
sha256: 9f980fc6a484bdef42a3f2f07bbb2a5929587824da8fd904bb9b98ef53bc8db5
---

# how-to-run-qwen-3-8-locally-with-the-superlinked-inference-engine

<!-- source: https://www.mindstudio.ai/blog/run-qwen-3-8-locally-sie -->

## What is the Superlinked Inference Engine?

The Superlinked Inference Engine (SIE) is an open-source, Apache 2.0 licensed inference server that runs the full stack of models an AI agent needs, from embedding models to full generation models, on your own hardware. Instead of running separate servers for different model types, SIE acts as a single server that loads whatever model you call and serves it with a configuration matched to your specific GPU. It’s available on GitHub, and it can also be run without local hardware through SIE Cloud, a hosted version of the same engine with free inference grants for people who don’t want to manage GPUs themselves.

## TL;DR

- The **Superlinked Inference Engine** is a single self-hosted server that runs embeddings and generation models from one install, avoiding the need to juggle separate inference stacks.
- Installation is minimal: setting up an environment and running one install command gets the engine running on a fresh Ubuntu machine in about a minute.
- Qwen 3.8’s **27 billion parameter model** landed in SIE with support for running it locally on day one of release, tested on an RTX Pro 6000.
- SIE ships **hardware-specific profiles** for the same model weights, including a safe default profile, an RTX Pro 6000 256K profile with speculative decoding, and separate H100/H200 profiles for data center cards.
- Selecting the right profile is a single flag in the generation call, which is what turns on features like a 256K context window and **speculative decoding** for extra speed.
- SIE only loads models on demand rather than keeping everything resident in VRAM, so nothing consumes GPU memory until you actually call generate.
- In a real test generating a full self-contained HTML page with ten animated tabs, the RTX Pro 6000 profile produced fast token-per-second throughput and finished the task in under 30 seconds.

## How do you install SIE from scratch?

Setting up SIE on a fresh machine is designed to be fast. On a clean Ubuntu system, the process comes down to two commands: one to set up the environment and one to install the engine itself. There’s no lengthy dependency chase or manual driver wrangling baked into the basic setup. Once the install finishes, you launch the inference engine (the same server handles both embedding and generation workloads), and a health check confirms it’s serving requests locally before you send it any real work.

This simplicity is part of the pitch. Most self-hosted inference setups require you to hand-tune batch sizes, context windows, quantization settings, and GPU-specific flags yourself. SIE tries to collapse that tuning work into a profile you select at call time rather than a system you configure by hand.

## What are hardware profiles and why do they matter?

A model like Qwen 3.8 isn’t a single fixed artifact once it’s inside SIE. The same weights get paired with different configuration profiles depending on the GPU you’re running on. These profiles are not different models, they’re different runtime configurations tuned for specific hardware:

- A **safe default profile** that runs on nearly any GPU with no speculative decoding, prioritizing compatibility over raw speed.
- An **RTX Pro 6000 256K profile** , tuned specifically for that card, with a 256,000-token context window and speculative decoding enabled.
- **H100 and H200 profiles** , aimed at squeezing additional throughput out of data center-class hardware.

The practical effect is that you don’t have to figure out the right context window, batch settings, or speculative decoding configuration for your card manually. You pick the profile that matches your GPU, and SIE applies the tuning that goes with it. That matters because getting these settings wrong is one of the most common ways people leave performance on the table when self-hosting large models.

## How does speculative decoding fit in?

Speculative decoding is a technique for speeding up text generation by using a smaller, faster “draft” model to propose several tokens ahead, which the larger target model then verifies in a batch rather than generating every token one at a time from scratch. When the draft’s guesses are accepted, generation moves faster than standard autoregressive decoding; when they’re rejected, the larger model falls back to its normal token-by-token process. It’s a way to get closer to the full model’s output quality while cutting down on the sequential computation that normally limits generation speed.

In SIE, speculative decoding isn’t something you configure by hand. It’s built into specific hardware profiles, like the RTX Pro 6000 256K profile, so turning it on is as simple as selecting that profile rather than tuning parameters yourself.

## How do you actually run Qwen 3.8 on SIE?

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

Once the engine is installed and healthy, running Qwen 3.8 comes down to a short script, roughly a dozen lines, that connects to the local SIE server and calls its generate function. The call specifies the model (Qwen 3.8 27B) and the profile to use (in this case, the RTX Pro 6000 256K profile). That single option flag determines whether you get a generic, compatibility-first run or one optimized for the full 256K context window with speculative decoding turned on.

One notable behavior: SIE doesn’t preload every model into VRAM at startup. Nothing occupies GPU memory until a generate call actually requests a specific model. When the first request comes in, SIE downloads and loads the model on demand, and VRAM usage climbs as the model and its context window get allocated. This on-demand loading is useful if you’re serving multiple models from one machine and don’t want them all resident in memory simultaneously.

A first “warm-up” run is worth doing before judging speed, since initial loading and caching add overhead that disappears on subsequent calls.

## Is running Qwen 3.8 on SIE fast in practice?

In a demonstrated test on an RTX Pro 6000, Qwen 3.8 27B was asked to generate a complete, self-contained HTML page containing ten tabs, each depicting a traditional grilled meat dish from a different, lesser-known country, complete with an animated flame and smoke effect per tab. This is a substantial single-shot generation task involving both code structure and creative detail across multiple sections.

The model completed the task with strong token-per-second throughput and a total wall time under 30 seconds. The output rendered correctly in-browser, with distinct meat shapes, skewer types, flame heights, and smoke effects varying across tabs, along with country-specific dish details (Uzbek shashlik, Georgian and Armenian skewer dishes, Azerbaijani tikka kebab, and others). The combination of fast generation and coherent, varied output is the core case for pairing a strong open model with a hardware-tuned inference profile: the model provides the knowledge and reasoning, and the tuned profile determines how quickly that gets served.

## Is self-hosting Qwen 3.8 with SIE worth it?

For teams that already run their own GPUs and want control over cost, latency, and data locality, SIE’s profile system removes a lot of the manual tuning that normally makes self-hosted inference painful. Rather than researching optimal context windows and speculative decoding setups for your specific card, you select a pre-built profile. For RTX Pro 6000, H100, or H200 owners, that’s a meaningful time save.

For people without dedicated hardware, SIE Cloud offers the same engine hosted, with free inference grants available, which lowers the barrier to trying Qwen 3.8 or other supported models without provisioning a GPU first.

## Frequently Asked Questions

### What is Qwen 3.8?

