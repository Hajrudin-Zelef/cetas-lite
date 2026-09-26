---
id: ai-industry-kb-2026-wave6/22-pricing-and-the-benchmark-landscape/part-5
title: "§22. Pricing and the Benchmark Landscape (part 5)"
domain: pricing-and-the-benchmark-landscape
role: deep-dive
task: pricing
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: ["2026-08-11", "2026-09", "2026-12-31"]
keywords: ["pricing", "astra", "claude", "cost", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10672, 10692]
section: "§22. Pricing and the Benchmark Landscape"
delta_of: ai-industry-kb-2026
sha256: 9bd8c46ad3fa1d6eb66eeb7905841869551add790e7dbf71e96f51ed1a0f4af2
---

# §22. Pricing and the Benchmark Landscape (part 5)

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

