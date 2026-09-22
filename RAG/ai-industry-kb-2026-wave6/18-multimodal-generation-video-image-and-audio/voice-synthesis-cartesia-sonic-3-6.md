---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/voice-synthesis-cartesia-sonic-3-6
title: "Voice synthesis — Cartesia Sonic-3.6"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["AWS", "Apple", "Google", "OpenAI", "Together AI"]
dates: ["2025-11-25", "2026-01-15", "2026-02-17", "2026-03-19", "2026-04", "2026-04-21", "2026-05", "2026-05-03", "2026-06-03", "2026-08", "2026-08-18", "2026-08-25", "2026-08-27", "2026-09"]
keywords: ["voice", "agents", "apache", "aws", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "diffusion", "fp8", "full-duplex"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8893, 9036]
section: "§18. Multimodal Generation: Video, Image, and Audio"
sha256: 07ea6f87576690dab9ebd3bbb5b3ff023b7ac3f0c3d8a976b18b9ddf3a4d167c
---

# Voice synthesis — Cartesia Sonic-3.6

### Voice synthesis — Cartesia Sonic-3.6

- [SECONDARY] Cartesia shipped Sonic-3.6 in August 2026 (MarkTechPost coverage dated 2026-08-18) as a streaming TTS model that then led both Artificial Analysis Speech Arenas. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/ ; https://slator.com/cartesia-launches-text-to-speech-model/)
- (single community coverage) [COMMUNITY] CONTRADICTION — Community docs dated August 27, 2026 suggest a slightly later public-availability date than the August 18 press coverage — release announcement versus availability; both dates preserved.
- (single vendor coverage) [VENDOR] Sonic runs on state space models (SSMs) rather than transformers; Cartesia claims sub-90ms model TTS latency and 100ms transcript latency for its Ink-2 speech-to-text model — vendor-stated model latency, not measured end-to-end round trips. (http://cartesia.ai/sonic — )
- [SECONDARY] Features: inline expression tags (e.g. `[laughter]` in the transcript), instant voice cloning from ~10 seconds of audio (up to 60 seconds on Sonic-3.6/newer for better accents), custom pronunciation dictionaries with IPA overrides, speed/volume/emotion parameters, native alphanumeric handling (order numbers, phone numbers, codes), and English/Hinglish demos. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/ ; https://slator.com/cartesia-launches-text-to-speech-model/)
- [SECONDARY] Sonic covers 44 languages (up from 42 at launch — Odia and Urdu added for Sonic-3.6 per Slator, a delta vs earlier 42-language coverage); Artificial Analysis normalizes Sonic 3.6 at $49.00 per 1M characters — half of ElevenLabs Eleven v3 ($100.00), above Speechify Simba 3.2 ($10.00). Cartesia bills 1 credit per character with plans from $5 (100K credits) to $299/month (8M credits); Scale at $299 includes ~10,667 TTS minutes and 15 concurrent requests; Line voice agents bill separately at $0.06/minute. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/ ; https://www.marktechpost.com/2026/09/21/best-voice-cloning-apis-in-2026-speaker-similarity-consent-checks-and-price-per-1m-characters/ ; https://slator.com/cartesia-launches-text-to-speech-model/)
- (single vendor coverage) [VENDOR] Sonic-3.6 can be deployed on-prem inside data centers (including air-gapped environments), in the customer's own VPC (AWS/GCP/Azure), or via OEM licensing — pitched at government, healthcare, financial services and data-sovereignty buyers. (http://cartesia.ai/sonic — )
- (single secondary coverage) [SECONDARY] In voice-cloning comparisons, Cartesia's Pro Voice Clones train on 30+ minutes and start on the $49 Startup plan; Fish Audio's comparable API is $15 per 1M UTF-8 bytes (CJK costs ~3x more per character). (https://www.marktechpost.com/2026/09/21/best-voice-cloning-apis-in-2026-speaker-similarity-consent-checks-and-price-per-1m-characters/)
- (single community coverage) [COMMUNITY] Real-world switchovers report Cartesia ~7x cheaper than ElevenLabs ($0.04/min vs $0.30/min) with ~150ms versus ~500ms latency, and direct pcm_mulaw 8kHz output avoiding transcoding for telephony (Twilio Media Streams). (https://github.com/wburns02/react-crm-api — )

### Voice synthesis — ElevenLabs, Sesame, Dia, MoshiVis

- (single unverified coverage) [UNVERIFIED] ElevenLabs v4 could not be verified as generally available from two independent sources as of September 2026; its status remains preview-only/unverified. No independent measurement confirms a v4 release — retain preview-only status.
- (single directional coverage) [DIRECTIONAL] Sesame CSM-1B, Dia and MoshiVis are the open/conversational voice models in the research notes for this section; their 2026 release details and benchmarks were not corroborated by two independent sources in this wave and require dedicated research before inclusion in the consolidation.

- [SECONDARY] V8 web interfaces were upgraded to support the new speed: an improved conversation mode ("talk in flow"), a Grid Mode for focusing on a single big set of images, and settings moved into sidebars so tweaks don't block the view; V8 supports multiple aspect ratios, --chaos, --weird, --exp, --raw, with backward compatibility for V7 personalization profiles, moodboards and srefs. (https://releasebot.io/updates/midjourney ; https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive)
- (single secondary coverage) [SECONDARY] Midjourney's 2026 price story: no radical change to the four-tier structure, but intensive use of --hd and --q 4 consumes GPU hours up to 4x faster, pushing professional users toward Pro/Mega — dayahimour's guide prices Basic at $10/month (3.5 GPU hours) and Mega at $120/month (60 hours). (https://github.com/dayahimour/ai-dayahimour/blob/HEAD/src/content/blog-en/midjourney-v8.md)
- (single secondary coverage) [SECONDARY] Users indicate V8 improves overall quality especially in commercial and cinematic contexts, but may lose some of the "artistic soul" that characterized V7 in painterly or fantasy styles — subjective quality regression preserved as user perception, not benchmark. (https://github.com/dayahimour/ai-dayahimour/blob/HEAD/src/content/blog-en/midjourney-v8.md)
- (single community coverage) [COMMUNITY] For developers on OpenAI's Images API, migrating gpt-image-1 → gpt-image-2 is mechanical: change `model="gpt-image-2"`, optionally add `quality="thinking"` for complex prompts, optionally request `n=8` for consistent series; editing/inpainting runs through the same endpoint family. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- (single community coverage) [COMMUNITY] Pre-release developer access to gpt-image-2 was available via third-party hosts (fal.ai, apiyi) before OpenAI's direct API GA in early May 2026. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- (single community coverage) [COMMUNITY] Per-token billing captures the "planning work" (prompt comprehension, reasoning steps, web-search results in Thinking mode) plus pixel output: a simple prompt costs less than a magazine cover with 5 cover lines and a hero photo. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- (single secondary coverage) [SECONDARY] Content policy for gpt-image-2 is the same as ChatGPT: no NSFW, no real persons, no copyrighted characters — a constraint on commercial workflows. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5)
- (single community coverage) [COMMUNITY] Free Suno users remain on v4.5-all while paid tiers unlock v5.5; during the Voices beta, costs were reduced to encourage testing; My Taste activates automatically based on user activity. (https://medium.com/@mkteam/suno-v5-5-what-is-new-and-how-to-use-it-via-api-studio-2356b3c2c3b0)
- [SECONDARY] Commercial-rights boundary on Suno: free = personal use only (no monetized YouTube, no client deliverables, no ads/branded content); paid plans unlock commercial rights for new songs — upgrading does not retroactively license old tracks. (https://sonilo.com/blog/suno-v5-5-pricing-commercial-use ; https://github.com/sanic732/sunoforge/blob/HEAD/files/DATA_SUNO_2026-07.md — community source)
- (single community coverage) [COMMUNITY] Suno credit arithmetic: a generation produces 2 tracks; Auto Split alone costs 50 credits — a full day of free-tier allowance — so stem work is the expensive operation, not generation. (https://github.com/sanic732/sunoforge/blob/HEAD/files/DATA_SUNO_2026-07.md)
- (single community coverage) [COMMUNITY] Cartesia third-party integration docs record a free tier of 20K credits and a Pro tier of $4/month for 100K credits (≈15–20 hours of voice-mode conversation), consistent with the vendor's credit-per-character billing at small scale. (https://github.com/mbailey/voicemode/blob/HEAD/docs/guides/cartesia-setup.md)
- (single secondary coverage) [SECONDARY] OpenAI's TTS comparison anchor: `tts-1` at $15/1M chars, `tts-1-hd` at $30/1M chars, GPT-4o Mini TTS with ~250–500ms latency — the ecosystem-native alternative that trades latency for simplicity. (https://onepin.ai/blog/cartesia-vs-openai-tts-2026)
- (single secondary coverage) [SECONDARY] Veo 3.1 Fast/Standard support "Ingredients to Video" reference-image workflows and native 4K upscaling (Vidguru integration noted as coming soon); it prioritizes photorealism and consistent physics over clip length or cost. (https://www.vidguru.ai/blog/veo-3-1-vs-grok-imagine-video-comparison.html)
- (single community coverage) [COMMUNITY] Benchmarked claims on Veo 3.1 vs rivals (kacky000 community comparison): Veo prompt adherence 87%, motion realism 92% — third-party benchmarked figures, not Google-vendor. (https://github.com/kacky000/aitoolpick/blob/HEAD/src/content/blog/sora-vs-veo-vs-runway-2026.md — community source)
- (single secondary coverage) [SECONDARY] Ideogram 4.0 community term page (earlyterms) corroborates the June 3, 2026 open-weight release and the non-commercial license gate — earlyterms is an independent aggregator, not Ideogram. (https://earlyterms.com/term/ideogram-4-0)
- [VENDOR] Ideogram's hosted plans (API/Pro tiers) gate the commercial license, and Turbo/Default/Quality image presets map to the $0.03/$0.06/$0.10 price points — hosted pricing is a vendor figure, not an independent measurement. (https://www.together.ai/models/ideogram-40 is independent reseller evidence; https://7minai.com/news/ideogram-4-open-weight/ is secondary reporting.)
- (single secondary coverage) [SECONDARY] The AI-music three-way comparison (March–April 2026) framed Lyria 3 Pro, Suno v5.5 and Udio as the current contest — with Udio's UMG-announced licensed platform slated for 2026 and fingerprinting/filtering measures expected. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/)
- (single secondary coverage) [SECONDARY] Runway's Standard plan includes third-party model access (Kling 3.0 Pro, Veo 3.1, Seedance 2.0) at no extra subscription — the marketplace model means credits, not subscriptions, are the unit of comparison across video vendors. (https://allaiwebsite.com/tool/runway-gen-4/)

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

- [SECONDARY] Basic $10/mo → 3.3 Fast GPU hours (~200 images); Standard $30/mo → 15 Fast GPU hours + unlimited Relax; Pro $60/mo → 30 Fast GPU hours + Stealth + 12 concurrent jobs; Mega $120/mo → 60 Fast GPU hours; annual billing ≈ 20% off. (https://github.com/linda-mhmd/ai-solutions-wiki/blob/HEAD/content/comparisons/midjourney-vs-dalle-vs-stable-diffusion.md ; https://aiwisepicks.com/tools/midjourney/)
- [SECONDARY] V8 premium jobs (--hd, --q 4, sref, moodboard) cost 4x standard at launch; V8 renders ~5x faster than V7 (30–60s → under 10s per image). (https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive ; https://releasebot.io/updates/midjourney)
- [SECONDARY] --hd mode renders natively at 2K resolution, eliminating the need for external upscaling in professional workflows. (https://aiforautomation.io/news/2026-03-19-midjourney-v8-5x-faster-4x-more-expensive)

### GPT Image 2 — cost ladder [COMMUNITY]

- [COMMUNITY] Per-token: $5/$10 per M text input/output, $8/$2 per M image input/cached input, $30 per M image output. (https://dev.to/tokenmixai/gpt-image-2-api-developer-guide-pricing-thinking-mode-and-production-integration-2026-28p5 ; https://blog.laozhang.ai/en/posts/gpt-image-2-api-release-date)
- [COMMUNITY] Per-image at 1024×1024: low $0.006, medium $0.053, high $0.211; 1024×1536 portrait: $0.005/$0.041/$0.165; scales roughly linearly above 1024². (https://github.com/shaharsha/claude-skills/blob/HEAD/skills/image-generation/reference/pricing.md)
- [COMMUNITY] Competitor anchor: Nano Banana Pro (`gemini-3-pro-image`) ≈ $0.134/image at ≈95% diagram-text accuracy versus gpt-image-2 ≈99%. (https://github.com/arthur-diego/linkedin-rag-content/blob/HEAD/docs/research/image-models-2026.md)

### Ideogram 4.0 — pricing [SECONDARY]

- [SECONDARY] Hosted: Turbo $0.03/image, Default $0.06/image, Quality $0.10/image; Together AI: $0.06/image, up to 2MP. (https://7minai.com/news/ideogram-4-open-weight/ ; https://www.together.ai/models/ideogram-40)
- [SECONDARY] Open-weight checkpoints: 9.3B parameters, single-stream DiT; HF gate requires license agreement (research/preview), nf4/fp8 variants behind the gate. (http://toknow.ai/posts/ideogram-4-open-weight-image-model-typography-layout/index.pdf ; https://7minai.com/news/ideogram-4-open-weight/)
- [SECONDARY] Native resolution 256–2048px (2K), no upscaling required for print-ready typography. (https://the-decoder.com/ideogram-4-0-drops-as-an-open-weight-model-with-native-2k-resolution-and-improved-text-rendering/)

### FLUX.2 — pricing ladder [VENDOR]

- [VENDOR] Per-megapixel starting prices: Klein 4B $0.014, Klein 9B $0.015, Pro $0.03 (T2I) / $0.045 (editing), Max $0.07, Flex $0.05; Dev local-only, non-commercial. (https://docs.bfl.ai/quick_start/pricing)
- [SECONDARY] CONTRADICTION — Community pricing mirror listed Flex at $0.06/$0.12 versus the current official from-$0.05 — drift preserved. (vendor mirror chain, community source)

### Suno v5.5 — credit/pricing metrics [SECONDARY]

- [SECONDARY] Free: $0, 50 credits/day (≈10 songs), v4.5 only, non-commercial; Pro: $8/mo annual (≈$10 monthly), 2,500 credits (≈500 songs), commercial rights; Premier: $24/mo annual (≈$30 monthly), 10,000 credits (≈2,000 songs). (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://sonilo.com/blog/suno-v5-5-pricing-commercial-use)
- [SECONDARY] Pro tier: 10 concurrent generations (priority queue), 30-minute uploads, 12 stems; Premier adds Suno Studio and fastest priority. (https://medium.com/@mkteam/suno-v5-5-what-is-new-and-how-to-use-it-via-api-studio-2356b3c2c3b0)
- [SECONDARY] Audio: 44.1kHz (from v5), max duration ~8 minutes; stems split into up to 12 individual tracks. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/ ; https://videoreddit.edu.vn/%F0%9F%8E%B8-suno-2026-the-v5-5-omni-voices-update/)

### Cartesia Sonic-3.6 — pricing metrics [SECONDARY]

- [SECONDARY] 1 credit per character of TTS output; plans: $5 (100K credits), $39–$49 tiers, up to $299/month (8M credits); Scale $299 ≈ 10,667 TTS minutes + 15 concurrent requests; Line voice agents $0.06/minute separate. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/ ; https://www.marktechpost.com/2026/09/21/best-voice-cloning-apis-in-2026-speaker-similarity-consent-checks-and-price-per-1m-characters/)
- [SECONDARY] Artificial Analysis normalized cost: Sonic 3.6 $49.00/1M chars vs ElevenLabs Eleven v3 $100.00/1M chars vs Speechify Simba 3.2 $10.00/1M chars. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/)
- [SECONDARY] Vendor-stated latency: sub-90ms TTS model latency; Ink-2 STT 100ms transcript latency. (http://cartesia.ai/sonic — vendor source)
- [SECONDARY] 44 languages; instant cloning from ~10s audio (up to 60s on Sonic-3.6/newer); Pro Voice Clones from $49 Startup plan on 30+ minutes of training. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/ ; https://www.marktechpost.com/2026/09/21/best-voice-cloning-apis-in-2026-speaker-similarity-consent-checks-and-price-per-1m-characters/)

### Benchmark signals (independent-ish) [SECONDARY]/[COMMUNITY]

- [COMMUNITY] Runway Gen-4.5: 1,247 Elo on Artificial Analysis T2V at late-2025 launch (then #1); displaced from top 10 by May 2026. (https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/best_video_generation_ai_models.md)
- [COMMUNITY] Gen-4.5 vs Sora 2 vs Veo 3.1 vs Kling 3.0 comparison Elo: 1,247 / 1,206 / 1,226 / n.a.; max duration 60s / 25s / 8s / 15s; max resolution 4K (Pro) / 1080p (Pro) / 1080p (Ultra) / 4K at 60fps. (https://adcreate.com/blog/runway-gen-4-5-review-features-pricing-2026)
- [COMMUNITY] gpt-image-2 debuted and remains #1 on LMArena/Artificial Analysis text-to-image and editing arenas (Aug 2026). (https://github.com/arthur-diego/linkedin-rag-content/blob/HEAD/docs/research/image-models-2026.md)
- [SECONDARY] Ideogram 4.0 reported #1 open-weight model on DesignArena (LMArena text-to-image, human votes); typography text error rates ~11–12% vs gpt-image-1.5's ~12.6%. (http://toknow.ai/posts/ideogram-4-open-weight-image-model-typography-layout/index.pdf)
- [SECONDARY] Sonic-3.6 led both Artificial Analysis Speech Arenas at its August 2026 launch. (https://www.marktechpost.com/2026/08/18/cartesia-ships-sonic-3-6-a-streaming-tts-model-that-now-leads-both-artificial-analysis-speech-arenas/)
- [SECONDARY] Suno v5.5 community ELO rankings and Reddit tests report clearer, more emotionally tuned outputs than v5. (https://www.cometapi.com/suno-v5-5-vs-lyria-3-pro-vs-udio-in-2026-which-ai-music-generator-is-best/)

