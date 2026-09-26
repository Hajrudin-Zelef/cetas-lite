---
id: ai-industry-kb-2026-wave6/13-anthropic/claude-fable-5-1-mythos-5-1-full-spec-sheet
title: "Claude Fable 5.1 / Mythos 5.1 — full spec sheet"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google"]
dates: ["2026-01-12", "2026-01-16", "2026-01-23", "2026-02-10", "2026-05", "2026-06", "2026-07", "2026-07-07", "2026-09-01"]
keywords: ["claude", "fable 5", "mythos 5", "agent", "agentic", "agents", "attribution", "benchmarks", "cyber", "opus 4", "pricing", "research"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6467, 6509]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: dd30a71ec13790c6de57f1a74fd93e3d04798f306ef0b75b36f037c029e321d6
---

# Claude Fable 5.1 / Mythos 5.1 — full spec sheet

### Claude Fable 5.1 / Mythos 5.1 — full spec sheet
- Released 2026-09-01 [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Same weights, different safeguards between Fable and Mythos [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Fable GA; Mythos restricted [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- 1M context [SECONDARY](https://servola.de/journal/anthropic-claude-fable-mythos-5-1-us-only-access-gap/)
- 128K output [SECONDARY](https://servola.de/journal/anthropic-claude-fable-mythos-5-1-us-only-access-gap/)
- June 2026 knowledge cutoff [SECONDARY](https://servola.de/journal/anthropic-claude-fable-mythos-5-1-us-only-access-gap/)
- Input/output pricing unchanged at $10/$50 [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75) [SECONDARY](https://abhs.in/blog/claude-fable-5-1-mythos-5-1-cache-reads-25-percent-cheaper-september-2026)
- Cache read cut 75%: from $1 to $0.25/M on 2026-09-01 [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75) [SECONDARY](https://abhs.in/blog/claude-fable-5-1-mythos-5-1-cache-reads-25-percent-cheaper-september-2026)
- Estimated 25% typical-workload savings from the cache cut [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Estimated up to 45% savings for agentic workloads with heavy cache reuse [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Vendor benchmarks (vendor claims): Terminal-Bench Science 0.1 — Fable 52.6% vs Fable 5 24.7% [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Vendor: Terminal-Bench 4.0 — Fable 55.8%, Mythos 60.9% [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- TB 4.0 ATTRIBUTION CAVEAT: one source says Anthropic did not publish a matching Mythos score, while multiple others report vendor 60.9% for Mythos — attribute carefully [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Vendor: HLE no tools 60.9%, with tools 65.0% [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Vendor: GDPval-AA v2 1853 [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Vendor: CursorBench 73.4% [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Vendor: OSWorld 77.9% (partial protocol) [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Safeguards reportedly cut cyber false positives 60% and benign biology false positives 85% (vendor claim) [SECONDARY](https://kahawatungu.com/anthropic-unveils-claude-fable-5-1-mythos-5-1-lower-costs-stronger-coding-expanded-science/)
- Fable can discover vulnerabilities but not develop exploits [SECONDARY](https://bitnewsbot.com/anthropic-drops-claude-fable-5-1-doubles/)
- Mythos is for vetted cyber/life-science users [SECONDARY](https://bitnewsbot.com/anthropic-drops-claude-fable-5-1-doubles/)

### Claude Cowork
- Launched 2026-01-12 as a research preview for Max subscribers in the Claude macOS desktop app [SECONDARY](https://mlq.ai/news/anthropic-launches-cowork-accessible-version-of-claude-code-for-non-developers/) [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-release)
- Pro rollout 2026-01-16 [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-release)
- Team/Enterprise rollout 2026-01-23 [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-release)
- Cowork grants Claude permission-based access to specific local folders [SECONDARY](https://aragonresearch.com/anthropic-claude-cowork/) [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-launch)
- Reads, edits, creates files [SECONDARY](https://aragonresearch.com/anthropic-claude-cowork/)
- Plans, executes in parallel, seeks clarification [SECONDARY](https://aragonresearch.com/anthropic-claude-cowork/)
- Built on the Claude Agent SDK [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-launch)
- Positioned as "Claude Code without the code" for non-developers [SECONDARY](https://mlq.ai/news/anthropic-launches-cowork-accessible-version-of-claude-code-for-non-developers/)
- Initial integrations: Google Drive [SECONDARY](https://mlq.ai/news/anthropic-launches-cowork-accessible-version-of-claude-code-for-non-developers/)
- Initial integrations: Canva [SECONDARY](https://mlq.ai/news/anthropic-launches-cowork-accessible-version-of-claude-code-for-non-developers/)
- Web tasks via Claude in Chrome [SECONDARY](https://mlq.ai/news/anthropic-launches-cowork-accessible-version-of-claude-code-for-non-developers/)
- Plugin system (skills, connectors, hooks, sub-agents) arrived ~May 2026 [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Web and mobile rollout 2026-07-07 [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Sessions run remotely by default (beta) on Anthropic-managed sandboxes [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Scheduled tasks run with no device online [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- By July 2026 Anthropic reported >90% of Cowork usage was not software development [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Business operations and content creation were the largest categories [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Windows support added 2026-02-10 per one secondary report [SECONDARY](https://adtools.org/buyers-guide/ai-news-anthropic-claude-cowork-release)
- Effort controls surfaced in Claude.ai/Cowork with Opus 4.8 [SECONDARY](https://www.ghacks.net/2026/05/30/anthropic-releases-claude-opus-4-8-with-effort-controls-and-dynamic-workflows-for-claude-code/)

