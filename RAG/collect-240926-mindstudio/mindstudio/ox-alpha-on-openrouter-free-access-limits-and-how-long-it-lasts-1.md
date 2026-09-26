---
id: collect-240926-mindstudio/mindstudio/ox-alpha-on-openrouter-free-access-limits-and-how-long-it-lasts-1
title: "ox-alpha-on-openrouter-free-access-limits-and-how-long-it-lasts"
domain: mindstudio
role: reference
task: reference
actors: ["China", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["agentic", "agents", "attention", "benchmark", "benchmarks", "compute", "context window", "cost", "glm", "gpu", "inference", "multimodal"]
source: docs/RAG/clean_en/mindstudio/ox-alpha-on-openrouter-free-access-limits-and-how-long-it-lasts.md
source_anchor: ""
source_lines: [1, 56]
sha256: 69f321efa05eb41fbf9b13d1fc5876a76aa2aea9b8507d2f4b837fc1bfefd2a9
---

# ox-alpha-on-openrouter-free-access-limits-and-how-long-it-lasts

<!-- source: https://www.mindstudio.ai/blog/ox-alpha-free-access-openrouter -->

## What is OX Alpha, and why is it free on OpenRouter?

OX Alpha is a stealth large language model that showed up on OpenRouter offering a 1 million token context window, multimodal input (text, image, and video), zero data retention, and a claimed daily capacity in the trillions of tokens, all at no cost. No company has officially attached its name to the model. OpenRouter’s own listing describes it only as a frontier model built for efficient coding, sustained agentic work, and real-world production use. That vagueness, combined with genuinely unusual usage limits, is exactly why it has generated so much attention among developers looking for a free way to run high-context, agentic coding workloads.

## TL;DR

- **OX Alpha is a stealth model** on OpenRouter with no confirmed publisher, offering a 1 million token context window and multimodal input across text, image, and video.
- **It’s currently free to use** , with rate limits described as near unlimited, which is unusual for a model performing in the same range as established frontier systems.
- **Most signals point to it being connected to the GLM family** , based on tokenizer analysis from jailbreaking researcher Pliny the Liberator and separate leaks referencing a “GLM 5.3 flash” checkpoint.
- **Benchmark results are inconsistent** , with some independent tests showing performance near GPT 5.5-5.6 territory and others showing it clearly underperforming once retested on contamination-resistant private benchmarks.
- **A prediction market reportedly assigned a 90% probability** that OX Alpha resolves to Zed.ai (the company behind GLM), though prediction markets aren’t guaranteed to be correct.
- **Free stealth models on OpenRouter historically don’t stay free or anonymous for long** , based on precedent with an earlier anonymous release later confirmed as GLM-5, so testing OX Alpha sooner rather than later is the practical move for anyone curious.
- **The “why is this free” question is unresolved** , with theories ranging from a marketing strategy to gather unbiased feedback, to speculation about a new compute partner, to unverified claims about continual learning breakthroughs.

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

## How much can you actually use for free?

The headline number circulating is a daily capacity figure in the range of 100 trillion tokens, which would make individual rate limits on OpenRouter extremely generous, described by early testers as “near unlimited.” That doesn’t mean there are zero constraints. OpenRouter typically imposes some form of per-user throttling on free and stealth models to prevent abuse, and free listings like this one are usually temporary by design. Historically, labs testing models anonymously on OpenRouter treat the free window as a limited-time public benchmark, not a permanent offering. Anyone building against OX Alpha right now should treat the access as provisional and avoid depending on it for anything in production.

## Why would a lab give away this much compute for free?

This is the part of the story that doesn’t have a clean answer. Running a 1 million context, multimodal model at near unlimited scale for free requires enormous inference compute, and most AI labs are currently constrained by GPU availability, not flush with spare capacity. A few explanations have circulated:

**It’s a blind evaluation strategy.** There’s direct precedent for this. GLM’s maker has publicly described a prior anonymous release, nicknamed “Pony Alpha,” as a deliberate move to strip away brand bias and let the model’s raw output quality drive the feedback. According to their own comments, the strategy worked: the earlier stealth model built a reputation purely on performance before being revealed as GLM-5, which the company said helped validate that Chinese LLMs could compete at the frontier level. If OX Alpha follows the same pattern, giving it away free and unbranded is a way to generate organic, unbiased benchmarking buzz before a formal launch.

**There’s new compute backing it.** Some observers have speculated about an undisclosed investment or new hardware partner enabling this level of free inference, since no major lab is publicly known to have surplus compute at this scale right now.

**Continual learning speculation.** A separate and unverified thread of discussion suggests OX Alpha might be tied to continual learning research, the idea that a model keeps updating from live usage rather than staying static after training. Safe Superintelligence, the lab founded by Ilya Sutskever, has been reported to be working on continual learning and was rumored to have upcoming breakthroughs. Some users claimed OX Alpha showed signs of being updated mid-deployment, which fed the theory. This remains speculation without confirmation, and continual learning in deployed LLMs is not something that has been demonstrated at this scale by any lab publicly.

## Is OX Alpha actually as good as GPT 5.5 or 5.6?

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

The benchmark picture is mixed, and that inconsistency is itself a useful data point. One early test run on the DeepSWE coding benchmark reportedly showed OX Alpha scoring around 80%, well ahead of comparison models. But when other testers ran their own private, contamination-resistant benchmarks, the model underperformed significantly, landing closer to parity with GPT 5.5-5.6 class models rather than beating them outright. Separate independent benchmarks placed it just below GPT 5.6 Sol-mid and just above GPT 5.5.

The practical takeaway: OX Alpha appears to be a strong, competitive model in the same tier as current frontier systems, not something that clearly leapfrogs them. Claims of it dramatically outperforming GPT-5-class models have not held up consistently across retests. Given that it hasn’t been tested widely through a stable API by independent researchers, benchmark numbers floating around should be treated as provisional.

## What model is OX Alpha actually built on?

The most consistent theory is that OX Alpha is connected to the GLM family, specifically referred to in some leaks as “GLM 5.3 flash.” Two pieces of evidence support this:

Pliny the Liberator, a researcher known for jailbreaking LLMs to extract system prompts and internal configuration, reportedly found tokenizer evidence linking OX Alpha to the GLM family. Jailbreaking a model can reveal its system prompt and behavioral instructions, which sometimes exposes clues about the underlying architecture or origin even when a model is deployed under a code name.

Separately, a prediction market reportedly assigned around a 90% probability that OX Alpha resolves to Zed.ai, the company behind GLM. Prediction markets aggregate the bets of people with inside knowledge or strong conviction, so a lopsided probability like that is a meaningful signal, though not proof. Other theories, including a link to Safe Superintelligence or an Nvidia collaboration, have circulated but lack comparable supporting evidence.

If the GLM connection holds, there’s a practical implication for anyone interested in local inference: a model performing near GPT 5.6 Sol-mid level that can eventually be run locally on capable hardware, with electricity as the main ongoing cost, would be a meaningful development for the open-weight ecosystem.

## Should you try OX Alpha now?

