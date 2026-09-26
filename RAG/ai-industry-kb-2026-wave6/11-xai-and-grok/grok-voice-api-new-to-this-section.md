---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/grok-voice-api-new-to-this-section
title: "Grok Voice API (new to this section)"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["Nvidia", "SpaceX", "xAI"]
dates: ["2023-05", "2023-07", "2024-11", "2025-03-07", "2025-04", "2025-12-30", "2026-01-06", "2026-02", "2026-02-02", "2026-05", "2026-05-03", "2026-05-15", "2026-06", "2026-07-15", "2026-08", "2026-09", "2026-09-22"]
keywords: ["grok", "voice", "accelerator", "acquisition", "agent", "agentic", "arr", "compute", "context window", "funding", "gpu", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5346, 5395]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: 442f43095682e22245b844803daeca2caaaea7506c909beb784263e2fe8aec07
---

# Grok Voice API (new to this section)

### Grok Voice API (new to this section)
- The Voice API covers real-time conversations, speech-to-text, and text-to-speech. [VENDOR, S44 — single source]
- Voice API pricing: Agent (speech-to-speech) $0.08/min; TTS $15.00 per 1M characters; STT batch $0.10/hour; STT streaming $0.20/hour. [VENDOR, S44 — single source]
- The "Voice Pricing" section names `grok-voice-think-fast-2.0` for speech-to-speech; `grok-voice-think-fast-1.0` no longer appears, though its retirement is not documented (absence not treated as retirement). [SECONDARY, S29 — single source]
- The Speech-to-Text API is generally available across 25 languages. [SECONDARY, S12 — single source]
- Grok Voice serves real-time voice mode in the Grok mobile app and in Tesla vehicles, with tool calling and real-time data access. [SECONDARY, S22 — single source]
- The Eve voice agent powered Grok 4's original full-audio interaction at launch. [SECONDARY, S5 — single source]

### Grok Build CLI (new to this section)
- Grok Build is xAI's agentic coding CLI, powered by Grok 4.5 at launch and Grok 4.6/4.7 thereafter. [SECONDARY, S3][SECONDARY, S26]
- It entered beta 2026-05-15 and left beta as 1.0 in early August 2026. [SECONDARY, S26 — single source]
- It was open-sourced 2026-07-15. [base §11 anchor; no new corroboration found in this wave — UNVERIFIED]
- Its API slug is `grok-build-0.1` with a 256K-token context window. [SECONDARY, S3][SECONDARY, S29]
- `grok-build-0.1` is API-key-only and is never advertised through OAuth/entitlement catalogs. [SECONDARY, S29 — single source]
- Grok 4.5 was free for a limited period inside Grok Build at launch. [SECONDARY, S10 — single source]
- The official models page alias rule, verbatim: "`<modelname>` is aliased to the latest stable version. `<modelname>-latest` is aliased to the latest version. `<modelname>-<date>` refers directly to a specific model release." [SECONDARY, S29 — single source]

