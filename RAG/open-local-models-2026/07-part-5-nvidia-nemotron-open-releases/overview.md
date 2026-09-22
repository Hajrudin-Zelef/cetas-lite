---
id: open-local-models-2026/07-part-5-nvidia-nemotron-open-releases/overview
title: "PART 5 — NVIDIA NEMOTRON (OPEN RELEASES)"
domain: part-5-nvidia-nemotron-open-releases
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "China", "DeepSeek", "Fireworks AI", "Google", "Meta", "Moonshot", "Nvidia", "OpenRouter", "Perplexity", "United States", "Z.ai"]
dates: ["2025-05", "2026-07"]
keywords: ["nvidia", "agent", "agentic", "agents", "aws", "bedrock", "benchmarks", "blackwell", "context window", "cost", "deepseek", "fine-tuning"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [577, 607]
section: "PART 5 — NVIDIA NEMOTRON (OPEN RELEASES)"
sha256: 3c81bf331a04026cc0d0cf69f066f84aa6615471f69fc395acc21646fe704338
---

# PART 5 — NVIDIA NEMOTRON (OPEN RELEASES)

**Vendor context:** NVIDIA's own model family, positioned as the open standard for **agentic AI workloads**. Differentiator vs Llama (and vs most open releases): NVIDIA publishes **weights + training data + RL environments + post-training recipes + fine-tuning code** — "fully open, not open in the marketing sense." Note NVIDIA previously built on Meta's weights (Llama-3.1-Nemotron 70B, Llama Nemotron reasoning family Apr/May 2025: Nano 4B / Super 49B / Ultra 253B); the Nemotron 3 generation is original NVIDIA architecture, optimized for NVIDIA GPUs (NVFP4 on Blackwell).

## 5.1 Nemotron 3 family (Nano / Super / Ultra)

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

