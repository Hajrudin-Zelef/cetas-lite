---
id: collect-250926-servers-hardware/servers-hardware/context-length-ollama
title: "context-length-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/clean4/context-length-ollama.md
source_anchor: ""
source_lines: [1, 18]
sha256: cb5b11a487f29fe5d2d1e2f95e4dac1d4c43a80f21fd027bbb791f591877eec6
---

# context-length-ollama

Ollama defaults to the following context lengths based on VRAM:

- < 24 GiB VRAM: 4k context
- 24-48 GiB VRAM: 32k context
- >= 48 GiB VRAM: 256k context

## Setting context length

Setting a larger context length will increase the amount of memory required to run a model. Ensure you have enough VRAM available to increase the context length. Cloud models are set to their maximum context length by default.
### App

Change the slider in the Ollama app under settings to your desired context length.
### CLI

If editing the context length for Ollama is not possible, the context length can also be updated when serving Ollama.
### Check allocated context length and model offloading

For best performance, use the maximum context length for a model, and avoid offloading the model to CPU. Verify the split under`PROCESSOR` using `ollama ps`.
