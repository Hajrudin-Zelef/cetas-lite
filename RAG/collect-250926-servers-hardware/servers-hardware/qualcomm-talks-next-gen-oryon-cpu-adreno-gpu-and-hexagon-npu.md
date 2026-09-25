---
id: collect-250926-servers-hardware/servers-hardware/qualcomm-talks-next-gen-oryon-cpu-adreno-gpu-and-hexagon-npu
title: "qualcomm-talks-next-gen-oryon-cpu-adreno-gpu-and-hexagon-npu"
domain: servers-hardware
role: reference
task: reference
actors: ["Qualcomm"]
dates: []
keywords: ["gpu", "accelerator", "agent", "agentic", "compute", "inference", "int4", "memory", "mixture of experts", "moe", "parameters", "prefill"]
source: docs/RAG/clean4/qualcomm-talks-next-gen-oryon-cpu-adreno-gpu-and-hexagon-npu.md
source_anchor: ""
source_lines: [1, 42]
sha256: 83e0283c97444587d2d6d5404b9c248242f47b997a920e563d25043e88c52255
---

# qualcomm-talks-next-gen-oryon-cpu-adreno-gpu-and-hexagon-npu

Ahead of its Snapdragon Summit later this month, Qualcomm disclosed the Oryon CPU, Adreno GPU, and Hexagon NPU for its next premium mobile Snapdragon platform. The NPU adds an Element Accelerator and larger shared memory, positioning it alongside the previously detailed Oryon CPU with 5 GHz Prime cores and FlexCache, and the Adreno GPU with Matrix Cores. Since it is the agentic AI era, Qualcomm is talking about the platform in terms of how AI runs on it in a bit of a teaser.


## Qualcomm Oryon CPU

The Oryon CPU orchestrates agentic work, coordinating planning, tool calls, and accelerator tasks. Our Computex 2025 coverage anticipated that newer mobile Oryon designs would move into PC chips, though the current architecture here remains mobile-focused.

The two Prime cores in the next-generation premium Snapdragon platform are rated at 5 GHz, a figure the company says marks the first mobile CPU to reach that frequency. Qualcomm attributes the design to its custom CPU microarchitecture, implementation, and subsystem.

FlexCache is a dynamically allocated cache pool shared by heterogeneous CPU cores on the new platform.

Prime cores can draw on the entire pool as workloads demand, keeping larger working sets cached and reducing system-memory accesses. Hopefully we will get more CPU details soon.

## Qualcomm Adreno GPU

The Adreno GPU comprises three slices clocked at 1.45 GHz, a command processor, and one 18 MB Adreno High Performance Memory (HPM) block. HPM serves as local graphics storage for working data such as tiles and frame buffers, reducing traffic to system memory.

Qualcomm says this delivers a 12% power improvement over the Snapdragon 8 Elite Gen 5 baseline.

Each of the three GPU slices contains Matrix Cores, which bring dedicated matrix and AI processing into the graphics pipeline. Adreno High Performance Memory (HPM) keeps working data nearby, reducing reliance on external memory and improving efficiency.

Neural Fusion provides AI rendering and super-resolution, integrating with Unity and Unreal Engine upscaling frameworks.

With Neural Fusion enabled, Qualcomm reports a 40 percent power savings in its internal Dragon Alley demo.

## Qualcomm Hexagon NPU

The NPU features a new Element Accelerator for transformer operations, plus vector and scalar extensions that handle AI math and agent decision or routing tasks. It supports context lengths up to 32K and includes KV-cache acceleration.

Shared memory grows by 50 percent, according to Qualcomm, though the company does not disclose the absolute capacity. Positioned alongside the tensor, vector, scalar, and element compute regions, this block keeps model state, context, and KV-cache near the accelerators, thereby reducing external-memory movement.

For INT4 models, Qualcomm reports up to 50 percent prefill uplift on its next-generation premium mobile Snapdragon platform versus the Snapdragon 8 Elite Gen 5. Prefill processes the input prompt before the model generates subsequent tokens.

Mixture of Experts (MoE) models up to 30 billion total parameters run on the Hexagon NPU, with roughly 3 billion active parameters routed per token generation step. Flash-to-memory expert management and caching handle the model.

The Sensing Hub captures voice input and routes it to the Oryon CPU, which orchestrates task execution.

Inference workloads are dispatched to the GPU or NPU, with a cloud path available for offloading.

## Final Words

The common thread is keeping data and computation close to the processing resources that use them, reducing reliance on slower system memory. We expect fairly substantial gains when this line finally makes its way into products. Hopefully we get to see them soon and also get a lot more detail on the chips.
