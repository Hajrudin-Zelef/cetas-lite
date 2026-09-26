---
id: collect-240926-mindstudio/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall-2
title: "how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "memory", "agent", "agents", "embedding", "embeddings", "mcp", "model context protocol", "research", "tool use"]
source: docs/RAG/clean_en/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall.md
source_anchor: ""
source_lines: [125, 229]
sha256: 41b3f2bc94f2995ef9da40caac8fd1866fafd5dc90d8cee872ea1045de9c210b
---

# how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall

For example: if you query GBrain about the authentication module, it doesn’t just return memories explicitly tagged with “auth.” It also returns memories about the session token implementation, the middleware that depends on auth, and the bug you fixed in the login flow three weeks ago — because those nodes are connected in the graph.

This relational recall is the key differentiator. For complex codebases where concepts are deeply interconnected, graph traversal produces more useful recall than pure embedding similarity.

### MCP Server Integration

GBrain exposes itself as an MCP (Model Context Protocol) server, which means Claude Code can call it directly during a session using tool use. The workflow looks like this:

1. Claude determines it needs prior context on a topic
2. It calls the GBrain MCP tool with a query
3. GBrain runs a graph search and returns relevant memories
4. Claude incorporates those memories into its response

This is real-time recall — not just upfront injection. Claude can query memory as needed throughout a session rather than relying on what was loaded at startup.

### Setting Up GBrain as an MCP Server

Claude Code’s MCP configuration lives in `.claude/mcp_config.json`. Adding GBrain as a server is a matter of registering the endpoint and providing auth credentials. Once registered, GBrain appears as a callable tool in Claude’s toolset, and Claude can invoke it autonomously when it determines it needs historical context.

## Building a Complete Memory Stack

You don’t have to choose between these tools — they address different layers and can work together.

### A Practical Architecture

Here’s a stack that covers all three layers:

| Layer | Tool | When It Runs | 
|---|---|---|
| Storage | Memarch | After each session, on memory write | 
| Injection | Hermes | At session start | 
| Recall | GBrain | Mid-session, on demand | 

**Workflow:**

1. During a session, Claude records significant decisions, bugs, and context changes to Memarch via a write hook or manual command.
2. When the next session starts, Hermes queries Memarch, scores memories against the current task, and injects relevant context into `CLAUDE.md` .
3. During the session, Claude can call GBrain’s MCP endpoint for deeper recall on specific topics.

This layered approach gives you coverage at every point in the session lifecycle — before, during, and after.

### A Simpler Starting Point

If the full stack feels like overkill, start with Hermes alone:

1. Manually write key memories to a structured JSON or YAML store after each session
2. Use Hermes to score and inject them at session start
3. Add Memarch or GBrain later when your memory store grows and retrieval quality degrades

Don’t over-engineer the memory layer before you have real volume. A few hundred well-structured memories will perform better than a sophisticated system with poorly formatted content.

## Common Mistakes When Adding Memory to Claude Code

A few patterns that consistently cause problems:

**Storing too much.** Memory dumps from full session transcripts are low signal. Store decisions, not dialogue. Store the outcome of a debugging session, not the debugging process.

**No memory hygiene.** Memories go stale. A decision you made in month one might be wrong by month three. Build a review cycle for old memories — either manual or automated — to mark outdated records.

**Context overloading.** Injecting 50 memories because you can fit them in context doesn’t mean you should. Token budget spent on marginally relevant memories is budget not available for the actual task. Be aggressive about trimming injection payloads.

**Skipping the write step.** The best memory system is useless if you never populate it. Build the habit of writing key context to memory at the end of every session. Some teams automate this with a session-end hook that prompts Claude to summarize what should be recorded.

**Treating CLAUDE.md as the only option.** CLAUDE.md is fine for static project context. It’s not a replacement for a real memory system on any project with more than a few weeks of history.

## How MindStudio Handles Persistent Memory for AI Agents

If you’re thinking about memory for Claude Code specifically, that’s one use case. But the broader problem — AI agents that lose context between sessions — shows up everywhere: customer support bots that forget users, research agents that repeat work, workflow agents that can’t build on past outputs.

MindStudio handles this at the platform level. When you build an agent in MindStudio, you get built-in data storage that persists across sessions, natively scoped to users, conversations, or global state. You don’t need to wire up a separate memory layer — it’s part of the agent builder.

For developers who want to extend this to external agents (including Claude Code), MindStudio’s Agent Skills Plugin exposes storage and retrieval as simple method calls. Your Claude Code agent can call `agent.storeMemory()` and `agent.recallMemory()` without managing the storage infrastructure yourself.

The platform also supports building fully autonomous background agents that can run on a schedule — which opens up patterns like nightly memory consolidation jobs that summarize, deduplicate, and restructure what an agent learned during the day.

You can try MindStudio free at mindstudio.ai.

## Frequently Asked Questions

### Does Claude Code have any built-in persistent memory?

The only native persistence is `CLAUDE.md` — a static file Claude reads at session start. It doesn’t update automatically, doesn’t grow based on session activity, and doesn’t support semantic retrieval. For any project with meaningful history, you need an external memory layer.

### What is MCP and how does it enable memory for Claude Code?

MCP (Model Context Protocol) is Anthropic’s standard for giving Claude access to external tools and data sources. Memory systems like GBrain can expose themselves as MCP servers, allowing Claude to call them directly during a session. This enables real-time recall — Claude can query memory mid-session rather than relying only on what was injected at startup. You configure MCP servers in `.claude/mcp_config.json`.

### What’s the difference between storage, injection, and recall?

These are three distinct operations in a memory system. **Storage** is writing memories to a persistent location after they’re created. **Injection** is loading relevant memories into context at the start of a session. **Recall** is retrieving memories on demand during a session, typically using semantic or graph search. A complete memory layer needs all three — many tools only address one or two.

### How do I decide what to store in Claude Code’s memory?

Focus on things that would take time to re-establish if Claude forgot them: architectural decisions, naming conventions, recurring bugs and their fixes, module relationships, and active task state. Avoid storing raw conversation transcripts or process details — they add noise without adding signal. A good rule of thumb: if you’d have to explain it again from scratch next session, it should be in memory.

### Can I use vector search for Claude Code memory?

Yes. Vector (embedding) search is a solid foundation for recall — you store memories as embeddings and retrieve by semantic similarity. Tools like Hermes use this for injection scoring. The limitation is that vector search returns individual similar records but misses relational context. Graph-based systems like GBrain can surface connected memories that a pure vector search would miss. For most projects, vector search is a good starting point.

### How much memory context should I inject per session?

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

