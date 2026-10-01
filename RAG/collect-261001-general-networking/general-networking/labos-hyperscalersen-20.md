---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-20
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Anthropic", "EU", "Google", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "xAI"]
dates: ["2025-06", "2026-04-08", "2026-05", "2026-07-09", "2026-08", "2026-08-05", "2026-09", "2026-09-02"]
keywords: ["agent", "agentic", "amd", "apache", "astra", "benchmark", "capex", "chatgpt", "claude", "compute", "consumer", "cyber"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [661, 729]
sha256: 17aba7f569078ff8a355fa8b50161aaccff524e7884bad0529fb420a2534607e
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

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

### 4.3 Muse Glimmer 30B — open weights (Apache 2.0)
- **Muse Glimmer 30B** released under **Apache 2.0** — Meta's open-weights offering in the window, continuing the Llama-era open strategy under the new Muse brand. [secondary: Track B]
- 30B parameters; positioned for on-device / edge / fine-tuning use cases. [secondary]

### 4.4 Capex, compute, MTIA
- **Meta capex** remained at historic highs through 2026 to fund AI data-center buildout (Track B; quarterly figures vary — pull from earnings for precision).
- **MTIA "Iris"** — Meta's next-gen AI training/inference chip program referenced in Track B (naming from roadmaps — [secondary]).
- Meta's compute fleet expansion (hundreds of thousands of GPUs) continued; **Nebius–Meta orders** noted in Track D (see §13).
- **Meta–Nebius**: Nebius reported Meta as a customer / order flow in 2026 (see §13.5). [secondary]

### 4.5 Products & partnerships
- **Meta AI** assistant continued rollout across Facebook, Instagram, WhatsApp, Messenger. [secondary]
- **Ray-Ban Meta / Oakley Meta** smart glasses — AI wearable line; 2026 iterations with deeper Muse integration (Track B).
- **Meta–Scale AI**: Meta's **$14.3B investment in Scale AI (June 2025, pre-window)** — Alexandr Wang joined Meta as Chief AI Officer; 2026 execution continued under the new **Meta Superintelligence Labs** structure. [secondary: Track B]
- **Meta–Nvidia**: continued GPU purchasing at scale. [secondary]
- **Meta–AMD / MTIA**: heterogeneous silicon strategy. [secondary]

### 4.6 Controversies & regulatory
- **EU AI Act / DSA**: Meta's AI products in scope; compliance posture through 2026 (Track B).
- **Llama-era data lawsuits** continued through 2026 (Track B — specific filings to verify).
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — Meta signatory status not confirmed in sources returned (gap).

### 4.7 Meta benchmark snapshot (as reported)

| Model | Benchmark | Score | Provenance |
|---|---|---|---|
| Muse Spark 1.3 | DeepSWE v1.1 | 75.4% | [vendor-reported] |
| GPT-6 Astra (for comparison) | DeepSWE v1.1 | 74.1% | [vendor-reported] |

---
### 4.8 Muse family — detailed release record (from Track B)

Meta Superintelligence Labs (MSL), led by Chief AI Officer **Alexandr Wang** (ex-Scale AI, joined via the ~$14B Scale AI investment/deal in 2025), pivoted Meta from pure open-weights to a dual strategy: closed API models (Muse Spark) + open-weights local models (Muse Glimmer). This is Meta's **first-ever paid model API** — a deliberate strategic shift from the Llama tradition [secondary] https://mlq.ai/news/meta-launches-muse-spark-11-api-its-first-paid-ai-model-to-challenge-anthropic-and-openai/.

**Muse Spark 1.0 (April 8, 2026)**: first major Meta AI model since the Scale AI deal; announced by Mark Zuckerberg at the Meta AI event; no public developer API at launch [secondary — CNBC] https://www.cnbc.com/2026/04/08/meta-debuts-first-major-ai-model-since-14-billion-deal-to-bring-in-alexandr-wang.html. Reported benchmark framing (vendor-adjacent aggregators): 77.4% coding vs. top rivals, 59% on Terminal-Bench 2.0; Meta claims parity/superiority on agentic and frontend/visual tasks vs. Claude Opus 4.6, GPT-5.4/5.5, Gemini 3.1 Pro — treat individual numbers as [unverified]. Independent commentary notes strength in visual/UI reasoning, weaker on heavy backend programming [unverified/secondary].

