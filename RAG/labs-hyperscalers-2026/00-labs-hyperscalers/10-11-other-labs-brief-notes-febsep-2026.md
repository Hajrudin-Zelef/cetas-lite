---
id: labs-hyperscalers-2026/00-labs-hyperscalers/10-11-other-labs-brief-notes-febsep-2026
title: "10.11 Other labs — brief notes (Feb–Sep 2026)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "Anthropic", "Cohere", "Google", "Hugging Face", "Inflection AI", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Poolside", "Sakana", "SpaceX", "TensorRT-LLM", "Together AI", "xAI"]
dates: ["2026-03", "2026-05", "2026-07", "2026-09"]
keywords: ["acquisition", "amd", "apache", "benchmarks", "blackwell", "capex", "cohere", "cost", "cyber", "distribution", "fugu", "funding"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1554, 1638]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 997cfa19b9c2dfdf909dbcccdc59578f58620bdcbafcf1073022b971a53381ae
---

# 10.11 Other labs — brief notes (Feb–Sep 2026)

### 10.11 Other labs — brief notes (Feb–Sep 2026)

**Liquid AI (MIT spinoff; liquid neural networks):**
- Founders: Daniela Rus, Ramin Hasani. Last confirmed financing: **$250M Series A (Dec 2024)** led by AMD at $2B+ valuation [independent, 2024 — context].
- **No significant verifiable Feb–Sep 2026 news found in this pass** — no new funding round, no major new model family. Include only with this caveat; flagged for follow-up.

**Character AI:**
- No significant verifiable Feb–Sep 2026 news found in this pass. Excluded from detailed coverage; flagged for follow-up.

