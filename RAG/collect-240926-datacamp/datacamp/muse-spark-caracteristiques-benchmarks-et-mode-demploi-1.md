---
id: collect-240926-datacamp/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi-1
title: "muse-spark-caracteristiques-benchmarks-et-mode-demploi"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "China", "DeepSeek", "Google", "Meta", "OpenAI"]
dates: ["2025-04", "2025-06-30", "2025-11", "2026-04-08"]
keywords: ["benchmark", "benchmarks", "muse", "agent", "agents", "chatgpt", "claude", "compute", "deepseek", "gemini", "inference", "latency"]
source: docs/RAG/clean_en/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi.md
source_anchor: ""
source_lines: [1, 90]
sha256: a59435beca1f1ba95e89b07b51bace8424e995dc162fa81d817cfd1bfd72b676
---

# muse-spark-caracteristiques-benchmarks-et-mode-demploi

<!-- source: https://www.datacamp.com/fr/blog/muse-spark -->

Cursus

We were publishing articles on Meta's Llama models (Llama 2, Llama 3, etc.) at a steady pace. Then Llama 4 arrived in April 2025 amid a firestorm of criticism: several media outlets and the company's outgoing AI director confirmed that benchmark results had been manipulated using specialized sub-models that were never released.

After that, the updates stopped. At the same time, Meta announced that Horizon Worlds would move to mobile only, effectively ending the VR version on which it had once staked the company's future. Taken together, it painted the picture of a company losing its footing on two fronts at once.

On April 8, 2026, Meta launched Muse Spark, the first model to come out of Meta Superintelligence Labs. The press release repeats the phrase "personal superintelligence" a bit too often. Once that marketing layer is stripped away, what emerges is a genuine model that puts Meta back in the race at the highest level.

To compare Meta's new model with one of its best current competitors, we recommend our guide Muse Spark vs Claude Opus 4.6. You can also read our Muse Glimmer guide and our tutorial on running Muse Glimmer locally.

## What is Muse Spark?

Muse Spark is a natively multimodal reasoning model that handles text, images, audio, and tools within a single architecture. It supports a visual chain of thought: the model can break down image-based problems step by step instead of producing a single answer. Multi-agent orchestration is also part of the equation, and we'll come back to that.

The early Llama models returned answers by matching patterns learned during training. Muse Spark thinks through the problem before responding. That's the real change.

### Who's in charge?

Meta Superintelligence Labs, or MSL, was created on June 30, 2025, when Mark Zuckerberg reorganized the company's AI operations. Alexandr Wang, former CEO of Scale AI, came on as Chief AI Officer; Meta had invested about $14 billion in Scale AI as part of the deal.

Nat Friedman, former CEO of GitHub, oversees product and applied research, and Shengjia Zhao, who co-created GPT-4 and o1 at OpenAI (the same o1 that Muse Spark is now compared to in benchmarks), is Chief Scientist.

A third factor matters: Yann LeCun, Meta's longtime Chief AI Scientist and the company's leading advocate for open source, left in November 2025. His departure followed organizational changes that had limited his role and the team's shift to closed development.

## What's new with Muse Spark (and why it matters)?

The highlights: reasoning modes, an overhauled training pipeline, and a deliberate focus on health. Let's go through them.

### Three reasoning modes

Muse Spark offers three ways to interact with it, and the distinction is worth understanding before you try it.

- **Instant** is the default mode for everyday questions. It responds quickly without extended reasoning, much like a standard chat model.
- **Thinking** uses an extended chain of thought. The model takes longer, breaks down intermediate steps, and generally performs better on hard problems. Most of the benchmark results cited here come from this mode.
- **Contemplating** is the most interesting. Details below.

A useful clarification right away: Contemplating mode rolled out gradually and wasn't available to everyone on launch day. If you don't see it yet, that's normal.

### Contemplating mode

Contemplating mode launches several reasoning agents in parallel, then combines their outputs into a single response. Where Gemini's Deep Think and OpenAI's GPT Pro mode extend reasoning by thinking longer, Muse Spark does it by thinking broader. More agents work simultaneously rather than one agent working longer.

Meta claims this approach produces comparable results with lower latency, since the agents operate in parallel rather than sequentially. Those latency figures have not yet been independently confirmed, but the benchmarks for Contemplating mode lead on several difficult evaluations (we'll come back to that).

This is an inference-time feature, not an architectural choice. The model itself doesn't change.

### Reinforcement learning scaling and thought compression

Meta rebuilt its training pipeline from scratch during the nine months of Muse Spark's development. The claims about reinforcement learning (RL) come from Meta's technical blog and have not been independently verified.

The most interesting point is a technique called thought compression. During RL training, the model is rewarded when it finds the right answer, but is also penalized for thinking time, which amounts to limiting output tokens. This induces three-phase behavior on complex tasks such as math problems.

First, the model improves by "thinking longer." Then the length penalty kicks in and forces the model to solve the same problems with far fewer tokens. At some point, it expands its reasoning again and surpasses its previous ceilings while using fewer tokens.

The practical consequence: the model learned to do more with less. This claim rests on Meta's training curves, which have not been independently validated.

### A compute requirement divided by 10

Meta claims its new architecture matches the performance of Llama 4 Maverick with ten times less training compute. This is a matter of architectural efficiency, not a ceiling for Muse Spark. Llama 4 Maverick scored 18 on the Artificial Analysis Intelligence Index. Muse Spark scored 52.

Artificial Analysis's token efficiency figures point in the same direction. Muse Spark used 58 million output tokens. GPT-5.4 used 120 million. Claude Opus 4.6 used 157 million.

### Health: a deliberate focus

Health is the benchmark territory where Muse Spark stands out the most, and that's no accident. Meta worked with over 1,000 physicians to curate training data specific to medical reasoning.

The model can generate interactive displays covering nutritional composition, medication information, and exercise physiology. On HealthBench Hard, Muse Spark scored 42.8 versus 40.1 for GPT-5.4 and 20.6 for Gemini 3.1 Pro. This gap with Gemini is confirmed under independent evaluation.

This is clearly Meta's answer to ChatGPT Health. Meta's argument for its competitiveness: the social context of 3 billion users, which would give it an edge in understanding how people actually ask their health questions. Whether this holds up for complex or atypical queries—beyond the common questions that dominate benchmarks—remains to be seen.

## And what about Llama in all this?

The developer community is asking a legitimate question, one that deserves a clear answer.

Muse Spark is not open-source. All Llama models up to Llama 4 shipped with weights that developers could download and run locally. Communities like r/LocalLLaMA were built on this. That use case is gone.

The reason cited by Meta is partly competitive: Chinese labs, including DeepSeek, used Llama's weights to accelerate their research. Wang stated that the company "hopes" to open-source future Muse models, without a timeline. "Hopes" carries a heavy weight here.

The Llama team was integrated into Wang's lab, and Llama 4 is the last model to come out of the old structure. Whether Llama continues alongside Muse or quietly fades away, Meta isn't saying.

## Muse Spark benchmark results

With Muse Spark, the benchmarks require distinguishing one important thing from the outset: given Meta's track record with Llama 4, keep the figures declared by the publisher clearly separate from those independently verified.

Here are the results in Thinking mode, which allow for the fairest comparisons.

Source: Meta Superintelligence Labs / ai.meta.com

