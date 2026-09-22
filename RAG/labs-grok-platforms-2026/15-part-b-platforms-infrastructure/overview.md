---
id: labs-grok-platforms-2026/15-part-b-platforms-infrastructure/overview
title: "PART B — PLATFORMS & INFRASTRUCTURE"
domain: part-b-platforms-infrastructure
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Google", "Hugging Face", "Intel", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2024-01", "2026-04-25", "2026-05", "2026-05-15", "2026-06-04", "2026-06-26", "2026-07"]
keywords: ["acquisition", "agent", "agentic", "agents", "amd", "deepseek", "glm", "gpu", "gpus", "inference", "intel", "kimi"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [476, 520]
section: "PART B — PLATFORMS & INFRASTRUCTURE"
sha256: 6048cea639289c3c1cf82e3fbdbed1547423e9e51d13ca1280a7ce633ba0534c
---

# PART B — PLATFORMS & INFRASTRUCTURE

## B.1 — OpenRouter: the multi-model exchange

Unified API gateway routing requests across **400+ LLMs** through a single OpenAI-compatible interface. In 2026 it became the industry's de-facto **model-release venue, usage-ranking authority, and agentic-infrastructure layer**.

- **Explosive growth:** weekly token usage up **~9,000× since January 2024** (from ~10B/week to **90+ trillion/week** by mid-2026). Weekly volume: 5T (Nov 2025) → 25T (May 2026) → **55T+ (Aug 2026)**. ~100T tokens/month.
- **Agents eat the platform:** agentic workloads grew **14× Feb–Aug 2026** (0.51T → 7.3T tokens/day); by early Aug 2026 agents were **~71% of all consumption**; agents consume **~5× more tokens per task** than humans, with **85%+ of agentic burn in cached prompts**.
- **The Chinese flip:** Chinese open models went from **~2% of OpenRouter usage (2025) to >50% (mid-2026)**; US frontier models fell from 70% to 30%. Crossover week: **Feb 9–15, 2026**.
- **Business:** **$113M Series B (May 2026)**, led by CapitalG — valuing it at **~$1.3B**; **$140M annualized revenue** by July 2026, growing **29% month-over-month**; **8M+ users**.

## B.2 — Hugging Face: the open-weights hub

- **NVIDIA acquisition (Sept 3, 2026):** Nvidia agreed to buy Hugging Face for **$12.93B** — Nvidia's second-largest deal. HF had been valued at **$4.5B** in a $235M round (2023).
- **Scale:** **18M+ developers/researchers/creators**; **3M+ models**, 500K datasets, 1M apps; **200,000+ companies**.
- **Download concentration (2026):** top-1,000 repos ≈ 2.54B monthly downloads (July 2026); **Qwen family dominates: 240.1M 30-day downloads (Sept 2026)** — 36.32% of text-gen top-1,000 share; Chinese open-weight models = **41% of HF downloads** (spring 2026).
- **Role in 2026 releases:** launch-day weights host for DeepSeek V4.1 Flash, GLM-5.2/5.3-Flash (MIT), MiMo-V2.6 (MIT), Kimi K3 (1.56 TB repo).

## B.3 — vLLM: the production workhorse

Origin: UC Berkeley. Core innovation: **PagedAttention** + continuous batching → 2–4× throughput, 100+ concurrent requests/GPU.

- **2026 state:** v0.21.0 (May 15, 2026). Broadest hardware support: NVIDIA, AMD, Intel, Google TPUs, Arm. Used internally by **Google Vertex AI and NVIDIA**. **Day-0 support for NVIDIA Nemotron 3 Ultra** (June 4, 2026).
- **Positioning:** the mature default for production API serving with SLA requirements.

## B.4 — SGLang: the performance-first challenger

Built by the **LMSYS team**. Core innovation: **RadixAttention** + Python frontend language + jump-forward constrained decoding.

- **2026 state:** deployed on **400,000+ GPUs worldwide**.
- **Day-0 model support as a competitive weapon:** first open-source stack with **Day-0 DeepSeek-V4 inference + RL training** (April 25, 2026). **SGLang v0.5.14 (June 26, 2026)**: native GLM-5.2, Kimi-K2.7-Code, MiMo, Nemotron-H.
- **Positioning:** dominates structured-generation and multi-turn agent workloads — **4.7× speedup in multi-turn conversations**, 99.8% JSON validity, **47% less GPU memory** vs vLLM-class in agent workflows.

## B.5 — vLLM vs SGLang (2026 decision guide)

| Dimension | vLLM | SGLang |
|---|---|---|
| Origin | UC Berkeley | LMSYS |
| Core trick | PagedAttention | RadixAttention (prefix tree) |
| Strength | High-concurrency API serving, predictable latency, maturity | Structured generation, multi-turn agents, prefix-cache reuse |
| Memory | baseline | ~47% less in agent workflows |
| Adoption signal | Vertex AI, NVIDIA internal | 400K+ GPUs, HF Endpoints native engine |

---

