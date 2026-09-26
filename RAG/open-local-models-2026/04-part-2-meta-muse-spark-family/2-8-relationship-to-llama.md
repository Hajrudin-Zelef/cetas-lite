---
id: open-local-models-2026/04-part-2-meta-muse-spark-family/2-8-relationship-to-llama
title: "2.8 Relationship to Llama"
domain: part-2-meta-muse-spark-family
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "Meta", "OpenAI"]
dates: []
keywords: ["llama", "agent", "agentic", "agi", "apache", "benchmark", "benchmarks", "consumer", "context window", "distillation", "gguf", "gpu"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [436, 460]
section: "PART 2 — META MUSE SPARK FAMILY"
sha256: 026307022fc0e5837149e6569f6d0775e847f07f60bfb08f20055c37241a4b2b
---

# 2.8 Relationship to Llama

| Attribute | Value |
|---|---|
| Architecture | **Dense 30B** (~29.6B), distilled from Muse Spark 1.2 via logit distillation; dedicated perception encoder; natively multimodal (text+image) |
| Context window | **120K+ tokens** (131K per one source) |
| License | **Apache 2.0** — commercial use, modification, redistribution; no MAU cutoff, no AUP |
| HF weights | `meta-models/Muse-Glimmer-30B` (separate `meta-models` org; +GGUF, -assistant, ExecuTorch-PTE variants) |
| Design goal | Agentic local model: fits **<20 GB** quantized on a single 24GB consumer GPU or Apple Silicon Mac; runs offline; **no first-party Meta API** |

**Benchmarks (Meta vendor-reported):** SWE-Bench Pro **51.2%**, AIME **94.7** (2026), GPQA **83.5**, SWE-Bench Verified **76%**, WildCodeBench **47.6%**; MCP Atlas **75.5–75.8** (sources vary slightly) vs Gemma4-31B 54.2 vs Qwen3.6-27B 62.5; OSWorld-Verified **65.9** vs Qwen3.6-27B 75.6. Bundled **DFlash** speculative drafter: RTX 5090 **74.9 → 233.4 tok/s** (3.1×); M5 Max 26.6 → 50.2 tok/s. vs Qwen3.6-27B: Qwen is the stronger pure coder (77.2% SWE-bench Verified); Glimmer wins on native multimodality and agent tooling.

## 2.8 Relationship to Llama

- Muse Spark is **not** an open-weights release and has **no migration path from Llama** ("fundamentally different deployment models"). Meta's stated rationale: Llama 4 (Apr 2025) failed commercially and its launch was marred by the LMArena benchmark scandal (acknowledged by Yann LeCun after his departure, Jan 2026).
- Alexandr Wang has promised to **open-source future (bigger) versions**; existing Llama models remain available but are in maintenance mode while frontier investment flows to Muse.
- One French source claims Muse Spark = instruction-tuned/RLHF version of a "Llama 5" base under a Llama 5 Community License — treat as unverified/single-source.

## 2.9 What distinguishes Muse Spark

- First paid Meta model API — strategic pivot from open weights to monetization under Wang/MSL; fast release cadence (1.0→1.3 in 5 months).
- Agentic-first design (context compaction over 1M tokens, multi-agent orchestration, computer use) with aggressive pricing aimed at high-volume agent workloads.
- Co-training of model + agent harness (Muse Spark 1.2 ↔ Muse Code).
- Strongest on: health/vision (1.0), long-horizon coding + long-context retrieval (1.3). Weakest on: abstract reasoning (ARC-AGI-2), open-ended agentic search vs OpenAI/Anthropic flagships.

---

