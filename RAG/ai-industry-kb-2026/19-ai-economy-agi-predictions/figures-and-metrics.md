---
id: ai-industry-kb-2026/19-ai-economy-agi-predictions/figures-and-metrics
title: "Figures and metrics"
domain: ai-economy-agi-predictions
role: deep-dive
task: finance
actors: ["Anthropic", "ByteDance", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "United States", "xAI"]
dates: ["2026-01-29", "2026-04", "2026-04-26", "2026-05", "2026-05-21", "2026-05-24", "2026-05-31", "2026-06-16", "2026-06-30", "2026-09-24", "2028-03"]
keywords: ["acquisition", "agent", "agentic", "agents", "agi", "arr", "capex", "compute", "consumer", "cost", "datacenter", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [9540, 9626]
section: "19. AI Economy & AGI Predictions"
sha256: d7992df941b492682fcdfd12e395eb410cbce4f2f668f37c215dc788db24f585
---

# Figures and metrics

## Figures and metrics

> Strict separation: these are **attributed forecasts, survey results, and secondary-reported business figures** — not measured realized value of the "AI revolution." Every entry carries its provenance inline.

### AGI forecast spread (2026 vintage)

- Altman: AGI 2026–2028; automated AI researcher by March 2028. (forecast)
- Amodei: "powerful AI" late 2026–2027. (forecast)
- Hassabis: 2029 ("real possibility," I/O May 2026) / 2031–2034 (tracker) / 3–10 years (Sept 2026 interview). (forecast, internally inconsistent)
- Musk: smarter than any individual human by end-2026; all of humanity combined 2030–2031. (forecast)
- LeCun: current approaches insufficient — dissent, no date. (forecast/dissent)
- [UNVERIFIED] The 2agi.me and skynetcountdown.com trackers are community aggregations; individual entries inherit their tracking methodology.

### Agent-economy figures (forecasts and surveys)

| Source | Figure | Nature |
|---|---|---|
| Gartner | 40% of business software ships task-completing agents by end-2026 | Forecast (secondary) |
| Gartner | 40%+ of enterprise agentic projects canceled by 2027 | Forecast, warning half (secondary) |
| Grand View Research | $10.9B AI-agents market in 2026, 49.6% CAGR | Market forecast (secondary) |
| IDC | 1.3B AI agents in enterprise workflows by 2028 | Forecast (secondary) |
| McKinsey | 15–20% bank cost cuts, $700–800B net savings | Forecast (secondary) |
| KPMG | $50B global agentic-AI spend in 2025 | Survey/spend estimate (secondary) |
| Wolters Kluwer | 44% of finance teams deploy agentic AI in 2026 (vs 6% prior year) | Survey (secondary) |
| index.dev | 88% favorable ROI, 4.3× ROI, payback <12 months | Vendor-side survey [VENDOR] |
| WEF/PYMNTS (Foxconn+BCG) | 80% of decision workflows automated, ~$800M value | Company-reported [VENDOR], unaudited |
| MCP adoption (April 2026 survey) | 78% of enterprise AI teams with ≥1 MCP-backed agent in production | Survey (secondary) |
| AAIF registry (May 24, 2026) | 9,652 server records; 97M monthly SDK downloads (Mar 2026) | Directional ecosystem stats |

### Video-generation price points (per-second API economics, 2026)

The surviving video businesses all converged on per-second API or subscription monetization — the economic layer beneath the capability story:

- **Gemini Omni Flash** — $0.10/sec of 720p output (API GA 2026-06-30; $17.50/M video tokens; $1.50/M input tokens; 50% batch discount). Comparator: Veo 3.1 Standard $0.40/sec at 1080p; Sora 2 $0.10–$0.70/sec. (source: wave2 §1.4)
- **Kling 3.0** — Standard $0.168/s without audio / $0.252/s with audio; Pro $0.224/s / $0.336/s. (source: Novita AI integration post, April 2026)
- **Seedance 2.0 API** — roughly $0.14 per 15-second clip (commercial PR, unconfirmed). [UNVERIFIED]
- **Project Genie** — consumer access gated behind Google AI Ultra $249.99/mo (consumer beta 2026-01-29; expansion 2026-05-21 with Street View/Maps grounding). Simulation surface monetized as premium subscription, not per-generation. (source: pulse2.com; soravideo.art)
- **Adobe Firefly** — 30+ video models by April 2026; rents inference, doesn't train foundation models. The aggregation model sidesteps the training-compute bill entirely. (source: baditaflorin deep-research compilation)
- The pattern: nobody in 2026 tried to monetize frontier video generation as a standalone consumer app after Sora's economics became known. [DIRECTIONAL]

### Enterprise adoption percentages (survey layer)

- **44% of finance teams** deploy agentic AI in 2026, up from 6% the prior year (Wolters Kluwer) — one of the largest single-year adoption deltas reported in the 2026 agent-economy corpus.
- **78% of enterprise AI teams** with at least one MCP-backed agent in production (April 2026 survey) — the integration-standard adoption figure behind the agent-spend story.
- **28% of Fortune 500** have implemented MCP servers; 67% of CTOs expect MCP as their default agent-integration standard within a year (2026 survey figures, secondary). [DIRECTIONAL]
- These are executive-belief numbers (class b), not measured ROI — they sit between the forecast layer and the measured layer.

### Video-generation economics (the 2026 quality→economics pivot)

- Sora: ~$5.4B reported annual compute spend vs ~$2.1M lifetime revenue (CIOL, unverified methodology) → consumer product discontinued 2026-04-26, API 2026-09-24. Measured fact: the shutdowns (OpenAI Help Center, primary). The $ figures: reported claims.
- Kling: $240M reported ARR Dec 2025 (single secondary compilation) — alive and growing. [UNVERIFIED]
- Adobe Firefly: 30+ video models by April 2026 (aggregation model; rents inference, doesn't train) — alive, aggregating.
- Project Genie: gated behind $249.99/mo Google AI Ultra plan (Jan 2026 consumer beta; May 2026 expansion) — expensive-subscription monetization.
- Measured lesson: the only consumer video product that died in 2026 had the worst reported economics. "Technology leadership ≠ product viability" is now an evidenced claim, not a slogan.

### Coding-agent revenue evidence (the closest thing to measured agent-economy data)

- **Cognition / Devin** [VENDOR] — Series D: >$1B at $26B (May 31, 2026); Series E: $2B at $48B (Sep 8, 2026); run-rate revenue $492M (May) → ~$900M (Sep), +83% in ~4 months. Enterprise customers claimed: Citi, Mercedes-Benz, Goldman Sachs, Dell, Santander, US Army/Navy, NVIDIA, GE Aerospace. Mercedes-Benz case: 8-month legacy modernization → 8 days (vendor). (sources: pulse2.com; en.cryptonomist.ch)
- **Cursor (Anysphere)** — $100M ARR (Jan 2025) → $500M (Jun 2025) → $1B (Nov 2025) → $2B (Feb 2026) → ~$4B annualized revenue (Jun 2026, ~$2.6B enterprise B2B); $2.3B Series D (Nov 2025) at $29.3B; ~2M+ DAU; 64–67% of Fortune 500 (Bloomberg-sourced). (source: getpanto.ai)
- **Replit** — ~$150M ARR (Sep 2025) → ~$525M (Sacra estimate, Apr 2026 — unconfirmed); $9B valuation on $400M Series D (Mar 11, 2026); 50M+ registered users; Agent 4 (Mar 2026) with parallel subagents. (source: wave2 §5.2)
- Caveats that travel with these numbers: all revenue figures are vendor-reported and unaudited; "annualized run-rate" annualizes a growth month; the Cursor $60B acquisition (June 16, 2026) is reported but unclosed with a disputed acquirer. Even so, these are actual invoices, not analyst models — they sit in the measured column with [VENDOR] labels.
- The ratio that matters: vendor-claimed agentic-product revenue (Devin ~$900M, Cursor ~$4B annualized, Replit ~$525M) is small against the forecast layer (KPMG's $50B agentic-AI spend, McKinsey's $700–800B bank savings) — but it is the *only* measured money in the agent economy. The forecast layer and the invoice layer are separated by roughly an order of magnitude. [DIRECTIONAL]

### Capex-cycle markers

- Training/inference spend split forecast: 80/20 training→inference inverting to 80/20 inference (Lenovo CEO, CES 2026). (forecast)
- Sora compute reallocation: official reason for shutdown ("side quests" memo) — measured official statement; directional support for inference-reallocation capex logic.
- Coding-agent revenue run-rates (vendor-reported, unaudited): Devin $492M (May) → ~$900M (Sep); Cursor ~$4B annualized (Jun). [VENDOR]

## Main actors

- **Sam Altman (OpenAI)** — the 2026–2028 AGI forecast holder; automated-AI-researcher-by-March-2028 claim; "delighted to be wrong" on the jobs apocalypse (May 2026); made the Sora shutdown call (compute reallocation to higher-priority products).
- **Dario Amodei (Anthropic)** — "powerful AI" late 2026–2027 ("a country of geniuses in a datacenter"); Jevons Paradox invocation at the JPMorgan/Dimon briefing (May 2026); "AI won't destroy jobs, it'll help companies grow."
- **Demis Hassabis (Google DeepMind)** — the elastic 2026 forecaster: 2029 (I/O), 2031–2034 (tracker), 3–10 years (Sept interview); "foothills of the singularity."
- **Elon Musk (xAI)** — the loudest 2026 doom-accelerationist pole: end-2026 individual-parity forecast (Davos); 2030–2031 collective superintelligence; "3 years of AGI+ leaves a 10–20% chance we all die" (Feb 2026, Forbes).
- **Yann LeCun (Meta)** — the standing dissent: scaling/LLM approaches insufficient; breakthroughs in world models, mental reasoning, and planning required.
- **Fidji Simo (OpenAI applications CEO)** — "side quests" memo behind the Sora shutdown; embodiment of the 2026 shift from prestige products to economically viable ones.
- **Marc Benioff (Salesforce)** — "agentic enterprise" as enterprise architecture, not a feature; EVA agentic concierge at Davos.
- **Satya Nadella (Microsoft)** — Davos theme: "production value over prototypes"; AI changing outcomes for people and communities.
- **Yuanqing Yang (Lenovo)** — CES 2026 inference-spend-inversion forecast; the hardware-side capex marker.
- **Market researchers (Gartner, Grand View, IDC, McKinsey, KPMG, Wolters Kluwer)** — authors of the agent-economy forecast layer; methodology caveats apply, primary reports not accessed.
- **Kuaishou (Kling), ByteDance (Seedance), Google (Veo/Genie)** — the commercial video leaders left standing after Sora's exit; the economic (not capability) winners of 2026.
- **Cognition (Devin), Anysphere (Cursor)** — coding-agent companies carrying the agentic-AI revenue evidence layer (vendor-reported run-rates, unaudited).
- **Syndicated PR machinery (ABNewswire, commercial integration blogs)** — not a lab, but an actor in the narrative audit: Kuaishou's Sept 11–12 "AI Director" wave turned a February model into a September "launch" in press coverage. The event-narrative loop in miniature — company positioning laundered into "industry trend."
- **Community trackers (skynetcountdown.com, 2agi.me, heyuan110 / awesome-aitools READMEs, baditaflorin deep-research compilation)** [COMMUNITY] — the de facto primary sources for several §19 figures (Seedance timelines, Hassabis tracker ranges). They carry the audit's loosest provenance but are often the *only* source; the RAG keeps them labeled rather than dropping the claims.
- **Enterprise IT buyers (unnamed, collective)** — the audience the whole forecast layer is priced for: the 44% finance-teams figure, the 78% MCP-in-production figure, and the index.dev 4.3× ROI figure are executive-belief instruments whose function is procurement justification, not scientific measurement. Reading them as belief instruments rather than facts is the audit's recommended stance.

