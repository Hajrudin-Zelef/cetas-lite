---
id: ai-industry-kb-2026-wave6/13-anthropic/claude-opus-4-7-4-8
title: "Claude Opus 4.7 / 4.8"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Glasswing", "Microsoft", "Stripe"]
dates: ["2026-05-28", "2026-06-09", "2026-06-10", "2026-06-22"]
keywords: ["claude", "opus 4", "aws", "bedrock", "benchmark", "benchmarks", "context window", "cost", "cyber", "distillation", "fable 5", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6368, 6409]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: ca9aaa11f7fe4324e9e6f16cd0114d2282c929cc4fd05c92fed6b69299a78774
---

# Claude Opus 4.7 / 4.8

### Claude Opus 4.7 / 4.8
- Opus 4.7 introduced a tokenizer yielding up to 35% more tokens for the same text than 4.6 — same nominal context window carries materially more effective text [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8 launched 2026-05-28 [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8: 1M context [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8 pricing: flat $5/M input, $25/M output [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Opus 4.8 Fast mode: $10/$50 at up to 2.5× speed [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- Fast-mode price cut: Opus 4.7's fast mode was reportedly $30/$150, so 4.8's $10/$50 is a 3× cut [SECONDARY](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- New in 4.8: Dynamic Workflows in Claude Code [SECONDARY](https://www.ghacks.net/2026/05/30/anthropic-releases-claude-opus-4-8-with-effort-controls-and-dynamic-workflows-for-claude-code/)
- New in 4.8: effort controls in Claude.ai/Cowork [SECONDARY](https://www.ghacks.net/2026/05/30/anthropic-releases-claude-opus-4-8-with-effort-controls-and-dynamic-workflows-for-claude-code/)
- Vendor benchmarks for 4.8 (vendor claims): SWE-bench Pro 69.2% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Vendor: SWE-bench Verified 88.6% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Vendor: Online-Mind2Web 84% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Vendor: USAMO 2026 96.7% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Vendor claim: 4.8 is four-times less likely than 4.7 to overlook flaws in its own generated code [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- One source reports an Artificial Analysis figure of 61.4 for 4.8 without a pinned index version — do NOT merge with versioned index scores [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-blog-anthropic-api-pricing.md)

### Claude Fable 5 / Mythos 5 — full spec sheet
- Same underlying model/architecture, different safeguards [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- 1M context [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- 128K max output [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- Always-on adaptive reasoning [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- Input pricing: $10/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Output pricing: $50/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 5-min cache write: $12.50/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 1-hr cache write: $20/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Cache read: $1/M (reduced to $0.25 only with 5.1, see below) [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- CONTEXT CONFUSION: one weak source says 128K context; multiple stronger sources say 1M context with 128K max output — use 1M/128K [UNVERIFIED](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)
- Fable routes offensive cyber, bio/chem, and distillation-sensitive requests to Opus 4.8 [SECONDARY](https://pondero.ai/news/2026-06-10-claude-fable-5-mythos-5-june-2026/)
- Vendor claims under 5% of sessions trigger routing [SECONDARY](https://pondero.ai/news/2026-06-10-claude-fable-5-mythos-5-june-2026/)
- Mandatory 30-day retention for Mythos-class traffic [SECONDARY](https://ai-checker.webcoda.com.au/articles/claude-mythos-fable-general-release-2026)
- No zero-data-retention at launch for Mythos-class [SECONDARY](https://ai-checker.webcoda.com.au/articles/claude-mythos-fable-general-release-2026)
- Mythos restricted through Project Glasswing (vetted customers) [SECONDARY](https://dev.to/itsemcy/claude-fable-5-explained-what-anthropics-new-mythos-class-model-means-186e)
- Public cloud availability: Claude API [SECONDARY](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)
- AWS Bedrock [SECONDARY](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)
- Vertex AI [SECONDARY](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)
- Microsoft Foundry [SECONDARY](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)
- Fable free on paid subscription products 2026-06-09 through 2026-06-22 [SECONDARY](https://github.com/beannation/wtclaude/blob/HEAD/site/src/content/blog/claude-fable-5-pricing-explained.md)
- API billed from day one (no free API window) [SECONDARY](https://github.com/beannation/wtclaude/blob/HEAD/site/src/content/blog/claude-fable-5-pricing-explained.md)
- The free window never completed — the June 12 suspension cut it short [SECONDARY](https://github.com/beannation/wtclaude/blob/HEAD/site/src/content/blog/claude-fable-5-pricing-explained.md)
- Vendor SWE-bench Pro figure 80.3% — vendor figure only [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- Partner/vendor testimony: Stripe reported a 50M-line Ruby migration compressed from two months to one day — testimony, not an independent benchmark [SECONDARY](https://pulse2.com/anthropic-launches-claude-fable-5-and-mythos-5-ai-models/)

