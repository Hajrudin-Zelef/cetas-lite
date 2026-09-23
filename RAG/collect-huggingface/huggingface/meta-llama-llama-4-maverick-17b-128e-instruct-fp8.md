---
id: collect-huggingface/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct-fp8
title: "Llama-4-Maverick-17B-128E-Instruct-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "Microsoft", "SGLang", "vLLM"]
dates: ["2024-08", "2025-04-05", "2026-09-23"]
keywords: ["fp8", "llama", "benchmark", "benchmarks", "compute", "gpu", "inference", "license", "moe", "multimodal", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/meta-llama-Llama-4-Maverick-17B-128E-Instruct-FP8.md
source_anchor: ""
source_lines: [1, 54]
sha256: 5be25f26f721c561086e29b4a15ec24b7a3d8c40731643dbbe8bd5df8f0ed680
---

# Llama-4-Maverick-17B-128E-Instruct-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Llama 4 Maverick is the larger of the two initial Llama 4 models from Meta, released April 5, 2025. Like Scout, it is a natively multimodal, auto-regressive language model using a Mixture-of-Experts (MoE) architecture with early fusion for native multimodality. Maverick is a 17B-parameter model with 128 experts, giving 17B activated parameters and 400B total parameters. It handles multilingual text and image input and outputs multilingual text and code, with a context length of 1M tokens and a knowledge cutoff of August 2024, trained on roughly 22 trillion tokens of multimodal data. The same 12 languages are supported (Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, Vietnamese), with pretraining covering 200 languages. This specific repository is the FP8-quantized variant: Meta releases Maverick in both BF16 and FP8, and the FP8 weights fit on a single H100 DGX host while maintaining quality. Reported evaluations were run on BF16 models. The card documents instruction-tuned benchmark scores: MMMU 73.4, MMMU Pro 59.6, MathVista 73.7, ChartQA 90.0, DocVQA 94.4, LiveCodeBench 43.4, MMLU Pro 80.5, GPQA Diamond 69.8, MGSM 92.3, and MTOB long-context chrF 54.0/46.4 (half book) and 50.8/46.7 (full book). Training used 2.38M GPU hours on H100-80GB, emitting ~645 tons CO2eq. Usage requires Transformers v4.51.0+ with `tp_plan="auto"`, or vLLM/SGLang. The model is governed by the Llama 4 Community License (commercial use, 700M-MAU threshold). The hub reports 402B params (BF16, F8_E4M3) and 52,188 monthly downloads.

## Key points

- Natively multimodal MoE model: 17B activated / 400B total parameters, 128 experts.
- FP8-quantized checkpoint that fits a single H100 DGX host while preserving quality.
- 1M-token context; ~22T training tokens; August 2024 knowledge cutoff.
- Instruction-tuned for chat, visual reasoning, and image understanding (up to 5 images).
- 12 supported languages; pretrained on 200 languages.
- Strong benchmarks: MMLU Pro 80.5, GPQA Diamond 69.8, MMMU 73.4, LiveCodeBench 43.4, DocVQA 94.4.
- Custom Llama 4 Community License; released April 5, 2025.
- Requires Transformers v4.51.0+ (`tp_plan="auto"`), vLLM, or SGLang.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | meta-llama |
| Model name | Llama-4-Maverick-17B-128E-Instruct-FP8 |
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
| Precision | FP8 (BF16 variant also released) |
| Tensor types | BF16, F8_E4M3 |
| Training compute | 2.38M GPU hours (H100-80GB) |
| GHG emissions | 645 tons CO2eq (location-based) |
| License | Llama 4 Community License |
| Key benchmarks | MMLU Pro 80.5; GPQA-D 69.8; MMMU 73.4; LiveCodeBench 43.4; DocVQA 94.4 |
| Downloads/month | 52,188 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This card documents the FP8 deployment variant of Llama 4 Maverick, a frontier open-weight multimodal MoE, and its single-host quantization strategy. It is a central reference for large MoE inference, FP8 quantization, and Meta's multimodal benchmark suite.
