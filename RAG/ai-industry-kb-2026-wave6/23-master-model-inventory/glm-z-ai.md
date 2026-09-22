---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/glm-z-ai
title: "GLM / Z.ai"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["EU", "Huawei", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "United States", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2024-04-18", "2024-12-04", "2025-04-05", "2025-04-30", "2025-07-11", "2025-08-29", "2025-10-27", "2025-12-09", "2026-01-27", "2026-02-12", "2026-03-11", "2026-03-16", "2026-03-18", "2026-04-01", "2026-04-07", "2026-04-12", "2026-04-20", "2026-04-22", "2026-04-23", "2026-04-29", "2026-05-20", "2026-05-29", "2026-06-01", "2026-06-12", "2026-06-16", "2026-06-30", "2026-07-05", "2026-07-16", "2026-07-21", "2026-07-31", "2026-08-03", "2026-08-10", "2026-08-14", "2026-08-26", "2026-08-28", "2026-09-02", "2026-09-21", "2026-09-22"]
keywords: ["glm", "agent", "apache", "ascend", "attribution", "incident", "int4", "kimi", "license", "llama", "mistral", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11168, 11239]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 2c24f17e611a16f3c94889edf1b508ca28bf2b43112303f044335b010b3bd5c6
---

# GLM / Z.ai

### GLM / Z.ai

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| GLM-5.1 | 2026-04-07 (one source: May 7 — unresolved) | 744B/40B [VENDOR headline] | 200K | MIT · open [VENDOR] | Active; trained on Huawei Ascend 910B [VENDOR] |
| GLM-5.2 | 2026-06-16/17 | 744B/40B (same base as 5.3) | 1M | MIT · open [VENDOR] | Active; HF's forensic tool in the 2026-07-21 incident [SECONDARY] |
| GLM-5.3 | 2026-08-14/18 (weights 2026-08-28/29) | 744B headline / 753.3B HF tensor / ~743B vLLM [VENDOR] | 1M | Bespoke GLM-5.3 License | Active |
| GLM-5.3-Flash | 2026-08-26 | — | — | MIT · open [VENDOR] | Active ("Ox Alpha" stealth) |
| GLM-5V-Turbo | 2026-04-01 | 744B/~40B MoE [SECONDARY] | 200K | Closed/proprietary | API-only; vision kept commercial [SECONDARY] |
| GLM-4.5-Air | 2026 | — | — | MIT [VENDOR] | Active |

### Kimi / Moonshot

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Kimi K2 | 2025-07-11/12 | 1T/32B (8 of 384 experts) [SECONDARY] | 128K | Modified MIT [VENDOR] | Active (weights) |
| Kimi K2.5 | 2026-01-27 | 1T–1.06T/32B [SECONDARY] | 256K | Modified MIT | Active; hosted API deprecated 2026-05-20 (weights remain) [SECONDARY] |
| Kimi K2.6 | 2026-04-20 | 1T/32B, 384 experts [SECONDARY] | 256K | Modified MIT | Active |
| Kimi K2.7 Code | 2026-06-12 | 1T/32B (384 experts), native INT4 [SECONDARY] | 256K | Modified MIT | Active; coding-specialized post-train |
| Kimi K3 | 2026-07 (price date 2026-07-16) | — | — | Modified MIT | Active; LMArena coding #1 at 1,679 Elo [SECONDARY] |

### LongCat / Meituan

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| LongCat-2.0 (Preview 2026-04-20) | 2026-06-30 (weights 2026-07-05) | 1T-class (vendor claim; "zero NVIDIA" not independently verified) | — | MIT · open [VENDOR] | Active |
| LongCat-Flash series (6 models) | 2025-08-29 → 2026-03-11 | e.g. Flash-Lite 68.5B/~3B [SECONDARY] | — | MIT | Retired 2026-05-29 (sunset changelog) [VENDOR] |
| LongCat-Flash-Omni | ~2025-11 | 560B/27B [SECONDARY] | 128K | MIT | Retired 2026-05-29 |

### MiniMax

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| MiniMax M2 | 2025-10-27 | 230B/10B [SECONDARY] | — | MIT · open | Active |
| MiniMax M2.5 | 2026-02-12 | 229.9B/9.8B [SECONDARY] | ~192–196K | Modified MIT (UI attribution) | Active |
| MiniMax M2.7 | announced 2026-03-18; weights ~2026-04-12 | 230B/~10B [SECONDARY] | 200K | Non-commercial modified license | Active; "open source" → "open weights" rollback [SECONDARY] |
| MiniMax M3 | 2026-06-01 | 428B/~23B, MSA [SECONDARY] | 1M | MiniMax Community (no US/EU/UK/SK local deployment) | Active; API $0.60/$2.40 (promo $0.30/$1.20) [SECONDARY] |
| MiniMax H3 (Hailuo 3) | API 2026-07-31; weights 2026-08-03 | 33B dense [SECONDARY] | Omni-modal | MiniMax Community | Active |

### MiMo / Xiaomi

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| MiMo-7B family | 2025-04-30 | 7B [SECONDARY] | — | Weights MIT / code Apache 2.0 | Active |
| MiMo-V2.5 | 2026-04-22 | 310B [SECONDARY] | 1M | MIT · open [VENDOR] | Active |
| MiMo-V2.5-Pro | 2026-04-22 | 1.02T/42B [SECONDARY] | 1M | MIT · open [VENDOR] | Active |
| MiMo-V2.6-Pro | announced 2026-09-21 (blog 2026-09-22) | 1.02T/42B [SECONDARY] | — | MIT · open [VENDOR] | Active; AA v4.3 46.32 [SECONDARY]; distinct generation from V2.5-Pro |
| MiMo-V2.6-Flash | 2026-09-21/22 | 310B/15B [SECONDARY] | — | MIT · open [VENDOR] | Active |
| MiMo-V2.5-TTS-Series | 2026-04-23 | — | — | — | Limited-time free launch [SECONDARY] |

### Meta

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Llama 3 / 3.1 / 3.2 / 3.3 | 2024-04-18 → 2024-12-04 | 8B–405B [VENDOR] | 8K–128K | Llama Community License | Legacy, weights available |
| Llama 4 Scout | 2025-04-05 | — | 10M | Llama 4 Community (EU multimodal exclusion) | Active |
| Llama 4 Maverick | 2025-04-05 | — | 1M | Llama 4 Community (EU multimodal exclusion) | Active |
| Llama 4 Behemoth | — | ~2T/~288B (preview) [SECONDARY] | — | — | Never released; effectively shelved [UNVERIFIED] |
| Muse Glimmer 30B | 2026-08-10 | 30B | — | Apache 2.0 · open | Active; Meta's first straight Apache-2.0 model [SECONDARY] |
| Muse Spark 1.1 | 2026 | — | — | Closed API | Active; SWE-bench Pro official 61.5% (w/ mini-SWE-agent) [SECONDARY] |
| Muse Spark 1.3 | 2026-09-02 | — | — | Closed API | Active; $1.25/$4.25; AA v4.3 48 [SECONDARY] |

### Mistral

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Mistral Large 3 | 2025-12 | Undisclosed | — | — | Active |
| Mistral Medium 3.5 | 2026-04-29/30 | — | — | — | Active |
| Mistral Small 4 | 2026-03-16 | 119B/6B [SECONDARY] | 256K | Apache 2.0 | Active |
| Devstral 2 | 2025-12-09/10 | — | — | — | Active (distinct from Devstral 2507, Jul 2025) |
| Ministral 3 (2026) | — | — | — | — | [UNVERIFIED] single-source; do not merge with Ministral 3B/8B |

