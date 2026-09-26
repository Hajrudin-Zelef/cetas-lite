---
id: collect-250926-servers-hardware/servers-hardware/qwen3-6-how-to-run-locally-unsloth-documentation-3
title: "qwen3-6-how-to-run-locally-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Nvidia", "Unsloth"]
dates: []
keywords: ["blackwell", "gpus", "memory", "moe", "nvfp4", "throughput"]
source: docs/RAG/clean4/qwen3-6-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [479, 489]
sha256: 185e1df21a747c96f4609ff4a25de302435ea0a0ddfb868ff01bfbdc64893fff
---

# qwen3-6-how-to-run-locally-unsloth-documentation

With this, Qwen3.6 27B can now do 140 tokens / s generation with UD-Q2_K_XL and Qwen3.6 35B-A3B 220 tokens / s generation! Some of the throughput numbers are noisy, so don't infer some quants are slower than others.

In terms of average speedup, we see a 1.4x for dense models at draft tokens = 2 and for the MoE around 1.15 to 1.2x.

We do not recommend more than 2 draft tokens because the acceptance rate drops precipitously from 83% to 50% with 4 draft tokens, and the forward passes for MTP become less beneficial.

These results make the trade-off simple: use Dynamic GGUFs for the best balance of memory and quality, use MTP when you want faster generation, and use NVFP4 on Blackwell GPUs for maximum throughput. If you want the easiest path, run the model in Unsloth Studio and keep the recommended defaults.

Last updated

Was this helpful?
