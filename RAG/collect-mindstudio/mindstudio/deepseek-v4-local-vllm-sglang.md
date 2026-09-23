---
id: collect-mindstudio/mindstudio/deepseek-v4-local-vllm-sglang
title: "How to Run DeepSeek V4 Pro Locally with vLLM or SGLang"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "DeepSeek", "Moonshot", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["deepseek", "sglang", "vllm", "agent", "agentic", "attention", "benchmark", "benchmarks", "fp4", "fp8", "glm", "gpu"]
source: docs/RAG/Collect RAG/02_mindstudio/deepseek-v4-local-vllm-sglang.md
source_anchor: ""
source_lines: [1, 51]
sha256: adfa8a639b062bd3c3f402894babdee9736cf1b665090b15d0a9a7abdb2fe3c1
---

# How to Run DeepSeek V4 Pro Locally with vLLM or SGLang

## Metadata

- **Source** : https://www.mindstudio.ai/blog/deepseek-v4-local-vllm-sglang
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide explains how to self-host **DeepSeek-V4-Pro-0813**, the full production release of DeepSeek's flagship reasoning and agentic coding model, which replaces the earlier preview build. It ships under an **MIT license with open weights**, allowing anyone with sufficient hardware to run it instead of calling DeepSeek's API. Running it locally matters for data control, avoiding per-token pricing on heavy agentic workloads, and wiring it into custom tooling — including DeepSeek's new agent harness — without depending on an external endpoint. The tradeoff is hardware: this is a large mixture-of-experts (MoE) model that expects multi-GPU serving infrastructure, not a laptop.

**vLLM** is DeepSeek's first documented serving path. The model card provides a launch command tuned for a single node with four GB300 GPUs, using `--kv-cache-dtype fp8`, `--block-size 256`, `--data-parallel-size 4`, `--enable-expert-parallel`, `--moe-backend deep_gemm_mega_moe`, an fp4 indexer cache in the attention config, and a DSpark speculative config (`num_speculative_tokens: 7`, greedy draft sampling). The fp8 KV cache and 256 block size keep memory manageable given the context length, while data/expert parallelism distributes experts across GPUs rather than replicating the full model.

**SGLang** is the second officially supported engine. Its command differs mainly in how parallelism and quantization are expressed: `--tp 4` for tensor parallelism, `--moe-runner-backend flashinfer_mxfp4` for an mxfp4 quantized MoE kernel, `--mem-fraction-static 0.90`, `--chunked-prefill-size 4096`, and `--speculative-algorithm DSPARK`. SGLang's cookbook extends this with hardware-tagged configurations (GB300), quantization (fp4), and strategy (low-latency).

**DSpark** is the speculative decoding method DeepSeek built for V4 Pro. Unlike traditional speculative decoding, it does not require a separate draft model checkpoint — both draft and target weights come from the same model file, so there is no extra model to download, load, or keep in sync. It is enabled as a single flag in both engines and is the default recommended local deployment path.

On benchmarks, V4 Pro placed among the strongest models overall but trailed the very largest models like **Qwen3-Max**. DeepSeek's own table shows V4 Pro-0813 improving substantially over the preview on coding-agent tasks (Terminal Bench 2.1 from 72.1 to 87.9; DeepSWE from 12.8 to 62.7), landing competitively against **GLM-5.2, Kimi K3, and Opus-4.8**. Hosted API pricing reportedly rose 2x–4x with this release, though it remains cheap relative to other frontier models. Reasoning effort is configurable at low, high, or max, with max recommending output lengths up to 384K tokens. The **DeepSeek Harness**, an open agentic coding framework, can register a locally-hosted V4 Pro instance as a custom provider instead of using the hosted API.

## Key points

- DeepSeek-V4-Pro-0813 is the MIT-licensed production successor to the V4 Pro preview, adding a DSpark speculative decoding module and stronger agentic/coding benchmark scores.
- vLLM and SGLang are the two officially documented serving engines, each with a published launch command and hardware recipe (reference: single 4x GB300 node).
- DSpark is a single-flag feature using the same checkpoint as both draft and target model, so no separate draft model is needed.
- Reference deployments use fp8 KV cache, expert parallelism, and MoE-specific backends.
- Reasoning effort is configurable at low/high/max, with max recommending output up to 384K tokens.
- The DeepSeek Harness can call a locally-hosted V4 Pro as a plug-in model.
- Independent benchmarking places V4 Pro among the strongest open models, trailing a few largest proprietary/open competitors on raw capability.

## Technical data / figures

| Item | Detail |
|---|---|
| Model | DeepSeek-V4-Pro-0813 |
| License | MIT (open weights) |
| Reference hardware | Single node, 4x GB300 GPUs |
| vLLM key flags | fp8 KV cache, block-size 256, DP=4, expert parallel, deep_gemm_mega_moe, fp4 indexer cache |
| SGLang key flags | tp=4, flashinfer_mxfp4, mem-fraction-static 0.90, chunked-prefill-size 4096, DSPARK |
| Speculative decoding | DSpark, num_speculative_tokens=7, greedy draft |
| Reasoning effort | low / high / max (up to 384K output tokens) |
| Terminal Bench 2.1 | 72.1 (preview) → 87.9 (0813) |
| DeepSWE | 12.8 (preview) → 62.7 (0813) |
| API price change | reportedly 2x–4x increase |

## Why this source matters for the RAG

It provides concrete, copy-ready deployment commands and hardware requirements for self-hosting DeepSeek V4 Pro, a frontier-class open-weight MoE model, making it a practical reference for local/private AI infrastructure. It also documents DSpark speculative decoding and the DeepSeek Harness integration, which are central to local agentic-coding economics.
