---
id: collect-tutoriels/tutoriels/mimo-2-6-vs-mimo-v2-6-2
title: "MiMo-2.6 vs MiMo-V2.6 : deux orthographes, une seule version, et un troisième nom qui correspond vraiment à un modèle différent"
domain: tutoriels
role: reference
task: article
actors: ["Alibaba", "Xiaomi"]
dates: ["2026-09"]
keywords: ["agent", "attention", "benchmark", "context window", "cost", "embedding", "fp8", "license", "mit license", "moe", "omni", "open-weight"]
source: docs/RAG/Collect RAG/07_tutoriels/mimo-2-6-vs-mimo-v2-6.md
source_anchor: ""
source_lines: [28, 68]
sha256: 750e872102840556c0e80006be16f11d6c60d63599499fbbbddb5edb50fceafd
---

# MiMo-2.6 vs MiMo-V2.6 : deux orthographes, une seule version, et un troisième nom qui correspond vraiment à un modèle différent

- MiMo-2.6 and MiMo-V2.6 are the same version written two ways; not two products.
- Real models: MiMo-V2.6-Pro (1,020B flagship) and MiMo-V2.6-Flash (309B efficiency), open-sourced 21 September 2026, MIT license.
- "2.6" is a generation index, not parameter count, size class, or capability tier.
- The dangerous collision is MiMo-V2-Flash (Dec 2025), a different download with different serving recipes and 256K context vs the new 1M.
- Only independently produced figure: MiMo-V2.6-Pro scores 46 on Artificial Analysis Intelligence Index v4.3.2 (highest open-weights entry).
- Xiaomi vendor figures: 71.9 DeepSWE v1.1 for Pro (vs 19.0 previous gen) and 67.9 for Flash.
- Training run streamed live; costs ~$2.62M (Pro) and ~$854K (Flash), 30 RL steps, ~750,000 trajectories, under six days; failures published too.
- Verify repo name (-RL), Safetensors param count, config.json drafter, and exact lowercase API IDs.

## Technical data / figures

| Item | MiMo-V2.6-Pro | MiMo-V2.6-Flash |
| --- | --- | --- |
| Claimed total params | 1,020B (1.02T) | 309B |
| Safetensors param count | 524B | 159B |
| Context window | 1M tokens | 1M tokens |
| Max positional embedding | 1,048,576 tokens | 1,048,576 tokens |
| Max output | 128K tokens | 128K tokens |
| Input modalities | text, image, video, audio | text, image, video, audio |
| License | MIT | MIT |
| Release | 21 Sept 2026 | 21 Sept 2026 |
| Weight size | — | 172.9 GB FP8, 65 shards |
| DeepSWE v1.1 (vendor) | 71.9 | 67.9 |
| Artificial Analysis Index | 46 | — |
| API identifier | `mimo-v2.6-pro` | `mimo-v2.6-flash` |

Generation timeline:

| Generation | Date | Key checkpoint / notes |
| --- | --- | --- |
| MiMo-V2 | Dec 2025 | MiMo-V2-Flash: 309B total / 15B active MoE, 256K context, hybrid attention, 128-token sliding window |
| MiMo-V2.5 | — | Previous generation; Pro scored 19.0 on DeepSWE v1.1 |
| MiMo-V2.6 | 21 Sept 2026 | Pro + Flash, omni-modal, 1M context, MIT |

Training run figures: Pro ~$2.62M, Flash ~$854K; 30 RL steps each; ~750,000 trajectories; completed in under six days. Dashboard DeepSWE: Pro 72.57 (mini-swe-agent, avg of 3) → final card 71.9; Flash 65.68 → 67.9.

Other IDs: `mimo-v2.6-pro-ultraspeed` (service tier), MiMo-V2.6-Distill-Qwen-9B. Config discrepancy: Flash card says 5-layer drafter predicting 7 tokens/pass, but config sets `num_nextn_predict_layers` = 3.

## Why this source matters for the RAG

It is a precise disambiguation guide for a confusingly named model family, with concrete repository-verification steps, benchmark provenance, and training-cost disclosures. It is valuable for RAG questions about MiMo/Xiaomi models, model naming pitfalls, and how to verify open-weight artifacts before deployment.
