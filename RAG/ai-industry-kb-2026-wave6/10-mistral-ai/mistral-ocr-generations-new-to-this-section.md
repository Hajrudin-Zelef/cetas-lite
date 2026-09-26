---
id: ai-industry-kb-2026-wave6/10-mistral-ai/mistral-ocr-generations-new-to-this-section
title: "Mistral OCR generations (new to this section)"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["AWS", "Apple", "Google", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Samsung"]
dates: ["2025-03-30", "2025-06", "2025-08", "2025-09", "2025-12", "2026-01", "2026-04", "2026-05", "2026-06", "2026-06-23", "2026-07-07", "2026-07-31", "2026-09-08", "2026-09-21"]
keywords: ["mistral", "agent", "agents", "apache", "arr", "benchmark", "blackwell", "chatgpt", "compute", "datacenter", "distribution", "foundry"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4772, 4824]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: a1cf2cf387957c460ded85725d86cc479433fd861aa31108148b41ef34856006
---

# Mistral OCR generations (new to this section)

### Mistral OCR generations (new to this section)
- Mistral OCR 3 launched in December 2025, unveiled as a Tuesday release during Mistral's December product offensive. [SECONDARY, S32][SECONDARY, S33]
- OCR 3 pricing: $2 per 1,000 pages, with a 50% discount for batch processing. [SECONDARY, S32][SECONDARY, S35]
- OCR 3 vendor claim: 74% win rate against competing products on forms, scanned documents, complex tables, and handwritten content. [VENDOR, S32 — single source]
- Mistral OCR 4 launched 2026-06-23. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 pricing: $4 per 1,000 pages, dropping to $2 per 1,000 pages with the Batch API. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 returns bounding boxes, typed-block classification (titles, tables, equations, signatures), and per-page and per-word confidence scores alongside extracted text. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 covers 170 languages across 10 language groups. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 ships as a single container for fully air-gapped on-premise deployment. [SECONDARY, S35][SECONDARY, S34]
- OCR 4 vendor claims: 72% average win rate in blind human evaluations against leading systems; top score on OlmOCRBench (85.20). [VENDOR, S35 — single source]
- Mistral itself notes the OCR benchmark figures are directional rather than definitive. [VENDOR, S35 — single source]
- OCR 4 distribution: Mistral API and Studio, Amazon SageMaker, Microsoft Foundry. [SECONDARY, S34 — single source]
- Mistral OCR 4.1 exists as a proprietary document-parsing model with bounding boxes, block labels, block-level confidence, 170 languages, at $4/1K pages. [SECONDARY, S31 — single source]
- Mistral held an OCR 4 production webinar on 2026-07-07 at 18:00 CET. [SECONDARY, S33 — single source]
- Anaqua, a legal software firm, measured roughly 4× faster throughput per page with OCR 4 versus its prior provider. [SECONDARY, S34 — single source]

### Codestral and model catalog pricing (new to this section)
- Codestral catalog pricing: $0.30 per million input tokens, $0.90 per million output tokens, 256K context. [SECONDARY, S19 — single source]
- A second provider listing shows Codestral-latest at $0.33/$0.99 with 131K context, updated 2026-09-21, added August 2025, 80+ languages supported. [SECONDARY, S18 — single source]
- Devstral Medium catalog pricing: approximately $0.50/$1.50. [SECONDARY, S19 — single source]
- Shieldstral 1.0 exists in Mistral's model catalog as a content-moderation model entry. [VENDOR, S2 — single source]
- Retired legacy entries: Magistral Small 1.2 and Mistral Small 3.2 (24B, v3.2, repo 2506) both carry published API retirement 2026-07-31. [SECONDARY, S31 — single source]
- Mixtral 8x22B (141B total / 39B active) carries published API retirement 2025-03-30. [SECONDARY, S31 — single source]

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


