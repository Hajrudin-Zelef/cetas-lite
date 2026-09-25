---
id: collect-250926-servers-hardware/servers-hardware/eb1c9b1e
title: "How to Run models with Unsloth Studio"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["agent", "agents", "claude", "gguf", "gpu", "gpus", "inference", "lora", "nvidia", "safetensors", "tool calling"]
source: docs/RAG/clean4/eb1c9b1e.md
source_anchor: ""
source_lines: [1, 52]
sha256: 131fce5d9879fe8bbb29e8a828b509225cc13f307a6b1d168276cfdd38dbe261
---

# How to Run models with Unsloth Studio

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/new/studio/chat.md).
# How to Run models with Unsloth Studio
Run AI models, LLMs and GGUFs locally with Unsloth Studio.
[Unsloth Studio](/docs/new/studio.md) lets you run AI models 100% offline on your computer. Run model formats like GGUF and safetensors from Hugging Face or from your local files.
* **Works on all MacOS, CPU, Windows, Linux, WSL setups! No GPU required**
* [**Self-healing tool calling**](#auto-healing-tool-calling)**,** advanced [**web search**](#advanced-web-search), [**code execution**](#code-execution)
* Use Unsloth as an OpenAI-compatible inference [**API endpoint**](/docs/basics/api.md) or connect a [provider](/docs/integrations/connections.md)
* Search + Download + Run + [Compare](#model-arena) any model like GGUFs, LoRA adapters, safetensors etc.
* [**Auto inference parameter**](#auto-parameter-tuning) tuning (temp, top-p etc.) and edit chat templates
* Upload images, audio, PDFs, code, DOCX and more file types to chat with.
### Using Unsloth Studio Chat
{% hint style="success" %}
Unsloth Studio Chat automatically works on **multi-GPU setups** for inference.
{% endhint %}
{% columns %}
{% column %}
#### Code execution
Unsloth Studio lets LLMs run Bash and Python, not just JavaScript. It also sandboxes programs like Claude Artifacts so models can test code, generate files, and verify answers with real computation.
This makes answers from models more reliable and accurate.
{% endcolumn %}
{% column %}

\.cache\huggingface\hub\` on Windows.
* **MacOS, Linux, WSL:** `~/.cache/huggingface/hub/`
* **Windows:** `%USERPROFILE%\.cache\huggingface\hub\`
If `HF_HUB_CACHE` or `HF_HOME` is set, use that location instead. On Linux and WSL, `XDG_CACHE_HOME` can also change the default cache root.
### **Unsloth not detecting or using my GPU**
If the model is not using your GPU specifically for Docker, try:
Pulling the latest image manually:
```bash
docker pull unsloth/unsloth:latest
```
* Start the container with GPU access:
* `docker run`: `--gpus all`
* Docker Compose: `capabilities: [gpu]`
* On Linux, make sure the NVIDIA Container Toolkit is installed.
* On Windows:
* Check that `nvcc --version` matches the CUDA version shown in `nvidia-smi`
* Follow:
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/new/studio/chat.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
