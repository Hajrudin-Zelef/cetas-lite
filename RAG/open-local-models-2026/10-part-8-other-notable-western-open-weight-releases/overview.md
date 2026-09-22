---
id: open-local-models-2026/10-part-8-other-notable-western-open-weight-releases/overview
title: "PART 8 — OTHER NOTABLE WESTERN OPEN-WEIGHT RELEASES"
domain: part-8-other-notable-western-open-weight-releases
role: deep-dive
task: model-release
actors: ["Alibaba", "EU", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "vLLM", "xAI"]
dates: ["2026-06"]
keywords: ["open-weight", "agent", "apache", "attention", "awq", "benchmark", "benchmarks", "compute", "dpo", "gguf", "grok", "kimi"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [639, 677]
section: "PART 8 — OTHER NOTABLE WESTERN OPEN-WEIGHT RELEASES"
sha256: bcfbf8807d8289b22cb436f2b43a4986fb809c4da39369a666499b3c6946a1d4
---

# PART 8 — OTHER NOTABLE WESTERN OPEN-WEIGHT RELEASES

## 8.1 Allen Institute for AI — OLMo 3 (Nov 20, 2025)

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

