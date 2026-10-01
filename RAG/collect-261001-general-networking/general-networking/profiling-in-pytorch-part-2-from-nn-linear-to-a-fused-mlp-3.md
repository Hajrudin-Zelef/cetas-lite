---
id: collect-261001-general-networking/general-networking/profiling-in-pytorch-part-2-from-nn-linear-to-a-fused-mlp-3
title: "bias=True would truly emulate the multiplication and addition"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["agent", "attention", "cost", "gpu", "hbm", "latency", "parameters"]
source: docs/RAG/collect-261001-general-networking/profiling-in-pytorch-part-2-from-nn-linear-to-a-fused-mlp.md
source_anchor: ""
source_lines: [196, 244]
sha256: 29899decde0becaeebcf1af9cd402ab56d38be6e124150c7bcb1ab72d2d1ad0a
---

# bias=True would truly emulate the multiplication and addition

```
from kernels import get_kernel
kernels_layers = get_kernel("kernels-community/liger-kernels", version=1).layers
kernels_geglu_mlp = kernels_layers.LigerGEGLUMLP(Config()).to(device, dtype=torch.bfloat16).eval()
```
The full script is here: `03_kernels_mlp.py`.

```
uv run 03_kernels_mlp.py --batch 64 --seq 128 --dim 768 --hidden 3072
uvx trace-util -f traces -b <hf_uname>/traces
```
Figure 12 shows the profile for the `LigerGEGLUMLP` layer using the Liger kernels from the Hub.

Writing kernels in Triton or CUDA is one problem and *shipping* them is another. The kernel has to be compiled for your exact combination of GPU architecture, CUDA version, and PyTorch version. This is the step that usually breaks ("works on my machine", missing `nvcc`, wrong Triton version).

The `kernels` library moves that build step off your machine. `get_kernel("kernels-community/liger-kernels", version=1)` downloads a **pre-built, version-pinned** kernel package from the Hugging Face Hub and caches it locally (here under `~/.cache/...kernels-community--liger-kernels`). The benefits are:

- The kernels are compiled once, in CI, for many architectures and version combinations. You download the right binary instead of compiling it yourself.
- `version=1` pins the exact build, so everyone running your script gets the same kernel. There is no "it got slower after I updated a package".
- The package exposes a `.layers` attribute with drop-in`nn.Module` s (like`LigerGEGLUMLP` ). You swap your module for theirs and nothing else in your model changes.

When we say "tuned", we mean two concrete things, and both are visible in the trace.

1. **The fusion is baked in.** The`LigerGEGLUMLP` forward is`down_proj(LigerGELUMulFunction.apply(gate_proj(x), up_proj(x)))` . The`LigerGELUMulFunction` runs a single Triton kernel,`_geglu_tanh_forward_kernel` , that computes`gelu(gate) * up` in one pass. This is exactly what we saw from`torch.compile` , where the intermediate never makes a round-trip through HBM. We get it here**without the compiler** , as shown in Figures 13 and 14 (no Dynamo guards, no compile latency, no recompilation risk).
2. **The launch parameters were chosen for the hardware.** The kernel does not guess its block size at random. Liger's`calculate_settings` picks them from the column count.

It is worth being honest about the trade-off here, because the raw numbers can be misleading. The Liger kernel runs in **92.8 µs**, while Inductor's fused kernel from the compile run was **89.4 µs**. At first glance the hand-written kernel looks slightly slower, but that comparison hides the cost that makes it worthwhile.

`torch.compile` specializes for a **static shape**. Inductor's `89.4 µs` kernel is fast precisely because it was generated for *this exact* `[8192, 3072]` problem. Change the batch size, the sequence length, or the hidden dimension, Dynamo re-traces, and you pay the compile cost all over again to get a new specialized kernel.

So the real choice is not "slow human kernel vs fast compiled kernel". It is **a fast generic kernel vs a kernel specialized for one particular input shape**. The Liger kernel takes one set of launch parameters and runs them for *any* shape with no recompilation. It gives up the last few microseconds that per-shape specialization would buy, in exchange for being robust to changing shapes.

The table below collects what each step changed on the GPU and what it left untouched.

| Setup | What changed | What stayed the same | 
|---|---|---|
| Eager `nn.Linear` | Baseline: bias add is already folded into the GEMM epilogue ( `addmm` ), so it is*one* cuBLAS kernel, not a matmul plus an add | — | 
| Compiled `nn.Linear` | A few CPU dispatch ops (the `aten::t` view bookkeeping) disappear | Same single cuBLAS GEMM kernel, byte-for-byte. Compile has nothing to fuse | 
| Eager MLP | 5 GPU kernels: 3 GEMMs + a GeLU + a mul. The `[8192, 3072]` intermediate makes a full round-trip through HBM | Each GEMM is still the same bias-free cuBLAS kernel as a standalone linear | 
| Compiled MLP | GeLU + mul + reshape collapse into **one** fused Triton kernel; the intermediate stays in registers. Pays compile pre-ops (Dynamo, guards) | The 3 GEMMs are untouched with identical cuBLAS kernel names | 
| Liger MLP | Same fusion, but baked into a hand-written Triton kernel with hardware-tuned launch params with **no** Dynamo, guards, or compile latency | The 3 GEMMs are still the same cuBLAS kernels | 

If there is one habit to carry forward, it is the one we practiced before every trace: **guess first, then look.** State what you expect the trace to contain, open it, and treat any mismatch as the most interesting thing on the screen.

This was the second stop in the **Profiling in PyTorch** series. In the next post we will keep climbing the ladder, moving from this MLP block towards the attention block and, eventually, a full model.

Thanks to Noe Flandre and Pedro Gabriel Gengo Lourenço for their reviews on the early draft of the post!

The blog post was polished using an LLM. This in no way means that we have let an agent run in the background and let it generate the blog. Some of us in the team are non-english speakers and think LLMs (which are mostly trained in the English Language) can rectify silly grammar mistakes or rephrase sentences that sound less intimidating and cleaner. Hope this helps with the idea of "why should I read, if this was LLM generated". 🤗
