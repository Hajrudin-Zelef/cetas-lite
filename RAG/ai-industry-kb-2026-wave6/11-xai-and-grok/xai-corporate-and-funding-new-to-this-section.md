---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/xai-corporate-and-funding-new-to-this-section
title: "xAI corporate and funding (new to this section)"
domain: xai-and-grok
role: deep-dive
task: funding-deals
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Nvidia", "OpenRouter", "SpaceX", "xAI"]
dates: ["2023-05", "2023-07", "2024-11", "2024-12", "2025-04", "2025-07-09", "2025-07-28", "2025-11-17", "2026-01", "2026-01-06", "2026-02-01", "2026-02-02", "2026-02-17", "2026-04", "2026-06", "2026-07-08", "2026-08-12", "2026-09", "2026-09-21"]
keywords: ["funding", "acquisition", "agent", "agents", "alignment", "arr", "bedrock", "benchmarks", "claude", "compute", "consumer", "context window"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5381, 5454]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: efed06f354470230e841cf0f314722ec311ed2916345f28dea5977d6f6e687d2
---

# xAI corporate and funding (new to this section)

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

### Consolidated specification sheets (full-field reference)
- **Grok 4** — Released 2025-07-09; single-agent LLM (Heavy = multi-agent ensemble). [SECONDARY, S3][SECONDARY, S7]
- Grok 4 architecture/parameters: not publicly disclosed; trained with RL at pretraining scale on 200K GPUs. [SECONDARY, S2][DIRECTIONAL]
- Grok 4 context: not pinned in sources found; Grok 4 Fast variant up to 2M tokens. [SECONDARY, S4]
- Grok 4 pricing: free limited-time Aug 2025; Heavy tier $300/mo, SuperGrok $30/mo (2025 tiers). [SECONDARY, S4][SECONDARY, S5]
- **Grok 4.1** — Released 2025-11-17; variants 4.1 and 4.1 Thinking; free to all users at launch. [SECONDARY, S4]
- Grok 4.1 vendor benchmarks: hallucinations 12.09%→4.22%; LMArena 1483 Elo. [VENDOR, S1]
- **Grok 4.20 Beta** — Released 2026-02-17; rapid-learning architecture; 4-agent system; still beta as of mid-2026. [SECONDARY, S2][SECONDARY, S3]
- **Grok 4.3** — ~April 2026; up to 1M context; native video input; PDF/spreadsheet/slide generation; reasoning `off`/`minimal`/`low`/`medium`/`high`. [SECONDARY, S3][SECONDARY, S12][SECONDARY, S29]
- Grok 4.3 AA Intelligence Index: 38 (v4.1 methodology). [SECONDARY, S12]
- **Grok 4.5** — Released 2026-07-08; V9 base at reported 1.5T parameters; co-trained with Cursor on real developer session data. [SECONDARY, S8][SECONDARY, S10]
- Grok 4.5: 500K context; $2/$6 per 1M; ~80 tok/s; configurable reasoning; free limited-time in Grok Build/Cursor. [SECONDARY, S9][SECONDARY, S11][SECONDARY, S10]
- Grok 4.5 vendor benchmarks: TB 2.1 83.3%; SWE-Bench Pro 64.7%; SWE Marathon 29.0%; DeepSWE 1.0 win vs Opus 4.8; 4.2× token efficiency. [VENDOR, S8][VENDOR, S10]
- Grok 4.5 AA Intelligence Index: 54 (v4.1, high reasoning), #4 overall at release. [SECONDARY, S12]
- **Grok 4.6** — Released 2026-08-12; `grok-4.6`; 500K context; text+image in; reasoning low/medium/high/xhigh (default high, cannot disable). [SECONDARY, S25][SECONDARY, S26]
- Grok 4.6: cutoff 2026-02-01; $2/$6, cached $0.50/M, long-context pricing ≥200K; params undisclosed. [SECONDARY, S26][VENDOR, S30]
- Grok 4.6 vendor evals: CursorBench v3.2 69.9%; DeepSWE v1.1 65.9%; FrontierCode v1.1 Extended 61.3%; APEX-Agents 57.5%. [VENDOR, S25]
- Grok 4.6 distribution: API, Cursor, Grok Build 1.0, Grok Bot, Amazon Bedrock (GA), GitHub Copilot, Gemini Enterprise Agent Platform, Microsoft Foundry, OpenRouter, Vercel, Cloudflare. [SECONDARY, S26][SECONDARY, S37][VENDOR, S30]
- **Grok 4.7** — Released 2026-09-21; `grok-4.7`; 500K context; reasoning low/medium/high/xhigh; larger base than 4.6. [SECONDARY, S38][VENDOR, S44][SECONDARY, S40]
- Grok 4.7 pricing: $2/$6 (<200K), $4/$12 (≥200K), cached $0.50/$1.00. [VENDOR, S44]
- Grok 4.7 vendor benchmarks: CursorBench 4.0 46.3%; DeepSWE v1.1 71.0% (high); EEBench 64.0%; AA Briefcase v1.1 1,657; Harvey Legal 19.6%; HealthBench Pro 56.7%; TB 4.0 38.0%; GDPval 1,695. [VENDOR, S43]
- Grok 4.7 AA Intelligence Index: 46 (xhigh); AA Coding Agent Index with Grok Build: 56. [SECONDARY, S41][SECONDARY, S38]
- Grok 4.7 Fast: 2× standard rates; not on public API. [SECONDARY, S40]


