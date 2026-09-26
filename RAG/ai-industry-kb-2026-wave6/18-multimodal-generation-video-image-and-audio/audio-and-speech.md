---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/audio-and-speech
title: "Audio and speech"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["Alibaba", "Google", "Groq", "Meta", "MiniMax", "Nvidia", "OpenAI", "Stability AI", "xAI"]
dates: ["2025-03-20", "2025-07", "2025-10", "2025-10-14", "2025-12-31", "2026-01", "2026-03", "2026-03-25", "2026-03-26", "2026-04", "2026-04-15", "2026-04-23", "2026-05", "2026-05-20", "2026-05-27", "2026-06", "2026-07-09", "2026-08-26", "2026-09", "2026-09-15"]
keywords: ["agent", "apache", "cost", "full-duplex", "gemini", "gemini 3.8", "grok", "latency", "leaderboard", "nvidia", "open-weight", "pricing"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8788, 8821]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: c7eff808ab346822d286e7e12ac8ebd597e6ff43e91ff79418228b9cd4faa5f8
---

# Audio and speech

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

### Video generation — Runway Gen-4.5 and GWM-1

