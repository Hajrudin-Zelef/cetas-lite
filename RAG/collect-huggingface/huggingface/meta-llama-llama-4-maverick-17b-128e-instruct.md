---
id: collect-huggingface/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct
title: "Llama-4-Maverick-17B-128E-Instruct - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "SGLang", "vLLM"]
dates: ["2024-08", "2025-04-05", "2026-09-23"]
keywords: ["llama", "attention", "benchmark", "benchmarks", "compute", "fp8", "gpu", "license", "moe", "multimodal", "parameters", "pretraining"]
source: docs/RAG/Collect RAG/03_huggingface/meta-llama-Llama-4-Maverick-17B-128E-Instruct.md
source_anchor: ""
source_lines: [1, 53]
sha256: a1d75ebcbf783a466e578556d405a98ac1a4ae74a220e0ba1cb2ce7c91ce4fee
---

# Llama-4-Maverick-17B-128E-Instruct - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/meta-llama/Llama-4-Maverick-17B-128E-Instruct
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Llama 4 Maverick is the larger of Meta's two initial Llama 4 models, released April 5, 2025. It is a natively multimodal, auto-regressive language model that uses a Mixture-of-Experts (MoE) architecture with early fusion for native multimodality. Maverick is a 17B-parameter model with 128 experts, yielding 17B activated parameters and 400B total parameters. It accepts multilingual text and image input and outputs multilingual text and code, with a context length of 1M tokens and an August 2024 knowledge cutoff, trained on roughly 22 trillion tokens of multimodal data from publicly available, licensed sources and Meta products/services (including public Instagram/Facebook posts and Meta AI interactions). Twelve languages are officially supported (Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, Vietnamese), while pretraining covered 200 languages. This repository holds the BF16 weights; Meta also releases an FP8 variant that fits a single H100 DGX host. The instruction-tuned model targets assistant-like chat and visual reasoning (recognition, image reasoning, captioning, image QA), tested for up to 5 input images. Instruction-tuned benchmarks: MMMU 73.4, MMMU Pro 59.6, MathVista 73.7, ChartQA 90.0, DocVQA 94.4, LiveCodeBench 43.4, MMLU Pro 80.5, GPQA Diamond 69.8, MGSM 92.3, and MTOB long-context chrF 54.0/46.4 (half book) and 50.8/46.7 (full book). Training used 2.38M GPU hours on H100-80GB, emitting ~645 tons CO2eq. Usage requires Transformers v4.51.0+ (`attn_implementation="flex_attention"`, bf16), or vLLM/SGLang. It is governed by the Llama 4 Community License (commercial use; 700M-MAU threshold). The hub reports 402B params (BF16) and 10,689 monthly downloads.

## Key points

- Natively multimodal MoE model: 17B activated / 400B total parameters, 128 experts.
- BF16 checkpoint (FP8 variant also released, fitting one H100 DGX host).
- 1M-token context; ~22T training tokens; August 2024 knowledge cutoff.
- Instruction-tuned for chat, visual reasoning, and image understanding (up to 5 images).
- 12 supported languages; pretrained on 200 languages.
- Strong benchmarks: MMLU Pro 80.5, GPQA Diamond 69.8, MMMU 73.4, LiveCodeBench 43.4, DocVQA 94.4.
- Custom Llama 4 Community License; released April 5, 2025.
- Requires Transformers v4.51.0+ (flex attention, bf16), vLLM, or SGLang.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | meta-llama |
| Model name | Llama-4-Maverick-17B-128E-Instruct |
| Architecture | Auto-regressive MoE, early-fusion multimodal |
| Total params | 400B |
| Active params | 17B |
| Experts | 128 |
| Context length | 1M tokens |
| Training tokens | ~22T |
| Knowledge cutoff | August 2024 |
| Input / output | Multilingual text + image / multilingual text + code |
| Languages | 12 supported (200 in pretraining) |
| Release date | April 5, 2025 |
| Precision | BF16 (FP8 variant also available) |
| Training compute | 2.38M GPU hours (H100-80GB) |
| GHG emissions | 645 tons CO2eq (location-based) |
| License | Llama 4 Community License |
| Key benchmarks | MMLU Pro 80.5; GPQA-D 69.8; MMMU 73.4; LiveCodeBench 43.4; DocVQA 94.4 |
| Downloads/month | 10,689 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This is the canonical model card for Llama 4 Maverick, the BF16 base of the FP8 deployment variant, documenting Meta's largest initial Llama 4 MoE. It anchors retrieval on multimodal MoE architectures, million-token context, and Meta's benchmark suite.
