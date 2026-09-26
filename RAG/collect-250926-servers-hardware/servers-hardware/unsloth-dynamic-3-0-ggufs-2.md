---
id: collect-250926-servers-hardware/servers-hardware/unsloth-dynamic-3-0-ggufs-2
title: "Unsloth Dynamic 3.0 GGUFs"
domain: servers-hardware
role: reference
task: reference
actors: ["Apple", "DeepSeek", "Unsloth"]
dates: []
keywords: ["gguf", "benchmark", "benchmarks", "deepseek", "int4", "llama", "llama.cpp", "moe", "perplexity", "pruning", "quantization", "research"]
source: docs/RAG/clean4/Unsloth Dynamic 3.0 GGUFs.md
source_anchor: ""
source_lines: [56, 131]
sha256: d70101788ecb23f86a7045ae386be6cc2d7a6c20822ca3d87043fc2d58f85bec
---

# Unsloth Dynamic 3.0 GGUFs

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

