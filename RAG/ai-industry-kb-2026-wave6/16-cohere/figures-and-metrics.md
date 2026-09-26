---
id: ai-industry-kb-2026-wave6/16-cohere/figures-and-metrics
title: "Figures and metrics"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Cohere", "CoreWeave", "DeepSeek", "Microsoft", "OpenAI"]
dates: ["2025-01-09", "2025-04-01", "2025-04-15", "2025-07", "2025-08-06", "2026-04-04", "2026-05", "2026-05-12", "2026-05-20"]
keywords: ["apache", "arr", "attention", "bedrock", "cohere", "context window", "copilot", "cost", "deepseek", "embedding", "license", "merger"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7972, 8064]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: f4700e9bd9272e1c7165124b0407a5b683eb87df8188f080d4b3ca091b383427
---

# Figures and metrics

## Figures and metrics
| Item | Detail | Provenance |
|---|---|---|
| Command A+ | launched 2026-05-20 | [VENDOR] |
| Command A+ params | 218B total / 25B active MoE | [VENDOR] |
| Command A+ license | Apache 2.0 (first for Cohere) | [VENDOR] |
| Command A+ languages | 48 | [VENDOR] |
| Command A+ context | 128K input / 64K generation | [VENDOR] |
| Command A+ deployment floor | 1× B200 or 2× H100 | [VENDOR] |
| Command A+ quantization | "near lossless" (not "lossless") | [VENDOR] |
| Predecessors (A / R / R+) | CC-BY-NC research-only | [SECONDARY] |
| Aya Expanse 8B / Aya Vision 8B API retirement | 2026-04-04 | [SECONDARY] |
| FedRAMP High | 2026-05-12 | [VENDOR] |
| Command R+ API | $3.00 / $15.00 per 1M | [SECONDARY] |

- No 2026 releases found: Command R7B successor, Embed 5, Rerank 5.
- Open item: Rerank 4 context — 32K [SECONDARY] vs 4K [COMMUNITY], unresolved.


### New verified metrics — expansion

| Claim | Value | Sources | Label |
|---|---|---|---|
| Command A+ total parameters | 218B | S1, S2 | [SECONDARY] |
| Command A+ active parameters per step | 25B (one source says 24B — see contradictions) | S1, S5 | [SECONDARY] |
| Command A+ experts / active | 128 experts, 8+1 active | S1 (single source) | [SECONDARY] |
| Command A+ input context | 128K tokens | S1, S3, S8 | [SECONDARY] |
| Command A+ max output | 64K tokens | S1, S3, S8 | [SECONDARY] |
| Command A+ license | Apache 2.0 | S1, S2 | [SECONDARY] |
| Command A+ release window | May 20–21, 2026 | S1, S5 | [SECONDARY] |
| Command A+ knowledge cutoff | April 1, 2025 | S1 (single source) | [SECONDARY] |
| Command A+ language coverage | 48 languages | S1 (single source) | [SECONDARY] |
| Command A+ W4A4 deployment | 2× H100 or 1× B200 | S1, S5 | [SECONDARY] |
| Command A+ hosted pricing (third-party mirror) | $2.50 / $10.00 per M input/output | S2 (single source) | [SECONDARY] |
| Command A parameters | 111B | S17 (single source) | [SECONDARY] |
| Command A context window | 256K tokens | S17, S17b | [SECONDARY] |
| Command A throughput vs R+ | ~150% higher (~32.8 tok/s out, TTFT ~706 ms) | S17 (single source) | [SECONDARY] |
| Command A hardware | 2× A100 or H100 minimum | S17 (single source) | [SECONDARY] |
| Command A / R+ price | $2.50 in / $10.00 out per M tokens | S17, S18 | [SECONDARY] |
| Command R+ launch price (04-2024) | $3.00 in / $15.00 out per M tokens | S18 (single source) | [SECONDARY] |
| Command R7B price | $0.0375 in / $0.15 out per M tokens | S18, S26 | [SECONDARY] |
| Embed v4 text price | $0.12 per M tokens | S22, S23 | [SECONDARY] |
| Embed v4 image price | $0.47 per M image tokens | S23 (single source) | [SECONDARY] |
| Embed v4 MTEB | 65.2 vs text-embedding-3-large 64.6 | S22 (single source) | [SECONDARY] |
| Embed v4 dimensions | 256 / 512 / 1024 / 1536 | S22 (single source) | [SECONDARY] |
| Embed v4 release | April 15, 2025 | S22 (single source) | [SECONDARY] |
| Cohere ARR (Feb 2026, disclosed) | ~$240M | S9, S11 | [SECONDARY] |
| Cohere ARR mid-2025 annualized | ~$100M | S11 (single source) | [SECONDARY] |
| Combined entity headline valuation | ~$20B | S9, S10 | [SECONDARY] |
| Cohere standalone valuation (Sept 2025) | $6.8–7B | S10, S11 | [SECONDARY] |
| Merger equity split | ~90% Cohere / ~10% Aleph Alpha | S10 (single source) | [SECONDARY] |
| Schwarz Digits structured financing | €500M / ~$600M (sources differ) | S10, S11 | [SECONDARY] |
| Series E size (reported) | $2.5B in progress / up to $3B in talks | S12, S11 | [SECONDARY] |
| Combined-entity valuation/ARR multiple | ~83x ($20B on $240M ARR) | S9 (single source) | [SECONDARY] |
| Command A+ W4A4 build self-hosting cost surface | $0 marginal token cost under Apache 2.0 | S2, S5 | [DIRECTIONAL] |
| Rerank 4 context (launch-era) | 32K tokens | S21 (single source) | [SECONDARY] |
| Rerank 4 context (2026 docs catalog) | 4K tokens | S8 (single source) | [COMMUNITY] |
| cohere-transcribe-03-2026 max audio | 25MB per input | S8 (single source) | [COMMUNITY] |
| Bedrock Command provisioned (no commit) | $49.50/hour | S24 (single source) | [SECONDARY] |
| Bedrock Command provisioned (6-month) | $23.77/hour | S24 (single source) | [SECONDARY] |
| Bedrock Rerank 3.5 on-demand | $2.00 per 1,000 queries | S24 (single source) | [SECONDARY] |
| Bedrock Embed 4 on-demand | $0.12 per M input tokens | S24 (single source) | [SECONDARY] |
| FedRAMP High authorization | May 2026 via Second Front Systems | S13b (single source) | [SECONDARY] |
| Legacy Command R+ 04-2024 pricing | $3.00 / $15.00 per M in/out | S18 (single source) | [SECONDARY] |
| Trial key free allowance | 1,000 calls/month, non-production | S18, S26 | [SECONDARY] |
| Command A+ vs GPT-5.5/Opus 4.7 size framing | "trillions of parameters" (third-party estimate, not verified) | S5 (single source) | [UNVERIFIED] |

