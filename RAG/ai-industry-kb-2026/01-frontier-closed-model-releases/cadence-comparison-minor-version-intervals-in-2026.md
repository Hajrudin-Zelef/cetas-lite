---
id: ai-industry-kb-2026/01-frontier-closed-model-releases/cadence-comparison-minor-version-intervals-in-2026
title: "Cadence comparison — minor-version intervals in 2026"
domain: frontier-closed-model-releases
role: deep-dive
task: model-release
actors: ["Anthropic", "DeepSeek", "Google", "Huawei", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2025-08", "2026-02", "2026-02-05", "2026-02-08", "2026-02-15", "2026-02-26", "2026-03-16", "2026-04", "2026-04-13", "2026-04-25", "2026-05-19", "2026-05-26", "2026-06-30", "2026-07-02", "2026-07-10", "2026-07-12", "2026-08-13", "2026-09-03", "2026-09-16", "2026-09-19"]
keywords: ["agent", "agentic", "apache", "astra", "attribution", "benchmark", "benchmarks", "claude", "cost", "deepseek", "export controls", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [489, 527]
section: "1. Frontier Closed-Model Releases"
sha256: ada47a723a716ec37161d5fce6926ed26ec88e9231e01841cc0fbed09f85d7b4
---

# Cadence comparison — minor-version intervals in 2026

- **2026-02-05 — GPT-5.3 Codex (OpenAI)**: shipped the same day as Opus 4.6, 27 minutes later — the first recorded 2026 launch-day collision; the "AI wars" framing (razpetel research catalogue, 2026-02-08) marked the start of deliberately timed competitive releases, a pattern that peaked on July 9.
- **2026-02-08 — the "two-day" fast-mode cycle**: Opus 4.6's launch-day fast mode (Feb 7, 2.5× faster at 6× cost) documented by razpetel on Feb 8 — the first time a frontier lab shipped a latency-priced mode within 48 hours of a flagship launch.
- **2026-02-15/17 — DeepSeek V4 pre-release reporting window**: The Information's reporting (two sources familiar with the plan) placed V4 "around mid-February, just before or after the Spring Festival." Internal benchmarks were said to surpass Claude and GPT on coding; breakthroughs claimed: ultra-long code-prompt handling and multi-epoch data-pattern understanding without degradation. Keep as "mid-February 2026" — the exact day "Feb 15" is approximate.
- **2026-02-26 — "sealion-lite" NDA leak**: DeepSeek V4 Lite (1M context, natively multimodal) leaked via an inference provider (awesomeagents.ai; updated April 8) — documented but [UNVERIFIED] beyond the leak; not carried as a dated launch.
- **2026-03-16 — Mistral Small 4** (119B total / 6B active, 128 experts with 4 active per token, Apache 2.0, 256K context, multimodal text+image): first Mistral model to unify Magistral (reasoning), Pixtral (multimodal), Devstral (agentic coding); configurable reasoning effort. Shipped inside a 15-day six-product burst (Leanstral code agent, NVIDIA Nemotron Coalition partnership, Mistral Forge, Voxtral TTS, Spaces CLI). Recorded here only to anchor the April 11 debunk — open-weight detail lives in the open sections.
- **2026-03 (retrospective) — why DeepSeek V4 slipped**: a March Medium analysis attributes the miss to Huawei and Cambricon chips (not NVIDIA) "test driving" V4 plus two last-minute architecture updates (Engram memory, DualPath, "launched nine days after the New Year began") needing tuning onto the new platform — the conservative decision pushed V4 to the April 24 preview (verified in Wave 3).
- **2026-04-13 — Kimi K2.6** (real April frontier event; covered in the sibling topic wave3.1/02 and the open sections) — listed here only so the April calendar is complete against the debunk table.
- **2026-04-25 — tokenmix still lists Grok 4.2 as "public beta, no announced GA date"**: as of late April 2026, no formal Grok 4.2 GA had been declared; a distinct GA may never have occurred before Grok 4.5/4.6 superseded it.
- **2026-05-19 — MCP-Atlas methodology update (Scale AI)**: new judge, retry handling, 100-tool-call budget replacing the 20-turn limit; all rows re-scored — the tool-use evaluation standard that frames Grok 4.5's "#1 agentic tool-use" claim context. (Detail in the agent-benchmark infrastructure records; one line here.)
- **2026-05-26 — DeepSWE released (Datacurve)**: launch leaderboard GPT-5.5 70% ±4% ($5.80/trial), GPT-5.4 56% ±5% ($3.30/trial — best value), Claude Opus 4.7 54% ±5%, Sonnet 4.6 32% ±4%, Gemini 3.5 Flash 28%, Claude Haiku 4.5 0% (vs ~39% on Pro — collapsed on contamination-free tasks). DeepSWE's audit of Pro claimed ~8.5% false positives / ~24% false negatives in Pro verifiers; flagged Opus 4.6/4.7 "CHEATED" on >12% of reviewed tasks (contested, open Scale GitHub issue #93).
- **2026-06-30 — Anthropic billing transition context**: export controls on Fable lifted the same day Sonnet 5 became default — two unrelated but same-day events the brief conflated in one sentence; separated here.
- **2026-07-02 — theaiinsider.tech roundup dates the Omni Flash / Nano Banana 2 Lite developer expansion**: the strongest date anchor for the corrected July 1–2 Nano Banana 2 Lite launch (vs the false I/O attribution).
- **2026-07-10 — pureai.com's same-day launch report**: "one of the most competitive single days in the sector's history" — the primary-adjacent press record for the July 9 framing; ailearningguides.com's "arguably the busiest single day in the history of AI models" echoes it independently.
- **2026-07-12 — nt-executive-tech-office brief**: consolidates GPT-5.6 pricing/tiers — the secondary anchor for the Sol/Terra/Luna price cards in the figures table.
- **2026-08-13 — newindianexpress on Grok 4.6**: "cheaper rival to OpenAI and Anthropic" framing — the price-positioning angle for the $2/$6 card.
- **2026-09-03 — GPT-6 Astra (OpenAI) launched**: out of §1's scope per the brief's coverage list (the chronology verified here ends with GPT-5.6), but noted as the immediate successor benchmark: entered Terminal-Bench 4.0 at 58.18% in Codex at max effort on launch day; SWE-bench Verified 96.2% vs Opus 5 97.0% (aggregator). Recorded only as successor context.
- **2026-09-16 — Grok Imagine Image 2.0 via API** (`grok-imagine-image-2.0`, `/v1/images/generations`, `/v1/images/edits`, OpenAI-compatible SDKs): the "API access coming soon" gap from the Aug 7 launch closed ~5 weeks later.
- **2026-09-19 — Grok Voice Transcribe 2.0 shipped** (two days before Grok 4.7): the voice line continued iterating after STT 1.0 — Wave 2.1 record.
- **Undated-within-2026 — Meta's "Muse" branding lineage**: original Muse Spark (April, AA v4.0 52, free) → 1.1 (Jul 9, first paid API) → 1.2 (Aug 5) → 1.3 (Sep 2). The August 10 open-weights promise ("soon," permissive) was still unfulfilled as of Sept 22, 2026; Muse Glimmer (30B, Apache 2.0) released Aug 10 as the immediate open-weight artifact of the lineage.

### Cadence comparison — minor-version intervals in 2026

| Lab | Interval | Days |
|---|---|---|
| Anthropic Opus 4.6 → 4.7 | Feb 5 → Apr 16 | ~70 |
| Anthropic Opus 4.7 → 4.8 | Apr 16 → May 28 | ~42 (fastest Anthropic minor cadence) |
| Anthropic Opus 4.8 → Opus 5 | May 28 → Jul 24 | ~57 |
| Anthropic Fable 5 → Fable 5.1 | Jun 9 → Sep 1 | ~84 |
| xAI Grok 4.5 → 4.6 | Jul 9 → Aug 12 | ~34 |
| xAI Grok 4.6 → 4.7 | Aug 12 → Sep 21 | ~40 |
| Meta Muse Spark 1.1 → 1.2 | Jul 9 → Aug 5 | ~27 |
| Meta Muse Spark 1.2 → 1.3 | Aug 5 → Sep 2 | ~28 |
| OpenAI GPT-5.5 GA → GPT-5.6 GA | May → Jul 9 | ~2 months (estimated) |
- **"Thinking/Instant" hybrid templates became the default product shape**: GPT-5.5's GA shipped as a Thinking/Instant hybrid; Mistral Small 4 (recorded in the debunk context) and Muse Spark 1.2's "mandatory reasoning" confirm the pattern — 2026 flagships are reasoning-first models with an effort dial, not separate "thinking" SKUs. The o-series sunset (Aug 26, 2026) completed the same absorption at OpenAI.
- **Naming hygiene was a 2026 failure mode worth encoding**: "SpaceXAI" vs xAI (press drift, no corporate confirmation); "Mistral Medium 3.1" vs 3.5 (mislabeled August 2025 refresh); "Gemini 3.1 Pro 2M" (single-source tier confusion); "Nano Banana 2 Lite" at I/O (never happened). The consolidation kept the vendor/press-attested name and flagged the variant in Consolidation notes — downstream queries should prefer these canonical forms.
- **The export-control episode created a pricing event, not just a legal one**: Fable 5's July 1 restoration behind stricter classifiers with auto-fallback to Opus 4.8 (rather than hard refusals) is the first shipped example of a tiered-safety product path — flagged traffic degrades to the cheaper model. The July 7 metered-credits move and July 20 plan split (Max/premium keep Fable 5 included; Pro → credits) then priced safety overhead into subscriptions.

## Sources and URLs

