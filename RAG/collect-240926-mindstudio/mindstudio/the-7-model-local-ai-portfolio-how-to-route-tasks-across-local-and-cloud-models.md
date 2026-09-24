---
id: collect-240926-mindstudio/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models
title: "the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "Microsoft", "Nvidia", "OpenAI", "vLLM"]
dates: ["2024-04"]
keywords: ["agent", "agentic", "agents", "apache", "benchmarks", "chatgpt", "claude", "compute", "cost", "deepseek", "embedding", "embeddings"]
source: docs/RAG/clean_en/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-.md
source_anchor: ""
source_lines: [1, 152]
sha256: 5c321299bbf25b46c526592fec4cc271d268f9620b49cf07c7db1cc33319eb93
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

For coding workflows, the routing looks like this: autocomplete and small edits go to the fast local model, repo-aware refactoring and test generation go to the coding generalist, architectural decisions and hard debugging go to the frontier fallback. The Continue VS Code extension handles the interface layer — it can point at any OpenAI-compatible endpoint, which means it works with your local Ollama server and your cloud API without changing your workflow.

For document work, the routing runs through the embedding model first. Retrieval happens locally, against your private vector store. Summarization and drafting go to the local generalist. Only the synthesis tasks that require genuine reasoning depth — the ones where a weaker model would produce a wrong answer, not just a worse one — go to the frontier.

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

For meeting capture, the entire pipeline stays local: Whisper for transcription, a local summarizer for extraction, your memory layer for storage. Jones describes this as one of the clearest wins in the stack. “Your decisions become searchable, your commitments become something you can retrieve and look at, your recurring conversations become part of a private institutional memory that you own.”

The memory layer underneath all of this is what makes routing durable over time. Jones built Open Brain — an open-source system combining SQL storage, embeddings, and an MCP server — specifically because the memory should belong to you, not the model provider. The inversion he describes: “In the cloud-first model, the AI service really wants to own your memory, and you visit your memory. In the personal compute model, you own the memory, and the models come to you.”

## What the Coding Sub-Stack Actually Looks Like

The coding slot deserves more detail because it’s where most builders will spend the most time, and where the routing decisions are most consequential.

The three-layer coding sub-stack: a small autocomplete model for fast loops (think sub-200ms response), a repo-aware model for refactoring and test generation (this is where Qwen and Gemma 4 compete), and a reasoning model for architectural decisions and hard debugging (GPT-OSS-120B or a frontier fallback).

Continue bridges the local runtime to your editor. Aider handles terminal-based editing for the cases where you want a more autonomous loop. The pattern Jones describes — “model plus tools plus repo plus context in a planning loop” — is the same whether you’re running local or cloud. The difference is that local inference absorbs the high-volume inner loop cheaply, and you reserve the expensive frontier calls for the decisions that actually require them.

If you’re building applications on top of this kind of multi-model stack, the orchestration layer matters as much as the models themselves. Platforms like MindStudio handle this orchestration at scale: 200+ models, 1,000+ integrations, and a visual builder for chaining agents and workflows — useful when you want to prototype routing logic without writing the plumbing from scratch.

On the code generation side, the abstraction question keeps moving. Tools like Remy take a different approach to the output layer: you write a spec — annotated markdown — and the full-stack app gets compiled from it. Backend, database, auth, deployment, all of it. The spec is the source of truth; the generated TypeScript is derived output. It’s a different answer to the question of how much of the stack you want to own directly versus derive from higher-level intent.

## The Hardware Constraint Nobody Mentions

The seven-model portfolio has a hardware dependency that’s easy to miss: you need enough unified memory to keep multiple models loaded simultaneously, or you pay a loading penalty every time you switch slots.

This is why the Mac mini M4 Pro with 64GB is the recommended entry point for local-first knowledge workers — not because it’s the fastest, but because unified memory means the GPU and CPU share the same pool, and 64GB is enough to keep your fast model, your generalist, and your embedding model warm at the same time. The Mac Studio M4 Max with 128GB gives you more headroom for larger models in the generalist and coding slots.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

The Nvidia DGX Spark’s 128GB of coherent unified memory matters for the same reason — it’s not just about total capacity, it’s about having a single addressable pool rather than the fragmented VRAM situation you get with dual RTX 5090s. Two RTX 5090s give you 64GB across cards, but that’s not one clean 64GB pool. Sharding across cards adds complexity and latency that works against the fast-routing architecture.

The buying rule Jones offers is worth repeating: “Don’t buy for the biggest model you read about. Buy the thing you’re going to run daily.”

## The Open-Weight Landscape Right Now

The model options in each slot have improved faster than most people predicted. Jones is direct about this: “Even a few months ago, open-source models couldn’t do a lot of what I just described at all.”

For the generalist slot, Llama 4 Scout and Maverick represent where the open ecosystem is headed — mixture-of-experts, multimodal, longer context. For the coding and reasoning slots, GPT-OSS-20B and GPT-OSS-120B are Apache 2.0 open-weight models you run on your own infrastructure. For multilingual work and tool use, Qwen has become a default. For smaller local deployments, Gemma 4 is specifically designed for open, local applications under a permissive license. If you want a detailed comparison of how Gemma 4 and Qwen perform against each other on local workflows, the Gemma 4 vs Qwen 3.5 open-weight comparison covers the benchmarks in depth.

The embedding slot is more stable — Qwen embedding models are the current recommendation, and the economics of embeddings (cheap, cacheable, rebuildable) mean you’re not under pressure to chase the latest release.

For teams thinking about cost reduction across the stack, the approach of routing cheaper tasks through local or free-tier models applies directly to this portfolio architecture — the local slots absorb the volume, the frontier slot handles the exceptions.

## Building the Portfolio vs. Buying a Model Appliance

The failure mode Jones keeps returning to is treating local AI as a single-model appliance. You buy a Mac mini, you install one model, you use it for everything, and then you’re disappointed when it can’t match GPT-4o on hard tasks and annoyed when it’s slow on simple ones.

The portfolio framing changes the question. You’re not asking “which model is best?” You’re asking “what is the mixture of models I need, and how do I route between them?”

That’s a more durable question. The specific models in each slot will change — probably every few months. The routing logic, the runtime layer, the memory architecture, the hardware — those are the investments that compound. A new model drops into the generalist slot without touching anything else. A better embedding model ships and you rebuild the index from your preserved raw data. A faster runtime replaces Ollama for a specific workload without breaking the interface layer.

If you’re evaluating how this kind of multi-model routing works in practice for agentic coding specifically, the comparison of Qwen 3.6 Plus and Claude Opus 4.6 on agentic coding tasks is a useful reference for understanding where the frontier models still pull ahead and where local models are closing the gap.

## One coffee. One working app.

You bring the idea. Remy manages the project.

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
