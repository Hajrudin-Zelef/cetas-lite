---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/overview
title: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "Z.ai"]
dates: ["2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-04-21", "2026-05", "2026-05-19", "2026-05-20", "2026-05-21", "2026-06-02", "2026-06-23", "2026-06-24", "2026-07", "2026-07-15", "2026-07-16", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-07-31", "2026-08-03", "2026-08-10", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-17", "2026-08-26", "2026-09", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-05", "2026-09-22"]
keywords: ["qwen", "agent", "agentic", "apache", "benchmark", "benchmarks", "claude", "cost", "datacenter", "fable 5", "fp8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1367, 1434]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 377629ef71a87f00a9e7ed0f8429d4e08c64c8ffa12b3649477400c16dd2a363
---

# 4. Chinese AI Labs — Deep File: Qwen, Seed, Ling
Keywords: Qwen3.5, Qwen3.5-397B-A17B, Qwen3.5-Plus, Qwen3.5-Omni, Qwen3.6-Plus, Qwen3.6-35B-A3B, Qwen3.7-Max, Qwen3.7-Plus, Qwen3.7 Flash, Qwen3.8, Qwen3.8-Max, Qwen3.8-2.4T-A95B, Qwen3.8-27B, Qwen3.8-Flash, Qwen3.8-Max-0902, qwen3.8-max license, Qwen Image 3.0, Ling 3.0 Flash, Seed 2.1 Turbo, Seedance 2.5, open weights, Apache 2.0, Gated Delta Networks, 1M context, $2/$6 per M tokens

## Summary

This section is the deep file on Alibaba's Qwen model line (February → September 2026) plus the two adjacent Chinese-lab items assigned to this part: ByteDance Seed 2.1 Turbo and InclusionAI (Ant Group) Ling 3.0 Flash. The Qwen chronology was verified claim-by-claim against live sources in wave 3.1 research; six claims verify cleanly, two verify with date corrections, and one verifies with a licensing qualification that changes how the RAG must classify the release.

The headline events: Qwen3.5-397B-A17B open weights (Apache 2.0) shipped 2026-02-16 alongside the proprietary hosted Qwen3.5-Plus; Qwen3.5-Omni launched 2026-03-30/31 (a date correction — it was reported as April in the brief), the first major Qwen flagship to ship API-only with no open weights; Qwen3.6-Plus shipped 2026-04-02 (proprietary) and Qwen3.6-35B-A3B followed 2026-04-16 as the open Apache tier; Qwen3.7-Max was unveiled 2026-05-20 at the Alibaba Cloud Summit in Hangzhou with public API on 2026-05-21/23, and Qwen3.7-Plus reached official release 2026-06-02 (preview 2026-05-19) as a proprietary agent-oriented multimodal model; Qwen3.7 Flash went public 2026-07-27 per OpenRouter. Qwen3.8 was announced 2026-07-19 at the World AI Conference in Shanghai as a 2.4-trillion-parameter sparse MoE — Alibaba's first Qwen model above 1T params — carrying the vendor's claim [VENDOR] that it sat "just behind Claude Fable 5," made with no benchmark table, model card, or license published at announcement. Qwen3.8-Max hit general availability 2026-08-03 at $2/$6 per million tokens; its downloadable weights landed 2026-08-12 under a custom restrictive "qwen3.8-max" license (not Apache 2.0), and the Qwen3.8-27B open weights followed 2026-08-14 (a second date correction — not August 12) under Apache 2.0. A Qwen3.8-Flash open-weight release (2026-08-26, 6B active) and the Qwen3.8-Max-0902 checkpoint (2026-09-02) complete the family through the knowledge cutoff.

Three structural findings shape this part. First, "open-weight" is a spectrum in the 2026 Qwen line, not a binary: Apache 2.0 for mid-tier models (3.5-397B-A17B, 3.6-35B-A3B, 3.8-27B), a custom restrictive license with revenue-sharing mechanics for the flagship Max tier (100M MAU or $20M monthly revenue triggers name display; MaaS/AI-assistant businesses past $50M trailing revenue need a separate commercial license), and fully closed API-only releases for Omni, 3.7-Max/Plus, and Qwen Image 3.0. Second, the Qwen3.8-Max open artifact is strictly not the hosted product: the published checkpoint is text-only and thinking-mode-only, dropping the vision input and the native 1M context the API offers — the RAG must treat artifact and service as separate entries. Third, 2026 shows a deliberate two-track Alibaba strategy — permissive mid-tier weights that build ecosystem adoption, hosted or custom-licensed terms for flagships that capture revenue — alongside an industry-wide cost bifurcation where datacenter flagships ($2/$6 per M) coexist with near-zero agent-workhorse pricing (Ling 3.0 Flash Fin at $0.06/$0.18 per M, free OpenRouter tiers).

### Verification posture

Of the nine Qwen claims checked in wave 3.1: six verify cleanly (Qwen3.5 pair, Qwen3.6-Plus, Qwen3.6-35B-A3B, Qwen3.7-Max, Qwen3.7-Plus, Qwen3.7 Flash, Qwen3.8 announcement and GA); two verify with date corrections (Qwen3.5-Omni 2026-03-30/31 not April; Qwen3.8-27B weights 2026-08-14 not August 12); one verifies with the licensing qualification (Qwen3.8-Max weights are real but custom-licensed, not Apache 2.0). New material beyond the brief: the Qwen3.8-Max open artifact's missing-feature gap (text-only, thinking-mode-only), Qwen3.8-Flash (2026-08-26, 6B active, open weights), Qwen Image 3.0 (2026-07-21, closed API-only), the Ling 3.0 Flash family, and Seed 2.1 Turbo (2026-06-24 with reserve). The July 2026 release cluster — Kimi K3 (07-16/17), Qwen3.8 (07-19), Qwen Image 3.0 (07-21), Ling 3.0 Flash (07-23), Qwen3.7 Flash (07-27), MiniMax H3 (07-31) — is treated as a competitive-dynamics datum, not just six isolated dates.

### Coverage and boundaries of this part

- Primary coverage: the full Qwen3.5/3.6/3.7/3.8 lineage (open, proprietary, omnimodal, image), benchmark reality checks for the 3.8 family, the Qwen3.8-Max licensing dissection, InclusionAI Ling 3.0 Flash family, ByteDance Seed 2.1 Turbo (with Seed 2.1/Pro base context), and the Qwen3.5-Omni strategy question.
- Deliberately abbreviated here: Seedance 2.5 (cross-reference §12; canonical home is wave 3/06-multimodal-revolution-check.md §1.3 — announced 2026-06-23 at Volcano Engine FORCE, global Dreamina launch 2026-07-31), Qwen3.8-Max-0902 detail (canonical home is wave 2.1/05 § — 2026-09-02 checkpoint, 4th preliminary in the September 5 Code Arena WebDev table), and the Feb→May 2026 Qwen3.5/3.6/3.7 skeleton (canonical home wave 3/02-openweight-chronology.md).
- Benchmark reality checks (vendor vs third-party vs aggregator figures, with labels) are kept in this section per the coordinator's dedup rule; plain chronology entries collapse to §2.

### Why Seed and Ling belong in this part

ByteDance's Seed 2.1 Turbo and Ant's Ling 3.0 Flash are not Alibaba models, but they are indexed here because they are the two efficiency-tier counterpoints the Qwen story needs: Seed Turbo ($0.50/$2.50, 262K, FORCE 2026-06-24) is the closed-lab answer to the agent-workload question from the same month Qwen3.7 shipped closed, and Ling 3.0 Flash (124B/5.1B, free tier, 2026-07-23) is the cost-floor evidence for the H2-2026 bifurcation thesis. Their canonical model-detail homes are wave 3.1/03 (Seed) and wave 3.1/04 (Ling); this part carries their dates, specs, pricing, and the economic framing, and nothing else.

### Label key used throughout

- [VENDOR] — figure or claim originating with the vendor (Alibaba/Qwen, ByteDance, Ant Group), not independently verified.
- [UNVERIFIED] — no primary or independent confirmation found as of 2026-09-22; includes rumor-phase items.
- [COMMUNITY] — secondary/community evidence (aggregators, community quantizations, benchmark tables from third parties).
- [DIRECTIONAL] — approximate or composite figures useful for ordering/ranking, not precision claims.

## Key dated facts

One-to-two-line facts; chronology detail collapses to §2 (the consolidated open-weight chronology) with cross-references.

- 2026-02-16 — Qwen3.5-397B-A17B (Apache 2.0 open weights, 397B/17B MoE, 262K ctx via YaRN→1M) + Qwen3.5-Plus (proprietary hosted, 1M ctx, tools). Canonical chronology: §2.
- 2026-03-30/31 — Qwen3.5-Omni (hosted API-only; Plus/Flash/Light; Thinker–Talker architecture; 113-language speech recognition). Date corrected from "April." Canonical chronology: §2.
- 2026-04-02 — Qwen3.6-Plus (proprietary API, agentic coding, 1M context). Canonical chronology: §2.
- 2026-04-16 — Qwen3.6-35B-A3B (open weights, 35B/3B MoE). Canonical chronology: §2.
- 2026-04-21 — Qwen3.6-Max-Preview (proprietary preview). Canonical chronology: §2.
- 2026-05-20 — Qwen3.7-Max unveiled (Hangzhou; API public 2026-05-21/23; proprietary, 1M ctx, 65K output). Canonical chronology: §2.
- 2026-05-19/2026-06-02 — Qwen3.7-Plus preview (2026-05-19) → official release (2026-06-02); proprietary agent-oriented multimodal model. Canonical chronology: §2.
- 2026-06-24 — ByteDance Seed 2.1 Turbo at Volcano Engine FORCE; 262K ctx; $0.50/$2.50 per M. Defensible date with single-source caveat on Turbo's presence at FORCE; Western catalog first-seen dates (2026-08-10/12) are aggregator artifacts, not releases.
- 2026-07-19 — Qwen3.8 announced (2.4T MoE, WAIC Shanghai); Qwen3.8-Max-Preview immediate via Token Plan at 10% pricing; vendor "just behind Claude Fable 5" [VENDOR].
- 2026-07-21 — Qwen Image 3.0 (closed, API-only; no weights/report/benchmarks at launch). Canonical chronology: §2.
- 2026-07-23 — InclusionAI Ling 3.0 Flash (124B/5.1B active MoE, 262K ctx; free tier on OpenRouter).
- 2026-07-27 — Qwen3.7 Flash public (per OpenRouter; QwenCloud snapshot 2026-07-15 is an identifier, not the launch date).
- 2026-08-03 — Qwen3.8-Max GA (2.4T/95B active [third-party-reported], 1M ctx, multimodal, $2/$6 per M; custom restrictive license).
- 2026-08-12 — Qwen3.8-Max downloadable weights (`Qwen/Qwen3.8-2.4T-A95B` + FP8; custom qwen3.8-max license; text-only, thinking-mode-only — strictly separate from the hosted multimodal/1M product).
- 2026-08-14 — Qwen3.8-27B weights (Apache 2.0; 27B dense; multimodal). Date corrected from "August 12."
- 2026-08-26 — Qwen3.8-Flash (open weights, 6B active) + Qwen3.8-Flash-Next (hosted only).
- 2026-09-02 — Qwen3.8-Max-0902 (new checkpoint; cross-ref §2 / wave 2.1).
- 2026-09-03 — InclusionAI Ling 3.0 Flash Fin (finance variant; $0.06/$0.18 per M).
- Date unpinned — Qwen3.6-27B dense variant exists in third-party benchmark tables (Terminal-Bench 2.1 60.7, AA Index 38) but no release date was pinned in research; flagged [DATE UNVERIFIED] in consolidation.
- 2026-07-15 — Beijing approved Apple Intelligence powered by Qwen in China (secondary, memeburn.com) — actor-context fact, not a model release.
- 2026-07-15 — QwenCloud snapshot identifier `qwen3.7-flash-2026-07-15` (snapshot, not public launch; public availability 2026-07-27 per OpenRouter).
- 2026-07-16/17 — Moonshot Kimi K3 launch window (2.8T), the immediate predecessor to the Qwen3.8 announcement; cross-ref §5 (Kimi line).
- 2026-07-23 — EqualOcean coverage of Qwen3.8 preview; Alibaba holds ~36% of Moonshot (MLQ) — the ownership link behind the July sequencing.
- 2026-07-31 — MiniMax H3 video model launch (Reuters) — the third July-cluster Chinese-lab release; cross-ref §7/§12, not duplicated here.
- 2026-08-11 — 302.AI head-to-head: Qwen Image 3.0 Pro vs GPT-Image-2 in infographic/layout scenarios [COMMUNITY].
- 2026-08-12 — NVIDIA deployment blog: "Alibaba released the open weights" + GB300 NVL72 reference deployment (>4,000 tok/s per GPU, FP8).
- 2026-08-10/12 — Western aggregator first-seen dates for Seed 2.1 Turbo (LLM Gateway PR #3580, NanoGPT, OpenRouter) — aggregator artifacts, not the release.
- 2026-08-17 — gigazine.net English coverage of the Qwen3.8-27B weight drop (dated three days after the 2026-08-14 official announcement).
- 2026-09-04 — AA Intelligence Index v4.3 composite: Qwen3.8-Max ≈ 40 (4th open) vs GLM-5.3 ≈ 45, Kimi K3 ≈ 44 [DIRECTIONAL].
- 2026-09-05 — Code Arena WebDev table (preliminary): Qwen3.8-Max-0902 4th — covered wave 2.1/05; recap only.
- 2026-09-22 — Seed 2.1 Turbo pricing $0.50/$2.50 per M verified current on OpenRouter/NanoGPT (as-of date for all pricing in this part unless otherwise sourced).

