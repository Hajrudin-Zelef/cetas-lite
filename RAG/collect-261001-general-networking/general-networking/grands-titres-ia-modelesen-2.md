---
id: collect-261001-general-networking/general-networking/grands-titres-ia-modelesen-2
title: "VOLET 1 — Vague 1 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Microsoft", "OpenAI", "Z.ai", "xAI"]
dates: ["2025-01", "2026-02", "2026-02-19", "2026-03-26", "2026-05-04", "2026-06", "2026-09", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "agi", "benchmark", "benchmarks", "claude", "compute", "consumer", "copilot", "cost", "deepseek"]
source: docs/RAG/collect-261001-general-networking/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [91, 183]
sha256: e21498d4056bed5d6a3b8eab2da378536dd9d58894e75383f57bf13007141e14
---

# VOLET 1 — Vague 1 (EN)

### What distinguished it in 2026
- **Terminal-first agentic specialist**: it dominated terminal-based workflows (Terminal-Bench 2.0 77.3% vs rivals) at 2.9× cheaper input pricing than Opus-class models — the price-efficiency leader for terminal-heavy engineering, with enterprise-grade LTS stability.

---

## 3. Gemini 3.1 Pro (Google DeepMind)

### Release & announcement
- **Released February 19, 2026 — in Preview** (`gemini-3.1-pro-preview`), ~3 months after Gemini 3 Pro (Nov 2025) — a fast turnaround by frontier standards.
- Preview → staged rollout: developer launch in Feb; consumer rollout accelerated May 4, 2026 in the Gemini app (expanded limits for AI Pro/Ultra subscribers) and NotebookLM for paid users. Gemini 3 Pro was discontinued on Vertex AI on March 26, 2026, pushing enterprise customers onto the 3.1 track.

### Architecture & technical specs
- **Architecture:** Mixture-of-Experts multimodal reasoning model with extended chain-of-thought ("thinking" / reasoning levels: Minimal, Low, Medium, High — configurable reasoning compute per request).
- **Context:** 1,048,576 input tokens (~1M; ≈1,500 A4 pages) → suitable for whole-codebase/long-context RAG. **Output:** 65,536 tokens per response.
- **Multimodality (native):** text, images, video, audio, PDF, code repositories as first-class inputs (e.g., ~45 min video with audio, 10 videos per prompt); text output only (no image/audio generation in preview). Live API not supported.
- Knowledge cutoff: January 2025 (API docs).
- Output speed: ~104–116 tokens/sec (well above peer median of ~72 t/s).
- **Dedicated agentic endpoint:** `gemini-3.1-pro-preview-customtools` — optimized for custom tool calls (view_file, search_code, bash) and multi-step agentic pipelines.

### Key benchmarks (reported values, Feb 2026)
| Benchmark | Score |
|---|---|
| **ARC-AGI-2** | **77.1%** — more than double Gemini 3 Pro (31.1%) |
| **SWE-Bench Verified** | **80.6%** (Google's figure; ~75% on third-party BenchLM standardization) |
| GPQA Diamond | 94.3% |
| MMMLU | 92.6% |
| MMMU / MMMU-Pro | ~62–64% / ~80–84% (independent trackers) |
| APEX-Agents | 33.5% (vs Claude Opus 4.6's 29.8%) |
- Headline claim: **1st place on 13 of 16** evaluated benchmarks (Google).

### API pricing & availability
- **$2.00 / 1M input, $12.00 / 1M output** for prompts ≤200K tokens; **$4.00 / $18.00** above 200K — **unchanged from Gemini 3 Pro**, making it a "free performance upgrade." Context caching supported (up to ~75% cost reduction); cached input ~$0.20/1M.
- Consumer: bundled in Google AI Pro and AI Ultra subscriptions (regional pricing varies).
- Channels: Gemini app, AI Studio, Gemini API, Vertex AI, Gemini CLI, Google Antigravity platform, Android Studio, NotebookLM, GitHub Copilot / VS Code extensions.

### Agentic / coding use cases
- **Creative & agentic coding:** code-based SVG animations, 3D visualizations, live aerospace dashboards, interactive design prototypes from text prompts alone.
- **Long-context RAG workflows:** entire codebases, contracts, research corpora in a single pass.
- Enterprise: Batch API, context caching, function calling, code execution, search grounding, structured outputs, URL context.

### What distinguished it in 2026
- **The reasoning-leap release at flat pricing**: generational jump (ARC-AGI-2 doubled, SWE-Bench 80.6%) with **zero price increase** — and the multimodal/reasoning flagship anchoring Google's entire 2026 ecosystem (Antigravity, AI Studio, NotebookLM, Android Studio), with a purpose-built agentic tool-calling endpoint.

---

## Comparison summary (February 2026 wave)

| Dimension | Claude Sonnet 5 "Fennec" | GPT-5.3-Codex | Gemini 3.1 Pro |
|---|---|---|---|
| **Release** | Feb 3, 2026 | Feb 5, 2026 | Feb 19, 2026 (Preview) |
| **Context** | 1M tokens | 400K tokens | 1M tokens |
| **SWE-Bench Verified** | **82.1%** | 56.8% (SWE-Bench Pro) | 80.6% |
| **Terminal-Bench 2.0** | 80.4% (2.1) | **77.3%** | — |
| **Headline reasoning** | Dev Team orchestration | Self-recursive training | ARC-AGI-2 **77.1%** |
| **Input / Output $ per 1M** | $3 / $15 | $1.75 / $14 | $2 / $12 |
| **Standout** | Agentic autonomy, best price/perf | Terminal workflows + LTS stability | Multimodal reasoning at flat pricing |
# VOLET 1 — Vague 1: Frontier AI Models (February → September 2026)
## Research fiche: Grok 4.20 · DeepSeek V4 / V4.1 Flash · GLM-5.2 / GLM-5.3 / GLM-5.3-Flash

*Research conducted 2026-09-22. All benchmark figures are vendor-reported unless marked as independently measured. Vendor-published numbers should be treated as upper bounds, not final verdicts.*

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

