---
id: collect-261001-general-networking/general-networking/profiling-in-pytorch-part-2-from-nn-linear-to-a-fused-mlp-2
title: "bias=True would truly emulate the multiplication and addition"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["decode", "gpu", "hbm", "memory", "research"]
source: docs/RAG/collect-261001-general-networking/profiling-in-pytorch-part-2-from-nn-linear-to-a-fused-mlp.md
source_anchor: ""
source_lines: [101, 195]
sha256: 6b0be1ad6ed9479a51d90136ea44038e49def061564326650b54c6cc24cf79aa
---

# bias=True would truly emulate the multiplication and addition

```
cutlass_80_wmma_tensorop_bf16_s161616gemm_bf16_32x32_32x1_tn_align8
```
If no transpose kernel ran, who taught the GEMM to read the weight matrix in transposed order? The answer is in the kernel's name. Look at the suffix:

```
cutlass_80_wmma_tensorop_bf16_s161616gemm_bf16_32x32_32x1_tn_align8
                                                          ^^
```
That `tn` is the layout descriptor. cuBLAS and CUTLASS precompile a *separate kernel binary* for each combination of input layouts.

`n` (non-transposed) and `t` (transposed) describe how a kernel walks its input during the inner loop. The dispatcher's job is to look at the input strides, decide which suffix combination matches, and pick the right precompiled kernel.

The kernel name in a profiler trace is a hash dump of the kernel's identity. If two runs show the same kernel name, the GPU is doing the same work. If they differ (e.g., `_tn_` vs `_nn_`, `bf16` vs `fp16`, or `s16816gemm` vs `s161616gemm`) then the GPU is doing different work, and the dispatcher took a different branch. Learning to read this name is one of the most useful habits when comparing traces.


In this section, we will profile a Multilayer Perceptron (MLP). To make this more interesting, we will profile a feed-forward network with the GeGLU activation variant (which is quite heavily used in practice). This is also our way of paying tribute to one of the greatest lines ever written in the history of deep learning research (Figure 6).

```
class SimpleGeGLUMLP(nn.Module):
    def __init__(self, dim, hidden):
        super().__init__()
        self.gate_proj = nn.Linear(dim, hidden, bias=False)
        self.up_proj = nn.Linear(dim, hidden, bias=False)
        self.down_proj = nn.Linear(hidden, dim, bias=False)
    def forward(self, x):
        g = self.gate_proj(x)
        u = self.up_proj(x)
        h = F.gelu(g, approximate="tanh")
        m = h * u
        y = self.down_proj(m)
        return y
```
You will find the entire script here: `03_simple_mlp.py`. Execute it like so:

```
uv run 03_simple_mlp.py --batch 64 --seq 128 --dim 768 --hidden 3072
uvx trace-util -f traces -b <hf_uname>/traces
```
Before we open the trace, let's think together about what we should expect to see. The `forward` function does a fair amount of computation, but most of it is already familiar to us.

We should expect three `aten::linear` dispatches, one for each `nn.Linear` layer. We should also expect two pointwise kernel launches, one for the GeLU and one for the multiplication. Forming this expectation before looking is the single most useful habit in the profiling journey: you read the trace to *confirm or break* a guess, not to form one from scratch.

From Figure 7 we can pat ourselves on the back, as our intuition was correct. Per forward pass (one `mlp_fwd`), the GPU runs exactly 5 kernels. Figure 8 highlights the "occupancy query" as seen in the CPU lane for the linear projection layers.

| Op | CPU op | GPU kernel | launches | 
|---|---|---|---|
| `gate_proj` | `aten::linear` | `ampere_bf16_s16816gemm_bf16_128x128_...` | occupancy query + cudaLaunchKernel | 
| `up_proj` | `aten::linear` | `ampere_bf16_s16816gemm_bf16_128x128_...` | occupancy query + cudaLaunchKernel | 
| `gelu` | `aten::gelu` | `vectorized_elementwise_kernel<4, GeluCUDAKernelImpl...>` | cudaLaunchKernel | 
| `h * u` | `aten::mul` | `vectorized_elementwise_kernel<4, ...MulFunctor...>` | cudaLaunchKernel | 
| `down_proj` | `aten::linear` | `ampere_bf16_s16816gemm_bf16_128x256_...` | occupancy query + cudaLaunchKernel | 

