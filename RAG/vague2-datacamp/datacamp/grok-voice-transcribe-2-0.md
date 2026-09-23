---
id: vague2-datacamp/datacamp/grok-voice-transcribe-2-0
title: "Grok Voice Transcribe 2.0 : précision, fonctions, tarifs et accès API"
domain: datacamp
role: reference
task: article
actors: ["Meta", "SpaceX", "xAI"]
dates: ["2026-09", "2026-09-23"]
keywords: ["grok", "voice", "agent", "benchmark", "benchmarks", "cost", "latency", "muse", "pricing", "training", "transcription"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/grok-voice-transcribe-2-0.md
source_anchor: ""
source_lines: [1, 51]
sha256: 5207f3aee7cc1593c7c1bd6bada93d2b32182a52535a101bebe2e6e00fc4017e
---

# Grok Voice Transcribe 2.0 : précision, fonctions, tarifs et accès API

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/grok-voice-transcribe-2-0
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Grok Voice Transcribe 2.0, SpaceXAI/xAI's second-generation speech-to-text model, which replaces Grok Voice Transcribe 1.0 at the same hourly price. It tops Artificial Analysis's public streaming speech-to-text ranking for accuracy. The model is built on the audio base model behind Grok Voice, which xAI says already handles tens of thousands of customer support calls daily, transcribes millions of hours of video narration, and powers the Grok assistant in Tesla vehicles. The upgrade is in training, not the API: it was trained on live, noisy, multilingual audio and post-trained on hard real-world conditions—unstable phone lines, overlapping voices, local accents, and dictated phone numbers or emails. xAI claims 2.0 is twice as accurate as 1.0 on its real-world evals.

Features are unchanged from 1.0. Users can transcribe files or live audio with one model (batch endpoint for files/URLs up to 500 MB in 12 audio formats including WAV, MP3, FLAC, MP4 containers, raw PCM, and G.711 phone codecs; WebSocket streaming with partial transcripts every ~500 ms). Every word carries start/end timestamps, and diarization adds a speaker label at no extra cost; multichannel mode transcribes up to 8 independent channels. Responses are a single JSON object with full text, detected language (BCP-47), audio duration, and word array. Up to 100 key terms per request (each up to 50 characters) can bias the model toward domain vocabulary; written formatting of numbers, dates, currencies, phone numbers, and emails is supported in 25 languages when the language parameter is set. Filler words are removed by default. Smart turn detection lets a voice agent wait for the end of an idea with a configurable confidence threshold (0.5 balanced, 0.7 for dictated numbers); a companion timeout forces turn end after fixed silence, and default endpointing triggers after 400 ms of silence. The model auto-detects language and handles mid-recording language changes in one pass.

On benchmarks, Grok Voice Transcribe 2.0 achieves a 2.7% WER on Artificial Analysis's streaming index (lowest of 34 streaming models as of 21 September 2026), vs Muse Voice Transcribe (Meta) 3.1%, ElevenLabs Scribe v2 Realtime 3.6%, and Grok 1.0 3.9%. Its largest lead is on VoxPopuli (1.4% WER, less than half of 1.0's 3.1%). The gap vs 1.0 on the index is 31%, not the "2x" the announcement implies. The trade-off is latency: 0.49 s to final transcript vs 0.37 s for 1.0 and 0.14 s for Scribe v2 Realtime. On internal production evals (8 kHz support-center telephony, Grok conversations, dictated identifiers, short voice commands in 19 languages), it improves on all four; the only published figure is short phrases, where WER drops from 20.6% to 6.8%.

Pricing is $0.10/hour batch and $0.20/hour streaming, identical to 1.0, with diarization, word timestamps, key-term biasing, and formatting included. Artificial Analysis lists Grok 2.0 at $3.33 per 1,000 minutes vs $3.00 for Meta's Muse and $6.50 for Scribe v2 Realtime and Deepgram Flux. The model is generally available via the xAI API (no waitlist), with model ID `grok-voice-transcribe-2.0`; 2.0 is becoming the default and 1.0 will be deprecated in coming weeks, so pin the model ID explicitly.

## Key points

- Grok Voice Transcribe 2.0 is xAI's new speech-to-text model, replacing 1.0 at the same price.
- It ranks first in accuracy among streaming models on Artificial Analysis (2.7% WER of 34 models).
- Biggest gains are on hard audio: phone lines, dictated account codes/emails, and short multilingual commands.
- Trade-off is latency: 0.49 s to final transcript vs 0.37 s for 1.0 and 0.14 s for ElevenLabs Scribe v2.
- Existing integrations upgrade without code changes, but explicitly set the model ID until the default switches.
- Features: diarization, word timestamps, 100 key terms, 25-language formatting, 8-channel, smart turn detection.
- Pricing: $0.10/hour batch, $0.20/hour streaming; among the lowest in the streaming ranking.

## Technical data / figures

| Model | WER index (final transcript) | Delay to final transcript |
| --- | --- | --- |
| Grok Voice Transcribe 2.0 | 2.7% | 0.49 s |
| Muse Voice Transcribe (Meta) | 3.1% | 0.16 s |
| ElevenLabs Scribe v2 Realtime | 3.6% | 0.14 s |
| Grok Voice Transcribe 1.0 | 3.9% | 0.37 s |
| Deepgram Flux | 7.4% | 0.02 s |

- Batch price: $0.10/hour audio; Streaming price: $0.20/hour audio.
- Cost per 1,000 minutes: Grok 2.0 $3.33; Muse Voice Transcribe $3.00; Scribe v2 Realtime / Deepgram Flux $6.50.
- Model ID: `grok-voice-transcribe-2.0` (previous: `grok-voice-transcribe-1.0`).
- Endpoints: REST `https://api.x.ai/v1/stt` (batch); WebSocket `wss://api.x.ai/v1/stt` (streaming).
- Batch limits: files up to 500 MB, 12 audio formats; streaming partials every ~500 ms.
- Internal short-phrase WER: 20.6% (1.0) → 6.8% (2.0); formatting in 25 languages.

## Why this source matters for the RAG

It provides current, detailed data on a leading streaming speech-to-text model, including WER/latency trade-offs, pricing, and API access—useful for voice-AI and transcription pipeline questions. It also documents the benchmark-versus-vendor-claim nuance, reinforcing the RAG's pattern of critically assessing performance figures.
