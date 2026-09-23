---
id: vague2-nerdykings/nerdykings/longcat-2-0-chine-sans-nvidia
title: "LongCat 2.0 : Le Modèle IA Chinois Entraîné Sans Un Seul GPU Nvidia"
domain: nerdykings
role: reference
task: article
actors: ["Alibaba", "China", "DeepSeek", "Huawei", "LongCat", "Meituan", "Moonshot", "Nvidia", "United States", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["gpu", "nvidia", "agent", "agents", "ascend", "attention", "benchmarks", "context window", "cost", "deepseek", "glm", "gpus"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/longcat-2-0-chine-sans-nvidia.md
source_anchor: ""
source_lines: [1, 69]
sha256: 4d24645c4891a8cb777e5d770acb9319b808be08454b5cbc89f866aaf1c6b807
---

# LongCat 2.0 : Le Modèle IA Chinois Entraîné Sans Un Seul GPU Nvidia

## Metadata

- **Source** : https://www.nerdykings.com/blog/longcat-2-0-chine-sans-nvidia.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Meituan, China's food-delivery champion (a mix of Uber Eats and DoorDash), has an AI lab called LongCat AI that published **LongCat 2.0**, an open-source 1.6-trillion-parameter model trained on over 50,000 Chinese AI accelerators — without a single Nvidia GPU. In the context of the US-China tech war, this is no small detail.

Meituan, the delivery company taking on Nvidia: when thinking of China's AI race, one thinks of DeepSeek, Alibaba's Qwen, Moonshot's Kimi, or Zhipu's GLM. LongCat 2.0 comes from none of them but from Meituan, a giant group in food delivery and local services, which quietly runs the LongCat AI lab.

Why replacing Nvidia is such a challenge: for years the US has limited China's access to the most advanced AI training chips. The problem isn't just raw chip power — Nvidia has an enormous advantage called **CUDA**, the entire software ecosystem built around its GPUs for nearly 20 years (libraries, tools, optimizations). Almost all modern AI was built around this environment. Replacing an Nvidia GPU with a Chinese chip isn't like swapping a graphics card in a PC: huge parts of the software infrastructure sometimes must be rewritten. Some reports estimate migrating large workloads to accelerators like Huawei Ascend can add considerable development time and cost.

That's why LongCat 2.0 draws attention: Meituan didn't run a small experiment on a few hundred chips. To train the model, they used an infrastructure of **more than 50,000 domestic AI accelerators**, with **more than 35 trillion tokens** used during pre-training — and not a single Nvidia GPU. Running a model on a Chinese chip is one thing; making tens of thousands of chips work together in a training run of this size is another.

At this scale, anything can break: chips can fail, or worse, start producing incorrect results. There are communication, memory, and synchronization problems, and the larger the cluster, the higher the probability of an issue. Meituan claims to have completed training without major problems, notably thanks to systems detecting **silent data corruption** — a chip producing incorrect calculations without announcing it. If unnoticed, such errors can progressively contaminate the whole training. Their infrastructure is reportedly able to detect and isolate these problems before too much damage.

A model tailored to its own hardware: Meituan didn't just adapt infrastructure — they adapted the model itself rather than replicating what works on Nvidia GPUs. LongCat 2.0 has 1.6 trillion parameters; activating all of them per generated token would be absurd in cost. So the model uses **Mixture of Experts (MoE)**: only about **48 billion parameters are active on average** per token. They went further: the model can decide that very simple tokens (a space, punctuation, ultra-predictable code structure) deserve almost no computation, unlike a complex math problem. Same logic on attention: LongCat 2.0 accepts up to **1 million tokens** of context, but constantly comparing each token to all of that million would explode cost. Hence **Long Cat Sparse Attention**: the model determines which context portions truly matter, and this search is designed to match how hardware accesses memory — working with grouped blocks rather than scattered fragments. The core idea: build a model that runs well on its own infrastructure rather than copying what works on Nvidia.

Is LongCat 2.0 actually good? Results are promising: it's particularly oriented toward coding and AI agents. On **SWE-bench Pro** it scores **59.5%**, and on **SWE-bench multilingual** it climbs to **77.3%**. Not the world's best model, but far from a mere technical demo proving Chinese chips can run something — it's genuinely competitive. And importantly, LongCat 2.0 is open source under the **MIT license** — weights can be freely downloaded, modified, and used.

Have US sanctions failed? This raises the question. For years, much of US strategy has been to slow China's AI development by limiting access to the best chips — and these restrictions genuinely created difficulties. The Nvidia ecosystem remains extremely mature, CUDA remains a huge advantage, and Chinese accelerators haven't suddenly become better than the best Nvidia chips. But there's an interesting side effect: the harder access to US technology becomes, the more Chinese companies have an incentive to massively invest in their own alternative — hardware, libraries, distributed systems, and model architectures themselves. LongCat 2.0 shows a new stage may be reached. The real question: how fast will this alternative ecosystem catch up on 20 years built around Nvidia?

The author's view: LongCat 2.0 isn't the world's best model, and Nvidia + CUDA remain a fortress. But strikingly, it isn't a known AI lab that produced this result — it's a food-delivery company. If Meituan can train a competitive 1.6-trillion-parameter coding model entirely without Nvidia, the recipe is becoming reproducible in China, not reserved for a handful of giants like Alibaba. The real issue, beyond benchmarks: if other Chinese labs reproduce this kind of large-scale training, US restrictions may have succeeded in slowing China short-term while paradoxically accelerating the construction of a fully independent Chinese ecosystem long-term — with consequences far more important than LongCat 2.0 itself.

## Key points

- **LongCat 2.0** is an open-source (MIT) 1.6-trillion-parameter model from Meituan's LongCat AI lab.
- Trained on **50,000+ Chinese AI accelerators** with **zero Nvidia GPUs** and **35+ trillion tokens**.
- Nvidia's real moat is **CUDA** — a ~20-year software ecosystem; switching hardware often requires rewriting infrastructure.
- Uses **MoE**: ~**48 billion active parameters** per token out of 1.6 trillion total.
- Simple tokens can trigger near-zero computation; **Long Cat Sparse Attention** handles up to **1M-token** context via hardware-friendly block access.
- Silent data corruption detection was key to completing the large-scale training run.
- Competitive coding/agent scores: SWE-bench Pro **59.5%**, SWE-bench multilingual **77.3%**.
- Raises the question of whether US chip sanctions are backfiring by accelerating an independent Chinese AI ecosystem.

## Technical data / figures

| Metric | Value |
|---|---|
| Model | LongCat 2.0 |
| Developer | Meituan (LongCat AI) |
| Total parameters | 1.6 trillion |
| Active parameters/token | ~48 billion (MoE) |
| Training accelerators | 50,000+ Chinese (no Nvidia) |
| Pre-training tokens | 35+ trillion |
| Context window | 1 million tokens |
| SWE-bench Pro | 59.5% |
| SWE-bench multilingual | 77.3% |
| License | MIT (open source) |
| Key technique | Long Cat Sparse Attention, MoE |
| Hardware challenge | CUDA replacement, silent data corruption |

- Key concepts: **Mixture of Experts**, **sparse attention**, **silent data corruption**, **CUDA**, **Huawei Ascend**
- Strategic context: US export restrictions on advanced AI chips to China

## Why this source matters for the RAG

This article documents a major milestone in hardware-independent AI training, with concrete scale metrics and architectural adaptations to non-Nvidia accelerators. It is valuable for a RAG knowledge base on AI geopolitics, hardware sovereignty, MoE/sparse-attention architectures, and the CUDA ecosystem.

## Source URL

https://www.nerdykings.com/blog/longcat-2-0-chine-sans-nvidia.html
