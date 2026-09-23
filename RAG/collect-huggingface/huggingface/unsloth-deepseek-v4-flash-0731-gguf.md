---
id: collect-huggingface/huggingface/unsloth-deepseek-v4-flash-0731-gguf
title: "DeepSeek-V4-Flash-0731-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "Unsloth"]
dates: ["2026-09-23"]
keywords: ["deepseek", "gguf", "agent", "agentic", "agents", "benchmark", "benchmarks", "consumer", "inference", "license", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-DeepSeek-V4-Flash-0731-GGUF.md
source_anchor: ""
source_lines: [1, 54]
sha256: 5eb6570cc0041fa4c3410356fe6a83bc34025aa6fbafebe22881221c294b6928
---

# DeepSeek-V4-Flash-0731-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/DeepSeek-V4-Flash-0731-GGUF
- **Site** : Hugging Face
- **Type** : Model card (quantized GGUF)
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is Unsloth AI's GGUF quantization of DeepSeek-V4-Flash-0731, the official DeepSeek-V4-Flash release with enhanced agentic capabilities and an attached DSpark speculative-decoding module. The quantized repo is a GGUF library (`unsloth`, `deepseek_v4`, `imatrix`, `conversational` tags) under the MIT license, with 284B params reported and architecture `deepseek4`. It is derived from the base model deepseek-ai/DeepSeek-V4-Flash-0731 and has 211,738 monthly downloads. The card reproduces DeepSeek's benchmark table, showing DeepSeek-V4-Flash-0731 outperforming the preview and DeepSeek-V4-Pro (Preview): Terminal Bench 2.1 = 82.7, NL2Repo = 54.2, Cybergym = 76.7, DeepSWE = 54.4, Toolathlon-Verified = 70.3, Agents' Last Exam = 25.2, AutomationBench Public = 25.1, DSBench-FullStack = 68.7, DSBench-Hard = 59.6.

Unsloth Dynamic 2.0 quantizations are offered, claimed to achieve superior accuracy versus other leading quants. Available variants include 1-bit (UD-IQ1_S 82.5GB, UD-IQ1_M 86.9GB), 2-bit (UD-IQ2_XXS 90.9GB, UD-IQ2_M 90.9GB, UD-Q2_K_XL 96.8GB), 3-bit (UD-IQ3_XXS 104GB, UD-IQ3_S 116GB, UD-Q3_K_M 128GB, UD-Q3_K_XL 128GB), 4-bit (UD-IQ4_XS 137GB, UD-IQ4_NL 137GB, UD-Q4_K_XL 155GB), and 8-bit (UD-Q8_K_XL 162GB). There are also two DSpark-specific files: DSPARK Q8_0 (10.9GB) and DSPARK BF16 (11.3GB). Unsloth notes that Q8 (UD-Q8_K_XL) is 162GB and only 7GB larger than Q4, making full-precision lossless use feasible. DSpark support allows up to 2x faster decoding.

The card provides extensive run instructions for many local apps: llama.cpp (install via script, winget, prebuilt binary, or source; `llama serve -hf unsloth/DeepSeek-V4-Flash-0731-GGUF:UD-Q4_K_XL`), LM Studio, Jan, Ollama (`ollama run hf.co/...`), Unsloth Studio (with High/Max thinking toggles), Pi, Docker Model Runner, Lemonade, Hermes Agent, Atomic Chat, and OpenClaw. This makes the model broadly accessible on consumer and workstation hardware. The recommended default quantization in examples is UD-Q4_K_XL. License remains MIT.

## Key points

- Unsloth Dynamic 2.0 GGUF quantization of DeepSeek-V4-Flash-0731 (284B params, architecture deepseek4).
- MIT license; 211,738 downloads/month; derived from deepseek-ai/DeepSeek-V4-Flash-0731.
- Rich quantization ladder from 1-bit UD-IQ1_S (82.5GB) to 8-bit UD-Q8_K_XL (162GB).
- Dedicated DSpark speculative-decoding files (Q8_0 10.9GB, BF16 11.3GB) enabling up to 2x faster decoding.
- Q8 is only ~7GB larger than Q4, so lossless full-precision local runs are practical.
- Runs across llama.cpp, Ollama, LM Studio, Jan, Unsloth Studio, Docker Model Runner, Lemonade, Hermes, OpenClaw, Pi.
- Reproduces DeepSeek benchmark table (Terminal Bench 2.1 82.7, DeepSWE 54.4, Cybergym 76.7).
- Unsloth Studio offers High/Max thinking toggles.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | DeepSeek-V4-Flash-0731-GGUF |
| Type | GGUF quantizations (Unsloth Dynamic 2.0) |
| Base model | deepseek-ai/DeepSeek-V4-Flash-0731 |
| Reported size | 284B params |
| Architecture | deepseek4 |
| License | MIT |
| 1-bit | UD-IQ1_S 82.5GB; UD-IQ1_M 86.9GB |
| 2-bit | UD-IQ2_XXS 90.9GB; UD-IQ2_M 90.9GB; UD-Q2_K_XL 96.8GB |
| 3-bit | UD-IQ3_XXS 104GB; UD-IQ3_S 116GB; UD-Q3_K_M/X_XL 128GB |
| 4-bit | UD-IQ4_XS 137GB; UD-IQ4_NL 137GB; UD-Q4_K_XL 155GB |
| 8-bit | UD-Q8_K_XL 162GB |
| DSpark files | DSPARK Q8_0 10.9GB; DSPARK BF16 11.3GB |
| Serving | llama.cpp, Ollama, LM Studio, Jan, Unsloth Studio, etc. |
| Key benchmarks | Terminal Bench 2.1 82.7; DeepSWE 54.4; Cybergym 76.7; Toolathlon-V 70.3 |
| Downloads/month | 211,738 |

## Why this source matters for the RAG

This card is the practical bridge between a frontier MoE model and local/consumer deployment, documenting the full quantization size/accuracy trade-off space and the DSpark acceleration path. It is essential for retrieval on GGUF quantization, local inference options, and hardware sizing.
