---
id: labs-hyperscalers-2026/00-labs-hyperscalers/3-12-personnel-google-deepmind
title: "3.12 Personnel (Google/DeepMind)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google", "Meta", "Microsoft", "OpenAI", "xAI"]
dates: ["2026-04-08", "2026-05", "2026-06-18", "2026-06-19", "2026-07-09", "2026-08", "2026-08-05", "2026-09", "2026-09-02"]
keywords: ["agent", "agi", "apache", "astra", "backlog", "chatgpt", "claude", "consumer", "distribution", "gemini", "gpt-6", "grok"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [648, 693]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 8caeb747373f80e1cace766a09fba285c7c5666fd6aa1f15dea4de7dfbb178e0
---

# 3.12 Personnel (Google/DeepMind)

### 3.12 Personnel (Google/DeepMind)

- **David Silver** (AlphaGo/AlphaZero/AlphaStar RL lead) left DeepMind ~Feb 2026 after ~13 years → founded Ineffable Intelligence; $1.1B seed at $5.1B valuation [secondary].
- **Noam Shazeer** → **OpenAI**, announced June 18, 2026 — under two years after Google's ~$2.7B Character.AI licensing deal brought him back [independent].
- **John Jumper** (AlphaFold co-creator, 2024 Nobel Chemistry) → **Anthropic**, announced June 19, 2026; AlphaFold researchers Jonas Adler and Alexander Pritzel followed [independent/secondary].
- **Aug 5, 2026 — leadership overhaul**: Demis Hassabis stepped back from day-to-day CEO of Google DeepMind → **Chairman of Google DeepMind + newly created Chief Scientist of Alphabet**; **Koray Kavukcuoglu** (CTO) → SVP Google DeepMind, owning Gemini model development, frontier research, Gemini app and developer teams, reporting to Pichai; Jeff Dean, Sanjay Ghemawat, Oriol Vinyals, Quoc Le left to found **Discovery Loop** (independent public-benefit corp; Google as founding investor + Cloud partner). Alphabet shares fell ~4–5.5% on the news [independent] https://www.businesstimes.com.sg/international/global/seismic-shift-google-ai-veterans-depart-biggest-reorganisation-ai-efforts.
- Talent-market data (Zeki Data via Fortune, Aug 2026): DeepMind's share of EMEA research/advanced-engineering hires fell from 49% (2022–23) to 18.6% (2025–26); arrivals-to-departures from ~12:1 (mid-2023) to ~2:1 (Q3 2026); of DeepMind leavers in past 12 months: 25% → Anthropic, 21% → Meta, 14% → OpenAI [secondary]. Of 29 named authors on the 2021 AlphaFold2 Nature paper, 13 have left DeepMind [secondary].
- Hassabis framing: "we have arrived at a pivotal moment in human history," AGI "close at hand"; he continues as CEO of Isomorphic Labs alongside the Alphabet Chief Scientist role [secondary].

### 3.13 Alphabet adoption / usage figures (vendor-reported unless noted)

Gemini API 16B tokens/min direct customer use (Q1 2026); 330 customers >1T tokens TTM; 35 >10T tokens TTM; gen-AI product revenue +800% YoY (Q1 2026); Gemini Enterprise 40% QoQ paid MAU growth; Workspace AI 11M+ paid seats (+76% YoY); Gemini app 950M MAU; Gemma 900M+ downloads; AI Mode in Search 1B MAU (May 2026); Veo 3 40M+ videos in first 7 weeks; Google Cloud backlog $240B (Apr) → $462B (Q1) → $514B (Q2); Cloud revenue +48% (Q4 2025) → +63% (Q1) → +82% (Q2) [vendor-reported].

### 3.7 Gemini Enterprise Agent Platform + DeepMind research frontier (from Track B)

**Gemini Enterprise Agent Platform:**
- Google's enterprise agent offering launched during the window as the counterpart to ChatGPT Work / Claude Cowork plays [secondary: Track B].
- Distribution note: **Grok 4.6+ became available on the Gemini Enterprise Agent Platform** from August 2026 — Google distributing a competitor's model on its own enterprise agent surface, a notable multi-model-routing signal [secondary].
- Positioned alongside Vertex AI and Antigravity as Google's three enterprise AI surfaces [secondary].

**DeepMind research & product frontier (2026):**
- **Project Astra** (universal AI assistant) continued development through 2026 [secondary: Track B].
- **Gemini Robotics** program ongoing [secondary: Track B].
- **Gemini in Search:** AI Mode expanded to ~200 countries and 98 languages (no subscription) by the May 2026 I/O keynote; AI Overviews expansion continued [secondary].
- **Nano Banana lineage:** image-generation models tracked alongside the 3.5 Flash launch; MAI-Image-2.5 claimed to surpass "Nano Banana Pro" at Microsoft Build (Microsoft-reported) [secondary/vendor-reported].
- **Gemini CLI** and open-source developer tooling continued through 2026 [secondary: Track B].
- **Gemini for Education / classroom tools** noted in Track B [secondary].
- **Android XR / Gemini on-device** privacy discussions (minor) [secondary: Track B].

## §4 — META

> **2026 at a glance — Meta:** no Llama 5 (explicitly unverified); the Muse brand took over — Spark 1.0 (Apr 8) → 1.1 API (Jul 9, Meta's first paid API) → 1.2 (Aug 5) → 1.3 (Sep 2), plus open-weights Glimmer 30B (Aug 10, Apache 2.0); Meta Superintelligence Labs under Alexandr Wang executed the post-Scale-AI strategy.

### 4.1 The Llama 5 question — explicitly unverified as of 22 Sept 2026
- **No first-party evidence of a Llama 5 release was found as of 22 September 2026.** Meta's public model lineup in the window remained **Muse Spark 1.0 → 1.3** (frontier) and **Muse Glimmer 30B** (open weights). Any claim of a "Llama 5" launch in Feb–Sep 2026 should be treated as **[unverified]** unless a Meta source is produced. [research finding]
- This marks a deliberate branding pivot: Meta's frontier models now carry the **"Muse"** name (Muse Spark, Muse Glimmer), distancing from the Llama lineage that defined 2023–2025. [secondary]

### 4.2 Muse Spark 1.0 → 1.3 (frontier releases, 2026 — corrected dates; full record in §4.8)
- **Muse Spark 1.0** — **April 8, 2026** (corrected; announced by Zuckerberg at Meta AI event; no public API at launch). [secondary: Track B]
- **Muse Spark 1.1** — **July 9, 2026** (corrected; API launch — Meta's first paid model API). [secondary: Track B]
- **Muse Spark 1.2** — **August 5, 2026** (corrected). [secondary: Track B]
- **Muse Spark 1.3** — **September 2, 2026** (corrected; specs largely unverified):
  - **DeepSWE v1.1 75.4%** — beating GPT-6 Astra's 74.1% on Meta's own-reported comparison (vendor-reported, scaffolding not disclosed). [vendor-reported via press]
  - Positioned as Meta's coding/agent flagship. [secondary]
- **Muse Spark** powers Meta's consumer AI surfaces (Meta AI assistant) and is offered to developers. [secondary]

