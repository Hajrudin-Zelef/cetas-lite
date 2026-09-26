---
id: ai-industry-kb-2026-wave6/10-mistral-ai/devstral-figures
title: "Devstral figures"
domain: mistral-ai
role: deep-dive
task: actor-profile
actors: ["Alibaba", "DeepSeek", "Mistral", "Moonshot", "Nvidia", "Samsung"]
dates: ["2025-03-30", "2025-06-10", "2025-06-30", "2025-07-10", "2025-08", "2025-08-12", "2025-12", "2025-12-02", "2025-12-09", "2026-03-16", "2026-04", "2026-04-28", "2026-04-29", "2026-05", "2026-05-22", "2026-06-23", "2026-07-07", "2026-07-10", "2026-07-24", "2026-07-31", "2026-09-08", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "apache", "arr", "benchmark", "blackwell", "compute", "datacenter", "deepseek", "gpus", "kimi", "mistral", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4951, 5029]
section: "§10. Mistral AI"
delta_of: ai-industry-kb-2026
sha256: 4f48009868316f2654223a2fbfc3ef5f89173587d18c8e08d5c87e88852e5c3a
---

# Devstral figures

### Devstral figures
- Small 1.1: 24B; Apache 2.0; 128K context; SWE-Bench Verified 53.6% vs Small 1.0 46.8%; #1 open model 2026-07-10. [SECONDARY, S21][SECONDARY, S24]
- Medium 2507: 61.6%+ SWE-Bench Verified (single source). [SECONDARY, S24]
- API devstral-small-2505: $0.10/$0.30 per 1M. [SECONDARY, S23]
- Runs on 1× RTX 4090 or Mac 32GB RAM. [SECONDARY, S22]

### Magistral figures
- Medium: AIME 2024 73.6%, 90% with majority voting @64 (vendor). [VENDOR, S26]
- Small: 70.7% / 83.3% (vendor). [VENDOR, S27]
- 8 reasoning languages; Flash Answers 10× throughput claim (vendor). [VENDOR, S26]
- Pricing: Medium $2/$8; Small $0.50/$1.50; 128K context. [SECONDARY, S19]
- Magistral Small 1.2 retired from API 2026-07-31. [SECONDARY, S31]

### OCR figures
- OCR 3: $2/1K pages, 50% batch discount; 74% win rate claim (vendor). [SECONDARY, S32]
- OCR 4: $4/1K pages, $2 with Batch API; 170 languages / 10 groups; OlmOCRBench 85.20; 72% blind-eval win rate (vendor). [SECONDARY, S35]
- Anaqua measured ~4× per-page throughput vs prior provider. [SECONDARY, S34]

### Corporate figures
- Series D: €3B at >€21B post-money (2026-09-08); Series C: €1.7B at €11.7B (Sep 2025); debt: $830M / €723M (Apr 2026). [SECONDARY, S13][SECONDARY, S16]
- Scale: ~1,000 employees; 20 countries; 125+ enterprise customers. [SECONDARY, S33][VENDOR, S13]
- Revenue: €200M (2025) → €1B target (2026); ~$400M ARR estimated Jan 2026 (community). [SECONDARY, S33][COMMUNITY, S13]
- Compute: 18,000 Blackwell GPUs, 44 MW datacenter near Paris (2026). [SECONDARY, S13]
- Emmi AI acquisition: ~€300M (May 2026). [SECONDARY, S13]
- Robostral Navigate: 8B params; 76.6% R2R-CE; 400K sim trajectories. [SECONDARY, S13]
- Vibe pricing: Free / Pro €14.99 / Team €24.99 (€19.99 annual) / Enterprise. [SECONDARY, S17][SECONDARY, S35]

## Main actors
- **Mistral AI** — French lab; ships a split licensing strategy: Apache 2.0 for small/local models, modified MIT for large coding models, proprietary for medium flagship tiers [SECONDARY].
- **Mistral Vibe CLI** — coding CLI shipped alongside Devstral 2 [SECONDARY].
- **Artificial Analysis** — no Mistral figures cited in this section; do not mix the AA no-cross-version rule here [DIRECTIONAL].

## Timeline and context
- **2024-10** — Ministral 3B/8B released — the established line the 2026 "Ministral 3" claim must NOT be merged with [SECONDARY].
- **2025-07-10** — Devstral 2507 released (Small 1.1 Apache 2.0 / Medium proprietary) [VENDOR].
- **2025-08** — Medium 3.1 refresh [SECONDARY].
- **2025-12** — Mistral Large 3 released [SECONDARY].
- **2025-12-09/10** — Devstral 2 + Devstral Small 2 + Mistral Vibe CLI released [SECONDARY].
- **2026-03-16** — Mistral Small 4 released (119B/6B, Apache 2.0, 256K) [SECONDARY].
- **2026-04-29/30** — Medium 3.5 released [SECONDARY].
- **2026-07-24** — Mistral signs the open-weights coalition letter [SECONDARY].
- The standing correction: any undated "Mistral Large 3 (2026)" or "Devstral 2 (2026)" citation in secondary press is the brief's error class — re-pin to December 2025 [DIRECTIONAL].


