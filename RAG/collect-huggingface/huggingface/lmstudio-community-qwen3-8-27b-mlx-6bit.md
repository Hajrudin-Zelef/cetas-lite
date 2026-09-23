---
id: collect-huggingface/huggingface/lmstudio-community-qwen3-8-27b-mlx-6bit
title: "Qwen3.8-27B-MLX-6bit - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Apple", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "agentic", "apache", "attention", "fine-tuning", "inference", "license", "parameters", "quantization", "reasoning", "research", "sglang"]
source: docs/RAG/Collect RAG/03_huggingface/lmstudio-community-Qwen3.8-27B-MLX-6bit.md
source_anchor: ""
source_lines: [1, 50]
sha256: cec1c47c1aec9b804e4d53bd6366448d96e8f951e4c902fa30d90493e9197b10
---

# Qwen3.8-27B-MLX-6bit - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/lmstudio-community/Qwen3.8-27B-MLX-6bit
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

lmstudio-community/Qwen3.8-27B-MLX-6bit is a 6-bit MLX quantized version of Qwen/Qwen3.8-27B, produced by the LM Studio team using mlx_vlm. It is part of the LM Studio Community Models highlight program. The quantization is optimized for Apple Silicon, enabling local inference of the 27B dense vision-language model on Mac hardware. The MLX variant weighs approximately 22.8 GB on disk.

The base model, Qwen3.8-27B, is the most capable generation in the Qwen open-model family to date and the compact, deployment-friendly dense member of the Qwen3.8 series. It is a native vision-language model (image-text-to-text) that understands images and videos, with flexible thinking control. It builds on the Qwen3.5 architectural foundation and delivers gains in coding, professional work, research, and long-horizon agentic tasks, with stronger autonomous planning and better handling of environment feedback.

Architecture details of Qwen3.8-27B: 27B parameters, hidden dimension 5120, 64 layers, vocabulary 248,320 (padded), and a hybrid layout of 16 × (3 × (Gated DeltaNet → FFN) → 1 × (Gated Attention → FFN)). It supports Multi-Token Prediction (MTP). Context length is 262,144 tokens natively, extensible up to 1,000,000 tokens. Thinking mode is on by default and can be disabled per request; reasoning depth is tunable with `reasoning_effort`, and reasoning context is retained via `preserve_thinking`.

The model is Apache-2.0 licensed. Because this is an MLX quant, it runs via mlx-vlm / mlx-lm on Apple Silicon and can also be served through vLLM, SGLang, or LM Studio. The base model card notes developer role support (for agentic tools like Codex), MTP for fast inference, improved tool calling, and fine-tuning in Unsloth Desktop.

## Key points

- 6-bit MLX quantization of Qwen3.8-27B, optimized for Apple Silicon; ~22.8 GB.
- Base model is a 27B dense native vision-language model (image + video).
- Context 262,144 native, extensible to 1,000,000 tokens.
- Flexible thinking control: `reasoning_effort`, `preserve_thinking`, disable per request.
- Apache-2.0 license; runs via mlx-vlm, LM Studio, vLLM, SGLang.
- 64 layers, hidden dim 5120, hybrid Gated DeltaNet/Gated Attention + FFN.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 27B (dense) |
| Active parameters | 27B |
| Quantization | 6-bit MLX |
| File size | ~22.8 GB |
| Architecture | Hybrid Gated DeltaNet + Gated Attention, Vision Encoder |
| Layers | 64 |
| Hidden dimension | 5120 |
| Vocabulary | 248,320 (padded) |
| Context length | 262,144 native; up to 1,000,000 |
| License | Apache-2.0 |
| Base model | Qwen/Qwen3.8-27B |
| Monthly downloads | ~4.4M |

## Why this source matters for the RAG

This card documents how to run a frontier dense vision-language model locally on Apple Silicon via MLX, a key scenario for privacy-conscious and offline deployments. It provides concrete quantization size figures and tooling paths (mlx-vlm, LM Studio). It supports the RAG's coverage of local AI hardware and model formats.
