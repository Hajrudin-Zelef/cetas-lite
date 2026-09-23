---
id: collect-mindstudio/mindstudio/gmktec-evo-x3-vs-evo-x2-pricing
title: "GMKtec EVO X3 vs EVO X2: Which Strix Halo Mini PC to Buy?"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["agent", "amd", "benchmark", "cost", "gpu", "gpus", "inference", "memory", "moe", "nvidia", "pricing", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/gmktec-evo-x3-vs-evo-x2-pricing.md
source_anchor: ""
source_lines: [1, 49]
sha256: aed0ae2f2690bbbe14e1ffa59ab3ed5d8dc08ca957cd24ad6908604d6220b536
---

# GMKtec EVO X3 vs EVO X2: Which Strix Halo Mini PC to Buy?

## Metadata

- **Source** : https://www.mindstudio.ai/blog/gmktec-evo-x3-vs-evo-x2-pricing
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article compares the **GMKtec EVO X3** and **EVO X2** Strix Halo mini PCs on price, ports, and Oculink eGPU support. Both machines run the same **AMD Strix Halo** chip with the same **128 GB of shared unified memory**, so raw CPU and GPU performance is identical; token-per-second speeds on models that fit in that memory are essentially the same. The real difference is one port: the **EVO X3 adds an Oculink connector** on the back, a PCIe link that lets you plug in an external desktop GPU. The EVO X2 lacks Oculink but comes with more onboard ports and a chassis that lays flat on a desk. The X3 costs about **$100 more** (the EVO X2 is listed at **$3,699**) and sits more awkwardly on its stand.

Strix Halo is AMD's APU architecture combining CPU and GPU on one die sharing a single memory pool of up to 128 GB, most of which can be allocated to the GPU for running LLMs or image generation. GMKtec's original EVO X2 was one of the first mini PCs to bring this AI-capable unified memory to a compact desktop box, beating even Nvidia's DGX Spark to market. The built-in GPU, while generous on memory, isn't as fast per token as a discrete desktop GPU. Oculink — a PCIe slot with an external connector running at **63 GB/s**, the same class of connection a desktop GPU uses internally — addresses that.

The key finding is that an external GPU only helps when the model fits in that GPU's VRAM. With an **RTX 5080 (16 GB)** over Oculink, a model that fit ran about **three times faster** than the built-in GPU, but once a model exceeded 16 GB, throughput collapsed to roughly **1.6 tokens/second**, while the Strix Halo APU kept producing around **11 tokens/second** on the same model. A larger **RTX Pro 6000 (96 GB VRAM)** changed the picture: a **32B dense model ran at ~69 tokens/second** (about six times the built-in GPU), and a **122B MoE model** (only ~10B active) ran even faster at about **121 tokens/second**.

The article also documents cross-vendor operation: Nvidia (CUDA) and AMD (ROCm) stacks don't normally interoperate, but **Vulkan** lets both GPUs run inference on the same model simultaneously. Tools like **LM Studio** let you choose CUDA, ROCm, or Vulkan as the runtime. In one demo, a single agent workload ran across both GPUs — the Nvidia card at ~87 tokens/second and the AMD chip at ~24 tokens/second in parallel. The most extreme test split a **122B model at Q8 (~130 GB on disk)** across both GPUs (~70% of layers on Nvidia, rest on AMD), achieving **~37-38 tokens/second**, with the Oculink link as the bottleneck. Setup wasn't plug-and-play: the mini PC refused to boot with the Pro 6000 attached, requiring booting without the GPU then hot-plugging it, plus manual BIOS memory allocation to the AMD GPU.

## Key points

- EVO X2 and EVO X3 share the same Strix Halo APU and 128 GB unified memory; token speeds are essentially identical.
- The EVO X3 adds an Oculink port (63 GB/s PCIe) enabling a full external GPU — unique among Strix Halo mini PCs.
- The EVO X3 costs about $100 more than the EVO X2 ($3,699); the X2 has more ports and a sturdier lay-flat chassis.
- An eGPU helps only while the model fits in its VRAM; beyond that, throughput collapses versus the built-in APU.
- A 96 GB RTX Pro 6000 runs a 32B dense model at ~69 tok/s and a 122B MoE at ~121 tok/s.
- Vulkan enables splitting one model across Nvidia and AMD GPUs simultaneously (~37-38 tok/s for a 122B Q8 model).

## Technical data / figures

| Item | Value |
|---|---|
| Shared memory | 128 GB unified |
| EVO X2 price | $3,699 |
| EVO X3 premium | ~$100 more |
| Oculink speed | 63 GB/s |
| RTX 5080 (16 GB) fit model | ~3× built-in GPU |
| RTX 5080 overflow model | ~1.6 tok/s (vs ~11 tok/s on APU) |
| RTX Pro 6000 (96 GB), 32B dense | ~69 tok/s (~6× built-in) |
| RTX Pro 6000, 122B MoE (~10B active) | ~121 tok/s |
| Dual-GPU split, 122B Q8 (~130 GB) | ~37-38 tok/s |
| Parallel agent demo | Nvidia ~87 tok/s + AMD ~24 tok/s |
| Cross-vendor runtime | Vulkan (LM Studio backend option) |

## Why this source matters for the RAG

It provides concrete, benchmark-backed guidance on Strix Halo mini PCs and eGPU expansion for local LLMs, including real throughput numbers and the VRAM-fit cliff. This is valuable for hardware recommendation and cost/performance questions in the local AI domain.
