---
id: collect-240926-huggingface/huggingface/unsloth-llama-4-maverick-17b-128e-instruct-gguf-hugging-face-1
title: "🦙 Run Unsloth Dynamic Llama 4 GGUF!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Meta", "Microsoft", "Mistral", "Unsloth", "vLLM"]
dates: ["2024-08", "2025-04-05"]
keywords: ["gguf", "llama", "decode", "distillation", "energy", "fine-tuning", "fp8", "gpu", "license", "llama.cpp", "memory", "mistral"]
source: docs/RAG/clean_en/huggingface/unsloth-llama-4-maverick-17b-128e-instruct-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 112]
sha256: fff2eac78b7d8991323bbf65bb8bb4fdba1df382a061d3925d8f8632ebb5b02a
---

# 🦙 Run Unsloth Dynamic Llama 4 GGUF!

<!-- source: https://huggingface.co/unsloth/Llama-4-Maverick-17B-128E-Instruct-GGUF -->

**See our collection for versions of Llama 4 including 4-bit & 16-bit formats.**
  

# 🦙 Run Unsloth Dynamic Llama 4 GGUF!

    *Read our Guide to see how to Fine-tune & Run Llama 4 correctly.*
  

| MoE Bits | Type | Disk Size | HF Link | Accuracy | 
|---|---|---|---|---|
| 1.78bit | IQ1_S | **122GB** | Link | Ok | 
| 1.93bit | IQ1_M | **128GB** | Link | Fair | 
| 2.42-bit | IQ2_XXS | **140GB** | Link | Better | 
| 2.71-bit | Q2_K_XL | **151B** | Link | Suggested | 
| 3.5-bit | Q3_K_XL | **193GB** | Link | Great | 
| 4.5-bit | Q4_K_XL | **243GB** | Link | Best | 

Currently text only is supported.

**Chat template/prompt format:**

```
<|header_start|>user<|header_end|>\n\nWhat is 1+1?<|eot|><|header_start|>assistant<|header_end|>\n\n
```
- Fine-tune Llama-4-Scout on a single H100 80GB GPU using Unsloth!
- Read our Blog about Llama 4 support: unsloth.ai/blog/llama4
- View the rest of our notebooks in our docs here.
- Export your fine-tuned model to GGUF, Ollama, llama.cpp, vLLM or 🤗HF.

| Unsloth supports | Free Notebooks | Performance | Memory use | 
|---|---|---|---|
| **GRPO with Llama 3.1 (8B)** | ▶️ Start on Colab | 2x faster | 80% less | 
| **Llama-3.2 (3B)** | ▶️ Start on Colab | 2.4x faster | 58% less | 
| **Llama-3.2 (11B vision)** | ▶️ Start on Colab | 2x faster | 60% less | 
| **Qwen2.5 (7B)** | ▶️ Start on Colab | 2x faster | 60% less | 
| **Phi-4 (14B)** | ▶️ Start on Colab | 2x faster | 50% less | 
| **Mistral (7B)** | ▶️ Start on Colab | 2.2x faster | 62% less | 

The Llama 4 collection of models are natively multimodal AI models that enable text and multimodal experiences. These models leverage a mixture-of-experts architecture to offer industry-leading performance in text and image understanding.

These Llama 4 models mark the beginning of a new era for the Llama ecosystem. We are launching two efficient models in the Llama 4 series, Llama 4 Scout, a 17 billion parameter model with 16 experts, and Llama 4 Maverick, a 17 billion parameter model with 128 experts.

**Model developer**: Meta

**Model Architecture:**  The Llama 4 models are auto-regressive language models that use a mixture-of-experts (MoE) architecture and incorporate early fusion for native multimodality. 

| Model Name | Training Data | Params | Input modalities | Output modalities | Context length | Token count | Knowledge cutoff | 
|---|---|---|---|---|---|---|---|
| Llama 4 Scout (17Bx16E) | A mix of publicly available, licensed data and information from Meta's products and services. This includes publicly shared posts from Instagram and Facebook and people's interactions with Meta AI. Learn more in our Privacy Center. | 17B (Activated) 109B (Total) | Multilingual text and image | Multilingual text and code | 10M | ~40T | August 2024 | 
| Llama 4 Maverick (17Bx128E) |  | 17B (Activated) 400B (Total) | Multilingual text and image | Multilingual text and code | 1M | ~22T | August 2024 | 

