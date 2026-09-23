---
id: collect-huggingface/huggingface/unsloth-gpt-oss-20b-gguf
title: "gpt-oss-20b-GGUF (Unsloth) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "agentic", "apache", "consumer", "fine-tuning", "inference", "license", "llama", "llama.cpp", "moe", "mxfp4", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-gpt-oss-20b-GGUF.md
source_anchor: ""
source_lines: [1, 44]
sha256: 735d6562ddd28adeae2f08b1570c01138ddf0ea1bc10336c447ae1684b25c874
---

# gpt-oss-20b-GGUF (Unsloth) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/gpt-oss-20b-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository by Unsloth AI provides GGUF quantizations of OpenAI's gpt-oss-20b model for local inference across llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang, Docker Model Runner and other compatible applications. The base model is a 21B-parameter Mixture-of-Experts (MoE) reasoning model with 3.6B active parameters, released by OpenAI under Apache 2.0 and trained on the harmony response format, which must be respected for correct operation. Unsloth's release uses their "Dynamic 2.0" quantization methodology, advertised as achieving superior accuracy and state-of-the-art quantization performance. A wide quant ladder is provided: Q2_K (11.5 GB), Q2_K_L (11.8 GB), Q3_K_S/Q3_K_M (11.5 GB), Q4_K_S/Q4_0/Q4_1/Q4_K_M (11.5–11.6 GB), the recommended UD-Q4_K_XL (11.9 GB), Q5_K_S/Q5_K_M (11.7 GB), Q6_K (12 GB) and UD-Q6_K_XL (12 GB), Q8_0 (12.1 GB) and UD-Q8_K_XL (13.2 GB), and F16 (13.8 GB). Note that the F32 quant is MXFP4 upcast to BF16 for every layer and is effectively unquantized. The model inherits gpt-oss-20b's highlights: permissive Apache 2.0 license, configurable reasoning effort (low/medium/high), full chain-of-thought access, fine-tunability, and agentic capabilities (function calling, web browsing, Python code execution, structured outputs). It can be run via `llama serve -hf unsloth/gpt-oss-20b-GGUF:UD-Q4_K_XL` or `ollama run hf.co/unsloth/gpt-oss-20b-GGUF:UD-Q4_K_XL`. Unsloth provides a free Colab notebook for fine-tuning gpt-oss-20b, and credits the llama.cpp team for quantization support. The repository has 532,675 downloads in the last month and 834 likes.

## Key points

- Unsloth Dynamic 2.0 GGUF quantizations of OpenAI gpt-oss-20b for local inference.
- Quant range from Q2_K (~11.5 GB) to F16 (13.8 GB), including UD-Q4_K_XL (recommended), UD-Q6_K_XL and UD-Q8_K_XL.
- F32 quant is MXFP4 upcast to BF16 and is unquantized.
- Base model: 21B total / 3.6B active MoE, Apache 2.0, harmony response format.
- Supported apps: llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang, Docker Model Runner, and more.
- Free Colab notebook available for fine-tuning gpt-oss-20b.
- 532,675 downloads last month; part of the Unsloth Dynamic 2.0 Quants collection.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | openai/gpt-oss-20b |
| Total / active params | 21B / 3.6B (MoE) |
| Format | GGUF (Unsloth Dynamic 2.0) |
| Quant variants | Q2_K, Q2_K_L, Q3_K_S/M, Q4_K_S, Q4_0/1, Q4_K_M, UD-Q4_K_XL, Q5_K_S/M, Q6_K, UD-Q6_K_XL, Q8_0, UD-Q8_K_XL, F16, F32 |
| Size range | ~11.5 GB (Q2_K) to 13.8 GB (F16) |
| Recommended quant | UD-Q4_K_XL (11.9 GB) |
| License | Apache 2.0 |
| Response format | OpenAI harmony |
| Reasoning levels | Low, medium, high |
| Downloads last month | 532,675 |
| Likes | 834 |

## Why this source matters for the RAG

This card documents the practical, community quantization path that makes a frontier open-weight reasoning model runnable on consumer hardware, including concrete file sizes and recommended quants. It is a key reference for local deployment and quantization trade-off analysis in the knowledge base.
