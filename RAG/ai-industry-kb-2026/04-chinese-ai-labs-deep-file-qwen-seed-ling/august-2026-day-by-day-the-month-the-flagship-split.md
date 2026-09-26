---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/august-2026-day-by-day-the-month-the-flagship-split
title: "August 2026, day by day — the month the flagship split"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "MiniMax", "Moonshot", "Nvidia", "OpenRouter"]
dates: ["2025-04", "2026-04-02", "2026-04-16", "2026-06-23", "2026-06-24", "2026-07-16", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-07-31", "2026-08", "2026-08-01", "2026-08-03", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-17", "2026-08-26", "2026-09-02", "2026-09-22"]
keywords: ["agent", "apache", "benchmark", "benchmarks", "claude", "cost", "datacenter", "disclosure", "fable 5", "fp8", "kimi", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1650, 1691]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 80ff0deb9164f5b8bedb16e7dcd28fab0cf6b33839bac2ade768e5e1ac470d94
---

# August 2026, day by day — the month the flagship split

ByteDance announced Seed 2.1, Seed 2.1 Pro, and a preview tier at the Volcano Engine FORCE conference on 2026-06-23/24. The official launch blog names no Turbo variant; the Turbo's presence at FORCE rests on a single June 24 trade report (DataNorth), while Western trackers only logged the model on August 10–12 (LLM Gateway PR #3580, NanoGPT catalog, OpenRouter) — first-seen dates, not releases (one engineering journal documents the "launched twice, 47 days apart" discrepancy). Defensible dating: 2026-06-24, with the single-source caveat kept visible. Verified specs via catalogs + LLM Gateway e2e mapping: 262K context (262,144); text/image/video in → text out; reasoning with full effort ladder (`none` disables thinking); tool calling; structured output; streaming; closed/proprietary (parameter count, cutoff, license all undisclosed [UNVERIFIED]). Pricing $0.50/$2.50 per M tokens (OpenRouter/NanoGPT, verified current 2026-09-22): the faster, lower-cost variant for high-throughput/latency-sensitive agent workloads vs the flagship Pro. ByteDance's launch materials used chart-only figures with superlatives and no absolute numbers in prose ("a masterclass in avoiding numbers in the text" — one journal); no independent Turbo benchmarks exist [UNVERIFIED]. Seedance 2.5 (video) is cross-referenced to §12 and not duplicated here.

### August 2026, day by day — the month the flagship split

