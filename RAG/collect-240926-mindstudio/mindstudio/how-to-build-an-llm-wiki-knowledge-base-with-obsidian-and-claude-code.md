---
id: collect-240926-mindstudio/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code
title: "Vault Schema"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agents", "attention", "claude", "context window", "cost", "fine-tuning", "llama", "research", "rlhf"]
source: docs/RAG/clean_en/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code.md
source_anchor: ""
source_lines: [1, 392]
sha256: db9b455241db17580fc37777c37afb52f2fd50f3728d5c1e4d6c6c3e0f0d0ca5
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

```
You are a knowledge base curator. Your job is to process new content 
and add it to the vault following the schema in _schema.md.
## Your Task
1. Read _schema.md to understand note formatting rules
2. Scan the Concepts/ folder to understand what's already in the vault
3. Read the provided source text
4. Create a Source note in Sources/[type]/ with proper frontmatter
5. Identify 3-5 key concepts from the source
6. For each concept:
   - Check if a Concept note already exists
   - If yes: add a link to the new Source note in the Related section
   - If no: create a stub Concept note
7. Update or create the relevant MOC note
8. Output a summary of what you created/modified
## Source Information
- Type: {type}
- URL/Path: {source}
- Raw Text: {content}
## Important Rules
- Never modify existing note content, only append
- Always use [[wikilink]] syntax
- Keep Source note summaries under 300 words
- Concept notes should be timeless — no dates in the content
```
Now you can run Claude Code with this prompt:

```
claude-code --prompt ingest_prompt.md \
  --var type=youtube \
  --var source="https://youtube.com/watch?v=..." \
  --var content="$(python extract_youtube.py URL)"
```
## Step 4: Automate the Ingestion Workflow

Running three commands every time you want to add something is friction. Wrap it in a single shell script:

```
#!/bin/bash
# ingest.sh
VAULT="/path/to/your/vault"
TYPE=$1
SOURCE=$2
case $TYPE in
  youtube)
    CONTENT=$(python "$VAULT/scripts/extract_youtube.py" "$SOURCE")
    ;;
  pdf)
    CONTENT=$(python "$VAULT/scripts/extract_pdf.py" "$SOURCE")
    ;;
  url)
    CONTENT=$(python "$VAULT/scripts/extract_url.py" "$SOURCE")
    ;;
  *)
    echo "Unknown type. Use: youtube, pdf, or url"
    exit 1
    ;;
esac
claude-code \
  --cwd "$VAULT" \
  --prompt "$VAULT/ingest_prompt.md" \
  --var type="$TYPE" \
  --var source="$SOURCE" \
  --var content="$CONTENT"
```
## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

Now adding something to your knowledge base looks like:

```
./ingest.sh youtube "https://youtube.com/watch?v=dQw4w9WgXcQ"
./ingest.sh pdf "/Downloads/attention-is-all-you-need.pdf"
./ingest.sh url "https://example.com/interesting-article"
```
That’s the entire user-facing workflow.

## Step 5: Improve Note Quality Over Time

The first pass of notes won’t be perfect. Here’s how to systematically improve them.

### Add a Review Step

Route all new notes to `Inbox/` first. Set Claude’s output path there, and do a quick review before moving notes to their permanent home. After a week, you’ll notice patterns in what needs fixing and can update your `ingest_prompt.md` accordingly.

### Teach Claude Your Vocabulary

Add a `glossary.md` file to your vault that lists terms you use in a specific way. If you call something “sparse attention” but the source calls it “local attention,” Claude will create two separate concepts instead of linking them. The glossary prevents this.

### Run Periodic Synthesis

Every few weeks, run a synthesis prompt that asks Claude to look across your Concepts folder and identify:

- Concepts that should be merged
- Missing links between related notes
- Concepts that have grown enough to warrant their own MOC

```
claude-code \
  --cwd "$VAULT" \
  --prompt "Review all notes in Concepts/ and suggest merges, 
            new links, and MOC candidates. Output a report to 
            Inbox/synthesis-report-$(date +%Y%m%d).md"
```
## Step 6: Handle Edge Cases

### Videos Without Transcripts

Some YouTube videos have auto-captions disabled. For these, use `yt-dlp` to download the audio and process it with a local Whisper model:

