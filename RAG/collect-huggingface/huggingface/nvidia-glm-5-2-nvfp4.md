---
id: collect-huggingface/huggingface/nvidia-glm-5-2-nvfp4
title: "GLM-5.2-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp4", "glm", "nvfp4", "agent", "attention", "benchmarks", "blackwell", "cost", "fp8", "inference", "license", "mit license"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-GLM-5.2-NVFP4.md
source_anchor: ""
source_lines: [1, 53]
sha256: 298042e377084a21a680e074e674d4849c2fa503e95c280dbc1ee492a53c26fb
---

# GLM-5.2-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/GLM-5.2-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

nvidia/GLM-5.2-NVFP4 is NVIDIA's NVFP4-quantized version of Z.ai's GLM-5.2, an auto-regressive Mixture-of-Experts language model for reasoning and coding that uses sparse attention (with an IndexShare indexer) to support a long context. NVIDIA quantized it with Model Optimizer (nvidia-modelopt v0.46.0). Reported model size is 381B params, architecture `GlmMoeDsaForCausalLM`; the base model has 753B total / 40B activated parameters. License is MIT, same as the base model, ready for commercial/non-commercial use.

The NVFP4 quantization covers only the weights and activations of the linear operators within transformer blocks in MoE experts; the shared expert is not quantized. It is ready for inference with SGLang and vLLM on NVIDIA Blackwell (B200/B300) on Linux. Input is text with context length up to 1M. It is intended for AI agent systems, chatbots, RAG systems, and other AI applications.

Evaluation (baseline = GLM-5.2-FP8): GPQA Diamond 89.39 vs 89.52, SciCode 49.04 vs 49.85, IFBench 75.81 vs 74.95, AA-LCR 70.13 vs 69.38, tau2-Bench Telecom 98.25 vs 97.9. NVFP4 is essentially at parity with FP8 and even slightly ahead on IFBench, AA-LCR, and tau2-Bench. Benchmarked at temperature=1.0, top_p=0.95 (GPQA used max_new_tokens=100000; others 64000). AA-LCR was measured with SGLang; other benchmarks with vLLM.

Usage examples: SGLang needs `transformers>=5.3.0` (`--quantization modelopt_fp4 --tool-call-parser glm47 --reasoning-parser glm45`); vLLM uses the `vllm/vllm-openai:v0.23.0` image with `--enable-expert-parallel --kv-cache-dtype fp8_e4m3`. The card includes full Model Card++ subcards (explainability, bias, safety & security, privacy). ~712K monthly downloads.

## Key points

- NVIDIA NVFP4 4-bit quant of GLM-5.2 (753B total / 40B active; 381B on disk).
- Quantized with nvidia-modelopt v0.46.0; shared expert kept unquantized.
- 1M-token context; sparse attention with IndexShare indexer.
- At parity with FP8: GPQA 89.39, SciCode 49.04, IFBench 75.81, AA-LCR 70.13, tau2-Bench 98.25.
- MIT license; runs on NVIDIA Blackwell (B200/B300) via SGLang and vLLM.
- Includes Model Card++ subcards for explainability, bias, safety, privacy.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 753B (base) |
| Active parameters | 40B |
| Disk size | 381B params (reported) |
| Architecture | GlmMoeDsaForCausalLM (MoE + sparse attention, IndexShare) |
| Quantization | NVFP4 (MoE expert linear ops only) |
| Quantization tool | nvidia-modelopt v0.46.0 |
| Context length | up to 1M |
| License | MIT |
| Runtime | SGLang, vLLM (NVIDIA Blackwell, Linux) |
| GPQA Diamond (FP8 / NVFP4) | 89.52 / 89.39 |
| SciCode (FP8 / NVFP4) | 49.85 / 49.04 |
| IFBench (FP8 / NVFP4) | 74.95 / 75.81 |
| AA-LCR (FP8 / NVFP4) | 69.38 / 70.13 |
| tau2-Bench Telecom (FP8 / NVFP4) | 97.9 / 98.25 |
| Monthly downloads | ~712K |

## Why this source matters for the RAG

This card documents enterprise-grade 4-bit NVFP4 quantization of Z.ai's 1M-context MoE flagship, with explicit FP8-vs-NVFP4 accuracy deltas showing near-lossless compression. It provides concrete serving commands for SGLang and vLLM on Blackwell hardware, directly relevant to cost and VRAM optimization. It complements the GLM-5.2-FP8 and GLM-5.3 entries in the knowledge base.
