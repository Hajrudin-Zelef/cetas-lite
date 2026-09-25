---
id: collect-250926-servers-hardware/servers-hardware/unsloth-dynamic-3-0-ggufs
title: "Unsloth Dynamic 3.0 GGUFs"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Apple", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "Unsloth", "vLLM"]
dates: []
keywords: ["gguf", "agent", "agentic", "agents", "benchmark", "benchmarks", "deepseek", "gpu", "inference", "int4", "kimi", "llama"]
source: docs/RAG/clean4/Unsloth Dynamic 3.0 GGUFs.md
source_anchor: ""
source_lines: [1, 243]
sha256: 74c20d4d8a1a154b6cd862afd2e95c42c64a36ce6251ce0f9e4f5e9598083917
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
[Sept 10, 2025 update:](/docs/basics/dynamic-3.0-ggufs/unsloth-dynamic-ggufs-on-aider-polyglot.md) You asked for tougher benchmarks, so here's Aider Polyglot results! Our Dynamic 3-bit DeepSeek V3.1 GGUF scores **75.6%**, surpassing many full-precision SOTA LLMs. [Read more.](/docs/basics/dynamic-3.0-ggufs/unsloth-dynamic-ggufs-on-aider-polyglot.md)
{% endhint %}
You can also view real-world use-case benchmarks conducted by Benjamin Marie for LiveCodeBench v6, MMLU Pro etc.:

You can see how Unsloth's GGUFs performs better than the non-Unsloth quants despite being \~8GB smaller.
Detailed analysis of our benchmarks and evaluation further below.
### 💡 What's New in Dynamic v2.0?
* **Revamped Layer Selection for GGUFs + safetensors:** Unsloth Dynamic 2.0 now selectively quantizes layers much more intelligently and extensively. Rather than modifying only select layers, we now dynamically adjust the quantization type of every possible layer, and the combinations will differ for each layer and model.
* Current selected and all future GGUF uploads will utilize Dynamic 2.0 and our new calibration dataset. The dataset contains more than >1.5M **tokens** (depending on model) and comprise of high-quality, hand-curated and cleaned data - to greatly enhance conversational chat performance.
* Previously, our Dynamic quantization (DeepSeek-R1 1.58-bit GGUF) was effective only for MoE architectures. **Dynamic 2.0 quantization now works on all models (including MOEs & non-MoEs)**.
* **Model-Specific Quants:** Each model now uses a custom-tailored quantization scheme. E.g. the layers quantized in Gemma 3 differ significantly from those in Llama 4.
* To maximize efficiency, especially on Apple Silicon and ARM devices, we now also add Q4\_NL, Q5.1, Q5.0, Q4.1, and Q4.0 formats.
To ensure accurate benchmarking, we built an internal evaluation framework to match official reported 5-shot MMLU scores of Llama 4 and Gemma 3. This allowed apples-to-apples comparisons between full-precision vs. Dynamic v2.0, **QAT** and standard **imatrix** GGUF quants.

