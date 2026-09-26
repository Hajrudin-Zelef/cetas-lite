---
id: collect-240926-mindstudio/mindstudio/how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang-2
title: "how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "Hugging Face", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "agents", "attention", "deepseek", "fine-tuning", "fp8", "gpu", "gpus", "inference", "latency", "memory"]
source: docs/RAG/clean_en/mindstudio/how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [102, 128]
sha256: e2d277aad57fe67744ebdcbbdde7ca4983e402f0efcfa3024e49e72f7efee84a
---

# how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang

But it’s explicitly a preview, and Tencent is upfront about the rough edges: the model tends to spend more time reasoning than necessary on straightforward tasks and over-verifies its own work, both signs of an undertuned reasoning policy. Tencent frames this the same way it framed the earlier Hy3 preview: ship early, gather feedback, iterate. If you need a stable, fully polished model for production today, that caveat matters. If you’re comfortable working with a frontier-class but still-rough open-weight release, and you have the 8-GPU infrastructure to run it, Hy4 preview is a legitimate option to evaluate against other current MoE flagships in its weight class.

## Frequently Asked Questions

### How many GPUs do I need to run Hy4 preview?

Tencent’s official deployment recipes for both vLLM and SGLang default to `--tensor-parallel-size 8`, meaning an 8-GPU setup is the documented baseline for serving this 770B-parameter MoE model, even using the FP8 quantized checkpoint.

### What’s the difference between Hy4-preview and Hy4-preview-FP8?

Hy4-preview is the full-precision release, while Hy4-preview-FP8 is an FP8 quantized version of the same model, intended to reduce memory footprint and speed up inference for production serving. Both official Docker deployment examples use the FP8 variant.

### Does Hy4 preview support speculative decoding?

Yes. The model includes a built-in MTP (multi-token prediction) layer with about 10B total parameters (0.7B activated) specifically for speculative decoding, which both the vLLM and SGLang deployment configs enable by default to reduce generation latency.

### Can I fine-tune Hy4 preview myself?

Tencent provides a complete finetuning pipeline documented in the model’s finetuning guide on Hugging Face, alongside AngelSlim, a separate toolkit for quantization and model compression if you need to shrink the model further after fine-tuning.

### What context length does Hy4 preview support?

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

The model card lists a context length of 1 million tokens, backed by an architecture using Gated DeepSeek Sparse Attention (Gated DSA) with IndexCache for efficient cross-layer index reuse at long context lengths.
