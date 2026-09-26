---
id: collect-240926-mindstudio/mindstudio/how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent-2
title: "how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent"
domain: mindstudio
role: reference
task: reference
actors: []
dates: ["2024-11-15"]
keywords: ["agent", "agents", "memory", "benchmark", "context window", "embedding", "embeddings", "inference", "throughput"]
source: docs/RAG/clean_en/mindstudio/how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent.md
source_anchor: ""
source_lines: [157, 349]
sha256: a57fb5283ecb984d63b1cae3063b4caea63dea30a9f7b2f712265d3790d3a329
---

# how-to-build-a-semantic-memory-system-for-ai-agents-without-hermes-agent

Once you have stored memories, retrieval needs to be fast and relevant.

### Basic Retrieval Pattern

```
def retrieve_relevant_memories(
    query: str,
    collection,
    n_results: int = 5,
    min_relevance_threshold: float = 0.7
) -> list[str]:
    
    query_embedding = embed_text(query)
    
    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=n_results,
        include=["documents", "distances"]
    )
    
    # Filter by relevance threshold
    memories = []
    for doc, distance in zip(results["documents"][0], results["distances"][0]):
        # ChromaDB returns L2 distance; lower is more similar
        similarity = 1 - distance
        if similarity >= min_relevance_threshold:
            memories.append(doc)
    
    return memories
```
### Hybrid Search for Better Recall

Pure semantic search sometimes misses exact matches. A hybrid approach combines vector similarity with keyword filtering:

1. Run semantic search to get top-N candidates
2. Apply BM25 or simple keyword scoring as a re-ranking step
3. Return the re-ranked top results

For most agent use cases, pure semantic search with a reasonable threshold (0.65–0.75) is sufficient. Add hybrid search when your agents work with lots of proper nouns, codes, or identifiers that semantic models don’t handle well.

### Temporal Decay

Not all memories are equally current. A preference the user stated six months ago may have changed. Add a recency factor to your scoring:

```
import math
from datetime import datetime
def score_with_recency(similarity: float, timestamp: datetime, decay_rate: float = 0.01) -> float:
    days_old = (datetime.now() - timestamp).days
    recency_score = math.exp(-decay_rate * days_old)
    return similarity * 0.7 + recency_score * 0.3
```
Tune the weights based on your use case. A coding assistant should weight recency lightly (older patterns are still valid). A customer support agent should weight it heavily (account status changes frequently).

## Memory Injection Strategies

Retrieved memories are only useful if you inject them into the agent’s context at the right time, in the right format.

### System Prompt Injection

The simplest approach: prepend retrieved memories to the system prompt before each agent turn.

```
def build_system_prompt(base_prompt: str, relevant_memories: list[str]) -> str:
    if not relevant_memories:
        return base_prompt
    
    memory_block = "\n".join([f"- {mem}" for mem in relevant_memories])
    
    return f"""{base_prompt}
## Relevant Memory
The following information from previous interactions is relevant to this task:
{memory_block}
Use this context to inform your responses, but don't reference it explicitly unless asked."""
```
## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

### Structured Memory Context

For agents with longer tasks, a more structured format helps the model parse and use memories correctly:

```
## Long-Term Memory (from previous sessions)
[User Preferences]
- Prefers bullet points over paragraphs in summaries
- Works in EST timezone, availability 9am–5pm
[Project Context]
- Current project: Q3 financial report automation
- Last completed: data ingestion pipeline (completed 2024-11-15)
- Next milestone: visualization layer
[Known Issues]
- API rate limits hit when batch size > 100 records
```
### Episodic vs. Semantic Memory

Borrowing from cognitive science, it helps to separate two memory types:

- **Episodic memory** : Specific events (“On Tuesday, the user rejected the CSV format proposal”)
- **Semantic memory** : General knowledge the agent has built up (“This user prefers JSON output formats”)

Store them in separate collections with different retrieval weights. Episodic memories are useful for continuity; semantic memories are useful for preference and behavioral adaptation.

### Avoiding Context Window Bloat

Memory injection can eat your context window fast. Keep these constraints in mind:

- Set a hard cap on injected memories (5–8 is usually enough)
- Compress older memories with periodic summarization runs
- Use importance scoring to prioritize high-signal memories over low-signal ones
- Consider a two-tier system: hot memory (recent, in-context) and cold memory (older, retrieved on demand)

## Managing Memory Over Time

A memory system that only grows will eventually degrade in quality and performance.

### Deduplication

Before storing a new memory, check if a semantically similar one already exists:

```
def should_store_memory(new_memory: str, collection, threshold: float = 0.92) -> bool:
    results = collection.query(
        query_embeddings=[embed_text(new_memory)],
        n_results=1
    )
    
    if not results["distances"][0]:
        return True
    
    similarity = 1 - results["distances"][0][0]
    return similarity < threshold
```
### Periodic Consolidation

Run a consolidation job (daily or weekly) that:

1. Clusters similar memories together
2. Generates a summary of each cluster
3. Replaces the cluster with the summary
4. Removes memories below an importance threshold

This is the rough equivalent of how sleep consolidates human memory — you’re compressing episodes into denser, more useful representations.

### Memory Importance Scoring

Not all memories are worth keeping. Add an importance score when storing:

- High: User stated explicit preference, agent made a significant decision, error with resolution
- Medium: Useful context, task status
- Low: Routine confirmations, obvious inferences

Drop low-importance memories after 30 days unless they’ve been accessed recently.

## Common Mistakes to Avoid

### Storing Too Much, Too Literally

The most common mistake is treating memory as a log. You don’t want a transcript — you want distilled signal. Always run a summarization/extraction step before storing.

### Ignoring Retrieval Quality

Building storage is the easy part. Retrieval quality determines whether the system actually helps. Test your semantic search with real agent queries. Tune your similarity thresholds. Add hybrid search if pure semantic search underperforms.

### Single Collection for Everything

Separate your memory types. User preferences, project context, error history, and episodic events should live in different collections with different retrieval strategies.

### No Memory Validation

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

Agents can store incorrect memories. Add a confidence layer: if the memory was derived from an uncertain inference rather than an explicit statement, tag it accordingly and weight it lower during retrieval.

### Forgetting Write Performance

For agents running at high frequency, write performance matters. ChromaDB in-process mode is fine for low-throughput agents. For anything that writes memories at high frequency, benchmark your setup or move to Qdrant with a persistent server.

## Frequently Asked Questions

### What is a semantic memory system for AI agents?

A semantic memory system stores agent experiences, observations, and context as vector embeddings and retrieves them using semantic similarity search. Unlike keyword search, it finds relevant memories based on meaning — so a query about “user formatting preferences” will surface memories about layout choices even if those exact words weren’t used.

### How is this different from just using a long context window?

Long context windows include everything — relevant or not. A semantic memory system selectively retrieves only the most relevant information for a given query. This means less noise, lower token costs, and the ability to retain knowledge across sessions without hitting context limits.

### Which vector database should I use for a local agent?

