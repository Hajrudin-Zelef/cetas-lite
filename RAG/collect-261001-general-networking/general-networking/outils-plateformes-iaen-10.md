---
id: collect-261001-general-networking/general-networking/outils-plateformes-iaen-10
title: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Baseten", "Cerebras", "China", "DeepSeek", "Google", "Hugging Face", "Microsoft", "Nvidia", "OpenRouter", "SpaceX", "Stripe", "Together AI", "Z.ai", "xAI"]
dates: ["2026-01", "2026-05", "2026-09"]
keywords: ["acquisition", "agent", "agents", "attribution", "aws", "benchmark", "claude", "consumer", "copilot", "cost", "deepseek", "disclosure"]
source: docs/RAG/collect-261001-general-networking/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [592, 660]
sha256: 0b6390a6eee6c591f366b3302e4ce19a134305951ea536255c4448d30c0e63e6
---

# Step 2 — AI Tools & Platforms (Feb–Sep 2026)

### 11.8 Items flagged but not verified (excluded pending confirmation)
- **Anthropic "Code with Claude" managed agents** (mid-May 2026) — referenced only via a news-aggregator slug; **no primary source located** — do not include without verification.
- **BMAD-METHOD** (open-source multi-agent spec-execution framework) — named in the SDD wave but not independently researched here.
- **Grok Bot** — official xAI harness name with no standalone documentation located.

---

## 12. Cross-tool comparison (September 2026 snapshot)

| Dimension | Claude Code | OpenCode | OpenClaw | Cursor | VS Code + Copilot | ZCode | Grok Build | DeepSeek Harness |
|---|---|---|---|---|---|---|---|---|
| Form factor | Terminal CLI + IDE ext. + remote | Terminal TUI (+ desktop beta) | Gateway + chat apps | AI-native IDE | Editor + agent mode | Desktop ADE (Electron) | Terminal CLI | Web UI + headless CLI + SDK |
| License | Proprietary (unconfirmed) | **MIT** | **MIT** | Proprietary | Proprietary | Proprietary client | Proprietary | **MIT** |
| Model lock-in | Claude only (official) | None (75+ providers) | None (any API/local) | Multi-model + Composer 2 | Multi-model (BYOK on Biz/Ent) | GLM-first, BYOK | Grok (xAI) | None (plugin-adapted) |
| Pricing entry | $20/mo | $0 + usage | $0 + own keys | $0 / $20/mo | $0 / $10/mo | $0 / $16.20/mo | SuperGrok Heavy early; 4.7 free in CLI | $0 (v0.1 preview) |
| Heavy tier | $200/mo | Zen PAYG / Go $10/mo | — | $200/mo | $100/mo | $144/mo | — | — |
| Standout 2026 feature | Cross-session messaging; plugin eval | Parallel agents; LSP grounding; Zen gateway | Atomic updates; masked credentials; 30 channels | Agent-first UI 3.0; cloud handoff | Browser tools GA; parallel sessions; AI credits | WeChat/Feishu remote steering; price | ACP support; parallel subagents in worktrees | Everything-is-a-plugin; append-only replayable session log |

**Platforms comparison:**

| Dimension | OpenRouter | Hugging Face |
|---|---|---|
| Role | Unified inference gateway (400+ models, 80+ providers) | Model/dataset/app hub (3M+ models, 1M+ Spaces) |
| Business model | ~5% platform fee on routed inference spend | Subscriptions (PRO $9 / Team $20 / Enterprise $50+) + usage billing; inference providers at cost |
| 2026 headline | Stripe acquisition announced Aug 19 (~$7.5B reported) | NVIDIA acquisition announced Sep 3 ($12.93B, close expected H1 2027) |
| Developer signal | Agents = 71% of token consumption (Aug 2026); Chinese open models >50% of usage | HF agent-usage dataset: Claude Code led July with 44.4% |

---

## 13. Open verification items (before final RAG consolidation)

### Coding agents & IDEs
1. Claude Code client license terms — verify from the official repo.
2. OpenCode "7.5M monthly developers" — vendor claim, unverified.
3. Cursor Pro allowance — "500 fast requests" (older docs) vs. "$20 credit pool ≈225 fast requests" (2026 credit docs); resolve against cursor.com/pricing.
4. ZCode current version (3.11.2?) and GLM-5.2 vs. GLM-5.3 tuning — verify at https://zcode.z.ai.
5. OpenClaw founder/foundation attribution — verify current stewardship.
6. Cursor funding/valuation and adoption numbers — not researched in this pass.
7. Copilot "Max" tier details and flex-credit policy changes — GitHub notes flex allotments may change; re-check the plans page.
8. DevToolPicks' claim of GitHub suspending new flat-rate Copilot signups — single secondary source; verify.
9. **SpaceXAI acquisition of Cursor** — single-source secondary claim; not confirmed by any official announcement; verify before final consolidation (belongs to step 3, labs/hyperscalers).

### Grok / DeepSeek Harness
10. DeepSeek Harness: exact GitHub repo URL, the arXiv paper ID (2608.25512), the `deepseek-v4-flash` default-model claim, and the four session-mode names — all currently **[secondary/unverified]**.
11. "Grok Code" as a product name is **unverified**; the real product is **Grok Build**.
12. xAI–SpaceX merger: no official announcement located; rests on secondary relays.
13. Grok 4.20 API GA (Mar 10, 2026): no official xAI launch post located.
14. Grok Collections (RAG building block): mentioned in one January 2026 overview; no official docs page located.
15. Grok Bot: official xAI harness name with no standalone documentation located.
16. API/consumer prices are snapshots dated September 21–22, 2026; xAI reprices frequently.

### Platforms
17. **OpenRouter current fee**: confirm the live platform-fee % and any post-acquisition Stripe pricing changes on openrouter.ai/docs/pricing. **[secondary only]**
18. **OpenRouter revenue figures**: $50M vs $140M annualized — conflicting secondary estimates; seek a company disclosure or reconcile dates.
19. **Stripe deal status**: closing expected "within weeks" of Aug 19 — check whether it has closed and whether terms were ever disclosed.
20. **NVIDIA/HF deal status**: definitive agreement signed; expected H1 2027 close pending regulators — track regulatory review progress.
21. **HF official pricing**: confirm PRO/Team/Enterprise prices and ZeroGPU quotas directly on huggingface.co/pricing (third-party guides disagree on quota minutes).
22. **Series B announcement URL**: https://openrouter.ai/announcements/series-b cited by community briefs; not directly fetched — verify.
23. **HF Inference Providers roster**: Baseten, Together AI, AWS, Google Cloud, Cerebras cited across sources; get the full current list from HF docs.

---

## 14. Collection metadata
- **Research waves:** Part A (coding agents & AI IDEs), Part B (OpenRouter & Hugging Face), Part C (Grok tooling, DeepSeek Harness, other tools).
- **Working files:** `etape2_partA.md`, `etape2_partB.md`, `etape2_partC.md` (kept alongside this consolidated file).
- **Priority sources used:** SiliconANGLE, VentureBeat, TechCrunch, Bloomberg, NYT, WSJ, The Register, MacRumors, official vendor blogs/repos/changelogs, Hugging Face blog, GitHub Changelog, SEC 8-K filings, community trackers.
- **Conventions kept:** per-item provenance tags; uncertainty flags; dates on every price, version, and star count; no direct comparison across different benchmark versions; vendor-reported vs independent numbers separated (e.g. Grok 4.7 CursorBench/DeepSWE vendor-reported vs Artificial Analysis independent).

*End of consolidated Step 2 report.*
