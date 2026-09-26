---
id: ai-industry-kb-2026-wave6/10-mistral-ai/consolidated-specification-sheets-full-field-reference
title: "Consolidated specification sheets (full-field reference)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Hugging Face", "Mistral", "Nvidia", "OpenAI"]
dates: ["2025-06-10", "2025-06-30", "2025-07-10", "2025-08-12", "2025-12", "2025-12-02", "2026-03-16", "2026-04-28", "2026-05-22", "2026-06-23", "2026-09-21"]
keywords: ["apache", "bedrock", "benchmarks", "compute", "distribution", "foundry", "gpt-5.6", "latency", "license", "mistral", "moe", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4825, 4869]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: d0eb73482ec3dbda98fe2118f2511b798e3091117b808fd8e4f39f46f03f8bdc
---

# Consolidated specification sheets (full-field reference)

### Consolidated specification sheets (full-field reference)
Fields marked "(base §10 anchor)" repeat the base section only as dated anchors inside a complete spec sheet; every other field is new.
- **Mistral Medium 3.1** — Released 2025-08-12/13 (v25.08); dense enterprise model; 131,072-token context (128K per Mistral docs); text/image/file input. [SECONDARY, S1][VENDOR, S2]
- Medium 3.1 knowledge cutoff: 2025-06-30; output limit 104.9K tokens. [SECONDARY, S1][SECONDARY, S3]
- Medium 3.1 pricing: $0.40/M in, $2.00/M out, $0.04/M cache read; license proprietary (API). [SECONDARY, S1][SECONDARY, S3]
- Medium 3.1 status: deprecated; published deprecation date 2026-05-22; replacement Medium 3.5. [VENDOR, S2]
- **Mistral Medium 3.5** — Released 2026-04-28 (official changelog; base §10 anchor Apr 29/30 corrected). [VENDOR, S4]
- Medium 3.5: 128B dense; 256K context; unified reasoning/coding/vision; configurable `reasoning_effort`. [SECONDARY, S5][SECONDARY, S31]
- Medium 3.5 license: Modified MIT; verbatim carve-out: no rights if global consolidated monthly revenue exceeds $20M — commercial license required above that bar. [SECONDARY, S36][SECONDARY, S41][SECONDARY, S5][SECONDARY, S31]
- Medium 3.5 parameters/layers: 128B disclosed; layer count not publicly disclosed. [DIRECTIONAL]
- **Mistral Large 3** — Released 2025-12-02 (base §10 anchor). [SECONDARY, S31]
- Large 3: 675B total / 41B active sparse MoE; 256K context; reasoning + vision. [SECONDARY, S31][SECONDARY, S7][SECONDARY, S48]
- Large 3 license: Apache 2.0; pricing $0.50/$1.50 per 1M; API + Bedrock + Azure AI Foundry. [SECONDARY, S31][SECONDARY, S8][SECONDARY, S7]
- Large 3 training tokens/compute: not publicly disclosed. [DIRECTIONAL]
- **Mistral Small 4** — Released 2026-03-16 (base §10 anchor). [SECONDARY, S9]
- Small 4: 119B total / 6–6.5B active MoE; 128 experts, 4 active/token; 256K advertised context. [SECONDARY, S9][SECONDARY, S31]
- Small 4 license: Apache 2.0; vendor claims −40% latency, 3× throughput vs Small 3. [SECONDARY, S9][VENDOR, S9]
- Small 4 vendor benchmarks: GPQA Diamond 71.2; MMLU-Pro 78.0. [VENDOR, S9]
- Small 4 distribution: Mistral API, AI Studio, Hugging Face, NVIDIA NIM. [SECONDARY, S10]
- **Ministral 3 family** — Released 2025-12-02 (correction to base §10). [SECONDARY, S31][SECONDARY, S13]
- Ministral 3: 3B / 8B / 14B dense edge models with vision; Apache 2.0; 131K–262K context. [SECONDARY, S31][SECONDARY, S13]
- Ministral 3 14B reasoning variant: 85% AIME 2025 (single source). [SECONDARY, S13]
- **Devstral Small 2507** — Released 2025-07-10 (base §10 anchor for Devstral 2507). [SECONDARY, S24]
- Devstral Small 1.1: 24B; Apache 2.0; 128K context; text-only (no vision encoder); All Hands AI collaboration. [SECONDARY, S21][SECONDARY, S23][SECONDARY, S22]
- Devstral Small 1.1: SWE-Bench Verified 53.6% (OpenHands); runs on 1× RTX 4090 / Mac 32GB. [SECONDARY, S21][SECONDARY, S22]
- Devstral Small 1.1 API: `devstral-small-2505` at $0.10/$0.30 per 1M. [SECONDARY, S23]
- **Devstral Medium 2507** — Released 2025-07-10; API-based enterprise model; 61.6%+ SWE-Bench Verified (single source). [SECONDARY, S24]
- **Magistral Small** — Released 2025-06-10; 24B; Apache 2.0; Hugging Face; multilingual chain-of-thought. [SECONDARY, S26][SECONDARY, S30]
- Magistral Small vendor benchmarks: 70.7% / 83.3% AIME 2024 (with majority voting). [VENDOR, S27]
- **Magistral Medium** — Released 2025-06-10; enterprise; preview in Le Chat/La Plateforme/SageMaker. [SECONDARY, S26][SECONDARY, S30]
- Magistral Medium vendor benchmarks: 73.6% AIME 2024, 90% with majority voting @64. [VENDOR, S26]
- Magistral Medium pricing: $2/M in, $8/M out; 128K context. [SECONDARY, S19]
- **Mistral OCR 3** — Released December 2025; $2/1K pages (50% batch discount); 74% win-rate claim (vendor). [SECONDARY, S32]
- **Mistral OCR 4** — Released 2026-06-23; $4/1K pages ($2 batch); 170 languages; bounding boxes + typed blocks + per-page/per-word confidence; single-container air-gapped. [SECONDARY, S35][SECONDARY, S34]
- **Codestral** — Catalog: $0.30/$0.90 per 1M, 256K context; alternate listing $0.33/$0.99, 131K context (2026-09-21). [SECONDARY, S19][SECONDARY, S18]


