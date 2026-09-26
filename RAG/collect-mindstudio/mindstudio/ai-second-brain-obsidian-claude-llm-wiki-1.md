---
id: collect-mindstudio/mindstudio/ai-second-brain-obsidian-claude-llm-wiki-1
title: "How to Build an AI Second Brain with Obsidian and Claude: The LLM Wiki Method"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["claude"]
source: docs/RAG/Collect RAG/02_mindstudio/ai-second-brain-obsidian-claude-llm-wiki.md
source_anchor: ""
source_lines: [1, 70]
sha256: ce857b7479b34c883ecaca11979f0f94d767e979b8577811521b51f4f544d9d3
---

# How to Build an AI Second Brain with Obsidian and Claude: The LLM Wiki Method

## Metadata

- **Source** : https://www.mindstudio.ai/blog/ai-second-brain-obsidian-claude-llm-wiki
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains the "LLM Wiki method": using a large language model (here Claude) as an active knowledge worker inside an Obsidian vault — not a one-off summarizer but a continuous editor that ingests, links, synthesizes, and briefs. The premise: most knowledge management systems are good at storing information and terrible at connecting it; the original "second brain" vision (Tiago Forte) was manual — humans did the tagging, linking, and synthesizing, which is cognitively expensive and abandoned within a month. Combining Claude with Obsidian's local-first, Markdown-based vault structure builds an AI second brain that maintains itself.

What an LLM Wiki is: a knowledge vault where the LLM acts as continuous editor, handling parsing and summarizing incoming content, extracting entities (people, concepts, projects, dates), creating bidirectional links between related notes, writing synthesis notes connecting ideas across sources, and generating daily/weekly briefings. Obsidian is the ideal host — a vault is just a folder of Markdown files on the local machine, no proprietary database, every note a plain text file any script/API can read and write. Claude is the ideal LLM: long context windows (up to 200K tokens on some versions), reliable structured formatting, clean Markdown output without heavy prompt engineering.

Folder architecture: /Inbox (raw unprocessed notes — the critical piece; everything enters here first), /Notes (processed permanent notes), /People (one note per person), /Projects, /Concepts (evergreen ideas), /Briefings (AI-generated daily/weekly summaries), /Templates. Note frontmatter standard: title, date_created, date_modified, tags, status (inbox | processed | evergreen — the pipeline trigger), related, source. Enable the Obsidian Local REST API plugin (community plugin, generates an API key) so external scripts read/write the vault programmatically via HTTP — the vault becomes a queryable database.

Connect Claude to the vault: a Python script that scans /Inbox, sends each note to Claude with a structured processing prompt (claude-opus-4-5, max_tokens 4096), parses the JSON response, creates/updates the processed note in the right folder, and updates frontmatter/backlinks. The processing prompt returns a JSON object: summary (2-3 sentences), key_concepts (3-7), people, action_items, related_topics, destination_folder, and processed_content (rewritten clean Markdown with H2 heading, summary, bullet key points, and [[wiki-links]] — Obsidian's internal link syntax, which builds the knowledge graph automatically).

Automating ingestion: Web Clipper extension (Chrome/Firefox) saves web pages to /Inbox; Readwise Reader syncs highlights/articles on a schedule; Zapier/Make watch email labels/senders and POST to the vault via REST API; Otter.ai/Fireflies.ai transcripts auto-pushed to /Inbox (Claude is good at extracting decisions/action items from transcripts); a daily note template auto-appears each morning.

Cross-linking engine — the highest-value feature: (1) after processing a new note, extract key concepts, search the vault for existing notes mentioning those concepts but not linking to the new note, send them to Claude with a targeted "add a [[link]]" prompt, write updated content back — bidirectional linking produces a connected knowledge graph. (2) Concept hub notes: weekly job identifies concepts appearing in 5+ notes without a dedicated Concept note, collects all mentioning notes, asks Claude to write a comprehensive synthesis note with links to all sources, saves to /Concepts/[topic].md — the compounding effect where the vault gets smarter.

Daily briefings: surface notes added/updated in last 24h, action items, recurring concepts (emerging themes), connections between recent and older content, plus a "resurface" item. Saved to /Briefings/YYYY-MM-DD.md. Scheduling: macOS launchd/cron — inbox processing every 2h during working hours, backlink updates nightly, concept synthesis weekly (Sunday night), briefing daily at 7am; Windows Task Scheduler; hosted options (Railway, Render) via server-side cron.

Common mistakes: over-engineering the folder structure (5-6 top-level folders is enough); using Claude for every read operation (use Obsidian search/Dataview/REST API for lookups; Claude only for transformation); skipping the frontmatter standard; not reviewing briefings; letting the inbox pile up (add an alert when inbox count exceeds a threshold).

FAQ highlights: LLM Wiki is continuous and push-based vs single-session pull-based chat; data leaves the machine when sent to Claude's API (API inputs not used to train by default; swap in a local open-source model via Ollama if data residency matters); cross-linking/synthesis becomes meaningfully useful around 100-200 processed notes; works with any Markdown tool with an API (Logseq close alternative, Notion via API with friction, Roam more closed); MindStudio offers a no-code visual path.

## Key points

- LLM Wiki method: Claude acts as a continuous editor (ingest, parse, extract entities, link, synthesize, brief) inside an Obsidian vault, not a one-off summarizer.
- Folder structure (/Inbox → /Notes, /People, /Projects, /Concepts, /Briefings) plus consistent YAML frontmatter (status field as pipeline trigger) are prerequisites for reliable automation.
- The Obsidian Local REST API plugin turns the vault into a queryable database for external scripts.
- Highest-value automation is bidirectional cross-linking: Claude links new notes to existing content AND updates existing notes to reference new ones — building a knowledge graph.
- Concept hub notes (concepts appearing in 5+ notes) create the compounding effect — the vault gets smarter over time.
- Daily briefings (recency-based) turn passive storage into an active thinking tool, useful from day one.
- Model used in examples: claude-opus-4-5, max_tokens 4096.

## Technical data / figures

| Folder | Purpose |
|---|---|
| /Inbox | Raw, unprocessed notes/clips — pipeline entry point |
| /Notes | Processed permanent notes |
| /People | One note per person |
| /Projects | Active project workspaces |
| /Concepts | Evergreen idea notes / hub notes |
| /Briefings | AI-generated daily/weekly summaries |
| /Templates | Note templates |

| Pipeline component | Detail |
|---|---|
| Frontmatter fields | title, date_created, date_modified, tags, status, related, source |
| Local REST API | Obsidian plugin; base URL http://localhost:27123 with bearer API key |
| Processing model | claude-opus-4-5, max_tokens 4096 |
| Processing prompt output | summary, key_concepts, people, action_items, related_topics, destination_folder, processed_content (with [[wiki-links]]) |
| Ingestion sources | Web Clipper, Readwise Reader, Zapier/Make (email), Otter.ai/Fireflies.ai (meetings), daily note template |
| Concept hub trigger | Concept appearing in 5+ notes without dedicated note |
| Useful threshold | 100–200 processed notes for meaningful cross-linking |
| Schedule example | Inbox every 2h; backlinks nightly; concepts weekly; briefing 7am |

## Why this source matters for the RAG

Complete blueprint for building a self-maintaining Markdown knowledge base that can serve as a RAG corpus: defines the folder taxonomy, frontmatter schema, ingestion pipeline, cross-linking/graph construction, and briefing generation that structure retrieval-ready content. Directly applicable to creating an "AI second brain" RAG system with traceable, wiki-linked notes.

## Related context from the article

