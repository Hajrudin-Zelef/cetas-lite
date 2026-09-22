---
id: labs-hyperscalers-2026/00-labs-hyperscalers/16-3-api-pricing-comparison-dated-per-1m-token-rates-standar
title: "16.3 API pricing comparison (dated per-1M-token rates; standard tiers)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: pricing
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Google", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Oracle", "SpaceX", "United States", "xAI"]
dates: ["2026-05"]
keywords: ["pricing", "amd", "astra", "aws", "capex", "claude", "cyber", "distribution", "energy", "gpt-5.6", "gpt-6", "gpu"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2142, 2195]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 3a3c0cf816cf90bde153e2ad0b39e6b88c2d783afd35db9a928a10de926b5852
---

# 16.3 API pricing comparison (dated per-1M-token rates; standard tiers)

### 16.3 API pricing comparison (dated per-1M-token rates; standard tiers)

| Model | Input | Output | Context | Date of rate | Provenance |
|---|---|---|---|---|---|
| Claude Opus 4.6 (launch) | $5.00 | $25.00 | 200K (1M beta) | 5 Feb 2026 | [vendor] |
| Claude Opus 4.6 (Aug 6) | $8.00 | $40.00 | + long-context 2x/1.5x | 6 Aug 2026 | [vendor] |
| Claude Sonnet 4.6 (launch) | $3.00 | $15.00 | 1M | 26 Mar 2026 | [vendor] |
| Claude Sonnet 4.6 (Aug 6) | $4.00 | $20.00 | + long-context 2x/1.5x | 6 Aug 2026 | [vendor] |
| GPT-5.3-Codex | $1.75 | $14.00 | 400K | 5 Feb 2026 | [secondary] |
| GPT-5.4 | $2.50 | $15.00 | 1M | 5 Mar 2026 | [secondary] |
| GPT-5.5 | $5.00 | $30.00 | 1M | 22 Apr 2026 | [secondary] |
| GPT-5.5 Pro | $30.00 | $180.00 | — | 22 Apr 2026 | [secondary] |
| GPT-5.6 Sol | $5.00 | $30.00 | 1.05M | 9 Jul 2026 | [secondary] |
| GPT-5.6 Sol (promo) | $4.00 | $20.00 | — | thru 21 Nov 2026 | [secondary] |
| GPT-5.6 Terra | $2.50 → $2.00 | $15 → $12 | — | Jul 2026 | [secondary] |
| GPT-5.6 Luna | $1.00 → **$0.20** | $6.00 → $1.20 | — | 30 Jul 2026 cut | [secondary] |
| GPT-5.6 Cyber | $12.50 | $75.00 | — | 2026 | [secondary] |
| GPT-6 Astra | $10.00 | $50.00 | 1M | 3 Sep 2026 | [secondary] |
| GPT-6 Astra Fast | $20.00 | $100.00 | — | 3 Sep 2026 | [secondary] |

**Pricing dynamics 2026:** flagship input-token prices held roughly flat ($5–10/1M) while mid-tier prices collapsed (Luna −80%); long-context premiums (2x/1.5x beyond thresholds) became standard at both labs; caching discounts (90% reads) and batch/flex (50% off) standardize effective pricing.

### 16.4 Infrastructure & capex comparison (2026)

| Company | Capex / infra commitment (2026) | Key programs | Provenance |
|---|---|---|---|
| OpenAI | Stargate $500B (multi-yr); Oracle $300B/5-yr/4.5GW; AWS $38B/7-yr reported | Stargate Texas (100K+ GPUs for Astra); SB Energy $1B; AMD 6GW/MI450; Broadcom 10GW co-design | [secondary] |
| Anthropic | $200B / 10 GW US gov partnership; $60B Project Apollo (Nvidia); $200B UK Stargate (Nvidia+Fujitsu) | AWS Trainium (500K units); Google TPU v7; Azure 1 GW Florida | [independent/secondary] |
| Google | Elevated capex (TPU + DC) | Ironwood TPU GA; TPU 8t/8i; Project Suncatcher (orbital, conceptual) | [secondary] |
| Meta | Historic-high capex | MTIA "Iris"; GPU fleet expansion; Nebius orders | [secondary] |
| Amazon | ~$220B capex (guided, highest hyperscaler) | Trainium3; Project Rainier ($38B AWS–OpenAI reported) | [secondary] |
| Microsoft | High capex (Azure) | Maia accelerators; Aion; $50B Anthropic cloud deal | [secondary] |
| xAI | Colossus expansion (multi-100K GPUs) | Colossus 2 planning; Starlink distribution | [secondary] |
| Oracle | ~$50B AI capex; 30K layoffs to fund it | 4.5GW OpenAI contract; RPO $638B (+363%) | [independent] |

### 16.5 Personnel comparison (notable 2026 moves)

| Person | From → To | Date | Note |
|---|---|---|---|
| Alexandr Wang | Scale AI → Meta (Chief AI Officer) | 2025 deal; 2026 execution | $14.3B Scale investment |
| Caitlin Kalinowski | OpenAI robotics → Anthropic | 2026 | Cited Pentagon deal |
| Noam Shazeer | Character.AI/Google → OpenAI | 18 Jun 2026 | Transformer co-author |
| Kevin Weil | OpenAI (VP Science) → departed | 17 Apr 2026 | OpenAI for Science dissolved |
| Bill Peebles | OpenAI (Sora) → departed | 17 Apr 2026 | Sora discontinued |
| Srinivas Narayanan | OpenAI (enterprise apps CTO) → departed | 17 Apr 2026 | — |
| Fidji Simo | OpenAI (CEO Applications) → medical leave | Apr 2026 | Brockman covering |
| Kate Rouch | OpenAI (CMO) → departed | Apr 2026 | Cancer recovery |
| Denise Dresser | Slack → OpenAI (CRO) → departed | 2026 (~8 months) | — |
| Brad Lightcap | OpenAI (COO) → special projects | 2026 | — |
| Ona/Gitpod team | Ona → OpenAI (Codex group) | 11 Jun 2026 | incl. CEO Johannes Landgraf |
| Tomoro (~150 engineers) | Tomoro → OpenAI (Deployment Co.) | May 2026 | Forward-deployed eng |
| WPP 1,500 engineers | WPP → Anthropic-embedded | 21 Sep 2026 | Creative OS partnership |

---
