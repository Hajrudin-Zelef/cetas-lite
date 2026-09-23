---
id: collect-mindstudio/mindstudio/run-qwen3-8-locally-hardware
title: "Can You Run Qwen3.8-2.4T-A95B Locally? Hardware Requirements Explained"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "attention", "compute", "consumer", "cost", "cost per token", "gpu", "gpus", "inference", "int4", "kv cache", "memory"]
source: docs/RAG/Collect RAG/02_mindstudio/run-qwen3-8-locally-hardware.md
source_anchor: ""
source_lines: [1, 52]
sha256: b8fd2aa3c2346326c09f0093b492fad211668c72fa922d6724b7d8851f22f1ea
---

# Can You Run Qwen3.8-2.4T-A95B Locally? Hardware Requirements Explained

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-qwen3-8-locally-hardware
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains what it actually takes to self-host **Qwen3.8-2.4T-A95B**, a 2.4-trillion-parameter mixture-of-experts (MoE) model with 95 billion parameters active per token. The short answer is that it is technically possible but practically reserved for multi-node GPU clusters. Even at aggressive quantization, the full weight set does not fit on a single consumer GPU or even a single high-end workstation; it is built for serving infrastructure using **vLLM** or **SGLang**.

The MoE architecture uses **512 total experts**, of which **10 routed experts plus 1 shared expert** activate per token — hence "A95B" (95B active of 2.4T total). The architecture interleaves **Gated DeltaNet** (a linear attention mechanism) with full Gated Attention layers across **92 layers**, following a repeating pattern of three DeltaNet+MoE blocks followed by one Gated Attention+MoE block, repeated 23 times. This creates a split personality: compute cost per token tracks the active count (~95B, manageable), but memory footprint tracks the total count (2.4T), because every expert must sit in memory (VRAM, host RAM, or fast storage) ready to be routed. In practice, paging experts from disk is too slow for interactive inference, so all 2.4T parameters need GPU memory or sharding across many GPUs.

VRAM math: at FP16/BF16 (2 bytes/parameter), 2.4T parameters require roughly **4.8 TB** just for weights, before KV cache, activations, or the multi-token prediction (MTP) head. Even at INT4 (~0.5 bytes/parameter with overhead), weights alone land in the **600 GB to 1 TB** range. No single GPU holds that much; even an 8-GPU server with 80 GB cards (640 GB total) would be tight or insufficient depending on quantization and KV cache needs. Realistic deployments use multiple multi-GPU nodes with tensor and expert parallelism splitting the 512 experts.

Context length compounds the problem: the model natively supports **262,144 tokens**, extensible to just over **1,000,000**. KV cache scales with context length, batch size, and attention heads. The Gated Attention uses **64 query heads and 4 key/value heads** at head dimension **256** — a grouped-query-style ratio that reduces KV cache size relative to full multi-head attention, but at million-token contexts the cache still adds meaningfully to the memory budget.

The model card explicitly lists **vLLM, SGLang, and TokenSpeed** as intended serving paths, supporting tensor, pipeline, and expert parallelism. Raw Transformers on a single machine is not realistic. For most developers, self-hosting is not worth it; the model card points to **Qwen Cloud's Qwen3.8-Max**, a hosted version on the same architecture that adds vision input, non-thinking mode, a default 1M-token context, and built-in tools. Self-hosting makes sense only for organizations with existing multi-node GPU clusters, strict data-residency/privacy requirements, or heavy customization needs.

## Key points

- Total parameters: 2.4T; only ~95B active per token via MoE routing across 512 experts (10 routed + 1 shared).
- Weights require terabytes of memory at full precision; hundreds of GB to ~1 TB even at INT4.
- Active-parameter compute cost resembles a dense ~95B model, so per-token speed is closer to that than a "2.4T" label suggests.
- vLLM and SGLang are the supported serving stacks, using tensor/expert parallelism across GPUs or nodes.
- Native context is 262,144 tokens, extensible to ~1,010,000, adding major KV cache overhead.
- Quantization helps but does not change the order of magnitude; still needs a multi-GPU cluster.
- Qwen Cloud's Qwen3.8-Max is the recommended access path for most users.

## Technical data / figures

| Spec | Value |
|---|---|
| Total parameters | 2.4 trillion (MoE) |
| Active parameters per token | ~95B |
| Experts | 512 total; 10 routed + 1 shared per token |
| Layers | 92 (3x DeltaNet+MoE : 1x Gated Attention+MoE, repeated 23x) |
| FP16/BF16 weight memory | ~4.8 TB |
| INT4 weight memory | ~600 GB – 1 TB |
| Native context | 262,144 tokens |
| Extended context | ~1,010,000 tokens |
| Attention | 64 query heads, 4 KV heads, head dim 256 |
| Serving engines | vLLM, SGLang, TokenSpeed |
| Hosted alternative | Qwen3.8-Max (Qwen Cloud) |

## Why this source matters for the RAG

It provides a clear, math-driven reality check on the hardware requirements of a 2.4T-parameter MoE model, useful for answering "can I run this locally?" with quantitative VRAM and context-length analysis. It also distinguishes compute-active parameters from total memory footprint, a core concept for understanding MoE deployment economics.
