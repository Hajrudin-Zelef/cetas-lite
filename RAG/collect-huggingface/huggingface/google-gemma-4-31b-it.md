---
id: collect-huggingface/huggingface/google-gemma-4-31b-it
title: "gemma-4-31B-it - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Google", "Hugging Face", "SGLang", "vLLM"]
dates: ["2025-01", "2026-09-23"]
keywords: ["apache", "attention", "benchmark", "benchmarks", "context window", "gguf", "license", "memory", "multimodal", "open weights", "parameters", "reasoning"]
source: docs/RAG/Collect RAG/03_huggingface/google-gemma-4-31B-it.md
source_anchor: ""
source_lines: [1, 52]
sha256: b00cff775ad5827775e3b69fd24af5f367ffb9ce6d83ff31b8c5b2d949fd5bd5
---

# gemma-4-31B-it - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/google/gemma-4-31B-it
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Gemma 4 31B is the largest dense model in Google DeepMind's Gemma 4 family of open multimodal models. Gemma 4 handles text and image input and generates text output (audio is supported only on the E2B, E4B, and 12B variants). The family includes both pre-trained and instruction-tuned open weights, supports a context window of up to 256K tokens, and maintains multilingual support across over 140 languages (35+ out of the box). Five sizes are offered—E2B, E4B, 12B, 26B A4B, and 31B—spanning dense and Mixture-of-Experts architectures. The 31B dense model has 30.7B total parameters, 60 layers, a 1024-token sliding window, a 262K vocabulary, a ~550M-parameter vision encoder, and supports text plus image (no audio). It uses a hybrid attention mechanism that interleaves local sliding-window attention with full global attention, ensuring the final layer is global; global layers use unified Keys/Values and Proportional RoPE (p-RoPE) to optimize memory for long contexts. Gemma 4 adds configurable thinking modes (triggered by `<|think|>` at the start of the system prompt), native system-prompt support, variable image resolution via token budgets (70/140/280/560/1120), video-as-frames processing, and native function calling. Benchmark results for the 31B (instruction-tuned): MMLU Pro 85.2%, AIME 2026 (no tools) 89.2%, LiveCodeBench v6 80.0%, Codeforces ELO 2150, GPQA Diamond 84.3%, Tau2 76.9%, HLE no tools 19.5% / with search 26.5%, BigBench Extra Hard 74.4%, MMMLU 88.4%, MMMU Pro 76.9%, MATH-Vision 85.6%, and MRCR v2 8-needle 128k 66.4%. Recommended sampling: temperature 1.0, top_p 0.95, top_k 64. Training data has a January 2025 cutoff. License is Apache 2.0; the model reports 9,199,819 monthly downloads. Technical report: arXiv 2607.02770.

## Key points

- Gemma 4 31B is the largest dense variant: 30.7B parameters, 60 layers, 256K context.
- Multimodal text + image (no audio); ~550M vision encoder; 262K vocabulary.
- Hybrid attention: interleaved local sliding-window (1024) and full global attention, with unified K/V and p-RoPE.
- Configurable thinking mode, native system prompts, variable image resolution, video, and native function calling.
- Top-tier benchmarks: MMLU Pro 85.2%, GPQA-D 84.3%, AIME 2026 89.2%, LiveCodeBench v6 80.0%, Codeforces ELO 2150, MMMU Pro 76.9%.
- Recommended sampling: temperature 1.0, top_p 0.95, top_k 64; January 2025 data cutoff.
- Apache 2.0 license; part of the five-size Gemma 4 family (E2B to 31B).
- Runs via Transformers, vLLM, SGLang, and GGUF quantizations.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | google |
| Model name | gemma-4-31B-it |
| Architecture | Dense hybrid-attention transformer (Gemma 4) |
| Total params | 30.7B |
| Layers | 60 |
| Sliding window | 1024 tokens |
| Context length | 256K tokens |
| Vocabulary | 262K |
| Modalities | Text, Image (no audio) |
| Vision encoder | ~550M params |
| Thinking mode | Configurable via `<|think|>` |
| Image token budgets | 70 / 140 / 280 / 560 / 1120 |
| Sampling | temperature 1.0, top_p 0.95, top_k 64 |
| Data cutoff | January 2025 |
| License | Apache 2.0 |
| Key benchmarks | MMLU Pro 85.2%; GPQA-D 84.3%; AIME 2026 89.2%; LiveCodeBench v6 80.0%; MMMU Pro 76.9% |
| Downloads/month | 9,199,819 |
| arXiv | 2607.02770 |

## Why this source matters for the RAG

This card documents the flagship dense model of Google's Gemma 4 family, with frontier-level reasoning and coding scores under a permissive Apache 2.0 license. It is a central reference for open multimodal models, hybrid attention, and thinking-mode configuration.
