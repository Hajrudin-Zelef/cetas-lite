---
id: collect-huggingface/huggingface/qwen-qwen-72b
title: "Qwen-72B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "China", "Hugging Face", "Meta", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["qwen", "attention", "benchmarks", "gpu", "int4", "license", "memory", "open-weight", "parameters", "research", "training"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen-72B.md
source_anchor: ""
source_lines: [1, 54]
sha256: 7ef3da339b69b8cdbb2e7e7d9fd3c7daa6ecff694d08d5f7ad7c757cc04bcd22
---

# Qwen-72B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen-72B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen-72B is the 72-billion-parameter base model of the Qwen (Tongyi Qianwen) series developed by Alibaba Cloud. It is a Transformer-based large language model pretrained on over 3 trillion tokens covering Chinese, English, multilingual text, code, and mathematics. A chat-aligned variant, Qwen-72B-Chat, was also released separately. This repository hosts the base pretrained model.

Key features include large-scale high-quality training corpora, competitive performance that surpassed existing open-source models on Chinese and English benchmarks at release, a comprehensive vocabulary of over 150K tokens (based on the GPT-4 BPE `cl100k_base` tokenizer, optimized for Chinese and multilingual), and a 32K context length enabled by extending the RoPE base. It uses RoPE relative position encoding, SwiGLU activation, and RMSNorm.

Architecture hyperparameters: 80 layers, 64 attention heads, model dimension 8192, vocabulary size 151,851, sequence length 32,768. Running in BF16/FP16 requires at least 144GB of GPU memory (e.g., 2×A100-80G or 5×V100-32G); the Int4 version needs at least 48GB (e.g., 1×A100-80G).

License: Tongyi Qianwen License Agreement (custom), open for research and commercial use, with a form required for commercial use. The model requires `transformers==4.32.0` and `trust_remote_code=True`. Benchmarks: Qwen-72B achieved an average of 66.4 across MMLU (77.4), C-Eval (83.3), GSM8K (78.9), MATH (35.2), HumanEval (35.4), MBPP (52.2), BBH (67.7), AGIEval (62.5), GaokaoBench (87.6), and CMMLU (83.6), outperforming LLaMA2-70B, Yi-34B, and InternLM-20B. Long-context PPL at 32K was 2.717.

## Key points

- 72B-parameter dense base LLM from Alibaba Cloud; 3T+ training tokens.
- 32K context length via extended RoPE base; 80 layers, d_model 8192.
- Vocabulary of 151,851 tokens (cl100k_base-based), strong multilingual support.
- Requires 144GB+ VRAM in BF16/FP16; 48GB+ for Int4.
- Custom Tongyi Qianwen License (research + commercial, with form).
- Qwen-72B avg score 66.4 across 10 benchmarks at release; beats LLaMA2-70B.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 72B |
| Active parameters | 72B (dense) |
| Layers | 80 |
| Attention heads | 64 |
| Model dimension | 8192 |
| Vocabulary size | 151,851 |
| Sequence length | 32,768 |
| License | Tongyi Qianwen License Agreement |
| Precision | BF16/FP16 (144GB+ VRAM); Int4 (48GB+) |
| Tokenizer | tiktoken (cl100k_base-based) |
| MMLU | 77.4 |
| C-Eval | 83.3 |
| GSM8K | 78.9 |
| HumanEval | 35.4 |
| Average (10 benchmarks) | 66.4 |
| Monthly downloads | ~4.1M |

## Why this source matters for the RAG

Qwen-72B is a historical reference point for large open-weight Chinese/English LLMs and provides a baseline against which newer Qwen generations (Qwen3.x) can be compared. Its detailed architecture, license, and hardware requirements are useful for understanding model scaling and multilingual tokenization. It enriches the knowledge base's longitudinal view of the Qwen family.
