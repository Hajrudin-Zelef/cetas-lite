---
id: collect-240926-mindstudio/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall-1
title: "how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "memory", "agent", "agents", "context window", "embeddings", "reasoning", "research"]
source: docs/RAG/clean_en/mindstudio/how-to-add-persistent-memory-to-claude-code-storage-injection-and-recall.md
source_anchor: ""
source_lines: [1, 124]
sha256: c7c4c0a1c03ab7dc1f1e5f31786296c83c649cd5e20d688d45586b33738c6a87
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

