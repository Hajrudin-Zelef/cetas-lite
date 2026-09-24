---
id: collect-240926-datacamp/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "OpenAI", "vLLM"]
dates: []
keywords: ["llama", "benchmark", "benchmarks", "claude", "consumer", "context window", "deepseek", "distillation", "fine-tuning", "gemini", "gpu", "gpus"]
source: docs/RAG/clean_en/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [1, 164]
sha256: 76f7cca4f98a266568df2f8f8ac1110b34051d28f795588e04014bbda4238700
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

Llama 4 Scout performs very well on a mix of reasoning, code, and multimodal benchmarks — especially given its reduced number of active parameters and its footprint on a single GPU.

Source: MetaAI

On image understanding, Scout leads its competitors: 88.8 on ChartQA and 94.4 on DocVQA (test), better than Gemini 2.0 Flash-Lite (73.0 and 91.2) and at or slightly above Mistral 3.1 and Gemma 3 27B.

On visual reasoning benchmarks like MMMU (69.4) and MathVista (70.7), it also leads the pack on the open-weight side, ahead of Gemma 3 (64.9; 67.6), Mistral 3.1 (62.8; 68.9), and Gemini Flash-Lite (68.0; 57.6).

In code, Scout scores 32.8 on LiveCodeBench, ahead of Gemini Flash-Lite (28.9) and Gemma 3 27B (29.7), but slightly behind Llama 3.3 (33.3). It is not a model primarily focused on code, but it holds its own.

On knowledge and reasoning, Scout reaches 74.3 on MMLU Pro and 57.2 on GPQA Diamond, surpassing all other open-weight models on both. These benchmarks favor long, multi-step reasoning, so Scout's performance is notable, especially at this scale.

Finally, its long-context capabilities show real potential. On MTOB (Massive Textual Overlap Benchmark), which evaluates the ability to translate between English and KGV, a low-resource language, it scores 42.2/36.6 on the half-book test and 39.7/36.3 on the full-book test. On the half-book, Gemini 2.0 Flash-Lite retains a slight advantage with 42.3, but Scout closes the gap on the full book, surpassing Gemini's 35.1/30.0.

### Llama Maverick Benchmarks

Maverick is the most balanced model in the Llama 4 lineup — and the benchmarks confirm it. Without aiming for Scout's context extremes or Behemoth's raw scale, it remains consistent across all key categories: multimodal reasoning, coding, language understanding, and long-context retention.

Source: MetaAI

In visual reasoning, Maverick scores 73.4 on MMMU and 73.7 on MathVista, ahead of Gemini 2.0 Flash (71.7 and 73.1) and GPT-4o (69.1 and 63.8). On ChartQA (image understanding), it reaches 90.0, slightly above Gemini's 88.3 and well above GPT-4o's 85.7. On DocVQA, Maverick reaches 94.4, matching Scout and ahead of GPT-4o's 92.8.

In coding, Maverick scores 43.4 on LiveCodeBench, above GPT-4o (32.3), Gemini Flash (34.5), and close to DeepSeek v3.1's 45.8.

On reasoning and knowledge, Maverick reaches 80.5 on MMLU Pro and 69.8 on GPQA Diamond, again ahead of Gemini Flash (77.6 and 60.1) and GPT-4o (no MMLU Pro score, 53.6 on GPQA). DeepSeek v3.1 retains a 0.7-point lead on MMLU Pro.

Maverick also performs well in multilingual understanding, with 84.6 on Multilingual MMLU, slightly above Gemini's 81.5. An asset for developers working across multiple languages or geographic regions.

On long-context evaluations (MTOB), Maverick scores 54.0/46.4 on the half-book and 50.8/46.7 on the full book — well ahead of Gemini's 48.4/39.8 and 45.5/39.6. These scores suggest that, even though it doesn't highlight its context length as much as Scout, it genuinely benefits from its extended window.

### Llama Behemoth Benchmarks

Behemoth hasn't been released yet, but its benchmark numbers are worth a look.

Source: MetaAI

