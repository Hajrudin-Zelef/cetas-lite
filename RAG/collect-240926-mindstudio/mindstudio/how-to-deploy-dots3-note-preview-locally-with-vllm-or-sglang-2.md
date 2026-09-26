---
id: collect-240926-mindstudio/mindstudio/how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang-2
title: "how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "agents", "fp8", "gpu", "memory", "nvidia", "speculative decoding", "tool calling"]
source: docs/RAG/clean_en/mindstudio/how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [99, 127]
sha256: 335353c1a424c9e4ab5c85b23dd3176110359fef7a4043223d0fc8bd44c2c91a
---

# how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang

Transformers support (tracked in Hugging Face pull request #47844) is good for quick local testing on a single GPU or machine, not for production multi-GPU serving. The setup requires installing compatible PyTorch and torchvision builds for your NVIDIA driver, plus `torchcodec` and FFmpeg if you want audio or video input support, then installing Transformers directly from the PR branch since it hasn’t merged yet.

A minimal example loads the FP8 checkpoint with `AutoModelForMultimodalLM` and runs `generate()` directly, which is useful for verifying the model loads and produces sensible output before committing to a full SGLang or vLLM deployment. For anything resembling real traffic or concurrent requests, though, the model card explicitly points to SGLang or vLLM for OpenAI-compatible serving.

## Frequently Asked Questions

### What GPU count does dots3-note preview need?

The documented configurations target a single 8-GPU node, specifically NVIDIA H100s in vLLM’s reference example, running the FP8 checkpoint with tensor parallelism and expert parallelism both set to 8.

### Does dots3-note preview support tool calling?

Yes. Both SGLang and vLLM support OpenAI-compatible tool/function calling through a dedicated `dots` tool-call parser flag (`--tool-call-parser dots` in SGLang, and `--enable-auto-tool-choice --tool-call-parser dots` in vLLM).

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

### What’s the difference between the BF16 and FP8 checkpoints?

The FP8 checkpoint (`dots-studio/dots3-note-prev-fp8`) is quantized and needs less GPU memory to serve, which is why it’s the recommended default for one-node deployment. The BF16 checkpoint offers full precision but needs considerably more memory for the same 280B-parameter architecture.

### Can I run dots3-note preview without a GPU cluster?

The officially documented deployment paths (SGLang and vLLM) assume an 8-GPU node. Transformers can load the model for basic single-machine testing, but it’s not positioned as a substitute for multi-GPU serving in the model’s own deployment guidance.

### What is MTP/NEXTN speculative decoding used for here?

It’s an optional feature that uses the model’s built-in Multi-Token Prediction layer to draft multiple tokens ahead and verify them together, which the documentation states can reduce time-per-output-token by more than 50% when enabled on either SGLang or vLLM.
