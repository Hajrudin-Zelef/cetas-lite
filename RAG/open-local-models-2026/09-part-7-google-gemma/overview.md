---
id: open-local-models-2026/09-part-7-google-gemma/overview
title: "PART 7 — GOOGLE GEMMA"
domain: part-7-google-gemma
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Apple", "Google", "vLLM"]
dates: ["2026-04-02"]
keywords: ["agent", "agentic", "apache", "benchmarks", "consumer", "embedding", "embeddings", "gemini", "gpu", "leaderboard", "license", "llama"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [619, 638]
section: "PART 7 — GOOGLE GEMMA"
sha256: 0f219a46377c5e0f306226234124010f354e47ea368a2c048d9b62825bda02d0
---

# PART 7 — GOOGLE GEMMA

## 7.1 Gemma 4 (April 2, 2026) — license watershed

- **The headline change:** Gemma 1/2/3 used the restrictive **Gemma Terms of Use** (commercial MAU thresholds, redistribution limits). **Gemma 4 ships under Apache 2.0** — unlimited commercial use, full redistribution. This is Google's first direct entry into the open-weight enterprise market against Llama.
- **Lineup (4 sizes, base + instruction-tuned on HF):**
  - **E2B:** 2.3B effective (5.1B w/ embeddings), 128K, 35 layers — text/image/audio
  - **E4B:** 4.5B effective (8B w/ embeddings), 128K, 42 layers — text/image/audio
  - **26B-A4B (MoE):** 25.2B total / 3.8B active (128 experts, 8 routed + 1 shared per token), 256K, 30 layers — text/image/video
  - **31B (dense):** 30.7B, 256K (vals.ai: 262K), 60 layers — text/image/video
  - 262K vocabulary; 140+ languages (up from 35+ in Gemma 3); built-in thinking/reasoning mode; native function calling; system-prompt support; video understanding up to 60s @1fps (a Gemma first).
- **Architecture notes:** E-series uses **Per-Layer Embeddings (PLE)** — a secondary embedding signal fed into every decoder layer — so 2.3B-active behaves like 5.1B params and fits <1.5GB at 2-bit quant (verified on Raspberry Pi 5). Ships **MTP (Multi-Token Prediction) draft models** for speculative decoding: up to 2× speedup with identical output quality.
- **HF weights:** `google/gemma-4-31B-it`, `google/gemma-4-26B-A4B-it`, `google/gemma-4-E4B-it`, `google/gemma-4-E2B-it`; Kaggle; Ollama (`ollama pull gemma4`); LM Studio; vLLM; MLX (Apple Silicon); llama.cpp; LiteRT-LM for on-device. **Free API via Google AI Studio** (31B/26B).
- **Benchmarks — Gemma 4 31B (vendor-reported, instruction-tuned):** MMLU Pro **85.2**; AIME 2026 **89.2** (no tools; vs 20.8 for Gemma 3 27B); LiveCodeBench v6 **80.0**; Codeforces ELO 2150 (vs 110 for Gemma 3); GPQA Diamond 84.3; Vision MMMU Pro 76.9; MATH-Vision 85.6; agentic **τ²-bench 86.4%** (vs 6.6% for Gemma 3 — a categorical change for agent use); vals.ai: #1 on SAGE (55.03%), 53.92% on SWE-bench Verified subset, 38.94% Vals Index; AA Intelligence Index **39**.
- **Benchmarks — 26B-A4B MoE:** MMLU Pro 82.6; AIME 2026 88.3; LiveCodeBench v6 77.1; Codeforces 1718 — ≈97% of dense-31B MMLU-Pro quality at ~12% of the FLOPs; runs on a 24GB consumer GPU.
- **Positioning:** derived from the Gemini 3 research lineage; the 31B ranks **#3 among all open models** on the Arena AI text leaderboard; E2B/E4B are the only Western edge models with native audio at that size class (no Llama 4/Qwen 3.5 equivalent); powers Gemini Nano 4 on Android.
- **Pricing:** free via Google AI Studio; third-party hosted pricing is low/near-zero; self-host on a single H100 (31B) or consumer GPU (26B MoE).

---

