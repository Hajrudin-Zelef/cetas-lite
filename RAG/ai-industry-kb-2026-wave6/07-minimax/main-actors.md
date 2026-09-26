---
id: ai-industry-kb-2026-wave6/07-minimax/main-actors
title: "Main actors"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "EU", "Hugging Face", "MiniMax", "Stability AI", "United States"]
dates: ["2025-10-27", "2025-12-30", "2026-01-01", "2026-01-09", "2026-02-12", "2026-02-13", "2026-03-13", "2026-03-18", "2026-04-12", "2026-06-01", "2026-06-03", "2026-06-06", "2026-06-17", "2026-06-18", "2026-07-31", "2026-08-02", "2026-08-03", "2026-08-26", "2026-09-19"]
keywords: ["agent", "agentic", "apache", "attention", "attribution", "deepseek", "gguf", "int4", "ipo", "lawsuit", "license", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3401, 3478]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 0eaba7c414df4858a78ada647deb2aa2893c40ee0384fcb44d17e277152b5e44
---

# Main actors

## Main actors
- **MiniMax** — Chinese AI lab; runs the MIT → modified-MIT → non-commercial → geo-excluding Community License tightening sequence across M2→M2.5→M2.7→M3 [SECONDARY].
- **Skyler Miao** — MiniMax R&D lead whose diagram originated the 9.7×/15.6× MSA figures [VENDOR].
- **Hailuo** — MiniMax's video product line/app; H3 is its third-generation branding [SECONDARY].
- **Hugging Face** — hosts the H3 open weights (2026-08-03) [SECONDARY].

## Timeline and context
- **2025-10-27** — MiniMax M2 released under plain MIT [SECONDARY].
- **2026-02-12** — M2.5 released; modified-MIT with "MiniMax M2.5" UI attribution requirement [SECONDARY].
- **2026-03-18** — M2.7 announced (API-only initially); weights ~2026-04-12 under non-commercial terms; MiniMax retreats from "open source" to "open weights" wording [SECONDARY].
- **2026-06-01** — MiniMax M3 released: 428B/~23B MoE, 1M context, MiniMax Sparse Attention, MiniMax Community License [SECONDARY].
- **2026-07-31** — Hailuo 3 API + Hailuo app release [SECONDARY].
- **2026-08-03** — Hailuo 3 weights on Hugging Face [SECONDARY].
- **2026-04-12 (approx)** — M2.7 weights window reported by MarkTechPost; the part-4 chronology centers weights around this date [SECONDARY].
- **2026-06-01** — the agent-times contradiction piece documented the gap between M2.7's "open source" claim and its non-commercial license [SECONDARY].
- The RITS Shanghai-NYU piece framed M3 as a frontier-coding model with 1M context and sparse attention [SECONDARY].

### Corrections applied (wave-6 discipline)
- **M3 date:** the brief's "March 13, 2026" is corrected to **June 1, 2026** — the resolved fact across the corpus [SECONDARY].
- **M3 license:** the one Part-4 "Apache 2.0" line is discarded; the resolved license is the custom MiniMax Community License [SECONDARY].
- **M2 date:** the Oct 23 outlier is discarded in favor of Oct 27 [SECONDARY].
- **M2.5→M3 license arc:** MIT → modified-MIT (attribution) → non-commercial → geo-excluding Community License — cite the full arc, not any single point, when characterizing MiniMax's 2026 posture [SECONDARY/DIRECTIONAL].


### New verified timeline entries — expansion (continued)

- **2026-06-01** — M3 launch; MiniMax stock +5% to HK$907.5 then −12.38% close; weights promised within ~10 days [SECONDARY].
- **2026-06-06/07** — Kilo independent audit: M3 vs Opus 4.8 bug-hunt [SECONDARY].


### New verified timeline entries — expansion (continued)

- **Late 2021** — MiniMax founded in Shanghai (Shanghai Xiyu Jizhi Technology Co., Ltd.) [COMMUNITY].
- **2026-01** — HK listing (0100.HK) [COMMUNITY].
- **2026-02-13** — M2.5 released (MIT, open weights): 230B/10B, SWE-Bench Verified 80.2% [SECONDARY] (the-decoder.com).
- **Early 2026-04** — M2.7 released (self-evolving, OpenClaw scaffold optimization) [SECONDARY] (particula.tech).
- **2026-04-12** — M2.7 officially open-sourced (weights on HF) [SECONDARY] (marktechpost.com).
- **2026-06-01** — M3 released (1M ctx, MSA, native multimodal) [COMMUNITY] (vibe-investing guide).


