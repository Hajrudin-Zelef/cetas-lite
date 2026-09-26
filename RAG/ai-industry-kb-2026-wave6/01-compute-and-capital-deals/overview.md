---
id: ai-industry-kb-2026-wave6/01-compute-and-capital-deals/overview
title: "§1. Compute and Capital Deals"
domain: compute-and-capital-deals
role: deep-dive
task: funding-deals
actors: ["AMD", "AWS", "Anthropic", "Cohere", "CoreWeave", "IREN", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "xAI"]
dates: ["2025-10-06", "2025-11", "2025-11-18", "2026-01-06", "2026-02", "2026-02-20", "2026-02-27", "2026-03", "2026-03-31", "2026-05-09", "2026-07"]
keywords: ["compute", "amd", "aws", "blackwell", "claude", "funding", "funding round", "governance", "gpu", "grok", "nvidia", "research"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [44, 98]
section: "§1. Compute and Capital Deals"
delta_of: ai-industry-kb-2026
sha256: 78ddd4d94a1857148255a6114a2beef0628dfc59184ddf7cb0e70e9ed6a0731a
---

# §1. Compute and Capital Deals

Keywords: nvidia openai investment, 30 billion nvidia openai, 100 billion LOI stalled, 110 billion funding round, 122 billion close, 852 billion post-money, 730 billion pre-money, cnbc nvidia 40 billion equity, coreweave nebius lumentum coherent marvell, corning 3.2 billion, iren 2.1 billion, warrant-based conditional ceilings, circular investment theme, wedbush matthew bryson, microsoft 5 billion anthropic, nvidia 10 billion anthropic, anthropic 30 billion azure commitment, anthropic series g, grace blackwell vera rubin, amd openai warrant, 6gw amd openai, 160 million share warrant, meta amd deal, anthropic amd july 2026, up-to investment, compute capital deals

## Summary

Five verified deal clusters anchor this section: NVIDIA's restructured OpenAI investment ($30B replacing a stalled $100B/10GW non-binding LOI, round closing at $122B), NVIDIA's early-2026 AI-equity binge (CNBC tally: more than $40B across seven public and ~24 private transactions), Wedbush's "circular investment" critique of that spending, the Microsoft–NVIDIA–Anthropic pact of 2025-11-18 (up to $5B + up to $10B equity alongside a strictly separate $30B Azure compute commitment), and the OpenAI–AMD 2025-10-06 compute-warrant deal (6GW, up to 160M shares ≈10%, $0.01 strike) whose structure Meta reused in February 2026. [SECONDARY]

This is a writing-only consolidation of wave6/01-compute-capital-deals.md; every claim carries the provenance it had in that source. Corroboration standard in the source file: four independent research passes, ≥2 sources per claim, cross-checked against wave5/01-startup-funding, wave4/02-hardware, and wave4/03-security-governance — no duplication. [SECONDARY]

"Up-to" and warrant-based ceilings are labeled as such wherever they appear, and Anthropic's two distinct $30B figures (a 2025 compute commitment vs a 2026 equity raise) are kept strictly separate. [DIRECTIONAL]

## Key dated facts

### NVIDIA ↔ OpenAI: the $30B restructure
- **2026-02-20** — FT reports NVIDIA's plan to invest up to $100B in OpenAI (the 10GW non-binding LOI) has stalled, and the deal is being reworked as a smaller $30B investment. [SECONDARY]
- **2026-02-27** — OpenAI announces a formal **$110B funding round at a $730B pre-money valuation**, with NVIDIA, Amazon, and SoftBank participating. [SECONDARY]
- **2026-03-31** — The round closes at **$122B, at an $852B post-money valuation** — the largest private funding round in the corpus as of March 2026. [SECONDARY]
- The round's close was accompanied by reporting that OpenAI's revenue had crossed **$2B per month** — the commercial context in which the $122B/$852B close was priced. [SECONDARY]
- The progression reads: non-binding $100B/10GW LOI (stalled) → $30B investment plan → $110B formal round at $730B pre-money → $122B final close at $852B post-money. These are successive states of one deal story, not independent events. [SECONDARY]
- Editorial framing ("NVIDIA ditches the $100B deal for the $30B bet") belongs to secondary outlets, not to any vendor statement — the vendors never used that language. [SECONDARY]

### NVIDIA's early-2026 AI equity binge (CNBC tally)
- **2026-05-09** — CNBC reports NVIDIA has already committed **more than $40B to AI equity deals in 2026**, comprising **seven public-company transactions plus about 24 private rounds**. [SECONDARY]
- The seven public deals: **CoreWeave, Nebius, Lumentum, Coherent, Marvell at $2B each**; **Corning up to $3.2B**; **IREN up to $2.1B**. [SECONDARY]
- Corning's $3.2B and IREN's $2.1B are **warrant-based conditional ceilings ("up to")**, not closed equity injections — a distinction that matters when comparing headline figures with realized cash. [SECONDARY]
- The Corning transaction's warrant structure is documented via Corning's SEC 8-K filing — the "up to $3.2B" ceiling is a conditional warrant, not an executed cash investment. [SECONDARY]
- The initial stalling of the $100B LOI was first reported via WSJ-syndicated coverage before FT's February 20, 2026 report formalized the $30B rework. [SECONDARY]

### The "circular investment" critique
- Wedbush analyst **Matthew Bryson** described NVIDIA's AI-equity deals as falling **"squarely into the circular investment theme"** — the criticism being that NVIDIA's customers are funded, in part, to buy more NVIDIA chips. [SECONDARY]
- This is **analyst criticism (Wedbush)**, not the reporting outlets' language: Reuters and CNBC reporter copy describes the transactions neutrally; the circularity framing belongs to Bryson specifically. [SECONDARY]

### Microsoft–NVIDIA–Anthropic: 2025-11-18
- **2025-11-18** — Microsoft and NVIDIA announce **up to $5B (Microsoft) and up to $10B (NVIDIA)** in Anthropic investment, alongside a strategic partnership covering cloud and silicon. [VENDOR]
- Anthropic committed to a **separate $30B Azure compute commitment** tied to the same pact, with deployments on **up to 1GW of NVIDIA Grace Blackwell and Vera Rubin infrastructure**. [VENDOR]
- Claude is available **across all three hyperscalers — AWS, GCP, and Azure**; the Azure leg of the pact did not make Claude exclusive to any one cloud. [VENDOR]
- The **$30B Azure compute commitment (November 2025) is distinct from Anthropic's $30B Series G equity raise (February 2026, at $380B post-money)** — one is a compute-purchasing commitment, the other is equity financing. They are different dollars, different counterparties, different dates. [VENDOR]

### OpenAI–AMD: 2025-10-06
- **2025-10-06** — OpenAI and AMD announce a 6GW AI-chip partnership; OpenAI receives a warrant for **up to 160 million AMD shares (≈10% of the company) at a $0.01 strike**, vesting in tranches tied to share-price gates up to $600. [VENDOR]
- **February 2026** — Meta reuses the same warrant-and-tranche template for its own AMD deal, confirming the structure became a market pattern rather than a one-off. [SECONDARY]
- Secondary coverage framed Meta's February 2026 AMD pact as a **$100B/6GW** deal built on the OpenAI warrant template — the template scales, even if the headline is coverage-framed. [SECONDARY]
- **July 2026** — Anthropic signs an AMD pact that contains **no warrant**, showing the template was not universally adopted. [SECONDARY]


### New verified facts — expansion

### xAI Series E — $20B at $230B, closed 2026-01-06 [SECONDARY]
- **2026-01-06** — xAI officially closed a **$20 billion Series E** at a **$230 billion post-money valuation**, upsized from an initial **$15 billion** target on "overwhelming investor demand" — roughly a 33% oversubscription. [SECONDARY]
- Investor roster (all [SECONDARY]): Valor Equity Partners, StepStone Group, Fidelity Management & Research, Qatar Investment Authority (QIA), Abu Dhabi's MGX, Baron Capital, and Saudi Arabia's HUMAIN — reported at **$3 billion** from HUMAIN alone. Nvidia and Cisco Investments came in as **strategic investors**; Nvidia was listed in the announcement as supporting "buildout of the largest GPU clusters in the world."
- Use of proceeds: expansion of the **Colossus** supercomputer in Memphis, Tennessee (brought online late 2024 with 100,000 NVIDIA H100s, reportedly past **1 million H100-equivalents by end-2025**), a new third data center near Memphis, and the transition to **"Colossus II"** — integrating NVIDIA's Rubin architecture and Cisco networking for **Grok 5** training.
- Stated five-year target: **50 million H100-equivalents deployed by 2030** — a ~50× scale-up from current capacity — per secondary coverage of xAI's infrastructure roadmap. [SECONDARY]
- Vendor announcement URL cited by secondary trackers: `x.ai/news/series-e` (referenced, not fetched). The $20B figure is corroborated by at least four independent secondary write-ups (FinancialContent, arturmarkus, aibuzzz/Medium, Blockonomi/Binance Square). [SECONDARY]

