---
id: collect-240926-mindstudio/mindstudio/how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method
title: "how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Meta", "Mistral"]
dates: ["2025-01-15"]
keywords: ["claude", "agent", "agents", "llama", "mistral", "research"]
source: docs/RAG/clean_en/mindstudio/how-to-build-an-ai-second-brain-with-obsidian-and-claude-the-llm-wiki-method.md
source_anchor: ""
source_lines: [1, 312]
sha256: db660517c8a2e5ef892865c0476eddad948906155a38af2f904a09922d50cffd
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

Connect your email to the pipeline using a service like Zapier or Make, which can watch for specific labels or senders and POST note content to your Obsidian vault via the REST API.

For meetings, tools like Otter.ai or Fireflies.ai produce transcripts that you can automatically push to `/Inbox`. Claude is particularly good at processing meeting transcripts — extracting decisions, action items, and context that you’d otherwise have to summarize manually.

### Daily Capture Automation

Build a simple daily note template that automatically appears in `/Inbox` each morning. Your morning capture (tasks, ideas, news items) goes here, and the evening processing run cleans it up, extracts key concepts, and links it to relevant existing notes.

## Build the Cross-Linking Engine

The most valuable feature of the LLM Wiki method isn’t note creation — it’s automatic cross-linking. This is what turns a collection of notes into a knowledge graph.

### Concept Extraction and Backlink Creation

When Claude processes a note, it identifies key concepts and adds `[[wiki-links]]` inline. But there’s a second step that most implementations miss: updating existing notes to link back to the new note.

This requires Claude to scan your existing vault for related content. Here’s the approach:

1. After processing a new note, extract its key concepts
2. Search your vault for existing notes that mention those concepts but don’t yet link to the new note
3. Send those existing notes to Claude with a targeted prompt: “Add a [[link]] to [[New Note Title]] wherever relevant in this note. Return only the updated note content.”
4. Write the updated content back to those files

This bidirectional linking happens automatically, and over time produces a genuinely connected knowledge graph that Obsidian’s graph view makes visually useful.

### Building Concept Hub Notes

As concepts appear repeatedly across multiple notes, Claude can generate “hub notes” — Concept notes that synthesize everything your vault knows about a topic.

Set up a weekly job that:

1. Identifies concepts that appear in 5+ notes but don’t have their own dedicated Concept note yet
2. Collects all notes mentioning that concept
3. Sends them to Claude with a synthesis prompt: “Based on these notes, write a comprehensive Concept note about [topic] that synthesizes the key ideas and links to all source notes.”
4. Saves the result to `/Concepts/[topic].md`

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

This is the LLM Wiki method’s real compounding effect. Your vault doesn’t just grow — it gets smarter as Claude synthesizes patterns across your accumulated knowledge.

## Generate Daily Briefings from Your Vault

The daily briefing is what makes this system genuinely useful day-to-day rather than just impressive in theory.

### What Goes Into a Good Briefing

A useful briefing should surface:

- Notes added or updated in the last 24 hours
- Any action items extracted from recent notes
- Concepts that appeared repeatedly this week (emerging themes)
- Connections Claude noticed between recent notes and older content
- A “resurface” item: one older note worth revisiting based on current context

### The Briefing Prompt

```
You have access to the following recent activity in a personal knowledge vault:
NEW NOTES (last 24 hours):
{new_notes_content}
UPDATED NOTES (last 24 hours):
{updated_notes_content}
RECURRING CONCEPTS THIS WEEK:
{top_concepts}
Generate a daily briefing in Markdown format:
1. 3-5 bullet summary of yesterday's key inputs
2. Open action items extracted from recent notes
3. One emerging theme you notice across recent notes
4. One connection between a recent note and an older note in the vault
5. One older note to resurface today (pick the most relevant to current themes)
Keep the tone direct and factual. No filler.
```
The briefing note gets saved to `/Briefings/YYYY-MM-DD.md` automatically. Over time, these briefings themselves become a searchable record of your intellectual activity.

