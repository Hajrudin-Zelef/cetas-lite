---
id: collect-250926-servers-hardware/servers-hardware/fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation
title: "fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Nvidia", "Unsloth", "Z.ai"]
dates: []
keywords: ["moe", "agentic", "attention", "benchmarks", "compute", "consumer", "deepseek", "fine-tuning", "glm", "gpu", "gpus", "inference"]
source: docs/RAG/clean4/fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 506]
sha256: f7e80a283ba5a4259142914d365546fe4b9fb8062ca9e81d169d60347fcbeda0
---

# fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation

We’re introducing ~12x faster Mixture of Experts (MoE) LLM training with **>35% less VRAM** and **~6x longer context** with our new MoE Triton kernels and new mathematical optimizations, all with no loss in accuracy.

- gpt-oss-20b fine-tunes in **12.8 GB VRAM** . Qwen3-30B-A3B (16-bit LoRA) uses 63GB.
- Our kernels work on both data-center (B200, H100), **consumer** and older GPUs (e.g., RTX 3090), and FFT, LoRA and QLoRA.

In collaboration with 🤗Hugging Face, we made all MoE training runs standardized with PyTorch’s new `torch._grouped_mm` function. Transformers v5 was recently optimized with ~6x faster MoE than v4 and Unsloth pushes this even further with custom Triton grouped‑GEMM + LoRA kernels for an **additional** ~2x speedup, >35% VRAM reduction and >6x longer context (12-30x overall speedup vs v4).

Try our Unsloth Notebooks for fast MoE training:

Alongside `torch._grouped_mm` (see ❓What is torch._grouped_mm?), we created custom Triton MoE kernels that can be even faster in some cases. They are also **backwards compatible** with older hardware like A100, and older PyTorch versions.

On A100, our **Triton kernels are ~2.5× faster** than `torch._grouped_mm`. The kernels also have a one‑time autotune step to pick the best kernel config.

Auto-tuning takes ~2 minutes once at the start of training, but can speed up the full run by 35% on A100 vs `_grouped_mm`, which is well worth it for longer runs.

The larger the model and more context you use, **the more pronounced the memory savings from our Unsloth kernels will be** (efficiency will scale exponentially).

Our main innovation is our **Split LoRA approach** for efficient MoE, which uses ~35% less memory and is 2x faster training when compared to Transformers v5 + `torch._grouped_mm`. Custom `torch._grouped_mm` + our Triton kernels are ~12-30x faster than Transformers v4.

Training MoE models in **4-bit** QLoRA isn’t recommended right now because BitsandBytes doesn’t support it. This isn’t specific to Unsloth. For now, use bf16 for LoRA or full fine-tuning.

Unsloth will auto select either the following backends depending on your hardware:

grouped_mm

`torch._grouped_mm` - available on T4s all the way until B200s, but optimized for H100s+.

unsloth_triton

Unsloth Triton kernels - which will turn on automatically for A100s, and older PyTorch versions.

native_torch

Native PyTorch. It's 12x slower, but our VRAM reductions are still there!

You can also toggle them yourself:

To enable faster MoE training, update Unsloth via `pip install --upgrade unsloth unsloth_zoo`

Previously, Mixture-of-Experts (MoE) weights were stored as a `ModuleList` of per‑expert linear layers. The only practical way to run a forward pass was a for‑loop over experts, which is expensive and suboptimal.

PyTorch recently introduced `grouped_mm` to address this exact bottleneck. In parallel, we provide our own MoE‑optimized Triton kernels. This also lines up with a key Transformers change: as of Transformers v5, expert weights are stored as a `single nn.Parameter`, making `grouped_mm` a natural fit for faster MoE training and inference.

So transformers 4.57.6 changed:

to transformers 5.0.0 style:

`torch._grouped_mm` works on GPUs starting with the NVIDIA T4, and we’ve verified it on H100, A100, B200, and RTX 6000 Pro, so support is broadly available.

We also previously introduced Unsloth Flex Attention for gpt-oss, and these optimizations should make it even more efficient.

Below is a comparison across sequence lengths for training speed and memory usage versus Transformers v5 (which already uses `torch._grouped_mm` for MoE). For **gpt-oss BF16 MoE training, we see 7x faster training and 36% VRAM reduction** on NVIDIA B200. For Qwen3-30B-A3B, it's 1.8x faster, and **GLM 4.7 Flash is 2.1x faster on RTX PRO 6000**. All benchmarks are done with LoRA rank = 64 and all LoRA modules on MoE layers (gate, up, down).