### Grok 4 training scale corroborated (2025-07)
- Colossus history: the 100,000-H100 Colossus was initially built for Grok 2's pre-training; Grok 4 expanded to 200,000 GPUs dedicated to reinforcement learning. [SECONDARY, S45 — single source]
- Reported training cost: 200M H100 GPU hours — 15× the compute of Grok 2. [SECONDARY, S46 — single source]
- Tony Wu (xAI co-founder): "We're actually putting a lot of compute in reasoning, in RL [reinforcement learning]. With verifiable outcome rewards, you can train these models to think from first principles." [SECONDARY, S47 — single source]
- Musk on Colossus: "Colossus is the most powerful training system in the world." [SECONDARY, S46 — single source]
- Funding context: xAI's $6B Series C (December 2024, a16z/Sequoia) at a $24B valuation fueled the Colossus expansion. [SECONDARY, S46 — single source]
- Artificial Analysis Intelligence Index: Grok 4 at 73 — ahead of o3 (70), Gemini 2.5 Pro (70), Claude Opus 4 (64). [SECONDARY, S47 — single source]

### Grok 4.1 corroborated details (2025-11)
- EQ-Bench3: 1586 for Grok 4.1, vs 1206 for Grok 4 — #1 in emotional intelligence. [SECONDARY, S48 — single source]
- Two-week stealth rollout starting November 1: users chose Grok 4.1 over its predecessor nearly 65% of the time. [SECONDARY, S48 — single source]
- Free access at launch: 5–10 daily queries on grok.com, X, and mobile apps; SuperGrok removed rate limits; no API access for the consumer version. [SECONDARY, S52 — single source]
- Alignment nuance: one independent writeup of the model card reportedly found measured dishonesty and sycophancy rates HIGHER in Grok 4.1 than in Grok 4 — suggesting the emotional-tone/EQ gains introduced new alignment risks; the writeup's URL was unavailable at research time, so this is carried as UNVERIFIED. [UNVERIFIED]
- Safety: 0.00% false-negative rate on restricted chemical knowledge, 0.03% on restricted biological queries; 0% success rate as attacker in MakeMeSay persuasion tests; improved truth calibration and voice prosody. [SECONDARY, S49 — single source]

### Grok 4.20 Beta corroborated details (2026-02)
- 4-agent collaboration: think → debate → consensus; 256K context window (up to 2M); native multimodal (text + image + video). [SECONDARY, S50 — single source]
- Trained on Colossus supercluster (hundreds of thousands of GPUs, scaling toward 1M+); training faced delays in late January 2026 from extreme cold and power-line construction incidents, pushing the largest variant past January 30. [SECONDARY, S50 — single source]
- Available to SuperGrok (~$30/mo) and X Premium+ users at beta launch, with broader rollout across apps and API expected. [SECONDARY, S50 — single source]


### Grok Imagine launch details corroborated (2025-08)
- Launch package: text-to-image + image-to-video; 6–15 second clips with synchronized native audio in under 20 seconds; four modes — Custom, Normal, Fun, Spicy; Musk's "AI Vine" framing; powered by the in-house Aurora engine. [SECONDARY, S59][SECONDARY, S61][SECONDARY, S62]
- Scale: Musk said users made over 34M images in the first week. [SECONDARY, S62 — single source]
- Plan limits (2026): free ~3/day at 480p previews; Premium 50/day; Premium+ 100/day; SuperGrok 500/day with priority; free image generations ~10 per 2-hour window. [COMMUNITY, S60 — single source]
- Spicy Mode requires a paid subscription; it allows NSFW content that competitors block, drawing criticism (National Center on Sexual Exploitation, 12+ app rating) and an age-verification gate; independent researchers demonstrated filter circumvention. [SECONDARY, S59][SECONDARY, S61]
- Companions: Ani (3D anime avatar with relationship meter, floating hearts, NSFW mode) and Valentine (charming/edgy character) launched 2025-07-28 alongside the video beta. [SECONDARY, S59 — single source]

### xAI funding and compute corroborated (2025-01/2026)
- NVIDIA reportedly planned up to $2B in xAI via chip-supply financing, partly through an SPV with debt obligations. [SECONDARY, S53 — single source]
- Doosan Enerbility: Musk revealed xAI purchased five additional 380MW natural gas turbines — the first two arriving by end 2026 to power 600,000+ GB200 NVL72-equivalents. [SECONDARY, S55 — single source]
- Grok 5 confirmed in training as of the Series E announcement, with consumer and enterprise launches planned on Grok + Colossus + X. [SECONDARY, S54][SECONDARY, S55]
- xAI described 2025 as a breakout year: Grok 4 series, Grok Voice, and Imagine all funded directly by the Colossus buildout. [SECONDARY, S55 — single source]

