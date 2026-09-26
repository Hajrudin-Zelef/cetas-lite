---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/part-9
title: "§22. Pricing and the Benchmark Landscape (part 9)"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: ["2026-05-23", "2026-07", "2026-07-24", "2026-08-11", "2026-08-13", "2026-08-16", "2026-08-19", "2026-08-26", "2026-09", "2026-09-10"]
keywords: ["benchmark", "pricing", "agents", "attention", "benchmarks", "claude", "compute", "cost", "deepseek", "gemini", "glm", "gpt-5.6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10858, 10909]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: f0d9eea60f6e0b3debb47ceb4ea92957366853c261f4b60eafffeb544f641d03
---

# §22. Pricing and the Benchmark Landscape (part 9)

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

