---
id: collect-240926-huggingface/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face-2
title: "coherelabs-cohere-transcribe-03-2026-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Apple", "Cohere", "EU", "Hugging Face", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["cohere", "apache", "benchmarks", "distribution", "leaderboard", "license", "nvidia", "qwen", "research", "transcription", "vllm", "voice"]
source: docs/RAG/clean_en/huggingface/coherelabs-cohere-transcribe-03-2026-hugging-face.md
source_anchor: ""
source_lines: [164, 238]
sha256: d75440c2ad420a989203609b11dea4078be540682d8e96d3a931fe220854b522
---

# coherelabs-cohere-transcribe-03-2026-hugging-face

| Model | Average WER | AMI | Earnings 22 | Gigaspeech | LS clean | LS other | SPGISpeech | Tedlium | Voxpopuli | 
|---|---|---|---|---|---|---|---|---|---|
| **Cohere Transcribe** | **5.42** | **8.15** | 10.84 | 9.33 | **1.25** | **2.37** | 3.08 | 2.49 | 5.87 | 
| Zoom Scribe v1 | 5.47 | 10.03 | 9.53 | 9.61 | 1.63 | 2.81 | **1.59** | 3.22 | **5.37** | 
| IBM Granite 4.0 1B Speech | 5.52 | 8.44 | **8.48** | 10.14 | 1.42 | 2.85 | 3.89 | 3.10 | 5.84 | 
| NVIDIA Canary Qwen 2.5B | 5.63 | 10.19 | 10.45 | 9.43 | 1.61 | 3.10 | 1.90 | 2.71 | 5.66 | 
| Qwen3-ASR-1.7B | 5.76 | 10.56 | 10.25 | **8.74** | 1.63 | 3.40 | 2.84 | **2.28** | 6.35 | 
| ElevenLabs Scribe v2 | 5.83 | 11.86 | 9.43 | 9.11 | 1.54 | 2.83 | 2.68 | 2.37 | 6.80 | 
| Kyutai STT 2.6B | 6.40 | 12.17 | 10.99 | 9.81 | 1.70 | 4.32 | 2.03 | 3.35 | 6.79 | 
| OpenAI Whisper Large v3 | 7.44 | 15.95 | 11.29 | 10.02 | 2.01 | 3.91 | 2.94 | 3.86 | 9.54 | 
| Voxtral Mini 4B Realtime 2602 | 7.68 | 17.07 | 11.84 | 10.38 | 2.08 | 5.52 | 2.42 | 3.79 | 8.34 | 

Link to the live leaderboard: Open ASR Leaderboard.

We observe similarly strong performance in human evaluations, where trained annotators assess transcription quality across real-world audio for accuracy, coherence and usability. The consistency between automated metrics and human judgments suggests that the model’s improvements translate beyond controlled benchmarks to practical transcription settings.

*Figure: Human preference evaluation of model transcripts. In a head-to-head comparison, 
annotators were asked to express preferences for generations which primarily preserved meaning - 
but also avoided hallucination, correctly identified named entities, 
and provided verbatim transcripts with appropriate formatting. 
A score of 50% or higher indicates that Cohere Transcribe was preferred on average in the comparison.*

## **per-language WERs**

*Figure: per-language error rate averaged over FLEURS, Common Voice 17.0, MLS and Wenet tests sets (where relevant for a given language). CER for zh, ja, ko — WER otherwise*

For more details and results:

- Technical blog post contains WERs and other quality metrics.
- Announcement blog post for more information about the model.
- English, EU and long-form transcription WERs/RTFx are on the Open ASR Leaderboard.

Cohere Transcribe is a performant, dedicated ASR model intended for efficient speech transcription.

Cohere Transcribe demonstrates best-in-class transcription accuracy in 14 languages. As a dedicated speech recognition model, it is also efficient, benefitting from a real-time factor up to three times faster than that of other, dedicated ASR models in the same size range. The model was trained from scratch, and from the outset, we deliberately focused on maximizing transcription accuracy while keeping production readiness top-of-mind.

- **Single language.** The model performs best when remaining in-distribution of a single, pre-specified language amongst the 14 in the range it supports. It does not feature explicit, automatic language detection and exhibits inconsistent performance on code-switched audio.
- **Timestamps/Speaker diarization.** The model does not feature either of these.
- **Silence.** Like most AED speech models, Cohere Transcribe is eager to transcribe, even non-speech sounds. The model thus benefits from prepending a noise gate or VAD (voice activity detection) model in order to prevent low-volume, floor noise from turning into hallucinations.

Cohere Transcribe is supported on the following libraries/platforms:

- `transformers` (see Quick Start above).
- `vLLM` (see vLLM integration above).
- `mlx-audio` for Apple Silicon.
- Rust implementation: `cohere_transcribe_rs`
- In the browser ✨**demo** ✨ (via`transformers.js` and WebGPU)
- Chrome extension: `cohere_transcribe_extension`
- `nano-cohere-transcribe` - blazing fast, particularly for long-form audio. e.g. 18 mins audio transcribed in 2.36s
- Whisper Memos (iOS App).
- Whisperian (Android App).
- hyprwhspr (Linux app).

If you have added support for the model somewhere not included above please raise an issue/PR!

If you find issues with any of these please raise an issue with the respective library.

For errors or additional questions about details in this model card, contact labs@cohere.com or raise an issue.

Terms of Use: We hope that the release of this model will make community-based research efforts more accessible, by releasing the weights of a highly performant 2 billion parameter model to researchers all over the world. This model is governed by an Apache 2.0 license.

To cite this model please use the following bibtex:

```
@misc{julian_mack_2026,
    author       = { Julian Mack and Ekagra Ranjan and Walter Beller-Morales and Bharat Venkitesh and Pierre Richemond },
    title        = { cohere-transcribe-03-2026 (Revision d96e814) },
    year         = 2026,
    url          = { https://huggingface.co/CohereLabs/cohere-transcribe-03-2026 },
    doi          = { 10.57967/hf/8653 },
    publisher    = { Hugging Face }
}
```
- Downloads last month
- 211,468
