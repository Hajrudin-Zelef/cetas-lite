---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-26
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Baseten", "EU", "Fireworks AI", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "TSMC", "United States", "xAI"]
dates: ["2026-03-17", "2026-04", "2026-05", "2026-07-31", "2026-09"]
keywords: ["3nm", "agents", "agi", "amd", "astra", "aws", "bedrock", "blackwell", "capex", "claude", "compute", "copilot"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [946, 1003]
sha256: dd31b584d5400173d4f6ee48831d54549e9ac5aa9d6584bed2ebbd6bb30f9bfa
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

- **Microsoft–NVIDIA–Anthropic deal (context, announced Nov 2025):** Microsoft to invest **up to $5B** in Anthropic; Nvidia up to $10B; Anthropic commits **$30B of Azure compute**, up to **1 GW** on Nvidia Grace Blackwell and Vera Rubin systems; Claude on Foundry/Copilot [secondary — AP reported].
- **Personnel / org — Copilot reorganization, March 17, 2026:** Nadella announced Ryan Roslansky, Perry Clarke, and Charles Lamanna would lead Microsoft 365 apps and the Copilot platform; **Suleyman retained** frontier-model/high-ambition work, reporting to Nadella [secondary].
- **AI demand figures (vendor-reported):** AI revenue run rate **$37B (+123% YoY)**; **20M paid Copilot seats (+250% YoY)** [vendor-reported][secondary].
- In an April 2026 MIT Technology Review interview, Suleyman cited compute drivers: chip perf up >7× in six years (Nvidia), **Maia 200 at 30% better performance per dollar**, 3× HBM data flow, NVLink-scale clusters; "warehouse-scale supercomputers, hundred-billion-dollar clusters, ten-gigawatt power draws" [secondary].
- **Capex:** Microsoft CY2026 capex guidance ~**$175–190B** (figures vary by source and quarter: earlier ~$190B; later ~$175B after an accounting change extending data-center depreciable life from 15→25 years — **not a spending cut**). Fiscal Q4 2026 (Apr–Jun) capex hit **$41B** (+70% YoY); Q3 FY26 (Jan–Mar) capex $31.9B [secondary]. ~$25B was added specifically to absorb higher component pricing [secondary].
- **Capacity:** management says Microsoft expects to remain **capacity-constrained at least through 2026**; constraint framed as time-to-power for new sites plus component pricing escalation [secondary].
- **Maia 200:** in production at **two US data centers (Iowa and Arizona)**; TSMC 3nm; 216 GB HBM3E; ~10 PFLOPs FP4; 30% better performance-per-dollar (per Suleyman) [secondary][vendor-reported]. Microsoft has stated custom silicon **adds** inference capacity rather than reducing Nvidia/AMD spend [secondary].
- **Maia 300** (upcoming): The Information reported Microsoft targets a **September 2026 public unveiling** and is in talks with TSMC for manufacturing capacity of **300,000+ units** (vs. tens of thousands of Maia 200s), with longer-term ambition of 1M+ units [secondary][unverified] — unveiling date falls within window; confirmation not found as of Sept 22.
- **Capex composition:** roughly **two-thirds of recent capex is short-lived assets** (GPUs, CPUs) and ~$4.7B of Q3 FY26 finance leases were primarily large data-center sites; ~50% of 2025–2026 capex described as short-lived assets (NVIDIA H100/B200, custom Maia 200 silicon) and ~50% long-lived (construction, land, power) [secondary].
- **Power procurement:** 835 MW PPA with Constellation for **Three Mile Island / Crane Clean Energy Center**; large nuclear PPAs; gas turbine slot reservations; PPA with **Helion Energy** targeting fusion-powered data centers by **2028** [secondary].
- **DoD classified-network AI (1 May 2026)** — Microsoft among eight companies (see §2.4) [secondary]. **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — Microsoft signatory [secondary].

### 6.7 Microsoft Foundry — model roster (2026)

> Renamed from "Azure AI Foundry" to **"Microsoft Foundry"** in 2026 (per Microsoft's Build recap). [official]

| Model family | Available on Foundry | Since |
|---|---|---|
| MAI-Thinking-1 | Private preview Jun 2, 2026 | Build 2026 |
| MAI-Code-1-Flash / Image-2.5 / Transcribe-1.5 / Voice-2 | Yes | Jun 2, 2026 |
| OpenAI GPT-5.x / 6 Astra | Yes (preferred: GPT-5.5→5.6 per Jul 2026) | Ongoing |
| Grok 4.6+ | Yes | Aug 2026 |
| Mistral Medium 3.5 + OCR 4 | Yes | Jul 21, 2026 (Microsoft–Mistral deal) |
| Phi-4-reasoning-vision-15B | Yes (also Hugging Face) | Mar 2026 |

- Third-party distribution of MAI via **OpenRouter, Fireworks AI, Baseten** — first time developers can tune Microsoft's weights off-Foundry [official].
- **GitHub Copilot** runs MAI-Code-1-Flash (from Jun 2) and Grok 4.6+ (from Aug) [secondary].
- **Microsoft 365 Copilot**: powered by GPT-5.5 then GPT-5.6 as preferred model per OpenAI (Jul 2026) [vendor-reported].

## §7 — AMAZON

> **2026 at a glance — Amazon:** $50B into OpenAI (announced Feb 27, completed Jul 31) — largest single corporate AI check ever, with AWS as exclusive third-party cloud for OpenAI Frontier and 2 GW of Trainium committed; Nova line wound down to KTLO (EOL Sep 14–30) as the Frontier Model Research division under Pieter Abbeel takes over; Alexa+ GA (Feb 4).

### 7.1 Nova consolidation
- **Amazon Nova** model family — 2026 consolidation: Nova became AWS's flagship in-house model line (text, multimodal, image/video generation) competing with Bedrock-hosted third-party models. [secondary: Track B]
- Nova positioning: cost-optimized frontier for AWS enterprise customers. [secondary]

### 7.2 The $50B OpenAI investment — announced Feb 27, 2026; completed July 31, 2026 (full record in §7.6)
- Amazon anchored OpenAI's $122B round with **$50B** — **$15B funded at close; $35B contingent on IPO or AGI by end-2028** (announced Feb 27, 2026; completed July 31, 2026 per SEC filing). Largest single corporate AI investment on record. [secondary] https://tech-insider.org/openai-122-billion-funding-round-852-billion-valuation-2026/ ; https://abhs.in/blog/openai-122-billion-852-billion-valuation-amazon-agi-clause-ipo-2026
- The contingent structure was widely read as establishing a firm IPO timeline for OpenAI. [secondary]

### 7.3 Bedrock & cloud
- **OpenAI models on Amazon Bedrock from 28 April 2026** (GPT-5.5, Codex, agents) — landmark, given the Microsoft exclusivity era had just ended. [secondary] https://www.vaasblock.com/research/microsoft-openai-exclusivity-end-copilot-moat-aws-bedrock-2026/
- **Project Rainier**: reported **$38B, 7-year AWS–OpenAI cloud deal** (hundreds of thousands of GPUs, fully deployed by end-2026); separate reporting claims a **$100B AWS deal** tied to the Amazon investment — [unverified]. [secondary] https://www.coinlive.com/news/openai-raises-122-billion-in-record-breaking-funding-round-at-852 (tweet citation)
- Bedrock continued hosting Anthropic (Claude), Meta, Mistral, and Amazon Nova models. [secondary]

### 7.4 Silicon & capex
- **Trainium3** — Amazon's third-gen AI training chip, ramped in 2026; positioned against Nvidia GPUs for internal + AWS workloads. [secondary: Track B]
- **~$220B capex** — Amazon's guided 2026 capital expenditure, the highest among hyperscalers, driven by AI infrastructure. [secondary: Track B]
- **Alexa+** — the LLM-rebuilt Alexa (launched 2025) continued rollout through 2026 with Nova-class models. [secondary]

### 7.5 Partnerships & regulatory
- **"A Call for Collective Action on Cyber Defense"** (27 Aug 2026) — AWS signatory. [secondary]
- EU AI Act / DSA compliance posture for AWS AI services (Track B).

---
### 7.6 Nova model family — detailed (from Track B)

