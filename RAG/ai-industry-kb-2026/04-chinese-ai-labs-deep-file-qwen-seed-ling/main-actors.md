---
id: ai-industry-kb-2026/04-chinese-ai-labs-deep-file-qwen-seed-ling/main-actors
title: "Main actors"
domain: chinese-ai-labs-deep-file-qwen-seed-ling
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-04", "2026-02", "2026-02-16", "2026-03-30", "2026-03-31", "2026-04", "2026-04-02", "2026-04-16", "2026-04-21", "2026-05", "2026-05-19", "2026-05-20", "2026-05-21", "2026-06-02", "2026-06-23", "2026-06-24", "2026-07", "2026-07-15", "2026-07-16", "2026-07-17", "2026-07-19", "2026-07-21", "2026-07-23", "2026-07-27", "2026-07-31", "2026-08", "2026-08-01", "2026-08-03", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-17", "2026-08-26", "2026-09", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "benchmark", "benchmarks", "blackwell", "claude", "cost", "datacenter", "disclosure"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1567, 1693]
section: "4. Chinese AI Labs — Deep File: Qwen, Seed, Ling"
sha256: 5529cb94955e465bdb7c2fb504d861797d8d45e599c0ed9f63ad8c08dda88137
---

# Main actors

## Main actors

### Alibaba Group / Alibaba Cloud / Qwen team

- Developer Shuai Bai named in connection with the Qwen3.8 2.4T announcement (first Qwen multimodal model above 1T params). Model Studio (QwenCloud) is the hosted-API surface; ModelScope is the domestic model hub mirroring Hugging Face publications. Token Plan, Qoder, and QoderWork carried the Qwen3.8-Max-Preview at 10% pricing. Beijing approved Apple Intelligence powered by Qwen in China on 2026-07-15 (memeburn.com) — distribution context. Alibaba holds ~36% of Moonshot (MLQ) — the ownership link to the Kimi line, relevant to the WAIC-week sequencing (Kimi K3 2026-07-17 → Qwen3.8 2026-07-19).

### Ant Group / InclusionAI

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

Every Qwen3.7 release was proprietary or hosted: Qwen3.7-Max unveiled 2026-05-20 at the Alibaba Cloud Summit/Apsara Conference in Hangzhou (API public 2026-05-21/23; 1M context, 65,536 output tokens; AA Index 56.6); Qwen3.7-Plus officially released 2026-06-02 (preview 2026-05-19 on the Qianwen site) as a multimodal agent-oriented model unifying GUI and CLI operations in a single closed loop — visual perception + screen reading + terminal code execution — claiming top-5 globally and #1 in China on Vision Arena [VENDOR]; scale known only through a third-party benchmark table (17B active [COMMUNITY]); Qwen3.7 Flash went public 2026-07-27 per OpenRouter (the QwenCloud `qwen3.7-flash-2026-07-15` identifier is a snapshot, not the launch date — cite OpenRouter for availability, QwenCloud separately for technical limits).

### Qwen3.8 (July–September 2026) — the 2.4T flagship and the license question

Announced 2026-07-19 at the World AI Conference in Shanghai, two days after Moonshot's Kimi K3 (2.8T; Alibaba holds ~36% of Moonshot per MLQ), Qwen3.8 is a 2.4-trillion-parameter sparse MoE — the first Qwen multimodal model above 1T params. The per-token active count was NOT disclosed at announcement (the most-flagged gap in independent technical reviews). Qwen3.8-Max-Preview was available immediately through the Token Plan, Qoder, and QoderWork at 10% of standard pricing (text/image/video/document input). General availability landed 2026-08-03 (Bloomberg, Quartz, MarkTechPost, zeronoise.ai digest, bitrouter spec): 2.4T total / ~95B active [third-party-reported], 1M context (991K max input, 131K max output), 512 experts (10 routed + 1 shared per token), 92 layers, flat pricing $2/M input / $6/M output (implicit cache reads $0.25/M, explicit $0.17/M; bitrouter CNY 12/36 base at ~7.2 CNY/USD).

#### The 2026-08-12 weight drop, dissected

Repositories `Qwen/Qwen3.8-2.4T-A95B` and `Qwen/Qwen3.8-2.4T-A95B-FP8` (official fine-grained FP8 quant, block size 128, mixed BF16/F8_E4M3) published on Hugging Face and ModelScope on 2026-08-12; NVIDIA's August 12 deployment blog confirms "Alibaba released the open weights." Three dimensions must not be collapsed:

