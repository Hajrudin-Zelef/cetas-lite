---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/v2-6-live-rl-training-metrics
title: "V2.6 Live RL training metrics"
domain: mimo-and-xiaomi
role: deep-dive
task: training
actors: ["China", "DeepSeek", "Hugging Face", "OpenRouter", "Xiaomi"]
dates: ["2025-04-30", "2025-05-30", "2026-03-11", "2026-03-18", "2026-03-19", "2026-04-22", "2026-04-23", "2026-05-27", "2026-06-30", "2026-09-21", "2026-09-22"]
keywords: ["training", "agent", "apache", "attention", "benchmark", "benchmarks", "cost", "deepseek", "disclosure", "inference", "license", "omni"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3930, 3997]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 7dbf657db6dfce45ab76f5903c3daa784a8d19b44c5d6ff339ebece44d0ad8be
---

# V2.6 Live RL training metrics

### V2.6 Live RL training metrics
- RL configuration: 1,568 prompts/step × 16 rollouts/prompt = 25,088 rollouts per step. [VENDOR, S24][SECONDARY, S25]
- Tokens per RL step: ~3.5B–3.7B. [VENDOR, S24]
- DeepSWE v1.1 (held-out, pre→post RL): Flash 48.8 → 65.68; Pro 58.4 → 72.57. [VENDOR, S21]
- V2.6-Pro vendor launch table: DeepSWE v1.1 71.9; AutomationBench v1.0.6 53.1; Toolathlon-Verified 76.9; Terminal-Bench 2.1 89.9; JobBench 62.0; CyberGym 94.0; MiMo Visual Coding 72.3. [VENDOR, S25]
- Reported training cost: Flash ~$0.85M; Pro ~$2.62M (~3.1×). [SECONDARY, S24]
- Checkpoint uploads: ~2026-09-21 15:39 UTC, ~18 seconds apart. [COMMUNITY, S24]
- Artificial Analysis v4.3 (cross-reference to base section): V2.6-Pro 46.32. [SECONDARY, S20]

### Xiaomi corporate figures
- AI investment commitment: ≥CNY60B (~$8.7B) over 3 years, announced 2026-03-19. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S11]
- 2026 AI R&D budget: >CNY16B (~$2.3B). [SECONDARY, S11]
- Five-year core-technology budget: CNY200B (chips, AI, OS). [SECONDARY, S11]
- Stock reaction: +~5% (Xiaomi 1810.HK). [SECONDARY, S12]

## Main actors
- **Xiaomi** — MiMo's owner; positions MiMo inside its "Human × Car × Home" ecosystem [SECONDARY].
- **MiMo** — Xiaomi's model family (text, omni, TTS variants) [SECONDARY].
- **Artificial Analysis** — the independent benchmark operator whose v4.3 Index scored V2.6-Pro at 46.32 [SECONDARY].
- **The "Hunter Alpha" watchers** — the community that wrongly suspected the stealth codename was DeepSeek V4; it was MiMo-V2-Pro [SECONDARY].

## Timeline and context
- **2025-04-30** — MiMo-7B family released under MIT (weights) / Apache 2.0 (repo code) [SECONDARY].
- **2026-03-18** — V2-Pro, V2-Omni, V2-TTS launched (proprietary); "Hunter Alpha"→V2-Pro, "Healer Alpha"→V2-Omni [SECONDARY].
- **2026-04-22** — V2.5 (310B, MIT) and V2.5-Pro (1.02T/42B, MIT) released; 1M context [SECONDARY].
- **2026-04-23** — V2.5-TTS-Series released, API-only, initially limited-time free [SECONDARY].
- **2026-09-21/22** — V2.6-Pro (natively omnimodal, 1.02T/42B, MIT, Live-RL-trained) and V2.6-Flash (310B/15B, MIT) released [SECONDARY].


### New verified timeline entries — expansion

- 2026-03-11: Anonymous "Hunter Alpha" appears on OpenRouter, tops daily usage charts. [SECONDARY, S27]
- 2026-03-18: Xiaomi reveals Hunter Alpha was MiMo-V2-Pro; stock +5.8%. [SECONDARY, S27][SECONDARY, S26]


