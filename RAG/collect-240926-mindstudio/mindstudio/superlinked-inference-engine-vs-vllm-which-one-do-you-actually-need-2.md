---
id: collect-240926-mindstudio/mindstudio/superlinked-inference-engine-vs-vllm-which-one-do-you-actually-need-2
title: "superlinked-inference-engine-vs-vllm-which-one-do-you-actually-need"
domain: mindstudio
role: reference
task: reference
actors: ["vLLM"]
dates: []
keywords: ["vllm", "agent", "apache", "embedding", "gpu"]
source: docs/RAG/clean_en/mindstudio/superlinked-inference-engine-vs-vllm-which-one-do-you-actually-need.md
source_anchor: ""
source_lines: [72, 80]
sha256: aadaa5ccff31014e0a19d5be992fea3e6533b851e98ef7481a0e1c7f161894e4
---

# superlinked-inference-engine-vs-vllm-which-one-do-you-actually-need

The core engine is Apache 2.0 licensed and free to run on your own hardware, from a laptop up to a Kubernetes cluster. A fully hosted option, SIE Cloud, is also available for teams that don’t want to manage their own GPU infrastructure.

### Can vLLM and SIE work together in the same application?

Yes. A common pattern is running SIE for the specialist layer (embedding, re-ranking, extraction) and vLLM for the large generative model, with the two communicating over a simple HTTP call, keeping the application architecture portable.

### Which one should I start with if I’m building an AI agent?

If your agent only needs to generate text from one large model, vLLM is the simpler starting point. If your agent needs multiple capabilities, such as embedding, re-ranking, or entity extraction, alongside generation, SIE avoids the need to stand up a separate server for each one.
