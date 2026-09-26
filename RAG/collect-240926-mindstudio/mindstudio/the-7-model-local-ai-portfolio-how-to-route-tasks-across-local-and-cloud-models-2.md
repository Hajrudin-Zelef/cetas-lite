---
id: collect-240926-mindstudio/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-2
title: "the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Microsoft", "Nvidia", "OpenAI"]
dates: []
keywords: ["agentic", "agents", "apache", "benchmarks", "claude", "compute", "cost", "embedding", "embeddings", "gpu", "inference", "latency"]
source: docs/RAG/clean_en/mindstudio/the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-.md
source_anchor: ""
source_lines: [66, 127]
sha256: eff562ce08b126afe09e1cfec030e187c9fec6570aa52dbb69cc18bad559e34d
---

# the-7-model-local-ai-portfolio-how-to-route-tasks-across-local-and-cloud-models-

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

