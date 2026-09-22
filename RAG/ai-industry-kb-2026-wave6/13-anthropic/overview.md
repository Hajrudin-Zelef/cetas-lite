---
id: ai-industry-kb-2026-wave6/13-anthropic/overview
title: "§13. Anthropic"
domain: anthropic
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "Glasswing", "OpenAI"]
dates: ["2025-05", "2026-02-05", "2026-02-17", "2026-04-07", "2026-04-16", "2026-05-28", "2026-06-09", "2026-06-12", "2026-06-15", "2026-06-30", "2026-07-24", "2026-08-05", "2026-08-11", "2026-09-01"]
keywords: ["agent", "agentic", "agi", "astra", "aws", "bedrock", "benchmark", "benchmarks", "claude", "cost", "distribution", "fable 5"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [6252, 6348]
section: "§13. Anthropic"
delta_of: ai-industry-kb-2026
sha256: fbb8d2a931096fdc28c0b64ef507b9520ff0b887ff40c994fe61c89ecffda8fb
---

# §13. Anthropic

Keywords: Anthropic, Claude, Claude Opus 4.6, Claude Sonnet 4.6, Claude Mythos Preview, Claude Opus 4.7, Claude Opus 4.8, Claude Fable 5, Claude Mythos 5, Claude Sonnet 5, Claude Opus 5, Claude Fable 5.1, Claude Mythos 5.1, Project Glasswing, Fennec, model retirement, Commerce Department suspension

## Summary
- Anthropic's 2026 cadence ran from Opus 4.6 (Feb) through a rapid-fire 4.7/4.8 sequence into the **Fable 5 / Mythos 5** dual-deployment generation (Jun 2026), then Sonnet 5, Opus 5, and the September **Fable 5.1 / Mythos 5.1** refresh. Model dates are [SECONDARY] unless noted.
- **Fable 5 and Mythos 5 are the same weights under different safeguards** — Mythos is the restricted deployment, gated through **Project Glasswing** [SECONDARY]; the split-deployment architecture is Anthropic's signature 2026 move.
- Anthropic runs an aggressive **60–90-day deprecation cadence**: Opus 4 retired 2026-06-15, Opus 4.1 retired 2026-08-05 [SECONDARY].
- Benchmark leadership as of Sept 2026 (version-pinned, [SECONDARY]): Fable 5.1 is **#1 on Terminal-Bench 4.0** (official 57.9%), ties Astra at **53** on AA Index v4.3, and Fable 5 leads the LMArena text board at ~1525; SWE-bench Verified is effectively saturated with Opus 5 at 97.0%.
- The **June 12, 2026 Commerce Department order suspending Fable 5 / Mythos 5 worldwide** (later re-enabled) is **[SECONDARY], single-source class** — export-control-shaped action against a closed flagship, not open weights.
- Anthropic did not sign the July 24 "Open Weights and American AI Leadership" letter and is described as "the most visible holdout" from the open-source cause since the July 21 HF incident [SECONDARY]; its CEO later said Anthropic "has never advocated banning open-weight models" [SECONDARY].
- **"Claude 4.62"** has zero source hits and is treated as garbled "4.6" — it is dropped, not carried as a model.

## Key dated facts
### Opus 4.6 / Sonnet 4.6 — the February pair
- **2026-02-05** — Claude **Opus 4.6** launched [SECONDARY].
- **2026-02-17** — Claude **Sonnet 4.6** launched [SECONDARY].
- These two are the dated 4.x anchor points; later 4.x versions (4.7, 4.8) are secondary/community-timeline sourced only.

### Mythos Preview and Project Glasswing
- **2026-04-07** — **Mythos Preview** launched [SECONDARY].
- **Project Glasswing** is the restricted-deployment program through which Mythos variants are gated [SECONDARY] — it becomes the distribution channel for the later Mythos 5 / 5.1 releases.
- AWS Bedrock hosts the Mythos 5 model card under this restricted program (link below).

### Opus 4.7 / Opus 4.8 — the spring sequence
- **2026-04-16** — Claude **Opus 4.7** [SECONDARY, community-timeline sourced].
- **2026-05-28** — Claude **Opus 4.8** [SECONDARY, community-timeline sourced].
- LMArena text (Sept 2026): Opus 4.8 at ~1512 [SECONDARY]; an Alibaba vendor table puts Opus 4.8 at 69.2 (vendor-reported class) [SECONDARY].
- Vals.ai's 2026-06-09 Fable 5 evaluation explicitly notes **Opus-4.8 fallback due to Fable 5 refusals** [SECONDARY] — refusals were load-bearing enough to substitute the judge model.

### The retirement cadence — Opus 4, Opus 4.1
- **2026-06-15** — Claude **Opus 4 retired** [SECONDARY].
- **2026-08-05** — Claude **Opus 4.1 retired** [SECONDARY].
- Pattern: a **60–90-day deprecation cadence** [SECONDARY] — operators pinning to dated Claude versions face forced migration roughly quarterly.

### Fable 5 + Mythos 5 — the dual deployment (June 9)
- **2026-06-09** — **Claude Fable 5 and Claude Mythos 5** launched: **same weights, different safeguards**; Mythos 5 is restricted through Glasswing [SECONDARY].
- This is the corpus's canonical dual-deployment example (see §21 licensing-politics): one checkpoint, two policy surfaces.
- The **Commerce Department suspension order of 2026-06-12** — three days after launch — suspended Fable 5 / Mythos 5 worldwide and was later re-enabled **[SECONDARY, single-source class]**; directionally consistent with BIS's closed-weight posture (closed weights are controllable; open weights, once distributed, cannot be recalled).
- LMArena text (Sept 2026): Fable 5 at **~1525 (#1)**, Opus 5 at 1522 [SECONDARY].
- Fable 5's vendor-reported Terminal-Bench 2.1 figures show harness/effort variance: 88.0 [VENDOR] vs 84.3 [COMMUNITY] vs 80.52 [SECONDARY] — the vals.ai figure is partly Opus-4.8-substituted (see above).

### Sonnet 5 — "Fennec"
- **2026-06-30** — Claude **Sonnet 5** launched [SECONDARY]; the internal codename **"Fennec"** is [SECONDARY].
- Pricing: **$2.00/$10.00** input/output, cache read $0.20 [VENDOR]; the planned Sept-1 rise to $3/$15 was **cancelled 2026-08-11** — the $2/$10 intro price was made permanent [SECONDARY].
- LMArena text: Sonnet 5 at 1479 [SECONDARY].

### Opus 5 — July flagship
- **2026-07-24** — Claude **Opus 5** launched **[SECONDARY]**.
- Pricing: **$5.00/$25.00**, cache read $0.50 [VENDOR].
- SWE-bench Verified: Opus 5 at **97.0%** [SECONDARY] — the board is effectively saturated as a frontier differentiator.
- Terminal-Bench 4.0: Opus 5 at **51.8%** (official runs) [SECONDARY]; vendor self-report: 52.3% [VENDOR].
- AA Index v4.3: Opus 5 at **51** [SECONDARY]; LMArena text: 1522 [SECONDARY].
- Of one observer's framing: Opus 5 "ranks seventh [on the Aug-31 text pull] and first on both task benchmarks — if you pick a model off LMArena for a coding agent, you are optimising for the wrong thing" [SECONDARY].

### Fable 5.1 + Mythos 5.1 — the September refresh
- **2026-09-01** — **Claude Fable 5.1 and Claude Mythos 5.1** launched [SECONDARY].
- Pricing: **$10.00/$50.00** input/output, cache read **$0.25**; Mythos 5.1 restricted through Glasswing [VENDOR/SECONDARY].
- AA Intelligence Index v4.3 (Sept 7): Fable 5.1 (max w/ fallback) ties GPT-6 Astra (max) at **53** [SECONDARY]; cost-per-task **$7.63** vs Astra's $3.26 — 57% more expensive for the same index score [SECONDARY].
- Terminal-Bench 4.0: **Fable 5.1 57.9%±3.8 (#1)** official runs [SECONDARY]; vendor self-reports: Fable 5.1 55.8%, **Mythos 5.1 60.9%** [VENDOR] — the highest vendor-claimed agentic-coding figure in the wave.
- SWE-bench Verified: Fable 5 Max at **95.0%** [SECONDARY].
- Note the no-cross-version rule: AA Index was revised three times in under a week (v4.1.1 Sept 3 → v4.2 Sept 4 → v4.3 Sept 7); scores from earlier versions are not comparable (see §22 benchmarks).

### The dropped "Claude 4.62"
- "Claude 4.62" has **zero source hits** and is treated as a probable garbling of "4.6" — it is dropped entirely, not carried as a model or a date.


### New verified facts — expansion

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