### Colossus compute infrastructure (new to this section)
- Colossus 1: 230,000 GPUs (150,000 H100s, 50,000 H200s, 30,000 GB200s), ~500 MW, built at a former Electrolux factory in Memphis, Tennessee in 122 days. [SECONDARY, S14][SECONDARY, S13]
- The initial 100,000-GPU H100 cluster was operational within 122 days of groundbreaking in 2024, then doubled to 200,000 GPUs in 92 additional days. [SECONDARY, S15 — single source]
- Colossus 2: 550,000 NVIDIA GB200/GB300 accelerators, the world's largest single-site AI training deployment, consuming over 1 GW. [SECONDARY, S14][SECONDARY, S13]
- First Colossus 2 racks went online in early May 2026; Musk announced the milestone 2026-05-03. [SECONDARY, S14 — single source]
- Each Colossus 2 node houses two GPUs, pushing the total accelerator count past one million across both sites. [SECONDARY, S14 — single source]
- Colossus 2 core facility: 1M sq ft warehouse on Tulane Road, Shelby County, Tennessee, with two adjacent 100-acre sites acquired 2025-03-07. [SECONDARY, S18 — single source]
- Power strategy: a dedicated power island in Southaven, Mississippi — natural-gas turbines buffered by Tesla Megapacks plus a small solar farm — creating a private utility for the Tennessee compute. [SECONDARY, S17][SECONDARY, S18]
- On 2025-12-30 Musk revealed the purchase of a third Memphis building ("MACROHARDRR," ~500 MW), bringing total site capacity to nearly 2 GW. [SECONDARY, S13][SECONDARY, S15]
- As of February 2026 the Memphis complex housed ~555,000 NVIDIA GPUs acquired for approximately $18B. [SECONDARY, S13][SECONDARY, S15]
- Musk's stated targets: scale Colossus 2 to 1.5 GW within months, longer-term 2 GW; third-building GPU deployment Q2–Q3 2026 toward a 1M-GPU target. [SECONDARY, S14][SECONDARY, S15]
- A five-year target of 50M H100-equivalent GPUs by 2030 was reported alongside the Series E. [SECONDARY, S21 — single source]
- Sustainability measures: an $80M greywater recycling facility (Colossus Water Recycle Plant) on 13 leased acres, protecting 4.75B gallons of aquifer water annually; the world's largest Tesla Megapack deployment for zero grid impact at peak. [SECONDARY, S16 — single source]
- Cooling water usage: 13M gallons/day, supplied by the world's largest ceramic membrane bioreactor. [SECONDARY, S15 — single source]
- Thermal management uses direct-to-chip liquid cooling developed with Supermicro. [SECONDARY, S14 — single source]
- Musk claims xAI installs up to 300,000 GPUs per month. [SECONDARY, S14 — single source]
- Training of Grok 5 had already begun on Colossus as of the May 2026 reporting — new corroboration that Grok 5 is in training though still unshipped as of 2026-09-22. [SECONDARY, S14 — single source]

### xAI corporate and funding (new to this section)
- xAI was founded May 2023 by Elon Musk. [SECONDARY, S20 — single source]
- Seed round July 2023: $134.7M. [SECONDARY, S20 — single source]
- Series C November 2024: $6B at a $24B valuation (investors: a16z, Sequoia, Fidelity, Prince Alwaleed). [SECONDARY, S20 — single source]
- X merger: all-stock combination of xAI and X Corp in March/April 2025 at an implied ~$113B valuation (sources differ on March vs April). [SECONDARY, S20][SECONDARY, S23]
- Series E closed 2026-01-06: $20B raised at a $230B valuation, exceeding the original $15B target by 33%. [SECONDARY, S21][SECONDARY, S24][SECONDARY, S22][SECONDARY, S53][SECONDARY, S54][SECONDARY, S55]
- Series E investors: NVIDIA (strategic, reportedly up to $2B via chip-supply financing), Valor Equity Partners, StepStone Group, Fidelity, Qatar Investment Authority, Abu Dhabi's MGX, Baron Capital Group, Cisco Investments. [SECONDARY, S22][SECONDARY, S24][SECONDARY, S53][SECONDARY, S56]
- Bloomberg confirmed the round; Musk publicly called an early $15B-only report "False" on X. [SECONDARY, S24 — single source]
- On 2026-02-02 SpaceX acquired xAI in an all-stock transaction implying a ~$1.25T combined entity (SpaceX $1T, xAI $250B off the $230B round); the AI unit was rebranded SpaceXAI. A September 2026 breaking report describes the same $1.25T deal as newly announced ahead of a potential record IPO — the date discrepancy is preserved in the ledger. [SECONDARY, S15][SECONDARY, S20][SECONDARY, S10][SECONDARY, S58]
- One funding analysis estimates ~$500M ARR in early 2026 with a $2B full-year target — a ~460x revenue multiple. [SECONDARY, S20 — single source]
- Reported user reach: ~600M monthly active users across the X and Grok apps. [SECONDARY, S22 — single source]
- Total capital raised across all rounds: ~$26.7B. [SECONDARY, S20 — single source]
- Mid-June 2026: SpaceX announced acquisition of Cursor's parent Anysphere at a $60B valuation, expected to close Q3 2026. [SECONDARY, S10 — single source]
- xAI's 2025 year-in-review (via Series E reporting): Colossus I and II ended 2025 with over 1M H100 GPU equivalents; Grok Voice serving millions of users; ~600M monthly active users across X and Grok apps; Grok Voice in Tesla vehicles. [SECONDARY, S22][SECONDARY, S54][SECONDARY, S55]

