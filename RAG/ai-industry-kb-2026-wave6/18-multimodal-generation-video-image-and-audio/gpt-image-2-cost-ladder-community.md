---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/gpt-image-2-cost-ladder-community
title: "GPT Image 2 — cost ladder [COMMUNITY]"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Together AI"]
dates: ["2026-03-19", "2026-05", "2026-08"]
keywords: ["cost", "agents", "benchmark", "claude", "diffusion", "fp8", "gemini", "gpu", "latency", "license", "omni", "open-weight"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8994, 9036]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: 01563a91970a2c7cd5f18a2605378348808c7d3c4688b2c403fa5949bbeae139
---

# GPT Image 2 — cost ladder [COMMUNITY]

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

