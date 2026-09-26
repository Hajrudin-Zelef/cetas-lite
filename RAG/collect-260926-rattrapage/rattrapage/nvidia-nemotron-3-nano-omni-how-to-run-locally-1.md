---
id: collect-260926-rattrapage/rattrapage/nvidia-nemotron-3-nano-omni-how-to-run-locally-1
title: "NVIDIA Nemotron 3 Nano Omni - How To Run Locally"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia", "Unsloth"]
dates: []
keywords: ["nvidia", "omni", "agentic", "gguf", "inference", "llama", "llama.cpp", "moe", "multimodal", "reasoning", "tool calling", "training"]
source: docs/RAG/lot-rattrapage/servers-reviews/NVIDIA Nemotron 3 Nano Omni - How To Run Locally.md
source_anchor: ""
source_lines: [1, 40]
sha256: 0301efbe12565d0f73eec73462db467653a124fd5e558468f70c3a30162ce444
---

# NVIDIA Nemotron 3 Nano Omni - How To Run Locally

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/nemotron-3-nano-omni.md).
# NVIDIA Nemotron 3 Nano Omni - How To Run Locally
Run & fine-tune Nemotron-3-Nano-Omni-30B-A3B locally on your device!
NVIDIA Nemotron-3-Nano-Omni-30B-A3B is an open 30B parameter, 3B active hybrid reasoning MoE model built for multimodal agentic workloads including **audio**, **video**, text, images and docs as input, with text output. The model runs on **25GB RAM** for 4-bit and 36GB for 8-bit.
With a **256K context**, Nemotron 3 Nano Omni is the **strongest omni** model for its size and the highest-efficiency open multimodal model. We collaborated with NVIDIA for day zero support!\
**GGUF:** [Nemotron-3-Nano-Omni-30B-A3B-Reasoning](https://huggingface.co/unsloth/Nemotron-3-Nano-30B-A3B-GGUF)
### ⚙️ Usage Guide
NVIDIA recommends these settings for inference:
{% columns %}
{% column %}
**Thinking mode:**
* `temperature = 0.6`
* `top_p = 0.95`
{% endcolumn %}
{% column %}
**Instruct mode:**
* `temperature = 0.2`
{% endcolumn %}
{% endcolumns %}
### Run Nemotron-3-Nano-Omni
Depending on your use-case you will need to use [different settings](#usage-guide). Some GGUFs end up similar in size because the model architecture (like [gpt-oss](/docs/models/gpt-oss-how-to-run-and-fine-tune.md)) has dimensions not divisible by 128, so parts can’t be quantized to lower bits. **GGUF:** [Nemotron-3-Nano-Omni-30B-A3B-Reasoning](https://huggingface.co/unsloth/Nemotron-3-Nano-30B-A3B-GGUF)
The 4-bit versions of the model requires \~25GB RAM. 8-bit requires 36GB. For these guides, we will be using `UD-Q4-K-XL` which is a good balance between size and accuracy.
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

