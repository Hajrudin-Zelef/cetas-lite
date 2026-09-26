---
id: collect-240926-mindstudio/mindstudio/what-is-speculative-decoding-how-draft-models-speed-up-llms-2
title: "what-is-speculative-decoding-how-draft-models-speed-up-llms"
domain: mindstudio
role: reference
task: reference
actors: ["SGLang", "vLLM"]
dates: []
keywords: ["gpu", "inference", "llama", "llama.cpp", "memory", "research", "sglang", "speculative decoding", "vllm"]
source: docs/RAG/clean_en/mindstudio/what-is-speculative-decoding-how-draft-models-speed-up-llms.md
source_anchor: ""
source_lines: [69, 97]
sha256: be73d67f72d593f92a0bfb70ba18320eed9f05935ec9c46ff5dcecf8d9d44e91
---

# what-is-speculative-decoding-how-draft-models-speed-up-llms

It’s already supported in mainstream inference engines. The video’s creator noted that beyond SGLang, speculative decoding support also extends to projects like vLLM and llama.cpp, which suggests it’s becoming a standard feature rather than a niche research trick.

The main limitation is that gains vary by workload. Speculative decoding shines when the draft model can reliably predict long stretches of the big model’s output, which tends to happen with more predictable or structured generations. Highly creative or unpredictable text may see smaller speedups, since the draft model’s guesses will be rejected more often.

## Frequently Asked Questions

### Does speculative decoding change the model’s output?

No. The large model still verifies and effectively generates every accepted token itself. Speculative decoding only changes how many forward passes are needed to produce that output, not the text produced.

### What is a draft model?

A draft model is a smaller, faster model used to propose likely next tokens ahead of time. It doesn’t need to be as accurate as the main model, since its guesses get checked and can be discarded cheaply if wrong.

### How is DFlash 2 different from standard speculative decoding?

## One coffee. One working app.

You bring the idea. Remy manages the project.

Standard speculative decoding drafts tokens somewhat sequentially and typically proposes one top candidate per position. DFlash 2 predicts all draft positions in parallel and keeps a pool of multiple candidate tokens per position, then uses a path selection step to pick the most coherent sequence through them.

### Does speculative decoding require more GPU memory?

Yes, some. Running a draft model alongside the base model adds VRAM overhead, though draft models are typically much smaller than the base model, so the added memory footprint is relatively modest compared to the base model’s own requirements.

### Is speculative decoding available in popular inference tools?

Yes. It’s supported in engines such as SGLang, and according to the video’s creator, similar support exists in vLLM and llama.cpp, making it accessible without custom infrastructure.
