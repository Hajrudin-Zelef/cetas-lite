---
id: collect-261001-general-networking/general-networking/modeles-ia-open-locauxen-4
title: "ÉTAPE 1 — Open / Local AI Models (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Huawei", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai"]
dates: ["2026-05", "2026-06-01", "2026-07-31", "2026-09-15"]
keywords: ["agent", "apache", "attention", "benchmarks", "chatgpt", "compute", "datacenter", "decode", "deepseek", "distribution", "fp4", "glm"]
source: docs/RAG/collect-261001-general-networking/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [171, 233]
sha256: 8c987c049e92488c2979fb5e9d07866564ff3890f9874f2ef7ad0c723b316e59
---

# ÉTAPE 1 — Open / Local AI Models (EN)

### MiniMax M3 — June 1, 2026
- **Architecture:** MoE **~428B total / ~23B active**; **MiniMax Sparse Attention (MSA)** — blockwise sparse attention on GQA: **9× prefill / 15× decode speedups vs M2 at 1M ctx, ~1/20 per-token compute** (arXiv:2606.13392); native multimodal (text+image+video from step one); three thinking modes; native MXFP4 + MXFP8.
- **Context:** **1M tokens** (guaranteed min 512K under load); 131K max output.
- **HF:** `MiniMaxAI/MiniMax-M3` (+ `-MXFP8`, 443.8GB serving checkpoint; BF16 854.2GB); MSA kernel open-sourced (`MiniMax-AI/MSA`); NVIDIA build.nvidia.com; OpenRouter `minimax/minimax-m3`; Anthropic-compatible endpoint.
- **License:** **"MINIMAX COMMUNITY LICENSE"** — non-commercial by default; commercial requires "Built with MiniMax M3" display + notice to api@minimax.io (≤$20M/yr) or **prior written authorization (>$20M/yr)**; Prohibited Uses appendix.
- **Benchmarks (vendor):** **SWE-bench Verified 80.5** (vs V4-Pro 80.6, K2.6 80.2, GPT-5.5 82.9, Opus 4.7 87.6); **SWE-bench Pro 59.0** (beats GPT-5.5 58.6); Terminal-Bench 2.1 66.0; BrowseComp 83.5; Video-MME v2 85.4; SVG-Bench 63.7 (beats Opus 4.7 62.3). One independent harness reproduced ~80.5 territory.
- **Pricing:** Standard ≤512K: **$0.30/M in / $1.20/M out** (cache read $0.06; "permanent 50% off" applied); >512K: $0.60/$2.40; Token Plans $20/$50/$120 per month.

