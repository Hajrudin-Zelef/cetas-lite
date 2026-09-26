---
id: collect-240926-mindstudio/mindstudio/run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm-2
title: "run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "vLLM"]
dates: []
keywords: ["qwen", "vllm", "agent", "agents", "consumer", "context window", "gpu", "gpus", "int4", "kv cache", "memory", "quantization"]
source: docs/RAG/clean_en/mindstudio/run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm.md
source_anchor: ""
source_lines: [75, 93]
sha256: 24d9a1de49d5eda48c26aecd926bd8748f20d9d3d3bd9e15512be00bb9fbd709
---

# run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm

### What quantization format does Qwen 3.8 Flash Next use for local deployment?

The commonly used local quant is INT4, specifically a W4A16 format produced by Vinnie AI, which keeps vision input support intact while reducing VRAM requirements enough to run on consumer GPUs.

### Can this run on fewer than four GPUs?

The documented setup assumes four RTX 3090s for a combined 96GB of VRAM to support the full context window and KV cache. Running on fewer GPUs would require reducing context length or other memory-hungry settings, and hasn’t been verified in this configuration.

### Why is async scheduling disabled in the vLLM config?

Async scheduling can increase throughput by around 10 tokens per second, but without an additional compatibility patch it introduces instability that eventually causes the server to crash, so it’s left off for a stable baseline.

### Does this setup work with other GPU architectures besides the 3090?

The build described here targets SM86, the Ampere architecture used by the RTX 3090. It may be adaptable to SM89 architectures with some modification, but that path isn’t the tested default.

### Why does vision support matter for a local agent model?

Vision input lets an agent process screenshots, UI renders, or diagrams as part of its reasoning, which matters for coding and content-generation agents that need to check their own visual output rather than working from text descriptions alone.
