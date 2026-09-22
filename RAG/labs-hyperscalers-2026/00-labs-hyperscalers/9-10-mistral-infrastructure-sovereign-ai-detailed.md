---
id: labs-hyperscalers-2026/00-labs-hyperscalers/9-10-mistral-infrastructure-sovereign-ai-detailed
title: "9.10 Mistral infrastructure / sovereign AI — detailed"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "EU", "Google", "Microsoft", "Mistral", "Nvidia", "United States"]
dates: ["2026-05", "2026-06", "2026-09-08"]
keywords: ["mistral", "agent", "blackwell", "capex", "compute", "energy", "funding", "gpus", "inference", "nvidia", "pricing", "research"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1347, 1406]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: f7e30c9c0176544cdca84cd12f02e3422549effc12760d2be5c300d1eb098f48
---

# 9.10 Mistral infrastructure / sovereign AI — detailed

### 9.10 Mistral infrastructure / sovereign AI — detailed

- **Bruyères-le-Châtel, Essonne**: 13,800 NVIDIA GB300 GPUs, 44 MW, operational Q2 2026 per company statements (site chosen Feb 2025); financed by the $830M debt raise [independent/secondary].
- **Les Ulis, France**: 10 MW **inference** data center, came online Q3 2026 [independent: VentureBeat interview with Timothée Lacroix] https://venturebeat.com/infrastructure/mistral-ai-wants-to-build-1-gigawatt-of-european-compute-by-2030-and-lock-in-customers-now.
- **Sweden**: 23 MW facility with EcoDataCenter (renewable energy, advanced cooling) [independent: VentureBeat].
- **Roadmap**: <200 MW operating today → 200 MW across Europe by end-2027 → **1 GW by 2030**; part of a stated **€4B investment strategy**. Epoch AI est.: ~$38B capex for a typical 1 GW facility — scale challenge flagged by VentureBeat [independent].
- **Mistral Compute** (with NVIDIA, announced VivaTech Jun 2025): 18,000 Grace Blackwell chips in Essonne, expanding continent-wide [secondary].
- **Microsoft deal (Jul 2026)**: multibillion-dollar funding of Mistral's European infrastructure; Azure Local route for sovereign deployments [independent: Reuters].
- **HUMAIN deal (Aug 2026)**: access to Saudi DC infrastructure for regional compute [secondary].
- Koyeb (AI infrastructure startup) acquired Feb 2026 — first M&A, per leaked timeline [unverified].
- Context: EU AI Act's most aggressive enforcement phase began Aug 2, 2026; June 2026 US restrictions on foreign access to two Anthropic models sharpened Europe's sovereignty push [independent/secondary].

### 9.11 Mistral personnel (2026)

- **Arthur Mensch** — co-founder & CEO; still in role (quoted on Series D, Microsoft deal, EDF deal; defended military-AI work May 2026). No departure [independent].
- **Guillaume Lample** — co-founder & Chief Science Officer; still in role (podcast interview Mar 2026 on hiring/research; retweeting company launches Jun 2026). No departure [secondary/independent].
- **Timothée Lacroix** — co-founder & CTO; active (quoted in Airbus release May 2026; VentureBeat infrastructure interview 2026). No departure [independent].
- **Johan Bergqvist** — CFO; public face of Series D (Reuters interview Sep 8, 2026) [independent].
- **Robotics chief departure — Sep 9, 2026** [secondary: Sifted via TechTimes] https://www.techtimes.com/articles/327117/20260909/mistral-robotics-chief-seeks-231-million-out-data-physical-ai-rivals.htm: Mistral's head of robotics is leaving to found an independent embodied-AI startup, reportedly seeking **€200M** (~$231M) to build robotics models on richer training data. **Name not published in accessible sources** — Sifted article text was paywall-scrambled; not identified here. Mark as name-unconfirmed [secondary, name unverified].
- **Brian Hall** (ex-Microsoft/Amazon/Google) joined as **CMO, Jun 2026** — sourced only from the leaked internal timeline doc [unverified].

