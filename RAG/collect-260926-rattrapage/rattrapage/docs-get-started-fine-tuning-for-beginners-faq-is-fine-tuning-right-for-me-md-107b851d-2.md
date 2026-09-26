---
id: collect-260926-rattrapage/rattrapage/docs-get-started-fine-tuning-for-beginners-faq-is-fine-tuning-right-for-me-md-107b851d-2
title: "FAQ + Is Fine-tuning Right For Me?"
domain: rattrapage
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["fine-tuning", "agent", "agents", "compute", "cost", "lora", "memory", "parameters", "qlora", "quantization", "training"]
source: docs/RAG/lot-rattrapage/fine-tuning/docs-get-started-fine-tuning-for-beginners-faq-is-fine-tuning-right-for-me-md-107b851d.md
source_anchor: ""
source_lines: [46, 66]
sha256: b96fd369f5d0ba9f3bc79f3c389e64a75033cb7214a68cb385fe3ff6d5a8eb67
---

# FAQ + Is Fine-tuning Right For Me?

* **LoRA (Low-Rank Adaptation)** – Fine-tunes only a small set of additional “adapter” weight matrices (in 16-bit precision), while leaving most of the original model unchanged. This significantly reduces the number of parameters that need updating during training.
* **QLoRA (Quantized LoRA)** – Combines LoRA with 4-bit quantization of the model weights, enabling efficient fine-tuning of very large models on minimal hardware. By using 4-bit precision where possible, it dramatically lowers memory usage and compute overhead.
We recommend starting with **QLoRA**, as it’s one of the most efficient and accessible methods available. Thanks to Unsloth’s [dynamic 4-bit](https://unsloth.ai/blog/dynamic-4bit) quants, the accuracy loss compared to standard 16-bit LoRA fine-tuning is now negligible.
### Experimentation is Key
There’s no single “best” approach to fine-tuning - only best practices for different scenarios. It’s important to experiment with different methods and configurations to find what works best for your dataset and use case. A great starting point is **QLoRA (4-bit)**, which offers a very cost-effective, resource-friendly way to fine-tune models without heavy computational requirements.
{% content-ref url="/pages/y6obKRSk8TwyjIrCjuGE" %}
[Hyperparameters Guide](/docs/get-started/fine-tuning-llms-guide/lora-hyperparameters-guide.md)
{% endcontent-ref %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/fine-tuning-for-beginners/faq-+-is-fine-tuning-right-for-me.md?ask=<question>&goal=<endgoal>
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
