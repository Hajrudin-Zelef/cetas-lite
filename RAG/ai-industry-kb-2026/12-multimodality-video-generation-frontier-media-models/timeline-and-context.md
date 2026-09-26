---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/timeline-and-context
title: "Timeline and context"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["Apple", "ByteDance", "China", "EU", "Google", "Meta", "MiniMax", "OpenAI", "United States", "xAI"]
dates: ["2025-08-05", "2025-09-30", "2025-12-16", "2026-01-29", "2026-02-05", "2026-02-12", "2026-04", "2026-04-26", "2026-05", "2026-05-12", "2026-06-23", "2026-07-01", "2026-07-31", "2026-09", "2026-09-15", "2026-09-17", "2026-09-24"]
keywords: ["agentic", "arr", "benchmark", "chatgpt", "compute", "consumer", "copyright", "cost", "distribution", "gemini", "grok", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6379, 6418]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 17fc333ab40c305d9a72db43c3ef5e4bc028c687776e81adc49eb03b744321e8
---

# Timeline and context

- **OpenAI:** Sora 2 (consumer product discontinued 2026-04-26, API cutoff 2026-09-24 planned; model survives in ChatGPT paid tiers; team → world-simulation research); GPT Image 1.5 family (native image generation); DALL-E 2/3 retired 2026-05-12.
- **Google / DeepMind:** Gemini Omni Flash (any-to-any video benchmark-setter, $0.10/s); Veo 3.1 (sparse verification — distributed, no surfaced 2026 launch record); Nano Banana image family (2 Lite entry tier ~2026-07-01/02; Pro via Gemini 3 Pro Image); Project Genie (2026 consumer beta + Street View grounding; no Genie 4); Agentic Video Understanding; SynthID watermarking and C2PA credentials by default on Omni Flash.
- **ByteDance Seed:** Seedance 2.0 (China launch 2026-02-12) and 2.5 (announced 2026-06-23, Dreamina 2026-07-31) — 30s single-pass, native 4K claims [VENDOR/secondary]; international launch paused over copyright disputes.
- **Kuaishou:** Kling 3.0 (global launch 2026-02-05) and Kling 3.0 Omni (O3) — 4K60, 6-shot storyboarding, 5-language lip sync, Motion Control AI; reported $240M ARR Dec 2025 [UNVERIFIED]; the revenue counterpoint to Sora.
- **Runway:** Gen-4.5 — commercially distributed via Adobe Firefly; claimed beaten by Seedance 2.0 on AA at launch [UNVERIFIED].
- **MiniMax:** H3 — open-weight video leader on at least one public ranking under a restrictive community license (US/EU/UK/SK excluded from local deployment; → §11).
- **HiDream.ai (Beijing):** HiDream-O1-Video-1.0 — company-reported native omnimodal video entrant (2026-09-17) [VENDOR].
- **Adobe Firefly:** aggregation model — 30+ video models by April 2026, rents inference, sidesteps the training-compute bill.
- **xAI:** Grok Imagine (image/video/voice via separate Imagine API) → §1; assembly-not-any-to-any stack.
- **Meta Movie Gen:** no primary material surfaced in this pass sufficient to add — deliberately omitted.
- **Dreamina (ByteDance):** consumer distribution surface for Seedance 2.5's global launch (2026-07-31) — the venue, not a separate model; international rollout is where the copyright-gating question applies.
- **Novita AI:** third-party API distributor documenting Kling 3.0's global launch (2026-02-05) and its per-second pricing — representative of the B2B distribution layer that made Kling's economics work.
- **Hypereal/Hyperreal:** commercial PR distributor for Seedance 2.0 API availability — representative of the commercial-secondary layer behind ByteDance's API pricing claims.
- **fal:** post-trained H3 Max variant (5s 768p in under 3s; 75% launch discount expired 2026-09-15, price $0.10→$0.40) — the post-training-as-a-service layer on top of open-weight video models.
- **the-decoder:** independent secondary outlet whose reporting made H3 "the first open model to top an AI video ranking" — the single best-corroborated leaderboard claim in this section.
- **ngram.com / CIOL:** analytical voices on the Sora shutdown — ngram's deprecation-page analysis (empty "replacement" column), CIOL's cost/revenue figures [UNVERIFIED] and the "side quests" reporting chain.
- **The likeness-rights estates (MLK Jr., Robin Williams):** not labs, but active constraints — their complaints shaped Sora's moderation posture and remain the concrete case law-adjacent evidence that likeness moderation is a product-cost factor in video generation.
- **Waymo** [UNVERIFIED connection]: claimed user of a fine-tuned Genie variant for driving simulation — secondary blogs only, no primary confirmation; included because it illustrates the world-model-to-simulation use case Genie is positioned for.

- **OpenAI Applications org (Fidji Simo):** the internal actor behind the Sora shutdown call — Simo's "side quests" memo to employees and Altman's redirection decision; the org-level evidence that video was deprioritized by the applications business, not by research.
- **Disney:** the reported $1B franchise-content partnership affected by the shutdown (WSJ: informed shortly before the announcement, no money changed hands) — the enterprise-partnership casualty that shows shutdown blast radius beyond consumers.
- **Volcano Engine (ByteDance):** the announcement and distribution venue for Seedance 2.0 (Feb 2026 launch tracking) and Seedance 2.5 (FORCE conference 2026-06-23) — ByteDance's cloud/API arm as the primary-source channel, in the absence of ByteDance-primary English materials.
- **Google Labs:** the distribution channel for Project Genie's consumer beta (2026-01-29) — Labs-as-launch-vehicle for the $249.99/mo-gated experiment.
- **C2PA / SynthID provenance standards:** Omni Flash shipped with C2PA credentials and SynthID watermarks on by default — provenance-by-default as a 2026 product norm for flagship video models.
- **ABNewswire (PR syndication):** the syndication channel for Kuaishou's September 2026 "AI Director" PR wave — representative of how company positioning becomes "industry trend" coverage in this domain (cf. the wave3 narrative audit, out of scope here).
- **humai.blog:** single-source origin of the "Sora resources redirected to Spud" claim [UNVERIFIED] — recorded as an actor only to pin the provenance of that one claim.
- **Novita AI and Hypereal/Hyperreal:** the third-party API-distribution layer (Kling 3.0 pricing/launch documentation; Seedance 2.0 API PR) — the commercial-secondary layer that makes Chinese labs' models globally consumable and is the actual source of most public per-second pricing.
- **Google AI Ultra ($249.99/mo):** not a model but a business model — the subscription gate behind Project Genie's consumer beta and its May 2026 global expansion; the price point at which interactive world models are currently monetized (two orders of magnitude above typical chatbot tiers).

## Timeline and context

### 12.13 Pre-2026 anchors (context, not 2026 events)

- **2025-08-05:** Google DeepMind announces Genie 3 (interactive world model; 720p/24fps; promptable world events).
- **2025-09-30:** OpenAI launches Sora 2 as a standalone iOS app (synchronized audio, Cameos, Sora 2 Pro; US/Canada first). Not 2026.
- **2025-12-16:** gpt-image-1.5 GA (~4x faster, 20% cheaper, RGBA-native).
- **2025-12:** Kling reportedly at $240M ARR [UNVERIFIED, single secondary compilation].

### 12.14 Dated release timeline (Feb → Sep 2026)

