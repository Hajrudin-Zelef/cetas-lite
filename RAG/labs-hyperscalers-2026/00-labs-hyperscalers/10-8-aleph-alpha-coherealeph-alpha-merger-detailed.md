---
id: labs-hyperscalers-2026/00-labs-hyperscalers/10-8-aleph-alpha-coherealeph-alpha-merger-detailed
title: "10.8 Aleph Alpha + Cohere–Aleph Alpha merger — detailed"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AMD", "AWS", "Cohere", "EU", "Hugging Face", "Inflection AI", "Microsoft", "Mistral", "Nvidia", "OpenRouter", "Perplexity", "Poolside", "Sakana", "SpaceX", "Stripe", "xAI"]
dates: ["2026-04", "2026-09", "2026-09-16"]
keywords: ["cohere", "merger", "agent", "amd", "apache", "arr", "aws", "bedrock", "compute", "consumer", "cost", "fugu"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1509, 1553]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 248bdd3461d5f5a82db9ab201c9b89f731fdf9a6491f674a21dfcadf406226ab
---

# 10.8 Aleph Alpha + Cohere–Aleph Alpha merger — detailed

### 10.8 Aleph Alpha + Cohere–Aleph Alpha merger — detailed

**Pre-merger 2026 news:** Strategic pivot (from Sept 2024, completed through 2025): away from frontier Luminous model API to **PhariaAI** platform/operating system for enterprise AI; Luminous model details deprecated [secondary]. **Pharia-1-LLM-7B-control**: 7B open-weight multilingual model (EN/DE/FR/ES/IT/PT/NL), 8K context, released under Open Aleph licence (non-commercial research/education; commercial requires review) — developed via partnership with Silo AI [secondary]. Late 2025: workforce reduced ~17% (~50 employees); co-founder Jonas Andrulis stepped down as CEO effective 1 Jan 2026 (left the company earlier in 2026) [independent — Reuters]. Sovereign AI contracts: German government and enterprise customers (Deutsche Bank, SAP, Bosch, Infineon named in coverage) [secondary]; revenue <€1M in 2023 per latest public accounts [independent — Reuters]. No 2026 funding round found.

**Merger with Cohere — the defining 2026 event:**
- First disclosed **April 2026** (AFP/FT reported; FT estimated ~$20B combined value) [independent] https://www.newagebd.net/post/telecom/297882/canadian-ai-startup-cohere-to-acquire-germanys-aleph-alpha.
- **Definitive merger agreement signed 16 Sept 2026** [independent — Reuters] https://www.reuters.com/legal/transactional/cohere-aleph-alpha-combine-target-enterprise-ai-market-2026-09-16/: combined entity operates as **Cohere**, dual HQ Toronto + Berlin, Heidelberg office becomes a research hub; co-chief Ilhan Scheer becomes COO; Samuel Weinbach (co-founder) becomes chief research officer [independent/secondary].
- Financial terms of the definitive deal not disclosed; subject to Canadian/German (possibly EU) regulatory approval, expected to close later in 2026 [independent].
- Schwarz Group (Lidl/Kaufland owner) investing **€500M**; plans €11–13B German data-center campus (up to 100,000 AI chips) via its StackIT cloud — the sovereign compute bet behind the deal [independent — Reuters via Tech Startups].
- Joint Command + Pharia model line unification targeted for Q4 2026 [secondary]. Positioning: first transatlantic "sovereign AI" stack for governments/regulated enterprises [independent].

### 10.9 AI21, Perplexity, and other labs — detailed

**AI21 (Israel):**
- **Jamba Mini 2** (Jan 2026) — `jamba-mini-2-2026-01` API alias `jamba-mini`; 52B total / 12B active, 256K context; **Apache 2.0** [secondary].
- **Jamba 3B / Jamba Reasoning 3B** — compact models (256K context), weights on Hugging Face (ai21labs), Apache 2.0, RLVR-trained; weights-only, no API endpoint [secondary]. Community coverage describes Jamba Reasoning 3B running on consumer devices (laptops/phones) [secondary].
- **AI21 Gateway (beta, Sept 2026)** — drop-in LLM endpoint that analyzes whole agent-run trajectories and adaptively routes/compacts context; claims up to 40% token-cost reduction (68% on one code-review agent test) [vendor-reported/secondary — single-source; treat cautiously].
- API deprecations: Jamba Mini 1.7 deprecated 1 Feb 2026; AWS Bedrock Jamba 1.5 variants EOL 26 Nov 2026 [secondary].
- **No Jamba Large 2 on the API as of 8 Sept 2026** [secondary]. No 2026 funding round found. No major 2026 enterprise wins found in this pass.

