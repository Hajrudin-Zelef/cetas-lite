---
id: collect-240926-huggingface/huggingface/unsloth-llama-3-2-3b-instruct-bnb-4bit-hugging-face
title: "unsloth-llama-3-2-3b-instruct-bnb-4bit-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face", "Meta", "Mistral", "Unsloth", "vLLM"]
dates: []
keywords: ["llama", "agentic", "attention", "benchmarks", "dpo", "fine-tuning", "gguf", "gqa", "inference", "license", "memory", "mistral"]
source: docs/RAG/clean_en/huggingface/unsloth-llama-3-2-3b-instruct-bnb-4bit-hugging-face.md
source_anchor: ""
source_lines: [1, 51]
sha256: 3071d8dcda1b6d2035a108914fbdc3f6f9d382869cfa6157acdfe78a0cf6f84a
---

# unsloth-llama-3-2-3b-instruct-bnb-4bit-hugging-face

<!-- source: https://huggingface.co/unsloth/Llama-3.2-3B-Instruct-bnb-4bit -->

## 
	
		
	
		***See our collection for all versions of Llama 3.2 including GGUF, 4-bit and original 16-bit formats.***
	

We have a free Google Colab Tesla T4 notebook for Llama 3.2 (3B) here: https://colab.research.google.com/drive/1T5-zKWM_5OD21QHwXHiV9ixTRR7k3iB9?usp=sharing

For more details on the model, please go to Meta's original model card

All notebooks are **beginner friendly**! Add your dataset, click "Run All", and you'll get a 2x faster finetuned model which can be exported to GGUF, vLLM or uploaded to Hugging Face.

| Unsloth supports | Free Notebooks | Performance | Memory use | 
|---|---|---|---|
| **Llama-3.2 (3B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Llama-3.1 (11B vision)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Llama-3.1 (8B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Phi-3.5 (mini)** | ▶️ Start on Colab | 2x faster | 50% less | 
| **Gemma 2 (9B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Mistral (7B)** | ▶️ Start on Colab | 2.2x faster | 62% less | 
| **DPO - Zephyr** | ▶️ Start on Colab | 1.9x faster | 19% less | 

- This conversational notebook is useful for ShareGPT ChatML / Vicuna templates.
- This text completion notebook is for raw text. This DPO notebook replicates Zephyr.
- * Kaggle has 2x T4s, but we use 1. Due to overhead, 1x T4 is 5x faster.

A huge thank you to the Meta and Llama team for creating and releasing these models.

The Meta Llama 3.2 collection of multilingual large language models (LLMs) is a collection of pretrained and instruction-tuned generative models in 1B and 3B sizes (text in/text out). The Llama 3.2 instruction-tuned text only models are optimized for multilingual dialogue use cases, including agentic retrieval and summarization tasks. They outperform many of the available open source and closed chat models on common industry benchmarks.

**Model developer**: Meta

**Model Architecture:** Llama 3.2 is an auto-regressive language model that uses an optimized transformer architecture. The tuned versions use supervised fine-tuning (SFT) and reinforcement learning with human feedback (RLHF) to align with human preferences for helpfulness and safety.

**Supported languages:**  English, German, French, Italian, Portuguese, Hindi, Spanish, and Thai are officially supported. Llama 3.2 has been trained on a broader collection of languages than these 8 supported languages. Developers may fine-tune Llama 3.2 models for languages beyond these supported languages, provided they comply with the Llama 3.2 Community License and the Acceptable Use Policy. Developers are always expected to ensure that their deployments, including those that involve additional languages, are completed safely and responsibly.

**Llama 3.2 family of models** Token counts refer to pretraining data only. All model versions use Grouped-Query Attention (GQA) for improved inference scalability.

**Model Release Date:** Sept 25, 2024

**Status:** This is a static model trained on an offline dataset. Future versions may be released that improve model capabilities and safety.

**License:** Use of Llama 3.2 is governed by the Llama 3.2 Community License (a custom, commercial license agreement).

Where to send questions or comments about the model Instructions on how to provide feedback or comments on the model can be found in the model README. For more technical information about generation parameters and recipes for how to use Llama 3.1 in applications, please go here.

- Downloads last month
- 68,937
