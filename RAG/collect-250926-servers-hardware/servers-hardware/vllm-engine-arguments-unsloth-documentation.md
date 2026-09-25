---
id: collect-250926-servers-hardware/servers-hardware/vllm-engine-arguments-unsloth-documentation
title: "vllm-engine-arguments-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "Unsloth", "vLLM"]
dates: []
keywords: ["vllm", "attention", "fp8", "gpu", "gpus", "inference", "kv cache", "llama", "lora", "memory", "quantization", "throughput"]
source: docs/RAG/clean4/vllm-engine-arguments-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 87]
sha256: 8f38196b9da7352524d54b0e84294854c7c35f7580f79ede69a7876dc11695e8
---

# vllm-engine-arguments-unsloth-documentation

vLLM engine arguments, flags, options for serving models on vLLM.

**--gpu-memory-utilization**

Default 0.9. How much VRAM usage vLLM can use. Reduce if going out of memory. Try setting this to 0.95 or 0.97.

**--max-model-len**

Set maximum sequence length. Reduce this if going out of memory! For example set **--max-model-len 32768**

**--quantization**

Use fp8 for dynamic float8 quantization. Use this in tandem with **--kv-cache-dtype**

**--kv-cache-dtype**

Use `fp8` for float8 KV cache to reduce memory usage by 50%.

**--port**

Default is 8000. How to access vLLM's localhost ie http://localhost:8000

**--api-key**

Optional - Set the password (or no password) to access the model.

**--tensor-parallel-size**

Default is 1. Splits model across tensors. Set this to how many GPUs you are using - if you have 4, set this to 4. 8, then 8. You should have NCCL, otherwise this might be slow.

**--pipeline-parallel-size**

Default is 1. Splits model across layers. Use this with **--pipeline-parallel-size**

**--enable-lora**

Enables LoRA serving. Useful for serving Unsloth finetuned LoRAs.

**--max-loras**

How many LoRAs you want to serve at 1 time. Set this to 1 for 1 LoRA, or say 16. This is a queue so LoRAs can be hot-swapped.

**--max-lora-rank**

Maximum rank of all LoRAs. Possible choices are `8`, `16`, `32`, `64`, `128`, `256`, `320`, `512`

**--dtype**

Allows `auto`, `bfloat16`, `float16` Float8 and other quantizations use a different flag - see `--quantization`

**--tokenizer**

Specify the tokenizer path like `unsloth/gpt-oss-20b` if the served model has a different tokenizer.

**--hf-token**

Add your HuggingFace token if needed for gated models

**--swap-space**

Default is 4GB. CPU offloading usage. Reduce if you have VRAM, or increase for low memory GPUs.

**--seed**

Default is 0 for vLLM

**--disable-log-stats**

Disables logging like throughput, server requests.

**--enforce-eager**

Disables compilation. Faster to load, but slower for inference.

**--disable-cascade-attn**

Useful for Reinforcement Learning runs for vLLM < 0.11.0, as Cascade Attention was slightly buggy on A100 GPUs (Unsloth fixes this)

For example to host Llama 3.3 70B Instruct (supports 128K context length) with Float8 KV Cache and quantization, try:

To enable LoRA serving for at most 4 LoRAs at 1 time (these are hot swapped / changed), first set the environment flag to allow hot swapping:

See our LoRA Hot Swapping Guide for more details.

Last updated

Was this helpful?
