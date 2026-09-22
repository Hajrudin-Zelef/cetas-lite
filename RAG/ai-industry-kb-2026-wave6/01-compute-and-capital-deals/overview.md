---
id: ai-industry-kb-2026-wave6/01-compute-and-capital-deals/overview
title: "§1. Compute and Capital Deals"
domain: compute-and-capital-deals
role: deep-dive
task: funding-deals
actors: ["AMD", "AWS", "Anthropic", "Cohere", "CoreWeave", "Crusoe", "IREN", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Oracle", "Perplexity", "SpaceX", "xAI"]
dates: ["2025-03", "2025-10", "2025-10-06", "2025-11", "2025-11-18", "2026-01-06", "2026-02", "2026-02-02", "2026-02-20", "2026-02-27", "2026-03", "2026-03-31", "2026-05", "2026-05-09", "2026-05-20", "2026-06", "2026-06-11", "2026-06-12", "2026-07", "2026-07-06", "2026-08", "2026-08-04", "2026-09-03", "2026-09-17"]
keywords: ["compute", "amd", "arr", "astra", "aws", "blackwell", "claude", "funding", "funding round", "governance", "gpu", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [44, 122]
section: "§1. Compute and Capital Deals"
sha256: b0219ad744a63a2e9b87c4f2b119fab69220c0c2e4a4f6126ebf05b00d591f71
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

### xAI absorbed by SpaceX — largest private merger on record, 2026-02-02 [SECONDARY]
- **2026-02-02** — SpaceX confirmed it had acquired xAI in an **all-stock deal**: xAI shareholders received **0.1433 SpaceX shares for every xAI share** held. The combined entity was reported at **~$1.25 trillion** (≈$1T SpaceX + ≈$250B xAI) — a deal CNBC called the **largest private merger on record**. [SECONDARY]
- **CONTRADICTION**: one diligence summary (inkeep/tech-ipos-kb-example) states xAI was valued at only **~$80B in the transaction** — against neuralwired's **~$250B** figure. Both are secondary; the ~$250B figure is consistent with the Series E $230B mark set four weeks earlier. Treat the $80B number as provisional/conflicting, not reconciled. [SECONDARY]
- By **May 2026**, Musk confirmed xAI would cease to exist as a standalone company and fold entirely into SpaceX; the "SpaceXAI" branding appeared on X on **July 6, 2026** — but several outlets noted the new branding had not yet appeared on the company's official website or legal filings, making it a branding event layered on an integration still catching up. [SECONDARY]
- Separate context: xAI had merged with X (Twitter) in March 2025 and held a Department of Defense contract. [SECONDARY]

### SPCX — SpaceX IPO, June 2026: $75B raised, the largest IPO in history [SECONDARY]
- SpaceX filed its S-1 on **2026-05-20** (ticker **SPCX**), targeting up to $75B at a valuation of at least ~$1.75–1.8T. Goldman Sachs led the offering; ~30% of the float was earmarked for retail investors. [SECONDARY]
- **2026-06-11** — priced at **$135/share**, raising **$75 billion** (underwriter options could push it to **$86 billion**) — the largest IPO in history, ahead of Saudi Aramco's 2019 $29.4B record. Shares began trading on Nasdaq on **2026-06-12**, closing the first day up 19% at ~$160.95, a ~$2.1T market cap. [SECONDARY]
- The registration statement lists **expansion of AI compute infrastructure as the first listed use of proceeds**, alongside Starship development and Starlink expansion. [SECONDARY]
- S-1 financials (per secondary compilations): 2025 revenue **$18.7B** (+33% YoY) on a **$4.9B net loss**; Q1 2026: **−$4.3B on $4.7B revenue**; total spend $20.7B in 2025 with **$12.7B on AI**, $4.2B on Starlink, $3.8B on space. [SECONDARY]
- **Q2 2026** (SEC filing, 2026-08-04): total revenue **$7.814B** (+92% YoY), segment split — Connectivity/Starlink **$4.291B** (+66%, operating income +$1.656B, the only profitable segment); **AI (the absorbed xAI business) $2.561B (+247% YoY, operating loss −$1.257B)**; Space/launch $962M (+29%, −$542M). Starlink: $11.4B 2025 revenue (~61% of total), $4.4B operating profit; 5M subscribers Feb 2026 → 10M shortly after; ARPU down ~18% to $81/month. [SECONDARY]
- Post-IPO noise: speculation about a SpaceX–Tesla merger was publicly floated by SpaceX president Gwynne Shotwell; no formal offer, no agreement — combined-entity speculation of >$3T is analyst talk, not a deal. [SECONDARY]

### Crusoe Series F — $3.9B at $30.9B, 2026-09-17 [SECONDARY]
- **2026-09-03** — Bloomberg scooped Crusoe raising **~$3B at ~$30B**, co-led by Atreides Management and Valor Equity Partners with Mubadala Capital participating, on the heels of a **$13 billion, five-year cloud contract with quantitative trading firm Jane Street**. [SECONDARY]
- **2026-09-17** — Crusoe formally announced a **$3.9 billion Series F** at **$30.9B post-money** (an oversubscribed initial closing of an anticipated larger round), co-led by **Atreides Management, Mubadala Capital, and Valor Equity Partners**. Participants: Founders Fund, GIC, **Nvidia**, Qatar Investment Authority, Oman Investment Authority, Radical Ventures, TPG, plus Fidelity, T. Rowe Price-advised accounts, Tiger Global, ARK Invest, Altimeter, Salesforce Ventures, Baillie Gifford, StepStone, and DPR Construction — sovereign, public-market, and industrial capital in one syndicate. [SECONDARY]
- Commercial facts (company-stated via secondary): **$140B+ total contracted value**; **6+ GW contracted capacity**, 1 GW operational; Crusoe Cloud bookings **>20× YoY**; Managed Inference >**$100M contracted ARR**; cloud customers named include Cognition, Figure, Perplexity. [SECONDARY]
- Customers: OpenAI (Crusoe's 1.2GW Abilene, Texas campus — where, per one secondary report, OpenAI trained its **Astra** model), Meta (two new data center contracts in Texas and Missouri), Microsoft, Oracle. [SECONDARY]
- The capital funds existing campuses plus modular **Crusoe Spark** "AI factory" units — truck-transportable modular data centers that can be connected to large power sources without major construction crews, partly to sidestep community opposition to hyperscale builds. [SECONDARY]
- IPO prep: talks with Goldman Sachs, Morgan Stanley, JPMorgan, and Bank of America (Axios, August 2026); no filing or timeline confirmed. Three new board members: Cloudflare CFO Thomas Seifert, Primary Digital Infrastructure CIO Bill Stein, and Redwood Materials founder/CEO JB Straubel. [SECONDARY]
- Prior raise: **$1.38B at ~$10B** in October 2025 — the Series F tripled the valuation in under a year. [SECONDARY]
- Reuters confirmation (2026-09-17): the Series F was **co-led by Atreides Management, Mubadala Capital and Valor Equity Partners** with Founders Fund, Nvidia, QIA among participants; proceeds fund existing programs and Crusoe's own **AI factories**. The same Reuters piece flags the sector-scale signal: **Blackstone + Alphabet's Crux AI** secured a **$22B chip loan** alongside a **$5B Blackstone equity investment** — debt financing for AI compute now operates at tens of billions per transaction. [SECONDARY]

