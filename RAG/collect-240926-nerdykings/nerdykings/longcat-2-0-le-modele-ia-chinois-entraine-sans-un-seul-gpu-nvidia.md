---
id: collect-240926-nerdykings/nerdykings/longcat-2-0-le-modele-ia-chinois-entraine-sans-un-seul-gpu-nvidia
title: "LongCat 2.0: The Chinese AI Model Trained Without a Single Nvidia GPU"
domain: nerdykings
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Huawei", "LongCat", "Meituan", "Moonshot", "Nvidia", "United States", "Z.ai"]
dates: []
keywords: ["gpu", "nvidia", "agents", "ascend", "attention", "benchmarks", "cost", "deepseek", "glm", "gpus", "kimi", "license"]
source: docs/RAG/clean_en/nerdykings/longcat-2-0-le-modele-ia-chinois-entraine-sans-un-seul-gpu-nvidia.md
source_anchor: ""
source_lines: [1, 53]
sha256: 8b384761f1d590c15316b73fdb061aef7572464f93d5959caf289472e0b0516a
---

# LongCat 2.0: The Chinese AI Model Trained Without a Single Nvidia GPU

<!-- source: https://www.nerdykings.com/blog/longcat-2-0-chine-sans-nvidia.html -->

# LongCat 2.0: The Chinese AI Model Trained Without a Single Nvidia GPU

Meituan. If this name means nothing to you, that's normal: it's the Chinese champion of food delivery, a sort of mix between Uber Eats and DoorDash. Except that this week, their AI lab released **LongCat 2.0**, an open source model with **1.6 trillion parameters** trained on more than **50,000 Chinese AI accelerators — without a single Nvidia GPU**. In the context of the technological war between China and the United States, that's no small detail.

## Meituan, the Delivery Company Taking on Nvidia

When we think of the AI race in China, we immediately think of DeepSeek, of Alibaba with Qwen, or of Moonshot with Kimi and Zhipu with GLM. But LongCat 2.0 doesn't come from any of these players. It comes from Meituan — a **gigantic Chinese group specializing in food delivery and local services**, imagine a mix between Uber Eats, DoorDash, and a dozen other services grouped into a single app. Except that they also, discreetly, have an artificial intelligence laboratory: LongCat AI. And this lab has just done something quite impressive.

## Why Replacing Nvidia Is Such a Challenge

Before going further, we need to understand why it's so difficult to do without Nvidia. For several years, the United States has been restricting China's access to the most advanced chips used to train AI models. The problem isn't just a question of raw chip power. Nvidia has a huge advantage called **CUDA**: an entire software ecosystem built around its GPUs for nearly 20 years — libraries, tools, optimizations. Practically the entire modern AI industry has been built around this environment.

Replacing an Nvidia GPU with a Chinese chip is therefore not like swapping a graphics card in a PC: sometimes you have to rewrite a huge part of the software infrastructure. Some reports estimate that migrating large workloads to accelerators like **Huawei Ascend** can add considerable development time and cost.

## 50,000 Chinese Chips, Zero Nvidia Cards

And that's precisely why LongCat 2.0 is generating so much buzz. Meituan isn't talking here about a small experiment run on a few hundred chips. To train the model, they used an infrastructure of more than **50,000 domestic AI accelerators**, with more than **35 trillion tokens** used during pre-training — and that, *without a single Nvidia GPU*.

Getting a Chinese chip to run a model is one thing. Getting tens of thousands of chips to work together during training of this size is a whole other story.

## At This Scale, Anything Can Break

At this scale, problems change in nature. One chip can fail, another can start producing incorrect results. There are also issues with communication between machines, memory, synchronization — and the bigger the cluster grows, the more the probability that a problem arises somewhere increases.

Meituan claims to have completed training without major issues, notably thanks to detection systems for what is called **silent data corruption**: a chip can start producing incorrect calculations without ever clearly announcing it. If no one notices, these errors can progressively contaminate the entire training run. Their infrastructure is reportedly capable of detecting and isolating this kind of problem before it does too much damage.

## A Model Tailor-Made for Its Own Hardware

But Meituan didn't just adapt its infrastructure: they also adapted the model itself, rather than trying to reproduce exactly what works on Nvidia GPUs. LongCat 2.0 has 1.6 trillion parameters — obviously, activating all of them for every generated token would be absurd in cost. The model therefore uses the **Mixture of Experts (MoE)** architecture: imagine an enormous number of small specialists, and depending on what it needs to process, the model only wakes up a small portion of them. Result: out of the 1.6 trillion parameters, **only about 48 billion are active on average** to process a token.

They even went further: the model can decide that certain very simple tokens — a space, punctuation, an ultra-predictable code structure — deserve almost no computation, unlike a complex mathematical problem. Same logic on the attention side: LongCat 2.0 accepts up to **1 million tokens of context**, but constantly comparing each token to the entirety of that million would make the cost explode. Hence the **Long Cat Sparse Attention**: the model determines which portions of the context are truly important, and this search is designed to match the way the hardware accesses memory — working with grouped blocks rather than going to fetch scattered crumbs everywhere. A detail that may seem technical, but the underlying idea is simple: **build a model that runs well on its own infrastructure**, rather than copying what works on Nvidia.

## Is LongCat 2.0 Really Any Good?

All of this would still be a lot less interesting if the model were completely useless. The results are rather promising: LongCat 2.0 is particularly oriented toward coding and AI agents. On **SWE-bench Pro**, it achieves 59.5%, and on **SWE-bench multilingual**, it climbs to 77.3%. We're not looking at the best model in the world, but we're far from a simple technical demonstration intended to prove that Chinese chips can run something: this is a genuinely competitive model. And an important detail: LongCat 2.0 is **open source, published under the MIT license** — the weights can be downloaded, modified, and used freely.

## Have American sanctions failed?

This inevitably raises the question: have American sanctions missed their mark? This is probably what makes this release more important than it appears. For several years now, a good part of American strategy has consisted of slowing down China's AI development by limiting its access to the best chips — and let's be honest, these restrictions have genuinely created difficulties. The Nvidia ecosystem remains extremely mature, CUDA remains a gigantic advantage, and Chinese accelerators haven't suddenly become better than the best Nvidia chips.

But there's an interesting side effect: the harder access to American technologies becomes, the more incentive Chinese companies have to invest massively in their own alternative — the hardware, the libraries, the distributed systems, but also the model architectures themselves. And LongCat 2.0 shows that we may be reaching a new stage. The real question becomes: how fast will this alternative ecosystem catch up to twenty years of construction around Nvidia?

## My take

Let's be honest: LongCat 2.0 isn't the best model in the world, and Nvidia + CUDA remain a fortress. But what strikes me is that it's not even a well-known AI lab putting out this result — it's **a food delivery company**. If Meituan is capable of training a 1.6-trillion-parameter model, competitive on coding, entirely without Nvidia, that means the recipe is becoming reproducible in China, not just reserved for a handful of giants like Alibaba.

And that's the real story, beyond the benchmarks: if other Chinese labs reproduce this kind of large-scale training, American restrictions may have succeeded in slowing China down in the short term… while paradoxically accelerating the construction of a fully independent Chinese ecosystem in the long term. And that could have far more important consequences than LongCat 2.0 itself. **Worth watching very, very closely.**

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.
