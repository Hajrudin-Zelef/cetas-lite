---
id: collect-huggingface/huggingface/unsloth-llama-3-3-70b-instruct-bnb-4bit
title: "Llama-3.3-70B-Instruct-bnb-4bit - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "SGLang", "Unsloth", "vLLM"]
dates: ["2023-12", "2024-12-06", "2026-09-23"]
keywords: ["llama", "attention", "benchmark", "benchmarks", "compute", "context window", "fine-tuning", "gguf", "gpu", "gqa", "inference", "license"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Llama-3.3-70B-Instruct-bnb-4bit.md
source_anchor: ""
source_lines: [1, 51]
sha256: 38bbf43baffbc4a246df7392af68b483d6f7a7aeabb7e674fe3d15cdf5acadb6
---

# Llama-3.3-70B-Instruct-bnb-4bit - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Llama-3.3-70B-Instruct-bnb-4bit
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's 4-bit bitsandbytes quantization of Meta's Llama 3.3 70B Instruct model. Llama 3.3 is a multilingual large language model from Meta, released December 6, 2024, offered in a single 70B text-only size. It is an auto-regressive language model using an optimized transformer architecture, aligned to human preferences for helpfulness and safety through supervised fine-tuning (SFT) and reinforcement learning with human feedback (RLHF). It is optimized for multilingual dialogue use cases and Meta states it outperforms many open and closed chat models on common industry benchmarks. The model uses Grouped-Query Attention (GQA) for improved inference scalability, supports eight languages (English, German, French, Italian, Portuguese, Hindi, Spanish, Thai), has 70B parameters, a 128K-token context window, and was pretrained on 15T+ tokens with a December 2023 knowledge cutoff. It also supports tool use through chat templates and function calling. Use is governed by the custom, commercial Llama 3.3 Community License. This Unsloth build is 4-bit precision (bitsandbytes) and reports tensor types F32, BF16, and U8, with 12,124 monthly downloads; the hub lists 71B params. Instruction-tuned benchmark results include MMLU (CoT) 86.0, MMLU Pro (CoT) 68.9, IFEval 92.1, GPQA Diamond 50.5, HumanEval 88.4, MBPP EvalPlus 87.6, MATH (CoT) 77.0, BFCL v2 77.3, and MGSM 91.1. Training used 7.0M GPU hours on H100-80GB, emitting ~2,040 tons CO2eq (total reported as 11,390 tons across the effort). Unsloth's notebooks claim roughly 2-2.4x faster fine-tuning with ~58% less memory. It loads via Transformers, vLLM, or SGLang, and can be further quantized to 8-bit/4-bit with bitsandbytes.

## Key points

- Unsloth 4-bit (bitsandbytes) quantization of Meta Llama 3.3 70B Instruct.
- Text-only auto-regressive transformer with Grouped-Query Attention; SFT + RLHF aligned.
- 70B parameters, 128K context, 15T+ training tokens, December 2023 cutoff.
- 8 supported languages; tool use / function calling supported.
- Custom, commercial Llama 3.3 Community License (llama3.3).
- Strong benchmarks: MMLU 86.0, MMLU Pro 68.9, HumanEval 88.4, IFEval 92.1, MGSM 91.1.
- Purpose: memory-efficient fine-tuning (~2x faster, 58% less memory) with GGUF/vLLM export.
- Loadable via Transformers, vLLM, SGLang; base model meta-llama/Llama-3.3-70B-Instruct.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Llama-3.3-70B-Instruct-bnb-4bit |
| Base model | meta-llama/Llama-3.3-70B-Instruct |
| Architecture | Auto-regressive optimized transformer (GQA) |
| Parameters | 70B (hub: 71B) |
| Context length | 128K tokens |
| Training tokens | 15T+ |
| Knowledge cutoff | December 2023 |
| Release date | December 6, 2024 |
| Languages | 8 supported |
| Quantization | 4-bit (bitsandbytes) |
| Tensor types | F32, BF16, U8 |
| License | Llama 3.3 Community License |
| Training compute | 7.0M GPU hours (H100-80GB) |
| Key benchmarks | MMLU 86.0; MMLU Pro 68.9; HumanEval 88.4; IFEval 92.1; GPQA-D 50.5; MGSM 91.1 |
| Downloads/month | 12,124 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This card documents a widely used 4-bit Unsloth quantization of Llama 3.3 70B Instruct, one of the strongest mid-size open chat models with tool use. It is a key reference for quantization, fine-tuning efficiency, and multilingual instruct benchmarks.
