---
id: ai-industry-kb-2026-wave6/10-mistral-ai/mistral-large-3-corroborated-specification-details-2025-12
title: "Mistral Large 3 corroborated specification details (2025-12)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["Alibaba", "China", "DeepSeek", "Mistral", "Nvidia", "OpenRouter", "SGLang", "Samsung", "Together AI", "vLLM"]
dates: ["2025-06-30", "2025-07-10", "2025-09", "2025-12", "2025-12-09", "2026-03", "2026-03-16", "2026-04-28", "2026-04-29", "2026-05-22", "2026-09-08"]
keywords: ["mistral", "agentic", "apache", "benchmarks", "compute", "consumer", "deepseek", "fp8", "gpus", "latency", "license", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4870, 4950]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: b951b05263ebded7213e168db52c42d0148dd76ae1e9ca418328a4dd1fb37a02
---

# Mistral Large 3 corroborated specification details (2025-12)

### Mistral Large 3 corroborated specification details (2025-12)
- Trained on 3,000 NVIDIA H200 GPUs. [SECONDARY, S47 — single source]
- Mistral API model IDs: `mistral-large-latest` or `mistral-large-2512`. [SECONDARY, S47 — single source]
- Independent benchmarks: MMLU ~85.5%; MMLU-Pro 73.11%; MATH-500 93.60% (high for a non-reasoning-specialized model). [SECONDARY, S47 — single source]
- LMSYS Elo ~1429.8; described as best-in-class agentic with frontier general performance. [COMMUNITY, S49 — single source]
- Self-hosting: FP8 runs on one 8×H200 node at 256K context; NVFP4 on 8×H100/A100; supported by vLLM (native), SGLang, TRT-LLM. [COMMUNITY, S49 — single source]
- Free access: Mistral AI "Experiment" plan (~1B tokens/month, 1 req/sec, no credit card); OpenRouter free tier (200 req/day); Together AI free tier; NVIDIA NIM free prototyping; Ollama Cloud free tier. [COMMUNITY, S49 — single source]
- Developer criticism: Large 3 priced at $0.50/$1.50 vs DeepSeek at $0.14/$2.28 — cheaper on input, pricier on output in 100K+ scenarios; critiqued as a generalist judged on specialist benchmarks. [SECONDARY, S47][SECONDARY, S48]

### Devstral 2 and Ministral 3 corroborated details (2025-12)
- Devstral 2 API: free until December 2025, then $0.40/M input, $2.00/M output; Devstral Small (2): $0.10/M input, $0.30/M output. [SECONDARY, S44][SECONDARY, S46]
- Devstral 2 requires at least four H100 GPUs or equivalent for deployment. [SECONDARY, S44 — single source]
- Devstral Small (2): 24B parameters, Apache 2.0, deployable on consumer GPUs/laptops. [SECONDARY, S43][SECONDARY, S44]
- Mistral Vibe CLI launched alongside: file manipulation, code searching, version control, command execution, persistent history, file-structure and Git-status scanning; coverage noted it is limited to Mistral's ecosystem. [SECONDARY, S44][SECONDARY, S45]
- Ministral 3: nine dense models — 3B, 8B, 14B × Base, Instruct, Reasoning variants — all Apache 2.0. [SECONDARY, S48 — single source]
- The Modified MIT carve-out for Devstral 2 explicitly covers derivatives, fine-tuned versions, and redistributed variants regardless of who hosts them — large enterprises cannot use the model even internally without a commercial license. [SECONDARY, S43 — single source]


### Magistral 1.2 updates corroborated (September 2025)
- Magistral 1.2 (Small 1.2 / Medium 1.2; API IDs `magistral-small-2509` and `magistral-medium-2509`) added a vision encoder for image analysis; Small 1.2 after quantization fits on a single RTX 4090 or a MacBook with 32GB RAM. [SECONDARY, S52][SECONDARY, S53]
- Medium 1.2 vendor benchmarks: AIME24 91.82% (vs DeepSeek-R1 91.40%, vs 73.59% in v1.0); AIME25 83.48% (vs DeepSeek 79.4%, vs 64.95% in v1.0); HMMT25 76.66%; GPQA Diamond 76.26% (vs 70.83% in v1.0, vs DeepSeek 81%); LiveCodeBench v5 75.00% (vs 59.36%); LiveCodeBench v6 68.50% (vs 50.29%); HLE text-only 11.76% (vs 8.99%). [VENDOR, S52 — single source]
- Small 1.2 vendor benchmarks: AIME24 86.14% (vs Qwen3-32B 81.40%, vs 70.68% in 1.0); AIME25 77.34% (vs 72.90%); GPQA Diamond 70.07% (vs 68.40%); LiveCodeBench v5 70.02% (vs 55.84%); v6 61.60% (vs 47.36%). [VENDOR, S52 — single source]
- Multimodal: Medium scored MMMU 70.0% (vs Mistral Medium 3's 65.0%, Small's 66.0%), MathVista 70.1%, MMMU-Pro standard 57.9%, MMMU-Pro vision 52.1%. [COMMUNITY, S54 — single source]
- Multilingual chain-of-thought: English, French, Spanish, German, Italian, Arabic, Russian, Simplified Chinese. [SECONDARY, S51 — single source]

