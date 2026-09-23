---
id: collect-huggingface/huggingface/unsloth-gemma-4-31b-it-qat-gguf
title: "gemma-4-31B-it-qat-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "apache", "attention", "benchmark", "benchmarks", "consumer", "context window", "inference", "license", "llama", "llama.cpp", "memory"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gemma-4-31B-it-qat-GGUF.md
source_anchor: ""
source_lines: [1, 52]
sha256: 6a27e4001f6f5d501047198dafc118b917d62ef8198ca51e7b0b2221c8450a73
---

# gemma-4-31B-it-qat-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gemma-4-31B-it-qat-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF repository for Google DeepMind's Gemma 4 31B instruction-tuned model optimized with Quantization-Aware Training (QAT). QAT preserves quality close to bfloat16 while dramatically reducing the memory required to load the model. Four QAT checkpoint families are described: unquantized QAT checkpoints (Q4_0, half-precision weights extracted from the QAT pipeline, for custom compilation/research, available for E2B, E4B, 12B, 26B A4B, 31B and their drafters); GGUF (Q4_0, ready-to-deploy for broad ecosystem compatibility); mobile-optimized wNa8o8 (targeted 2-bit decoding layers, optimized KV caches, static activations for VRAM savings, for E2B/E4B); and compressed tensors w4a16 (for optimized vLLM inference, for E2B/E4B/12B/31B). The underlying Gemma 4 31B is the largest dense model in the family, with 30.7B total parameters, 60 layers, a 1024-token sliding window, a 256K-token context window, a 262K vocabulary, and a ~550M-parameter vision encoder; it supports text and image (no audio). It uses hybrid attention interleaving local sliding-window attention with full global attention (final layer global), with unified Keys/Values and Proportional RoPE (p-RoPE) on global layers. Gemma 4 adds configurable thinking modes (via `<|think|>` in the system prompt), native system-prompt support, variable image resolution (token budgets 70/140/280/560/1120), video-as-frames, and native function calling. Benchmark highlights for 31B: MMLU Pro 85.2%, AIME 2026 89.2%, LiveCodeBench v6 80.0%, Codeforces ELO 2150, GPQA Diamond 84.3%, Tau2 76.9%, HLE no tools 19.5% / with search 26.5%, BigBench Extra Hard 74.4%, MMMLU 88.4%, MMMU Pro 76.9%, MATH-Vision 85.6%, and MRCR v2 8-needle 128k 66.4%. The repository ships a Multi-Token Prediction (MTP) drafter at the repo root (`mtp-gemma-4-31B-it.gguf`, a near-lossless smart Q4_0) for speculative decoding, auto-discovered by recent llama.cpp. The main quant is UD-Q4_K_XL (17.3GB); MTP files: Q4_0 280MB, Q8_0 515MB, BF16 955MB. Run with `--spec-type draft-mtp --spec-draft-n-max 4`. License is Apache 2.0; 576,727 monthly downloads.

## Key points

- Unsloth QAT GGUF of Gemma 4 31B, the family's largest dense model (30.7B parameters).
- Quantization-Aware Training preserves near-bfloat16 quality at dramatically lower memory; main quant UD-Q4_K_XL (17.3GB).
- 256K context, 1024-token sliding window, hybrid attention with p-RoPE; text + image modalities.
- Ships an MTP (Multi-Token Prediction) drafter for speculative decoding without changing output.
- Four QAT formats: unquantized Q4_0, GGUF Q4_0, mobile wNa8o8, and vLLM compressed-tensors w4a16.
- Top benchmarks: MMLU Pro 85.2%, GPQA-D 84.3%, AIME 2026 89.2%, LiveCodeBench v6 80.0%, Codeforces ELO 2150.
- Configurable thinking mode, native system prompts, variable image resolution, native function calling.
- Apache 2.0 license; runs via llama.cpp, vLLM, SGLang, Ollama, LM Studio.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | gemma-4-31B-it-qat-GGUF |
| Base model | google/gemma-4-31B-it (QAT) |
| Architecture | Gemma 4 dense, hybrid attention |
| Total params | 30.7B |
| Layers | 60 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image |
| Quantization | QAT GGUF (Q4_0 based) |
| Main quant | UD-Q4_K_XL (17.3GB) |
| MTP drafter | Q4_0 280MB / Q8_0 515MB / BF16 955MB |
| Speculative decoding | `--spec-type draft-mtp --spec-draft-n-max 4` |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 85.2%; GPQA-D 84.3%; AIME 2026 89.2%; LiveCodeBench v6 80.0% |
| Downloads/month | 576,727 |

## Why this source matters for the RAG

This card documents Quantization-Aware Training deployment of Gemma 4 31B via Unsloth GGUFs, including MTP speculative decoding and multiple QAT formats. It is a key reference for QAT, efficient dense-model inference, and speculative decoding on consumer hardware.