All future GGUF uploads will utilize Unsloth Dynamic 2.0, and our Dynamic 4-bit safe tensor quants will also benefit from this in the future.
## 📊 Why KL Divergence?
[Accuracy is Not All You Need](https://arxiv.org/pdf/2407.09141) showcases how pruning layers, even by selecting unnecessary ones still yields vast differences in terms of "flips". A "flip" is defined as answers changing from incorrect to correct or vice versa. The paper shows how MMLU might not decrease as we prune layers or do quantization,but that's because some incorrect answers might have "flipped" to become correct. Our goal is to match the original model, so measuring "flips" is a good metric.

{% hint style="info" %}
**KL Divergence** should be **one of the gold standards for reporting quantization errors** as per the research paper "Accuracy is Not All You Need". **Using perplexity is incorrect** since output token values can cancel out, so we must use KLD or harder benchmarks like [Aider](/docs/basics/dynamic-3.0-ggufs/unsloth-dynamic-ggufs-on-aider-polyglot.md).
{% endhint %}
The paper also shows that interestingly KL Divergence is highly correlated with flips, and so our goal is to reduce the mean KL Divergence whilst increasing the disk space of the quantization as less as possible.
## ⚖️ Calibration Dataset Overfitting
Most frameworks report perplexity and KL Divergence using a test set of Wikipedia articles. However, we noticed using the calibration dataset which is also Wikipedia related causes quants to overfit, and attain lower perplexity scores. We utilize [Calibration\_v3](https://gist.github.com/bartowski1182/eb213dccb3571f863da82e99418f81e8) and [Calibration\_v5](https://gist.github.com/tristandruyen/9e207a95c7d75ddf37525d353e00659c/) datasets for fair testing which includes some wikitext data amongst other data. **Also instruct models have unique chat templates, and using text only calibration datasets is not effective for instruct models** (base models yes). In fact most imatrix GGUFs are typically calibrated with these issues. As a result, they naturally perform better on KL Divergence benchmarks that also use Wikipedia data, since the model is essentially optimized for that domain.
To ensure a fair and controlled evaluation, we do not to use our own calibration dataset (which is optimized for chat performance) when benchmarking KL Divergence. Instead, we conducted tests using the same standard Wikipedia datasets, allowing us to directly compare the performance of our Dynamic 2.0 method against the baseline imatrix approach.
## :1234: MMLU Replication Adventure
* Replicating MMLU 5 shot was nightmarish. We **could not** replicate MMLU results for many models including Llama 3.1 (8B) Instruct, Gemma 3 (12B) and others due to **subtle implementation issues**. Llama 3.1 (8B) for example should be getting \~68.2%, whilst using incorrect implementations can attain **35% accuracy.**
* Llama 3.1 (8B) Instruct has a MMLU 5 shot accuracy of 67.8% using a naive MMLU implementation. We find however Llama **tokenizes "A" and "\_A" (A with a space in front) as different token ids**. If we consider both spaced and non spaced tokens, we get 68.2% (+0.4%)
* Interestingly Llama 3 as per Eleuther AI's [LLM Harness](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/lm_eval/tasks/llama3/instruct/mmlu/_continuation_template_yaml) also appends **"The best answer is"** to the question, following Llama 3's original MMLU benchmarks.
* There are many other subtle issues, and so to benchmark everything in a controlled environment, we designed our own MMLU implementation from scratch by investigating [github.com/hendrycks/test](https://github.com/hendrycks/test) directly, and verified our results across multiple models and comparing to reported numbers.
## :sparkles: Gemma 3 QAT Replication, Benchmarks
The Gemma team released two QAT (quantization aware training) versions of Gemma 3:
1. Q4\_0 GGUF - Quantizes all layers to Q4\_0 via the formula `w = q * block_scale` with each block having 32 weights. See [llama.cpp wiki ](https://github.com/ggml-org/llama.cpp/wiki/Tensor-Encoding-Schemes)for more details.
2. int4 version - presumably [TorchAO int4 style](https://github.com/pytorch/ao/blob/main/torchao/quantization/README.md)?
We benchmarked all Q4\_0 GGUF versions, and did extensive experiments on the 12B model. We see the **12B Q4\_0 QAT model gets 67.07%** whilst the full bfloat16 12B version gets 67.15% on 5 shot MMLU. That's very impressive! The 27B model is mostly nearly there!

Metric

1B

4B

12B

27B

MMLU 5 shot

26.12%

55.13%

67.07% (67.15% BF16)

70.64% (71.5% BF16)

Disk Space

0.93GB

2.94GB

7.52GB

16.05GB

Efficiency*

1.20

10.26

5.59

2.84

We designed a new **Efficiency metric** which calculates the usefulness of the model whilst also taking into account its disk size and MMLU 5 shot score:
$$
\text{Efficiency} = \frac{\text{MMLU 5 shot score} - 25}{\text{Disk Space GB}}
$$
{% hint style="warning" %}
We have to **minus 25** since MMLU has 4 multiple choices - A, B, C or D. Assume we make a model that simply randomly chooses answers - it'll get 25% accuracy, and have a disk space of a few bytes. But clearly this is not a useful model.
{% endhint %}
On KL Divergence vs the base model, below is a table showcasing the improvements. Reminder the closer the KL Divergence is to 0, the better (ie 0 means identical to the full precision model)
| Quant | Baseline KLD | GB | New KLD | GB |
| --------- | ------------ | ----- | -------- | ----- |
| IQ1\_S | 1.035688 | 5.83 | 0.972932 | 6.06 |
| IQ1\_M | 0.832252 | 6.33 | 0.800049 | 6.51 |
| IQ2\_XXS | 0.535764 | 7.16 | 0.521039 | 7.31 |
| IQ2\_M | 0.26554 | 8.84 | 0.258192 | 8.96 |
| Q2\_K\_XL | 0.229671 | 9.78 | 0.220937 | 9.95 |
| Q3\_K\_XL | 0.087845 | 12.51 | 0.080617 | 12.76 |
| Q4\_K\_XL | 0.024916 | 15.41 | 0.023701 | 15.64 |
If we plot the ratio of the disk space increase and the KL Divergence ratio change, we can see a much clearer benefit! Our dynamic 2bit Q2\_K\_XL reduces KLD quite a bit (around 7.5%).
Truncated table of results for MMLU for Gemma 3 (27B). See below.
1. **Our dynamic 4bit version is 2GB smaller whilst having +1% extra accuracy vs the QAT version!**
2. Efficiency wise, 2bit Q2\_K\_XL and others seem to do very well!
| Quant | Unsloth | Unsloth + QAT | Disk Size | Efficiency |
| -------------- | --------- | ------------- | --------- | ---------- |
| IQ1\_M | 48.10 | 47.23 | 6.51 | 3.42 |
| IQ2\_XXS | 59.20 | 56.57 | 7.31 | 4.32 |
| IQ2\_M | 66.47 | 64.47 | 8.96 | 4.40 |
| Q2\_K\_XL | 68.70 | 67.77 | 9.95 | 4.30 |
| Q3\_K\_XL | 70.87 | 69.50 | 12.76 | 3.49 |
| **Q4\_K\_XL** | **71.47** | **71.07** | **15.64** | **2.94** |
| **Google QAT** | | **70.64** | **17.2** | **2.65** |
Click here for Full Google's Gemma 3 (27B) QAT Benchmarks:
| Model | Unsloth | Unsloth + QAT | Disk Size | Efficiency |
| -------------- | --------- | ------------- | --------- | ---------- |
| IQ1\_S | 41.87 | 43.37 | 6.06 | 3.03 |
| IQ1\_M | 48.10 | 47.23 | 6.51 | 3.42 |
| IQ2\_XXS | 59.20 | 56.57 | 7.31 | 4.32 |
| IQ2\_M | 66.47 | 64.47 | 8.96 | 4.40 |
| Q2\_K | 68.50 | 67.60 | 9.78 | 4.35 |
| Q2\_K\_XL | 68.70 | 67.77 | 9.95 | 4.30 |
| IQ3\_XXS | 68.27 | 67.07 | 10.07 | 4.18 |
| Q3\_K\_M | 70.70 | 69.77 | 12.51 | 3.58 |
| Q3\_K\_XL | 70.87 | 69.50 | 12.76 | 3.49 |
| Q4\_K\_M | 71.23 | 71.00 | 15.41 | 2.98 |
| **Q4\_K\_XL** | **71.47** | **71.07** | **15.64** | **2.94** |
| Q5\_K\_M | 71.77 | 71.23 | 17.95 | 2.58 |
| Q6\_K | 71.87 | 71.60 | 20.64 | 2.26 |
| Q8\_0 | 71.60 | 71.53 | 26.74 | 1.74 |
| **Google QAT** | | **70.64** | **17.2** | **2.65** |
## :llama: Llama 4 Bug Fixes + Run
We also helped and fixed a few Llama 4 bugs:
* Llama 4 Scout changed the RoPE Scaling configuration in their official repo. We helped resolve issues in llama.cpp to enable this [change here](https://github.com/ggml-org/llama.cpp/pull/12889)
* Llama 4's QK Norm's epsilon for both Scout and Maverick should be from the config file - this means using 1e-05 and not 1e-06. We helped resolve these in [llama.cpp](https://github.com/ggml-org/llama.cpp/pull/12889) and [transformers](https://github.com/huggingface/transformers/pull/37418)
* The Llama 4 team and vLLM also independently fixed an issue with QK Norm being shared across all heads (should not be so) [here](https://github.com/vllm-project/vllm/pull/16311). MMLU Pro increased from 68.58% to 71.53% accuracy.
* [Wolfram Ravenwolf](https://x.com/WolframRvnwlf/status/1909735579564331016) showcased how our GGUFs via llama.cpp attain much higher accuracy than third party inference providers - this was most likely a combination of the issues explained above, and also probably due to quantization issues.
As shown in our graph, our 4-bit Dynamic QAT quantization deliver better performance on 5-shot MMLU while also being smaller in size.
### Running Llama 4 Scout:
To run Llama 4 Scout for example, first clone llama.cpp:
```bash
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
-DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON -DLLAMA_CURL=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
Then download out new dynamic v 2.0 quant for Scout:
```python
# !pip install huggingface_hub hf_transfer
import os
os.environ["HF_HUB_ENABLE_HF_TRANSFER"] = "1"
from huggingface_hub import snapshot_download
snapshot_download(
repo_id = "unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF",
local_dir = "unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF",
allow_patterns = ["*IQ2_XXS*"],
)
```
And and let's do inference!
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
--model unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF/Llama-4-Scout-17B-16E-Instruct-UD-IQ2_XXS.gguf \
--threads 32 \
--ctx-size 16384 \
--n-gpu-layers 99 \
-ot ".ffn_.*_exps.=CPU" \
--seed 3407 \
--prio 3 \
--temp 0.6 \
--min-p 0.01 \
--top-p 0.9 \
-no-cnv \
--prompt "<|header_start|>user<|header_end|>\n\nCreate a Flappy Bird game.<|eot|><|header_start|>assistant<|header_end|>\n\n"
```
{% endcode %}
{% hint style="success" %}
Read more on running Llama 4 here:
{% endhint %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/dynamic-3.0-ggufs.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
