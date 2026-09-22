---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/poolside
title: "Poolside"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "Baidu", "ByteDance", "Cohere", "DeepSeek", "Falcon", "Google", "Groq", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Perplexity", "Poolside", "StepFun", "TII", "Xiaomi", "Z.ai", "xAI"]
dates: ["2025-02-26", "2025-03-13", "2025-04-29", "2025-04-30", "2025-07-08", "2025-11-30", "2025-12-15", "2025-12-31", "2026-01-05", "2026-02-12", "2026-03-18", "2026-03-31", "2026-04-20", "2026-04-22", "2026-05", "2026-05-20", "2026-06-01", "2026-06-24", "2026-06-28", "2026-07-05", "2026-07-21", "2026-07-23", "2026-07-24", "2026-07-27", "2026-07-31", "2026-08-02", "2026-08-06", "2026-08-10", "2026-08-12", "2026-08-13", "2026-08-14", "2026-08-26", "2026-08-31", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-10", "2026-09-21", "2026-09-22", "2026-10-15"]
keywords: ["apache", "astra", "benchmark", "benchmarks", "claude", "cohere", "cost", "deepseek", "distribution", "embeddings", "glm", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11317, 11428]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 3c72efe747f51617bddb53aa4c80d5e0e4b246161904edb68a8bd326eab19fe2
---

# Poolside

### Poolside

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Laguna S 2.1 | 2026-07-21 | 118B/8B MoE [VENDOR] | 1M (1,048,576) | OpenMDW-1.1 · open | Active; TB 2.1 70.2% [VENDOR]; single DGX Spark |
| Laguna XS 2.1 | ~2026-07 | — | — | OpenMDW-1.1 · open | Active; smaller sibling |

### ByteDance

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Seed 2.1 Turbo | 2026-06-24 | Undisclosed | 262,144 | Proprietary | Active; $0.50/$2.50 [SECONDARY, single-sourced] |
| Seed 2.1 Pro (260628) | 2026-06-28 | Undisclosed | — | Proprietary | Active |

### InclusionAI

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Ling 3.0 Flash (+ Fin/Sante/VL) | 2026-07-23 (Fin 2026-09-03; Sante 2026-09-04) | 124B/~5.1B MoE [SECONDARY] | 262,144 | License UNVERIFIED (family historically Apache 2.0) | Active; free :free route on OpenRouter (rate-limited) |

### StepFun / Baidu

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Step 5 Preview | 2026-09 | — | — | API-only | Active; $1.00/$2.70; weights promised 2026-10-15 [SECONDARY] |
| ERNIE 5.0 | 2026 | — | — | API | Active; $0.85/$3.40 [SECONDARY] |
| ERNIE 5.1 | 2026 | — | — | API | Active; $0.59/$2.65 [SECONDARY] |

### Perplexity

| Model | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|
| Perplexity 2026 models (Sonar line) | 2026 | — | — | Proprietary | 2026 model posture verified; no flagship open release [SECONDARY] |

### Others

| Model | Lab | Date | Params (total/active) | Context | License/weights | Status |
|---|---|---|---|---|---|---|
| Ideogram 4 | Ideogram | 2026 | — | — | Code Apache 2.0; weights non-commercial | Active; weights NOT open [SECONDARY] |
| Hunyuan3D 3.0 | Tencent | 2026 | — | — | Free hosted access | Not open weights (open checkpoints in 2.x/Omni) [SECONDARY] |
| Hy3 Preview / Hy3 full / Hy4 Preview | Tencent | 2026 | — | — | Preview: Tencent Hy Community License; full/Hy4: Apache 2.0 | Do not label Hy3 Preview Apache [SECONDARY] |
| Baichuan M3-235B | Baichuan | 2026 | 235B (Qwen3-235B base) [SECONDARY] | — | Apache 2.0 | Medical model; active |
| Midjourney V8.1 | Midjourney | 2026-04 (14 vs 30 — unresolved) | — | — | Proprietary | Alpha; date conflict retained [SECONDARY] |
| Eleven v3 / v4 | ElevenLabs | v3 active; v4 previewed 2026-06 | — | — | Proprietary | v4 not shipped as of Sept 2026 [SECONDARY] |
| Runway Gen-4.5 / Gen-5 | Runway | Gen-4.5 2025-12 | — | — | Proprietary | Gen-4.5 flagship; Gen-5 not shipped [SECONDARY] |
| PlayAI / PlayHT engine | PlayAI | shut down 2025-12-31 | — | — | — | Dead; Meta acqui-hire Jul 2025; Groq retired the engine [SECONDARY] |


### New verified metrics — expansion

- Entry count: ~116 model entries across 16 families (OpenAI 8, Anthropic 6, Google 8, xAI 2, DeepSeek 6, Qwen 16, Z.ai 6, Moonshot 3, MiniMax 4, Xiaomi/Meituan 5, Meta/Microsoft/HF 12, Mistral 15, embeddings/rerankers 19, other 6).
- License distribution across entries: MIT ~14, Apache 2.0 ~30, custom/community ~8, proprietary API-only ~30, non-commercial (CC BY-NC/Qwen Research) ~4, UNVERIFIED ~30 — nearly a third of entries have unverified license fields [DIRECTIONAL — count of this file's entries].
- Status distribution: current ~70, superseded/retired ~12 (Medium 3, Small 3.1, Pixtral Large, DeepSeek V4 alias, Phi-3, Mistral OCR legacy), staged/not-yet-open ~3 (GLM-5.3 non-Flash, K3 pre-7/27), UNVERIFIED ~30 [DIRECTIONAL].
- Benchmark-version hygiene: every scored entry names the benchmark version (DeepSWE 2026-09-22 vs May 2026; Terminal-Bench 2.1 official vs Ante harness; MTEB vs MTEB-BR vs MMTEB) — no merged generations [SECONDARY].

## Main actors

- **DeepSeek** — MIT across the V4 line; V4.1 Flash = cheapest frontier-class API.
- **Alibaba** — densest 2026 cadence (3.5 → 3.6 → 3.7 closed → 3.8 Max); vision closed, text open.
- **Z.ai** — 744B-class open MoEs (MIT for Flash/5.1); bespoke 5.3 license; vision closed.
- **Moonshot** — K2→K3 Modified MIT lineage; coding-arena #1.
- **Meituan** — LongCat Flash consolidation into MIT-licensed 2.0.
- **MiniMax** — license tightening ladder; M3/H3 under geo-excluding Community License.
- **Xiaomi** — MiMo MIT lineage; V2.6-Pro as a distinct generation.
- **Meta** — gated Llama 4 line vs Apache-2.0 Glimmer 30B; closed Muse Spark.
- **NVIDIA** — Nemotron 3 family; Model Factory license ($6B); Nemotron 4 in development.
- **Anthropic / OpenAI / Google / xAI** — closed flagships with dated price tables (§22).
- **TII, Cohere, Poolside, ByteDance, InclusionAI, StepFun, Baidu, Perplexity** — secondary labs with verified 2026 activity.

## Timeline and context

- **2025-12-15** — NVIDIA Nemotron 3 Nano (31.6B/3.2B MoE, first MoE in the line).
- **2026-01-05** — Falcon-H1R / Falcon-H1 Arabic (TII active in 2026).
- **2026-02-12** — MiniMax M2.5 (modified MIT); M2.7 announced 2026-03-18.
- **2026-03-31** — Gemma 4 weights (Apache 2.0).
- **2026-04-20** — Kimi K2.6; LongCat-2.0-Preview.
- **2026-04-22** — MiMo-V2.5 / V2.5-Pro (MIT).
- **2026-05-20** — Cohere Command A+ (Apache 2.0).
- **2026-06-01** — MiniMax M3 (428B/~23B, Community license); Nemotron 3 Ultra (550B/55B).
- **2026-07-05** — LongCat-2.0 weights published (MIT).
- **2026-07-21** — Poolside Laguna S 2.1 (OpenMDW-1.1).
- **2026-08-10** — Meta Muse Glimmer 30B (Apache 2.0).
- **2026-08-12** — Qwen3.8-Max weights (custom license).
- **2026-08-13** — DeepSeek V4-Pro-0813.
- **2026-09-02** — Muse Spark 1.3; Qwen3.8-Max-0902 refresh.
- **2026-09-03** — GPT-6 Astra launch.
- **2026-09-10** — DeepSeek V4.1 Flash GA.
- **2026-09-21** — Grok 4.7 launch; MiMo-V2.6-Pro announced.


### New verified timeline entries — expansion

- 2024-12: Phi-4 (14B) MIT [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- 2025-02-26: Phi-4-mini + Phi-4-multimodal [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- 2025-03-13: Gemma 3 [SECONDARY]. Source: https://themunicheye.com/google-launches-gemma-3-open-source-ai-model-13085
- 2025-04-29: Qwen3 dense family [SECONDARY]. Source: http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- 2025-04-30: Phi-4-reasoning family [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- 2025-07-08: SmolLM3-3B [SECONDARY]. Source: https://github.com/huggingface/blog/blob/HEAD/smollm3.md
- 2026-04-20: Kimi K2.6 [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- 2026-06-01: MiniMax M3 [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- 2026-07-05: LongCat-2.0 MIT [SECONDARY]. Source: https://github.com/XUEXUE-XING/minimax-h3-webui/pull/12
- 2026-07-24: Opus 5/Sonnet 5/Haiku 4.5 launch; DeepSeek V4 alias retirement [SECONDARY]. Sources: https://medium.com/@automation.labs/opus-5-sonnet-5-haiku-4-5-which-claude-model-for-which-job-bce5e8346233 and wave6/03-model-weights-wave3.md
- 2026-07-27: Kimi K3 weights commitment date [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- 2026-07-31: DeepSeek V4.1 Flash (0731) [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- 2026-08-02: MiniMax H3 license effective [SECONDARY]. Source: https://github.com/onigirikiller/minimax-h3-webui
- 2026-08-06: MiniMax Music 3 license effective [SECONDARY]. Source: https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3
- 2026-08-14: GLM-5.3 launch (closed weights) [SECONDARY]. Source: https://emergent.sh/learn/what-is-glm-5-3
- 2026-08-26: GLM-5.3-Flash MIT [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- 2026-08-31: Mistral Medium 3 retired [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- 2026-09-21/22: MiMo V2.6-Pro [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- 2025-11-30: Mistral Small 3.1 retired (predates 2026 — included for line hygiene) [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/

