---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/overview
title: "§18. Multimodal Generation: Video, Image, and Audio"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Alibaba", "Apple", "ByteDance", "EU", "Google", "Groq", "Meta", "MiniMax", "Nvidia", "OpenAI", "Stability AI", "xAI"]
dates: ["2025-03-20", "2025-03-26", "2025-05-20", "2025-07", "2025-07-28", "2025-08-27", "2025-09", "2025-09-19", "2025-09-30", "2025-10", "2025-10-14", "2025-11-25", "2025-12", "2025-12-03", "2025-12-31", "2026-01", "2026-01-15", "2026-01-26", "2026-02-17", "2026-03", "2026-03-17", "2026-03-25", "2026-03-26", "2026-03-30", "2026-04", "2026-04-14", "2026-04-15", "2026-04-21", "2026-04-23", "2026-04-26", "2026-04-30", "2026-05", "2026-05-20", "2026-05-27", "2026-06", "2026-06-03", "2026-06-10", "2026-07", "2026-07-03", "2026-07-09", "2026-07-21", "2026-07-24", "2026-07-30", "2026-08-26", "2026-08-31", "2026-09", "2026-09-15", "2026-09-22", "2026-09-24"]
keywords: ["multimodal", "agent", "apache", "chatgpt", "consumer", "cost", "full-duplex", "gemini", "gemini 3.8", "gpu", "grok", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8749, 8819]
section: "§18. Multimodal Generation: Video, Image, and Audio"
sha256: 201e164460164000e8a678f1e6df3ce30662427f188e38bcd7c22c7fbcd83753
---

# §18. Multimodal Generation: Video, Image, and Audio

Keywords: multimodal generation, video generation, image generation, text-to-speech, Sora 2, Runway Gen-4.5, ElevenLabs, Suno v5.5, FLUX.2, Ideogram 4, Midjourney V8, Luma Ray 3, Cartesia Sonic-3.6, Dia, TTS, voice cloning, music generation, image editing, video editing, audio tags

## Summary
- **Sora 2 launched 2025-09-30** — the brief's 2026 date is corrected to 2025 [SECONDARY]; the consumer app/web reportedly ended 2026-04-26 and the API end was reported for 2026-09-24, both single-source-family and needing corroboration [SECONDARY].
- **Runway has no Gen-5 as of 2026-09-22**: Gen-4.5 (December 2025) remains the flagship, with Aleph as the in-context video-editing line [SECONDARY].
- **ElevenLabs v4 was preview-only** (ElevenSummit Warsaw, June 2026); Eleven v3 remains the flagship speech model — no GA, model ID, or endpoint as of September 2026 [SECONDARY].
- **PlayAI/PlayHT is dead**: Meta's July 2025 acqui-hire preceded the December 31, 2025 shutdown; Groq retired and repointed the engine [SECONDARY].
- **Ideogram 4** (June 3, 2026) is a 9.3B typography-layout image model whose code/pipeline is Apache 2.0 but whose weights sit under a non-commercial Ideogram agreement — the "Apache 2.0" shorthand for the weights is contradicted [SECONDARY].
- **Midjourney's V8.1 alpha date is unresolved** (2026-04-14 vs 2026-04-30); both dates are retained pending resolution [SECONDARY]; V8.2 became the default 2026-07-24 [SECONDARY].
- **Hunyuan3D 3.0 means free hosted access, not open weights** — the open checkpoints concentrate in the 2.x/Omni lines; the Hy3 Preview license split is Tencent Hy Community License for the preview vs Apache 2.0 for full Hy3/Hy4 Preview [SECONDARY].
- Image generation consolidates toward editing-capable models (FLUX.2 [dev] reference editing, GPT Image 2 editing endpoints, Photon multi-image referencing), while audio consolidates on licensed-data business models (Suno, Udio, ElevenLabs Music v2) [DIRECTIONAL].

