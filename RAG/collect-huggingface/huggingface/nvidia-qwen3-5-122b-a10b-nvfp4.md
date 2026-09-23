---
id: collect-huggingface/huggingface/nvidia-qwen3-5-122b-a10b-nvfp4
title: "Qwen3.5-122B-A10B-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp4", "nvfp4", "agent", "apache", "blackwell", "cost", "fp8", "gpu", "inference", "license", "memory", "mixture of experts"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-Qwen3.5-122B-A10B-NVFP4.md
source_anchor: ""
source_lines: [1, 54]
sha256: fede3846c1f4d2031c150243a5b2a85c3219d2be77a3f50d7bc2ffe753eb4b5e
---

# Qwen3.5-122B-A10B-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/Qwen3.5-122B-A10B-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

nvidia/Qwen3.5-122B-A10B-NVFP4 is NVIDIA's NVFP4-quantized version of Alibaba's Qwen3.5-122B-A10B, a Mixture-of-Experts auto-regressive language model with 122B total parameters and 10B activated per token. It was quantized using NVIDIA Model Optimizer (nvidia-modelopt v0.44.0) and is designed for commercial/non-commercial deployment in AI agent systems, chatbots, RAG systems, and other AI-powered applications. NVIDIA notes it is a third-party model developed to a third party's requirements.

The quantization converts the weights and activations of the linear operators within the transformer blocks of the MoE from 16 bits to 4 bits (NVFP4), reducing disk size and GPU memory requirements by approximately 4×. It is ready for inference with vLLM and supports NVIDIA Blackwell hardware on Linux. The model accepts text, image, and video inputs with a context length up to 262K tokens, and outputs text.

The quantization calibration used cnn_dailymail and NVIDIA's Nemotron-Post-Training-Dataset-v2. Evaluation datasets include MMMU Pro, GPQA Diamond, tau2_bench_telecom, SciCode, AA-LCR, and IFBench. The card provides a direct FP8 vs. NVFP4 accuracy comparison: MMMU Pro 75.90 vs. 75.55, GPQA Diamond 87.37 vs. 86.77, SciCode 42.16 vs. 41.79, AA-LCR 65.5 vs. 67.13, IFBench 70.91 vs. 70.80. The baseline is Qwen3.5-122B-A10B-FP8, benchmarked at temperature=0.6, top_p=0.95, max tokens 64000.

License is Apache-2.0. Serving example uses the `nvcr.io/nvidia/vllm:26.04-py3` container with `vllm serve ... --quantization modelopt_fp4 --kv-cache-dtype fp8 --tensor-parallel-size 1 --reasoning-parser qwen3 --enable-auto-tool-choice --tool-call-parser qwen3_coder`. The model card also documents limitations (potential bias/toxic language inherited from web data) and ethical considerations for image/video inputs.

## Key points

- NVIDIA NVFP4 4-bit quantization of Qwen3.5-122B-A10B MoE (122B total / 10B active).
- Quantized with nvidia-modelopt v0.44.0; ~4× reduction in disk/GPU memory.
- Multimodal input (text, image, video); context up to 262K tokens.
- Near-lossless: within ~0.6 points of FP8 across MMMU Pro, GPQA Diamond, SciCode, IFBench.
- Runs on vLLM with NVIDIA Blackwell on Linux; Apache-2.0 license.
- Calibrated on cnn_dailymail and Nemotron-Post-Training-Dataset-v2.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 122B (MoE) |
| Active parameters | 10B |
| Architecture | Qwen3.5-122B-A10B Mixture of Experts (Transformers) |
| Quantization | NVFP4 (weights + activations of MoE linear ops) |
| Quantization tool | nvidia-modelopt v0.44.0 |
| Context length | up to 262K |
| Inputs | Text, Image (RGB), Video (MP4/WebM) |
| Output | Text |
| License | Apache-2.0 |
| Runtime | vLLM (Blackwell, Linux) |
| MMMU Pro (FP8 / NVFP4) | 75.90 / 75.55 |
| GPQA Diamond (FP8 / NVFP4) | 87.37 / 86.77 |
| SciCode (FP8 / NVFP4) | 42.16 / 41.79 |
| AA-LCR (FP8 / NVFP4) | 65.5 / 67.13 |
| IFBench (FP8 / NVFP4) | 70.91 / 70.80 |
| Monthly downloads | ~1.4M |

## Why this source matters for the RAG

This card is a rigorous reference for 4-bit NVFP4 quantization of a large MoE multimodal model, with explicit FP8-vs-NVFP4 accuracy deltas. It documents enterprise-grade serving on NVIDIA Blackwell with vLLM, relevant to cost and VRAM optimization. It strengthens the knowledge base's coverage of quantization and inference optimization.
