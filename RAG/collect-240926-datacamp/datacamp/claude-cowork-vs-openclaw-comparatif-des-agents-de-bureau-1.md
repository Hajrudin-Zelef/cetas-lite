---
id: collect-240926-datacamp/datacamp/claude-cowork-vs-openclaw-comparatif-des-agents-de-bureau-1
title: "Curriculum"
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Microsoft", "Mistral", "SGLang", "vLLM", "xAI"]
dates: ["2026-02-24", "2026-04-09"]
keywords: ["agent", "agentic", "agents", "claude", "cohere", "cost", "deepseek", "gemini", "grok", "guardrails", "incident", "license"]
source: docs/RAG/clean_en/datacamp/claude-cowork-vs-openclaw-comparatif-des-agents-de-bureau.md
source_anchor: ""
source_lines: [1, 69]
sha256: d1eafe6be423fb4f757e5a23df7f179625c9e9ba671f7744a7b08b63d4b023ae
---

# Curriculum

<!-- source: https://www.datacamp.com/fr/blog/claude-cowork-vs-openclaw -->

# Curriculum

You have a stack of interviews to synthesize, a Monday report that nobody wants to put together, and a computer you'd rather shut down at 6 p.m. Two tools promise to take that work off your hands: Claude Cowork, Anthropic's agentic mode in the Claude desktop app, and OpenClaw, a self-hosted gateway that connects your messaging apps to AI agents.

They compete, but they're not playing in the same league. Cowork sells you an agent with controlled permissions and subsidized usage, ready to download and connect. OpenClaw hands you a Node process, an MIT license, and full responsibility for whatever happens next.

In this article, I compare Claude Cowork and OpenClaw on unattended persistence, default security settings, real cost at three usage tiers, setup friction, model lock-in, and enterprise administration. To go deeper on each tool, see our Claude Cowork tutorial, our OpenClaw guide, and our roundup of Claude Cowork alternatives.

## In brief

- Persistence is the real difference: OpenClaw runs as a 24/7 service with cron, while Cowork's "close your laptop" promise applies to web and mobile in beta, not to the desktop app that accesses your local folders.
- The security models are inverted, and only one has a documented incident: OpenClaw grants broad system access by default, and its skills marketplace was exploited in the ClawHavoc incident.
- Choose Claude Cowork if you want permission guardrails, admin controls, and a $17/month bill you can forget about.
- Choose OpenClaw if you need unattended scheduled tasks, local models via Ollama, or control over where your data lives.

## What is Claude Cowork?

Claude Cowork is Anthropic's agentic work mode: you give Claude a goal, it operates in the folders and connected tools you choose, and hands you a finished result for review. It runs on macOS, Windows (x64 and arm64), Linux, and ChromeOS, with web and mobile in beta. In the Claude desktop app, a toggle below the input window lets you switch between Cowork and the usual chat mode.

The principle Anthropic highlights on Cowork's product page is: "say what, not how." Claude prioritizes connectors first, falls back to your browser if needed, and only takes control of your screen as a last resort (computer use remains in research preview). It also breaks large jobs into segments run in parallel, for simultaneous rather than sequential drafting and research.

Customization comes through plugins, which bundle skills (domain knowledge), connectors (Amplitude, Microsoft 365, Google Drive, Slack, and others), and sub-agents into a single install. Anthropic shipped private plugin marketplaces for admins on February 24, 2026, and enterprise deployment controls on April 9, 2026. We detail file organization, batch conversions, and Chrome automation in our hands-on Cowork tutorial.

## What is OpenClaw?

OpenClaw is a self-hosted gateway that connects chat apps to AI agents, developed openly by the OpenClaw Foundation, a nonprofit organization, under the MIT license. You run a Gateway process on your machine or a server, which becomes the bridge between an always-available agent and a large number of messaging services: Discord, Google Chat, iMessage, Matrix, Microsoft Teams, Signal, Slack, Telegram, WhatsApp, Zalo, and WebChat are supported. The Gateway is the source of truth for sessions, routing, and channel connections.

