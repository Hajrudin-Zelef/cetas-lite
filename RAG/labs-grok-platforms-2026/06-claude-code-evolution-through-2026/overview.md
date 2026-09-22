---
id: labs-grok-platforms-2026/06-claude-code-evolution-through-2026/overview
title: "5. Claude Code — evolution through 2026"
domain: claude-code-evolution-through-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["claude", "agent", "agentic", "agents", "aws", "bedrock", "copilot", "export controls", "fable 5", "foundry", "mcp", "memory"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [140, 181]
section: "5. Claude Code — evolution through 2026"
sha256: 490aa4e7915a5d8d71eb3ea1d9552f01222875ff1411aae24d3fe61c06b63e44
---

# 5. Claude Code — evolution through 2026

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

