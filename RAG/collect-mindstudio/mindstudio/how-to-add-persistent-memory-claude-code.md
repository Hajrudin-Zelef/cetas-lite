---
id: collect-mindstudio/mindstudio/how-to-add-persistent-memory-claude-code
title: "How to Add Persistent Memory to Claude Code: Storage, Injection, and Recall"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["claude", "memory", "agent", "agents", "context window", "embeddings", "lean", "mcp", "model context protocol", "reasoning", "research", "tool use"]
source: docs/RAG/Collect RAG/02_mindstudio/how-to-add-persistent-memory-claude-code.md
source_anchor: ""
source_lines: [1, 80]
sha256: 0d399241d77aea97366a9d2b8c8d9135e1e40d6a3b2848c5142b276ac77ac290
---

# How to Add Persistent Memory to Claude Code: Storage, Injection, and Recall

## Metadata

- **Source** : https://www.mindstudio.ai/blog/how-to-add-persistent-memory-claude-code
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article addresses Claude Code's persistent memory gap: out of the box it remembers nothing between sessions. No record of architecture decisions, renames, or debugging history. The only native persistence is CLAUDE.md — a static file read at session start. It doesn't update automatically, doesn't grow or reorganize based on session activity, and doesn't support semantic retrieval — "a sticky note, not a memory system." What gets lost between sessions: architectural decisions and reasoning, bug fixes and causes, file structure changes/refactors, custom patterns/conventions, task history, external research surfaced during sessions.

The three pillars of persistent memory (each independently addressable): Storage — where memories live and how they're structured (flat file, database, vector store, or knowledge graph); Injection — how the right memories load into context at session start or during (not all memories are relevant all the time, so a selection mechanism is needed); Recall — how memories are retrieved mid-session (exact lookups and semantic similarity search).

Three tools representing distinct approaches:
- Memarch (storage): organizes memories into typed records rather than freeform text (a decision memory differs from a bug memory). Tagging with category and metadata fields makes retrieval reliable. Deduplication at write time — merging related records instead of creating copies, preventing accumulation noise. Best for: high volume of decisions/conventions, queryable-by-type memory, storage that grows cleanly without manual curation.
- Hermes (injection): sits between the memory store and the session. At session start, scores stored memories against current task/project context and injects the highest-relevance subset. Relevance scoring combines recency weighting, tag matching, and semantic similarity (embeddings vs task description); output is a ranked, trimmed set sized to a target context budget. Integration via a startup hook — a script runs before the session begins, builds the injection payload, writes to CLAUDE.md or a context file Claude reads at initialization. Injection stays transparent.
- GBrain (semantic recall): uses a knowledge-graph structure (memories = nodes, relationships = edges) rather than a flat vector store. Graph traversal surfaces related context a pure vector search would miss (querying "auth module" also returns session-token implementation, middleware dependencies, login-flow bug). Exposes itself as an MCP (Model Context Protocol) server, so Claude Code calls it directly mid-session via tool use (.claude/mcp_config.json registration). Real-time recall, not just upfront injection.

Building a complete memory stack — the layers are complementary: Storage = Memarch (after each session, on memory write); Injection = Hermes (at session start); Recall = GBrain (mid-session, on demand). Workflow: Claude records decisions to Memarch via write hook; Hermes queries Memarch, scores against current task, injects into CLAUDE.md; Claude calls GBrain's MCP endpoint for deeper recall. Simpler start: manual structured JSON/YAML store + Hermes alone; add Memarch/GBrain later as volume grows. Don't over-engineer before real volume — "a few hundred well-structured memories outperform a sophisticated system with poorly formatted content."

Common mistakes: storing too much (store decisions, not dialogue); no memory hygiene (memories go stale — build a review cycle); context overloading (be aggressive about trimming injection payloads); skipping the write step (build a session-end hook); treating CLAUDE.md as the only option.

Practical heuristics: inject under 20-30% of effective context window with memory; 10-20 high-relevance memories per session is a good balance. MCP config lives in .claude/mcp_config.json.

## Key points

- Claude Code's only built-in persistence is static CLAUDE.md — no automatic persistence, growth, or semantic retrieval between sessions.
- A complete memory layer needs three things: storage, injection (session start), and recall (mid-session).
- Memarch = structured storage (typed records + dedup at write time); Hermes = relevance-scored injection at startup; GBrain = graph-based recall via MCP.
- The tools address different layers and can be combined: Memarch for storage, Hermes for injection, GBrain for recall.
- Vector search returns similar records but misses relational context; graph traversal (GBrain) surfaces connected memories.
- Keep memory lean: store decisions not dialogue; inject 10-20 high-relevance memories or under 20-30% of context window.
- Start simple (Hermes + structured store), add layers as memory volume grows.

## Technical data / figures

| Pillar | Role | Tool |
|---|---|---|
| Storage | Where/how memories live | Memarch (typed records, dedup) |
| Injection | Loading relevant memories at session start | Hermes (relevance scoring) |
| Recall | Retrieving memories mid-session | GBrain (knowledge graph, MCP server) |

| Memarch traits | Details |
|---|---|
| Record typing | Decisions, bugs, conventions, dependencies structured differently |
| Dedup | At write time; merges related records |
| Best when | High volume of decisions/conventions; query by type; clean growth |

| Hermes traits | Details |
|---|---|
| Scoring | Recency weighting + tag matching + semantic similarity |
| Output | Ranked, trimmed set sized to context budget |
| Integration | Startup hook writes payload to CLAUDE.md/context file |

| GBrain traits | Details |
|---|---|
| Structure | Knowledge graph (nodes + edges); relational recall |
| Integration | MCP server registered in .claude/mcp_config.json; called as a tool mid-session |
| Example | Query "auth module" → also returns token implementation, middleware, login-flow bug |

| Heuristic | Value |
|---|---|
| Injected memory budget | <20–30% of effective context window |
| High-relevance memories per session | 10–20 |
| MCP config location | .claude/mcp_config.json |

## Why this source matters for the RAG

Defines the three-layer memory architecture (storage/injection/recall) that underlies persistent RAG agents, with concrete tooling choices (typed records, relevance-scored injection, graph-based recall). Directly relevant to building retrieval-and-recall layers for Claude Code agents and to comparing vector vs graph retrieval for knowledge bases.

## Related context from the article

- Hermes (MindStudio orchestration) and Hermes Agent (Ollama-based) are distinct projects both mentioned in this source.
- MCP = Anthropic's standard for giving Claude access to external tools/data sources.
- MindStudio platform provides agent.storeMemory()/agent.recallMemory() and nightly memory-consolidation background agents.
- GBrain relational recall example: querying about a module surfaces connected memories a pure vector search would miss.