### 9.12 Mistral API pricing — La Plateforme (as of Sep 2026)

> Official docs page (docs.mistral.ai pricing) **could not be fetched** for this report (fetch failed); figures below are from secondary compilations and are **not officially verified**. Treat all as [secondary].

| Model | Input / 1M tokens | Output / 1M tokens |
|---|---|---|
| Mistral Small 4 | $0.15 | $0.60 |
| Mistral Medium 3 | ~$0.40 | ~$2.00 |
| Mistral Large 3 | $2.00 | $6.00 |
| Codestral | $0.30 | $0.90 |
| Mistral Nemo | $0.02 | $0.10 |
| Magistral Small | ~$0.10 | ~$0.30 |
| Ministral 3B | $0.04 | $0.04 |
| Voxtral TTS | $16.00 / 1M characters | — |
| Voxtral Transcribe | ~$0.003 / minute | — |
| OCR 4 API | $4 / 1,000 pages (batch $2; Document AI $5) | — |

- **Mistral Medium 3.5 API pricing**: not found in sources — unknown, flag for follow-up [unverified].
- Le Chat/Vibe subscription: Free / **€14.99 Pro** / **€24.99 per user Team (€19.99 annual)** / Enterprise custom; taxes extra [independent].

### 9.13 Mistral key sources

- Funding (Reuters, Sep 8 2026): https://www.reuters.com/world/europe/french-ai-company-mistral-hits-24-billion-valuation-funding-round-2026-09-08/
- Debt raise / Essonne DC: https://www.globalbankingandfinance.com/frances-mistral-raises-830-million-debt-ai-data-centre/
- 1 GW plan (VentureBeat): https://venturebeat.com/infrastructure/mistral-ai-wants-to-build-1-gigawatt-of-european-compute-by-2030-and-lock-in-customers-now
- Microsoft deal (Reuters via SRN, Jul 21 2026): https://srnnews.com/microsoft-to-fund-mistrals-european-ai-expansion-in-multibillion-dollar-deal/
- HUMAIN partnership (Aug 24 2026): https://www.unite.ai/mistral-and-humain-team-on-sovereign-ai-for-saudi-arabia-and-the-region/
- EDF partnership (May 28 2026, official): https://presse-edf.fr/download?n=CP%20PARTENARIAT%20EDF%20MISTRAL_VENG.pdf&picid=13616
- Airbus + BMW (May 28 2026): https://www.businesstimes.com.sg/companies-markets/telcos-media-tech/mistral-signs-airbus-and-bmw-it-brings-ai-manufacturing
- Cloudera (Sep 10 2026): https://www.manilatimes.net/2026/09/10/tmt-newswire/globenewswire/cloudera-and-mistral-partner-to-bring-specialized-sovereign-intelligence-to-enterprise-data/2422454
- Medium 3.5 launch: https://www.testingcatalog.com/mistral-ai-unveils-medium-3-5-model-and-work-mode-for-le-chat/
- Voxtral Transcribe 2 (VentureBeat): https://venturebeat.com/technology/mistral-drops-voxtral-transcribe-2-an-open-source-speech-model-that-runs-on
- Voxtral TTS (TechCrunch): https://techcrunch.com/2026/03/26/mistral-releases-a-new-open-source-model-for-speech-generation/
- OCR 4: https://www.fonearena.com/blog/485781/mistral-ocr-4-features.html
- Workflows (VentureBeat): https://venturebeat.com/technology/mistral-ai-launches-workflows-a-temporal-powered-orchestration-engine-already-running-millions-of-daily-executions
- Vibe rebrand (The Decoder): https://the-decoder.com/mistral-rebrands-lechat-as-vibe-betting-its-chatbots-future-is-as-a-full-blown-work-agent/
- Robotics chief exit (Sifted via TechTimes): https://www.techtimes.com/articles/327117/20260909/mistral-robotics-chief-seeks-231-million-out-data-physical-ai-rivals.htm
- "Le Chaton Fat" hoax debunk: https://explainx.ai/blog/le-chaton-fat-mistral-ai-viral-hoax-meme-2026

