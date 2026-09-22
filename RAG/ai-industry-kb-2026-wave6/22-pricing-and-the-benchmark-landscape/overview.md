---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/overview
title: "§22. Pricing and the Benchmark Landscape"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "Cohere", "DeepSeek", "Glasswing", "Google", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "StepFun", "Z.ai", "xAI"]
dates: ["2026-02-19", "2026-04-23", "2026-06-24", "2026-07-09", "2026-07-16", "2026-07-21", "2026-08-11", "2026-08-12", "2026-08-18", "2026-08-21", "2026-08-26", "2026-09", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-07", "2026-09-10", "2026-09-21", "2026-11-21", "2026-12-31"]
keywords: ["benchmark", "pricing", "agent", "astra", "claude", "cohere", "cost", "deepseek", "fable 5", "gemini", "glm", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10484, 10558]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 3afc4c2dbb80fa689bfa67d8f3de993d096a74eed4cd52d14b0e0723f71facf2
---

# §22. Pricing and the Benchmark Landscape

Keywords: API pricing, USD per 1M tokens, cache-read pricing, context cliffs, free tiers, Artificial Analysis Intelligence Index v4.3, LMArena, SWE-bench, Terminal-Bench 4.0, FACTS Suite, Vals AI

## Summary

The 2026 price war's defining feature is a **20× spread inside one model generation** (GPT-5.6 Sol $4/$20 vs Luna $0.20/$1.20). The cheapest frontier-class API as of September 2026 is **DeepSeek V4.1 Flash at off-peak $0.15/$0.60** — $0.27 per Artificial Analysis Index task, the cheapest measured on the board [VENDOR]. Price moves are dated because they decay in weeks: Luna fell from $1/$6 (GA 2026-07-09) to $0.20/$1.20 by early September; Sol was cut 20/33% on 2026-08-21 (promo through 2026-11-21); Sonnet 5's planned 2026-09-01 rise to $3/$15 was cancelled (permanent $2/$10); GPT-6 Astra launched 2026-09-03 at $10/$50 with a 272K-token cliff [VENDOR]. Context-window pricing tiers are standard (Grok 4.7 and Gemini 3.1 Pro both double at 200K). The benchmark world had its most volatile week ever in early September 2026: Artificial Analysis revised the Intelligence Index **three times in under a week (v4.1.1 on 2026-09-03 → v4.2 on 2026-09-04 → v4.3 on 2026-09-07)** — **no cross-version comparison is valid** [SECONDARY]. On v4.3: Claude Fable 5.1 and GPT-6 Astra tie at **53**; among open weights, Qwen3.8-Max (0902) leads China at **45** [SECONDARY]. SWE-bench Verified is saturated (Opus 5 at 97.0%); SWE-bench Pro's official Scale board is led by Muse Spark 1.1 + mini-SWE-agent at 61.5% [SECONDARY]. Terminal-Bench 4.0 (66 tasks): Fable 5.1 57.9%, official #1 — not comparable to TB 2.1 scores [SECONDARY].

## Key dated facts

### Full price table, September 2026 (USD per 1M tokens; cache = cache-read/hit)

All rows via published rate cards or vendor-sourced trackers; each row pinned to its price date because 2026 prices decay in weeks.

| Model | Input /1M | Output /1M | Cache read | Price date | Access |
|---|---|---|---|---|---|
| GPT-6 Astra | $10.00 | $50.00 | $1.00 | 2026-09-03 (launch) | Closed API [VENDOR] |
| Claude Fable 5.1 | $10.00 | $50.00 | $0.25 | Sept 2026 | Closed API [VENDOR] |
| Claude Mythos 5.1 | $10.00 | $50.00 | $0.25 | 2026-09-01 | Restricted (Glasswing) [VENDOR] |
| Claude Opus 5 | $5.00 | $25.00 | $0.50 | Jul 2026 | Closed API [VENDOR] |
| GPT-5.5 / 5.5 Pro | $5.00 / $30.00 | $30.00 / $180.00 | $0.50 | 2026-04-23 | Closed API [VENDOR] |
| GPT-5.6 Sol | $4.00 (promo; std $5.00) | $20.00 (promo; std $30.00) | $0.40 | 2026-08-21 cut, promo thru 2026-11-21 | Closed API [VENDOR] |
| Gemini 3.1 Pro (≤200K) | $2.00 | $12.00 | $0.20 | 2026-02-19 | Closed API [VENDOR] |
| GPT-5.6 Terra | $2.00 | $12.00 | $0.20 | 2026-09-02 | Closed API [VENDOR] |
| Claude Sonnet 5 | $2.00 (intro, permanent) | $10.00 | $0.20 | 2026-08-11 (rise cancelled) | Closed API [VENDOR] |
| Grok 4.7 (≤200K) | $2.00 | $6.00 | — | 2026-09-21 (launch) | Closed API [VENDOR] |
| Grok 4.6 (short ctx) | $2.00 | $6.00 | $0.50 | 2026-08-12 | Closed API [VENDOR] |
| Kimi K3 | $3.00 | $15.00 | — | 2026-07-16 | API + Modified-MIT weights [VENDOR] |
| Qwen3.8-Max | $2.00 | $6.00 | — | 2026-08-12 | API + custom-license weights [VENDOR] |
| GLM-5.3 | $1.40 | $4.40 | 81% discount | 2026-08-18 | API + GLM-5.3-License weights [VENDOR] |
| Muse Spark 1.3 | $1.25 | $4.25 | — | 2026-09-02 | Closed API [VENDOR] |
| Claude Haiku 4.5 | $1.00 | $5.00 | $0.10 | Sept 2026 | Closed API [VENDOR] |
| StepFun Step 5 Preview | $1.00 | $2.70 | $0.05 | Sept 2026 | API-only [SECONDARY] |
| MiMo-V2.5-Pro | $0.80 | $3.20 | $0.16 (80% off) | Sept 2026 | API + MIT weights [SECONDARY] |
| Gemini 3.7/3.8 Flash | $0.75 | $3.75 | $0.075 | intro thru 2026-12-31 | Closed API [VENDOR] |
| GPT-5.4-mini | $0.75 | $4.50 | $0.075 | Sept 2026 | Closed API [SECONDARY] |
| Gemini 3.5 Flash-Lite | $0.30 | $2.50 | — | Sept 2026 | Closed API [SECONDARY] |
| Qwen3.7-Plus | $0.276 | — | — | Sept 2026, Global ≤256K non-thinking | Closed API [VENDOR] |
| GPT-5.6 Luna | $0.20 | $1.20 | $0.02 | 2026-09-02 (was $1/$6 at GA 2026-07-09) | Closed API [VENDOR] |
| GLM-5.3-Flash | $0.15 | — | — | 2026-08-26 | API + MIT weights [VENDOR] |
| DeepSeek V4.1 Flash (off-peak) | $0.15 | $0.60 | $0.003 | 2026-09-10 | API + MIT weights [VENDOR] |
| DeepSeek V4.1 Flash (peak) | $0.30 | $1.20 | $0.006 | 2026-09-10 | API + MIT weights [VENDOR] |
| Poolside Laguna S 2.1 | $0.09 | $0.18 | — | 2026-07-21 | API + OpenMDW-1.1 weights [VENDOR] |
| GPT-4o | $2.50 | $10.00 | $1.25 | Sept 2026 | Closed API (legacy) [VENDOR] |
| Cohere Command R+ | $3.00 | $15.00 | — | Sept 2026 | Closed API [SECONDARY] |
| MiniMax M3 | $0.60/$2.40 (std; $0.30/$1.20 launch promo) | — | — | 2026 (exact date not pinned) | API [SECONDARY] |
| ERNIE 5.0 | $0.85 | $3.40 | — | 2026 | API [SECONDARY] |
| ERNIE 5.1 | $0.59 | $2.65 | — | 2026 | API [SECONDARY] |
| ByteDance Seed 2.1 Turbo | $0.50 | $2.50 | — | 2026-06-24 | API, proprietary [SECONDARY] |

- The "OpenAI shelved flagship Astra Aug 7" Medium assertion is **[UNVERIFIED]** and conflicts with the documented GPT-6 Astra launch of 2026-09-03 — not used as fact.

### Cheapest frontier-class and price spreads

- DeepSeek V4.1 Flash off-peak ($0.15/$0.60) is the cheapest frontier-class API on the board [VENDOR].
- Artificial Analysis's cost-per-Index-task metric (captures verbosity, not just rate cards) puts V4.1 Flash at **$0.27/task** — the cheapest measured — vs Step 5 Preview $0.71, DeepSeek V4 Pro $0.67, Grok 4.6 $1.86, Kimi K3 $2.00, GPT-6 Astra $3.26, Fable 5.1 $7.63 [SECONDARY, intelligentliving citing AA, Sept 2026]. Open-weight models occupy 5 of the 7 cheapest per-task slots.
- **Input spread: 111×** ($0.09 → $10.00 per 1M) [VENDOR].
- **20× spread inside one generation**: GPT-5.6 Sol $4/$20 vs Luna $0.20/$1.20 [VENDOR].
- Fable 5.1 costs 66× V4.1 Flash off-peak on input, 83× on output [VENDOR].

### Context-window pricing tiers (standard in 2026)

- Grok 4.7: $2/$6 **≤200K tokens**, $4/$12 above (2026-09-21) [VENDOR]; Grok 4.6: long-context pricing starts at 200K [VENDOR].
- Gemini 3.1 Pro: $2/$12 ≤200K, **$4/$18 the moment a single prompt exceeds 200K** (standing surcharge, not a repricing) [VENDOR].
- GPT-6 Astra: $10/$1 cached/$50 below 272K; **2× input/cache, 1.5× output above 272K** [VENDOR].
- Batch/Flex runs at roughly **half** standard rate; Fast/Priority mode at **double** — one model now carries six or more distinct prices [SECONDARY, Medium Sept 2026].
- Anthropic cache reads at $0.25/M (Fable/Mythos 5.1) = 97.5% off input — the industry-wide push to make long-context reasoning affordable [VENDOR].

### Free tiers and reseller arbitrage (gateway-shaped, not unlimited)

- Kimi: free K3 quota with a Kimi account [VENDOR].
- OpenRouter **:free routes** — InclusionAI Ling 3.0 Flash (rate-limited), the full NVIDIA Nemotron 3 family (Ultra/Super/Nano/3.5) [COMMUNITY].
- MiMo-V2.5-TTS-Series launched **limited-time free** (2026-04-23) [SECONDARY].
- Reseller arbitrage: Qubax lists GPT-5.6 Luna at $0.015/$0.09 vs retail $0.20/$1.20 (~91% off) and Gemini 3.7 Flash at ~91% off retail [SECONDARY] — the effective floor for API buyers is often the reseller, not the vendor.

