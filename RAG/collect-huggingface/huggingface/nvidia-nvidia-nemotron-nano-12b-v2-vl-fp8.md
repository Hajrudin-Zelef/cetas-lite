---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-nano-12b-v2-vl-fp8
title: "NVIDIA-Nemotron-Nano-12B-v2-VL-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["China", "Hugging Face", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: ["2025-10-28", "2026-09-23"]
keywords: ["fp8", "nvidia", "benchmarks", "embedding", "foundry", "inference", "license", "parameters", "quantization", "reasoning", "tensorrt", "training"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-Nano-12B-v2-VL-FP8.md
source_anchor: ""
source_lines: [1, 45]
sha256: 02f164dfe5fa4170e0d5cfe4a6d5d47c7ce7d265c3cf0b7048dce273ecc76e15
---

# NVIDIA-Nemotron-Nano-12B-v2-VL-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-Nano-12B-v2-VL-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-Nano-VL-12B-V2-FP8 is the FP8-quantized version of the NVIDIA Nemotron Nano VL V2 vision-language model, an autoregressive VLM using an optimized transformer architecture. It combines the C-RADIOv2-H vision encoder with the NVIDIA-Nemotron-Nano-12B-v2 language encoder and supports single-image inference. The model is quantized with TensorRT Model Optimizer and trained on commercial images across all three training stages. It supports image summarization, text-image analysis, OCR, interactive Q&A on images, and text chain-of-thought reasoning, and is aimed at AI foundry enterprise customers. Input is RGB images (up to a 12-tile layout constraint, each tile 512x512, supporting resolutions up to 3072x1024, with no alpha channel) plus text, with a context length up to 128K. Output is text up to 128K. Supported languages include German, Spanish, French, Italian, Korean, Portuguese, Russian, Japanese, Chinese and English. It runs on vLLM (modelopt quantization) on H100 SXM 80GB, with Linux support. Benchmarks (FP8 vs BF16): AI2D 87.6% vs 87.1%, OCRBenchV2 61.8% vs 62.0%, OCRBench 85.4% vs 85.6%, ChartQA 89.4% vs 89.7%, DocVQA val 94.3% vs 94.4%. Training data totaled 39,486,703 samples across 270 datasets (27.7 TB), mixing public, internal, synthetic and crawled text/image/video data. It is governed by the NVIDIA Open License Agreement with backbone NVIDIA-Nemotron-Nano-12B-v2; released October 28, 2025. Reported model size 13B params; 123,428 downloads last month.

## Key points

- FP8-quantized vision-language model: C-RADIOv2-H vision encoder + Nemotron-Nano-12B-v2 LLM.
- Single-image inference; supports summarization, OCR, image Q&A and text CoT reasoning.
- Context up to 128K; 10 supported languages; 12-tile image layout (max 3072x1024).
- Quantized with TensorRT Model Optimizer; runs via vLLM on H100 SXM 80GB.
- Benchmarks (FP8): AI2D 87.6%, DocVQA 94.3%, ChartQA 89.4%, OCRBench 85.4%.
- License: NVIDIA Open Model License; released 2025-10-28; 13B params.

## Technical data / figures

| Attribute | Value |
|---|---|
| Parameters | 13B (reported) |
| Architecture | Transformer VLM: C-RADIOv2-H (vision) + Nemotron-Nano-12B-v2 (language) |
| Context length | Up to 128K |
| Input | Image (RGB, 12-tile max) + Text |
| Image resolution | 512x512 tiles; up to 3072x1024 (6x2 layout) |
| Languages | German, Spanish, French, Italian, Korean, Portuguese, Russian, Japanese, Chinese, English |
| Precision | FP8 (F8_E4M3); BF16 reference |
| Runtime | vLLM (modelopt quantization) |
| Hardware | H100 SXM 80GB |
| AI2D / DocVQA / ChartQA / OCRBench | 87.6% / 94.3% / 89.4% / 85.4% (FP8) |
| Training data | 39.49M samples, 270 datasets, 27.7 TB |
| License | NVIDIA Open Model License |
| Release date | 2025-10-28 |

## Why this source matters for the RAG

This card documents a compact FP8-quantized vision-language model tailored for document understanding and image Q&A, providing concrete quantization impact data (FP8 vs BF16) and image-tiling constraints. It complements the embedding and parse models to cover the full NVIDIA document-processing stack.
