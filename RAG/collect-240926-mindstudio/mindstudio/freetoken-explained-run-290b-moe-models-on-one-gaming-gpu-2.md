---
id: collect-240926-mindstudio/mindstudio/freetoken-explained-run-290b-moe-models-on-one-gaming-gpu-2
title: "freetoken-explained-run-290b-moe-models-on-one-gaming-gpu"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["gpu", "moe", "agent", "agents", "attention", "benchmark", "claude", "compute", "consumer", "gpus", "kv cache", "quantization"]
source: docs/RAG/clean_en/mindstudio/freetoken-explained-run-290b-moe-models-on-one-gaming-gpu.md
source_anchor: ""
source_lines: [69, 83]
sha256: a03f76fce639537e97066b372220252171c1da249f3c428da9d1ee8bebc58b48
---

# freetoken-explained-run-290b-moe-models-on-one-gaming-gpu

### Do I need a multi-GPU setup to use FreeToken?

No. The entire point of FreeToken is to avoid needing multiple GPUs. A single consumer or workstation GPU, combined with enough system RAM to hold the rest of the model, is sufficient. VRAM usage scales with how many experts are cached and the size of the KV cache, not the full model size.

### How does FreeToken decide whether to use CPU or GPU for an expert?

It runs a one-time benchmark comparing CPU compute speed against PCIe data transfer speed for the model’s specific quantization format. If computing an expert locally on the CPU is faster than streaming it to the GPU, it uses hybrid mode. If streaming is faster, it uses offload mode. This is chosen automatically before serving starts.

### Does using FreeToken require manual configuration?

Minimal configuration is needed. The server automatically selects the attention backend, MoE backend, and cache allocation based on the model checkpoint and detected GPU specs. Users mainly need to download a model, run the benchmark, and start the server.

### Can FreeToken be used with tools I already have, like Claude Code?

Yes. FreeToken can act as a local model server that existing coding agents point to instead of a cloud API. Agents such as Claude Code, Codex, and others can be configured to send requests to the local FreeToken server, letting you use a locally hosted model within a familiar agent interface.
