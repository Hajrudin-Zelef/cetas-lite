---
id: labs-hyperscalers-2026/00-labs-hyperscalers/2-5-infrastructure-compute-stargate
title: "2.5 Infrastructure & compute (Stargate)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: funding-deals
actors: ["AMD", "Anthropic", "Broadcom", "CISA", "California", "Crusoe", "EU", "Google", "Nvidia", "OpenAI", "Oracle", "SpaceX", "United States"]
dates: ["2026-01", "2026-03", "2026-03-05", "2026-04", "2026-06", "2026-07", "2026-08"]
keywords: ["compute", "accelerator", "agent", "agents", "agi", "amd", "antitrust", "astra", "benchmark", "blackwell", "capex", "chatgpt"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [457, 493]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 790d6cde692d8e1201905e464884065049c612273d4ac70c964fbb5ec82593c0
---

# 2.5 Infrastructure & compute (Stargate)

#### Lawsuits, regulation
- **Nippon Life lawsuit** — filed **4 March 2026** (N.D. Illinois, Chicago); insurer accused ChatGPT of **unauthorized practice of law**, seeking $300K compensatory + $10M punitive; one of the first cases of its kind. [independent: Reuters] https://www.reuters.com/legal/legalindustry/openai-hit-with-lawsuit-claiming-chatgpt-acted-an-unlicensed-lawyer-2026-03-05/
- **AI slowdown antitrust class action** — filed **18 Sept 2026** (N.D. California) against Anthropic, OpenAI, SpaceXAI, Google (see §1.4). [independent: AP] https://www.cnbctv18.com/technology/anthropic-openai-and-google-sued-over-alleged-deal-to-slow-ai-development-19994529.htm ; https://www.wvlt.tv/2026/09/20/lawsuit-says-anthropic-openai-spacexai-google-made-illegal-agreement-ai-slowdown/
- **Executive Order 14409** — "Promoting Advanced Artificial Intelligence Innovation and Security," signed **2 June 2026**: voluntary framework giving federal agencies **up to 30 days pre-release access** to "covered frontier models"; NSA-led classified benchmarking within 60 days; explicitly no mandatory licensing. GPT-5.6 was the first model shipped under it. Trump later rejected broader AI-regulation calls and planned an AI task force. [secondary, multiple] https://github.com/supwils/swil-news/blob/HEAD/NEWS/ai-tech/en/2026-06-03_ai-tech-digest.md ; https://savvymonknewsletter.com/p/first-anthropic-now-openai-washington-is-gating-frontier-ai-customer-by-customer
- **EU AI Act**: GPAI systemic-risk enforcement went live **2 August 2026** (fines up to 3% turnover / €15M for incident-reporting failures; up to 7% for prohibited practices). EC confirmed **OpenAI filed its first EU AI Act incident report** (confirmed 7 Sept 2026 by spokesperson Thomas Regnier); timing questions raised by Nightingale Collective research on a "DseWiki" incident. [secondary] https://www.techtimes.com/articles/326933/20260908/openai-files-first-eu-ai-act-incident-report-chief-scientist-admits-monitoring-gap.htm ; https://www.techtimes.com/articles/322604/20260801/eu-engages-openai-anthropic-after-ai-models-hacked-real-companies-fines-take-effect-sunday.htm
- EU DSA: Commission **notified OpenAI of VLOP designation risk 7 April 2026** (ChatGPT ~75M EU users; fines up to 6% revenue). [secondary] https://knowaiuse.com/eu-openai-digital-services-act-dsa/

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