The three GEMMs each do an extra `cudaOccupancyMaxActiveBlocksPerMultiprocessor` call before the launch. We have a separate section on this in Part 1, you can find it here. That is cuBLAS sizing the grid. The pointwise ops (GeLU and mul) launch directly, with no occupancy query. So "a linear" is actually query + launch, while "a pointwise op" is just launch.

The `aten::t`, `aten::transpose`, `aten::reshape`, `aten::view`, `aten::as_strided`, and `aten::_unsafe_view` ops launch zero kernels. They show `0.000us` of CUDA time in the table (Figure 9) because they only rewrite tensor metadata (shape and stride) on the CPU. A reader scanning the table sees around six op names per linear, but only one of them (`mm`) ever reaches the GPU.

The MLP flattens `[batch, seq, dim]` to `[batch * seq, dim]` for the matmul. In our command-line invocation we used 64 for `batch` and 128 for `seq`, so that's where the `8192` (`batch * seq = 64 * 128`) below comes from.

From the trace:

| Linear | `aten::mm` input dims | M·K·N | cuBLAS kernel | avg CUDA | 
|---|---|---|---|---|
| `gate_proj` | `[8192,768] x [768,3072]` | `8192·768·3072` | `…128x128…stages_32x5_tn` | 0.19ms | 
| `up_proj` | `[8192,768] x [768,3072]` | `8192·768·3072` | `…128x128…stages_32x5_tn` | 0.19ms | 
| `down_proj` | `[8192,3072] x [3072,768]` | `8192·3072·768` | `…128x256…stages_64x3_tn` | 0.17ms | 

All three GEMMs have the same FLOP count, `2·8192·768·3072 ≈ 38.7 GFLOP` each, yet `down_proj` is about `10%` faster. Same work, different shape (`N=768` instead of `3072`), so cuBLAS picks a different tile (`128×256`, with a deeper `stages_64x3` pipeline) that gets better reuse for that shape.

If you want to learn more about tiling in depth, here is a great resource to get started with.


This is exactly why the table had two GEMM rows (Figure 9): the `128x128` row is gate+up and the `128x256` row is down.

Before compiling the `forward` method and visualizing it, let's do the mental exercise again of asking ourselves what we expect to see in the trace. This is a fun experiment, and an important one to repeat every time you profile something yourself. Always build on your intuition, and the moment something does not match, stop and figure out why.

```
uv run 03_simple_mlp.py --batch 64 --seq 128 --dim 768 --hidden 3072 --compile
uvx trace-util -f traces -b <hf_uname>/traces
```
In eager mode, each `nn.Linear` was expanded into a chain of dispatcher ops (`aten::linear` → `aten::t` → `aten::transpose` → `aten::matmul` → `aten::reshape` → `aten::mm`). Those are the high-level wrappers that ATen walks through before reaching the real GEMM. `torch.compile` removes that chain.

By the time the compiled graph runs, there is no linear, no matmul, no transpose or reshape and those metadata ops were folded into how `mm` is called. We can see three bare `aten::mm` external calls (Figure 10). The proof that it is the same GEMM is that the kernel names are byte-for-byte identical to eager: `...128x128...stages_32x5_tn` for gate and up, and `...128x256...stages_64x3_tn` for down.

This is the headline of the whole compile lesson. The two eager pointwise kernels (GeLU and mul) plus a reshape collapsed into one kernel, `triton_poi_fused__unsafe_view_gelu_mul_0` (Figure 11). Let's decode the name:

- `triton` : generated by Inductor's Triton backend (not cuBLAS, not ATen).
- `poi` : pointwise (Inductor tags pointwise kernels`poi` , reductions`red` , and persistent reductions`per` ).
- `fused__unsafe_view_gelu_mul` : the ops it merged: the`_unsafe_view` (reshape), the GeLU, and the mul.
- `0` : the unique id within the graph.

Why is this a win? In eager mode, the intermediate `h = gelu(g)` is a full `[8192, 3072]` bf16 tensor (around 50 MB) that the GeLU kernel writes to HBM and the mul kernel immediately reads back. Fusion keeps it in registers (memory that resides inside the chip and are closer than the HBM). The Triton kernel reads `g` and `u` once, computes `gelu(g) * u`, and writes the result once. One whole round trip of the intermediate through global memory is gone.

So far we have let PyTorch (eager) and the compiler (`torch.compile`) pick our kernels. Now we plug in a kernel that a human expert wrote and tuned by hand. We use the `LigerGEGLUMLP` layer, that we can easily fetch from the Hugging Face Hub with the `kernels` library.

