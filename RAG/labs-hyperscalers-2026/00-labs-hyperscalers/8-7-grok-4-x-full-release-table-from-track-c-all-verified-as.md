---
id: labs-hyperscalers-2026/00-labs-hyperscalers/8-7-grok-4-x-full-release-table-from-track-c-all-verified-as
title: "8.7 Grok 4.x — full release table (from Track C; all verified as shipped)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-02-02", "2026-05", "2026-05-15", "2026-06-12"]
keywords: ["grok", "grok 4", "agent", "agentic", "agents", "agi", "astra", "bedrock", "benchmark", "benchmarks", "claude", "compute"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1139, 1191]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 0db87efd1878243d7ef678fd66c164d6b88217d0ce63a168e686414997e66ea0
---

# 8.7 Grok 4.x — full release table (from Track C; all verified as shipped)

### 8.7 Grok 4.x — full release table (from Track C; all verified as shipped)

| Model | Release date | Context | API price (in/out per 1M) | Notes |
|---|---|---|---|---|
| Grok 4 / Grok 4 Heavy | Jul 9–10, 2025 | 256k | $3 / $15 | Single-agent + multi-agent (Heavy) tiers; SuperGrok Heavy $300/mo introduced [secondary] |
| Grok 4.1 / 4.1 Thinking | Nov 17, 2025 | n/a | $0.20 / $0.50 ($0.05 cached) | Silent rollout Nov 1–14, 2025; focus on EQ, creative writing, ~3× hallucination cut [secondary] |
| Grok 4.20 (= "Grok 4.2") | Feb 17, 2026 | 256k | $1.25 / $2.50 | Public beta; "rapid-learning" architecture, 4-agent parallel reasoning, medical-document feature. Branding "4.2" vs "4.20" used inconsistently across sources [secondary] |
| Grok 4.3 | Beta Apr 17, 2026 (SuperGrok Heavy); API Apr 30, 2026 | 1M | $1.25 / $2.50 ($0.20 cached) | Video input (up to 5 min), native PDF/XLSX/PPTX output, reasoning always-on; 8 legacy models retired May 15, 2026 [secondary] |
| Grok 4.5 | Jul 8, 2026 | 500k | $2 / $6 | 1.5T-param MoE "V9" architecture; co-trained with Cursor session data [vendor-reported] |
| Grok 4.6 | Aug 12, 2026 | 500k | $2 / $6 | Long-running agents focus; distributed to Copilot, Bedrock, Foundry, Gemini Enterprise Agent Platform later in August [secondary] |
| Grok 4.7 | **Sep 21, 2026** | 500k | $2 / $6 | 2.1T params; supplemental SpaceX engineering-data training; released by post-merger "SpaceXAI" [secondary] |
| Grok 5 | Not released | n/a | n/a | Training on full SpaceX historical corpus targeted before year-end [unverified — ByteVyte, single source] |

**Version check:** Grok 4, 4.1, 4.2/4.20, 4.3, 4.5, 4.6, 4.7 are all real shipped versions. Numbering skips 4.4 (no public release under that name found — likely internal). "4.20" and "4.2" refer to the same Feb 2026 release.

### 8.8 Grok benchmarks by model (vendor-reported vs independent)

**Grok 4 (Jul 2025):** Vendor-reported (xAI livestream) [vendor-reported]: HLE 25.4% (Grok 4, no tools), 44.4% (Grok 4 Heavy, multi-agent + tools); GPQA 87.5–88.9%; AIME Heavy 100%; ARC-AGI-2 16.2% (Heavy) / 15.9%. Independent [independent]: Artificial Analysis Intelligence Index 73, ahead of OpenAI o3 (70), Gemini 2.5 Pro (70), Claude Opus 4 (64) at the time.

**Grok 4.1 (Nov 2025):** Vendor claims [vendor-reported]: LMArena Text Arena #1 at 1,483 Elo (31 pts ahead of nearest rival); EQ-Bench3 1,586 (vs 1,206 for Grok 4); hallucination rate ~12.09% → ~4.22% on real-world queries; FActScore errors 9.89% → <3%. One day after the consumer launch, xAI added API access at $0.20/1M input ($0.05 cached) / $0.50/1M output — notably cheap, likely an enticement/price-cut play [secondary].

**Grok 4.20/4.2 (Feb 2026):** 4-agent parallel architecture ("adversarial consensus"; up to 16 agents in Heavy mode), claimed 65% hallucination reduction (12% → 4.2%), medical-document analysis via photo upload [secondary]. Medical feature had no published clinical validation as of Feb 2026 [secondary]. Early independent estimate of LMArena Elo 1505–1535 (provisional, never confirmed) [unverified].

**Grok 4.3 (Apr–May 2026):** Pricing cut at launch: output $15.00 → $2.50 (−83%), input $3.00 → $1.25 (−58%) vs the 4.20 era [vendor-reported]. Independent [independent]: Artificial Analysis Intelligence Index 53 (some reports cite 38 for "Grok 4.3 (high)" — apparent index-version/config discrepancy); ranked #37 of 154, behind GPT-5.5 (60) and Claude Opus 4.7 (57) [secondary]. AA put it on the Pareto frontier of cost-per-intelligence; benchmark-suite run cost $395, ~20% below Grok 4.20 despite 44% more output tokens [secondary].

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

