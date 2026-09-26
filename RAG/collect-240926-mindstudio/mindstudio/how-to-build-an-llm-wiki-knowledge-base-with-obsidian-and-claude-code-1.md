---
id: collect-240926-mindstudio/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code-1
title: "Vault Schema"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agents", "attention", "claude", "fine-tuning", "research", "rlhf"]
source: docs/RAG/clean_en/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code.md
source_anchor: ""
source_lines: [1, 191]
sha256: 9d8791fc60e96a60e3cd02e129aaf262a12f09bf335cd8cde0a75033770b0926
---

# Vault Schema

<!-- source: https://www.mindstudio.ai/blog/how-to-build-llm-wiki-knowledge-base-obsidian-claude-code -->

## The Problem with How We Consume Information

Most people read an article, watch a tutorial, or skim a PDF — and then forget 80% of it within a week. The information disappears, and finding it again means re-reading the same source from scratch.

Building an LLM wiki knowledge base solves this. Instead of passive consumption, you process every piece of content you encounter into structured, searchable, interlinked notes — automatically. Andrej Karpathy has publicly described this kind of system as a core part of how he learns, and with tools like Claude Code and Obsidian, you can build one in an afternoon.

This guide shows you exactly how to do it: ingest YouTube transcripts, PDFs, and web pages, then use Claude Code to synthesize them into a self-growing wiki inside your Obsidian vault.

## What an LLM Wiki Knowledge Base Actually Is

The idea is simple. You have a vault of markdown notes (Obsidian is built for this). Every time you encounter useful content — a video, a research paper, a blog post — an LLM processes it and writes a structured note that fits into your existing knowledge graph.

The “wiki” part matters. Unlike a flat file system or a list of bookmarks, a wiki creates links between concepts. If you’ve already got a note on “attention mechanisms” and you process a new YouTube video on transformers, the new note should reference your existing one. Over time, you get a network of interconnected ideas, not a pile of documents.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

The “LLM” part is what makes it scalable. Manually summarizing and linking notes is slow. Claude Code can do it in seconds — reading your existing vault, understanding what’s already there, and writing new notes that slot in coherently.

Karpathy’s vision for this kind of system emphasizes that the knowledge base should grow with minimal friction. The harder it is to add something, the less often you’ll do it. The system should make ingestion almost effortless.

## Architecture Overview

Before writing a single line of code, it helps to understand the three layers of this system.

### Layer 1: The Vault (Obsidian)

Obsidian stores everything as local markdown files. This matters for two reasons: Claude Code can read and write them directly without any API integration, and you own your data completely. The vault becomes your “database” of structured notes.

A well-organized vault for an LLM wiki typically uses:

- A `Sources/` folder for raw processed content
- A `Concepts/` folder for synthesized ideas
- A `MOCs/` folder (Maps of Content) for index notes that link related topics
- A consistent frontmatter schema (title, tags, date, source URL)

### Layer 2: The Ingestion Pipeline

This is where Claude Code lives. You point it at a source — a YouTube URL, a PDF file, a web page — and it handles:

1. Extracting the raw text
2. Reading your existing vault to understand context
3. Writing a structured note with links to related concepts
4. Updating relevant MOC notes

### Layer 3: The Growth Loop

The real power comes from iteration. Each new note makes the next one better, because Claude has more context to work with. A note on “RLHF” becomes richer once there are already notes on “reward modeling” and “fine-tuning” linking into it.

## Prerequisites

You don’t need to be a developer to follow this, but you should be comfortable running commands in a terminal.

**What you’ll need:**

- Obsidian installed with a vault set up
- Claude Code installed (`npm install -g @anthropic-ai/claude-code` )
- Python 3.8+ (for the transcript/PDF extraction scripts)
- `yt-dlp` for YouTube transcript extraction (`pip install yt-dlp` )
- `PyMuPDF` for PDF processing (`pip install pymupdf` )
- An Anthropic API key configured for Claude Code

