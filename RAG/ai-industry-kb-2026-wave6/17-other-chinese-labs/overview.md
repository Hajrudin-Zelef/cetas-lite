---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/overview
title: "§17. Other Chinese Labs"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Baidu", "ByteDance", "China", "DeepSeek", "Falcon", "LongCat", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "Perplexity", "Poolside", "StepFun", "TII", "Z.ai"]
dates: ["2025-06", "2026-01-05", "2026-01-22", "2026-03", "2026-04", "2026-05-08", "2026-06", "2026-06-12", "2026-06-23", "2026-06-24", "2026-06-30", "2026-07", "2026-07-21", "2026-07-23", "2026-07-31", "2026-08", "2026-08-13", "2026-08-28", "2026-09", "2026-09-18", "2026-09-20", "2026-09-22", "2026-09-27", "2026-10-15"]
keywords: ["acquisition", "agent", "apache", "benchmark", "cost", "deepseek", "glm", "kimi", "license", "licenses", "llama", "nvidia"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8248, 8319]
section: "§17. Other Chinese Labs"
sha256: c74a66582efe3265dd2a222554b16b7bfe47b43a1862e99e6e4dfe998b82ed05
---

# §17. Other Chinese Labs

Keywords: StepFun, Step 5 Preview, Baidu, ERNIE 5.0, ERNIE 5.1, ByteDance, Seed 2.1 Turbo, Seedance 2.5, Seed-Coder, InclusionAI, Ling 3.0 Flash, Falcon, TII, Falcon H1R, Falcon Perception, Poolside, Laguna S 2.1, Model Factory, Tencent, Hy4 Preview, openPangu 2.0, 01.AI, Perplexity, Sonar, R1-1776, DeepSeek R2

## Summary
- This section collects the Chinese-lab and adjacent delta facts not owned by other sections: StepFun, Baidu, ByteDance, InclusionAI, Falcon (TII/UAE), Poolside, Tencent, openPangu, 01.AI, and Perplexity. DeepSeek/Qwen/GLM/Kimi/LongCat/MiniMax/MiMo facts belong to §§2, 3, 7, 8; one line on DeepSeek R2 below only.
- **StepFun Step 5 Preview** (2026-09-20, majority date) is **API-only as of 2026-09-22** — BF16 weights promised 2026-10-15; do **not** call it open-weight [SECONDARY].
- Baidu's **ERNIE 5.0** (GA 2026-01-22) and **ERNIE 5.1** (~2026-05-08) carry an **unresolved pricing conflict** ($0.85/$3.40 vs $0.59/$2.65) [SECONDARY] — presented as-is, not merged.
- Falcon had **real 2026 releases** (H1R/H1 Arabic Jan, Perception + OCR ~April, Perception-300M July) — the "no 2026 Falcon releases" framing is corrected [SECONDARY].
- **Tencent Hy4 Preview** (2026-08-28, 770B/49B, Apache 2.0) and **openPangu 2.0** (announced 2026-06-12, staged from 2026-06-30; Pro 505B/18B, Flash 92B/6B, 512K) are the two large-preview open stories of the summer — but openPangu's official license is **unverified** (do not call Apache/MIT) [SECONDARY].
- NVIDIA **licensed Poolside's Model Factory for $6B** (+$1B investment) — not an acquisition, not a Laguna license; **Laguna S 2.1** (118B/8B, OpenMDW-1.1, $0.09/$0.18) is the cheapest open-weight API with a published rate [SECONDARY/VENDOR].
- **01.AI exited foundation-model building** [SECONDARY]; **Perplexity shipped no verified new-from-scratch 2026 model** — Sonar is a Llama 3.3 70B fine-tune, R1-1776 a DeepSeek-R1 fine-tune, and Sonar chat-completions endpoints retire 2026-09-27 [SECONDARY].

## Key dated facts
### StepFun Step 5 Preview — API-only, September 2026
- **2026-09-20** — StepFun **Step 5 Preview** launched (majority date; one source said 2026-09-18 — the conflict is noted, 09-20 retained) [SECONDARY].
- **API-only as of 2026-09-22**; BF16 weights promised **2026-10-15** — do not call it open-weight and do not claim live/open status for the promised weights [SECONDARY].
- Pricing: **$1.00/$2.70** input/output per million, cache read $0.05 [SECONDARY]; cost-per-Index-task **$0.71** (vs DeepSeek V4 Pro $0.67, V4.1 Flash $0.27) [SECONDARY].
- Master log position: Step 5 Preview vs weights — API-only as of Sept 22 ($1.00/$2.70); weights promised Oct 15 — do not cite as open-weights.

