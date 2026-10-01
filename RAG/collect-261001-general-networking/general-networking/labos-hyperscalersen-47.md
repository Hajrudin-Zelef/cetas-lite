---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-47
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Baseten", "Cerebras", "Cohere", "CoreWeave", "EU", "Fireworks AI", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nebius", "Nvidia", "OpenAI", "Perplexity", "Sakana", "Together AI", "United States", "xAI"]
dates: ["2026-05", "2026-05-13"]
keywords: ["agent", "agentic", "agents", "agi", "arr", "astra", "benchmark", "chatgpt", "claude", "cohere", "fable 5", "fugu"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2314, 2423]
sha256: 2c0637d8670a3a37c6ecf6ef48d54b4c555e81edf8f4cba4af1bc873e6c7fedb
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**SWE-bench lineage (vendor-reported unless noted; harness versions differ):**
| Model | Benchmark | Score | Provenance |
|---|---|---|---|
| GPT-5.3-Codex | SWE-Bench Pro | 56.8% | [vendor-reported] |
| Muse Spark 1.3 | DeepSWE v1.1 | 75.4% | [vendor-reported] (vs GPT-6 Astra 74.1% same table) |
| GPT-6 Astra | DeepSWE v1.1 | 74.1% | [vendor-reported] |
| Grok 4.5 | SWE-Bench Pro | 64.7% | [vendor-reported] |
| Grok 4.6 | DeepSWE v1.1 | 65.9% | [vendor-reported] |
| Grok 4.6 | FrontierCode v1.1 Ext | 61.3% | [vendor-reported] |
| Grok 4.7 | DeepSWE v1.1 (high-effort) | 71.0% | [vendor-reported] |
| Mistral Medium 3.5 | SWE-Bench Verified | 77.6% | [vendor-reported] |
| Sakana Fugu Ultra v2 | DeepSWE | 74.3 | [vendor-reported] |
| Anthropic Opus 5 | SWE-bench (harness A) | 96–97% | [vendor-reported — conflicts with 74.8% below] |
| Anthropic Opus 5 | SWE-bench (harness B) | 74.8% | [vendor-reported — likely different harness/config; see §18] |

**Terminal-Bench lineage (vendor-reported; 2.0 vs 2.1 vs 3.0 vs 4.0 NOT comparable):**
| Model | Version | Score |
|---|---|---|
| GPT-5.3-Codex | 2.0 | 77.3% |
| GPT-5.5 | 2.0 | 82.7% |
| GPT-5.6 Sol | 2.1 | 88.8% (91.9% ultra) |
| Gemini 3.5 Flash | 2.1 | 76.2% |
| Gemini 3.6 Flash | 2.1 | 78.0% |
| Grok 4.5 | 2.1 | 83.3% |
| Grok 4.6 | 3.0 | 26% |
| GPT-6 Astra | 4.0 | 57.9% |
| Fable 5 (vendor) | vendor harness | 88.0% |
| Fable 5 (Vals AI independent) | independent re-run | 80.52% |

**Agentic / computer-use (vendor-reported):**
| Model | Benchmark | Score |
|---|---|---|
| GPT-5.5 | OSWorld-Verified | 78.7% |
| GPT-6 Astra | OSWorld 2.0 | 72.6% |
| Gemini 3.6 Flash | OSWorld-Verified | 83.0% |
| Grok 4.3-era | AA-Briefcase Elo | 1,546 (4.20) → 1,546-class |
| Grok 4.6 | AA-Briefcase Elo | 1,577 |
| Grok 4.7 | AA-Briefcase Elo | 1,657 |
| Grok 4.5 | AutomationBench-AA | 51% (#1 at launch) |

**Reasoning / knowledge (vendor-reported):**
| Model | Benchmark | Score |
|---|---|---|
| Gemini 3.1 Pro | ARC-AGI-2 | 77.1% |
| Gemini 3.5 Flash | ARC-AGI-2 | 77.1% |
| Gemini 3.5 Flash | MCP Atlas | 83.6% |
| Gemini 3.5 Flash | CharXiv Reasoning | 84.2% |
| Gemini 3.6 Flash | CharXiv Reasoning | 85.2% (no tools) / 89.4% (with tools) |
| Gemini 3.6 Flash | GDM-MRCR v2 | 91.8% (128k) / 54.0% (1M) |
| MAI-Thinking-1 | AIME 2025 / 2026 | 97.0% / 94.5% |
| Mistral Small 4 | GPQA Diamond | 71.2% |
| Mistral Small 4 | MMLU-Pro | 78.0% |
| Grok 4.6 | GDPVal-AA v2 Elo | 1,753 |
| Gemini 3.6 Flash | GDPVal-AA v2 Elo | 1,421 |

**Coding-agent leaderboards (independent, dated):**
- Artificial Analysis Coding Agent Index: Grok 4.7 = 56 (+9 vs 4.6) on Sep 21, 2026 first pass. [independent]
- CursorBench: Grok 4.6 v3.2 69.9% → Grok 4.7 v4.0 46.3% (version change; not comparable). [vendor-reported]

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

### 16.9 Master model index — every model in this file (lab · release date · access · status)

