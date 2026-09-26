---
id: collect-240926-mindstudio/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram-3
title: "Install Ollama (macOS/Linux)"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["llama", "claude", "consumer", "context window", "gemini", "gpu", "inference", "llama.cpp", "memory", "multimodal", "quantization", "reasoning"]
source: docs/RAG/clean_en/mindstudio/how-to-add-vision-capabilities-to-a-local-ai-agent-without-blowing-your-vram.md
source_anchor: ""
source_lines: [289, 315]
sha256: ea54b385417d0eaf522032d11cced634d3ccb2a4fdea71ccf2b396aa5c4be0bd
---

# Install Ollama (macOS/Linux)

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