### Baidu ERNIE 5.0 / 5.1 — GA and the pricing conflict
- **2026-01-22** — **ERNIE 5.0** general availability [SECONDARY].
- **~2026-05-08** — **ERNIE 5.1** [SECONDARY].
- **Pricing conflict, unresolved:** ERNIE 5.0 at **$0.85/$3.40** vs ERNIE 5.1 at **$0.59/$2.65** [SECONDARY] — unclear whether the difference is version-driven or market-driven; the two figures are **not merged**.

### ByteDance — Seed 2.1 Turbo, Seedance 2.5, Seed-Coder
- **2026-06-24** — **Seed 2.1 Turbo** [SECONDARY].
- **Seedance 2.5**: announced **2026-06-23**, global availability **2026-07-31** [SECONDARY] — ByteDance's video-generation line cadence.
- **Seed-Coder is June 2025, not 2026** [SECONDARY] — the year correction stands.

### InclusionAI Ling 3.0 Flash
- **2026-07-23** — **Ling 3.0 Flash**: **124B total / ~5.1B active** [SECONDARY].
- License is **[UNVERIFIED]** — do not assert a license.
- Available on OpenRouter **:free** routes (rate-limited) [COMMUNITY].

### Falcon (TII) — real 2026 releases
- **2026-01-05** — **Falcon H1R / H1 Arabic** [SECONDARY].
- **~April 2026** — **Falcon Perception + OCR** (approximately April; exact date uncertain) [SECONDARY].
- **July 2026** — **Falcon Perception-300M** [SECONDARY].
- Correction held: Falcon had **real 2026 releases** — the "Falcon = contradicted (no releases)" reading is resolved in favor of actual Jan/Jul 2026 shipments.

### Poolside Laguna S 2.1 — July 2026
- **2026-07-21** — **Laguna S 2.1** released: **118B total / 8B active**, **OpenMDW-1.1** license [VENDOR — Poolside announcement].
- API pricing **$0.09/$0.18** per million input/output [VENDOR] — the cheapest open-weight API with a published rate in the wave.
- Vendor framing: "the West's most capable open-weight model" [VENDOR] — labeled as vendor language.

### The NVIDIA–Poolside deal — license, not acquisition
- NVIDIA **licensed Poolside's Model Factory for $6B and invested $1B** [SECONDARY].
- It did **not** acquire Poolside and did **not** license Laguna itself — both the acquisition framing and the Laguna-licensing framing are corrected.
- (NVIDIA-actor detail duplicated from §15 for RAG locality; §15 owns the coalition framing.)

### Tencent Hy4 Preview — August 2026
- **2026-08-13** — Hy4 Preview announced; **2026-08-28** — **Tencent Hy4 Preview** released: **770B total / 49B active**, **>1M context**, **Apache 2.0**; secondary coverage frames it as built "through deep model-product co-design" [SECONDARY].
- Do not confuse with **Hy3 Preview**, which sits under the Tencent Hy Community License — the Hy3-Preview/Hy4-Preview license split is the correction to hold; the completed **Hy3** also uses Apache 2.0 [SECONDARY].
- Hy4 is one of three 2026 Apache-2.0 releases above 190B total parameters (Hy4 770B, openPangu Pro 505B, Step 3.7 Flash 198B), plus dots3-note (280B) — permissive licenses now attach to frontier-scale parameter counts, not just small models [DIRECTIONAL].
- Preview-tier access with reported gating; context-length and benchmark caveats were recorded in the wave's Part 1 detail and carry [SECONDARY] provenance — pin them before citing specifics.

### openPangu 2.0 — June 2026
- **2026-06-12** — **openPangu 2.0** announced; **staged release from 2026-06-30** [SECONDARY].
- **Pro: 505B/18B**; **Flash: 92B/6B**; **512K** context [SECONDARY].
- The **official license is unverified** — do **not** call it Apache or MIT [SECONDARY].

### 01.AI — exit from foundation models
- **01.AI has exited foundation-model building** [SECONDARY]; the company now fine-tunes existing Chinese open models for an enterprise **"Boss AI"** positioning [SECONDARY].
- One more 2025–2026 lab-level strategy shift alongside the license-tightening trend (see §21).

### Perplexity — fine-tunes, not a 2026 base model
- **Sonar is a Llama 3.3 70B fine-tune** [SECONDARY]; **R1-1776 is a DeepSeek-R1 fine-tune** [SECONDARY] — Perplexity's models are derivatives, not from-scratch builds.
- **Agent API launched March 2026** [SECONDARY].
- **Sonar chat-completions endpoints are scheduled for retirement 2026-09-27** [SECONDARY].
- **No verified new-from-scratch 2026 Perplexity model** exists in the corpus.

