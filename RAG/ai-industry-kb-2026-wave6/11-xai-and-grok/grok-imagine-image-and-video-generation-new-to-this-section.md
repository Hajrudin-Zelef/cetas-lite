---
id: ai-industry-kb-2026-wave6/11-xai-and-grok/grok-imagine-image-and-video-generation-new-to-this-section
title: "Grok Imagine — image and video generation (new to this section)"
domain: xai-and-grok
role: deep-dive
task: multimodal
actors: ["ByteDance", "xAI"]
dates: ["2025-08-04", "2025-08-07", "2025-10", "2026-01", "2026-01-21", "2026-01-29", "2026-01-30", "2026-05", "2026-07-05", "2026-08-07", "2026-11-02"]
keywords: ["grok", "video generation", "diffusion", "distribution", "leaderboard", "parameters", "text-to-image", "text-to-video"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5328, 5345]
section: "§11. xAI and Grok"
delta_of: ai-industry-kb-2026
sha256: c6bbb934d6b86719ec644cffb5d209e2cf5327d51eec49f4a9fa0d8f0d4ccf6c
---

# Grok Imagine — image and video generation (new to this section)

### Grok Imagine — image and video generation (new to this section)
- Grok Imagine was officially announced 2025-08-04 (beta, paid: SuperGrok and Premium+; Android access from 2025-08-07); xAI removed the paywall 2025-08-07 making image-to-video free for all Grok app users. One timeline places a separate October 2025 debut — the discrepancy is preserved. [SECONDARY, S59][SECONDARY, S61][SECONDARY, S37]
- It is integrated directly in the X app, giving it distribution across hundreds of millions of users. [SECONDARY, S37 — single source]
- Under the hood it uses xAI's proprietary Aurora model for text-to-image generation. [SECONDARY, S37][SECONDARY, S36]
- The Aurora engine is an autoregressive image/model network processing text, image, video, and audio tokens jointly — replacing diffusion-transformer approaches — which yields native audio-video sync in a single generation step; a January 2026 third-party API launch announcement independently describes Aurora as xAI's internal autoregressive image model. [SECONDARY, S36][SECONDARY, S81]
- On 2026-01-21 xAI extended maximum video length from 5 seconds to 10 seconds with improved visual quality, smoothness, and audio synchronization. [SECONDARY, S33 — single source]
- The Grok Imagine API launched 2026-01-29 (third-party API platform Pixazo announced availability 2026-01-30) for text-to-video, image-to-video, and prompt-driven video edits with synchronized audio. [SECONDARY, S35][SECONDARY, S81]
- API parameters: clip lengths 1–15 seconds, 480p or 720p resolution, aspect ratios including 16:9, 4:3, 1:1, 9:16, 3:4, 3:2, 2:3 (and 20:9 per launch posts). [SECONDARY, S35 — single source]
- Requests are processed as deferred jobs with SDK auto-polling. [SECONDARY, S35 — single source]
- xAI claimed the #1 position in Artificial Analysis text-to-video rankings and reported human side-by-side preference on IVEBench at 1280×720 versus Kling o1 and Runway Aleph. [VENDOR, S35 — single source]
- In late May 2026 Grok Imagine video debuted atop the Artificial Analysis Video Arena image-to-video leaderboard (Elo 1404 ±6), displacing ByteDance Seedance 2.0. [SECONDARY, S36 — single source]
- Elon Musk posted "Done with Grok Imagine" on 2026-07-05, signaling the end of the core development cycle and a move to a polished/stable phase. [SECONDARY, S37 — single source]
- Grok Imagine Image 2.0: in the August 7, 2026 Arena snapshot cited by xAI, it ranked second globally in both text-to-image generation and image editing. [SECONDARY, S32 — single source]
- The Imagine API documents five model ids: `grok-imagine-image`, `grok-imagine-image-2.0`, `grok-imagine-image-quality`, `grok-imagine-video`, `grok-imagine-video-1.5`. [SECONDARY, S29 — single source]
- `grok-imagine-image-quality` is scheduled to retire 2026-11-02 in favor of `grok-imagine-image-2.0`, with a dedicated November 2 migration guide. [SECONDARY, S29 — single source]
- Grok Imagine video 1.5 generates 720p video at 24fps with native synchronized audio and topped third-party image-to-video leaderboards on release. [SECONDARY, S12 — single source]
- xAI also shipped 21 multilingual voices alongside the Imagine updates. [SECONDARY, S12 — single source]

