---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/minimax-m3-june-1-2026
title: "MiniMax M3 — June 1, 2026"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "EU", "MiniMax", "Nvidia", "OpenAI", "OpenRouter", "United States"]
dates: ["2026-06-01", "2026-07-31", "2026-09-15"]
keywords: ["attention", "benchmarks", "chatgpt", "compute", "decode", "distribution", "gpu", "gqa", "latency", "license", "moe", "multimodal"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [171, 197]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: d5a23a537c1952b661f60043cbcc58a2847b2526b6ca9b5ed08808a74b8a79fa
---

# MiniMax M3 — June 1, 2026

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