- 2026-08-01: MarkTechPost launch-day coverage window for the late-July releases still landing (MiniMax H3 coverage dated Aug 1 belongs to the July cluster's tail).
- 2026-08-03: Qwen3.8-Max GA — Bloomberg, Quartz, MarkTechPost, zeronoise.ai, bitrouter registry spec all pin the date; $2/$6 flat pricing; 1M context; 512 experts.
- 2026-08-11: 302.AI publishes the Qwen Image 3.0 Pro vs GPT-Image-2 head-to-head in infographic/layout scenarios — the only independent comparative data point on Image 3.0 [COMMUNITY].
- 2026-08-12: Qwen3.8-Max weights drop (`Qwen/Qwen3.8-2.4T-A95B` + FP8 on HF/ModelScope) the same day NVIDIA publishes its GB300 NVL72 deployment blog confirming "Alibaba released the open weights" — hardware and weights landing together is the tell that the artifact was staged for datacenter partners first.
- 2026-08-14: Qwen3.8-27B weights under Apache 2.0 (official @Alibaba_Qwen announcement: "We promised open weights for Qwen3.8. Now, time to meet them!") — the sentence itself confirms the two-step structure: Max weights were the promise kept on the 12th, the 27B was the democratization event on the 14th.
- 2026-08-17: gigazine.net English coverage — the date that proves the brief's "Aug 12" for the 27B was wrong.
- 2026-08-26: Qwen3.8-Flash (open weights, 6B active) + Qwen3.8-Flash-Next (hosted, Qwen4-generation preview) — the efficiency tier closing out the month.

### Qwen3.8-Max-0902 (2026-09-02) — recap only

Newer Max checkpoint dated 2026-09-02; 4th (preliminary) in the September 5 Code Arena WebDev table vs Claude Opus 5 Max (contracollective.com comparison) [COMMUNITY]. Canonical home: wave 2.1/05-open-models-delta.md — this part records only the date, the checkpoint's existence, and the cross-reference. No license change, no spec change, no independent benchmark package beyond the Code Arena table at cutoff.

### Naming guard

Qwen3.8-Max ≠ Qwen3-8B (the 8.2B Apache-2.0 dense model released April 2025) — search engines conflate them; consolidation indexing must keep them separate.

### The WAIC announcement (2026-07-19): what was and wasn't disclosed

- Announced at the World AI Conference in Shanghai; coverage within 96 hours from MLQ, MarkTechPost, Dataconomy, EqualOcean, Straits Times, memeburn — the best-attested single date in the Qwen 2026 line.
- Disclosed: 2.4T total parameters, sparse MoE, first Qwen multimodal model above 1T params (developer Shuai Bai), text/image/video/document input, immediate Qwen3.8-Max-Preview via Token Plan/Qoder/QoderWork at 10% of standard pricing, the vendor's "just behind Claude Fable 5" positioning [VENDOR].
- NOT disclosed: active-per-token count (the gap every independent technical review flagged), no benchmark table, no model card, no license. Independent test results did not exist at announcement.
- Sequencing: two days after Moonshot's Kimi K3 (2.8T) — Alibaba holds ~36% of Moonshot (MLQ) — so the announcement reads as a same-family one-upmanship as much as a competitive response; EqualOcean explicitly frames the preview as promising an open-weight release, a promise the 2026-08-12 drop kept under custom terms.
- RAG guidance: treat the announcement and the GA (2026-08-03) as distinct events with distinct disclosure states — the "just behind Fable 5" claim belongs to the announcement, not to the GA benchmark package.

### The two-track strategy, quarter by quarter

Q1 2026 (Feb–Mar): the pattern is set at launch — permissive open weights (397B-A17B, Apache 2.0) ship the same day as the proprietary hosted flagship (Plus, 1M ctx). By late March, Omni extends the hosted track into omnimodal territory with no weights at all, and The Information reads it as a possible open-source retreat — but the open track never stops.

Q2 2026 (Apr–Jun): Qwen3.6 repeats the pattern (proprietary Plus 2026-04-02, open 35B-A3B 2026-04-16) while the entire Qwen3.7 generation ships closed — Max unveiled at the May Hangzhou summit, Plus official June 2 as an agent-oriented GUI+CLI model, and the Turbo-adjacent efficiency story left to ByteDance (Seed 2.1 family at FORCE, June 23/24) and Ant's InclusionAI (Ling 3.0 Flash, July 23 — technically Q3 but conceived in the same efficiency wave).

Q3 2026 (Jul–Sep): the July cluster — Kimi K3 (2026-07-16/17, 2.8T, Moonshot, ~36% Alibaba-owned), Qwen3.8 announced (2026-07-19, 2.4T), Qwen Image 3.0 closed (2026-07-21), Ling 3.0 Flash free (2026-07-23), Qwen3.7 Flash public (2026-07-27), MiniMax H3 (2026-07-31). Five Chinese labs shipping frontier-class or frontier-adjacent releases within fifteen days. August resolves the Max question in two steps: GA on the 3rd ($2/$6, 1M ctx, multimodal) and the weight drop on the 12th under the custom license with the text-only/thinking-only artifact caveat; the Apache-2.0 27B on the 14th is the actual democratization event. Late August adds the efficiency tier (3.8-Flash open, Flash-Next hosted preview of Qwen4-generation architecture). September adds the 0902 checkpoint and the finance-specialized Ling Fin.

The net arc: Alibaba moved from binary open/closed (Qwen3.5's clean Apache-vs-hosted split) to a three-way license taxonomy (permissive / custom-restrictive / closed-API) while holding both tracks open all year. The "shift away from open-source" narrative is only accurate for the flagship modality frontier (omni, image, Max-hosted); the mid-tier open-weight pipeline accelerated in parallel. The correct consolidation framing is "two tracks converging on terms-as-monetization," not "open then closed."

### Qwen3.5-Omni and the open-weight strategy question (deep dive)

Omni matters less as a model than as a strategy signal — it is the first major Qwen flagship to ship with no open weights, and the template for everything after it (Qwen3.7-Max/Plus, Image 3.0, the Max weight drop's custom license). The Information's "potential shift away from a strategy that centers on open-source models" framing is directionally right but incomplete: what actually happened is modality-gated openness. Text mid-tier models stayed Apache-2.0 all year; the models that crossed a product-maturity threshold in a new modality (omnimodal understanding, image generation, agent GUI/CLI control) shipped hosted or custom-licensed. The license follows the modality's monetizability, not the parameter count — the 2.4T Max got weights (under custom terms) while the undisclosed-scale Omni got none, because a downloadable omni model is a product in a way a text checkpoint is an ingredient.

