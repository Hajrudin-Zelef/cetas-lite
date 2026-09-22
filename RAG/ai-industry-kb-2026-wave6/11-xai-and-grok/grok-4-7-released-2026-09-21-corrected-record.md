---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/grok-4-7-released-2026-09-21-corrected-record
title: "Grok 4.7 (released 2026-09-21 — corrected record)"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["Anthropic", "ByteDance", "Microsoft", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2025-03-07", "2025-08-04", "2025-08-07", "2025-08-28", "2025-09", "2025-10", "2025-12-30", "2026-01", "2026-01-21", "2026-01-29", "2026-01-30", "2026-02", "2026-05", "2026-05-03", "2026-05-15", "2026-07-05", "2026-07-15", "2026-08", "2026-08-07", "2026-09-21", "2026-09-22", "2026-11-02"]
keywords: ["grok", "grok 4", "accelerator", "agent", "agentic", "astra", "benchmark", "benchmarks", "compute", "context window", "copilot", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5296, 5380]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: 0c92095fc672aa5b7dba2aa151298737ee1add5f1d2b4cdedc32e35db4c017f6
---

# Grok 4.7 (released 2026-09-21 — corrected record)

### Grok 4.7 (released 2026-09-21 — corrected record)
- Grok 4.7 was released 2026-09-21, 40 days after Grok 4.6 — correcting the prior brief's September 15 dating. [SECONDARY, S38][SECONDARY, S39][SECONDARY, S40]
- Context window: 500K tokens — correcting the prior brief's 1M figure. [VENDOR, S44][SECONDARY, S40]
- Pricing: $2.00/M input, $6.00/M output under 200K input tokens; $4.00/M input, $12.00/M output at or above 200K; cached input $0.50/$1.00. [VENDOR, S44][SECONDARY, S39]
- This corrects the prior brief's $2.50/$7.50 and $0.63 cached figures. [SECONDARY, S39][VENDOR, S44]
- Reasoning: configurable `low` / `medium` / `high` / `xhigh` — the prior brief's fifth "ultra" level is not corroborated. [SECONDARY, S40][VENDOR, S44]
- It runs on a larger base model than Grok 4.6 with a longer reinforcement-learning run on a harder task mix weighted toward multi-hour problems; xAI added self-verification of its own output during training. [SECONDARY, S39][SECONDARY, S42][SECONDARY, S75]
- Musk claimed pre-launch that 4.7 would run on a 2.1-trillion-parameter base with SpaceX company data in the training mix and "exceed all current models." [SECONDARY, S76 — single source]
- It was trained to natively understand the Grok Bot harness, improving conversational and general knowledge-work performance. [SECONDARY, S39][SECONDARY, S42]
- New safeguard stack: 3.3% pass-through rate for risky prompts on the internal HackerBench v0.3 test. [SECONDARY, S42 — single source]
- Vendor benchmark table (effort: 4.7 at xhigh, 4.6 at high, GPT-5.6 Sol and Fable 5.1 at max): CursorBench 4.0 46.3% (4.6: 40.4%, GPT-5.6 Sol: 41.7%, Fable 5.1 Max: 51.8%). [VENDOR, S43][SECONDARY, S39]
- DeepSWE v1.1: 71.0% at high effort (4.6: 65.2%, GPT-5.6 Sol: 72.7%, Fable 5.1 Max: 70.0%). [VENDOR, S43][SECONDARY, S39]
- EEBench (electrical engineering): 64.0% (4.6: 53.0%, GPT-5.6 Sol: 39.4%, Fable 5.1 Max: 56.4%). [VENDOR, S43][SECONDARY, S39]
- AA Briefcase v1.1 (multi-hour office work): 1,657 (4.6: 1,546, GPT-5.6 Sol: 1,487, Fable 5.1: 1,678). [VENDOR, S43][SECONDARY, S39]
- Harvey Legal Agent Benchmark: 19.6% (4.6: 15.8%, GPT-5.6 Sol: 2.5%, Fable 5.1 Max: 6.7%). [VENDOR, S43][SECONDARY, S39]
- HealthBench Professional: 56.7% (4.6: 48.5%, GPT-5.6 Sol: 60.5%, Fable 5.1 Max: 62.1%). [VENDOR, S43][SECONDARY, S39]
- Terminal-Bench 4.0: 38.0% (4.6: 20.3%, GPT-5.6 Sol: 37.3%, Fable 5.1 Max: 57.9%). [VENDOR, S43][SECONDARY, S39]
- GDPval: Elo 1,695 (Fable 5.1: 1,735, Grok 4.6: 1,605, GPT-6 Astra: 1,542). [VENDOR, S43][SECONDARY, S39]
- Artificial Analysis Intelligence Index: 46 at xhigh reasoning — +2 over Grok 4.6 (high), but below the then-leading score of 53. [SECONDARY, S41][SECONDARY, S38]
- With Grok Build, Grok 4.7 scores 56 on the AA Coding Agent Index, up from 47 for Grok 4.6 (dev.to also notes it passed GPT-5.6 Sol and trails only Fable 5.1, GPT-6 Astra, Opus 5). [SECONDARY, S38][SECONDARY, S79][SECONDARY, S76]
- Grok 4.7 uses ~81,000 output tokens per Intelligence Index task at xhigh, vs ~36,000 for Grok 4.6 (high) and ~27,000 for GPT-6 Astra (max); independent testing reports ~240M tokens for 4.7 at max vs ~94M for 4.6 at high — a token-cost caveat on the $2/$6 sticker price. [SECONDARY, S41][SECONDARY, S74]
- A Grok 4.7 Fast variant costs twice the standard rates and is not available through the public API. [SECONDARY, S40 — single source]
- Available at launch in Cursor, Grok Build, the Grok API, GitHub Copilot rollout, and third-party gateways. [SECONDARY, S42][SECONDARY, S38]
- Served at the same price and speed as Grok 4.6, per launch coverage. [SECONDARY, S42 — single source]


### Grok Code Fast 1 corroborated (2025-08-28)
- Launched 2025-08-28/29 as an agentic coding model built from scratch on a new architecture with a programming-heavy pre-training corpus and post-training on real-world PRs; free for a limited period (through early September 2025) via launch partners GitHub Copilot, Cursor, Cline, Roo Code, Kilo Code, opencode, Windsurf. [SECONDARY, S65][SECONDARY, S66][SECONDARY, S67]
- API pricing: $0.20/1M input, $1.50/1M output, $0.02/1M cached input; 256K context. Shipped in stealth under codename "sonic". [SECONDARY, S65][SECONDARY, S66][SECONDARY, S68]
- Vendor benchmark: 70.8% SWE-Bench-Verified subset via xAI's internal harness; combined benchmark + human evaluation methodology; xAI notes benchmarks may not capture end-to-end agentic usability. [VENDOR, S65][SECONDARY, S66]
- One 2026 community taxonomy notes Grok Code Fast 1 is "mostly historical now," with xAI's coding guidance reduced to the Grok 4.6 flagship line. [COMMUNITY, S69 — single source]

### Grok Imagine — image and video generation (new to this section)
- Grok Imagine was officially announced 2025-08-04 (beta, paid: SuperGrok and Premium+; Android access from 2025-08-07); xAI removed the paywall 2025-08-07 making image-to-video free for all Grok app users. One timeline places a separate October 2025 debut — the discrepancy is preserved. [SECONDARY, S59][SECONDARY, S61][SECONDARY, S37]
- It is integrated directly in the X app, giving it distribution across hundreds of millions of users. [SECONDARY, S37 — single source]
- Under the hood it uses xAI's proprietary Aurora model for text-to-image generation. [SECONDARY, S37][SECONDARY, S36]
- The Aurora engine is an autoregressive image/model network processing text, image, video, and audio tokens jointly — replacing diffusion-transformer approaches — which yields native audio-video sync in a single generation step; a January 2026 third-party API launch announcement independently describes Aurora as xAI's internal autoregressive image model. [SECONDARY, S36][SECONDARY, S81]
- On 2026-01-21 xAI extended maximum video length from 5 seconds to 10 seconds with improved visual quality, smoothness, and audio synchronization. [SECONDARY, S33 — single source]
- The Grok Imagine API launched 2026-01-29 (third-party API platform Pixazo announced availability 2026-01-30) for text-to-video, image-to-video, and prompt-driven video edits with synchronized audio. [SECONDARY, S35][SECONDARY, S81]
- API parameters: clip lengths 1–15 seconds, 480p or 720p resolution, aspect ratios including 16:9, 4:3, 1:1, 9:16, 3:4, 3:2, 2:3 (and 20:9 per launch posts). [SECONDARY, S35 — single source]
- Requests are processed as deferred jobs with SDK auto-polling. [SECONDARY, S35 — single source]
- xAI claimed the #1 position in Artificial Analysis text-to-video rankings and reported human side-by-side preference on IVEBench at 1280×720 versus Kling o1 and Runway Aleph. [VENDOR, S35 — single source]
- In late May 2026 Grok Imagine video debuted atop the Artificial Analysis Video Arena image-to-video leaderboard (Elo 1404 ±6), displacing ByteDance Seedance 2.0. [SECONDARY, S36 — single source]
- Elon Musk posted "Done with Grok Imagine" on 2026-07-05, signaling the end of the core development cycle and a move to a polished/stable phase. [SECONDARY, S37 — single source]
- Grok Imagine Image 2.0: in the August 7, 2026 Arena snapshot cited by xAI, it ranked second globally in both text-to-image generation and image editing. [SECONDARY, S32 — single source]
- The Imagine API documents five model ids: `grok-imagine-image`, `grok-imagine-image-2.0`, `grok-imagine-image-quality`, `grok-imagine-video`, `grok-imagine-video-1.5`. [SECONDARY, S29 — single source]
- `grok-imagine-image-quality` is scheduled to retire 2026-11-02 in favor of `grok-imagine-image-2.0`, with a dedicated November 2 migration guide. [SECONDARY, S29 — single source]
- Grok Imagine video 1.5 generates 720p video at 24fps with native synchronized audio and topped third-party image-to-video leaderboards on release. [SECONDARY, S12 — single source]
- xAI also shipped 21 multilingual voices alongside the Imagine updates. [SECONDARY, S12 — single source]

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

