---
id: collect-240926-nerdykings/nerdykings/nvidia-nemotron-two-tower-2-4x-plus-rapide-mais-a-quel-prix
title: "Nvidia Nemotron Two Tower: 2.4x Faster, But At What Cost?"
domain: nerdykings
role: reference
task: reference
actors: ["DeepSeek", "Nvidia"]
dates: []
keywords: ["cost", "nvidia", "agent", "agents", "attention", "compute", "deepseek", "gpus", "inference", "latency", "memory", "mixture of experts"]
source: docs/RAG/clean_en/nerdykings/nvidia-nemotron-two-tower-2-4x-plus-rapide-mais-a-quel-prix.md
source_anchor: ""
source_lines: [1, 47]
sha256: 11e7a9da09a0d04e65a3d986f3ef7eaa563d473ca9a6b10fc7f908e8f156c100
---

# Nvidia Nemotron Two Tower: 2.4x Faster, But At What Cost?

<!-- source: https://www.nerdykings.com/blog/nvidia-nemotron-two-tower-llm-plus-rapide.html -->

# Nvidia Nemotron Two Tower: 2.4x Faster, But At What Cost?

In early July, Nvidia released an architecture called **Nemotron Two Tower**, with a simple promise: making a language model **about 2.4 times faster**, without starting from scratch or sacrificing the bulk of its capabilities. Except that by digging into the numbers, we discover a bill that isn't being sold to you in the headline: to speed up a model, Nvidia must practically load two of them into memory. Here's what's hidden behind that 2.4x.

## The shift in philosophy behind this paper

For years, American labs have mainly sought to build **ever-bigger models**, trained on ever more data, with ever more GPUs. But for a few months now, **DeepSeek** has been pushing a different approach: getting more out of the same base, thanks to architecture, compression, and inference optimization — we already talked about it with Dual Path and DSpark. And this Nvidia paper fits exactly into that logic: instead of predicting only the next token, a model learns to anticipate several at once.

## How Two Tower works

Most current language models generate text **autoregressively**: the model produces a token, then uses that token to compute the next one. It must finish step 9 before starting step 10. This is very reliable, especially for code or mathematics, because each decision builds on everything that has already been validated. You can compare it to an extremely fast typewriter, but one that must always type the keys one by one.

For a classic chatbot, this isn't dramatic — the text is generated faster than you can read it anyway. But for **an agent** that calls a model dozens of times, uses a tool, reads the result, thinks, and starts again, a few seconds lost at each step quickly turns a simple task into an endless workflow.

That's where Nvidia comes in. **Nemotron Two Tower** is based on Nemotron 3 Nano 30B-A3B, a hybrid model that mixes Mamba, classic attention, and Mixture of Experts. Nvidia duplicates this base into two towers: the **context tower**, which reads the query and keeps a clean representation of everything that has already been generated, and the **denoising tower**, which learns to generate several tokens in parallel. The simplest image is that of a *writer* who knows the already-validated chapter perfectly, and a *proofreader* who receives the next paragraph in the form of 16 blank slots to fill in at the same time.

The base model had been pre-trained on about 25 billion tokens. Nvidia did not redo that colossal training: the company froze the first tower and trained only the second, on about 2.1 billion tokens. At the start of each block, the model creates 16 masked positions; the denoising tower tries to predict them simultaneously, looking in both directions within the block. The positions still uncertain are re-masked and recalculated in the next iteration, until all 16 positions are filled. The system therefore remains **autoregressive at the block scale, but parallel within each block**.

## 2.42x faster, but not on everything

In the configuration tested by Nvidia — blocks of 16 tokens on two H100 GPUs — Two Tower generates text **about 2.42 times faster** than the base autoregressive version, with **98.7% of overall quality preserved** according to Nvidia. On MMLU, the score barely moves: 78.56 for the classic version versus 78.24 for Two Tower. On the other hand, on **HumanEval** (Python code), the score goes from about 79.3 to 75.6, and mathematics also loses a few points.

This is fairly easy to explain: for tasks that require an extremely rigorous sequence like code or math, parallel generation lets more small errors slip through. In a program, a single parenthesis or the choice between a loop and a conditional immediately changes all the following tokens. Nvidia also allows switching back to a purely autoregressive mode when precision matters more than speed.

## The constraint that cannot be exceeded

The model is trained to work with blocks of 16 tokens, and that is a *hard* limit. With blocks of 64 tokens, HumanEval collapses to around 20% and GSM8K drops to almost zero. Parallelism works, but only within a very controlled window — it's not a button you can push at will to go even faster.

## The real bill: the hardware tax

And this is where the catchy headline hides the essential point. Each tower contains the weight of a model of about **30 billion parameters**. Even if only a few billion are active per token thanks to Mixture of Experts, both sets of weights must remain loaded in memory. At full precision, the system requires about **138 GB of VRAM**, on an official configuration with two H100s. It's quite ironic: the architecture reduces generation time, but it nearly doubles the amount of weights to store. And the 2.42x figure describes a precise, well-optimized configuration — nothing guarantees that an average user will get that result on their own hardware.

## For whom this really changes the game

For a classic chat, displaying 16 tokens per burst isn't necessarily more pleasant than smooth word-by-word streaming. The real potential lies elsewhere: in **agents**. An agent that has to make 30 calls to the model to complete a task sees each small latency gain repeated 30 times — and become very noticeable. With a speedup close to 2.4x, some workflows could go from several minutes to a much more acceptable time. On the other hand, for writing sensitive code, solving complex mathematics, or serving thousands of simultaneous requests in the cloud, the advantage is clearly less obvious.

## My opinion

Honestly, I don't think Two Tower is the universal future of language models. The hardware tax is really too high, and the drop in code performance shows that token-by-token generation is going to be very hard to replace, especially when absolute sequential precision is required.

But this paper remains interesting for a reason that goes far beyond Nvidia: it means that **American labs are finally adopting a different philosophy**. Stop always making things bigger, and start optimizing what we already have — exactly what DeepSeek has been repeating for six months. So the real question isn't whether Two Tower will replace autoregressive everywhere. It's how much longer the American giants will take to realize that the race for size has its limits, while others are already optimizing what exists.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.
