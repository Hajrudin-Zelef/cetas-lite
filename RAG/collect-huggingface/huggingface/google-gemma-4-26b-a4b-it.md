---
id: collect-huggingface/huggingface/google-gemma-4-26b-a4b-it
title: "gemma-4-26B-A4B-it - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face"]
dates: ["2025-01", "2026-09-23"]
keywords: ["apache", "attention", "benchmark", "benchmarks", "context window", "inference", "license", "memory", "moe", "multimodal", "open weights", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/google-gemma-4-26B-A4B-it.md
source_anchor: ""
source_lines: [1, 54]
sha256: d00133cd7097b19acd7f8ad9c7d61e1e80161577182fd1db1b1f83b7d770e17e
---

# gemma-4-26B-A4B-it - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google/gemma-4-26B-A4B-it
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Gemma 4 is a family of open models built by Google DeepMind, and this card covers the 26B A4B instruction-tuned variant. Gemma 4 models are multimodal, handling text and image input and generating text output (audio is supported on E2B, E4B, and 12B). The release includes open weights in both pre-trained and instruction-tuned variants, features a context window of up to 256K tokens, and supports over 140 languages (35+ out of the box). The family ships in five sizes: E2B, E4B, 12B, 26B A4B, and 31B, spanning dense and Mixture-of-Experts (MoE) architectures. The 26B A4B MoE has 25.2B total parameters with 3.8B active parameters, 30 layers, a 1024-token sliding window, a 262K vocabulary, and 8 active / 128 total experts plus 1 shared expert; it supports text and image (no audio). Its hybrid attention interleaves local sliding-window attention with full global attention (final layer global), with unified K/V and Proportional RoPE (p-RoPE) on global layers for long-context memory efficiency. Gemma 4 adds configurable thinking modes (enabled via `<|think|>` in the system prompt), native system-prompt support, variable image resolution (token budgets 70/140/280/560/1120), video as frames, and native function calling. Benchmark highlights for 26B A4B: MMLU Pro 82.6%, AIME 2026 (no tools) 88.3%, LiveCodeBench v6 77.1%, Codeforces ELO 1718, GPQA Diamond 82.3%, Tau2 68.2%, HLE no tools 8.7%, MMMLU 86.3%, MMMU Pro 73.8%, MATH-Vision 82.4%, MRCR v2 128k 44.1%. Recommended sampling: temperature 1.0, top_p 0.95, top_k 64. Training data has a January 2025 cutoff. License is Apache 2.0, and the model reports 10,486,181 monthly downloads. Technical report: arXiv 2607.02770.

## Key points

- Gemma 4 26B A4B is a multimodal MoE: 25.2B total / 3.8B active parameters, 8 active / 128 total experts (+1 shared).
- 256K-token context; hybrid local sliding-window + global attention; p-RoPE and unified K/V.
- Handles text and image input (no audio, unlike E2B/E4B/12B); 262K vocabulary.
- Configurable thinking mode (`<|think|>`), native system-prompt support, and native function calling.
- Variable image resolution via token budgets 70/140/280/560/1120; video processed as frames.
- Strong benchmarks: MMLU Pro 82.6%, GPQA-D 82.3%, AIME 2026 88.3%, LiveCodeBench v6 77.1%, MMMU Pro 73.8%.
- Apache 2.0 license; January 2025 data cutoff; arXiv 2607.02770.
- Recommended sampling: temperature 1.0, top_p 0.95, top_k 64.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google |
| Model name | gemma-4-26B-A4B-it |
| Architecture | Hybrid-attention MoE (Gemma 4) |
| Total params | 25.2B |
| Active params | 3.8B |
| Experts | 8 active / 128 total + 1 shared |
| Layers | 30 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image |
| Vision encoder | ~550M params |
| Thinking mode | Configurable via `<|think|>` |
| Image token budgets | 70 / 140 / 280 / 560 / 1120 |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| Data cutoff | January 2025 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 82.6%; GPQA-D 82.3%; AIME 2026 88.3%; LiveCodeBench v6 77.1%; MMMU Pro 73.8% |
| Downloads/month | 10,486,181 |
| arXiv | 2607.02770 |

## Why this source matters for the RAG

This is the core model card for Google's Gemma 4 26B A4B, a state-of-the-art open multimodal MoE with hybrid attention, thinking modes, and Apache 2.0 licensing. It is a central reference for efficient MoE inference, long-context multimodal modeling, and open-weight frontier models.
