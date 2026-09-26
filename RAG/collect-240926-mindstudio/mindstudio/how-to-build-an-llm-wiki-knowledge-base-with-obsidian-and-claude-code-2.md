---
id: collect-240926-mindstudio/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code-2
title: "Vault Schema"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agents", "attention", "claude", "context window", "cost", "llama"]
source: docs/RAG/clean_en/mindstudio/how-to-build-an-llm-wiki-knowledge-base-with-obsidian-and-claude-code.md
source_anchor: ""
source_lines: [192, 383]
sha256: bb17b553d347a165aad9abc5614ea283e23f2f09c2bd28fbbf29ea7fabcff380
---

# Vault Schema

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

