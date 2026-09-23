---
id: collect-huggingface/huggingface/unsloth-llama-4-maverick-17b-128e-instruct-gguf
title: "Llama-4-Maverick-17B-128E-Instruct-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Meta", "SGLang", "Unsloth", "vLLM"]
dates: ["2024-08", "2026-09-23"]
keywords: ["gguf", "llama", "context window", "fine-tuning", "gpu", "license", "llama.cpp", "memory", "moe", "multimodal", "parameters", "quantization"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Llama-4-Maverick-17B-128E-Instruct-GGUF.md
source_anchor: ""
source_lines: [1, 49]
sha256: c71db645d2f171562f6a3c2fb1747230f1129482a21cca8db58e561b16035b13
---

# Llama-4-Maverick-17B-128E-Instruct-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Llama-4-Maverick-17B-128E-Instruct-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth's GGUF quantization repository for Meta's Llama 4 Maverick 17B 128E Instruct model, produced with Unsloth Dynamic v2.0 (claimed to outperform other leading quantization methods in accuracy). Llama 4 Maverick is a natively multimodal, auto-regressive MoE model with 17B activated and 400B total parameters across 128 experts, ~22T training tokens, an August 2024 knowledge cutoff, and a 1M-token context window. This repository is the quantized text-only variant: the card notes "Currently text only is supported," so the multimodal vision path is not exposed through these GGUF files. It is governed by the Llama 4 Community License. Because the full model is large, the quantization sizes are substantial. Highlighted MoE-bit builds and disk sizes: UD-IQ1_S (1.78-bit, 122GB, "Ok"), UD-IQ1_M (1.93-bit, 128GB, "Fair"), UD-IQ2_XXS (2.42-bit, 140GB, "Better"), UD-Q2_K_XL (2.71-bit, 151GB, "Suggested"), UD-Q3_K_XL (3.5-bit, 193GB, "Great"), and UD-Q4_K_XL (4.5-bit, 243GB, "Best"). The full hardware-compatibility table spans 1-bit (UD-TQ1_0 106GB, UD-IQ1_S 121GB) through 16-bit (BF16 801GB), including Q8_0 (426GB), Q6_K (329GB), and Q4_K_M (243GB). The hub reports 401B params, architecture llama4, and 6,767 monthly downloads. Usage is via llama.cpp (`llama serve`/`llama cli`), LM Studio, Jan, Ollama, vLLM, SGLang, Docker Model Runner, and other local apps, with the example quant UD-Q4_K_XL. The chat template is `<|header_start|>user<|header_end|>\n\n...<|eot|><|header_start|>assistant<|header_end|>\n\n`. Unsloth also documents fine-tuning Llama 4 with a single H100 80GB GPU.

## Key points

- Unsloth Dynamic v2.0 GGUF quantizations of Meta Llama 4 Maverick 17B 128E Instruct.
- Text-only in this repository; vision capabilities are not exposed.
- 401B-parameter MoE (17B activated, 128 experts); 1M-token context; ~22T tokens.
- Large quant range: 1-bit UD-TQ1_0 106GB up to 16-bit BF16 801GB; recommended UD-Q4_K_XL 243GB.
- Accuracy ratings guide selection: IQ1_S "Ok" up to Q4_K_XL "Best".
- Runs with llama.cpp, LM Studio, Jan, Ollama, vLLM, SGLang, and Docker.
- Fine-tunable with Unsloth tooling.
- Llama 4 Community License; 6,767 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Llama-4-Maverick-17B-128E-Instruct-GGUF |
| Base model | meta-llama/Llama-4-Maverick-17B-128E-Instruct |
| Architecture | MoE (llama4), text-only in GGUF |
| Total params | 401B |
| Active params | 17B (128 experts) |
| Context length | 1M tokens |
| Quantization | Unsloth Dynamic v2.0 GGUF (1-bit to 16-bit) |
| Recommended quant | UD-Q4_K_XL (4.5-bit, 243GB) |
| Smallest quant | UD-TQ1_0 (106GB) |
| Largest quant | BF16 (801GB) |
| License | Llama 4 Community License |
| Runtimes | llama.cpp, LM Studio, Ollama, vLLM, SGLang |
| Downloads/month | 6,767 |
| arXiv | 2204.05149 |

## Why this source matters for the RAG

This card documents local GGUF quantizations of Llama 4 Maverick, showing the large memory requirements and Unsloth Dynamic v2.0 accuracy trade-offs for a 400B-class MoE. It is a key reference for quantized deployment of frontier-scale open models.
