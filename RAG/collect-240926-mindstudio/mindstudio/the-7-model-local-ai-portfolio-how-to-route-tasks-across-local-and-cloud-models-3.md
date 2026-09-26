---
id: collect-240926-mindstudio/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-3
title: "the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-"
domain: mindstudio
role: reference
task: reference
actors: ["Apple", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["agentic", "compute", "cost", "embedding", "gguf", "gpu", "latency", "llama", "llama.cpp", "memory", "nvidia", "pricing"]
source: docs/RAG/clean_en/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-.md
source_anchor: ""
source_lines: [128, 152]
sha256: 1b71530c1210ff9f694b7c4c32f370d475926adc311b2570aec452d98f4921f4
---

# the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-

The personal AI computer, as Jones frames it, is not a nostalgia play. It’s a routing system. Some work stays local because it’s private, cheap, repeated, or context-heavy. Some work goes to the cloud because it’s rare, hard, or genuinely requires the frontier. The power is in deciding — not defaulting to whatever the cloud provider’s pricing model incentivizes.

Seven models. One stack. You decide what goes where.

## Frequently Asked Questions

### What are the seven slots in the portfolio?

The seven slots are a fast local model for cheap calls, a strong local generalist, a coding model, an embedding model, a speech model, a vision model, and a frontier cloud fallback. The idea is to build around model classes rather than specific model names, so you can swap the weights in each slot as better options arrive. The routing logic, runtime, and memory layer stay stable while only the models change.

### Why is running a single model for everything a problem?

A single model creates two failure modes at once: you overpay in latency and compute for cheap tasks like classification, summarization, and autocomplete, and you’re underpowered for tasks that need reasoning depth, such as architectural decisions, hard debugging, and long-document synthesis. A fast small local model can handle the repeated inner-loop calls of an agentic workflow at essentially zero marginal cost. Frontier cloud models should be reserved for rare, hard, high-value work.

### How much memory do I need to run multiple models at once?

The portfolio depends on having enough unified memory to keep multiple models loaded simultaneously, otherwise you pay a loading penalty each time you switch slots. A Mac mini M4 Pro with 64GB is the recommended entry point because unified memory lets the CPU and GPU share one pool, enough to keep your fast model, generalist, and embedding model warm. The Mac Studio M4 Max with 128GB gives more headroom, and the Nvidia DGX Spark’s 128GB coherent unified memory matters for the same reason — a single addressable pool rather than fragmented VRAM across dual RTX 5090s.

### Which runtimes should I use to run these models?

llama.cpp is the foundation — it created the GGUF format and runs across CPU, Apple Metal, CUDA, and Vulkan. Ollama sits on top for daily use with a clean CLI, local server, model registry, and OpenAI-compatible surface, while LM Studio handles evaluation and MLX is the Apple-native performance path. vLLM comes in when you’re serving real workloads on Nvidia hardware and need batching and throughput.

### How should coding tasks be split across models?

The coding slot is really a sub-portfolio of three layers: a small autocomplete model for fast loops, a repo-aware model for refactoring and test generation, and a reasoning model for architectural decisions and hard debugging. Autocomplete and small edits go to the fast local model, refactoring and test generation to the coding generalist, and architectural or hard debugging work to GPT-OSS-120B or a frontier fallback. Continue bridges the local runtime to VS Code by pointing at any OpenAI-compatible endpoint, and Aider handles terminal-based editing.
