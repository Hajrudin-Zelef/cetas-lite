---
id: labs-hyperscalers-2026/00-labs-hyperscalers/16-3-api-pricing-comparison-flagship-frontier-models-per-1m-
title: "16.3 API pricing comparison — flagship & frontier models (per 1M tokens, input/output; as reported in-window)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: model-release
actors: ["Anthropic", "Cohere", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Sakana", "xAI"]
dates: ["2026-05-19"]
keywords: ["pricing", "agentic", "apache", "astra", "benchmark", "chatgpt", "claude", "cohere", "consumer", "cost", "deflation", "fable 5"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [2196, 2311]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 80581ab5dabf1ee3de7f7b6c15cd990ab8ac8dbeaa5213493862615869c58540
---

# 16.3 API pricing comparison — flagship & frontier models (per 1M tokens, input/output; as reported in-window)

### 16.3 API pricing comparison — flagship & frontier models (per 1M tokens, input/output; as reported in-window)

> Prices move frequently; all figures dated to their reporting. Cached-input and batch discounts omitted except where noted. [secondary] compilations unless marked [official]/[vendor-reported].

**Anthropic (Claude)**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Claude Opus 4.6 | $5.00 | $25.00 | 1M (beta) | Raised to $8.00 input Aug 6, 2026 |
| Claude Sonnet 4.6 | $3.00 | $15.00 | 1M | Raised to $4.00 input Aug 6, 2026 |
| Claude Fable 5 | — | — | — | Pricing not returned in-window (gap) |
| Claude Cowork (product) | $20–100/mo seats | — | — | Agentic OS tool, Apr 9 2026 |

**OpenAI**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| GPT-5.3-Codex | $1.75 | $14.00 | 400K | Feb 5, 2026 |
| GPT-5.4 / Thinking / Pro | $2.50 | $15.00 | 1M | Mar 5, 2026 |
| GPT-5.5 | $5.00 | $30.00 | 1M | Apr 22, 2026 |
| GPT-5.6 Sol | $5.00 | $30.00 | 1.05M | Jul 9, 2026 GA |
| GPT-5.6 Terra | $2.50 | $15.00 | 1.05M | Jul 9, 2026 GA (later cut) |
| GPT-5.6 Luna | $1.00 → **$0.20** | $6.00 | 1.05M | 80% input cut Jul 30, 2026 |
| GPT-6 Astra | $10.00 | $50.00 | 1M | Sep 3, 2026 limited preview |
| GPT-Rosalind | — | — | — | Life-sciences; pricing n/r (gap) |
| Codex (product) | Token-based billing from Apr 2, 2026; pay-as-you-go Codex-only seats from Apr 3 | — | — | — |
| ChatGPT tiers | Free (ads) / Go $8/mo / Plus / Pro $100 (Apr 9) / Pro $200 closed Sep 10 / Team→Business $20/seat/mo | — | — | — |

**Google (Gemini)**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Gemini 3.1 Pro | $2.00 | $12.00 | 1M | Feb 19, 2026 [official model-card table] |
| Gemini 3.1 Flash-Lite | $0.25 | — | — | ~2.5× faster; Feb 19, 2026 |
| Gemini 3.5 Flash | $1.50 | $9.00 | 1M (64K out) | May 19, 2026 [official] |
| Gemini 3.6 Flash | $1.50 | $7.50 | 1M (64K out) | Model card Jul 21, 2026 [official] |

**Meta (Muse)**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Muse Spark 1.1+ API | — | — | — | Meta's first paid API (Jul 9, 2026); list pricing n/r (gap) |
| Muse Glimmer 30B | open weights (Apache 2.0) | — | — | Self-host; Aug 10, 2026 |

**xAI (Grok)**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Grok 4.20 ("4.2") | $1.25 | $2.50 | 256K | Feb 17, 2026 |
| Grok 4.3 | $1.25 | $2.50 ($0.20 cached) | 1M | Apr 30, 2026; output cut −83% at launch |
| Grok 4.5 | $2.00 | $6.00 | 500K | Jul 8, 2026 |
| Grok 4.6 | $2.00 | $6.00 | 500K | Aug 12, 2026; cached input $0.03→$0.05 |
| Grok 4.7 | $2.00 | $6.00 | 500K | Sep 21, 2026 |
| Grok TTS API | $4.20 / 1M chars | — | — | Realtime voice $3/hour |
| SuperGrok / Heavy | $30/mo / $300/mo | — | — | Consumer tiers |

**Mistral**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Mistral Small 4 | $0.15 | $0.60 | 256K | Mar 16, 2026; Apache 2.0 |
| Mistral Medium 3 | ~$0.40 | ~$2.00 | — | [secondary] compilation |
| Mistral Medium 3.5 | — | — | 256K | Pricing n/r (gap) |
| Mistral Large 3 | $2.00 | $6.00 | — | Dec 2025; [secondary] compilation |
| Codestral | $0.30 | $0.90 | — | [secondary] |
| Mistral Nemo | $0.02 | $0.10 | — | [secondary] |
| Ministral 3B | $0.04 | $0.04 | — | [secondary] |
| Voxtral Transcribe | ~$0.003/min | — | — | [secondary] |
| Voxtral TTS | $16.00/1M chars | — | — | [secondary] |
| OCR 4 | $4/1,000 pages ($2 batch; $5 Document AI) | — | — | [secondary] |
| Le Chat → Vibe | Free / €14.99 Pro / €24.99 Team (€19.99 annual) / Enterprise | — | — | [independent] |

**Independent labs**
| Model | Input | Output | Context | Notes |
|---|---|---|---|---|
| Cohere Command A+ | — | — | 128K in / 64K out | API pricing n/r (gap); Apache 2.0 weights |
| Sakana Fugu Ultra (orig) | $5.00 | $30.00 | — | $0.50/M cached; Jun 22, 2026 |
| Sakana Fugu Max | $2.00 | $6.00 | 1M | Fixed regardless of length; Sep 11–12, 2026 |
| Sakana Fugu Ultra v2 | $5.00 | $30.00 | — | $0.50/M cached; $10/$45/$1.00 above 272K |
| AI21 Jamba Mini 2 | — | — | 256K | Pricing n/r in this pass (gap); Apache 2.0 |
| Thinking Machines Inkling | free on Tinker at launch | — | 1M | Per-token list pricing n/r (gap) |

**Price-per-intelligence observations (as reported, not independent verification):**
- Artificial Analysis put Grok 4.3 on the Pareto frontier of cost-per-intelligence ($395 benchmark-suite run, ~20% below Grok 4.20). [independent]
- Grok 4.5: $0.31–$0.49 per completed Intelligence-Index task; ~14k output tokens/task vs ~67k for Opus 4.8 (token-efficiency pitch). [independent/vendor-reported]
- Gemini 3.5 Flash: ~40% below Gemini 3.1 Pro on input pricing; positioned for agentic workloads at ~4× output tok/s. [vendor-reported]
- OpenAI's Luna 80% input cut (Jul 30, 2026) and Grok 4.3's 83% output cut (Apr 30, 2026) show frontier API deflation through 2026. [vendor-reported]

### 16.4 Model specification matrix — flagship releases Feb–Sep 2026

| Model | Lab | Release | Params (total/active) | Context | Weights license | Input modalities |
|---|---|---|---|---|---|---|
| Claude Opus 4.6 | Anthropic | Feb 5, 2026 | n/d | 1M (beta) | Closed | text/image/code |
| Claude Sonnet 4.6 | Anthropic | Mar 26, 2026 | n/d | 1M | Closed | text/image/code |
| Claude Fable 5 | Anthropic | 2026 (in-window) | n/d | n/d | Closed | text/code |
| GPT-5.3-Codex | OpenAI | Feb 5, 2026 | n/d | 400K | Closed | text/code |
| GPT-5.5 | OpenAI | Apr 22, 2026 | n/d | 1M | Closed | text/image/code |
| GPT-5.6 Sol/Terra/Luna | OpenAI | Jul 9, 2026 | n/d | 1.05M | Closed | text/image/code |
| GPT-6 Astra | OpenAI | Sep 3, 2026 | n/d | 1M | Closed | text/code (+ tools) |
| GPT-Rosalind | OpenAI | Apr 16, 2026 | n/d | n/d | Closed | text (life-sci) |
| Gemini 3.1 Pro | Google | Feb 19, 2026 | n/d | 1M | Closed | text/image/audio/video/code |
| Gemini 3.5 Flash | Google | May 19, 2026 | n/d | 1M | Closed | text/image/video/audio/PDF |
| Gemini 3.6 Flash | Google | Jul 21, 2026 (card) | n/d | 1M | Closed | multimodal |
| Muse Spark 1.3 | Meta | Sep 2, 2026 | n/d | n/d | Closed (paid API) | text/code |
| Muse Glimmer 30B | Meta | Aug 10, 2026 | 30B | n/d | Apache 2.0 | text |
| MAI-Thinking-1 | Microsoft | Jun 2, 2026 | 35B active | 256K | Closed (Foundry) | text |
| MAI-Code-1-Flash | Microsoft | Jun 2, 2026 | 5B | n/d | Closed | code |
| Phi-4-reasoning-vision-15B | Microsoft | Mar 4–5, 2026 | 15B | n/d | Open weights | text/image |
| Grok 4.5 | xAI | Jul 8, 2026 | 1.5T MoE ("V9") | 500K | Closed | text/code |
| Grok 4.7 | xAI | Sep 21, 2026 | 2.1T | 500K | Closed | text/code |
| Mistral Small 4 | Mistral | Mar 16, 2026 | 119B / ~6–6.5B | 256K | Apache 2.0 | text/image in |
| Mistral Medium 3.5 | Mistral | Apr 28–29, 2026 | 128B dense | 256K | Modified MIT | text |
| Mistral Large 3 | Mistral | Dec 2, 2025 | 675B / 41B | n/d | Apache 2.0 | text |
| Inkling | Thinking Machines | Jul 15, 2026 | 975B / 41B | 1M | Apache 2.0 | text/image/audio |
| Inkling-Small (preview) | Thinking Machines | Jul 15, 2026 | 276B / 12B | n/d | Planned open | text/image/audio |
| Command A+ | Cohere | May 20–21, 2026 | 218B / 25B | 128K/64K | Apache 2.0 | text/image |
| Tiny Aya | Cohere | Feb 2026 | 3.35B | n/d | Open weights | text (70+ langs) |
| Nemotron 3 Nano/Super/Ultra | NVIDIA | 2026 | 30B / 100B / 500B | 1M (some) | NVIDIA Open Model License | text/code |
| Jamba Mini 2 | AI21 | Jan 2026 | 52B / 12B | 256K | Apache 2.0 | text |

n/d = not disclosed. Closed labs (Anthropic, OpenAI, Google, xAI) disclose no parameter counts for 2026 flagships; open-weight labs do.

