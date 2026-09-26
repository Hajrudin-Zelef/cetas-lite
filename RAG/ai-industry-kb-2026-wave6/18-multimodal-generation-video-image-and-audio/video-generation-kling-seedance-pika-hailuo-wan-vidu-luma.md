---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/video-generation-kling-seedance-pika-hailuo-wan-vidu-luma
title: "Video generation — Kling, Seedance, Pika, Hailuo, Wan, Vidu, Luma"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["ByteDance", "Google"]
dates: ["2025-10-28", "2025-12", "2025-12-11", "2026-02-05", "2026-04-23", "2026-05", "2026-08", "2026-08-25"]
keywords: ["video generation", "agent", "agents", "cost", "distribution", "leaderboard", "mcp", "pricing", "robotics", "text-to-video", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8822, 8842]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: 4e5119b076c94412070a63b2d7cd0eb2370b216ddf737c50ebd17d80fdb5f586
---

# Video generation — Kling, Seedance, Pika, Hailuo, Wan, Vidu, Luma

- [SECONDARY] Runway launched Gen-4.5 in late 2025 and topped the Artificial Analysis text-to-video leaderboard at launch with 1,247 Elo; by May 2026 it had been displaced by Seedance 2.0, HappyHorse-1.0 and the Kling 3.0 / Veo 3.1 cluster and no longer appeared in the top 10. (https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_video_generation_ai_models.md ; https://indianexpress.com/article/technology/artificial-intelligence/runway-launches-its-first-world-model-as-it-upgrades-gen-4-5-with-native-audio-10416477/)
- [SECONDARY] TechCrunch (December 11, 2025) reported a Gen 4.5 update adding native audio (lip sync, environmental SFX) and long-form multi-shot generation — up to one-minute videos with character consistency, native dialogue and background audio, editable dialogue and multi-shot videos of any length — released "earlier in the month", i.e. December 2025. (https://techcrunch.com/2025/12/11/runway-releases-its-first-world-model-adds-native-audio-to-latest-video-model ; https://indianexpress.com/article/technology/artificial-intelligence/runway-launches-its-first-world-model-as-it-upgrades-gen-4-5-with-native-audio-10416477/)
- [SECONDARY] Runway positioned Gen-4.5 as the foundation for GWM-1, its General World Model family: GWM-Worlds (interactive prompt/image-driven world exploration), GWM-Robotics (synthetic data for robot policy training), GWM-Avatars (realistic human avatars); the company said the three would eventually merge into one model. (https://techcrunch.com/2025/12/11/runway-releases-its-first-world-model-adds-native-audio-to-latest-video-model ; https://indianexpress.com/article/technology/artificial-intelligence/runway-launches-its-first-world-model-as-it-upgrades-gen-4-5-with-native-audio-10416477/)
- (single secondary coverage) [SECONDARY] GWM-Worlds was reported as real-time interactive simulation at 24 fps and 720p resolution, pitched at gaming and at training agents to navigate the physical world. (https://techcrunch.com/2025/12/11/runway-releases-its-first-world-model-adds-native-audio-to-latest-video-model)
- [COMMUNITY] Gen-4.5 control surface: Motion Brush (paint which parts of an image move and how), scene consistency from a single reference image, keyframes, video-to-video, and a Gen-4 Turbo variant for fast drafts. (https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_video_generation_ai_models.md ; https://github.com/dayahimour/ai-dayahimour/blob/HEAD/src/content/blog-en/runway-gen-4-5.md)
- (single secondary coverage) [SECONDARY] As of August 2026 Runway hosted third-party models inside the same credit system — Google Veo 3.1, ByteDance Seedance 2.5 and others — functioning as a multi-model video marketplace; the platform also exposes Act-Two (character animation), Aleph 2.0 (footage editing), a Studio editor, a Runway Agent for marketing-campaign workflows, and an MCP server for external agents. (https://aiweekly.co/learning-ai/generative-ai/how-to-use-runway)
- [SECONDARY] Per-generation cost for Gen-4.5 is 12 credits per second of video (standard plans), i.e. a 5–10 second clip costs 60–120 credits; Gen-4 Turbo is 5 credits/second. (https://aiweekly.co/learning-ai/generative-ai/how-to-use-runway ; https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_video_generation_ai_models.md)
- (single secondary coverage) [SECONDARY] CONTRADICTION — AdCreate's review states Gen-4.5 costs 25 credits/second while AI Weekly (verified against official docs 2026-08-25) states 12 credits/second; the difference may reflect plan-era or pricing changes, so both figures are preserved.

### Video generation — Kling, Seedance, Pika, Hailuo, Wan, Vidu, Luma

- (single secondary coverage) [SECONDARY] CONTRADICTION — Kling 3.0 release is reported as February 4 versus February 5, 2026 across sources; and native 4K at launch conflicts with an April 23, 2026 platform 4K date. Both date pairs are preserved.
- [DIRECTIONAL] Kling 3.0 pricing varies by gateway, resolution and audio setting — no single canonical price exists across sources.
- (single secondary coverage) [SECONDARY] CONTRADICTION — Seedance 2.5 resolution reporting conflicts: 480p/720p in some catalogs versus 1080p in others, with 4K appearing in marketing but unsupported in the product.
- (single secondary coverage) [SECONDARY] ByteDance's Seedance 2.0/2.5 sits in the same third-party marketplace interfaces (Runway) as Veo 3.1, indicating its commercial API distribution as of 2026. (https://aiweekly.co/learning-ai/generative-ai/how-to-use-runway)
- (single secondary coverage) [SECONDARY] CONTRADICTION — Pika free-tier rights conflict across sources: watermarked non-commercial versus claimed commercial/no-watermark.
- (single secondary coverage) [SECONDARY] Hailuo 2.3 was an October 28, 2025 release and reads as a legacy product by August 2026 — two model generations behind the then-current Hailuo releases.
- (single directional coverage) [DIRECTIONAL] Vidu Q3 plan and credit allowances conflict across catalog sources; Luma's legacy Dream Machine tiers must be distinguished from the current Luma App / Ray 3.2 pricing.

### Image generation — Midjourney V8.x

