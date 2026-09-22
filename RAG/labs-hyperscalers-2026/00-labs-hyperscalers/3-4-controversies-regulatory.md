---
id: labs-hyperscalers-2026/00-labs-hyperscalers/3-4-controversies-regulatory
title: "3.4 Controversies & regulatory"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: regulation
actors: ["Anthropic", "EU", "Google", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-02-19", "2026-03", "2026-05-19", "2026-06", "2026-07", "2026-07-21"]
keywords: ["agent", "agentic", "agi", "antitrust", "benchmark", "benchmarks", "claude", "cyber", "distribution", "gemini", "gpt-5.6", "grok"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [544, 596]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 0d35ecadf2206ca7d879c5f3dc60a6b5b8048637f95ba1d0bbb8dcb242ea497b
---

# 3.4 Controversies & regulatory

### 3.4 Controversies & regulatory
- **EU AI Act**: Google a GPAI provider in scope; systemic-risk obligations from 2 Aug 2026 (see §2.4 for the framework). [independent]
- **AI slowdown antitrust class action (18 Sept 2026)** names Google alongside Anthropic, OpenAI, SpaceXAI (see §1.4/§2.4). [independent: AP]
- "A Call for Collective Action on Cyber Defense" (27 Aug 2026) — Google signatory. [secondary]
- **Android XR / Gemini on-device** privacy discussions (Track B — minor).

### 3.5 Google benchmark snapshot (as reported)

| Model | Artificial Analysis Intelligence Index | Notes |
|---|---|---|
| Gemini 3.1 Pro | 66 @ $2.70/task [independent, index at the time] | Feb 19, 2026 launch (corrected) |
| Gemini 3.5 Flash | 57 @ $1.60/task [independent, index at the time] | May 19, 2026 launch (corrected) |

> Scores above were measured on different index revisions and are not directly comparable to each other or to scores in §§1–2.

---
### 3.6 Gemini 3.x — corrected release record (from Track B)

> Track B research corrected several dates/figures used in condensed summaries. The record below supersedes §3.1 where they differ.

- **Gemini 3.1 Pro — February 19, 2026** (not "March 2026"): GA via Gemini API, Vertex AI, and Antigravity platform; 1M-token context; natively multimodal (text, image, audio, video, code). Benchmarks (vendor-reported via press): **77.1% ARC-AGI-2**; leading 13 of 16 major benchmarks at launch incl. GPQA and HumanEval/MMLU families. API pricing: **$2.00 / 1M input, $12.00 / 1M output** (consistent with Google's own model-card comparison table [official]). Released alongside **Gemini 3.1 Flash-Lite**: ~2.5× faster responses, **$0.25 / 1M input**. [secondary] https://24rows.com/google-gemini-3-1-pro-launches-dominates-13-of-16-major-ai-benchmarks/
- **Gemini 3.5 Flash — May 19, 2026 (Google I/O keynote)** (not "June 2026"): 1M context; 64K max output; text/image/video/audio/PDF inputs. Positioned for agentic workloads: ~4× faster output tok/s than other frontier models. Launch benchmarks (vendor-reported): **Terminal-Bench 2.1 76.2%, MCP Atlas 83.6%, CharXiv Reasoning 84.2%, ARC-AGI-2 77.1%** — stated to beat Gemini 3.1 Pro on coding/agentic benchmarks. API pricing: **$1.50 / 1M input, $9.00 / 1M output** (~40% below 3.1 Pro; confirmed by Google's model-card table [official]). Distribution: default model for Google AI Ultra subscribers in Gemini app; AI Mode in Search (expanded to ~200 countries, 98 languages, no subscription); Antigravity agent platform; Gemini API in AI Studio. At announcement Google reported processing **>3.2 quadrillion tokens/month** across its surfaces [vendor-reported]. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/46-gemini-3-5-flash-google-io-2026.md
- **Gemini 3.6 Flash — model card published July 21, 2026** (not "September"): official model card [official] https://deepmind.google/models/model-cards/gemini-3-6-flash/ — the Gemini 3 series "workhorse": better coding/knowledge-work/multimodal than 3.5 Flash, better token efficiency; built on 3.5 Flash; 1M context, 64K output; knowledge cutoff March 2026. Pricing: **$1.50 / 1M input, $7.50 / 1M output** [official]. Benchmarks (vendor-reported, official table): **SWE-Bench Pro 58.7%, DeepSWE v1.1 49%, Terminal-Bench 2.1 78.0%, MLE-Bench 63.9%, GDPVal-AA v2 Elo 1421, OSWorld-Verified 83.0%, CharXiv Reasoning 85.2% (no tools) / 89.4% (with tools), GDM-MRCR v2 91.8% (128k avg) / 54.0% (1M pointwise)** [official]. Google compares against GPT-5.6 Luna, Grok 4.5, Claude Sonnet 5 in the same table — cross-vendor numbers are Google's own evaluations, not independently verified. Frontier Safety: did not reach any Critical Capability Levels; cyber remained below CCL after extra testing; note "previous models in the Gemini 3 series reached the alert threshold for cyber" [official].
- **Gemini 3.5 Flash-Lite — July 2026**: model card results dated July 2026 [official] https://deepmind.google/models/model-cards/gemini-3-5-flash-lite/. Pricing: **$0.30 / 1M input, $2.50 / 1M output** [official]. Benchmarks (vendor-reported): SWE-Bench Pro 54.2%, Terminal-Bench 2.1 54.0%, MLE-Bench 39.2%, OSWorld-Verified 74.0%, CharXiv 74.5%/76.5%, GDM-MRCR v2 72.2% (128k) / 21.3% (1M) [official].
- **Gemini Omni — May 19, 2026 (I/O)**: new multimodal-generation model — video output from prompts with refinement steps. First release **Gemini Omni Flash** available to Google AI subscribers via Gemini app and YouTube Shorts. [secondary] https://blockchain.news/news/google-io-2026-key-announcements
- **Gap:** no in-window Gemini 3.x Pro-tier release after 3.1 Pro (Feb 19) was found; a "Gemini 3.5 Pro expected ~June 2026" was predicted by one secondary timeline but no launch was confirmed. [research finding]

**Google API pricing summary (per 1M tokens, Google's own tables):**

| Model | Input | Output | Provenance |
|---|---|---|---|
| Gemini 3.1 Pro | $2.00 | $12.00 | [official] |
| Gemini 3.1 Flash-Lite | $0.25 | (n/a) | [secondary] |
| Gemini 3.5 Flash | $1.50 | $9.00 | [official] |
| Gemini 3.6 Flash | $1.50 | $7.50 | [official] |
| Gemini 3.5 Flash-Lite | $0.30 | $2.50 | [official] |

**Gemini 3 benchmark scoreboard (vendor-reported, official July 2026 model cards; do not mix with independent evals):**

| Benchmark | 3.6 Flash | 3.5 Flash | 3.1 Pro |
|---|---|---|---|
| SWE-Bench Pro | 58.7% | 55.1% | 54.2% |
| DeepSWE v1.1 | 49% | 37% | 12% |
| Terminal-Bench 2.1 | 78.0% | 76.2% | 73.8% |
| MLE-Bench | 63.9% | 49.7% | 42.6% |
| GDPVal-AA v2 (Elo) | 1421 | 1349 | 965 |
| OSWorld-Verified | 83.0% | 78.4% | 76.2% |
| CharXiv Reasoning (no tools / with tools) | 85.2% / 89.4% | 84.2% / 84.9% | 83.3% / 83.2% |
| GDM-MRCR v2 128k avg | 91.8% | 77.3% | 84.9% |
| GDM-MRCR v2 1M pointwise | 54.0% | 26.6% | 26.3% |

Per Google's own table, 3.6 Flash led OSWorld-Verified and CharXiv among all models; trailed GPT-5.6 Luna / Grok 4.5 / Claude Sonnet 5 on SWE-Bench Pro, DeepSWE, Terminal-Bench 2.1, MLE-Bench, and GDPVal-AA — cross-vendor cells are Google's evaluations [official].

