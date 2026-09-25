---
id: collect-250926-servers-hardware/servers-hardware/diffusiongemma-how-to-run-locally
title: "DiffusionGemma - How to Run Locally"
domain: servers-hardware
role: reference
task: reference
actors: ["Google", "Unsloth"]
dates: []
keywords: ["diffusion", "agent", "agentic", "agents", "benchmark", "benchmarks", "fine-tuning", "gguf", "gpu", "inference", "latency", "llama"]
source: docs/RAG/clean4/DiffusionGemma - How to Run Locally.md
source_anchor: ""
source_lines: [1, 102]
sha256: 8ecc8da29d6ae0ce49d351ea0a00a5e33f4637f84249da363ea3401b8958f1c8
---

# DiffusionGemma - How to Run Locally

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/diffusiongemma.md).
# DiffusionGemma - How to Run Locally
DiffusionGemma **26B-A4B** is Google DeepMind’s new open **multimodal** model, built on the [Gemma 4](/docs/models/gemma-4.md) MoE architecture. With support for **256K context**, **140+ languages**, DiffusionGemma is designed for **high-speed text generation** across text, video and image inputs. DiffusionGemma can run locally on **18GB RAM**, and [fine-tuning](#fine-tune-diffusiongemma) is now supported via [Unsloth](https://github.com/unslothai/unsloth).
Instead of standard token-by-token decoding, DiffusionGemma uses **diffusion generation** to produce outputs in parallel and gradually refine them into a final answer - similar to diffusion image models, but for text. Run the model via [Unsloth Studio](/docs/new/studio.md) or llama.cpp. On a RTX 6000, DiffusionGemma can reach **2000+ tokens/s**. **GGUF:** [diffusiongemma-26B-A4B-it-GGUF](https://huggingface.co/unsloth/diffusiongemma-26B-A4B-it-GGUF)
Run DiffusionGemmaFine-tune DiffusionGemma
{% hint style="success" %}
**Jun 12:** You can now run DiffusionGemma via [Unsloth Studio](#unsloth-studio-guide) ✨ with 1.8x faster inferece!
{% endhint %}
### Usage Guide
DiffusionGemma is designed for users who need faster generation than standard models. It is suited to fast local inference, long-context doc analysis, image/video understanding, OCR and document parsing, code generation, tool use, agentic workflows, and low-latency small-batch inference.
Unlike standard Gemma 4 models, DiffusionGemma requires a diffusion-aware inference runtime. Autoregressive settings such as `temperature`, `top_p`, and `top_k` alone are not enough to reproduce the recommended behavior without the required diffusion sampler.

🦥 Unsloth Studio Guide🦙 Llama.cpp Guide
### 🦥 Unsloth Studio Guide
{% hint style="success" %}
You can now run DiffusionGemma via [Unsloth Studio](#unsloth-studio-guide) ✨. Ensure you use [`v0.1.463-beta`](https://github.com/unslothai/unsloth/tree/v0.1.462-beta) or `2026.6.6`.
{% endhint %}
DiffusionGemma can now be run and trained in [Unsloth Studio](/docs/new/studio.md), our new open-source web UI for local AI. Unsloth Studio lets you run models locally on **MacOS**, **Windows**, Linux and:
{% columns %}
{% column %}
* Search, download, [run GGUFs](/docs/new/studio.md#run-models-locally) and safetensor models
* [**Self-healing** tool calling](/docs/new/studio.md#execute-code--heal-tool-calling) + **web search**
* [**Code execution**](/docs/new/studio.md#run-models-locally) (Python, Bash)
* [Automatic inference](https://unsloth.ai/docs/desktop#feature-deep-dive) parameter tuning (temp, top-p, etc.)
* Fast CPU + GPU inference via llama.cpp
* [Train LLMs](/docs/new/studio.md#no-code-training) 2x faster with 70% less VRAM
{% endcolumn %}
{% column %}
" %}
[final answer]
```
To disable thinking, remove the `<|think|>` token from the system prompt. When thinking is disabled, the model may still emit an empty thought channel:
```
<|channel>thought
[final answer]
```
For multi-turn conversations, do **not** include previous hidden thoughts in the conversation history. Only include the final assistant response before the next user turn.
## DiffusionGemma Best Practices
### Multimodal Prompting
DiffusionGemma supports interleaved multimodal inputs, including text and images. Video can be processed as sequences of image frames.
For best results with multimodal prompts, place image or frame content before text instructions. Example:
```
[image]
Describe the chart and summarize the key trend.
```
For document parsing, OCR, chart understanding, UI understanding, or small text extraction, use a higher visual token budget.
Supported visual token budgets:
| Visual Token Budget | Best For |
| ------------------- | ------------------------------------------------ |
| 70 | Fast classification, simple captioning |
| 140 | Lightweight visual QA |
| 280 | General image understanding |
| 560 | OCR, charts, UI screenshots |
| 1120 | Dense documents, small text, detailed extraction |
For video-style inputs, DiffusionGemma can process up to **60 seconds** when sampled at **1 frame per second**.
### Sampling Notes
DiffusionGemma is not a normal next-token-only model. It generates a block of tokens, called a **canvas**, by repeatedly refining noisy token predictions. The generation process works roughly as follows:
1. The encoder processes the prompt and builds a context cache.
2. The decoder receives a 256-token generation canvas.
3. The diffusion sampler iteratively denoises the canvas.
4. Confident tokens are selected and preserved.
5. Uncertain tokens are renoised and refined again.
6. Once the canvas is complete, it is appended to the context.
7. The model continues with the next canvas.
This block-autoregressive approach allows DiffusionGemma to generate many tokens in fewer forward passes than a standard autoregressive model.
## Benchmarks
DiffusionGemma is optimized for speed and multimodal reasoning, though standard Gemma 4 is stronger on conventional reasoning benchmarks.
| Benchmark | DiffusionGemma 26B-A4B | Gemma 4 26B-A4B |
| ------------------- | ---------------------: | --------------: |
| MMLU Pro | 77.6% | 82.6% |
| AIME 2026 no tools | 69.1% | 88.3% |
| LiveCodeBench v6 | 69.1% | 77.1% |
| Codeforces ELO | 1429 | 1718 |
| GPQA Diamond | 73.2% | 82.3% |
| Tau2 Average | 56.2% | 68.2% |
| HLE no tools | 11.0% | 8.7% |
| HLE with search | 11.9% | 17.2% |
| BigBench Extra Hard | 47.6% | 64.8% |
| MMMLU | 81.5% | 86.3% |
| Long Context Benchmark | DiffusionGemma 26B-A4B | Gemma 4 26B-A4B |
| ----------------------------- | ---------------------: | --------------: |
| MRCR v2 8 needle 128K average | 32.0% | 44.1% |
**Vision benchmarks:**
| Vision Benchmark | DiffusionGemma 26B-A4B | Gemma 4 26B-A4B |
| --------------------------------- | ---------------------: | --------------: |
| MMMU Pro | 54.3% | 73.8% |
| OmniDocBench 1.5, lower is better | 0.319 | 0.149 |
| MATH-Vision | 70.5% | 82.4% |
| MedXPertQA MM | 49.0% | 58.1% |
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/models/diffusiongemma.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
