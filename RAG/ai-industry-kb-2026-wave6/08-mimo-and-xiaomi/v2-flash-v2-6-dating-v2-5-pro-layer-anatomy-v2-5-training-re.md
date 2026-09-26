---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/v2-flash-v2-6-dating-v2-5-pro-layer-anatomy-v2-5-training-re
title: "V2-Flash, V2.6 dating, V2.5-Pro layer anatomy, V2.5 training recipe (corroborated 2026-09)"
domain: mimo-and-xiaomi
role: deep-dive
task: training
actors: ["Alibaba"]
dates: ["2025-04-30", "2025-05", "2025-05-30", "2025-12-17", "2026-03-18", "2026-04", "2026-04-22", "2026-04-23", "2026-05-27", "2026-06-30", "2026-09-21", "2026-09-22"]
keywords: ["training", "agentic", "apache", "attention", "benchmark", "context window", "cost", "fine-tuning", "fp4", "fp8", "gqa", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3837, 3929]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: aa9232d290f145e1a2b2abcec3c4610821ed724f69b9f2f3d62e5f026d649005
---

# V2-Flash, V2.6 dating, V2.5-Pro layer anatomy, V2.5 training recipe (corroborated 2026-09)

### V2-Flash, V2.6 dating, V2.5-Pro layer anatomy, V2.5 training recipe (corroborated 2026-09)
- MiMo-V2-Flash (launched 2025-12-17): open-source MoE with 309B total / 15B active parameters, trained on 27T tokens with FP8 mixed precision, hybrid attention at a 5:1 SWA:GA ratio — a distinct ratio from V2-Pro's reported 7:1 and V2.5-Pro's 6:1. [SECONDARY, S4 — single source]
- Knowledge cutoffs: MiMo-V2-Pro, V2-Omni, V2.5, and V2.5-Pro all carry a May 2025 cutoff per the model table. [SECONDARY, S4 — single source]
- The model table dates the V2.6-Flash / V2.6-Pro pair to 2026-09-21 — corroborating the base section's V2.6 timing. [SECONDARY, S4 — single source]
- MiMo-7B-RL-0530: scaled the fine-tuning dataset from 500,000 to 6 million instances and extended the RL window from 32,000 to 48,000 tokens; AIME 2024 improved from 68.2 to 80.1. [SECONDARY, S4 — single source]
- V2.5-Pro layer anatomy (vendor HF card): 70 layers (1 dense + 69 MoE); 10 full-attention + 60 SWA layers; 128 attention heads, 8 KV heads (GQA); 384 routed experts, 8 per token; MoE intermediate 2048; dense intermediate 16384 (layer 0 only); head dim 192 (QK) / 128 (V). [VENDOR, S30 — single source]
- V2.5 layer anatomy (vendor HF card): 48 layers (1 dense + 47 MoE); 9 full-attention + 39 SWA layers; 64 heads; 8 KV heads (global) / 4 (SWA); 256 routed experts, 8 per token; 128-token SWA window; 1M max context; 3 MTP layers. [VENDOR, S30 — single source]
- V2.5 five-stage training recipe: (1) text pretraining on 48T tokens; (2) projector warmup aligning in-house audio/visual encoders; (3) multimodal pretraining on cross-modal data; (4) agentic post-training with context 32K→1M. [SECONDARY, S19 — single source]
- V2.5-Pro standard API pricing: $1 per million input tokens, $3 per million output tokens. [SECONDARY, S42 — single source]
- `mimo-v2.5-pro-ultraspeed`: FP4-quantized variant at 1000 tokens/second, no vision/audio. [SECONDARY, S15 — single source]
- CNY price card corroboration (2026-06): mimo-v2.5 at ¥0.02/¥1/¥2 (cache-hit/in/out); mimo-v2.5-pro at ¥0.025/¥3/¥6; TTS free (limited time); ASR ¥0.5/hour — consistent with USD tiers. [SECONDARY, S15 — single source]
- TTS/ASR model ids: `mimo-v2.5-tts` (voice control), `mimo-v2.5-tts-voiceclone`, `mimo-v2.5-tts-voicedesign`, `mimo-v2.5-asr` (bilingual + dialects, lyrics). [SECONDARY, S15 — single source]
- The aidirectortoolkit catalog confirms V2 identifiers auto-routed to V2.5 endpoints with V2.5 pricing, full deprecation 2026-06-30. [SECONDARY, S15][SECONDARY, S16]

