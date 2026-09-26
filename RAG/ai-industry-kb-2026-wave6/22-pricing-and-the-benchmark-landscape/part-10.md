---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/part-10
title: "§22. Pricing and the Benchmark Landscape (part 10)"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["AWS", "Anthropic", "DeepSeek", "Glasswing", "Google", "OpenAI", "Z.ai"]
dates: ["2025-10-15", "2026-04", "2026-05-28", "2026-06-30", "2026-07-24", "2026-08-31", "2026-09", "2026-09-01"]
keywords: ["pricing", "agents", "astra", "bedrock", "claude", "cost", "deepseek", "fable 5", "glm", "gpt-5.6", "gpt-6", "luna"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10910, 10941]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 15f55a4a17f81bde6e3f030d6d39ff21931ae24766e9ca6768ac70beed4d31de
---

# §22. Pricing and the Benchmark Landscape (part 10)

- Claude September 2026 rate card: Fable 5.1 $10/$50 (cache read $0.25, 5-min write $12.50), Opus 5 $5/$25 (read $0.50, write $6.25), Sonnet 5 $2/$10 (read $0.20, write $2.50), Haiku 4.5 $1/$5 (read $0.10, write $1.25) — per million tokens, standard tier [SECONDARY]. Source: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- Sonnet 5 pricing contradiction RESOLVED: $2/$10 was introductory pricing through August 31, 2026; $3/$15 took effect September 1, 2026 (cache read $0.30, write $3.75/1h $6.00) — both figures were correct, dated [SECONDARY]. Source: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- Claude release dates: Fable 5.1 September 1, 2026; Opus 5 July 24, 2026; Sonnet 5 June 30, 2026; Haiku 4.5 October 15, 2025 [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Context and output limits: Fable 5.1, Opus 5, Sonnet 5 all 1M context / 128K max output; Haiku 4.5 200K context / 64K max output [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Thinking and effort: Fable 5.1 thinking always on (low to max effort); Opus 5 and Sonnet 5 on by default with effort levels low to max; Haiku 4.5 effort not supported, token-budget thinking only [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Opus 5 "Fast mode" runs at Fable's price ($10/$50) — a premium speed tier within the Opus line [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Anthropic effort economics: dropping Opus 5 from default to medium effort costs ~2 points on long-horizon coding for half the price; low costs ~8 points for a quarter of the price; on research/knowledge work the curve is nearly flat [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Anthropic's cost-optimization ordering: caching first, then input/output hygiene, then batch, then effort, and only then a change of model [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Prompt-caching multipliers: 5-minute cache writes 1.25x base input, 1-hour writes 2x, cache reads 0.1x — stacking with batch and residency modifiers [SECONDARY]. Source: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- Worked cache example: parking a 500K-token codebase in Opus 4.6's cache costs $2.50 once (500K × $5/M × 1.25) then $0.25 per re-read (500K × $0.50/M × 0.1) — 1,000 requests/day becomes a reasonable line item [SECONDARY]. Source: https://pecollective.com/tools/anthropic-api-pricing/
- Generational price history: Opus 4.6 at $5/$25 is 3x cheaper than Opus 4.1 and Opus 4 at $15/$75; the Sonnet family held $3/$15 steady across four generations (3.7 through 4.6); Haiku 4.5 at $1/$5 versus legacy Haiku 3 at $0.25/$1.25 [SECONDARY]. Source: https://pecollective.com/tools/anthropic-api-pricing/
- Opus 4.8 (May 28, 2026) $5/$25 with Fast Mode at $10/$50 — down from $30/$150 Fast Mode on Opus 4.7 [SECONDARY]. Source: https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-blog-anthropic-api-pricing.md
- Tokenizer cost trap: Opus 4.7's new tokenizer generates up to 35% more tokens for the same input text versus Opus 4.6 — per-token prices unchanged but effective per-request cost rises up to 35% on migration; 4.7→4.8 is config-only with no additional penalty [SECONDARY]. Source: https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-blog-anthropic-api-pricing.md
- Deprecation ledger: Opus 4.1 deprecated; Opus 4 retired except on Google Cloud; Sonnet 4 retired except Bedrock/GCP; Haiku 3 deprecated April 2026; Haiku 3.5 retired except Bedrock/GCP [SECONDARY]. Sources: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md and https://pecollective.com/tools/anthropic-api-pricing/
- OpenAI September 2026 short-context rates: gpt-6-astra $10/$50, gpt-5.6-sol $4/$20, gpt-5.6-terra $2/$12, gpt-5.6-luna $0.20/$1.20, gpt-5.5 $5/$30, gpt-5.5-pro $30/$180, gpt-5.4 $2.50/$15, gpt-5.4-pro $30/$180 [SECONDARY]. Source: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- OpenAI long-context rates (2x input, higher output): gpt-6-astra $20/$75, gpt-5.6-sol $8/$30, gpt-5.6-terra $4/$18, gpt-5.6-luna $0.40/$1.80, gpt-5.5 $10/$45, gpt-5.5-pro $60/$270, gpt-5.4 $5/$22.50 [SECONDARY]. Source: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- GPT-5.6 Sol contradiction update: short-context is $4/$20 and long-context $8/$30 — the §22 "$5/$30" figure's basis remains unresolved (neither short nor long); do not quote $5/$30 without a dated source [SECONDARY].
- Blended 3:1 cost ladder: Fable 5.1 $20/M, Opus 5 $10/M, Sonnet 5 $4/M, Haiku 4.5 $2/M — output on Fable costs 10x Haiku, each tier roughly doubling [SECONDARY]. Source: https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained

**Additional facts, seventh tranche:**

- Anthropic Batch API: 50% cheaper across all models — anything nobody is waiting for goes through batch at half price [SECONDARY]. Source: https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-blog-anthropic-api-pricing.md
- Opus 4.8 Fast Mode at $10/$50 versus Opus 4.7's $30/$150 — a 3x Fast Mode price cut between generations [SECONDARY]. Source: https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-blog-anthropic-api-pricing.md
- Claude Mythos 5 pricing matches Fable 5 ($10 in, $12.50 5-min write, $20 1-hour write, $1 read, $50 out) — limited availability via Glasswing [SECONDARY]. Source: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- GPT-5.5 Pro and GPT-5.4 Pro: $30/$180 short-context, $60/$270 long-context — the "pro" reasoning tiers at 6x flagship input [SECONDARY]. Source: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- Cache-write pricing ladder (Anthropic): 5-minute writes cost 1.25x base input, 1-hour writes 2x — Fable 5.1's 1-hour write is $20/MTok versus $10 base [SECONDARY]. Source: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- DeepSeek's 5M-token free API grant versus Anthropic's no-free-tier posture — the onboarding cost asymmetry for new developers [SECONDARY]. Sources: https://felloai.com/deepseek-pricing/ and https://pecollective.com/tools/anthropic-api-pricing/
- Price-per-point on SWE-bench Pro (September): DeepSeek-V4-Flash-Max $0.18/M output at 52.6% versus Fable 5 $50/M at 80.0% — ~280x the output price for ~1.5x the score [SECONDARY — derived]. Source: https://www.morphllm.com/swe-bench-pro
- The August→September 2026 price drops (DeepSeek-V4-Pro-Max $3.20→$2.60, GLM-5.2 $3.00→$2.40, GLM-5.1 $4.40→$3.50) all postdate DeepSeek's August 16 hike — third-party cuts against a first-party increase [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro

**Additional facts, eighth tranche:**

