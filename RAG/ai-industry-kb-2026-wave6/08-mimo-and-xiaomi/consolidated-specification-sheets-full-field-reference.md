---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/consolidated-specification-sheets-full-field-reference
title: "Consolidated specification sheets (full-field reference)"
domain: mimo-and-xiaomi
role: deep-dive
task: actor-profile
actors: ["China", "Hugging Face", "Xiaomi"]
dates: ["2025-04-30", "2025-05-30", "2026-03-18", "2026-04-22", "2026-04-23", "2026-05-27", "2026-06-30", "2026-09-21", "2026-09-22"]
keywords: ["apache", "attention", "benchmarks", "compute", "cost", "distillation", "fp8", "license", "moe", "multimodal", "omni", "parameters"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3708, 3791]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 703b3c586e1fc62fcd9c767efdda89cf0458449d1235054e516716014147c94d
---

# Consolidated specification sheets (full-field reference)

### Consolidated specification sheets (full-field reference)
Fields marked "(base §8 anchor)" repeat the base section only as dated anchors inside a complete spec sheet; every other field is new.

- **MiMo-7B-RL-0530** — Checkpoint update, not a new base model. [VENDOR, S1][SECONDARY, S2]
- Release/GA date: 2025-05-30 (Hugging Face). [VENDOR, S1]
- Base model: MiMo-7B, released 2025-04-30, ~25T pretraining tokens (base §8 anchor). [SECONDARY, S2][SECONDARY, S4]
- Architecture: dense 7B reasoning model; layer count not publicly disclosed. [VENDOR, S1]
- Parameters: 7B total (dense; active = total). [VENDOR, S1]
- Context: RL training window 48K tokens (up from 32K). [VENDOR, S1]
- Training tokens/compute: base ~25T (base §8 anchor); post-training on ~130K math/code RL problems and 6M-instance SFT. [SECONDARY, S2][VENDOR, S1]
- License: MIT weights / Apache-2.0 repository code (base §8 anchor). [VENDOR, S1]
- API pricing: none found — open-weights release with no first-party API price card. [DIRECTIONAL]
- Benchmarks, dated 2025-05-30, vendor: AIME 2024 80.1; AIME 2025 70.2; LiveCodeBench v5 60.9; LiveCodeBench v6 52.2; GPQA Diamond 60.6. [VENDOR, S1]
- Integrations: Hugging Face (XiaomiMiMo organization). [VENDOR, S1]
- Retirement: none announced. [DIRECTIONAL]
- **MiMo V2-Pro** — Flagship MoE, released 2026-03-18 (base §8 anchor). [SECONDARY, S5][SECONDARY, S7]
- Architecture: mixture-of-experts with hybrid/mixed attention; layer count not publicly disclosed. [SECONDARY, S5]
- Parameters: ~1T total / ~42B active per token. [SECONDARY, S5][SECONDARY, S7]
- Context: 1,048,576 tokens; maximum output ~32K tokens. [SECONDARY, S7][SECONDARY, S5]
- Training tokens/compute: not publicly disclosed. [DIRECTIONAL]
- License: MIT (base §8 anchor). [SECONDARY, S5]
- API pricing (launch rate card): $1/M input, $3/M output up to 256K; $2/M input, $6/M output from 256K to 1M. [SECONDARY, S5]
- API pricing (base §8 record): $0.80/$3.20 — preserved as a separate dated record; contradiction documented below. [SECONDARY, S5]
- Benchmarks: vendor launch benchmarks not independently corroborated; AA score of 49 excluded (no methodology version). [excluded per rule]
- Integrations: MiMo Studio; Xiaomi Browser; WPS/Kingsoft Office; OpenClaw; OpenCode; Cline. [SECONDARY, S5][SECONDARY, S7]
- Retirement: legacy V2 identifiers deprecated and auto-routed, fully retired 2026-06-30 Beijing time. [SECONDARY, S16]
- **MiMo V2-Omni** — Multimodal model, released 2026-03-18 (base §8 anchor). [SECONDARY, S6][SECONDARY, S7]
- Architecture: multimodal MoE; internal details not publicly disclosed. [SECONDARY, S6]
- Parameters: not publicly disclosed. [DIRECTIONAL]
- Context: 256K tokens. [SECONDARY, S6]
- Input modalities: text, image, video, audio. [SECONDARY, S6][SECONDARY, S7]
- Long-audio understanding: beyond 10 hours continuous. [SECONDARY, S6]
- Training tokens/compute: not publicly disclosed. [DIRECTIONAL]
- License: MIT (base §8 anchor). [SECONDARY, S6]
- API pricing: $0.40/M input, $2/M output (reported). [SECONDARY, S6]
- Retirement: legacy V2 identifiers retired 2026-06-30 Beijing time. [SECONDARY, S16]
- **MiMo V2-TTS** — Speech model, released 2026-03-18 (base §8 anchor). [SECONDARY, S6][SECONDARY, S7]
- Architecture: joint speech-text modeling; proprietary audio tokenizer; multi-codebook acoustic model. [SECONDARY, S6]
- Speech pretraining: >100M hours (vendor claim). [VENDOR, S6]
- Parameters: not publicly disclosed. [DIRECTIONAL]
- License: MIT (base §8 anchor). [SECONDARY, S6]
- API pricing: API variants remained limited-time free per provider logs. [SECONDARY, S16]
- V2.5-TTS variant released 2026-04-23 (base §8 anchor). [SECONDARY, S17]
- **MiMo V2.5-Pro** — Sparse MoE flagship, released 2026-04-22 (base §8 anchor). [VENDOR, S13][SECONDARY, S19]
- Architecture: sparse MoE; alternating sliding-window/global attention at 6:1 ratio; 128-token local window; lightweight MTP module. [VENDOR, S13]
- Parameters: 1.02T total / 42B active. [VENDOR, S13][SECONDARY, S19]
- Layers: not publicly disclosed. [DIRECTIONAL]
- Context: 1M tokens; API maximum output 128K. [SECONDARY, S14]
- Training: ~27T tokens; FP8 precision; native 32K pretraining sequence before extension. [VENDOR, S13]
- License: MIT (base §8 anchor). [SECONDARY, S19]
- API pricing (2026-05-27 rate card): $0.435/M input, $0.87/M output; cache hit $0.0036/M. [SECONDARY, S14][SECONDARY, S16]
- Integrations: OpenClaw; OpenCode; Cline. [SECONDARY, S5]
- Status as of 2026-09-22: current. [DIRECTIONAL]
- **MiMo V2.5** — Sparse multimodal MoE, released 2026-04-22 (base §8 anchor). [SECONDARY, S18][SECONDARY, S19]
- Architecture: sparse multimodal MoE; separate vision and audio encoders with projection modules. [SECONDARY, S13]
- Parameters: 310B total / 15B active. [SECONDARY, S18][SECONDARY, S19]
- Layers: not publicly disclosed. [DIRECTIONAL]
- Context: 1M tokens; API maximum output 128K. [SECONDARY, S14]
- Training: ~48T tokens (reported; tension with V2.5-Pro's 27T documented below). [SECONDARY, S13]
- License: MIT (base §8 anchor). [SECONDARY, S19]
- API pricing (2026-05-27 rate card): $0.14/M input, $0.28/M output; cache hit $0.0028/M. [SECONDARY, S14][SECONDARY, S15]
- Status as of 2026-09-22: current. [DIRECTIONAL]
- **MiMo V2.6-Pro** — Live-RL flagship, released 2026-09-21/22 (base §8 anchor). [SECONDARY, S21][SECONDARY, S22]
- Architecture: MoE; RL recipe = fully asynchronous GRPO with Groupwise Reward Synthesis, Groupwise Advantage Redistribution, and MOPD2 on-policy distillation. [VENDOR, S21][VENDOR, S24]
- Parameters: not publicly disclosed; secondary coverage frames it as 1T-class. [SECONDARY, S24]
- Layers: not publicly disclosed. [DIRECTIONAL]
- Context: up to 1M tokens in the RL run. [VENDOR, S24]
- RL configuration: 1,568 prompts/step × 16 rollouts/prompt; ~3.5B–3.7B tokens/step. [VENDOR, S24][SECONDARY, S25]
- RL training cost: ~$2.62M. [SECONDARY, S24]
- License: MIT (base §8 anchor). [SECONDARY, S21]
- Benchmarks (vendor launch table): DeepSWE v1.1 71.9; AutomationBench v1.0.6 53.1; Toolathlon-Verified 76.9; Terminal-Bench 2.1 89.9; JobBench 62.0; CyberGym 94.0; MiMo Visual Coding 72.3. [VENDOR, S25]
- Benchmarks (held-out RL deltas): DeepSWE v1.1 58.4 → 72.57. [VENDOR, S21]
- Third-party: Artificial Analysis v4.3 score 46.32 (base §8 cross-reference). [SECONDARY, S20]
- Artifacts: technical report, RL environments, and training code released. [SECONDARY, S22][SECONDARY, S21]
- **MiMo V2.6-Flash** — Cheaper RL sibling, released 2026-09-21/22 (base §8 anchor). [SECONDARY, S21][SECONDARY, S23]
- Architecture: MoE; same asynchronous GRPO RL recipe as Pro. [VENDOR, S21]
- Parameters: not publicly disclosed; smaller than Pro (exact undisclosed). [DIRECTIONAL]
- Context: 1M tokens. [VENDOR, S24]
- RL training cost: ~$0.85M (~3.1× cheaper than Pro). [SECONDARY, S24]
- Benchmarks (held-out RL deltas): DeepSWE v1.1 48.8 → 65.68. [VENDOR, S21][SECONDARY, S25]
- License: MIT (base §8 anchor). [SECONDARY, S21]
- Positioning: cheaper alternative released alongside Pro. [SECONDARY, S21]


