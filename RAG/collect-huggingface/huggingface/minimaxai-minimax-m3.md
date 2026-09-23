---
id: collect-huggingface/huggingface/minimaxai-minimax-m3
title: "MiniMax-M3 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["AMD", "Hugging Face", "MiniMax", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agentic", "agents", "attention", "benchmarks", "compute", "decode", "gqa", "inference", "latency", "license", "memory", "moe"]
source: docs/RAG/Collect RAG/03_huggingface/MiniMaxAI-MiniMax-M3.md
source_anchor: ""
source_lines: [1, 55]
sha256: b045eb9401f3b2ec62630e3fd51341f4d910207d4c6ab6774a7064a70801dcc0
---

# MiniMax-M3 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/MiniMaxAI/MiniMax-M3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiniMax-M3 is a native multimodal model from MiniMax with 1M-token context, approximately 428B total parameters and ~23B activated parameters. It is a Mixture-of-Experts (MoE) vision-language model (Image-Text-to-Text, Transformers, custom code) that undergoes mixed-modality training from the very first step, enabling deeper semantic fusion across text, image, and video.

Key highlight: M3 introduces MiniMax Sparse Attention (MSA), a high-performance sparse attention operator for million-token contexts. Compared with GQA, MSA dramatically reduces attention compute and memory while preserving quality. M3 delivers 9x prefill and 15x decode speedups versus M2 at 1M context, reducing per-token compute to 1/20. It achieves frontier-level performance across long-horizon agentic benchmarks, excelling in both coding and cowork (agentic collaboration).

M3 supports three reasoning modes via the `thinking` parameter: `enabled` (reasoning always on), `adaptive` (model decides when reasoning is beneficial), and `disabled` (minimum latency / maximum throughput). Recommended inference parameters: temperature=1.0, top_p=0.95.

Deployment: Hugging Face Transformers (AutoProcessor/AutoModelForMultimodalLM with trust_remote_code), SGLang, vLLM, KTransformers, unsloth, and ROCm ATOM (with MXFP4/MXFP8 usage guide). Hub results: Apex Agents 27.7, SWE-Bench Verified 80.5, SWE-Bench Pro 59, LHTB mean reward 38.5, MMMU Pro 78.1, Video-MME v2 85.4. License: minimax-community (MiniMax Community License). Weights in BF16/F32, 427B hub size. Paper: arXiv 2606.13392 (MiniMax Sparse Attention). 58 community quantizations and 14 finetunes.

## Key points

- Native multimodal MoE: ~428B total / ~23B activated, 1M-token context.
- Mixed-modality training from the first step across text, image, video.
- MiniMax Sparse Attention (MSA): 9x prefill and 15x decode speedup vs M2 at 1M context; per-token compute 1/20.
- Three reasoning modes: enabled / adaptive / disabled via `thinking` parameter.
- Frontier agentic performance in coding and cowork; SWE-Bench Verified 80.5.
- MiniMax Community License; BF16/F32 weights.
- Deploy via Transformers, SGLang, vLLM, KTransformers, unsloth, ROCm ATOM.
- Paper: arXiv 2606.13392.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | MiniMaxAI |
| Model name | MiniMax-M3 |
| Architecture | MoE multimodal (text/image/video) |
| Total params | ~428B (hub: 427B) |
| Activated params | ~23B |
| Context length | 1M tokens |
| Attention | MiniMax Sparse Attention (MSA) |
| Efficiency | 9x prefill / 15x decode vs M2 at 1M; 1/20 per-token compute |
| Reasoning modes | enabled / adaptive / disabled |
| License | MiniMax Community License |
| Weight formats | BF16, F32 |
| Benchmarks | SWE-Bench Verified 80.5, SWE-Bench Pro 59, MMMU Pro 78.1, Video-MME v2 85.4 |
| Sampling | temperature=1.0, top_p=0.95 |
| Paper | arXiv 2606.13392 |
| Downloads/month | 167,692 |

## Why this source matters for the RAG

This card documents MiniMax's flagship open-weight multimodal MoE and its sparse attention design for million-token contexts, including detailed speedup figures and reasoning-mode controls. It is essential reference for retrieval on sparse attention, long-context multimodal models, and MoE efficiency.
