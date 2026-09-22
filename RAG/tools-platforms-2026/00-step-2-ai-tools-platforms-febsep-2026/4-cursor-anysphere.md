---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/4-cursor-anysphere
title: "4. Cursor (Anysphere)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Hugging Face", "OpenAI", "SpaceX", "xAI"]
dates: ["2025-06", "2026-03", "2026-04", "2026-04-02", "2026-07-10", "2026-09-18"]
keywords: ["acquisition", "agent", "agents", "claude", "cost", "funding", "grok", "grok 4", "license", "pricing", "research", "throughput"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [160, 211]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: a7bd656c9574e9eba4a83bacca25494606809d9eced6b5a50a5b336db6fa6467
---

# 4. Cursor (Anysphere)

## 4. Cursor (Anysphere)

### 4.1 Overview
AI-native code editor (VS Code fork) by **Anysphere**. In 2026 it pivoted to an agent-first product **[independent — The Decoder; blazetrends]**.

### 4.2 Latest versions
- **3.21.13** — released September 18, 2026 (per community download trackers; **secondary — unverified vs. official changelog**).
- 3.20.21 (Sep 13, 2026); 3.18.25 (Sep 1, 2026); 3.18.9 (Aug 28, 2026) **[secondary — cursor-download trackers]**.
- **Cursor 3.0** — launched April 2, 2026: interface rebuilt from scratch around agents; new "Agents Window" (`Cmd+Shift+P → Agents Window`); parallel agents across repo boundaries; unified sidebar for local and cloud agents; drag cloud↔local session handoff; integrated browser for agents; new diff view with staging/commit/PR management; plugin marketplace (MCPs, skills, subagents); traditional IDE layout still available as an option **[independent — The Decoder, April 2026; blazetrends]**.

### 4.3 Features (2026)
- **Composer 2** (March 2026): Cursor's own RL-trained coding model; frontier-level performance at ~4× generation speed; most interactive turns under 30 seconds **[secondary — rapidevelopers.com citing Cursor]**.
- **Automations:** always-on agents triggered by schedules or Slack/Linear/GitHub/PagerDuty/webhooks; agents spin up cloud sandboxes with configured MCPs/models **[secondary — rapidevelopers.com]**.
- **Agent Skills** (2.4+): structured knowledge/workflows via skill files **[secondary — igmguru]**.
- **JetBrains IDE support** via the Agent Client Protocol (IntelliJ, PyCharm, WebStorm) **[secondary — rapidevelopers.com]**.
- 30+ marketplace plugins (Atlassian, Datadog, GitLab, Glean, Hugging Face, monday.com, PlanetScale, …) **[secondary — rapidevelopers.com]**.
- **Cursor CLI** with its own changelog (Aug 11, 2026): steer running turns via Enter, subagent transcripts, Explore subagent model choice, sticky skills/custom modes, durable goals (`/goal`) **[official-via-mirror — CLI changelog mirror]**.
- Rules migration: Cursor 3.11 (July 10, 2026) splits `.cursorrules`/`.mdc` rules into AGENTS.md + skills + subagents + plugins; nested `AGENTS.md` supported **[secondary — start-debugging blog]**.
- Bugbot: AI code-review add-on for GitHub PRs (separate pricing) **[secondary — aivexify]**.

### 4.4 Pricing (snapshot — note conflicting reports)

| Plan | Price | Reported allowance |
|---|---|---|
| Hobby (free) | $0 | ~2,000 completions, 50 slow premium requests; 7-day Pro trial |
| Pro | $20/mo ($16/mo annual) | Credit pool: older docs say 500 fast requests/mo; 2026 credit-system docs say ~$20 credit pool (≈225 fast requests), unlimited slow, unlimited completions |
| Pro+ | $60/mo | 3× Pro credits |
| Ultra | $200/mo | 20× Pro credits, max throughput |
| Teams | $40/user/mo ($32 annual) | Pro credits/seat + admin controls; premium seats $120 reported |
| Enterprise | Custom | SSO, audit logs, advanced admin |
| Bugbot add-on | $40/user/mo | Unlimited reviews on up to 200 PRs |

- **Credit system:** since June 2025, Cursor bills dollar-denominated credits; one standard request ≈ $0.04 baseline, but models consume credits at different rates (Claude Opus burns faster than GPT-5.x Mini); overage charges apply beyond the pool in Auto mode **[secondary — aitooldiscovery.com, Aug 2026]**.
- **⚠ Discrepancy:** older guides (incl. 2026 H1) quote "500 fast requests/month" for Pro; newer credit-based guides quote a $20 credit pool (~225 fast requests). Present both as a snapshot conflict **[secondary]**.

### 4.5 Adoption & company
- Widely adopted flagship AI editor; r/cursor community 182k+ members **[secondary — aitooldiscovery.com]**.
- Competitive pressure: community reporting that OpenCode is displacing Cursor among cost-sensitive developers **[secondary — synapse-news]**; funding/valuation figures not confirmed in this research — **not reported here**.
- **Unverified claim:** an AI release-tracker summary of xAI's Grok 4.7 launch states SpaceXAI "bought [the] coding company [Cursor] earlier in the year" and that Grok 4.7 received supplemental training on anonymized Cursor workflow data — **no official acquisition announcement located; must be verified** (see §13).

### 4.6 License & availability
- **Proprietary** (closed-source client); available for Windows, macOS, Linux **[secondary]**.

### 4.7 Key sources
- https://the-decoder.com/new-cursor-3-ditches-the-classic-ide-layout-for-an-agent-first-interface-built-around-parallel-ai-fleets/
- https://blazetrends.com/cursor-3-0-ditches-vs-code-roots-for-full-ai-agent-dashboard/
- https://www.aitooldiscovery.com/guides/cursor-ai-pricing
- https://aivexify.com/cursor-ai-pricing-in-2026/
- https://github.com/oslook/cursor-ai-downloads (version tracker, secondary)

---

