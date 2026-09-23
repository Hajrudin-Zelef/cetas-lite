---
id: collect-huggingface/huggingface/unsloth-deepseek-v4-flash-vision-exp-gguf
title: "DeepSeek-V4-Flash-Vision-Exp-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "OpenAI", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["deepseek", "gguf", "agent", "agents", "attention", "benchmark", "benchmarks", "inference", "license", "llama", "llama.cpp", "mit license"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-DeepSeek-V4-Flash-Vision-Exp-GGUF.md
source_anchor: ""
source_lines: [1, 54]
sha256: 3ec5953da021e76b2a6a64e23d4c6178f5346db0fc342ca5cda650771bbbab0d
---

# DeepSeek-V4-Flash-Vision-Exp-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/DeepSeek-V4-Flash-Vision-Exp-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of DeepSeek-V4-Flash-Vision-Exp, DeepSeek's first experimental multimodal model in the DeepSeek-V4 family. The quantized repo is tagged Transformers, GGUF, unsloth, deepseek, imatrix, and conversational, licensed under MIT, derived from deepseek-ai/DeepSeek-V4-Flash-Vision-Exp, with 284B params and architecture `deepseek4`, and 34,728 monthly downloads. The card reproduces DeepSeek's benchmark table comparing the vision model to DeepSeek-V4-Flash-0731 and Opus-4.8. Text agent scores: Terminal Bench 2.1 = 83.9, NL2Repo = 57.7, Cybergym = 75.3, DeepSWE = 59.3, Toolathlon-Verified = 75.9, DSBench-Hard = 63.6, AutomationBench Public = 25.7. Multimodal agent scores: ApexBench pass@1 = 36.5 (vs 26.2 when the text-only model ignores images), Agents' Last Exam = 27.3, Chartography = 64.3, ZeroBench pass@5 = 35.0.

Unsloth Dynamic 3.0 quantization is used here (noting superior accuracy versus other leading quants). Available variants span 1-bit (UD-IQ1_S 82.4GB, UD-IQ1_M 86.9GB), 2-bit (UD-IQ2_XXS 90.7GB, UD-Q2_K_XL 96.8GB), 3-bit (UD-IQ3_XXS 103GB, UD-IQ3_S 114GB, UD-Q3_K_XL 128GB), 4-bit (UD-IQ4_XS 137GB, UD-Q4_K_XL 155GB), and 8-bit (UD-Q8_K_XL 162GB). As with the text variant, Q8 is 162GB and only ~7GB larger than Q4, so lossless full-precision use is practical.

A key implementation note is provided: image input requires llama.cpp build b10766 or later, the first release containing DeepSeek-V4 vision support (PRs #28133 and #28154). The card documents the model's repository layout (tokenizer, prompt encoding reference, minimal PyTorch inference covering the vision encoder and aligner, DFlash attention, MoE, Hyper-Connections, and DSpark), and that prompt encoding supports both OpenAI-style JSON content blocks and compact `<image>path</image>` TXT notation. The card supplies run instructions for llama.cpp (`llama serve -hf unsloth/DeepSeek-V4-Flash-Vision-Exp-GGUF:UD-Q4_K_XL`), LM Studio, Jan, Ollama, Unsloth Studio (High/Max thinking toggles), Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, and OpenClaw. License is MIT.

## Key points

- Unsloth Dynamic 3.0 GGUF quantization of the experimental DeepSeek-V4-Flash-Vision-Exp multimodal model.
- 284B params, architecture deepseek4, MIT license, derived from deepseek-ai/DeepSeek-V4-Flash-Vision-Exp.
- Quantization ladder 1-bit UD-IQ1_S (82.4GB) through 8-bit UD-Q8_K_XL (162GB).
- Image input requires llama.cpp build b10766+ (vision support via PRs #28133/#28154).
- Reproduces benchmark table: Terminal Bench 2.1 83.9, DeepSWE 59.3, ApexBench 36.5, Chartography 64.3, ZeroBench 35.0.
- Supports OpenAI-style JSON and `<image>path</image>` TXT prompt encoding.
- Broad local-app support (llama.cpp, Ollama, LM Studio, Jan, Unsloth Studio, Docker Model Runner, Lemonade, Hermes, OpenClaw, Pi).
- 34,728 downloads/month; Unsloth Studio High/Max thinking toggles.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | DeepSeek-V4-Flash-Vision-Exp-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 3.0) |
| Base model | deepseek-ai/DeepSeek-V4-Flash-Vision-Exp |
| Reported size | 284B params |
| Architecture | deepseek4 |
| License | MIT |
| 1-bit | UD-IQ1_S 82.4GB; UD-IQ1_M 86.9GB |
| 2-bit | UD-IQ2_XXS 90.7GB; UD-Q2_K_XL 96.8GB |
| 3-bit | UD-IQ3_XXS 103GB; UD-IQ3_S 114GB; UD-Q3_K_XL 128GB |
| 4-bit | UD-IQ4_XS 137GB; UD-Q4_K_XL 155GB |
| 8-bit | UD-Q8_K_XL 162GB |
| Vision requirement | llama.cpp build b10766 or later |
| Serving | llama.cpp, Ollama, LM Studio, Jan, Unsloth Studio, etc. |
| Key benchmarks | Terminal Bench 2.1 83.9; DeepSWE 59.3; ApexBench 36.5; Chartography 64.3; ZeroBench 35.0 |
| Downloads/month | 34,728 |

## Why this source matters for the RAG

This card documents how to run an experimental frontier multimodal MoE locally via GGUF, including the specific llama.cpp build requirement for vision. It is valuable for retrieval on multimodal local inference, quantization trade-offs, and hardware sizing for large vision-language models.
