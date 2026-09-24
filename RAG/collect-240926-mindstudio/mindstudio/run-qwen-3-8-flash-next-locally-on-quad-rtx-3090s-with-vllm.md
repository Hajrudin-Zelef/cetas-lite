---
id: collect-240926-mindstudio/mindstudio/run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm
title: "run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "vLLM"]
dates: []
keywords: ["qwen", "vllm", "agent", "agentic", "agents", "attention", "consumer", "context window", "gpu", "gpus", "int4", "kv cache"]
source: docs/RAG/clean_en/mindstudio/run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm.md
source_anchor: ""
source_lines: [1, 93]
sha256: c68ce6140850c12e357ce2518a1e0e98d2eebed5346b31d55a9e77379596b76a
---

# run-qwen-3-8-flash-next-locally-on-quad-rtx-3090s-with-vllm

<!-- source: https://www.mindstudio.ai/blog/qwen-3-8-flash-next-local-agent-setup -->

## What is Qwen 3.8 Flash Next?

Qwen 3.8 Flash Next is a vision-capable model from Alibaba’s Qwen team, listed on Hugging Face under `Qwen/Qwen3.8-Flash-Next` with an `image-text-to-text` pipeline and the `qwen4_exp` architecture tag. The published safetensors weights ship across 131 files, which points to a large full-precision footprint, but the model becomes practical on consumer hardware once quantized. The version used for local agentic work is an INT4 (W4A16) quant, produced by Vinnie AI, that keeps vision support intact while shrinking VRAM requirements enough to fit on a multi-GPU 3090 rig. That combination, strong output quality at Q4 plus vision, is why it’s drawing attention as a faster alternative to smaller 27B-class local models.

## TL;DR

- **Qwen 3.8 Flash Next** runs locally as an INT4 (W4A16) quant from Vinnie AI that retains vision support, unlike some earlier quant options.
- A **quad RTX 3090 setup** (96GB total VRAM) can serve the model with a full 131,072-token context window using vLLM.
- Hitting good speeds requires **128GB of system RAM** as a practical minimum, since the engram load pushes system memory usage close to 100GB during startup.
- The working vLLM launch config uses `max-model-len 131072` ,`max-num-seqs 2` ,`max-num-batched-tokens 2048` ,`gpu-memory-utilization 0.96` ,`cuda-graph-mode full` , and`enable-prefix-caching` .
- **Async scheduling should stay disabled** ; enabling it can add roughly 10 tokens per second but introduces crashes without an additional patch set.
- Agentic tool calling needs `enable-auto-tool-choice` with the`qwen3xml` tool-calling parser to get reliable function-call behavior out of a Hermes-style agent.
- Real-world throughput on a quad-3090 rig lands around 55 to 61 tokens per second at large context depths, roughly double what earlier setups on the same hardware achieved with llama.cpp.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

## Why run this model instead of something smaller?

Local LLM users running 3090s have spent a long time choosing between speed and capability. Smaller dense models (the 27B-class options many people defaulted to) run fast but plateau on quality. Larger mixture-of-experts style models are strong but often too slow or too VRAM-hungry to be useful for agentic loops, where a model has to call tools, wait, re-read context, and generate again many times per task.

Qwen 3.8 Flash Next threads that needle. It’s a larger model than the previous go-to options, but its INT4 quant doesn’t degrade output quality as much as expected, and it retains vision input, which matters for any agent that needs to look at screenshots, diagrams, or rendered UI as part of its workflow. For anyone doing iterative coding or content-generation agent loops locally, that’s a meaningfully different experience than being stuck making single-shot chat requests.

## What hardware do you actually need?

The baseline hardware profile for this setup is specific and non-negotiable in a few places:

- **GPUs** : four RTX 3090s, giving 96GB of combined VRAM. This is what allows the model to load with a large KV cache and still hit the full context window.
- **System RAM** : 128GB is the recommended floor. During model load, system memory usage climbs to around 100GB before settling, so anything close to that ceiling risks failure.
- **CUDA/drivers** : CUDA and NVCC 13.0 or higher (13.3+ preferred), with current NVIDIA drivers installed.
- **Architecture target** : the build is compiled for SM86 (the Ampere generation that the 3090 belongs to). Newer cards on SM89 architectures could likely work with some changes to the build, but that’s not the tested path.