## Key dated facts
### Image generation
- **2025-03-26** — Ideogram 3.0 launched with Ideogram Canvas (closed, paid tiers) [SECONDARY].
- **2025-05-20** — Google Imagen 4 announced at I/O: standard $0.04/image, Ultra $0.06, Fast $0.02; 2K output; SynthID watermarking [SECONDARY]; an Imagen 4 retirement date remains [UNVERIFIED] — no precise deprecation found.
- **2025-11-25** — FLUX.2 [dev] released: 32B text-to-image with single- and multi-reference editing; BFL license terms [VENDOR]; the **FLUX.2 [klein]** family (**2026-01-15**): 4B and 9B four-step distilled models; klein 4B Apache 2.0, 9B different terms, dev downloadable open-weight but commercial use requires a BFL license [VENDOR].
- **2026-02-17** — Recraft V4 reportedly launched: four variants (raster 1024², Pro 2048², Vector, Vector Pro); editable SVG output as the differentiator; V4.1 reportedly May 2026 [SECONDARY]; Elo 1172 / 72% as a single data point [DIRECTIONAL].
- **2026-03-17** — Midjourney V8 alpha [SECONDARY]; **V8.1 alpha April 2026** — date discrepancy **2026-04-14 vs 2026-04-30**, both retained [SECONDARY]; V8.1 default 2026-06-10/11; V8.2 default 2026-07-24; GPU-native rewrite, ~5× faster, native 2K via `--hd`, improved text; no public first-party REST API [SECONDARY].
- **2026-04-21** — OpenAI GPT Image 2 released (snapshot gpt-image-2-2026-04-21): generation + editing endpoints; $5/$8/$30 per 1M tokens (text-in/image-in/image-out); 1024² images at $0.006/$0.053/$0.211 [SECONDARY]; the GPT Image 2.5 "Sunburst" claim is [UNVERIFIED] single weak source.
- **2026-06-03** — Ideogram 4 released: 9.3B single-stream DiT, Qwen3-VL-8B text encoder, native 2K, structured JSON prompts, bounding boxes, color control; ~0.97 English OCR, 47.9% typography preference [DIRECTIONAL]; code/pipeline Apache 2.0 but weights under a non-commercial Ideogram agreement [SECONDARY].
- Luma Photon / Photon Flash: 1080p images $0.015 ($0.002 Flash), character consistency, multi-image referencing; no 2026 version bump [SECONDARY]; text rendering weak per review coverage [DIRECTIONAL].

### Video generation
- **2025-07-28** — Wan 2.2 released (Apache 2.0 open-weight): T2V-A14B/I2V-A14B MoE, TI2V-5B; up to 5s 720p [SECONDARY]; Wan 2.2 Animate (open weights/code) 2025-09-19 [SECONDARY]; Wan 2.5/2.6 remain closed/API-first [UNVERIFIED].
- **2025-08-27** — PixVerse V5: 360p–1080p, 5/8s, start/end frames, effects, Turbo/HD tiers [SECONDARY]; V6 (2026-03-30) single weak source [UNVERIFIED]; V5 modality coverage internally contradicted (image-only vs T2V/I2V) [SECONDARY].
- **2025-09-30** — **Sora 2 launched** (not 2026 — corrected from the brief): synchronized audio/dialogue, improved physics, Cameos; iOS/social app + Pro for ChatGPT Pro [SECONDARY].
- **2025-12-03** — HunyuanVideo released: 13B+ open code/weights under the Tencent Hy Community License (bounded below 100M monthly users, outside EU/UK/SK); no 2025–2026 descendant found [SECONDARY].
- **2026-01-26** — Luma Ray3.14: native 1080p, 4× faster/3× cheaper at 720p; dropped Character Reference and HDR/EXR; "reasoning" video model line [SECONDARY]; by September 2026 Ray 3.2 is current (5/10s T2V/I2V, 1080p, Modify up to 20s, 16 keyframes, 16-bit EXR/ACES, no generated audio) [SECONDARY].
- **2026-04-26** — Sora consumer app/web reportedly ended [SECONDARY] — single-source-family, needs corroboration; API end reported for 2026-09-24 [SECONDARY] — single-source-family.
- **2026-07-03** — Vidu S1 unveiled: real-time voice-driven interactive video from a single image, 540p/25fps, peak 42fps [SECONDARY] — syndicated coverage flagged, needs an independent source.
- Pika 2.5 flagship: 480p/720p/1080p, 5 or 10s clips; Pikaframes 2–5 keyframes, 20–25s; 2026-08-31 pricing: free 80 credits/mo, paid from $10/mo, a 5s clip at 12/20/40 credits by tier [SECONDARY].
- Kling 3.0/O3: 5–15s clips, multi-shot, native 4K reported [SECONDARY]; Kling 2.6→3.0/O3 line; **the Kling pricing block is contradicted across the corpus — do not cite any Kling price** [DIRECTIONAL].
- Runway Gen-4/Gen-4 Turbo + Aleph in-context video editing; API credits at 5/12/15 per second by model [SECONDARY]; **no Gen-5 as of 2026-09-22** — Gen-4.5 (December 2025) is the flagship [SECONDARY]; Aleph v1 sunset 2026-07-30, Aleph 2.0 replacement [UNVERIFIED]; Runway pivoting toward world models, $5.3B valuation per May 2026 coverage [SECONDARY].
- Hailuo 02: September 2025 integration coverage already named MiniMax-Hailuo-02; by July 2026 Hailuo 2.3 is the flagship (NCR architecture, native 1080p, T2V/I2V, 10s, no native audio) [SECONDARY].
- Pointer: MiniMax H3/Hailuo 3 video detail lives in §7; ByteDance Seedance 2.5 in §17; Google Gemini Omni/Veo in §14; Qwen Image 3.0 in wave6/04 (July 21, 2026, closed/API-only) [DIRECTIONAL].

