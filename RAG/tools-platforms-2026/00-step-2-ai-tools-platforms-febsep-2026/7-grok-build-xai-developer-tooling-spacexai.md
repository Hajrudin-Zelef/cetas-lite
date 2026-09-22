---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/7-grok-build-xai-developer-tooling-spacexai
title: "7. Grok Build / xAI developer tooling (SpaceXAI)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "DeepSeek", "Google", "Microsoft", "OpenAI", "OpenRouter", "SpaceX", "United States", "xAI"]
dates: ["2026-01", "2026-03-10", "2026-03-16", "2026-05-06", "2026-05-15", "2026-05-21", "2026-05-25", "2026-05-28", "2026-06", "2026-07-09", "2026-08", "2026-08-12", "2026-08-13", "2026-09-21"]
keywords: ["grok", "agent", "agentic", "agents", "astra", "attribution", "bedrock", "benchmark", "benchmarks", "claude", "copilot", "cost"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [279, 398]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 85c9b5dc8f24022860651fc8505a3acfcce5c2e8c78ece002a7cb49165aed7d5
---

# 7. Grok Build / xAI developer tooling (SpaceXAI)

## 7. Grok Build / xAI developer tooling (SpaceXAI)

### 7.1 xAI API — current model catalog and pricing (official)
Source: SpaceXAI official API pricing page, crawled within hours of this report (`docs.x.ai/developers/pricing`) — **[official]**. Prices in USD per 1M tokens. Models with long-context pricing bill the long-context rates for **all** tokens in a request once the prompt reaches the model's long-context threshold (≥ 200K tokens). Requests to the US regional endpoint run inference in the United States and are billed at 1.1× the global token rates (10% premium).

| Model | Context | Input | Cached input | Output | Long-context (in / cached / out) |
|---|---|---|---|---|---|
| grok-4.7 | 500K | $2.00 | $0.50 | $6.00 | $4.00 / $1.00 / $12.00 |
| grok-build-0.1 | 256K | $1.00 | $0.20 | $2.00 | $2.00 / $0.40 / $4.00 |
| grok-4.6 | 500K | $2.00 | $0.50 | $6.00 | $4.00 / $1.00 / $12.00 |
| grok-4.5 | 500K | $2.00 | $0.30 | $6.00 | $4.00 / $0.60 / $12.00 |
| grok-4.3 | 1M | $1.25 | $0.20 | $2.50 | $2.50 / $0.40 / $5.00 |
| grok-4.20-multi-agent-0309 | 1M | $1.25 | $0.20 | $2.50 | $2.50 / $0.40 / $5.00 |
| grok-4.20-0309-reasoning | 1M | $1.25 | $0.20 | $2.50 | $2.50 / $0.40 / $5.00 |
| grok-4.20-0309-non-reasoning | 1M | $1.25 | $0.20 | $2.50 | $2.50 / $0.40 / $5.00 |

**API characteristics** **[official]**/**[vendor-reported]**:
- OpenAI-compatible REST endpoint (`https://api.x.ai/v1`); drop-in migration from OpenAI/Anthropic SDKs by swapping base URL and key.
- Function/tool calling, structured outputs, vision (image input), image generation (reported ~$0.07/image — **[secondary/unverified]**), real-time web and X search.
- Third-party routing (OpenRouter) was observed billing Grok 4.7 at $1.60/$4.80 (standard tier) — route prices can sit below list price **[secondary/unverified]**.

### 7.2 Grok 4.x release timeline (developer-relevant)
- **Grok 4.3** — opened to all API developers **May 6, 2026** (beta since April 17 behind SuperGrok Heavy). Pricing cut at launch: output $15.00 → $2.50 (-83%), input $3.00 → $1.25 (-58%). 1M-token context, native video input (up to 5 min: mp4/mov/webm), document generation (PDF, XLSX, PPTX). Five legacy models retired May 15, 2026 **[secondary/unverified]** (reported by Awesome Agents).
- **Grok 4.20** — API general availability reported **March 10, 2026** (three dated SKUs: `-0309-reasoning`, `-0309-non-reasoning`, `-multi-agent-0309`) **[secondary/unverified]** (reported by Venice model datasheet; no official xAI launch post located).
- **Grok 4.5** — public API **July 9, 2026**, after private testing inside SpaceX/Tesla engineering teams **[secondary/unverified]** (AI Cost Estimator). Positioned as flagship coding/agentic model at $2/$6.
- **Grok 4.6** — released **August 12, 2026**, announced by the official @SpaceXAI account **[vendor-reported]** (relay by pasqualepillitteri.it). Not a new base training run: reuses a 1.5-trillion-parameter "V9" base (per the same relay — **unverified**), improved via supervised fine-tuning + reinforcement learning; emphasis on long-horizon agents, autonomous coding, self-verification. Day-one availability in **Cursor**, **Grok Build**, **Grok Bot** agents, and the API; 2× usage included in Cursor and Grok Build for the first week.
- **Grok 4.7** — released **September 21, 2026** (one day before this report), announced by @SpaceXAI **[vendor-reported]** (relays: Metaverse Post, Unite.AI, The Decoder, XenoSpectrum). Key facts:
  - New, larger base than 4.6; extended RL run on harder, multi-hour tasks; pretraining data cutoff June 2026, supplemental data as late as August 2026. Parameter count not disclosed in the model card; a "2.1 trillion" figure attributed to Musk's pre-launch remarks is **unconfirmed** **[secondary/unverified]** (AI Release Tracker).
  - Vendor-reported benchmarks: CursorBench 4.0 46.3% (vs 40.4% for 4.6; vs GPT-5.6 Sol 41.7%), DeepSWE v1.1 high-effort 71.0% (vs 65.2% for 4.6; vs GPT-5.6 Sol 72.7%). Independent Artificial Analysis numbers: Intelligence Index v4.3.2 = 46 (vs 53 for Claude Fable 5.1 and GPT-6); Terminal-Bench 4.0 agentic coding = 26% (vs 60% GPT-6 Astra, 55% Claude Fable 5.1) **[independent]** (via The Decoder).
  - New safeguard stack: 3.3% pass-through of risky dual-use prompts on internal HackerBench v0.3; tops LatchBio's biosafety benchmark at 62.4% **[vendor-reported]**; invite-only red-team access for select cybersecurity partners.
  - Pricing unchanged at $2/$6; a **fast variant** offers 2× output speed at 2× price ($4/$12). Free access inside Grok Build at x.ai/build **[vendor-reported]**/**[secondary]**.
  - Distribution at launch: xAI API, Grok Build, Cursor (all plan tiers), Grok Bot harness, third-party harnesses, routers/cloud platforms; phased rollout to **GitHub Copilot** (Pro, Pro+, Max, Business, Enterprise) began the same day **[secondary/unverified]** (XenoSpectrum).

### 7.3 Grok Build — xAI's agentic coding CLI
- **Launch:** early beta **May 14–15, 2026**, initially restricted to **SuperGrok Heavy** subscribers (reported $300/month) **[vendor-reported]** (relays: AlternativeTo, fonearena, Crypto Briefing). Windows PowerShell installer released **May 25, 2026**.
- **Install commands** (as reported): macOS/Linux `curl -fsSL https://x.ai/cli/install.sh | bash`; Windows `irm https://x.ai/cli/install.ps1 | iex`.
- **Features** **[vendor-reported]**: interactive, mouse-capable fullscreen TUI; plan mode (AI proposes a plan, user approves/comments/rewrites before execution); clean diff tracking; parallel subagents with deep git-worktree integration; headless mode (`-p` flag) for scripts, bots, and CI pipelines; respects existing repo conventions out of the box (`AGENTS.md`, plugins, hooks, skills, MCP servers); **full ACP (Agent Client Protocol) support** for building custom bots and agent orchestration apps (ACP is the open agent-client protocol originated by Zed); built-in `/feedback` command.
- **Model:** powered by `grok-build-0.1`, a purpose-built agentic coding model (256K context, 100+ tok/s reported throughput). Public API beta **May 28, 2026** (no subscription required) at $1/$2 per 1M tokens **[secondary/unverified]** (GitHub issue tracker relay; consistent with the official pricing table above).
- **Availability evolution:** early beta → SuperGrok Heavy only; Grok 4.6 (Aug 12) and Grok 4.7 (Sep 21) both shipped day-one inside Grok Build; Grok 4.7 free inside Grok Build per xAI's launch messaging.

### 7.4 Grok inside OpenCode (integration, May 21, 2026)
On **May 21, 2026**, xAI rolled out an integration letting Grok subscribers use xAI models directly inside **OpenCode** (open-source, provider-agnostic terminal coding agent; reported 160K+ GitHub stars, 75+ providers supported) — no separate API key required; an existing SuperGrok or X Premium subscription acts as the credential **[secondary/unverified]** (Memeburn). Strategic note: the move followed Anthropic's January 2026 blocking of third-party tools from Claude Pro/Max subscriptions, which had pushed developers toward model-agnostic harnesses.

### 7.5 Voice and retrieval APIs
- **Grok Text-to-Speech API** — launched **March 16, 2026** **[vendor-reported]** (relay: BaseNor, citing @xai). 5 voices (Ara, Eve, Leo, Rex, Sal), 20+ languages with auto-detection + BCP-47 codes, inline expressive speech tags (pauses, laughter, whispers, emphasis); output formats MP3, WAV, PCM (Linear16), G.711 μ-law/A-law.
- **Voice Agent API** — launched late 2025 (realtime voice apps) **[secondary/unverified]**.
- **Grok Collections** — RAG/document-retrieval building block **[secondary/unverified]** (mentioned in a January 2026 Medium overview; no official docs page located — **verify before final consolidation**).

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
