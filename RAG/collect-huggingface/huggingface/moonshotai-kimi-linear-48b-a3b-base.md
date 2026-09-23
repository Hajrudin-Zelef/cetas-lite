---
id: collect-huggingface/huggingface/moonshotai-kimi-linear-48b-a3b-base
title: "Kimi Linear 48B-A3B Base - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Moonshot", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "attention", "benchmarks", "inference", "kv cache", "latency", "license", "memory", "mit license", "moe", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-Linear-48B-A3B-Base.md
source_anchor: ""
source_lines: [1, 54]
sha256: 319d0e70e065020bab024cea700401d88c0660cbfc6bfde7da5a95488994f5f2
---

# Kimi Linear 48B-A3B Base - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-Linear-48B-A3B-Base
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi Linear is a hybrid linear attention architecture introduced by Moonshot AI that aims to outperform traditional full attention methods across short, long, and reinforcement-learning scaling regimes. Its core component is Kimi Delta Attention (KDA), a refined version of the Gated DeltaNet (arXiv 2412.06464) that introduces a more efficient gating mechanism to optimize the use of finite-state RNN memory. The base checkpoint, Kimi-Linear-48B-A3B-Base, is a sparse Mixture-of-Experts (MoE) model with 48B total parameters and 3B activated parameters, supporting a 1M-token context length.

The architecture is hybrid, mixing KDA linear-attention layers with global MLA (multi-head latent attention) layers at a roughly 3:1 ratio, which reduces KV cache memory requirements by up to 75% and boosts decoding throughput by up to 6x for contexts as long as 1M tokens. On MMLU-Pro at 4K context the model achieves 51.0 performance with similar speed to full attention, while on RULER at 128K context it shows Pareto-optimal performance (84.3) with a 3.98x speedup; it also reaches 6.3x faster TPOT (time per output token) than MLA at 1M-token sequences. The checkpoints were trained on 5.7T tokens.

The model is released under the MIT license, is quantizable (safetensors weights in BF16/F32, 49B hub-reported size), and the KDA kernel is open-sourced in the FLA (Flash Linear Attention) library. Both a Base and an Instruct variant (Kimi-Linear-48B-A3B-Instruct) were released. Inference requires Python >= 3.10, torch >= 2.6, and fla-core >= 0.4.0, loaded via Transformers with trust_remote_code=True; deployment is supported through the latest vLLM (e.g., vllm serve with --tensor-parallel-size 4 and --max-model-len 1048576). The paper is arXiv 2510.26692. It logs about 124,559 monthly downloads.

## Key points

- Hybrid linear attention architecture built on Kimi Delta Attention (KDA), refining Gated DeltaNet with fine-grained gating.
- MoE model: 48B total / 3B activated parameters with 1M-token context length.
- 3:1 KDA-to-global-MLA ratio cuts KV cache by up to 75% and decoding latency by up to 6x.
- Trained on 5.7T tokens; open-sourced Base and Instruct checkpoints.
- Pareto-optimal RULER (128K) score 84.3 with 3.98x speedup; 6.3x faster TPOT vs MLA at 1M tokens.
- MIT license; KDA kernel open-sourced in FLA library.
- Requires python>=3.10, torch>=2.6, fla-core>=0.4.0, Transformers with trust_remote_code.
- Deployable via latest vLLM with up to 1M max-model-len.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-Linear-48B-A3B-Base |
| Architecture | Hybrid linear attention (KDA + global MLA), MoE |
| Total params | 48B |
| Activated params | 3B |
| Context length | 1M tokens |
| Training tokens | 5.7T |
| License | MIT |
| Weight formats | BF16, F32 |
| Hub-reported size | 49B params |
| KDA ratio | 3:1 KDA-to-global MLA |
| Benchmarks | MMLU-Pro 51.0 (4K), RULER 84.3 (128K, 3.98x speedup) |
| Requirements | python>=3.10, torch>=2.6, fla-core>=0.4.0 |
| Deployment | vLLM (tensor-parallel-size 4, max-model-len 1048576) |
| Paper | arXiv 2510.26692 |
| Downloads/month | 124,559 |

## Why this source matters for the RAG

This card documents a frontier open-weight linear-attention MoE architecture that challenges full-attention transformers for long-context efficiency, including hard performance figures, gating design, and deployment specifics. It is essential reference material for retrieval on linear attention, long-context MoE models, KV-cache efficiency, and open-source attention kernels.
