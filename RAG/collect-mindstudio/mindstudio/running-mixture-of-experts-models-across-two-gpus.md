---
id: collect-mindstudio/mindstudio/running-mixture-of-experts-models-across-two-gpus
title: "Splitting a 122B MoE Model Across an Nvidia and AMD GPU with Vulkan"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Hugging Face", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["amd", "gpu", "moe", "nvidia", "compute", "consumer", "distribution", "gpus", "inference", "inference engine", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/02_mindstudio/running-mixture-of-experts-models-across-two-gpus.md
source_anchor: ""
source_lines: [1, 53]
sha256: a63113d81aaf2223816653661f42e475b92ebb6509cb05ada2b1e753f19e317a
---

# Splitting a 122B MoE Model Across an Nvidia and AMD GPU with Vulkan

## Metadata

- **Source** : https://www.mindstudio.ai/blog/running-mixture-of-experts-models-across-two-gpus
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how to run a single **122B-parameter MoE model** split across mismatched **Nvidia and AMD GPUs** using **Vulkan** and **llama.cpp**. Splitting a model across two vendors means dividing its layers between devices so each holds and processes part of the network, even though Nvidia relies on **CUDA**, AMD relies on **ROCm**, and the two don't interoperate. The workaround is **Vulkan**, a cross-platform graphics and compute API both vendors support, letting a single inference engine like llama.cpp address both cards and load a model too big for either alone.

The practical reason to combine an Nvidia GPU with an AMD APU is memory. Large MoE models often exceed what a single consumer or prosumer GPU can hold. AMD's **Strix Halo** APUs share a unified memory pool of up to **128 GB** between CPU and GPU, letting a mini PC load models that would otherwise require a multi-GPU server, but the integrated GPU is slower than a discrete card for models that fit in less memory. An **Oculink** port — external PCIe running at roughly **63 GB/s** — lets a full-size eGPU (including Nvidia's flagship **RTX Pro 6000 with 96 GB VRAM**) attach to a machine otherwise limited to integrated graphics. Combining both memory pools lets you fit models neither device could hold alone.

Under the hood, transformer weights are organized into layers, and llama.cpp distributes those layers across devices. In the demonstrated setup, roughly **70% of layers** were placed on the Nvidia RTX Pro 6000 and the rest on the AMD Strix Halo iGPU. Every token passes through both sets of layers in sequence, so the GPUs work as pipeline stages rather than in parallel. Because the eGPU attaches over Oculink rather than a full internal PCIe slot, that link becomes the throughput ceiling. Memory allocation for the AMD APU's GPU must be set manually in BIOS (96 GB in the tested configuration); left automatic, the system may not properly allocate memory for the eGPU.

The MoE design explains surprising results. A dense 32B model activates all 32B parameters per token; a 122B MoE may activate only ~10B per token (labeled "**A10B**" on Hugging Face), routing each token to a subset of experts. In tests, the **122B MoE at Q4 (~76.5 GB on disk) ran at 121 tokens/second on the RTX Pro 6000**, noticeably faster than the 32B dense model on the same card. A 35B MoE model ran close to six times faster than an old-style dense model of similar size on the integrated GPU.

The split earns its keep specifically when a model is too large for any single GPU but small enough to fit across two. A **130 GB Q8-quantized model** that fit on neither the Nvidia GPU (96 GB) nor the AMD APU alone became runnable by dividing it — about **88 GB on the RTX Pro 6000 and 34 GB on the Strix Halo APU** — producing **37 to 38 tokens/second**. That's slower than a natively fitting model, but it's the difference between running a very large, higher-quality quantization locally versus not at all. Running **eight concurrent requests** scaled throughput to around **80 tokens/second combined**, with the Nvidia card drawing only about **130-150 watts** and the AMD APU about **76 watts**. Splitting only changes where layers execute, not the math, so output quality depends on the model and quantization, not on hardware distribution.

## Key points

- Vulkan is the bridge letting llama.cpp split one model across Nvidia and AMD, since CUDA and ROCm can't share an inference job.
- A 122B MoE (Q4, ~76.5 GB) ran at 121 tok/s on a single RTX Pro 6000 (96 GB) — faster than a smaller 32B dense model.
- MoE models activate only a fraction of parameters per token (~10B active of 122B, "A10B").
- A 130 GB Q8 model was split ~88 GB Nvidia / 34 GB AMD Strix Halo APU for 37-38 tok/s.
- The split is layer-based and proportional to memory (~70% on the larger GPU); every token passes through both in sequence.
- Oculink (63 GB/s) is the bottleneck for cross-GPU inference; BIOS memory allocation must be manual.
- 8 concurrent requests scaled to ~80 tok/s combined at low power (Nvidia ~130-150 W, AMD ~76 W).

## Technical data / figures

| Item | Value |
|---|---|
| Model split | 122B MoE, Q8 (~130 GB) |
| Active params | ~10B ("A10B") |
| Q4 model size | ~76.5 GB on disk |
| RTX Pro 6000 VRAM | 96 GB |
| 122B MoE Q4 throughput | 121 tok/s |
| Split memory | ~88 GB Nvidia / 34 GB AMD |
| Split throughput (Q8) | 37-38 tok/s |
| 8 concurrent requests | ~80 tok/s combined |
| Power draw | Nvidia ~130-150 W; AMD ~76 W |
| Oculink bandwidth | ~63 GB/s |
| BIOS allocation (tested) | 96 GB to iGPU |
| Cross-vendor API | Vulkan (vs CUDA / ROCm) |

## Why this source matters for the RAG

It documents a concrete dual-vendor local inference technique with precise throughput and memory figures, useful for questions about running models larger than any single GPU. It also explains the MoE active-parameter principle that underpins much of the local AI hardware discussion.
