---
id: labs-hyperscalers-2026/00-labs-hyperscalers/2-7-shared-cross-cutting-themes-anthropic-vs-openai
title: "2.7 Shared / cross-cutting themes (Anthropic vs OpenAI)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["Anthropic", "EU", "Google", "Hugging Face", "OpenAI"]
dates: ["2026-02-19", "2026-03", "2026-03-31", "2026-04-22", "2026-05-19", "2026-06", "2026-07-21", "2026-09"]
keywords: ["agent", "agentic", "agents", "astra", "benchmark", "capex", "chatgpt", "claude", "compute", "export controls", "fable 5", "funding"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [494, 543]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: f96627436e401a123ca83085cf6c1ce9599b69909517d0a9fce97611fa4c3fc5
---

# 2.7 Shared / cross-cutting themes (Anthropic vs OpenAI)

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

### 3.2 Revenue, users, market position
- Gemini app usage and market share data tracked: Gemini **27.7% of mobile AI-app share (Mar 2026)** vs ChatGPT 46.4%, Claude 10.3%. [independent: SmashingApps infographic] https://www.smashingapps.com/ai-market-share-2026-infographic/
- Google continued reporting strong AI-driven Cloud growth through 2026 (Track B).
- **Anthropic investment**: Google committed **additional $10B to Anthropic (Sept 2026)** as part of Anthropic's Series H (May 28) extension round — Google's total Anthropic investment ~$20B+ (pre-window $3B+ plus extension). [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md ; https://datafloq.com/openai-set-the-ai-valuation-record-in-march-anthropic-broke-it-by-may/
- Note: the Google–Anthropic cloud/compute relationship deepened in 2026 (Anthropic runs on Google TPUs in part; see §1.5). [secondary]

### 3.3 Infrastructure & silicon
- **Ironwood TPU** — **General Availability March 31, 2026** (corrected; full record in §3.6)
- **TPU 8t/8i** — announced **April 22, 2026** (corrected; full record in §3.6)
- **Alphabet capex** remained elevated through 2026 to fund TPU + data-center buildout (Track B; specific quarterly figures to be pulled from earnings — gap noted).
- **Project Suncatcher** (Google's orbital data-center concept) noted in Track B as a long-horizon infrastructure bet (flag: conceptual/early).

