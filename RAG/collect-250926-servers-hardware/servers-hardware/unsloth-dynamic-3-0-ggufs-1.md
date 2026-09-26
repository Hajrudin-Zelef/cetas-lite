---
id: collect-250926-servers-hardware/servers-hardware/unsloth-dynamic-3-0-ggufs-1
title: "Unsloth Dynamic 3.0 GGUFs"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Unsloth"]
dates: []
keywords: ["gguf", "agentic", "benchmarks", "inference", "kimi", "llama", "llama.cpp", "mistral", "perplexity", "quantization", "reasoning", "tool calling"]
source: docs/RAG/clean4/Unsloth Dynamic 3.0 GGUFs.md
source_anchor: ""
source_lines: [1, 55]
sha256: c7b9960cad1f664eeb1ac1e48f9d0d4f5ae55973c714e13fa988f5819a103529
---

# Unsloth Dynamic 3.0 GGUFs

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/basics/dynamic-3.0-ggufs.md).
# Unsloth Dynamic 3.0 GGUFs
[**Unsloth**](https://github.com/unslothai/unsloth) **Dynamic v3.0** is the next iteration of our Dynamic quantization and a major improvement over Dynamic v2.0.
Today, we’re releasing [**Qwen3.8-27B**](/docs/models/qwen3.8.md) Dynamic v3.0 quants that deliver **>10% top-1 better accuracy at the same size** compared to **every other provider**. This is an update of our first shared **early preview** version of Dynamic v3.0. The new 3.0 GGUFs work with most inference engines including **llama.cpp** and [**Unsloth Desktop**](/docs/desktop.md).
{% columns %}
{% column width="41.66666666666667%" %}
Dynamic v3.0 overall preserves more model quality while keeping the same size, with stronger results across metrics like **Divergence-300** @32 and **KL Divergence**.
Also a huge thanks to all your support! We saw over 5.1 million Unsloth Qwen3.8 downloads in just 5 days!
{% endcolumn %}
{% column width="58.33333333333333%" %}
{% endcolumn %}
{% endcolumns %}
Our new methodology composes of many new features and improvements. We now use a much higher-quality imatrix calibration dataset from diverse sources. The dataset is refined for **agentic coding, chat**, and multilingual performance. We also improved **layer selection** and introduced many more quantization techniques to preserve as much model quality as possible.
We **do not train on the imatrix calibration dataset**, and we do NOT use **QAT** or **QAD**. Everything is done through **post-training quantization**. Our imatrix file used is available for the community to test, evaluate, and use. We encourage researchers and developers to create variations and fine-tunes of Qwen3.8 using our Unsloth quants/imatrix. You can read our [overfitting analysis](#not-overfitting) as well.
* We also removed the MTP module from smaller quants under `UD-Q2_K_XL` (8.37GB and lower) to converse around 500MB of disk space - you can use the `Q4_0` MTP separate module if needed
* We also made some smaller UD-1bit quants with `UD-IQ1_S` being 6.2GB (without MTP) which retain around 72% top-1 accuracy yet being 89% smaller.
* `UD-Q2_K_XL` is around +8% more accurate on top-1 than the next best and it's 9.83GB and managed to create a working HTML program with 1 small JS bug - previously it would break.
### 🔀 Divergence-300 @32
We generally report top-1 accuracy like how for Kimi-K3 "Dynamic 1-bit reaches **\~78.9%** top-1 accuracy while being **62% smaller**." However top-1 is an argmax on 1 prediction, so it's not really effective on gauging actual inference.
We created a dataset of 300 held out examples (NOT in calibration dataset) from Terminal-Bench 2.1 + DeepSWE + Harbor + MathArena 2025-26 + non-Latin/long-doc prompts and we did greedy argmax decoding for 32 tokens for BF16 vs all quants and providers. See [overfitting analysis](#not-overfitting) for more details on overfitting.
This allows us to gauge if there is overfitting and if quant outputs are similar to BF16's trajectories over multiple tokens. This is a better metric than top-1 accuracy since we extend KLD top-1 to more like KLD top-1 at 32 tokens.
### :question:1-bit should not be used for agentic use-cases
As seen in [#divergence-300-32](#divergence-300-32 "mention"), there is a sharp drop off from UD-Q2\_K\_XL to UD-IQ2\_S for 32 token prediction from around 25% accuracy to under 8-10%. This sharp drop off means tool calling and non thinking modes break down. Some issues and mitigations if using the 1-bit:
1. **Excessive looping**\
You will see a lot of looping when using quants below UD-Q2\_K\_XL - use `presence_penalty = 1.5` in all cases (or higher)
2. **Empty responses**\
Always enable thinking at least on low reasoning for 1-bit quants - non reasoning modes cause the model to not even output anyway
3. **Agentic use cases and tool calling**\
Do not use the model for tool calling - only **general knowledge is retained** when quantized heavily, and the model will either fail to call tools, keep calling tools or not even call them.
4. **General Knowledge works**\
A Top-1 recovery of 77% is not a replacement for Divergence-300 @32 at 8% which is a better indicator for actual inference workloads - you can use the model for very short general knowledge fact questions, but best to use UD-Q2\_K\_XL.
### 🔀 KL Divergence Benchmarks
We ran KLD benchmarks for all providers as well and report Top-1 and KLD mean. At all levels especially on the smaller quant sizes, Unsloth UD-3 quants get up to +10% extra top-1 accuracy at the same disk space!
All plots remove the MTP head from the x axis when calculating disk space to provide a fair comparison to everyone.
### :dove:Not Overfitting
When comparing to our older UD-2 on unseen Wikitext and Code, we show great improvement on KLD - the bigger ones not so much, so we still use our old UD-2 for the larger quants - we plan to experiment and improve them as well!
We also control for overfitting by using totally different datasets for calibration and remove all leakages as much as possible. We test KLD on these unseen datasets, and also we do NOT do QAD / QAT, just pure PTQ so overfitting is less of a concern vs other QAD / QAT approaches.
Similarly [#divergence-300-32](#divergence-300-32 "mention") uses an unseen dataset of 300 prompts from DeepSWE, Terminal Bench and others, and acts as another dataset to gauge overfitting - and shows our new UD-3 methods do not overfit.
***
## Dynamic v2.0 (Old)
We're introducing [Unsloth](https://github.com/unslothai/unsloth) Dynamic v2.0 quantization - a major upgrade to our previous quants. This new method outperforms leading quantization methods and sets new benchmarks for [Aider Polyglot](/docs/basics/dynamic-3.0-ggufs/unsloth-dynamic-ggufs-on-aider-polyglot.md), 5-shot MMLU and KL Divergence.
This means you can now run + fine-tune [quantized LLMs](/docs/models/tutorials.md) while preserving as much accuracy as possible! You can run the 2.0 GGUFs on most inference engines like llama.cpp, [Unsloth Studio](/docs/new/studio.md) etc.
{% columns %}
{% column %}
**Apr 20, 2026 Update:** See our new GGUF Benchmarks for [Qwen3.6](/docs/models/qwen3.6.md#unsloth-gguf-benchmarks) and [Gemma 4](/docs/models/gemma-4.md#unsloth-gguf-benchmarks).
[Feb 27, 2026 Update:](/docs/models/qwen3.5/gguf-benchmarks.md) **Qwen3.5** is out and we fixed some tool-calling chat template issues and benchmarked every GGUF on perplexity & KL Divergence. [See benchmarks!](/docs/models/qwen3.5/gguf-benchmarks.md)
The **key advantage** of using the [Unsloth package](https://github.com/unslothai/unsloth) and quants is our active role in fixing bugs in major models. We've collaborated directly with teams behind [Qwen3](https://www.reddit.com/r/LocalLLaMA/comments/1kaodxu/qwen3_unsloth_dynamic_ggufs_128k_context_bug_fixes/), [Meta (Llama 4)](https://github.com/ggml-org/llama.cpp/pull/12889), [Mistral (Devstral)](https://app.gitbook.com/o/HpyELzcNe0topgVLGCZY/s/xhOjnexMCB3dmuQFQ2Zq/~/changes/618/basics/tutorials-how-to-fine-tune-and-run-llms/devstral-how-to-run-and-fine-tune), [Google (Gemma 1–3)](https://news.ycombinator.com/item?id=39671146) and [Microsoft (Phi-3/4)](https://simonwillison.net/2025/Jan/11/phi-4-bug-fixes), contributing fixes that increase accuracy.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
{% hint style="success" %}
Unsloth Dynamic GGUFs can now be run in [Unsloth Studio](/docs/new/studio.md) ✨
{% endhint %}
{% hint style="success" %}