**Supported languages:** Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, and Vietnamese. 

**Model Release Date:** April 5, 2025

**Status:** This is a static model trained on an offline dataset. Future versions of the tuned models may be released as we improve model behavior with community feedback.

**License**: A custom commercial license, the Llama 4 Community License Agreement, is available at: https://github.com/meta-llama/llama-models/blob/main/models/llama4/LICENSE

**Where to send questions or comments about the model:** Instructions on how to provide feedback or comments on the model can be found in the Llama README. For more technical information about generation parameters and recipes for how to use Llama 4 in applications, please go here.

Please, make sure you have transformers `v4.51.0` installed, or upgrade using `pip install -U transformers`.

```
from transformers import AutoTokenizer, Llama4ForConditionalGeneration
import torch
model_id = "meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8"
tokenizer = AutoTokenizer.from_pretrained(model_id)
messages = [
    {"role": "user", "content": "Who are you?"},
]
inputs = tokenizer.apply_chat_template(messages, add_generation_prompt=True, return_tensors="pt", return_dict=True)
model = Llama4ForConditionalGeneration.from_pretrained(
    model_id,
    tp_plan="auto",
    torch_dtype="auto",
)
outputs = model.generate(**inputs.to(model.device), max_new_tokens=100)
outputs = tokenizer.batch_decode(outputs[:, inputs["input_ids"].shape[-1]:])
print(outputs[0])
```
**Intended Use Cases:** Llama 4 is intended for commercial and research use in multiple languages. Instruction tuned models are intended for assistant-like chat and visual reasoning tasks, whereas pretrained models can be adapted for natural language generation. For vision, Llama 4 models are also optimized for visual recognition, image reasoning, captioning, and answering general questions about an image. The Llama 4 model collection also supports the ability to leverage the outputs of its models to improve other models including synthetic data generation and distillation. The Llama 4 Community License allows for these use cases. 

**Out-of-scope**: Use in any manner that violates applicable laws or regulations (including trade compliance laws). Use in any other way that is prohibited by the Acceptable Use Policy and Llama 4 Community License. Use in languages or capabilities beyond those explicitly referenced as supported in this model card**.

**Note:

1. Llama 4 has been trained on a broader collection of languages than the 12 supported languages (pre-training includes 200 total languages). Developers may fine-tune Llama 4 models for languages beyond the 12 supported languages provided they comply with the Llama 4 Community License and the Acceptable Use Policy. Developers are responsible for ensuring that their use of Llama 4 in additional languages is done in a safe and responsible manner.

2. Llama 4 has been tested for image understanding up to 5 input images. If leveraging additional image understanding capabilities beyond this, Developers are responsible for ensuring that their deployments are mitigated for risks and should perform additional testing and tuning tailored to their specific applications.

**Training Factors:** We used custom training libraries, Meta's custom built GPU clusters, and production infrastructure for pretraining. Fine-tuning, quantization, annotation, and evaluation were also performed on production infrastructure.

**Training Energy Use:**  Model pre-training utilized a cumulative of **7.38M** GPU hours of computation on H100-80GB (TDP of 700W) type hardware, per the table below. Training time is the total GPU time required for training each model and power consumption is the peak power capacity per GPU device used, adjusted for power usage efficiency. 

## 
	
		
	
		**Training Greenhouse Gas Emissions:** Estimated total location-based greenhouse gas emissions were **1,999 tons** CO2eq for training. Since 2020, Meta has maintained net zero greenhouse gas emissions in its global operations and matched 100% of its electricity use with clean and renewable energy; therefore, the total market-based greenhouse gas emissions for training were 0 tons CO2eq.
	

| Model Name | Training Time (GPU hours) | Training Power Consumption (W) | Training Location-Based Greenhouse Gas Emissions (tons CO2eq) | Training Market-Based Greenhouse Gas Emissions (tons CO2eq) | 
|---|---|---|---|---|
| Llama 4 Scout | 5.0M | 700 | 1,354 | 0 | 
| Llama 4 Maverick | 2.38M | 700 | 645 | 0 | 
| Total | 7.38M | - | 1,999 | 0 | 

## The methodology used to determine training energy use and greenhouse gas emissions can be found here. Since Meta is openly releasing these models, the training energy use and greenhouse gas emissions will not be incurred by others.

