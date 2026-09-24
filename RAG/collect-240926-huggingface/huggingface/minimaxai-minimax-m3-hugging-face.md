---
id: collect-240926-huggingface/huggingface/minimaxai-minimax-m3-hugging-face
title: "minimaxai-minimax-m3-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "MiniMax", "SGLang", "vLLM"]
dates: []
keywords: ["agentic", "agents", "attention", "benchmarks", "compute", "decode", "gqa", "inference", "latency", "leaderboard", "memory", "multimodal"]
source: docs/RAG/clean_en/huggingface/minimaxai-minimax-m3-hugging-face.md
source_anchor: ""
source_lines: [1, 58]
sha256: 3ffbe8ef7b1b45e86025d46120290e694006f607bf2904cc698d467958a365f7
---

# minimaxai-minimax-m3-hugging-face

<!-- source: https://huggingface.co/MiniMaxAI/MiniMax-M3 -->

MiniMax-M3 is a native multimodal model with 1M context. It has ~428B parameters and ~23B activated parameters.

**Highlights:**

- **Native Multimodality:** M3 undergoes mixed-modality training from the very first step, enabling deeper semantic fusion across text, image, and video.
- **Context Scaling via Sparse Attention:** M3 introduces MiniMax Sparse Attention (MSA) to improve long context efficiency. M3 delivers 9× prefill and 15× decode speedups compared to M2 at 1M context, reducing per-token compute to 1/20.
- **Coding & Cowork Capability:** M3 achieves frontier-level performance across long-horizon agentic benchmarks, excelling in both coding and cowork.


M3 is powered by **MiniMax Sparse Attention (MSA)**, a high-performance sparse attention operator designed for million-token contexts. Compared with GQA, MSA dramatically reduces the attention compute and memory footprint while preserving model quality.


📄 Read the technical report: arXiv:2606.13392 · Hugging Face Papers


M3 supports three reasoning modes through the `thinking` parameter:

- **`enabled`** — Reasoning is always enabled.
- **`adaptive`** — M3 automatically determines when additional reasoning is beneficial.
- **`disabled`** — Reasoning is disabled to minimize latency and maximize throughput.

Download the model:

```
hf download MiniMaxAI/MiniMax-M3 --local-dir MiniMax-M3
```
We recommend the following inference frameworks to serve the model:

- SGLang - see SGLang cookbook.
- vLLM - see vLLM recipes.
- Transformers - see Transformers docs.
- KTransformers - see KTransformers MiniMax-M3 tutorial.
- unsloth - see tutorial
- ATOM - see MiniMax-M3 MXFP4/MXFP8 Usage Guide

We recommend the following parameters for best performance: `temperature=1.0`, `top_p=0.95`.

Contact us at model@minimax.io.

- Downloads last month
- 167,808

## Spaces using MiniMaxAI/MiniMax-M3 52

## Collection including MiniMaxAI/MiniMax-M3

## Paper for MiniMaxAI/MiniMax-M3

- mercor/apex-agents · Apex Agents View evaluation results leaderboard
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results leaderboard
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results    leaderboard  59<sup>*</sup>
- IntelligenceLab/Long-Horizon-Terminal-Bench leaderboard
- Lhtb View evaluation results source
- Lhtb Solved View evaluation results source
- MME-Benchmarks/Video-MME-v2 · Video Mme V2 View evaluation results leaderboard
- MMMU/MMMU_Pro · Mmmu Pro Standard 10 Options View evaluation results leaderboard
