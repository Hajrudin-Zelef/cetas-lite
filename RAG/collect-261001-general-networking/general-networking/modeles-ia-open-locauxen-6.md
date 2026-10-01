---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-6
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States"]
dates: ["2025-04", "2025-04-05", "2026-01", "2026-04", "2026-04-08", "2026-07", "2026-07-09", "2026-07-31", "2026-08", "2026-08-05", "2026-08-10", "2026-08-31", "2026-09", "2026-09-02", "2026-09-04", "2026-09-08"]
keywords: ["agent", "agentic", "amd", "apache", "attention", "attribution", "aws", "bedrock", "benchmark", "benchmarks", "capex", "claude"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [338, 392]
sha256: 9103f8f3f5506aac090b2fc391db3c8f928ca84f9d3230ca429b571158ca7b8a
---

# ÉTAPE 1 — Open / Local AI Models (EN)

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

1. **April 2026: Llama brand retired** — Meta pivots from open weights to the closed, API-monetized **Muse Spark** line (first proprietary Meta model ever). "Meta's last open release was Llama 4, and Llama 4's adoption was 'disappointing' by Meta's own admission."
2. **July 2026: Meta Model API goes public/paid** — Meta joins the metered-API revenue model it once undercut, with an aggressive Contributor tier ($0.10/$0.20) that trades price for training-data rights.
3. **August 10, 2026: the reversal** — **Muse Glimmer 30B under Apache 2.0** (genuinely OSI-approved, no MAU carve-out) plus a promise to open-source **Muse Spark 1.2's** weights. The strategy is now **two-track: closed API for enterprise revenue, open weights for ecosystem leverage** — and Zuckerberg's accompanying essay reframes open weights as a US-competitiveness/national-security argument.
4. **Infrastructure scale-up underpins it:** ~$14B Scale AI investment, up to **$60B AMD chip deal** (Feb 2026), **$115–135B AI capex** reported around earnings (single weak source; treat as indicative).

---

# PART 2 — META MUSE SPARK FAMILY

*(Merged from two independent research passes; discrepancies across outlets are flagged inline rather than smoothed over.)*

## 2.1 What it is

**Muse Spark** is Meta's proprietary frontier reasoning-model family and the **first product of Meta Superintelligence Labs (MSL)**, the division led by Chief AI Officer **Alexandr Wang** (joined Meta from Scale AI in 2025; Meta invested ~$14–14.3B in Scale AI to bring him over). Debuted **April 8, 2026**. Internally codenamed **"Avocado"**.

It is a **natively multimodal reasoning model**: text, image, video, PDF, audio in; text out. Built around long-horizon agentic work: tool use, computer use, MCP servers, custom skills, multi-agent orchestration (main agent delegating to parallel subagents, or running as a subagent). Meta says it was built after a **9-month rebuild of its AI stack "from scratch"** (new infrastructure, new architecture, new data pipelines).

## 2.2 Release timeline (through Sept 2026)

| Version | Date | Notable |
|---|---|---|
| **Muse Spark 1.0** | 2026-04-08 | First MSL model; Meta AI app/meta.ai free tier; API invite-only |
| **Muse Spark 1.1** | 2026-07-09 | First **paid Meta Model API** (public preview, US developers at launch); 1M context; Thinking mode in Meta AI app |
| **Muse Spark 1.2** | 2026-08-05 | Co-released with **Muse Code** terminal coding agent (co-trained with it); global public-preview API expansion; OpenRouter listing; API out of beta 2026-08-31 (subscription option added) |
| **Muse Glimmer 30B** (sibling, open weights) | ~2026-08-10 | 30B dense local agentic model, Apache 2.0, HF: `meta-models/Muse-Glimmer-30B` |
| **Muse Spark 1.3** | 2026-09-02 | Coding/long-context flagship; two variants — `muse-spark-1.3` (xhigh, public), `muse-spark-1.3-max` (limited partner preview); reasoning_effort up to "max" (live 2026-09-04); paid API opened Sept 3 |
| **"Muse" consumer app** | 2026-09-08 | Personal agent app running Muse Spark 1.3 on per-user Meta-hosted VMs; free / $20/mo / $100/mo tiers |

**Meta Model API** (Meta's first metered inference API) public + paid since **July 9, 2026**.

## 2.3 Specs

