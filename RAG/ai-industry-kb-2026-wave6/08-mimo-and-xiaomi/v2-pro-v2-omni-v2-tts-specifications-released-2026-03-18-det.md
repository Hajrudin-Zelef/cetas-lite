---
id: ai-industry-kb-2026-wave6/08-mimo-and-xiaomi/v2-pro-v2-omni-v2-tts-specifications-released-2026-03-18-det
title: "V2-Pro / V2-Omni / V2-TTS specifications (released 2026-03-18; details beyond base section)"
domain: mimo-and-xiaomi
role: deep-dive
task: actor-profile
actors: ["China", "Xiaomi"]
dates: ["2026-03-18", "2026-04-22", "2026-05-27", "2026-06-30"]
keywords: ["omni", "agent", "agentic", "attention", "context window", "fp8", "multimodal", "parameters", "pretraining", "pricing", "throughput", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3625, 3672]
section: "§8. MiMo and Xiaomi"
delta_of: ai-industry-kb-2026
sha256: 34219541557dc699ab216d4c6dc9b39934861542556657fc39b6bc5581dcc6f3
---

# V2-Pro / V2-Omni / V2-TTS specifications (released 2026-03-18; details beyond base section)

### V2-Pro / V2-Omni / V2-TTS specifications (released 2026-03-18; details beyond base section)
- V2-Pro is a mixture-of-experts model with approximately 1T total parameters and approximately 42B active parameters per token. [SECONDARY, S5][SECONDARY, S7]
- V2-Pro uses a hybrid/mixed attention architecture rather than uniform attention; one secondary reports a 7:1 hybrid ratio for V2-Pro specifically (see tension note). [SECONDARY, S5][SECONDARY, S26]
- V2-Pro supports a 1,048,576-token context window with a maximum output length of approximately 32K tokens. [SECONDARY, S7][SECONDARY, S5]
- V2-Pro's launch API pricing was $1 per million input tokens and $3 per million output tokens for prompts up to 256K tokens. [SECONDARY, S5][SECONDARY, S26][SECONDARY, S27][SECONDARY, S28]
- V2-Pro's launch pricing rose to $2 input / $6 output per million tokens for prompts from 256K to 1M tokens. [SECONDARY, S5][SECONDARY, S26][SECONDARY, S27]
- V2-Omni accepts text, image, video, and audio input modalities in a single model. [SECONDARY, S6][SECONDARY, S7]
- V2-Omni's context window is 256K tokens (one secondary reports 262K — preserved as a catalog discrepancy). [SECONDARY, S6][SECONDARY, S37][SECONDARY, S38][SECONDARY, S40]
- V2-Omni's reported API pricing is $0.40 per million input tokens and $2 per million output tokens, with $0.08/M cache reads. [SECONDARY, S6][SECONDARY, S37][SECONDARY, S8][SECONDARY, S39]
- V2-Omni supports long-audio understanding beyond 10 hours of continuous audio (claimed industry-leading duration). [SECONDARY, S6][SECONDARY, S8][SECONDARY, S38][SECONDARY, S40]
- V2-TTS performs joint speech-text modeling in a single architecture rather than cascading separate models. [SECONDARY, S6][SECONDARY, S7]
- V2-TTS uses a proprietary (non-open) audio tokenizer. [SECONDARY, S6 — single source]
- V2-TTS uses a multi-codebook acoustic model design. [SECONDARY, S6 — single source]
- V2-TTS speech pretraining consumed over 100 million hours of speech data, per vendor claims reported at launch (one secondary reports "hundreds of millions of hours"). [VENDOR, S6][SECONDARY, S8][SECONDARY, S38]
- The "Hunter" stealth codename was publicly tied to the V2 trio after community detection of anonymous "Hunter Alpha" checkpoints. [SECONDARY, S8][SECONDARY, S7]
- Hunter Alpha processed more than 1T tokens of community testing traffic in earlier reports, later revised above 1.5T tokens as the stealth period extended. [SECONDARY, S8 — range across report dates]
- V2-Pro is integrated into MiMo Studio, Xiaomi's model development and deployment workbench. [SECONDARY, S5][SECONDARY, S7]
- V2-Pro is integrated into the Xiaomi Browser AI assistant surface (native integrations targeted at the Chinese market; core API globally accessible). [SECONDARY, S5][SECONDARY, S37]
- V2-Pro is integrated into WPS / Kingsoft Office AI features (Word, Excel, PPT, PDF generation). [SECONDARY, S7][SECONDARY, S8][SECONDARY, S37]
- V2-Pro/V2.5-Pro are supported in the OpenClaw agent harness; a limited-time one-week free testing period ran through OpenClaw, OpenCode, KiloCode, Blackbox, and Cline at launch. [SECONDARY, S5][SECONDARY, S37][SECONDARY, S39]
- V2-Pro/V2.5-Pro are supported in the OpenCode agent harness. [SECONDARY, S5][SECONDARY, S37][SECONDARY, S39]
- V2-Pro/V2.5-Pro are supported in the Cline coding agent. [SECONDARY, S5][SECONDARY, S37][SECONDARY, S39]

