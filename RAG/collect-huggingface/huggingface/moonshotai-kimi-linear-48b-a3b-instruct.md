---
id: collect-huggingface/huggingface/moonshotai-kimi-linear-48b-a3b-instruct
title: "Kimi Linear 48B-A3B Instruct - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Moonshot", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "attention", "benchmarks", "context window", "decode", "inference", "kv cache", "license", "memory", "mit license", "moe", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-Linear-48B-A3B-Instruct.md
source_anchor: ""
source_lines: [1, 52]
sha256: 1a0801974ea1801e87669b76c2a8372dbff9a5489565e349f895e73cb820a429
---

# Kimi Linear 48B-A3B Instruct - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-Linear-48B-A3B-Instruct
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi Linear 48B-A3B-Instruct is Moonshot AI's experimental MoE model demonstrating the Kimi Linear hybrid linear-attention architecture (arXiv 2510.26692, "Kimi Linear: An Expressive, Efficient Attention Architecture"). It has 48B total parameters with 3B activated, a 1M-token context window, and is released under the MIT license with ~183K downloads/month. The architecture's core is Kimi Delta Attention (KDA), a refined version of Gated DeltaNet (arXiv 2412.06464) that introduces a more efficient gating mechanism to optimize the use of finite-state RNN memory. The model is a 3:1 hybrid of KDA to global MLA layers, which reduces memory usage while maintaining or surpassing full-attention quality. Two checkpoints were released (Base and Instruct), both trained on 5.7T tokens.

Performance/hardware efficiency: Kimi Linear reduces KV cache requirements by up to 75% and boosts decoding throughput by up to 6x for contexts as long as 1M tokens. On MMLU-Pro at 4K context it achieves 51.0 with speed similar to full attention; on RULER at 128K it shows Pareto-optimal performance (84.3) with a 3.98x speedup; and it achieves a 6.3x faster time-per-output-token (TPOT) compared to MLA at long sequence lengths (1M tokens). The KDA kernel is open-sourced in FLA (flash-linear-attention, `fla/ops/kda`).

Usage: requires `python >= 3.10`, `torch >= 2.6`, and `fla-core >= 0.4.0` (`pip install -U fla-core`); load via `AutoModelForCausalLM.from_pretrained(..., trust_remote_code=True)`. Deployment example uses vLLM: `vllm serve moonshotai/Kimi-Linear-48B-A3B-Instruct --port 8000 --tensor-parallel-size 4 --max-model-len 1048576 --trust-remote-code`. It is a conversational text-generation model (architecture tag `kimi_linear`), with 9 fine-tunes and 24 quantizations downstream. Citation is the Kimi Linear paper by Zhang et al., 2025.

## Key points

- Experimental 48B-A3B MoE showcasing the Kimi Linear hybrid linear-attention architecture.
- Core: Kimi Delta Attention (KDA), a refined Gated DeltaNet with fine-grained gating.
- 3:1 KDA-to-global-MLA hybrid; 1M context; trained on 5.7T tokens.
- Up to 75% smaller KV cache and up to 6x faster decoding at 1M tokens.
- RULER 128K: 84.3 with 3.98x speedup; 6.3x faster TPOT vs MLA at 1M.
- MIT license; KDA kernel open-sourced in FLA.
- Requires fla-core >=0.4.0; vLLM deployment with 1M max-model-len.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-Linear-48B-A3B-Instruct |
| Architecture | MoE with hybrid linear attention (3:1 KDA : global MLA) |
| Total params | 48B |
| Active params | 3B |
| Hub model size | 49B params |
| Context length | 1M (1,048,576) |
| Training tokens | 5.7T |
| Core mechanism | Kimi Delta Attention (refined Gated DeltaNet) |
| License | MIT |
| Efficiency | KV cache −75%; decode up to 6x; TPOT 6.3x vs MLA (1M ctx) |
| Benchmarks | MMLU-Pro (4k) 51.0; RULER (128k) 84.3 @ 3.98x speedup |
| Requirements | python ≥3.10, torch ≥2.6, fla-core ≥0.4.0 |
| Paper | arXiv 2510.26692 (also 2412.06464) |
| Downloads/month | ~183,165 |

## Why this source matters for the RAG

This card documents Moonshot's open experimental linear-attention architecture (KDA), providing concrete efficiency numbers (KV-cache reduction, decode speedup, TPOT) and a 1M-context small-MoE deployment recipe. It is a valuable, citable source on hybrid linear attention and efficient long-context inference for the RAG.