**Perplexity AI:**
- Revenue: annualized revenue >$750M by Aug 2026, up from <$250M at start of 2026 (~$450M in March; $750M+ by August) [secondary — The Information via aggregators]. ~780M queries/month cited by CEO [secondary].
- **Funding talks (Aug 2026): Nvidia in discussions to participate in a round at $30B+ valuation** — 50%+ above the ~$20B valuation of late 2025 [secondary — The Information exclusive 23 Aug 2026 via aggregators; unconfirmed]. Total raised to date >$1.5B; $750M Microsoft Azure agreement signed [secondary]. CEO Aravind Srinivas has floated a 2028 IPO target [secondary].
- Product: Perplexity Computer (cloud-based agent for computer tasks) is the key growth driver; Sonar API; Comet browser went free Oct 2025; advertising discontinued Feb 2026 (subscription-first pivot) [secondary].
- Note: classified as an "AI search" company rather than a foundation-model lab; mark product-company nuance.

**Liquid AI (MIT spinoff, liquid neural networks):** No 2026 funding or major new model family found in this pass. Last confirmed financing: $250M Series A (Dec 2024) led by AMD at $2B+ valuation [independent, 2024]. Founders: Daniela Rus, Ramin Hasani. Mark as: **no significant verifiable Feb–Sep 2026 news** — include only with this caveat.

**Character AI / Inflection AI / Poolside:** No significant verifiable Feb–Sep 2026 news found in this pass. Excluded; flagged for follow-up.

### 10.10 Cross-lab comparison — September 2026 snapshot (from Track C)

| Lab | HQ | Flagship model (Sep 2026) | Open weights? | Headline 2026 event |
|---|---|---|---|---|
| xAI (SpaceXAI) | USA (Memphis/Starbase) | Grok 4.7 (2.1T params, $2/$6 per 1M) | No (API only) | Merged into SpaceX (Feb); SPCX IPO raised $75B (Jun); Colossus ~2 GW |
| Mistral AI | France | Mistral Medium 3.5 (128B dense, mod-MIT); Large 3 (Dec 2025) | Mostly yes (Apache 2.0 / mod-MIT) | €3B Series D at €21B+ (Sep 8); Microsoft multibillion deal (Jul); HUMAIN (Aug) |
| Thinking Machines | USA | Inkling (975B/41B MoE, Apache 2.0, Jul 15) | Yes (Inkling) | Inkling open release; Nvidia 1 GW deal (Mar); talks of $5–6B raise at $40B (unconfirmed) |
| Sakana AI | Japan | Fugu Max ($2/$6) / Fugu Ultra v2 ($5/$30) orchestration APIs | Partial (orchestrator, hosted) | Fugu Max/Ultra v2 launch (Sep); NVIDIA Nemotron pool (Aug); Nature AI Scientist (Mar) |
| Cohere | Canada (+Berlin) | Command A+ (218B/25B MoE, Apache 2.0, May 20–21) | Yes (Command A+, Tiny Aya) | Definitive merger with Aleph Alpha signed Sep 16; $240M ARR |
| AI21 | Israel | Jamba Mini 2 (Jan, Apache 2.0); Gateway beta (Sep) | Partial | No 2026 round found; Gateway beta (Sep) |
| Aleph Alpha | Germany | PhariaAI platform (pivot) | Partial (Pharia-1-LLM-7B-control) | Merging into Cohere; Schwarz Group €500M + €11–13B DC bet |
| Perplexity | USA | Sonar API; Perplexity Computer agent | No | >$750M annualized revenue; Nvidia talks for $30B+ round (Aug, unconfirmed) |

**Merger wave (2026):** xAI merged into SpaceX (Feb), Cohere and Aleph Alpha signed a definitive merger (Sep 16), Stripe acquired OpenRouter (announced Aug 19), and NVIDIA signed a definitive agreement to acquire Hugging Face (Sep 3). **Sovereign AI** became a structuring force in Europe (Mistral's €4B infra strategy; Cohere–Aleph Alpha as the "transatlantic sovereign stack") and the Middle East (Mistral–HUMAIN; Saudi PIF).