### Mistral Series D (announced 2026-09-08)
- €3B raised at >€21B post-money valuation — per Mistral, the largest equity fundraising round ever completed by a European technology company. [SECONDARY, S55][SECONDARY, S56]
- Led by Samsung Electronics; co-leads Scaleup Europe Fund (managed by EQT) and PSG Equity; Nvidia and ASML participated alongside Advent, BlackRock-managed funds, Bpifrance, a16z, General Catalyst, Lightspeed, Salesforce Ventures, DST Global, Index Ventures, and the Grand Duchy of Luxembourg. [SECONDARY, S55][SECONDARY, S57]
- Series C comparison: €1.7B at €11.7B post-money (September 2025), led by ASML — Series D is a ~1.8× step-up. [SECONDARY, S55][SECONDARY, S56]
- Capital use: frontier research, owned compute capacity, infrastructure, international expansion; ~€4B committed to French/European data centers; target 1GW online by 2030. [SECONDARY, S56][SECONDARY, S57]
- Business footprint: 20 countries; 125+ enterprise customers including Airbus, ASML, HSBC. [SECONDARY, S56][SECONDARY, S57]
- Preceded by $830M in debt financing in March 2026 for data-center buildouts near Paris. [SECONDARY, S57 — single source]

## Figures and metrics
| Model | Release | Params | Context | License |
|---|---|---|---|---|
| Mistral Large 3 | 2025-12 | — | — | — [SECONDARY] |
| Mistral Medium 3.1 | 2025-08 refresh | — | — | — [SECONDARY] |
| Mistral Small 4 | 2026-03-16 | 119B / 6B active | 256K | Apache 2.0 [SECONDARY] |
| Mistral Medium 3.5 | 2026-04-29/30 | — | — | — [SECONDARY] |
| Devstral Small 1.1 (2507) | 2025-07-10 | 24B | — | Apache 2.0 [VENDOR] |
| Devstral Medium (2507) | 2025-07-10 | — | — | Proprietary [VENDOR] |
| Devstral 2 | 2025-12-09/10 | 123B dense | 256K | Modified MIT [SECONDARY] |
| Devstral Small 2 | 2025-12-09/10 | 24B | 256K | Apache 2.0 [SECONDARY] |
| "Ministral 3" | 2026 | — | — | — [UNVERIFIED] |


### New verified metrics — expansion

### Magistral 1.2 / Series D additional figures
- Magistral Medium 1.2: AIME24 91.82%; AIME25 83.48%; GPQA Diamond 76.26%; LCB v5 75.00%; LCB v6 68.50%; HLE 11.76%. [VENDOR, S52]
- Magistral Small 1.2: AIME24 86.14%; AIME25 77.34%; GPQA 70.07%; LCB v5 70.02%; v6 61.60%. [VENDOR, S52]
- Series D: €3B at >€21B post; Series C €1.7B at €11.7B; 125+ enterprise customers; 20 countries; €4B datacenters; 1GW by 2030; $830M debt Mar 2026. [SECONDARY, S55][SECONDARY, S56][SECONDARY, S57]

### Large 3 / Devstral 2 additional figures
- Large 3: MMLU ~85.5%; MMLU-Pro 73.11%; MATH-500 93.60%; LMSYS Elo ~1429.8; 3,000 H200 training GPUs; $0.50/$1.50. [SECONDARY, S47][COMMUNITY, S49]
- Devstral 2: 123B; 256K; $0.40/$2.00 (free until Dec 2025); ≥4 H100s; Modified MIT $20M/month carve-out incl. derivatives. [SECONDARY, S44][SECONDARY, S43]
- Devstral Small 2: 24B; Apache 2.0; $0.10/$0.30. [SECONDARY, S43][SECONDARY, S44]
- Ministral 3: 3B/8B/14B × Base/Instruct/Reasoning; Apache 2.0. [SECONDARY, S48]

### Medium 3.5 additional figures
- Vendor: SWE-Bench Verified 77.6%; T3-Telecom 91.4%. [VENDOR, S37]
- API: $1.50/M input, $7.50/M output. [SECONDARY, S41]
- License carve-out: $20M global consolidated monthly revenue. [SECONDARY, S36][SECONDARY, S41]
- LLM Radar: quality index 39/100; speed 162 tok/s; blended price $3.00/M. [SECONDARY, S38]


### Medium 3.1 / 3.5 figures
- Medium 3.1: 131,072-token context; $0.40/M in, $2.00/M out, $0.04/M cache read; cutoff 2025-06-30; deprecated 2026-05-22. [SECONDARY, S1][SECONDARY, S3][VENDOR, S2]
- Medium 3.5: 128B dense; 256K context; Modified MIT ($20M/month carve-out); vendor SWE-Bench Verified 77.6%, T3-Telecom 91.4%. [SECONDARY, S5][SECONDARY, S31][VENDOR, S37]
- Medium 3.5 changelog date: 2026-04-28 (corrects existing §10's Apr 29/30). [VENDOR, S4]

### Large 3 / Small 4 figures
- Large 3: 675B total / 41B active; 256K context; Apache 2.0; $0.50/$1.50. [SECONDARY, S31][SECONDARY, S7][SECONDARY, S8][SECONDARY, S47]
- Small 4: 119B total / 6–6.5B active; 128 experts, 4 active/token; 256K advertised context. [SECONDARY, S9][SECONDARY, S31]
- Small 4 vendor claims: −40% completion latency, 3× throughput vs Small 3; GPQA Diamond 71.2; MMLU-Pro 78.0. [VENDOR, S9]
- Small 4 self-host guidance: ~4× H100, 2× H200, or 1× DGX B200. [SECONDARY, S12]

### Ministral 3 figures
- 3B / 8B / 14B; Apache 2.0; 131K–262K context; 14B reasoning variant 85% AIME 2025 (single source). [SECONDARY, S31][SECONDARY, S13]

