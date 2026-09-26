---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/grok-4-6-released-2026-08-12-details-beyond-base-section
title: "Grok 4.6 (released 2026-08-12; details beyond base section)"
domain: xai-and-grok
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2025-08-28", "2025-09", "2026-02-01", "2026-08", "2026-08-12", "2026-08-18", "2026-09-21"]
keywords: ["grok", "grok 4", "agent", "agentic", "agents", "astra", "bedrock", "benchmark", "benchmarks", "context window", "copilot", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5275, 5327]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: f49939b62df305d5f4b8799caf747f9b0503766dc0236c2ff640551a996e6842
---

# Grok 4.6 (released 2026-08-12; details beyond base section)

### Grok 4.6 (released 2026-08-12; details beyond base section)
- Grok 4.6 was released 2026-08-12 as a SpaceXAI reasoning model for coding, long-running agents, interactive/visual projects, and knowledge work. [SECONDARY, S25][SECONDARY, S26]
- API identifier: `grok-4.6`; text and image input, text output. [SECONDARY, S25][SECONDARY, S26]
- Context window: 500,000 tokens; prompts at or above 200,000 tokens use long-context pricing. [SECONDARY, S25][SECONDARY, S26]
- Four configurable reasoning-effort levels: `low`, `medium`, `high` (default), `xhigh`; reasoning cannot be disabled. [SECONDARY, S25][SECONDARY, S26]
- The `xhigh` level requires xai-sdk 1.18.0 or later (older SDKs reject the value client-side). [SECONDARY, S27 — single source]
- Knowledge cutoff: 2026-02-01. [SECONDARY, S26 — single source]
- Pricing: $2/M input, $6/M output; cached input $0.50/M. [VENDOR, S30 — single source]
- Vendor evaluations: CursorBench v3.2 69.9%; DeepSWE v1.1 65.9%; FrontierCode v1.1 Extended 61.3%; APEX-Agents 57.5%. [VENDOR, S25 — single source]
- Training: a longer supplemental training run than Grok 4.5 received — not a larger base model — using curated model-generated data, high-quality engineering data, and an improved optimizer and training recipe. [SECONDARY, S31 — single source]
- Grok 4.5 was used to regenerate supervised fine-tuning trajectories across reasoning-effort levels, agent harnesses, and domains (STEM, software engineering, knowledge work), with problematic traces filtered by model-based checks. [SECONDARY, S31 — single source]
- Reinforcement learning followed in agentic environments covering knowledge work, general coding, web development, computer-aided design, and kernel optimization. [SECONDARY, S25][SECONDARY, S31]
- Reported behavioral effect: more self-testing and verification on longer trajectories (vendor observation from internal testing, not independently measured). [VENDOR, S31 — single source]
- No parameter count was published for Grok 4.6. [SECONDARY, S31 — single source]
- The official SpaceXAI models overview recommends Grok 4.6 for code and chat and calls it the most intelligent and fastest model xAI has built. [SECONDARY, S29 — single source]
- Generally available on Amazon Bedrock (Bedrock model card dated 2026-08-18); xAI's announcement frames Bedrock as a first-class channel. Bedrock coverage describes Grok 4.6 as an MoE design of roughly 1.5T parameters with a 500K context window and four selectable reasoning levels. [VENDOR, S30][SECONDARY, S77][SECONDARY, S78]
- Later in August 2026 it was extended to GitHub Copilot, the Gemini Enterprise Agent Platform, and Microsoft Foundry. [SECONDARY, S37 — single source]
- Live in Cursor, Grok Build 1.0, and Grok Bot agents at launch; distributed through OpenRouter, Vercel, and Cloudflare. [SECONDARY, S26 — single source]
- xAI documentation recommends the `prompt_cache_key` parameter to keep multi-turn request routing stable and cache hits reliable. [SECONDARY, S26 — single source]
- Third-party provider Pi maps its `minimal` thinking level onto xAI's `low` for Grok 4.6. [SECONDARY, S28 — single source]

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

