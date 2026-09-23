---
id: collect-huggingface/huggingface/unsloth-glm-5-3-flash-gguf
title: "GLM-5.3-Flash-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Apple", "Hugging Face", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "glm", "agent", "agentic", "agents", "attention", "benchmarks", "claude", "fp8", "inference", "license", "llama"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-GLM-5.3-Flash-GGUF.md
source_anchor: ""
source_lines: [1, 48]
sha256: c4c0517818e80b62249a8265c124c595ec66ae72cfb9efb286dd1e5f7028060c
---

# GLM-5.3-Flash-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/GLM-5.3-Flash-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/GLM-5.3-Flash-GGUF is the GGUF quantized release of zai-org/GLM-5.3-Flash by Unsloth, using Unsloth Dynamic 3.0 for superior quantization accuracy that outperforms other leading quants. GLM-5.3-Flash is the first natively multimodal model in the GLM-5 series: 320B total parameters with just 18B active, outperforming GLM-5.2 at one-tenth the price while approaching Claude Opus 4.8 on coding and agentic benchmarks. Architecture highlights include a hybrid sparse + linear attention design (first in the GLM series) and Manifold-Constrained Hyper-Connections (mHC), trained on a 30T-token multimodal corpus. License is MIT.

This GGUF repository provides quantization from 1-bit up to 16-bit for local inference on CPU/Apple Silicon via llama.cpp. File sizes: UD-IQ1_S 93.1 GB, UD-IQ1_M 97.6 GB; 2-bit UD-IQ2_XXS 102 GB, UD-Q2_K_XL 109 GB; 3-bit UD-IQ3_XXS 120 GB, UD-Q3_K_XL 148 GB; 4-bit UD-IQ4_XS 157 GB, UD-Q4_K_XL 200 GB; 5-bit UD-Q5_K_XL 240 GB; 6-bit UD-Q6_K_XL 292 GB; 8-bit Q8_0 341 GB; 16-bit BF16 642 GB.

It runs via llama.cpp (requires Unsloth's llama.cpp PR #27754), Unsloth Desktop, LM Studio, Jan, Ollama (`ollama run hf.co/unsloth/GLM-5.3-Flash-GGUF:UD-Q4_K_XL`), vLLM, Docker Model Runner, and agent tools like Pi, Hermes, OpenClaw, and Lemonade. The default recommended quant is UD-Q4_K_XL. ~779K monthly downloads. GLM-5.3-Flash supports `reasoning_effort` (low/high/max, default max) for thinking budget control.

## Key points

- Unsloth Dynamic 3.0 GGUF quants of GLM-5.3-Flash (320B total / 18B active).
- Range from 1-bit (93.1 GB) to 16-bit BF16 (642 GB); default UD-Q4_K_XL (200 GB).
- MIT license; multimodal GLM-5 with hybrid sparse+linear attention and mHC.
- Runs via llama.cpp (Unsloth PR), LM Studio, Jan, Ollama, vLLM, Unsloth Desktop.
- Approaches Claude Opus 4.8 on coding/agentic benchmarks at ~1/10th price.
- Requires Unsloth's llama.cpp PR #27754 for local run.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 320B (321B reported) |
| Active parameters | 18B |
| Quantization | GGUF (Unsloth Dynamic 3.0) |
| 1-bit (UD-IQ1_S / IQ1_M) | 93.1 GB / 97.6 GB |
| 2-bit (IQ2_XXS / Q2_K_XL) | 102 GB / 109 GB |
| 3-bit (IQ3_XXS / Q3_K_XL) | 120 GB / 148 GB |
| 4-bit (IQ4_XS / Q4_K_XL) | 157 GB / 200 GB |
| 5-bit / 6-bit | UD-Q5_K_XL 240 GB / UD-Q6_K_XL 292 GB |
| 8-bit / 16-bit | Q8_0 341 GB / BF16 642 GB |
| License | MIT |
| Architecture | glm5next (hybrid attention + mHC) |
| Monthly downloads | ~779K |

## Why this source matters for the RAG

This card documents the full quantization ladder for running a 320B multimodal model locally, which is critical for VRAM/RAM planning in local AI deployments. It shows exact GGUF file sizes per bit level and the tooling ecosystem (llama.cpp, Ollama, LM Studio, agents). It complements the FP8/MLX quant entries and strengthens the knowledge base's local-inference coverage.
