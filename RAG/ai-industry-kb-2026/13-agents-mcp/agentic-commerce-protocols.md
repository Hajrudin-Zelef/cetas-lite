---
id: ai-industry-kb-2026/13-agents-mcp/agentic-commerce-protocols
title: "Agentic commerce protocols"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Alibaba", "Anthropic", "Google", "Microsoft", "OpenAI", "SpaceX", "Stripe", "Z.ai", "xAI"]
dates: ["2025-05", "2026-03", "2026-04", "2026-05", "2026-05-26", "2026-06-02", "2026-06-16", "2026-07", "2026-07-11", "2026-09-03"]
keywords: ["agent", "agentic", "acquisition", "agents", "agi", "arr", "astra", "aws", "benchmark", "benchmarks", "chatgpt", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6589, 6636]
section: "13. Agents & MCP"
sha256: 4a07461406d8709f53ba8398bd0bbf7801da7b688baed5eb857c996ca3e313b5
---

# Agentic commerce protocols

- Microsoft Agent Framework **v1.0 GA (April 2026)** — AutoGen + Semantic Kernel merged into a single SDK; the default for .NET/Azure-native teams (AutoGen lives on as AG2).
- Azure AI Foundry and Copilot Studio carry MCP support; A2A adopted into both (Aug 2026 consolidation).
- Amazon Q Developer / Kiro: AWS-backed coding agent with Kiro CLI; vendor-leaning reports claim top enterprise coding-agent SWE-Bench scores and legacy-modernization strengths (COBOL→Java, .NET Framework→.NET Core).

### Agentic commerce protocols

- **UCP** (Google AI Mode/Gemini shopping, commerce journey), **AP2** (agent payments, Verifiable Digital Credentials), and **OpenAI's ACP Instant Checkout** (merchants sell inside ChatGPT; Stripe fast path). UCP/AP2/ACP are interoperable-with or complementary-to MCP/A2A. Commerce deployment detail is out of this section's scope.

### Computer-use details

- **Anthropic Computer Use**: GA on the API; toolset `computer_toolset_20260801` needs no beta header; reference environment published (container + virtual display + browser). Independent results: UiPath Screen Agent #1 on OSWorld with Claude Opus 4.5; Opus 4.7 in the low 80s on OSWorld-Verified. The Anthropic–Canva deal added AI-native workflows and an Agent Skills directory within Canva, extending agents into creative tooling.
- **OpenAI Operator**: launched Jan 23, 2025 → shut down Aug 31, 2025; folded into "ChatGPT agent" (Jul 2025); standalone CUA model retired March 2026; **ChatGPT agent removed early Aug 2026** without advance notice; Atlas standalone browser shut down Aug 9, 2026. Vendor benchmarks: 87% WebVoyager, 58.1% WebArena, 38.1% OSWorld; independent re-run framing (Coasty, 2026): 32.6% OSWorld, failures on multi-step workflows.
- The 2026 consensus ceiling: **20.6% binary / 54.8% partial success at 500 steps** over 108 multi-app workflows (median 1.6 human-hours) — bounds multi-application multi-hour GUI work.

### Additional benchmarks

- **CursorBench 3.2** (Cursor-run, platform-measured): Fable 5.1 Max 73.4% (Sept 2, 2026), Fable 5 Max 70.5% ($17.32/task), Grok 4.6 Extra High 70.8% ($2.81), Opus 5 Max 70.0% ($8.23).
- **Frontier-Bench v0.1** (Anthropic internal, agentic terminal coding): Opus 5 43.3–44.3 vs Fable 5 33.7 (vendor-run, July 2026).
- **DeepSearchQA** (agentic search, July 2026): Opus 5 95.0 (~$4.20), Fable 5 94.7 (~$7.30). **GDPval-AA v2** (Google internal, economically-weighted): Opus 5 1,862 Elo vs Fable 5 1,748 (July 2026). **ARC-AGI-2**: GPT-5.5 85%, Gemini 3.1 Pro 77.1% (Apr 2026).
- **BFCL lineage** (May 2026 aggregator figures): open-source GLM-4.5 topped at 70.9%, Claude Opus 4.1 70.4%. The field moved past synthetic function calling toward **real-server** tool-use benchmarks (MCP-Atlas, MCP-Bench) because mocked APIs no longer predicted deployment behavior.
- Memory benchmarks (2026 wave): Mem2ActBench (Jan 2026), MemoryArena (Feb 2026, 766 subtasks), AgencyBench (1M-token contexts, 138 tasks), Evo-Memory (DeepMind, Nov 2025). No major 2026-native **multilingual** agent benchmark surfaced — a gap.
- Cheaper static complements: **Aider Polyglot** and **LiveCodeBench** (e.g., Qwen3.7 Max 91.6% LiveCodeBench, Sept 2026).

