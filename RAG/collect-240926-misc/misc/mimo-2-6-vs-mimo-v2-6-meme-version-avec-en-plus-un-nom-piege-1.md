---
id: collect-240926-misc/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege-1
title: "mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege"
domain: orcarouter
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Moonshot", "OpenAI", "Xiaomi", "Z.ai"]
dates: ["2025-12", "2026-09-21"]
keywords: ["agent", "astra", "attention", "benchmark", "benchmarks", "claude", "context window", "cost", "embedding", "fable 5", "fp8", "glm"]
source: docs/RAG/clean_en/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege.md
source_anchor: ""
source_lines: [1, 52]
sha256: 7cdf6e0f8ae7c18b358c06f7da10f6e84bca585d4abd0fa0186fee30ea94c869
---

# mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege

<!-- source: https://www.orcarouter.ai/fr/blog/mimo-2-6-vs-mimo-v2-6 -->

There is no confrontation here, and saying so quickly is the most useful thing this page can do. **MiMo-2.6** and **MiMo-V2.6** are the same version written two ways — a dropped hyphen, a letter that the vendor's own documents keep and that most secondary write-ups do not. The models that actually exist under this name are **MiMo-V2.6-Pro**, the 1.02-trillion-parameter flagship, and **MiMo-V2.6-Flash**, the 309-billion-parameter efficiency checkpoint, both released and open-sourced by Xiaomi on September 21, 2026. If you landed here because a search engine treated the two spellings as two products, the answer is one sentence long: this is a single product line, and the name you typed decides nothing.

What does decide something is the third name in this article's problem-space title — the one that looks like this family without being part of it. That collision cost people very real hardware this month, and it deserves more of your attention than the hyphen does.

## How Xiaomi numbers these things

The "2.6" is a generation index, not a parameter count, not a size class, not a capability tier. Xiaomi's line runs MiMo-V2, then MiMo-V2.5, then MiMo-V2.6, and each generation ships as a small set of checkpoints rather than a single artifact. That is the design decision underneath the whole confusion: the family name gets reused across members, so it can be used honestly to mean the series, and lazily to mean whichever member the writer had in mind.

Read the sequence, and the shape of this release comes into focus more clearly than the announcement page lets on.

• MiMo-V2, December 2025 — the generation whose flagship checkpoint was MiMo-V2-Flash. A sparse mixture of experts, 309 billion parameters total, 15 billion active, a 256K context window, and a hybrid attention scheme interleaving sliding-window attention with full attention at a 128-token sliding window.

• MiMo-V2.5 — the generation before this one, and the baseline Xiaomi measures itself against. Its Pro checkpoint logged 19.0 on DeepSWE v1.1 in the vendor's own table.

• MiMo-V2.6, September 21, 2026 — two native omnimodal checkpoints, published to the model hub eighteen seconds apart at 15:39:33 and 15:39:51 UTC, both open-weight, both carrying an MIT license tag, both accompanied by a technical report. A 1M-token context window on both, a max positional embedding of 1,048,576 tokens in the config, up to 128K output tokens, and, on input, text, images, video, and audio.

• The rest of the set — MiMo-V2.6-Pro-UltraSpeed, a serving tier built from the Pro weights rather than a distinct model; a desktop client; MiMo-V2.6-Distill-Qwen-9B; and, per Xiaomi, more than seven thousand reinforcement-learning task environments plus an end-to-end training framework and mini-harnesses released alongside the weights.

Now notice what "2.6" does not tell you. It does not tell you the count, the license, the modality, or whether the thing you are about to download will fit on your node. The gap between 309B total and 1.02T total sits entirely inside one generation number.

## The name that matches a genuinely different model

MiMo-V2-Flash shipped in December 2025, and Xiaomi's release blog post describing it outlines a mixture-of-experts model with 309 billion parameters total, 15 billion active, a hybrid attention scheme alternating sliding-window attention with full attention, and a 128-token sliding window. Compare that against the bullet list above.

The 2025 checkpoint and the 2026 one share the same overall structure, the same activation budget, and the same sliding-window size — nine months apart, from different training runs, with different capabilities and a different benchmark table. A search for "MiMo V2.6 Flash" will cheerfully hand you pages about MiMo-V2-Flash, complete with a 256K context figure and an input price around a tenth of a dollar per million tokens that belongs to the older model. If you are sizing a node, that is not a cosmetic error. The two are different downloads, with different serving recipes, and the new one claims four times the context.

That is the collision to guard against, and it has nothing to do with whether you write the V. The trap is a version string that differs by a single character between a shipped model and a superseded one — the same failure mode that gets MiMo-V2.5-Pro's numbers passed around as if they described the new flagship. Checking the repository, not the name, is what settles it.

## What actually changed in 2.6?

Once the naming question is out of the way, the announcement carries four substantive claims, and they are worth separating by who made them.

The measurable number: MiMo-V2.6-Pro sits at **46 on the Artificial Analysis Intelligence Index**, version 4.3.2. That score was produced by Artificial Analysis using its own evaluation harness, not by Xiaomi, which makes it the one figure in this announcement a reader does not have to discount. On that same leaderboard, the model is the highest-ranked open-weight entry, above GLM-5.3 at 45 and Kimi K3 at 44, and below Claude Fable 5.1 and GPT-6 Astra at 53, and Claude Opus 5 at 51.

Xiaomi's own, and much higher: Xiaomi's tables place MiMo-V2.6-Pro at 71.9 on DeepSWE v1.1, versus 19.0 for the previous generation's Pro checkpoint, with Flash at 67.9. Read that as a claim about how far a single reinforcement-learning run moved the model, rather than as a position in a ranking — the numbers come from Xiaomi's harness, Xiaomi's grader, and an offline run that was not submitted anywhere and has not been reproduced.

The unusual part: the training run was streamed live. Xiaomi published a live dashboard throughout September showing costs, tokens, and benchmarks as they moved, then posted the bills — roughly $2.62 million for the Pro run and $854,000 for the Flash run, thirty RL steps per run and about 750,000 trajectories, all finished in under six days. Labs almost never publish this, and Xiaomi published the failures alongside it: a restart from a GPU out-of-memory caused by expert load imbalance, a network failure between the training cluster and the evaluator deployment, and a Flash restart after an infrastructure error went undetected for about three hours.

The structural part: the post-training pipeline ships with the download. Xiaomi says the RL environments, training framework, and harnesses are released with the weights. Whether what shipped is enough to reproduce a 71.9 is a separate question, and it gets settled by the people who try, not by the release notes.

One detail that applies specifically to the streamed numbers. The dashboard had Pro at 72.57 on DeepSWE with a mini-swe-agent harness averaged over three runs; the finalized model card lists 71.9 for the released weights. Flash moved the other way, from 65.68 to 67.9. If you are quoting dashboard readings from last week, they describe training snapshots rather than the artifacts you would download today.

## A small check before you commit to anything named MiMo V2.6

The naming problem does not get solved by being careful about the letter V. It gets solved by checking four things in the repository rather than in the title.

• Does the repo name end in `-RL`? Both released checkpoints do, and it does not mean this is an adapter over a base model. It marks the post-training lineage — these are the full checkpoints from the streaming run.

• What does the Safetensors block say about parameter count? It says 159B for Flash and 524B for Pro. Neither matches the 309B and 1.02T in the model cards. Xiaomi has not explained the discrepancy. The size comes from the released weight files, which total 172.9 GB of FP8 data across 65 shards for Flash.

