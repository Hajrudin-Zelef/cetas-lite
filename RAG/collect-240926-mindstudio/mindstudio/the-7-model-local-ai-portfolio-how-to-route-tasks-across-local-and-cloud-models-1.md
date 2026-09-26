---
id: collect-240926-mindstudio/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-1
title: "the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "Microsoft", "Nvidia", "OpenAI", "vLLM"]
dates: ["2024-04"]
keywords: ["agent", "agentic", "agents", "apache", "chatgpt", "claude", "compute", "cost", "deepseek", "embedding", "embeddings", "gguf"]
source: docs/RAG/clean_en/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-.md
source_anchor: ""
source_lines: [1, 65]
sha256: 5f1b96c94a7b111d27d0e51b739d712a497e41e05019df65146b59fea043b990
---

# the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-

<!-- source: https://www.mindstudio.ai/blog/7-model-local-ai-portfolio-routing-local-cloud -->

## Seven Models, One Stack: The Local AI Portfolio That Actually Routes Work Correctly

Most people running local AI are doing it wrong. They install Ollama, pull a single model, and then wonder why it feels slow on simple tasks and inadequate on hard ones. The answer isn’t better hardware — it’s a better portfolio. Specifically, a seven-model portfolio: a fast local model, a strong local generalist, a coding model, an embedding model, a speech model, a vision model, and a frontier cloud fallback. That’s the architecture. The rest is routing.

This framing comes from Nate Jones, who published a detailed breakdown of the full personal AI stack after testing the Mac mini M4 Pro, Mac Studio M4 Max, RTX 5090, and Nvidia DGX Spark against real workloads. His central argument isn’t about hardware. It’s about how you think about model selection. “I would not build a personal AI computer around a single model name,” he says. “I would build around model classes for particular workloads.”

That’s the shift. You’re not picking a favorite chatbot. You’re building a tool cabinet.

## The Problem With Running One Model for Everything

A single model creates two failure modes simultaneously. You’re overpaying — in latency and compute — for cheap tasks like classification, summarization, and autocomplete. And you’re underpowered for the tasks that actually need reasoning depth: architectural decisions, hard debugging, synthesis across long documents.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

The economics are obvious once you see them. A fast, small local model can handle the inner loop of an agentic workflow — the repeated, cheap calls — at essentially zero marginal cost. A frontier cloud model like GPT-4o or Claude Opus should be reserved for the rare, hard, high-value work. Everything in between gets routed to the right specialist.

Jones puts it plainly: “The principle should be you own the runtime, and you only rent the cloud model in exceptional cases.”

The routing question, then, is the real engineering problem. And it starts with knowing what each slot in the portfolio is actually for.

## The Seven Slots, Defined

**Slot 1: Fast local model for cheap calls.** This is your workhorse for high-frequency, low-complexity tasks — classification, short summarization, intent detection, routing decisions themselves. You want something that loads fast and responds in under a second on your hardware. On Apple Silicon, a 7B or 8B quantized model running through Ollama fits this slot well. The goal is throughput, not depth.

**Slot 2: Strong local generalist.** This is the model you reach for when the task is real but not exceptional — drafting, research synthesis, document Q&A, longer reasoning chains. Llama 4 Scout and Llama 4 Maverick are the current reference points here. They’re mixture-of-experts models, which means the active parameter count per token is lower than the total model size suggests. Longer context, multimodal support, more deployment nuance. Qwen models have also become a default for this slot, particularly for agents, coding, and multilingual work.

**Slot 3: Coding model.** This is not one model — it’s a sub-portfolio. Jones is specific about this: “You don’t want one model doing everything. You want a small autocomplete model, a repo-aware editor model, and a deeper reasoning model for architectural changes, for debugging, for migrations.” The tooling here is Continue for VS Code (which points at any OpenAI-compatible local endpoint) and Aider for terminal-based editing. GPT-OSS-20B and GPT-OSS-120B — OpenAI’s Apache 2.0 open-weight reasoning models — are worth evaluating for the deeper reasoning slot in this sub-stack.

**Slot 4: Embedding model.** This is the most underrated slot. Embeddings are cheap to run, easy to cache, and central to any private retrieval system. Qwen embedding models are the current recommendation for local RAG. The key architectural point: your raw documents and your embeddings should live separately in your database. When a better embedding model ships — and they keep shipping — you can rebuild the index without losing your source data. Postgres with pgvector is the production default for this layer; SQLite with sqlite-vec works for personal-scale setups.

**Slot 5: Speech model.** Whisper. Local transcription is fast, private, and — if you own the hardware — essentially free at scale. Jones describes the workflow: “Local Whisper plus a local summarizer means you can record and transcribe and summarize and extract decisions and create tasks and store that result in your memory layer. No audio ever leaves the machine, no per-hour transcription bill.” Run it on every call for a year and your decisions become searchable.

**Slot 6: Vision model.** Local vision models are now good enough for document screenshots, chart extraction, and personal media search. Not for all visual reasoning — but for a meaningful slice of daily work. This slot was largely empty a year ago. It isn’t now.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

**Slot 7: Frontier cloud fallback.** This is not a failure mode — it’s a design choice. Hard synthesis, frontier coding problems, tasks that genuinely require the best available model. The personal AI computer isn’t anti-cloud. It’s anti-dependence. You keep one frontier subscription or API account for the work that deserves it. The rest stays local.

## Why the Model Layer Ages Out Faster Than Everything Else

Here’s the non-obvious part of this architecture: the model list is the least durable component of the stack.

Jones makes this point explicitly. In April 2024, DeepSeek previewed V4 with Pro and Flash variants. Gemma 4 pushed serious capability into smaller local models under a more permissive license. GPT-OSS-20B and GPT-OSS-120B appeared as Apache 2.0 open-weight reasoning models — not the ChatGPT API, but weights you run on infrastructure you control. “Any model list you make today starts aging right away,” Jones says. “That’s the point.”

This is why the portfolio framing matters more than the specific model names. If you build around model classes — fast, generalist, coding, embedding, speech, vision, cloud — you can swap the specific model in each slot as better options arrive. The routing logic stays stable. The runtime stays stable. The memory layer stays stable. Only the weights change.

The runtime layer is what makes this possible. llama.cpp is the foundation — it created the GGUF format, runs across CPU, Apple Metal, CUDA, and Vulkan, and underpins most of what you’ll use. Ollama sits on top for daily use: clean CLI, local server, simple model registry, OpenAI-compatible surface. LM Studio handles model evaluation. MLX is the Apple-native performance path for Apple Silicon. vLLM enters the picture when you’re serving real workloads on Nvidia hardware and need batching and throughput for a team or internal product.

The practical default: Ollama for daily use, LM Studio for evaluation, MLX on Mac, vLLM when serving becomes infrastructure. Notice that none of this is about which model is best this week.

## The Routing Logic Nobody Talks About

Building the portfolio is the easy part. Routing between slots correctly is where most setups fall apart.

The routing question has two dimensions: task complexity and data sensitivity. Simple, repeated, private tasks go local. Rare, hard, high-value tasks go to the frontier. The interesting cases are in the middle.

