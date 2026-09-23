---
id: collect-huggingface/huggingface/coherelabs-cohere-transcribe-03-2026
title: "cohere-transcribe-03-2026 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Apple", "China", "Cohere", "Hugging Face", "Nvidia", "OpenAI", "vLLM"]
dates: ["2026-03-26", "2026-09-23"]
keywords: ["cohere", "apache", "benchmarks", "inference", "leaderboard", "license", "nvidia", "open-weight", "qwen", "vllm", "voice"]
source: docs/RAG/Collect RAG/03_huggingface/CohereLabs-cohere-transcribe-03-2026.md
source_anchor: ""
source_lines: [1, 50]
sha256: 483f7920743953b2fbf84aa5b1676b9bbf980574c898d628d1927feec4f27e2b
---

# cohere-transcribe-03-2026 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/CohereLabs/cohere-transcribe-03-2026
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Cohere Transcribe is an open-source release from Cohere and Cohere Labs of a 2B-parameter dedicated audio-in, text-out automatic speech recognition (ASR) model supporting 14 languages. Architecture: a large Conformer-based encoder-decoder — a Conformer encoder extracts acoustic representations from log-Mel spectrograms (audio auto-resampled to 16 kHz; stereo averaged to mono), followed by a lightweight Transformer decoder for token generation. Trained from scratch with a supervised cross-entropy objective on output tokens. Languages: European (English, French, German, Italian, Spanish, Portuguese, Greek, Dutch, Polish), APAC (Chinese/Mandarin, Japanese, Korean, Vietnamese), MENA (Arabic). License: Apache 2.0. DOI 10.57967/hf/8653.

Results (English ASR leaderboard as of 2026-03-26): average WER 5.42 — best-in-class among compared models (Zoom Scribe v1 5.47, IBM Granite 4.0 1B Speech 5.52, NVIDIA Canary Qwen 2.5B 5.63, Qwen3-ASR-1.7B 5.76, ElevenLabs Scribe v2 5.83, Kyutai STT 2.6B 6.40, OpenAI Whisper Large v3 7.44, Voxtral Mini 4B 7.68). Per-dataset: AMI 8.15, Earnings 22 10.84, Gigaspeech 9.33, LibriSpeech clean 1.25 / other 2.37, SPGISpeech 3.08, Tedlium 2.49, Voxpopuli 5.87. RTFx 524.88. Human-preference evaluations confirm strong real-world quality. Per-language WERs reported over FLEURS, Common Voice 17.0, MLS, Wenet.

Usage: natively supported in transformers>=5.4.0 (AutoProcessor + CohereAsrForConditionalGeneration); long-form audio auto-chunked by the feature extractor (transcribing a 55-min call with high RTFx); punctuation control (default on, punctuation=False for lower-case no-punct); batched inference mixing short/long audio; language code for non-English. Production serving via vLLM 0.19.0 (vllm serve ... --trust-remote-code; /v1/audio/transcriptions endpoint). Ecosystem: transformers, vLLM, mlx-audio (Apple Silicon), Rust port, browser/WebGPU demo, Chrome extension, iOS/Android apps. Limitations: no automatic language detection (best in single language), no timestamps or speaker diarization, and benefits from a VAD/noise gate to avoid hallucinating on silence. Gated access (contact info required). Hub: 2B params BF16, 211,332 downloads/month, 39 quantizations, 19 finetunes.

## Key points

- 2B dedicated ASR model (Conformer encoder + lightweight Transformer decoder), trained from scratch.
- 14 languages across Europe, APAC, and MENA; audio auto-resampled to 16 kHz, stereo averaged.
- Best-in-class English WER: 5.42 average (beats Whisper Large v3 at 7.44).
- RTFx 524.88; real-time factor up to 3x faster than similar-size dedicated ASR models.
- Long-form audio auto-chunking; punctuation control; batched inference.
- Apache 2.0 license; natively supported in transformers>=5.4.0 and vLLM.
- No language auto-detection, no timestamps/diarization; VAD recommended.
- Hub: 2B params, 211,332 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | CohereLabs |
| Model name | cohere-transcribe-03-2026 |
| Architecture | Conformer encoder-decoder |
| Total params | 2B |
| Input | Audio waveform -> log-Mel spectrogram (16 kHz) |
| Output | Transcribed text |
| Languages | 14 (en, fr, de, it, es, pt, el, nl, pl, zh, ja, ko, vi, ar) |
| License | Apache 2.0 |
| Weight formats | BF16 |
| Benchmarks | Avg WER 5.42 (English ASR leaderboard); RTFx 524.88 |
| Requirements | transformers>=5.4.0; vLLM 0.19.0 for serving |
| Downloads/month | 211,332 |

## Why this source matters for the RAG

This card documents a state-of-the-art open-weight ASR model with full leaderboard numbers, language coverage, and practical long-form/batched serving guidance. It is essential reference material for retrieval on speech recognition, Conformer encoder-decoder designs, and production ASR deployment.
