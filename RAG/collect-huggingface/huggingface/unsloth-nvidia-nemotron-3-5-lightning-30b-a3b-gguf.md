---
id: collect-huggingface/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf
title: "NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF (Unsloth) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "nvidia", "agent", "attention", "benchmark", "benchmarks", "consumer", "license", "llama", "llama.cpp", "moe", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF.md
source_anchor: ""
source_lines: [1, 44]
sha256: 5e1f00634922fabd57b8c881db853bc4bc6f8cbad6d74490973c7804bf5af8cd
---

# NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF (Unsloth) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository by Unsloth AI hosts GGUF quantizations of NVIDIA's Nemotron-3.5-Lightning-30B-A3B model, built with Unsloth's Dynamic 2.0 quantization methodology, advertised as achieving superior accuracy and outperforming other leading quants. The base model is a 30B-parameter Mixture-of-Experts (MoE) hybrid (Mamba-2 + MoE + Attention) reasoning model with 3B active parameters, released by NVIDIA under the OpenMDW-1.1 license, with up to 1M token context (256K for single-H100 deployment) and support for English (plus coding languages), Spanish, French, German, Italian and Japanese. The card reproduces the full reference model card content, including the benchmark table comparing BF16 Lightning against Qwen 3.6 35B A3B, Gemma 4 26B A4B, Nemotron 3 Nano, Nemotron 3 Super and GPT-OSS 20B: Lightning BF16 scores include MMLU Pro 81.94, GPQA Diamond 75.44, SWE-bench Verified 51.56, SWE-bench Multilingual 39.33, Terminal-Bench 2.1 24.58, IFBench 71.88 and AA-LCR 52.00. The reference model includes DSpark, DFlash and MTP speculative decoding. Recommended sampling is temperature 1.0 and top_p 0.95; reasoning is configurable via the chat template (`enable_thinking`). Usage is supported via llama.cpp (`llama serve -hf unsloth/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF:UD-Q4_K_M`), Ollama, LM Studio, Jan, vLLM, SGLang, Docker Model Runner and others. The GGUF files use the imatrix tag and the recommended quant is UD-Q4_K_M.

## Key points

- Unsloth Dynamic 2.0 GGUF quantizations of NVIDIA Nemotron-3.5-Lightning-30B-A3B.
- Base model: 30B total / 3B active MoE hybrid (Mamba-2 + MoE + Attention), OpenMDW-1.1.
- Up to 1M context; 6 languages; reasoning toggled via `enable_thinking`.
- Reference benchmarks (BF16): MMLU Pro 81.94, SWE-bench Verified 51.56, GPQA Diamond 75.44, IFBench 71.88.
- Recommended quant UD-Q4_K_M; runs on llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang.
- Recommended sampling: temperature 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 |
| Total / active params | 30B / 3B (MoE hybrid) |
| Format | GGUF (Unsloth Dynamic 2.0, imatrix) |
| Recommended quant | UD-Q4_K_M |
| Context length | Up to 1M tokens |
| Languages | English, Spanish, French, German, Italian, Japanese |
| License | OpenMDW-1.1 |
| MMLU Pro (BF16) | 81.94 |
| SWE-bench Verified (BF16) | 51.56 |
| GPQA Diamond (BF16) | 75.44 |
| Recommended sampling | temp 1.0, top_p 0.95 |
| Runtimes | llama.cpp, Ollama, LM Studio, Jan, vLLM, SGLang, Docker Model Runner |

## Why this source matters for the RAG

This card documents the local, consumer-hardware deployment path for NVIDIA's agent-focused Lightning model through community GGUF quantization. It is a key reference for running a frontier reasoning model locally while preserving tool-calling and thinking behavior.
