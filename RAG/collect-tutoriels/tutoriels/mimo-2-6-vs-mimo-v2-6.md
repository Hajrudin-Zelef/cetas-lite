---
id: collect-tutoriels/tutoriels/mimo-2-6-vs-mimo-v2-6
title: "MiMo-2.6 vs MiMo-V2.6 : deux orthographes, une seule version, et un troisième nom qui correspond vraiment à un modèle différent"
domain: tutoriels
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Moonshot", "OpenAI", "Xiaomi", "Z.ai"]
dates: ["2025-12", "2026-09", "2026-09-23"]
keywords: ["agent", "astra", "attention", "benchmark", "benchmarks", "claude", "context window", "cost", "embedding", "fable 5", "fp8", "glm"]
source: docs/RAG/Collect RAG/07_tutoriels/mimo-2-6-vs-mimo-v2-6.md
source_anchor: ""
source_lines: [1, 68]
sha256: b61b088fa5c1654f7928bded04d7ca7425e65ae81161791d7e462fe949fd09a9
---

# MiMo-2.6 vs MiMo-V2.6 : deux orthographes, une seule version, et un troisième nom qui correspond vraiment à un modèle différent

## Metadata

- **Source** : https://www.orcarouter.ai/fr/blog/mimo-2-6-vs-mimo-v2-6
- **Site** : OrcaRouter
- **Type** : Article
- **Language** : fr
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This OrcaRouter article by Alistair Wren (published 22 September 2026) clarifies a naming confusion around Xiaomi's MiMo models. Its key message up front: MiMo-2.6 and MiMo-V2.6 are the same version written two ways — a dropped hyphen and a letter ("V") that the vendor's own documents keep and most secondary coverage drops. The models that actually exist under this name are MiMo-V2.6-Pro, the 1,020-billion-parameter flagship, and MiMo-V2.6-Flash, the 309-billion-parameter efficiency checkpoint, both released and open-sourced by Xiaomi on 21 September 2026. If a search engine treated the two spellings as two products, the answer is that it is a single product line, and the spelling you typed decides nothing. What does decide something is the third name in the article's problem-space title — one that looks like it belongs to this family but does not. That collision has cost people real hardware this month.

The article explains how Xiaomi numbers things: "2.6" is a generation index, not a parameter count, size class, or capability tier. The line follows MiMo-V2, then MiMo-V2.5, then MiMo-V2.6, each generation shipped as a small set of checkpoints rather than a single artifact. This design decision underlies the confusion: the family name is reused across members, so it can honestly designate the series and lazily designate whichever member the author had in mind. Reading the sequence: MiMo-V2 (December 2025) had a flagship checkpoint MiMo-V2-Flash — a sparse mixture of experts, 309B total parameters, 15B active, 256K context, hybrid attention interleaving sliding-window attention and full attention with a 128-token sliding window. MiMo-V2.5 is the previous generation and the benchmark baseline Xiaomi measures against; its Pro checkpoint scored 19.0 on DeepSWE v1.1 in the vendor's own table. MiMo-V2.6 (21 September 2026) shipped two natively omni-modal checkpoints, published on the model hub 18 seconds apart at 15:39:33 and 15:39:51 UTC, both open access, both MIT-licensed, both with a technical report. Both have a 1M-token context window, a maximum positional embedding of 1,048,576 tokens in the config, up to 128K output tokens, and text, image, video and audio input. The rest of the set: MiMo-V2.6-Pro-UltraSpeed (a service tier built from the Pro weights, not a distinct model), a desktop client, MiMo-V2.6-Distill-Qwen-9B, and, per Xiaomi, more than seven thousand RL task environments plus an end-to-end training framework and mini-harnesses published with the weights. The point: "2.6" tells you nothing about count, license, modality, or whether the download fits your node. The gap between 309B total and 1.02T total lies entirely within one generation number.

The truly different model is MiMo-V2-Flash (December 2025), which has the same general structure, same activation budget, and same sliding-window size as the 2026 checkpoint but is a different download with different serving recipes, and the new one claims four times the context. The trap is a version string differing by one character between a shipped and a superseded model — the same failure mode that circulates MiMo-V2.5-Pro numbers as if they described the new flagship. You resolve it by checking the repository, not the name.

What actually changed in 2.6: four substantive claims, distinguished by who made them. (1) The measurable figure: MiMo-V2.6-Pro scores 46 on the Artificial Analysis Intelligence Index v4.3.2, produced by Artificial Analysis with its own harness — the only figure a reader need not discount. In that ranking it is the highest-rated open-weights entry, above GLM-5.3 at 45 and Kimi K3 at 44, and below Claude Fable 5.1 and GPT-6 Astra at 53 and Claude Opus 5 at 51. (2) The vendor's own, much higher figure: Xiaomi's tables place MiMo-V2.6-Pro at 71.9 on DeepSWE v1.1, versus 19.0 for the previous Pro checkpoint, with Flash at 67.9 — read as a claim about RL progress, not a ranking position. (3) The unusual case: the training run was streamed live. Xiaomi published a live dashboard throughout September showing costs, tokens and benchmarks, then released the accounts — about $2.62 million for the Pro run and $854,000 for the Flash run, thirty RL training steps each, roughly 750,000 trajectories, all completed in under six days. Labs almost never publish this, and Xiaomi also published the failures: a GPU out-of-memory restart caused by expert load imbalance, a network outage between the training cluster and the evaluator deployment, and a Flash restart after an infrastructure error went undetected for about three hours. (4) The structural aspect: the post-training pipeline is part of the download — Xiaomi claims the RL environments, training framework and harnesses are published with the weights. Whether that suffices to reproduce 71.9 is a separate question for those who try. One detail applies to the live-streamed figures: the dashboard showed Pro at 72.57 on DeepSWE with a mini-swe-agent harness averaged over three runs; the finalized model card shows 71.9 for the published weights. Flash moved the other way, from 65.68 to 67.9.

