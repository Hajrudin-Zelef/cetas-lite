---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/part-7
title: "§18. Multimodal Generation: Video, Image, and Audio (part 7)"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Google", "OpenAI"]
dates: ["2026-03-19", "2026-04", "2026-05", "2026-06-03"]
keywords: ["benchmark", "chatgpt", "cost", "gpu", "grok", "latency", "license", "open-weight", "pricing", "reasoning", "voice"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8909, 8927]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: a31e472a254ebca7ffaf364226567980eccd7886dc7f0df796d952751ec13919
---

# §18. Multimodal Generation: Video, Image, and Audio (part 7)

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

