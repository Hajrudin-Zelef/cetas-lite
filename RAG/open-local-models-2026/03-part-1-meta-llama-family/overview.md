---
id: open-local-models-2026/03-part-1-meta-llama-family/overview
title: "PART 1 — META LLAMA FAMILY"
domain: part-1-meta-llama-family
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter"]
dates: ["2025-04", "2025-04-05", "2026-01", "2026-04-08", "2026-07-09", "2026-07-31", "2026-08", "2026-08-05", "2026-08-10", "2026-09", "2026-09-02", "2026-09-08"]
keywords: ["llama", "agent", "apache", "attention", "attribution", "aws", "bedrock", "benchmark", "benchmarks", "claude", "consumer", "context window"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [287, 360]
section: "PART 1 — META LLAMA FAMILY"
sha256: 209532ffbc83e10a94ca4b8c99f2c0c4430baee1f545c35c0a2c2bd0eb00f20c
---

# PART 1 — META LLAMA FAMILY

## 1.1 Timeline overview (through September 2026)

| Date | Event |
|---|---|
| **2025-04-05** | Meta releases **Llama 4 Scout** and **Llama 4 Maverick** — first natively-multimodal, open-weight MoE models in the Llama family. Llama 4 **Behemoth** (288B-active teacher) announced as still training. |
| 2025-04 | **Benchmark controversy**: the LMArena entry `Llama-4-Maverick-03-26-Experimental` turns out to be an unreleased variant scoring far above the public weights; later publicly acknowledged. |
| 2025-06 | Meta forms **Meta Superintelligence Labs (MSL)** under Chief AI Officer **Alexandr Wang**, following the ~$14B Scale AI investment. |
| **2026-01** | Departing chief AI scientist Yann LeCun publicly acknowledges the Llama 4 benchmark manipulation. |
| **2026-04-08** | **Muse Spark** launches — first model from MSL, **closed-weight/proprietary** (Meta's first ever proprietary model), replacing the Llama brand. Internally codenamed **"Avocado"**. |
| 2026-07-09 | **Meta Model API** goes public and paid — Meta's first metered inference API. |
| 2026-07 | **Muse Spark 1.1** ships. |
| **2026-08-05** | **Muse Spark 1.2** + **Muse Code** (terminal coding agent, macOS/Linux beta). |
| **2026-08-10** | **Muse Glimmer 30B** — open weights under **Apache 2.0**; Meta announces **Muse Spark 1.2's weights will also be open-sourced**. Zuckerberg publishes an essay framing open weights as a safeguard against AI concentration. |
| 2026-09-02 | **Muse Spark 1.3** ships (xhigh public, max partner preview); paid API opened Sept 3. |
| 2026-09-08 | **"Muse" consumer app** launches — personal agent app on per-user Meta-hosted VMs. |
| Sept 2026 | Current frontier: **muse-spark-1.3**. **No Llama 5 exists** — the newest Meta model on OpenRouter is still `meta-llama/llama-4-maverick` (Apr 2025). Behemoth was never released. |

## 1.2 Llama 4 Scout

**Release date:** April 5, 2025

| Attribute | Value |
|---|---|
| Architecture | Mixture-of-Experts, natively multimodal (early fusion; text+image in) |
| Parameters | **109B total / 17B active per token, 16 experts** |
| Context window | **10M tokens** — largest open-weight context at release |
| Training data | ~40T tokens, text-focused corpus |
| Hardware | Fits on a single NVIDIA H100 (FP8) |
| License | **Llama 4 Community License** (gated; accept on Hugging Face) |
| HF weights | `meta-llama/Llama-4-Scout-17B-16E-Instruct` (+ `-Instruct-Original` native format, base `-17B-16E`); NVIDIA FP8/FP4 builds `nvidia/Llama-4-Scout-17B-16E-Instruct-FP8` |

**Benchmarks (Meta-reported, April 2025):** MMLU **79.6%**, MMLU Pro **74.3%**, GPQA Diamond **57.2%**, MATH **50.3%**, LiveCodeBench **32.8%**, MMMU **69.4%**, MGSM **90.6%**.
**Benchmarks (Artificial Analysis, independent):** Intelligence Index v4.3 **6.5** (earlier AA estimate **8**); GPQA Diamond **58.7%**, SciCode **21.3%**, Terminal-Bench 2.1 **3.7%**, AIME 2025 **14%**. Output speed **100.1 tok/s** median, TTFT **0.83s**.
**API pricing:** AA median across providers **$0.19 / $0.68** per 1M in/out tokens; OpenRouter-range ≈ **$0.15–$0.20 in / $0.50–$0.80 out**.
**What distinguishes it:** the 10M-token context window in an open-weight model (∼78× GPT-4o's 128K), single-H100 deployability, and best-in-class multimodal quality at its size tier (beat Gemma 3, Gemini 2.0 Flash-Lite, Mistral 3.1 per Meta).

## 1.3 Llama 4 Maverick

**Release date:** April 5, 2025

| Attribute | Value |
|---|---|
| Architecture | MoE, natively multimodal (early fusion); distilled partly from Behemoth |
| Parameters | **~400B total / 17B active per token, 128 experts** |
| Context window | **1M tokens** |
| Training data | ~22T tokens, multimodal incl. Meta content |
| License | **Llama 4 Community License** |
| HF weights | `meta-llama/Llama-4-Maverick-17B-128E-Instruct` (+ `-Original` native); Ollama `llama4:maverick` |

**Benchmarks (Meta-reported, April 2025):** MMLU **85.5%**, MMLU Pro **80.5%**, GPQA Diamond **69.8%**, MATH **61.2%**, LiveCodeBench **43.4%**, MMMU **73.4%**; MBPP pass@1 **77.6%**; HumanEval **62.0** (third-party AWS Bedrock table).
**Benchmarks (independent):** AA Intelligence Index v4.3 **9.3** (AA estimate **10**); AA GPQA Diamond **67.1%**, SciCode **31.7%**, Terminal-Bench 2.1 **7.9%**, AIME 2025 **19.3%**. **SWE-bench Verified 21.0%**; **SWE-bench Pro 5.24%** — a weak coding showing. LMArena Elo **1,417** — *but see the benchmark controversy below*. Output speed **111.8 tok/s**, TTFT **0.90s**.
**API pricing:** AA median **$0.26 / $0.91** per 1M in/out; OpenRouter **`meta-llama/llama-4-maverick` $0.20 / $0.80**; cheapest providers from **$0.15 / $0.60**.
**What distinguishes it:** the strongest open-weight multimodal model of its generation (Meta claimed it beat GPT-4o and Gemini 2.0 Flash on broad benchmarks; comparable to DeepSeek V3 on reasoning/coding at <½ the active parameters), 1M context, 200+ languages. Coding was its weak spot vs. Chinese peers — Qwen3.5-122B-A10B beats it by +16.8 pts on GPQA Diamond and +34.6 on LiveCodeBench.

## 1.4 Llama 4 Behemoth — never released

Announced April 5, 2025 as still training: **288B active parameters / 16 experts, ~2T total**, intended as the teacher for Scout/Maverick and Meta's answer to GPT-4.5/Claude 3.7/Gemini 2.0 Pro on STEM. As of September 2026 it has **never been publicly released and was effectively shelved** — mid-training MoE-routing/attention issues at 2T scale; Kimi K2.6 took the "best open-weight flagship on SWE-bench" slot Behemoth was designed for.

## 1.5 License: the Llama Community License (through 2026)

- **Terms:** free for research and commercial use, but with restrictions that make it non-open-source by OSI standards: an **Acceptable Use Policy**, attribution requirements, a **"Llama" naming/trademark** clause, and — crucially — a clause requiring a **separate license from Meta for entities with >700M monthly active users**. The same MAU carve-out has applied since Llama 2.
- Weights are **gated on Hugging Face**: users must accept the license before download. NVIDIA's FP8/FP4 redistributions inherit the same license (no separate HF auth needed).
- **No "Community License 2.0" was issued** — instead, Meta leapfrogged it: **Muse Glimmer (Aug 2026) shipped under plain Apache 2.0** with no MAU carve-out and no AUP, and Meta announced Muse Spark 1.2's weights will follow. This is the real 2026 licensing shift: from Llama Community License to genuine OSI-approved licensing.

## 1.6 Llama 5 — does not exist (as of Sept 2026)

Multiple independent sources confirm **no Llama 5 has been released**: "the newest Meta model in the [OpenRouter] catalog, since there is no Llama 5" (catalog read 2026-07-31); "Llama 5 (600B+, April 8 2026) does not exist… HF meta-llama holds no Llama-5 weights at all"; "As of August 2026, Meta has not released Llama 5". Notably, the internal codename **"Avocado" became Muse Spark, not Llama 5**. Some trackers forecast a possible 2027 Llama 5; consensus put a 2026 ship below 20%.

**The benchmark controversy (relevant context):** Meta submitted an unreleased experimental variant (`Llama-4-Maverick-03-26-Experimental`) to LMArena that scored far above the public model; in January 2026 Yann LeCun publicly acknowledged the manipulation. This damaged Meta's open-weight credibility and is widely cited as a driver of the 2026 strategy pivot.

## 1.7 Meta's 2026 open-weight strategy shifts — summary

