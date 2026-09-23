---
id: collect-huggingface/huggingface/minimaxai-minimax-h3
title: "MiniMax-H3 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "China", "EU", "Hugging Face", "MiniMax", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["inference", "license", "multimodal", "omni", "open-weight", "sglang", "video generation", "vllm"]
source: docs/RAG/Collect RAG/03_huggingface/MiniMaxAI-MiniMax-H3.md
source_anchor: ""
source_lines: [1, 54]
sha256: b59c9b31b5116e6c9ba0f2988a0bf147d2570e31e6f44b479751ac908f05d089
---

# MiniMax-H3 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/MiniMaxAI/MiniMax-H3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiniMax H3 is a general-purpose, omni-modal generative system from MiniMax (Hailuo AI) that unifies understanding of text, images, video, and audio contexts and generates video with native stereo audio at resolutions up to 2K and durations up to 15 seconds. Output specs: 4-15s duration, wide aspect-ratio support (21:9 to 9:16), default 768p shorter-side resolution (2K via H3-Regenerate-2K), 24 FPS, 32 kHz stereo audio, and stable dialogue in 11 languages (Arabic, Chinese, English, French, German, Italian, Japanese, Korean, Portuguese, Russian, Spanish).

The system is composed of three modules: (1) H3-Context-IR, a hosted preprocessing/orchestration system that parses free-form multimodal inputs into a structured Context Intermediate Representation (not open-sourced, API provided); (2) H3-Base, the open-sourced omni-transformer that generates 768p audio-video; and (3) H3-Regenerate-2K, which regenerates the 768p output at 2K in-context (not yet open-sourced). H3-Base comes as two task-specific checkpoints: FL2VA (text-to-audio-video, first/last-frame-to-video) and Ref2VA (reference-to-audio-video with up to 9 images, 3 video clips, 3 audio clips, max 12 files). Architecture: H3-Omni-Transformer is a 33B-parameter dense single-stream transformer (~13B params in AdaLN branches, cacheable, not needed for inference-only); text encoded by H3-Encoder (full pretrained Qwen3-VL-32B, hidden states from layer 50); visual inputs via H3-VisualVAE (temporally causal, spatial 16x, temporal 4x, 24 channels, f16t4d24, with ViT-based decoder); audio via H3-AudioVAE (per-channel, 32 kHz to 40 Hz latent tokens). Positional encoding is 3D Multimodal RoPE (MM-RoPE) over (t,h,w). The released weights are CFG-distilled.

License: MiniMax H3 Community License Agreement (applicable territory restricted; license form required for USA/EU/UK/South Korea). Weights in BF16/F32, 33B hub size. Deployment via SGLang, vLLM, diffusers, and ComfyUI; Full 2K workflow combines local H3-Base with H3-Context-IR and H3-Regenerate-2K APIs. Hub reports 3.66M monthly downloads, 143 finetunes, 63 quantizations.

## Key points

- Omni-modal generative system: video with native stereo audio, up to 2K / 15s / 24 FPS.
- Three-module system: H3-Context-IR (hosted), H3-Base (open), H3-Regenerate-2K (hosted).
- H3-Omni-Transformer: 33B dense single-stream, ~13B in cacheable AdaLN branches.
- H3-Encoder = pretrained Qwen3-VL-32B (layer-50 hidden states); VisualVAE f16t4d24 + AudioVAE.
- Two open checkpoints: FL2VA (text/first-last-frame) and Ref2VA (multi-reference).
- 3D MM-RoPE positional encoding; released weights are CFG-distilled.
- MiniMax H3 Community License; deploy via SGLang, vLLM, diffusers, ComfyUI.
- 3.66M downloads/month; 143 finetunes, 63 quantizations on the Hub.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | MiniMaxAI |
| Model name | MiniMax-H3 |
| Type | Omni-modal (video + stereo audio generation) |
| Transformer | H3-Omni-Transformer, 33B dense |
| AdaLN params | ~13B (cacheable, not required for inference) |
| Text encoder | H3-Encoder (Qwen3-VL-32B, layer 50) |
| VisualVAE | f16t4d24, spatial 16x, temporal 4x, 24 channels |
| AudioVAE | per-channel, 32 kHz -> 40 Hz latents |
| Positional encoding | 3D MM-RoPE (t, h, w) |
| Output | up to 2K, 4-15s, 24 FPS, 32 kHz stereo |
| Checkpoints | FL2VA (BF16), Ref2VA (BF16), CFG-distilled |
| License | MiniMax H3 Community License |
| Weight formats | F32, BF16 |
| Languages | 11 stable (incl. en, zh, ar, ru, es) |
| Deployment | SGLang, vLLM, diffusers, ComfyUI |
| Downloads/month | 3,664,216 |

## Why this source matters for the RAG

This card documents a frontier open-weight audio-video generation system, detailing a modular three-stage architecture (context IR, base generation, 2K regeneration) and the VAE/transformer stack behind it. It is essential reference material for retrieval on video generation, audio-video models, omni-modal systems, and open-weight creative-model licensing.
