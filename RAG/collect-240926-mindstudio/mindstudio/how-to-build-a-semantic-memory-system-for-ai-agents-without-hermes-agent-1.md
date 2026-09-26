---
id: collect-240926-mindstudio/mindstudio/how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent-1
title: "how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "OpenAI"]
dates: []
keywords: ["agent", "agents", "memory", "attention", "claude", "context window", "cost", "embedding", "embeddings", "latency", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent.md
source_anchor: ""
source_lines: [1, 156]
sha256: 2367e1702dfe0ad863c9c431f2c414217adc17cb3143db4f16d9659be846493d
---

# how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent

<!-- source: https://www.mindstudio.ai/blog/build-semantic-memory-system-ai-agents-without-hermes -->

## Why AI Agents Keep Forgetting (And How to Fix It)

Memory is the difference between an AI agent that’s genuinely useful and one you have to re-explain everything to every single session. If you’ve spent any time building agents, you’ve hit this wall. The model reasons well. The tool calls work. But the moment context resets, it’s back to square one.

This is exactly what made the Hermes Agent experiment interesting. It showed — clearly, repeatedly — that when you give an AI agent a structured semantic memory system, its performance on complex, multi-step tasks improves substantially. Agents stopped repeating themselves. They built on prior reasoning. They started feeling less like stateless APIs and more like persistent collaborators.

The good news: you don’t need Hermes Agent to get that. You can build a **semantic memory system for AI agents** from scratch using local vector databases, embedding pipelines, and smart injection logic. This guide walks through exactly how to do that inside environments like Claude Code or any custom agent framework.

## What Hermes Agent Actually Demonstrated

Hermes Agent — built on top of Ollama and local models — gained attention for its emphasis on persistent, context-aware memory. The core insight wasn’t revolutionary in theory, but it was compelling in practice: agents that could recall relevant past interactions performed better on reasoning tasks than those operating from scratch.

What it specifically demonstrated:

- **Retrieval-augmented context** matters more than raw model size in long-horizon tasks
- Storing memories as vector embeddings, rather than raw text logs, allows for *semantic* retrieval — not just keyword matching
- Agents with memory made fewer contradictory decisions across sessions
- The retrieval step itself could be lightweight enough to run locally without meaningful latency

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

The limitation of Hermes Agent was mostly about portability and flexibility. It worked within its own ecosystem. If you’re building on Claude Code, LangChain, CrewAI, or a custom Python agent, you need the same capability without the lock-in.

## The Anatomy of a Semantic Memory System

Before writing any code, it’s worth understanding the three functional layers of a semantic memory system.

### Layer 1: Storage

Where memories live. In a semantic system, this is a vector database — a store that holds numerical representations of text (embeddings) alongside the original content and metadata.

### Layer 2: The Embedding Pipeline

The process of converting text (agent observations, user inputs, decisions, outcomes) into vector embeddings. This is what enables semantic search rather than keyword matching. Two sentences that mean the same thing but use different words will have similar vectors.

### Layer 3: Retrieval and Injection

When the agent needs to act, it queries the memory store using the current context as a search key. The most semantically relevant memories are retrieved and injected into the agent’s prompt or context window.

These three layers are independent. You can swap out the vector DB, change embedding models, or redesign the injection logic without touching the others. That modularity is intentional — your memory system should be infrastructure, not a bottleneck.

## Choosing a Local Vector Database

Running memory locally matters for several reasons: latency, privacy, cost, and control. Here are the main options worth considering.

### ChromaDB

ChromaDB is probably the most accessible starting point. It runs in-process (no separate server needed for small deployments), has a clean Python API, and integrates natively with LangChain and LlamaIndex.

**Good for:** prototyping, local development, small to medium agent deployments

**Limitations:** not ideal for high-concurrency production workloads

```
import chromadb
client = chromadb.Client()
collection = client.create_collection("agent_memory")
collection.add(
    documents=["User prefers concise responses", "Project deadline is Friday"],
    ids=["mem_001", "mem_002"]
)
results = collection.query(
    query_texts=["What are the user's preferences?"],
    n_results=3
)
```
### Qdrant

Qdrant is a purpose-built vector search engine with a Rust core, which makes it fast. It supports filtering, named vectors, and payload indexing. You can run it locally via Docker with a single command.

**Good for:** production-grade local deployments, agents that need filtered retrieval

**Limitations:** requires Docker; slightly more setup than Chroma

### FAISS

Facebook AI Similarity Search (FAISS) is a library rather than a database. It’s extremely fast for pure similarity search but doesn’t handle persistence, metadata, or CRUD operations natively. You’d wrap it yourself or use it through a higher-level library.

**Good for:** high-performance similarity search when you’re managing storage separately

**Limitations:** no built-in persistence; you handle serialization

### LanceDB

A newer entrant, LanceDB stores vectors in a columnar format on disk. It’s serverless by design and has good performance for agents that need to scan large memory stores.

**For most builders starting out, ChromaDB gets you running in under 10 minutes. Move to Qdrant when you need more control.**

## Building Your Embedding Pipeline

The embedding pipeline converts raw text into vectors. The model you use determines how well your semantic search performs.

### Choosing an Embedding Model

For local use, these are the practical options:

- **`nomic-embed-text`** via Ollama — solid general-purpose embeddings, runs locally, 768 dimensions
- **`mxbai-embed-large`** via Ollama — higher quality, 1024 dimensions, slightly slower
- **`text-embedding-3-small`** via OpenAI API — excellent quality if you’re okay with API calls, cost is minimal
- **`sentence-transformers/all-MiniLM-L6-v2`** via HuggingFace — fast, lightweight, works entirely offline

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

For a Claude Code environment where you want full local operation:

```
import ollama
def embed_text(text: str) -> list[float]:
    response = ollama.embeddings(
        model="nomic-embed-text",
        prompt=text
    )
    return response["embedding"]
```
### What to Embed

This is a design decision that significantly affects memory quality. Don’t just dump raw conversation logs.

**Embed these:**

- Summarized observations (“User confirmed the report format should be CSV, not JSON”)
- Agent decisions with context (“Chose to use the search tool because the query required current data”)
- Task outcomes (“Completed email draft — user approved with minor edits”)
- User preferences revealed during interaction
- Errors and how they were resolved

**Avoid embedding:**

- Raw, unprocessed transcripts (too noisy, retrieval quality degrades)
- Ephemeral tool outputs that aren’t meaningful long-term
- Duplicate or near-duplicate entries (use deduplication logic)

### Memory Extraction Step

Before storing, have the agent (or a lightweight summarizer) extract the memory-worthy content from a conversation turn:

```
def extract_memory(conversation_turn: str, llm) -> str:
    prompt = f"""Extract the key factual information, preferences, or decisions from this interaction that would be useful to remember in future sessions. Be concise — one to three sentences maximum.
Interaction: {conversation_turn}
Memory:"""
    return llm.complete(prompt)
```
This step dramatically improves retrieval quality because you’re storing clean, dense signal rather than raw conversation noise.

## Implementing Semantic Search