### Scheduling the Pipeline
```

On macOS, a launchd plist or cron job can run your processing script on whatever schedule makes sense:

- **Inbox processing** : Every 2 hours during working hours
- **Backlink updates** : Nightly
- **Concept synthesis** : Weekly (Sunday night)
- **Daily briefing** : Each morning at 7am

On Windows, Task Scheduler does the same job. If you want a hosted option that doesn’t depend on your local machine being on, a simple server-side cron via Railway, Render, or similar services works well.

## Common Mistakes to Avoid

### Over-Engineering the Folder Structure

Five or six top-level folders is enough. People who create deeply nested taxonomies before their vault has 50 notes are building for a system that doesn’t exist yet. Start simple.

### Using Claude for Every Read Operation

Claude should process incoming notes and generate synthesis. It shouldn’t be in the loop for simple search or retrieval. Use Obsidian’s built-in search, Dataview plugin queries, or the REST API for lookups — save Claude calls for actual transformation work.

### Skipping the Frontmatter Standard

Inconsistent frontmatter breaks automation. Define your schema once, create a template, and enforce it from day one. Every note that deviates requires manual cleanup later.

### Not Reviewing Briefings

The daily briefing only creates value if you read it. Block five minutes each morning. If you consistently skip it, make the briefing shorter or change what it surfaces — don’t abandon the system.

### Letting the Inbox Pile Up

If processing runs fail and you don’t notice for a week, your inbox floods. Add a simple alert (email or Slack notification) when inbox note count exceeds a threshold. This is a five-minute addition to your pipeline and worth it.

## Frequently Asked Questions

### What is an AI second brain?

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

An AI second brain is a personal knowledge management system where an AI model actively organizes, connects, and synthesizes your notes — rather than you doing that work manually. Unlike traditional note-taking apps, the AI continuously processes incoming information, creates links between related ideas, and surfaces relevant content when you need it.

### Do I need coding experience to build this?

Basic Python familiarity helps for the custom scripting approach described here. However, platforms like MindStudio let you configure the same pipeline visually without code, which makes this accessible to non-developers. The key prerequisites are comfort with API keys and understanding how Obsidian’s plugin ecosystem works.

### How is this different from just chatting with Claude about my notes?

Chatting with Claude is a single-session, pull-based interaction. The LLM Wiki method is continuous and push-based — Claude automatically processes every note that enters your vault, maintains cross-links, and generates briefings on a schedule. You’re not asking Claude questions; Claude is actively maintaining your knowledge base as a background process.

### Is my data private if Claude is processing my notes?

Your notes leave your local machine when they’re sent to Claude’s API. Anthropic’s usage policies describe how API data is handled — API inputs are not used to train models by default. If data residency is a concern, you can run an open-source model locally (via Ollama with LLaMA or Mistral) and swap it into the same pipeline architecture described here. Performance will be lower, but your data never leaves your machine.

### How many notes does a vault need before this becomes useful?

The cross-linking and synthesis features become meaningfully useful around 100–200 processed notes. Before that threshold, Claude can still process and organize notes, but the graph connections are sparse. The daily briefing is useful from day one since it’s based on recency, not vault size.

### Can this work with tools other than Obsidian?

The LLM Wiki method works with any Markdown-based tool that exposes an API or accepts file system writes. Logseq (which uses a similar local Markdown format) is a close alternative. Notion works if you use their API, though the lack of local file access adds friction. Roam Research has a more closed architecture that makes automation harder. Obsidian is the path of least resistance for this approach.

## Key Takeaways

- The LLM Wiki method uses Claude as an active knowledge worker inside your Obsidian vault — not a one-off summarizer, but a continuous processor that ingests, links, and synthesizes.
- A clean folder structure and consistent frontmatter schema are prerequisites that make everything else work reliably.
- The highest-value automation is bidirectional cross-linking: Claude not only links new notes to existing content, but updates existing notes to reference new ones.
- Daily briefings generated from vault activity turn passive note storage into an active thinking tool.
- For teams or non-developers, MindStudio provides the infrastructure to run this pipeline without custom scripting — building the same Claude-powered automation through a visual workflow builder.

If you want to build the AI knowledge system described here without wrestling with infrastructure, MindStudio is worth exploring. The visual agent builder handles the Claude API connections, scheduling, and multi-step logic that make this pipeline work reliably over time.
