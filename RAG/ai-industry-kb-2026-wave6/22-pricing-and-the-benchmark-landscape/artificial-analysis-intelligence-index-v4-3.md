---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/artificial-analysis-intelligence-index-v4-3
title: "Artificial Analysis Intelligence Index v4.3"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Meta", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "Z.ai", "xAI"]
dates: ["2026-04", "2026-07", "2026-07-12", "2026-08-10", "2026-08-11", "2026-08-14", "2026-08-19", "2026-08-26", "2026-08-28", "2026-08-31", "2026-09", "2026-09-03", "2026-09-04", "2026-09-07", "2026-09-17", "2026-09-21", "2026-09-22", "2026-12-31", "2027-01-01"]
keywords: ["agent", "agentic", "agents", "astra", "benchmark", "benchmarks", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10559, 10692]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 56eda52efe53c545c1a37964379db620bd15a2987518328694de0c0eb199548f
---

# Artificial Analysis Intelligence Index v4.3

### Artificial Analysis Intelligence Index v4.3

- Published **2026-09-07** — the third revision in under a week (v4.1.1 on 2026-09-03 → v4.2 on 2026-09-04 → v4.3 on 2026-09-07), an accelerated rollout of changes planned for Index v5 [SECONDARY].
- Methodology changes: **Terminal-Bench upgraded v2.1→v4.0**; **τ³-Banking replaced by AutomationBench-AA** (built with Zapier: 657 held-out business-workflow tasks; any guardrail violation zeroes the task); **private-test weighting 40%→45%** [SECONDARY]. Category weights unchanged: Agents 30%, Coding 20%, General 30%, Scientific Reasoning 20%.
- Leaders (v4.3): **Claude Fable 5.1 (max w/ fallback) = GPT-6 Astra (max) = 53 (tie)**; Claude Opus 5 51; Muse Spark 1.3 48 [SECONDARY]. Cost-per-task at the top: Astra $3.26 vs Fable $7.63 (57% cheaper for the same index score) [SECONDARY].
- Open-weight scores (v4.3): **Qwen3.8-Max (0902) 45** [SECONDARY, r/LocalLLaMA via alextech, 2026-09-17 — the 0902 checkpoint reclaimed the China lead], GLM-5.3 44.9, Kimi K3 43.8, GLM-5.3-Flash 42, DeepSeek V4.1-Flash 40, DeepSeek V4-Pro-0813 36 [SECONDARY].
- Later report: MiMo-V2.6-Pro scored **46.32** on the AA Index v4.3 (announced 2026-09-21/22, after the Sept 7 snapshot) — highest open-weight figure reported, ahead of Kimi K3 and Qwen3.8-Max [SECONDARY, unite.ai/VentureBeat]; treat as a separate dated claim, not a merged re-ranking.
- **Scale caution**: one tracker reports Qwen3.8-Max at **58** and Kimi K3 (max) at **60** "on the Intelligence Index," and benchlm's Sept-21 AA mirror shows GPT-5.6 Sol at **58.9%** — a different AA presentation scale (percentage view) than the 53-point Index view. **Never mix the 53-point and 58.9%-style figures.**
- **Standing rule**: scores from v4.1/v4.2 are not comparable to v4.3 — GPQA-Diamond was dropped as solved in the rebase; absolute scores fell. "Any production decision made using last week's leaderboard was made on a methodology that no longer applies" [SECONDARY, techtimes].

### LMArena, September 2026

