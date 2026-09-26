---
id: collect-240926-huggingface/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face-2
title: "meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["California", "Meta", "Microsoft"]
dates: ["2024-08", "2025-04-05"]
keywords: ["llama", "attention", "decode", "distillation", "energy", "fine-tuning", "gpu", "license", "moe", "multimodal", "parameters", "pretraining"]
source: docs/RAG/clean_en/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face.md
source_anchor: ""
source_lines: [53, 152]
sha256: d141222d978942dc48171193da1f50446815707675df9cf782c2153f89e25e4b
---

# meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face

7. **Governing Law and Jurisdiction**. This Agreement will be governed and construed under the laws of the State of California without regard to choice of law principles, and the UN Convention on Contracts for the International Sale of Goods does not apply to this Agreement. The courts of California shall have exclusive jurisdiction of any dispute arising out of this Agreement.

Log in or Sign Up to review the conditions and access this model content.

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

**Intended Use Cases:** Llama 4 is intended for commercial and research use in multiple languages. Instruction tuned models are intended for assistant-like chat and visual reasoning tasks, whereas pretrained models can be adapted for natural language generation. For vision, Llama 4 models are also optimized for visual recognition, image reasoning, captioning, and answering general questions about an image. The Llama 4 model collection also supports the ability to leverage the outputs of its models to improve other models including synthetic data generation and distillation. The Llama 4 Community License allows for these use cases. 

**Out-of-scope**: Use in any manner that violates applicable laws or regulations (including trade compliance laws). Use in any other way that is prohibited by the Acceptable Use Policy and Llama 4 Community License. Use in languages or capabilities beyond those explicitly referenced as supported in this model card**.

**Note:

1. Llama 4 has been trained on a broader collection of languages than the 12 supported languages (pre-training includes 200 total languages). Developers may fine-tune Llama 4 models for languages beyond the 12 supported languages provided they comply with the Llama 4 Community License and the Acceptable Use Policy. Developers are responsible for ensuring that their use of Llama 4 in additional languages is done in a safe and responsible manner.

2. Llama 4 has been tested for image understanding up to 5 input images. If leveraging additional image understanding capabilities beyond this, Developers are responsible for ensuring that their deployments are mitigated for risks and should perform additional testing and tuning tailored to their specific applications.

Please, make sure you have transformers `v4.51.0` installed, or upgrade using `pip install -U transformers`.

```
from transformers import AutoProcessor, Llama4ForConditionalGeneration
import torch
model_id = "meta-llama/Llama-4-Maverick-17B-128E-Instruct"
processor = AutoProcessor.from_pretrained(model_id)
model = Llama4ForConditionalGeneration.from_pretrained(
    model_id,
    attn_implementation="flex_attention",
    device_map="auto",
    torch_dtype=torch.bfloat16,
)
url1 = "https://huggingface.co/datasets/huggingface/documentation-images/resolve/0052a70beed5bf71b92610a43a52df6d286cd5f3/diffusers/rabbit.jpg"
url2 = "https://huggingface.co/datasets/huggingface/documentation-images/resolve/main/datasets/cat_style_layout.png"
messages = [
    {
        "role": "user",
        "content": [
            {"type": "image", "url": url1},
            {"type": "image", "url": url2},
            {"type": "text", "text": "Can you describe how these two images are similar, and how they differ?"},
        ]
    },
]
inputs = processor.apply_chat_template(
    messages,
    add_generation_prompt=True,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
).to(model.device)
outputs = model.generate(
    **inputs,
    max_new_tokens=256,
)
response = processor.batch_decode(outputs[:, inputs["input_ids"].shape[-1]:])[0]
print(response)
print(outputs[0])
```
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

**Overview:** Llama 4 Scout was pretrained on ~40 trillion tokens and Llama 4 Maverick was pretrained on ~22 trillion tokens of multimodal data from a mix of publicly available, licensed data and information from Meta’s products and services. This includes publicly shared posts from Instagram and Facebook and people’s interactions with Meta AI.

**Data Freshness:** The pretraining data has a cutoff of August 2024.

