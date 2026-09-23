---
id: collect-huggingface/huggingface/unsloth-qwen3-8-27b-nvfp4
title: "Qwen3.8-27B-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["nvfp4", "qwen", "agentic", "apache", "attention", "fine-tuning", "fp8", "gguf", "license", "parameters", "quantization", "reasoning"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Qwen3.8-27B-NVFP4.md
source_anchor: ""
source_lines: [1, 50]
sha256: 5720c0745b1f971782e5b25115f7edc21e79ef5dea5440b26c5509b3d5fdc25d
---

# Qwen3.8-27B-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Qwen3.8-27B-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/Qwen3.8-27B-NVFP4 is Unsloth's NVFP4 (4-bit floating point) quantized version of Qwen/Qwen3.8-27B. It uses Unsloth Dynamic V3.0 (preview) for state-of-the-art quantization performance. The card notes it works with vLLM, and with SGLang only at v0.5.19 because the `lm_head` is FP8 rather than 4-bit. This enables running and fine-tuning the 27B model locally, including inside Unsloth Desktop.

The base model, Qwen3.8-27B, is the most capable generation in the Qwen open-model family to date, built on the Qwen3.5 architectural foundation. It is a native vision-language model (understands images and videos) designed for coding, professional work, research, and long-horizon agentic tasks. Highlights include stronger autonomous planning, better handling of environment feedback, broad harness compatibility, flexible thinking control, and native vision-language understanding.

Architecture: 27B dense parameters, hidden dimension 5120, 64 layers, vocabulary 248,320 (padded), hybrid layout of 16 × (3 × (Gated DeltaNet → FFN) → 1 × (Gated Attention → FFN)). It uses Gated DeltaNet linear attention (48 V heads, 16 QK heads) and Gated Attention (24 Q heads, 4 KV heads). FFN intermediate dimension is 17,408. It supports Multi-Token Prediction (MTP). Context length is 262,144 tokens natively, extensible to 1,000,000 tokens.

Thinking mode is on by default and can be disabled per request; reasoning depth is tunable with `reasoning_effort`, and reasoning context from historical messages is retained via `preserve_thinking`. Recommended sampling: thinking mode `temperature=1.0, top_p=0.95, top_k=20`; instruct mode `temperature=0.7, top_p=0.80, top_k=20`. License is Apache-2.0. The card also highlights developer role support (Codex-style agentic tools), improved tool calling with nested object parsing, and Unsloth Desktop fine-tuning.

## Key points

- Unsloth NVFP4 (4-bit) dynamic quantization of Qwen3.8-27B.
- Base model: 27B dense native vision-language model (image + video).
- Works with vLLM; SGLang only v0.5.19 (FP8 `lm_head`).
- Context 262,144 native, extensible to 1,000,000 tokens.
- Flexible thinking control (`reasoning_effort`, `preserve_thinking`).
- Apache-2.0; supports local run and fine-tuning via Unsloth Desktop.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 27B (dense) |
| Active parameters | 27B |
| Quantization | NVFP4 (Unsloth Dynamic V3.0 preview) |
| Architecture | Hybrid Gated DeltaNet + Gated Attention, Vision Encoder |
| Layers | 64 |
| Hidden dimension | 5120 |
| FFN intermediate dimension | 17,408 |
| Vocabulary | 248,320 (padded) |
| Context length | 262,144 native; up to 1,000,000 |
| License | Apache-2.0 |
| Base model | Qwen/Qwen3.8-27B |
| Monthly downloads | ~3.2M |

## Why this source matters for the RAG

This card documents 4-bit NVFP4 quantization for a frontier dense vision-language model, a critical technique for reducing VRAM requirements and enabling local deployment. It captures practical serving constraints (vLLM vs. SGLang compatibility) and Unsloth's tooling. It complements the MLX and GGUF quant entries in the knowledge base.
