---
id: collect-261001-rattrapage/rattrapage/github-moonshotai-kimi-vl-kimi-vl-mixture-of-experts-vision-language-model-for-multimodal--1
title: "If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation=\"flash_attention_2\""
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Meta", "Moonshot", "OpenAI", "vLLM"]
dates: []
keywords: ["agent", "benchmark", "context window", "cost", "deepseek", "fine-tuning", "inference", "kimi", "llama", "memory", "moe", "multimodal"]
source: docs/RAG/collect-261001-rattrapage/github-moonshotai-kimi-vl-kimi-vl-mixture-of-experts-vision-language-model-for-multimodal-reasoning-.md
source_anchor: ""
source_lines: [1, 57]
sha256: a1c365ecd45b821e79fae795cf3e06289d6f7a9b80074dd32f2805634166cc34
---

# If flash-attn has been installed, it is recommended to set torch_dtype=torch.bfloat16 and attn_implementation="flash_attention_2"

We present **Kimi-VL**, an efficient open-source Mixture-of-Experts (MoE) vision-language model (VLM) that offers **advanced multimodal reasoning, long-context understanding, and strong agent capabilities**—all while activating only **2.8B** parameters in its language decoder (Kimi-VL-A3B).

Kimi-VL demonstrates strong performance across challenging domains: as a general-purpose VLM, Kimi-VL excels in multi-turn agent interaction tasks (e.g.,OSWorld), achieving state-of-the-art results comparable to flagship models. Furthermore, it exhibits remarkable capabilities across diverse challenging vision language tasks, including college-level image and video comprehension, optical character recognition (OCR), mathematical reasoning, multi-image understanding, and etc.

In comparative evaluations, it effectively competes with cutting-edge efficient VLMs such as GPT-4o-mini, Qwen2.5-VL-7B, and Gemma-3-12B-IT, while surpassing GPT-4o in several specialized domains.

Kimi-VL also advances the pareto frontiers of multimodal models in processing long contexts and perceiving clearly: Equipped with a 128K extended context window, Kimi-VL can processes long and diverse inputs, achieving impressive scores of 64.5 on LongVideoBench, and 35.1 on MMLongBench-Doc; Its native-resolution vision encoder, MoonViT, further allows it to see and understand ultra-high-resolution visual inputs, achieving 83.2 on InfoVQA and 34.5 on ScreenSpot-Pro, while maintaining lower computational cost with common visual inputs and general tasks.

Building on this foundation, we introduce an advanced long-thinking variant: **Kimi-VL-Thinking**. Developed through long chain-of-thought (CoT) supervised fine-tuning (SFT) and reinforcement learning (RL), this model exhibits strong long-horizon reasoning capabilities. It achieves scores of 61.7 on MMMU, 36.8 on MathVision, and 71.3 on MathVista while maintaining the compact 2.8B activated LLM parameter footprint, setting a new standard for efficient yet capable multimodal **thinking** models.

*Besides original model variants, we also provide a new Kimi-VL-A3B-Thinking-2506 variant with several new or improved abilities:*


The model adopts an MoE language model, a native-resolution visual encoder (MoonViT), and an MLP projector, as illustrated in the following image.

- 2025.06.21: Release of Kimi-VL-A3B-Thinking-2506: Tech Blog & Cookbook, 🤗 Hugging Face
- 2025.04.15: vLLM has supported Kimi-VL deployment. See #16387 for details.
- 2025.04.14: LLaMA-Factory has supported Kimi-VL finetuning. See #7719 for details.

🤗 For common general multimodal perception and understanding, OCR, long video and long document, video perception, and OS-agent uses, we recommend `Kimi-VL-A3B-Instruct` for efficient inference; meanwhile, our new thinking version, `Kimi-VL-A3B-Thinking-2506` also has excellent multimodal perception, long video and long document and OS-agent grounding abilities while achieving better multimodal reasoning skills. See this blog for more information.

| **Model** | **#Total Params** | **#Activated Params** | **Context Length** | **Download Link** | 
|---|---|---|---|---|
| 🔥Kimi-VL-A3B-Thinking-2506 | 16B | 3B | 128K | 🤗 Hugging Face | 
| Kimi-VL-A3B-Instruct | 16B | 3B | 128K | 🤗 Hugging Face | 
| Kimi-VL-A3B-Thinking (deprecated) | 16B | 3B | 128K | 🤗 Hugging Face | 

Note

Recommended parameter settings:

- For **Thinking models** , it is recommended to use`Temperature = 0.8` .
- For **Instruct models** , it is recommended to use`Temperature = 0.2` .

🤗 We serve our model demo in Hugging Face spaces:


- Chat with
**Kimi-VL-A3B-Thinking-2506**👀🤔🗺️🎬📖🖥️ (*unifying thinking, general understanding, puzzle solving, agent, video, PDF*) model on Chat Web.

As an efficient model, Kimi-VL can robustly handle diverse tasks (fine-grained perception, math, college-level problems, OCR, agent, etc) across a broad spectrum of input forms (single-image, multi-image, video, long-document, etc).

A brief comparison with existing 10B-level dense VLMs and DeepSeek-VL2 (A4.5B):

With effective long-thinking abilities, Kimi-VL-A3B-Thinking (2504 version) can match the performance of 30B/70B frontier open-source VLMs on MathVision benchmark:

```
conda create -n kimi-vl python=3.10 -y
conda activate kimi-vl
pip install -r requirements.txt
```
Note

If you encounter Out-of-Memory or want to speed up inference, please install **flash-attn** with `pip install flash-attn --no-build-isolation`.

We introduce how to use our model at inference stage using transformers library. It is recommended to use python=3.10, torch=2.5.1, and transformers=4.51.3 as the development environment.

