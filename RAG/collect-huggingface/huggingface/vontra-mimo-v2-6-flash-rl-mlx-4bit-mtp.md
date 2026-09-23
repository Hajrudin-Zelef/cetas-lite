---
id: collect-huggingface/huggingface/vontra-mimo-v2-6-flash-rl-mlx-4bit-mtp
title: "MiMo-V2.6-Flash-RL-MLX-4bit-MTP - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Apple", "Hugging Face", "OpenAI", "Xiaomi"]
dates: ["2026-09-23"]
keywords: ["attention", "decode", "fp8", "license", "memory", "mit license", "moe", "mxfp4", "quantization", "safetensors", "throughput"]
source: docs/RAG/Collect RAG/03_huggingface/Vontra-MiMo-V2.6-Flash-RL-MLX-4bit-MTP.md
source_anchor: ""
source_lines: [1, 52]
sha256: a8df5dd2be36817bd5376d733a3ea21213768947c93369673942008abb9efe2a
---

# MiMo-V2.6-Flash-RL-MLX-4bit-MTP - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Vontra/MiMo-V2.6-Flash-RL-MLX-4bit-MTP
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is a tested Apple Silicon (MLX) conversion of XiaomiMiMo/MiMo-V2.6-Flash-RL, published by the community user Vontra. It converts the text backbone of MiMo-V2.6-Flash-RL for the MLX framework: dense projections use 4-bit affine quantization with group size 64, while the model's native MXFP4 MoE experts remain in their original group-size-32 format, yielding an average of 4.257 bits per weight. The checkpoint includes the native three-layer MTP payload (converted to 4-bit affine weights) in mtp/model_mtp.safetensors, plus the upstream five-layer DFlash drafter, vision encoder, audio encoder, and audio tokenizer packaged so they need no second download.

Measured on a 256 GB M3 Ultra Mac Studio (oMLX 0.7.0.dev2, MLX 0.32.2, mlx-lm 0.31.3): sustained generation 59.4 tok/s (128-token decode), prompt processing 477.1 tok/s at 512 tokens and 562.8 tok/s at 2,048 tokens, peak unified memory 164.3 GB short-context / 166.8 GB at 2,048 tokens, quantized text model about 154 GiB, complete repository about 160 GiB. A 256 GB Mac is recommended. Smoke tests produced correct arithmetic, factual explanations, and coherent Python.

Usage: pip install "mlx-lm>=0.31.3" then python -m mlx_lm generate (or mlx_lm.chat / mlx_lm.server for an OpenAI-compatible endpoint). Quantization notes: MiMo-V2.6 stores fused attention tensors in tensor-parallel order with FP8 scale grids padded per shard; this conversion reconstructs those shards before quantization (skipping this produces broken output). Current oMLX/MLX text generation runs the target model correctly but does not yet execute MTP, DFlash, vision, or audio paths; those assets are packaged for MiMo-aware runtimes. The upstream model (309B total / 15B active, MIT license) is credited to the Xiaomi MiMo team; this repo is a community MLX conversion with measured results.

## Key points

- Community MLX conversion of XiaomiMiMo/MiMo-V2.6-Flash-RL for Apple Silicon.
- 4-bit affine quantization (group size 64) for dense weights; MXFP4 experts kept at group size 32.
- Average 4.257 bits per weight; ~154 GiB text model on disk, ~160 GiB full repo.
- Packages native MTP payload, DFlash drafter, vision/audio encoders for MiMo-aware runtimes.
- Measured 59.4 tok/s generation and ~560 tok/s prompt processing on 256 GB M3 Ultra.
- Fused-attention shards must be reconstructed before quantization to avoid broken output.
- MIT license (upstream); run via mlx-lm generate/chat/server.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Vontra (community) |
| Model name | MiMo-V2.6-Flash-RL-MLX-4bit-MTP |
| Base model | XiaomiMiMo/MiMo-V2.6-Flash-RL |
| Framework | MLX |
| Quantization | 4-bit affine (dense, gs64) + native MXFP4 experts (gs32) |
| Bits/weight (avg) | 4.257 |
| Total params | 309B (upstream) |
| Active params | 15B (upstream) |
| Text model size | ~154 GiB |
| Repository size | ~160 GiB |
| Speed (M3 Ultra 256GB) | 59.4 tok/s gen; 477–563 tok/s prompt processing |
| Peak memory | 164–167 GB |
| License | MIT (upstream) |
| Hardware | 256 GB Apple Silicon Mac recommended |
| Downloads/month | 1,126 |

## Why this source matters for the RAG

This card documents a real-world quantization path for a flagship MoE on Apple Silicon, including measured throughput/memory figures and critical conversion pitfalls (tensor-parallel shard reconstruction). It is valuable reference material for retrieval on MLX conversions, 4-bit MoE quantization, and local Mac deployment of large open models.
