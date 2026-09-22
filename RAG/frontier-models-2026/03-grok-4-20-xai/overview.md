---
id: frontier-models-2026/03-grok-4-20-xai/overview
title: "1. Grok 4.20 (xAI)"
domain: grok-4-20-xai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Google", "OpenAI", "xAI"]
dates: ["2026-02", "2026-06", "2026-09"]
keywords: ["grok", "grok 4", "agent", "agents", "agi", "benchmark", "benchmarks", "claude", "compute", "consumer", "gemini", "gpu"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [154, 196]
section: "1. Grok 4.20 (xAI)"
sha256: 37553465a5784e565fd48d48ab31cfb353c577e86980c59cb798ab205c8e08ff
---

# 1. Grok 4.20 (xAI)

## 1.1 Release & status
- **Public beta launched 17 February 2026.** Elon Musk announced it in a post (no formal launch window — "it was live"); users had to manually select "Grok 4.2" in the model menu.
- **Rapid-learning architecture:** first Grok that improves continuously after deployment — capabilities updated weekly based on user feedback, with release notes published alongside each update. Musk stated Grok 4.20 would be "an order of magnitude smarter and faster" than Grok 4 by the time the beta concludes.
- **Still in beta as of mid-2026**, accessible to SuperGrok subscribers; by September 2026 the API lists dated snapshot variants (`grok-4.20-multi-agent-0309`, `grok-4.20-0309-reasoning`, `grok-4.20-0309-non-reasoning`), indicating the line is in active production use while successor `grok-4.3` rolls out in stages.
- Prior cadence: Grok 4 (9 Jul 2025, native tool use + real-time search, RL at pretraining scale on 200K GPUs), Grok 4 Heavy (9 Jul 2025), Grok 4.1 (17 Nov 2025, emotional-intelligence upgrade). Grok 4.20 arrived ~3 months after Grok 4.1.

## 1.2 Architecture: four-agent system (multi-agent)
- **Not one monolithic model — four specialized agents** that think in parallel and debate in real time before answering:
  - **Grok** — lead orchestration
  - **Harper** — research and fact-checking
  - **Benjamin** — logical verification
  - **Lucas** — creative synthesis
- Agents cross-check outputs and reach **"adversarial consensus"** before the lead agent synthesizes the final response.
- Headline safety metric: **hallucination rate dropped from ~12% to ~4.2% (−65%)**.
- **"Heavy" mode scales to 16 agents** for demanding tasks, drawing on xAI's 200,000-GPU Colossus supercluster.
- **Context:** 256K window, up to 2M (per third-party analysis).
- **Native multimodal:** text + image + video; video generation and revamped image generation added to the platform in early 2026; Grok Imagine 1.5 (image/video engine) as of June 2026.
- **Training:** reinforcement learning at pretraining scale on the 200K-GPU Colossus supercluster.
- **New capability highlights at beta:** medical document analysis via photo upload; improved engineering reasoning; financial reasoning (strong showing in live trading simulations, e.g. Alpha Arena: +34.59% returns in one report, ~12.11% average / 47% max in a stealth "mystery model" debut per crypto press).
- **Musk roadmap note:** Grok 5 (next flagship, aimed at AGI-level breakthroughs, reportedly ~6T parameters — 2× Grok 4) was expected "in a few months" from early 2026.

## 1.3 Benchmarks
- **LMArena Elo: estimated 1505–1535 (provisional).** Grok 4.1 Thinking already sat at 1483; the multi-agent council + extra inference-time compute was projected to add 20–60 Elo, potentially #1 overall. No official provisional ranking at beta time (internal/beta status).
- **ForecastBench: #2.**
- **Safety Overfit tests: strong (per NextBigFuture).**
- **Live trading competitions (Alpha Arena): +34.59% returns while competitors posted losses** (vendor-adjacent claim; treat with caution).
- **Comparative press (mid-2026):** OpenAI GPT-5.4 "wins on reliability and reasoning; Grok 4.20 wins on personality and speed" — both described as the most human-feeling assistants either company had shipped.

## 1.4 License & availability
- **Closed/proprietary.** No open weights.
- **Consumer access:** X Premium+ and SuperGrok subscribers (manual model selection at beta); SuperGrok Heavy ($300/mo) unlocks Grok 4 Heavy multi-agent (16-agent) mode. Free tier limited; X Premium bundles include Grok access.
- **Developer API (api.x.ai / console.x.ai):** dated model IDs — `grok-4.20-multi-agent-0309`, `grok-4.20-0309-reasoning`, `grok-4.20-0309-non-reasoning`. Dated IDs are the only route to a guaranteed version (consumer apps don't disclose which variant answered).
- **API pricing (2026, per 1M tokens):** input $1.25 / cached input $0.20 / output $2.50 — same as `grok-4.3`; 1M context. (Dedicated coding model `grok-build-0.1`: $1.00 in / $2.00 out, 256K context.) xAI changes API prices with little warning.
- **Subscription vs API billing are separate products:** API usage is pay-per-token and does NOT draw on SuperGrok subscription quotas.

## 1.5 Position vs competitors (Feb 2026)
- Multi-agent "council" approach differentiated it from monolithic frontier models (Claude, GPT, Gemini); hallucination-reduction narrative targeted enterprise reliability concerns.
- Weaknesses at launch: limited official benchmark suite; beta-only availability; personality-forward brand that critics saw as a liability.

---

