---
id: ai-industry-kb-2026/13-agents-mcp/the-78-retraction-figures-governance
title: "The 78% retraction (figures governance)"
domain: agents-mcp
role: deep-dive
task: regulation
actors: ["Anthropic", "Google", "Moonshot", "OpenAI", "SpaceX", "StepFun", "Z.ai", "xAI"]
dates: ["2026-03-25", "2026-04", "2026-05-24", "2026-08-13", "2026-09", "2026-09-11", "2026-09-20", "2026-09-21"]
keywords: ["governance", "agent", "agentic", "agents", "astra", "benchmark", "claude", "cost", "fable 5", "glm", "gpt-5.6", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6637, 6680]
section: "13. Agents & MCP"
sha256: daedf7b2822fd96ec8b473afdf9c225ceeaa4f252bd552b6edc51e8d0834eb2e
---

# The 78% retraction (figures governance)

- **Guardrail pattern (2026 consensus)**: sandboxed shell by default (Codex CLI: bubblewrap/Seatbelt/sandbox; Claude Code: permission prompts; OpenHands: Docker; Jules: ephemeral GCP VMs); plan-review gates; approval policies; MCP allowlists. CrewAI's analysis of 1.7B agentic workflows: the winning pattern is a "deterministic backbone with intelligence deployed where it matters."
- **Vendor long-horizon claims** (all under human steering): OpenAI "Harness Engineering" (Feb 2026) — 3 engineers steered Codex through ~1,500 PRs to ship a million lines over 5 months; Cursor — hundreds of concurrent agents ran for weeks producing a million-line browser; Anthropic (Mar 2026) — Claude compiled the Linux kernel across ~2,000 sessions, multi-day scientific workflows; Spotify Honk — 1,500+ PRs, ~50% of Spotify updates via agents; Atlassian HULA (12,000 engineers): plans for 79% of work items, 82% approved, 59% of HULA PRs merged — but 54% of engineers said code had defects without human review, 67% said it didn't solve the task without intervention.
- **Empirical caution**: "Why Do Multi-Agent LLM Systems Fail?" (Cemri et al., ICLR 2025) — ~79% of 14 failure modes from specification/coordination; LangChain survey — evaluation/observability the lowest-rated stack parts; Google DORA — higher AI adoption associates with **−7.2% delivery stability**.

### The 78% retraction (figures governance)

- Wave 2 §7.1 carried "78% of enterprise AI teams with at least one MCP-backed agent in production (April 2026 survey)". **Verdict: no traceable source; retracted by at least one publication that carried it** (LLM Book 2026 edition: "Do not use it"). This is a retraction, not a nuance — the figure must not appear in the consolidated document.
- **Replacement**: Stacklok (2026-01, n=100 senior technical leaders, software cohort): **41% in production (29% limited + 12% broad), 30% pilot, 29% planning/evaluating**. This is the best-sourced production-adoption figure.
- Additional corroborating enterprise figures (secondary, methodology mixed): 28% of Fortune 500 have implemented MCP servers; 67% of CTOs expect MCP as their default integration standard within a year; Gartner forecasts 75% of API gateway vendors with MCP capabilities and 40% of enterprise apps with task-specific agents by end of 2026 [SECONDARY analyst projections, not outcomes].
- FastMCP (Python, Prefect) claims ~1M daily downloads and ~70% of MCP servers [VENDOR claim].

### September 2026 agent-model competition

- **Cognition SWE-2 (2026-09-11)** [VENDOR]: post-trained from Moonshot AI's **Kimi K3 (2.8T base)** via RL across three effort levels (medium/high/max) — claimed trained in one run. **FrontierCode 1.1 Main 50.0% vs Fable 5.1 50.9%** (benchmark: would a human maintainer merge the AI-written PR), at up to **70% lower running cost**; SWE-2 medium beat predecessor SWE-1.7 with 58% fewer turns and 81% lower cost. On harder agentic ground: **Terminal-Bench 4.0 27.3%** vs Fable 5.1 55.8% / GPT-6 Astra 57.9%. No standalone weights or public per-token price sheet published.
- **StepFun Step 5 (600B, ~2026-09-20)** [SECONDARY, aggregator-sourced]: 1M-token context; Internal StepCodeBench (553 repos, 33 languages): 49.0 avg-at-four; DeepSWE v1.1 67.7% (vs Kimi K3 67.5, GLM-5.3 66.9, GPT-6 Astra 74.1, Opus 5 74.0), ProgramBench 80.5, TB4 33.3, ALE-CLI 29.5.
- **Grok 4.7 (2026-09-21)** [VENDOR/company-reported]: **DeepSWE v1.1 71.0%** (xhigh config — edging Fable 5.1 max 70.0%, trailing GPT-5.6 Sol max 72.7%); CursorBench 4.0 46.3%; HealthBench Professional 56.7%; Harvey Legal Agent Benchmark 19.6%; TB4.0 37.58% (xAI's Grok Build harness; company announcements report 38.0% vs 20.3% for 4.6 — keep both with provenance tags; the gap is harness/rounding noise). Pricing: **$2/$6 per million input/output tokens**, fast variant at double price and double speed. Naming note: several September reports call the company "SpaceXAI" — press-level naming drift; record as reported, not as corporate confirmation.

## Figures and metrics

### MCP scale figures (canonical series)

| Metric | Value | Date | Provenance |
|---|---|---|---|
| Combined SDK monthly downloads at launch | ~2M | 2024-11 | Anthropic ecosystem reporting (canonical) |
| Combined SDK monthly downloads | **97M** | 2026-03-25 | VERIFIED (secondary report citing SDK registry data; corroborated) |
| Combined SDK monthly downloads | **~110M** | 2026-06 | VERIFIED (Dev Summit India reporting) |
| Official registry, latest server records | **9,652** | 2026-05-24 | VERIFIED (registry survey) |
| Official registry, server-version records | 28,959 | 2026-05-24 | VERIFIED (registry survey) |
| Servers across 33 registries | **12,000+** | 2026-04 | VERIFIED (field report) |
| Independent census estimate | ~17,468 | 2026-03 (Q1) | VERIFIED |
| Private/internal enterprise servers | 3–4× public counts | 2026 | [DIRECTIONAL/estimation] |
| AAIF membership | 247 orgs, 8 platinum sponsors | 2026-08-13 | VERIFIED (per wave 2) |

Growth rate: ~4,750% over 16 months (Nov 2024 → Mar 2026). **Do not use**: ~500M/month (single-aggregator, methodology inconsistent with the registry-anchored 97M→110M series); the 78% production figure (retracted); one-migration-article claim of 1B cumulative per TS/Python SDK (uncorroborated).

### Stacklok enterprise adoption (VERIFIED verbatim, the canonical replacement)

- 41% in some form of production = **29% limited production + 12% broad production**
- 30% pilot
- 29% planning or evaluating
- Methodology: n=100 senior technical leaders, software cohort; report 2026-01; parallel reports for Retail and Financial Services sectors. Software-industry cohort sub-slice: 26% planning, 30% pilot, 26% limited, 19% broad (= 45% production).

### Benchmark scoreboard (selected, provenance-tagged)

