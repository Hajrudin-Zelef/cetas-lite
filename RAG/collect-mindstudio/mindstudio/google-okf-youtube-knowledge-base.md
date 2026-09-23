---
id: collect-mindstudio/mindstudio/google-okf-youtube-knowledge-base
title: "Turn Any YouTube Channel Into a Searchable AI Knowledge Base with OKF"
domain: mindstudio
role: reference
task: article
actors: ["Google"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "compute", "cost", "mcp", "memory", "reasoning"]
source: docs/RAG/Collect RAG/02_mindstudio/google-okf-youtube-knowledge-base.md
source_anchor: ""
source_lines: [1, 58]
sha256: 1a2320aa79bc636a1e942de2857781c9c41903015a6c60a96229f56295d1b7ca
---

# Turn Any YouTube Channel Into a Searchable AI Knowledge Base with OKF

## Metadata

- **Source** : https://www.mindstudio.ai/blog/google-okf-youtube-knowledge-base
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how Google's Open Knowledge Format (OKF) can turn an entire YouTube channel's back catalog into a structured, cross-referenced, agent-readable knowledge base. OKF is a specification from Google for structuring knowledge bases that AI agents can navigate on their own. It grew out of Andrej Karpathy's "LLM Wiki" concept — a wiki built and organized by a language model instead of by hand — and turns it into a shared standard, analogous to how MCP standardized agent-to-tool calls and A2A standardized agent-to-agent communication. With OKF, a knowledge base built by one person follows a predictable structure any agent can read, making knowledge bases shareable objects rather than one-off personal projects.

Applied to YouTube, a channel's hundreds of videos can become an archive an agent can search, summarize, and cite with timestamps — instead of scrubbing hours of footage, you ask a question and get an answer pulled from the exact video and moment where it was covered.

The structure is wiki-like but entirely plain markdown. At the top is an index file — the first thing an agent reads — laying out the "themes" and the high-level map of what questions the archive answers. From there the agent drills into two main folders: concepts and entities. Concepts are ideas, frameworks, or patterns discussed across multiple videos (e.g., a named workflow or a recurring critique). Entities are concrete tools, protocols, or products mentioned repeatedly (e.g., a specific agent framework or coding tool). Each concept/entity gets its own markdown file aggregating everything said about it across the channel, with links back to specific videos and timestamps. Underneath sits a raw folder with full markdown transcripts of every video, timestamped sentence by sentence — the ground-truth layer enabling exact citations. Because every file is markdown and linked, opening the base in Obsidian renders an actual graph view showing how ideas relate across the corpus.

Canonicalization is the core step: an LLM reads across all transcripts to determine what deserves a dedicated file. People describe the same idea differently (acronyms, fuller phrases), so keyword matching alone fails; an LLM must reason about meaning to merge variant phrasings into one canonical file. The reverse also applies: a tool mentioned once in passing doesn't get an entity page — only concepts/entities recurring across multiple videos earn a file. This filtering matters for scale, keeping the base dense with signal rather than unmanageable.

Usage: querying looks like normal conversation with an agent, except answers come with sourced, timestamped citations. Broad questions (walk through a multi-step process) trace through index, concept, and entity files; narrow factual questions resolve into short summaries with links into the original source.

Setup: accessing an existing OKF base is a one-step prompt (point an agent at the repository/README; it clones the base and learns OKF structure if needed). Building from scratch: extract transcripts (markdown, with timestamps), canonicalize across all transcripts, generate linked concept/entity files with citations, build the index. Skills can take just a channel URL as input. The catch: YouTube's Data API only exposes transcripts for channels you own, so arbitrary channels require third-party transcript services (paid = more reliable, free = less consistent).

Trade-off: real compute and token cost upfront (canonicalization reasons across the whole corpus, not document-by-document), in exchange for a permanently queryable, source-cited archive. OKF differs from vector-DB RAG: a vector DB retrieves chunks by semantic similarity without inherent structure, while OKF is explicitly organized (index, canonicalized concepts/entities, linked graph) so an agent navigates relationships rather than retrieving isolated passages.

## Key points

- OKF is Google's standard for agent-readable knowledge bases, extending Karpathy's LLM Wiki concept into a shareable, reproducible format (alongside MCP and A2A as agent standards).
- Structure: index file → concepts folder → entities folder → raw folder of timestamped markdown transcripts, all plain markdown.
- Canonicalization (LLM reasoning across all transcripts) merges variant phrasings and filters one-off mentions; only recurring concepts/entities get files — this keeps the base navigable.
- YouTube transcripts become the raw material; answers come with timestamped citations pointing to the exact video and moment.
- Accessing an existing base = a single prompt; building one = a defined extraction + canonicalization pipeline (transcript services vary in cost/reliability for non-owned channels).
- OKF differs from vector-DB RAG: explicit structure and navigable relationships vs semantic chunk retrieval.
- Cost: meaningful LLM compute/token cost upfront for a permanently queryable, source-cited archive.

## Technical data / figures

| Layer | Content |
|---|---|
| Index file | High-level themes/map of what the archive answers |
| Concepts folder | Ideas, frameworks, patterns recurring across videos (one file each) |
| Entities folder | Concrete tools, protocols, products mentioned repeatedly (one file each) |
| Raw folder | Full markdown transcripts, timestamped sentence by sentence (ground truth) |
| Canonicalization | LLM merges variant phrasings; filters one-off mentions; only recurring items get files |
| Citation model | Answers cite exact video + timestamp; viewable as graph in Obsidian |
| Transcript source | YouTube Data API (own channels) vs third-party services (paid/more reliable, free/less consistent) |

## Why this source matters for the RAG

Describes a structured, citation-preserving knowledge-base format (OKF) that is a direct alternative or complement to vector-database RAG — useful for building RAG on video/transcript content with traceable sources. Provides the canonicalization and index/entity/concept architecture that can be replicated for any multi-document RAG corpus, plus a comparison of OKF vs vector retrieval that informs architecture choice.

## Related context from the article

- OKF extends Karpathy's LLM Wiki concept, same lineage as the Google Wiki Skill agent memory work.
- Positioned alongside MCP (agent-tool) and A2A (agent-agent) as Google-defined agent standards.
- Building a base for hundreds of videos requires corpus-wide reasoning (canonicalization), not per-document processing.