## Figures and metrics
| Model | Params (total/active) | Context | License | Release |
|---|---|---|---|---|
| MiMo-7B | 7B | — | MIT weights / Apache-2.0 repo code | 2025-04-30 [SECONDARY] |
| V2-Pro / V2-Omni / V2-TTS | — | — | Proprietary | 2026-03-18 [SECONDARY] |
| V2.5 | 310B | 1M (256K std) | MIT | 2026-04-22 [SECONDARY] |
| V2.5-Pro | 1.02T / 42B | 1M (256K std) | MIT | 2026-04-22 [SECONDARY] |
| V2.6-Pro | 1.02T / 42B | 1M | MIT | 2026-09-21/22 [SECONDARY] |
| V2.6-Flash | 310B / 15B | — | MIT | 2026-09-22 window [SECONDARY] |

- V2.6-Pro Live RL: ~6 days, 30 steps, ~750K trajectories, ~$2.62M cost [SECONDARY].
- V2.5-Pro API pricing: $0.80/$3.20 per million input/output, $0.16 cache reads (80% off) [SECONDARY].
- Artificial Analysis Intelligence Index v4.3 score for V2.6-Pro: **46.32** [SECONDARY]. Never compare this with scores from other AA methodology versions; v4.1.1→v4.2→v4.3 are not cross-comparable.
- V2.5-TTS launch pricing: limited-time free API access [SECONDARY].

| Variant | Release | Access | Price |
|---|---|---|---|
| mimo-v2.5-tts | 2026-04-23 | API-only | limited-time free [SECONDARY] |
| mimo-v2.5-tts-voicedesign | 2026-04-23 | API-only | limited-time free [SECONDARY] |
| mimo-v2.5-tts-voiceclone | 2026-04-23 | API-only | limited-time free [SECONDARY] |

- V2.5-TTS keeps voice API-gated while text weights go open — the same product/control split as the proprietary V2-Pro/V2-Omni line [SECONDARY/DIRECTIONAL].


### New verified metrics — expansion

### V2.5-Pro / V2.5 / V2-Flash additional figures
- V2.5-Pro: 70 layers (1 dense + 69 MoE); 384 experts, 8/token; 128 heads, 8 KV (GQA). [VENDOR, S30]
- V2.5: 48 layers; 256 experts, 8/token; 64 heads; 128-token SWA window. [VENDOR, S30]
- V2.5-Pro standard pricing $1/$3 per 1M; ultraspeed FP4 variant 1000 tok/s. [SECONDARY, S42][SECONDARY, S15]
- V2-Flash: 309B/15B; 27T tokens; 5:1 hybrid attention. [SECONDARY, S4]
- RL-0530: SFT 500K→6M instances; RL window 32K→48K; AIME 2024 68.2→80.1. [SECONDARY, S4]
- CNY card: v2.5 ¥0.02/¥1/¥2; v2.5-pro ¥0.025/¥3/¥6; TTS free limited; ASR ¥0.5/hr. [SECONDARY, S15]

### V2-Omni / V2-TTS / V2.5-Pro additional figures
- V2.5-Pro: MMLU 89.4%; MATH 86.2%; GSM8K 99.6%; GPQA Diamond 66.7%; AA Intelligence Index 54 (April 2026, methodology unstated). [SECONDARY, S41]
- V2.5-Pro: 384 experts, 8/token; 131K max output. [SECONDARY, S41]
- V2-Omni: MM-BrowserComp 52.0; GDPval-AA 1435 Elo. [SECONDARY, S38]
- V2-Omni cache reads $0.08/M; V2-Pro cache reads $0.20/M (≤256K), $0.40/M (256K–1M); cache writes temporarily free. [SECONDARY, S39]
- MiMo-VL-7B-RL: 59.4 OlympiadBench; 56.1 OSWorld-G; 35/40 tasks vs Qwen2.5-VL-7B. [SECONDARY, S36]

