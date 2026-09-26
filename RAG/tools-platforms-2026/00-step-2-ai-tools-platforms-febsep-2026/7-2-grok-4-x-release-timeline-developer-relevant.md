---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/7-2-grok-4-x-release-timeline-developer-relevant
title: "7.2 Grok 4.x release timeline (developer-relevant)"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI", "OpenRouter", "SpaceX", "xAI"]
dates: ["2026-01", "2026-03-10", "2026-03-16", "2026-05-06", "2026-05-15", "2026-05-21", "2026-05-25", "2026-05-28", "2026-06", "2026-07-09", "2026-08", "2026-08-12", "2026-09-21"]
keywords: ["grok", "grok 4", "agent", "agentic", "agents", "astra", "benchmark", "benchmarks", "claude", "copilot", "cost", "cybersecurity"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [284, 326]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 15bee3fce75df6e5754e193f262f4211af7dc7d13c36dac68ad8f0e32653e963
---

# 7.2 Grok 4.x release timeline (developer-relevant)

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

