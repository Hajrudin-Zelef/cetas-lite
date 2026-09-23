---
id: collect-huggingface/huggingface/unsloth-llama-3-2-3b-instruct-bnb-4bit
title: "Llama-3.2-3B-Instruct-bnb-4bit - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "SGLang", "Unsloth", "vLLM"]
dates: ["2023-12", "2024-09-25", "2026-09-23"]
keywords: ["llama", "agentic", "alignment", "attention", "benchmark", "benchmarks", "fine-tuning", "gguf", "gqa", "inference", "license", "memory"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Llama-3.2-3B-Instruct-bnb-4bit.md
source_anchor: ""
source_lines: [1, 49]
sha256: 4dbae6bdeb9246450589e7a7f6e9b6a238e137718b819ff947b665b54ab6c135
---

# Llama-3.2-3B-Instruct-bnb-4bit - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Llama-3.2-3B-Instruct-bnb-4bit
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's 4-bit bitsandbytes quantization of Meta's Llama 3.2 3B Instruct model. Llama 3.2 is a collection of multilingual large language models (LLMs) from Meta, released September 25, 2024, in 1B and 3B text-only sizes alongside larger vision models. The 3B instruction-tuned text model is an auto-regressive language model using an optimized transformer architecture, aligned with human preferences for helpfulness and safety via supervised fine-tuning (SFT) and reinforcement learning with human feedback (RLHF). It is optimized for multilingual dialogue use cases, including agentic retrieval and summarization, and Meta states it outperforms many open and closed chat models on common industry benchmarks. It uses Grouped-Query Attention (GQA) for improved inference scalability and supports eight officially supported languages: English, German, French, Italian, Portuguese, Hindi, Spanish, and Thai. The model has 3B parameters, a context length of 128K tokens, and a knowledge cutoff of December 2023. Use is governed by the custom, commercial Llama 3.2 Community License. This Unsloth build is provided in 4-bit precision (bitsandbytes), tagged with 4-bit and bitsandbytes, and reports tensor types F32, BF16, and U8, with 69,887 monthly downloads. Unsloth's purpose here is to enable faster, more memory-efficient fine-tuning: their notebooks claim roughly 2.4x faster training and 58% less memory for Llama 3.2 3B, with export to GGUF, vLLM, or the Hugging Face Hub. It can be loaded with Transformers via pipeline or AutoModelForCausalLM, or served with vLLM/SGLang. The card defers benchmark details and full model information to Meta's original model card.

## Key points

- Unsloth 4-bit (bitsandbytes) quantization of Meta Llama 3.2 3B Instruct.
- Text-only, auto-regressive optimized transformer; SFT + RLHF aligned.
- 3B parameters, 128K context, December 2023 knowledge cutoff.
- 8 officially supported languages; optimized for multilingual dialogue, retrieval, and summarization.
- Custom, commercial Llama 3.2 Community License (llama3.2).
- Purpose: memory-efficient fine-tuning (~2.4x faster, 58% less memory) with GGUF/vLLM export.
- Loadable via Transformers, vLLM, or SGLang; base model is meta-llama/Llama-3.2-3B-Instruct.
- 69,887 downloads/month; tensor types F32/BF16/U8.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Llama-3.2-3B-Instruct-bnb-4bit |
| Base model | meta-llama/Llama-3.2-3B-Instruct |
| Architecture | Auto-regressive optimized transformer (GQA) |
| Parameters | 3B |
| Context length | 128K tokens |
| Training alignment | SFT + RLHF |
| Knowledge cutoff | December 2023 |
| Release date | September 25, 2024 |
| Languages | 8 supported |
| Quantization | 4-bit (bitsandbytes) |
| Tensor types | F32, BF16, U8 |
| License | Llama 3.2 Community License |
| Fine-tuning speed | ~2.4x faster, 58% less memory (Unsloth) |
| Downloads/month | 69,887 |

## Why this source matters for the RAG

This card documents a widely used 4-bit Unsloth quantization of Llama 3.2 3B Instruct, a small, efficient multilingual chat model popular for local and fine-tuning workflows. It is a key reference for lightweight open models, bitsandbytes quantization, and memory-efficient fine-tuning.