This is a demanding but not exotic setup. It doesn’t require enterprise GPUs or NVLink bridges, just four consumer 3090s and enough system memory to hold the model as it loads into RAM before being distributed across the GPUs.

## How do you configure vLLM for this setup?

The specific vLLM launch flags matter more than usual here, because a few of them are the difference between a working agent and a system that silently produces worse results or crashes under load. The core configuration includes:

- `--max-model-len 131072` : sets the maximum context window to 131,072 tokens.
- `--max-num-seqs 2` : limits concurrent sequences, tuned for single-agent workloads rather than high-throughput multi-user serving.
- `--max-num-batched-tokens 2048` : caps batch size for token processing.
- `--gpu-memory-utilization 0.96` : an aggressive setting that pushes VRAM usage close to the ceiling to maximize context and cache space.
- `--kv-cache-dtype auto` : lets vLLM pick the appropriate precision for the KV cache automatically.
- `--cuda-graph-mode full` : enables full CUDA graph capture for faster repeated execution.
- `--enable-prefix-caching` : reuses computation across requests that share a prompt prefix, which helps agentic loops that repeatedly reference the same system prompt or file context.
- `--enable-auto-tool-choice` with`--tool-call-parser qwen3xml` : required for reliable function/tool calling in an agent framework.
- **No async scheduling** : this is deliberately left off. Turning it on can add close to 10 tokens per second in throughput, but without an additional patch applied on top of the base setup, it introduces instability that eventually crashes the server.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

Two other settings worth tuning: a compression threshold (a starting value around 0.5, with headroom to push toward 0.65) and a context length setting around 131,072 tokens with a KV cache large enough to use the full window. On a quad 3090 rig, that combination has been shown to fit the full 262-something thousand token combined context and cache budget while still sustaining generation speeds in the 55 to 61 tokens per second range.

## How fast is it in practice?

On a quad RTX 3090 host, prompt processing has been observed climbing from roughly 399 to over 500 tokens as context builds, with generation throughput landing around 59 tokens per second in agentic use, dropping slightly with async scheduling disabled but staying stable. That’s roughly double the throughput seen running the same class of model through llama.cpp on identical hardware, which is the main reason to go through the vLLM setup process instead of sticking with a more familiar llama.cpp workflow.

For agentic work specifically, that speed difference compounds. An agent loop that calls tools, re-reads file state, and generates new code or content dozens of times per session benefits far more from a 2x throughput gain than a single chat response would.

## Is this setup worth the effort?

For anyone already running multiple 3090s, yes. The hardware most people already have (3090s bought during or after the GPU shortage years) continues to be viable for genuinely capable local agent work, provided the software stack keeps catching up. The fact that a quad-3090 setup can run a vision-capable model with a full 131K context window and sustain generation speeds suitable for iterative agent loops is a meaningful sign that these GPUs still have real headroom left in them.

The caveat is that some of the performance gains involve patches and configurations that are still moving quickly. Faster KV cache quantization approaches exist (one 3090 owner reported hitting close to 200 tokens per second with additional KV cache quantization steps), but those come with more setup complexity and are less stable as a repeatable recipe. The configuration described here is meant to be a stable, repeatable baseline rather than the theoretical ceiling.

## Frequently Asked Questions

### What quantization format does Qwen 3.8 Flash Next use for local deployment?

The commonly used local quant is INT4, specifically a W4A16 format produced by Vinnie AI, which keeps vision input support intact while reducing VRAM requirements enough to run on consumer GPUs.

### Can this run on fewer than four GPUs?

The documented setup assumes four RTX 3090s for a combined 96GB of VRAM to support the full context window and KV cache. Running on fewer GPUs would require reducing context length or other memory-hungry settings, and hasn’t been verified in this configuration.

### Why is async scheduling disabled in the vLLM config?

Async scheduling can increase throughput by around 10 tokens per second, but without an additional compatibility patch it introduces instability that eventually causes the server to crash, so it’s left off for a stable baseline.

### Does this setup work with other GPU architectures besides the 3090?

The build described here targets SM86, the Ampere architecture used by the RTX 3090. It may be adaptable to SM89 architectures with some modification, but that path isn’t the tested default.

### Why does vision support matter for a local agent model?

Vision input lets an agent process screenshots, UI renders, or diagrams as part of its reasoning, which matters for coding and content-generation agents that need to check their own visual output rather than working from text descriptions alone.
