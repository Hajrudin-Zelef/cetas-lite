---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/figures-and-metrics
title: "Figures and metrics"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Apple", "Google", "OpenAI"]
dates: ["2025-11-25", "2026-01-15", "2026-02-17", "2026-04", "2026-04-21", "2026-05-03", "2026-06-03", "2026-08-25"]
keywords: ["apache", "benchmark", "chatgpt", "claude", "cost", "full-duplex", "gemini", "gemini 3.8", "latency", "license", "open-weight", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8928, 8993]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: d716fa2be49e90925d55e8453cb810b1a07882e34f20c66ed25b1d65c28fc918
---

# Figures and metrics

- (single secondary coverage) [SECONDARY] OpenAI's Sora 2 strengths (per AdCreate's 2026 review): remarkably natural lighting and physics, a Disney partnership with a character-cameo feature, and mobile iOS/Android apps — but the $20/month ChatGPT Plus plan limited users to 5-second 720p watermarked clips, with genuine production use requiring the $200/month Pro plan. (https://adcreate.com/blog/runway-gen-4-5-review-features-pricing-2026)
- (single secondary coverage) [SECONDARY] Midjourney's Stealth mode — required for confidential client projects — is gated on the Pro plan ($60/month) and above; Standard users have unlimited Relax but no private mode. (https://aicomparison.ai/midjourney-vs-dalle-3/)
- (single secondary coverage) [SECONDARY] Third-party route to GPT Image 2 without an OpenAI key: TwoShot offers free GPT Image 2 credits on signup (no credit card, no OpenAI API key) through its creative-assistant interface, with side-by-side comparison against Midjourney. (https://twoshot.app/gpt-image-2/)
- (single secondary coverage) [SECONDARY] Cartesia's Sonic-3 achieved a ~40ms time-to-first-audio benchmark across streaming WebSocket connections, with audio chunks streaming nearly instantaneously as upstream LLM tokens are generated — the profile that made it the default for high-frequency voice bots and gaming NPCs. (https://onepin.ai/blog/cartesia-vs-openai-tts-2026)
- (single secondary coverage) [SECONDARY] OpenAI's Images API pricing ladder effect: at 1024×1024, gpt-image-2 costs $0.006 (low) vs $0.211 (high) — a ~35x spread between draft and delivery quality that lets teams route exploration to cheap tiers and only pay high for final text-fidelity work. (https://github.com/shaharsha/claude-skills/blob/HEAD/skills/image-generation/reference/pricing.md)
- (single community coverage) [COMMUNITY] The genblaze connector audit (April 2026) documented that gpt-image-2 edits accept masks natively while the Responses API `"type": "image_generation"` tool path was deliberately deferred as out of scope — integration detail useful for developers wiring the model. (https://github.com/backblaze-labs/genblaze/blob/HEAD/docs/exec-plans/completed/openai-image-model-expansion.md)

- (single community coverage) [COMMUNITY] Runway shipped a Gen-4 update on May 3, 2026 adding native audio (lip sync and environmental SFX), social-media templates (vertical, subtitle-ready), and new API hooks for hybrid pipelines combining Runway with Veo and Seedance. (https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_video_generation_ai_models.md)

## Figures and metrics
| Image model | Release | Params | License (verified split) |
|---|---|---|---|
| FLUX.2 [dev] | 2025-11-25 | 32B | open-weight download; commercial needs BFL license [VENDOR] |
| FLUX.2 [klein] 4B | 2026-01-15 | 4B | Apache 2.0 [VENDOR] |
| Ideogram 4 | 2026-06-03 | 9.3B | code/pipeline Apache 2.0; weights non-commercial agreement [SECONDARY] |
| Midjourney V8.1 | 2026-04 (14 or 30) | — | proprietary, no public API [SECONDARY] |
| GPT Image 2 | 2026-04-21 | — | API-only; $5/$8/$30 per 1M tok [SECONDARY] |
| Recraft V4 | 2026-02-17 | — | API; editable SVG differentiator [SECONDARY] |

| Speech/TTS | Status | Latency / cost |
|---|---|---|
| ElevenLabs v3 | flagship (v4 preview-only) | Flash v2/v2.5 ~75ms TTFA [SECONDARY] |
| Cartesia Sonic-3.6 | leads AA speech arenas | ~40ms TTFA; $49/1M chars [SECONDARY] |
| Dia 1.6B | open-weight TTS | Apache 2.0 [SECONDARY] |
| Sesame CSM-1B | open voice model | Apache 2.0 [SECONDARY] |
| Kyutai Moshi | open full-duplex | ~200ms [SECONDARY] |

- GPT Image 2 per-image cost at 1024²: $0.006 / $0.053 / $0.211 depending on quality tier; Batch API 50% off [SECONDARY].
- Luma Photon Flash 1080p image: $0.002 (vs $0.015 standard Photon) [SECONDARY].
- Gemini 3.8 Extended Thinking speech: $0.005/$0.018 per minute in/out [SECONDARY].
- Hume octave-2: $7.60/1M chars PAYG; ~100–200ms [SECONDARY].


### New verified metrics — expansion

### Veo 3.1 — pricing tiers (USD per second, with audio) [SECONDARY]

| Tier | 720p | 1080p (single source) | 4K | Reference images | Extension | Label |
| --- | --- | --- (single source) | --- | --- | --- | --- |
| Lite | $0.05 | $0.08 (single source) | — | single I2V | no | [SECONDARY] |
| Fast | $0.15 ($0.10 after Apr 7) | $0.15 ($0.12 after Apr 7) (single source) | yes | up to 3 | yes, ~20x | [SECONDARY] |
| Standard | $0.40 | $0.40 (single source) | $0.60 | up to 3 | yes | [SECONDARY] |

Source: https://www.cometapi.com/what-is-google-veo-3-1-lite/ ; https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator (Standard $0.40/s 1080p with audio, Fast $0.10/s 720p).

- [SECONDARY] A 5-second 1080p clip with audio costs ≈ $2.00 at Standard rates; 8-second clips ≈ $3.20 (Standard) or ≈ $0.80–$1.20 (Fast/Lite). (https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator ; https://max-productive.ai/blog/google-veo-3-1-release/)
- [SECONDARY] Older launch-era reporting priced Veo 3.1 at $0.75/second including audio (apatero, ~late 2025/early 2026); the 2026 lineup cut effective rates via Lite/Fast — pricing history, not contradiction. (https://apatero.com/blog/google-veo-31-complete-guide-ai-video-audio-2025)
- [SECONDARY] Google AI Studio / Gemini Advanced access runs ~$20/month with generation limits; Vertex AI enterprise users get volume discounts. (https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator ; https://max-productive.ai/blog/google-veo-3-1-release/)
- [SECONDARY] Free tier (April 2026): 10 generations/month per Google account, 720p, 8s clips, via Google Vids. (https://www.veo3ai.io/blog/google-veo-3-1-vs-seedance-comparison-2026)

### Runway — credit pricing [SECONDARY]

| Plan | Annual $/mo | Monthly $/mo (single source) | Credits/month | Highlights | Label |
| --- | --- | --- (single source) | --- | --- | --- | --- |
| Free | $0 | $0 (single source) | 125 (one-time) | Gen-4 image-to-video, 720p export, 3 projects | [SECONDARY] |
| Standard | $12 | $15 (single source) | 625 | Gen-4.5, 4K export, Kling 3.0 Pro + Veo 3.1 access | [SECONDARY] |
| Pro | $28 | $35 (single source) | 2,250 | ProRes export, custom voices, 500GB storage | [SECONDARY] |
| Unlimited | $76 | $95 (single source) | 2,250 + unlimited Explore | unlimited Explore-mode generations | [SECONDARY] |
| Enterprise | custom | custom (single source) | custom | SSO, security, success program | [SECONDARY] |

Sources: https://allaiwebsite.com/tool/runway-gen-4/ ; https://aiweekly.co/learning-ai/generative-ai/how-to-use-runway
- [SECONDARY] 625 credits = 25s of Gen-4.5, 52s of Gen-4, or 125s of Gen-4 Turbo; annual billing saves 20% vs monthly. (https://allaiwebsite.com/tool/runway-gen-4/)
- [SECONDARY] Gen-4.5: 12 credits/second (AI Weekly, verified 2026-08-25 against official docs) vs 25 credits/second (AdCreate review) — preserved contradiction. (https://aiweekly.co/learning-ai/generative-ai/how-to-use-runway ; https://adcreate.com/blog/runway-gen-4-5-review-features-pricing-2026)

### Midjourney V8.x — subscription metrics [SECONDARY]

