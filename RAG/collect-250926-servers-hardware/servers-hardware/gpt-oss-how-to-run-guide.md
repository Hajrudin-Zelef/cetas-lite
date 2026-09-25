---
id: collect-250926-servers-hardware/servers-hardware/gpt-oss-how-to-run-guide
title: "gpt-oss: How to Run Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "OpenAI", "Unsloth"]
dates: []
keywords: ["agentic", "apache", "fine-tuning", "gguf", "inference", "license", "llama", "llama.cpp", "memory", "reasoning", "research", "tool use"]
source: docs/RAG/clean4/gpt-oss How to Run Guide.md
source_anchor: ""
source_lines: [1, 27]
sha256: 16293672d38415035a02561e611184a4e0e19e63f773e0fc33665bd6502c36f4
---

# gpt-oss: How to Run Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/gpt-oss-how-to-run-and-fine-tune.md).
# gpt-oss: How to Run Guide
Run & fine-tune OpenAI's new open-source models!
OpenAI releases '**gpt-oss-120b'** and '**gpt-oss-20b'**, two SOTA open language models under the Apache 2.0 license. Both 128k context models outperform similarly sized open models in reasoning, tool use, and agentic tasks. You can now run & fine-tune them locally with Unsloth!
Run gpt-oss-20bRun gpt-oss-120bFine-tune gpt-oss
> [**Fine-tune**](#fine-tuning-gpt-oss-with-unsloth) **gpt-oss-20b for free with our** [**Colab notebook**](https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/gpt-oss-$20B$-Fine-tuning.ipynb)
Trained with [RL](/docs/get-started/reinforcement-learning-rl-guide.md), **gpt-oss-120b** rivals o4-mini and **gpt-oss-20b** rivals o3-mini. Both excel at function calling and CoT reasoning, surpassing o1 and GPT-4o.
For best performance, make sure your total available memory (unified mem + VRAM + system RAM) exceeds the size of the quantized model file you’re downloading. If it doesn’t, llama.cpp can still run via SSD/HDD offloading, but inference will be slower.
#### **gpt-oss - Unsloth GGUFs:**
{% hint style="success" %}
**Includes Unsloth's** [**chat template fixes**](#unsloth-fixes-for-gpt-oss)**. For best results, use our uploads & train with Unsloth!**
{% endhint %}
* 20B: [gpt-oss-**20B**](https://huggingface.co/unsloth/gpt-oss-20b-GGUF)
* 120B: [gpt-oss-**120B**](https://huggingface.co/unsloth/gpt-oss-120b-GGUF)
## :scroll:Unsloth fixes for gpt-oss
{% hint style="info" %}
Some of our fixes were pushed upstream to OpenAI's official model on Hugging Face. [See](https://huggingface.co/openai/gpt-oss-20b/discussions/94/files)
{% endhint %}
OpenAI released a standalone parsing and tokenization library called [Harmony](https://github.com/openai/harmony) which allows one to tokenize conversations to OpenAI's preferred format for gpt-oss.
Inference engines generally use the jinja chat template instead and not the Harmony package, and we found some issues with them after comparing with Harmony directly. If you see below, the top is the correct rendered form as from Harmony. The below is the one rendered by the current jinja chat template. There are quite a few differences!

&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
