---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/11-other-notable-ai-developer-tools-platforms-febsep-2026
title: "11. Other notable AI developer tools / platforms (Feb–Sep 2026)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Baseten", "Cerebras", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "Microsoft", "Nvidia", "OpenAI", "OpenRouter", "SpaceX", "Stripe", "Together AI", "United States", "Z.ai", "xAI"]
dates: ["2026-01", "2026-02", "2026-05", "2026-05-19", "2026-05-20", "2026-05-24", "2026-06", "2026-06-08", "2026-06-11", "2026-06-21", "2026-07", "2026-09"]
keywords: ["acquisition", "agent", "agentic", "agents", "attribution", "aws", "bedrock", "benchmark", "claude", "consumer", "copilot", "cost"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [562, 660]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 6946bb1fdca07ef4c36c5366506fce2ecc6d1f397011b72595b3dd9a5799f3c8
---

# 11. Other notable AI developer tools / platforms (Feb–Sep 2026)

## 11. Other notable AI developer tools / platforms (Feb–Sep 2026)

### 11.1 GitHub Spec Kit — late February 2026 **[secondary/unverified]**
Open-source (MIT) toolkit that structures AI-assisted development into phases, each producing artifacts that feed the next — the flagship of the **spec-driven development (SDD)** wave that displaced "vibe coding" by mid-2026. Reported **111K GitHub stars and 9.8K forks as of June 11, 2026**, described as the fastest-growing dev-tooling repo in recent memory (via the atomsbaza research compendium on GitHub).
- Source: https://github.com/atomsbaza/my-superpowers/blob/HEAD/docs/research/agentic-ai/2026-06-21-agentic-ai-coding-agent-trends.md

### 11.2 Google Antigravity 2.0 — May 19, 2026 (Google I/O) **[vendor-reported]**
Google's agent-first development platform, no longer an in-product experiment but a standalone production-grade offering. Four components shipped simultaneously: standalone **desktop app** (macOS/Windows/Linux) for multi-agent orchestration, **CLI**, **SDK** (programmatic access to the agent harness behind Google's consumer products; supports self-hosted deployment), and the **Gemini Enterprise Agent Platform** (managed sandboxed execution via the Gemini API). Powered by **Gemini 3.5 Flash** by default; native voice commands; integrations with Google AI Studio, Android Studio, Firebase. Positioned directly against Claude Code, Cursor, and OpenAI's Codex. Sources: Google Developers Blog, SiliconANGLE, MarkTechPost.
- Sources: https://github.com/archie0125/synapse-news/blob/HEAD/src/content/articles/2026-05-24-google-antigravity-20-agentic-dev-platform-en.md ; https://github.com/kovalovme/ai-news/blob/HEAD/news/2026-05-20/tools.md

### 11.3 AWS Kiro — June 2026 (AWS Summit New York) **[secondary/unverified]**
Agentic IDE powered by **Claude Sonnet + Amazon Nova via Bedrock**. Spec-driven: takes a prompt, generates a formal EARS-notation requirements doc with acceptance criteria, waits for human review, then generates code. (Via atomsbaza compendium — verify against AWS announcements before final consolidation.)
- Source: https://github.com/atomsbaza/my-superpowers/blob/HEAD/docs/research/agentic-ai/2026-06-21-agentic-ai-coding-agent-trends.md

### 11.4 Cline "Spec Driven for Enterprise" — June 8, 2026 **[secondary/unverified]**
**LG CNS × Cline** (US open-source AI coding-agent developer) partnership: agentic platform automating the full enterprise system lifecycle (requirements analysis, design, coding, QA) with per-stage specialized agents. Announced via PRNewswire relay (vir.com.vn).
- Source: https://vir.com.vn/lg-cns-and-cline-launch-agentic-ai-platform-for-enterprise-system-development-and-operations-154352.html

### 11.5 Z.ai ZCode — early July 2026 **[secondary/unverified]**
See §6 (full entry). Noted as a pricing disruption even as real-world robustness remained unproven (WindowsForum relay, July 2026).
- Source: https://windowsforum.com/windows-news.4/z-ai-zcode-launch-free-agentic-ai-coding-desktop-on-windows-macos-linux.433572/

### 11.6 Huawei Cloud CodeArts Agent — July 2026 (open beta, Thailand) **[vendor-reported]**
Launched at Huawei Cloud Summit Thailand 2026 as part of Huawei Cloud's **Agentic Infrastructure**: project-level code generation, completion, R&D knowledge queries, unit-test generation; "Agent Team" mode for multi-agent collaboration; specification-driven development workflow. The wider Agentic Infrastructure adds a UnifiedBus AI cluster service, petabyte-scale agentic memory storage, the AgentSphere runtime, and CCE VolcanoNext scheduling (TechNode/TNGlobal).
- Source: https://technode.global/2026/07/24/chinas-huawei-cloud-launches-agentic-ai-infrastructure-coding-agent-beta-in-thailand/

### 11.7 Augment Code Cosmos — mid-September 2026 **[secondary/unverified]**
**Augment Code** launched **Cosmos**, a team-level AI development platform: instead of agents per developer, a shared software-delivery surface where a validator agent checks specs against codebase and company practices, and corrections (e.g. via Slack) persist as shared cross-session memory for all agents and team members. Framing from the company: "2024 was chat, 2025 is agents, 2026 is agents for teams" (completeaitraining.com).
- Source: https://completeaitraining.com/news/augment-code-launches-cosmos-to-extend-agentic-ai-coding/

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
