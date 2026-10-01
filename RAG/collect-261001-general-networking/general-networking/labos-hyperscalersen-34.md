---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-34
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "EU", "Google", "Microsoft", "Mistral", "Nvidia", "United States"]
dates: ["2026-05", "2026-05-22", "2026-05-28", "2026-06"]
keywords: ["agents", "blackwell", "capex", "compute", "copilot", "cyber", "cybersecurity", "energy", "foundry", "funding", "gpus", "inference"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1332, 1389]
sha256: bcafe3f3b66a77dae872d9f46f7802d5e8645a16f040cf569b5c8b905c91d36d
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

| Date | Partner | Terms / scope | Provenance |
|---|---|---|---|
| Feb 2026 | **Accenture** | Strategic partnership: Accenture as enterprise deployment/SI partner | [unverified — leaked timeline doc only] |
| Mar 17, 2026 | **NVIDIA** (GTC) | Forge launch; GB300 hardware supply; "Nemotron Coalition" joint frontier-model work (leaked doc) | [secondary] |
| May 22, 2026 | **Emmi** (acq.) | Austrian physics-AI/simulation startup acquired; ~€300M figure [unverified]; feeds "Mistral for Industrial Engineering" stack | [independent: TPS/Bloomberg lineage; price unverified] |
| May 28, 2026 | **Airbus** | 5-year agreement across commercial aircraft, helicopters, defence, space; full product suite licences; on-premises deployment; access to Mistral research teams + roadmap influence; use cases: tech-doc automation, simulation/optimisation, onboard/edge AI, defence cyber + code assistance | [independent: Bloomberg/Business Times; vendor: Airbus press] |
| May 28, 2026 | **BMW Group** | Central partner for BMW "Large Industry Model" (LIM): multimodal reasoning on engineering data; crash simulation | [independent: TPS; leaked doc mentions 1 PB simulation data — unverified] |
| May 28, 2026 | **EDF** | 5-year partnership: AI for nuclear engineering/maintenance, EPR2 construction; conversational agents over fleet "technical memory"; data stays EDF-owned on sovereign cloud/EDF DCs; explicitly excludes plant control systems | [official: EDF press release] https://presse-edf.fr/download?n=CP%20PARTENARIAT%20EDF%20MISTRAL_VENG.pdf&picid=13616 |
| Jun 2026 | **SAP** | Sovereign AI stack for French/German government (reported late-2025 origin) | [unverified — leaked doc] |
| Jul 21, 2026 | **Microsoft** | Multibillion-dollar deal (Reuters): Microsoft funds Mistral's European compute infra; Azure customers can build on **Mistral's French data centers**; Medium 3.5 + OCR 4 added to Azure AI Foundry; Medium 3.5 in Copilot Studio; open models runnable via Azure Local in independent DCs. Joint interview: Brad Smith + Arthur Mensch | [independent: Reuters] https://srnnews.com/microsoft-to-fund-mistrals-european-ai-expansion-in-multibillion-dollar-deal/ |
| Aug 24, 2026 | **HUMAIN** (Saudi PIF) | Strategic collaboration: AI infrastructure, model development/localisation (cybersecurity, voice, **Arabic frontier models**), joint go-to-market in Saudi; Mistral to explore HUMAIN DC capacity; deal valued in **hundreds of millions of euros** | [secondary: Unite.AI, YourStory; joint announcement] https://www.unite.ai/mistral-and-humain-team-on-sovereign-ai-for-saudi-arabia-and-the-region/ |
| Sep 10, 2026 | **Cloudera** | Strategic partnership: Mistral frontier/open-weight models on Cloudera hybrid data platform; sovereign/on-prem/edge/air-gapped deployment for regulated industries | [vendor-reported: GlobeNewswire joint release] |

Also in orbit: CMA CGM named as industrial-stack customer and Workflows user; Stellantis, TotalEnergies, SNCF, Siemens, Veolia listed as existing partners/clients in May 2026 summit coverage [secondary].

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

