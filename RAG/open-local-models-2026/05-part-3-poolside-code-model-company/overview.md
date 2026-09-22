---
id: open-local-models-2026/05-part-3-poolside-code-model-company/overview
title: "PART 3 — POOLSIDE — CODE-MODEL COMPANY"
domain: part-3-poolside-code-model-company
role: deep-dive
task: reference
actors: ["Anthropic", "Apple", "Baseten", "China", "CoreWeave", "Hugging Face", "Meta", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "SGLang", "United States", "vLLM"]
dates: ["2024-06", "2025-05", "2026-04-28", "2026-05-22", "2026-07-02", "2026-07-09", "2026-07-21", "2026-08-21", "2026-09-02"]
keywords: ["acquisition", "agent", "agentic", "agents", "apache", "attention", "benchmarks", "claude", "compute", "cost", "datacenter", "funding"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [461, 522]
section: "PART 3 — POOLSIDE — CODE-MODEL COMPANY"
sha256: 741957757a9df1e2a0d7507ce62af33c6434777458656e2aa102bb3283e24f1c
---

# PART 3 — POOLSIDE — CODE-MODEL COMPANY

## 3.1 Company & business model

- **Founded 2023** by **Jason Warner** (ex-GitHub CTO) and **Eiso Kant** (ex-Athenian, acquired by Linux Foundation 2023). Paris-based (moved HQ from US to Paris ~Aug 2024).
- **Business model:** enterprises pay to run powerful coding models **on their own hardware** (self-hosted/VPC, inside the customer's security boundary, "no information leaves") rather than sending prompts to closed APIs. Core customers: government, defense, highly regulated industries; pitch framed around Western enterprises/governments avoiding data transfers to overseas vendors.
- **Model Factory:** internal platform automating architecture search + **RL from code execution feedback (RLCEF)** — models "improve by completing millions of tasks in tens of thousands of real-world software projects." Enables ~5-week release cadence and ~60-day train-to-release cycles.

## 3.2 Funding timeline

- ~2024: $126M initial raise (BCV-led, with Redpoint, Felicis et al.) around Paris HQ move; June 2024 reports of a $400M raise at a $2B valuation (TechCrunch via SiliconANGLE).
- **Oct 2024: $500M Series B at $3B valuation** (Nvidia, eBay participated).
- **Apr 2026: planned $2B Series C at $14B valuation fell through** after CoreWeave withdrew from a joint Texas datacenter project.
- **2026-08-21: NVIDIA deal — $6B non-exclusive license for the Model Factory + $1B investment at $12B pre-money valuation; 109 employees (most of technical staff; "fewer than 115 people built the model" per Kant) hired by NVIDIA; founders stay; described as neither acquisition nor acquihire.** Poolside's investor letter: missed a 6-week window to raise $2B for a 40,000-GB300 cluster, lost the cluster; frontier training now constrained by datacenter space and contracted compute. Infra arm (PIC) spun out Jan 2026, building a 1.2GW Texas datacenter toward 7GW.

## 3.3 Model lineup (2026: the Laguna family; Malibu = legacy)

**Legacy: Malibu / Point (pre-2026, enterprise-only, now superseded)**
- **Malibu** (as of May 2025): flagship model for deep reasoning via chat — code generation, test writing, refactoring, documentation; **>1M token context**; built from scratch, fine-tuned on each customer's environment via RLCEF, deployed inside customer security boundaries.
- **Point**: real-time code completion model, ~100K context.
- **Malibu 2.2** (referenced in partner materials): dense architecture, 128K context, enterprise license, for low-latency IDE completion/inline editing.
- Status in 2026: no Malibu family on poolside.ai/models or the blog — treated as **legacy/superseded by Laguna**.

**Laguna XS.2** — released **2026-04-28**. 33.4B total / 3B active (per one guide; others say undisclosed), 256K context, Apache 2.0, single-GPU/Apple Silicon friendly. SWE-Bench Verified **64%** (vendor). Superseded by XS 2.1 (retired from API/OpenRouter 2026-07-09).

**Laguna M.1** — released **2026-04-28**. 225B total / 23B active MoE, 262,144 (256K) context, enterprise availability. Benchmarks: SWE-Bench Verified **65.4%** (vendor; one blog claims 72.5% for M.1 — flagged), SWE-Bench Multilingual **63.1%**, SWE-Bench Pro (public) **49.2%**, Terminal-Bench 2.0 **44.9%**.

**Laguna XS 2.1** — released **2026-07-02**. 33B total / 3B active MoE, 256K context, **OpenMDW-1.1** license, free on HF + free OpenRouter tier; runs on a single GPU. SWE-Bench Multilingual **63.1%** (+5.4pp vs XS.2), Terminal-Bench 2.0 **37.5%**.

**Laguna S 2.1** — released **2026-07-21**. Flagship; "the West's most capable open-weight model" (GlobeNewswire release).
- **Params/architecture:** 118B total / **8B active** MoE — 256 routed experts + 1 shared expert, grouped-query attention, interleaved sliding-window layers (per HF model card). Trained via Model Factory on 4,000 Nvidia H200 GPUs in <4 weeks; training began 2026-05-22 → 60-day cycle.
- **Context:** 1,048,576 (1M) tokens. Inference scales with the 8B active params; runs on a single **Nvidia DGX Spark** (desktop AI machine); official 4-bit Apple MLX weights (71.9GB download — real target: 128GB Mac); 4-bit GGUF ~75GB.
- **License:** **OpenMDW-1.1** — a new permissive license for model weights backed by the Linux Foundation; commercial use, modification, redistribution allowed.
- **Hugging Face:** `huggingface.co/poolside/Laguna-S-2.1` — open weights, day one. Also official Apple MLX weights.
- **Benchmarks (vendor, run with Poolside's `pool` agent harness, reproducible):** Terminal-Bench 2.1 **70.2%**; SWE-Bench Multilingual **78.5%**; SWE-Bench Pro (public) **59.4%**; DeepSWE v1.1 **40.4%**; Toolathlon Verified **49.7%**. Comparables: Tencent Hy3 71.7% TB2.1; Nemotron 3 Ultra 56.4%; Claude Sonnet 5 80.4% / 63.2% Pro; GPT-5.6 Luna Max 82.5% TB.
- **Availability & pricing:** HF download; **OpenRouter** `poolside/laguna-s-2.1` (paid, full 1M context) at **$0.10 / 1M input, $0.20 / 1M output** — an order of magnitude below frontier pricing; free tier `poolside/laguna-s-2.1:free` (capped 262,144 context), plus free XS 2.1 / M.1 routes; Baseten model library; Vercel AI Gateway; vLLM/SGLang/Ollama/llama.cpp support; free demo chat at chat.poolside.ai (no login); platform.poolside.ai API.

## 3.4 What distinguishes Poolside

- Only Western lab shipping frontier-adjacent **open-weight agentic coding models** at this cadence (positioned explicitly against Chinese open-weight dominance and US closed APIs).
- "Token economics" thesis: sparse MoE means inference cost follows active params; on hardest benchmarks the model burns a mean ~249K completion tokens/trajectory in thinking mode — self-hosting, not per-token API, is the economic answer for enterprise agents.
- Behavioral training emphasis ("working habits": verification, persistence) over raw scale.
- Open-ended future post-NVIDIA deal: founders pivoting the remaining company toward an undisclosed direction; the thesis that intelligence-bound problems commoditize while experiment-bound problems hold value ("the world's most valuable scientific discovery engine").

## 3.5 Side-by-side (Sept 2026 snapshot)

| | Muse Spark 1.3 (Meta) | Laguna S 2.1 (Poolside) |
|---|---|---|
| Release | 2026-09-02 | 2026-07-21 |
| Params | Undisclosed (proprietary) | 118B total / 8B active MoE |
| Context | 1M | 1M |
| Weights/license | Proprietary, API-only | Open weights, OpenMDW-1.1 (commercial OK) |
| HF | No | Yes (`poolside/Laguna-S-2.1`) |
| AA Intelligence Index | 61 (xhigh) / 62 (max) | n/a (coding-specialist) |
| Terminal-Bench 2.x | 88.8% (2.1) | 70.2% (2.1) |
| DeepSWE v1.1 | 75.4% | 40.4% |
| API price in/out per 1M | $1.25 / $4.25 | $0.10 / $0.20 (OpenRouter paid) |
| Channel | Meta AI app, WhatsApp/IG/FB, Meta Model API, Muse Code, OpenRouter | Self-host, HF, OpenRouter (free+paid), Baseten, Vercel AI Gateway |
| Audience | Consumers + developers/enterprise | Developers, enterprises, governments (self-host) |

---

