---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/part-3
title: "§22. Pricing and the Benchmark Landscape (part 3)"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Anthropic", "Google", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-07", "2026-08-11", "2026-08-19", "2026-08-26", "2026-08-31", "2026-09", "2026-09-22", "2026-12-31", "2027-01-01"]
keywords: ["benchmark", "pricing", "agent", "agentic", "agents", "benchmarks", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10613, 10639]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 9d9439ee750033ce244cf36b7a26cda5f4107a6f95538e4af21947731ccc0f25
---

# §22. Pricing and the Benchmark Landscape (part 3)

- OpenAI's 272K long-context rule (2026): above 272K prompt tokens, the ENTIRE request bills at 2x input and 1.5x output for the GPT-5.4/5.5/5.6 families — not just the tokens beyond the threshold [SECONDARY]. Sources: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md and https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- GPT-5.6 explicit cache write costs 1.25x normal input — writing cache is not free, only reading it is discounted [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- GPT-5.6 Sol pricing contradiction ledger: July 2026 contract research lists $4/$20 per 1M; some September 2026 secondary pages list $5/$30 — whether this is a September price increase, a provider-specific markup, or confusion with a "Max" tier is unresolved; check OpenAI's current price page before quoting [SECONDARY — contradiction]. Sources: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md versus September 2026 secondary compilations
- GPT-5.6 Terra: $8/$24 per 1M; GPT-5.6 Luna: $4/$12 per 1M (July 2026 contract research) [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- GPT-5.5 Pro: $15/$120 per 1M; Codex model: $12/$60 per 1M (July 2026) [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- GPT-5.6 Sol context tiers: 400K standard, 1M extended — the 272K price threshold sits below both, so long-context bills at 2x/1.5x well before the context limit [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- Anthropic Claude Code pricing (2026-08-11 analysis): Opus 5 $5/$25; Fable 5 $10/$50 per 1M [SECONDARY]. Sources: https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md and https://www.morphllm.com/claude-code-pricing
- Claude Sonnet 5 pricing contradiction ledger: several sources report the introductory $2/$10 ended 2026-08-31 and the rate became $3/$15; a newer source reports Anthropic made $2/$10 permanent — the current rate must be read from Anthropic's pricing page, not secondary compilations [SECONDARY — contradiction]. Sources: https://www.morphllm.com/claude-code-pricing versus https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Claude Haiku 4.5: $1/$5 per 1M — the budget Claude tier [SECONDARY]. Source: https://www.morphllm.com/claude-code-pricing
- Claude pricing modifiers: explicit cache write 1.25x input; batch 0.5x; code execution $0.05/hour; MCP tool use billed in tool units [SECONDARY]. Source: https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md
- Claude Code breakeven analysis (Aug 2026): at Opus 5 $5/$25, heavy agentic coding sessions run $10–50/day per developer — the unit economics that drove the "Claude Code pricing breakeven" discourse [SECONDARY]. Source: https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md
- Gemini 3.6/3.7 Flash introductory pricing through 2026-12-31: $0.75 input / $0.075 cached input / $3.75 output; from 2027-01-01: $1.50 / $0.15 / $7.50 — the "50% introductory price cut" framing [SECONDARY]. Sources: https://github.com/caudena/beam_weaver/blob/HEAD/docs/partners/google.md and https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
- The existing §22 row showing Gemini Flash at $3.75 output is current PROMOTIONAL pricing; $7.50 is scheduled standard pricing from 2027 — they are not contradictory, they are dated tiers [SECONDARY]. Source: https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- Gemini 3.5 Flash (older generation) output at $9 per 1M — the requested $9 figure belongs to 3.5 Flash, not 3.6/3.7 Flash [SECONDARY]. Source: https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- Gemini 3.8 Flash tiering (Sept 2026): Standard $0.75/$0.075/$3.75 | Pro $4.50/$0.45/$18.00 | Lite $0.30/$0.03/$1.20 | Max $9/$0.90/$36 — a 4-tier Flash lineup spanning 30x input-price range [SECONDARY]. Sources: https://github.com/aaronmarchant96-max/rei-ai/blob/HEAD/docs/CACHE_PRICING_LANDSCAPE.md and https://github.com/2389-research/dippin-lang/commit/437a8a2889b92a4842582195f2a525836ad93645
- Gemini 3.8 Flash Pro context: 2M context standard; the whole Flash family is priced for the "tokenomics" era where caching makes long-context cheap [SECONDARY]. Source: https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- Z.ai first-party list prices (direct): GLM-5.3 $1.40 input / $0.26 cached / $4.40 output; GLM-5.3-Flash $0.15 / $0.03 / $0.50 per 1M [VENDOR]. Sources: https://openrouter.ai/z-ai/glm-5.3 and https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.5/2026-08-19-glm-5.3-ga.md
- OpenRouter aggregator rates for GLM models run LOWER than Z.ai list and change over time — aggregator pricing and first-party list pricing are different numbers; the existing §22 rows mixing them should be separated [SECONDARY]. Source: https://openrouter.ai/z-ai/glm-5.3
- GLM-5.3-Flash on OpenRouter: "the most popular model on OpenRouter by token usage" — price drives adoption, adoption drives the benchmark sample [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.2 on OpenRouter (Aug 26, 2026 changelog): $0.60/$1.70 — the pre-5.3 price anchor [SECONDARY]. Source: https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.8/2026-08-26-glm-5.3-flash-vision.md
- Mistral Large 3: $0.50/$1.50 per 1M, $0.05 cached input — the current official-synced rate (aipricing.guru 2026-09-22 dataset); older $2/$6 rows on stale pages are superseded [SECONDARY]. Sources: https://www.aipricing.guru/mistral-ai-pricing/ and https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- Mistral Medium 3.5: $1.50/$7.50, $0.15 cached; Mistral Small 4: $0.15/$0.60, $0.015 cached; Codestral: $0.30/$0.90; Devstral 2: $0.40/$2.00; Devstral Small 2: $0.10/$0.30; Magistral Medium: $2.00/$5.00; Magistral Small: ~$0.10/$0.30 [SECONDARY]. Sources: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/ and https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral Ministral 3 edge pricing: 3B $0.10/$0.10 (128K), 8B $0.15/$0.15 (256K), 14B $0.20/$0.20 (256K) [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral specialized: OCR 4 $4/1,000 pages (OCR) / $5/1,000 (Document AI); Voxtral Mini Transcribe 2 $0.003/min; Voxtral Small $0.004/min + $0.10/$0.40 text; TTS $16/M chars; Mistral Embed $0.10/M; Codestral Embed $0.15/M; Moderation $0.10/M [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Mistral agent tool charges: web search $30/1,000 calls; code execution $30/1,000 calls; image generation $100/1,000 images; premium news $50/1,000 calls; document libraries $3/1K pages OCR + $1/1M indexing + $0.01/call [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- Le Chat plans: Free $0 (~25 messages/day soft cap, Small model); Pro $14.99/month (~150/day, all models); Team $24/user/month (5-user minimum) [SECONDARY]. Source: https://www.smashingapps.com/mistral-ai-review/
- Kimi K3 API pricing: $3 input / $0.30 cache hit / $15 output per 1M — cache-hit input is 10x cheaper than base input [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
