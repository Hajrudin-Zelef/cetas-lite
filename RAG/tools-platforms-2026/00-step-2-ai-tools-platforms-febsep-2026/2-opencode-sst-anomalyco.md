---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/2-opencode-sst-anomalyco
title: "2. OpenCode (SST / anomalyco)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Anthropic", "DeepSeek", "EU", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "United States", "Z.ai", "xAI"]
dates: ["2026-01", "2026-03", "2026-06", "2026-06-24", "2026-07", "2026-07-27", "2026-08-14", "2026-08-30", "2026-09", "2026-09-05", "2026-09-19"]
keywords: ["agent", "agents", "chatgpt", "claude", "consumer", "copilot", "cost", "deepseek", "glm", "grok", "kimi", "license"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [79, 132]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: c323d3775d3863b88eef6ede6aa8ad0525ae85aabb992ab472488b9adc183241
---

# 2. OpenCode (SST / anomalyco)

## 2. OpenCode (SST / anomalyco)

### 2.1 Overview
Terminal-first, model-agnostic, **open-source (MIT)** AI coding agent — one of the most-starred developer tools of 2026 **[independent/secondary — gcn.com, nerdleveltech.com]**.

### 2.2 Latest releases
- **v1.18.29** — current release tag as of September 5, 2026 **[secondary — GitHub commit pin reference]**.
- v1.18.18 in-tree on August 14, 2026; v1.18.8 stable referenced ~July 2026; v1.17.10 dated June 24, 2026 **[secondary — kaidera-ai research, datacamp.com]**.
- A **V2 beta** dual-track is mentioned in community config docs (September 2026) **[secondary — unverified]**.
- **Repository moved from `sst/opencode` to `anomalyco/opencode`** (docs and package fields updated) **[official-via-mirror — research doc citing repo state Aug 14, 2026]**.

### 2.3 Architecture
- **Headless server + TUI client**, shipped as a single binary / npm package (`opencode`); terminal-first design with desktop app in beta **[official-via-mirror]**.
- **OpenTUI**: native Zig terminal-UI core with TypeScript bindings and a SolidJS reconciler (`anomalyco/opentui`, ~13k stars, MIT) **[official-via-mirror]**.
- **Model-agnostic:** 75+ LLM providers via the models.dev registry; per-project config in `opencode.json` (or `~/.opencode/config.json`); `/model` switches models mid-session; supports Ollama/local models; in January 2026, after Anthropic blocked third-party tools from using consumer Claude subscriptions through unofficial channels, OpenCode added other subscription options and its own gateway **[secondary — datacamp.com]**.
- **Agents:** build / plan / general agent types; **parallel agent execution** (June 2026) — multiple agents working on different parts of the same codebase simultaneously **[secondary — synapse-news]**.
- **Project memory:** plain-file `AGENTS.md` (OpenCode's equivalent of CLAUDE.md), generated via `/init` **[secondary — devtoollab.com]**.
- **Compiler-aware:** integrates the Language Server Protocol so the agent is grounded in live diagnostics — cited as a differentiator vs. paid rivals **[independent — gcn.com]**.
- **LSP + local execution** make it the default pick for teams that cannot/will not ship code to a vendor cloud **[independent — gcn.com]**.

### 2.4 Pricing & the Zen / Go gateways
- **The tool itself is free (MIT).** Spend is model usage: pay-as-you-go API tokens, or existing subscriptions (Copilot, ChatGPT Plus/Pro) OpenCode can authenticate against **[secondary — nerdleveltech.com]**.
- **OpenCode Zen** — the team's curated, tested model gateway at `https://opencode.ai/zen` (`/v1/chat/completions`, `/v1/responses` endpoints), OpenAI-compatible, single API key (`OPENCODE_API_KEY`). Pay-as-you-go credits (auto-reload); sold at cost with no markup, plus processing fees; includes rotating **free models** (e.g. `deepseek-v4-flash-free`, `mimo-v2.5-free`, `qwen3.6-plus-free`, `minimax-m3-free`, `nemotron-3-ultra-free`, `north-mini-code-free`, stealth `big-pickle`). Paid models include DeepSeek V4, GLM-5.x, Kimi K2.5/K2.6, MiniMax M2.x, Grok Build. Teams/spend caps/BYO OpenAI/Anthropic keys supported **[official-via-mirror — docs.docker.com; secondary — dysonharness research, July 27, 2026]**.
- **OpenCode Go** — fixed-price subscription: **$5 first month, then $10/mo**, with dollar caps (~$12/5h, $30/week, $60/month); **open-weight coding models only**; hosted in US, EU, Singapore; optional "use balance" falls back to Zen credits **[secondary — dysonharness, July 27, 2026]**.
- Free-tier note: free models are promotional and can disappear; free/stealth models may have data used for training **[secondary — itsfree.ai]**.

### 2.5 Adoption
- GitHub stars: **~197k (Aug 14) → 201.4k (Aug 26) → 208.2k (early Sep)** on star-ranking trackers **[secondary]**; ~25k forks, ~900 contributors, ~5.2k open issues **[secondary — GitHub API via kaidera-ai research]**.
- Claimed **7.5M+ monthly developers** **[vendor-reported/secondary — gcn.com, nerdleveltech.com]** — **unverified; treat as marketing**.
- Android companion app ("OpenCode: AI Coding Agent") by third party VIBE TECHNOLOGIES, LLC (~7.6k downloads, 4.67★, updated Aug 14, 2026) **[secondary — AppBrain]**.
- Community forks exist (e.g. "Gorilla OpenCode" fork emphasizing token efficiency) **[secondary — GitHub]**.

### 2.6 License & availability
- **MIT**, TypeScript; install via `curl -fsSL https://opencode.ai/install | bash` or Homebrew (`opencode-ai/tap/opencode`); npm; desktop app beta at opencode.ai/download **[official-via-mirror]**.

### 2.7 Key sources
- https://github.com/anomalyco/opencode (official repo)
- https://opencode.ai/docs/zen/ (official docs)
- https://docs.docker.com/ai/docker-agent/providers/opencode-zen/
- https://gcn.com/7-5-over-million-developers-now/20116/
- https://www.datacamp.com/blog/what-is-opencode
- http://nerdleveltech.com/opencode-open-source-ai-coding-agent-explained

---

## 3. OpenClaw (OpenClaw Foundation — verify)

### 3.1 Overview
**Open-source (MIT), self-hosted personal AI agent** — a "shared agent control plane" rather than a coding-only tool. Formerly known as **Clawdbot / Moltbot** **[secondary — milvus-io community blog]**. Tagline: "Your own personal AI assistant. Any OS. Any platform." One architecture doc attributes the project to Peter Steinberger and notes a move to an open-source foundation **[secondary — skyclaw docs, March 2026; unverified]**.

### 3.2 Latest releases
- **2026.9.5** — npm `latest` tag, September 19, 2026: **atomic updates** (Gateway keeps running while preparing the next version; rolls back to last working config on failure), plugin hot reload, conversation sharing (read-only access for teammates), expanded GPT Live, shared browser pages, conversation archiving, specialist-agent setup. 4,179 PRs, 64 direct commits, 502–503 contributing accounts **[independent — MarkTechPost, 2026-09-19; AgentRiot]**.
- **2026.8.1 ("OpenClaw 2.0")** — August 30, 2026: major security overhaul — masked credential prompts (secret values never enter chat transcript or model context), explicit per-session permission modes, file access tied to recorded workspace/worktree, team roles, shared credential store, protected network access limiting secret use to approved hosts, clearer approval screens; plus conversation search, paired devices + cloud workers, durable session progress, interactive dashboards, rebuilt browser Control UI, shared multiplayer sessions. Built by 933 contributors across 16,000+ PRs **[independent — CyberPress; Analytics Insight]**.

