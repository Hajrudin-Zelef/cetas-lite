---
id: collect-240926-huggingface/huggingface/unsloth-llama-3-3-70b-instruct-bnb-4bit-hugging-face-1
title: "First, define a tool"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Google", "Hugging Face", "Meta", "Mistral", "Unsloth", "vLLM"]
dates: ["2023-12", "2024-12-06"]
keywords: ["attention", "benchmarks", "distillation", "dpo", "fine-tuning", "gguf", "gqa", "inference", "license", "llama", "memory", "mistral"]
source: docs/RAG/clean_en/huggingface/unsloth-llama-3-3-70b-instruct-bnb-4bit-hugging-face.md
source_anchor: ""
source_lines: [1, 131]
sha256: 6646cae1316d1ccb8b9a5a71451aaefd59a58e73b1b20d9097dbcb09cf827204
---

# First, define a tool

<!-- source: https://huggingface.co/unsloth/Llama-3.3-70B-Instruct-bnb-4bit -->

## 
	
		
	
		***See our collection for all versions of Llama 3.3 including GGUF, 4-bit and original 16-bit formats.***
	

We have a free Google Colab Tesla T4 notebook for Llama 3.1 (8B) here: https://colab.research.google.com/github/unslothai/notebooks/blob/main/nb/Llama3.1_(8B)-Alpaca.ipynb

For more details on the model, please go to Meta's original model card

All notebooks are **beginner friendly**! Add your dataset, click "Run All", and you'll get a 2x faster finetuned model which can be exported to GGUF, vLLM or uploaded to Hugging Face.

