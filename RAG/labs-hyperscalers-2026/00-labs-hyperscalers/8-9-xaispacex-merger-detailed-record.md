---
id: labs-hyperscalers-2026/00-labs-hyperscalers/8-9-xaispacex-merger-detailed-record
title: "8.9 xAI–SpaceX merger — detailed record"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-02-02", "2026-06-12"]
keywords: ["merger", "agent", "agentic", "agents", "astra", "bedrock", "benchmarks", "claude", "compute", "copilot", "cost", "distribution"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1164, 1191]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: cca114b66febaae37a524bf5dbc7c28adf2c55ea6ab0e85d3c9531f11c54ef37
---

# 8.9 xAI–SpaceX merger — detailed record

**Grok 4.5 (Jul 8, 2026):** Built on 1.5T-param MoE "V9" foundation (3× the 500B v8-small that had served production Grok traffic in early 2026) [vendor-reported]. Vendor-reported [vendor-reported]: **64.7% SWE-Bench Pro** (behind Opus 4.8 69.2%, Fable 5 80.4%, ahead of GPT-5.5 58.6%); **83.3% Terminal-Bench 2.1** (tying GPT-5.5 83.4%); #1 AutomationBench-AA at 51% at ~1/4 cost per task of rivals; ~14k output tokens/task vs ~67k for Opus 4.8 (token-efficiency pitch). Independent [independent]: Artificial Analysis Intelligence Index 54 — #4 overall behind Claude Fable 5 (60), Claude Opus 4.8 (56), GPT-5.5 (55); #1 of 28 on agentic tool use; $0.31–$0.49 per completed Intelligence-Index task [secondary]. **Training-data controversy:** Grok 4.5 used real Cursor developer session data; users reportedly worked in Cursor for ~a year without knowing sessions would train a model for a company that hadn't acquired their tool at the time [secondary — one reporter's finding, unverified by other outlets].

**Grok 4.6 (Aug 12, 2026):** "Significantly better at difficult tasks and knowledge work," trained for long-running agents with curated model-generated reasoning data + agentic RL [secondary]. Vendor-reported [vendor-reported]: AA Intelligence Index 61 (matching GPT-5.6 Sol; trailing Claude Opus 5 and Fable 5 Max); GDPVal-AA v2 Elo 1753 (behind only Claude Opus 5); CursorBench v3.2 69.9%; DeepSWE v1.1 65.9%; FrontierCode v1.1 Extended 61.3%; APEX-Agents 57.5%; APEX-SWE 56.4%; Terminal-Bench v3.0 26%; AA-Briefcase 1577; Harvey LAB (Vals) 15.8%. Trade-offs reported by third parties: strongest long-horizon agent xAI had shipped, but regressions in speed/token-efficiency vs 4.5 (per-task cost roughly doubled via token inflation; cached-input price $0.03 → $0.05) [secondary — single-source assessment]. Independent eval that did NOT make the launch post: AA-Omniscience 48.2% accuracy; 65.7% non-hallucination rate (~1 in 3 inventions on unknowns) — flagged by a Medium analysis, Aug 2026 [unverified-independent].

