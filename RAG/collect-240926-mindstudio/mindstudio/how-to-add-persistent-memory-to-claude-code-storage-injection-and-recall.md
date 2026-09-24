---
id: collect-240926-mindstudio/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall
title: "how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "memory", "agent", "agents", "context window", "embedding", "embeddings", "lean", "mcp", "model context protocol", "reasoning", "research"]
source: docs/RAG/clean_en/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall.md
source_anchor: ""
source_lines: [1, 240]
sha256: 56cf78d648e415c0f7f69378e04095144bf970467ba56af54f420900ad33080e
---

# how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall

<!-- source: https://www.mindstudio.ai/blog/how-to-add-persistent-memory-claude-code -->

## Claude Code’s Memory Problem Is Worse Than You Think

Claude Code is an impressive coding agent out of the box. But there’s a gap most people hit within the first few sessions: it remembers nothing.

Every time you start a new Claude Code session, you’re starting from scratch. No record of the architecture decisions you made last Tuesday. No awareness that you renamed the authentication module. No memory of the debugging rabbit hole that took three hours and ended with a one-line fix. The persistent memory problem is one of the biggest practical friction points when using Claude Code for real, ongoing projects.

This guide covers how to fix that — through proper storage, context injection, and semantic recall. Tools like Memarch, Hermes, and GBrain represent three distinct approaches to solving this, and understanding the difference between them will help you build a memory layer that actually holds up in production.

## What Claude Code Actually Remembers (And What It Doesn’t)

Before adding any memory tooling, it helps to understand what Claude Code does and doesn’t retain natively.

### CLAUDE.md: The Only Built-In Persistence

Claude Code reads a `CLAUDE.md` file at the root of your project on session start. This is the only real persistent memory it has out of the box. You can put anything in it — project context, tech stack, naming conventions, known issues — and Claude will pick it up.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

But `CLAUDE.md` is a flat text file. It doesn’t update automatically. It doesn’t grow or reorganize based on what you do during sessions. It doesn’t retrieve relevant context based on what you’re currently working on. It’s a sticky note, not a memory system.

### Context Window Limits

Claude’s context window is large, but it’s session-scoped. Once you close a session, everything that happened in that conversation is gone. You can paste in prior context manually, but that doesn’t scale.

### What Gets Lost

Here’s what actually disappears between sessions without a memory layer:

- Architectural decisions and the reasoning behind them
- Bug fixes and what caused the bugs
- File structure changes and refactors
- Custom patterns, conventions, or shortcuts you’ve established
- Task history and what was completed, in-progress, or abandoned
- External research you surfaced during a session

That’s not a minor gap. For any project running longer than a few days, this is a real productivity drain.

## The Three Pillars of Persistent Memory

Any serious memory system for Claude Code needs to handle three things independently:

**Storage** — Where memories live and how they’re structured. This could be a flat file, a database, a vector store, or a knowledge graph.

**Injection** — How the right memories get loaded into Claude’s context at the start of a session or during one. Not all memories are relevant all the time, so you need a mechanism to decide what to include.

**Recall** — How memories are retrieved when Claude needs them mid-session. This includes both exact lookups and semantic similarity search.

Different tools focus on different layers of this stack. Understanding which layer a tool operates on helps you decide what to reach for and how to combine tools effectively.

## Memarch: Structuring What Gets Stored

Memarch is built around the storage problem. The core insight it’s designed around is that unstructured memory is nearly useless at scale. If you dump session transcripts into a file and try to search it later, you’ll hit noise, contradictions, and retrieval failures.

### Memory as Typed Records

Memarch organizes memories into typed records rather than freeform text. A decision memory looks different from a bug memory. A convention memory gets structured differently from a dependency note. This typing is important because it determines how memories get summarized, deduplicated, and retrieved later.

At a practical level, this means when you record a memory through Memarch, you’re tagging it with a category and a set of metadata fields — not just appending text to a log. That structure makes retrieval far more reliable.

### Deduplication and Merging

One problem with naive memory systems is accumulation. If you note the same architectural pattern five times across five sessions, you get five copies of that memory. Memarch handles deduplication at write time, merging related records and updating existing ones rather than creating new copies.

This matters for injection quality. If your context gets loaded with ten nearly-identical memories, you’re wasting token budget and introducing noise.

### When to Use Memarch

Memarch works best when:

- Your project has a high volume of decisions, conventions, and evolving structure
- You need memory to be queryable by type (e.g., “show me all architecture decisions”)
- You want storage that grows cleanly over time without manual curation

## Hermes: Getting Memory Into Context

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Storage is only useful if the right memories actually reach Claude. That’s the injection problem, and it’s where Hermes focuses.

### Context Injection Patterns

There are a few ways to inject memory into a Claude Code session:

1. **Prepend to CLAUDE.md** — Write relevant memories into`CLAUDE.md` before the session starts, based on what task is active.
2. **System prompt injection** — Load memories into the system prompt dynamically.
3. **In-session injection** — Surface memories mid-session when Claude detects it’s entering a domain where prior context exists.

Hermes operates primarily as an injection layer, sitting between your memory store and the Claude Code session. When you start a session, Hermes scores stored memories against the current task or project context and injects the highest-relevance subset.

### Relevance Scoring

Raw injection — loading everything into context — defeats the purpose. Token budgets are real, and flooding context with tangentially related memories is worse than loading nothing.

Hermes applies relevance scoring using a combination of:

- **Recency weighting** — More recent memories score higher by default
- **Tag matching** — Memories tagged with the current task type rank up
- **Semantic similarity** — Memory embeddings are compared against the current task description

The output is a ranked, trimmed set of memories sized to fit a target context budget. You can configure how aggressive the trimming is.

### Hooking Into Claude Code’s Startup

The practical integration path with Hermes is via a startup hook. Claude Code supports running a script before a session begins — you can use that to call Hermes, build the injection payload, and write it to `CLAUDE.md` or a context file that Claude reads at initialization.

This keeps the injection transparent. Claude doesn’t need to know it’s reading injected memories vs. static documentation. It just gets context.

## GBrain: Semantic Recall During Sessions

Storage and injection handle what Claude knows at the start of a session. But what about mid-session recall? That’s where GBrain enters.

### Graph-Based Memory

GBrain uses a knowledge graph structure rather than a flat vector store. Memories are nodes. Relationships between them are edges. When you recall a memory, GBrain can traverse the graph — not just finding the closest match, but also surfacing related context that a pure vector search would miss.

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

There’s no universal answer, but a practical heuristic is to stay under 20–30% of your effective context window with injected memory. That leaves enough headroom for the actual task, codebase context, and Claude’s outputs. Most projects hit a good balance with 10–20 high-relevance memories per session rather than trying to inject everything.

## Key Takeaways

- Claude Code’s built-in memory is limited to static `CLAUDE.md` — there’s no automatic persistence between sessions.
- A complete persistent memory layer needs three things: **storage** (where memories live),**injection** (loading relevant memories into context at session start), and**recall** (querying memories mid-session).
- Memarch handles structured storage with typed records and deduplication. Hermes handles context injection with relevance scoring. GBrain handles graph-based semantic recall via MCP.
- These tools address different layers and can be combined — start with the layer that creates the most friction and add others as your memory store grows.
- Common failure modes: storing too much noise, skipping the write step, and over-injecting into context. Keep memory lean and high-signal.

If you want persistent memory handled at the platform level — without managing your own storage infrastructure — MindStudio builds it into the agent builder, with integrations that extend to external agents including Claude Code.
