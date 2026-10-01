---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-3
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Microsoft", "OpenAI", "Stripe", "United States"]
dates: ["2026-06-09", "2026-06-30", "2026-07-31", "2026-08-10", "2026-09-01", "2026-09-10"]
keywords: ["accelerator", "agentic", "agi", "bedrock", "benchmarks", "claude", "cost", "cyber", "distillation", "fable 5", "fine-tuning", "foundry"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [91, 112]
sha256: 1997e8671f7d19411689f51162b17e40b5a3be8cb4939dab362cd0d29aa95338
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

#### Claude Sonnet 5 — June 30, 2026
- Replaced Sonnet 4.6; became the default model for Free and Pro subscription tiers, Claude Code, and API default selection. [independent] https://www.techtimes.com/articles/319409/20260701/claude-sonnet-5-ships-anthropic-default-agentic-performance-closes-opus-gap.htm
- API ids/context: `claude-sonnet-5`, 1M context, 128K max output. [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Pricing: introductory $2/M input / $10/M output through Aug 31, 2026 (then list price $2 in / $10 out per later price tables — no step-up confirmed). [independent] https://memeburn.com/claude-sonnet-5-launched-with-near-opus-performance-at-half-the-price/ ; [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Anthropic-reported benchmarks (launch): agentic coding 63.2% (vs Opus 4.8's 69.2%), SWE-bench Verified 85.2% (vendor-reported), SWE-bench Pro 63.2%, Terminal-Bench 2.1 80.4%, OSWorld Verified 81.2%, GDPval-AA ~1607–1609 Elo, Artificial Analysis Intelligence Index ~53. [vendor-reported] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/ and [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Vals AI independent: Terminal-Bench 2.1 74.53%. [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- BenchLM snapshots: HLE 57.4% (rank #7 as of 2026-09-10); BenchLM "BenchAlign" composite ~65, rank #29–33 (July 31, 2026); coding #9 / agentic #15; BrowseComp 84.7%. [secondary] https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Safety: Anthropic claims lower rate of undesirable behaviors vs Sonnet 4.6, "safer to use in agentic contexts" (hallucination reduction claimed). [vendor-reported via press] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/
- API behavioral changes (Aug 2026 changelog): adaptive thinking on by default; manual extended-thinking budgets and non-default sampling parameters now return errors; new tokenizer producing ~30% more tokens for the same text. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md

#### Claude Fable 5 + Claude Mythos 5 — June 9, 2026
- Fable 5: first publicly available "Mythos-class" model, tier above Opus; Mythos 5 is the same underlying weights with cyber safeguards lifted, restricted to Project Glasswing partners (cyberdefense, critical-infrastructure orgs; ~150 organizations in 15+ countries, a US-government collaboration) and limited US government channels. [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/ ; [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md
- Availability channels: Claude API, Amazon Bedrock, Google Cloud, Microsoft Foundry. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md
- Pricing: $10/M input / $50/M output for both; 90% prompt-cache input discount; most expensive GA Anthropic model (~2× Opus output). In Claude Code, Fable selectable via `/model fable` or `best` alias, requires v2.1.170+; 1M context on API; not available under zero-data-retention; Anthropic requires 30-day data retention on all Mythos-class traffic. [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md ; [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Subscription availability: free on Pro/Max/Team/seat Enterprise June 9 → June 22 only; from June 23 requires usage credits; Anthropic said it intended to restore as standard subscription model "as quickly as we can," no date given. [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md
- Anthropic-reported benchmarks (launch): SWE-bench Verified 95.0% (Vals AI independent: 95.0%), SWE-bench Pro 80.3% (leading at launch; contested — produced on Anthropic scaffolding), GPQA Diamond 92.6%, Terminal-Bench 2.1 88.0% (Vals AI independent: 80.52%), τ²-Bench 98.5%, FrontierCode 29.3% (vs 13.4% Opus 4.8, 5.7% GPT-5.5), ExploitBench (Mythos 5) 78%. Artificial Analysis Intelligence Index: #1 across public models, quality score 64.9 at launch. [vendor-reported/secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md ; https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Safety architecture: classifiers route cyber/bio/distillation flagged requests to Opus 4.8 fallback; triggers in <5% of sessions (Anthropic claim); 319-page system card. [vendor-reported via secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md
- **Covert-degradation controversy (June 10–11, 2026):** Simon Willison published "If Claude Fable Stops Helping You, You'll Never Know" documenting that for "frontier LLM development" queries (pretraining pipelines, distributed training infra, accelerator design), the system card described INVISIBLE safeguards — prompt modification, steering vectors, parameter-efficient fine-tuning — with no user notification. On June 11 Anthropic reversed the policy: "we made the wrong tradeoff," committing to make all capability restrictions visible and explained. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/68-anthropic-fable-covert-capability-degradation.md
- Adoption datapoint: Stripe used Fable 5 to migrate a 50M-line Ruby codebase in one day (task estimated at two months for a team). [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Suspension (see §1.5) June 12 → lifted ~June 30. [independent] https://www.rappler.com/technology/united-states-lifts-curbs-anthropic-fable-mythos-ai-models/
- **Fable 5.1 released September 1, 2026:** `claude-fable-5-1`, 1M context, $10/$50 per Mtok with $0.25/Mtok cache reads; became default Fable model in Claude Code v2.1.257. BenchLM HLE 65.0% (rank #1 as of 2026-09-10); ARC-AGI-2 90.0%. [secondary] https://github.com/florianbruniaux/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md ; https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md

