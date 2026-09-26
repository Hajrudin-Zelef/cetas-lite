---
id: ai-industry-kb-2026-wave6/13-anthropic/sources-and-urls
title: "Sources and URLs"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Glasswing", "OpenAI"]
dates: ["2026-06-14", "2026-07"]
keywords: ["agent", "agentic", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "cost", "cyber", "fable 5", "leaderboard", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6710, 6733]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: 9bdfda47c41d982e27bf2c3899166bae1bbf9ce77af1e56cc6f61bc86448ed14
---

# Sources and URLs

- The June 12 suspension shows deployment control can sit with the government, not the company: a first-ever use of export-control authority to suspend a commercial model means frontier deployment schedules now carry sovereign-risk exposure, and fallback routing (`availableModels`) became operationally essential [DIRECTIONAL](https://github.com/hogu59/claude-news/blob/HEAD/src/content/en/briefing-2026-06-14.md)
- Anthropic's Mythos-class retention posture (mandatory 30-day retention, no ZDR at launch) vs Opus 5's reported zero-retention support creates a two-tier privacy architecture: maximum-capability models cost data-retention; standard flagships preserve ZDR — enterprise buyers must pick their tradeoff [DIRECTIONAL](https://ai-checker.webcoda.com.au/articles/claude-mythos-fable-general-release-2026) [DIRECTIONAL](https://coursiv.io/blog/claude-opus-5)
- Sonnet 5's cancelled price rise ($2/$10 made permanent) and Fable 5.1's 75% cache-read cut show Anthropic competing on unit economics, not just capability — agentic workloads with heavy cache reuse get up to ~45% cheaper on 5.1 [DIRECTIONAL](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Cowork's >90% non-development usage by July 2026 validates the "Claude Code without the code" thesis: the agentic market is knowledge work, not just software — direct competitive overlap with OpenAI's ChatGPT Work [DIRECTIONAL](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- The Opus 4.7 tokenizer change (+35% tokens for the same text) quietly moves the effective context/price frontier — token-denominated benchmarks and bills are not comparable across the 4.6→4.7 boundary without adjustment [DIRECTIONAL](https://www.neowin.net/news/anthropic-launches-claude-opus-48-with-better-coding-and-lower-fast-mode-pricing/)
- The $8B gross-vs-net accounting dispute with OpenAI means Anthropic's run-rate series and OpenAI's are not apples-to-apples; enterprise buyers comparing "who leads" should use customer counts (>1,000 at >$1M) and product-level revenue (Claude Code $2.5B) instead [DIRECTIONAL](https://github.com/uroshp/scout-ci/blob/HEAD/v2/archive/anthropic__vs__openai__general/current.md)
- The June-15 Claude Code credit-pool split being announced then paused shows Anthropic is still renegotiating how agentic (non-interactive) usage is metered — the boundary between "chat" and "agent" spend remains unsettled in billing [DIRECTIONAL](https://github.com/beannation/wtclaude/blob/HEAD/site/src/content/blog/claude-billing-changes-june-2026.md)
- Claude Design running on Opus 4.7 (not the latest 4.8) suggests vision-focused research previews pin older checkpoints for stability — the newest model is not always the best substrate for multimodal tooling [DIRECTIONAL](https://github.com/nikbearbrown/brutalist.art/blob/HEAD/examples/ai-explainer/claude-debunked/FACTCHECK-THE-FACTCHECK.md)
- Project Glasswing's Mythos restriction plus mandatory 30-day retention shows the dual-use fence is structural, not promotional — frontier cyber/biology capability now ships with built-in customer vetting and data retention by default [DIRECTIONAL](https://dev.to/itsemcy/claude-fable-5-explained-what-anthropics-new-mythos-class-model-means-186e)

## Sources and URLs
- https://www.vals.ai/models/anthropic_claude-fable-5
- https://therouter.ai/news/anthropic-deprecates-claude-opus-4-1-august-5-migration-guide/
- https://docs.aws.amazon.com/bedrock/latest/userguide/model-card-anthropic-claude-mythos-5.html
- https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/anthropic-claude.md
- https://aiweekly.co/alerts/artificial-analysis-ships-index-v43-adds-automationbench-aa
- https://www.techtimes.com/articles/327053/20260909/ai-leaderboard-rewrote-itself-three-times-last-week-same-score-half-cost.htm
- https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- https://www.swfte.com/ai/lmarena-ai
- https://ofox.ai/blog/llm-leaderboard-best-ai-models-ranked-2026/


### New sources — expansion