### Medium 3.5 license, pricing, and hosting — corroborated (2026-04/05)
- License portfolio context: Apache 2.0 covers Small 3/3.1/3.2/4, Devstral (incl. Small 1.1/2), Magistral Small, Mistral Large 3, Ministral 3, Voxtral Small/Mini/Realtime, Leanstral; Modified MIT covers Medium 3.5 and Devstral 2; research/non-commercial covers Codestral 22B (MNPL), Mistral Large 2 (MRL), Ministral 8B (MRL), Pixtral Large (MRL), Voxtral TTS (CC BY-NC 4.0). [SECONDARY, S39 — single source]
- API pricing: $1.50 per million input tokens, $7.50 per million output tokens — roughly 40% cheaper on input and 50% on output than OpenAI's GPT-5.6 Terra ($2.50/$15). [SECONDARY, S41 — single source]
- Hosting: NVIDIA build.nvidia.com endpoints, NVIDIA NIM containerized microservice, Le Chat Work Mode globally, and la Plateforme API. [SECONDARY, S37 — single source]
- Model card details: Mistral3 multimodal transformer, 24+ languages (EN, FR, DE, ES, IT, ZH, AR), reasoning/coding/function-calling/JSON/vision; training-data composition not published. [SECONDARY, S38 — single source]
- Critical reception: ML professor Pedro Domingos (via Decrypt) mocked the posture — "Only Mistral brags about how much worse its [model] is" — and questioned whether Mistral's strategy serves European AI; developer reaction to pricing plus license carve-outs stayed combative. [SECONDARY, S37 — single source]