```
yt-dlp -x --audio-format wav -o /tmp/audio.wav "$URL"
whisper /tmp/audio.wav --output_format txt
```
Then pipe the text output into your normal ingestion flow.

### Large PDFs

Academic papers and books can exceed Claude’s context window. Split them first:

```

```
def chunk_text(text, max_chars=50000):
    chunks = []
    while len(text) > max_chars:
        split_point = text.rfind('\n', 0, max_chars)
        chunks.append(text[:split_point])
        text = text[split_point:]
    chunks.append(text)
    return chunks
```
Process each chunk separately, then run a consolidation step that merges the resulting notes.

### Duplicate Detection

Before creating a new note, Claude should check for duplicates. Add this to your prompt:

```
Before creating any note, check if a similar note already exists 
by searching for the source URL in existing frontmatter. 
If found, update the existing note instead of creating a new one.
```
## Common Mistakes to Avoid

### Skipping the Schema File

Without clear formatting instructions, Claude’s output will be inconsistent. Notes will have different section names, link styles, and frontmatter fields. The schema file is the single most important thing to get right early.

### Processing Too Much Too Fast

Resist the urge to dump 200 PDFs into the system on day one. Start with 10–20 sources. Let the vault develop its initial structure, review the output, and refine your prompts before scaling up.

### Not Linking Back to Sources

Source notes without proper URLs become orphaned knowledge — you won’t know where an idea came from. Always include the original URL or file path in frontmatter.

### Ignoring Concept Sprawl

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

If you’re not careful, you’ll end up with 500 one-line stub notes that never develop. Run synthesis passes regularly, and set a rule: any concept note under 100 words that hasn’t been updated in 30 days should be merged into a related note or deleted.

## Frequently Asked Questions

### Does this work with private or paywalled content?

For YouTube, it works with any video that has captions (auto or manual). Paywalled articles are trickier — you’d need to copy-paste the text manually and pipe it into your script as a text file. PDFs from paywalled journals work fine once you have the file locally.

### How much does Claude Code cost to run?

Claude Code uses Anthropic’s API, billed per token. A typical YouTube video ingestion (60-minute video, ~10,000 words of transcript) costs roughly $0.10–$0.30 depending on the model you choose. Processing a 20-page PDF is usually under $0.10. For personal use, the monthly cost is typically $5–$20 depending on how much you ingest.

### Can I use a different LLM instead of Claude?

Yes. The architecture is model-agnostic — the extraction scripts and Obsidian structure don’t care which model you use. Claude Code is convenient because it has strong file system awareness and handles long-context documents well, but you could adapt the prompts for GPT-4o or local models via Ollama. For local-first setups, Llama 3 70B handles this task reasonably well.

### How do I handle content in languages other than English?

Claude handles multilingual content well. You can either process notes in their original language or ask Claude to translate and process them into English. Add a `language` field to your frontmatter schema and a note in your prompt about how to handle non-English sources.

### Will this work with audio or video files I record myself?

Yes. Use Whisper to transcribe any audio or video file locally, then pipe the transcript text into your normal ingestion flow. The rest of the pipeline doesn’t care where the text came from.

### How does this compare to tools like Mem or Notion AI?

Tools like Mem and Notion AI are hosted services that handle ingestion and retrieval for you. This approach is local-first, free beyond API costs, and fully customizable. The tradeoff is setup time. If you want full control over your data and the structure of your knowledge base, the DIY approach wins. If you want something that works in five minutes, a hosted tool is easier.

## Key Takeaways

- An LLM wiki knowledge base pairs Obsidian’s local markdown vault with Claude Code’s file system awareness to build a self-growing network of interconnected notes.
- The schema file is the most important piece — it’s what keeps Claude’s output consistent across hundreds of notes.
- Three extraction scripts (YouTube, PDF, URL) plus a single shell script reduce the ingestion workflow to one command.
- Periodic synthesis runs keep the vault clean and prevent concept sprawl.
- For teams, MindStudio can add a no-code trigger layer on top of the same pipeline, connecting it to Slack, Notion, or any other tool where links get shared.

The system compounds. After a month of consistent ingestion, you’ll have a knowledge base that surfaces connections you never would have noticed manually — and adding new content only takes seconds.