- Text board (Sept 2026 snapshot): **Claude Fable 5 ~1525 (#1)**, Opus 5 1522, GPT-5.6 Sol 1514, Opus 4.8 1512, Grok 4.5 1499, Gemini 3.1 Pro Preview 1500, **Kimi K3 ~1500**, GLM-5.2 1483, Sonnet 5 1479, DeepSeek V4 Pro 1462, Qwen 3.7 Max 1455, V4 Flash 1441, Llama 4 Maverick 1352 [SECONDARY]. Volume: 10M+ evaluations/month.
- **Coding (Frontend Code) arena: Kimi K3 #1 at 1,679 Elo** — the first open model to top a board outright [SECONDARY].
- Gaming/neutrality concerns are now mainstream: (1) models optimize specifically for Arena — human vote biases reward polish over task success ("if you pick a model off LMArena for a coding agent, you are optimising for the wrong thing") [SECONDARY]; (2) conflict of interest — Arena is a **$1.7B-valuation** startup ($150M Series A; $250M raised in 7 months) backed by **OpenAI, Google, and Anthropic**, and a **~$100M annualized-revenue evaluation business** [SECONDARY]; (3) Arena's factuality-weighted ranking (25% factual-accuracy weight over 2M+ labeled claims) moved GPT-5.5 **up 13 places to #7** and cratered Muse Spark **13 places to #20**, with Fable 5 slipping to #2 [SECONDARY, cryptobriefing ~July 2026] — the default Elo ordering is preference, not truth.

### SWE-bench

- SWE-bench **Verified** (Vals leaderboard, 2026-08-19): **Claude Opus 5 97.0%**, GPT-5.6 Sol Max 96.2%, Claude Fable 5 Max 95.0% — **effectively saturated** as a frontier differentiator [SECONDARY].
- SWE-bench **Pro** (Scale's official standardized public board): **Muse Spark 1.1 + mini-SWE-agent 61.5%** leads the public board [SECONDARY].
- Vendor-reported, non-comparable harness numbers: Claude Fable 5 80.0%, GPT-5.6 Sol 64.6% (OpenAI model card) [SECONDARY] — **never mix vendor-aggregate with the Scale public board**.
- The three number classes must never be mixed: (1) Scale public standardized, 731 tasks; (2) Scale commercial/private standardized, 276 tasks; (3) vendor-aggregate/custom harness.

### Terminal-Bench 4.0

- Released **Sept 1–2, 2026** (announced by tbench.ai **2026-08-28**): 66 tasks (19 fixed, 8 removed as saturated/refusal-prone/publicly solved), mini-SWE-agent harness, 3 runs/task, 5 trials/task, 8-hour agent timeout [SECONDARY].
- Official tbench.ai runs: **Claude Fable 5.1 57.9%±3.8 (#1)**, Opus 5 51.8%, Fable 5 44.5%, **GLM-5.3 41.8%**, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1%±3.4 [SECONDARY].
- Anthropic's own TB-4.0 numbers: **Mythos 5.1 60.9%**, Fable 5.1 55.8%, Opus 5 52.3% [VENDOR].
- BenchLM's Sept-21 mirror: **GPT-6 Astra 58.18%** leads its snapshot; alextech (Sept 22): GPT-6 Astra 60%, Fable 5.1 55%, **DeepSeek V4.1 Flash 27%**, **Grok 4.7 26%** [SECONDARY]. Different snapshots/dates; confidence intervals overlap — attribute per source.
- **TB 4.0 scores are NOT comparable with TB 2.1** (where V4.1 Flash took #1 at 90.6% on the Sept-11 update) — different task set, harness, and difficulty.

### Vals AI and FACTS

- Vals (vals.ai) is the 2026-era independent benchmark operator running verified leaderboards; its SWE-bench Verified board (2026-08-19) is the source for the saturation figures above [SECONDARY].
- Google DeepMind's **FACTS Benchmark Suite** (announced early 2026; initial leaderboard run Dec 2025): four dimensions — Parametric, Search, Multimodal, Grounding v2; 3,513 public examples plus a private Kaggle holdout; LLM-judge ensemble (models rate their own outputs 3.23pp higher on average) [VENDOR/SECONDARY].
- Initial run: **Gemini 3 Pro 68.8** (Search 83.8, Parametric 76.4, Multimodal 46.1), Gemini 2.5 Pro 62.1, GPT-5 61.8, Grok 4 53.6, Claude 4.5 Opus 51.3 [VENDOR/SECONDARY].
- **The ~70% factuality ceiling: no model has broken 70%** — the headline finding [SECONDARY, VentureBeat et al.]. For RAG builders the Search sub-score (Gemini 3 Pro 83.8) is the procurement-relevant number; Multimodal sub-scores (~46–47% even at the top) are the caution flag.

### Who leads what, as of September 2026 (version-pinned)

| Category (version pinned) | Leader |
|---|---|
| Composite intelligence (AA Index **v4.3**, 2026-09-07) | **Fable 5.1 = GPT-6 Astra (53)** — tie; Astra 57% cheaper per task [SECONDARY] |
| Open-weight composite (AA **v4.3**) | Qwen3.8-Max (0902) 45 [SECONDARY] > GLM-5.3 44.9 > Kimi K3 43.8 > GLM-5.3-Flash 42 > V4.1-Flash 40 > V4-Pro-0813 36 |
| Human preference text (LMArena, Sept 2026) | Claude Fable 5 (~1525) [SECONDARY] |
| Human preference coding (LMArena Frontend Code) | Kimi K3 (1,679) — first open #1 [SECONDARY] |
| Agentic terminal (Terminal-Bench **4.0**, official) | Claude Fable 5.1 (57.9%); vendor-self-reported: Mythos 5.1 60.9% [SECONDARY/VENDOR] |
| Issue-resolution coding (SWE-bench Verified) | Saturated — Opus 5 97.0%; live board is SWE-bench Pro (Muse Spark 1.1 61.5% official) [SECONDARY] |
| Factuality/grounding (FACTS Suite, Dec 2025 run) | Gemini 3 Pro (68.8); ceiling ~70% unbroken [VENDOR/SECONDARY] |
| Cost per unit of intelligence (AA $/Index-task, Sept 2026) | DeepSeek V4.1 Flash ($0.27) — cheapest measured [SECONDARY] |


### New verified facts — expansion

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


- Mistral Medium 3 → 3.5 was a price INCREASE: legacy Medium 3 $0.40/$2.00 versus Medium 3.5 $1.50/$7.50 — a 3.75x input jump alongside the capability upgrade; "3.5" is not a discount tier [SECONDARY]. Sources: https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/ versus https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- Three 2026 pricing strategies are visible: penetration (GLM-5.3-Flash $0.15 input, Qwen3.7 Turbo $0.12), promo-window (Gemini Flash 50% off through 2026-12-31), tier-ladder (OpenAI Sol/Terra/Luna, Gemini Lite/Standard/Pro/Max, Mistral Small/Medium/Large) — labs pick one as their primary weapon [DIRECTIONAL].
- Mistral's 4x list-price cut on Large 3 ($2→$0.50) is the clearest documented 2026 evidence of margin compression at the frontier — flagship prices fall while capabilities rise [SECONDARY]. Source: https://www.aipricing.guru/mistral-ai-pricing/
- Effective-price arithmetic (derived, 90% cache-hit assumption): Kimi K3 effective input = 0.9×$0.30 + 0.1×$3 = $0.57/M; GLM-5.3 effective input = 0.9×$0.26 + 0.1×$1.40 = $0.37/M — at high cache-hit rates the $3 vs $1.40 list gap compresses to $0.57 vs $0.37 [DIRECTIONAL — derived from sourced rates].
- Cross-model cache comparison: a cache-HIT Kimi K3 call ($0.30) costs 2x an UNCACHED GLM-5.3-Flash call ($0.15) — cache discounts don't erase list-price gaps [DIRECTIONAL — derived].
- Output-price spread, derived: GLM-5.3-Flash $0.50 → Gemini 3.8 Flash Lite $1.20 → Mistral Small 4 $0.60 → Gemini 3.7 Flash $3.75 (promo) → GLM-5.3 $4.40 → Kimi K3 $15 → Qwen3.7 Max Thinking $16 → GPT-5.6 Terra $24 → Fable 5 $50 → GPT-5.5 Pro $120 — a 240x output-price range [DIRECTIONAL — derived from sourced list prices].
- GPT-5.6 cache-write cost, derived: at $4/M input, writing 1M cache tokens costs $5 (1.25x) — the first write of a long context costs more than the tokens themselves at list [DIRECTIONAL — derived].
- The 272K cliff reprices the 1M context tier: Sol's 1M context exists, but tokens 272K–1M bill at 2x/1.5x — headline context and priced context are different products [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- No September 2026 secondary source in this pass documents OpenAI priority-tier pricing — the priority row is a gap, not a zero [UNVERIFIED].


- Pricing-source reliability ranking for 2026 verification: (1) first-party pricing pages (check monthly); (2) partner docs and changelogs (agenthub, beam_weaver); (3) community price trackers with history (aipricing.guru, since 2026-04); (4) community compilations (kzinmr ai-topics, secondtalent); (5) AI-generated guides (devtk, curlscape — useful but verify) [DIRECTIONAL].
- aipricing.guru hides flat price charts until a change is detected — the Large 3 $2→$0.50 cut is exactly the class of event it surfaces; a flat chart means "no detected change", not "verified current" [SECONDARY]. Source: https://www.aipricing.guru/mistral-ai-pricing/
- The secondtalent Mistral table preserves the pre-cut price structure (Large 3 $2/$6, Medium 3 $0.40/$2.00) — valuable as a historical anchor, dangerous as a current quote [SECONDARY]. Source: https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
- The 2026 price stack has five layers: list price → cache terms (read discount, write surcharge) → batch discount (0.5x) → effort multiplier (1.5x–2x) → context threshold (OpenAI 272K cliff) — a quote that names fewer than five layers is incomplete [DIRECTIONAL].
- Worked monthly budget, derived (10M input + 2M output tokens, no cache): GLM-5.3-Flash $1.50+$1.00=$2.50 | Mistral Small 4 $1.50+$1.20=$2.70 | Gemini 3.7 Flash promo $7.50+$7.50=$15 | Kimi K3 $30+$30=$60 | GPT-5.6 Terra $80+$48=$128 | Fable 5 $100+$100=$200 | GPT-5.5 Pro $150+$240=$390 — a 156x monthly-cost spread for the same token volume [DIRECTIONAL — derived from sourced list prices].
- With 90% cache hit on input: Kimi K3 drops to $5.70+$30=$35.70; GLM-5.3 to $3.74+$8.80=$12.54 — caching moves K3 from 24x to 14x the GLM-5.3-Flash budget [DIRECTIONAL — derived].
- Within-family ratios, derived: Fable 5 costs 2x Opus 5 ($10 vs $5 input); Haiku 4.5 is 5x cheaper than Opus 5; GPT-5.6 Terra is 2x Luna ($8 vs $4); Gemini Flash Max is 30x Flash Lite ($9 vs $0.30); Qwen3.7 Turbo ($0.12) is 125x cheaper than GPT-5.5 Pro ($15) [DIRECTIONAL — derived].
- The next verification pass should resolve: (1) GPT-5.6 Sol $4/$20 vs $5/$30; (2) Sonnet 5 $2/$10 vs $3/$15; (3) Qwen3.7 Plus $0.32 vs $0.42; (4) Mistral Nemo $0.02/$0.10 vs $0.15/$0.15; (5) K3's $20M revenue prong; (6) missing prices for GPT-6 Astra, Grok 4.5/4.6, Mythos 5/5.1, DeepSeek V4 Pro, OpenAI priority tier [SECONDARY].
- Batch is the cheapest tier on both OpenAI and Anthropic (0.5x) — any non-urgent workload not using batch is overpaying by 2x [SECONDARY]. Sources: https://github.com/pydantic/genai-prices/pull/625 and https://github.com/antseed/antseed/blob/HEAD/apps/website/blog/2026-08-11-claude-code-pricing-breakeven.md

