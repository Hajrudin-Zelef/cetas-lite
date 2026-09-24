---
id: collect-240926-nerdykings/nerdykings/diffusionblocks-sakana-ai-divise-la-memoire-d-entrainement-par-4
title: "DiffusionBlocks: Sakana AI Divides Training Memory By 4"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Nvidia", "OpenAI", "Sakana"]
dates: []
keywords: ["diffusion", "memory", "sakana", "training", "chatgpt", "claude", "deepseek", "gpu", "gpus", "llama", "nvlink", "parameters"]
source: docs/RAG/clean_en/nerdykings/diffusionblocks-sakana-ai-divise-la-memoire-d-entrainement-par-4.md
source_anchor: ""
source_lines: [1, 49]
sha256: 07c2096b65fc5055275c1acc8a334efdf8e432c392dcc51fa8792367b6546a0a
---

# DiffusionBlocks: Sakana AI Divides Training Memory By 4

<!-- source: https://www.nerdykings.com/blog/sakana-diffusionblocks-vram-entrainement-ia.html -->

# DiffusionBlocks: Sakana AI Divides Training Memory By 4

Training a powerful artificial intelligence today requires dozens, even hundreds of interconnected GPUs, and above all an enormous amount of memory. A team of researchers from the Japanese laboratory **Sakana AI** has just published a method called **DiffusionBlocks**, capable of reducing this required memory by 3, 4, or even more — without sacrificing performance. The basic idea is almost disconcertingly simple: split the model into several blocks and train each block independently. Let me explain how it works, what it produces in practice, and why this avenue could change the way we train tomorrow's AIs.

## Why training an AI consumes so much memory

When you use a model like ChatGPT, it mainly performs a **forward pass**: your text travels through the network's layers and the model produces a response. Training, on the other hand, is much heavier. The model first performs this forward pass, compares the result to the expected answer, then goes back in the opposite direction to determine which part of the network to correct — this is **backpropagation**. To be able to go back, the system must keep in memory the intermediate results of each layer (the activations), the model's weights, the gradients used to correct them, and the optimizer's states.

Concretely, for a model with 10 billion parameters, the weights alone already represent about 40 GB in FP32. Adding gradients, optimizer states, and activations, the required memory far exceeds 100 GB. And when several GPUs share the load, they must additionally constantly exchange their information, which adds yet another communication bottleneck.

## DiffusionBlocks: training the model block by block

It is this observation that pushed Sakana AI to ask a fairly radical question: are we really obliged to train all the model's layers at the same time? With DiffusionBlocks, the researchers split the network into several large blocks — for example, a 12-layer transformer can be divided into three blocks of four layers. Instead of loading all 12 layers and computing the gradients of the entire model at once, the system trains a single block at a time, then moves on to the next once that one is finished.

Since the memory used now depends mainly on the size of a single block and not the entire model, dividing the network into three blocks theoretically reduces memory by about three times, and with four blocks, the reduction can reach about four times. But for this to work, each block still needs to know precisely what it must learn — and this is where the diffusion principle that gives the method its name comes in.

## The link with image generation by diffusion

In an image generation model, one generally starts from pure noise, which the model progressively denoises: it first recovers the broad shapes, then builds the structure, then adds the fine details. The researchers at Sakana AI noticed that the passage through the layers of a transformer could be seen in the same way: each group of layers receives imperfect information, performs a small correction, then transmits a slightly cleaner state to the next group.

With DiffusionBlocks, each block therefore becomes a specialist in a precise stage of this denoising: the first learns to work with very noisy information, the second processes an already more structured version, the last focuses on the details and produces the final result. And above all, each block can learn its task without waiting for the results of the others — a bit like an assembly line where each team has a well-defined mission, without needing to keep the entire line in memory at all times.

## Results surprisingly close — sometimes better — than classical training

On paper, the idea may seem too simple to work. Yet the first results are solid. On an image classification test with three blocks, classical training obtained about 60.5% accuracy, versus 59.3% for DiffusionBlocks — with a theoretical memory footprint about three times smaller.

On image generation, the result is even more surprising: on ImageNet, the normally trained model obtained an FID score of 12.09 (the lower this score, the better the quality and diversity of the generated images), versus 10.63 for DiffusionBlocks. The method that uses less memory therefore slightly surpassed classical training. The researchers also tested a small language model inspired by Llama 2, separated into four blocks, with performance that remains close to classical training. The method is therefore not limited to image generation.

## The limit: too many blocks, and it degrades

One might think that it would then be enough to split a model into 20 or 30 blocks to save as much memory as possible. But the experiments show a clear limit: the more one increases the number of blocks, the fewer layers each block contains, and the more it loses its ability to understand the transformation it must perform. On image generation, two or three blocks generally give the best results, and from six blocks onward, quality degrades noticeably. There is therefore a real trade-off between memory saved and each block's ability to properly learn its task.

## What it would change if the method scales up

If DiffusionBlocks ever works on models with several billion parameters, the consequences could be significant. Researchers and companies with fewer resources could train or customize much larger models with more affordable GPUs. The method could also simplify distributed training: since the blocks train independently, they no longer need to constantly exchange their activations and gradients with the other GPUs. Each machine could process a different block with much less communication — and therefore potentially reduce dependence on ultra-fast connections like NVLink or InfiniBand, somewhat in the spirit of what we recently saw with Dual Path at DeepSeek.

We should remain cautious: the experiments presented here are still very far from the scale of models like GPT, Claude, or DeepSeek, and the researchers themselves acknowledge that scaling up to much larger models remains to be demonstrated. This method also mainly concerns training — to quickly run an autoregressive LLM, you still need to load all of its layers into memory. DiffusionBlocks therefore does not, as it stands, make it possible to run a giant model on a small graphics card.

## My opinion

Despite this limitation, DiffusionBlocks calls into question a rule that was almost taken for granted in modern AI: the idea that all the layers of a network must necessarily be trained together. Sakana AI shows that by transforming each group of layers into a specialist in a denoising step, one can train a network block by block while maintaining competitive performance. It's the same kind of bet found in recent attempts to rethink the transformer architecture: at some point, having more GPUs is no longer enough, you also have to question the way models are trained.

Let's be honest: for now, this is promising research, not yet an industrial revolution. But if it truly manages to scale to large models, it could make AI training much more accessible, and shake up part of the current GPU economy.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.
