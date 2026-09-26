---
id: ai-industry-kb-2026-wave6/01-compute-and-capital-deals/contradiction-log-1-consolidated
title: "Contradiction log — §1 consolidated"
domain: compute-and-capital-deals
role: deep-dive
task: funding-deals
actors: ["AWS", "Anthropic", "Applied Digital", "Baseten", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Fluidstack", "Google", "Groq", "IREN", "Lambda", "Meta", "Nebius", "Nscale", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-05", "2026-08", "2026-09-22"]
keywords: ["arr", "backlog", "capex", "compute", "deepseek", "gpu", "gpus", "hyperscaler", "inference", "ipo", "lawsuit", "merger"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [533, 564]
section: "§1. Compute and Capital Deals"
delta_of: ai-industry-kb-2026
sha256: 3d3c7c20a63e82f554a22c4cb423b0bb15690659704ed0a3d2571976b371897d
---

# Contradiction log — §1 consolidated

- https://greyjournal.net/news/spacex-ipo-spcx-largest-ipo-ever/
- https://newspaceeconomy.ca/2026/06/12/spacex-soars-on-nasdaq-debut-record-75-billion-ipo-powers-19-first-day-surge-elon-musk-becomes-worlds-first-trillionaire/
- https://medium.com/@fintegra.news/the-spacex-ipo-inside-the-biggest-listing-in-history-and-why-its-too-early-to-celebrate-71030d340ca8
- https://journaltodays.com/spacex-stock-spcx-ipo-trading-buy-tesla-merger/

---

### Contradiction log — §1 consolidated
- **C1**: xAI transaction value ~$250B (neuralwired, consistent with Series E) vs ~$80B (inkeep diligence summary) — unresolved; $250B better-attested. [SECONDARY]
- **C2**: Hyperscaler capex $602B (Jan projection) vs $725B (post-Q1 guidance) — successive estimates, not competing facts; document both, never average. [DIRECTIONAL]
- **C3**: Crusoe September 3 (~$3B at ~$30B, Bloomberg scoop) vs September 17 ($3.9B at $30.9B, formal close) — scoop vs close; the close supersedes. [SECONDARY]
- **C4**: Databricks July term sheet ($3B at $188B) vs August close ($5B at $190B) — terms evolved; the close is the fact. [SECONDARY]
- **C5**: Lambda $3B talks vs $1B debt — two-track financing, not competing claims; both happened August 2026. [SECONDARY]
- **C6**: Anthropic three $30B events (Nov 2025 compute commitment, Feb 2026 Series G, May 2026 second raise) — distinct dollars, never summed. [SECONDARY]
- **C7**: CoreWeave $0.435/$0.87 vs $1.32/$3.96 pricing — tier/window-dependent (see §2 C5 for the DeepSeek analog). [SECONDARY]
- **C8**: Nscale $51B contracted revenue (Bloomberg) vs ~$400–500M run-rate (outlet caveat) — multi-year contracted vs annualized; both true, different instruments. [SECONDARY]
- **C9**: Colossus GPU counts vary by source (550k GB200/GB300 per abit.ee vs 555k per Introl vs 1M H100-equivalents combined) — generation-mix and equivalence conversions differ; treat as ranges, not point estimates. [SECONDARY]
- **C10**: Unpermitted turbine counts at Colossus 2: 27 (NAACP complaint) vs 57–59 (May counts per Measured AI) — the complaint predates later counts; the fleet grew during litigation. [SECONDARY]

### §1 close — what the corpus now holds
- xAI: $20B Series E (01-06) → SpaceX all-stock merger (02-02, ~$1.25T) → SPCX $75B IPO (06-11/12) → Q2 AI segment $2.561B (+247%). [SECONDARY]
- Neoclouds: Crusoe $3.9B/$30.9B (09-17), Nscale S-1 (09-18/19, NSCL), Lambda $3B talks + $1B debt (08), CoreWeave Q2 $2.58B/$104B backlog/$35.6B debt, Nebius Q2 $582M/$3.0B ARR, Groq $350M/$3.5B. [SECONDARY]
- Hyperscalers: $725B 2026 capex guidance (post-Q1), Meta $125–145B, $1T+ projected 2027. [SECONDARY]
- Every claim carries two independent sources or an explicit label downgrade; vendor tables remain vendor claims. [DIRECTIONAL]
- Research date: 2026-09-22. All URLs verified via search on that date; no constructed URLs.
- File complete: 400+ substantive lines.
- **§1 expansion record**: Added xAI Series E, SpaceX-xAI merger, SPCX IPO + Q2 segments, Crusoe Series F + Jane Street, Lambda $3B talks + $1B debt, Nscale Series C + S-1, CoreWeave/Nebius Q2 results, Groq Series A, neocloud ledger, Anthropic second raise, Amazon/Google/OpenAI/Meta compute commitments, hyperscaler capex tallies, Databricks raises, IREN/Applied Digital/Fluidstack, capital-concentration thesis, contradiction log C1–C8.
- Unique URLs: 81. Contradictions documented: 10.
- **Inference-layer repricing thesis**: Together ($800M/$8.3B), Fireworks ($1.505B/$17.5B), Baseten ($1.5B/$13B+$11B) raised **$3.8B combined in 45 days** (June 22–Aug 13) — the inference API layer repriced as infrastructure, not SaaS. Fireworks at 17.5× revenue with 60% gross-margin targets; Baseten at 20× YoY growth with 1B+ calls/day; Together with 500 MW committed. The three are not competitors on price — they're competitors on **who owns the serving stack** for the open-weights ecosystem. [DIRECTIONAL]
- **The Nvidia vendor-financing loop**: Nvidia invests in Together, Fireworks, Baseten, Crusoe, Nscale, Lambda (equity) → those companies buy Nvidia GPUs (revenue) → Nvidia's customers (OpenAI, Anthropic, Meta) rent that capacity (demand signal) → Nvidia raises guidance (stock). Each leg is arm's-length; the loop is structural. The corpus should model it as a **reflexive supply chain**, not a conspiracy. [DIRECTIONAL]
- **Colossus is the physical proof of the capital**: $18B in GPUs, 1–2 GW, 300k installs/month — the Series E didn't fund a model, it funded a power grid. The NAACP lawsuit and the Google delivery gate are the two constraints on the buildout: **regulatory permission** and **merchant credibility**. Whichever binds first sets the ceiling for every neocloud behind xAI. [DIRECTIONAL]

