---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-11
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "China", "EU", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Poolside", "United States", "vLLM", "xAI"]
dates: ["2024-12-12", "2025-04-05", "2025-11-20", "2025-12-02", "2025-12-15", "2026-03-11", "2026-03-16", "2026-04-02", "2026-04-29", "2026-06", "2026-06-04", "2026-07-21", "2026-08-10", "2026-09-02"]
keywords: ["agent", "apache", "attention", "awq", "benchmark", "benchmarks", "compute", "dpo", "gguf", "grok", "kimi", "leaderboard"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [643, 709]
sha256: 2e830cb7a14949f0914de35cb836d1b4785fd20f6720a2578b22f85132851bd1
---

# ÉTAPE 1 — Open / Local AI Models (EN)

- **Sizes/variants:** 7B & 32B; Olmo 3-Base (7B, 32B), Olmo 3-Think (7B, 32B — first Ai2 reasoning release), Olmo 3-Instruct (7B). Context 65K (doubled in pre/mid-training; sliding-window attention on 3 of 4 layer types).
- **License:** Apache 2.0. **Fully open** — the key differentiator: Dolma 3 training data (5.9T tokens for 32B), Dolci post-training stack, OlmoRL (SFT→DPO→RLVR), checkpoints, all public. Marketed as "best American-made open source model at this scale" and "best 7B Western instruct/thinking model."
- **Benchmarks:** Olmo 3-Think 32B — MATH ≈ 96.1%; HumanEval+ ≈ 91.4% (marginally ahead of Qwen 3 32B); AIME-style low-to-mid 70s; MMLU mid-80s; competitive with Qwen 3-32B-Thinking on ~1/6 the training tokens. Olmo 3-Think 7B: AIME 2025 70.7% (AA).
- **Positioning:** the transparency-first lab's answer to open-weight models; reproducibility and research use over raw scale.

## 8.2 IBM — Granite 4.0 family (Oct 2025 → 2026)

- **Architecture:** hybrid **Mamba-2 state-space + transformer** ("H" variants) plus traditional dense transformer fallbacks; linear compute scaling, constant memory with context length, >70% RAM reduction claimed on long-context workloads.
- **Releases:** Granite 4.0 core family (Oct 2025): H-Micro/Micro 3B, Small; **Granite 4.0 Nano** (Oct 29, 2025): 350M / 1B (plus transformer-only alternates) — run in-browser via Transformers.js, llama.cpp/vLLM/MLX compatible; **Granite-4.0-3B-Vision** (Mar 27, 2026): document/chart/table extraction VLM (`<chart2csv>`, `<tables_json>` task tags); **Granite 4.0 1B Speech** (Mar 2026): compact multilingual STT/translation (EN/FR/DE/ES/PT/JA + EN↔IT/ZH), two-pass design.
- **License:** Apache 2.0 throughout. First open model family with **ISO/IEC 42001** AI-management certification; cryptographically signed checkpoints; HackerOne bug bounty.
- **Benchmarks (IBM-reported):** Granite 4.0 H 1B — IFEval **78.5** (vs Qwen3 1.7B 73.1, Gemma 3 1B 59.3); Berkeley Function Calling Leaderboard v3 **54.8** (vs Qwen3 52.2, Gemma 3 16.3); Granite 4.0 1B Speech **#1 on OpenASR leaderboard** (avg WER 5.52, RTFx 280; LibriSpeech Clean 1.42). Strong on Stanford HELM.
- **Positioning:** enterprise trust + efficiency play (EY, Lockheed Martin testing); available on watsonx.ai, HF (`ibm-granite`), Docker Hub, Kaggle, NVIDIA NIM, Dell Pro AI Studio.

## 8.3 Nous Research — Hermes 4 (2026 release cycle)

- **What:** community open-weight lab's "hybrid reasoning" family — neutrally aligned (minimal refusals), toggleable `<think>`-tag reasoning, trained on ~5M samples / 19B tokens of newly synthesized reasoning + instruction data (Distilabel pipeline) + RLHF, per the 40-page technical report (arXiv 2508.18255).
- **Sizes/licensing:** reported sizes vary across coverage (8B/14B/35B/70B; base architecture consistent with the Llama 3.x generation); weights + GGUF/AWQ quants on HF under `NousResearch` (e.g., the `hermes-4-collection`); no API gate, no registration.
- **Benchmarks (mixed sourcing — treat as directional):** Hermes 4 35B ≈ 83 on GPQA Diamond per a Delphi Digital industry comparison (vs Qwen 3.5 88.4, Kimi K2.5 87.6); 405B-class variant reported at 96.3% MATH-500 / 81.9% AIME'24 in reasoning mode; top score on Nous's own RefusalBench (57.1% vs GPT-4o 17.67%).
- **Adjacent:** Nous shipped **Hermes Agent** (open-source MIT autonomous CLI agent, persistent memory, GEPA self-improving skill loop — ICLR 2026 oral) — evidence of the lab's agent-first strategy.

## 8.4 OpenAI — GPT-OSS (Aug 2025; relevant baseline)

- **gpt-oss-20b / gpt-oss-120b:** OpenAI's first open-weight release since GPT-2; Apache 2.0; hybrid reasoning (configurable effort); text-only.
- **Benchmarks:** AA Intelligence Index **33** for the 120B (June 2026); beaten by Mistral Small 4 on the AA LCR long-context reasoning benchmark; Nemotron 3 Nano beats it on RULER and throughput (2.2×).
- **Why included:** the US-lab open-weight baseline that 2026 releases are measured against.

## 8.5 Gaps / negative findings

- **Snowflake Arctic:** no 2026 successor found; Arctic 2 (Dec 2024) remains the latest — excluded.
- **Databricks DBRX successors:** nothing found in 2026 — excluded.
- **xAI:** no open-weight release in 2026 found; Grok line remains closed — excluded.
- **Apertus (ETH Zurich/EPF Lausanne, Sep 2025, 70B, Apache 2.0)** — first EU-AI-Act-compliant LLM, fully open; Swiss/European but worth a cross-reference for the "sovereign open" theme.

---

# PART 9 — MASTER COMPARISON TABLE (Sept 2026 snapshot)

| Model | Vendor | Release | Params | Context | License | HF weights | AA Intelligence Index | SWE-bench Verified | API in/out per 1M |
|---|---|---|---|---|---|---|---|---|---|
| Llama 4 Maverick | Meta | 2025-04-05 | 400B / 17B act | 1M | Community License | ✅ | 9.3 | 21.0% | $0.20 / $0.80 |
| Llama 4 Scout | Meta | 2025-04-05 | 109B / 17B act | 10M | Community License | ✅ | 6.5 | — | $0.19 / $0.68 |
| Muse Spark 1.3 | Meta | 2026-09-02 | undisclosed | 1M | Proprietary (API-only) | ❌ | 61 (xhigh) / 62 (max) | — | $1.25 / $4.25 |
| Muse Glimmer 30B | Meta | 2026-08-10 | 30B dense | 120K+ | Apache 2.0 | ✅ | — | 76% (vendor) | n/a (local) |
| Laguna S 2.1 | Poolside | 2026-07-21 | 118B / 8B act | 1M | OpenMDW-1.1 | ✅ | — (coding) | 59.4% Pro (vendor) | $0.10 / $0.20 |
| Mistral Large 3 | Mistral | 2025-12-02 | 675B / 41B act | 256K | Apache 2.0 | ✅ | — | — | $0.50 / $1.50 |
| Mistral Medium 3.5 | Mistral | 2026-04-29 | 128B dense | 256K | Modified MIT | ✅ | — | 77.6% (vendor) | via La Plateforme |
| Mistral Small 4 | Mistral | 2026-03-16 | 119B / 6B act | 262K | Apache 2.0 | ✅ | — | — | via La Plateforme |
| Nemotron 3 Ultra | NVIDIA | 2026-06-04 | 550B / 55B act | 1M | OpenMDW-1.1 | ✅ | **48** (top US open) | 70.7 (vendor) | $0.50 / $2.50 |
| Nemotron 3 Super | NVIDIA | 2026-03-11 | 120B / 12B act | 1M | NVIDIA Open Model License | ✅ | 36 | — | via NIM |
| Nemotron 3 Nano | NVIDIA | 2025-12-15 | 31.6B / 3.2B act | 1M | NVIDIA Open Model License | ✅ | — | — | via NIM |
| Gemma 4 31B | Google | 2026-04-02 | 30.7B dense | 256K | Apache 2.0 | ✅ | 39 | ~54% (vals.ai subset) | Free (AI Studio) |
| Gemma 4 26B-A4B | Google | 2026-04-02 | 25.2B / 3.8B act | 256K | Apache 2.0 | ✅ | — | — | Free (AI Studio) |
| OLMo 3 Think 32B | Ai2 | 2025-11-20 | 32B | 65K | Apache 2.0 | ✅ | — | — | self-host |
| Granite 4.0 H 1B | IBM | 2025-10 | 1B | long | Apache 2.0 | ✅ | — | — | self-host |
| Phi-4 | Microsoft | 2024-12-12 | 14B dense | 16K | MIT | ✅ | — | — | self-host |
| gpt-oss-120b | OpenAI | 2025-08 | 120B | 128K | Apache 2.0 | ✅ | 33 | — | self-host |

*AA Intelligence Index values are from AA v4.x snapshots (methodology-version dependent); vendor benchmarks are marked where relevant.*

---

# PART 10 — CROSS-CUTTING OBSERVATIONS (for the RAG)

## 10.1 Meta vs Chinese open-weights (2026)

The open-weight frontier in 2026 is **Chinese-led**, and Meta is no longer in the flagship race:

