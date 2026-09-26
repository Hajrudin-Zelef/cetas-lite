---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/overview
title: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "MiniMax", "Moonshot", "OpenRouter"]
dates: ["2026-02-16", "2026-03-30", "2026-04-02", "2026-04-16", "2026-05", "2026-05-19", "2026-05-20", "2026-05-21", "2026-06-02", "2026-06-23", "2026-06-24", "2026-07", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-07-31", "2026-08-03", "2026-08-12", "2026-08-14", "2026-08-26", "2026-09", "2026-09-02", "2026-09-22"]
keywords: ["qwen", "agent", "apache", "benchmark", "claude", "cost", "datacenter", "fable 5", "kimi", "license", "moe", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1367, 1402]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 979077d2c2c13ed83f317a4f525cd99e43114e6f20e6f263f66c087703b0aab4
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

