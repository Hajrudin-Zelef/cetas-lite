---
id: collect-huggingface/huggingface/zai-org-glm-5-2-fp8
title: "GLM-5.2-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Huawei", "Hugging Face", "MiniMax", "OpenAI", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "glm", "agents", "ascend", "attention", "benchmark", "benchmarks", "claude", "deepseek", "gemini", "latency", "license"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-5.2-FP8.md
source_anchor: ""
source_lines: [1, 53]
sha256: c3feac6c7f31793db218aa67a5804927fcf43cb58ec6e12a420d734e3edd75c7
---

# GLM-5.2-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-5.2-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-5.2-FP8 is the FP8-quantized release of Z.ai's GLM-5.2 flagship model for long-horizon tasks. GLM-5.2 marks a substantial leap in long-horizon capability over GLM-5.1 and, for the first time, delivers that capability on a solid 1M-token context. It is licensed under MIT (pure open, no regional limits). The card notes that a newer version, GLM-5.3, is available. Reported model size is 753B params, architecture `glm_moe_dsa` (MoE with dynamic sparse attention), tensors BF16/F8_E4M3/F32.

GLM-5.2's new capabilities: solid 1M context for stable long-horizon work; advanced coding with flexible thinking effort levels; an improved architecture proposing IndexShare (arXiv 2603.12201), which reuses the same indexer across every four sparse attention layers, reducing per-token FLOPs by 2.9x at 1M context; and an improved MTP layer for speculative decoding increasing acceptance length by up to 20%.

Benchmarks: HLE 40.5, HLE (w/ tools) 54.7, AIME 2026 99.2, GPQA-Diamond 91.2, SWE-bench Pro 62.1, NL2Repo 48.9, DeepSWE 46.2, ProgramBench 63.7, Terminal Bench 2.1 (Terminus-2) 81.0, FrontierSWE 74.4, PostTrainBench 34.3, SWE-Marathon 13.0, MCP-Atlas 76.8, Tool-Decathlon 48.2. Compared against Qwen3.7-Max, MiniMax M3, DeepSeek-V4-Pro, Claude Opus 4.8, GPT-5.5, and Gemini 3.1 Pro.

Deployment: SGLang (v0.5.13.post1+), vLLM (v0.23.0+), Transformers, KTransformers, Unsloth, and Ascend NPU platforms (vLLM-Ascend, xLLM, SGLang). It supports `reasoning_effort` levels for balancing performance and latency. ~717K monthly downloads.

## Key points

- FP8 version of GLM-5.2, Z.ai's 1M-context flagship MoE (glm_moe_dsa).
- IndexShare indexer reuse cuts per-token FLOPs by 2.9x at 1M context.
- MIT license (pure open, no regional limits); 753B params.
- AIME 2026: 99.2; GPQA-Diamond: 91.2; Terminal Bench 2.1: 81.0.
- SWE-bench Pro 62.1; DeepSWE 46.2; HLE 40.5 (w/ tools 54.7).
- Served via SGLang, vLLM, Transformers, KTransformers, Unsloth, Ascend NPU.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 753B |
| Architecture | MoE with dynamic sparse attention (glm_moe_dsa) |
| Context length | Solid 1M tokens |
| License | MIT |
| Precision | BF16 / F8_E4M3 / F32 (FP8 quant) |
| AIME 2026 | 99.2 |
| GPQA-Diamond | 91.2 |
| HLE / HLE w/ tools | 40.5 / 54.7 |
| SWE-bench Pro | 62.1 |
| NL2Repo | 48.9 |
| DeepSWE | 46.2 |
| Terminal Bench 2.1 | 81.0 |
| FrontierSWE (Dominance) | 74.4 |
| MCP-Atlas (Public Set) | 76.8 |
| Monthly downloads | ~717K |

## Why this source matters for the RAG

GLM-5.2-FP8 documents the official FP8 baseline of Z.ai's first 1M-context flagship, key for understanding the GLM-5.2→5.3 evolution and for NVFP4 comparisons. Its IndexShare architecture note and detailed benchmark methodology (footnotes) reduce hallucination risk about long-context performance. It anchors the RAG's coverage of long-horizon open-weight agents.
