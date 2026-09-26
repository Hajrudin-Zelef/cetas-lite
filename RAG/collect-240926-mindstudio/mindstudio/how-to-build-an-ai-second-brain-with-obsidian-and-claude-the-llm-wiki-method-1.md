---
id: collect-240926-mindstudio/mindstudio/how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method-1
title: "how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2025-01-15"]
keywords: ["claude", "research"]
source: docs/RAG/clean_en/mindstudio/how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method.md
source_anchor: ""
source_lines: [1, 163]
sha256: aead7306e17686f4aecd8910adbb84f6676dac9b6f700ab60451985bf165a251
---

# how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method

<!-- source: https://www.mindstudio.ai/blog/ai-second-brain-obsidian-claude-llm-wiki -->

## Why Your Notes App Is Failing You

Most knowledge management systems share the same flaw: they’re good at storing information and terrible at connecting it. You clip an article, jot down a meeting insight, paste a research summary — and three weeks later, none of it surfaces when you actually need it.

The promise of a “second brain” has been around since Tiago Forte popularized the concept. But the original vision was still largely manual. You, the human, were responsible for tagging, linking, and synthesizing. That’s cognitively expensive work that most people abandon within a month.

What’s changed is Claude. Combined with Obsidian’s local-first, Markdown-based vault structure, you can now build an AI second brain that actually maintains itself — ingesting new information, cross-linking related ideas, and generating daily briefings without you manually curating every note. This is what practitioners are calling the **LLM Wiki method**: using a large language model as the active knowledge worker inside your personal knowledge base.

This guide walks through exactly how to build that system.

## What the LLM Wiki Method Actually Means

The term “LLM Wiki” describes a knowledge vault where an LLM — in this case, Claude — acts as a continuous editor, not just a one-off summarizer.

Traditional wikis are written and maintained by humans. An LLM Wiki flips that model: you feed raw information (notes, transcripts, articles, emails, meeting recordings), and the LLM handles:

- **Parsing and summarizing** incoming content
- **Extracting entities** (people, concepts, projects, dates)
- **Creating bidirectional links** between related notes
- **Writing synthesis notes** that connect ideas across different sources
- **Generating daily or weekly briefings** from your vault’s recent activity

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

Obsidian is the ideal host for this system because its vault is just a folder of Markdown files on your local machine. There’s no proprietary database to wrestle with — every note is a plain text file that any script or API call can read and write.

Claude is the ideal LLM for this because it handles long context windows well (up to 200K tokens on some versions), follows structured formatting instructions reliably, and produces clean Markdown output without heavy prompt engineering.

## Set Up Your Obsidian Vault for Automation

Before any AI does anything useful, your vault needs a structure that makes automation predictable.

### Folder Architecture

A functional LLM Wiki vault uses a small number of top-level folders:

```
/Inbox          — Raw, unprocessed notes and clips
/Notes          — Processed permanent notes
/People         — One note per person (contacts, colleagues, thinkers)
/Projects       — Active project workspaces
/Concepts       — Evergreen idea notes
/Briefings      — AI-generated daily/weekly summaries
/Templates      — Note templates for consistency
```
The `/Inbox` folder is the critical piece. Everything enters here first. Your automation pipeline watches this folder, processes each file, and moves it to the right location. This separation keeps raw input away from your curated knowledge.

### Note Frontmatter Standards

Consistent YAML frontmatter is what makes AI processing reliable. Every note should include:

```
---
title: "Note Title"
date_created: 2025-01-15
date_modified: 2025-01-15
tags: []
status: inbox  # inbox | processed | evergreen
related: []
source: ""
---
```
The `status` field is your pipeline trigger. Notes with `status: inbox` are waiting to be processed. The AI changes this to `processed` after it’s done, and you can manually promote notes to `evergreen` when they’ve matured.

### Enable the Obsidian Local REST API

The Obsidian Local REST API plugin is what lets external scripts read and write your vault programmatically. Install it from the Obsidian community plugins directory, then enable it and note the API key it generates.

With this running, your vault becomes a queryable database. You can list all notes in a folder, read note content, create new notes, and update existing ones — all via simple HTTP requests from any script.

## Connect Claude to Your Vault

With your vault structure in place, you need a processing layer that sits between Obsidian and Claude’s API.

### The Core Processing Script

The backbone of your LLM Wiki is a script that:

1. Scans `/Inbox` for unprocessed notes
2. Sends each note to Claude with a structured processing prompt
3. Parses Claude’s response
4. Creates or updates the processed note in the appropriate folder
5. Updates frontmatter and backlinks

Here’s the basic Python structure:

```
import anthropic
import requests
import json
OBSIDIAN_BASE_URL = "http://localhost:27123"
OBSIDIAN_API_KEY = "your-api-key"
client = anthropic.Anthropic(api_key="your-claude-api-key")
def get_inbox_notes():
    response = requests.get(
        f"{OBSIDIAN_BASE_URL}/vault/Inbox/",
        headers={"Authorization": f"Bearer {OBSIDIAN_API_KEY}"}
    )
    return response.json()["files"]
def process_note(note_path):
    # Read the raw note
    response = requests.get(
        f"{OBSIDIAN_BASE_URL}/vault/{note_path}",
        headers={"Authorization": f"Bearer {OBSIDIAN_API_KEY}"}
    )
    raw_content = response.text
    
    # Send to Claude for processing
    message = client.messages.create(
        model="claude-opus-4-5",
        max_tokens=4096,
        messages=[{
            "role": "user",
            "content": PROCESSING_PROMPT.format(content=raw_content)
        }]
    )
    
    return json.loads(message.content[0].text)
```
### The Processing Prompt

The quality of your LLM Wiki depends heavily on this prompt. Here’s a robust starting point:

```
You are a knowledge management assistant. Process the following raw note and return a JSON object with these fields:
- "summary": A 2-3 sentence summary of the main idea
- "key_concepts": Array of 3-7 core concepts mentioned
- "people": Array of any people referenced
- "action_items": Array of any tasks or follow-ups mentioned
- "related_topics": Array of topics this note connects to
- "destination_folder": Where this note should live (Notes/People/Projects/Concepts)
- "processed_content": The full note rewritten in clean Markdown with:
  - A clear H2 heading
  - Summary section
  - Key points as bullets
  - [[wiki-links]] added for any concepts, people, or projects that should have their own notes
  
Raw note:
{content}
Return only valid JSON. No additional text.
```
### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

The `[[wiki-links]]` instruction is crucial — this is Obsidian’s internal link syntax, and having Claude insert these automatically is what builds your knowledge graph without manual effort.

## Automate Data Ingestion from Multiple Sources

A second brain that only accepts manually created notes isn’t much of an upgrade. The LLM Wiki method becomes genuinely powerful when you pipe external sources directly into your `/Inbox`.

### Web Clips via Browser Extension

Obsidian’s Web Clipper extension (available for Chrome and Firefox) can save any web page directly to your vault. Configure it to always drop clips into `/Inbox` with `status: inbox` in the frontmatter. From there, your processing script picks them up automatically.

For heavier research workflows, tools like Readwise Reader can sync highlights and articles to your vault on a schedule, giving you a continuous feed of curated reading into your `/Inbox`.

### Email and Meeting Notes