### Audio and speech
- **2025-03-20** — OpenAI STT refresh (gpt-4o-transcribe $0.006/min, mini $0.003/min); TTS unchanged through 2026: gpt-4o-mini-tts ($0.015/min), tts-1 ($15/1M), tts-1-hd ($30/1M), 13 voices, no cloning, 4096-char cap [SECONDARY].
- **2026-03-25** — Google Lyria 3 Pro: 3-minute songs, SynthID watermarking [SECONDARY].
- **2026-03-26** — Suno v5.5: Voices (Pro/Premier, identity verification), Custom Models (6+ tracks, up to 3/account), My Taste; 12-stem WAV+MIDI export, Song Editor inpainting, 44.1kHz stereo; no v6 three weeks later; free tier v4.5-only; WMG settled only one major; ~100M users, $2.45B valuation [DIRECTIONAL].
- **2026-04-15** — Gemini 3.1 Flash TTS: audio tags, $1/1M in, $20/1M out, 80+ locales [SECONDARY].
- **2026-05-27** — ElevenLabs Music v2 launched: licensed training data, genre-switching mid-track, section inpainting, $0.15/min, 5-minute max, 44.1kHz, commercial from Starter+ [SECONDARY].
- **2026-06** — ElevenLabs v4 **previewed only** at ElevenSummit Warsaw ("performance acting"): no model ID, no endpoint, no release date [SECONDARY]; flagship remains Eleven v3 (audio tags, emotion/intent/accent) [SECONDARY]; deprecated *_v1 and scribe_v1 removed 2026-07-09; Flash v2/v2.5 low-latency ~75ms TTFA [SECONDARY]; Dubbing V2 mid-2026 [SECONDARY].
- **2026-08** — Cartesia Sonic-3.6: leads both AA speech arenas; SSM architecture; sub-90ms TTS, ~40ms TTFA; 42 languages; instant voice cloning from ~10s; $49/1M chars normalized (vs ElevenLabs $100) [SECONDARY]; Line voice agent ~$0.06/min [SECONDARY].
- **2026-08-26** — Gemini 3.5 Transcribe: WER 4.0% streaming / 2.6% non-streaming, 85 languages, 3 speakers [SECONDARY].
- **2026-09-15** — Gemini 3.8 Live / Extended Thinking: speech-to-speech with reasons+tools while talking; Extended Thinking #1 AA speech-to-speech (82.6, τ-Voice 68.6%), $0.005/$0.018 per min in/out [SECONDARY].
- Udio licensing: UMG settled with Udio October 2025 (new licensed service planned, partner-only as of May 2026); WMG deal followed; Udio still v1.5/Allegro v1.5 with no successor as of September 2026 [SECONDARY].
- **2025-12-31** — PlayAI/PlayHT shut down after Meta's July 2025 acqui-hire; Groq retired the engine and repointed playai-turbo at Canopy Labs' Orpheus [SECONDARY].
- Open-voice stack: Sesame CSM-1B (Apache 2.0, ~1M hours English training, near-human on short blind snippets); CSM-3B/8B commercial, no open timeline as of April 2026 [SECONDARY]; Nari Labs Dia 1.6B (Apache 2.0 open-weight, multi-speaker dialogue, nonverbals, zero-shot cloning, English) and Dia2 streaming (1B/2B checkpoints, HF Transformers) [SECONDARY]; Kyutai Moshi (open-source full-duplex ~200ms, Mimi codec); NVIDIA PersonaPlex 7B (June 2026, built on Moshi, ~0.24s interruption latency) [SECONDARY].
- Hume octave-2: 50% cheaper than v1, emotion-aware, natural-language description instructions, 11 languages, ~100–200ms, $7.60/1M chars PAYG; EVI + Expression Measurement [SECONDARY]; Google DeepMind licensing deal [DIRECTIONAL single-family].
- Adjacent music: Stable Audio 3.0 (May 20, 2026; four sizes, three open-weight, >6-minute tracks, licensed content); MiniMax music-2.6 GA [COMMUNITY]; Tencent SongGeneration (March 2026, open-weight) [COMMUNITY].
- Open-weight TTS leaderboard: Kokoro-82M (#1 TTS Arena, Apache 2.0), Chatterbox (MIT), Qwen3-TTS (Apache 2.0) [SECONDARY]; MiMo-V2.5-TTS-Series covered in wave6/04 (April 23, 2026, three API-only variants, limited-time free) [DIRECTIONAL].


### New verified facts — expansion

### Video generation — Google Veo 3.1

- [SECONDARY] Google Veo 3.1 was released in October 2025 (announced October 14, 2025) as the upgrade to Veo 3, adding native audio generation (dialogue, sound effects, ambient soundscapes), enhanced realism and stronger prompt adherence over Veo 3. (https://max-productive.ai/blog/google-veo-3-1-release/ ; https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator) — VENDOR-independent corroboration via Google's official blog cited by both.
- [SECONDARY] Veo 3.1 generates 4–8 second clips (per-clip maximum 8 seconds) at 720p or 1080p, with native synchronized audio; scene extension in Google Flow chains clips beyond 60 seconds. (https://apatero.com/blog/google-veo-31-complete-guide-ai-video-audio-2025 ; https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator)
- [SECONDARY] Google's 2026 lineup is a three-tier product: Veo 3.1 Lite (high-volume, cost-sensitive), Veo 3.1 Fast (speed/quality balance) and Veo 3.1 Standard (premium cinematic quality); model IDs include `veo-3.1-lite-generate-preview`, `veo-3.1-fast-generate-preview` and `veo-3.1-generate-preview`. (https://www.cometapi.com/what-is-google-veo-3-1-lite/ ; https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator)
- [SECONDARY] Lite supports 720p/1080p only (no 4K), single image-to-video, no video extension or first/last-frame control; Fast and Standard add up to 3 reference images, first/last frames, and video extension (up to ~20x / ~148s total on Fast). (https://www.cometapi.com/what-is-google-veo-3-1-lite/)
- [SECONDARY] Access surface: Gemini API, Vertex AI, Gemini app, Google Flow, YouTube Shorts and Google Vids; outputs carry SynthID watermarking. (https://blog.buildfastwithai.com/google-veo-3-1-ai-video-generator ; https://apatero.com/blog/google-veo-31-complete-guide-ai-video-audio-2025)
- (single secondary coverage) [SECONDARY] In April 2026 Google opened a permanent free tier for personal Google accounts: 10 video generations per month, 720p, up to 8 seconds per clip, text-to-video and photo-to-video inside Google Vids, one-click YouTube publish — no credit card, no subscription. (https://www.veo3ai.io/blog/google-veo-3-1-vs-seedance-comparison-2026)
- (single secondary coverage) [SECONDARY] CONTRADICTION — An earlier secondary report (buildfastwithai, ~April 2026) said Veo 3.1 "does not have a meaningful free tier" and only Gemini Advanced users got some credits; the veo3ai comparison (April 2026) reports a genuinely free 10-generations/month tier opened in April 2026. These conflict on free access — the later tier launch date explains the difference; both are preserved rather than averaged.
- (single secondary coverage) [SECONDARY] xAI's Grok Imagine Video launched January 2026: 2–15 second clips, native audio included, ~30-second generation time, multiple aspect ratios (16:9, 9:16, 4:3, 3:4, 2:3, 3:2, 1:1), at a reported $0.20 pricing. (https://www.vidguru.ai/blog/veo-3-1-vs-grok-imagine-video-comparison.html)
- (single community coverage) [COMMUNITY] May 2026 community comparison status: Sora app shut down; Sora API kept until September 2026; no free Sora access since January 2026; Veo active with free tier; Runway active and fully supported. (https://github.com/kacky000/aitoolpick/blob/HEAD/src/content/blog/sora-vs-veo-vs-runway-2026.md)

