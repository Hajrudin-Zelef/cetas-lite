---
id: collect-huggingface/huggingface/nvidia-llama-4-scout-17b-16e-instruct-nvfp4
title: "Llama-4-Scout-17B-16E-Instruct-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["EU", "Hugging Face", "Meta", "Microsoft", "Nvidia", "TensorRT-LLM"]
dates: ["2025-07-28", "2026-09-23"]
keywords: ["fp4", "llama", "nvfp4", "scout", "agent", "agents", "blackwell", "gpu", "gpus", "inference", "license", "memory"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-Llama-4-Scout-17B-16E-Instruct-NVFP4.md
source_anchor: ""
source_lines: [1, 51]
sha256: 904fee86c0abeb492ce4dcb56d55e76d105c74ddefb47c266839529b9430a80c
---

# Llama-4-Scout-17B-16E-Instruct-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/Llama-4-Scout-17B-16E-Instruct-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is NVIDIA's post-training quantized (NVFP4) version of Meta's Llama 4 Scout 17B 16E Instruct model. Llama 4 Scout is an auto-regressive language model using a Mixture-of-Experts (MoE) architecture with early fusion for native multimodality. NVIDIA quantized the weights and activations of the linear operators within transformer blocks to the FP4 data type, ready for inference with TensorRT-LLM. This reduces bits per parameter from 16 to 4, cutting disk size and GPU memory requirements by approximately 3.3x. The model is intended for developers who want off-the-shelf pre-quantized models for AI agent systems, chatbots, RAG systems, and other AI applications. It accepts multilingual text and up to 5 images, outputs multilingual text and code, and supports a context length up to 1M tokens. Quantization was performed with nvidia-modelopt v0.33.0, using cnn_dailymail as the calibration dataset. The model runs on NVIDIA Blackwell, Hopper, and Ampere GPUs under Linux, using TensorRT-LLM (test hardware: B200). The evaluation dataset list includes MMMU Pro, GPQA Diamond, HLE, LiveCodeBench, SciCode, HumanEval, AIME 2024, and MATH-500. Accuracy is well preserved versus the BF16 reference: MMLU Pro 74 (vs 75), GPQA Diamond 56 (vs 57), HLE Challenge 4 (vs 4), SciCode 24 (vs 26), MATH-500 81 (vs 82), and AIME 2024 31 (vs 30). The model is commercially and non-commercially usable under the NVIDIA Open Model License, with the additional Llama 4 Community License terms ("Built with Llama"). Deployment geography is global except the European Union. The hub reports 56B params (BF16, F8_E4M3, U8) and 6,364 monthly downloads. Release date on Hugging Face: July 28, 2025.

## Key points

- NVIDIA NVFP4 post-training quantization of Meta Llama 4 Scout 17B 16E Instruct.
- Quantizes linear-operator weights and activations to FP4, ~3.3x smaller disk/GPU memory.
- Optimized for TensorRT-LLM inference on NVIDIA Blackwell, Hopper, and Ampere GPUs.
- Quantized with nvidia-modelopt v0.33.0; calibration on cnn_dailymail; tested on B200.
- Multilingual text + up to 5 images input; 1M-token context; text and code output.
- Near-parity accuracy: MMLU Pro 74, GPQA-D 56, MATH-500 81, AIME 2024 31.
- NVIDIA Open Model License plus Llama 4 Community License; EU excluded from deployment.
- Target use: AI agents, chatbots, RAG systems, and other applications.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | nvidia |
| Model name | Llama-4-Scout-17B-16E-Instruct-NVFP4 |
| Base model | meta-llama/Llama-4-Scout-17B-16E-Instruct |
| Architecture | Transformers / Llama4 MoE |
| Quantized size | 56B params (hub) |
| Quantization | NVFP4 (weights + activations of linear ops) |
| Quantization tool | nvidia-modelopt v0.33.0 |
| Calibration dataset | cnn_dailymail |
| Memory reduction | ~3.3x vs 16-bit |
| Context length | up to 1M tokens |
| Input | Multilingual text + up to 5 images |
| Runtime | TensorRT-LLM |
| Hardware | NVIDIA Blackwell, Hopper, Ampere |
| Eval (FP4 vs BF16) | MMLU Pro 74/75; GPQA-D 56/57; MATH-500 81/82; AIME 2024 31/30 |
| License | NVIDIA Open Model License + Llama 4 Community License |
| Release date | July 28, 2025 |
| Downloads/month | 6,364 |

## Why this source matters for the RAG

This card documents an NVIDIA-optimized FP4 deployment path for Llama 4 Scout, with near-lossless accuracy and ~3.3x memory savings for TensorRT-LLM. It is a key reference for NVIDIA hardware quantization, FP4 inference, and production RAG deployment.
