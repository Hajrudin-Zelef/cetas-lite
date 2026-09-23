---
id: collect-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8
title: "Nemotron 3 Nano Omni 30B-A3B Reasoning FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-04-28", "2026-09-23"]
keywords: ["fp8", "omni", "reasoning", "gpu", "license", "llama", "llama.cpp", "moe", "multimodal", "nvfp4", "nvidia", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8.md
source_anchor: ""
source_lines: [1, 47]
sha256: 1b2f2778bfa83bfe4031b40d06f8c05721983fc598bd5d5e7f81aea2af50467e
---

# Nemotron 3 Nano Omni 30B-A3B Reasoning FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA Nemotron 3 Nano Omni is a multimodal large language model that unifies video, audio, image, and text understanding for enterprise-grade Q&A, summarization, transcription and document intelligence. It extends the Nemotron Nano family with integrated video+speech comprehension, GUI understanding, OCR and speech transcription, targeting meeting recordings, media and entertainment assets, training videos and complex business documents. The architecture is a Mamba2-Transformer hybrid Mixture-of-Experts (MoE) backbone (31B total parameters, ~3B active per token) combined with a C-RADIOv4-H vision encoder and a Parakeet speech encoder. It supports up to 256K tokens of context, English only, and accepts video (mp4 up to 2 minutes), audio (wav/mp3 up to 1 hour), images (RGB) and text, producing text output with JSON, reasoning chains, tool calling and word-level transcription timestamps. Reasoning is on by default and toggled via `enable_thinking`. This FP8 checkpoint is 33 GB, versus 62 GB for BF16 and 21 GB for NVFP4; minimum GPU for FP8 is one L40S 48GB (RTX Pro 6000 or B200 recommended). It runs on vLLM (0.20.0), SGLang (BF16 only), TensorRT-LLM, TensorRT Edge-LLM, llama.cpp and Ollama, with deployment guidance for DGX Spark, Jetson Thor and RTX hardware. Recommended parameters are temperature 0.6 / top_p 0.95 / max_tokens 20480 / reasoning_budget 16384 for thinking mode, and temperature 0.2 / top_k 1 / max_tokens 1024 for instruct mode. Training used ~717.0B tokens across 1395 datasets (354.6M data points), with heavy synthetic data generation. It is governed by the NVIDIA Open Model Agreement and is available for commercial use.

## Key points

- Any-to-any multimodal enterprise model: video, audio, image, text in; text out.
- Mamba2-Transformer hybrid MoE: 31B total / ~3B active parameters.
- Encoders: C-RADIOv4-H (vision), Parakeet (speech); 256K context; English only.
- FP8 checkpoint at 33 GB (BF16 62 GB, NVFP4 21 GB); min GPU one L40S 48GB.
- Reasoning on by default (`enable_thinking`), supports tool calling and word timestamps.
- Runtimes: vLLM 0.20.0, SGLang (BF16), TensorRT-LLM, llama.cpp, Ollama.
- Training data ~717B tokens across 1395 datasets; NVIDIA Open Model Agreement.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 31B (Mamba2-Transformer hybrid MoE) |
| Active parameters | ~3B per token |
| Architecture | NemotronH_Nano_Omni_Reasoning_V3 |
| Context length | 256K tokens |
| Modalities in | Video, Audio, Image, Text |
| Modality out | Text |
| Encoders | CRADIO v4-H (vision), Parakeet (speech) |
| Precision | FP8 (33 GB) |
| Other precisions | BF16 (62 GB), NVFP4 (21 GB) |
| Min GPU (FP8) | 1× L40S 48GB |
| Language | English only |
| License | NVIDIA Open Model Agreement |
| Training data | ~717.0B tokens, 1395 datasets |
| Release date | 2026-04-28 |

## Why this source matters for the RAG

This card documents NVIDIA's omni-modal enterprise model, providing structured detail on hybrid Mamba2/MoE architecture, multimodal encoders and FP8/NVFP4 deployment trade-offs. It is a key reference for video, audio and document-intelligence use cases in the knowledge base.
