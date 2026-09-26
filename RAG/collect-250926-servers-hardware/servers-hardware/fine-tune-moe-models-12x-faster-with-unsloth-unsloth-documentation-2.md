---
id: collect-250926-servers-hardware/servers-hardware/fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation-2
title: "fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Unsloth", "Z.ai"]
dates: []
keywords: ["moe", "attention", "compute", "deepseek", "fine-tuning", "glm", "gpus", "lora", "memory", "parameters", "qwen", "training"]
source: docs/RAG/clean4/fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation.md
source_anchor: ""
source_lines: [277, 506]
sha256: 8dfde8085dba5fe1349654b3237ee197b9ea7d3bffce72652536a476ad48fb03
---

# fine-tune-moe-models-12x-faster-with-unsloth-unsloth-documentation

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
