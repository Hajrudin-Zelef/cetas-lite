---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/bytedance-seed-team-volcano-engine
title: "ByteDance (Seed team / Volcano Engine)"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "Google", "Mistral", "Nvidia", "OpenRouter", "Unsloth"]
dates: ["2026-02", "2026-02-16", "2026-03-30", "2026-03-31", "2026-04", "2026-04-02", "2026-04-16", "2026-04-21", "2026-06-23", "2026-07", "2026-07-15", "2026-07-19", "2026-07-23", "2026-07-27", "2026-08-03", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-17", "2026-08-26", "2026-09", "2026-09-03", "2026-09-04"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "claude", "cost", "gemini", "gguf", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1575, 1619]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: df244d743e7f4424df09b4da256753bc78fec0ac42c878d9c7332d5303da4d3a
---

# ByteDance (Seed team / Volcano Engine)

- InclusionAI is Ant Group's AI lab; the Ling 3.0 Flash family is its 2026 efficient-agent flagship. Ling 3.0 Flash (2026-07-23), Flash VL (September 2026), Flash Fin (2026-09-03, built with financial institutions and domain experts), Flash Sante (2026-09-04). Significance for the RAG: the Ling line is the agent-economy cost floor — 124B-class MoE at near-zero/free API pricing, aimed at high-frequency agentic workflows (coding agents, document processing, long multi-turn conversations) rather than frontier reasoning.

### ByteDance (Seed team / Volcano Engine)

- Seed 2.1 family announced 2026-06-23/24 at the Volcano Engine FORCE conference: base Seed 2.1, Seed 2.1 Pro, preview tier, and (per a single trade report) Seed 2.1 Turbo. Official June 23/24 launch blog names no Turbo variant — the Turbo's presence at FORCE rests on one June 24 trade report (DataNorth) plus Western aggregator catalog dates (Aug 10–12) that are first-seen dates, not releases. ByteDance positions Seed 2.1 around "productivity delivery" — end-to-end software delivery from planning to verification — rather than raw benchmark leadership (secondary review). Seedance 2.5 (video, cross-ref §12) belongs to the same lab but is canonical elsewhere.

### Secondary and community actors

- OpenRouter: primary public-dating source for Qwen3.7 Flash (2026-07-27) and Ling 3.0 Flash (2026-07-23); serves Ling 3.0 Flash free tier; model IDs `inclusionai/ling-3.0-flash:free`, `qwen-image-3.0` also served via OpenRouter.
- Artificial Analysis: Intelligence Index (Qwen3.7-Max 56.6; Qwen3.8-Max ≈ 40, 4th open, v4.3 Sept 4, 2026; Qwen3.8-27B 52, size-class leader; Qwen3.6-27B 38; Ling 3.0 Flash VL 41.2).
- NVIDIA: August 12, 2026 deployment blog confirming "Alibaba released the open weights" + GB300 NVL72 reference deployment.
- Unsloth: community GGUF quantizations of Qwen3.8-2.4T-A95B [COMMUNITY].
- Coverage outlets that carried independent verification: Bloomberg (GA date 2026-08-03), Quartz, MarkTechPost, Dataconomy, EqualOcean, Straits Times, MLQ, tech-insider.org, gigazine.net, miraflow.ai, binance NS3.AI roundup, codenewsletter.ai, theaitrack.com, pulsemark.ai, eWeek, AlternativeTo, Brief.news, The Information (strategy framing).

### What each outlet verified (for the RAG's provenance layer)

- Bloomberg (2026-08-03): Qwen3.8-Max GA date — the single strongest primary-adjacent date anchor in the family. Quartz and zeronoise.ai (digest dated 2026-08-03) corroborate.
- MarkTechPost: three load-bearing dates — Qwen3.5-Omni 2026-03-30, Qwen3.8 preview 2026-07-19, Max GA 2026-08-03.
- gigazine.net (2026-08-17): English coverage of the Qwen3.8-27B drop, three days after the official 2026-08-14 announcement — pins the correction against the brief's "Aug 12."
- NVIDIA (2026-08-12 deployment blog): confirms "Alibaba released the open weights" and publishes the GB300 NVL72 reference configuration — the only vendor-neutral hardware floor on record.
- memeburn.com: two actor-context facts — Beijing's 2026-07-15 approval of Apple Intelligence powered by Qwen in China, and a Qwen3.8 explainer.
- tech-insider.org: Qwen3.8-Flash-Next architecture/benchmark coverage; explicit "Licensing: Closed, API-only" verdict on Qwen Image 3.0.
- LatestLY (2026-08-26): Qwen3.8-Flash open-weight launch detail.
- contracollective.com: Qwen3.8-Max-0902 vs Claude Opus 5 Max Code Arena WebDev comparison (September 5 table) [COMMUNITY].
- Medium analyses: practical-llm-systems ("what the numbers actually say" on Qwen 3.8 benchmarks), infoinletdev.five (6B-active vs 17B-active on Flash-Next), leucopsis (Qwen3.5-Plus review), 302.AI (Image 3.0 Pro head-to-head vs GPT-Image-2, 2026-08-11) [COMMUNITY].
- thecherrycreeknews.com: the artifact-vs-service framing ("an available model is a service you rent on the vendor's terms; an open-weight model is an artifact you hold") [COMMUNITY].
- The Information: the "potential shift away from open-source-first" strategy framing for Qwen3.5-Omni, via community analysis [SECONDARY].
- codenewsletter.ai: Audio-Visual Vibe Coding coverage for Omni.

## Timeline and context

### Qwen3.5 (February 2026) — the open-weight flag-planting

The year opened for Qwen with a dual launch on 2026-02-16: the Apache 2.0 open weights of Qwen3.5-397B-A17B and the proprietary hosted Qwen3.5-Plus with 1M default context, built-in tools, and adaptive tool use. The open 397B-A17B ran a hybrid Gated Delta Network + sparse MoE architecture (10 routed + 1 shared expert per token, ~4.3% activation, 512 experts) with native vision-language and 262K context extendable to 1M via YaRN. This dual-track pattern — permissive mid-tier weights beside a hosted flagship — recurs through every Qwen generation in 2026 and is the business-model signature of Alibaba's lab strategy. Cross-ref §2 for the full Feb→May skeleton (established in wave 3/02).

### Qwen3.5-Omni (2026-03-30/31, not April) — the API-only turn

The brief's "April" dating for Qwen3.5-Omni is wrong: Multiple outlets (MarkTechPost 2026-03-30, AlternativeTo 2026-03-31, eWeek, Brief.news) place the launch in the last week of March. Omni is the first major Qwen flagship with no open weights at all (API-only, Offline + Realtime APIs), following the Qwen3.5-Plus pattern and anticipating the Wan2.5/2.6 video-model closures. The Information coverage, cited in community analysis, frames this as "a potential shift away from a strategy that centers on open-source models." Technically it is Alibaba's native end-to-end omnimodal model — text/image/audio/video in one pipeline with a Thinker–Talker architecture and Hybrid-Attention MoE (not a text model with bolt-on audio encoders) — with tiers Plus/Flash/Light, 256K context on the flagship, >10 hours of audio per request, 100M+ hours of audiovisual pretraining, 113 languages for speech recognition and 36 for speech generation, plus voice cloning via API. Plus beats Gemini 3.1 Pro on audio benchmarks per vendor claims [VENDOR]. Distinctive use case: "Audio-Visual Vibe Coding" — generating functional websites/games from spoken descriptions plus camera input (codenewsletter.ai), an agentic-multimodal pattern unlike text-only coding.

### Qwen3.6 (April 2026) — proprietary Plus, open 35B, preview Max

Qwen3.6-Plus shipped 2026-04-02 as a proprietary API (agentic coding, 1M context, ¥2/¥12 on Bailian). Qwen3.6-35B-A3B followed 2026-04-16 as the generation's open tier (35B/3B MoE, open weights). A dense Qwen3.6-27B exists in third-party benchmark tables (Terminal-Bench 2.1 60.7, AA Index 38) but its release date was not pinned in research — flagged [DATE UNVERIFIED]. Qwen3.6-Max-Preview (2026-04-21) kept the Max line in private preview. Cross-ref §2.

### Qwen3.7 (May–July 2026) — the closed generation

