---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/corroborated-launch-details-and-new-benchmarks-2026-03-2026-
title: "Corroborated launch details and new benchmarks (2026-03 → 2026-09)"
domain: mimo-and-xiaomi
role: deep-dive
task: benchmark
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Baidu", "China", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-04-30", "2025-05", "2025-05-30", "2025-12-17", "2026-03-11", "2026-03-18", "2026-03-19", "2026-04", "2026-04-02", "2026-04-22", "2026-04-23", "2026-05-27", "2026-06-30", "2026-09-21", "2026-09-22"]
keywords: ["benchmark", "benchmarks", "agent", "agentic", "amd", "apache", "attention", "aws", "claude", "context window", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3792, 3944]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: f384ac5c7842005e09b69e6422266e69599282058995d732c85d0cf21ba4146b
---

# Corroborated launch details and new benchmarks (2026-03 → 2026-09)

### Corroborated launch details and new benchmarks (2026-03 → 2026-09)
- On 2026-03-11 an anonymous model called "Hunter Alpha" appeared on OpenRouter with no developer listed and quickly climbed to the top of daily usage charts; on 2026-03-18 Xiaomi revealed Hunter Alpha was MiMo-V2-Pro — a deliberate stealth launch to let performance speak without brand bias. [SECONDARY, S27][SECONDARY, S8]
- The MiMo team is led by Luo Fuli, an engineer who came from the DeepSeek R1 project. [SECONDARY, S27][SECONDARY, S32]
- Xiaomi's stock (1810.HK) jumped 5.8% on the V2-Pro launch day. [SECONDARY, S26][SECONDARY, S44]
- MiMo-V2-Pro is currently closed source; MiMo's lead (Fuli Luo) said Xiaomi intends to open-source a variant of V2-Pro with no specified timeline. [SECONDARY, S26][SECONDARY, S44][SECONDARY, S46]
- MiMo-V2-Pro scored 78% on SWE-bench Verified, against Claude Opus 4.6's 80.8% and Claude Sonnet 4.6's 79.6%. [SECONDARY, S26][SECONDARY, S44]
- On ClawEval (OpenClaw's agentic benchmark) V2-Pro hits 61.5, approaching Opus 4.6's 66.3. [SECONDARY, S26][SECONDARY, S38]
- On PinchBench V2-Pro sits third globally at ~81.0 (one report cites 84.0), behind Opus 4.6 and its sibling V2-Omni. [SECONDARY, S26][SECONDARY, S38]
- Artificial Analysis Intelligence Index placement for V2-Pro varies by reporting date: #8 globally (Decrypt), #9 globally and #3 in China (ChinaBizInsider), #10 with score 49 tied with GPT-5.2 Codex (Particula) — dated snapshots, not a contradiction. [SECONDARY, S26][SECONDARY, S8][SECONDARY, S27]
- V2-Pro GDPval-AA Elo: 1,426 (GLM-5: 1,406, Kimi K2.5: 1,283; Claude Sonnet 4.6: 1,633) — highest recorded Chinese-origin model in that category per the reporting secondary. [SECONDARY, S43][SECONDARY, S45]
- V2-Pro AA hallucination rate: 30% (vs 48% for V2-Flash); AA Omniscience index: +5 (GLM-5: +2, Kimi K2.5: -8); full AA Intelligence Index run: 77M output tokens (vs 109M GLM-5, 89M K2.5), costing $348 vs $2,304 for GPT-5.2 and $2,486 for Claude Opus 4.6. [SECONDARY, S43]
- V2-Omni multimodal benchmarks: BigBench Audio 94.0, MMAU-Pro 69.4, FutureOmni 66.7 (community guide figures). [COMMUNITY, S48 — single source]
- One-week launch free API access was extended to 2026-04-02 due to demand. [SECONDARY, S46]
- V2-Omni was codenamed "Healer Alpha" on OpenRouter before the official release. [SECONDARY, S46][SECONDARY, S47]
- MiMo-V2-Flash: open weights under the MIT license (weights + inference code on Hugging Face); Pro and Omni are proprietary; TTS is proprietary with no publicly available weights. [SECONDARY, S46]
- MiMo-V2-TTS capabilities: emotional transitions, mid-sentence tone shifts, singing with accurate pitch, and synthesis of regional dialects (Sichuan, Cantonese, Henan, Taiwanese); positioned as the voice completing the Pro-plans/Omni-executes agent loop for Xiaomi's phone/vehicle/home ecosystem. [SECONDARY, S46][COMMUNITY, S48]
- V2-Pro launch context: stealth "Hunter Alpha" ran over 1.5T tokens on OpenRouter before Xiaomi acknowledged it; one community transcript notes an independent AI Index figure of 49 with V2-Pro scoring one point below GLM-5 — vs the #8/#9/#10 dated snapshots preserved above. [SECONDARY, S50][COMMUNITY, S1]
- V2-Pro cache reads cost $0.20 per million tokens on the standard tier ($0.40/M at 256K–1M); cache writes were temporarily free at launch. [SECONDARY, S27][SECONDARY, S28][SECONDARY, S41]
- V2-Omni context is reported as 262K tokens in one secondary vs the 256K figure in launch coverage — a minor catalog discrepancy preserved as stated. [SECONDARY, S27][SECONDARY, S6]
- The official V2.5-Pro spec page confirms: FP8 (E4M3) mixed precision, Base variant at 256K vs full model at 1M context, learnable attention-sink bias preserving long-context performance. [VENDOR, S29][VENDOR, S30]
- The official 3-stage post-training paradigm: (1) SFT for instruction following; (2) domain-specialized training where separate teacher models are each RL-optimized for math, safety, agentic tool-use, etc.; (3) Multi-Teacher On-Policy Distillation (MOPD) merging specialist capabilities into one student. [VENDOR, S29][SECONDARY, S31]
- V2.5-Pro open weights are on Hugging Face (XiaomiMiMo/MiMo-V2.5-Pro, FP8) and ModelScope under the MIT license — corroborating the permissive-license claim. [SECONDARY, S32][VENDOR, S30]
- Free access tiers for V2.5-Pro: Xiaomi's own API, OpenRouter free (200 req/day), Kenari, UnoRouter, AIHubMix. [SECONDARY, S32 — single source]
- LMSYS Elo ~1462.5, reported as top-3 among open-weight models. [SECONDARY, S32 — single source]
- GDPVal-AA Elo: 1581, reported as surpassing Kimi K2.6 and GLM 5.1. [SECONDARY, S33 — single source]
- ClawEval Pass^3: 64% at roughly 70K tokens per trajectory — 40–60% fewer tokens than Claude Opus 4.6, Gemini 3.1 Pro, and GPT-5.4 for comparable capability. [SECONDARY, S33 — single source]
- Reported agent-index placement: #1 among open-source models, #5 globally including closed-source. [SECONDARY, S33 — single source]
- Day-0 hardware support: Alibaba Pingtouge (Zhenwu 810E), AWS Trainium2, AMD ROCm, Baidu Kunlun Core, Enflame, Muxi, Tianshu Zhixin; frameworks SGLang and vLLM. [SECONDARY, S33 — single source]
- Self-hosting requirement: native FP8 weights need ~600GB+ VRAM — 8×H100-80GB minimum (capped at ~256K context), 8×H200 for full 1M; SGLang first-class, vLLM supported. [SECONDARY, S32 — single source]
- An independent sglang-jax cookbook recipe validates V2.5-Pro serving on TPU v6e/v7x with tensor + expert parallelism + sharded attention. [SECONDARY, S34 — single source]


### V2-Omni, V2-TTS, V2-Pro details (corroborated 2026-09)
- V2-Pro: ~1 trillion total parameters with 42B active (MoE); supports extended thinking via `<think>` tags; 32K max output tokens. [SECONDARY, S40 — single source]
- V2.5-Pro: 384 experts with 8 active per token. [SECONDARY, S41 — single source]
- V2.5-Pro max output: 131K tokens. [SECONDARY, S41 — single source]
- V2.5-Pro benchmarks (April 2026 provider documentation): MMLU 5-shot 89.4%; MATH 4-shot 86.2%; GSM8K 8-shot 99.6%; GPQA Diamond 66.7%. [SECONDARY, S41 — single source]
- V2.5-Pro AA Intelligence Index: 54 (#1 of 85, tied with GPT-5.4) per April 2026 provider docs — methodology version unstated, so kept separate from version-pinned scores. [SECONDARY, S41 — single source]
- V2-Omni: native unified processing of text, images, video, and audio through dedicated modality-specific encoders feeding into a shared backbone. [SECONDARY, S38 — single source]
- V2-Omni: MM-BrowserComp 52.0 (reported as outperforming Gemini 3 Pro in some reports); GDPval-AA 1435 Elo (reported ahead of Gemini 3 Pro). [SECONDARY, S38 — single source]
- V2-Omni launch demo: autonomous end-to-end browser operation via OpenClaw — researching "Xiaomi 17" on Xiaohongshu, comparing offers on JD.com, escalating to human customer service for bargaining, then adding to cart and placing the order. [SECONDARY, S8 — single source]
- V2-TTS: precise multi-granular emotional control — emotion and tone transitions mid-sentence, pitch-accurate singing synthesis, native dialect synthesis (Sichuanese, Henan, Cantonese, Taiwanese accents), role-play styles; infers punctuation and emphasis markers from text without manual annotation. [SECONDARY, S37][SECONDARY, S8]
- MiMo-VL-7B-RL: outperforms Qwen2.5-VL-7B on 35 of 40 evaluated tasks; 59.4 on OlympiadBench (surpassing models up to 78B parameters); 56.1 on OSWorld-G for GUI grounding. [SECONDARY, S36 — single source]


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

