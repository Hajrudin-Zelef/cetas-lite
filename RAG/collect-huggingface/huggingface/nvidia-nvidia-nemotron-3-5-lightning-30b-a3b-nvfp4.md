---
id: collect-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4
title: "NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Hugging Face", "Nvidia"]
dates: ["2026-08-11", "2026-09-23"]
keywords: ["nvfp4", "nvidia", "agent", "agents", "attention", "benchmarks", "blackwell", "fine-tuning", "fp8", "gpu", "gpus", "inference"]
source: docs/RAG/Collect RAG/03_huggingface/nvidia-NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4.md
source_anchor: ""
source_lines: [1, 47]
sha256: c8c322f10f42f9d597765b38050023eb6698f981d8abc94c89e296223c267b30
---

# NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 is a 30B-parameter Mixture-of-Experts (MoE) hybrid language model from NVIDIA with 3B active parameters, using interleaved Mamba-2 and MoE layers plus select Attention layers. It is designed for long-running autonomous agents, sub-agent workhorse deployments and efficient local inference on personal hardware, supporting context lengths up to 1M tokens. Supported languages are English (and coding languages), Spanish, French, German, Italian and Japanese. The model was pre-trained with over 20T tokens using an NVFP4 recipe, followed by continued pre-training of Multi-Token Prediction (MTP) layers, supervised fine-tuning, GRPO reinforcement learning, and finally post-training quantization (Four Over Six NVFP4 W4A16 on experts, FP8 per-tensor dynamic scales on Mamba projections and KV cache). It ships alongside several speculative decoding methods: DSpark (recommended default for DGX Spark and low-concurrency data-centre serving), DFlash, and MTP. It runs on a single DGX Spark (GB10) or a single H100, and is validated on Blackwell (GB200, RTX 5090) and Hopper (H100/H200), with W4A16 support extending to Ampere GPUs. Recommended sampling is temperature 1.0 and top_p 0.95. Benchmarks (NVFP4 vs BF16): MMLU Pro 81.62 / 81.94, GPQA Diamond 75.57 / 75.44, HLE 10.47 / 11.72, SciCode 31.38 / 32.60, SWE-bench Verified 52.80 / 51.56, Terminal-Bench 2.1 23.46 / 24.58, IFBench 72.88 / 71.88, AA-LCR 49.19 / 52.00. It is released under the OpenMDW License Agreement version 1.1 and is ready for commercial use. Release date August 11, 2026.

## Key points

- 30B total / 3B active MoE hybrid (Mamba-2 + MoE + Attention) with MTP.
- Context up to 1M tokens; six supported languages; NVFP4 checkpoint (W4A16 on some hardware).
- Ships with DSpark, DFlash and MTP speculative decoding; DSpark recommended default.
- Single-GPU deployment on DGX Spark (GB10) or H100; also GB200, RTX 5090, A100 (W4A16).
- Trained on 20T+ tokens with NVFP4 recipe; 5-stage pipeline including GRPO RL and PTQ.
- Benchmarks: MMLU Pro 81.62, SWE-bench Verified 52.80, GPQA Diamond 75.57, IFBench 72.88.
- License: OpenMDW-1.1; release date 2026-08-11.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 30B |
| Active parameters | 3B |
| Architecture | MoE hybrid (Mamba-2 + MoE + Attention) + MTP |
| Context length | Up to 1M tokens |
| Languages | English, Spanish, French, German, Italian, Japanese |
| Precision | NVFP4 (W4A16 on Hopper/Ampere) |
| Single-GPU deployment | 1× DGX Spark (GB10) or 1× H100 |
| Speculative decoding | DSpark (default), DFlash, MTP |
| MMLU Pro | 81.62 |
| SWE-bench Verified | 52.80 |
| GPQA Diamond | 75.57 |
| IFBench (loose) | 72.88 |
| License | OpenMDW-1.1 |
| Release date | 2026-08-11 |

## Why this source matters for the RAG

This card documents NVIDIA's compact agent-focused 3.5 Lightning model and its advanced speculative decoding stack, providing key data on efficient local deployment for autonomous agents. It is important for RAG systems that need low-latency, long-context local inference.
