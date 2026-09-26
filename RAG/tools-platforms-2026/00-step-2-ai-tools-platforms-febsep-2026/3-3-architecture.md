---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/3-3-architecture
title: "3.3 Architecture"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Apple", "Google", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "cost", "inference", "license", "memory", "pricing"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [133, 159]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: cfc98c2bc3192223e36b5289e4f233c1257600936199af42fcf23768f284ac58
---

# 3.3 Architecture

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
