---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers-2
title: "deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "DeepSeek"]
dates: []
keywords: ["deepseek", "consumer", "gpu", "inference", "memory", "moe", "parameters", "quantization", "qwen", "throughput", "tokens per second"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers.md
source_anchor: ""
source_lines: [60, 83]
sha256: 9ee90ef7224384d3bb6b1ca2d5f5d029f3c873aad7a283960d4772fff24ca601
---

# deepseek-v4-flash-on-one-rtx-3090-real-tokens-per-second-numbers

DeepSeek V4 Flash is a mixture-of-experts large language model from DeepSeek AI. It activates only a subset of its total parameters per token, which makes it far more practical to run on consumer hardware with system RAM offload than an equivalently sized dense model would be.

### Can I run DeepSeek V4 Flash with only 128GB of RAM?

Based on hands-on testing, 128GB was not enough. The model ran successfully once system RAM was increased to 192GB, and something around 156 to 168GB is a more realistic practical minimum depending on your exact configuration and quantization.

### Does RAM speed actually affect performance, or just capacity?

Speed matters, not just capacity. The test system ran DDR4 at 2400MT/s, slower than common current standards, and RAM bandwidth directly limits how fast expert weights can be streamed to the GPU during inference. Faster DDR4 (3200MT/s) or DDR5 memory should meaningfully increase tokens per second.

### Why did the dense Qwen 3.8 27B model fail while the MoE model worked?

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

Dense models activate all their parameters for every token, so they don’t benefit from the same selective-loading tricks that make MoE models offload-friendly. In this case, the Qwen 27B BF16 model also hit a specific software bug in FreeToken’s beta desktop app, throwing a generic error and failing to start entirely.

### What kind of speed can I expect for local mixture-of-experts models on one GPU?

Expect interactive chat speed rather than API-level throughput. Testing showed roughly 8.8 to 11 tokens per second for DeepSeek V4 Flash on a single RTX 3090, varying based on which experts were active and whether the model was accessed through a server backend or a desktop client.