Configuration lives in a single file at `~/.openclaw/openclaw.json`, and a browser-based control interface runs locally at `http://127.0.0.1:18789/`. The documentation is clear about the target audience: developers and power users who want a personal AI assistant reachable from anywhere without entrusting their data to a hosted service. TechRadar counts more than 50 official integrations, plus ClawHub, the community skills registry.

To get it running, our OpenClaw installation tutorial covers setup and channel pairing, and our OpenClaw with Ollama tutorial covers the fully local path.

## Claude Cowork vs OpenClaw: a point-by-point comparison

Here's the summary version before we get into the criteria that really decide things. Note how few lines are about "power."

| Dimension | Claude Cowork | OpenClaw | 
|---|---|---|
| Unsupervised execution | Scheduled tasks on the cloud path; web and mobile in beta. Desktop sessions depend on your machine's sleep state | Always-on daemon with cron and webhooks | 
| Default permissions | Scope limited to folder, approval before significant actions, and deletion subject to validation | Full system access with extended permissions by default | 
| Installation | Download the desktop app, sign in, choose folders; QR pairing for mobile | Node 26 recommended (22.22.3+, 24.15+ or 25.9+), JSON config, WSL2 on Windows | 
| Cost | $17/month Pro annual ($20 monthly), $100 Max 5x, $200 Max 20x, $20/seat Team | Free software; about $3 per million input tokens and $15 per million output tokens on Claude Sonnet, plus hosting | 
| Model choice | Anthropic models only | Claude, Gemini, GPT/Codex, Grok, Mistral, DeepSeek, Cohere, Qwen, plus Ollama, LM Studio, vLLM, SGLang | 
| Where to interact with the agent | Claude desktop app; web and mobile in beta | Any paired messaging channel, plus the local control interface | 
| Enterprise administration | Per-team RBAC, spending caps, per-department permissions, OpenTelemetry to your SIEM | Nothing out of the box; up to you to harden the environment | 
| License | Proprietary, subscription only | MIT, community-driven | 

### Persistence and unsupervised scheduling

Cowork's product page promises: "Close your laptop, it keeps going" and highlights tasks that can be scheduled at any cadence, without supervision. This promise concerns the cloud path, and the page itself states that web and mobile are in beta.

The desktop app is the surface that accesses your local folders and applications, and Anthropic says so in its FAQ: the desktop app adds what web and mobile cannot reach. Our hands-on testing of Cowork Dispatch showed that the desktop experience is tied to the session and stops as soon as the Mac goes to sleep. One of the tips we give is to go into System Settings and change the sleep timeout.

OpenClaw doesn't have that asterisk. The Gateway installs as a service and runs in the background, with cron and webhooks in its toolbox; a 7 a.m. task starts whether you're at your desk or not. If your machine goes to sleep, that's your power setting, not a product limitation.

Frame this honestly: if your work lives in local folders and you want it to run overnight, OpenClaw does it today, and Cowork does it in beta on the cloud surface, which cannot touch your local file system. Anthropic will close this gap, and the beta badge suggests it's underway. For now, the "fire and forget" promise is Cowork's weakest point.

### Security, permissions and accountability

Cowork sits on the cautious side by default:

- Access limited to explicitly selected folders and tools
- Deletion always requires your approval
- Can be configured to show its plan and wait for validation before significant actions
- Per-app permission when computer use is enabled, with the ability to stop at each step

OpenClaw, by contrast, runs by default with extended permissions. Our own comparison of Cowork Dispatch and OpenClaw clearly describes it as offering full system access, which can become a risk if you don't configure it carefully. The risk isn't theoretical: as explained in our OpenClaw deep dive, the ClawHavoc incident turned access to the skills marketplace into credential theft. Anything you install from ClawHub inherits the agent's scope: shell, files, browser.

