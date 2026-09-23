---
id: collect-mindstudio/mindstudio/how-to-build-llm-wiki-knowledge-base-obsidian-claude-code
title: "How to Build an LLM Wiki Knowledge Base with Obsidian and Claude Code"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["claude", "attention", "cost", "fine-tuning", "llama", "rlhf"]
source: docs/RAG/Collect RAG/02_mindstudio/how-to-build-llm-wiki-knowledge-base-obsidian-claude-code.md
source_anchor: ""
source_lines: [1, 77]
sha256: 1d95b19aaea4f301328e2ce5b3412fd25ee692d88869eacd1962d39ee7568e75
---

# How to Build an LLM Wiki Knowledge Base with Obsidian and Claude Code

## Metadata

- **Source** : https://www.mindstudio.ai/blog/how-to-build-llm-wiki-knowledge-base-obsidian-claude-code
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains how to build a self-growing "LLM wiki" knowledge base using Obsidian and Claude Code, ingesting YouTube transcripts, PDFs, and web pages into structured, searchable, interlinked notes. The premise: most people forget ~80% of what they read within a week, and finding it again means re-reading from scratch. Karpathy has publicly described such a system as core to how he learns; with Claude Code and Obsidian you can build one in an afternoon.

What an LLM wiki is: a vault of markdown notes where an LLM processes every piece of useful content into a structured note fitting an existing knowledge graph. The "wiki" part matters — links between concepts (a new video on transformers should reference an existing "attention mechanisms" note), producing a network of interconnected ideas, not a pile of documents. The "LLM" part makes it scalable — Claude Code reads the existing vault, understands what's there, and writes notes that slot in coherently. Karpathy's vision: the knowledge base should grow with minimal friction.

Architecture — three layers: Layer 1 the Vault (Obsidian, local markdown files — Claude Code reads/writes directly, you own the data): /Sources (raw processed content), /Concepts (synthesized ideas), /MOCs (Maps of Content index notes), plus consistent frontmatter (title, tags, date, source URL). Layer 2 the Ingestion Pipeline (Claude Code extracts raw text, reads the vault for context, writes structured notes with links, updates MOC notes). Layer 3 the Growth Loop — each new note makes the next one better because Claude has more context ("RLHF" note becomes richer once "reward modeling" and "fine-tuning" notes exist).

Prerequisites: Obsidian with a vault, Claude Code (npm install -g @anthropic-ai/claude-code), Python 3.8+, yt-dlp (YouTube transcripts), PyMuPDF (PDF processing), Anthropic API key; optional Dataview and Templater plugins.

Setup steps: (1) Create folders /Sources/YouTube, /Sources/PDFs, /Sources/URLs, /Concepts, /MOCs, /Inbox, plus a critical _schema.md at the vault root defining note types (Source Note: frontmatter + Summary, Key Ideas, Quotes, Related Concepts wikilinks; Concept Note: Definition, Context, Examples, Related Concepts; MOC Note: Overview, Notes in cluster) and linking rules (always use [[wikilink]], prefer linking to Concept over Source notes, create stub concept notes). (2) Write three extraction scripts: YouTube (yt-dlp or youtube-transcript-api), PDF (PyMuPDF/fitz), URL (requests + BeautifulSoup, removing nav/footer/scripts). (3) Write the ingest_prompt.md giving Claude the curation workflow: read _schema.md, scan Concepts/, create Source note, identify 3-5 key concepts, check for existing concept notes (add link or create stub), update/create MOC, output summary; rules: never modify existing content only append, always [[wikilink]], Source summaries under 300 words, Concept notes timeless. Run via `claude-code --prompt ingest_prompt.md --var type=... --var source=... --var content="$(python extract_....py URL)"`. (4) Wrap in a single ingest.sh shell script so the whole user-facing workflow is one command per source type. (5) Improve quality over time: route new notes to /Inbox for review, add a glossary.md (prevents duplicate concepts when you and sources use different terms), run periodic synthesis (merge candidates, missing links, MOC candidates). (6) Edge cases: videos without transcripts → yt-dlp audio + local Whisper; large PDFs → chunk_text (max 50000 chars) then consolidate; duplicate detection via source-URL search in frontmatter.

Common mistakes: skipping the schema file (single most important thing); processing too much too fast (start with 10-20 sources); not linking back to sources (orphaned knowledge); ignoring concept sprawl (merge/delete concept notes under 100 words not updated in 30 days).

FAQ figures: a 60-min video (~10,000 words transcript) ingestion costs ~$0.10-0.30; a 20-page PDF under $0.10; typical personal monthly cost $5-20. Architecture is model-agnostic (adaptable to GPT-4o or local models via Ollama; Llama 3 70B handles the task reasonably). Multilingual content works (add language frontmatter field). Compares favorably to hosted Mem/Notion AI for control, at the cost of setup time.

## Key points

- LLM wiki = Obsidian markdown vault + Claude Code file-system awareness → self-growing network of interconnected notes from YouTube/PDF/URL sources.
- Three layers: Vault (Sources/Concepts/MOCs + Inbox), Ingestion Pipeline (Claude Code), Growth Loop (each note improves the next).
- The _schema.md file is the most important piece — it keeps Claude's output consistent across hundreds of notes (note types, sections, linking rules).
- Three extraction scripts (YouTube, PDF, URL) plus one ingest.sh shell script reduce ingestion to a single command.
- Bidirectional linking and concept stub creation build the knowledge graph automatically.
- Periodic synthesis runs prevent concept sprawl and find merge/link/MOC candidates.
- Costs: ~$0.10–0.30 per video ingestion, under $0.10 per PDF, $5–20/month typical personal use.

## Technical data / figures

| Layer | Role |
|---|---|
| Vault (Obsidian) | Local markdown; /Sources, /Concepts, /MOCs, /Inbox; _schema.md + frontmatter |
| Ingestion pipeline | Claude Code: extract → read vault → write structured note → update MOCs |
| Growth loop | New notes give Claude more context; wiki compounds |

| Note type | Sections |
|---|---|
| Source Note | Frontmatter; Summary; Key Ideas; Quotes; Related Concepts (wikilinks) |
| Concept Note | Definition; Context; Examples; Related Concepts |
| MOC Note | Overview; Notes in this cluster (wikilinks) |

| Ingestion cost | Value |
|---|---|
| 60-min video (~10K words) | ~$0.10–$0.30 |
| 20-page PDF | < $0.10 |
| Typical personal monthly | $5–$20 |
| Large-PDF chunk size | max 50,000 chars per chunk |
| Concept note cleanup rule | <100 words & not updated in 30 days → merge or delete |
| Recommended starting volume | 10–20 sources |

| Tool | Purpose |
|---|---|
| yt-dlp / youtube-transcript-api | YouTube transcripts |
| PyMuPDF (fitz) | PDF text extraction |
| requests + BeautifulSoup | Web page extraction |
| Whisper (local) | Videos without captions |

## Why this source matters for the RAG

Full implementation blueprint for a local-first, self-growing RAG knowledge base built on wiki-linked markdown: covers ingestion (YouTube/PDF/URL), schema design, note types, cross-linking, and synthesis. Directly useful for constructing and maintaining a retrieval corpus with traceable sources (frontmatter URLs) and for understanding Karpathy's LLM-wiki architecture that underpins OKF and related patterns.

## Related context from the article

- Karpathy's LLM wiki vision: knowledge base should grow with minimal friction.
- Glossary.md prevents duplicate concepts from vocabulary mismatch.
- /Inbox review step keeps quality high before content enters the main knowledge base.
- Model-agnostic architecture (GPT-4o or Ollama local models work too).
