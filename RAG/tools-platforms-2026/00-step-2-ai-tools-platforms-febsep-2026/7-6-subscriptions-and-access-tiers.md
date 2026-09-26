---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/7-6-subscriptions-and-access-tiers
title: "7.6 Subscriptions and access tiers"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "DeepSeek", "Google", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-08-13"]
keywords: ["agent", "attribution", "bedrock", "benchmarks", "claude", "copilot", "cost", "deepseek", "governance", "gpt-6", "grok", "grok 4"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [327, 398]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: ed61fcfdeb7d03c9aa945314efe27cf1d1874fa4493103e2aad34405bc5dd05b
---

# 7.6 Subscriptions and access tiers

### 7.6 Subscriptions and access tiers
- **SuperGrok Heavy** — reported at **$300/month**; was the gate for Grok 4.3 beta, Grok Build early beta, and Grok 4.20-era "multi-agent" usage **[secondary/unverified]** (multiple relays; consistent across sources).
- Standard **SuperGrok** (~$30/mo class) and **X Premium+** also serve as credentials for integrations such as Grok-in-OpenCode **[secondary/unverified]**.

### 7.7 Corporate/brand context relevant to tooling
- xAI now announces products via the **@SpaceXAI** account, the name used after the **xAI–SpaceX merger** earlier in 2026 **[vendor-reported]**/**[secondary]**.
- **Grok Bot** — referenced in official launch messaging as the conversational-agent harness Grok 4.6/4.7 natively understand; no standalone product page located **[vendor-reported]** (details thin — verify).

### 7.8 Key sources
- https://docs.x.ai/developers/pricing (official pricing)
- https://mpost.io/xai-ships-grok-4-7-with-new-safeguard-stack-independent-benchmarks-confirm-gains-flag-doubled-token-consumption/
- https://www.unite.ai/spacexai-releases-grok-4-7-for-coding-and-knowledge-work/
- https://the-decoder.com/xai-launches-grok-4-7-at-bargain-prices-but-benchmarks-reveal-a-wide-gap-to-claude-and-gpt-6/
- https://xenospectrum.com/en/xai-grok-4-7-pricing-copilot/
- https://aireleasetracker.com/model/xai/grok-4.7
- https://awesomeagents.ai/news/xai-grok-4-3-api-launch/
- https://ai-cost-estimator.com/blog/grok-4-5-public-launch-spacexai-vs-claude-opus-gpt-5-6-cost-per-task
- https://pasqualepillitteri.it/en/news/10915/grok-4-6-released-cursor-grok-build-api
- https://cryptobriefing.com/xai-grok-cli-windows-powershell/
- https://www.fonearena.com/blog/482869/xai-grok-build-coding-agent-features.html
- https://alternativeto.net/news/2026/5/xai-launches-grok-build-in-beta-a-powerful-coding-agent-and-cli-for-advanced-engineering/
- https://memeburn.com/how-to-use-grok-in-opencode-full-setup-guide-with-xai-api/
- https://www.basenor.com/blogs/news/xai-launches-grok-text-to-speech-api-5-voices-20-languages

---

## 8. DeepSeek Harness (`dsh`)

### 8.1 What it is
**DeepSeek Harness** is an **open-source (MIT) agent harness** (not a model): the engineering layer that turns a language model into a working agent — context management, tool use, file read/write, terminal execution, session logging, permissions, observability. DeepSeek's framing, widely echoed: **Agent = Model + Harness** — the model reasons, the harness does everything else.
- **Released:** **August 13, 2026**, announced by the official @deepseek_ai account on X **[vendor-reported]** (relays: pasqualepillitteri.it, Memeburn, Medium/Generative AI). Published as the GitHub repository `deepseek-harness` under the DeepSeek AI org — **repo URL itself [secondary/unverified], verify**.
- **Version status:** v0.1 **Developer Preview** — DeepSeek warns that compatibility-breaking changes should be expected.
- **Language:** TypeScript monorepo (web UI, headless runner, sandbox/execution layers, plugin framework), with unusually direct architecture documentation **[secondary/unverified]** (Medium/Generative AI, Sep 2026).
- Before launch it was known informally as **"DeepSeek Code"** (Memeburn).
- **Positioning:** the direct open-source challenger to Anthropic's Claude Code and OpenAI's Codex (both closed-source) — the open MIT license is the strategic point.

### 8.2 Architecture
- **"Everything is a plugin":** the model adapter, tool registry, session log, and the agent loop itself are all swappable plugins — deployers can replace models, storage, security policies, tools, or the entire agent loop without touching the core framework.
- **Cordis kernel:** powered by the **Cordis meta-framework** (third-party engine, `cordiverse/cordis` on GitHub); design described in the paper *"A Programming Paradigm for Spatiotemporal Composability"* (arXiv:2608.25512) **[secondary/unverified]** (attribution via Medium explainer — verify the arXiv ID).
- **Session modes (four):** Standard (full agent: files, shell, search, subagents, workflows), Code/PTC (the model writes one typed program instead of many tool round-trips), Minimal, Creator (the agent drafts new presets from conversation) **[secondary/unverified]**.
- **Session log:** every run is recorded in an **append-only session log** that can be resumed, forked, and replayed — repeatedly cited as the standout design idea.

### 8.3 Interfaces and model support
- **Local web UI:** `npx @deepseek-ai/dsh web` → serves at `http://127.0.0.1:3080` (npm package `@deepseek-ai/dsh`, terminal command `dsh`).
- **Headless CLI** for single tasks (CI/automation-friendly) and a **Python SDK**.
- **Model-agnostic:** works with DeepSeek's own API (default model reported as `deepseek-v4-flash`), Anthropic, OpenAI, Google Vertex, Azure, Amazon Bedrock, any OpenAI-compatible endpoint, and local Ollama models — provider is a replaceable plugin, not the system's identity **[secondary/unverified]** (community setup guides).

### 8.4 Adoption (date-stamped — figures move weekly)

| Date | Reported scale | Source class |
|---|---|---|
| Aug 13, 2026 (launch day) | ~27K–50K stars within hours | secondary |
| ~Aug 14 | 70K+ stars within a day | secondary |
| ~Aug 16–17 | ~95K–135K stars within 2–4 days; ~8.8K forks at day 4 | secondary |
| Early Sep 2026 | 208K+ stars, ~24.2K forks | secondary |
| Sep 16, 2026 | 225K+ stars | secondary |

Community plugin counts are reported between "a few thousand catalogued" and "6,000+ repos" depending on the counter and date **[secondary/unverified]**.
One analysis notes the harness performs best with DeepSeek's own models and flags plugin-system and contribution-governance concerns **[secondary/unverified]** (Geeky Gadgets).

### 8.5 Caveats and open questions
- **Preview status:** release candidates only; DeepSeek explicitly promises breaking changes. Not production-hardened.
- **To verify before final consolidation:** exact GitHub repo URL, the arXiv paper ID, the `deepseek-v4-flash` default-model claim, and the four session-mode names — all currently **[secondary/unverified]**.

### 8.6 Key sources
- https://pasqualepillitteri.it/en/news/11027/deepseek-harness-mit-claude-code-rival
- https://memeburn.com/deepseek-v4-pro-launches-with-major-agent-upgrades-and-open-source-harness/
- https://www.geeky-gadgets.com/deepseek-harness-local-ai/
- https://generativeai.pub/i-argued-the-harness-was-the-product-then-deepseek-open-sourced-theirs-d092f8b3dcde
- https://medium.com/@axotopia/the-machine-wears-the-harness-now-you-just-talk-0a951e02fd88

---
