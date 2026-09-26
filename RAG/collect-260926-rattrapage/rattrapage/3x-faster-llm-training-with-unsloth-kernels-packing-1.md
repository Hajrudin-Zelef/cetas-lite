---
id: collect-260926-rattrapage/rattrapage/3x-faster-llm-training-with-unsloth-kernels-packing-1
title: "3x Faster LLM Training with Unsloth Kernels + Packing"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Unsloth"]
dates: ["2023-12", "2024-03"]
keywords: ["training", "attention", "benchmarks", "embedding", "fine-tuning", "gpu", "gpus", "memory", "pretraining", "rotary", "throughput"]
source: docs/RAG/lot-rattrapage/fine-tuning/3x Faster LLM Training with Unsloth Kernels + Packing.md
source_anchor: ""
source_lines: [1, 65]
sha256: 1ff5a6d12326def696357eae704a5170159156eb1793f921e95ee98457af6a41
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

