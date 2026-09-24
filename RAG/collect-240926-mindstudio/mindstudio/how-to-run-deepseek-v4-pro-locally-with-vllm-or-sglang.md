---
id: collect-240926-mindstudio/mindstudio/how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang
title: "how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["deepseek", "sglang", "vllm", "agent", "agentic", "agents", "attention", "benchmark", "claude", "fp4", "fp8", "glm"]
source: docs/RAG/clean_en/mindstudio/how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [1, 104]
sha256: c4c3a11447c7f74dc42fc9997820e82bb65ff67c7a84bc0a1f4048e1eea50cef
---

# how-to-run-deepseek-v4-pro-locally-with-vllm-or-sglang

<!-- source: https://www.mindstudio.ai/blog/deepseek-v4-local-vllm-sglang -->

## What is DeepSeek V4 Pro and why run it locally?

DeepSeek-V4-Pro-0813 is the full release version of DeepSeek’s flagship reasoning and agentic coding model, replacing the earlier preview build. It ships under an MIT license with open weights, which means anyone with sufficient hardware can self-host it instead of calling DeepSeek’s API. Running it locally matters if you care about data control, want to avoid per-token pricing on heavy agentic workloads, or need to wire it into your own tooling (including DeepSeek’s new agent harness) without depending on an external endpoint. The tradeoff is hardware: this is a large mixture-of-experts model that expects multi-GPU serving infrastructure, not a laptop.

## TL;DR

- **DeepSeek-V4-Pro-0813** is the production successor to the V4 Pro preview, adding a DSpark speculative decoding module and notably stronger agentic benchmark scores across coding and tool-use tasks.
- **vLLM and SGLang** are the two officially documented serving engines, each with a specific launch command and hardware recipe published alongside the model card.
- **DSpark speculative decoding** is a single-flag feature in both engines, using the same checkpoint as both draft and target model rather than requiring a separate draft model.
- **Reference deployments target multi-GPU nodes** , such as a single 4x GB300 node, using fp8 KV cache, expert parallelism, and MoE-specific backends.
- **Reasoning effort is configurable** at low, high, or max levels, with max effort recommending output lengths up to 384K tokens for local deployment.
- **The DeepSeek Harness** , a separate open-source agentic coding framework, can call locally-hosted V4 Pro as a plug-in model rather than only DeepSeek’s hosted API.
- **Independent benchmarking** places V4 Pro among the strongest open models on agentic and coding tasks, though it trails a few of the largest proprietary and open competitors on raw capability.

## One coffee. One working app.

You bring the idea. Remy manages the project.

## How do you set up vLLM for DeepSeek V4 Pro?

vLLM is DeepSeek’s first documented serving path, and the model card provides a working launch command tuned for a single node with four GB300 GPUs. The core command looks like this:

```
vllm serve deepseek-ai/DeepSeek-V4-Pro-0813 \
  --trust-remote-code --kv-cache-dtype fp8 --block-size 256 \
  --data-parallel-size 4 --enable-expert-parallel \
  --moe-backend deep_gemm_mega_moe \
  --attention-config '{"use_fp4_indexer_cache": true}' \
  --speculative-config '{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}'
```
A few things stand out. The `fp8` KV cache dtype and 256 block size are chosen to keep memory footprint manageable given the model’s context length. `--data-parallel-size 4` combined with `--enable-expert-parallel` reflects the mixture-of-experts architecture, where experts get distributed across GPUs rather than replicating the full model on each device. The `moe-backend` flag points to a specialized kernel (`deep_gemm_mega_moe`) built for this class of MoE workload, and the `attention-config` flag enables an fp4 indexer cache, a memory optimization tied to DeepSeek’s attention implementation.

DeepSeek publishes a fuller vLLM recipe covering other hardware configurations beyond the single 4x GB300 example, so smaller or larger clusters have documented starting points rather than requiring trial and error.

## How do you set up SGLang for DeepSeek V4 Pro?

SGLang is the second officially supported engine, and its launch command differs mainly in how it expresses parallelism and quantization:

```
sglang serve \
  --trust-remote-code \
  --model-path deepseek-ai/DeepSeek-V4-Pro-0813 \
  --tp 4 \
  --moe-runner-backend flashinfer_mxfp4 \
  --speculative-algorithm DSPARK \
  --mem-fraction-static 0.90 \
  --chunked-prefill-size 4096 \
  --swa-full-tokens-ratio 0.1
```
Here `--tp 4` sets tensor parallelism across four GPUs, and `--moe-runner-backend flashinfer_mxfp4` selects an mxfp4 quantized MoE kernel through FlashInfer. The `--mem-fraction-static 0.90` flag reserves most of GPU memory for static allocation, and `--chunked-prefill-size 4096` controls how prompt processing is batched to balance latency and throughput. SGLang’s cookbook documentation extends this further with configurations tagged by hardware (GB300), quantization (fp4), and strategy (low-latency), giving operators a menu rather than a single fixed setup.