### New verified timeline entries — expansion

- **2025-12-30** — Reuters: MiniMax targets up to $539M Hong Kong IPO (25.4M shares, HK$151–165; cornerstone Alibaba, ADIA) [SECONDARY].
- **2026-01-01** — Year-end Hong Kong IPO rush coverage (MiniMax among Chinese AI firms listing) [SECONDARY] (bworldonline.com).
- **2026-01-09** — Planned HK debut; ~$6.5B implied valuation [SECONDARY] (reuters.com).
- **2026-06-03** — MLQ.ai quotes MiniMax on ~100T interleaved text+image training tokens for M3 (disputed — see contradiction log) [SECONDARY].
- **2026-06-17** — MarkTechPost piece mixes 109B/3T research-model figures into M3 coverage [SECONDARY] (techtimes.com).
- **2026-06-18** — TechTimes clarifies: 109B/3T = MSA research testbed; production M3 token budget undisclosed as of that date [SECONDARY].
- **2026-07-31** — MiniMax H3 (Hailuo 3.0) announced; teased #MiniMaxH3 the prior day [SECONDARY].
- **2026-08-02** — Third-party H3 model card reviewed; Stability AI H3Max (paid speed-optimized fine-tune) noted [SECONDARY] (ai-tools-mcp card).
- **2026-08-03** — H3 weights on Hugging Face; ComfyUI support merged same day; GGUF/INT4/NVFP4 quants within 24h [SECONDARY] (runaihome.com).
- **2026-08-26** — Independent H3 pricing/positioning analysis ($7.80/min vs $20–22/min rivals) [SECONDARY] (justbeingresourceful.com).
- **2026-09-19** — aifoss.dev licensing survey: revenue-cap/MAU-cap/regional-exclusion clause taxonomy directly applicable to MiniMax terms [SECONDARY].

## Implications
1. MiniMax is the corpus's load-bearing case study for the 2026 license-tightening trend: capability goes up while license freedom goes down, within the same lab (MIT → attribution → non-commercial → geo-exclusion) [DIRECTIONAL].
2. M3's 1M context plus sparse-attention claims target agentic/coding workloads; verify MSA figures against vendor measurement conditions before citing — three different vendor numbers circulate [DIRECTIONAL].
3. The MiniMax Community License's US/EU/UK/South Korea local-deployment exclusion is the most geographically aggressive community license in the corpus — a procurement blocker for Western on-prem use, distinct from Llama's narrower EU-multimodal carve-out [SECONDARY/DIRECTIONAL].
4. H3 shows MiniMax extending open weights to the video line (third Hailuo generation), not just text models [SECONDARY].
5. Cite M3 pricing only with both the standard and promotional figures labeled; the master contradictions log explicitly flags the gap [SECONDARY].
6. The one outlier M2 date (Oct 23) is discarded in favor of Oct 27; cite the corrected date only [SECONDARY].


### New verified implications — expansion (continued)

- **SCMP's framing ("challenges ByteDance")** confirms the H3 competitive set is domestic as much as Western — Kling and Seedance are the price anchors, not just Sora/Veo [DIRECTIONAL].


### New verified implications — expansion (continued)

- **"Open weights" with a territorial license** is a new category in this knowledge base — the weights are downloadable but not legally runnable in the four biggest developer markets; enterprises must read the Applicable Territory clause, not just the "open" headline [DIRECTIONAL].
- **The 2K regeneration module being API-only** means the marketed capability is never actually in the open weights — the classic open-core split applied to a media model [DIRECTIONAL].
- **The lawsuit's May 26 discovery ruling** is the legal context for every MiniMax video license decision; the May timing (2 months before H3) explains the territorial design [DIRECTIONAL].
- **Video is the last closed-model stronghold**: "one of the last corners of generative AI where closed, expensive models still had a comfortable lead" — H3's price/partial-open playbook is the same one DeepSeek ran on text [DIRECTIONAL].


### New verified implications — expansion (continued)