| Unsloth supports | Free Notebooks | Performance | Memory use | 
|---|---|---|---|
| **Llama-3.2 (3B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Llama-3.2 (11B vision)** | ▶️ Start on Colab | 2x faster | 60% less | 
| **Qwen2 VL (7B)** | ▶️ Start on Colab | 1.8x faster | 60% less | 
| **Qwen2.5 (7B)** | ▶️ Start on Colab | 2x faster | 60% less | 
| **Llama-3.1 (8B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Phi-3.5 (mini)** | ▶️ Start on Colab | 2x faster | 50% less | 
| **Gemma 2 (9B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Mistral (7B)** | ▶️ Start on Colab | 2.2x faster | 62% less | 

- This Llama 3.2 conversational notebook is useful for ShareGPT ChatML / Vicuna templates.
- This text completion notebook is for raw text. This DPO notebook replicates Zephyr.
- * Kaggle has 2x T4s, but we use 1. Due to overhead, 1x T4 is 5x faster.

A huge thank you to the Meta and Llama team for creating and releasing these models

The Meta Llama 3.3 multilingual large language model (LLM) is a pretrained and instruction tuned generative model in 70B (text in/text out). The Llama 3.3 instruction tuned text only model is optimized for multilingual dialogue use cases and outperform many of the available open source and closed chat models on common industry benchmarks.

**Model developer**: Meta

**Model Architecture:** Llama 3.3 is an auto-regressive language model that uses an optimized transformer architecture. The tuned versions use supervised fine-tuning (SFT) and reinforcement learning with human feedback (RLHF) to align with human preferences for helpfulness and safety. 

|  | Training Data | Params | Input modalities | Output modalities | Context length | GQA | Token count | Knowledge cutoff | 
|---|---|---|---|---|---|---|---|---|
| Llama 3.3 (text only) | A new mix of publicly available online data. | 70B | Multilingual Text | Multilingual Text and code | 128k | Yes | 15T+ | December 2023 | 

**Supported languages:** English, German, French, Italian, Portuguese, Hindi, Spanish, and Thai.

**Llama 3.3 model**. Token counts refer to pretraining data only. All model versions use Grouped-Query Attention (GQA) for improved inference scalability.

**Model Release Date:** 

- **70B Instruct: December 6, 2024**

**Status:** This is a static model trained on an offline dataset. Future versions of the tuned models will be released as we improve model safety with community feedback.

**License** A custom commercial license, the Llama 3.3 Community License Agreement, is available at: https://github.com/meta-llama/llama-models/blob/main/models/llama3_3/LICENSE

Where to send questions or comments about the model Instructions on how to provide feedback or comments on the model can be found in the model README. For more technical information about generation parameters and recipes for how to use Llama 3.3 in applications, please go here.

**Intended Use Cases** Llama 3.3 is intended for commercial and research use in multiple languages. Instruction tuned text only models are intended for assistant-like chat, whereas pretrained models can be adapted for a variety of natural language generation tasks. The Llama 3.3 model also supports the ability to leverage the outputs of its models to improve other models including synthetic data generation and distillation. The Llama 3.3 Community License allows for these use cases. 

**Out-of-scope** Use in any manner that violates applicable laws or regulations (including trade compliance laws). Use in any other way that is prohibited by the Acceptable Use Policy and Llama 3.3 Community License. Use in languages beyond those explicitly referenced as supported in this model card**.

**Note: Llama 3.3 has been trained on a broader collection of languages than the 8 supported languages. Developers may fine-tune Llama 3.3 models for languages beyond the 8 supported languages provided they comply with the Llama 3.3 Community License and the Acceptable Use Policy and in such cases are responsible for ensuring that any uses of Llama 3.3 in additional languages is done in a safe and responsible manner.

This repository contains two versions of Llama-3.3-70B-Instruct, for use with transformers and with the original `llama` codebase.

Starting with `transformers >= 4.43.0` onward, you can run conversational inference using the Transformers `pipeline` abstraction or by leveraging the Auto classes with the `generate()` function.

Make sure to update your transformers installation via `pip install --upgrade transformers`.

See the snippet below for usage with Transformers:

```
import transformers
import torch
model_id = "meta-llama/Llama-3.3-70B-Instruct"
pipeline = transformers.pipeline(
    "text-generation",
    model=model_id,
    model_kwargs={"torch_dtype": torch.bfloat16},
    device_map="auto",
)
messages = [
    {"role": "system", "content": "You are a pirate chatbot who always responds in pirate speak!"},
    {"role": "user", "content": "Who are you?"},
]
outputs = pipeline(
    messages,
    max_new_tokens=256,
)
print(outputs[0]["generated_text"][-1])
```
LLaMA-3.3 supports multiple tool use formats. You can see a full guide to prompt formatting here.

Tool use is also supported through chat templates in Transformers. Here is a quick example showing a single simple tool:

```
# First, define a tool
def get_current_temperature(location: str) -> float:
    """
    Get the current temperature at a location.
    
    Args:
        location: The location to get the temperature for, in the format "City, Country"
    Returns:
        The current temperature at the specified location in the specified units, as a float.
    """
    return 22.  # A real function should probably actually get the temperature!
# Next, create a chat and apply the chat template
messages = [
  {"role": "system", "content": "You are a bot that responds to weather queries."},
  {"role": "user", "content": "Hey, what's the temperature in Paris right now?"}
]
inputs = tokenizer.apply_chat_template(messages, tools=[get_current_temperature], add_generation_prompt=True)
```
You can then generate text from this input as normal. If the model generates a tool call, you should add it to the chat like so:

```
tool_call = {"name": "get_current_temperature", "arguments": {"location": "Paris, France"}}
messages.append({"role": "assistant", "tool_calls": [{"type": "function", "function": tool_call}]})
```
and then call the tool and append the result, with the `tool` role, like so:

```
messages.append({"role": "tool", "name": "get_current_temperature", "content": "22.0"})
```
After that, you can `generate()` again to let the model use the tool result in the chat. Note that this was a very brief introduction to tool calling - for more information,
see the LLaMA prompt format docs and the Transformers tool use documentation.

The model checkpoints can be used in `8-bit` and `4-bit` for further memory optimisations using `bitsandbytes` and `transformers`

See the snippet below for usage:

