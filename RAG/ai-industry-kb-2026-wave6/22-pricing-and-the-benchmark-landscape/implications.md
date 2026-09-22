---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/implications
title: "Implications"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["AWS", "Anthropic", "DeepSeek", "Glasswing", "Google", "Meta", "Mistral", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2025-10-15", "2026-04", "2026-05-23", "2026-05-28", "2026-06-30", "2026-07", "2026-07-24", "2026-08-11", "2026-08-13", "2026-08-16", "2026-08-19", "2026-08-26", "2026-08-29", "2026-08-31", "2026-09", "2026-09-01", "2026-09-10", "2026-12-31"]
keywords: ["agentic", "agents", "astra", "attention", "bedrock", "benchmark", "benchmarks", "claude", "compute", "cost", "deepseek", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10787, 10950]
section: "§22. Pricing and the Benchmark Landscape"
sha256: 7f5c2dd89cc642e0b39778c9197ad47b6b6a55960d7c9dab36e2936f26adc512
---

# Implications

## Implications

1. Price is now a per-model, per-context, per-time-of-day object — a single model carries 6+ prices; rate cards decay in weeks. Any pricing analysis without a price date is stale on arrival.
2. The durable pricing facts: the 20× intra-generation spread (Sol/Luna), the 200K context cliffs, and the ~90% cache-read discounts.
3. Cost-per-intelligence, not tokens-per-dollar, is the 2026 procurement metric — AA's $/Index-task (V4.1 Flash $0.27 vs Fable 5.1 $7.63) captures verbosity and is the number enterprise buyers actually feel; open-weight models dominate the cheap end [SECONDARY].
4. The benchmark regime is unstable by design — three AA revisions in one week; SWE-bench Verified saturated; TB 4.0 a hard reset; Arena factuality-weighting reshuffles the board. Version-pin every number; treat any undated benchmark claim as suspect [DIRECTIONAL].
5. SWE-bench Pro's official Scale board (Muse Spark 1.1 61.5%) is the live coding signal; vendor-aggregate numbers belong to a different class and must never be mixed with it [SECONDARY].
6. Factuality remains the binding constraint — the ~70% FACTS ceiling is unbroken; Multimodal sub-scores (~46–47%) are the caution flag for vision-grounded deployments [SECONDARY].
7. Free tiers are gateway-shaped, and reseller arbitrage (~91% off) sets the effective floor — procurement should benchmark against resellers, not just list prices [SECONDARY].


### New verified implications — expansion

- The 2026 pricing war is fought on three axes at once: list price (Mistral Large 3's $2→$0.50 cut), promotional windows (Gemini Flash through 2026-12-31), and caching architecture (10x cache-read discounts) — quoting a single number without the axis is meaningless [DIRECTIONAL].
- OpenAI's 272K whole-request repricing means context-window headroom is a billing cliff, not a gradient: a 273K-token request costs double a 271K one on input [SECONDARY].
- The aggregator-vs-first-party split (GLM on OpenRouter vs Z.ai direct) means the same model has two prices; benchmarks citing "cost" must name the route [SECONDARY].
- Cache-write surcharges (GPT-5.6 1.25x, Claude explicit 1.25x) make "caching" a two-sided price: workloads that write more than they re-read can pay MORE with caching on [SECONDARY].
- The effort knob is a price knob (1.5x–2x): benchmark leaderboards at max effort quote the most expensive configuration, not the default one [SECONDARY].


- Price-per-benchmark-point is the missing leaderboard column: the Sept 2026 DeepSWE board's top scorers include a $2.36 model and a $13.41 model at the same 74% — ranking by score alone hides a 6x efficiency spread [DIRECTIONAL].
- The Terminal-Bench 114x cost ratio for equal scores (DeepSeek V4.1 Flash vs GPT-5.5+Codex) is the strongest 2026 evidence that harness and model selection dominate benchmark economics more than raw capability [DIRECTIONAL].
- Promotional pricing (Gemini Flash through 2026-12-31) is a demand-shaping tool: the $3.75→$7.50 step is pre-announced, so adoption built on promo pricing faces a known 2x cost increase — budget for the scheduled rate, not the promo [SECONDARY].
- The 272K whole-request repricing makes OpenAI's long context a two-tier product in practice: under 272K it is one price, over it another — with the 400K/1M context tiers sitting above the price cliff [SECONDARY].


- The §22 pricing table should carry a "verified" date per row: a 10-week-old price and a current price look identical in a table but behave differently in a budget [DIRECTIONAL].
- Aggregator pricing (OpenRouter) and first-party pricing (Z.ai direct) diverge enough to change model selection — the cheaper route for the same weights is a procurement decision, not a technical one [SECONDARY].
- The budget-tier cluster ($0.12–$0.30 input across four labs) is where 2026's volume now lives: benchmark leadership at $10–$15 input no longer sets the default choice [DIRECTIONAL].


- Quote effective price, not list price: at 90% cache-hit rates the ranking changes — procurement that compares list input prices is comparing the wrong number [DIRECTIONAL].
- The 240x output spread means output-heavy workloads (agents, reasoning) face a different price war than input-heavy ones (RAG, long-context) — optimize the price of the token type you actually burn [DIRECTIONAL].
- Medium 3→3.5's 3.75x price increase alongside Large 3's 4x cut shows Mistral repricing the line, not just discounting it — model numbers and prices move independently [SECONDARY].


- The 156x monthly-budget spread ($2.50 → $390 for identical token volume) means model selection is now primarily a financial decision with technical constraints, not the reverse [DIRECTIONAL].
- A pricing table without verification dates is a liability: the 10-week age spread across §22's rows is normal for 2026, and the table should show it [DIRECTIONAL].
- The five-layer price stack means "the price" of a model is a function of workload shape (cache-hit rate, urgency, effort, context length) — there is no single number [DIRECTIONAL].

## Sources and URLs

- https://www.cloudzero.com/blog/token-based-pricing/
- https://binaryverseai.com/llm-pricing-comparison/
- https://intuitionlabs.ai/articles/ai-api-pricing-comparison-grok-gemini-openai-claude
- https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- https://medium.com/@blueblud/the-truth-about-ai-api-token-pricing-from-gpt-4s-30-to-gpt-5-6-s-0-20-and-why-it-went-back-up-5645994ccbca
- https://www.intelligentliving.co/step-5-preview-cheapest-frontier-ai/
- https://qubax.ai/blog/2026-08-29-gemini-37-flash-vs-gpt-56-luna-cost-efficiency-comparison
- https://aiweekly.co/alerts/artificial-analysis-ships-index-v43-adds-automationbench-aa
- https://www.techtimes.com/articles/327053/20260909/ai-leaderboard-rewrote-itself-three-times-last-week-same-score-half-cost.htm
- https://officechai.com/ai/artificial-analysis-updates-intelligence-index-twice-in-2-days-fable-5-1-gpt-6-astra-now-tied-for-first-place/
- https://benchlm.ai/benchmarks/artificialanalysis
- https://benchlm.ai/benchmarks/terminal-bench-4
- https://www.alextech.ai/en/news/qwen38-max-0902-reclaims-top-spot-in-chinas-ai-leaderboard/
- https://www.alextech.ai/en/news/grok-47-undercuts-rivals-but-trails-gpt-6-on-agentic-coding/
- https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- https://www.swfte.com/ai/lmarena-ai
- https://www.swfte.com/ai/leaderboard
- https://ofox.ai/blog/llm-leaderboard-best-ai-models-ranked-2026/
- https://news.cleartechai.com/the-leaderboard-you-cant-game-funded-by-the-companies-it-ranks/
- https://cryptobriefing.com/arena-factuality-rankings-language-models/
- https://sherwood.news/tech/ai-leaderboard-maker-lmarena-hits-usd1-7-billion-valuation/
- https://www.predictionmarketnetwork.com/article/ai-models-already-surpass-1520-arena-score-ahead-of-deadline-183135-20260916
- https://reportwire.org/the-70-factuality-ceiling-why-googles-new-facts-benchmark-is-a-wake-up-call/
- https://venturebeat.com/ai/the-70-factuality-ceiling-why-googles-new-facts-benchmark-is-a-wake-up-call
- https://awesomeagents.ai/leaderboards/hallucination-benchmarks-leaderboard/


### New sources — expansion

- https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- https://github.com/pydantic/genai-prices/pull/625
- https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md
- https://www.morphllm.com/claude-code-pricing
- https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- https://github.com/caudena/beam_weaver/blob/HEAD/docs/partners/google.md
- https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
- https://github.com/aaronmarchant96-max/rei-ai/blob/HEAD/docs/CACHE_PRICING_LANDSCAPE.md
- https://github.com/2389-research/dippin-lang/commit/437a8a2889b92a4842582195f2a525836ad93645
- https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.5/2026-08-19-glm-5.3-ga.md
- https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.8/2026-08-26-glm-5.3-flash-vision.md
- https://openrouter.ai/z-ai/glm-5.3
- https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- https://emergent.sh/learn/what-is-glm-5-3
- https://www.aipricing.guru/mistral-ai-pricing/
- https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- https://curlscape.com/blog/mistral-api-pricing-2026
- https://aiworldtoday.com/guides/mistral-ai-pricing
- https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
- https://www.smashingapps.com/mistral-ai-review/
- https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- https://deepswe.datacurve.ai/


- https://aiworldtoday.com/guides/mistral-ai-pricing
- https://www.smashingapps.com/mistral-ai-review/

**Additional facts, fifth tranche:**

- DeepSeek V4 Pro 0813 GA (August 13, 2026): 1.6T-parameter MoE, 49B active per token, hybrid attention cutting single-token compute to 27% and KV-cache to 10% of the V3.2 generation; 1M context, 384K max output; concurrency capped at 500 requests for the Pro endpoint [SECONDARY]. Source: https://github.com/full-stack-assets/wireandlogic/blob/HEAD/content/posts/deepseek-v4-pro-0813-release.mdx
- DeepSeek price HIKE (effective August 16, 2026): the flat rates in effect since May 23, 2026 (when the promotional discount became permanent) were replaced by peak/off-peak billing — V4 Pro peak $1.32 input / $3.96 output, off-peak $0.66 / $1.98; V4 Flash peak $0.44 / $1.32, off-peak $0.22 / $0.66; peak = 2x off-peak [SECONDARY]. Sources: http://www.techtimes.com/articles/324764/20260817/deepseek-v4-api-prices-quadruple-peak-what-developers-pay-starting-now.htm and https://felloai.com/deepseek-pricing/
- Pre-hike baselines for the ledger: V4 Pro $0.435/$0.87 (cache-hit input $0.003625), V4 Flash $0.14/$0.28 (cache-hit $0.0028) — the July 2026 official pricing docs [SECONDARY]. Source: http://coworker.ai/blog/deepseek-api-pricing
- Peak windows: 01:00–04:00 and 06:00–10:00 UTC; off-peak is everything else; DeepSeek gave ten days' notice (warning August 6, final figures August 13, effective August 16) [SECONDARY]. Source: https://www.morphllm.com/deepseek-api
- Percentage increases at peak: V4 Pro output +355% ($0.87 → $3.96), V4 Pro input +203% ($0.435 → $1.32), V4 Flash output +371% ($0.28 → $1.32), V4 Flash input +214% ($0.14 → $0.44); cache-hit input up to +1,100% on some tiers [SECONDARY]. Source: http://www.techtimes.com/articles/324764/20260817/deepseek-v4-api-prices-quadruple-peak-what-developers-pay-starting-now.htm
- Post-hike economics (September 2026): V4 Pro at $0.435/$0.87 registry-row pricing is STALE — the September 17 pricing page reads `deepseek-v4-pro` at $1.32/$3.96 peak; the September 14 announced reroute of V4 Pro to V4.1 Flash was CANCELLED [SECONDARY]. Source: https://benchlm.ai/deepseek/api-pricing
- Model-ID retirements: `deepseek-chat` and `deepseek-reasoner` aliases reached retirement July 24, 2026 (now point to V4 Flash non-thinking/thinking modes); the `deepseek-v4-flash` ID was retired September 10, 2026 — use `deepseek-flash` [SECONDARY]. Source: https://benchlm.ai/deepseek/api-pricing
- DeepSeek V4 Pro 0813 vendor benchmarks: 80.6% SWE-bench Verified and ahead of Claude Opus 4.8 on Terminal-Bench 2.1 — vendor-reported, not independently verified [VENDOR]. Source: https://github.com/full-stack-assets/wireandlogic/blob/HEAD/content/posts/deepseek-v4-pro-0813-release.mdx
- Free tier: web chat at chat.deepseek.com is free for individual users; every new API account gets a 5M-token free grant; DeepSeek lists no separate batch tier [SECONDARY]. Sources: https://felloai.com/deepseek-pricing/ and https://benchlm.ai/deepseek/api-pricing
- Quality-per-dollar framing (September 2026): V4 Pro costs ~17x less than GPT-5.6 Terra ($2.50/$15) on output and ~28x less than GPT-5.6 Sol or Claude Opus 4.8 ($30/$25 output); V4.1 Flash undercuts GPT-5.6 Luna 5x on output; the caveat is Western flagships still score higher overall [SECONDARY]. Source: https://benchlm.ai/deepseek/api-pricing
- DeepSeek-V3.2-Exp is 685B total MoE under an MIT license; DeepSeek-R1 is 671B total with 37B activated — the open-weight side of the DeepSeek line [SECONDARY]. Source: https://www.morphllm.com/deepseek-api
- Third-party reseller pricing (ModelsLab, September 2026): V4 Flash Latest $0.07, V4 Flash 0731 $0.09, V3.2 $0.33, V3.1 $0.60, R1 Distill Llama 70B $0.80, V4 Pro 0813 $1.97, V4 Pro 0423 $2.40 — reseller, not official [SECONDARY]. Source: https://modelslab.com/deepseek-api.md
- Contradiction ledger: the §22 "DeepSeek V4 Pro $0.435/$0.87" and "V4.1 Flash $0.30/$1.20" rows are PRE-HIKE (July 2026) figures; post-August-16 the official rates are peak/off-peak as above — do not quote the flat rates as current [SECONDARY].

**Additional facts, sixth tranche (Anthropic/OpenAI September 2026 rate cards):**

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

- Fable 5.1's $10/$50 rate card versus Fable 5's identical $10/$50 — the point release carried no price change [SECONDARY]. Sources: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md and https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- Sonnet 5's September 1 price rise ($2/$10 → $3/$15) coincided with the end of its introductory window — introductory pricing as a launch tactic, not a permanent tier [SECONDARY]. Source: https://github.com/chaz-clark/make-ai-agents/blob/HEAD/source-docs/anthropic_prompt_caching.md
- GPT-6 Astra's $10/$50 short-context card matches Fable 5.1's $10/$50 exactly — flagship price convergence between OpenAI and Anthropic in September 2026 [SECONDARY]. Source: https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- The cheapest Pro-board output ($0.18/M, DeepSeek-V4-Flash-Max) versus the cheapest Claude output ($5/M, Haiku 4.5) — a 28x gap at the bottom of the market [SECONDARY — derived]. Source: https://www.morphllm.com/swe-bench-pro

**Additional facts, ninth tranche:**

- Haiku 4.5 at $1/$5 is 4x the legacy Haiku 3's $0.25/$1.25 — the "fastest, most cost-efficient" tier still repriced upward on capability [SECONDARY]. Source: https://pecollective.com/tools/anthropic-api-pricing/