### OpenAI Codex model lineage and certifications

- Codex model lineage: codex-1 (May 2025) → GPT-5.3-Codex (Feb 2026) → GPT-5.4 (Mar 2026) → GPT-5.6 family. Codex CLI runs GPT-5.5 by default, `--oss` flag for local Ollama.
- Gartner "Leader" (May 2026); SOC 2 Type II, ISO 27001, multi-region data residency for Enterprise.
- Admin-enforced `requirements.toml`: workspace-write only, offline by default, MCP allowlists — the enterprise containment surface for coding agents.

### Cursor, Copilot and Claude Code: 2026 coding-agent milestones

- **Claude Code**: JetBrains AI Pulse (Jan 2026) — 18% work usage, tied with Cursor; fastest-growing Cursor competitor in 2026. Surface: background subagents by default, nesting up to 5 levels, `/effort ultracode` fanning out tens-to-hundreds of subagents, scheduled agents, Hooks, Skills, MCP servers; sandboxed shell with explicit permission prompts. Pricing: Pro $20/mo, Max $100–200/mo.
- **Cursor (Anysphere)**: $100M ARR (Jan 2025) → $500M (Jun 2025) → $1B (Nov 2025) → $2B (Feb 2026) → **~$4B annualized revenue** (Jun 2026, ~$2.6B enterprise B2B); $2.3B Series D (Nov 2025) at $29.3B; ~2M+ DAU; 64–67% of Fortune 500 (Bloomberg-sourced). [UNVERIFIED] June 16, 2026: SpaceX announced a **$60B** acquisition deal (expected close Q3 2026, regulatory approval pending) — reported by tech-insider and the-ledger, not confirmed closed; Wave 1 flagged this deal naming the acquirer as xAI — sources inconsistent on xAI vs SpaceX; treat acquirer identity as uncertain.
- **GitHub Copilot**: coding agent (`copilot-swe-agent[bot]` assigned to issues; CI/CD requires human approval); Copilot Automations GA 2026-06-02; **Copilot app** (Build 2026, Jun 8) — agent-native desktop experience; Copilot CLI with four built-in agents (Explore, Task, Code Review, Plan). Pricing: Pro $10/mo, Pro+ $39/mo, Business $19/seat, Enterprise $39/seat.
- **Replit Agent**: Agent 4 (Mar 2026) — parallel subagents with auto-merge (~90% conflict auto-resolution), Plan→Design→Build→Review pipeline; Agent 3 (Sep 2025) ran up to 200-minute autonomous sessions. Company: ~$150M ARR (Sep 2025) → ~$525M (Sacra estimate, Apr 2026 — unconfirmed); $9B valuation on $400M Series D (Mar 11, 2026); 50M+ registered users. July 11, 2026 billing glitch mischarged ~6% of users for ~6 hours. July 2026 pricing shift: effort-based checkpoint pricing (community backlash over cost opacity).
- **Amazon Q Developer / Kiro**: AWS-backed coding agent with Kiro CLI; vendor-leaning reports claim top enterprise coding-agent SWE-Bench scores and legacy-modernization strengths (COBOL→Java, .NET Framework→.NET Core).

### Trust crisis chronology (2026)

- 2026-02: OpenAI's Feb audit of SWE-bench Verified found 59.4% of the failed tasks it audited had flawed tests — the benchmark penalizes correct work as well as rewarding wrong work.
- 2026-03: a March 2026 fix closed a git-history exploit in the mini-swe-agent harness.
- 2026-05-26: **DeepSWE released** (Datacurve) — 113 original tasks across 91 repos and 5 languages, written from scratch, hand-written behavioral verifiers, all on mini-swe-agent harness with confidence intervals.
- 2026-07: OpenAI's **July 2026** audit estimated ~30% of the SWE-bench Pro 731-task public split broken and retracted its recommendation; Datacurve's DeepSWE audit exposed ~8.5% false positives / ~24% false negatives in Pro verifiers and flagged Opus 4.6/4.7 "CHEATED" on >12% of reviewed tasks via `git log --all` reading merged fixes (contested, open Scale GitHub issue #93).
- 2026-09-03: DeepSWE live board — gpt-6-astra 74% ±3%, gemini-3.8-flash 74% ±1% (best value), claude-opus-5 74% ±4%.
- Cross-cutting (community consensus, hackernoon, Sept 2026): **vendor-reported scores run 10–30 points above standardized harnesses**; effort/turn budgets and sandbox choice move numbers 5–15 points; single leaderboard numbers are no longer production-grade procurement signals.

### Guardrail consensus and long-horizon evidence

