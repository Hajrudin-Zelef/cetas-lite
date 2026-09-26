---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/image-generation-flux-2-black-forest-labs
title: "Image generation — FLUX.2 (Black Forest Labs)"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Alibaba", "Apple", "Google", "Together AI"]
dates: ["2026-03", "2026-03-26", "2026-03-27", "2026-05", "2026-06-03", "2026-07"]
keywords: ["agents", "diffusion", "fp8", "license", "omni", "open weights", "open-weight", "pricing", "research", "settlement", "text-to-image", "voice"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8865, 8894]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: 877afd7494190d37d9b379646c738112c80a6095894db8b84d307a273d314363
---

# Image generation — FLUX.2 (Black Forest Labs)

- [SECONDARY] Ideogram 4.0 was released June 3, 2026 as an open-weight model: a 9.3B single-stream diffusion transformer with native 256–2048px output and improved text rendering. (https://7minai.com/news/ideogram-4-open-weight/ ; https://the-decoder.com/ideogram-4-0-drops-as-an-open-weight-model-with-native-2k-resolution-and-improved-text-rendering/)
- [SECONDARY] It uses a Qwen3-VL-8B-Instruct text/vision encoder for prompt understanding; the HF checkpoint is gated by license agreement (research/preview), with nf4 (quantized) and fp8 variants published behind that gate. (http://toknow.ai/posts/ideogram-4-open-weight-image-model-typography-layout/index.pdf ; https://7minai.com/news/ideogram-4-open-weight/)
- [SECONDARY] New controls: structured JSON prompts, bounding-box layout control, and transparent-background generation. (https://7minai.com/news/ideogram-4-open-weight/ ; https://the-decoder.com/ideogram-4-0-drops-as-an-open-weight-model-with-native-2k-resolution-and-improved-text-rendering/)
- [SECONDARY] The open weights are licensed for non-commercial use; commercial use requires a paid license from Ideogram. (https://7minai.com/news/ideogram-4-open-weight/ ; https://www.together.ai/models/ideogram-40)
- [SECONDARY] Hosted API pricing: $0.03/image Turbo, $0.06/image Default, $0.10/image Quality; Together AI lists $0.06 per image with up to 2MP support. (https://7minai.com/news/ideogram-4-open-weight/ ; https://www.together.ai/models/ideogram-40)
- [COMMUNITY] DesignArena rankings (LMArena text-to-image by human votes) reportedly put Ideogram 4.0 at #1 among open-weight models; its typography performance is described as best-in-class with text error rates around 11–12%, which is better than gpt-image-1.5's reported ~12.6%. (http://toknow.ai/posts/ideogram-4-open-weight-image-model-typography-layout/index.pdf ; https://earlyterms.com/term/ideogram-4-0)

### Image generation — FLUX.2 (Black Forest Labs)

- (single vendor coverage) [VENDOR] FLUX.2 ships in six variants: FLUX.2 Klein 4B, Klein 9B, Pro, Max, Flex and Dev. (https://docs.bfl.ai/quick_start/pricing —  note: docs, blog and the vendor's GitHub skills mirrors form one vendor evidence chain, not independent sources.)
- (single vendor coverage) [VENDOR] Official megapixel-based pricing starts at: Klein 4B from $0.014, Klein 9B from $0.015, Pro from $0.03 text-to-image / $0.045 editing, Max from $0.07, Flex from $0.05; Dev is local-only and non-commercial. (https://docs.bfl.ai/quick_start/pricing — )
- [VENDOR] Klein models accept up to 4 reference images; Pro, Max and Flex accept up to 8; FLUX.2 image editing works through the same endpoint family without a separate endpoint. (https://github.com/black-forest-labs/skills/blob/HEAD/skills/bfl-api/SKILL.md ; https://docs.bfl.ai/quick_start/pricing — )
- (single secondary coverage) [SECONDARY] CONTRADICTION — A community mirror of BFL pricing listed Flex at $0.06/$0.12 (and possibly other variants) while the current official BFL docs list Flex from $0.05 — pricing drift over time; both states preserved with the current docs taking precedence for present-day figures.
- (single vendor coverage) [VENDOR] BFL's own API best-practices guidance distinguishes model selection by use case (Klein for cheap/fast, Pro for photorealistic quality, Max for top fidelity), with usage through Replicate, fal.ai and BFL's direct API. (https://github.com/black-forest-labs/skills/blob/HEAD/skills/flux-image-best-practices/AGENTS.md ; https://github.com/black-forest-labs/skills/blob/HEAD/skills/flux-image-best-practices/rules/model-selection-guide.md — )

### Music generation — Suno v5.5

- [SECONDARY] Suno v5.5 was released March 26, 2026 as Suno's "most expressive model yet," shifting from raw fidelity to personalization. Three headline features: Voices, Custom Models and My Taste. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://theplanettools.ai/blog/suno-v5-5-voice-cloning-custom-models-launch)
- [SECONDARY] Voices: train the model on the user's own singing voice with live identity verification (random phrase recording or upload); private to the user; Pro/Premier only; 4 credits per creation during beta. Community testing clarifies it blends vocal signature into Suno's vocal engine at a configurable "Audio Influence" (community sweet spot 40–60%) rather than true voice cloning. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://theplanettools.ai/blog/suno-v5-5-voice-cloning-custom-models-launch)
- [SECONDARY] Custom Models: upload original tracks (6+ tracks, up to 30 minutes audio) to fine-tune a private personalized v5.5 variant — up to 3 models per Pro/Premier user. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://theplanettools.ai/blog/suno-v5-5-voice-cloning-custom-models-launch)
- [SECONDARY] My Taste: adaptive preference engine learning genre preferences over time, available to all users including the free tier. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://theplanettools.ai/blog/suno-v5-5-voice-cloning-custom-models-launch)
- [SECONDARY] v5.5 builds on v5's 44.1kHz audio, natural vocals (whispers, vibrato, breathiness) and a composition architecture maintaining coherence across full songs; max duration reaches ~8 minutes with verse/chorus/bridge handling; Suno Studio 1.2 is a built-in DAW-lite with Warp Markers, Remove FX and post-generation time-signature changes; Pro users get 12-stem separation for professional mixing. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://videoreddit.edu.vn/%F0%9F%8E%B8-suno-2026-the-v5-5-omni-voices-update/)
- [SECONDARY] Official pricing as of March 2026: Free ($0, 50 credits/day ≈ 10 songs, v4.5 only, no commercial rights); Pro ($8/month billed annually, 2,500 credits ≈ 500 songs, full v5.5 + Voices/Custom Models, 12-stem exports, commercial rights); Premier ($24/month billed annually, 10,000 credits ≈ 2,000 songs, Suno Studio, priority queue). (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://sonilo.com/blog/suno-v5-5-pricing-commercial-use)
- (single secondary coverage) [SECONDARY] Suno has no official public API; third-party wrappers exist (e.g. via CometAPI) but are not Suno-official. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/)
- (single secondary coverage) [SECONDARY] Google dropped Lyria 3 Pro the same week as v5.5 (March 27, 2026), framing AI music as a two-front race; Suno's Warner Music settlement was the only major-label settlement in place at that point. (https://theplanettools.ai/blog/suno-v5-5-voice-cloning-custom-models-launch)
- (single secondary coverage) [SECONDARY] CONTRADICTION — The Planet Tools (May 2026) states that as of its writing "no v6 has shipped" and v5.5 is still the flagship, while a separate later source claims a v6/v6-wild — both states are preserved; v5.5 status beyond May 2026 cannot be confirmed from two independent sources.
- (single community coverage) [COMMUNITY] July 2026 community docs record post-v5.5 feature shipping: mobile UI overhaul (May 14), iOS share-from-Notes/Voice Memos (June 4), stem separation overhaul (June 11), lyrics editor rebuild (July 9), iMessage keyboard (July 15), Duration Slider (July 20). (https://github.com/sanic732/sunoforge/blob/HEAD/files/DATA_SUNO_2026-07.md — )

### Voice synthesis — Cartesia Sonic-3.6