**Optional but useful:**

- The Obsidian Dataview plugin for querying your vault
- The Templater plugin for consistent note formatting

## Step 1: Set Up Your Obsidian Vault Structure

Create a new vault or use an existing one. Inside it, create these folders:

```
/YourVault
  /Sources
    /YouTube
    /PDFs
    /URLs
  /Concepts
  /MOCs
  /Inbox
```
The `Inbox/` folder is where Claude will drop new notes before they’re fully integrated. This gives you a review step before anything enters your main knowledge base — useful when you’re starting out and want to check the output quality.

Next, create a `_schema.md` file at the root of your vault. This is critical. Claude Code will read this file to understand how your notes should be formatted.

Here’s a minimal schema file:

```
# Vault Schema
## Note Types
### Source Note
- Frontmatter: title, source_url, date_processed, tags, type: "source"
- Sections: Summary, Key Ideas (bullet list), Quotes, Related Concepts (wikilinks)
### Concept Note  
- Frontmatter: title, tags, type: "concept"
- Sections: Definition, Context, Examples, Related Concepts (wikilinks)
### MOC Note
- Frontmatter: title, tags, type: "moc"
- Sections: Overview, Notes in this cluster (wikilinks)
## Linking Rules
- Always use [[wikilink]] syntax to reference other notes
- Prefer linking to Concept notes over Source notes
- If a concept note doesn't exist yet, create it as a stub
```
## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

This file acts as a style guide for Claude. The more specific you are here, the more consistent your output will be.

## Step 2: Write the Extraction Scripts

You need three small scripts to pull raw text from each source type. These aren’t meant to be production-grade — they just need to work.

### YouTube Transcript Extractor

```
# extract_youtube.py
import sys
import subprocess
import json
def get_transcript(url):
    result = subprocess.run(
        ['yt-dlp', '--write-auto-sub', '--skip-download', 
         '--sub-format', 'json3', '-o', '/tmp/transcript', url],
        capture_output=True, text=True
    )
    # Parse the downloaded subtitle file
    # Return clean transcript text
    ...
if __name__ == "__main__":
    url = sys.argv[1]
    transcript = get_transcript(url)
    print(transcript)
```
For production use, `youtube-transcript-api` is cleaner (`pip install youtube-transcript-api`). It handles most public videos without needing yt-dlp:

```
from youtube_transcript_api import YouTubeTranscriptApi
import re
def get_video_id(url):
    match = re.search(r'(?:v=|youtu\.be/)([^&\n?#]+)', url)
    return match.group(1) if match else None
def fetch_transcript(url):
    video_id = get_video_id(url)
    transcript = YouTubeTranscriptApi.get_transcript(video_id)
    return ' '.join([entry['text'] for entry in transcript])
```
### PDF Extractor

```
# extract_pdf.py
import sys
import fitz  # PyMuPDF
def extract_pdf(filepath):
    doc = fitz.open(filepath)
    text = ''
    for page in doc:
        text += page.get_text()
    return text
if __name__ == "__main__":
    print(extract_pdf(sys.argv[1]))
```
### URL Extractor

For web pages, use `requests` and `BeautifulSoup`:

```
# extract_url.py
import sys
import requests
from bs4 import BeautifulSoup
def fetch_page(url):
    response = requests.get(url, timeout=10)
    soup = BeautifulSoup(response.text, 'html.parser')
    # Remove nav, footer, scripts
    for tag in soup(['nav', 'footer', 'script', 'style']):
        tag.decompose()
    return soup.get_text(separator='\n', strip=True)
if __name__ == "__main__":
    print(fetch_page(sys.argv[1]))
```
## Step 3: Write the Claude Code Prompt

This is the most important part. The quality of your knowledge base depends heavily on the instructions you give Claude.

Create a file called `ingest_prompt.md` in your vault root:

