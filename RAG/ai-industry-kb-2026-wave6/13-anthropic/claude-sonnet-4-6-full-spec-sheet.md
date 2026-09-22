---
id: ai-industry-kb-2026-wave6/13-anthropic/claude-sonnet-4-6-full-spec-sheet
title: "Claude Sonnet 4.6 — full spec sheet"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Glasswing", "Microsoft", "Stripe", "United States"]
dates: ["2025-08", "2026-02", "2026-05-28", "2026-06-09", "2026-06-10", "2026-06-12", "2026-06-14", "2026-06-20", "2026-06-22", "2026-06-30"]
keywords: ["claude", "aws", "bedrock", "benchmark", "benchmarks", "context window", "cost", "cyber", "disclosure", "distillation", "fable 5", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6349, 6425]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: 28ae9b31965974cf6d3e9ba77aa6516d0127e52f2ef3f9ecbdc47b7b697cf2a6
---

# Claude Sonnet 4.6 — full spec sheet

### Claude Sonnet 4.6 — full spec sheet
- Released February 2026 as part of the Claude 4.6 lineup [SECONDARY](https://www.sostav.ru/blogs/278670/84700)
- Model ID `claude-sonnet-4-6` [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Context 1M tokens [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Max output 16K per one spec table [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- MAX-OUTPUT DISCREPANCY: another source lists 64K max output — flag the discrepancy, do not cite one figure without caveat [SECONDARY](https://www.sostav.ru/blogs/278670/84700)
- Knowledge cutoff August 2025 [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Input pricing: $3/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md) [SECONDARY](https://github.com/obrelix/what-lurks-within/blob/HEAD/claude-opus-4.6-prompt-engineering-guide.md)
- Output pricing: $15/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 5-min cache write: $3.75/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 1-hr cache write: $6/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Cache read: $0.30/M (−90%) [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Reported speed ~40–60 t/s [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Vision supported [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Tool use supported [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Extended thinking supported [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Default model in Claude.ai at launch [SECONDARY](https://www.sostav.ru/blogs/278670/84700)
- Claude Code available on Pro plans and above at Sonnet 4.6 launch [SECONDARY](https://www.sostav.ru/blogs/278670/84700)

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

### June 12, 2026 export-control suspension (full episode)
- On 2026-06-12 at 5:21pm ET the US Commerce Department issued an export-control directive ordering Anthropic to suspend Fable 5 and Mythos 5 for all foreign nationals [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md) [SECONDARY](https://9to5mac.com/2026/07/01/claude-fable-5-cleared-to-return-as-us-lifts-anthropics-export-control-restriction/)
- The order covered even foreign-national Anthropic employees [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Cited national security; the government letter gave no stated specifics [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Anthropic had ~90 minutes to comply [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- With no API-level way to verify citizenship, Anthropic disabled both models for every user worldwide [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Fable 5 was offline by ~5:45pm ET [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Trigger per reporting: Amazon researchers found a jailbreak technique (asking the model to "read a codebase and fix software flaws" could identify exploitable vulnerabilities) [SECONDARY](https://github.com/na-e/claude_code_daily_learning/blob/HEAD/entries/2026-06-20.md)
- Amazon CEO Andy Jassy reportedly flagged it to the administration; Amazon is Anthropic's largest investor (~$13B stake) — this chain is [SECONDARY] reporting, not Anthropic-confirmed [SECONDARY](https://github.com/na-e/claude_code_daily_learning/blob/HEAD/entries/2026-06-20.md)
- Anthropic's objection: Fable 5 had strong safeguards; pulling an entire model over a narrow vulnerability sets a precedent that would halt frontier deployments industry-wide [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Anthropic said it had "not even received a disclosure of a concerning non-universal potential jailbreak that led to a harmful result" [SECONDARY](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Controls lifted 2026-06-30: Anthropic cleared to release Mythos 5 to 100+ US institutions [SECONDARY](https://9to5mac.com/2026/07/01/claude-fable-5-cleared-to-return-as-us-lifts-anthropics-export-control-restriction/)
- Fable 5 restoration for general availability followed — ~19 days offline [SECONDARY](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- This was the first time the US government used export-control authority to forcibly suspend a commercial AI model [SECONDARY](https://github.com/na-e/claude_code_daily_learning/blob/HEAD/entries/2026-06-20.md)
- NOTE: per requester instruction, all facts in this episode carry [SECONDARY] labels.

