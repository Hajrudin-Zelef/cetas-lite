---
id: collect-huggingface/huggingface/unsloth-glm-5-3-gguf
title: "GLM-5.3-GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "glm", "agent", "agents", "benchmarks", "cyber", "deepseek", "fable 5", "gpt-5.6", "inference", "kimi", "license"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-GLM-5.3-GGUF.md
source_anchor: ""
source_lines: [1, 49]
sha256: 723b27710444dc0cd84009165b885ed940207415c5476494e628e8ebb4e3e36d
---

# GLM-5.3-GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/GLM-5.3-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/GLM-5.3-GGUF is the GGUF quantized release of zai-org/GLM-5.3 by Unsloth, based on Unsloth Dynamic 3.0 GGUFs for superior quantization accuracy. GLM-5.3 uses the same base model as GLM-5.2 (architecture `glm_moe_dsa`, 754B reported params, license glm-5.3 custom); every gain comes from post-training. It is the most capable open-weights model for coding, with a 50% improvement over GLM-5.2 on Z.ai's in-house Code Bench, plus open-source SOTA on Terminal Bench 3.0 and Agents' Last Exam. It also shows emergent cyber capability: SOTA on CyberGym for vulnerability discovery, more than doubling GLM-5.2 on exploitation benchmarks.

This repository provides GGUF quants from 1-bit to 16-bit for local inference. File sizes: 1-bit UD-IQ1_S 217 GB, UD-IQ1_M 228 GB; 2-bit UD-IQ2_M 239 GB, UD-Q2_K_XL 254 GB; 3-bit UD-IQ3_XXS 282 GB, UD-Q3_K_XL 343 GB; 4-bit UD-IQ4_XS 365 GB, UD-Q4_K_XL 467 GB; 5-bit UD-Q5_K_XL 562 GB; 6-bit UD-Q6_K_XL 684 GB; 8-bit Q8_0 801 GB; 16-bit BF16 1.51 TB. The default recommended quant is UD-Q4_K_XL.

It runs via llama.cpp, LM Studio, Jan, Ollama, vLLM, SGLang, Docker Model Runner, Unsloth Desktop (with Low/High/Max thinking toggles), and agent tools (Pi, Hermes, OpenClaw, Lemonade). Benchmarks vs Kimi K3, DeepSeek-V4 Pro, Qwen3.8-Max, Opus 4.8, Fable 5, GPT-5.6 Sol: Terminal Bench 2.1 88.2, Terminal Bench 3.0 28.3, DeepSWE 66.9, NL2Repo 58.0, FrontierSWE 78.1, CyberGym 84.5, ExploitBench 54.4, Toolathlon Verified 73.0, AutomationBench 48.2, Agents' Last Exam 28.5, HLE w/ tools 62.5, GDPval 1769. `reasoning_effort` (low/high/max) controls thinking budget. ~562K monthly downloads.

## Key points

- Unsloth Dynamic 3.0 GGUF quants of GLM-5.3 (754B params, glm_moe_dsa).
- Quants from 1-bit (217 GB) to 16-bit BF16 (1.51 TB); default UD-Q4_K_XL (467 GB).
- Most capable open-weights coding model; open-source SOTA on Terminal Bench 3.0 and ALE.
- Emergent cyber capability: CyberGym 84.5, ExploitBench 54.4.
- Custom GLM-5.3 license; runs via llama.cpp, Ollama, LM Studio, vLLM, SGLang, Unsloth Desktop.
- `reasoning_effort` low/high/max toggles in Unsloth Desktop.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 754B |
| Architecture | glm-dsa (glm_moe_dsa) |
| Quantization | GGUF (Unsloth Dynamic 3.0) |
| 1-bit (UD-IQ1_S / IQ1_M) | 217 GB / 228 GB |
| 2-bit (IQ2_M / Q2_K_XL) | 239 GB / 254 GB |
| 3-bit (IQ3_XXS / Q3_K_XL) | 282 GB / 343 GB |
| 4-bit (IQ4_XS / Q4_K_XL) | 365 GB / 467 GB |
| 5-bit / 6-bit | UD-Q5_K_XL 562 GB / UD-Q6_K_XL 684 GB |
| 8-bit / 16-bit | Q8_0 801 GB / BF16 1.51 TB |
| License | glm-5.3 (custom) |
| Terminal Bench 2.1 / 3.0 | 88.2 / 28.3 |
| CyberGym / ExploitBench | 84.5 / 54.4 |
| Monthly downloads | ~562K |

## Why this source matters for the RAG

This card documents the complete GGUF quantization ladder for running a 754B model locally, essential for hardware sizing and offline inference planning. It connects the top-tier GLM-5.3 model with the local tooling ecosystem (llama.cpp, Ollama, Unsloth Desktop). It strengthens the knowledge base's coverage of local AI deployment for frontier-scale models.
