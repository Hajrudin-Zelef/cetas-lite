---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/overview
title: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "China", "DeepSeek", "Google", "Huawei", "Hugging Face", "Microsoft", "Nvidia", "OpenRouter", "SpaceX", "Stripe", "United States", "xAI"]
dates: ["2025-02", "2026-04", "2026-04-21", "2026-05", "2026-06", "2026-08-07", "2026-08-08", "2026-09", "2026-09-18", "2026-09-22"]
keywords: ["acquisition", "agent", "agentic", "agents", "attribution", "aws", "bedrock", "claude", "context window", "copilot", "cost", "deepseek"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [1, 68]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 8d36ef38a8c2c37770cf3e40678a8e5deaddd1e54c4e46c2559c588dd3f4ef78
---

# Step 2 — AI Tools & Platforms (Feb–Sep 2026)
## Consolidated research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 2 (Outils & plateformes IA)
**Coverage date:** February → September 22, 2026
**Consolidation:** September 22, 2026
**Status:** Research snapshot. Prices are dated snapshots; re-verify against official pricing pages before use. Star counts move weekly.

### Scope
- Coding agents & AI IDEs: Claude Code, OpenCode, OpenClaw, Cursor, VS Code + GitHub Copilot, ZCode, Grok Build, DeepSeek Harness
- Distribution platforms: OpenRouter, Hugging Face
- Other notable developer tools: GitHub Spec Kit, Google Antigravity 2.0, AWS Kiro, Cline × LG CNS, Huawei Cloud CodeArts Agent, Augment Code Cosmos

### Provenance legend
- **[official]** — vendor blog, docs, changelog, repository, or regulatory filing owned by the tool's maker.
- **[vendor-reported]** — figures claimed by the vendor (pricing, user counts, performance) without independent audit.
- **[independent]** — reputable third-party press/analysis (VentureBeat, SiliconANGLE, The Register, TechCrunch, Bloomberg, NYT, WSJ, MacRumors, etc.).
- **[secondary]** — community sources, docs mirrors, trackers, aggregator blogs; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

### Cross-cutting notes for RAG indexing
- **2026 = the year of the agent harness.** Products converged on the same patterns: parallel subagents, plan/approval modes, git-worktree isolation, `AGENTS.md` project memory, MCP/plugin extensibility, remote/mobile steering, and headless/CI modes.
- **Spec-driven development (SDD)** displaced "vibe coding" as the dominant workflow narrative from mid-2026 (GitHub Spec Kit, AWS Kiro, CodeArts Agent).
- **Naming correction:** no official xAI product named "Grok Code" was found. The verified xAI product is **Grok Build** (CLI + model `grok-build-0.1`). "Grok Code" appears only informally in community writing.
- **"Jev"** (from the step-1 roadmap) was previously identified as a proprietary decision model from US-based TypeSafe AI — not open-weight, not Chinese — and does not appear in this step.
- Multiple large M&A moves hit this space in 2026 (see §10–11): Stripe→OpenRouter (announced Aug 19), NVIDIA→Hugging Face (announced Sep 3), plus a **single-source, unverified** claim of a SpaceXAI→Cursor acquisition. All pending acquisitions are regulatory-uncertain.

---

## 1. Claude Code (Anthropic)

### 1.1 Overview
Terminal-first agentic coding tool by Anthropic, originally released February 2025 alongside Claude Sonnet 3.7 **[independent — VentureBeat]**. Evolved from a Node.js CLI into a platform shipping as native Rust-compiled binaries, with a plugin marketplace, agent teams, and cloud/mobile remote-control surfaces **[secondary — toknow.ai, May 2026]**.

### 1.2 Latest versions (as of September 2026)
- **2.1.277** — September 18, 2026 (latest changelog entry found) **[official — generated from https://github.com/anthropics/claude-code CHANGELOG.md]**.
- **2.1.224** — August 7, 2026: cross-session messaging on macOS/Linux via two new tools, `ListAgents` and `SendMessage`; messages stay local on the same machine (never reach Anthropic servers); cross-machine messaging only as replies through Remote Control; also shipped `claude self-hosted-runner` for Team/Enterprise plans **[independent — MacRumors, 2026-08-08]**.
- **2.1.0** — infrastructure release: hooks for agents/skills/slash commands (`PreToolUse`, `PostToolUse`, `Stop`), hot reload for skills in `~/.claude/skills`, forked sub-agent context (`context: fork`), wildcard tool permissions (e.g. `Bash(npm *)`) **[independent — VentureBeat]**.
- June 2026 train (2026.6): nested sub-agents up to 3 levels deep, `fallbackModel` configuration (up to three fallback models), community tool marketplace, usage attribution with per-agent/per-task cost breakdowns, scoped permissions, streaming agent logs (beta), agent checkpointing (beta), inline cost budgets, custom agent templates, multi-repo orchestration (beta) **[secondary — SitePoint, June 2026]**.
- September 2026 entries: AGENTS.md support (2.1.277 — in projects without CLAUDE.md, Claude Code reads AGENTS.md instead); `claude plugin eval` (2.1.269); `/output-style` command (2.1.269); `CLAUDE_CODE_WORKFLOW_MAX_CONCURRENT_AGENTS` up to 256 (2.1.269); per-command `allowed_domains` in auto mode with sandboxing (2.1.271); fast mode in Remote sessions (2.1.271); gateway hint headers for LLM gateways (2.1.273) **[official — CHANGELOG via community mirrors]**.

### 1.3 Features & architecture
- **Agent runtime:** terminal CLI + IDE extensions + Remote Control via claude.ai/code; subagents, Agent Teams ("Claude Cowork"), background agents, plans mode (with "Ultraplan" in newer builds), auto mode classifier (billing docs at code.claude.com), `/rewind` (undo agent runs), `/plan`, `/btw`, vim mode, voice mode (reported 20 languages), Playwright MCP / computer use, DESIGN.md support, MEMORY.md persistent memory **[secondary — toknow.ai]**.
- **Integrations:** MCP servers (with mid-session disconnect notifications and reconnection), first-party GitHub Action (runs the same runtime in CI), skills in `.claude/skills/` with `SKILL.md` frontmatter, plugin system with marketplace, OpenTelemetry metrics, Bedrock / Vertex AI / Foundry / LLM-gateway support (Claude apps gateways, model discovery timeout config), sandbox mode on Linux/macOS, managed settings (`requiredMinimumVersion`/`requiredMaximumVersion`, `allowManagedMcpServersOnly`, `deniedMcpServers`) **[official — CHANGELOG; secondary — SitePoint]**.
- **Models:** deep integration with the Claude Opus 4.x / Sonnet 4.x family; 1M-token context window in beta on newer Claude models (reported for Opus 4.6/Sonnet 4.6) **[secondary]**. Mid-2026 references cite Claude Opus 4.5 as flagship and Opus 5 as the current flagship API model **[independent/secondary]**.

### 1.4 Pricing (snapshot — re-check https://claude.com/pricing before budgeting)

| Plan | Price | Claude Code included? |
|---|---|---|
| Free | $0 | No |
| Pro | $20/mo ($17/mo annual) | Yes — base usage tier |
| Max 5x | $100/mo | Yes — 5× Pro usage/session |
| Max 20x | $200/mo | Yes — 20× Pro usage/session |
| Team standard | $25/seat/mo ($20 annual, 2-seat min) | Yes — ~1.25× Pro |
| Team premium | $125/seat/mo ($100 annual) | Yes — ~6.25× Pro, plus Claude Cowork |
| Enterprise | $20/seat + usage at API rates | Yes |
| API (Console) | Pay per token | Works via API key |

API rates (per MTok, Aug–Sep 2026): Claude Opus 5 $5 in / $25 out / $0.50 cache-read; Claude Sonnet 5 $3/$15/$0.30 (raised from $2/$10 on Sep 1, 2026); Claude Haiku 4.5 $1/$5/$0.10; Claude Fable 5 $10/$50/$1.00 **[secondary — morphllm.com, verified Sep 19, 2026]**.
**Notable event:** ~April 21, 2026 Anthropic briefly tested removing Claude Code from the Pro plan for ~2% of new signups, then restored it; the reason cited was flat-rate plan economics breaking under long-running agentic sessions **[secondary — devtoolpicks.com, Sep 21, 2026]**.
**Policy:** since April 2026, Anthropic explicitly blocks subscription tokens from being used by third-party CLI tools and runners — use a real API key for anything outside the official Claude Code client **[secondary]**.

### 1.5 Availability & adoption
- Available wherever Claude subscriptions/API are sold; GitHub repo `anthropics/claude-code` (public, 52 listed contributors; internal team larger) **[secondary]**.
- Enterprise deal reported: Allianz making Claude Code available to all employees with custom agents (one of Anthropic's first major enterprise deals of 2026) **[independent — Medium/Cogni Down Under citing Jan 2026 timing — treat as secondary]**.
- License: not stated in collected sources — **not confirmed as open source; assume proprietary**.

