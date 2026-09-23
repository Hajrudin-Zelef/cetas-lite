---
id: collect-huggingface/huggingface/qwen-qwen3-5-122b-a10b-fp8
title: "Qwen3.5-122B-A10B-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "qwen", "agent", "agentic", "agents", "apache", "attention", "benchmark", "cost", "license", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.5-122B-A10B-FP8.md
source_anchor: ""
source_lines: [1, 52]
sha256: d3115153a2e89d8ec5f9db98896ca1fd070d7f5d3931bc62f39668a411d35198
---

# Qwen3.5-122B-A10B-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.5-122B-A10B-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.5-122B-A10B-FP8 is the FP8-quantized release of Qwen3.5-122B-A10B, a flagship multimodal model of the Qwen3.5 generation. It is a causal language model with a vision encoder (image-text-to-text), integrating breakthroughs in multimodal learning, architectural efficiency, reinforcement learning at scale, and global accessibility. The FP8 weights use fine-grained quantization with block size 128, with performance nearly identical to the original; the official FP8 checkpoint is also the baseline used by NVIDIA for its NVFP4 quant.

Architecture: 122B total parameters with 10B activated, hidden dimension 3072, padded vocabulary 248,320, 48 layers, hybrid layout of 12 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)). Gated DeltaNet: 64 V heads, 16 QK heads. Gated Attention: 32 Q heads, 2 KV heads, RoPE dim 64. MoE: 256 experts, 8 routed + 1 shared, expert intermediate dimension 1024. Multi-step MTP. Context length is 262,144 tokens natively, extensible to 1,010,000 tokens via YaRN.

Highlights: unified vision-language foundation with early fusion training, efficient hybrid Gated Delta Network + sparse MoE, scalable RL generalization across million-agent environments, support for 201 languages and dialects, and near-100% multimodal training efficiency. It operates in thinking mode by default (can be disabled). Benchmark highlights: MMLU-Pro 86.7, GPQA Diamond 86.6, SWE-bench Verified 72.0, Terminal-Bench 2 49.4, TAU2-Bench 79.5, MMMU 83.9, OmniDocBench1.5 89.8, VideoMME (w/sub) 87.3. License is Apache-2.0.

## Key points

- FP8 MoE multimodal model: 122B total / 10B activated.
- 48 layers, 256 experts (8 routed + 1 shared); Gated DeltaNet + Gated Attention.
- Native context 262,144, extensible to 1,010,000 via YaRN.
- Unified vision-language foundation; supports 201 languages.
- Thinking mode by default; strong agentic and tool-calling performance.
- Apache-2.0; served via SGLang, vLLM, KTransformers, Transformers.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 122B |
| Active parameters | 10B |
| Architecture | Hybrid Gated DeltaNet + Gated Attention + MoE + Vision Encoder |
| Layers | 48 |
| Hidden dimension | 3072 |
| Experts | 256 (8 routed + 1 shared) |
| Context length | 262,144 native; up to 1,010,000 |
| Vocabulary | 248,320 (padded) |
| Languages | 201 |
| License | Apache-2.0 |
| Quantization | FP8 (fine-grained, block 128) |
| MMLU-Pro | 86.7 |
| GPQA Diamond | 86.6 |
| SWE-bench Verified | 72.0 |
| MMMU | 83.9 |
| OmniDocBench1.5 | 89.8 |

## Why this source matters for the RAG

Qwen3.5-122B-A10B-FP8 is the official FP8 baseline for one of the largest open-weight multimodal MoE models, central to understanding the Qwen3.5 family and its quantization chain. Its broad benchmark suite across knowledge, coding, agents, vision, and multilingual tasks makes it a strong reference point. It anchors cost/performance comparisons with the NVIDIA NVFP4 variant also covered in this knowledge base.