| Claim | Value | Sources | Label |
|---|---|---|---|
| Command A release window | March 13–16, 2025 | S28, S29 | [SECONDARY] |
| Command A context / languages | 256K tokens / 23 languages | S28, S29 | [SECONDARY] |
| Command A attention design | 3 sliding-window layers (4096) + 1 global | S28 (single source) | [SECONDARY] |
| Command A claimed streaming speed (100K ctx) | 73 tok/s vs 38 (GPT-4o) / 32 (DeepSeek-V3) | S29 (single source) | [VENDOR] |
| Command A ADI2 dialect score | 24.7 vs 15.9 (GPT-4o) / 15.7 (DeepSeek-V3) | S29 (single source) | [VENDOR] |
| Command A HF config context | 128K default (256K supported, configurable) | S30 (single source) | [VENDOR] |
| Command A research license | CC-BY-NC + Acceptable Use Policy | S30 (single source) | [VENDOR] |
| Command A Vision | 112B, 6 languages, limited tool use | S32 (single source) | [COMMUNITY] |
| Command A+ HF model ID pattern | CohereLabs/command-a-plus-05-2026-{w4a4,bf16} | S32 (single source) | [COMMUNITY] |
| Command A+ API model name | command-a-plus-05-2026 | S32, S8 | [COMMUNITY] |
| Command A+ languages | 48 | S32, S1 | [SECONDARY] |
| North early access | January 9, 2025 | S33, S34 | [SECONDARY] |
| North GA | August 6, 2025 (~7 months later) | S35 (single source) | [SECONDARY] |
| North launch customers (GA) | Dell Technologies, Royal Bank of Canada | S35 (single source) | [SECONDARY] |
| North vendor eval vs Copilot/Vertex | wins all four categories (internal eval) | S33 (single source) | [VENDOR] |
| North document parsing formats | PDF, PPT, DOCX, XLSX | S34 (single source) | [SECONDARY] |
| CoreWeave training speedup claim | 3x faster | S36 (single source) | [VENDOR] |
| CoreWeave cluster | GB200 NVL72 (early production) | S36 (single source) | [VENDOR] |
| Cohere valuation July 2025 | $5.5B ($500M round) | S33 (single source) | [SECONDARY] |
| Cohere valuation Aug 2025 | $6.8B (Series D extension) | S31 (single source) | [SECONDARY] |
| Command A+ vs Command A (2025) params | 218B/25B vs 111B dense | S32 (single source) | [COMMUNITY] |
| Manufacturing-showcase claimed downtime cut | −50–65% (community demo) | S38 (single source) | [COMMUNITY] |
| Manufacturing-showcase defect escape | 35% → <8% (community demo) | S38 (single source) | [COMMUNITY] |

