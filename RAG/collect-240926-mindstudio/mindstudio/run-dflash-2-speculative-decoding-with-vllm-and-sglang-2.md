---
id: collect-240926-mindstudio/mindstudio/run-dflash-2-speculative-decoding-with-vllm-and-sglang-2
title: "run-dflash-2-speculative-decoding-with-vllm-and-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "benchmark", "benchmarks", "diffusion", "distribution", "gpu", "memory", "nvidia", "qwen", "speculative decoding", "throughput"]
source: docs/RAG/clean_en/mindstudio/run-dflash-2-speculative-decoding-with-vllm-and-sglang.md
source_anchor: ""
source_lines: [88, 108]
sha256: 182bd788037a07f0a6cc49289633ca0309c0c4e9f0fa840fb2631459eb641afd
---

# run-dflash-2-speculative-decoding-with-vllm-and-sglang

The published benchmarks use a single NVIDIA H200 with FlashAttention 3. The model card doesn’t specify a minimum VRAM requirement below that, but since you’re loading a 27-billion-parameter target model plus a smaller draft model, you need a GPU with substantial high-bandwidth memory, which is why H200-class hardware was used for evaluation.

### Does DFlash 2 change the model’s output quality?

No. Decoding is lossless: greedy decoding produces output identical to what the target model alone would generate, and sampling preserves the target model’s original probability distribution. DFlash 2 only changes how fast that output is produced, not what it is.

### Can I use DFlash 2 with a model other than Qwen3.8-27B?

Based on the model card, this specific checkpoint is trained as a draft model for Qwen/Qwen3.8-27B and isn’t described as compatible with other target models. It functions only inside a speculative decoding server paired with that particular target model.

### How is DFlash 2 different from Qwen3.8’s built-in multi-token prediction?

Qwen3.8 ships with its own seven-token MTP head for speculative drafting. In the published comparisons, DFlash 2 achieved higher acceptance length and higher throughput than that built-in MTP head across every benchmark task and concurrency level tested, thanks to its block-diffusion drafting approach and candidate-path selector.

### Is DFlash 2 available in stable releases of vLLM and SGLang?

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

Not yet. SGLang requires installing from the project’s GitHub source, and vLLM requires installing from a specific open pull request branch. Neither is part of a tagged, stable release at the time of the model’s publication, so setup involves building from source rather than a simple pip install.