On benchmarks with a strong STEM component, Behemoth excels. It reaches 95.0 on MATH-500 — above Gemini 2.0 Pro (91.8) and well ahead of Claude Sonnet 3.7 (82.2). On MMLU Pro, Behemoth scores 82.2, while Gemini Pro gets 79.1 (no score reported for Claude). And on GPQA Diamond, another benchmark that values factual depth and precision, Behemoth reaches 73.7, ahead of Claude (68.0), Gemini (64.7), and GPT-4.5 (71.4).

In multilingual understanding, Behemoth scores 85.8 on Multilingual MMLU, slightly ahead of Claude Sonnet (83.2) and GPT-4.5 (85.1). Significant scores for developers worldwide, and Behemoth currently leads this category.

In visual reasoning, Behemoth reaches 76.1 on MMMU, ahead of Gemini (71.8), Claude (72.7), and GPT-4.5 (74.4). Even though this isn't its main focus, it remains competitive against leading multimodal models.

In code generation, Behemoth scores 49.4 on LiveCodeBench. A score significantly higher than Gemini 2.0 Pro (36.0).

## How to Access Llama 4

Llama 4 Scout and Llama 4 Maverick are both available now under Meta's open-weight license. You can download them directly from the official Llama website or via Hugging Face.

To access the models through Meta's services, you can interact with Meta AI on several platforms: WhatsApp, Messenger, Instagram, and Facebook. Access currently requires signing in with a Meta account, and there is no standalone API endpoint for Meta AI — at least not yet.

If you're considering integrating the models into your own applications or infrastructure, keep the license clause in mind: if your product or service exceeds 700 million monthly active users, you'll need to obtain a separate authorization from Meta. The models otherwise remain usable for research, experimentation, and most commercial uses.

## Conclusion

Scout introduces unprecedented context length on a single GPU. Maverick rivals larger models on reasoning, code, and multimodal tasks. And Behemoth, still in training, shows how teacher models can shape more efficient and deployable variants.

The open-weight space has never been this competitive. DeepSeek, Qwen, Gemma, and soon OpenAI, are all advancing with solid releases. Llama 4 continues Meta's effort to offer scalable and open models for a wide range of uses.

## FAQs

### Is there an API for Llama 4?

Meta has not released an official API for Llama 4. However, third-party providers may offer API access to Llama 4.

### Can I fine-tune Llama 4 Scout or Maverick on my own data?

Yes, both models are open-weight and can be fine-tuned.

### Can I run Llama 4 models locally?

You can run Scout locally if you have access to a high-end GPU (such as an A100 or H100). Maverick is significantly larger and generally requires multiple GPUs or distributed infrastructure. For lightweight testing, quantized versions of Scout may work on consumer hardware with tools like llama.cpp or vLLM.

### What are the hardware requirements for Llama 4 Scout?

Scout is designed to fit on a single H100 GPU. That said, depending on context length and batch size, you might run smaller versions or quantized models on lower-tier GPUs like the A100 or even an RTX 4090, with reduced performance.

### Is Llama 4 multilingual?

Yes — Maverick and Behemoth show very good results on multilingual benchmarks like Multilingual MMLU. Although Meta hasn't published per-language details, early benchmarks suggest good performance on major non-English languages.

### Can I use Llama 4 in commercial products?

Yes, unless your company or product exceeds 700 million monthly active users, in which case you'll need to obtain a special license from Meta. For most startups, researchers, and individual developers, the standard license applies.

### Can I distill my own model from Llama Behemoth?

Not yet. Behemoth has not been released, and there's no indication of when Meta will make it public. That said, Meta used Behemoth internally to distill Scout and Maverick: if it is released, it could serve as a basis for further distillations.

### What is the difference between Llama 3.1, Llama 3.3, and Llama 4?

Llama 3.1 and 3.3 were dense models with limited or no multimodal support. Llama 4 moves to a mixture-of-experts architecture and adds native multimodal training. Scout and Maverick also incorporate longer context windows and improved post-training techniques.

I am a copywriter and writer and I cover blogs, tutorials, and news about AI, making sure everything conforms to a solid content strategy and SEO best practices. I have written data science courses on Python, statistics, probability, and data visualization. I have also published an award-winning novel and I spend my free time writing screenplays and making films.
