---
id: collect-240926-mindstudio/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram-2
title: "Install Ollama (macOS/Linux)"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["llama", "agent", "agents", "context window", "decode", "gpu", "inference", "latency", "llama.cpp", "memory", "mistral", "quantization"]
source: docs/RAG/clean_en/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram.md
source_anchor: ""
source_lines: [143, 288]
sha256: 863179e8e1902e8954cccb2b96cee1f6ee66acab11d6232137ac0c4599c7212c
---

# Install Ollama (macOS/Linux)

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