**Inflection AI:**
- No significant verifiable Feb–Sep 2026 news found in this pass. (Context: Microsoft's 2024 acqui-hire of Mustafa Suleyman and much of the Inflection team preceded the window; Suleyman now leads Microsoft AI and unveiled the MAI family Jun 2, 2026 — see §6.) Excluded from detailed coverage; flagged for follow-up.

**Poolside:**
- No significant verifiable Feb–Sep 2026 news found in this pass. Excluded; flagged for follow-up.

**Why this matters for the RAG:** the 2026 narrative is dominated by scaled incumbents and well-funded challengers (Anthropic, OpenAI, xAI/SpaceXAI, Google, Meta, Microsoft, Mistral, Thinking Machines, Cohere). Smaller 2023–2024-vintage labs largely went quiet — itself a signal of the capital-intensity bar rising. [research finding]

## §11 — NVIDIA NIM

> **2026 at a glance — NVIDIA NIM:** Nemotron 3 (Nano 30B / Super 100B / Ultra 500B, NVIDIA Open Model License) positioned as the open-weight challenger stack; NIM containers via AI Enterprise/Build/DGX Cloud; Vera Rubin ≥1 GW committed to Thinking Machines; definitive agreement to acquire Hugging Face (signed Sep 3, close H1 2027).

### 11.1 NIM platform (2026)
- **NVIDIA NIM** (NVIDIA Inference Microservices) — continued as NVIDIA's standard inference-serving layer through 2026: containerized optimized model endpoints deployable on-prem, in cloud, and at the edge. [secondary: Track D]
- 2026 expansion: broader model catalog (open and proprietary models packaged as NIMs), tighter integration with **NeMo**, **TensorRT-LLM**, and **Dynamo** (NVIDIA's distributed inference framework). [secondary]
- NIM remained NVIDIA's strategic answer to inference-API competition: own the serving layer wherever the GPUs are. [secondary]

### 11.2 NVIDIA → Hugging Face definitive acquisition agreement — 3 September 2026
- **NVIDIA and Hugging Face signed a definitive acquisition agreement on 3 September 2026**; closing expected **H1 2027** (regulatory approvals pending). [secondary: Track D — full detail in §17 M&A tracker]
- Strategic significance: NVIDIA absorbs the central open-model hub (models, datasets, Spaces, inference endpoints) — vertical integration from silicon to model distribution. [secondary]
- Follows the July 2026 Hugging Face breach (see §2.4) — security posture of the combined entity under scrutiny. [secondary]

### 11.3 NVIDIA partnerships & investments (2026)
- **$30B pure-equity investment in OpenAI's March 2026 round** (see §2.2). [secondary]
- **DoD classified-network AI (1 May 2026)** — NVIDIA among eight companies. [secondary]
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — NVIDIA signatory (via the 100-company letter; listed among AI companies in some coverage — verify signatory list). [secondary]
- Continued GPU supply relationships with all hyperscalers; **Blackwell ramp** through 2026; **Rubin** roadmap (2026–2027). [secondary: Track D]
- NVIDIA's inference-market position: NIM + GPUs + networking (NVLink/Infiniband/Spectrum-X) as the full stack. [secondary]

---

### 11.4 NVIDIA NIM — detailed (from Track D)

**NIM (Nvidia Inference Microservices):** containerized GPU-accelerated inference microservices, part of NVIDIA AI Enterprise ($4,500 per GPU/year [secondary]); also consumable via NVIDIA Build (credits) and third-party clouds [secondary].

**Nemotron 3 family (launched 2026)** [independent — VentureBeat Oct 14 2025 pre-announcement; 2026 GA coverage] https://venturebeat.com/ai/nvidia-unveils-nemotron-3-open-reasoning-model-families:
- Three sizes: Nano (30B), Super (100B), Ultra (500B); open weights under **NVIDIA Open Model License** [vendor-reported].
- "Hybrid Mamba-Transformer" architecture; 1M-token context for some variants; trained on 60T tokens (synthetic-heavy) [vendor-reported].
- VentureBeat's framing: Nvidia positioning as an open-weight challenger to GPT-5.5 / Gemini 3.5 / Grok 4.5 class models [independent].
- Benchmarks (vendor-reported, Oct 2025 announcement): competitive with leading open models on reasoning/coding evals [vendor-reported].
- **License nuance:** Nemotron 3 open weights carry the NVIDIA Open Model License — review commercial terms vs Apache 2.0 before treating as unrestricted open source [vendor-reported].
- **Nemotron in the wild:** Sakana's Fugu Max (Sept 2026) folds Nemotron family models into its orchestration pool via Aug 2026 NVIDIA collaboration [secondary].

**NIM deployment model:** self-hosted on own GPUs, via NVIDIA DGX Cloud, or third-party (e.g., Together AI serves some NIM containers) [secondary]. Nemotron CC / Nano / Super / Ultra available as NIM endpoints [secondary]. Hugging Face integration: NVIDIA's agreement to acquire Hugging Face (definitive agreement, Sep 3, 2026 — see §17) would make HF the primary distribution channel for NIM/Nemotron artifacts — strategic, not yet closed [independent].

**Rubin / Vera Rubin (2026 roadmap):** next-gen GPU platform; Thinking Machines' March 2026 Nvidia deal includes procurement of at least **1 GW of Vera Rubin systems**, deployment starting 2027 [independent]. Industry cost estimates for 1 GW-class capacity ~$50B [secondary].

**DGX Spark (personal AI supercomputer):** shipping 2026 — GB10 Grace Blackwell Superchip desktop; NIM-compatible local inference target [secondary].

**NIM pricing:** no public per-token price list located; enterprise licensing via AI Enterprise ($4,500/GPU/yr) [secondary]. Mark: **no published per-token NIM price list as of Sep 22, 2026** — research gap.

### 11.5 Key sources — NVIDIA
- Nemotron 3 (VentureBeat): https://venturebeat.com/ai/nvidia-unveils-nemotron-3-open-reasoning-model-families
- Thinking Machines Nvidia deal: https://siliconangle.com/2026/03/30/thinking-machines-nvidia-multi-year-deal

### 11.6 NVIDIA NIM model catalog (2026 snapshot)

| NIM artifact | Type | License | Availability |
|---|---|---|---|
| Nemotron 3 Nano (30B) | Open reasoning model | NVIDIA Open Model License | NIM container; HF |
| Nemotron 3 Super (100B) | Open reasoning model | NVIDIA Open Model License | NIM container; HF |
| Nemotron 3 Ultra (500B) | Open reasoning model | NVIDIA Open Model License | NIM container; HF |
| Nemotron CC / Nano / Super / Ultra endpoints | Hosted inference | — | NVIDIA Build / DGX Cloud |
| Third-party NIMs (e.g., via Together AI) | Partner-served | Per-model | Partner clouds |

- **Consumption paths:** self-hosted (AI Enterprise $4,500/GPU/yr), NVIDIA Build (credits), DGX Cloud, third-party clouds. [secondary]
- **Strategic note:** the Sep 3, 2026 definitive agreement to acquire Hugging Face would make HF the primary distribution channel for NIM/Nemotron artifacts — pending close H1 2027. [independent]
- **Gap:** no public per-token NIM price list as of Sep 22, 2026.

### 11.7 Rubin / Vera Rubin roadmap + DGX Spark (from Track D)

- **Vera Rubin platform:** NVIDIA's next-gen GPU platform after Blackwell; the March 2026 Thinking Machines deal commits ≥1 GW of Vera Rubin systems with deployment starting 2027 [independent].
- **Cost scale:** industry estimates ~$50B for 1 GW-class AI capacity [secondary] — the capex bar for next-gen training clusters.
- **DGX Spark:** personal AI supercomputer (GB10 Grace Blackwell Superchip desktop) shipping 2026; NIM-compatible local inference target for developers running Nemotron-class models at the desk [secondary].
- **Strategic read:** NVIDIA is simultaneously (a) selling the picks and shovels (Rubin), (b) giving away the models (Nemotron open weights), and (c) buying the distribution (Hugging Face, pending) — a full-stack capture of the open-model inference path. [research finding]

