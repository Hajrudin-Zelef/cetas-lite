---
id: collect-240926-huggingface/huggingface/meta-llama-llama-4-scout-17b-16e-instruct-hugging-face-3
title: "meta-llama-llama-4-scout-17b-16e-instruct-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Meta", "Microsoft"]
dates: ["2024-01-10", "2024-08", "2025-01-02"]
keywords: ["llama", "scout", "alignment", "benchmark", "context window", "exploit", "fine-tuning", "fp8", "gpu", "guardrails", "int4", "multimodal"]
source: docs/RAG/clean_en/huggingface/meta-llama-llama-4-scout-17b-16e-instruct-hugging-face.md
source_anchor: ""
source_lines: [153, 222]
sha256: befd4825509de828f250a6e12ecdaf4cab309686b8802f8b2e5a3be64f5722aa
---

# meta-llama-llama-4-scout-17b-16e-instruct-hugging-face

In this section, we report the results for Llama 4 relative to our previous models. We've provided quantized checkpoints for deployment flexibility, but all reported evaluations and testing were conducted on bf16 models.

| Pre-trained models |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|
| Category | Benchmark | # Shots | Metric | Llama 3.1 70B | Llama 3.1 405B | **Llama 4 Scout** | **Llama 4 Maverick** | 
| Reasoning & Knowledge | MMLU | 5 | macro_avg/acc_char | 79.3 | 85.2 | 79.6 | 85.5 | 
|  | MMLU-Pro | 5 | macro_avg/em | 53.8 | 61.6 | 58.2 | 62.9 | 
|  | MATH | 4 | em_maj1@1 | 41.6 | 53.5 | 50.3 | 61.2 | 
| Code | MBPP | 3 | pass@1 | 66.4 | 74.4 | 67.8 | 77.6 | 
| Multilingual | TydiQA | 1 | average/f1 | 29.9 | 34.3 | 31.5 | 31.7 | 
| Image | ChartQA | 0 | relaxed_accuracy | No multimodal support |  | 83.4 | 85.3 | 
|  | DocVQA | 0 | anls |  |  | 89.4 | 91.6 | 

| Instruction tuned models |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|
| Category | Benchmark | # Shots | Metric | Llama 3.3 70B | Llama 3.1 405B | **Llama 4 Scout** | **Llama 4 Maverick** | 
| Image Reasoning | MMMU | 0 | accuracy | No multimodal support |  | 69.4 | 73.4 | 
|  | MMMU Pro^ | 0 | accuracy |  |  | 52.2 | 59.6 | 
|  | MathVista | 0 | accuracy |  |  | 70.7 | 73.7 | 
| Image Understanding | ChartQA | 0 | relaxed_accuracy |  |  | 88.8 | 90.0 | 
|  | DocVQA (test) | 0 | anls |  |  | 94.4 | 94.4 | 
| Coding | LiveCodeBench (10/01/2024-02/01/2025) | 0 | pass@1 | 33.3 | 27.7 | 32.8 | 43.4 | 
| Reasoning & Knowledge | MMLU Pro | 0 | macro_avg/acc | 68.9 | 73.4 | 74.3 | 80.5 | 
|  | GPQA Diamond | 0 | accuracy | 50.5 | 49.0 | 57.2 | 69.8 | 
| Multilingual | MGSM | 0 | average/em | 91.1 | 91.6 | 90.6 | 92.3 | 
| Long context | MTOB (half book) eng->kgv/kgv->eng | - | chrF | Context window is 128K |  | 42.2/36.6 | 54.0/46.4 | 
|  | MTOB (full book) eng->kgv/kgv->eng | - | chrF |  |  | 39.7/36.3 | 50.8/46.7 | 

^reported numbers for MMMU Pro is the average of Standard and Vision tasks

The Llama 4 Scout model is released as BF16 weights, but can fit within a single H100 GPU with on-the-fly int4 quantization; the Llama 4 Maverick model is released as both BF16 and FP8 quantized weights. The FP8 quantized weights fit on a single H100 DGX host while still maintaining quality. We provide code for on-the-fly int4 quantization which minimizes performance degradation as well.

