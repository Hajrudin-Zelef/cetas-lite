---
id: ai-industry-kb-2026/01-frontier-closed-model-releases/overview
title: "1. Frontier Closed-Model Releases"
domain: frontier-closed-model-releases
role: deep-dive
task: model-release
actors: ["Anthropic", "DeepSeek", "Glasswing", "Google", "Meta", "Mistral", "OpenAI", "United States", "xAI"]
dates: ["2025-12", "2026-02", "2026-02-05", "2026-02-08", "2026-03-05", "2026-03-24", "2026-04-11", "2026-04-23", "2026-05", "2026-06-25", "2026-07-09", "2026-07-30", "2026-08-26", "2026-09-02", "2026-09-22", "2026-10-14"]
keywords: ["agents", "claude", "compute", "cost", "cybersecurity", "deepseek", "export controls", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [230, 253]
section: "1. Frontier Closed-Model Releases"
sha256: 54477578194e4af390d10f5c2f142f0de65d6c485560ef015bc2ed5c05a8411f
---

# 1. Frontier Closed-Model Releases
Keywords: GPT-5.5 Spud, GPT-5.6 Sol Terra Luna, Claude Opus 4.6, Claude Opus 4.8, Claude Fable 5, Claude Mythos 5, Claude Sonnet 4.6, Claude Sonnet 5, Claude Opus 5, Gemini 3.1 Pro, Gemini Omni Flash, Grok 4.5, Grok 4.6, Grok 4.7, Muse Spark 1.2, Grok Imagine Image 2.0, Grok STT 1.0, Artificial Analysis Intelligence Index, most competitive day July 9 2026, April 11 2026 debunk, US export controls frontier models, Project Glasswing

## Summary

The closed-frontier release year ran **February 2026 → September 22, 2026**, anchored by three labs' flagship lines: **OpenAI** (GPT-5.5 "Spud" GA May 2026; GPT-5.6 Sol/Terra/Luna GA July 9, 2026), **Anthropic** (Opus 4.6 → Opus 4.7 → Opus 4.8 → Fable 5/Mythos 5 → Sonnet 5 → Opus 5 → Fable 5.1/Mythos 5.1), and **xAI** (Grok 4.2 beta → Grok 4.5 → Grok 4.6 → Grok 4.7), with **Google** contributing Gemini 3.1 Pro (February 19) and Gemini Omni Flash (announced at Google I/O 2026, developer rollout ~July 1–2), and **Meta** shipping the Muse Spark lineage (original April → 1.1 July 9 → 1.2 August 5 → 1.3 September 2, 2026). **July 9, 2026 is independently documented as one of the most competitive single days in AI history**: GPT-5.6's three-tier family and Grok 4.5 went public the same day, with Claude Fable 5 freshly restored from its export-control suspension. The headline debunk of this section: **the "April 11, 2026 grouped release wave" (Mistral Small 4 / Medium 3.1 / Large 3, Grok 4.1 / 4.2, DeepSeek V3.2, Grok 4.6, Grok STT 1.0, Grok Imagine Image 2.0, Muse Spark 1.2) is FALSE** — every dated element of it is contradicted, misdated, or mislabeled; the real events scatter across December 2025, February, March, April 29–30, July 23, and August 7–12, 2026. Release cadence accelerated sharply: Anthropic shipped Opus 4.7 → 4.8 in 41 days, and Opus 4.6's launch-day was undercut 27 minutes later by GPT-5.3 Codex. The first **US export-control suspension of frontier models** (Fable 5 / Mythos 5, June 12 → July 1) is treated as an incident and lives in §17; only a one-line cross-reference is kept here.

## Key dated facts

### OpenAI

- **2026-03-05 — GPT-5.4 launches** (baseline before Spud), an unusually fast six weeks before the next pretraining window — context for the Spud claim.
- **2026-03-24 — "Spud" (GPT-5.5) pretraining completed.** Sam Altman confirmed to employees the model was "a few weeks" from release and described it publicly as "a very strong model that could really accelerate the economy." OpenAI president Greg Brockman, on the Big Technology podcast, called it "two years of research" with a "big model feel," explicitly framed as not incremental. Whether it would ship as GPT-5.5 or GPT-6 was unconfirmed (it depended on the performance leap over GPT-5.4); it shipped as GPT-5.5, not GPT-6. The Information first reported the "Spud" codename on March 24, consistent with the food-codename tradition (Strawberry → o1, Opal, Spud). Polymarket priced ~78% release probability by April 30 and >95% by June 30. [VENDOR] vendor-confirmation via Altman/Brockman statements; cadence narrative secondary-reported.
- **2026-04-23 — GPT-5.5 "Thinking Instant" released** (~3–6-week safety-evaluation window after Spud pretraining), tracking the prediction markets.
- **2026-05 (GA) — GPT-5.5 GA** with the "Thinking/Instant" hybrid template. AA Intelligence Index 60.2 [VENDOR/comm] before being overtaken by Opus 4.8 (61.4).
- **2026-06-25/26 — GPT-5.6 government-reviewed preview** with ~20 approved partners under a cybersecurity executive order requiring frontier labs to submit powerful models for review before public release. This **contradicts and supersedes** the Wave 2 file's "GPT-5.6 Sol (preview July 30, 2026)" — the July 30 date is discarded (possibly a later staged-rollout milestone or an error).
- **2026-07-09 — GPT-5.6 family GA: Sol, Terra, Luna.** Sol $5 ($0.50 cached) / $30 per M; Terra $2.50 ($0.25 cached) / $15; Luna $1 ($0.10 cached) / $6 — Luna was the cheapest model ever released by a major lab at the time. All ~1.05M-token context, 128K max output, vision. Sol is the flagship for the hardest problems: **91.9% Terminal-Bench 2.1 in Ultra mode** (highest publicly recorded at the time) [VENDOR]. Terra: near-GPT-5.5 performance at half the cost. Luna: only **41.3% long-context recall** — the price/quality trade-off is explicit. Sol adds Pro/Ultra compute modes on the same underlying model; Ultra coordinates 4–16 parallel reasoning agents. New API surface: per-generation effort normalization, reasoning_mode, reasoning_context, prompt_cache_options.
- **2026-08-26 — o3 sunset** (context).
- **2026-10-14 — GPT-5.5 retirement** [DIRECTIONAL] planned.
- **2026-02-05 — GPT-5.3 Codex dropped 27 minutes after Claude Opus 4.6's launch**, creating the launch-day "AI wars" narrative of the month (razpetel catalogue, 2026-02-08).
- [UNVERIFIED] One secondary source (humai.blog) claims OpenAI discontinued Sora video generation and redirected resources to Spud — single-source; carried with flag only.

### Anthropic

