---
id: collect-mindstudio/mindstudio/qwen-3-8-flash-next-local-agent-setup
title: "Run Qwen 3.8 Flash Next Locally on Quad RTX 3090s with vLLM"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "vllm", "agent", "agentic", "agents", "attention", "consumer", "context window", "gpu", "gpus", "int4", "kv cache"]
source: docs/RAG/Collect RAG/02_mindstudio/qwen-3-8-flash-next-local-agent-setup.md
source_anchor: ""
source_lines: [1, 61]
sha256: 91a2661eec5b7a3d8ed1b2405a7efc7324166b7f0180184b8b04a7a559e90ccc
---

# Run Qwen 3.8 Flash Next Locally on Quad RTX 3090s with vLLM

## Metadata

- **Source** : https://www.mindstudio.ai/blog/qwen-3-8-flash-next-local-agent-setup
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article is a practical guide to running Qwen 3.8 Flash Next locally with vLLM on four RTX 3090s. Qwen 3.8 Flash Next is a vision-capable model from Alibaba's Qwen team, listed on Hugging Face under `Qwen/Qwen3.8-Flash-Next` with an image-text-to-text pipeline and the `qwen4_exp` architecture tag. The published safetensors weights ship across 131 files (a large full-precision footprint), but the model becomes practical on consumer hardware once quantized. The version used for local agentic work is an INT4 (W4A16) quant produced by Vinnie AI that keeps vision support intact while shrinking VRAM requirements enough to fit a multi-GPU 3090 rig. Strong output quality at Q4 plus vision is why it draws attention as a faster alternative to smaller 27B-class local models.

The article explains the motivation: 3090 users have long chosen between speed and capability. Smaller dense 27B-class models run fast but plateau on quality; larger MoE-style models are strong but often too slow or VRAM-hungry for agentic loops (calling tools, waiting, re-reading context, generating repeatedly). Qwen 3.8 Flash Next threads the needle — larger than previous go-to options, but its INT4 quant doesn't degrade output quality much and it retains vision input, which matters for agents that inspect screenshots, diagrams, or rendered UI.

Hardware requirements are specific: four RTX 3090s (96GB combined VRAM) to load the model with a large KV cache and hit the full context window; 128GB system RAM as the recommended floor (system memory climbs to around 100GB during load); CUDA and NVCC 13.0+ (13.3+ preferred) with current NVIDIA drivers; and a build compiled for SM86 (Ampere). Newer SM89 cards could likely work with changes but that isn't the tested path. It doesn't require enterprise GPUs or NVLink, just four consumer 3090s.

The vLLM launch flags matter more than usual. Core configuration: `--max-model-len 131072` (max context 131,072 tokens), `--max-num-seqs 2` (concurrent sequences, tuned for single-agent workloads), `--max-num-batched-tokens 2048`, `--gpu-memory-utilization 0.96` (aggressive, maximizes context/cache), `--kv-cache-dtype auto`, `--cuda-graph-mode full`, `--enable-prefix-caching` (reuses computation across shared prompt prefixes, helps agentic loops), and `--enable-auto-tool-choice` with `--tool-call-parser qwen3xml` (required for reliable function calling with a Hermes-style agent). Async scheduling is deliberately left off: enabling it can add close to 10 tokens/second but introduces instability that eventually crashes the server without an additional patch. Two other settings: a compression threshold starting around 0.5 with headroom toward 0.65, and a context length around 131,072 with a KV cache large enough for the full window.

On performance: prompt processing was observed climbing from roughly 399 to over 500 tokens as context builds, with generation around 59 tokens/second in agentic use, dropping slightly with async scheduling disabled but staying stable — roughly double the throughput of llama.cpp on identical hardware. The article notes real-world throughput lands around 55–61 tokens/second at large context depths. For agentic work the speed difference compounds across dozens of tool calls per session. The article concludes the setup is worth it for anyone already running multiple 3090s: a quad-3090 rig can run a vision-capable model with full 131K context at speeds suitable for iterative agent loops. Caveats: some performance gains involve moving patches/configs, and faster KV cache quantization approaches (one owner reported near 200 tokens/second) add setup complexity and instability, so the described config is a stable baseline rather than the theoretical ceiling.

## Key points

- Qwen 3.8 Flash Next is a vision-capable model (`Qwen/Qwen3.8-Flash-Next`, `qwen4_exp`); local deployment uses a Vinnie AI INT4 (W4A16) quant retaining vision.
- Requires four RTX 3090s (96GB VRAM) and 128GB system RAM (system usage ~100GB at load).
- Full 131,072-token context window achievable with vLLM.
- Key flags: max-model-len 131072, max-num-seqs 2, max-num-batched-tokens 2048, gpu-memory-utilization 0.96, cuda-graph-mode full, enable-prefix-caching.
- Tool calling needs enable-auto-tool-choice with the qwen3xml parser.
- Async scheduling left off (adds ~10 tok/s but crashes without a patch).
- Throughput: ~55–61 tokens/second generation at large context; prompt processing ~399→500+; ~2x llama.cpp.
- Build targets SM86 (Ampere); SM89 may work with changes but untested.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Qwen 3.8 Flash Next |
| HF repo | Qwen/Qwen3.8-Flash-Next |
| Architecture tag | qwen4_exp |
| Weights files | 131 safetensors |
| Quant | INT4 (W4A16) by Vinnie AI |
| GPUs | 4x RTX 3090 (96GB VRAM) |
| System RAM | 128GB minimum |
| CUDA/NVCC | 13.0+ (13.3+ preferred) |
| Arch target | SM86 (Ampere) |
| Max context | 131,072 tokens |
| max-num-seqs | 2 |
| max-num-batched-tokens | 2048 |
| gpu-memory-utilization | 0.96 |
| Compression threshold | ~0.5 (up to ~0.65) |
| Tool-call parser | qwen3xml |
| Generation speed | ~55–61 tokens/sec |
| Prompt processing | ~399→500+ tokens |
| Async scheduling | Disabled (stability) |

## Why this source matters for the RAG

It is a concrete, reproducible local-deployment recipe with exact vLLM flags, hardware requirements, and measured throughput for a vision-capable agentic model. It is directly actionable for local-AI hardware and serving decisions, and its performance data enables comparison with cloud and other local options.

