---
id: collect-240926-nerdykings/nerdykings/dspark-deepseek-accelere-son-modele-de-78-sans-le-toucher
title: "DSpark: DeepSeek Accelerates Its Model by 78% Without Touching It"
domain: nerdykings
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "OpenAI"]
dates: []
keywords: ["deepseek", "agent", "agentic", "benchmark", "chatgpt", "claude", "gpu", "inference", "kv cache", "license", "llama", "memory"]
source: docs/RAG/clean_en/nerdykings/dspark-deepseek-accelere-son-modele-de-78-sans-le-toucher.md
source_anchor: ""
source_lines: [1, 92]
sha256: 4da0fe5c78c4d38dbd218f552f4e36917f47d924574d8f32d9883766e735644e
---

# DSpark: DeepSeek Accelerates Its Model by 78% Without Touching It

<!-- source: https://www.nerdykings.com/blog/dspark-deepseek-decodage-speculatif.html -->

# DSpark: DeepSeek Accelerates Its Model by 78% Without Touching It

DeepSeek has done it again. **With DSpark, their V4 Pro model generates up to 78% faster** — no new model, no retraining, without them having touched a single line of the weights. And it's the 5th time in 6 months they've pulled the *"we optimize instead of scaling up"* trick on us. Here's how it works and why it's seriously worrying the competition.

## Why an LLM is slow (by default)

Before understanding DSpark, you need to really grasp why an LLM is slow by default. When you send a question to ChatGPT or Claude, the model doesn't generate the answer all at once. It builds it **one word at a time**. Concretely: it generates the first word, looks at what it just produced, generates the second word, looks again, generates the third, etc. This is called **autoregressive generation**.

The thing is, to generate each new word, the model has to *reread its entire internal memory* — the famous KV cache. So the longer the answer, the more each additional word costs. And meanwhile, the GPU is working at **20-30% of its capacity**. You're paying for H100s at $30,000 a pop so they can wait on memory.

## Speculative decoding: betting on 8 words at once

Now, imagine if instead of generating one word, checking, generating the next, rechecking... we could **bet on the next 8 words at once** and just verify them in parallel. That's exactly the idea behind speculative decoding.

The principle is almost too simple to be true. You take two models: a big ultra-powerful one (your *target* model) and a small much faster one (your *draft* model). The small one proposes a burst of 8 words, the big one verifies all 8 proposals in **a single parallel pass**. If it validates: you've generated 8 words for the price of one. If it rejects starting from word 5: you keep the first 4 and start over. It's exactly like when your smartphone keyboard suggests the rest of your sentence — you accept it as a block if it's correct, otherwise you correct it.

## The historical flaw of both approaches

The problem is that all existing systems had a structural flaw:

- **Autoregressive drafts (Eagle 3).** Precise, the large model often validates. But generating the draft takes time → you gain on one side, you lose on the other.
- **Parallel drafts (10 Flash).** They spit out the 8 words ultra fast. But since each word is generated without looking at its neighbors, the draft becomes incoherent. **The acceptance rate collapses** on the last words of the block.

Nobody could get both. That's exactly what DSpark has just unlocked.

## The DSpark trick: the Markov head

DSpark's main idea is to take the best of both worlds with an architecture they call **semi-autoregressive**. The small draft model guesses several words at once, in parallel. Except that to avoid incoherence, DeepSeek adds its clever idea: **the Markov head**.

It's a mini layer that runs just before the final word selection. For each word, it looks *only* at the word right before it. Not the whole sentence. Not the whole context. Just the left neighbor. And that's enough to correct plenty of local errors. If the model just chose "faster", the Markov head increases the chances that the next word is "than" or something coherent, and lowers the chances of words that make no sense there.

Really clever technical detail: this Markov head is built with a **rank-256 factorization**. Instead of doing a huge massive computation over the entire vocabulary, they compress the operation so it costs almost nothing. Result: the draft becomes much cleaner, stays fast, and the large model accepts it much more often. They didn't make the model bigger, *they just made the draft less dumb*.

