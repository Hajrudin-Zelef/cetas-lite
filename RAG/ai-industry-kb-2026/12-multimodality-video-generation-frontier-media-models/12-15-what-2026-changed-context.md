---
id: ai-industry-kb-2026/12-multimodality-video-generation-frontier-media-models/12-15-what-2026-changed-context
title: "12.15 What 2026 changed (context)"
domain: multimodality-video-generation-frontier-media-models
role: deep-dive
task: multimodal
actors: ["ByteDance", "China", "Google", "Hugging Face", "MiniMax", "OpenAI", "United States"]
dates: ["2026-01-29", "2026-02-05", "2026-02-12", "2026-03-24", "2026-04-26", "2026-05-12", "2026-05-19", "2026-05-21", "2026-06-23", "2026-06-30", "2026-07-01", "2026-07-31", "2026-08-02", "2026-08-31", "2026-09-02", "2026-09-11", "2026-09-15", "2026-09-17", "2026-09-22", "2026-09-24"]
keywords: ["agent", "agentic", "benchmark", "chatgpt", "consumer", "diffusion", "gemini", "license", "omni", "open weights", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6419, 6455]
section: "12. Multimodality — Video Generation & Frontier Media Models"
sha256: e427f575da08c42fd43f823677b5366732596dd5e926b07ce69fd8a240485b5b
---

# 12.15 What 2026 changed (context)

| Date (2026) | Event |
|---|---|
| 2026-01-29 | **Project Genie** consumer beta via Google Labs, gated on Google AI Ultra ($249.99/mo) |
| 2026-02-05 | **Kling 3.0** global launch (15s clips, native multilingual audio) |
| 2026-02-12 | **Seedance 2.0** China launch (quad-modal refs, native audio, 2K) |
| 2026-02 (month) | Last month with all three consumer video products live: Sora (still up), Kling 3.0 (new), Seedance 2.0 (China) |
| 2026-03-24 | **Sora shutdown announced** (official Sora account); Sora 2 API deprecation notice to developers |
| 2026-04 (month) | **Adobe Firefly passes 30+ video models** (Veo 3.1, Runway Gen-4.5, Kling 3.0 distributed) — aggregation as a business model |
| 2026-04-26 | **Sora web/app discontinued** [VERIFIED — OpenAI Help Center] |
| 2026-08-31 | OpenAI Help Center discontinuation article updated (~2026-08-31) — confirms 2026-04-26 web/app death, 2026-09-24 API death, sora.chatgpt.com/sunset export, permanent data deletion |
| 2026-09-02 | Vendor-exit migration guide (dev.to) retrieved quoting OpenAI's provider-owned discontinuation notice — API end 2026-09-24 restated |
| 2026-05-12 | **DALL-E 2/3 sunset** — pipeline-era generation retired for GPT Image family |
| 2026-05-19 | **Google I/O 2026** — Gemini Omni (Flash) unveiled ("Nano Banana for video") |
| 2026-05-21 | **Project Genie expansion** — Street View/Maps Imagery Grounding (US first), broader global AI Ultra access |
| 2026-06-23 | **Seedance 2.5 announced** at Volcano Engine FORCE (30s single-pass, native 4K, 50 references [VENDOR/secondary]) |
| 2026-06-30 | **Gemini Omni Flash API GA** ($0.10/s 720p; $1.50/M input tokens) |
| ~2026-07-01/02 | **Gemini Omni Flash developer rollout + Nano Banana 2 Lite launch** (corrected dates — not I/O) |
| 2026-07-31 | **Seedance 2.5 global launch on Dreamina**; MiniMax H3 API launch (Hailuo) |
| 2026-08-02/03 | **MiniMax H3 open weights** land on Hugging Face (sources differ by one day) |
| 2026-08 (month) | **StreamArena** released (streaming video-agent benchmark) |
| 2026-09 (~early) | **Google Agentic Video Understanding** announced (−88% tokens, +~7% accuracy) |
| 2026-09-11/12 | **Kling 3.0 "AI Director" PR wave** — repositioning of the February model, not a new release |
| 2026-09-15 | fal H3 Max 75% launch discount expires ($0.10→$0.40 per 5s clip) |
| 2026-09-17 | **HiDream-O1-Video-1.0 announced** (native omnimodal video, company-reported) |
| 2026-09-24 | **Sora API permanent shutdown** — PLANNED, future relative to 2026-09-22; reverification required |

### 12.15 What 2026 changed (context)

- The bottleneck moved from quality to unit economics: the only consumer video product that died was the one with the worst reported economics.
- The commercial leaders are ByteDance (Seedance), Kuaishou (Kling), and Google (Veo/Genie) — two Chinese labs and one US lab — while OpenAI, the 2024–2025 video flagship, exited the consumer product entirely.
- Generation commoditized into aggregation (Firefly, 30+ models) while the any-to-any frontier moved to single-pass synchronized audio-video (Omni Flash, H3, HiDream claims) and agentic perception (AVU) moved the economics of video *understanding*.
- **The "announced vs shipped vs PR" trichotomy** became the section's governing discipline: announced (Genie beta dates, Seedance 2.5 specs), shipped (Omni Flash API GA, H3 weights, Kling pricing via Novita), PR (Kling's September "AI Director" wave, HiDream's launch PR). Secondary roundups routinely promote PR to "launch" and announcements to "availability"; the dated-facts section above demotes each to its evidence level.
- **China-first launch sequencing** is now a structural feature of the video race: Seedance 2.0 (China-only, Feb 2026), Seedance 2.5 (Dreamina global, Jul 2026), Kling 3.0 (global from day one, Feb 2026), H3 (open weights global, Aug 2026, but license-restricted in the West). "Global launch" means different availability sets for different labs — the availability geography must be recorded alongside the date.
- **The 30s/4K/lip-sync bundle** (Seedance 2.5 claims, Kling 3.0 O3 PR) is the 2026 answer to "what does the next video model add": not a new architecture announcement, but longer single-pass duration, higher native resolution, and better audio fidelity — capability stacking on the existing diffusion-transformer substrate, marketed through repositioning waves rather than version numbers.

## Implications

