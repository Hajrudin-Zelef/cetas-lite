---
id: collect-250926-servers-hardware/servers-hardware/3x-faster-llm-training-with-unsloth-kernels-packing
title: "3x Faster LLM Training with Unsloth Kernels + Packing"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: ["2023-12", "2024-03"]
keywords: ["training", "agent", "agents", "attention", "benchmarks", "embedding", "fine-tuning", "flash attention", "gpu", "gpus", "llama", "memory"]
source: docs/RAG/clean4/3x Faster LLM Training with Unsloth Kernels + Packing.md
source_anchor: ""
source_lines: [1, 125]
sha256: ab6585bfed67016e6191f92178f119400f7160bf4fcb7259c7d62208d93c8d56
---

# 3x Faster LLM Training with Unsloth Kernels + Packing

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/blog/3x-faster-training-packing.md).
# 3x Faster LLM Training with Unsloth Kernels + Packing
Learn how Unsloth increases training throughput and eliminates padding waste for fine-tuning.
Unsloth now supports up to **5× faster** (typically 3x) training with our new custom **RoPE and MLP Triton kernels**, plus our new smart auto packing. Unsloth's new kernels + features not only increase training speed, but also further **reduces VRAM use (30% - 90%)** with no accuracy loss. [Unsloth GitHub](https://github.com/unslothai/unsloth)\
\
This means you can now train LLMs like [Qwen3](/docs/models/tutorials/qwen3-how-to-run-and-fine-tune.md)-4B not only on just **3GB VRAM**, but also 3x faster.
Our auto [**padding-free**](#padding-free-by-default) uncontaminated packing is smartly enabled for all training runs without any changes, and all fast attention backends (FlashAttention 3, xFormers, SDPA). [Benchmarks](#analysis-and-benchmarks) show training losses match non-packing runs **exactly**.
* **2.3x faster QK Rotary Embedding** fused Triton kernel with packing support
* Updated SwiGLU, GeGLU kernels with **int64 indexing for long context**
* **2.5x to 5x faster uncontaminated packing** with xformers, SDPA, FA3 backends
* **2.1x faster padding free, 50% less VRAM**, 0% accuracy change
* Unsloth also now has improved SFT loss stability and more predictable GPU utilization.
* This new upgrade works **for all training methods** e.g. full fine-tuning, pretraining etc.
### :drum:Fused QK RoPE Triton Kernel with packing
Back in December 2023, we introduced a RoPE kernel coded up in Triton as part of our Unsloth launch. In March 2024, a community member made end to end training 1-2% faster by optimizing the RoPE kernel to allow launching a block for a group of heads. See [PR 238](https://github.com/unslothai/unsloth/pull/238).
One issue is for each Q and K, there are 2 Triton kernels. We merged them into 1 Triton kernel now, and enabled variable length RoPE, which was imperative for padding free and packing support. This makes the RoPE kernel in micro benchmarks **2.3x faster on longer context lengths**, and 1.9x faster on shorter context lengths.
We also eliminated all clones and contiguous transpose operations, and so **RoPE is now fully inplace**, reducing further GPU memory. Note for the backward pass, we see that `sin1 = -sin1` since:
```
Q * cos + rotate_half(Q) * sin
is equivalent to
Q * cos + Q @ R * sin
where R is a rotation matrix [ 0, I]
[-I, 0]
dC/dY = dY * cos + dY @ R.T * sin
where R.T is again the same [ 0, -I]
but the minus is transposed. [ I, 0]
```
### :railway\_car:Int64 Indexing for Triton Kernels
During 500K long context training which we introduced in [500K Context Training](/docs/blog/500k-context-length-fine-tuning.md), we would get CUDA out of bounds errors. This was because MLP kernels for SwiGLU, GeGLU had int32 indexing which is by default in Triton and CUDA.
We can't just do `tl.program_id(0).to(tl.int64)` since training will be slightly slower due to int64 indexing. We instead make this a `LONG_INDEXING: tl.constexpr` variable so the Triton compiler can specialize this. This allows shorter and longer context runs to both run great!
{% code overflow="wrap" %}
```python
block_idx = tl.program_id(0)
if LONG_INDEXING:
offsets = block_idx.to(tl.int64) * BLOCK_SIZE + tl.arange(0, BLOCK_SIZE).to(tl.int64)
n_elements = tl.cast(n_elements, tl.int64)
else:
offsets = block_idx * BLOCK_SIZE + tl.arange(0, BLOCK_SIZE)
```
{% endcode %}
### :abacus:Why is padding needed & mathematical speedup
Computers and GPUs cannot process different length datasets, so we have to pad them with 0s. This causes wastage. Assume we have a dataset of 50% short sequences S, and 50% long sequences L, then in the worst case, padding will cause token usage to be $$\text{batchsize} \times L$$ since the longest sequence length dominates.
By packing multiple examples into a single, long one-dimensional tensor, we can eliminate a significant amount of padding. In fact we get the below token usage:
$$
\text{Token Usage} = \frac{\text{batchsize}}{2}L+\frac{\text{batchsize}}{2}S
$$
By some math and algebra, we can work out the speedup via:
$$
\text{Speedup} = \frac{\text{batchsize} \times L}{\frac{\text{batchsize}}{2}L+\frac{\text{batchsize}}{2}S} = 2 \frac{L}{L + S}
$$
By assuming $$S\rightarrow0$$ then we get a 2x theoretical speedup since $$2 \frac{L}{L + 0} = 2$$
By changing the ratio of 50% short sequences, and assuming we have MORE short sequences, for eg 20% long sequences and 80% short sequences, we get $$\frac{L}{0.2L + 0.8S}\rightarrow\frac{L}{0.2L}=5$$ so 5x faster training! This means packing's speedup depends on how short rows your dataset has (the more shorter, the faster).
### :clapper:Padding-Free by Default
In addition to large throughput gains available when setting `packing = True` in your `SFTConfig` , we will **automatically use padding-free batching** in order to reduce padding waste improve throughput and increases tokens/s throughput, while resulting in the ***exact same loss*** as seen in the previous version of Unsloth.
For example for Qwen3-8B and Qwen3-32B, we see memory usage decrease by 60%, be 2x faster, and have the same exact loss and grad norm curves!

### :spades:Uncontaminated Packing 2-5x faster training
Real datasets can contain different sequence lengths, so increasing the batch size to 32 for example will cause padding, making training slower and use more VRAM.
{% hint style="success" %}
In the past, increasing `batch_size` to large numbers (>32) will make training SLOWER, not faster. This was due to padding - we can now eliminate this issue via `packing = True`, and so training is FASTER!
{% endhint %}
When we pack multiple samples into a single one-dimensional tensor, we keep sequence length metadata around in order to properly mask samples, without leaking attention between samples. We also need the RoPE kernel described in [#fused-qk-rope-triton-kernel-with-packing](#fused-qk-rope-triton-kernel-with-packing "mention") to allow reset position ids.
{% columns %}
{% column width="41.66666666666667%" %}

The first graph (above) plots progress on `yahma/alpaca-cleaned` with `max_length = 2048`, Unsloth new with packing + kernels (maroon) vs. Unsloth old (gray). Both are trained with `max_steps = 500`, but we plot the x-axis in wall-clock time. Notice that we train on nearly 40% of an epoch in the packed case in the same amount of steps (and only a bit more wall-clock time) that it takes to train less than 5% of an epoch in the unpacked case.
Similarly, the 2nd graph (above) plots loss from the same runs, this time plotted with training steps on the x-axis. Notice that the losses match in scale and trend, but the loss in the packing case is less variable since the model is seeing more tokens per training step.
### :sparkles:How to enable packing?
**Update Unsloth first and padding free is done by default**! So all training is immediately 1.1 to 2x faster with 30% less memory usage at least and 0 change in loss curve metric!
{% code overflow="wrap" %}
```bash
pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth
pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth_zoo
```
{% endcode %}
We also support Flash Attention 3 via Xformers, SDPA support, Flash Attention 2, and this works on old GPUs (Tesla T4, RTX 2080) and new GPUs like H100s, B200s etc! Sample packing works *regardless of choice of attention backend or model family*, so enjoy the same speedups previously had with these fast attention implementations!
If you want to enable explicit packing, then add `packing = True` to enable up to 5x faster training!
{% hint style="warning" %}
Note `packing=True` will change the training loss and will make the dataset number of rows truncated, since multiple short sequences are packed into 1 sequence. You might see the number of examples in the dataset shrink.
To not get different training loss numbers, simply set `packing=False` and we will enable auto padding-free, which already makes training faster!
{% endhint %}
```python
from unsloth import FastLanguageModel
from trl import SFTTrainer, SFTConfig
model, tokenizer = FastLanguageModel.from_pretrained(
"unsloth/Qwen3-14B",
)
trainer = SFTTrainer(
model = model,
processing_class = tokenizer,
train_dataset = dataset,
args = SFTConfig(
per_device_train_batch_size = 1,
max_length = 4096,
…,
packing = True, # required to enable sample packing!
),
)
trainer.train()
```
All our notebooks are automatically faster (no need to do anything). See [Unsloth Notebooks](/docs/get-started/unsloth-notebooks.md)
{% columns %}
{% column %}
Qwen3 14B faster:
{% embed url="" %}
{% endcolumn %}
{% column %}
Llama 3.1 Conversational faster:
{% embed url="" %}
{% endcolumn %}
{% endcolumns %}
Thank you! If you're interested, see our [500K Context Training](/docs/blog/500k-context-length-fine-tuning.md) blog, [Memory Efficient RL](/docs/get-started/reinforcement-learning-rl-guide/memory-efficient-rl.md) blog and [Long Context gpt-oss](/docs/models/gpt-oss-how-to-run-and-fine-tune/long-context-gpt-oss-training.md) blog for more topics on kernels and performance gains!
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/blog/3x-faster-training-packing.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
