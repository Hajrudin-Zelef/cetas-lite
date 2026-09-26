---
id: collect-240926-datacamp/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore-2
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: datacamp
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "OpenAI", "vLLM"]
dates: []
keywords: ["llama", "benchmark", "benchmarks", "claude", "consumer", "deepseek", "gemini", "gpu", "gpus", "license", "llama.cpp", "mistral"]
source: docs/RAG/clean_en/datacamp/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [72, 164]
sha256: a19c189a38752e7b538a9477c9e2e26562ea68e539420e64e48aee22d92dfe7e
---

# llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore

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