### MiniMax H3 — announced July 31, 2026 (weights ~Aug 2026) ⚠ NOT an LLM
**H3-Omni-Transformer, 33B dense single-stream transformer for native joint video+audio generation** (4–15s clips, 24 FPS, **native 32 kHz stereo audio in the same forward pass**, 768p native / 2K via H3-Regenerate-2K, 11 dialogue languages). Two open checkpoints (FL2VA, Ref2VA); H3-Context-IR and H3-Regenerate-2K closed. **First open model to top an AI video ranking** (AA: #1 Video Editing, #2 Text-to-Video, #3 Image-to-Video). **License: "MiniMax-H3 Community License" with "Applicable Territory" geographic restriction** (reported to effectively lock out US/EU — procurement risk). Commercial only under $20M revenue. HF `MiniMaxAI/MiniMax-H3`; fits single RTX 5090-class GPU.

---

## PART G — "JEV": identity resolved ⚠ NOT a Chinese open model

- **What:** Jev is a **proprietary American AI model** by **TypeSafe AI** (San Francisco, founded 2024 by **Diogo Almeida** — ex-OpenAI: RLHF, InstructGPT, ChatGPT, GPT-4 — Erik Gafni, Sasha Sheng).
- **Release:** **September 15, 2026**, limited early access (waitlisted), ~2-year stealth exit; **$40M seed led by DCVC** (~$200M valuation, per Forbes).
- **Category:** first of a new class of **"System One models"** (Kahneman's fast/intuitive System 1 vs deliberative System 2 LLMs).
- **Key difference:** **not a generative LLM** — generates **no natural-language text**. Takes program state + typed questions, returns **typed values with probability/confidence** for direct software consumption. Three primitives: **Choice** (distribution over options), **Score** (probability-weighted rubric score), **Noul** (yes/no probability). All questions answered in one parallel pass.
- **Specs:** 70–500 ms latency; 64K context (32K state + longest question); text-only at launch. Trained with **RLCD (Reinforcement Learning for Calibrated Decisions)** — distinct from RLHF/RLVR.
- **License:** **proprietary / closed** — managed API only, **no open weights**. Routes: `jev-latest`, `jev-preview`, `jev-1.13.0`. Python + JS/TS SDKs.
- **Pricing:** **$0.042 per 1M input tokens, output free**.
- **Open-source clone:** **Laya** (independent researcher, Sept 2026) — open weights, claimed 6–8× faster, benchmarked against Jev.
- **RAG filing:** does NOT belong in the open-weight Chinese track. File under a separate "new model categories / System One" track. Own Wikipedia page: https://en.wikipedia.org/wiki/Jev_(AI_model).

---

## CROSS-FAMILY COMPARISON (Sept 2026)

| Model | Release | Params (total/active) | Context | License | API $/M in/out | SWE-bench Verified | AA Index |
|---|---|---|---|---|---|---|---|
| Qwen3.5-397B-A17B | Feb 2026 | 397B / 17B | 262K→1M+ | Apache 2.0 | — | 76.2 | — |
| Qwen3.6-27B | Apr 2026 | 27B dense | 256K | Apache 2.0 | — | 77.2 | — |
| Qwen3.6-Max-Preview | Apr 2026 | 35B / 3B | 260K | Proprietary | — | — | 52 (v4.1.1) |
| Qwen3.7-Max | May 2026 | >1T | 1M | Proprietary | $2.50/$7.50 | — | 56.6 (v4.1.1) |
| Qwen3.8-Max | Aug 2026 | 2.4T / 95B | 1M | Custom (open) | $2 / $6 | — (Pro 67.7) | 45 (v4.3) |
| DeepSeek V4-Pro | Apr 2026 | 1.6T / 49B | 1M | MIT | $0.66/$1.98 | — | 36 |
| DeepSeek V4-Flash | Apr 2026 | 284B / 13B | 1M | MIT | $0.22/$0.66 | — | 35 |
| DeepSeek V4.1 Flash | Sep 2026 | 552B+196B / 8–16B | 1M | MIT | $0.15/$0.60 | — | ~39–40 |
| Kimi K2.6 | Apr 2026 | 1T / 32B | 256K | Modified MIT | $0.95/$4.00 | 80.2 | 54 |
| Kimi K3 | Jul 2026 | 2.8T / ~16×896e | 1M | Custom ($20M MaaS gate) | $3 / $15 | — | 44 (v4.3) |
| GLM-5.2 | Jun 2026 | 753B / 40B | 1M | MIT | $1.40/$4.40 | — (Pro 62.1) | 51 |
| GLM-5.3 | Aug 2026 | 753B / 40B | 1M | Custom ($10B review) | ~$2.15 blended | — | 45 (v4.3) |
| GLM-5.3-Flash | Aug 2026 | 320B / 18B | 1M | MIT | $0.15/$0.50 | — | 57 (v4.1.1) |
| MiMo-V2.5-Pro | Apr 2026 | 1.02T / 42B | 1M | MIT | $1.00/$3.00* | — (Pro 57.2) | — |
| MiMo-V2.6-Pro | Sep 2026 | 1.02T / 42B | 1M | MIT | $0.435/$0.87 | — (DeepSWE 71.9) | 46 (v4.3) |
| MiMo-V2.6-Flash | Sep 2026 | 309B / 15B | 1M | MIT | $0.14/$0.28 | — | — |
| MiniMax M2.5 | Feb 2026 | 229B / 10B | 196K | MINIMAX license | $0.30/$1.20 | 80.2 | — |
| MiniMax M2.7 | Mar 2026 | 230B / 10B | 200K | Non-commercial | $0.30/$1.20 | — (Pro 56.22) | 50 |
| MiniMax M3 | Jun 2026 | 428B / 23B | 1M | Community ($20M gate) | $0.30/$1.20 | 80.5 | — |
| Jev (TypeSafe) | Sep 2026 | n/a (decision model) | 64K | Proprietary | $0.042/in, out free | n/a | n/a |

\*V2.5 launch-era pricing; later revised downward.

## STRUCTURAL FINDINGS (for RAG metadata)
1. **License tightening is universal** across Chinese open-weight families: MIT/Apache → custom licenses with revenue-gated MaaS clauses ($20M: Kimi K3, MiniMax M3; $10B security review: GLM-5.3; territorial: MiniMax H3). Counter-example: GLM-5.3-Flash stayed plain MIT.
2. **Sovereign-hardware narrative**: GLM-5.3-Flash serves entirely on ~100K domestic Chinese accelerators (Cambricon/Huawei/Moore Threads) — built in ~13 days with Infra-Agent automation; MiMo-V2.6's RL ran partly on Chinese chips.
3. **Staged open-weight release**: GLM-5.3's 2-week safety hold is the first explicit pause→harden→ship precedent.
4. **1M context is table stakes** for flagships (K3, GLM-5.x, M3, MiMo-V2.x, Qwen Max, DeepSeek V4.x); Kimi K2.x (256K) and MiniMax M2.x (~200K) are the exceptions.
5. **"Open" now means datacenter-scale**: flagship open models are trillion-parameter MoEs; small dense models (Qwen3.6/3.8-27B) serve the single-GPU tier.
6. **Architectural trend**: off plain attention — Gated DeltaNet hybrids (Alibaba), compressed-sparse attention + FP4 KV (DeepSeek), KDA linear attention (Moonshot, Z.ai), MiniMax Sparse Attention, encoder-decoder (DeepSeek V4.1 Flash), Engram conditional memory (DeepSeek).
7. **Meta out**: Llama 5 delayed to 2027, Meta pivoted to closed Muse Spark — Chinese labs own the open frontier.

