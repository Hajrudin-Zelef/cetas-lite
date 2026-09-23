---
id: collect-mindstudio/mindstudio/run-fuse-1-lite-locally
title: "How to Run fuse-1 Lite Locally: VRAM, Setup, and Formats"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "vLLM"]
dates: ["2026-09-23"]
keywords: ["alignment", "compute", "consumer", "cost", "distillation", "fine-tuning", "gguf", "gpu", "llama", "llama.cpp", "memory", "moe"]
source: docs/RAG/Collect RAG/02_mindstudio/run-fuse-1-lite-locally.md
source_anchor: ""
source_lines: [1, 51]
sha256: faa140839314228c468ab82b2a514cd94a484123216ed4f9dd50e50d7142ab50
---

# How to Run fuse-1 Lite Locally: VRAM, Setup, and Formats

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-fuse-1-lite-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide covers running **fuse-1 Lite**, a **5.72B parameter coding model** that fits comfortably on consumer hardware, with VRAM needs from about **3.36 GB in 4-bit quantized form** up to roughly **12 GB in full bfloat16 precision**. It combines a small **2.6B host model (LiquidAI's LFM2.5-2.6B)** with **960 coding-specialized experts** pulled from a much larger Qwen model, preserving coding capability while keeping the footprint small. It is available across five deployment formats: native transformers, bitsandbytes quantization, GGUF, MLX, and vLLM.

fuse-1 Lite is built differently from typical fine-tunes. Instead of distillation or fine-tuning, it takes actual expert weights from **Qwen3.6-35B-A3B** (a MoE model) and grafts **960 experts (32 per layer across 30 layers)** directly onto LFM2.5-2.6B's decoder layers as frozen residual additions. A small router (**~2 million trainable parameters**, including a per-layer scale factor) is trained to decide which transplanted experts fire for a given token. The host model and experts stay frozen; only the router and scale values are updated. Training took just **300 steps (~8 minutes on an L4 GPU)** on **55 examples**, at a reported total cost of about **$3**. The model behaves like LFM2.5 most of the time and only activates coding experts for code-relevant tokens; experts are effectively disabled (scale near zero) in **19 of 30 layers**, concentrating activity in layers tied to algorithmic reasoning, logic flow, and code synthesis, peaking at **layer 19**.

VRAM by backend: transformers bfloat16 ~12 GB (L4, A10G, RTX 4090); transformers 8-bit 6.00 GB (T4, L4, RTX 3060); transformers 4-bit NF4 3.36 GB (T4, RTX 3060, M2 Pro); vLLM bfloat16 ~12 GB (A10G, A100, H100); MLX float16 ~12 GB (M1 Pro+, M2, M3, M4); llama.cpp F16 GGUF ~11.4 GB; llama.cpp Q4_K_M GGUF ~4 GB.

**bitsandbytes** is the easiest entry point: standard transformers/torch/bitsandbytes with no custom builds, using 4-bit NF4 double-quantized with bfloat16 compute dtype. **GGUF and MLX require custom code**: the GGUF release uses a custom `fuse3` architecture that stock llama.cpp cannot load, requiring a patched fork with a C++ graph builder (`src/models/fuse3.cpp`) and a Python converter; MLX needs a custom `fuse3_mlx.py` file. **vLLM** support comes via a plugin that registers a `Fuse3ForCausalLM` class through the ModelRegistry, requiring `--mamba-cache-mode align` because LFM2's short-convolution layers need different cache alignment. A notable feature is `set_coding_enabled(False)`, which disables the transplanted expert layers at runtime, reverting to plain LFM2.5 behavior without reloading weights. `trust_remote_code=True` is always required due to the custom `Fuse3ForCausalLM` architecture.

## Key points

- fuse-1 Lite is a 5.72B MoE model: 960 coding experts from Qwen3.6-35B-A3B grafted onto a frozen LFM2.5-2.6B host, with only a lightweight router trained.
- VRAM ranges from 3.36 GB (4-bit NF4) to ~12 GB (bfloat16).
- Five deployment paths: transformers+bitsandbytes, GGUF, MLX, vLLM, and native transformers.
- GGUF and MLX need custom code (patched llama.cpp fork / custom MLX model file).
- 4-bit bitsandbytes is the cheapest entry point, running on a T4, RTX 3060, or M2 Pro.
- A runtime toggle (`set_coding_enabled`) disables coding experts to use it as a general chat model.
- Training cost was ~$3: 300 steps of router training on 55 examples.

## Technical data / figures

| Backend | Precision | VRAM/Memory | Hardware tier |
|---|---|---|---|
| Transformers | bfloat16 | ~12 GB | L4, A10G, RTX 4090 |
| Transformers | 8-bit | 6.00 GB | T4, L4, RTX 3060 |
| Transformers | 4-bit (NF4) | 3.36 GB | T4, RTX 3060, M2 Pro |
| vLLM | bfloat16 | ~12 GB | A10G, A100, H100 |
| MLX | float16 | ~12 GB | M1 Pro+, M2, M3, M4 |
| llama.cpp | F16 GGUF | ~11.4 GB | Any CPU/GPU |
| llama.cpp | Q4_K_M GGUF | ~4 GB | Any CPU/GPU |

- Host model: LFM2.5-2.6B (LiquidAI), 2.6B
- Expert source: Qwen3.6-35B-A3B (960 experts, 32/layer x 30 layers)
- Router: ~2M trainable params; 300 steps, ~8 min on L4, 55 examples, ~$3
- Active expert layers: 11 of 30 (disabled in 19), peak at layer 19

## Why this source matters for the RAG

It documents an unusual "expert transplant" architecture and gives precise, per-backend VRAM figures and setup commands for a small local coding model, useful for local deployment guides. It also illustrates custom-architecture handling (patched llama.cpp, vLLM plugins, `trust_remote_code`) that generalizes to other non-standard local models.
