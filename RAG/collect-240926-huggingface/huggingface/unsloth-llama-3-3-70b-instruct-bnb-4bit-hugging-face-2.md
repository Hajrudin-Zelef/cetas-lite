---
id: collect-240926-huggingface/huggingface/unsloth-llama-3-3-70b-instruct-bnb-4bit-hugging-face-2
title: "First, define a tool"
domain: huggingface
role: reference
task: reference
actors: ["Meta"]
dates: ["2023-12"]
keywords: ["agentic", "alignment", "benchmark", "decode", "energy", "exploit", "fine-tuning", "gpu", "guardrails", "llama", "pretraining", "quantization"]
source: docs/RAG/clean_en/huggingface/unsloth-llama-3-3-70b-instruct-bnb-4bit-hugging-face.md
source_anchor: ""
source_lines: [132, 208]
sha256: 7ed88baa30c38c0cf10a4e43c1de174c1ba49dfb5f69e5597b4098e86d1312c1
---

# First, define a tool

```
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
model_id = "meta-llama/Llama-3.3-70B-Instruct"
quantization_config = BitsAndBytesConfig(load_in_8bit=True)
quantized_model = AutoModelForCausalLM.from_pretrained(
    model_id, device_map="auto", torch_dtype=torch.bfloat16, quantization_config=quantization_config)
tokenizer = AutoTokenizer.from_pretrained(model_id)
input_text = "What are we having for dinner?"
input_ids = tokenizer(input_text, return_tensors="pt").to("cuda")
output = quantized_model.generate(**input_ids, max_new_tokens=10)
print(tokenizer.decode(output[0], skip_special_tokens=True))
```
To load in 4-bit simply pass `load_in_4bit=True`

Please, follow the instructions in the repository.

To download Original checkpoints, see the example command below leveraging `huggingface-cli`:

```
huggingface-cli download meta-llama/Llama-3.3-70B-Instruct --include "original/*" --local-dir Llama-3.3-70B-Instruct
```
**Training Factors** We used custom training libraries, Meta's custom built GPU cluster, and production infrastructure for pretraining. Fine-tuning, annotation, and evaluation were also performed on production infrastructure.

**Training Energy Use** Training utilized a cumulative of **39.3**M GPU hours of computation on H100-80GB (TDP of 700W) type hardware, per the table below. Training time is the total GPU time required for training each model and power consumption is the peak power capacity per GPU device used, adjusted for power usage efficiency. 

## 
	
		
	
		**Training Greenhouse Gas Emissions** Estimated total location-based greenhouse gas emissions were **11,390** tons CO2eq for training. Since 2020, Meta has maintained net zero greenhouse gas emissions in its global operations and matched 100% of its electricity use with renewable energy, therefore the total market-based greenhouse gas emissions for training were 0 tons CO2eq.
	

|  | Training Time (GPU hours) | Training Power Consumption (W) | Training Location-Based Greenhouse Gas Emissions (tons CO2eq) | Training Market-Based Greenhouse Gas Emissions (tons CO2eq) | 
|---|---|---|---|---|
| Llama 3.3 70B | 7.0M | 700 | 2,040 | 0 | 

## The methodology used to determine training energy use and greenhouse gas emissions can be found here. Since Meta is openly releasing these models, the training energy use and greenhouse gas emissions will not be incurred by others.

**Overview:** Llama 3.3 was pretrained on ~15 trillion tokens of data from publicly available sources. The fine-tuning data includes publicly available instruction datasets, as well as over 25M synthetically generated examples. 

**Data Freshness:** The pretraining data has a cutoff of December 2023.

In this section, we report the results for Llama 3.3 relative to our previous models.

| Category | Benchmark | # Shots | Metric | Llama 3.1 8B Instruct | Llama 3.1 70B Instruct | Llama-3.3 70B Instruct | Llama 3.1 405B Instruct | 
|---|---|---|---|---|---|---|---|
|  | MMLU (CoT) | 0 | macro_avg/acc | 73.0 | 86.0 | 86.0 | 88.6 | 
|  | MMLU Pro (CoT) | 5 | macro_avg/acc | 48.3 | 66.4 | 68.9 | 73.3 | 
| Steerability | IFEval |  |  | 80.4 | 87.5 | 92.1 | 88.6 | 
| Reasoning | GPQA Diamond (CoT) | 0 | acc | 31.8 | 48.0 | 50.5 | 49.0 | 
| Code | HumanEval | 0 | pass@1 | 72.6 | 80.5 | 88.4 | 89.0 | 
|  | MBPP EvalPlus (base) | 0 | pass@1 | 72.8 | 86.0 | 87.6 | 88.6 | 
| Math | MATH (CoT) | 0 | sympy_intersection_score | 51.9 | 68.0 | 77.0 | 73.8 | 
| Tool Use | BFCL v2 | 0 | overall_ast_summary/macro_avg/valid | 65.4 | 77.5 | 77.3 | 81.1 | 
| Multilingual | MGSM | 0 | em | 68.9 | 86.9 | 91.1 | 91.6 | 

As part of our Responsible release approach, we followed a three-pronged strategy to managing trust & safety risks:

- Enable developers to deploy helpful, safe and flexible experiences for their target audience and for the use cases supported by Llama.
- Protect developers against adversarial users aiming to exploit Llama capabilities to potentially cause harm.
- Provide protections for the community to help prevent the misuse of our models.

Llama is a foundational technology designed to be used in a variety of use cases, examples on how Meta’s Llama models have been responsibly deployed can be found in our Community Stories webpage. Our approach is to build the most helpful models enabling the world to benefit from the technology power, by aligning our model safety for the generic use cases addressing a standard set of harms. Developers are then in the driver seat to tailor safety for their use case, defining their own policy and deploying the models with the necessary safeguards in their Llama systems. Llama 3.3 was developed following the best practices outlined in our Responsible Use Guide, you can refer to the Responsible Use Guide to learn more.

Our main objectives for conducting safety fine-tuning are to provide the research community with a valuable resource for studying the robustness of safety fine-tuning, as well as to offer developers a readily available, safe, and powerful model for various applications to reduce the developer workload to deploy safe AI systems. For more details on the safety mitigations implemented please read the Llama 3 paper. 
**Fine-tuning data**

We employ a multi-faceted approach to data collection, combining human-generated data from our vendors with synthetic data to mitigate potential safety risks. We’ve developed many large language model (LLM)-based classifiers that enable us to thoughtfully select high-quality prompts and responses, enhancing data quality control. 
**Refusals and Tone**

Building on the work we started with Llama 3, we put a great emphasis on model refusals to benign prompts as well as refusal tone. We included both borderline and adversarial prompts in our safety data strategy, and modified our safety data responses to follow  tone guidelines. 

**Large language models, including Llama 3.3, are not designed to be deployed in isolation but instead should be deployed as part of an overall AI system with additional safety guardrails as required.** Developers are expected to deploy system safeguards when building agentic systems. Safeguards are key to achieve the right helpfulness-safety alignment as well as mitigating safety and security risks inherent to the system and any integration of the model or system with external tools.

As part of our responsible release approach, we provide the community with safeguards that developers should deploy with Llama models or other LLMs, including Llama Guard 3, Prompt Guard and Code Shield. All our reference implementations demos contain these safeguards by default so developers can benefit from system-level safety out-of-the-box. 

