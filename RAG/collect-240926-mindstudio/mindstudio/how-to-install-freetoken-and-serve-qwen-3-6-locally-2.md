---
id: collect-240926-mindstudio/mindstudio/how-to-install-freetoken-and-serve-qwen-3-6-locally-2
title: "how-to-install-freetoken-and-serve-qwen-3-6-locally"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "MiniMax", "Z.ai"]
dates: []
keywords: ["qwen", "agent", "agents", "benchmark", "claude", "compute", "consumer", "deepseek", "glm", "gpu", "gpus", "mixture of experts"]
source: docs/RAG/clean_en/mindstudio/how-to-install-freetoken-and-serve-qwen-3-6-locally.md
source_anchor: ""
source_lines: [86, 106]
sha256: b2ebcc5744c4c5fd713943e05236b706392514602dcb15a64441da9c2a3b41dd
---

# how-to-install-freetoken-and-serve-qwen-3-6-locally

No. The entire premise of FreeToken is enabling models that would normally need several high-end GPUs to run on a single consumer GPU, by keeping most of the model in system RAM and only streaming or computing the experts actually needed for each token.

### What models can I serve with FreeToken?

The demo used Qwen 3.6, but FreeToken is built around mixture of experts architectures broadly, which includes models like GLM, DeepSeek, and MiniMax. Swapping models is a matter of downloading a different checkpoint and pointing the serve command at it.

### What’s the difference between hybrid mode and offload mode?

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

Hybrid mode splits expert computation between CPU and GPU, chosen when CPU compute is fast enough relative to PCIe transfer speed. Offload mode streams needed experts over PCIe to the GPU instead, chosen when PCIe transfer outperforms CPU compute for that format. FreeToken’s benchmark step picks between the two automatically.

### Can I use FreeToken with coding agents like Claude Code?

Yes. FreeToken supports several coding agents, including Claude Code, Codex, Hermes agent, OpenClaw, and OpenCode. You launch the agent with a command pointing it at your local FreeToken server instead of a cloud API, and it installs automatically if not already present.
