---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-3
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 3)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic"]
dates: ["2026-06-30", "2026-07-31", "2026-08-10", "2026-09-10"]
keywords: ["agentic", "benchmarks", "claude", "cost", "opus 4", "parameters", "pricing", "sonnet 5"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [91, 100]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 3b4ceeec3f3054c000142ded09c4ea82911cfbb7f4ca1377f308daff010cb215
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 3)

#### Claude Sonnet 5 — June 30, 2026
- Replaced Sonnet 4.6; became the default model for Free and Pro subscription tiers, Claude Code, and API default selection. [independent] https://www.techtimes.com/articles/319409/20260701/claude-sonnet-5-ships-anthropic-default-agentic-performance-closes-opus-gap.htm
- API ids/context: `claude-sonnet-5`, 1M context, 128K max output. [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Pricing: introductory $2/M input / $10/M output through Aug 31, 2026 (then list price $2 in / $10 out per later price tables — no step-up confirmed). [independent] https://memeburn.com/claude-sonnet-5-launched-with-near-opus-performance-at-half-the-price/ ; [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Anthropic-reported benchmarks (launch): agentic coding 63.2% (vs Opus 4.8's 69.2%), SWE-bench Verified 85.2% (vendor-reported), SWE-bench Pro 63.2%, Terminal-Bench 2.1 80.4%, OSWorld Verified 81.2%, GDPval-AA ~1607–1609 Elo, Artificial Analysis Intelligence Index ~53. [vendor-reported] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/ and [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Vals AI independent: Terminal-Bench 2.1 74.53%. [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- BenchLM snapshots: HLE 57.4% (rank #7 as of 2026-09-10); BenchLM "BenchAlign" composite ~65, rank #29–33 (July 31, 2026); coding #9 / agentic #15; BrowseComp 84.7%. [secondary] https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Safety: Anthropic claims lower rate of undesirable behaviors vs Sonnet 4.6, "safer to use in agentic contexts" (hallucination reduction claimed). [vendor-reported via press] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/
- API behavioral changes (Aug 2026 changelog): adaptive thinking on by default; manual extended-thinking budgets and non-default sampling parameters now return errors; new tokenizer producing ~30% more tokens for the same text. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md

