---
id: collect-240926-nerdykings/nerdykings/deepseek-dual-path-doubler-les-perfs-ia-sans-gpu
title: "DeepSeek Dual Path: Doubling AI Performance Without GPUs"
domain: nerdykings
role: reference
task: reference
actors: ["China", "DeepSeek", "Nvidia", "United States"]
dates: []
keywords: ["deepseek", "gpu", "gpus", "agent", "agentic", "decode", "inference", "kv cache", "memory", "nvidia", "prefill", "throughput"]
source: docs/RAG/clean_en/nerdykings/deepseek-dual-path-doubler-les-perfs-ia-sans-gpu.md
source_anchor: ""
source_lines: [1, 77]
sha256: caea6ca39845e47f5eb709af77c0d71f2ecee23c2ad723d43ea98c75f0225042
---

# DeepSeek Dual Path: Doubling AI Performance Without GPUs

<!-- source: https://www.nerdykings.com/blog/deepseek-dual-path-gpu.html -->

# DeepSeek Dual Path: Doubling AI Performance Without GPUs

Companies spend hundreds of millions of dollars on AI servers filled with ultra-powerful GPUs. **And a large part of the time, these GPUs run at 20-30% of their capacity** — not because they're weak, but because they're waiting for data. DeepSeek has just published a paper, **Dual Path**, that solves a large part of the problem *without buying a single new GPU*.

## How an LLM Really Generates Text

Before understanding the trick, you need to understand what happens under the hood. When you send a prompt to an LLM, two very distinct things happen:

**Phase 1 — Prefill.** The model reads your message all at once, analyzes all your tokens in parallel, calculates the state of each word in its memory, and stores it all in the KV cache. It's brutal, it's dense, it demands enormous power. GPUs run at 90-95% utilization.

**Phase 2 — Decode.** The model generates the response *one token at a time*. And that's where everything gets complicated: to generate each new word, it must re-read the entire KV cache from memory. Again and again. This requires very little computation but an enormous amount of memory bandwidth. Result: **GPUs drop to 20-30% utilization**, all that power spinning its wheels waiting for memory to keep up.

The perverse thing: decode accounts for **95% of a request's lifetime**. For a 300-token response, that's 20 milliseconds of prefill versus 9 seconds of decode. You're paying for H100s at gold prices so they can twiddle their thumbs almost the entire time.

## The First Attempt: Disaggregating Prefill and Decode

The industry started correcting this with **disaggregation**. Instead of a single fleet of servers doing both, they're separated: specialized machines for prefill, others for decode. Each does its own job, no more interference. In production, this yields throughput gains of **2 to 7×**.

But this solution created another problem. Obviously, that would have been too simple.

## The Agentic Wall: 98.7% of Context Already Exists

A classic chatbot is simple: you ask a question, you get an answer, it's over. Agentic is very different. An autonomous agent runs in a loop, executes code, calls tools, reasons over dozens or even hundreds of iterations. With each turn, the context accumulates.

DeepSeek analyzed its own production traces. Result: an average agentic session is **157 iterations with a cumulative context of 32,000 tokens**. But with each turn, the agent only adds **409 new tokens on average**. In other words, **98.7% of the context already exists in previous turns**.

The KV cache is almost always there somewhere on disk. The prefill phase barely *computes* anything anymore: it spends its time fetching that cache from storage and loading it into memory. And there, wall.

In the classic architecture, it's always the prefill machine that has to fetch the data. That machine's storage network card is **saturated**. Meanwhile, the decode machines (which represent the majority of the cluster) have their own storage network card **completely inactive**. The pipes are empty, overall throughput collapses.

## Dual Path: Two Simultaneous Paths

The insight from the DeepSeek, Beijing, and Tsinghua researchers is ridiculously simple: *why is it that only the prefill machine is allowed to fetch data from storage?*

Dual Path is exactly that: **two simultaneous paths for loading the KV cache**.

- **Path A — classic.** The prefill machine fetches data from storage as usual, processes it all layer by layer, and sends the result to the decode machine via the inter-GPU network.
- **Path B — the detour.** In parallel, the decode machine uses its own storage network cards (until now inactive) to fetch the cache history itself directly from storage.

And there, clever: since the decode machine already has the bulk of the historical cache after prefill, the prefill machine **no longer needs to send it gigabytes**. It only sends the small incremental cache of new tokens. The decode merges it and off it goes.

## Avoiding Breaking Everything: Model Priority

You may have already spotted the risk: if Dual Path starts sending a lot of data over the network, it can compete with the communications the model needs to generate its response. Poorly managed, everything slows down.

The solution: **model traffic always takes priority**. It's like a highway with a lane reserved for emergencies. The model's important data goes first. Dual Path only uses the remaining bandwidth when the model is less busy.

The system monitors in real time which machines are least loaded. If the classic path is congested → switch to the other. If the other is slower → return to classic.

## The Results

In production:

- **Machine utilization goes from 40% to 80%**
- **Inference throughput doubles**
- **The first token arrives 56% faster**
- **Zero new GPUs purchased**

## The Broader Context: Optimizing Around GPUs

Dual Path doesn't come out of nowhere. It's one piece of a much broader movement in AI infra: *stop treating GPUs as the only problem, optimize everything around them*. Because in an AI data center, there's also memory, SSDs, network, the file system, cache. If even one of these pieces is a bottleneck, the whole machine slows down.

That's why systems like **HiCache** (organizes the KV cache across 3 tiers: GPU memory → server RAM → persistent storage), **Mooncake** (Best Paper at FAST 2025, exploits underutilized resources in a cluster), and **3FS** (DeepSeek's open-source distributed file system, designed for AI workloads) are appearing in parallel.

## My take: China is responding with software

There's obviously a geopolitical context behind all this. Since 2022, the United States has restricted China's access to the most advanced Nvidia GPUs (A100, H100). Chinese labs are under enormous pressure: **getting more performance with less raw hardware**.

Dual Path is exactly that: *a software response to a hardware constraint*. Instead of saying "we'll buy more GPUs," DeepSeek says "we'll make better use of what we already have." That's what makes the paper fascinating. **We're not talking about a smarter model — we're talking about smarter infrastructure.**

China has just proven that you can double the performance of an entire infrastructure with just code. It's a wake-up call that will have to be reckoned with.

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.
