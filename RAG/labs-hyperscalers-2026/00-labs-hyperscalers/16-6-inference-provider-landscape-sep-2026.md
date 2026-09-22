---
id: labs-hyperscalers-2026/00-labs-hyperscalers/16-6-inference-provider-landscape-sep-2026
title: "16.6 Inference-provider landscape (Sep 2026)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Baseten", "Cerebras", "Cohere", "CoreWeave", "EU", "Fireworks AI", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nebius", "Nvidia", "OpenAI", "Perplexity", "Sakana", "Together AI", "United States", "xAI"]
dates: ["2026-05", "2026-05-13"]
keywords: ["inference", "agent", "agents", "arr", "astra", "chatgpt", "claude", "cohere", "fable 5", "fugu", "gemini", "gpt-5.6"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2373, 2421]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 5c0737d7a48c8c0dc87168ef2eb9124457de8dfe4790b4f5e07e39d74afd931e
---

# 16.6 Inference-provider landscape (Sep 2026)

### 16.6 Inference-provider landscape (Sep 2026)

| Provider | Model | Differentiator | Status |
|---|---|---|---|
| Groq (GroqCloud) | LPU inference | Deterministic low-latency; ~$0.35/$0.79 per 1M (Llama 3.3 70B class); free tier | Operating; EU expansion 2026 |
| Cerebras Inference | Wafer-scale WSE-3 | Very high tok/s on Llama-class | Public since May 13, 2026 IPO |
| Together AI | GPU serverless | Open-model catalog; served Inkling at launch | Series C Jul 1, 2026 |
| Fireworks AI | GPU serverless | Function-calling/compound AI; served Inkling | Series D Jul 16, 2026 |
| SambaNova | Full-stack (chips+cloud) | Samba-1 / Composition of Experts | Series F Jul 8, 2026 |
| Nebius | GPU cloud + serverless | EU/US buildout; reported Meta orders | Public (NBIS) |
| CoreWeave | Reserved GPU capacity | Debt-financed fleet; Sep 17, 2026 converts | Public (CRWV) |
| NVIDIA NIM | Containers (AI Enterprise $4,500/GPU/yr) | Nemotron 3 Nano/Super/Ultra; DGX Cloud | Operating |
| Baseten / Modal / Databricks | Serverless | Served Inkling at launch | Operating |
| FreeLLMAPI | Open-source router | 34 providers / 7.4B monthly tokens (repo-reported) | ToS risk flagged |

### 16.7 Revenue, users & adoption scoreboard (Sep 2026)

| Company / product | Metric | Date | Provenance |
|---|---|---|---|
| Anthropic total | $25B annualized run rate (enterprise 80%) | Aug 2026 | [independent] |
| Claude Code | $2.5B run rate | May–Aug 2026 | [independent] |
| OpenAI total | $25B (Feb, CFO-confirmed) → >$40B (Aug, Bloomberg) | 2026 | [independent] |
| ChatGPT | ~1B active users cited | Jul 2026 | [vendor-reported] |
| Codex + ChatGPT Work | 10M combined users | Jul 21, 2026 | [vendor-reported] |
| Codex | 5M+ weekly users | Jun 2026 | [vendor-reported] |
| Gemini app | 27.7% mobile AI-app share (Mar 2026) | Mar 2026 | [independent: SmashingApps] |
| ChatGPT | 46.4% mobile AI-app share (Mar 2026) | Mar 2026 | [independent] |
| Claude | 10.3% mobile AI-app share (Mar 2026) | Mar 2026 | [independent] |
| Vibe (ex-Le Chat) Android | ~1.9M total downloads (+48K/30d) | Sep 2026 | [secondary: AppBrain] |
| Sakana Fugu | Pay-as-you-go from Jun 22, 2026 | 2026 | [vendor-reported] |
| Cohere | $240M ARR run rate | 2026 | [independent: Reuters] |
| Perplexity | >$750M annualized revenue; ~780M queries/mo | Aug 2026 | [secondary] |
| Google | >3.2 quadrillion tokens/month processed | May 2026 | [vendor-reported] |
| Hugging Face agents | Claude Code 44.4% of agent actions; Codex 10.4%→20.8% (Apr–Jul) | Jul 2026 | [independent: HF dataset] |

### 16.8 Frontier release cadence — days between flagship launches (Feb–Sep 2026)

| Lab | Flagship releases in-window | Cadence notes |
|---|---|---|
| Anthropic | Opus 4.6 (Feb 5) → Sonnet 4.6 (Mar 26) → Fable 5 (in-window) | ~7-week major cadence; Fable 5 date n/r |
| OpenAI | GPT-5.3-Codex (Feb 5) → GPT-5.4 (Mar 5) → GPT-5.5 (Apr 22) → GPT-5.6 (Jul 9) → GPT-6 Astra (Sep 3) | 28d → 48d → 78d → 56d; accelerating major-version cadence |
| Google | Gemini 3.1 Pro (Feb 19) → 3.5 Flash (May 19) → 3.6 Flash (card Jul 21) | ~89d → ~63d; flash-tier refresh quickening |
| Meta | Muse Spark 1.0 (Apr 8) → 1.1 (Jul 9) → 1.2 (Aug 5) → 1.3 (Sep 2) | 92d → 27d → 28d; sharp acceleration mid-window |
| Microsoft | MAI family (Jun 2, all seven at once) | Single-drop strategy at Build 2026 |
| xAI | Grok 4.20 (Feb 17) → 4.3 (Apr 17/30) → 4.5 (Jul 8) → 4.6 (Aug 12) → 4.7 (Sep 21) | ~59d → ~82d → ~35d → ~40d |
| Mistral | Voxtral T2 (Feb 4) → Small 4 (Mar 16) → Medium 3.5 (Apr 28) → OCR 4 (~Jun) | ~40d → ~43d; steady monthly-ish drops |

**Read:** OpenAI and xAI shipped 4–5 flagship iterations each in ~7.5 months; Meta compressed from quarterly to monthly; Google's flash tier refreshed twice. The industry-wide cadence in 2026 is roughly one frontier-class release per lab per 4–8 weeks. [research finding — derived from dated entries in §15]

