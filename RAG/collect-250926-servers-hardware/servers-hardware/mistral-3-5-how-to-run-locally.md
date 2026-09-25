---
id: collect-250926-servers-hardware/servers-hardware/mistral-3-5-how-to-run-locally
title: "Mistral 3.5 - How To Run Locally"
domain: servers-hardware
role: reference
task: reference
actors: ["Mistral", "Nvidia", "Unsloth"]
dates: ["2026-05-01"]
keywords: ["mistral", "agent", "agentic", "agents", "context window", "gguf", "inference", "llama", "llama.cpp", "memory", "multimodal", "nvidia"]
source: docs/RAG/clean4/Mistral 3.5 - How To Run Locally.md
source_anchor: ""
source_lines: [1, 76]
sha256: 4e13d62d183f90301558771abc2ee114a89ccdb388a21fbf7fc0e0683b6b542e
---

# Mistral 3.5 - How To Run Locally

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/mistral-3.5.md).
# Mistral 3.5 - How To Run Locally
Guide for Mistral Mistral 3.5 models, to run or fine-tune locally on your device
Mistral releases Mistral-Medium-3.5-128B, their new dense 128B parameter, multimodal, hybrid reasoning model. It supports text and image input, text output, a 256K context window and excels at reasoning, coding, long-context, tool use, agentic workflows, and multimodal doc/image understanding.
Mistral Medium 3.5 offers highly competitive performance for models 5x its size. Run locally on \~64GB RAM. GGUF: [Mistral-Medium-3.5-128B-GGUF](https://huggingface.co/unsloth/Mistral-Medium-3.5-128B-GGUF)
{% hint style="success" %}
**May 1, 2026 Update:** We worked with Mistral to fix Mistral Medium 3.5 inference affecting some implementations, and released updated GGUFs with the fix (**NOT related to Unsloth** or our quants). The issue was caused by a YaRN parsing quirk affecting several implementations, including `transformers` and `llama.cpp`. Changing `mscale_all_dim` from `1` to `0` resolved it. We also fixed `mmproj` files not being generated correctly.
**Mistral has now pushed our fixes to their official repo!**
{% endhint %}
### Usage Guide
{% hint style="info" %}
Vision for GGUFs it now supported for now. Support will come later.
{% endhint %}
Table: Mistral Medium 3.5 recommended hardware requirements. Units are total memory: RAM + VRAM, or unified memory.
| Mistral 3.5 | 3-bit | 4-bit | 8-bit |
| --------------- | ----- | ----- | ---------- |
| Medium 3.5 128B | 64 GB | 80 GB | 128-170 GB |
{% hint style="info" %}
Your total available memory should at least exceed the size of the quantized model you download. If it does not, llama.cpp can still run with partial RAM / disk offload, but generation will be slower. You will also need more memory for long context, larger batches, tool-heavy agent runs and image prompts.
{% endhint %}
#### Recommended Settings
Use Mistral's recommended reasoning settings:
* `reasoning_effort="none"` → fast instant replies, chat, extraction and simple instructions.
* `reasoning_effort="high"` → reasoning mode, recommended for complex prompts, coding, research, math and agentic usage.
Recommended sampling defaults:
* Use `temperature = 0.7` for `reasoning_effort="high"`.
* Use `temperature = 0.0` to `0.7` for `reasoning_effort="none"`, depending on the task.
* Keep repetition and presence penalties disabled or at `1.0` unless you see looping.
* Maximum context length of `262,144`
#### **Reasoning Mode**
Mistral Medium 3.5 supports instant instruct mode and reasoning mode with a 'high' option.
To enable high reasoning for llama.cpp / llama-server:
```bash
--chat-template-kwargs '{"reasoning_effort":"high"}'
```
To disable reasoning:
```bash
--chat-template-kwargs '{"reasoning_effort":"none"}'
```
If you're on Windows PowerShell, use:
```powershell
--chat-template-kwargs "{\"reasoning_effort\":\"none\"}"
```
## Run Mistral 3.5 Tutorials
Because Mistral Medium 3.5 is a dense 128B model, the recommended starting point is Dynamic 4-bit GGUFs for local inference. GGUF: `unsloth/Mistral-Medium-3.5-128B-GGUF`
Run in Unsloth StudioRun in llama.cpp
{% hint style="warning" %}
Currently no multimodal/vision GGUF works in **Ollama** due to separate `mmproj` vision files. Use llama.cpp compatible backends.
Do NOT use **CUDA 13.2** as you may get gibberish outputs. NVIDIA is working on a fix.
{% endhint %}
### 🦥 Unsloth Studio Guide
For this tutorial, we will be using [Unsloth Studio](/docs/new/studio.md), which is our new web UI for running and training LLMs. With Unsloth Studio, you can run models and input **audio**, image and text locally on **Mac, Windows**, and Linux and:
{% columns %}
{% column %}
* Search, download, [run GGUFs](/docs/new/studio.md#run-models-locally) and safetensor models
* **Compare** models **side-by-side**
* [**Self-healing** tool calling](/docs/new/studio.md#execute-code--heal-tool-calling) + **web search**
* [**Code execution**](/docs/new/studio.md#run-models-locally) (Python, Bash)
* [Automatic inference](https://unsloth.ai/docs/desktop#feature-deep-dive) parameter tuning (temp, top-p, etc.)
* [Train LLMs](/docs/new/studio.md#no-code-training) 2x faster with 70% less VRAM
{% endcolumn %}
{% column %}

---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/models/mistral-3.5.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
