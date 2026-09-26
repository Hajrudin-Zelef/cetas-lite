---
id: ai-industry-kb-2026-wave6/13-anthropic/claude-opus-4-6-full-spec-sheet
title: "Claude Opus 4.6 — full spec sheet"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic"]
dates: ["2025-05", "2025-08", "2026-02", "2026-02-05"]
keywords: ["claude", "opus 4", "agent", "agi", "aws", "bedrock", "benchmarks", "cost", "foundry", "gemini", "pricing", "tool use"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6321, 6367]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: ddf9687f33dea8c2a5e0399300061aca75b6bf1457ea8a157cded6ab81a37342
---

# Claude Opus 4.6 — full spec sheet

### Claude Opus 4.6 — full spec sheet
- Released 2026-02-05 [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Model ID `claude-opus-4-6` [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Context 1M tokens (initially beta, reportedly GA in March) [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29) [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Max output 128K [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Knowledge cutoff May 2025 [SECONDARY](https://tech-insider.org/claude-opus-vs-sonnet-vs-haiku-2026/)
- Input pricing: $5/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Output pricing: $25/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 5-min cache write: $6.25/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- 1-hr cache write: $10/M [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- Cache read: $0.50/M (−90%) [SECONDARY](https://github.com/mikahniehaus/claudeboost/blob/HEAD/.claudeboost/knowledge/batch-api-cost-about-claude-pricing.md)
- LONG-CONTEXT SURCHARGE CONTRADICTION: sources report either $10/$37.50 or $7.50/$37.50 above 200K input — unresolved; date every figure [UNVERIFIED](https://aithinkerlab.com/claude-opus-4-6-vs-opus-4-5-benchmarks-pricing-adaptive-thinking/)
- Available on Claude.ai [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Available on the Anthropic API [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Available on AWS Bedrock [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Available on Vertex AI [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Available on Azure Foundry [SECONDARY](https://dev.to/bokaai/claude-opus-46-a-first-person-review-from-an-ai-agent-actually-running-on-it-1l29)
- Features: Agent Teams [SECONDARY](https://github.com/gofinkle/pilot-shell/blob/HEAD/docs/site/src/content/blog/claude-opus-4-6.md)
- Features: context compaction [SECONDARY](https://github.com/gofinkle/pilot-shell/blob/HEAD/docs/site/src/content/blog/claude-opus-4-6.md)
- Features: adaptive thinking [SECONDARY](https://github.com/gofinkle/pilot-shell/blob/HEAD/docs/site/src/content/blog/claude-opus-4-6.md)
- Vendor benchmarks (vendor claims, not independent): Terminal-Bench 2.0 65.4% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Vendor: ARC-AGI-2 68.8% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Vendor: OSWorld 72.7% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Vendor: MRCR v2 76% [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- Vendor: GDPval-AA 1606 Elo [SECONDARY](https://www.humai.blog/claude-opus-4-6-vs-gpt-5-2-vs-gemini-3-pro-which-ai-model-should-you-actually-use-in-2026/)
- CACHING ANECDOTE: one community report claims `claude-opus-4-6` did not support prompt caching on their account (cache_write=0 in direct API tests) — single anecdotal account, contradicts the documented cache pricing table; treat as [COMMUNITY] only [COMMUNITY](https://github.com/guru-labs-ai/ai-sponsor/commit/383a64c8a28ebf9de3b1ca1ce3fa8e26b07b7bdc)
- OUTLIER PRICE: one outlier secondary source lists Opus 4.6 at ~$15/$75 — contradicts the documented $5/$25 table; treat the $15/$75 figure as erroneous [UNVERIFIED](https://www.sostav.ru/blogs/278670/84700)

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

