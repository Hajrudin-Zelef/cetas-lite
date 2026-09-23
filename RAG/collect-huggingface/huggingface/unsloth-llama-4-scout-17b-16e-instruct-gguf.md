---
id: collect-huggingface/huggingface/unsloth-llama-4-scout-17b-16e-instruct-gguf
title: "Llama-4-Scout-17B-16E-Instruct-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "Microsoft", "SGLang", "Unsloth", "vLLM"]
dates: ["2024-08", "2026-09-23"]
keywords: ["gguf", "llama", "scout", "consumer", "context window", "fine-tuning", "gpu", "license", "llama.cpp", "moe", "multimodal", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Llama-4-Scout-17B-16E-Instruct-GGUF.md
source_anchor: ""
source_lines: [1, 49]
sha256: b372d7619bf2535ec6260c98b034016c344d5c566b1302bca8b36fbe5e77191c
---

# Llama-4-Scout-17B-16E-Instruct-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF quantization repository for Meta's Llama 4 Scout 17B 16E Instruct model, produced with Unsloth Dynamic v2.0, which the authors state achieves superior accuracy versus other leading quantization methods. Llama 4 Scout is a natively multimodal, auto-regressive MoE model with 17B activated and 109B total parameters across 16 experts, ~40T training tokens, an August 2024 knowledge cutoff, and a 10M-token context window. This repository is the quantized text-only variant: the card notes "Currently text only is supported," and the underlying vision capabilities are not exposed through these GGUF files. It is governed by the Llama 4 Community License. The card provides a detailed quantization table with MoE-bit, type, disk size, and accuracy guidance. Highlighted recommended builds: UD-IQ1_S (1.78-bit, 33.8GB, "Ok"), UD-IQ1_M (1.93-bit, 35.4GB, "Fair"), UD-IQ2_XXS (2.42-bit, 38.6GB, "Better"), UD-Q2_K_XL (2.71-bit, 42.2GB, "Suggested"), UD-Q3_K_XL (3.5-bit, 52.9GB, "Great"), and UD-Q4_K_XL (4.5-bit, 65.6GB, "Best"). The full hardware-compatibility table spans 1-bit (UD-IQ1_S 32.5GB, UD-TQ1_0 29.3GB) through 16-bit (BF16 216GB), including Q8_0 (115GB), Q6_K (88.4GB), and Q4_K_M (65.4GB). The hub reports 108B params, architecture llama4, and 31,099 monthly downloads. Usage is via llama.cpp (`llama serve`/`llama cli`), LM Studio, Jan, Ollama, vLLM, SGLang, Docker Model Runner, and many local apps, with the example quant UD-Q4_K_XL. The chat template is `<|header_start|>user<|header_end|>\n\n...<|eot|><|header_start|>assistant<|header_end|>\n\n`. Unsloth also highlights fine-tuning Llama 4 Scout on a single H100 80GB GPU.

## Key points

- Unsloth Dynamic v2.0 GGUF quantizations of Meta Llama 4 Scout 17B 16E Instruct.
- Text-only in this repository; vision capabilities are not exposed.
- 108B-parameter MoE (17B activated, 16 experts); 10M-token context; ~40T tokens.
- Quant range from 1-bit (UD-IQ1_S 32.5GB) to 16-bit BF16 (216GB); recommended UD-Q4_K_XL 65.6GB.
- Accuracy ratings guide selection: IQ1_S "Ok" up to Q4_K_XL "Best".
- Runs with llama.cpp, LM Studio, Jan, Ollama, vLLM, SGLang, and Docker.
- Fine-tunable on a single H100 80GB GPU per Unsloth.
- Llama 4 Community License; 31,099 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Llama-4-Scout-17B-16E-Instruct-GGUF |
| Base model | meta-llama/Llama-4-Scout-17B-16E-Instruct |
| Architecture | MoE (llama4), text-only in GGUF |
| Total params | 108B |
| Active params | 17B (16 experts) |
| Context length | 10M tokens |
| Quantization | Unsloth Dynamic v2.0 GGUF (1-bit to 16-bit) |
| Recommended quant | UD-Q4_K_XL (4.5-bit, 65.6GB) |
| Smallest quant | UD-TQ1_0 (29.3GB) / UD-IQ1_S (32.5GB) |
| Largest quant | BF16 (216GB) |
| License | Llama 4 Community License |
| Runtimes | llama.cpp, LM Studio, Ollama, vLLM, SGLang |
| Downloads/month | 31,099 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This card documents the practical local-deployment quantizations of Llama 4 Scout, including Unsloth Dynamic v2.0's accuracy/size trade-offs. It is a key reference for running large MoE models on consumer or single-GPU hardware via GGUF.
