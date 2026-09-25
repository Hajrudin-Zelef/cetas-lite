---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-guide
title: "Fine-tuning LLMs Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "OpenAI", "Unsloth", "vLLM"]
dates: []
keywords: ["fine-tuning", "agent", "agents", "awq", "distillation", "dpo", "fp8", "gguf", "inference", "inference engine", "llama", "llama.cpp"]
source: docs/RAG/clean4/Fine-tuning LLMs Guide.md
source_anchor: ""
source_lines: [1, 71]
sha256: 9650c9ea7afd50a3c3746b9405fd2d96380fff922f44872579d539ba33ca0d65
---

# Fine-tuning LLMs Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/get-started/fine-tuning-llms-guide.md).
# Fine-tuning LLMs Guide
Learn all the basics and best practices of fine-tuning. Beginner-friendly.
## 1. What Is Fine-tuning?
Fine-tuning / training / post-training models customizes its behavior, enhances + injects knowledge, and optimizes performance for domains and specific tasks. For example:
* OpenAI’s **GPT-5** was post-trained to improve instruction following and helpful chat behavior.
* The standard method of post training is called Supervised Fine-Tuning (SFT). Other methods include preference optimization (DPO, ORPO), distillation and [Reinforcement Learning (RL)](/docs/get-started/reinforcement-learning-rl-guide.md) (GRPO, GSPO), where an "agent" learns to make decisions by interacting with an environment and receiving **feedback** in the form of **rewards** or **penalties**.
With [Unsloth](https://github.com/unslothai/unsloth), you can fine-tune or do RL for free on Colab, Kaggle, or locally with just 3GB VRAM by using our [notebooks](https://docs.unsloth.ai/get-started/unsloth-notebooks). By fine-tuning a pre-trained model on a dataset, you can:
* **Update + Learn New Knowledge**: Inject and learn new domain-specific information.
* **Customize Behavior**: Adjust the model’s tone, personality, or response style.
* **Optimize for Tasks**: Improve accuracy and relevance for specific use cases.
**Example fine-tuning or RL use-cases**:
* Enables LLMs to predict if a headline impacts a company positively or negatively.
* Can use historical customer interactions for more accurate and custom responses.
* Fine-tune LLM on legal texts for contract analysis, case law research, and compliance.
You can think of a fine-tuned model as a specialized agent designed to do specific tasks more effectively and efficiently. **Fine-tuning can replicate all of RAG's capabilities**, but not vice versa.
{% columns %}
{% column %}
#### :question:What is LoRA/QLoRA?
In LLMs, we have model weights. Llama 70B has 70 billion numbers. Instead of changing all 70B numbers, we instead add thin matrices A and B to each weight, and optimize those. This means we only optimize 1% of weights. LoRA is when the original model is 16-bit unquantized while QLoRA quantizes to 4-bit to save 75% memory.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### Fine-tuning misconceptions:
You may have heard that fine-tuning does not make a model learn new knowledge or RAG performs better than fine-tuning. That is **false**. You can train a specialized coding model with fine-tuning and RL while RAG can’t change the model’s weights and only augments what the model sees at inference time. Read more FAQ + misconceptions [here](https://unsloth.ai/docs/get-started/pages/HP82bIzgldwxWk3OSzVy#fine-tuning-vs.-rag-whats-the-difference):
{% content-ref url="/pages/HP82bIzgldwxWk3OSzVy" %}
[FAQ + Is Fine-tuning Right For Me?](/docs/get-started/fine-tuning-for-beginners/faq-+-is-fine-tuning-right-for-me.md)
{% endcontent-ref %}
> [**Introducing Unsloth Studio:** ](/docs/new/studio.md) Our new open-source web UI for training and running models. This means you can now fine-tune models with no-code and have observability and automatic dataset creation features.

Reminder Unsloth itself provides **2x faster inference** natively as well, so always do not forget to call `FastLanguageModel.for_inference(model)`. If you want the model to output longer responses, set `max_new_tokens = 128` to some larger number like 256 or 1024. Notice you will have to wait longer for the result as well!
### Saving + Deployment
For saving and deploying your model in desired inference engines like Ollama, vLLM, Open WebUI, you will need to use the LoRA adapter on top of the base model. We have designated guides for each framework:
{% content-ref url="/pages/gEugERiAw2ztDNt98JVR" %}
[Inference & Deployment](/docs/basics/inference-and-deployment.md)
{% endcontent-ref %}
{% columns %}
{% column %}
If you’re running inference on a single device (like a laptop or Mac), use llama.cpp to convert to GGUF format to use in Ollama, llama.cpp, LM Studio etc:
{% content-ref url="/pages/T7ZPf3SNAwDykZNgXptE" %}
[GGUF & llama.cpp](/docs/basics/inference-and-deployment/saving-to-gguf.md)
{% endcontent-ref %}
{% endcolumn %}
{% column %}
If you’re deploying an LLM for enterprise or multi-user inference for FP8, AWQ, use vLLM:
{% content-ref url="/pages/fhJtaLFFXVsGnbMUiACo" %}
[vLLM](/docs/basics/inference-and-deployment/vllm-guide.md)
{% endcontent-ref %}
{% endcolumn %}
{% endcolumns %}
We can now save the fine-tuned model as a small 100MB file called a LoRA adapter like below. You can instead push to the Hugging Face hub as well if you want to upload your model! Remember to get a Hugging Face [token](https://huggingface.co/settings/tokens) and add your token!

After saving the model, we can again use Unsloth to run the model itself! Use `FastLanguageModel` again to call it for inference!
## 8. We're done!
You've successfully fine-tuned a language model and exported it to your desired inference engine with Unsloth!
To learn more about fine-tuning tips and tricks, head over to our blogs which provide tremendous and educational value:
If you need any help on fine-tuning, you can also join our Discord server [here](https://discord.gg/unsloth) or [Reddit r/unsloth](https://www.reddit.com/r/unsloth/). Thanks for reading and hopefully this was helpful!
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/fine-tuning-llms-guide.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