1. **Weights published**: real — BF16 + official FP8 + community Unsloth GGUF quants [COMMUNITY]. First 2T-class downloadable model.
2. **License permissiveness**: NOT Apache 2.0. The HF repo lists a custom "qwen3.8-max" license — past 100M MAU or $20M monthly revenue you must display the model name; MaaS/AI-assistant businesses past $50M trailing revenue need a separate commercial license; internal use exempt (aicodingdir.com, explainx.ai). Reporting says Alibaba plans revenue-sharing with large commercial users of the weights. Consolidation must classify this as "open-weight with custom restrictive license," alongside Llama Community and the GLM-5.3 License — never in an Apache-2.0/MIT index.
3. **Feature parity**: the open checkpoint is text-only and thinking-mode-only — it drops vision input and the native 1M context the hosted API offers (explainx.ai; HF discussion thread opened within hours calling it a "huge disappointment," comparing it to DLC paywalling). "Open weights" here means the capability ceiling is public and auditable; the full product experience is API-gated. Community commentary (thecherrycreeknews.com) frames it: an available model is a service rented on the vendor's terms; an open-weight model is an artifact you hold.

The developer-relevant open artifact of the drop is really Qwen3.8-27B, not the 2.4T Max: 27B dense native multimodal, Apache 2.0, 262K native context extensible to 1M, supported in Transformers/vLLM/SGLang/Docker Model Runner with an official FP8 quant — ~14–16 GB VRAM at 4-bit, RTX 4090/5090 class (aicodingdir.com). The Max remains a datacenter-only artifact: BF16 ~4.89 TB, FP8 roughly half, cheapest usable quant ~450 GB RAM+VRAM; NVIDIA reference deployment on a GB300 NVL72 rack (72 Blackwell Ultra GPUs), >4,000 tok/s per GPU in FP8.

### Qwen3.8-Flash (2026-08-26) and Qwen3.8-Max-0902 (2026-09-02)

Qwen3.8-Flash (2026-08-26): multimodal MoE, 6B active, open weights, 262K→1M via YaRN, $0.16/$0.47 per M — training cost ~1/9 of Qwen3.7-Plus [VENDOR]; completes the 3.8 family at the efficient end. Qwen3.8-Flash-Next (same day): hosted-only variant on QwenCloud/Model Studio, a Qwen4-generation 125B/6B-active architecture preview with Alibaba-reported SWE-bench Pro 62.5 / GPQA Diamond 91.7 [VENDOR]. Qwen3.8-Max-0902 (2026-09-02): newer Max checkpoint, 4th preliminary in the September 5 Code Arena WebDev table — canonical home wave 2.1/05; recap only. The Max lineage reads: Qwen3-Max (1T, Sept 2025, estimated) → Qwen3.6-Max-Preview (Apr 2026, closed) → Qwen3.7-Max (May 2026, closed) → Qwen3.8-Max-Preview (Jul 19, 2026) → Qwen3.8-Max GA (Aug 3, 2026) → Qwen3.8-Max-0902 (Sept 2, 2026).

### Qwen Image 3.0 (2026-07-21) — closed departure

`qwen-image-3.0` / `qwen-image-3.0-pro` shipped on Model Studio (also OpenRouter) 2026-07-21 with no open weights, no technical report, and no published benchmarks — a departure from Qwen-Image 1.0's Apache 2.0 launch; closed, API-only (tech-insider.org). Stated differentiation vs Google's Nano Banana 2 (Sept 2026) and OpenAI's GPT Image 2.5 (Sept 8, 2026) is dense text/structured layouts, not aesthetics: 4,500-token instruction input, legible text down to ~10 px, dense single-pass multi-panel layouts (infographics, newspapers, storyboards), 12 languages, LaTeX math notation, web-page/game-interface/livestream-overlay simulation; max 2048×2048; 1–3 reference images for editing. Framed as "industrial-grade visual foundation" — "not just pursuing good-looking, it is pursuing useful." A 302.AI head-to-head (Aug 11, 2026) compared it against GPT-Image-2 in infographic/layout scenarios [COMMUNITY]; no public eval suite exists [UNVERIFIED].

### Ling 3.0 Flash (July–September 2026) — the agent-economy cost floor

