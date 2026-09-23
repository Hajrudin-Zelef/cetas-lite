---
id: collect-huggingface/huggingface/xiaomimimo-mimo-v2-5-pro
title: "MiMo-V2.5-Pro - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "OpenAI", "SGLang", "Xiaomi", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "attention", "benchmark", "benchmarks", "distillation", "fp8", "gqa", "inference", "license", "mit license", "moe", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/XiaomiMiMo-MiMo-V2.5-Pro.md
source_anchor: ""
source_lines: [1, 59]
sha256: 3b49b4fd122dbee0d647471488598bdce2cda70c64c7fcd16673c7083d4e236b
---

# MiMo-V2.5-Pro - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/XiaomiMiMo/MiMo-V2.5-Pro
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiMo-V2.5-Pro is an open-source Mixture-of-Experts (MoE) language model from Xiaomi MiMo, with 1.02T total parameters and 42B active parameters, up to 1M-token context length. It uses the hybrid attention architecture and 3-layer Multi-Token Prediction (MTP) introduced in MiMo-V2-Flash, designed for the most demanding agentic, complex software engineering, and long-horizon tasks, sustaining thousands of tool calls with strong instruction following and coherence. Architecture: 70 layers (1 dense + 69 MoE), hidden size 6144, 128 attention heads, 8 KV heads (GQA), head dims QK/V = 192/128, 384 routed experts (8 per token, MoE intermediate 2048), dense intermediate 16384 (layer 0), SWA window 128, 60 SWA + 10 GA layers (6:1 ratio), 3 MTP layers with dense FFNs.

Training: pre-trained on 27T tokens with FP8 mixed precision and native 32k sequence length; the 6:1 SWA/GA interleave reduces KV-cache storage by nearly 7x with a learnable attention-sink bias; post-training uses a three-stage paradigm — SFT, Domain-Specialized RL (math, safety, agentic tool-use teachers), and Multi-Teacher On-Policy Distillation (MOPD). MTP triples output speed during inference and accelerates RL rollout.

Benchmarks (base): BBH 88.4, MMLU 89.4, MMLU-Pro 68.5, DROP 86.3, ARC-C 97.2, HellaSwag 89.8, GPQA-Diamond 66.7, GSM8K 99.6, MATH 86.2, AIME 24&25 37.3, HumanEval+ 75.6, MBPP+ 74.1, LiveCodeBench v6 39.6, C-Eval 91.5, GlobalMMLU 83.6. Long-context: on GraphWalks (OpenAI BFS benchmark) it scores 0.56 BFS / 0.92 Parents at 512k and 0.37/0.62 at 1M, vs V2 Pro collapsing to 0.00 at 1M. Post-training results: SWE-Bench Verified 78.9, SWE-Bench Pro 57.2.

License: MIT. Weights in FP8 (E4M3) mixed precision. Deployment via SGLang (EAGLE speculative decoding, deepep) and vLLM with sampling temperature=1.0, top_p=0.95. Available on Hugging Face, ModelScope, AI Studio, and the Xiaomi MiMo API platform; DeepInfra hosts it as an inference provider.

## Key points

- Open-source MoE: 1.02T total / 42B active, 1M-token context.
- Hybrid SWA/GA attention (6:1, window 128) reduces KV-cache ~7x with attention-sink bias.
- 3-layer MTP modules triple output speed and accelerate RL rollout.
- Pre-trained on 27T tokens with FP8 mixed precision at native 32k seq length.
- Three-stage post-training: SFT, domain-specialized RL, Multi-Teacher On-Policy Distillation.
- Strong long-context GraphWalks results at 512k/1M where V2 Pro collapses.
- MIT license; FP8 (E4M3) mixed weights.
- Deploy via SGLang/vLLM; recommended temp 1.0, top_p 0.95.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | XiaomiMiMo |
| Model name | MiMo-V2.5-Pro |
| Architecture | Sparse MoE, hybrid SWA/GA + MTP |
| Total params | 1.02T |
| Activated params | 42B |
| Layers | 70 (1 dense + 69 MoE) |
| Hidden size | 6144 |
| Attention heads / KV heads | 128 / 8 (GQA) |
| Routed experts | 384 (8 per token) |
| SWA / GA layers | 60 / 10 |
| SWA window | 128 |
| MTP layers | 3 (dense FFN) |
| Context length | 1M tokens |
| Training tokens | 27T (FP8 mixed) |
| License | MIT |
| Weight precision | FP8 (E4M3) mixed |
| Benchmarks | MMLU 89.4, GSM8K 99.6, GPQA-D 66.7, SWE-Bench Verified 78.9 |
| Sampling | temperature=1.0, top_p=0.95 |
| Downloads/month | 22,213 |

## Why this source matters for the RAG

This card documents Xiaomi's flagship open-weight MoE model and its hybrid SWA/GA + MTP design that sustains 1M-token contexts efficiently, plus a three-stage RL post-training pipeline. It is critical reference for retrieval on long-context MoE architectures, multi-token prediction, KV-cache efficiency, and agentic post-training.
