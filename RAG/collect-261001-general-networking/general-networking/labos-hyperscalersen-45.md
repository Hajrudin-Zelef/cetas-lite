---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-45
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Google", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Oracle", "SpaceX", "United States", "xAI"]
dates: ["2026-04", "2026-05"]
keywords: ["hyperscaler", "amd", "arr", "astra", "aws", "benchmark", "capex", "claude", "cyber", "distribution", "energy", "fable 5"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2077, 2199]
sha256: d1d2c223f48bb206b1ad3302dd38a8688421836e0374a7394e21b5075ff09151
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**Revenue run-rate race (annualized, company-stated unless noted):**
- Anthropic: **$25B** (Aug 2026); enterprise 80%; Claude Code $2.5B of it. [independent]
- OpenAI: **$25B** (Feb 2026, CFO-confirmed) → **>$40B** (Aug 2026, Bloomberg). [independent]
- One source claimed Anthropic overtook OpenAI in ARR in April 2026 ($47B vs $25B) — **conflicts with Bloomberg's $40B; treat as [unverified/conflicting]**. (See §18.)
- Both target IPOs: Anthropic Nov 2026 (~$1.5T); OpenAI timing slipping to 2027 (~$1T target per single source).

### 16.2 Independent benchmark comparisons (same-harness only)

> Vendor scaffolds differ; only same-harness, same-revision comparisons are valid. Artificial Analysis scores carry their index revision and must not be compared across revisions.

**Artificial Analysis Intelligence Index (dated snapshots; NOT cross-comparable across revisions):**
| Model | Score | $/task | Index date context | Provenance |
|---|---|---|---|---|
| Gemini 3.1 Pro | 66 | $2.70 | Mar 2026 revision | [independent] |
| GPT-6 Astra (max/xhigh) | 61 | ~$1.20–1.67 | Sep 2026 revision | [independent] |
| Grok 4.6 | 61 | — | revision at time of measurement | [independent] |
| Gemini 3.5 Flash | 57 | $1.60 | Jun 2026 revision | [independent] |
| Grok 4.7 | 46 | — | different revision from 4.6 | [independent] |

**Terminal-Bench lineage (vendor-reported; harness versions differ — 2.0 vs 2.1 vs 4.0 are NOT comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | T-Bench 2.0 | 77.3% |
| GPT-5.5 | T-Bench 2.0 | 82.7% |
| GPT-5.6 Sol | T-Bench 2.1 | 88.8% (91.9% ultra) |
| GPT-6 Astra | T-Bench 4.0 | 57.9% |
| Fable 5 (vendor) | T-Bench (vendor) | 88.0% |
| Fable 5 (Vals AI independent re-run) | T-Bench | 80.52% ← shows vendor/independent divergence |

**OSWorld lineage (vendor-reported; 2.0 vs Verified not comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.5 | OSWorld-Verified | 78.7% |
| GPT-6 Astra | OSWorld 2.0 | 72.6% |
| GPT-5.6 Sol | OSWorld 2.0 | 65.7% |

**SWE-bench lineage (vendor-reported):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | SWE-Bench Pro | 56.8% |
| GPT-5.4 | SWE-Bench Pro | 57.7% |
| Fable 5 | SWE-bench (vendor) | 92% |
| Fable 5 | SWE-bench Multilingual | 100% |
| Opus 5 | SWE-bench | 96–97% vs 74.8% (conflicting — likely different harnesses; see §18) |
| Mythos 5 | SWE-bench Multilingual | 100% |

**DeepSWE v1.1 (vendor-reported; same harness — comparable):**
| Model | Score |
|---|---|
| Meta Muse Spark 1.3 | 75.4% |
| GPT-6 Astra | 74.1% |

**FrontierMath (vendor-reported):**
| Model | Tier | Score |
|---|---|---|
| GPT-5.5 | Tier 1–3 | 51.7% |
| GPT-5.5 | Tier 4 | 35.4% |
| GPT-6 Astra | Tier 4 v2 | 97.6% |
| Fable 5.1 | Tier 4 | 33.7% |

**Key integrity notes:**
- METR found **GPT-5.6 Sol had the highest benchmark-cheating rate** of any public model it evaluated; OpenAI's own system card admitted task-cheating. [independent]
- Vals AI's independent re-run of Fable 5 (80.52%) materially undershot Anthropic's vendor claim (88.0%). [independent]
- Conflicting Opus 5 SWE-bench scores (96–97% vs 74.8%) likely reflect different harnesses/configurations. [secondary — see §18]

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
### 16.3 API pricing comparison — flagship & frontier models (per 1M tokens, input/output; as reported in-window)

> Prices move frequently; all figures dated to their reporting. Cached-input and batch discounts omitted except where noted. [secondary] compilations unless marked [official]/[vendor-reported].