We fine-tuned unsloth/gpt-oss-20b-BF16 for benchmarking. Unsloth is 7x faster and uses 36% less VRAM at 16K context lengths. Transformers v5 + TRL goes out of memory whilst Unsloth does not. Also the speed up increases with sequence length in this case thanks to our Unsloth's Flex Attention implementation, and our MoE kernels.

1024

275.35

376.99

40.91

43.88

1.4x

6.76%

2048

292.88

696.57

41.83

44.93

2.4x

6.89%

4096

370.30

1785.89

43.68

49.86

4.8x

12.39%

8192

712.33

5226.86

47.43

73.80

7.3x

35.73%

16384

1775.80

**OOM**

55.13

**OOM**

N/A

N/A

On an **NVIDIA B200**, we see **~1.7x speedup and ~35% better memory efficiency with Qwen3-30B-A3B LoRA**, with memory savings improving further at longer sequence lengths.

Qwen3-Next and Coder surprisingly fit on a single B200 GPU in bf16 LoRA.

On H100 GPU, we perform significantly better than the baseline getting up to **1.77x speed up** in training while also saving ~5.3GB when fine tuning at 4K context length. While we seamlessly scale to 8192 context lengths, Transformers v5 + TRL OOMs at 8K. Notice that we use less memory at 8K than the baseline does at 4K so we can keep pushing the context length further.

1024

366.3

628.3

80.88

104.80

1.7x

2.06%

2048

467.0

745.3

80.88

104.81

1.6x

2.57%

4096

711.6

975.5

80.89

104.80

1.4x

5.08%

8192

1376.6

1633.5

80.90

104.81

1.2x

9.17%

16384

3182.2

3407.9

85.53

116.61

1.1x

15.26%

Unsloth achieves **2.6x faster throughput with >15% less VRAM** across all batch sizes for GLM 4.7 Flash. GLM 4.7 Flash is a 30B MoE (3B active parameters) agentic & coding model and employs a configuration similar to the DeepSeek MoE style, featuring 64 routed experts and 1 shared expert. We benchmarked Unsloth MoE training vs the new optimized Transformers v5.

Use our new Colab notebook for GLM 4.7 Flash below:

512

1145.0

2992.1

57.81

60.89

2.6x

6.51%

1024

1298.9

3323.3

58.76

62.55

2.6x

6.22%

2048

1831.9

4119.3

60.09

67.32

2.3x

9.46%

4096

2883.9

5646.1

63.34

76.78

2x

14.83%

In Transformers/PEFT, the usual approach is to **merge the LoRA adapter into the base weight** and then run the MoE computation (especially since MoE often uses `nn.Parameter` instead of `nn.Linear`). The problem is that this merge effectively **materializes the LoRA delta (for all the experts)** `lora_B @ lora_A.t`, which is **very memory-hungry**.

Unsloth avoids that. We previously used the same idea to optimize generic LoRA training and inference, and we’ve now applied it to **MoE + LoRA** as well. The math is identical, so the loss, gradients, and outputs stay the same. The only change is **the order of operations**, made possible by matrix-multiplication associativity. With this reordering, we get major speedups and memory reductions.

Training MoE models in **4-bit** QLoRA isn’t recommended right now because BitsandBytes doesn’t support it. This isn’t specific to Unsloth. For now, use bf16 for LoRA or full fine-tuning.

These optimizations are **enabled by default** when training MoE models with Unsloth (notably Qwen-3 MoE, gpt-oss, and the models mentioned above). You can switch implementations via the `UNSLOTH_MOE_BACKEND` environment variable: either `torch._grouped_mm` **Triton kernels** or a **basic PyTorch for-loop**, depending on compatibility and preference. We default to `grouped_mm` for the best performance and broad support.

LoRA is a parameter-efficient fine-tuning method: instead of updating the full weight matrix, you train a low-rank “adapter” with far fewer parameters, which drastically reduces optimizer memory.

If the original weight has shape **(m, n)**, LoRA adds two trainable matrices with shapes **(m, r)** and **(r, n)**. Their product is **(m, n)**, but you only track optimizer states and gradients for:

- `m*r + r*n` parameters (LoRA) instead of
- `m*n` parameters (full fine-tuning)

For typical MLP layers, `m ≈ 4096, n ≈ 12k, and r ≈ 64`, that’s roughly **~1M LoRA parameters vs ~48M full parameters -** about **~2%,** often with minimal to no accuracy loss.

