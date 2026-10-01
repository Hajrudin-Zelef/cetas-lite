---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-10
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Alibaba", "Apple", "China", "DeepSeek", "Fireworks AI", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Perplexity", "United States", "Z.ai", "vLLM"]
dates: ["2026-04-02", "2026-07"]
keywords: ["agent", "agentic", "agents", "apache", "aws", "bedrock", "benchmarks", "consumer", "context window", "cost", "deepseek", "distribution"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [583, 642]
sha256: 971cdcf3adfef77a990ae253854b5cd0ba38135884985c8ef8b848d660570db6
---

# ÉTAPE 1 — Open / Local AI Models (EN)

- **Releases:** Nano — Dec 15, 2025; Super — Mar 11, 2026 (GTC); Ultra — Jun 4, 2026 (Computex/GTC Taipei keynote).
- **Architecture:** hybrid **Mamba-Transformer MoE** (Nemotron-H / latent MoE). Learned MLP router: Nano activates 6 of 128 experts per forward pass (~10% active).
- **Sizes:** Nano 31.6B total / ~3.2–3.6B active; Super 120B total / ~12B active; Ultra **550B total / 55B active** (~90% sparse; 108 layers, 8192 dim, 512 experts/layer, top-22 routing). Ultra pretrained on 20T tokens (Warmup-Stable-Decay); post-training SFT → RLVR → MOPD → MTP Boosting.
- **Context window:** **1M tokens** (Ultra NVFP4; BF16 HF checkpoint ships 262K); Super 1M (256K HF); Nano 1M (262K HF).
- **License:** NVIDIA Open Model License (Nano family); Ultra governed by **OpenMDW-1.1** (per NVIDIA API docs) — both permit commercial use and redistribution.
- **HF weights:** `nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16` (+ NVFP4 and FP8 variants), Nano/Super equivalents under `nvidia/`. Also on AWS Bedrock, Google Cloud, NVIDIA NIM, OpenRouter, Together, Fireworks, DeepInfra, SageMaker JumpStart.
- **Features:** reasoning ON/OFF control with configurable thinking-token budget; built-in MTP head for speculative decoding; reasoning-optimized agent scaffolds.
- **Benchmarks — Nemotron 3 Ultra** (HF model card + NVIDIA API docs):
  - Artificial Analysis **Intelligence Index 48** — highest of any **US-origin open-weight model** (vs Gemma 4 31B: 39, Nemotron 3 Super: 36, gpt-oss-120b: 33; China's Kimi K2.6: 54, DeepSeek V4 Pro: ~56)
  - SWE-Bench Verified **70.7** (HF card; NVIDIA docs: 71.9 BF16 / 69.7 NVFP4); range 65–70.4% across five agent harnesses (Pi, OpenHands, Hermes, OpenCode, Mini SWE Agent)
  - SWE-Bench Multilingual 67.7; Terminal-Bench 2.1 56.4; GPQA (no tools) 87.0; LiveCodeBench v6 89.0; IOI 2025 570.0; RULER @1M 94.7; IFBench 81.7; τ-bench V3 avg 70.9; PinchBench 90.0; BrowseComp 44.4; HLE 26.7
  - **Throughput:** 300+ tok/s (DeepInfra pre-release, BF16); NVFP4 on GB200: 5.9× throughput vs GLM-5.1-754B, 4.8× vs Kimi K2.6 at 8K/64K; up to 30% lower cost-to-completion than peers
- **Benchmarks — Nemotron 3 Nano:** 3.3× throughput vs Qwen3-30B-A3B and 2.2× vs GPT-OSS-20B (8K in/16K out, single H200); beats both on RULER; 4× throughput of Nemotron 2 Nano; 60% fewer reasoning tokens; Artificial Analysis ranks Nemotron 3 most efficient among same-size open models. 13 early enterprise adopters (Accenture, CrowdStrike, Palantir, Perplexity).
- **API pricing (June–July 2026):** Ultra — OpenRouter $0.50/$2.50 (+$0.15 cached), DeepInfra $0.50/$2.20, Fireworks NVFP4 $0.60/$2.40 ($0.12 cached), Together $0.60/$3.60. Blended AA rate ≈ $0.52/1M.
- **Positioning vs Llama:** NVIDIA is building a Llama-independent open stack — original architectures, open data/recipes, GPU-co-designed efficiency — aimed at enterprise agentic workloads (coding agents, RAG, orchestration) where cost-per-task and long context matter more than chat benchmarks. Meta's Llama line was in maintenance mode by mid-2026, leaving NVIDIA as the leading US open-weight lab on the Intelligence Index.

## 5.2 Other 2026 Nemotron releases (brief)

- **Nemotron 3 Nano Omni** (Apr 28, 2026): 30B-A3B multimodal (text+image+audio+video), NVIDIA Open Model License.
- **Nemotron-Cascade-2-30B-A3B** (Mar 19, 2026): 30B-A3B, cascade RL, thinking+instruct unified.
- **Nemotron 3.5 Content Safety** (Jun 4, 2026): open-weight safety model, 128K context.
- **Nemotron Nano 2 12B** (Dec 2, 2025): hybrid Mamba-Transformer, 131K context. (Nano 2 9B: Aug 2025.)

---

# PART 6 — MICROSOFT PHI

**Status: incremental in 2026 — no Phi-5.** Two independent 2026 surveys confirm "Phi-5 unreleased" as of mid-2026.

- **Phi-4-reasoning-vision-15B** (Mar 2026) — current latest. Selective reasoning (model decides when to think), ~200B multimodal tokens; fits 12GB VRAM. MIT license.
- **Phi-4** (Dec 12, 2024, 14B dense, MIT); **Phi-4-mini** (3.8B); **Phi-4-multimodal** (5.6B, unified text+speech+vision). All MIT-licensed, HF + Ollama (`ollama pull phi-4`) + LM Studio.
- **Positioning as small local models:** Microsoft's SLM line is the reference "runs anywhere" family. **Foundry Local** reached GA (announced at Build 2026, milestone Apr 9): a ~20MB embeddable runtime (Windows/macOS Apple Silicon/Linux x64, no cloud, no per-token cost, OpenAI-compatible API, ONNX Runtime backend claiming 3.9× throughput over llama.cpp). Its curated GA catalog ships **Phi**, Qwen, DeepSeek, Mistral, Whisper — Phi is the first-party small-model default for on-device Windows/agent scenarios.
- **Note:** Phi-4-class models are now outclassed by newer small open models (e.g., Qwen3.5-9B, Gemma 4 E4B) on benchmarks; their value is MIT licensing + Microsoft distribution, not leadership.

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

# PART 8 — OTHER NOTABLE WESTERN OPEN-WEIGHT RELEASES

## 8.1 Allen Institute for AI — OLMo 3 (Nov 20, 2025)

