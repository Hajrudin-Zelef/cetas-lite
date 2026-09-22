---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/part-e-mimo-xiaomi-mimo-ai-team-led-by-luo-fuli-ex-deepseek
title: "PART E — MIMO (Xiaomi; MiMo AI team led by Luo Fuli, ex-DeepSeek)"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Xiaomi", "xAI"]
dates: ["2025-05", "2026-02-12", "2026-03-18", "2026-04-22", "2026-06", "2026-06-01", "2026-07-31", "2026-09-15", "2026-09-21"]
keywords: ["deepseek", "agent", "agentic", "attention", "aws", "bedrock", "benchmarks", "chatgpt", "compute", "cost", "decode", "distribution"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [137, 197]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: 667acf0761d85970c2a33b6b84318cdc21bcc8893cb984308ebabaf8aa5ecfdb
---

# PART E — MIMO (Xiaomi; MiMo AI team led by Luo Fuli, ex-DeepSeek)

## PART E — MIMO (Xiaomi; MiMo AI team led by Luo Fuli, ex-DeepSeek)

**Context:** First model MiMo-7B (MIT, Apr 30, 2025). CEO Lei Jun announced ≥**$8.7B** AI investment over 3 years (Mar 2026). MiMo Code (open-source terminal coding agent) and HarnessX shipped June 2026. V2.6's RL reportedly ran in part on **Chinese chips rather than Nvidia GPUs**.

### MiMo-V2.5 / V2.5-Pro — April 22, 2026
V2.5 = **310B total**; V2.5-Pro = **1.02T total / 42B active** MoE. 1M context; May 2025 cutoff. **MIT** (V2.5 made the Pro tier open; predecessor March-2026 MiMo-V2-Pro was proprietary). Agentic coding / long-horizon SE; "harness awareness". Benchmarks (vendor): SWE-bench Pro 57.2 with 40–60% fewer tokens than Opus 4.6; DeepSWE v1.1 19.0 (Pro). Launch pricing (overseas): Pro $1.00/$3.00 (≤256K), doubling for 256K–1M; cache reads $0.20–$0.40/M; base V2.5 $0.40/$2.00. (Pricing later revised downward before V2.6.)

### MiMo-V2.6-Flash — September 21, 2026
**309B total / 15B active** MoE. **1M context**, 128K max output; natively omnimodal (text/image/audio/video in, text out). **MIT**; HF `XiaomiMiMo/MiMo-V2.6-Flash-RL` (65 shards, ungated) + ModelScope. API: **$0.14/M in, $0.28/M out**, cached $0.0028/M; 50% batch discount. Benchmarks (vendor): DeepSWE v1.1 **65.68–67.9**, AutomationBench **52.3–52.7** (Flash *outscores* Pro here), Terminal Bench 2.1 87.6, CyberGym 95.1, Toolathlon-Verified 73.6.

### MiMo-V2.6-Pro ⭐ — September 21, 2026 (top open-weight model)
- **Architecture:** **1.02T total / 42B active**; **frozen-router MoE**; **hybrid attention**; **5-layer MTP speculative decoder**; RL via **fully asynchronous GRPO** — 1,568 prompts × 16 rollouts/step, 3.5–3.7B tokens/step, 750k trajectories, 30 RL steps in <6 days.
- **Context:** **1M tokens**, 128K max output; native multimodal in.
- **License:** **MIT**; HF `XiaomiMiMo/MiMo-V2.6-Pro-RL` (ungated) + ModelScope. Training cost (Xiaomi-reported): ~**$2.62M** Pro RL / $0.85M Flash; training livestreamed publicly at mimo.xiaomi.com/rl.
- **Benchmarks (vendor):** AA Index v4.3/v4.3.2: **46.32 (→46)** — best open-weight; ahead of Kimi K3, Qwen3.8 Max; tied with closed Grok 4.7 (46); above Grok 4.6 (44), Gemini 3.8 Flash (41), DeepSeek V4.1 Flash (39). AA measures **$0.13/task**, ~134 tok/s. DeepSWE v1.1 **71.9** (from 58.4 pre-RL; vs Opus 5 74.0, Sol ~73.0). AutomationBench **53.1** (beats all listed closed rivals). Terminal Bench 2.1 **89.9**. GDPval-AA 2.1 Elo 1673 (vs Opus 5 1708). MiMo Visual Coding 72.3 (vs Opus 5 70.0). Weak spots (vendor's own table): ProgramBench 26.5 (vs Opus 5 37.0); Terminal Bench 4.0 34.9 (vs 49.0); ExploitBench 47.9 (vs Sol 78.5).
- **Pricing:** **$0.435/M uncached in, $0.87/M out**; cached $0.0036/M; 50% batch discount. Xiaomi: "1/20 to 1/60 of overseas models at the same intelligence level."
- **Why it leads:** (1) scaled RL on verifiable long-horizon agentic tasks (30 steps / ~750k trajectories / 6 days); (2) frozen-router MoE + hybrid attention + MTP at 42B/15B active params; (3) MIT + ungated HF + released RL environments (7,000+ task envs); (4) V2.5-era pricing kept with higher intelligence → cost Pareto record ($0.13/task).
- **Pricing note:** the brief's "$0.20/$0.70" could NOT be matched to any MiMo tier — likely confusion with V2.5-era cache-read pricing ($0.20/M) or Grok 4.1 Fast. Treat as unverified for MiMo.

### MiMo-V2.6-Pro-UltraSpeed (serving tier, not separate weights)
Accelerated inference tier: up to **20× faster output at the same quality**; 10× price → **$4.35/$8.70** (cached $0.036/M). MiMo API / MiMo Desktop. No Flash UltraSpeed tier documented.

---

## PART F — MINIMAX (Shanghai, est. 2021; also Hailuo video, MiniMax Agent)

**License trajectory (tightest of the families):** M2 (MIT) → M2.1 (modified MIT) → M2.5 (frontmatter "modified-mit" but file is "MINIMAX MODEL LICENSE" — trust the file) → M2.7 (non-commercial) → M3 (Community License, non-commercial default; <$20M notice, >$20M prior written authorization) → H3 (community license + **"Applicable Territory" geographic restriction**).

### MiniMax M2.5 — February 12, 2026
MoE **229B total / ~10B active**; RL in "hundreds of thousands of real-world environments" (Forge framework). 196K context (AWS Bedrock card). Variants: Standard (50 tok/s), **Lightning (100 tok/s)**. License: MINIMAX MODEL LICENSE. Benchmarks: SWE-bench Verified **80.2** (Droid scaffold 79.7 vs Opus 4.6 78.9), BrowseComp 76.3. Pricing: $0.30/$1.20, cache read $0.03/M; "intelligence too cheap to meter" (~$1/hr).

### MiniMax M2.7 — March 18, 2026
**230B / 10B active** MoE; ~200K context; text-only. **Non-commercial** license (commercial needs prior written authorization + "Built with MiniMax M2.7" display). Benchmarks: SWE-Pro **56.22%** (matching GPT-5.3-Codex), Terminal Bench 2 ~57.0, AA Index **50** (+8 over M2.5 in <1 month). Pricing $0.30/$1.20. **First model to participate in its own development** — 100+ rounds autonomous scaffold self-optimization, +30% internal gain, zero human intervention. Caveat: per-task verbosity can run ~3× headline rate (~87M vs 26M median tokens).

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

