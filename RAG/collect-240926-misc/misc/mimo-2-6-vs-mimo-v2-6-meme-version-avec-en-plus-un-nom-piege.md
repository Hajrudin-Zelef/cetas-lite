---
id: collect-240926-misc/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege
title: "mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege"
domain: orcarouter
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Moonshot", "OpenAI", "Xiaomi", "Z.ai"]
dates: ["2025-12", "2026-09-21"]
keywords: ["agent", "astra", "attention", "benchmark", "benchmarks", "claude", "context window", "cost", "embedding", "fable 5", "fp8", "glm"]
source: docs/RAG/clean_en/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege.md
source_anchor: ""
source_lines: [1, 71]
sha256: e06676d348953248264ef27d738a1713e51d07ad354f82c2386f0415c17ac2de
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

• What does `config.json` say about the drafter? The Flash card describes a five-layer speculative drafter predicting seven tokens per pass; the config sets `num_nextn_predict_layers` to 3. The runtime reads the config.

• Which identifier are you actually sending? Xiaomi's platform lists explicit lowercase identifiers — `mimo-v2.6-pro`, `mimo-v2.6-flash`, `mimo-v2.6-pro-ultraspeed`. There is no plain `mimo-v2.6` among them. Xiaomi's own documentation also notes that users migrated from the invite-only test need to change the model name to continue, which is the same lesson delivered by the vendor.

## Where a router has a place, and where it doesn't

The honest position on this release is that the MiMo-V2.6 checkpoints are a self-hosting decision first, before they are an API decision. They are MIT-licensed downloads first; the choice between the 309B checkpoint and the 1.02T one is a choice between two GPU bills more than between two measuring instruments. We do not route any Xiaomi model, and nothing here should be read as suggesting otherwise — there is no MiMo-V2.6-Pro route on our platform, and the availability mention in Xiaomi's own materials points to Xiaomi's own channels.

What a router is actually for in this situation is the other half of the decision. If you are self-hosting a 1.02-trillion-parameter checkpoint to serve a workload that only sometimes needs frontier capability, the useful architecture is a cheap local path for the bulk of traffic and a hosted path for the tail — and it is on the hosted path that **a single API key for 200+ models, at provider list price passed through with 0% margin** does something concrete, because the frontier models you would benchmark MiMo-V2.6-Pro against change their pricing without warning, and a pass-through means the change reaches you the same day. **Automatic failover** is the other half: an unproven or freshly released model is exactly what you want on a path that can switch rather than on a production path of its own.

This is a statement about how to place this release in a stack, not a claim about its hosting.

## What to watch

Three things would change how this release reads.

An independent run of DeepSWE or Terminal Bench on MiMo-V2.6-Flash would do more for the family than any new announcement, because Flash is the member with no third-party score and the one most teams would actually run. An explanation of the parameter-count and drafter discrepancies would remove the last reason to distrust the weight data, that is, the number that determines the hardware. And an independent attempt to reproduce the RL pipeline from the published environments is the real test of the claim Xiaomi is actually making — that the interesting artifact here is not the checkpoint but the training loop that produced it.

Until then, the practical summary of MiMo-2.6 vs MiMo-V2.6 is brief. One version, two spellings, two checkpoints, one of which bears the name of the flagship model in common usage. Type either spelling into a search bar and you will find this generation. Type either into an API call and nothing will respond, because neither is an identifier — and that, rather than the missing V, is the distinction that matters.
