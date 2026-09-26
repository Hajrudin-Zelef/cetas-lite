---
id: collect-240926-mindstudio/mindstudio/how-to-run-maple-preview-locally-on-a-mac-mini-m4-2
title: "how-to-run-maple-preview-locally-on-a-mac-mini-m4"
domain: mindstudio
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "agentic", "benchmarks", "gpus", "inference", "inference engine", "license", "mit license", "open-weight", "reasoning", "tokens per second", "tool use"]
source: docs/RAG/clean_en/mindstudio/how-to-run-maple-preview-locally-on-a-mac-mini-m4.md
source_anchor: ""
source_lines: [63, 71]
sha256: 9ac274916bd2d9304e16fb68b919bd8bee43ed37ff834a3108cc5c8296cca3bb
---

# how-to-run-maple-preview-locally-on-a-mac-mini-m4

The official Transformers implementation depends on Triton and FlashAttention, which target CUDA GPUs. The benchmarked 218 tokens per second on a Mac mini M4 comes from a separate on-device runtime, not the default Transformers code path, so Mac users need an Apple Silicon compatible inference engine to reproduce that speed.

### Is Maple-Preview good for agentic tasks or tool use?

Not particularly, at least in this preview. DeepGrove notes it received minimal post-training for agentic tasks and only small-scale general reinforcement learning. It’s tuned for raw reasoning benchmarks like AIME, HMMT, GPQA-D, and code reasoning rather than multi-step agent workflows.

### What license is Maple-Preview released under?

It’s released under the MIT license, meaning it’s free to use, modify, and redistribute, including for commercial purposes, without the restrictions some other open-weight models attach.
