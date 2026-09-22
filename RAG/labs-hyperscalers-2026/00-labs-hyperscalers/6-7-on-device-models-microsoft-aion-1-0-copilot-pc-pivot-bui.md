---
id: labs-hyperscalers-2026/00-labs-hyperscalers/6-7-on-device-models-microsoft-aion-1-0-copilot-pc-pivot-bui
title: "6.7 On-device models: Microsoft Aion 1.0 / Copilot+ PC pivot (Build 2026)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Anthropic", "Baseten", "Fireworks AI", "Google", "Hugging Face", "Intel", "Microsoft", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "Qualcomm", "TSMC", "United States", "xAI"]
dates: ["2025-10", "2026-03-17", "2026-04", "2026-04-27", "2026-05", "2026-06-02", "2026-07", "2026-07-29", "2026-09"]
keywords: ["copilot", "3nm", "agent", "agentic", "agents", "agi", "amd", "astra", "aws", "bedrock", "blackwell", "capex"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [914, 974]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: d9311dad13dd733b8e0ec677c1637e028b640d1fb8b006c4ed3444acbbc30431
---

# 6.7 On-device models: Microsoft Aion 1.0 / Copilot+ PC pivot (Build 2026)

### 6.7 On-device models: Microsoft Aion 1.0 / Copilot+ PC pivot (Build 2026)

- **Aion 1.0** announced **June 2, 2026** at Build: on-device small language model family for Windows 11. **Aion 1.0 Instruct** (preview at launch): small SLM for summarization and rewriting. **Aion 1.0 Plan**: 14B reasoning and tool-calling model, 32K context, "coming in the following months." Open weights scheduled for **Hugging Face in July 2026** (whether shipped by Sept 22 not confirmed) [secondary] https://runaihome.com/blog/microsoft-aion-1-0-windows-local-ai-copilot-pc-2026/.
- Hardware: NPU path requires Copilot+ PC with ≥40 TOPS NPU (Qualcomm Snapdragon X Elite 45 TOPS; Intel Lunar Lake 45–48 TOPS); AMD Ryzen AI (50 TOPS) qualified on hardware but Aion NPU support deferred; runs via Windows ML / Windows AI Foundry; also runs on discrete GPUs [secondary].
- **Strategic pivot at Build 2026:** Microsoft de-emphasized the Copilot+ PC brand. CEO Satya Nadella told developers "you now have the full scope of GPUs that you can get to" when writing AI software for Windows ML; the agentic future of Windows was presented as running locally across a wider device range, including NVIDIA-powered hardware — not requiring a Copilot+ PC NPU [secondary].
- **New hardware at Build 2026:** Surface Laptop Ultra and **Surface RTX Spark Dev Box** (NVIDIA RTX Spark chip, targeting professional AI workloads well above 40 TOPS) [secondary].

### 6.8 Microsoft Foundry / Azure AI (2026)

- **Foundry catalog scale:** 11,000+ models from multiple providers per Nadella at Build 2026 (Foundry documentation counts >1,900 core models; counts above 10,000 in third-party tallies include open-weight variants) [official][secondary]. Catalog size figures vary by counting method.
- **Fireworks AI generally available on Foundry** (announced June 2, 2026): single-platform experience with enterprise governance and Azure data residency [official].
- **Frontier Tuning** (private preview, June 2, 2026): applies RL within the customer's compliance boundary to tune agents on their own data/workflows [official].
- **Anthropic models on Foundry:** Claude Sonnet 4.5, Haiku 4.5, and Opus 4.1 available in public preview in Microsoft Foundry with Azure billing (per Anthropic announcement; reported in Aug 2026 coverage) [secondary].
- **Foundry new models added (2026 digest):** NVIDIA Nemotron-3-Super-120B-A12B, IBM Granite-4.0-1b-Speech, Sarvam-105B; voice-native agents in public preview; user-scoped persistent memory for agents [secondary].
- **Build 2026 platform announcements:** Foundry Agent Service deep-dives; **Microsoft IQ context layer** (Work IQ, Web IQ, Foundry IQ); **Agent 365** governance SDK (Entra/Defender extended to agents); **Azure HorizonDB** (GPU-accelerated Fabric); **Project Solara** (hardware) [secondary].

### 6.9 Microsoft–OpenAI relationship — detailed 2026 record

**October 2025 restructuring (context):** OpenAI restructured into a public-benefit corporation (OpenAI Group PBC) valued at **$500B**; Microsoft holds ~**27%** stake (~$135B) on $13.8B invested — ~10× return [independent — Al Jazeera, Reuters].

**April 27, 2026 — commercial terms renegotiated (joint announcement):**
- **Microsoft is no longer the exclusive licensee** of OpenAI models; OpenAI may serve products on AWS, Google Cloud, Oracle; Azure remains "primary cloud partner" — products ship first on Azure "unless Microsoft cannot and chooses not to support the necessary capabilities" [secondary; AP via SRN News].
- **Revenue share:** Microsoft no longer pays OpenAI a share on Azure-resold AI products (eliminated). OpenAI continues paying Microsoft a share through 2030 — 20%, now subject to a fixed total cap (cap amount undisclosed) [secondary].
- **AGI clause removed:** the provision that would have terminated Microsoft's IP rights upon OpenAI's board declaring AGI was eliminated; Microsoft's IP license now non-exclusive with a hard expiration in **2032** [secondary].
- **Equity unchanged:** 27% stake retained [secondary].
- **Immediate consequence:** within ~24 hours, OpenAI models (GPT-5.5, GPT-5.4 preview; Codex limited preview) became available on **Amazon Bedrock**; Amazon separately disclosed a planned **$50B** investment in OpenAI [secondary] https://srnnews.com/microsoft-cuts-openai-revenue-share-in-a-fresh-step-to-loosen-their-ai-alliance/.
- Wedbush analyst Dan Ives: agreement "puts OpenAI on a strong path forward to going public through IPO" (reportedly targeted late 2026); also important for Microsoft's "tech independence" in Copilot [secondary].

**OpenAI fundraising context in 2026:** Feb 27, 2026 — OpenAI announced a **$110B funding round** ($50B Amazon + $30B SoftBank + $30B Nvidia). Pre-money valuation reported as $730B (FT/GulfNews) vs. $840B (aiforautomation) — inconsistent [secondary]. OpenAI subsequently closed a **$122B round at $852B post-money** [secondary]; as of **Sept 16, 2026**, FT reports talks of another round at ~**$1.2T** valuation ahead of a possible IPO [secondary via Reuters]. ChatGPT approaching **1B weekly active users** (reported July 29, 2026) [secondary].

### 6.10 Microsoft investments, personnel, infrastructure (2026)

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

