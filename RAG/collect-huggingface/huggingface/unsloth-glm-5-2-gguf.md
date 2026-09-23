---
id: collect-huggingface/huggingface/unsloth-glm-5-2-gguf
title: "GLM-5.2-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "glm", "agent", "agents", "attention", "benchmarks", "inference", "license", "llama", "llama.cpp", "mcp", "mit license"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-GLM-5.2-GGUF.md
source_anchor: ""
source_lines: [1, 50]
sha256: 17776cd0d555b09bc2f5f0aaa517cb25dbac0affc68eeb3f443c2a109bc3160b
---

# GLM-5.2-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/GLM-5.2-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/GLM-5.2-GGUF is the GGUF quantized release of zai-org/GLM-5.2 by Unsloth, based on Unsloth Dynamic 2.0 GGUFs. GLM-5.2 is Z.ai's flagship model for long-horizon tasks, delivering a solid 1M-token context for the first time, with advanced coding and flexible thinking effort levels. Architecture is `glm_moe_dsa` (MoE with dynamic sparse attention), 754B reported params, MIT license. Key improvements over GLM-5.1 include the IndexShare proposal (arXiv 2603.12201) which reuses the same indexer across every four sparse attention layers, reducing per-token FLOPs by 2.9x at 1M context, and an improved MTP layer for speculative decoding that increases acceptance length by up to 20%.

This repository provides GGUF quants from 1-bit to 16-bit for local inference on llama.cpp-compatible tools. File sizes: 1-bit UD-IQ1_S 217 GB, UD-IQ1_M 228 GB; 2-bit UD-IQ2_XXS 238 GB, UD-IQ2_M 239 GB, UD-Q2_K_XL 254 GB; 3-bit UD-IQ3_XXS 282 GB, UD-IQ3_S 309 GB, UD-Q3_K_M 343 GB, UD-Q3_K_XL 343 GB; 4-bit UD-IQ4_XS 365 GB, UD-IQ4_NL 373 GB, UD-Q4_K_S 436 GB, UD-Q4_K_M 466 GB, UD-Q4_K_XL 467 GB; 5-bit UD-Q5_K_S 527 GB, UD-Q5_K_M 561 GB, UD-Q5_K_XL 562 GB; 6-bit UD-Q6_K 626 GB, UD-Q6_K_XL 684 GB; 8-bit Q8_0 801 GB, UD-Q8_K_XL 820 GB; 16-bit BF16 1.51 TB. Default recommended quant is UD-Q4_K_M.

Benchmarks: HLE 40.5, HLE (w/ tools) 54.7, AIME 2026 99.2, GPQA-Diamond 91.2, SWE-bench Pro 62.1, NL2Repo 48.9, DeepSWE 46.2, ProgramBench 63.7, Terminal Bench 2.1 81.0/82.7, FrontierSWE 74.4, PostTrainBench 34.3, SWE-Marathon 13.0, MCP-Atlas 76.8, Tool-Decathlon 48.2. Runs via llama.cpp, LM Studio, Jan, Ollama, vLLM, Docker Model Runner, Unsloth Studio (High/Max thinking toggles), and agent tools. ~350K monthly downloads.

## Key points

- Unsloth Dynamic 2.0 GGUF quants of GLM-5.2 (754B params, glm_moe_dsa).
- Full quant range 1-bit (217 GB) to 16-bit BF16 (1.51 TB); default UD-Q4_K_M (466 GB).
- 1M-token context; IndexShare reduces per-token FLOPs by 2.9x.
- MIT license; AIME 2026 99.2, GPQA-Diamond 91.2, Terminal Bench 2.1 81.0.
- Runs via llama.cpp, Ollama, LM Studio, vLLM, Unsloth Studio, agents.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 754B |
| Architecture | glm-dsa (glm_moe_dsa) |
| Context length | Solid 1M tokens |
| Quantization | GGUF (Unsloth Dynamic 2.0) |
| 1-bit (UD-IQ1_S / IQ1_M) | 217 GB / 228 GB |
| 2-bit (IQ2_XXS / IQ2_M / Q2_K_XL) | 238 GB / 239 GB / 254 GB |
| 3-bit (IQ3_XXS / IQ3_S / Q3_K_M) | 282 GB / 309 GB / 343 GB |
| 4-bit (IQ4_XS / IQ4_NL / Q4_K_M / Q4_K_XL) | 365 GB / 373 GB / 466 GB / 467 GB |
| 5-bit (Q5_K_S / Q5_K_M / Q5_K_XL) | 527 GB / 561 GB / 562 GB |
| 6-bit (Q6_K / Q6_K_XL) | 626 GB / 684 GB |
| 8-bit (Q8_0 / Q8_K_XL) | 801 GB / 820 GB |
| 16-bit BF16 | 1.51 TB |
| License | MIT |
| AIME 2026 / GPQA-Diamond | 99.2 / 91.2 |
| Monthly downloads | ~350K |

## Why this source matters for the RAG

This card captures the most complete GGUF quantization matrix for Z.ai's 1M-context flagship, providing exact file sizes for local deployment planning. It documents the IndexShare architecture advance and MIT licensing that made 753B-scale local inference feasible. It is a key reference for quantization trade-off questions in the RAG knowledge base.
