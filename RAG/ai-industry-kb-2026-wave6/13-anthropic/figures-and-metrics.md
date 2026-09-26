---
id: ai-industry-kb-2026-wave6/13-anthropic/figures-and-metrics
title: "Figures and metrics"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Glasswing"]
dates: ["2026-02-05", "2026-02-17", "2026-04-07", "2026-04-16", "2026-05-28", "2026-06-09", "2026-06-15", "2026-06-30", "2026-07", "2026-07-24", "2026-08-05", "2026-09-01"]
keywords: ["agentic", "agi", "attribution", "benchmark", "claude", "compute", "cost", "fable 5", "gemini", "mythos 5", "opus 4", "opus 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6594, 6655]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: f4e8cc9c36eb1d7dfcde1e9d377093ebcef36046d6d765900e78329fd244413b
---

# Figures and metrics

## Figures and metrics
| Model | Launch | Input / Output per 1M | Note |
|---|---|---|---|
| Opus 4.6 | 2026-02-05 | — | 4.x anchor [SECONDARY] |
| Sonnet 4.6 | 2026-02-17 | — | 4.x anchor [SECONDARY] |
| Mythos Preview | 2026-04-07 | — | Glasswing-gated [SECONDARY] |
| Opus 4.7 | 2026-04-16 | — | Community-timeline [SECONDARY] |
| Opus 4.8 | 2026-05-28 | — | Community-timeline [SECONDARY] |
| Fable 5 | 2026-06-09 | — | Same weights as Mythos 5 [SECONDARY] |
| Mythos 5 | 2026-06-09 | — | Restricted via Glasswing [SECONDARY] |
| Sonnet 5 | 2026-06-30 | $2.00 / $10.00 (cache $0.20) | Codename "Fennec" [SECONDARY/VENDOR] |
| Haiku 4.5 | Sept 2026 | $1.00 / $5.00 (cache $0.10) | Light tier [SECONDARY] |
| Opus 5 | 2026-07-24 | $5.00 / $25.00 (cache $0.50) | SWE-bench Verified 97.0% [SECONDARY/VENDOR] |
| Fable 5.1 | 2026-09-01 | $10.00 / $50.00 (cache $0.25) | AA v4.3: 53 (tie) [SECONDARY/VENDOR] |
| Mythos 5.1 | 2026-09-01 | $10.00 / $50.00 (cache $0.25) | TB 4.0 vendor 60.9% [VENDOR] |

- Retirements: Opus 4 — 2026-06-15; Opus 4.1 — 2026-08-05; cadence ~60–90 days [SECONDARY].
- Benchmark snapshot (all [SECONDARY] unless noted, version-pinned): TB 4.0 official — Fable 5.1 57.9% #1, Opus 5 51.8%, Fable 5 44.5%; AA v4.3 — Fable 5.1 53 (tie), Opus 5 51; LMArena text — Fable 5 ~1525 #1, Opus 5 1522, Opus 4.8 ~1512, Sonnet 5 1479.


### New verified metrics — expansion

- Opus 4.6 vendor: Terminal-Bench 2.0 65.4% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Opus 4.6 vendor: ARC-AGI-2 68.8% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Opus 4.6 vendor: OSWorld 72.7% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Opus 4.6 vendor: MRCR v2 76% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Opus 4.6 vendor: GDPval-AA 1606 Elo [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Opus 4.8 vendor: SWE-bench Pro 69.2% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Opus 4.8 vendor: SWE-bench Verified 88.6% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Opus 4.8 vendor: Online-Mind2Web 84% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Opus 4.8 vendor: USAMO 2026 96.7% [SECONDARY](https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/55-anthropic-claude-opus-4-8-release.md)
- Fable 5 vendor: SWE-bench Pro 80.3%; routing to Opus 4.8 claimed <5% of sessions [SECONDARY](https://github.com/ombharatiya/ai-system-design-guide/pull/14)
- Sonnet 5 vendor: SWE-bench Pro 63.2% [SECONDARY](https://medium.com/@abhishek.ssntpl/claude-sonnet-5-fennec-what-anthropic-actually-shipped-c6a50fab1bbe)
- Sonnet 5 vendor: OSWorld-Verified 81.2% [SECONDARY](https://medium.com/@abhishek.ssntpl/claude-sonnet-5-fennec-what-anthropic-actually-shipped-c6a50fab1bbe)
- Sonnet 5 vendor: Terminal-Bench 2.1 80.4% [SECONDARY](https://medium.com/@abhishek.ssntpl/claude-sonnet-5-fennec-what-anthropic-actually-shipped-c6a50fab1bbe)
- Sonnet 5 vendor: BrowseComp 25 84.7% [SECONDARY](https://medium.com/@abhishek.ssntpl/claude-sonnet-5-fennec-what-anthropic-actually-shipped-c6a50fab1bbe)
- Sonnet 5 vendor: GDPval-AA v2 1618 [SECONDARY](https://medium.com/@abhishek.ssntpl/claude-sonnet-5-fennec-what-anthropic-actually-shipped-c6a50fab1bbe)
- Opus 5 vendor: Frontier-Bench v0.1 43.3 (vs 18.9 for Opus 4.8) [SECONDARY](https://github.com/premdesai/pilot-shell/blob/HEAD/docs/docusaurus/blog/2026-07-24-claude-opus-5.md)
- Opus 5 vendor: GDPval-AA v2 1861 [SECONDARY](https://github.com/premdesai/pilot-shell/blob/HEAD/docs/docusaurus/blog/2026-07-24-claude-opus-5.md)
- Opus 5 vendor: ARC-AGI-3 30.2 [SECONDARY](https://github.com/premdesai/pilot-shell/blob/HEAD/docs/docusaurus/blog/2026-07-24-claude-opus-5.md)
- Opus 5 Artificial Analysis: v4.1.1 = 63; v4.3 ≈ 50.7/51 — never merge versions [SECONDARY](https://opentools.ai/llms/claude-opus-5)
- Fable 5.1 vendor: TB Science 0.1 Fable 52.6% (vs Fable 5 24.7%) [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Fable 5.1 vendor: TB 4.0 Fable 55.8% / Mythos 60.9% (attribution caveat applies) [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Fable 5.1 vendor: HLE 60.9% no tools / 65.0% with tools [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Fable 5.1 vendor: GDPval-AA v2 1853 [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Fable 5.1 vendor: CursorBench 73.4% [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Fable 5.1 vendor: OSWorld 77.9% (partial protocol) [SECONDARY](https://www.tradingkey.com/analysis/stocks/us-stocks/262145687-anthropic-claude-fable-mythos-performance-cache-reduction-tradingkey)
- Full cache-pricing ladder (documented): Opus 4.6/4.7/4.8 $5 input | $6.25 (5m) | $10 (1h) | $0.50 read [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Sonnet 4.6 cache ladder: $3 input | $3.75 (5m) | $6 (1h) | $0.30 read [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Fable/Mythos 5 cache ladder: $10 input | $12.50 (5m) | $20 (1h) | $1.00 read → 5.1 read $0.25 [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md) [SECONDARY](https://aiweekly.co/alerts/anthropic-ships-fable-51-and-mythos-51-cuts-cache-reads-75)
- Sonnet 5 cache ladder: $2 input | $10 output (introductory-made-permanent) [SECONDARY](https://www.worthview.com/claude-sonnet-5-is-here-anthropics-most-agentic-sonnet-model-closes-the-gap-with-opus-4-8/)
- Revenue run rate: $9B (end-2025) → $30B (Apr 2026) → $47B (May) → $65B (end-Jul); Q2 2026 $11.5B prelim; Claude Code $2.5B run rate by Feb 2026 [SECONDARY](https://www.ainvest.com/news/65-billion-question-2608/)
- Cowork adoption signal: >90% of Cowork usage non-software-development by July 2026 [SECONDARY](https://github.com/adnova-group/muster/blob/HEAD/docs/research/claude-cowork.md)
- Compute economics: $0.71 compute cost per revenue dollar (Q1 2026) → projected $0.56 (Q2); gross margin −94% (2024) → 44–60% range (2026) [SECONDARY](https://www.ainvest.com/news/15-000x-cost-number-anthropic-26-96-safety-headline-graded-1-2608/)

## Main actors
- **Anthropic** — dual-deployment architecture (Fable open API / Mythos restricted); 60–90-day deprecation cadence; TB 4.0 and LMArena-text leader; did not sign the July 24 letter [SECONDARY].
- **Project Glasswing** — the restricted-deployment program gating Mythos variants [SECONDARY].
- **Fable 5 / 5.1** — the API flagship line: same weights as Mythos, different safeguards [SECONDARY].
- **Mythos 5 / 5.1** — the restricted counterpart; highest vendor-claimed TB 4.0 score (60.9%) [VENDOR].
- **The endoflife-date project** — tracks the Claude retirement schedule (link below) [COMMUNITY].

