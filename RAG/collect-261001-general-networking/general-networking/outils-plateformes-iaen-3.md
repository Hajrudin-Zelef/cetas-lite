---
id: collect-261001-general-networking/general-networking/outils-plateformes-iaen-3
title: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Hugging Face", "Microsoft", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2025-06", "2026-03", "2026-04", "2026-04-02", "2026-07-10", "2026-09-18"]
keywords: ["acquisition", "agent", "agents", "claude", "cost", "funding", "grok", "grok 4", "inference", "license", "memory", "pricing"]
source: docs/RAG/collect-261001-general-networking/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [133, 202]
sha256: 2ba9ec2cc0452dbb1c4658364ba6482ef308866f9accf690e69a6b8fd9b3b9ff
---

# Step 2 — AI Tools & Platforms (Feb–Sep 2026)

### 3.3 Architecture
- **Hub-and-spoke:** central **Gateway** (Node.js/TypeScript) routes tasks to LLM backends, executes skills (file ops, shell, browser automation, API calls), and returns structured results **[secondary — architecture docs]**.
- **Channels:** one Gateway serves WhatsApp, Telegram, Discord, Slack, iMessage, Signal, Matrix, Microsoft Teams, Google Chat, Zalo, and more via channel plugins (~30 total) **[official-via-mirror — openclaw-android-assistant docs]**.
- **Surfaces:** web Control UI (browser dashboard), CLI, native macOS app, iOS/Android nodes, Apple Watch companion (v2026.2.19) **[independent/secondary]**.
- **Memory/state:** conversations, long-term memory, and skills stored as plain Markdown/YAML under the workspace and `~/.openclaw`; heartbeat daemon (`systemd`/`LaunchAgent`, default every 30 min, hourly with Anthropic OAuth) drives autonomous checks via `HEARTBEAT.md`; cron jobs, webhooks, subagents, cloud workers **[secondary — milvus-io blog]**.
- **Models:** any cloud provider (Anthropic, OpenAI, Google) or local via Ollama/LM Studio/OpenAI-compatible servers — fully local inference possible **[secondary — milvus-io blog]**.
- **ClawHub:** plugin/skill registry; 2.0 shows security audit info before install, requires force flag for arbitrary executable sources; plugin SDK with migration requirements **[independent — Analytics Insight; secondary — 2.0 release notes]**.
- **Security history:** v2026.2.17/2.19 patched a real-world infostealer targeting OpenClaw config files (documented by Hudson Rock), plus credential-theft path OC-09, config traversal, and session permission bugs **[secondary — YouTube breakdown citing Hudson Rock]**.

### 3.4 Pricing, licensing, availability
- **Free; self-hosted only** — no first-party paid hosted service. Only cost is your own model API keys. Requires Node 24.16+ or 26.1+ (Node 26 recommended) **[official-via-mirror; independent — Composio comparison]**.
- **License: MIT** (core Gateway; "OpenClaw Foundation" named as steward in one source) **[secondary — aleksandr-bogdanov/imprnt, citing openclaw.ai, Aug 26, 2026]**.
- Install: one-liner `curl -fsSL https://openclaw.ai/install.sh | bash` or npm; native macOS ZIP/DMG signed+notarized for 2026.8.1; Android via Google Play **[secondary — release notes mirror]**.

### 3.5 Adoption
- **~387.6k–389.5k GitHub stars** (Aug–Sep 2026 trackers) **[secondary]**; Composio (Sep 2026) reports ~389.5k stars vs. Hermes Agent ~244.7k — Hermes leads on OpenRouter token volume instead **[independent — Composio]**.
- Compared by Composio (Sep 2026) as the "shared agent control plane" vs. Hermes Agent's "persistent, self-improving agent runtime"; OpenClaw rated stronger on human collaboration, teams/shared infrastructure; sandboxing off by default (vs. Hermes' smart approvals) **[independent — Composio]**.

### 3.6 Key sources
- https://www.marktechpost.com/2026/09/19/openclaw-releases-2026-9-5/
- https://agentriot.com/news/ai-agents/openclaw-2026-9-5-atomic-updates
- https://cyberpress.org/openclaw-2-0-released-with-enhanced-ai-agent-security/
- https://www.analyticsinsight.net/artificial-intelligence/openclaw-explained-how-its-viral-ai-agent-platform-is-being-built-and-secured
- https://composio.dev/content/openclaw-vs-hermes-agent
- https://github.com/openclaw/openclaw (official repo)

---
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