The article gives four things to verify in the repository rather than the title: whether the repo name ends in -RL (both published checkpoints do; it marks post-training lineage, not an adapter); what the Safetensors block says about parameter count (159B for Flash and 524B for Pro — neither matches the 309B and 1.02T of the model cards, an unexplained discrepancy; the size comes from published weight files totalling 172.9 GB of FP8 data across 65 shards for Flash); what config.json says about the drafter (the Flash card describes a five-layer speculative drafter predicting seven tokens per pass; the config sets `num_nextn_predict_layers` to 3, and the runtime reads the config); and which identifier you actually send (Xiaomi's platform lists explicit lowercase IDs — `mimo-v2.6-pro`, `mimo-v2.6-flash`, `mimo-v2.6-pro-ultraspeed` — with no plain `mimo-v2.6`). Xiaomi's own docs also note that users migrated from the invite test must change model names to continue.

The article's honest position: the MiMo-V2.6 checkpoints are first a self-hosting decision, then an API decision. They are MIT-licensed downloads; choosing between the 309B and 1.02T checkpoint is a choice between two GPU bills more than two measurement instruments. OrcaRouter routes no Xiaomi model, and nothing in the article should be read as suggesting otherwise — there is no MiMo-V2.6-Pro route on the platform. Where a router helps is the other half of the decision: a cheap local path for most traffic and a hosted path for the tail, where one API key for 200+ models at vendor list price with 0% markup matters because frontier model prices change without warning. It closes with three things to watch: an independent DeepSWE or Terminal Bench run on MiMo-V2.6-Flash (which has no third-party score), an explanation of the parameter-count and drafter discrepancies, and an independent attempt to reproduce the RL pipeline. The practical summary: one version, two spellings, two checkpoints, one of which carries the flagship name in common usage. Type either spelling into a search bar and you'll find this generation; type either into an API call and nothing will answer, because neither is an identifier — and that, not the missing V, is the distinction that matters.

## Key points

- MiMo-2.6 and MiMo-V2.6 are the same version written two ways; not two products.
- Real models: MiMo-V2.6-Pro (1,020B flagship) and MiMo-V2.6-Flash (309B efficiency), open-sourced 21 September 2026, MIT license.
- "2.6" is a generation index, not parameter count, size class, or capability tier.
- The dangerous collision is MiMo-V2-Flash (Dec 2025), a different download with different serving recipes and 256K context vs the new 1M.
- Only independently produced figure: MiMo-V2.6-Pro scores 46 on Artificial Analysis Intelligence Index v4.3.2 (highest open-weights entry).
- Xiaomi vendor figures: 71.9 DeepSWE v1.1 for Pro (vs 19.0 previous gen) and 67.9 for Flash.
- Training run streamed live; costs ~$2.62M (Pro) and ~$854K (Flash), 30 RL steps, ~750,000 trajectories, under six days; failures published too.
- Verify repo name (-RL), Safetensors param count, config.json drafter, and exact lowercase API IDs.

## Technical data / figures

| Item | MiMo-V2.6-Pro | MiMo-V2.6-Flash |
| --- | --- | --- |
| Claimed total params | 1,020B (1.02T) | 309B |
| Safetensors param count | 524B | 159B |
| Context window | 1M tokens | 1M tokens |
| Max positional embedding | 1,048,576 tokens | 1,048,576 tokens |
| Max output | 128K tokens | 128K tokens |
| Input modalities | text, image, video, audio | text, image, video, audio |
| License | MIT | MIT |
| Release | 21 Sept 2026 | 21 Sept 2026 |
| Weight size | — | 172.9 GB FP8, 65 shards |
| DeepSWE v1.1 (vendor) | 71.9 | 67.9 |
| Artificial Analysis Index | 46 | — |
| API identifier | `mimo-v2.6-pro` | `mimo-v2.6-flash` |

Generation timeline:

| Generation | Date | Key checkpoint / notes |
| --- | --- | --- |
| MiMo-V2 | Dec 2025 | MiMo-V2-Flash: 309B total / 15B active MoE, 256K context, hybrid attention, 128-token sliding window |
| MiMo-V2.5 | — | Previous generation; Pro scored 19.0 on DeepSWE v1.1 |
| MiMo-V2.6 | 21 Sept 2026 | Pro + Flash, omni-modal, 1M context, MIT |

Training run figures: Pro ~$2.62M, Flash ~$854K; 30 RL steps each; ~750,000 trajectories; completed in under six days. Dashboard DeepSWE: Pro 72.57 (mini-swe-agent, avg of 3) → final card 71.9; Flash 65.68 → 67.9.

Other IDs: `mimo-v2.6-pro-ultraspeed` (service tier), MiMo-V2.6-Distill-Qwen-9B. Config discrepancy: Flash card says 5-layer drafter predicting 7 tokens/pass, but config sets `num_nextn_predict_layers` = 3.

## Why this source matters for the RAG

It is a precise disambiguation guide for a confusingly named model family, with concrete repository-verification steps, benchmark provenance, and training-cost disclosures. It is valuable for RAG questions about MiMo/Xiaomi models, model naming pitfalls, and how to verify open-weight artifacts before deployment.
