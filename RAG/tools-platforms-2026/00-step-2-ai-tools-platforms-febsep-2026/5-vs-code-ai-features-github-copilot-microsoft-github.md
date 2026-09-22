---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/5-vs-code-ai-features-github-copilot-microsoft-github
title: "5. VS Code AI features / GitHub Copilot (Microsoft & GitHub)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Anthropic", "China", "Google", "Microsoft", "OpenAI", "Z.ai"]
dates: ["2026-03", "2026-05", "2026-06", "2026-06-01", "2026-07", "2026-07-02", "2026-07-04", "2026-07-30", "2026-07-31", "2026-08", "2026-08-31"]
keywords: ["copilot", "agent", "agentic", "agents", "claude", "cloud agent", "gemini", "glm", "inference", "pricing", "research", "revenue"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [212, 278]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 4106efd631b0da98e9095ddf420478ccffa310ec000344b0d199feaeae327683
---

# 5. VS Code AI features / GitHub Copilot (Microsoft & GitHub)

## 5. VS Code AI features / GitHub Copilot (Microsoft & GitHub)

### 5.1 Copilot modes in VS Code (2026)
Three distinct modes: **Ask** (conversational, no edits), **Edit** (targeted diffs), **Agent** (autonomous: plans, edits files, runs terminal commands, observes output, iterates) **[secondary — dev.to guide]**.

### 5.2 Recent release highlights (official GitHub Changelog)
- **March 2026 (VS Code 1.100):** faster agent edits, automatic follow-up fixes when new diagnostics appear, apply-patch editing format for GPT-4.1/o4-mini, Anthropic replace-string tool for Claude Sonnet 3.7/3.5, `#githubRepo` code-search tool, smarter prompt caching, conversation summarization **[independent — AIntelligenceHub]**.
- **April–May 2026 (v1.116–v1.119):** semantic indexing in all workspaces; experimental `/chronicle` (query own chat history from a local DB); deferred tool loading; **BYOK extends to Copilot Business/Enterprise**; group policies for agent-reachable domains; Copilot CLI sessions monitorable/steerable remotely from GitHub.com or mobile app **[official — GitHub Changelog via RSS mirror]**.
- **June 2026 (v1.123–v1.127):** **browser tools GA (July 1)** — agents browse/click through live apps, enabled by default for paid Copilot; **parallel agent sessions**; **1M-token context windows** with Anthropic and OpenAI models; session sync across machines; credit-visibility features **[independent — TechTimes]**.
- **July 2026 (v1.127–v1.131):** Agents window (public preview) redesign; multi-chat support for Claude (multiple chats per session, each with own history/title/model); fork into peer chat; **worktrees with any harness** (Copilot, Claude, or Codex sessions in isolated Git worktrees); subagent tracking (model, elapsed time, active tool call); PR updates from chat (failed CI, review comments) **[official — https://github.blog/changelog/2026-07-30-github-copilot-in-visual-studio-code-july-2026-releases/]**.
- **August 2026 (v1.132–v1.135):** side-by-side chats; `/btw` side-conversation sharing primary context/prompt cache; prompt timeline; **Agent Plugins 1.0 standard** (portable agent plugins across VS Code and compatible clients); open Agents window without GitHub sign-in (experimental, Claude API key); switch model providers mid Claude session (Anthropic ↔ Copilot subscription); continue external agent sessions; experimental `/rubber-duck` second-opinion command **[official — https://github.blog/changelog/2026-08-31-github-copilot-in-vs-code-august-2026-releases/]**.

### 5.3 Copilot pricing — GitHub AI Credits (since June 1, 2026)
GitHub replaced premium requests with **GitHub AI Credits (1 credit = $0.01)** billed on actual token usage (input/output/cached) at per-model API rates. Seat prices unchanged **[independent — Second Talent, Sep 11, 2026; Guickly, Aug 2026]**.

| Plan | Price | Included AI credits |
|---|---|---|
| Free | $0 | Limited (2,000 completions) |
| Pro | $10/mo | $15 ($10 base + $5 flex) |
| Pro+ | $39/mo | $70 ($39 base + $31 flex); premium models incl. Opus |
| Max | $100/mo | $200 ($100 base + $100 flex) |
| Business | $19/seat/mo | $19 pooled across org |
| Enterprise | $39/seat/mo | $39 pooled across org |

- Code completions + Next Edit Suggestions remain **unlimited on paid plans and consume no credits**; chat, agent mode, code review, cloud agent, CLI draw credits. Business/Enterprise credits pool org-wide; no rollover; admin spend caps **[independent — Second Talent]**.
- Copilot code review also consumes GitHub Actions minutes (two billing surfaces) **[secondary — future-stack-reviews]**.
- **Adoption signal:** in the April–June 2026 quarter (first month of usage-based billing), Copilot revenue grew **>60% QoQ** **[independent — Second Talent, citing quarter]**.

### 5.4 Key sources
- https://github.blog/changelog/2026-08-31-github-copilot-in-vs-code-august-2026-releases/
- https://github.blog/changelog/2026-07-30-github-copilot-in-visual-studio-code-july-2026-releases/
- https://Www.techtimes.com/articles/319982/20260709/github-closes-agentic-loop-vs-code-browser-tools-go-ga-sessions-run-parallel.htm
- https://www.secondtalent.com/resources/github-copilot-statistics/
- https://www.guickly.com/blogs/github-copilot-pricing-2026

---

## 6. ZCode (Z.ai / Zhipu AI)

### 6.1 What it is (and is not)
- **ZCode is the desktop "Agentic Development Environment" of Z.ai** (Beijing-based lab, formerly Zhipu AI), launched **July 2, 2026** as the native environment for its GLM coding models **[independent — VentureBeat]**.
- **Not a model** — the application layer around GLM-5.2 (analogous to Claude Code vs. Claude Opus): desktop interface, long-horizon task structure, workspace management, remote-control surfaces **[secondary — tokenkarma]**.
- An **Electron desktop app** (macOS, Windows, Linux) where the agent conversation is the interface center, surrounded by file manager, terminal, Git panel, and live browser preview **[secondary — pondero.ai]**.
- Tuned for **long-horizon "Goal" tasks**: user describes an outcome; the agent plans, edits, runs checks, iterates; sensitive commands require confirmation **[secondary — pondero.ai; community — ai-agent-map]**.
- **Remote steering:** monitor/direct a running session from **WeChat, Feishu, or Telegram** on a phone **[independent — VentureBeat; secondary — pondero.ai]**.
- **BYOK** for third-party models (Claude Code, Codex, Gemini reported); one community source also lists OpenCode as a supported agent **[secondary — ai-agent-map]**.
- **Proprietary client**; the GLM model line is separately open-weighted (GLM-5.2 described as MIT-licensed by one source) **[secondary — pondero.ai]**.

### 6.2 Pricing
- **Free to download**; revenue via GLM Coding Plan subscriptions: **Lite $16.20/mo**, **Max $144/mo** — undercutting Cursor ($20/$200) **[secondary — tokenkarma; pondero.ai]**.
- Free tier exists as a **promotional offer** (not committed permanent) **[secondary — tokenkarma]**.
- Launch promo: **1.5× usage-quota bonus** for Coding Plan subscribers through **July 31, 2026**; off-peak token consumption at **0.67×** coefficient **[secondary — cryptonwz/VentureBeat-derived]**.
- GLM-5.2 (unveiled June 13–16, 2026): 1M-token lossless context; Z.ai claims inference costs ~1/5 cheaper than comparable competitors **[vendor-reported]**.
- Noted in industry coverage as a pricing disruption even as real-world robustness remained unproven (WindowsForum relay, July 2026) **[secondary]**.

### 6.3 Version & uncertainties
- Community agent map lists **current version 3.11.2** and describes it as tuned around **GLM-5.3** — but launch coverage (July 2026) centers on **GLM-5.2**; **version/model drift is unverified — flag as uncertain** **[secondary — ai-agent-map]**.
- Gartner's ~$10B agentic-coding market estimate quoted in launch coverage **[independent — VentureBeat, citing Gartner]**.
- No adoption/user figures confirmed in this research.

### 6.4 Key sources
- https://venturebeat.com/technology/z-ai-launches-zcode-to-challenge-cursor-claude-code-and-github-copilot-in-ai-coding
- https://tokenkarma.app/blog/zcode-ai-coding-assistant-cursor-alternative-2026/
- https://pondero.ai/news/2026-07-04-zcode-launch/
- https://github.com/weijt606/ai-agent-map/blob/HEAD/agents/zcode.md

---
