---
id: collect-huggingface/huggingface/google-gemma-4-26b-a4b-it-qat-q4-0-gguf
title: "gemma-4-26B-A4B-it-qat-q4_0-gguf - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "apache", "attention", "benchmark", "benchmarks", "context window", "inference", "license", "llama", "llama.cpp", "memory", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/google-gemma-4-26B-A4B-it-qat-q4_0-gguf.md
source_anchor: ""
source_lines: [1, 52]
sha256: bdb533c9f3acae1c694953f7aea449a95d0b2e76de91daa16a1fa6f2938aee3d
---

# gemma-4-26B-A4B-it-qat-q4_0-gguf - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google/gemma-4-26B-A4B-it-qat-q4_0-gguf
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Google DeepMind's official GGUF release of the Gemma 4 26B A4B instruction-tuned model optimized with Quantization-Aware Training (QAT). QAT preserves quality close to bfloat16 while dramatically reducing the memory required to load the model. The card describes four QAT checkpoint families: unquantized QAT checkpoints (Q4_0, half-precision weights extracted from the QAT pipeline for custom downstream compilation and research, available for E2B, E4B, 12B, 26B A4B, 31B and their drafter models); GGUF (Q4_0, ready-to-deploy formats for broad ecosystem compatibility); mobile-optimized wNa8o8 (a custom schema with targeted 2-bit decoding layers, optimized KV caches, and static activations for VRAM savings, for E2B/E4B); and compressed tensors w4a16 (serialized for optimized vLLM inference, for E2B/E4B/12B/31B). It also notes an Assistant Compatibility requirement: when using multi-token prediction (speculative decoding) with an assistant model alongside a QAT target, the assistant must also be a QAT checkpoint with the same precision. The underlying Gemma 4 26B A4B is a Mixture-of-Experts model with 25.2B total parameters and 3.8B active parameters, 30 layers, a 1024-token sliding window, a 256K-token context window, a 262K vocabulary, and 8 active / 128 total experts plus 1 shared expert; it supports text and image, with a ~550M-parameter vision encoder. It uses hybrid attention (local sliding window + full global, final layer global) with unified K/V and Proportional RoPE (p-RoPE). Gemma 4 adds configurable thinking modes (`<|think|>`), native system prompts, variable image resolution (token budgets 70/140/280/560/1120), video-as-frames, and native function calling. Benchmark highlights: MMLU Pro 82.6%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, Codeforces ELO 1718, GPQA Diamond 82.3%, Tau2 68.2%, MMMLU 86.3%, MMMU Pro 73.8%, MATH-Vision 82.4%, and MRCR v2 128k 44.1%. This repository provides a single 4-bit Q4_0 GGUF (14.4 GB); recommended sampling is temperature 1.0, top_p 0.95, top_k 64. License is Apache 2.0; 557,225 monthly downloads. Technical report: arXiv 2607.02770.

## Key points

- Official Google QAT GGUF of Gemma 4 26B A4B instruction model, 4-bit Q4_0 (14.4 GB).
- QAT preserves near-bfloat16 quality at dramatically reduced memory.
- MoE architecture: 25.2B total / 3.8B active parameters, 8 active / 128 total experts (+1 shared).
- 256K context, 1024-token sliding window, hybrid attention with p-RoPE; text + image.
- Four QAT formats: unquantized Q4_0, GGUF Q4_0, mobile wNa8o8, and vLLM compressed-tensors w4a16.
- Speculative-decoding assistant must be a QAT checkpoint at the same precision for compatibility.
- Benchmarks: MMLU Pro 82.6%, GPQA-D 82.3%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, MMMU Pro 73.8%.
- Apache 2.0 license; runs via llama.cpp, vLLM, SGLang, Ollama, LM Studio.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google |
| Model name | gemma-4-26B-A4B-it-qat-q4_0-gguf |
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
| Quantization | QAT GGUF Q4_0 (14.4 GB) |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 82.6%; GPQA-D 82.3%; AIME 2026 88.3%; LiveCodeBench v6 77.1% |
| Downloads/month | 557,225 |
| arXiv | 2607.02770 |

## Why this source matters for the RAG

This is the official Google QAT GGUF for Gemma 4 26B A4B, establishing the reference 4-bit deployment format and the assistant-compatibility rule for speculative decoding. It is a central reference for QAT, GGUF deployment, and official quantized releases.
