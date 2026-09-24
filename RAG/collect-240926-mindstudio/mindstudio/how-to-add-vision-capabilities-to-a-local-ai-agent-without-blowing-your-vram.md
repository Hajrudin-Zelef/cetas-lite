---
id: collect-240926-mindstudio/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram
title: "Install Ollama (macOS/Linux)"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "agent", "agents", "claude", "consumer", "context window", "cost", "decode", "gemini", "gpu", "gpus", "inference"]
source: docs/RAG/clean_en/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram.md
source_anchor: ""
source_lines: [1, 315]
sha256: edd1fe0fc2e4bfc38d8f55dda08dd7db61183ab6bb470ea0d958fb2d7cd0a5fb
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

```
import requests
import base64
import json
OLLAMA_URL = "http://localhost:11434/api/generate"
def encode_image(image_path):
    with open(image_path, "rb") as f:
        return base64.b64encode(f.read()).decode("utf-8")
def vision_extract(image_path, extraction_prompt):
    payload = {
        "model": "minicpm-v",
        "prompt": extraction_prompt,
        "images": [encode_image(image_path)],
        "stream": False
    }
    response = requests.post(OLLAMA_URL, json=payload)
    return response.json()["response"]
def text_reason(context, user_query):
    prompt = f"""You have been given the following content extracted from an image:
---
{context}
---
User question: {user_query}
Answer the question based on the extracted content above."""
    
    payload = {
        "model": "mistral",
        "prompt": prompt,
        "stream": False
    }
    response = requests.post(OLLAMA_URL, json=payload)
    return response.json()["response"]
def visual_agent(image_path, user_query):
    # Step 1: Extract information from image
    extraction_prompt = "Describe all text, data, and visual elements in this image in detail."
    extracted_content = vision_extract(image_path, extraction_prompt)
    
    # Step 2: Reason over extracted content
    final_response = text_reason(extracted_content, user_query)
    
    return final_response
# Usage
result = visual_agent("dashboard_screenshot.png", "What metrics are underperforming?")
print(result)
```
This is bare-bones but functional. You can extend it with conversation history, multiple images, and task-specific extraction prompts.

## Handling Specific Use Cases

The basic pipeline works differently depending on what you’re processing. Here’s how to adapt it for common scenarios.

### Screenshots and UI Images

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

Screenshots are usually high-contrast with clear text, which vision models handle well. The key is giving the vision model a specific extraction prompt rather than a generic one.

Instead of: `"What's in this image?"`

Use: `"Extract all text, numbers, and labels visible in this screenshot. Organize by section if possible. Also note any status indicators, color coding, or highlighted items."`

Specific prompts produce structured output that’s much easier for your text LLM to reason over.

### PDF Documents

PDFs require an extra preprocessing step — converting pages to images. The `pdf2image` Python library handles this cleanly:

```
from pdf2image import convert_from_path
def process_pdf(pdf_path, user_query):
    pages = convert_from_path(pdf_path, dpi=150)
    
    all_extracted_text = []
    
    for i, page in enumerate(pages):
        page_path = f"/tmp/page_{i}.png"
        page.save(page_path, "PNG")
        
        extracted = vision_extract(
            page_path, 
            "Extract all text from this document page, preserving structure and formatting."
        )
        all_extracted_text.append(f"[Page {i+1}]\n{extracted}")
    
    combined_content = "\n\n".join(all_extracted_text)
    return text_reason(combined_content, user_query)
```
For long PDFs, you may want to process only relevant pages or use a chunking strategy. Passing 50 pages of extracted text into a single context window will hit token limits.

### Charts and Graphs

Charts are the hardest case for small vision models. Ask specifically what you need:

`"This is a chart. State the chart type, the title, all axis labels, the data series names, and describe the trend or key values visible in the data."`

For precise numerical extraction from charts, small models can struggle. If accuracy is critical, consider preprocessing charts with dedicated tools or reserving the larger vision model for these cases.

### Handwritten Notes and Forms

Handwriting varies widely in legibility. Use higher DPI when scanning (300+ DPI), and set expectations appropriately. The prompt should acknowledge uncertainty:

`"Transcribe any handwritten text in this image. If a word is unclear, indicate it with [unclear]. Transcribe printed text exactly."`

## Managing VRAM Efficiently

Running two models simultaneously creates memory pressure. Here are practical strategies to stay within your budget.

**Sequential loading, not concurrent.** If you’re VRAM-constrained, don’t try to keep both models loaded simultaneously. Load the vision model, run inference, unload it, then load the text model. Ollama handles model loading automatically — if you send a request to `minicpm-v` and then to `mistral`, it will manage swapping. The tradeoff is latency: each model load takes 5–15 seconds.

**Use Ollama’s `keep_alive` parameter.** If you’re processing batches of images, set `keep_alive` to a higher value to keep the vision model warm between requests. Set it to `0` to unload immediately after a single use.

