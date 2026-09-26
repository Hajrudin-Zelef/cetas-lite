---
id: ai-industry-kb-2026-wave6/18-multimodal-generation-video-image-and-audio/voice-synthesis-elevenlabs-sesame-dia-moshivis
title: "Voice synthesis — ElevenLabs, Sesame, Dia, MoshiVis"
domain: multimodal-generation-video-image-and-audio
role: deep-dive
task: multimodal
actors: ["AWS"]
dates: ["2026-08", "2026-08-18", "2026-08-27", "2026-09"]
keywords: ["voice", "agents", "aws", "benchmarks", "latency", "parameters", "research", "sovereignty"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8895, 8908]
section: "§18. Multimodal Generation: Video, Image, and Audio"
delta_of: ai-industry-kb-2026
sha256: 528285315607012d0f32dc8c4fa687cb952accd5bff170d7a0848f6cd381f665
---

# Voice synthesis — ElevenLabs, Sesame, Dia, MoshiVis

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

