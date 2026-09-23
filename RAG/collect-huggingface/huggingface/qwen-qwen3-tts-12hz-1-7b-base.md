---
id: collect-huggingface/huggingface/qwen-qwen3-tts-12hz-1-7b-base
title: "Qwen3-TTS-12Hz-1.7B-Base - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "China", "Hugging Face", "MiniMax", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "agents", "apache", "benchmark", "benchmarks", "fine-tuning", "inference", "latency", "license", "multimodal", "omni", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3-TTS-12Hz-1.7B-Base.md
source_anchor: ""
source_lines: [1, 49]
sha256: 6e97abb8aa9002b4734ce2dd58c22d84f43b310784d9b766248d166ef265badc
---

# Qwen3-TTS-12Hz-1.7B-Base - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3-TTS-12Hz-1.7B-Base
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3-TTS-12Hz-1.7B-Base is the base text-to-speech model of the Qwen3-TTS family, released under Apache-2.0. It is designed for rapid voice cloning from a short (3-second) user audio clip and can be used to fine-tune other TTS models. The Qwen3-TTS family covers 10 major languages (Chinese, English, Japanese, Korean, German, French, Russian, Portuguese, Spanish, Italian) plus multiple dialectal voice profiles, with strong contextual understanding that adapts tone, speaking rate, and emotional expression based on instructions and text semantics.

The model uses a discrete multi-codebook LM architecture (universal end-to-end), powered by the self-developed Qwen3-TTS-Tokenizer-12Hz. This tokenizer achieves efficient acoustic compression and high-dimensional semantic modeling of speech, preserving paralinguistic and acoustic-environment information for high-fidelity reconstruction through a lightweight non-DiT architecture. A key feature is the Dual-Track hybrid streaming generation architecture, which supports both streaming and non-streaming generation with end-to-end synthesis latency as low as 97ms (first audio packet immediately after a single character). It is instruction-aware, supporting natural-language control over timbre, emotion, and prosody.

The model is roughly 2B parameters (labeled 1.7B) and ships as BF16 Safetensors. It is used via the `qwen-tts` Python package (Python 3.12, FlashAttention 2 recommended), the `qwen-tts-demo` web UI, or vLLM-Omni (offline inference; online serving planned). DashScope APIs are available for custom voice, voice clone, and voice design.

Benchmarks: on Seed-TTS test-zh/test-en WER, Qwen3-TTS-12Hz-1.7B-Base scores 0.77/1.24, competitive with CosyVoice 3 and MiniMax-Speech. It leads multilingual content consistency and speaker similarity in many languages and performs strongly on cross-lingual and long speech generation. The Qwen3-TTS-Tokenizer-12Hz achieves PESQ_WB 3.21, STOI 0.96, UTMOS 4.16, SIM 0.95.

## Key points

- 1.7B base TTS model for 3-second rapid voice cloning and fine-tuning.
- 10 languages; discrete multi-codebook LM architecture, no DiT bottleneck.
- Dual-Track streaming: latency as low as 97ms; supports streaming and non-streaming.
- Powered by Qwen3-TTS-Tokenizer-12Hz; instruction-controlled timbre/emotion/prosody.
- Apache-2.0 license; run via `qwen-tts`, web UI, or vLLM-Omni.
- Seed-TTS WER 0.77 (zh) / 1.24 (en); tokenizer PESQ_WB 3.21, SIM 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | ~1.7B (reported 2B on HF) |
| Architecture | Discrete multi-codebook LM (Dual-Track streaming) |
| Tokenizer | Qwen3-TTS-Tokenizer-12Hz |
| Languages | 10 major + dialects |
| Streaming latency | as low as 97ms |
| License | Apache-2.0 |
| Precision | BF16 |
| Seed-TTS WER (zh / en) | 0.77 / 1.24 |
| Tokenizer PESQ_WB | 3.21 |
| Tokenizer STOI / UTMOS / SIM | 0.96 / 4.16 / 0.95 |
| Monthly downloads | ~3.9M |

## Why this source matters for the RAG

Qwen3-TTS-12Hz-1.7B-Base extends the knowledge base beyond text LLMs into speech synthesis and voice cloning, a key multimodal capability for local AI pipelines. Its detailed benchmark tables and architecture notes provide authoritative grounding on modern TTS quality and latency. It supports questions about building voice-enabled agents and audio content generation.
