---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/overview
title: "§14. Google: Gemini and Gemma"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Cohere", "Google", "Meta", "OpenAI", "xAI"]
dates: ["2025-05", "2025-12-04", "2026-01-13", "2026-01-15", "2026-02", "2026-02-19", "2026-03-31", "2026-04", "2026-04-02", "2026-04-16", "2026-05-19", "2026-06-03", "2026-06-05", "2026-06-10", "2026-06-30", "2026-08-13", "2026-08-27", "2026-12-31"]
keywords: ["gemini", "agentic", "apache", "astra", "attribution", "benchmark", "claude", "cohere", "consumer", "gemini 3.8", "gpt-6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6793, 6863]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: bdd6bd650fedb99327f4b5566f1f58100426e849e3f151c1e66c3bc10a415184
---

# §14. Google: Gemini and Gemma

Keywords: Google, Gemini, Gemini 3 Deep Think, Gemini 3.1 Pro, Gemini 3.1 Pro Preview, Gemini Omni, Gemini Omni Flash, Gemini Omni 1.1 Flash, Veo 3, Nano Banana 2, Genie 3, Project Genie, Gemma 4, Gemma 4-MTP, Gemma 4 12B Unified, MedGemma 1.5, TranslateGemma, QAT, DiffusionGemma, FACTS Suite, Google AI Ultra

## Summary
- **Gemini 2.0 Ultra never existed** — the corpus's phantom-model correction stands; do not cite it.
- Gemini 3 Deep Think (Dec 2025) → Gemini 3.1 Pro Preview (Feb 2026) → **Gemini Omni**, announced at Google I/O 2026-05-19 [SECONDARY], is the 2026 Gemini arc; Gemini 3.5 Pro stayed internal/unreleased, and "Omni Pro" is teased-only **[UNVERIFIED]**.
- Gemini pricing is the undercut play: 3.1 Pro at $2/$12 ≤200K with a standing **$4/$18 surcharge above 200K** [VENDOR]; 3.7/3.8 Flash at $0.75/$3.75 intro through 2026-12-31 (doubles Jan 1, 2027) [SECONDARY].
- The **Gemma 4** family (weights 2026-03-31, announced 2026-04-02) is Google's first Gemma generation off the custom Gemma Terms — **Apache 2.0** across E2B, E4B, 26B-A4B, and 31B [SECONDARY].
- Media models anchor the consumer story: **Veo 3** (May 2025) with native synchronized audio, the **Nano Banana 2** image tier (2 / 2 Lite / Pro), and **Genie 3 / Project Genie** [SECONDARY].
- Google DeepMind's **FACTS Benchmark Suite** (initial run Dec 2025) puts Gemini 3 Pro at 68.8 with a **~70% factuality ceiling no model has broken** [VENDOR/SECONDARY] — the headline grounding finding of the wave.
- Gemma corrections to hold: **Gemma 4 12B Unified** (2026-06-03) was **not** an I/O highlight; there is **no Gemma TTS** and **no Gemma 4.5**; TranslateGemma ships under Gemma Terms, not Apache.

## Key dated facts
### Gemini 3 Deep Think — the 2025 anchor
- **2025-12-04** — Gemini 3 Deep Think released [SECONDARY].
- Bundled in **Google AI Ultra at $249.99/month** [SECONDARY] — the subscription wrapper for Google's reasoning-tier access.

### Gemini 3.1 Pro Preview — February 2026
- **2026-02-19** — **Gemini 3.1 Pro Preview** released [SECONDARY]; there is **no separate "Preview Vision"** variant — the name correction stands.
- Pricing [VENDOR]: **$2.00/$12.00** per million input/output **at or below 200K tokens**; the moment a single prompt exceeds 200K, pricing steps to **$4.00/$18.00** — a standing surcharge, not a repricing.
- LMArena text (Sept 2026): Gemini 3.1 Pro Preview at ~1500 [SECONDARY].
- Context-window pricing tiers are standard in 2026; Gemini's 200K cliff is its instance (Grok 4.7 doubles at 200K; GPT-6 Astra cliffs at 272K) [DIRECTIONAL].

### Gemini 3.5 Pro — internal only
- Gemini 3.5 Pro is **internal/unreleased only** [SECONDARY] — no public API, no weights; do not cite it as a shipped model.

### Gemini Omni — the I/O 2026 generation
- **2026-05-19** — **Gemini Omni** announced at Google I/O [SECONDARY].
- **2026-06-30** — **Omni Flash preview** [SECONDARY].
- **2026-08-27** — **Omni 1.1 Flash**: 40-second scene extension, first-last-frame control, 4K upscaling [SECONDARY].
- **Omni Pro** is **teased-only [UNVERIFIED]** — no launch, no specs; do not cite it as a model.

### Media models — Veo 3, Nano Banana 2, Genie 3
- **Veo 3** (May 2025) shipped with **native synchronized audio** [SECONDARY] — video generation with audio in a single pass.
- **Nano Banana 2** ships in three tiers — **2 / 2 Lite / Pro** [SECONDARY].
- **Genie 3** and **Project Genie** are Google's world-model / interactive-generation track [SECONDARY].

### MedGemma 1.5 and TranslateGemma
- **2026-01-13** — **MedGemma 1.5** released [SECONDARY].
- **2026-01-15** — **TranslateGemma** released under **Gemma Terms, not Apache** [SECONDARY] — the domain models did not follow Gemma 4 onto Apache 2.0.

### Gemma 4 — Apache 2.0, March/April 2026
- **2026-03-31** — Gemma 4 **weights** released; **2026-04-02** — public announcement [SECONDARY].
- Sizes: **E2B, E4B, 26B-A4B, 31B** — all **Apache 2.0** [SECONDARY].
- This is the **first Gemma generation off the custom Gemma Terms** — a 2026 Apache-2.0 first alongside Meta's Glimmer 30B, Cohere's Command A+, and OpenAI's gpt-oss (see §21).
- **2026-04-16** — **Gemma 4-MTP** (multi-token prediction): ~3× speedup [SECONDARY].
- **2026-06-03** — **Gemma 4 12B Unified** [SECONDARY] — this was **not an I/O highlight**; the I/O-attribution claim is corrected.
- **2026-06-05** — **QAT** (quantization-aware training) release [SECONDARY].
- **2026-06-10** — **DiffusionGemma** [SECONDARY].
- **There is no Gemma TTS and no Gemma 4.5** — both are dropped, not carried as models.
- HF download scale (Sept 2026, [VENDOR] via Alibaba-cited comparison): Qwen downloaded **2.045B times in 7 months of 2026 — roughly 4.9× Google's and 9× Meta's** [SECONDARY] — the permissive-license scale gap.

### The flash-pricing tier
- **Gemini 3.7/3.8 Flash**: **$0.75/$3.75** input/output, cache $0.075, introductory pricing **effective 2026-08-13, doubling Jan 1, 2027** [SECONDARY] — promo pricing with a stated expiry, priced to undercut.
- **Gemini 3.5 Flash-Lite**: **$0.30/$2.50** [SECONDARY].
- Reseller arbitrage: Gemini 3.7 Flash listed at ~91% off retail by one reseller [SECONDARY].
- Terminal-Bench 4.0: Gemini 3.8 Flash at **19.1%±3.4** [SECONDARY] — a weak agentic-coding showing for the flash tier.

### FACTS — the grounding benchmark
- Google DeepMind's **FACTS Benchmark Suite** (announced early 2026; initial leaderboard run **Dec 2025**): four dimensions — Parametric, Search, Multimodal, Grounding v2; 3,513 public examples plus a private Kaggle holdout; multi-judge design (models rate their own outputs 3.23pp higher on average — the multi-judge design is load-bearing) [VENDOR/SECONDARY].
- Initial run: **Gemini 3 Pro 68.8** overall (Search 83.8, Parametric 76.4, Multimodal 46.1), then Gemini 2.5 Pro 62.1, GPT-5 61.8, Grok 4 53.6, Claude 4.5 Opus 51.3 [VENDOR/SECONDARY].
- **The ~70% factuality ceiling: no model has broken 70%** — the headline finding [SECONDARY]; for RAG builders the Search sub-score (83.8) is the procurement-relevant number, and Multimodal (~46–47% even at the top) is the caution flag.
- The original FACTS Grounding (Jan 2025) is a different instrument (1,719 examples) — **do not compare with the Suite's Grounding v2** (2,104 items, harder).

### Coalition posture
- Google is reported (community sources only) to have **signed the July 24 open-weights letter shortly after publication**, while absent from the original 25 — **[SECONDARY], community-sourced only**; the original-25 list is the verified fact.


### New verified facts — expansion