### V2.5 and V2.5-Pro specifications (released 2026-04-22; details beyond base section)
- V2.5-Pro is a sparse mixture-of-experts model with 1.02T total parameters and 42B active parameters. [VENDOR, S29][VENDOR, S30][SECONDARY, S31][SECONDARY, S19]
- The 1.02T figure distinguishes V2.5-Pro from the approximately 1T figure reported for V2-Pro at launch; the two are distinct checkpoints. [VENDOR, S13][SECONDARY, S5]
- V2.5-Pro alternates sliding-window attention with global attention at a reported 6:1 ratio. [VENDOR, S29][VENDOR, S30][SECONDARY, S31]
- V2.5-Pro's local sliding window is 128 tokens. [VENDOR, S29][VENDOR, S30]
- V2.5-Pro was pretrained on approximately 27T tokens. [VENDOR, S29][VENDOR, S30][SECONDARY, S31][SECONDARY, S19]
- V2.5-Pro training used FP8 precision (E4M3 mixed). [VENDOR, S29][VENDOR, S30][SECONDARY, S31]
- V2.5-Pro was natively pretrained at 32K sequence length before context extension to 1M. [VENDOR, S29][SECONDARY, S31]
- V2.5-Pro includes a lightweight 3-layer MTP (multi-token prediction) module with dense FFNs, roughly tripling output throughput and accelerating RL rollouts. [VENDOR, S29][VENDOR, S30][SECONDARY, S31]
- V2.5 (non-Pro) is a sparse multimodal mixture-of-experts model with 310B total parameters and 15B active parameters. [SECONDARY, S18][SECONDARY, S19]
- V2.5 was trained on 48T tokens — substantially more than V2.5-Pro's reported 27T (tension preserved, not resolved). [SECONDARY, S13][SECONDARY, S33][SECONDARY, S19]
- V2.5 uses self-developed pretrained vision and audio encoders rather than a single fused encoder. [SECONDARY, S13][SECONDARY, S33]
- V2.5 aligns modalities through lightweight projection modules for cross-modal fusion on top of the separate encoders. [SECONDARY, S13][SECONDARY, S33]
- Post-2026-05-27 provider pricing for V2.5-Pro is $0.435 per million input tokens. [SECONDARY, S14][SECONDARY, S16]
- Post-2026-05-27 provider pricing for V2.5-Pro is $0.87 per million output tokens. [SECONDARY, S14][SECONDARY, S16]
- V2.6-Pro pricing (unchanged from V2.5 series): $0.0036 cache-hit input / $0.435 cache-miss input / $0.87 output per million tokens; Flash $0.0028/$0.14/$0.28; Pro-UltraSpeed 10× (20× output speed at same quality). [SECONDARY, S14][SECONDARY, S51][SECONDARY, S54]
- Post-2026-05-27 provider pricing for V2.5 is $0.14 input / $0.28 output per million tokens. [SECONDARY, S14][SECONDARY, S15]
- V2.5 cache-hit pricing is $0.0028 per million tokens in the post-May rate card. [SECONDARY, S14 — single source]
- MiMo ASR pricing is approximately $0.074 per hour for overseas usage or ¥0.5 per hour domestically. [SECONDARY, S16 — single source]
- V2.5-TTS API variants remained limited-time free at the time of the provider log update. [SECONDARY, S16 — single source]
- Legacy V2 model identifiers were deprecated and auto-routed to V2.5 endpoints. [SECONDARY, S16 — single source]
- Legacy V2 identifiers were fully retired on 2026-06-30 Beijing time. [SECONDARY, S16 — single source]
- V2.5 API endpoints expose 1M-token context with a 128K-token maximum output. [SECONDARY, S14 — single source]
- Third-party benchmarking coverage describes V2.5-Pro as among the most efficient and affordable models at "agentic claw" coding-agent tasks. [SECONDARY, S19 — single source]

