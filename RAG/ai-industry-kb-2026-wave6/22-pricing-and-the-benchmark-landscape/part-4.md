---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/part-4
title: "§22. Pricing and the Benchmark Landscape (part 4)"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "Z.ai", "xAI"]
dates: ["2026-04", "2026-07", "2026-07-12", "2026-08-10", "2026-08-11", "2026-08-14", "2026-08-19", "2026-09", "2026-09-22", "2026-12-31"]
keywords: ["benchmark", "pricing", "astra", "benchmarks", "claude", "cost", "deepseek", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10640, 10671]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: c23773c86c0cf5a8c35d98a5e126aa146be9520e7d236750c1afe583b97718ed
---

# §22. Pricing and the Benchmark Landscape (part 4)

- Qwen3.7 Max Thinking: $4/M input, $16/M output, $1.20/M cached, $0.50/M vector [SECONDARY]. Source: https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- Qwen3.7 Plus: $0.32/M input, $1.28/M output (ayinedjimi PDF lists $0.42) — the two compilations differ; flag as provider/date variance [SECONDARY — minor contradiction]. Sources: https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md versus https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Turbo: $0.12/M input [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Moonshot Kimi context-caching rates live-verified 2026-08-14; DeepSeek/GLM/Grok/Kimi cache-read catalog verified against official pages 2026-08-10 — the caching-price landscape is audited, not assumed [COMMUNITY]. Sources: https://github.com/aaronmarchant96-max/rei-ai/blob/HEAD/docs/CACHE_PRICING_LANDSCAPE.md and https://github.com/2389-research/dippin-lang/commit/437a8a2889b92a4842582195f2a525836ad93645
- Reasoning-effort cost correlation: higher effort tiers cost 1.5x–2x more — the "effort" knob is a price knob [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf


- Mistral OCR version pricing: aiworldtoday lists Mistral OCR at $2/1,000 pages (legacy OCR $8); curlscape lists OCR 4 at $4/1,000 pages (OCR) and $5/1,000 (Document AI) — three generations, three prices; quote the version [SECONDARY]. Sources: https://aiworldtoday.com/guides/mistral-ai-pricing versus https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral Large 3's price cut ($2/$6 → $0.50/$1.50) happened between the April 2026 tracking start and the September 2026 snapshot — a 4x list-price cut in under six months on a flagship [SECONDARY]. Sources: https://www.aipricing.guru/mistral-ai-pricing/ and https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
- pydantic genai-prices PR #625 added GPT-5.6 Standard and Batch rows to the community pricing database — 2026 pricing data is maintained as code, in PRs, not PDFs [SECONDARY]. Source: https://github.com/pydantic/genai-prices/pull/625
- The July 2026 OpenAI pricing research lives in a "cpa-usage-lens" repo — a cost-tracking (FinOps) tool; pricing research is now embedded in spend-management tooling, not analyst reports [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- "Tokenomics" is the 2026 term for pricing-as-architecture: devtk's Gemini 3.8 Flash guide is titled "Tokenomics, Pricing & Practical Guide" — caching, batch, and effort tiers are one system [SECONDARY]. Source: https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- The 272K rule, worked: a 300K-token GPT-5.6 Sol prompt bills all 300K input tokens at 2x ($8/M effective at the $4 list) — the cliff is at the threshold, and the whole request reprices [SECONDARY — arithmetic on sourced rule]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- The alt-f4-llc facts file pairs Claude v5 model costs with benchmarks in one document — cost and capability are now tracked together at the source [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- The beam_weaver Google partner doc carries Gemini pricing for partners — pricing flows through partnership channels, not just public pages [SECONDARY]. Source: https://github.com/caudena/beam_weaver/blob/HEAD/docs/partners/google.md
- Verification-date map (when each price below was confirmed): OpenAI tiers 2026-07-12; Claude Code 2026-08-11; cache landscape 2026-08-10/14; GLM changelogs 2026-08-19/26; Mistral snapshot 2026-09-22; Gemini Flash tiers 2026-09-22 [SECONDARY]. Sources: as listed in NEW SOURCES.
- No verified 2026-09 list prices were found in this pass for: GPT-6 Astra, Grok 4.5/4.6, Claude Mythos 5/5.1, DeepSeek V4 Pro, Gemini 3.5 Flash-Lite — gaps, not zeros [UNVERIFIED].
- Qwen3.6 Max Preview: $7.80/M input — the priciest Qwen tier documented, above Max Thinking's $4 [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf


- Corrections to the existing §22 September 2026 table from this pass: (a) Gemini 3.6/3.7 Flash $3.75 output is PROMOTIONAL through 2026-12-31, scheduled $7.50 from 2027 — the table should show both with dates; (b) Mistral Large 3 is $0.50/$1.50, not $2/$6 — the $2/$6 rows are stale; (c) the $9 output figure belongs to Gemini 3.5 Flash, not 3.6/3.7 Flash; (d) GLM rows must separate Z.ai list prices from OpenRouter aggregator rates [SECONDARY].
- devtk's 2026 OpenAI pricing guide independently documents the GPT-5.6 family tier matrix (Sol/Terra/Luna) and the 272K long-context rule — second confirmation of the contract research [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- The Kimi K3 vs DeepSeek V4 Flash (0731) price comparison is published on kimi-k2.org, a Kimi-aligned domain — vendor-adjacent pricing claims should be read with that affiliation in mind [DIRECTIONAL]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- Z.ai's direct list prices here were captured via OpenRouter's provider page and the agenthub changelogs — first-party numbers reproduced in third-party docs, not scraped from z.ai itself; treat as [VENDOR]-via-secondary [SECONDARY]. Sources: https://openrouter.ai/z-ai/glm-5.3 and https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.5/2026-08-19-glm-5.3-ga.md
- The Claude Code breakeven analysis models per-developer daily spend, not per-token cost — the FinOps framing is seats × days, not tokens [SECONDARY]. Source: https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md
- Pricing verification checklist (2026): (1) date the price — list prices move quarterly; (2) name the provider — first-party vs aggregator; (3) name the tier — standard/batch/priority; (4) check cache terms — read discount vs write surcharge; (5) check context thresholds — OpenAI's 272K cliff; (6) check promo windows — Gemini Flash through 2026-12-31; (7) check model generation — 3.5 vs 3.6/3.7 Flash are different prices [DIRECTIONAL].
- The ayinedjimi Qwen3.7 price table (Max Thinking $4/$16, Plus $0.32–0.42, Turbo $0.12, 3.6 Max Preview $7.80) is dated to the Qwen3.7 generation (2026) — Qwen's internal spread is 65x from Turbo to 3.6 Max Preview [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Within-provider spreads, derived: Gemini Flash Lite→Max 30x input ($0.30→$9); Mistral Small 4→Large 3 3.3x ($0.15→$0.50); GPT-5.6 Sol→GPT-5.5 Pro 3.75x ($4→$15); Qwen Turbo→3.6 Max Preview 65x ($0.12→$7.80) [DIRECTIONAL — derived from sourced list prices].
- Cross-provider budget-tier cluster: GLM-5.3-Flash $0.15, Mistral Small 4 $0.15, Gemini 3.8 Flash Lite $0.30, Qwen3.7 Turbo $0.12 — four labs' cheapest usable tiers sit within a 2.5x band [DIRECTIONAL — derived].
- The September 2026 corpus contains no verified price for GPT-6 Astra anywhere — the newest OpenAI flagship has benchmark scores (DeepSWE 74%) but no public price in the sources pulled [UNVERIFIED].