InclusionAI (Ant Group's AI lab) shipped Ling 3.0 Flash 2026-07-23: 124B / ~5.1B active MoE, 262K context, 32K max output, reasoning + tool calling, text in/out (no vision), distributed as a free tier on OpenRouter (`inclusionai/ling-3.0-flash:free`) — which explains its benchmark presence as a new entrant. Intelligence 22.6 / coding 50.6–55.6 on aggregators [COMMUNITY]: modest vs frontier, competitive on cost. Siblings: Ling 3.0 Flash VL (vision, September 2026, AI Index 41.2), Ling 3.0 Flash Sante (2026-09-04, free variant, secondary-sourced), and Ling 3.0 Flash Fin (2026-09-03): the first finance-specialized model in the Ant Ling family, built with financial institutions and domain experts, 124B/5.1B active, 256K context, 236K max output, function calling + reasoning, $0.06/M input / $0.18/M output — the cheapest finance-domain model of the window [SECONDARY vendor-adjacent]. The Ling line, paired with Qwen3.8-Flash ($0.16/$0.47), is the evidence base for 2026's bifurcation: datacenter flagships vs ultra-cheap agent workhorses.

### Seed 2.1 Turbo (2026-06-24) — dated with reserve

ByteDance announced Seed 2.1, Seed 2.1 Pro, and a preview tier at the Volcano Engine FORCE conference on 2026-06-23/24. The official launch blog names no Turbo variant; the Turbo's presence at FORCE rests on a single June 24 trade report (DataNorth), while Western trackers only logged the model on August 10–12 (LLM Gateway PR #3580, NanoGPT catalog, OpenRouter) — first-seen dates, not releases (one engineering journal documents the "launched twice, 47 days apart" discrepancy). Defensible dating: 2026-06-24, with the single-source caveat kept visible. Verified specs via catalogs + LLM Gateway e2e mapping: 262K context (262,144); text/image/video in → text out; reasoning with full effort ladder (`none` disables thinking); tool calling; structured output; streaming; closed/proprietary (parameter count, cutoff, license all undisclosed [UNVERIFIED]). Pricing $0.50/$2.50 per M tokens (OpenRouter/NanoGPT, verified current 2026-09-22): the faster, lower-cost variant for high-throughput/latency-sensitive agent workloads vs the flagship Pro. ByteDance's launch materials used chart-only figures with superlatives and no absolute numbers in prose ("a masterclass in avoiding numbers in the text" — one journal); no independent Turbo benchmarks exist [UNVERIFIED]. Seedance 2.5 (video) is cross-referenced to §12 and not duplicated here.

### August 2026, day by day — the month the flagship split

- 2026-08-01: MarkTechPost launch-day coverage window for the late-July releases still landing (MiniMax H3 coverage dated Aug 1 belongs to the July cluster's tail).
- 2026-08-03: Qwen3.8-Max GA — Bloomberg, Quartz, MarkTechPost, zeronoise.ai, bitrouter registry spec all pin the date; $2/$6 flat pricing; 1M context; 512 experts.
- 2026-08-11: 302.AI publishes the Qwen Image 3.0 Pro vs GPT-Image-2 head-to-head in infographic/layout scenarios — the only independent comparative data point on Image 3.0 [COMMUNITY].
- 2026-08-12: Qwen3.8-Max weights drop (`Qwen/Qwen3.8-2.4T-A95B` + FP8 on HF/ModelScope) the same day NVIDIA publishes its GB300 NVL72 deployment blog confirming "Alibaba released the open weights" — hardware and weights landing together is the tell that the artifact was staged for datacenter partners first.
- 2026-08-14: Qwen3.8-27B weights under Apache 2.0 (official @Alibaba_Qwen announcement: "We promised open weights for Qwen3.8. Now, time to meet them!") — the sentence itself confirms the two-step structure: Max weights were the promise kept on the 12th, the 27B was the democratization event on the 14th.
- 2026-08-17: gigazine.net English coverage — the date that proves the brief's "Aug 12" for the 27B was wrong.
- 2026-08-26: Qwen3.8-Flash (open weights, 6B active) + Qwen3.8-Flash-Next (hosted, Qwen4-generation preview) — the efficiency tier closing out the month.

### Qwen3.8-Max-0902 (2026-09-02) — recap only

Newer Max checkpoint dated 2026-09-02; 4th (preliminary) in the September 5 Code Arena WebDev table vs Claude Opus 5 Max (contracollective.com comparison) [COMMUNITY]. Canonical home: wave 2.1/05-open-models-delta.md — this part records only the date, the checkpoint's existence, and the cross-reference. No license change, no spec change, no independent benchmark package beyond the Code Arena table at cutoff.

### Naming guard

Qwen3.8-Max ≠ Qwen3-8B (the 8.2B Apache-2.0 dense model released April 2025) — search engines conflate them; consolidation indexing must keep them separate.

### The WAIC announcement (2026-07-19): what was and wasn't disclosed

- Announced at the World AI Conference in Shanghai; coverage within 96 hours from MLQ, MarkTechPost, Dataconomy, EqualOcean, Straits Times, memeburn — the best-attested single date in the Qwen 2026 line.
- Disclosed: 2.4T total parameters, sparse MoE, first Qwen multimodal model above 1T params (developer Shuai Bai), text/image/video/document input, immediate Qwen3.8-Max-Preview via Token Plan/Qoder/QoderWork at 10% of standard pricing, the vendor's "just behind Claude Fable 5" positioning [VENDOR].
- NOT disclosed: active-per-token count (the gap every independent technical review flagged), no benchmark table, no model card, no license. Independent test results did not exist at announcement.
- Sequencing: two days after Moonshot's Kimi K3 (2.8T) — Alibaba holds ~36% of Moonshot (MLQ) — so the announcement reads as a same-family one-upmanship as much as a competitive response; EqualOcean explicitly frames the preview as promising an open-weight release, a promise the 2026-08-12 drop kept under custom terms.
- RAG guidance: treat the announcement and the GA (2026-08-03) as distinct events with distinct disclosure states — the "just behind Fable 5" claim belongs to the announcement, not to the GA benchmark package.

### The two-track strategy, quarter by quarter

Q1 2026 (Feb–Mar): the pattern is set at launch — permissive open weights (397B-A17B, Apache 2.0) ship the same day as the proprietary hosted flagship (Plus, 1M ctx). By late March, Omni extends the hosted track into omnimodal territory with no weights at all, and The Information reads it as a possible open-source retreat — but the open track never stops.

Q2 2026 (Apr–Jun): Qwen3.6 repeats the pattern (proprietary Plus 2026-04-02, open 35B-A3B 2026-04-16) while the entire Qwen3.7 generation ships closed — Max unveiled at the May Hangzhou summit, Plus official June 2 as an agent-oriented GUI+CLI model, and the Turbo-adjacent efficiency story left to ByteDance (Seed 2.1 family at FORCE, June 23/24) and Ant's InclusionAI (Ling 3.0 Flash, July 23 — technically Q3 but conceived in the same efficiency wave).

Q3 2026 (Jul–Sep): the July cluster — Kimi K3 (2026-07-16/17, 2.8T, Moonshot, ~36% Alibaba-owned), Qwen3.8 announced (2026-07-19, 2.4T), Qwen Image 3.0 closed (2026-07-21), Ling 3.0 Flash free (2026-07-23), Qwen3.7 Flash public (2026-07-27), MiniMax H3 (2026-07-31). Five Chinese labs shipping frontier-class or frontier-adjacent releases within fifteen days. August resolves the Max question in two steps: GA on the 3rd ($2/$6, 1M ctx, multimodal) and the weight drop on the 12th under the custom license with the text-only/thinking-only artifact caveat; the Apache-2.0 27B on the 14th is the actual democratization event. Late August adds the efficiency tier (3.8-Flash open, Flash-Next hosted preview of Qwen4-generation architecture). September adds the 0902 checkpoint and the finance-specialized Ling Fin.

The net arc: Alibaba moved from binary open/closed (Qwen3.5's clean Apache-vs-hosted split) to a three-way license taxonomy (permissive / custom-restrictive / closed-API) while holding both tracks open all year. The "shift away from open-source" narrative is only accurate for the flagship modality frontier (omni, image, Max-hosted); the mid-tier open-weight pipeline accelerated in parallel. The correct consolidation framing is "two tracks converging on terms-as-monetization," not "open then closed."

### Qwen3.5-Omni and the open-weight strategy question (deep dive)

Omni matters less as a model than as a strategy signal — it is the first major Qwen flagship to ship with no open weights, and the template for everything after it (Qwen3.7-Max/Plus, Image 3.0, the Max weight drop's custom license). The Information's "potential shift away from a strategy that centers on open-source models" framing is directionally right but incomplete: what actually happened is modality-gated openness. Text mid-tier models stayed Apache-2.0 all year; the models that crossed a product-maturity threshold in a new modality (omnimodal understanding, image generation, agent GUI/CLI control) shipped hosted or custom-licensed. The license follows the modality's monetizability, not the parameter count — the 2.4T Max got weights (under custom terms) while the undisclosed-scale Omni got none, because a downloadable omni model is a product in a way a text checkpoint is an ingredient.

Technical differentiators worth indexing for the multimodal RAG: Thinker–Talker architecture with Hybrid-Attention MoE across all modalities — native end-to-end, not cascaded text-model-plus-encoders; 256K context; >10 hours of audio and >400 seconds of 720p video per request; 100M+ hours of audiovisual pretraining; 113 languages speech recognition / 36 speech generation; voice cloning via API; tiers Plus/Flash/Light. The distinctive agentic use case is Audio-Visual Vibe Coding: functional websites/games generated from spoken descriptions plus live camera input (codenewsletter.ai) — multimodal agency rather than text-only coding, and the clearest preview of where Qwen3.7-Plus's GUI/CLI agent loop was headed three months later.