## Confidence-planned verification

But DSpark doesn't stop there. Second problem solved: in production, if for each request you send a block of 16 words to verify and the large model systematically rejects the last 10, you've wasted 10 words of computation for nothing. At DeepSeek's scale, that's dozens of H100s spinning in a vacuum permanently.

The solution is called **confidence-planned verification**. Before asking the large model to verify the block, we evaluate word by word the probability that each word survives verification. They use two tools in sequence:

- **A confidence head** that outputs a score for each word of the draft.
- **An algorithm called Sequential Temperature Scaling** that recalibrates the scores (because neural networks tend to be overconfident). The calibration error goes from 3-8% to about 1%.

The system monitors GPU load in real time: when they're free, it verifies large blocks. When the system is saturated, it shortens the block and only sends the words most likely to be accepted. The power is concentrated *only where it's worth it*.

## The production results

On DeepSeek's servers:

- **+57 to +78% generation speed** per user on DeepSeek V4 Pro
- **+60 to +85% speed** on DeepSeek V4 Flash
- **Zero degradation in response quality**
- On Qwen 3 4B, DSpark accepts **26-31% more words** than Eagle 3 and 16-18% more than 10 Flash

And the craziest thing in all of this: DeepSeek has **open-sourced everything under the MIT license**. The project is called DeepSpec. Full code, data preparation, training, evaluation, everything is on GitHub. Anyone can deploy DSpark on their own Qwen, Gemma or Llama model without paying a single cent.

## 5 releases in 6 months: the DeepSeek pattern

Let's take two minutes and look at the patterns. In 6 months, DeepSeek has given us:

- **January — MHC.** A new way to stabilize the model's internal flows.
- **February — Engram.** The conditional memory that offloads to RAM.
- **May — DeepSeek V4.** A million tokens at a cut-rate price.
- **June — Dual Path.** GPU optimization without new hardware.
- **Late June — DSpark.** Speculative decoding in production.

Each time, the same spirit: *you don't make the model bigger, you just make it smarter*. It's the Chinese philosophy in the face of American GPU sanctions. When you don't have access to the latest generation of H100s, you find other solutions. You learn to make do with what you have. And along the way, you demolish the commercial margins of Western labs that have to charge a lot because they spend hundreds of millions on hardware.

Concretely: DeepSeek V4 Pro is now **7× cheaper than GPT 5.5** and **6× cheaper than Claude Opus 4.7** for producing almost the same result. With DSpark on top, it's another 80% faster.

## The blind spots you need to know about

Let's be honest: DeepSeek V4 Pro still has **two major blind spots**.

**1. Long agentic tasks.** On Terminal Bench 2.0 — a benchmark where the agent has to compile code, configure servers, and train models over several hours — DeepSeek V4 Pro scores 67.9%. GPT 5.5 scores 82.7%. **A 15-point gap is massive**. For true dev autonomy, the proprietary ecosystem remains ahead.

**2. Factual hallucinations.** On pure general-knowledge questions, V4 Pro has a hallucination rate of **94%** when asked about a fact it doesn't know. Almost every time, it prefers to make something up rather than admit its ignorance. For legal, medical, or critical factual work: it's not the right choice.

DSpark accelerates a very powerful model, but *it doesn't fix its flaws*. You need to know what you're getting into.

## My take: a pace no one else can keep

DeepSeek is playing a completely different game from the other players. They're proving that you can **divide inference costs by 7, nearly double the speed, and give all of it away for free to the open-source community**.

The real signal from DSpark isn't just the performance, it's the *timing*. They released DeepSeek V4 in April and only 2 months later, they're delivering an infrastructure optimization that multiplies everything. **It's an iteration pace that no one else can keep.**

We're seeing a split take shape in the AI market. On one side, proprietary models that justify themselves on high-end niches (extreme agentic autonomy, zero medical hallucination, multimodality). On the other, **the DeepSeek ecosystem becoming the new center of gravity for absolutely everything else**.

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.
