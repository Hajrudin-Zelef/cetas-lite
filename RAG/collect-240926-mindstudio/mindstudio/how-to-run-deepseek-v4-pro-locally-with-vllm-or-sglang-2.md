---
id: collect-240926-mindstudio/mindstudio/how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang-2
title: "how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["deepseek", "sglang", "vllm", "agent", "agents", "gpu", "gpus", "inference", "license", "mit license", "open source", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [78, 104]
sha256: daf9f6c2352ed0eaef6f454a917ef7ef8df5aeb8786125b845a346f4ac771352
---

# how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang

Because the harness treats models as pluggable providers under an MIT license, a locally-hosted V4 Pro instance running under vLLM or SGLang can be registered as a custom provider instead of pointing at DeepSeek’s hosted API. This is useful for anyone who wants the harness’s agent tooling, permission modes (read-only, workspace, full access), and plugin ecosystem while keeping inference entirely on their own infrastructure.

## Frequently Asked Questions

### What hardware do you need to run DeepSeek V4 Pro locally?

DeepSeek’s official recipes reference multi-GPU nodes, with the primary documented example being a single node with four GB300 GPUs for both vLLM and SGLang. Both recipes link out to additional hardware configurations for different cluster sizes.

### Does DeepSeek V4 Pro require a separate draft model for speculative decoding?

No. DSpark speculative decoding draws both draft and target weights from the same released checkpoint, so there’s no additional draft model to download or maintain.

### What’s the difference between reasoning effort levels in V4 Pro?

The model supports low, high, and max reasoning effort settings, controlling how much deliberation it performs before answering. For high and max effort, DeepSeek recommends allowing output lengths up to 384K tokens locally.

### Can I use DeepSeek V4 Pro with agent frameworks besides DeepSeek’s own harness?

Yes. Since it’s served through standard OpenAI-compatible endpoints via vLLM or SGLang, it can be wired into most agent frameworks that support custom model providers, not just DeepSeek’s own harness.

### Is DeepSeek V4 Pro fully open source?

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

The repository and model weights are released under the MIT license, which permits commercial use, modification, and self-hosting without licensing fees.
