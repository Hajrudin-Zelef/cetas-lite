---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4
title: "NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-01-28", "2026-09-23"]
keywords: ["nvfp4", "nvidia", "attention", "benchmark", "benchmarks", "distillation", "fp8", "kv cache", "license", "moe", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4.md
source_anchor: ""
source_lines: [1, 46]
sha256: 79cfddc85a2328c4c1153e8e11ea7f99eec4f54ed4752bc8be947ac42cef24a6
---

# NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 is the NVFP4-quantized version of Nemotron-Nano-3-30B-A3B, a large language model trained from scratch by NVIDIA as a unified model for reasoning and non-reasoning tasks. It uses a hybrid Mixture-of-Experts (MoE) architecture with 23 Mamba-2 layers, 23 MoE layers and 6 Attention (grouped query attention, 2 groups) layers, 52 layers total. Each MoE layer has 128 routed experts plus 1 shared expert, with 6 experts activated per token. The model has 30B total parameters and 3.5B active parameters. It supports up to 1M tokens of context (default HF config 256K), and supports English, German, Spanish, French, Italian and Japanese. Reasoning is configurable via the chat template (`enable_thinking`), defaulting to on. The model was pre-trained on 25T tokens and post-trained with SFT and synchronous GRPO RL. This NVFP4 variant was produced with post-training quantization and KV cache quantized to FP8, using a selective strategy that keeps attention layers and the Mamba layers feeding them in BF16, followed by Quantization-Aware Distillation (QAD) for accuracy recovery. Benchmarks (NVFP4 vs BF16 vs FP8): MMLU-Pro 77.4 / 78.3 / 78.1, AIME25 86.7 / 89.1 / 87.7, GPQA 71.9 / 73.0 / 72.5, LiveCodeBench v6 65.4 / 68.3 / 67.6, TauBench V2 average 45.6 / 49.0 / 47.0, IFBench 70.7 / 71.5 / 72.2, MMLU-ProX 57.8 / 59.50 / 59.6. It is governed by the NVIDIA Nemotron Open Model License and is ready for commercial use. Released January 28, 2026. Run via Transformers (4.57.3), vLLM (>=0.12.0), TensorRT-LLM and SGLang.

## Key points

- NVFP4-quantized 30B MoE reasoning model: 30B total / 3.5B active parameters.
- Hybrid Mamba2-Transformer MoE: 23 Mamba-2 + 23 MoE + 6 Attention layers, 52 total.
- 128 routed experts + 1 shared per MoE layer, 6 active per token.
- Context up to 1M tokens (default 256K); six supported languages.
- Selective PTQ keeps attention/Mamba layers in BF16, then QAD accuracy recovery; FP8 KV cache.
- Runtimes: Transformers, vLLM, TensorRT-LLM, SGLang; NVIDIA Nemotron Open Model License.
- Benchmarks: MMLU-Pro 77.4, AIME25 86.7, GPQA 71.9, LiveCodeBench v6 65.4.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 30B |
| Active parameters | 3.5B |
| Architecture | Mamba2-Transformer Hybrid MoE (NemotronH) |
| Layers | 52 (23 Mamba-2, 23 MoE, 6 Attention) |
| Experts | 128 routed + 1 shared; 6 active per token |
| Context length | Up to 1M tokens (default 256K) |
| Languages | English, German, Spanish, French, Italian, Japanese |
| Precision | NVFP4 (KV cache FP8) |
| License | NVIDIA Nemotron Open Model License |
| MMLU-Pro | 77.4 |
| AIME25 (no tools) | 86.7 |
| GPQA (no tools) | 71.9 |
| Release date | 2026-01-28 |

## Why this source matters for the RAG

This card documents NVIDIA's efficient NVFP4-quantized Nano model and the quantization-aware distillation pipeline used to recover accuracy, which is directly relevant to deployment efficiency discussions. It provides structured benchmark comparisons across BF16, FP8 and NVFP4 precisions.
