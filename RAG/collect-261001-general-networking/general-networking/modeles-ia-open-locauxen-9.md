---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-9
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Meta", "Mistral", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2025-05", "2025-06", "2025-06-10", "2025-12-02", "2026-03-16", "2026-05"]
keywords: ["acquisition", "agent", "agentic", "apache", "benchmarks", "blackwell", "claude", "context window", "cost", "embeddings", "fine-tuning", "gemini"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [527, 582]
sha256: c9efe74660ef200d13738c9a07738f69722e3dc9b1d28f1cf2d8857a48025e49
---

# ÉTAPE 1 — Open / Local AI Models (EN)

## 4.1 Mistral Large 3 — flagship open-weight MoE
- **Release:** December 2, 2025. 10-model "Mistral 3" release.
- **Parameters/architecture:** 675B total / 41B active per token, sparse MoE. Trained from scratch on ~3,000 NVIDIA H200 GPUs.
- **Context window:** 256K–262K tokens. Multimodal (text + vision), multilingual.
- **License:** Apache 2.0.
- **HF weights:** `mistralai/Mistral-Large-3-675B-Instruct-2512`. API id `mistral-large-latest` / `mistral-large-2512`. Runs on a single 8×A100/H100 node with vLLM.
- **Benchmarks:** LMArena debut at **#2 open-source non-reasoning model**. MMLU ≈ 85.5% (8-language); MMLU-Pro 73.11%; MATH-500 93.60%; HumanEval ≈ 92% pass@1; GPQA Diamond ≈ 43.9% (non-reasoning model; strong on broad knowledge, weaker on hard reasoning).
- **API pricing (La Plateforme, May 2026):** $0.50 / $1.50 per 1M input/output tokens.
- **Distinguisher:** largest Apache 2.0 open-weight model in the world at release; frontier non-reasoning quality at ~½ the API cost of closed frontier models; sovereign-European, enterprise-focused positioning.

## 4.2 Mistral Medium 3.5 — flagship "merged" model
- **Release:** April 29/30, 2026.
- **Parameters/architecture:** 128B **dense** transformer + vision encoder trained from scratch (variable image sizes/aspect ratios). Self-hostable on 4 GPUs.
- **Context window:** 256K tokens.
- **License:** **Modified MIT** (a license change worth flagging — several commentators noted Mistral was "loose with the word open source" here; not Apache 2.0).
- **HF weights:** open weights published; API + Le Chat/Vibe. Powers Mistral's Vibe CLI coding agent and Le Chat Work mode.
- **Benchmarks (vendor-reported):** SWE-Bench Verified **77.6%** (ahead of its own Devstral 2 ~72% and Qwen 3.5 397B-A17B ~75%); τ³-Telecom 91.4%; τ³-Airline 72.0; τ³-Retail 76.1.
- **Distinguisher:** Mistral's "first flagship merged model" — it **replaces/retires three lines**: Devstral 2 (coding), Magistral (reasoning), and Medium 3.1. One weight set with configurable per-request `reasoning_effort` (`none` for instant chat → `high` for deep agentic reasoning with `[THINK]` traces).

## 4.3 Mistral Small 4 — efficient unified MoE
- **Release:** March 16, 2026.
- **Parameters/architecture:** 119B total / **~6B active** per token (8B incl. embeddings/output) sparse MoE. Inference stack co-developed with NVIDIA (day-0 NIM, vLLM, SGLang).
- **Context window:** 262K tokens. Multimodal.
- **License:** Apache 2.0.
- **Benchmarks (vendor-reported):** "close to the level of Mistral Medium 3.1 and Mistral Large 3, particularly in MMLU Pro"; Artificial Analysis **LCR (Live Code Reasoning) 0.72** — beats OpenAI's GPT-OSS 120B; Qwen 3.5 122B and Qwen 3-next 80B still beat it on LiveCodeBench. Key claim: shortest instruct-mode outputs of any model tested (2.1K chars vs 14.2K Claude Haiku, 23.6K GPT-OSS 120B) → ~3× throughput vs Mistral Small 3, lower inference cost.
- **Distinguisher:** unifies instruct + reasoning + vision + coding in one open-weight model; accuracy-per-token efficiency play for the 80–120B class.

## 4.4 Devstral 2 — agentic coding model (retired)
- **Release:** December 9–10, 2025. Co-developed with All Hands AI (OpenHands).
- **Parameters:** 123B (plus Devstral Small 2: 24B). Context 262K. License Apache 2.0.
- **Benchmarks:** SWE-Bench ≈ 46.8% (per Red Hat intel doc; another third-party summary cites ~72% SWE-Bench Verified — figures conflict across harnesses, treat as directional).
- **Status:** retired/replaced by Mistral Medium 3.5 (Apr 2026) in the Vibe CLI.

## 4.5 Magistral — reasoning models (retired)
- **Magistral Small (24B):** June 10, 2025; Apache 2.0; open on HF. **Magistral Medium:** enterprise, API/premier access (not open weights).
- **Benchmarks (vendor-reported, June 2025):** Magistral Medium AIME-24 73.6% (90% with majority voting); underperformed Gemini 2.5 Pro and Claude Opus 4 on GPQA Diamond / AIME / LiveCodeBench; 10× faster responses claimed in Le Chat; strong multilingual coverage.
- **API pricing:** Magistral Medium $2/$5 per 1M in/out (2025); catalogs list magistral-medium $2.00/$8.00, magistral-small $0.50/$1.50.
- **Status:** retired into Mistral Medium 3.5 (Apr 2026) for Le Chat reasoning.

## 4.6 Ministral 3 — edge dense models
- **Release:** December 2, 2025. Sizes 3B / 8B / 14B dense; base, chat, and reasoning variants. Context 131K–262K. Apache 2.0, HF open.
- **Benchmarks:** 14B reasoning variant hits **85% on AIME 2025** with far fewer tokens than comparable models.

## 4.7 Other 2026 Mistral open releases (brief)
- **Voxtral TTS** (Mar 23, 2026): open-weight TTS on Ministral 3B backbone; 9 languages, zero-shot voice cloning; **CC BY-NC 4.0** (non-commercial — not permissive). Voxtral Mini Transcribe / Small for speech-to-text, mostly open.
- **Robostral Navigate** (8B, 2026): embodied navigation VLM, single RGB camera, **76.6% on R2R-CE**; sim-trained on 400K trajectories; Apache 2.0 — first Mistral robotics model, part of the Physical AI push (with Emmi AI acquisition, May 2026).
- **Leanstral** (Mar 16, 2026): Mistral Labs release for formal proof engineering (Lean 4).

---

# PART 5 — NVIDIA NEMOTRON (OPEN RELEASES)

**Vendor context:** NVIDIA's own model family, positioned as the open standard for **agentic AI workloads**. Differentiator vs Llama (and vs most open releases): NVIDIA publishes **weights + training data + RL environments + post-training recipes + fine-tuning code** — "fully open, not open in the marketing sense." Note NVIDIA previously built on Meta's weights (Llama-3.1-Nemotron 70B, Llama Nemotron reasoning family Apr/May 2025: Nano 4B / Super 49B / Ultra 253B); the Nemotron 3 generation is original NVIDIA architecture, optimized for NVIDIA GPUs (NVFP4 on Blackwell).

## 5.1 Nemotron 3 family (Nano / Super / Ultra)

