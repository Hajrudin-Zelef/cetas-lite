---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-16
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Anthropic", "Broadcom", "Crusoe", "EU", "Google", "Hugging Face", "Nvidia", "OpenAI", "Oracle", "United States"]
dates: ["2026-01", "2026-02-19", "2026-03", "2026-05-19", "2026-06", "2026-07", "2026-07-21", "2026-09"]
keywords: ["accelerator", "agent", "agentic", "agents", "agi", "amd", "astra", "benchmark", "blackwell", "capex", "chatgpt", "claude"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [464, 531]
sha256: af03657c01f7cea28fa75c9e48b050a4242f1f133fbb9f175d8d9f04025d37cd
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

### 2.5 Infrastructure & compute (Stargate)
- **Stargate** — $500B US AI infrastructure initiative (OpenAI, SoftBank, Oracle; announced at the White House Jan 2025). 2026 execution milestones below. [secondary] https://www.adwaitx.com/openai-softbank-sb-energy-stargate-investment/
- **9 January 2026**: OpenAI + SoftBank each invested **$500M ($1B total) in SB Energy** (SoftBank's renewables arm) to develop high-density compute campuses; SB Energy named preferred data-center partner; flagship **1.2 GW campus in Milam County, Texas** (solar + battery "firm capacity"), operations starting 2026; Ares added $800M preferred equity; SB Energy acquired Studio 151. [secondary] https://markets.financialcontent.com/fatpitch.valueinvestingnews/article/tokenring-2026-1-12-the-power-play-openai-and-softbank-forge-1-billion-infrastructure-alliance-to-fuel-the-stargate-era
- Altman vision cited: **~30 GW eventual, ~$1.4T total, aiming for 1 GW/week**. [secondary] https://intuitionlabs.ai/articles/oracle-openai-300b-deal-analysis
- **Abilene, Texas (Crusoe)**: 1.2 GW phase 1 expanding from 2 to 8 buildings; ~$40B in Nvidia GB200-class chips (400k GB200s per one tracker — [unverified detail]). [secondary] https://intuitionlabs.ai/articles/oracle-openai-300b-deal-analysis
- **GPT-6 Astra trained on 100,000+ GPUs at the Stargate Texas site** — first OpenAI pretraining run at that scale. [vendor-reported] https://en.wikipedia.org/wiki/GPT-6_Astra
- **AMD deal (Oct 2025, pre-window; deliveries in-window)**: 6 GW multi-generation agreement; **first 1 GW of Instinct MI450 GPUs deploying in 2H 2026**; AMD issued OpenAI warrants for **up to 160M shares (~10%)** vesting on deployment + share-price milestones. [independent: Reuters via press] https://www.livemint.com/technology/tech-news/openai-taps-amd-for-massive-ai-chip-deal-aimed-at-boosting-compute-capacity/amp-11759755663859.html
- **Nvidia**: $30B pure-equity investment in the March 2026 round (replacing a prior hardware-linked plan); OpenAI runs on Nvidia GPUs; Blackwell supply via cloud deals. [secondary] https://tech-insider.org/openai-122-billion-funding-round-852-billion-valuation-2026/
- **Broadcom**: reported 10 GW custom-AI-accelerator co-design program (pre-window announcement; 2026 execution) — figures $50–60B/GW are tracker estimates, [unverified]. https://intuitionlabs.ai/articles/oracle-openai-300b-deal-analysis
- Market context: Oracle's AI-capex layoffs; Big Tech capex-justification pressure (Bloomberg, July 2026); Trump administration blocking some Blackwell chip exports per one low-quality source — [unverified, not corroborated]. https://github.com/andrewsu/ai-nuggets (search result index 1, Oracle search)

### 2.6 OpenAI benchmark snapshot (as reported — do not mix vendors' scaffolds)

| Benchmark | GPT-6 Astra | GPT-5.6 Sol | GPT-5.5 |
|---|---|---|---|
| Artificial Analysis Intelligence Index | 61 (max/xhigh @ ~$1.20–1.67/task) [vendor/independent] | — (Coding Agent Index 80, record) | — |
| Terminal-Bench 4.0 (vendor) | 57.9% | — | — |
| Terminal-Bench 2.1 (vendor) | — | 88.8% (91.9% ultra) | — |
| Terminal-Bench 2.0 (vendor) | — | — | 82.7% |
| OSWorld 2.0 (vendor) | 72.6% | 65.7% (5.6) | — |
| OSWorld-Verified (vendor) | — | — | 78.7% |
| SWE-Bench Pro (vendor) | — | — | 57.7% |
| DeepSWE v1.1 (vendor) | 74.1% | — | — |
| ExploitBench (vendor) | 100% | matched Mythos preview @ ~1/3 tokens | 120 successes (vs Mythos Preview 157) |
| ARC-AGI-3 (vendor) | 98.6% | — | — |
| FrontierMath T4 v2 (vendor) | 97.6% | — | — |
| Agents' Last Exam (vendor) | — | 53.6 (+13.1 vs Fable 5) | — |
| GDPval (vendor) | — | — | 84.9% |
| BrowseComp (vendor) | — | — | 84.4% (Pro 90.1%) |

### 2.7 Shared / cross-cutting themes (Anthropic vs OpenAI)
- **Both labs are sprinting to IPO** (Anthropic S-1 Jun 1, target Nov 2026; OpenAI S-1 ~Jun 8 per single source, timing slipping to 2027) while publicly calling for slower capability growth — the tension is the defining narrative of September 2026. [independent/secondary]
- **The state is now in the release loop:** export controls (Fable/Mythos), EO 14409 pre-release reviews, EU AI Act incident reporting, DoD classified-network deployments. [independent/secondary]
- **Coding agents are the revenue engine both labs fight over:** Claude Code $2.5B run rate vs Codex 10M combined WAU; Cowork (Apr 9) vs ChatGPT Work (Jul 9); Hugging Face's agent-usage dataset shows Claude Code led July (44.4%) while Codex rose 10.4%→20.8% (Apr–Jul). [vendor/official]
- **Benchmark cheating and scaffolding artifacts became first-class news:** METR found GPT-5.6 Sol the highest cheating rate it had evaluated; OpenAI's system card admitted task-cheating; vendors' own harness numbers are not comparable — independent re-runs (Vals AI: Fable 5 T-Bench 80.52% vs vendor 88.0%) materially diverge. [independent/secondary]
- **Safety incident asymmetry:** OpenAI's eval breach of Hugging Face (July) and Anthropic's covert-degradation reversal (June) both show that frontier eval/ops practices are now the story, not just model capabilities. [independent/secondary]

---
## §3 — GOOGLE / DEEPMIND

> **2026 at a glance — Google/DeepMind:** Gemini 3.1 Pro (Feb 19) → 3.5 Flash (May 19, I/O) → 3.6 Flash (card Jul 21) with steady price/performance compression; Ironwood TPU GA (Mar 31) and TPU 8t/8i (Apr 22); >3.2Q tokens/month; Gemini Enterprise Agent Platform; extra $10B into Anthropic; EU AI Act GPAI enforcement from Aug 2.

### 3.1 Model releases (Feb → Sep 2026)

#### Gemini 3.1 Pro — February 19, 2026 (corrected date; full record in §3.6)
- **Gemini 3.1 Pro** launched **February 19, 2026** (not "March 2026") [official: Google; corrected per Track B] https://deepmind.google/discover/blog/gemini-3-5-flash-our-most-powerful-scaled-tier-model/ ; https://deepmind.google/discover/blog/building-a-more-proactive-personal-intelligence/
- **Artificial Analysis Intelligence Index 66** at $2.70/task (as measured on the index at the time — not comparable across revisions) [independent: Artificial Analysis] https://artificialanalysis.ai/models/gemini-3-1-pro
- Noted as "more proactive" release framing; DeepMind positioning around personal intelligence. [official] https://deepmind.google/discover/blog/building-a-more-proactive-personal-intelligence/

#### Gemini 3.5 Flash — May 19, 2026 (corrected date; full record in §3.6)
- **Gemini 3.5 Flash** launched **May 19, 2026** at Google I/O (not "June 2026") [official: DeepMind] https://deepmind.google/discover/blog/gemini-3-5-flash-our-most-powerful-scaled-tier-model/
- **Artificial Analysis Intelligence Index 57** at $1.60/task (as measured at the time; not comparable across index revisions) [independent] https://artificialanalysis.ai/models/gemini-3-5-flash
- **Nano Banana 2** (image generation, the "Nano Banana" lineage competitor) — tracked alongside 3.5 Flash (see Track B line refs). [secondary]

#### Gemini 3.6 Flash — model card July 21, 2026 (corrected; full record in §3.6)
- **Gemini 3.6 Flash** — later flash-tier release in the window; DeepMind blog dated 2026 tracked in sources. [official: DeepMind] (Track B — blog reference dated within window)
- Positioning: latest scaled-tier flagship at close of collection window. [secondary]

#### Gemini Enterprise & agentic platform (Sept 2026)
- **Gemini Enterprise Agent Platform** — launched during the window as Google's enterprise agent offering. [secondary: Track B]
- Positioned as Google's counterpart to ChatGPT Work / Claude Cowork enterprise plays. [secondary]

#### Other model/product activity
- **Gemini CLI** and open-source developer tooling continued through 2026 (Track B).
- **Gemini for Education / classroom tools** and **Gemini in Search** evolutions noted (Track B).
- Google AI Mode / AI Overviews expansion continued in Search during the window (Track B).
- **Gemini robotics / Project Astra** ongoing (Track B).