### New verified timeline entries — expansion

- 2025-03-30: Mixtral 8x22B API retirement (published date; legacy self-hosting only thereafter). [SECONDARY, S31]
- 2025-06-10: Magistral reasoning family announced at London Tech Week — Small (24B, Apache 2.0) and Medium (enterprise preview). [SECONDARY, S26][SECONDARY, S30]
- 2025-06: Mistral AI Campus announced with NVIDIA and France's state investment bank. [SECONDARY, S13]
- 2025-07-10: Devstral 2507 released — Small 1.1 (24B, Apache 2.0, 53.6% SWE-Bench Verified) and Medium (API, 61.6%+). [SECONDARY, S24][SECONDARY, S21]
- 2025-08-12/13: Mistral Medium 3.1 released (v25.08), knowledge cutoff 2025-06-30. [SECONDARY, S2][SECONDARY, S1]
- 2025-09: Magistral Small 1.2 (2509) released with multimodal reasoning and [THINK] tokens. [SECONDARY, S31]
- 2025-09: Series C — €1.7B at €11.7B post-money, led by ASML. [SECONDARY, S13][SECONDARY, S32]
- 2025-12-02: Mistral Large 3 (675B/41B, Apache 2.0) and Ministral 3 family (3B/8B/14B, Apache 2.0) released; Devstral 2 (2512) generation begins. [SECONDARY, S31][SECONDARY, S13][SECONDARY, S32]
- 2025-12: Mistral OCR 3 launched at $2/1K pages during the December product offensive. [SECONDARY, S32][SECONDARY, S33]
- 2026-01: Community intelligence brief estimates ~$400M ARR. [COMMUNITY, S13]
- 2026-04-28: Mistral Medium 3.5 released per official changelog (128B dense, Modified MIT) — corrects the existing §10 dating. [VENDOR, S4]
- 2026-04: $830M (€723M) debt financing. [SECONDARY, S13]
- 2026-05: Emmi AI acquired for ~€300M. [SECONDARY, S13]
- 2026-05-22: Mistral Medium 3.1 deprecated; Medium 3.5 named as replacement. [VENDOR, S2]
- 2026-05 (approx): Le Chat rebranded as Vibe with Work/Code/Chat modes. [SECONDARY, S17][SECONDARY, S35]
- 2026-06-23: Mistral OCR 4 launched ($4/1K pages, $2 batch; 170 languages; single-container air-gapped). [SECONDARY, S35][SECONDARY, S34]
- 2026-07-07 18:00 CET: OCR 4 production webinar. [SECONDARY, S33]
- 2026-07-31: Published API retirement for Magistral Small 1.2 and Mistral Small 3.2. [SECONDARY, S31]
- 2026-09-08: Series D announced — €3B at >€21B post-money, led by Samsung, co-led by EQT Scaleup Europe Fund and PSG Equity. [SECONDARY, S13][SECONDARY, S14][SECONDARY, S15][SECONDARY, S16]
- 2026-09-21: Provider listing updates Codestral-latest to $0.33/$0.99 (131K context). [SECONDARY, S18]
- 2026-09-22 (research date): Vibe freemium structure in place; Mistral Compute datacenter launching 2026; Modified MIT verbatim carve-out extracted ($20M/month revenue threshold). [SECONDARY, S17][SECONDARY, S13][SECONDARY, S36]

## Implications
1. The brief's Mistral dates were systematically off — Large 3 is December 2025, not April 2026; Devstral 2 is December 2025, not 2026; Medium 3.1 is an August 2025 refresh. Wave 6's corrections discipline applies: pin year with month whenever citing Mistral [DIRECTIONAL].
2. Mistral's licensing mirrors the industry two-track: Apache 2.0 for the small/local line (Devstral Small 1.1/2, Small 4) vs modified-MIT or proprietary for the larger tiers (Devstral 2, Devstral Medium) — the most capable weights carry the most strings [SECONDARY/DIRECTIONAL].
3. "Ministral 3" (2026) stays [UNVERIFIED] until corroborated — do not treat it as a release [DIRECTIONAL].
4. Deep-dive licensing details belong to §21; per-model one-liners here only (delta discipline).
5. The Devstral 2 / Vibe CLI bundle (Dec 2025) is the coding-model artifact to cite for Mistral's 2026-relevant developer story — not the unverified "Ministral 3" line [SECONDARY].
6. When comparing Mistral coding claims against DeepSeek/Qwen/Kimi coding claims, pin benchmark version and harness — vendor-reported Devstral numbers are [VENDOR], never to be mixed with standardized boards [DIRECTIONAL].


### New verified implications — expansion