The key difference in enabling speculative decoding between the two engines is naming: vLLM uses a JSON `--speculative-config` block, SGLang uses a plain `--speculative-algorithm DSPARK` flag. Functionally they do the same thing.

## What is DSpark speculative decoding and why does it matter?

DSpark is the speculative decoding method DeepSeek built specifically for V4 Pro. Speculative decoding normally works by having a smaller, faster “draft” model propose several tokens ahead, which the larger target model then verifies in a single pass, cutting down the number of expensive full forward passes needed to generate text. What’s notable about DSpark is that it does not require a separate draft model checkpoint. The target and draft weights come from the same model file, so there’s no extra model to download, load into memory, or keep in sync with the target model’s tokenizer and behavior.

In both vLLM and SGLang, DSpark is enabled with `num_speculative_tokens` set to a small number (7 in the vLLM example) and a greedy draft sampling method. Because it’s a single-flag feature built into the officially released weights, it’s the default recommended path for local deployment rather than an advanced or experimental option.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

## Is DeepSeek V4 Pro worth running locally versus using the API?

The honest answer depends on workload and scale. On independent benchmark rankings referenced around its release, V4 Pro placed among the strongest models overall, though it trailed the very largest models like Qwen3-Max, which arrived just days earlier. DeepSeek’s own benchmark table shows V4 Pro-0813 improving substantially over the V4 Pro preview across coding-agent tasks (Terminal Bench 2.1 jumped from 72.1 to 87.9, DeepSWE from 12.8 to 62.7), and landing competitively against GLM-5.2, Kimi K3, and Opus-4.8 depending on the specific benchmark.

Pricing on DeepSeek’s hosted API rose significantly with this release, reportedly by 2x to 4x compared to prior versions, though it reportedly remains cheap relative to other frontier-class models. For teams running heavy, sustained agentic coding workloads (the kind that can burn through tens of millions of tokens on a single complex task, as observed in early testing of the accompanying DeepSeek Harness), the economics of self-hosting become more attractive at scale, assuming the multi-GPU hardware investment is already available or justified elsewhere.

For lighter or occasional use, the hosted API removes the need to manage GPU clusters, MoE-specific kernels, and speculative decoding configuration entirely.

## How does local deployment interact with the DeepSeek Harness?

DeepSeek also released an open agentic coding framework called DeepSeek Harness, built around the idea that tools, skills, sessions, and even entire coding agents (Codex, Claude Code) can be plugged in and swapped. It ships with a local web app (`ds web`) and a settings panel where users configure model providers.

Because the harness treats models as pluggable providers under an MIT license, a locally-hosted V4 Pro instance running under vLLM or SGLang can be registered as a custom provider instead of pointing at DeepSeek’s hosted API. This is useful for anyone who wants the harness’s agent tooling, permission modes (read-only, workspace, full access), and plugin ecosystem while keeping inference entirely on their own infrastructure.

## Frequently Asked Questions

### What hardware do you need to run DeepSeek V4 Pro locally?

DeepSeek’s official recipes reference multi-GPU nodes, with the primary documented example being a single node with four GB300 GPUs for both vLLM and SGLang. Both recipes link out to additional hardware configurations for different cluster sizes.

### Does DeepSeek V4 Pro require a separate draft model for speculative decoding?

No. DSpark speculative decoding draws both draft and target weights from the same released checkpoint, so there’s no additional draft model to download or maintain.

### What’s the difference between reasoning effort levels in V4 Pro?

The model supports low, high, and max reasoning effort settings, controlling how much deliberation it performs before answering. For high and max effort, DeepSeek recommends allowing output lengths up to 384K tokens locally.

### Can I use DeepSeek V4 Pro with agent frameworks besides DeepSeek’s own harness?

Yes. Since it’s served through standard OpenAI-compatible endpoints via vLLM or SGLang, it can be wired into most agent frameworks that support custom model providers, not just DeepSeek’s own harness.

### Is DeepSeek V4 Pro fully open source?

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

The repository and model weights are released under the MIT license, which permits commercial use, modification, and self-hosting without licensing fees.