### Additional benchmark and pricing figures
- MiMo-V2-Pro: SWE-bench Verified 78%; ClawEval 61.5; PinchBench 81.0 (3rd). [SECONDARY, S26]
- MiMo-V2.5-Pro: LMSYS Elo ~1462.5 (top-3 open-weight); GDPVal-AA Elo 1581; ClawEval Pass^3 64% at ~70K tokens/trajectory. [SECONDARY, S32][SECONDARY, S33]
- V2-Pro cache reads $0.20/M (standard tier); launch promo: temporarily free cache writes. [SECONDARY, S27][SECONDARY, S28]
- V2.5-Pro self-host floor: ~600GB+ VRAM (8×H100-80GB for 256K; 8×H200 for 1M). [SECONDARY, S32]


### MiMo-7B-RL-0530 vendor benchmark deltas (2025-05-30)
- AIME 2024: 68.2 → 80.1 (+11.9 points). [VENDOR, S1]
- AIME 2025: 55.4 → 70.2 (+14.8 points). [VENDOR, S1]
- LiveCodeBench v5: 57.8 → 60.9 (+3.1 points). [VENDOR, S1]
- LiveCodeBench v6: 49.3 → 52.2 (+2.9 points). [VENDOR, S1]
- GPQA Diamond: 54.4 → 60.6 (+6.2 points). [VENDOR, S1]
- SFT dataset scale: ~500K → ~6M instances (12×). [VENDOR, S1]
- RL context window: 32K → 48K tokens (1.5×). [VENDOR, S1]
- RL problem count: ~130,000 math/code problems. [SECONDARY, S2]
- MiMo-VL-7B training data: ~2.4T tokens across 4 stages. [SECONDARY, S2]

### V2-Pro / V2-Omni launch pricing (2026-03-18 rate card)
- V2-Pro ≤256K context: $1/M input, $3/M output. [SECONDARY, S5]
- V2-Pro 256K–1M context: $2/M input, $6/M output. [SECONDARY, S5]
- V2-Omni: $0.40/M input, $2/M output (256K context). [SECONDARY, S6]
- V2-TTS speech pretraining: >100M hours (vendor claim). [VENDOR, S6]

### V2.5 series pricing evolution (2026-05-27 rate card)
- V2.5-Pro: $0.435/M input, $0.87/M output (down from $1/$3 at launch under 256K). [SECONDARY, S14][SECONDARY, S16]
- V2.5-Pro cache hit: $0.0036/M tokens. [SECONDARY, S14]
- V2.5: $0.14/M input, $0.28/M output. [SECONDARY, S14][SECONDARY, S15]
- V2.5 cache hit: $0.0028/M tokens. [SECONDARY, S14]
- MiMo ASR: ~$0.074/hour overseas, ¥0.5/hour domestic. [SECONDARY, S16]
- V2.5 endpoints: 1M context, 128K max output. [SECONDARY, S14]

### V2.5 / V2.5-Pro training and architecture figures
- V2.5-Pro: 1.02T total / 42B active parameters. [VENDOR, S13]
- V2.5-Pro: ~27T pretraining tokens, FP8, 6:1 global-to-sliding-window attention, 128-token local window. [VENDOR, S13]
- V2.5: 310B total / 15B active parameters, ~48T training tokens. [SECONDARY, S13][SECONDARY, S18]
- V2.5-Pro vs V2.5 active-parameter ratio: 42B vs 15B (2.8×). [SECONDARY, S13][SECONDARY, S18 — arithmetic on sourced figures]

