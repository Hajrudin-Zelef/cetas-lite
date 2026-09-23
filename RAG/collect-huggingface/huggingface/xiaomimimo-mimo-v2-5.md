---
id: collect-huggingface/huggingface/xiaomimimo-mimo-v2-5
title: "MiMo-V2.5 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "SGLang", "Xiaomi", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "attention", "benchmarks", "context window", "distillation", "fp8", "kv cache", "license", "mit license", "moe", "multimodal", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/XiaomiMiMo-MiMo-V2.5.md
source_anchor: ""
source_lines: [1, 59]
sha256: 3a9c003c3c7d8c5d31da3cae5686d234d843434e85d6d7f46f09f72008b66b1d
---

# MiMo-V2.5 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/XiaomiMiMo/MiMo-V2.5
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiMo-V2.5 is a native omnimodal model from Xiaomi MiMo with strong agentic capabilities, supporting text, image, video, and audio understanding within a unified architecture. Built on the MiMo-V2-Flash backbone with dedicated vision and audio encoders, it is a sparse MoE with 310B total / 15B activated parameters and up to 1M-token context. It inherits the hybrid SWA/GA attention design (5:1 ratio, 128 sliding window, learnable attention-sink bias) which reduces KV-cache storage by nearly 6x. Architecture: 48 layers (1 dense + 47 MoE), hidden size 4096, 64 attention heads, 8 KV heads (GA) / 4 (SWA), head dims QK/V = 192/128, 256 routed experts (8 per token), MoE intermediate 2048, 39 SWA + 9 full-attention layers, 3 MTP layers (329M params) for speculative decoding.

Vision encoder: 729M-param MiMo ViT (28 layers: 24 SWA + 4 full) with hybrid window attention. Audio encoder: 261M-param audio transformer (24 layers: 12 SWA + 12 full) initialized from MiMo-Audio-Tokenizer. Training: ~48T tokens total with FP8 mixed precision, following text pre-training, projector warmup, multimodal pre-training, SFT & agentic post-training (context progressively extended 32K -> 256K -> 1M), then RL & Multi-Teacher On-Policy Distillation (MOPD). The context window supports up to 1M tokens.

License: MIT. Weights in F32/BF16/F8_E4M3. Deployment via SGLang (recommended cookbook; --reasoning-parser qwen3, --tool-call-parser mimo, fp8 quantization) and vLLM, with recommended sampling temperature=1.0, top_p=0.95. Hub results include SWE-Bench Pro 56.1 and Terminal-Bench 2.0 65.8. The card also notes an important config update: config.json/tokenizer_config.json were updated post-release and older downloads must be re-pulled. Community has produced 29 quantizations. Also available on ModelScope, AI Studio, and API platform.

## Key points

- Native omnimodal sparse MoE: 310B total / 15B activated; text, image, video, audio.
- Hybrid SWA/GA attention (5:1, window 128) cuts KV cache ~6x via attention-sink bias.
- 729M ViT vision encoder + 261M audio encoder from MiMo-Audio.
- 3-layer MTP modules (329M params) for speculative decoding and RL acceleration.
- Trained on ~48T tokens with FP8 mixed precision; context to 1M tokens.
- Post-training: SFT, agentic RL, Multi-Teacher On-Policy Distillation (MOPD).
- MIT license; F32/BF16/FP8 weights; config files were updated post-release.
- Deploy via SGLang (fp8, mimo parsers) or vLLM; temp 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | XiaomiMiMo |
| Model name | MiMo-V2.5 |
| Architecture | Sparse MoE, hybrid SWA/GA + MTP |
| Total params | 310B |
| Activated params | 15B |
| Layers | 48 (1 dense + 47 MoE) |
| Hidden size | 4096 |
| Attention heads / KV heads | 64 / 8 (GA), 4 (SWA) |
| Routed experts | 256 (8 per token) |
| SWA / GA layers | 39 / 9 |
| SWA window | 128 |
| MTP layers | 3 (329M params) |
| Vision encoder | 729M MiMo ViT (24 SWA + 4 full) |
| Audio encoder | 261M audio transformer (12 SWA + 12 full) |
| Context length | Up to 1M tokens |
| Training tokens | ~48T (FP8 mixed) |
| License | MIT |
| Weight formats | F32, BF16, F8_E4M3 |
| Benchmarks | SWE-Bench Pro 56.1, Terminal-Bench 2.0 65.8 |
| Sampling | temperature=1.0, top_p=0.95 |
| Downloads/month | 282,074 |

## Why this source matters for the RAG

This card documents the accessible 15B-active omnimodal sibling of MiMo-V2.5-Pro, including the full hybrid-attention backbone, dual encoders, and the 5-stage training recipe with context extension to 1M. It is key reference material for retrieval on omnimodal MoE models, hybrid attention, and multi-token prediction.