**Grok 4.7 (Sep 21, 2026) — current flagship:** Architecture: new larger base model, **2.1T parameters (+40% vs 4.6's 1.5T)**, longer RL run on harder multi-hour tasks; 500k context; same $2/$6 pricing [vendor-reported]. Training-data differentiator: supplemental SpaceX engineering data — Starlink telemetry, manufacturing records, engineering failure logs — pitched as reasoning better about hardware/physical systems [secondary]. Timeline: Musk walked back release estimates ≥5 times since late July [secondary]. Vendor-reported benchmarks [vendor-reported]: CursorBench 4.0 46.3% (vs 40.4% for 4.6; vs GPT-5.6 Sol 41.7%), DeepSWE v1.1 high-effort 71.0% (vs 65.2% for 4.6; vs GPT-5.6 Sol 72.7%). Early independent [independent]: Artificial Analysis first pass (same day) — Intelligence Index 46 (+2 vs Grok 4.6), Coding Agent Index 56 (+9 vs 4.6), AA-Briefcase Elo 1,657 (+111 vs 4.6); GDPval and AA-Briefcase: 2nd to Claude Fable 5.1; EEBench: 2nd to GPT-6 Astra [secondary]. **INDEX-VERSION DISCREPANCY: Grok 4.6 = 61 (Aug 2026, earlier index version) vs Grok 4.7 = 46 (Sep 2026, index v4.3.2). These numbers are NOT directly comparable across index revisions.** Live at launch in: Grok app, Cursor, Grok Build, xAI API, third-party routers/cloud platforms; no waitlist [secondary].

### 8.9 xAI–SpaceX merger — detailed record

- **Announcement:** Feb 2, 2026 — SpaceX acquired xAI in an all-stock deal [secondary: Reuters, Bloomberg] https://www.reuters.com/business/musks-spacex-merge-with-xai-combined-valuation-125-trillion-bloomberg-news-2026-02-02/. First reported: Reuters, week of Jan 26, 2026; Bloomberg first reported the plan to merge ahead of IPO [secondary].
- **Deal structure:** all-stock merger; xAI investors receive **0.1433 shares of SpaceX per xAI share**; some xAI executives could opt for cash at **$75.46/xAI share**; combined entity priced ~$527/share privately [secondary: Reuters, Economic Times, SlashGear].
- **Valuation:** SpaceX $1T + xAI $250B = **$1.25T combined** — described as one of the largest tech mergers ever [secondary: Reuters, Bloomberg].
- **Scope:** folded xAI, Grok, and the X social platform under one SpaceX cap table; Tesla NOT part of the deal despite speculation [secondary].
- **Stated rationale (Musk):** (1) funding solution for xAI's compute-heavy cash burn; (2) long-term thesis that **space-based data centers** become the cheapest place to run AI compute within 2–3 years, bypassing Earth's power/cooling/environmental constraints [secondary].
- **Status as of Sep 22, 2026:** deal closed (Feb 2026); entity now publicly branded **"SpaceXAI"** in product announcements (Grok 4.5, 4.6, 4.7 press) [secondary]; combined company went public June 12, 2026 (see §8.11).
- **Implications:** Grok Bot (launched Aug 11, 2026) — "a team of always-on agents with their own computer" — shows operational leverage of the merged stack [secondary]. Grok 4.6/4.7 distribution expanded to Microsoft Foundry, Amazon Bedrock, GitHub Copilot, Gemini Enterprise Agent Platform — enterprise reach beyond the pre-merger xAI API [secondary]. Orbital data centers remain a stated ambition, not demonstrated capacity — no satellite-compute hardware milestones reported as of Sep 22, 2026.

### 8.10 Colossus compute cluster — detailed

- **Colossus 1** (Memphis, TN): ~230,000 GPUs (incl. 30k GB200s), ~500 MW — operational [secondary] https://introl.com/blog/xai-colossus-2-gigawatt-expansion-555k-gpus-january-2026.
- **Colossus 2** (Memphis/Southaven): ~550,000 GB200s/GB300s targeting ~1 GW; "world's first gigawatt-scale AI training supercluster" per company descriptions [secondary].
- **Dec 30, 2025:** Musk announced purchase of a third Memphis building ("MACROHARDRR"), bringing the complex toward **~2 GW total capacity** and **555,000+ GPUs (~$18B spend)** [secondary, citing Musk posts; Reuters reported ultimate target of 2 GW / 1,000,000 GPUs].
- **Power:** purpose-built private power — acquired former Duke Energy natural gas plant site in Southaven, MS (mid-2025); gas turbines + Tesla Megapacks + small solar; near TVA Southaven gas plant [secondary].
- **Regulatory/legal (2026):** on Apr 14, 2026, the Mississippi NAACP + environmental groups (SELC/Earthjustice) filed a **Clean Air Act lawsuit** alleging **27 unpermitted methane gas turbines** at Colossus 2 in Southaven, MS — a ~200 MW makeshift plant feeding the training cluster with no federal air permit; framed as the first major federal test of AI compute vs the Clean Air Act [secondary] https://tech-insider.org/xai-colossus-2-naacp-lawsuit-illegal-gas-turbines-memphis-2026/. Outcome as of Sep 2026 not found — **unresolved**.
- **Financing:** Bloomberg reported (late 2025) a ~$20B equity+debt package for Colossus 2 with **Nvidia as strategic investor (up to $2B)** via an SPV that buys Nvidia GPUs and leases them to xAI; $7–8B equity + up to $12B debt [secondary: Tom's Hardware citing Bloomberg] https://www.tomshardware.com/pc-components/gpus/nvidia-backs-20-billion-xai-chip-deal.
- **Vendors:** Dell ~$5B GPU-server deal reported; Cisco networking hardware for Colossus II buildout; Introl as infrastructure contractor [secondary].
- **Second site:** xAI reportedly built a second data center in Atlanta (~$700M in chips/equipment) [secondary — thin sourcing, treat as unverified].
- **Forward guidance:** possible expansion toward 3 GW by late 2026 discussed [secondary]; one vendor-adjacent source claims 1M GPUs target [unverified].

