---
id: ai-industry-kb-2026/03-chinese-ai-labs-deep-file-glm-kimi-minimax/minimax
title: "MiniMax"
domain: chinese-ai-labs-deep-file-glm-kimi-minimax
role: deep-dive
task: actor-profile
actors: ["Anthropic", "ByteDance", "China", "DeepSeek", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-01", "2026-04-13", "2026-04-20", "2026-06-23", "2026-06-24", "2026-07", "2026-07-16", "2026-07-27", "2026-07-31", "2026-08", "2026-08-02", "2026-08-10", "2026-08-12", "2026-08-14", "2026-08-18", "2026-08-20", "2026-08-26", "2026-09-22"]
keywords: ["agent", "attention", "benchmark", "compute", "consumer", "cybersecurity", "deepseek", "fp8", "glm", "ipo", "kimi", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1039, 1096]
section: "3. Chinese AI Labs — Deep File: GLM, Kimi, MiniMax"
sha256: 3e9c239f9f1fdfe0d456ad8ecf043d76d081ee686918881691c555079f9e1a06
---

# MiniMax

- **2026-04-13 — Kimi K2.6 preview opens** (an eight-day preview; single SEO-page sourcing for the exact start, corroborated by the GA dates).
- **2026-04-20/21 — Kimi K2.6 general availability**: April 20 per awesomeagents / tpsreport / verdent; April 21 per datanorth and kimi-k2.org. Pin GA as **2026-04-20/21** and mention the preview phase — all three April dates are real phases, not errors.
- **2026-07-16 — Kimi K3 launch**; **~2026-07-27 — K3 open weights** (chronology canonical in wave3/02 Claim 4; kept here as 1–2 lines plus cross-reference per the dedup rule). "Beats Opus 4.8" is benchmark-dependent (GDPval-AA v2: 1687 vs 1600; AA Intelligence ~57).

### MiniMax

- **January 2026 — MiniMax lists publicly in Hong Kong**, the second of China's "AI tigers" to list after Z.AI (Reuters).
- **Early July 2026 — Reuters reports (person with direct knowledge) that MiniMax was developing a 2.7-trillion-parameter language model** alongside H3 [UNVERIFIED]. Only the H3 half materialized. No confirmation, preview, or leak under any other name by 2026-09-22.
- **2026-07-31 — MiniMax H3 = Hailuo 3.0 = "Hailuo 03" launched**: API (model ID `MiniMax-H3`) and the consumer Hailuo AI app. Date carries a ±1-day flag (Reuters via srnnews.com reported July 30/31). Consolidation merges the three names into ONE entry.
- **2026-08-02/03 — H3 open weights** (canonical detail in wave2.1/07 §1.1; Reuters had said "within days" of the July 31 launch).

### ByteDance (Seed)

- **2026-06-23/24 — Volcano Engine FORCE conference**: ByteDance announces Seed 2.1, Seed 2.1 Pro, and a preview tier. The official launch blog (seed.bytedance.com) names **no Turbo variant**; one trade report (DataNorth, June 24) places Seed 2.1 Pro **and** Seed 2.1 Turbo at FORCE. Pin **2026-06-24** for the Turbo with a single-source caveat.
- **2026-06-23 — Seedance 2.5 announced** at FORCE; **2026-07-31 — global launch on Dreamina** (canonical home wave3/06-multimodal-revolution-check.md §1.3; dating correction there supersedes wave2.1 §1.3's "~early Aug"). Not duplicated here.
- **2026-08-10/12 — Seed 2.1 Turbo first-seen on Western trackers** (LLM Gateway PR #3580, NanoGPT catalog, OpenRouter). These are aggregator first-seen dates, not release dates — one engineering journal documents the "launched twice, 47 days apart" discrepancy explicitly. Never present them as the release.

### GLM-5.3 rumor-phase chronology (standing rule for the KB)

- Early August 2026: speculation only — any entry dated here stays [UNVERIFIED] rumor.
- 2026-08-14: Z.ai announcement → first confirmed sighting; confirmed window for attributable claims is **August 14–29, 2026**.
- Never merge a pre-announcement speculation date with a post-announcement confirmation as if they were one event.

### Access and catalog milestones

- **2026-04-20/21** — `moonshotai/Kimi-K2.6` live on Hugging Face (Modified MIT) at GA; Moonshot API serving at $0.60/M input; vLLM / SGLang / KTransformers deployment paths documented in launch coverage.
- **2026-08-14** — GLM-5.3 launch access opens through Z.ai's API, the GLM Coding Plan ($18/mo base tier), and the ZCode agent (API per-token pricing follows 2026-08-18).
- **2026-08-20** — anonymous "Ox Alpha" tops OpenRouter usage charts (retrospectively GLM-5.3-Flash); **2026-08-26** — GLM-5.3-Flash announced (320B/18B MoE, MIT, $0.15/M; canonical home wave1/07 §2.4).
- **2026-08-10** — LLM Gateway PR #3580 adds Seed 2.1 Turbo (first-seen evidence, not a release); **2026-08-12** — NanoGPT catalog and OpenRouter listings log the Turbo (first-seen, not release).
- **2026-09-22** — verification cutoff for this file: the MiniMax 2.7T LLM rumor and the Moonshot HK IPO filing rumor are both still unconfirmed.

## Figures and metrics

### Parameter counts and architecture

#### GLM-5.2 — the three-way reconciliation (pin 753B total / ~40B active)

- **743.4B** = backbone parameter total derived from `config.json` (what vLLM reports as "~743B") [COMMUNITY — config-derived from shipped weights; deterministic math, but no Z.ai architecture paper exists to confirm].
- **753.3B** = including the **9.95B MTP speculative-decoding block** (what NVIDIA's model card reports as "753B").
- **39.3B** = active parameters per token (vLLM "39B" / official "A40B").
- **"744B"** = Z.ai's official naming shorthand ("744B-A40B"), identified by a detailed practitioner write-up as the **FP8-build / VRAM figure**, not the architectural count.
- **Consolidation guidance**: state **753B total / ~40B active** and footnote Z.ai's "744B-A40B" shorthand as the FP8-build figure, with the config-derived backbone at 743.4B. Do not present "744B" as architectural fact without the footnote. Wave3/02's "753B (not 744B)" framing was directionally right; this file adds the reconciliation and the FP8 origin of "744B".
- License: **MIT** open weights on Hugging Face. Context: **native 1M-token input**; **131,072 max output tokens** (docs.z.ai lists 128K rounded). Text-only — no vision, confirmed independently as the model's binding constraint.

#### GLM-5.2 — IndexShare sparse attention

- One attention indexer reused across every **4 sparse-attention layers**: **21 of 78 layers** carry a real indexer; the rest reuse.
- Z.ai reports **~2.9× per-token FLOP reduction at 1M context** [VENDOR].
- Built on **DeepSeek Sparse Attention** (MLA + lightning indexer, top-K 2048). Compiled with per-layer shapes in the tensorsharp architecture page.

#### GLM-5.3 — same base, post-training-only gains

- Z.ai states GLM-5.3 keeps the exact **743-billion-parameter MoE base** of GLM-5.2; every capability gain comes from scaled post-training — more RL environments, more task diversity, longer trajectories, more RL compute. **No architectural change, no new pretraining** (consistent with wave2's framing: "scaled post-training/RL on long-running tasks rather than bigger pre-training").
- Always-on reasoning with three effort levels (low/high/max); thinking can no longer be disabled — a **breaking change** from GLM-5.2. Text-only (no vision).
- Targeting: long-horizon software engineering + cybersecurity / vulnerability detection.

#### Kimi K2.6 — full spec sheet

