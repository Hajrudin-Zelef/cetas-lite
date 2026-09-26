---
id: collect-240926-mindstudio/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram-1
title: "Install Ollama (macOS/Linux)"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "agent", "agents", "consumer", "context window", "cost", "gpus", "inference", "memory", "mistral", "multimodal", "parameters"]
source: docs/RAG/clean_en/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram.md
source_anchor: ""
source_lines: [1, 142]
sha256: 1b7785eb46a61dde97e24947d78cd7f38f1eaabbbfaf117f6343327071d4501a
---

# Install Ollama (macOS/Linux)

<!-- source: https://www.mindstudio.ai/blog/add-vision-to-local-ai-agent-small-model -->

## The VRAM Wall That Stops Most Local AI Setups

Running a local LLM is increasingly practical. Models like Mistral, Llama 3, and Phi-3 run well on consumer hardware. But the moment you need vision — reading a screenshot, parsing a scanned PDF, extracting data from an image — most setups hit a wall.

Naive solutions are expensive. Loading a full multimodal model like LLaVA-34B or a large GPT-4V equivalent locally can demand 24–40GB of VRAM. That's not a consumer setup. That's a workstation or a server.

But there's a smarter approach: instead of cramming everything into one giant multimodal model, you separate vision from reasoning. A lightweight vision model handles image interpretation. Your existing text LLM handles reasoning and response generation. The result is a capable local AI agent that adds vision capabilities without blowing your VRAM budget.

This guide explains how to build that architecture — from choosing a small vision model to wiring the pipeline together for real-world tasks like reading screenshots and processing PDFs.

## Why Keeping Vision Local Actually Matters

Before getting into the how, it's worth being clear about why you'd bother running vision locally at all.

**Privacy is usually the real driver.** If you're processing screenshots of internal dashboards, scanned contracts, medical records, or any document containing sensitive data, sending those images to a cloud API creates a data exposure risk. Even if the API provider claims not to store inputs, you're still transmitting the data.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

**Cost adds up fast.** Cloud vision APIs charge per image or per token. At scale — hundreds or thousands of documents per day — the bill grows quickly. A local setup has upfront hardware costs but near-zero marginal cost per inference.

**Offline reliability matters for some use cases.** If your agent runs in an air-gapped environment, a factory floor with spotty connectivity, or on a laptop while traveling, cloud dependency is a non-starter.

The challenge is that most consumer GPUs have 8–16GB of VRAM. That's tight when you're already running a capable text model. Fitting both a text LLM and a vision model into that budget requires choosing models carefully.

## How the Dual-Model Architecture Works

The core insight is that vision and reasoning don't need to happen in the same model.

Most multimodal models are actually doing something conceptually simple: they convert an image into a rich text description or structured extraction, then use language modeling to answer questions about it. You can replicate that pipeline with two separate, smaller models instead of one large one.

Here's the flow:

1. **User provides an image** (screenshot, PDF page, photo, etc.)
2. **Vision model receives the image + a prompt** asking it to describe, extract, or interpret specific content
3. **Vision model outputs text** — a description, extracted table data, transcribed text, or structured JSON
4. **Text LLM receives that output** as context, along with the original user query
5. **Text LLM generates the final response** using its reasoning capabilities

The vision model acts as a translator. It converts visual information into language. Your text LLM never touches pixels — it only sees the linguistic output.

This split has real advantages:

- Each model is optimized for its task
- You can swap out either model independently
- Total VRAM usage is much lower than a single large multimodal model
- You can run the vision model on CPU if needed, since it only runs for a few seconds per image

## Choosing a Lightweight Vision Model

Not all vision models are created equal, and the small-model space has improved significantly. Here are the most practical options as of 2025.

### MiniCPM-V

MiniCPM-V from OpenBMB is one of the most capable small vision models available. The 2B variant runs comfortably on 4–6GB of VRAM. The 8B variant (MiniCPM-V 2.6) needs more headroom — around 10–16GB depending on quantization — but handles complex documents, charts, and multi-page images well.

Key strengths:

- Strong OCR performance (reading text from screenshots and documents)
- Handles multiple images in a single context window
- Good at structured extraction tasks
- Available through Ollama as `minicpm-v`

For most screenshot and PDF use cases, the 2B model is sufficient and leaves plenty of VRAM for your text LLM.

### Moondream

Moondream2 is a 1.8B parameter model specifically designed for efficiency. It loads fast, uses minimal memory (under 4GB VRAM), and handles basic image description and question-answering well.

It's not the best at dense document parsing, but for tasks like "what's in this screenshot" or "describe the chart" it performs well above its size. Available via Ollama as `moondream`.

### LLaVA Variants

LLaVA (Large Language and Vision Assistant) has several community-maintained variants. LLaVA-Phi-3-Mini combines Microsoft's Phi-3 Mini with vision capability in a package that fits in 4–8GB of VRAM.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

LLaVA models tend to be more verbose in their descriptions, which can be helpful or noisy depending on your use case. They're widely supported across inference frameworks.

### Qwen2-VL (2B)

Alibaba's Qwen2-VL at 2B parameters is worth mentioning. It handles multi-language documents well and has strong performance on charts and infographics. Available through Ollama and compatible with most inference setups.

### A Practical Comparison

| Model | Size | VRAM (approx.) | Best For |
|---|---|---|---|
| MiniCPM-V 2B | 2B | 4–6GB | OCR, screenshots, documents |
| Moondream2 | 1.8B | 3–4GB | General image Q&A |
| LLaVA-Phi-3-Mini | 3.8B | 5–8GB | Detailed descriptions |
| Qwen2-VL 2B | 2B | 4–6GB | Multilingual docs, charts |
| MiniCPM-V 2.6 | 8B | 10–16GB | Complex documents, multi-image |

If you’re running on 8GB of VRAM total, target models in the 3–4GB range to leave room for your text LLM. If you have 16GB, you have more flexibility.

## Setting Up the Pipeline with Ollama

Ollama is the easiest way to run both your vision model and text model locally. It handles model management, provides a consistent API, and supports hot-swapping between models.

### Step 1: Install Ollama and Pull Your Models

```
# Install Ollama (macOS/Linux)
curl -fsSL https://ollama.com/install.sh | sh
# Pull a vision model
ollama pull minicpm-v
# Pull your text LLM (example: Mistral)
ollama pull mistral
```
For Windows, download the Ollama installer from the official site.

### Step 2: Test the Vision Model

Ollama’s API accepts base64-encoded images directly. Here’s a quick test with curl:

```
curl http://localhost:11434/api/generate -d '{
  "model": "minicpm-v",
  "prompt": "Extract all text visible in this image. Return it as plain text, preserving structure.",
  "images": ["'$(base64 -i screenshot.png)'"]
}'
```
If you see text from your screenshot in the response, the vision model is working correctly.

### Step 3: Build the Orchestration Layer

The orchestration layer is what makes this a proper agent. It needs to:

1. Accept an image + user query as input
2. Send the image to the vision model with an appropriate extraction prompt
3. Capture the vision model’s text output
4. Construct a new prompt combining that output with the user’s original query
5. Send the combined prompt to the text LLM
6. Return the final response

Here’s a minimal Python implementation:

