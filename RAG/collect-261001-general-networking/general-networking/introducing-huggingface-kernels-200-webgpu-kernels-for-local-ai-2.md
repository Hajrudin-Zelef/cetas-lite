---
id: collect-261001-general-networking/general-networking/introducing-huggingface-kernels-200-webgpu-kernels-for-local-ai-2
title: "introducing-huggingface-kernels-200-webgpu-kernels-for-local-ai"
domain: general-networking
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["gpu", "benchmark", "embedding", "gpus", "inference"]
source: docs/RAG/collect-261001-general-networking/introducing-huggingface-kernels-200-webgpu-kernels-for-local-ai.md
source_anchor: ""
source_lines: [75, 100]
sha256: 7eb107ab115ae1f8cd17c6abc69c90556c12af5ee53ddcaf7d793bd01906745c
---

# introducing-huggingface-kernels-200-webgpu-kernels-for-local-ai

Some individual wins were much bigger. A particularly difficult bilinear Einsum case (`i,ij,j` with size 4096) ran in 0.136 ms with our kernel versus 1,396 ms with ORT WebGPU: more than **10,000x faster**. A row-wise CumSum over `[256, 4096]` was **301x faster**, at 0.016 ms versus 4.784 ms. These are unusual cases rather than the speedups you should expect everywhere, but they show how much a specialized kernel can help when a general implementation hits a slow path.

We timed the work done on the GPU itself, leaving out setup such as loading kernels, creating sessions, uploading inputs, compiling shaders, and reading outputs back. Very short workloads are naturally harder to measure, and small cases can benefit from the GPU cache, so these numbers are best read as a useful comparison rather than a promise for every application.

They are also results for individual operations, not complete models. Exact performance will change across GPUs and browsers, which is why Fleet is so important for building a broader picture.

We are also working with the ONNX Runtime team to upstream these improvements so they can benefit the broader ONNX Runtime Web ecosystem.

WebGPU performance varies across GPUs, browsers, and drivers, so results from one machine only tell part of the story. Fleet lets anyone run correctness and performance checks in the browser and see how the kernels behave on their hardware.

With consent, each run privately contributes evidence that helps us spot device-specific failures, compare variants, and improve selection rules. The goal is simple: use broad, real-world coverage to make the kernels faster and more reliable for everyone.

The initial 207 kernels are a starting point, not the end state. Publishing kernels independently on the Hub gives us a common place to inspect contracts, compare implementations, reproduce correctness checks, and improve performance without embedding every shader directly into every runtime.

The collection is also part of the Hub's broader kernel ecosystem: on the Kernels page, the WebGPU kernels sit alongside kernels for CUDA, ROCm, Metal, and other platforms, and can be filtered, sorted, and explored like any other artifact on the Hub.

The pieces reinforce one another:

1. Kernel repositories define transparent, versioned operation contracts.
2. `@huggingface/kernels` makes those operations straightforward to load and run from JavaScript.
3. Fleet crowdsources real-world evidence across a much broader range of devices than a conventional benchmark lab can cover.
4. Every contributed run can reveal failures, guide tuning, improve variant selection, and help validate future kernel versions.

This is the low-level foundation for the next steps in our browser inference stack. We are excited to connect these kernels to higher-level model tooling, continue expanding operation coverage, and make fast local inference easier to use across the WebAI ecosystem.

Explore the WebGPU kernel collection, try `@huggingface/kernels`, and join the Fleet to contribute evidence from your device and help us make the kernels better for everyone.
