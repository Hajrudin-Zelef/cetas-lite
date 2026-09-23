---
id: vague2-datacamp/datacamp/claude-cowork-vs-openclaw
title: "Claude Cowork vs OpenClaw : persistance, permissions et coût réel"
domain: datacamp
role: reference
task: article
actors: ["AWS", "Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Mistral", "OpenAI", "OpenRouter", "SGLang", "vLLM", "xAI"]
dates: ["2026-09-23"]
keywords: ["claude", "agent", "agentic", "agents", "bedrock", "chatgpt", "cohere", "cost", "deepseek", "funding", "gemini", "grok"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/claude-cowork-vs-openclaw.md
source_anchor: ""
source_lines: [1, 58]
sha256: 4a2a57136110b5aa46cf35611ad62ac5cd4b547ff0393a81acf48ebedb8d97fc
---

# Claude Cowork vs OpenClaw : persistance, permissions et coût réel

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/claude-cowork-vs-openclaw
- **Site** : DataCamp
- **Type** : Article (comparatif)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp comparison contrasts two desktop agents: **Claude Cowork** (Anthropic's agentic mode inside the Claude desktop app) and **OpenClaw** (a self-hosted, MIT-licensed gateway connecting chat apps to AI agents, developed by the nonprofit OpenClaw Foundation). They compete but are not equivalent: Cowork sells a permission-controlled agent with subsidized usage, while OpenClaw hands you a Node process and full responsibility.

**Persistence is the real differentiator.** Cowork's "close your laptop, it keeps working" promise applies to the cloud path (web and mobile in beta), not the desktop app that accesses local folders. DataCamp's hands-on Cowork Dispatch tests showed desktop sessions stop when the Mac sleeps. OpenClaw's Gateway installs as a service and runs continuously with cron and webhooks; a 7 a.m. task runs whether you're at your desk or not.

**Security is inverted, and only one has a documented incident.** Cowork defaults to caution: access limited to explicitly selected folders/tools, deletion always requires approval, plan confirmation before significant actions, per-app permission for computer use. OpenClaw defaults to extended system access (full shell/files/browser). The **ClawHavoc incident** turned access to the skills marketplace into credential theft; anything installed from ClawHub inherits the agent's scope. But Cowork isn't flawless: its controls are configurable, not automatic, and Anthropic's pricing page notes Cowork activity isn't yet captured in audit logs or the Compliance API.

**Cost.** Cowork is flat-rate: Pro $17/mo annual ($20 monthly), Max 5x $100, Max 20x $200, Team $20/seat (2–150 people), Enterprise custom. Cowork consumes usage limits faster than Chat because it coordinates subagents and tool calls. OpenClaw software is free but costs tokens (~$3/M input, $15/M output on Claude Sonnet) plus hosting (managed hosting from ~$9.99/mo). A key actionable point: if you already pay for Claude Pro/Max, you can generate a config token via the Claude Code CLI and run OpenClaw on that subscription instead of pay-per-token.

**Model choice.** Cowork is Anthropic-only (Fable, Opus, Sonnet, Haiku). OpenClaw is provider-agnostic: Claude, Gemini, OpenAI/ChatGPT/Codex, Grok, Mistral, DeepSeek, Cohere, Qwen, plus OpenRouter, LiteLLM, Amazon Bedrock, Vertex AI, and local execution via Ollama, LM Studio, vLLM, SGLang.

**Install friction.** Cowork: download, sign in, point at a folder, QR pairing for mobile. OpenClaw: Node 26 recommended (22.22.3+/24.15+/25.9+), API key, config at `~/.openclaw/openclaw.json`, WSL2 on Windows, control UI at `http://127.0.0.1:18789/`.

**Channels & admin.** Cowork lives in the Claude app (web/mobile beta). OpenClaw reaches you in Discord, Google Chat, iMessage, Matrix, Teams, Signal, Slack, Telegram, WhatsApp, Zalo, WebChat, plus local control UI. Cowork wins enterprise admin: org-wide disable, RBAC by team, spend caps, per-department tool permissions, OpenTelemetry to your SIEM.

**Together:** the author uses Cowork for interactive sensitive local documents (guardrails worth the lock-in) and OpenClaw for scheduled overnight local tasks, funding both with one Claude subscription.

## Key points

- Persistence is the true differentiator: OpenClaw runs 24/7 as a service; Cowork's unsupervised path is cloud beta.
- Security models are inverted: Cowork restricts by default; OpenClaw grants full system access by default.
- ClawHavoc exploited OpenClaw's skills marketplace for credential theft; Cowork's controls are configurable, not automatic.
- Cowork pricing: Pro $17/mo annual, Max $100/$200, Team $20/seat.
- OpenClaw is free (MIT) but pays per token plus hosting; can run on a Claude subscription token.
- Cowork is Anthropic-only; OpenClaw supports many providers and local models.
- Cowork offers enterprise admin (RBAC, spend caps, OpenTelemetry); OpenClaw offers none by default.
- Combining both is recommended since their weaknesses are complementary.

## Technical data / figures

| Dimension | Claude Cowork | OpenClaw |
|---|---|---|
| Unsupervised run | Cloud scheduling; web/mobile beta | Always-on daemon with cron/webhooks |
| Default permissions | Folder-scoped, approval, deletion requires validation | Full system access |
| Install | Download app, sign in, pick folders, QR pairing | Node 26 (22.22.3+/24.15+/25.9+), JSON config, WSL2 on Windows |
| Cost | $17/mo Pro annual ($20 monthly), Max $100/$200, Team $20/seat | Free software; ~$3/M input + $15/M output (Sonnet) + hosting (~$9.99/mo managed) |
| Models | Anthropic only | Claude, Gemini, GPT/Codex, Grok, Mistral, DeepSeek, Cohere, Qwen + local |
| Interaction | Claude desktop app; web/mobile beta | 10+ chat channels + local control UI |
| Enterprise admin | RBAC, spend caps, per-department permissions, OpenTelemetry | None standard |
| License | Proprietary subscription | MIT |

Key terms: `~/.openclaw/openclaw.json`, control UI `http://127.0.0.1:18789/`, ClawHub skills registry, ClawHavoc incident.

## Why this source matters for the RAG

It provides a detailed, security-conscious comparison of two prominent desktop agent approaches, with concrete pricing, permission defaults, and a documented security incident. It is valuable for questions on agent persistence, sandboxing/permissions, self-hosting, and Claude Cowork alternatives.
