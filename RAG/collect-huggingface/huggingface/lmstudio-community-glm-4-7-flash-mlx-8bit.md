---
id: collect-huggingface/huggingface/lmstudio-community-glm-4-7-flash-mlx-8bit
title: "GLM-4.7-Flash-MLX-8bit - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Apple", "Hugging Face", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["glm", "agent", "agentic", "benchmarks", "consumer", "inference", "license", "mit license", "moe", "parameters", "quantization", "qwen"]
source: docs/RAG/Collect RAG/03_huggingface/lmstudio-community-GLM-4.7-Flash-MLX-8bit.md
source_anchor: ""
source_lines: [1, 49]
sha256: 29875af06f5ccd70e214439186b1b0e2d35a27881909ef77c0ed986ef88431b3
---

# GLM-4.7-Flash-MLX-8bit - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/lmstudio-community/GLM-4.7-Flash-MLX-8bit
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

lmstudio-community/GLM-4.7-Flash-MLX-8bit is an 8-bit MLX quantized version of zai-org/GLM-4.7-Flash, produced by the LM Studio team using mlx_lm. It is part of the LM Studio Community Models highlight program and is optimized for Apple Silicon, enabling local inference of the 30B-A3B MoE model on Mac hardware. The MLX variant is approximately 31.8 GB on disk.

The base model, GLM-4.7-Flash, is a 30B-A3B Mixture-of-Experts language model from Z.ai, positioned as the strongest model in the 30B class and a lightweight deployment option balancing performance and efficiency. It is MIT-licensed. Benchmarks: AIME 25 91.6, GPQA 75.2, LiveCodeBench v6 64.0, HLE 14.4, SWE-bench Verified 59.2, tau2-Bench 79.5, BrowseComp 42.8, leading the 30B class on agentic and reasoning tasks. It supports a "Preserved Thinking" mode for multi-turn agentic tasks.

Because this is an MLX quant, it runs via mlx-lm (load, generate, chat REPL, or OpenAI-compatible server), LM Studio, and other MLX tools on Apple Silicon. It can also be served via vLLM, SGLang, or Docker Model Runner. The base card notes deployment via vLLM/SGLang with `--tool-call-parser glm47` and `--reasoning-parser glm45`. Reported model size is 30B params with U32/BF16/F32 tensors. License is MIT (inherited from the base model). ~212K monthly downloads.

## Key points

- 8-bit MLX quantization of GLM-4.7-Flash, optimized for Apple Silicon; ~31.8 GB.
- Base model: 30B-A3B MoE; "strongest in the 30B class" (MIT license).
- AIME 25 91.6; SWE-bench Verified 59.2; tau2-Bench 79.5; BrowseComp 42.8.
- Supports Preserved Thinking mode for multi-turn agent tasks.
- Runs via mlx-lm, LM Studio, vLLM, SGLang, Docker Model Runner.
- glm47 tool-call parser / glm45 reasoning parser.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 30B (MoE, glm4_moe_lite) |
| Active parameters | 3B (A3B) |
| Quantization | 8-bit MLX |
| File size | ~31.8 GB |
| License | MIT |
| Precision (tensors) | U32 / BF16 / F32 |
| AIME 25 | 91.6 |
| GPQA | 75.2 |
| LiveCodeBench v6 | 64.0 |
| SWE-bench Verified | 59.2 |
| tau2-Bench | 79.5 |
| BrowseComp | 42.8 |
| Monthly downloads | ~212K |

## Why this source matters for the RAG

This card documents local deployment of a strong 30B-class MoE model on Apple Silicon via MLX, covering both performance specs and tooling. It shows the GLM family is reachable for local, privacy-friendly inference, complementing the Qwen MLX entry. It supports questions about running frontier-class reasoning models on consumer Mac hardware.
