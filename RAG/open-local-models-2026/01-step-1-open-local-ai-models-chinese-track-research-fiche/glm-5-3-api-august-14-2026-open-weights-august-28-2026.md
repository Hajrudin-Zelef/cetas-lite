---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/glm-5-3-api-august-14-2026-open-weights-august-28-2026
title: "GLM-5.3 — API August 14, 2026; open weights August 28, 2026"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: actor-profile
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Huawei", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM", "xAI"]
dates: ["2025-05", "2026-02-12", "2026-03-18", "2026-04-22", "2026-06", "2026-08-14", "2026-08-26", "2026-08-28", "2026-09-18", "2026-09-21"]
keywords: ["glm", "open weights", "agent", "agentic", "attention", "aws", "bedrock", "benchmarks", "compute", "cost", "cyber", "deepseek"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [118, 170]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: 7748c9817b27b6b9d2a60a29db8bcdade7d18b84871766f1343bc1822b6d1148
---

# GLM-5.3 — API August 14, 2026; open weights August 28, 2026

### GLM-5.3 — API August 14, 2026; open weights August 28, 2026
- **Architecture:** 753B/~40B active MoE (78 layers); **same base as GLM-5.2 — gains from post-training only**; text-only; 1M context.
- **Staged release:** weights **withheld ~2 weeks for "safety evaluation and hardening"** (emergent offensive-cyber capability, CyberGym 84.5); shipped on HF (`zai-org/GLM-5.3`, `-BF16`) **Aug 28, 2026**, exactly the promised 14 days; safety work reportedly found 2,436 vulns across 269 OSS projects. 756GB native FP8; fits 8×H200, not 8×80GB H100.
- **License:** **NOT MIT.** Custom **GLM-5.3 License**: MIT-style permissions **plus security-review condition for MaaS operators with >$10B trailing-12-month revenue** (~500× looser than Kimi K3's $20M trigger).
- **Benchmarks:** DeepSWE **66.9** (vs 5.2's 46.2), CyberGym 84.5, AutomationBench 48.8 (vs Opus 4.8 41.0), AA v4.3 **45**. AA measures ~$2.15/M blended.
- **Distinguishing:** first explicit **staged open-weight release (pause → harden → ship)**; cyber-capability framing drove both the delay and the license change.

### GLM-5.3-Flash — August 26, 2026 (stealth-tested as "Ox Alpha")
- **Architecture:** **320B total / 18B active** MoE (45 layers); hybrid **KDA linear + NoPE sparse MLA** (~3× less attention compute, 4.4× smaller KV via IndexPool); **natively multimodal** (first in GLM-5 family); **1M context**; native FP8 (~306 GiB).
- **License:** **MIT**. HF: `zai-org/GLM-5.3-Flash` + `-BF16` (321B params).
- **Benchmarks:** DeepSWE v1.1 **63.4** (vs Opus 4.8 58.0), AutomationBench 48.8 (vs Opus 4.8 41.0), Terminal-Bench 2.1 **84.3** (vs Opus 4.8 85.0 — near-parity), AA Index v4.1.1 **57** at ~$0.045/task. Vision is the weak flank.
- **Pricing:** **$0.15/M in, $0.50/M out, $0.03/M cached**; launch half-price promo through Sept 9 (~1/33 in, ~1/50 out of Opus 4.8's $5/$25).
- **The 100,000-chip story:** stealth-tested as **"Ox Alpha"** on OpenRouter/OpenCode — became **most-used model on both, >62T tokens in 6 days (31% of OpenRouter weekly traffic)**. All inference on **>100,000 domestic Chinese accelerators** (Cambricon, Huawei, Moore Threads) with custom SGLang-based disaggregated engine; **Infra Agent powered by GLM-5.3 did most optimization** — first boot on domestic silicon to all production traffic in **~13 days, 3.2× throughput vs baseline**. Serving via SGLang, vLLM, TokenSpeed, KTransformers.

### GLM-5.3-FlashX — September 18, 2026
**Identical weights to GLM-5.3-Flash** (320B/18B, 1M) — a pure **inference-configuration** variant for the domestic-chip fleet. Output **200 tok/s** (~3.2–5× faster than Flash depending on metric); license/pricing presumed same as Flash.

---

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

