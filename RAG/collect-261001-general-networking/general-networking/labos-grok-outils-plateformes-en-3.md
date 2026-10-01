---
id: collect-261001-general-networking/general-networking/labos-grok-outils-plateformes-en-3
title: "VOLET 1 — Vague 2 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "alignment", "astra", "aws", "bedrock", "benchmark", "claude", "copilot", "cost", "cyber"]
source: docs/RAG/collect-261001-general-networking/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [142, 203]
sha256: d47a890249d027261027e106a941612233ba1019b6c2f310ebd488660ce89a07
---

# VOLET 1 — Vague 2 (EN)

## 5.1 Product timeline (2026)
| Date | Milestone |
|---|---|
| Feb 5 | **Opus 4.6** launch; **agent teams** in research preview (multiple independent sessions sharing work via team structure) |
| Feb 20 | **Claude Code Desktop** review tools (preview running app, PR review, CI-failure response); move sessions between Desktop / web / mobile / CLI; **Claude Code Security** limited research preview (vuln scanning + patch proposals, Team/Enterprise) |
| Feb 25 | **Remote Control** research preview (Pro/Max first): continue local sessions from phone/tablet/browser via claude.ai/code or mobile app; execution stays on local machine |
| Feb 25 – Mar 19 | Auto memory for session context; named Remote Control sessions; `--channels` preview (MCP servers push events into sessions) |
| Mar 24 | **Auto mode** research preview: classifier-gated permissions — safe tool calls auto-approved, destructive/irreversible ones blocked or prompted; the middle path between default approvals and `--dangerously-skip-permissions` |
| May 28 | Opus 4.8 release (coding/long-agent Opus update) |
| Jun 2 | **Dynamic workflows** preview: task-specific multi-agent workflows composed on the fly |
| Jun 9 | Fable 5 / Mythos 5 announced; Fable 5 becomes default Claude Code model for Pro/Max |
| Jun 12 – Jul 1 | Fable 5 pulled under export controls; restored |
| Jun 22 | Auto-mode safety guards (destructive git/terraform commands blocked); `/config key=value`; implicit agent teams; `Tool(param:value)` permission rules |
| Jun 30 | **Claude Sonnet 5** release |
| Jul 10 | **Auto mode GA** |
| Jul 24 | **Opus 5** becomes default Opus model; dynamic workflow-size settings; sandbox network allowlisting; deeper nested-subagent delegation |
| Aug 4 | Focus view (VS Code); Linux/WSL sandbox credential masking; `/fork` worktrees |
| Aug 7 | **Self-hosted runners** (`claude self-hosted-runner`, Team/Enterprise); **cross-session messaging** (`SendMessage`/`ListAgents` across machines) |
| Aug 10 | Plugin install from SHA-256-pinned HTTPS archives |
| Aug 13 | Default subagent forking (inherits conversation + prompt cache); GitLab workflows |
| Aug 14 | **Auto mode becomes default** for new sessions (Pro/Max/Team); Enterprise/API within a month. Rationale: users approved 97% of prompts (permission fatigue); auto mode blocked 89% of deliberately dangerous commands vs 13.6% caught by humans (1,053-action study) |
| Sep 7 | Friendly model names in `/model`; auto-compact tuned for 1M-context models |
| Sep 22 | **Opus 5.5** (`claude-opus-5-5`) becomes default Opus model; changelog notes 1M context, $4/$20, $0.20 cache reads |

## 5.2 Dev Team mode (multi-agent orchestration)
- Shipped with the Sonnet 5 "Fennec" wave (Feb 2026): a **Manager agent decomposes a high-level goal and spawns specialized sub-agents** (Backend, QA, Infrastructure/Researcher) that work in parallel on different files — e.g., Backend writes an API route while QA generates unit tests — then the Manager reconciles conflicts and presents the final PR.
- Community skill/plugin ecosystem formalized the pattern: Architect → Coder → Tester → Docs pipelines with plan-approval gates (`dev-team` skills, plan/dev/full/auto modes).
- Enables bug-report → write → test → verify patch loops and full-feature builds from briefs — described as a "human parity" milestone for junior-to-mid-level dev work.

## 5.3 Integrations & surfaces
- **GitHub Copilot**: Opus 5.5 day-one (Sept 22) — VS Code, Visual Studio, Copilot CLI, coding agent, JetBrains, Xcode, Eclipse; billed at provider list pricing.
- **Cloud platforms**: AWS/Bedrock, Google Cloud (incl. Agent Platform), Microsoft Foundry, Azure — all carrying Opus 5.5 at launch.
- **IDE surfaces**: VS Code extension (Focus view), JetBrains, Xcode, Eclipse via Copilot; Slack `@Claude` mention → Claude Code web session with PR link back in thread.
- **Enterprise**: Admin API (beta) for org/member/role management; Managed Agents with effort settings; Enterprise Frontier Safeguards (fall 2026).

## 5.4 What distinguishes Claude Code in 2026
- From terminal CLI to **full multi-surface agent platform** (CLI, Desktop, web, mobile, IDE, Slack) with session portability; from per-action approvals to **classifier-gated auto mode** as default; from single-agent to **agent teams / dynamic multi-agent workflows** as a first-class primitive — the reference implementation of the agentic-coding loop every lab now copies.

---

# 6. Cross-cutting: Anthropic's 2026 model ladder (Sept 22 snapshot)

| Tier | Model | Release | In / Out ($/1M) | Context | Signature |
|---|---|---|---|---|---|
| Mythos (restricted) | Mythos 5.1 | Sep 1 | — (Glasswing only) | 1M | Uncapped frontier; vetted partners |
| Fable (public flagship) | Fable 5.1 | Sep 1 | $10 / $50 | 1M | Science agents, alignment; watermarking |
| Opus (premium) | **Opus 5.5** | **Sep 22** | **$4 / $20** | 1M | Fable-class at 40% less; default Opus |
| Opus (prev) | Opus 5 | Jul 24 | $5 / $25 | 1M | SWE-bench Verified 96% |
| Sonnet | Sonnet 5 "Fable"? "Fennec" | Feb 3 (+Jun 30 build) | $3 / $15 | 1M | Dev Team; Opus-killer |
| Haiku | Haiku 5.x | — | — | — | Fast/cheap tier |

*(Sonnet 5.5 / Haiku 5.5 announced as coming in the weeks after Sept 22.)*

# 7. Caveats & methodology notes
- Opus 5.5 figures are **<24h old** at research time — all vendor-reported, none independently replicated (Digital Applied).
- AA Intelligence Index absolute values differ across trackers/snapshots (58 vs 66 for Fable 5.1-class) — record methodology version.
- "Opus 5.2" was never a product — don't cite as a released model.
- Fable 5.1 benchmark rows where safeguards triggered scored zeros or Opus fallbacks — understates raw capability on cyber/bio suites.
- Opus 5.5's per-task token appetite (119K vs Astra's 17K on AA tasks) can erase sticker-price savings — cost-per-task, not cost-per-token, is the decision metric.

---