- 2025-04-30: MiMo-7B base release (base section; anchor for the RL update below). [SECONDARY, S2][SECONDARY, S4]
- 2025-05-30: MiMo-7B-RL-0530 checkpoint posted to Hugging Face with 6M-instance SFT and 48K RL context. [VENDOR, S1]
- 2026-03-18: MiMo V2-Pro, V2-Omni, and V2-TTS released (base section; specification details added above). [SECONDARY, S5][SECONDARY, S6][SECONDARY, S7]
- 2026-03-19: Lei Jun announces ≥CNY60B / ~$8.7B AI investment over three years; Xiaomi shares rise ~5%. [SECONDARY, S9][SECONDARY, S10][SECONDARY, S11][SECONDARY, S12]
- 2026-04-22: MiMo V2.5 and V2.5-Pro released (base section; architecture and pricing deltas added above). [SECONDARY, S13][SECONDARY, S19]
- 2026-04-23: V2.5-TTS released (base section anchor). [SECONDARY, S17]
- 2026-05-27: Provider rate-card revision takes effect — V2.5-Pro $0.435/$0.87, V2.5 $0.14/$0.28, cache-hit tiers introduced. [SECONDARY, S14][SECONDARY, S16]
- 2026-06-30 (Beijing time): Legacy V2 model identifiers fully retired after deprecation/auto-routing period. [SECONDARY, S16]
- 2026-09-21 ~15:39 UTC: V2.6-Pro and V2.6-Flash Hugging Face checkpoints appear, ~18 seconds apart. [COMMUNITY, S24]
- 2026-09-21/22: V2.6-Pro and V2.6-Flash officially released with Live RL training disclosure, technical report, environments, and code. [SECONDARY, S21][SECONDARY, S22]
- 2026-09-22 (research date): MiClaw smartphone agent in closed beta; V2-TTS API variants still limited-time free per provider logs. [SECONDARY, S8][SECONDARY, S16]

## Implications
1. MiMo is the 2026 counterexample to license tightening: Xiaomi ships its strongest generations (V2.5-Pro, V2.6-Pro, V2.6-Flash) under plain MIT while the proprietary line (V2-Pro/V2-Omni/V2-TTS) stays closed — a two-track strategy, open weights for community reach and proprietary for product control [DIRECTIONAL].
2. V2.6-Pro's public "Live RL" (30 steps, ~750K trajectories, $2.62M) is unusually transparent post-training disclosure; treat cost/trajectory figures as secondary-sourced claims, not audited facts [SECONDARY/DIRECTIONAL].
3. AA v4.3 46.32 positions V2.6-Pro at the top of open-weight rankings on that index version — pin the methodology version whenever citing it, and never compare across v4.1.1/v4.2/v4.3 [SECONDARY].
4. Artifact-scope discipline matters for MiMo-7B: weights MIT vs repository code Apache 2.0 — quote the right artifact [SECONDARY].
5. V2.5-Pro and V2.6-Pro share a parameter profile but are distinct generations; consolidation or benchmarking work must not merge them [SECONDARY].
6. The "top open-weights model" framing of the V2.6 launch is press language — the citable, version-pinned figure is AA v4.3 46.32, nothing else [SECONDARY].


### New verified implications — expansion

- The MiMo-7B-RL-0530 update shows Xiaomi iterating on small reasoning models through post-training scale (12× SFT data, 1.5× RL context) rather than base-model retraining — the largest gains landed on math benchmarks (+11.9/+14.8 AIME points) while code gains were smaller (+3.1/+2.9 LiveCodeBench). [VENDOR, S1]
- Xiaomi's pricing trajectory is aggressively downward: V2-Pro launched at $1/$3 (≤256K) and $2/$6 (256K–1M), while V2.5-Pro fell to $0.435/$0.87 within five weeks of its own launch — a deliberate undercutting pattern consistent with the reported agent-pricing pressure narrative around Hunter Alpha. [SECONDARY, S5][SECONDARY, S14][SECONDARY, S8]
- The V2.6 "Live RL" disclosure — asynchronous GRPO at ~3.5B tokens per step with released environments and code — is unusually transparent for a frontier open-weights lab and directly targets the reproducibility gap that closed labs maintain. [VENDOR, S21][SECONDARY, S22]
- The ~$0.85M (Flash) and ~$2.62M (Pro) reported RL training costs reframe V2.6 as a post-training story: the base models already existed, and the headline capability came from a comparatively cheap RL phase — a pattern that favors labs with strong infrastructure over labs with the largest pretraining budgets. [SECONDARY, S24]
- The V2 deprecation hygiene is notable: legacy identifiers were auto-routed to V2.5 endpoints before the hard 2026-06-30 retirement, a migration pattern that reduces breaking changes for API consumers. [COMMUNITY, S16] The Beijing-time cutoff itself is single-sourced to provider logs.
- The V2.5 token-count tension (48T for the 310B model vs 27T for the 1.02T Pro) is consistent with a strategy of over-training smaller models for inference efficiency, but neither vendor nor secondary sources state this explicitly, so it remains a hypothesis, not a claim.
- V2-TTS's >100M-hour speech pretraining claim, if taken at face value, would make it one of the largest speech-pretraining runs publicly claimed by any lab. [VENDOR — single source at launch reporting; treat as a vendor claim until the technical report is verified.]
- The 6:1 sliding-window-to-global attention ratio with a 128-token local window in V2.5-Pro is one of the most specific open-weights attention disclosures of 2026, useful for inference-engine implementers even though layer counts remain undisclosed.
- MiClaw's closed beta is Xiaomi's first disclosed attempt at a system-level smartphone agent, distinct from app-level assistants — its integration into Xiaomi OS rather than a standalone app suggests the company views the OS, not the chatbot, as the agent platform.

