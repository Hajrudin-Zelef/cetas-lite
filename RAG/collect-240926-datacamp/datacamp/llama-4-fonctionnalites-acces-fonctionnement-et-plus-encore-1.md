---
id: collect-240926-datacamp/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore-1
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Meta", "Microsoft", "OpenAI"]
dates: []
keywords: ["llama", "benchmark", "benchmarks", "context window", "deepseek", "distillation", "fine-tuning", "gemini", "gpu", "inference", "license", "moe"]
source: docs/RAG/clean_en/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [1, 71]
sha256: 01e3ad96fdc636e1c9d5be7613784855599e2342991ed5796f3601d50fd5c98c
---

# llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore

<!-- source: https://www.datacamp.com/fr/blog/llama-4 -->

Course

Meta has just announced the Llama 4 suite of models, which includes two models already available — Llama 4 Scout and Llama 4 Maverick — and a third still in training: Llama 4 Behemoth.

The Scout and Maverick variants are available now, openly released under Meta's usual open-weight license — with a notable caveat: if your services exceed 700 million monthly active users, you must obtain a separate license from Meta, which may or may not be granted at its discretion.

Llama Scout supports a context window of 10 million tokens, the largest of any publicly released model. Llama Maverick is a general-purpose model that targets GPT-4o, Gemini 2.0 Flash, and DeepSeek-V3. Llama Behemoth, still in training, serves as a high-capacity teacher model.

In this introductory article, I offer an overview of the Llama 4 suite. Our team has already tested the model, and I recommend these tutorials if you want to go further:

**Update:** after about a year of silence, Meta has released its latest LLM. For an overview, see our guide on Muse Spark.

We keep our readers informed of the latest AI advances via *The Median*, our free Friday newsletter that decodes the week's key news. Subscribe and stay up to date in a few minutes per week:


## What is Llama 4?

Llama 4 is Meta's new family of large language models. The release includes two models already available: Llama 4 Scout and Llama 4 Maverick; and a third, Llama 4 Behemoth, still in training.

Source: Meta AI

Llama 4 introduces substantial improvements. In particular, it incorporates a mixture-of-experts (MoE) architecture, which aims to improve efficiency and performance by activating only the components needed for a given task (we'll come back to this shortly). This design marks a shift toward more scalable and specialized AI models.

Llama 4 extends Meta's strategy of releasing open-weight models — but with a caveat. If your company operates services totaling more than 700 million monthly active users, you will need to obtain a separate license from Meta, which may not be granted. Despite this constraint, the release remains a major event in the open-weight ecosystem, even though that ecosystem has evolved very rapidly in recent months.

If Llama 2 and 3 once defined the category, Llama 4 now arrives in a far more competitive field. DeepSeek has established itself with strong reasoning capabilities. Alibaba's Qwen series stands out on multilingual and code benchmarks. Google's Gemma models advance on the same ground with smaller, more efficient architectures. And very recently, OpenAI announced its intention to release an open-weight model, an inflection unthinkable a year ago.

Let's now look at the details of each model.

## Llama Scout

Llama 4 Scout is the lightest model in the new suite, but it is arguably the most intriguing. It runs on a single H100 GPU and supports a context window of 10 million tokens. This makes it the most context-hungry open-weight model to date and potentially the most useful for tasks such as multi-document summarization, reasoning over long code, and activity analysis.

Scout has 17 billion active parameters, organized via 16 experts, for a total of about 109 billion parameters. It was pre-trained and post-trained with a context window of 256 K, but Meta claims it generalizes well beyond that (to be verified). In practice, this paves the way for workflows encompassing entire codebases, session histories, or legal documents — all processed in a single pass before.

Architecturally, Scout relies on Meta's MoE framework, where only a subset of parameters is activated per token — unlike dense models such as GPT-4o, where all parameters are activated. The result: computational efficiency and high scalability.

Beyond architecture, Meta highlights Scout's multimodal capabilities. It was pre-trained on text, image, and video data with early fusion, which allows it to natively handle combinations of textual and visual prompts. On image-rich tasks such as visual grounding and VQA (visual question answering), Scout outperforms all previous Llama models — and holds its own against much larger systems.

In short, Scout is designed for versatility and scale. It runs efficiently, accepts more input than any previously released open-source model, and performs well on both text and image tasks. We will soon test this 10 M context limit — and we'll report back.

## Llama Maverick

Llama 4 Maverick is the generalist of the lineup: a large-scale multimodal model designed for performance in conversation, reasoning, image understanding, and code. While Scout pushes the boundaries of context length, Maverick prioritizes balanced, high output quality across all tasks. It is Meta's answer to GPT-4o, DeepSeek-V3, and Gemini 2.0 Flash.

Maverick has the same number of 17 billion active parameters as Scout, but with a larger MoE configuration: 128 experts and a total of about 400 billion parameters. Like Scout, it uses an MoE architecture that activates only part of the model per token — reducing inference costs while increasing capacity. The model runs on a single H100 DGX host, but can also be deployed in distributed inference for large-scale applications.

Meta adopted a different approach to post-training, combining lightweight supervised fine-tuning, online reinforcement learning, and direct preference optimization. The goal: sharpening performance on difficult prompts without over-constraining the model. To do this, Meta discarded more than 50% of the training examples deemed "easy" by earlier versions of Llama and built a training program focused on harder reasoning, coding, and multimodal tasks.

Maverick was also co-distilled from Llama 4 Behemoth, Meta's much larger internal model, which improved performance without additional training costs. According to Meta, this distillation chain led to a clear improvement in reasoning and conversation quality.

## Llama Behemoth

Llama 4 Behemoth is by far Meta's most powerful and largest model — but it is not yet available. Still in training, Behemoth is not a reasoning model in the same sense as DeepSeek-R1 or OpenAI's o3, which are designed and optimized for multi-step chain-of-thought reasoning.

Based on available information, it also does not appear to be designed as a direct-use product. It acts instead as a teacher model, used to distill and shape Scout and Maverick. Once released, it could also allow others to distill their own models.

Behemoth has 288 billion active parameters, organized via 16 experts, for a total approaching 2 trillion parameters. Meta designed an entirely new training infrastructure to support Behemoth at this scale. It introduces asynchronous reinforcement learning, curriculum sampling based on prompt difficulty, and a new distillation loss function that dynamically balances "soft" and "hard" targets.

Behemoth's post-training also required a different recipe. Meta discarded more than 95% of SFT examples to focus on difficult prompts and directed reinforcement learning toward complex reasoning, coding, and multilingual scenarios. Sampling from varied system instructions helped the model generalize, while dynamic filtering eliminated low-value prompts during RL training.

## Llama 4 Benchmarks

Meta published internal benchmark results for each of the Llama 4 models, comparing them to previous Llama variants as well as several competing open-weight models and state-of-the-art models.

In this section, I present the highlights of the benchmarks for Scout, Maverick, and Behemoth, according to Meta's figures. As always, be cautious with self-reported benchmarks; they nevertheless offer a useful first overview of each model's performance by task and their position in the current landscape. Let's start with Scout.

### Llama Scout Benchmarks

