---
id: ai-industry-kb-2026/01-frontier-closed-model-releases/cross-check-fabricated-rumor-items-correctly-excluded-from-t
title: "Cross-check: fabricated-rumor items correctly excluded from the timeline"
domain: frontier-closed-model-releases
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Huawei", "Meta", "Microsoft", "Mistral", "OpenAI", "United States", "Xiaomi", "xAI"]
dates: ["2025-04", "2025-04-28", "2025-08", "2025-11-17", "2025-12", "2025-12-01", "2026-02", "2026-02-05", "2026-02-10", "2026-02-17", "2026-02-18", "2026-03-16", "2026-04", "2026-04-11", "2026-04-16", "2026-04-17", "2026-04-29", "2026-05-28", "2026-07-01", "2026-07-23", "2026-08-05", "2026-08-07", "2026-08-12", "2026-08-26"]
keywords: ["agent", "apache", "attribution", "benchmarks", "claude", "deepseek", "distribution", "fable 5", "gemini", "grok", "grok 4", "memory"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [297, 327]
section: "1. Frontier Closed-Model Releases"
sha256: 7d3ffd472674b5174314152d9af831e049ec81df8f8abfe7c4396f49d72d7a81
---

# Cross-check: fabricated-rumor items correctly excluded from the timeline

The claim that Mistral Small 4, Mistral Medium 3.1, Mistral Large 3, Grok 4.1, Grok 4.2, DeepSeek V3.2, Grok 4.6, Grok STT 1.0, Grok Imagine Image 2.0, and Meta Muse Spark 1.2 all shipped in a grouped wave around April 11, 2026 is **FALSE — every element is contradicted, misdated, or mislabeled**:

| Brief claim | Real fact | Correct date |
|---|---|---|
| "Apr 11: Mistral Small 4" | Small 4 (119B/6B active, Apache 2.0, 256K ctx) shipped March 16, 2026 | 2026-03-16 |
| "Apr 11: Mistral Large 3" | Large 3 (675B/41B active, Apache 2.0) is a December 2025 MoE flagship | 2025-12 |
| "Apr 11: Mistral Medium 3.1" | Mislabeled — Medium 3.1 is an August 2025 refresh; the real April model is **Medium 3.5** (128B dense, $1.50/$7.50, modified MIT, SWE-Bench Verified 77.6%) | 2026-04-29/30 |
| "Apr 11: Grok 4.1" | Grok 4.1 shipped November 17, 2025 | 2025-11-17 |
| "Apr 11: Grok 4.2" | Grok 4.20 Beta launched as a public beta February 17–18, 2026; no April launch existed | 2026-02-17/18 |
| "Apr 11: DeepSeek V3.2" | V3.2 shipped December 1, 2025 (V3.2-Exp Sept 29, 2025) | 2025-12-01 |
| "Apr 11: Grok STT 1.0" | Unversioned STT/TTS APIs April 17–19, 2026; versioned 1.0 on July 23, 2026 | 2026-04-17/19 → 2026-07-23 |
| "Apr 11: Grok Imagine Image 2.0" | Image 2.0 shipped August 7, 2026 | 2026-08-07 |
| "Apr 11: Grok 4.6" | Grok 4.6 shipped August 12, 2026 (post-training on V9 base) | 2026-08-12 |
| "Apr 11: Meta Muse Spark 1.2" | Muse Spark 1.2 shipped August 5, 2026, co-launched with Muse Code | 2026-08-05 |
| Real April 2026 frontier events that fit the window | xAI STT/TTS APIs (Apr 17–19); Mistral Medium 3.5 (Apr 29–30); OpenAI Codex "for everything" (Apr 16); Microsoft Agent Framework v1.0 (Apr 2026) | Apr 2026 |

### Cross-check: fabricated-rumor items correctly excluded from the timeline

- **"Qwen3-30B-A3B" February 10, 2026 release — CONTRADICTED.** This is a Qwen3 April 2025 SKU (30.5B/3.3B active, catalog April 28, 2025); no February 2026 release exists. The brief attached the wrong name to the right week (the real mid-February event was Qwen3.5-397B-A17B). Dropped from the consolidated timeline.
- **DeepSeek V4 "first mention" mid-February 2026 — VERIFIED as pre-release reporting** (The Information: V4 around mid-February, pre/post Spring Festival; internal benchmarks said to surpass Claude/GPT on coding). The slip to the April 24 preview is attributed to a Huawei/Cambricon platform transposition plus last-minute architecture updates (Engram memory, DualPath). Fabricated February leaks (83.7% SWE-bench Verified / 99.4% AIME 2026 screenshots — Epoch AI confirmed fabricated; "Hunter Alpha" = a Xiaomi 1T model; "sealion-lite" NDA leak Feb 26) are documented but not carried as facts.
- **Galaxy S26 (Galaxy Unpacked, San Francisco Feb 25 local / Feb 26 Asia)** is a hardware launch (Snapdragon 8 Elite Gen 5, "The Next AI Phone" branding) — not a closed frontier model release; recorded here only because it was in the February brief, and belongs to a devices chapter, not §1.

- **2026-04-16 — OpenAI Codex "for everything"**: the real mid-April OpenAI frontier event (the sibling to the debunked wave) — positioned Codex as a general agent platform, not just a coding tool; covered in the agent-framework records.
- **2026-04 (undated) — Microsoft Agent Framework v1.0**: a real April framework release that sits alongside the model calendar; recorded here only to complete the April picture.
- **2026-01 (background) — Stanford 2026 AI Index (data current Feb 2026)**: Claude 4.5 Opus in high-reasoning mode at ~76.8% SWE-bench Verified — the saturation baseline that 2026's 88–97% figures must be read against.
- **Frontier release "default" mechanics converged**: Sonnet 4.6 (default Feb 17) → Sonnet 5 (default Jun 30) → Fable 5 (default Claude Code model Pro/Max from Jun 9). The "becomes default" event is now a dated, checkable milestone — this section records it as a first-class fact (Jun 30) rather than a marketing note.

- **2026-02-18 — Sonnet 4.6 press-day bundle**: the 365i.co.uk FAQ, NYU Shanghai recap, and tokenring/sandiego pieces all date February 18 (US) for what Anthropic shipped Feb 17 — a one-day press-cycle lag the consolidator should expect in secondary sourcing; primary-adjacent dates (gofinkle's 2026-02-05 Opus 4.6 post, maxritter's 2026-05-28 Opus 4.8 post) are the anchors.
- **2026-07-01 — withAssetRevalidation-style cache lessons aside, the Gemini developer-rollout of ~July 1–2 is itself anchored by theaiinsider.tech 2026/07/02 roundup and gigazine's 20260701 piece** — two independent date stamps against the false I/O attribution for Nano Banana 2 Lite and the Omni Flash dev rollout.
- **2026-08-26 — Claude in Chrome GA for paid plans (Aug 26, 2026)**: same-day as the o3 sunset; a distribution milestone for the Claude line rather than a model release — recorded for timeline completeness, not benchmarked here.

