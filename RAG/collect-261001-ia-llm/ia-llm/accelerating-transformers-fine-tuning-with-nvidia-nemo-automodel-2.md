---
id: collect-261001-ia-llm/ia-llm/accelerating-transformers-fine-tuning-with-nvidia-nemo-automodel-2
title: "From nemo_automodel/components/moe/parallelizer.py"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["moe", "attention", "benchmark", "benchmarks", "consumer", "cost", "deepseek", "fine-tuning", "flash attention", "gpu", "gpus", "inference"]
source: docs/RAG/collect-261001-ia-llm/accelerating-transformers-fine-tuning-with-nvidia-nemo-automodel.md
source_anchor: ""
source_lines: [100, 163]
sha256: 48d9c52cbd8f21acd8cb30eeae245809ea4cc75c35a1be01d1d9d47bfa1967fd
---

# From nemo_automodel/components/moe/parallelizer.py

1. **Expert Parallelism reduces memory pressure.** EP=8 distributes expert weights across GPUs, cutting the per-GPU MoE footprint by 8x. For Qwen3, this drops peak memory from 68.2 GiB to 48.1 GiB (-29%). For Nemotron Nano, it drops from 62.1 GiB to 42.5 GiB (-32%), freeing headroom for larger batch sizes or longer sequences.
2. **DeepEP fuses communication with computation.** Instead of separate AllGather/ReduceScatter collectives for expert routing, DeepEP fuses token dispatch and combines into optimized GPU kernels, overlapping communication with expert computation.
3. **TransformerEngine kernels accelerate core operations.** TE's fused attention, linear layers, and RMSNorm implementations provide consistent speedups over their PyTorch/Flash Attention equivalents across all layer types, not just MoE layers.

One of the most impactful features in Transformers v5 is the experts_implementation parameter, which includes three expert backends:

| Backend | Description | Best for | 
|---|---|---|
| eager | For-loop over selected experts | Debugging, compatibility, and correctness. Also available for v4. | 
| batched_mm | Duplicates expert params, single batched GEMM via torch.bmm | Small inputs, fast with torch.compile. Added for v5 | 
| grouped_mm | Orders tokens by expert, single grouped GEMM via torch.nn.functional.grouped_mm | Training (memory efficient, no param duplication). Added for v5. | 

The grouped_mm backend is the key training optimization: instead of looping over experts one by one, it sorts tokens by their assigned expert and executes a single fused grouped matrix multiplication.

NeMo AutoModel takes this further. For models with custom implementations, it uses DeepEP fused all-to-all dispatch combined with grouped GEMM kernels and TransformerEngine linear layers. The progression looks like:

```
v4 (eager for-loop) → v5 (grouped_mm) → NeMo AutoModel (DeepEP + GMM + TE)
```
In NeMo AutoModel, the expert backend is configured through BackendConfig:

```
from nemo_automodel.components.models.common.utils import BackendConfig
backend = BackendConfig(
    attn="te",           # TransformerEngine attention
    linear="te",         # TransformerEngine linear layers
    experts="torch_mm",  # Grouped expert matmul
    dispatcher="deepep", # DeepEP fused all-to-all
)
```
Transformers v5 also ships an Expert Parallelism path. It shards expert weights across GPUs. The GroupedGemmParallel style loads only each device's local experts, and RouterParallel routes tokens and combines results with an all_reduce. It's neatly built on v5's existing tensor-parallel machinery. Enabling it makes the model's tp_plan return its expert plan, so expert parallelism shares the device budget with data parallelism (ep × dp = world_size). For the single-node 30B benchmarks here, we found plain data-parallel v5 (dp=8, ep=1) to be the fastest v5 configuration, so that's the v5 setup we report.

NeMo AutoModel takes a complementary approach tuned for multi-GPU MoE training. It makes EP its own parallelism dimension, a dedicated moe_mesh alongside (rather than carved from) the data-parallel mesh, using PyTorch's DTensor with Shard(0). Because the expert mesh is orthogonal to data parallelism, the two compose on the same devices. On 8 GPUs NeMo AutoModel runs ep=8 and dp=8 together, so every GPU trains on its own data shard while holding only 1/8 of the experts. Expert weights are physically sharded across GPUs along the expert dimension.

```
# From nemo_automodel/components/moe/parallelizer.py
from torch.distributed.tensor import Shard, distribute_tensor
# Each GPU holds only 1/ep_size of the expert weights
distribute_tensor(param, device_mesh, [Shard(0)])
```
With ep_size=8 on 8 GPUs, each GPU holds only 1/8 of the expert parameters. For a model like Nemotron-3-Nano-30B-A3B with ~55 GiB of expert weights, EP reduces the per-GPU expert footprint from ~55 GiB to ~6.8 GiB, making training possible where FSDP-only approaches run out of memory.

On top of EP, NeMo AutoModel integrates DeepEP that fuses the token routing into optimized GPU kernels, and delivers significant speedups when combined with grouped GEMM for grouped expert computation. In our large-scale MoE benchmarks, DeepEP + grouped GEMM reduced cost per iteration by 47% on the full DeepSeek V3 671B model compared to all-gather + looped expert baselines.

Transformers v5 also introduced a dynamic weight loading system through WeightConverter and WeightRenaming. This enables MoE checkpoint to be stored in fused 3D tensors for more efficient execution. The WeightConverter applies composable operations to transform checkpoint tensors on-the-fly during from_pretrained().

NeMo AutoModel is a direct consumer of this v5 API. Over 20 model types use this mechanism through MODELS_REQUIRING_TENSOR_MERGING, including Mixtral, Qwen2 MoE, Qwen3 MoE, DeepSeek V2/V3, OLMoE, and more. The conversions are fully reversible: save_pretrained() produces standard HF-format checkpoints that any downstream tool can load.

To try NeMo AutoModel, please visit our official documentation page to get started.

For more details, see:

- NeMo AutoModel HuggingFace API Compatibility Guide
- NeMo AutoModel Model Coverage
- NeMo AutoModel Performance Summary
- NeMo AutoModel on HuggingFace

NVIDIA NeMo AutoModel is the natural next step for HuggingFace users scaling up model training. By building directly on Transformers v5, AutoModel provides a zero-friction upgrade path: change one import line and get a model instance that is more than three times as fast.

On Qwen3-30B-A3B and Nemotron 3 Nano 30B-A3B, this delivers 3.4-3.7x higher training throughput with 29-32% less GPU memory compared to the best Transformers v5 configuration. And because true Expert Parallelism shards experts across GPUs, the same path scales up to full fine-tuning a 550B model like Nemotron 3 Ultra across 16 nodes, the regime where Expert Parallelism becomes essential to fit the model in memory. Because NeMo AutoModel checkpoints are standard HF-format safetensors, you can deploy them on inference frameworks like vLLM and SGLang.

The code, configs, and benchmark scripts are all available in the NeMo AutoModel repository.

Core contributors to this work, listed alphabetically by last name: Adil Asif, Hemil Desai, Alexandros Koumparoulis, and Huiying Li.
