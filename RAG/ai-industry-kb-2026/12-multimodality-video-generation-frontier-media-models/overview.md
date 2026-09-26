---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/overview
title: "12. Multimodality — Video Generation & Frontier Media Models"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["ByteDance", "China", "EU", "Google", "MiniMax", "OpenAI", "United States"]
dates: ["2025-09-30", "2026-01-29", "2026-02-05", "2026-02-12", "2026-04", "2026-04-26", "2026-05-12", "2026-05-19", "2026-05-21", "2026-06-23", "2026-07-01", "2026-07-31", "2026-09", "2026-09-17", "2026-09-22", "2026-09-24"]
keywords: ["multimodal", "video generation", "agentic", "attribution", "chatgpt", "compute", "consumer", "cost", "distribution", "gemini", "leaderboard", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6152, 6173]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: 7485a2ff2cf886c1fd1832c7824531864c156dd69bc2be25acff80903582e213
---

# 12. Multimodality — Video Generation & Frontier Media Models
Keywords: Sora 2, Sora shutdown, Gemini Omni Flash, Project Genie, Genie 3, Seedance 2.5, Seedance 2.0, Kling 3.0, Veo 3.1, Runway Gen-4.5, Nano Banana 2 Lite, GPT Image 1.5, DALL-E sunset, HiDream-O1-Video-1.0, MiniMax H3, synchronized audio generation, any-to-any video, native multimodal video, AI video economics, agentic video understanding, world models, text-to-video leaderboards, Adobe Firefly

## Summary

- By 2026-09-22 the video-generation industry had inverted its 2024–2025 narrative: quality kept climbing, but the only consumer product that died — OpenAI's Sora — died on unit economics, while monetization converged on B2B/API, aggregation, and expensive subscription gating.
- OpenAI's Sora 2 launched 2025-09-30 (not 2026), its web/app was killed 2026-04-26, and its Videos API was slated for permanent shutdown on 2026-09-24 — a date still in the future as of 2026-09-22 and requiring post-date reverification; the Sora 2 model itself survives inside ChatGPT paid tiers and the team continues as a world-simulation research unit.
- Google's Project Genie had a distribution year, not a model year: consumer beta 2026-01-29 via Google Labs behind the $249.99/mo AI Ultra plan, and a 2026-05-21 expansion with Street View/Maps Imagery Grounding; no Genie 4 or new numbered model was found as of 2026-09-22.
- Google I/O 2026 (2026-05-19) unveiled Gemini Omni Flash — the closed flagship for any-to-any video output (single-pass video with synchronized audio, conversational editing) — with the nuance that audio-reference inputs and scene extension were not supported at preview; Nano Banana 2 Lite shipped ~2026-07-01/02 in the developer rollout, not at I/O.
- The text-to-video leadership became a China-led, closed-model race: ByteDance's Seedance 2.0 (China launch 2026-02-12) and 2.5 (announced 2026-06-23 at Volcano Engine FORCE, global Dreamina launch 2026-07-31) with 30-second single-pass clips and native 4K; Kuaishou's Kling 3.0 (global launch 2026-02-05), whose September 11–12 "AI Director" PR wave was repositioning, not a new model.
- The open-weight side produced exactly one leaderboard winner: MiniMax H3 (API 2026-07-31, open weights Aug 2–3, 2026) topped Artificial Analysis Video Editing (With Audio) under a restrictive community license that excludes US/EU/UK/South Korea from local deployment (→ §11 for the open-weight taxonomy); HiDream-O1-Video-1.0 (2026-09-17, company-reported) claims native omnimodal video with synchronized audio.
- Native multimodal output generation consolidated: DALL-E 2/3 sunset 2026-05-12 in favor of the GPT Image 1.5 family (~4x faster, 20% cheaper, RGBA-native) is the structural evidence that stitched pipelines are being retired; agentic video understanding (Google's September 2026 announcement: −88% tokens, +~7% accuracy, −66% cost) moved the video economics conversation from generation to perception.
- The dominant business lesson of 2026 in this domain: technology leadership ≠ product viability. Survivors monetize APIs (Kling, Seedance), aggregate others' models (Adobe Firefly, 30+ models by April 2026), or gate behind premium subscriptions (Genie at $249.99/mo); the failed product (Sora) reported ~$5.4B annual compute spend against ~$2.1M lifetime revenue [UNVERIFIED, secondary].
- Attribution risk is concentrated in this section: Sora 2's launch year (2025, not 2026), Nano Banana 2 Lite's venue (~July dev rollout, not I/O), Seedance 2.5's timing (June/July, not August), and Kling's September wave (repositioning, not a launch) were all misdated by secondary roundups — each is corrected with primary-adjacent evidence in the dated facts below.
- The any-to-any frontier split in September 2026 into shipped-outputs (Omni Flash, H3: synchronized audio in one pass), claimed-outputs (HiDream, Seedance 2.5 specs: [VENDOR/secondary]), and roadmap-inputs (Omni Flash audio references) — the section keeps that tripartition explicit wherever the "single forward pass" claim appears.
- Two deliberate product sacrifices define the section's two poles: OpenAI killed working image products (DALL-E) to consolidate onto native multimodality, and killed a best-in-class video product (Sora) on economics — architecture-first consolidation on one side, viability-first triage on the other.
- The 2026-09-24 Sora API cutoff is the only future-dated fact in this part file; everything else is retrospective as of 2026-09-22 and should hold without reverification.

## Key dated facts

### 12.1 Sora 2: launched 2025-09-30, killed as a product in 2026

