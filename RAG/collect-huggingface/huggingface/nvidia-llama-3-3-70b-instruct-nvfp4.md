---
id: collect-huggingface/huggingface/nvidia-llama-3-3-70b-instruct-nvfp4
title: "Llama-3.3-70B-Instruct-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "Nvidia", "TensorRT-LLM"]
dates: ["2026-09-23"]
keywords: ["fp4", "llama", "nvfp4", "benchmarks", "blackwell", "gpu", "inference", "license", "memory", "nvidia", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-Llama-3.3-70B-Instruct-NVFP4.md
source_anchor: ""
source_lines: [1, 44]
sha256: 54a178e256575fc3d10a472840812ea524f5832b4e1d148abb1008e17c521a57
---

# Llama-3.3-70B-Instruct-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/Llama-3.3-70B-Instruct-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The NVIDIA Llama 3.3 70B Instruct FP4 (NVFP4) model is the quantized version of Meta's Llama 3.3 70B Instruct autoregressive language model, produced by NVIDIA using TensorRT Model Optimizer (modelopt v0.23.0). It is a third-party model: NVIDIA explicitly states the model is not owned or developed by NVIDIA but built to a third party's requirements, and refers users to the original Meta-Llama-3.3-70B-Instruct card. The model was obtained by quantizing the weights and activations of the linear operators within transformer blocks to FP4 data type, ready for inference with TensorRT-LLM. This optimization reduces bits per parameter from 16 to 4, cutting disk size and GPU memory requirements by approximately 3.3x, and reduces reported model size to 41B params. Input is text with a context length of up to 128K tokens; output is text. Supported runtime is TensorRT-LLM and supported hardware is NVIDIA Blackwell (tested on B200), with Linux as the preferred OS. Calibration used the cnn_dailymail dataset. Accuracy benchmarks (BF16 vs FP4): MMLU 83.3 vs 81.1, GSM8K CoT 95.3 vs 92.6, ARC Challenge 93.7 vs 93.3, IFEVAL 92.1 vs 92.0. Licensing is governed by the NVIDIA Open Model License as well as the original Llama 3.3 license (llama3.3). It is ready for commercial and non-commercial use, with 318,033 downloads in the last month.

## Key points

- FP4 (NVFP4) quantized version of Meta Llama 3.3 70B Instruct by NVIDIA Model Optimizer.
- Reduces bits per parameter from 16 to 4 (~3.3x smaller disk/GPU memory); 41B reported size.
- Third-party model: not owned by NVIDIA; see Meta-Llama-3.3-70B-Instruct card.
- Context up to 128K tokens; text in/text out.
- Runs on TensorRT-LLM; Blackwell hardware (tested B200); calibration on cnn_dailymail.
- Accuracy: MMLU 81.1, GSM8K 92.6, ARC Challenge 93.3, IFEVAL 92.0 (FP4).
- License: NVIDIA Open Model License + llama3.3; 318,033 downloads last month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | meta-llama/Llama-3.3-70B-Instruct |
| Parameters | 70B (41B reported quantized size) |
| Architecture | Transformers (Llama 3.3) |
| Precision | NVFP4 (FP4 weights/activations) |
| Context length | Up to 128K tokens |
| Runtime | TensorRT-LLM |
| Hardware | NVIDIA Blackwell (B200) |
| Calibration | cnn_dailymail |
| MMLU / GSM8K CoT / ARC / IFEVAL | 81.1 / 92.6 / 93.3 / 92.0 (FP4) |
| License | NVIDIA Open Model License + llama3.3 |
| Quantization tool | TensorRT Model Optimizer v0.23.0 |

## Why this source matters for the RAG

This card documents a widely used, memory-optimized NVFP4 deployment of Llama 3.3 70B, showing the concrete accuracy/size trade-off of 4-bit quantization. It is a practical reference for serving a mainstream 70B model on Blackwell-class hardware with TensorRT-LLM.
