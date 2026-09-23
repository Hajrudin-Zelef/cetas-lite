---
id: collect-huggingface/huggingface/meta-llama-llama-4-scout-17b-16e-instruct
title: "Llama-4-Scout-17B-16E-Instruct - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "Microsoft", "SGLang", "vLLM"]
dates: ["2024-08", "2025-04-05", "2026-09-23"]
keywords: ["llama", "scout", "benchmark", "benchmarks", "compute", "context window", "gguf", "gpu", "int4", "license", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/meta-llama-Llama-4-Scout-17B-16E-Instruct.md
source_anchor: ""
source_lines: [1, 53]
sha256: 50d92311f1b0ae956bb23d8f407d969fb057fc94d4b82bbb7fb7b50385b644f5
---

# Llama-4-Scout-17B-16E-Instruct - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/meta-llama/Llama-4-Scout-17B-16E-Instruct
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Llama 4 Scout is one of two initial models in Meta's Llama 4 collection, a family of natively multimodal, auto-regressive language models that use a Mixture-of-Experts (MoE) architecture with early fusion for multimodality. Scout is a 17B-parameter model with 16 experts, featuring 17B activated parameters and 109B total parameters. It accepts multilingual text and image input and produces multilingual text and code output, with a context length of up to 10M tokens and a knowledge cutoff of August 2024. It was pretrained on roughly 40 trillion tokens drawn from publicly available, licensed data and information from Meta's products and services (including public Instagram/Facebook posts and Meta AI interactions). Twelve languages are officially supported: Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, and Vietnamese, though pretraining covered 200 languages. The instruction-tuned model is designed for assistant-like chat and visual reasoning, including visual recognition, image reasoning, captioning, and image question answering, and is tested for up to 5 input images. Released April 5, 2025, it uses the custom Llama 4 Community License (commercial use permitted, with a 700M-MAU threshold for additional licensing). Scout is released as BF16 weights and can fit within a single H100 GPU using on-the-fly int4 quantization; reported evaluations were conducted on BF16. Training used 5.0M GPU hours on H100-80GB hardware, emitting an estimated 1,354 tons CO2eq. Instruction-tuned benchmark highlights: MMMU 69.4, MMMU Pro 52.2, MathVista 70.7, ChartQA 88.8, DocVQA 94.4, LiveCodeBench 32.8, MMLU Pro 74.3, GPQA Diamond 57.2, MGSM 90.6, and MTOB long-context chrF 54.0/46.4 (half book). It runs via Transformers ≥4.51.0, vLLM, SGLang, and quantized GGUF builds.

## Key points

- Natively multimodal MoE model: 17B activated / 109B total parameters, 16 experts.
- Very large 10M-token context window; ~40T training tokens; August 2024 knowledge cutoff.
- Instruction-tuned for chat, visual reasoning, captioning, and image QA (tested up to 5 images).
- 12 officially supported languages; pretrained on 200 languages.
- Custom Llama 4 Community License; release date April 5, 2025.
- Released as BF16; fits a single H100 via on-the-fly int4 quantization.
- Strong instruction benchmarks: MMMU Pro 52.2, GPQA Diamond 57.2, MMLU Pro 74.3, LiveCodeBench 32.8.
- Runs with Transformers ≥4.51.0, vLLM, SGLang, and GGUF quantizations.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | meta-llama |
| Model name | Llama-4-Scout-17B-16E-Instruct |
| Architecture | Auto-regressive MoE, early-fusion multimodal |
| Total params | 109B |
| Active params | 17B |
| Experts | 16 |
| Context length | 10M tokens |
| Training tokens | ~40T |
| Knowledge cutoff | August 2024 |
| Input / output | Multilingual text + image / multilingual text + code |
| Languages | 12 supported (200 in pretraining) |
| Release date | April 5, 2025 |
| Precision | BF16 (on-the-fly int4 quantization supported) |
| Training compute | 5.0M GPU hours (H100-80GB) |
| GHG emissions | 1,354 tons CO2eq (location-based) |
| License | Llama 4 Community License |
| Key benchmarks | MMLU 79.6; MMLU Pro 74.3; GPQA-D 57.2; MMMU 69.4; DocVQA 94.4; LiveCodeBench 32.8 |
| Downloads/month | 162,876 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This is the primary model card for Llama 4 Scout, documenting a frontier open-weight multimodal MoE with a 10M-token context and the Llama 4 Community License. It anchors retrieval on large-context open models, MoE efficiency, and Meta's multimodal benchmarks.
