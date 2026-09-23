---
id: collect-huggingface/huggingface/qwen-qwen3-6-35b-a3b
title: "Qwen3.6-35B-A3B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-04", "2026-09-23"]
keywords: ["qwen", "agentic", "apache", "attention", "benchmark", "benchmarks", "fp8", "gguf", "license", "llama", "llama.cpp", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.6-35B-A3B.md
source_anchor: ""
source_lines: [1, 47]
sha256: b24ca0557fa5e892f1880de23031de52a401b328542f9693ca63687db52b0356
---

# Qwen3.6-35B-A3B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.6-35B-A3B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.6-35B-A3B is the first open-weight variant of the Qwen3.6 series, released by the Qwen team in April 2026. It is a causal vision-language model (image-text-to-text) built on a sparse Mixture-of-Experts (MoE) architecture with 35B total parameters and only ~3B activated per token, making it extremely efficient to serve. The model emphasizes agentic coding, frontend workflows, repository-level reasoning, and real-world utility, and introduces a "thinking preservation" option that retains reasoning context from historical messages to streamline iterative development and reduce token overhead. Architecturally it has a hidden dimension of 2048, 40 layers arranged as 10 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)), 256 experts with 8 routed + 1 shared active, expert intermediate dimension 512, and a padded vocabulary/output size of 248,320. It supports text, image, and video input, has a native context length of 262,144 tokens (extensible to 1,010,000 with YaRN), and is trained with multi-step MTP (multi-token prediction). Benchmarks show SWE-bench Verified 73.4, SWE-bench Multilingual 67.2, SWE-bench Pro 49.5, Terminal-Bench 2.0 51.5, GPQA 86.0, AIME26 92.7, LiveCodeBench v6 80.4, MMLU-Pro 85.2, plus vision scores such as MMMU 81.7, MMMU-Pro 75.3, RealWorldQA 85.3, OmniDocBench1.5 89.9, and VideoMMMU 83.7. It is released under Apache 2.0 and is compatible with Transformers, vLLM, SGLang, and KTransformers.

## Key points

- MoE vision-language model: 35B total / ~3B active parameters, 40 layers, 256 experts (8 routed + 1 shared).
- Native 262,144-token context, extensible to ~1,010,000 tokens via YaRN RoPE scaling.
- Agentic coding focus with "thinking preservation" for retaining reasoning across historical turns.
- Multimodal: text, image, and video input; padded vocabulary of 248,320.
- Strong agentic/coding benchmarks: SWE-bench Verified 73.4, Terminal-Bench 2.0 51.5, AIME26 92.7, GPQA 86.0.
- Apache 2.0 license; supports SGLang, vLLM, KTransformers, and Transformers serving.
- Recommended sampling: general thinking temp 1.0/top_p 0.95; coding temp 0.6; non-thinking temp 0.7/top_p 0.80.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 35B |
| Active parameters | 3B |
| Architecture | MoE, causal LM with vision encoder |
| Hidden dimension | 2048 |
| Layers | 40 |
| Experts | 256 total; 8 routed + 1 shared active |
| Expert intermediate dim | 512 |
| Vocabulary | 248,320 (padded) |
| Context length | 262,144 native; up to 1,010,000 with YaRN |
| Modalities | Text, Image, Video |
| License | Apache 2.0 |
| Quantization | BF16 base; community GGUF (llama.cpp/LM Studio/Jan/Ollama); FP8 checkpoint |
| MTP | Trained with multi-steps |
| Downloads last month | 3,175,372 |

## Why this source matters for the RAG

This card documents a frontier open-weight sparse MoE multimodal model that competes directly with Gemma 4 and serves as a key comparison point for coding, agentic and long-context workloads. Its detailed architecture and benchmark tables provide high-value, structured data for retrieval and model-selection tasks.