**Quantized models.** Most models on Ollama come quantized by default (Q4 or Q5). This reduces memory usage by 50–75% compared to full precision with minimal quality loss for most tasks. If you’re pulling a model and it offers multiple quantization levels, Q4_K_M is generally the best balance.

**CPU offloading.** Tools like llama.cpp (which Ollama uses under the hood) support partial GPU offloading. If you have 8GB VRAM and a model needs 10GB, you can offload some layers to CPU RAM. This slows inference but keeps the model functional.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

**Consider CPU-only for the vision model.** Vision inference on a small 2B model is fast — even on CPU, extraction from a single image typically takes 10–30 seconds. If you need to save all GPU VRAM for your reasoning LLM, run the vision model on CPU using a lightweight framework like llama.cpp directly.

## Common Mistakes and How to Avoid Them

**Generic extraction prompts.** The biggest quality issue in dual-model pipelines is vague vision prompts. “What do you see?” produces vague answers. Specific prompts produce usable data. Always tailor your extraction prompt to the document type.

**Ignoring image quality.** Small vision models struggle with blurry, low-resolution, or low-contrast images. A 72 DPI screenshot that’s been compressed twice will produce poor extractions. When possible, process images at 150–300 DPI and avoid aggressive JPEG compression.

**No validation step.** Vision models hallucinate, especially small ones. If you’re extracting numbers, names, or specific data points, add a validation step. Either have the text LLM confirm what it read makes sense, or implement structured output parsing with error handling.

**Context window overflow.** A dense, text-heavy PDF page can produce 2,000–4,000 tokens of extracted text. Multiply that across many pages and you’ll hit the context limit of your text LLM. Plan for chunking or summarization at the vision stage.

**Assuming the vision model understands intent.** The vision model doesn’t know what you’re going to do with its output. It just describes what it sees. If you want structured JSON, ask for structured JSON. If you want a specific field extracted, name it explicitly.

## Frequently Asked Questions

### Can I run both models at the same time on 8GB of VRAM?

It’s tight. Two 2B models in 4-bit quantization each use roughly 2–3GB of VRAM, so in theory both fit with room to spare. In practice, model loading and inference create memory spikes. Ollama will automatically swap models, which adds latency. For a smoother experience, target a vision model under 3GB and a text model under 4GB, or accept the swap latency.

### What’s the best small vision model for reading text from screenshots?

MiniCPM-V 2B has strong OCR-like performance for a model its size. It consistently outperforms Moondream on dense text extraction tasks. If you’re primarily reading text from UI screenshots or document images, MiniCPM-V is the best starting point among the sub-4GB options.

### Can this approach handle multi-page documents?

Yes, but you need to process each page separately and then combine the extracted text before passing it to your reasoning model. For long documents, implement a retrieval step — extract all pages, then search for relevant sections based on the user’s query before passing content to the LLM. This avoids context window limits.

### How accurate is text extraction compared to dedicated OCR tools?

Modern small vision models rival traditional OCR tools like Tesseract for clean, printed text. They’re often better at understanding document structure (tables, headers, columns). They’re worse at handwriting, damaged documents, and very small text. For high-stakes extraction where errors are costly, combine vision model output with validation logic.

### Do I need a GPU to run these vision models?

No. Moondream and MiniCPM-V can run on CPU using llama.cpp. Inference is slower — 30–90 seconds per image on a modern CPU — but that’s often acceptable if you’re processing batches or running in a background workflow rather than real-time. GPU inference typically takes 2–8 seconds for the same models.

### What about using a vision model through an API instead of running it locally?

If privacy isn’t a concern, API-based options like GPT-4o Vision, Claude 3.5 Sonnet, or Gemini Flash Vision are more capable than small local models and simpler to integrate. The local dual-model architecture makes sense specifically when you need privacy, offline capability, or want to avoid per-image API costs at scale.

## Key Takeaways

- **Separate vision from reasoning.** A small vision model (2–4GB VRAM) extracts visual content into text. Your text LLM reasons over that text. Two small models fit where one large multimodal model won’t.
- **MiniCPM-V and Moondream are strong starting points.** Both run well on consumer hardware via Ollama and handle the most common use cases — screenshots, PDFs, charts — competently.
- **Prompt specificity matters more than model size.** Telling the vision model exactly what to extract produces dramatically better results than generic prompts.
- **PDFs need preprocessing.** Convert pages to images with`pdf2image` before passing to the vision model, and plan for context length when working with long documents.
- **Quantized models are the default.** Q4_K_M quantization cuts memory use by 60–70% with minimal quality loss — use it.

If you want to wire this pipeline into a complete automated workflow — with triggers, integrations, and no hand-coded orchestration — MindStudio’s support for local models makes it straightforward to build on top of what you’ve set up locally.
