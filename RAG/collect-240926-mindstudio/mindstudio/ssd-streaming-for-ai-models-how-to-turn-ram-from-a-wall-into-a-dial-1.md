---
id: collect-240926-mindstudio/mindstudio/ssd-streaming-for-ai-models-how-to-turn-ram-from-a-wall-into-a-dial-1
title: "ssd-streaming-for-ai-models-how-to-turn-ram-from-a-wall-into-a-dial"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["attention", "compute", "consumer", "cost", "deepseek", "inference", "latency", "memory", "mixture of experts", "moe", "parameters", "quantization"]
source: docs/RAG/clean_en/mindstudio/ssd-streaming-for-ai-models-how-to-turn-ram-from-a-wall-into-a-dial.md
source_anchor: ""
source_lines: [1, 98]
sha256: 7f5e051f085dee21251de15822cea4915e47658ca0269ab8bb986faca2a2afb7
---

# ssd-streaming-for-ai-models-how-to-turn-ram-from-a-wall-into-a-dial

<!-- source: https://www.mindstudio.ai/blog/ssd-streaming-ai-models-ram-dial -->

## When RAM Becomes the Bottleneck, Not the Solution

If you’ve ever tried running a large language model locally, you’ve probably hit the same wall. You find a model you want to use — maybe a Mixtral variant, a DeepSeek model, or something else with real capability — check the requirements, look at your available RAM, and discover you’re about 20GB short. The model doesn’t run at a reduced quality. It just doesn’t run.

That binary outcome is the core frustration of local AI deployment. RAM has always functioned as an on/off switch: either the whole model fits, or nothing works. SSD streaming flips that logic entirely. By storing model weights on disk and loading them on demand, it turns RAM from a fixed ceiling into an adjustable dial. Dwarf Star is one of the more interesting implementations of this idea, specifically targeting the expert weights in Mixture of Experts (MoE) models.

This article explains how SSD streaming works, why MoE architectures make it practical, what trade-offs you’re accepting, and who actually benefits from it.

## The Real Problem with Running Large Models Locally

Most people frame the local AI problem as “my computer isn’t powerful enough.” That’s not quite right. The more precise framing is: model weights have to live somewhere your processor can access them quickly, and RAM is the fastest accessible space in that chain.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

When you load a model, its weights get copied into RAM. Every inference — every token generated — reads from those weights repeatedly. If the weights don’t fit in RAM, the system starts reading from swap space (disk), and inference slows to a crawl or crashes entirely.

For smaller models (7B parameters, quantized), this is manageable. A 4-bit quantized 7B model runs comfortably in 6–8GB of RAM. But the models that people actually want to run locally — the ones with strong reasoning, broad knowledge, and real-world utility — tend to be 30B, 70B, or larger. A 70B model at 4-bit quantization still needs around 40GB of RAM.

Most consumer machines top out at 16–32GB. That gap isn’t small.

### The Scale Problem Gets Worse with Capability

There’s an uncomfortable relationship between model capability and size. Broadly speaking, more parameters means better performance on complex tasks. The models people reach for when they want serious results — multi-step reasoning, code generation, nuanced writing — are the ones that won’t fit.

This creates a practical tiering problem. You can run the models you need for simple tasks, but the moment you need something harder, you’re looking at API calls, cloud costs, or a hardware upgrade.

SSD streaming is an attempt to break that relationship by changing where model weights live, not by shrinking the model itself.

## Mixture of Experts: Why Not All Weights Are Created Equal

To understand why SSD streaming is practical at all, you need to understand Mixture of Experts (MoE) architecture. It’s the key that makes on-demand weight loading viable.

In a standard dense model, every layer processes every input. If a model has 70 billion parameters, all 70 billion are involved in generating each token. That’s why large dense models need so much RAM — everything has to be loaded and accessible.

MoE models work differently. Instead of one large feedforward network in each layer, they have multiple smaller networks called “experts.” A routing mechanism looks at each input token and decides which experts to activate — typically just two or four out of eight, sixteen, or more available.

### What This Means Practically

A model like Mixtral 8x7B has eight expert networks per layer. For any given token, only two of them fire. The other six sit idle.

Total parameter count is high — around 46 billion parameters — but active parameters per forward pass are much lower, closer to 12–13 billion. That’s why MoE models can punch above their inference weight: you get the parameter breadth of a large model with the compute cost of a smaller one.

But here’s the implication that makes SSD streaming interesting: if only two experts are active at any given moment, you don’t need all eight experts in RAM simultaneously. You only need the two that are about to be used.

This is the architectural opening that Dwarf Star exploits.

## How SSD Streaming Actually Works

The core idea is straightforward: keep the expert weights on your SSD instead of in RAM, and load them into memory only when the router selects them.

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

Dwarf Star stores the expert weight tensors on disk. The non-expert components of the model — the attention layers, layer normalization, routing weights — stay in RAM because they’re used for every token and need to be fast. But the expert weights, which represent a large fraction of total parameters, get offloaded to disk.

When inference runs and the router selects experts for a given layer, the system fetches those expert weights from the SSD, runs the computation, and moves on. The weights don’t need to stay loaded — they’ll only be needed again if that expert is selected for a future token.

### The Math Behind It

Take Mixtral 8x7B as an example. Total weights are roughly 90GB in full precision, around 26GB at 4-bit quantization. A large chunk of that is expert weights.

If you store experts on disk and only pull the active two per layer per token, your RAM requirement drops dramatically. The non-expert weights — the part that actually needs to stay in memory — can fit comfortably on a machine with 16–24GB of RAM.

You’re not running a smaller model. You’re running the same model, with the same output quality, but with a much smaller RAM footprint.

### Prefetching Matters

A naive implementation would have noticeable latency: pick an expert, wait for the disk read, compute, repeat. That would be slow enough to make the approach impractical.

Better implementations use prefetching. Because the routing decision for layer N is known before layer N+1 executes, the system can start loading the next layer’s experts from disk while the current computation runs. With a fast NVMe SSD (which can sustain 5–7+ GB/s sequential reads), and with careful prefetching logic, the disk reads can largely happen in parallel with compute.

The result is that SSD streaming doesn’t feel like constant disk thrashing — it feels like slightly slower inference than native RAM, not like reading from a hard drive.

## SSD Speed Is Not Optional

This only works if your storage is fast. A traditional SATA SSD tops out around 550 MB/s sequential read. A modern NVMe SSD using PCIe 4.0 hits 5,000–7,000 MB/s. That’s a 10x difference, and it’s the difference between a streaming approach that’s workable and one that makes inference painfully slow.

If you’re considering running Dwarf Star or similar expert-streaming setups, your storage tier matters as much as your RAM:

- **NVMe PCIe 4.0 or 5.0** : This is where streaming inference becomes genuinely usable. Read speeds are fast enough that prefetching keeps the pipeline fed.
- **NVMe PCIe 3.0** : Still viable, but you’ll notice the latency more. Tokens per second will be lower.
- **SATA SSD** : Too slow for comfortable inference on most MoE models. You’ll likely bottleneck here before you bottleneck on compute.
- **HDD** : Not a realistic option. Seek times alone make this unworkable.

This is worth flagging because “runs on your laptop” can mean very different things. A machine with a fast NVMe drive and 16GB RAM can have a genuinely good experience. The same RAM with a slower drive won’t.