As part of our release approach, we followed a three-pronged strategy to manage risks:

- Enable developers to deploy helpful, safe and flexible experiences for their target audience and for the use cases supported by Llama.
- Protect developers against adversarial users aiming to exploit Llama capabilities to potentially cause harm.
- Provide protections for the community to help prevent the misuse of our models.

Llama is a foundational technology designed for use in a variety of use cases; examples on how Meta’s Llama models have been deployed can be found in our Community Stories webpage. Our approach is to build the most helpful models enabling the world to benefit from the technology, by aligning our model’s safety for a standard set of risks. Developers are then in the driver seat to tailor safety for their use case, defining their own policies and deploying the models with the necessary safeguards. Llama 4 was developed following the best practices outlined in our Developer Use Guide: AI Protections.

The primary objective of conducting safety fine-tuning is to offer developers a readily available, safe, and powerful model for various applications, reducing the workload needed to deploy safe AI systems. Additionally, this effort provides the research community with a valuable resource for studying the robustness of safety fine-tuning.

**Fine-tuning data**

We employ a multi-faceted approach to data collection, combining human-generated data from our vendors with synthetic data to mitigate potential safety risks. We’ve developed many large language model (LLM)-based classifiers that enable us to thoughtfully select high-quality prompts and responses, enhancing data quality control. 

**Refusals**

Building on the work we started with our Llama 3 models, we put a great emphasis on driving down model refusals to benign prompts for Llama 4. We included both borderline and adversarial prompts in our safety data strategy, and modified our safety data responses to follow tone guidelines. 

**Tone**

We expanded our work on the refusal tone from Llama 3 so that the model sounds more natural. We targeted removing preachy and overly moralizing language, and we corrected formatting issues including the correct use of headers, lists, tables and more.

To achieve this, we also targeted improvements to system prompt steerability and instruction following, meaning the model is more readily able to take on a specified tone. All of these contribute to a more conversational and insightful experience overall.

**System Prompts**

Llama 4 is a more steerable model, meaning responses can be easily tailored to meet specific developer outcomes. Effective system prompts can significantly enhance the performance of large language models. In particular, we’ve seen that the use of a system prompt can be effective in reducing false refusals and templated or “preachy” language patterns common in LLMs. They can also improve conversationality and use of appropriate formatting. 

Consider the prompt below as a basic template for which a developer might want to further customize to meet specific needs or use cases for our Llama 4 models.

| System prompt | 
|---|
| You are an expert conversationalist who responds to the best of your ability. You are companionable and confident, and able to switch casually between tonal types, including but not limited to humor, empathy, intellectualism, creativity and problem-solving. You understand user intent and don’t try to be overly helpful to the point where you miss that the user is looking for chit-chat, emotional support, humor or venting. Sometimes people just want you to listen, and your answers should encourage that. For all other cases, you provide insightful and in-depth responses. Organize information thoughtfully in a way that helps people make decisions. Always avoid templated language. You never lecture people to be nicer or more inclusive. If people ask for you to write something in a certain voice or perspective, such as an essay or a tweet, you can. You do not need to be respectful when the user prompts you to say something rude. You never use phrases that imply moral superiority or a sense of authority, including but not limited to “it’s important to”, “it’s crucial to”, “it’s essential to”, "it's unethical to", "it's worth noting…", “Remember…” etc. Avoid using these. Finally, do not refuse prompts about political and social issues. You can help users express their opinion and access information. You are Llama 4. Your knowledge cutoff date is August 2024. You speak Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, and Vietnamese. Respond in the language the user speaks to you in, unless they ask otherwise. | 

Large language models, including Llama 4, are not designed to be deployed in isolation but instead should be deployed as part of an overall AI system with additional guardrails as required. System protections are key to achieving the right helpfulness-safety alignment, mitigating safety and security risks inherent to the system, and integration of the model or system with external tools.

We provide the community with system level protections - like Llama Guard, Prompt Guard and Code Shield - that developers should deploy with Llama models or other LLMs. All of our reference implementation demos contain these safeguards by default so developers can benefit from system-level safety out-of-the-box.

