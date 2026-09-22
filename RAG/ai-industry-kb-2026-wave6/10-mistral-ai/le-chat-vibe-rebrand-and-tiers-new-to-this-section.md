---
id: ai-industry-kb-2026-wave6/10-mistral-ai/le-chat-vibe-rebrand-and-tiers-new-to-this-section
title: "Le Chat → Vibe rebrand and tiers (new to this section)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Apple", "Google", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "Samsung"]
dates: ["2025-06", "2025-06-10", "2025-06-30", "2025-07-10", "2025-08-12", "2025-09", "2025-12", "2025-12-02", "2026-01", "2026-03-16", "2026-04", "2026-04-28", "2026-05", "2026-05-22", "2026-06", "2026-06-23", "2026-09-08", "2026-09-21"]
keywords: ["agent", "agents", "apache", "arr", "bedrock", "benchmarks", "blackwell", "chatgpt", "compute", "datacenter", "distribution", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4796, 4869]
section: "§10. Mistral AI"
sha256: e9421bec672ee0f10511514ce1fb1e49018a46f9b076990e971ef846cfb3f96f
---

# Le Chat → Vibe rebrand and tiers (new to this section)

### Le Chat → Vibe rebrand and tiers (new to this section)
- Mistral rebranded Le Chat as Vibe in approximately May 2026, unifying chat, work automation, and coding under one license. [SECONDARY, S17][SECONDARY, S35]
- Vibe modes: Work (long-running multi-step tasks across Google Workspace, Outlook, SharePoint, Slack, GitHub after plan sign-off), Code (remote coding agents in isolated sandboxes, reviewable PRs, parallel sessions, VS Code extension), and Chat. [SECONDARY, S35 — single source]
- Vibe availability: chat.mistral.ai web, iOS/Android mobile apps; the code interface is separate at code.mistral.ai. [SECONDARY, S17][SECONDARY, S35]
- Vibe tiers: Free; Pro €14.99/month; Team €24.99/user/month (or €19.99 with annual billing); Enterprise by quote. [SECONDARY, S17][SECONDARY, S35]
- Students receive Pro at half price. [SECONDARY, S17 — single source]
- Team tier adds storage, domain verification, exports, and administration rather than higher message quotas. [SECONDARY, S35 — single source]
- Vibe's single-license bundle of work automation plus coding agent is positioned against OpenAI's ChatGPT Work bundle, with Mistral's self-host and sovereignty posture as the differentiator. [SECONDARY, S35 — single source]

### Mistral corporate: funding, compute, enterprise (new to this section)
- Series D announced 2026-09-08: €3B raised at post-money valuation above €21B. [SECONDARY, S13][SECONDARY, S14][SECONDARY, S15][SECONDARY, S16]
- Series D led by Samsung Electronics; co-led by the Scaleup Europe Fund (EQT) and PSG Equity. [SECONDARY, S13][SECONDARY, S16]
- Series D participants included NVIDIA and ASML as existing investors, plus Advent, BlackRock-managed funds, Bpifrance, a16z, and Luxembourg's sovereign fund. [SECONDARY, S13][SECONDARY, S16]
- Mistral described Series D as the largest equity round ever raised by a European technology company. [SECONDARY, S15][SECONDARY, S13]
- Series C (September 2025): €1.7B raised at €11.7B post-money, led by ASML. [SECONDARY, S13][SECONDARY, S32]
- Debt financing (inaugural, March/April 2026): $830M from seven banks (BNP Paribas, Crédit Agricole CIB, HSBC, MUFG) for 13,800 NVIDIA GB300 GPUs at Bruyeres-le-Chatel near Paris, 44 MW, operational Q2 2026. [SECONDARY, S13][SECONDARY, S47][SECONDARY, S48]
- Stated use of Series D capital: frontier research, training compute, infrastructure, commercialization, and international expansion. [SECONDARY, S13][SECONDARY, S14]
- Mistral reported operations in 20 countries and more than 125 enterprise customers (named: Airbus, ASML, HSBC); enterprise wins include Airbus 5-year deal and BMW partnership. [VENDOR, S13][COMMUNITY, S45]
- Headcount approximately 1,000 employees. [SECONDARY, S33 — single source]
- Revenue: €200M in 2025; targeting €1B in 2026 (Le Monde); one internal research profile cites ~$400M ARR (Jan 2026), ~$6.4B total raised, $23B valuation on the June 2026 round. [SECONDARY, S33][COMMUNITY, S45]
- Mistral Compute: 44 MW datacenter near Paris; the debt round funds 13,800 GB300s, with a broader 18,000-GPU Blackwell plan cited; Mistral Compute announced at VivaTech 2025 with Macron and Jensen Huang. [SECONDARY, S13][SECONDARY, S47][COMMUNITY, S45]
- Mistral AI Campus announced June 2025 with NVIDIA and France's state investment bank. [SECONDARY, S13 — single source]
- Emmi AI (Austrian physics-simulation startup, large engineering models for fluid dynamics/structural impacts) acquired May 2026 for approximately €300M; Koyeb (AI infra) acquired 2026-02 as first M&A. [SECONDARY, S13][COMMUNITY, S45][SECONDARY, S49]
- Robostral Navigate: 8B-parameter embodied navigation model, Apache 2.0, 76.6% on R2R-CE, trained on 400K simulated trajectories, single RGB camera. [SECONDARY, S13][COMMUNITY, S45]
- Agents API: workflow orchestration with MCP support and a Document Library for RAG. [SECONDARY, S13 — single source]
- Forge: custom model-training platform for enterprise customers. [SECONDARY, S13 — single source]
- A community intelligence brief estimated ~$400M ARR in January 2026. [COMMUNITY, S13 — single source]


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


