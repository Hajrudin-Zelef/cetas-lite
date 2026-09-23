---
id: collect-huggingface/huggingface/coherelabs-north-small-translate-1-0-w4a16
title: "North-Small-Translate-1.0-w4a16 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Cohere", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "attention", "benchmarks", "blackwell", "cohere", "embeddings", "fp4", "fp8", "gpus", "license", "moe", "nvfp4"]
source: docs/RAG/Collect RAG/03_huggingface/CohereLabs-North-Small-Translate-1.0-w4a16.md
source_anchor: ""
source_lines: [1, 55]
sha256: a485765351a6d6b1b56f11e54eb909d1a3fadd034d2f2b2441bc46380677a748
---

# North-Small-Translate-1.0-w4a16 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/CohereLabs/North-Small-Translate-1.0-w4a16
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository holds the NVFP4 W4A16 checkpoint of Cohere's North Small Translate model, an open-weights research release of a sparse Mixture-of-Experts model with 25B active parameters and 218B total parameters specialized for high-quality machine translation across 50 languages. It is one of three production-served quantizations (BF16, FP8, NVFP4 W4A16); the W4A16 variant uses 4-bit weights with 16-bit activations in the compressed-tensors format, with a group size of 16 and FP8 scales. Only the MoE experts are quantized — the attention projections, routers, and output head remain at higher precision — since experts hold the overwhelming majority of parameters, capturing nearly all savings while leaving precision-sensitive parts untouched. This reduces the footprint from roughly 437 GB to 131 GB. Because only weights are quantized, this format does not require native FP4 hardware and runs on pre-Blackwell GPUs such as Hopper and Ada.

Architecture details are inherited from the base model: decoder-only sparse MoE, 128 experts with 8 activated per token plus shared experts, interleaved sliding-window attention (window 4096, RoPE) with global attention without positional embeddings at a 3:1 ratio (Command A design), sigmoid-activated router normalized over selected top-k, post-trained for translation quality. Context length: 16K input & 16K output. Languages: 50.

Usage: this checkpoint is designed to be served with vLLM and is not compatible with transformers (no native 4-bit support for this format in transformers) — the BF16 weights should be used in transformers. vLLM serving: uv pip install vllm and cohere_melody>=0.9.0, then `vllm serve CohereLabs/North-Small-Translate-1.0-w4a16 -tp 2 --max-model-len 32768 --tool-call-parser cohere_command4 --reasoning-parser cohere_command4 --enable-auto-tool-choice`. Greedy decoding recommended (production default). Minimum hardware: 1 x B200 (Blackwell) or 2 x H100 (Hopper). Evaluation: WMT26 all-languages 83.60, rising to 84.36 with the agentic multi-pass translation workflow.

License: CC BY-NC 4.0 (non-commercial) + Cohere Labs Acceptable Use Policy; gated access. Hub: 125B params, BF16 + U8 tensors, ~30 downloads/month.

## Key points

- NVFP4 W4A16 quantization of North Small Translate (4-bit weights, 16-bit activations).
- 25B active / 218B total MoE; 128 experts (8 per token) + shared experts.
- Experts quantized only; attention/routers/head kept high precision; 437 GB -> 131 GB.
- Runs on pre-Blackwell GPUs (Hopper, Ada) — no native FP4 hardware required.
- vLLM-only; not compatible with transformers (use BF16 checkpoint there).
- Requires cohere_melody>=0.9.0 and cohere_command4 parsers; greedy decoding.
- 50 languages; 16K in / 16K out; WMT26 83.60 (84.36 multi-pass).
- CC BY-NC 4.0 + Acceptable Use Policy; gated access; min 1 x B200 or 2 x H100.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | CohereLabs |
| Model name | North-Small-Translate-1.0-w4a16 |
| Base model | CohereLabs/North-Small-Translate-1.0 |
| Quantization | NVFP4 W4A16 (compressed-tensors, group size 16, FP8 scales) |
| Total params | 218B |
| Activated params | 25B |
| Experts | 128 (8 per token) + shared experts |
| Context length | 16K input / 16K output |
| Footprint | ~131 GB (from ~437 GB) |
| Languages | 50 |
| License | CC BY-NC 4.0 + Acceptable Use Policy |
| Tensor types | BF16, U8 |
| Hardware | 1 x B200 or 2 x H100 |
| Benchmarks | WMT26 83.60 (84.36 multi-pass) |
| Downloads/month | 30 |

## Why this source matters for the RAG

This card documents a production-grade 4-bit quantization path for a 218B translation MoE, detailing the expert-only quantization strategy, hardware requirements, and vLLM serving setup. It is key reference material for retrieval on W4A16/NVFP4 quantization, MoE compression, and efficient open-weight model serving.
