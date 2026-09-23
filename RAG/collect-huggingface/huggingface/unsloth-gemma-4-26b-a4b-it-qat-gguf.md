---
id: collect-huggingface/huggingface/unsloth-gemma-4-26b-a4b-it-qat-gguf
title: "gemma-4-26B-A4B-it-qat-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "apache", "attention", "benchmark", "benchmarks", "inference", "license", "llama", "llama.cpp", "memory", "moe", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gemma-4-26B-A4B-it-qat-GGUF.md
source_anchor: ""
source_lines: [1, 53]
sha256: 459e22946852e6291d2eebe5000a454e56e24b637f76bf5a646d46ddd8f67996
---

# gemma-4-26B-A4B-it-qat-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gemma-4-26B-A4B-it-qat-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF repository for Google DeepMind's Gemma 4 26B A4B instruction-tuned model optimized with Quantization-Aware Training (QAT). QAT preserves quality close to bfloat16 while dramatically reducing memory requirements to load the model. Four QAT checkpoint families are described: unquantized QAT checkpoints (Q4_0, half-precision weights extracted from the QAT pipeline, for custom compilation/research, available for E2B, E4B, 12B, 26B A4B, 31B and their drafters); GGUF (Q4_0, ready-to-deploy formats for broad ecosystem compatibility); mobile-optimized wNa8o8 (a custom schema with targeted 2-bit decoding layers, optimized KV caches, and static activations for VRAM savings, for E2B/E4B); and compressed tensors w4a16 (QAT checkpoints serialized for optimized vLLM inference, for E2B/E4B/12B/31B). The underlying Gemma 4 26B A4B is a Mixture-of-Experts model with 25.2B total parameters and 3.8B active parameters, 30 layers, a 1024-token sliding window, a 256K-token context, a 262K vocabulary, and 8 active / 128 total experts plus 1 shared expert; it supports text and image (no audio). It uses hybrid attention (local sliding window + full global, final layer global) with unified K/V and p-RoPE. Gemma 4 adds configurable thinking modes, native system prompts, variable image resolution, and native function calling. Benchmark highlights for 26B A4B: MMLU Pro 82.6%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, Codeforces ELO 1718, GPQA Diamond 82.3%, Tau2 68.2%, MMMLU 86.3%, MMMU Pro 73.8%, MATH-Vision 82.4%, MRCR v2 128k 44.1%. The repository ships a Multi-Token Prediction (MTP) drafter at the repo root (`mtp-gemma-4-26B-A4B-it.gguf`, a near-lossless smart Q4_0) for speculative decoding; recent llama.cpp auto-discovers it. Main quant is UD-Q4_K_XL (14.2GB); MTP files: Q4_0 252MB, Q8_0 462MB, BF16 855MB. Run with `--spec-type draft-mtp --spec-draft-n-max 4`. License is Apache 2.0; 630,941 monthly downloads.

## Key points

- Unsloth GGUF of Gemma 4 26B A4B instruction model optimized with Quantization-Aware Training (QAT).
- QAT preserves near-bfloat16 quality at dramatically lower memory; main quant UD-Q4_K_XL (14.2GB).
- MoE architecture: 25.2B total / 3.8B active parameters, 8 active / 128 total experts (+1 shared).
- 256K context, 1024-token sliding window, hybrid attention with p-RoPE; text + image modalities.
- Ships an MTP (Multi-Token Prediction) drafter for speculative decoding with no output change.
- Four QAT formats: unquantized Q4_0, GGUF Q4_0, mobile wNa8o8, and vLLM compressed-tensors w4a16.
- Benchmarks: MMLU Pro 82.6%, GPQA-D 82.3%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, MMMU Pro 73.8%.
- Apache 2.0 license; runs via llama.cpp, vLLM, SGLang, Ollama, LM Studio.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | gemma-4-26B-A4B-it-qat-GGUF |
| Base model | google/gemma-4-26B-A4B-it (QAT) |
| Architecture | Gemma 4 MoE, hybrid attention |
| Total params | 25.2B |
| Active params | 3.8B |
| Experts | 8 active / 128 total + 1 shared |
| Layers | 30 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image |
| Quantization | QAT GGUF (Q4_0 based) |
| Main quant | UD-Q4_K_XL (14.2GB) |
| MTP drafter | Q4_0 252MB / Q8_0 462MB / BF16 855MB |
| Speculative decoding | `--spec-type draft-mtp --spec-draft-n-max 4` |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 82.6%; GPQA-D 82.3%; AIME 2026 88.3%; LiveCodeBench v6 77.1% |
| Downloads/month | 630,941 |

## Why this source matters for the RAG

This card documents Quantization-Aware Training deployment of Gemma 4 26B A4B via Unsloth GGUFs, including MTP speculative decoding and multiple QAT formats. It is a key reference for QAT, efficient MoE inference, and speculative decoding.
