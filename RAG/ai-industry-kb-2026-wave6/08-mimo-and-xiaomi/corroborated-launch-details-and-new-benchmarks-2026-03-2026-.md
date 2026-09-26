---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/corroborated-launch-details-and-new-benchmarks-2026-03-2026-
title: "Corroborated launch details and new benchmarks (2026-03 → 2026-09)"
domain: mimo-and-xiaomi
role: deep-dive
task: benchmark
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Baidu", "China", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-03-11", "2026-03-18", "2026-04", "2026-04-02"]
keywords: ["benchmark", "benchmarks", "agent", "agentic", "amd", "attention", "aws", "claude", "cost", "deepseek", "distillation", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3792, 3836]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 032ba92930f03e7d66f2f5f36b6a1da43f9fdea0290b1dafabb1709ed9c2a8db
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


