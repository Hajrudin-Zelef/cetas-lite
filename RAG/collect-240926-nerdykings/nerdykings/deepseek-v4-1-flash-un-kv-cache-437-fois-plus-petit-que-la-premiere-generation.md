---
id: collect-240926-nerdykings/nerdykings/deepseek-v4-1-flash-un-kv-cache-437-fois-plus-petit-que-la-premiere-generation
title: "DeepSeek V4.1 Flash: A KV Cache 437 Times Smaller Than The First Generation"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["deepseek", "kv cache", "agents", "attention", "benchmark", "claude", "gpu", "memory", "mixture of experts", "parameters", "quantization"]
source: docs/RAG/clean_en/nerdykings/deepseek-v4-1-flash-un-kv-cache-437-fois-plus-petit-que-la-premiere-generation.md
source_anchor: ""
source_lines: [1, 43]
sha256: 4af36b744a989300a8998d26a112c5cc3d4e43756bc90a94d04802b382f63803
---

# DeepSeek V4.1 Flash: A KV Cache 437 Times Smaller Than The First Generation

<!-- source: https://www.nerdykings.com/blog/deepseek-v4-1-flash-kv-cache-437x.html -->

# DeepSeek V4.1 Flash: A KV Cache 437 Times Smaller Than The First Generation

DeepSeek has just released **V4.1 Flash**, and the model is quite surprising: first because it's extremely fast, but above all because on certain tests, it manages to compete with high-end models. Except that's not the real story. What DeepSeek managed to do is drastically reduce the amount of memory needed to handle very long contexts — their **KV cache** is about **4 times smaller** than that of V4 Flash, and up to **437 times smaller per token** than that of their very first generation, all while maintaining a context that can reach 1 million tokens. I'll explain how they did it, and why this could well show where the future of LLMs is heading.

## A fast model, but one that doesn't win everywhere

Let's start by calming things down. On DeepSWE, V4.1 Flash reaches **74.2 points** versus **74** for Claude in the reported results: that's impressive for a model in this category. But be careful, it doesn't win everywhere. On Terminal-Bench, DeepSeek drops to **31.2** versus **51.8** for Claude. So no, DeepSeek hasn't suddenly created the best model in the world. On the other hand, they've done something much more interesting.

## The real problem with big models: the KV cache

When you chat with an AI, the model must constantly keep information about what happened previously in memory. That's the **KV cache**. With a conversation of a few messages, no problem. But today, we're creating agents that are given entire projects: hundreds of files, a terminal's history, tools, images. V4.1 Flash can work with contexts of up to **1 million tokens**, and the more that context grows, the more the memory needed for the KV cache becomes a real problem. That's exactly where DeepSeek made a huge leap: the KV cache of V4.1 Flash is about **4 times smaller** than that of V4 Flash — and about **437 times smaller per token** than that of DeepSeek's first generation, in just a few years.

## CSA2: only 4 layers out of 40 write the global memory

The big innovation behind this is called **CSA2**, for *Compressed Sparse Attention*. A Transformer is made up of many layers, and traditionally, each of them produces and retains its own information in the KV cache. DeepSeek asked itself: do we really need all the layers to independently store all that information? With CSA2, the answer is no. On V4.1 Flash, only **four layers** — layers 2, 8, 14 and 20 — write a persistent global KV state, and all the others reuse the information already produced instead of keeping their own version.

DeepSeek combines this with a rather unusual structure: V4.1 Flash uses **40 layers split into two parts**, 20 encoder layers and 20 decoder layers. When you give your huge prompt to the model, the encoder first processes it and produces, in a way, a global memory of that context. When the decoder starts generating the response, it can read that memory directly instead of recreating new global keys and values at each layer. Add a few other compression and quantization techniques, and the result is huge.

## The result: 3,514 bytes per token → 890

On V4 Flash, we were at about **3,514 bytes per token**. With V4.1 Flash, only **890 bytes**: almost four times less. And that matters enormously, because GPU memory is one of the most precious resources when running these enormous models — it's the same logic of efficiency as Dual Path on the GPU side, or Engram on the memory side: DeepSeek has been tackling hardware bottlenecks for months rather than raw size.

## Careful: this is not a small model

One thing still needs to be clarified: V4.1 Flash is absolutely not a small model. Its backbone contains about **552 billion parameters**. Fortunately, it's a *Mixture of Experts*: they're not all used at the same time, and only about **8 billion parameters** are active per token during prompt processing, then about **16 billion** during generation. But if you want to download the full checkpoint, we're still talking about roughly **511 GB**.

Another very important point: V4.1 now has **native visual understanding**. You can send it text, but also images — for example a user interface, a screenshot or even a video game menu — then ask it to analyze what it sees and generate the code needed to reproduce the behavior. This obviously becomes very interesting for agents, which become much more versatile than a simple chatbot.

## The flaw: it thinks way too much

Our little model seems almost perfect, but there's a problem: it thinks way too much. The maximum effort mode can consume up to **2.5 times more output tokens**. And the funny part is that it isn't necessarily useful: on a writing test, the standard mode gets **87.13 out of 100**, the maximum mode about **86.75** — practically the same, or even slightly less. Except that generation takes **73% more time** and costs about **60% more**. Sometimes, thinking more is absolutely useless.

## My opinion

V4.1 Flash is a fairly interesting, even surprising model, because you can clearly see that DeepSeek is working on a very specific problem: how to run gigantic models with gigantic context, without blowing up the required memory. And in a few years, they've managed to get a KV cache per token up to 437 times smaller than their first-generation model. Let's be honest: today, the model remains substantial — 552 billion parameters and 511 GB is not something you're going to run on your PC.

But it's the trajectory that matters. If the trend continues, these models will become less and less expensive to run, and that's where the real race is being played: not the one for the best score on a benchmark, but the one for who will be able to make agents work on millions of context tokens without the GPU memory bill exploding. On that front, DeepSeek clearly has a lead worth watching closely.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.
