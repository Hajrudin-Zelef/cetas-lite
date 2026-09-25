---
id: collect-250926-servers-hardware/servers-hardware/how-to-run-models-with-unsloth-studio-unsloth-documentation
title: "how-to-run-models-with-unsloth-studio-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "Nvidia", "OpenAI", "Unsloth", "vLLM"]
dates: []
keywords: ["claude", "fine-tuning", "gguf", "gpu", "gpus", "inference", "llama", "llama.cpp", "lora", "multimodal", "nvidia", "parameters"]
source: docs/RAG/clean4/how-to-run-models-with-unsloth-studio-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 110]
sha256: e6954a32b511e28ba47dd5fc17b327333f98a7888908ab5566f297b37517ac05
---

# how-to-run-models-with-unsloth-studio-unsloth-documentation

Unsloth Studio lets you run AI models 100% offline on your computer. Run model formats like GGUF and safetensors from Hugging Face or from your local files.

- **Works on all MacOS, CPU, Windows, Linux, WSL setups! No GPU required**
- Upload images, audio, PDFs, code, DOCX and more file types to chat with.

Unsloth Studio Chat automatically works on **multi-GPU setups** for inference.

Unsloth Studio not only allows tool calling, but also auto-fixes malformed or broken tool-calls by 50%.

This means you'll always get inference outputs **without** broken tool calling. 

E.g. Qwen3.5-4B searched 20+ websites and cited sources, with web search happening inside its thinking trace.

Unsloth's unlimited and secure web search actually visits pages directly to collect relevant information and data and doesn't just scan through website summaries. This provides outputs much more accurate / in-depth info and context. Search uses DuckDuckGo's private and secure API.

You can now use local LLMs via tools like Claude Code and Codex by connecting it to Unsloth's API endpoint. This means you'll be able to directly run Qwen and Gemma models in those tools with Unsloth's inference which includes features like self-healing tool-calling, websearch etc.

Inference parameters like **temperature**, **top-p**, **top-k**, **MTP** are automatically pre-set for new models like Qwen3.5 so you can get the best outputs without worrying about settings. You can also adjust parameters manually and edit the system prompt.

Context length adjustment is no longer necessary with llama.cpp’s smart auto context, which uses only the context you need without loading anything extra.

Unsloth connects to OpenAI, Anthropic, Ollama, llama.cpp, vLLM, and others.

Add API keys or model server URLs, then use external models in the same chat interface as local + cloud models. Run with prompt caching, tool-calling, thinking, and provider-native features like OpenAI's web search and code execution.

You can search and download any model via Hugging Face or use local files.

Unsloth supports a wide range of model types, including **GGUF**, vision-language, and text-to-speech models. Run the latest models like Qwen3.5 or NVIDIA Nemotron 3.

Upload images, audio, PDFs, code, DOCX and more file types to chat with.

Unsloth offers several unique features that improve tool calling, including:

- Tool calls across all models in Unsloth are **30% to 80% more accurate** .
- Web search retrieves actual web content instead of only summaries.
- The maximum number of allowed tool calls is **more than 25.**
- Tool calls terminate more reliably, reducing loops and repeated calls.
- Improved tool-call healing and deduplication logic helps prevent XML from leaking into outputs.

See test results with `unsloth/Qwen3.5-4B-GGUF (UD-Q4_K_XL)` with web search, code execution, and thinking enabled:

XML leaks in response

10/10

0/10

URL fetches used

0

4/10 runs

Runs with correct song names

0/10

2/10

Avg tool calls

5.5

3.8

Avg response time

12.3s

9.8s

Unsloth Chat lets you compare any two models side-by-side using the same prompt. E.g. compare the base model and LoRa adapter. Inference will firstly load for one model, then the second one (parallel inference is being worked on).

After training, you can compare the base and fine-tuned models side by side with the same prompt to see what changed and whether results improved.

This workflow makes it easy to see how your fine-tuning changed the model’s responses and whether it improved results for your use case.

Unsloth Studio Chat auto works on **multi-GPU setups** for inference.

**Apr 1 update:** You can now select an existing folder for Unsloth to detect from.

**Mar 27 update:** Unsloth Studio now **automatically detects older / pre-existing models** downloaded from Hugging Face, LM Studio etc.

**Manual instructions:** Unsloth Studio detects models downloaded to your Hugging Face Hub cache `(C:\Users{your_username}.cache\huggingface\hub)`. If you have GGUF models downloaded through LM Studio, note that these are stored in `C:\Users\{your_username}.cache\lm-studio\models` **OR**`C:\Users{your_username}\lm-studio\models` and are not visible to llama.cpp by default - you will need to move or copy those .gguf files into your Hugging Face Hub cache directory (or another path accessible to llama.cpp) for Unsloth Studio to load them.

After fine-tuning a model or adapter in Unsloth, you can export it to GGUF and run local inference with **llama.cpp** directly in Unsloth Chat. Unsloth Studio is powered by llama.cpp and Hugging Face.

Unsloth Chat supports multimodal inputs directly in the conversation. You can attach documents, images, or audio as additional context for a prompt.

This makes it easy to test how a model handles real-world inputs such as PDFs, screenshots, or reference material. Files are processed locally and included as context for the model.

You can delete old model files either from the bin icon in model search or by removing the relevant cached model folder from the default Hugging Face cache directory. By default, Hugging Face uses `~/.cache/huggingface/hub/` on macOS/Linux/WSL and `C:\Users\<username>\.cache\huggingface\hub\` on Windows.

- **MacOS, Linux, WSL:**`~/.cache/huggingface/hub/`
- **Windows:**`%USERPROFILE%\.cache\huggingface\hub\`

If `HF_HUB_CACHE` or `HF_HOME` is set, use that location instead. On Linux and WSL, `XDG_CACHE_HOME` can also change the default cache root.

If the model is not using your GPU specifically for Docker, try:

Pulling the latest image manually:

- Start the container with GPU access: 
  - `docker run` :`--gpus all`
  - Docker Compose: `capabilities: [gpu]`
- On Linux, make sure the NVIDIA Container Toolkit is installed.

Last updated

Was this helpful?