MoE layers are different because you have **E expert MLPs in parallel**, so any per‑expert change (like adding LoRA) scales across all experts.

Take **Qwen3‑30B‑A3B**: hidden size **m=2048**, intermediate size **n=768**, **E=128** experts with **k=8** activated per token. Per expert:

- `gate_proj` and`up_proj` :`(m, n) = (2048, 768)`
- `down_proj` :`(n, m) = (768, 2048)`

With **LoRA rank r=64**, each projection adds `r*(m+n)=64*(2048+768)=180,224` parameters per expert (≈ `11%` of a `2048×768` matrix). The core issue is that `r/n = 64/768` is large compared to typical MLP setups, for e.g., `r/n = 64/25600` in Qwen3-32B of similar size.

If you materialize this across *all* experts, memory adds up quickly. And since `gate_proj` and `up_proj` are often fused as `gate_up_proj`, you typically materialize both together, roughly doubling the overhead/peak memory.

**In terms of memory, for a sequence length s, E experts and** **k** **chosen, we have the following common for both approaches**

This is where things start to diverge. For peft’s approach we have

For Unsloth’s split LoRA approach, we perform the following operations

Now lets take the case of Qwen3-30B-A3B.

`E = 128, k = 8, m = 2048, n = 768.` Plugging all these in , we get `s < 32K.` 

**In terms of compute, for a sequence length** **s****,** **E** **experts and top** **k** **chosen, we're doing:**

In case of Unsloth split lora that we mentioned, we have

The point till where the Split LoRA from analytical perspective  is better is when `s > Emn/k(m+n)` which is in the order of `16K` tokens for Qwen3-30B-A3B style model.

Finally, some speedups come from **reduced memory traffic**: modern GPUs are often **bandwidth‑bound**, so transferring less data can matter more than FLOPs. A rough speedup estimate is `Emn / [k·s·(m+n)]`, so it depends strongly on **s, E, k**, and the matrix shapes.

Unsloth supports faster MoE training for Qwen, gpt-oss, DeepSeek and GLM models:

- **Qwen3** (Thinking and Instruct): VL • 2507 • Coder
- **gpt-oss** : 20B • 120B • safeguard
- **GLM** : 4.5 • 4.6 • 4.6-Air • 4.7 • 4.7-Flash
- **DeepSeek** : V3 • R1 • V3.1 • V3.2

We may have not uploaded some MoE models, but Unsloth should still support them.

Training Speed including vs Transformers v4

1024

275.35

376.99

2111.18

1.37x

2048

292.88

696.57

2626.80

2.38x

4096

370.30

1785.89

4027.93

4.82x

8192

712.33

5226.86

8513.52

7.34x

16384

1775.80

OOM

OOM

N/A

**Memory VRAM usage**

1024

40.91

43.88

89.75

6.76%

2048

41.83

44.93

90.47

6.89%

4096

43.68

49.86

92.72

12.39%

8192

47.43

73.80

100.3

35.73%

16384

55.13

OOM

OOM

N/A

1. As part of our MoE release, we also made **Gemma-3 now use Flex-Attention** by default, and this works in float16 settings as well (there were infinities which we solved a while back).**Gemma-3 now uses O(N) memory and not O(N^2) memory, and trains >3x faster** (scales even better with context length). Previous Unsloth versions would OOM.

1K

20.1 GB

20.1 GB

0 GB (0%)

2K

21.5 GB

21.1 GB

0.3 GB (2%)

4K

27.7 GB

23.3 GB

4.5 GB (16%)

8K

52.3 GB

27.5 GB

24.8 GB (47%)

16K

OOM

36.0 GB

--

24K

OOM

44.6 GB

--

32K

OOM

53.1 GB

--

48K

OOM

38.4 GB

--

64K

OOM

44.7 GB

--

1. Vision fine-tuning now accepts mixed data of only images and text data!
2. `trl==0.27.1` and`transformers==5.1.0` are supported well - previous coverage was 30% of all our 120 notebooks, but now we have >80% coverage - we plan to make it 100% over the next few days.

To enable faster MoE training, update Unsloth via `pip install --upgrade unsloth unsloth_zoo`

We thank the Hugging Face team for collaborating with us on improving MoE training for the community.

We also sincerely thank the torchao team, especially Vasily Kuznetsov (vkuzo) for working helping us enabling grouped_mm support for float16 to get it work on T4 and backward compatibility with A100.

Last updated

Was this helpful?
