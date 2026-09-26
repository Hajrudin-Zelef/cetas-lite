---
id: collect-mindstudio/mindstudio/build-semantic-memory-system-ai-agents-without-hermes-1
title: "How to Build a Semantic Memory System for AI Agents Without Hermes Agent"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Hugging Face", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "memory", "claude", "context window", "cost", "embedding", "embeddings", "latency", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/build-semantic-memory-system-ai-agents-without-hermes.md
source_anchor: ""
source_lines: [1, 71]
sha256: 3b4e3e4b5844b8434bd86b200e425c8b2ab55a06273f54e33a9264327bc2abac
---

# How to Build a Semantic Memory System for AI Agents Without Hermes Agent

## Metadata

- **Source** : https://www.mindstudio.ai/blog/build-semantic-memory-system-ai-agents-without-hermes
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide shows how to build a semantic memory system for AI agents from scratch using local vector databases, embedding pipelines, and smart injection logic — demonstrating that Hermes Agent (built on Ollama and local models) proved memory matters, but the capability can be replicated inside Claude Code or any custom agent framework without lock-in.

What Hermes Agent demonstrated: (1) retrieval-augmented context matters more than raw model size in long-horizon tasks; (2) storing memories as vector embeddings (rather than raw text logs) enables semantic retrieval, not just keyword matching; (3) agents with memory made fewer contradictory decisions across sessions; (4) the retrieval step could run locally without meaningful latency. Its limitation was portability/flexibility — it worked within its own ecosystem.

Anatomy of a semantic memory system — three independent, swappable functional layers: (1) Storage — a vector database holding embeddings alongside original content and metadata; (2) The embedding pipeline — converting text (agent observations, user inputs, decisions, outcomes) into vectors so semantically-similar sentences have similar vectors; (3) Retrieval and injection — querying the store with the current context as search key and injecting the most relevant memories into the agent's prompt/context.

Local vector database options: ChromaDB — most accessible, runs in-process (no separate server for small deployments), clean Python API, native LangChain/LlamaIndex integration; good for prototyping/local dev/small-medium deployments, not ideal for high-concurrency production. Qdrant — purpose-built Rust-core vector search engine, fast, supports filtering, named vectors, payload indexing; runs locally via Docker; production-grade. FAISS (Facebook AI Similarity Search) — library not database, extremely fast pure similarity search but no native persistence/metadata/CRUD — wrap it yourself. LanceDB — newer, columnar on-disk vectors, serverless. Recommendation: start with ChromaDB (running in under 10 minutes), move to Qdrant for more control.

Embedding pipeline: local model options — nomic-embed-text via Ollama (768 dims, solid general-purpose), mxbai-embed-large via Ollama (1024 dims, higher quality, slower), text-embedding-3-small via OpenAI API (excellent quality, minimal cost), sentence-transformers/all-MiniLM-L6-v2 via HuggingFace (fast, lightweight, fully offline). What to embed: summarized observations ("User confirmed report format should be CSV, not JSON"), agent decisions with context, task outcomes, user preferences revealed in interaction, errors and how they were resolved. What to avoid: raw unprocessed transcripts (noisy), ephemeral tool outputs, duplicate entries (use deduplication). Memory extraction step: have the agent/summarizer extract memory-worthy content from each conversation turn (1-3 sentences max) before storing — dramatically improves retrieval quality.

Semantic search: basic retrieval pattern queries with n_results=5 and min_relevance_threshold=0.7 (ChromaDB L2 distance; similarity = 1 - distance). Hybrid search (vector + BM25/keyword re-ranking) recommended when agents handle many proper nouns, codes, or identifiers. Temporal decay: add a recency factor — score = similarity*0.7 + recency_score*0.3, with recency_score = exp(-decay_rate*days_old); a coding assistant weights recency lightly, a customer support agent heavily.

Memory injection: system prompt injection — prepend retrieved memories to the system prompt per turn ("## Relevant Memory" block); structured memory context format (User Preferences / Project Context / Known Issues) for longer tasks; episodic vs semantic memory separation (specific events vs general knowledge) in different collections with different weights; avoid context window bloat — hard cap injected memories at 5-8, compress older memories via periodic summarization, use importance scoring, consider hot (recent in-context) vs cold (on-demand) two-tier memory.

Managing memory over time: deduplication before storing (semantic similarity threshold ~0.92 → skip storing near-duplicates); periodic consolidation (daily/weekly) clustering similar memories, generating a cluster summary, replacing clusters with summaries, removing below-importance memories — "the rough equivalent of how sleep consolidates human memory"; importance scoring (High: explicit preference/significant decision/error+resolution; Medium: useful context/task status; Low: routine confirmations — drop after 30 days unless recently accessed).

Common mistakes: storing too much too literally (memory is distilled signal, not a log); ignoring retrieval quality (test with real queries, tune thresholds, add hybrid search); single collection for everything (separate preference/project/error/episodic collections); no memory validation (add confidence layer tagging uncertain inferences, weight lower); forgetting write performance at high frequency (ChromaDB in-process fine for low-throughput; move to Qdrant persistent server).

## Key points

- Semantic memory = three independent layers: storage (vector DB), embedding pipeline, retrieval + injection.
- ChromaDB is the fastest local starting point; Qdrant (Docker) for production-grade/filtered retrieval.
- Always summarize/extract signal before storing — never log raw conversations.
- Inject 3–8 relevant memories per agent turn using similarity thresholds, not fixed counts.
- Separate episodic (events) from semantic (preferences/patterns) memories in different collections.
- Add temporal decay, deduplication, and periodic consolidation to keep memory quality high.
- Local embedding models: nomic-embed-text (768d), mxbai-embed-large (1024d), all-MiniLM-L6-v2 (offline).

## Technical data / figures

| Vector DB | Strengths | Limitations |
|---|---|---|
| ChromaDB | In-process, clean Python API, LangChain/LlamaIndex native | Not for high-concurrency production |
| Qdrant | Rust core, fast, filtering, named vectors, payload indexing | Requires Docker; more setup |
| FAISS | Very fast pure similarity search | No built-in persistence/metadata/CRUD |
| LanceDB | Columnar on-disk, serverless | Newer entrant |

| Embedding model | Dimensions | Notes |
|---|---|---|
| nomic-embed-text (Ollama) | 768 | Solid general-purpose, local |
| mxbai-embed-large (Ollama) | 1024 | Higher quality, slower |
| text-embedding-3-small (OpenAI API) | — | Excellent quality, minimal cost |
| all-MiniLM-L6-v2 (HuggingFace) | — | Fast, lightweight, fully offline |

| Parameter | Typical value |
|---|---|
| Retrieved memories per turn | 3–8 |
| Relevance threshold | 0.65–0.75 (pure semantic); 0.7 example |
| Deduplication similarity threshold | ~0.92 |
| Recency scoring | similarity*0.7 + recency*0.3 (decay_rate ~0.01) |
| Embedding recommendation | Summarize before embedding; never raw transcripts |

## Why this source matters for the RAG

Provides the complete, code-level architecture for local semantic memory — storage, embedding, retrieval, injection, and lifecycle management — which is the retrieval half of a RAG system applied to agent memory. Directly useful for designing RAG pipelines (vector DB selection, embedding models, hybrid search, temporal decay) and for persistent-memory agent design in Claude Code or custom frameworks.

## Related context from the article

