---
id: collect-mindstudio/mindstudio/run-dflash2-vllm-sglang-locally
title: "Run DFlash 2 Speculative Decoding with vLLM and SGLang"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["sglang", "speculative decoding", "vllm", "benchmark", "benchmarks", "cohere", "compute", "diffusion", "distribution", "gpu", "inference", "nvidia"]
source: docs/RAG/Collect RAG/02_mindstudio/run-dflash2-vllm-sglang-locally.md
source_anchor: ""
source_lines: [1, 56]
sha256: 6483cce96dc0d18239fa19b9d05d7a6df625c4562b55bd268161e7f78ad527ee
---

# Run DFlash 2 Speculative Decoding with vLLM and SGLang

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-dflash2-vllm-sglang-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide shows how to set up **DFlash 2**, a lossless draft model for **Qwen3.8-27B**, using **vLLM** or **SGLang** for up to **3.4x** faster inference. DFlash 2 is a draft model for speculative decoding: instead of guessing one token at a time, it predicts a whole **block of tokens** in a single forward pass, keeps multiple candidates at each position, and uses a lightweight selector to pick one coherent path, which the target model then verifies in one step. On a single **NVIDIA H200**, this pushed throughput up to **3.43x** over plain autoregressive decoding at low concurrency. Decoding stays lossless: greedy output matches the target model exactly, and sampled output preserves the same distribution.

DFlash 2 changes how the draft is generated. Rather than drafting sequentially, it uses a **block-diffusion** approach: it predicts an entire block of tokens at once and retains the top candidates at every position, then a small **selector module** traces a single consistent path through them (rather than taking the greedy pick at each slot independently). This matters because naive parallel drafting produces blocks where later tokens don't cohere with earlier ones. The backbone also uses **"two-tap dynamic convolutions"**, credited with preventing draft quality from decaying toward the tail end of each block — a known weakness in earlier drafters. The practical effect shows in **acceptance length**, the average tokens accepted per verification step: on **GSM8K, DFlash 2 hit 5.46** vs **5.02 for Qwen3.8's built-in MTP head** and **4.36 for the DSpark drafter**, with similar gaps on MATH-500, HumanEval, MBPP, and MT-Bench.

Hardware and models: benchmarks run on a single **NVIDIA H200** with **FlashAttention 3** for both target and draft. The target is **Qwen/Qwen3.8-27B** (27B parameters), and the draft checkpoint (**Qwen3.8-27B-DFlash2**) is mirrored on Hugging Face under both **z-lab** and **incoai** namespaces. It is not a general-purpose model you can prompt directly; it only functions inside a speculative decoding server alongside its matching target model. Evaluation used a speculation **block size of 8** (7 draft tokens per verification step), **temperature 1.0**, **top-p 0.95**, **top-k 20** (Qwen3.8's official recommended settings), and "**xhigh**" reasoning effort, capped at **4096 new tokens** per generation.

SGLang setup requires building from GitHub source rather than a tagged release:
`pip install "sglang[all] @ git+https://github.com/sgl-project/sglang.git#subdirectory=python"`, then launching with `--model-path Qwen/Qwen3.8-27B --speculative-algorithm DFLASH --speculative-draft-model-path incoai/Qwen3.8-27B-DFlash2 --speculative-num-draft-tokens 8`. The `--speculative-algorithm DFLASH` flag selects the block-diffusion drafting path. vLLM support currently lives on an open **pull request** (not main), installed via `pip install -U "vllm @ git+https://github.com/vllm-project/vllm.git@refs/pull/52816/head"`, then served with `--speculative-config '{"method": "dflash", "model": "incoai/Qwen3.8-27B-DFlash2", "num_speculative_tokens": 7}'`. Note vLLM uses **7** speculative tokens (not 8), reflecting the 7 draft tokens generated per step within the block-size-8 scheme. Since the integration depends on an unmerged PR, expect it to shift or break.

On value: gains depend heavily on **concurrency**. At concurrency 1, DFlash 2 delivered 3.43x on GSM8K, 3.34x on MATH-500, and 3.11x on HumanEval. At concurrency 8, gains stay strong (2.27x-2.85x). At **concurrency 32**, GPU compute is saturated, so speculative decoding has less idle capacity: DFlash 2 still beats plain autoregressive (1.01x-1.45x), but the built-in MTP head and DSpark fall **below 1.0x** on several tasks — slower than running the target model alone. Compared directly, DFlash 2 wins across the board on acceptance length and throughput at every concurrency level and benchmark. The catch is **maturity**: both integrations require building from source or an unmerged PR rather than a stable pip package.

## Key points

- DFlash 2 is a block-diffusion draft model proposing several tokens per step, verified together by the target model.
- It's designed specifically for Qwen/Qwen3.8-27B and ships as a companion checkpoint, not a standalone LLM.
- Decoding is lossless: greedy matches the target model; sampling preserves the distribution.
- Benchmarks on one H200: up to 3.43x at concurrency 1, dropping to ~1.0x-1.45x at concurrency 32.
- It outperforms Qwen3.8's built-in MTP head and the DSpark drafter on acceptance length and throughput.
- Setup: SGLang from source; vLLM from an open pull request branch — neither in a stable release.
- Two-tap dynamic convolutions keep draft quality from decaying late in each block.

## Technical data / figures

| Item | Value |
|---|---|
| Target model | Qwen/Qwen3.8-27B |
| Draft checkpoint | Qwen3.8-27B-DFlash2 (z-lab, incoai) |
| GPU | Single NVIDIA H200, FlashAttention 3 |
| Block size | 8 |
| vLLM spec tokens | 7 |
| Sampling | temp 1.0, top-p 0.95, top-k 20 |
| Reasoning effort | xhigh |
| Max new tokens | 4096 |
| Acceptance length (GSM8K) | 5.46 (MTP 5.02, DSpark 4.36) |
| Speedup concurrency 1 | 3.43x (GSM8K), 3.34x (MATH-500), 3.11x (HumanEval) |
| Speedup concurrency 8 | 2.27x-2.85x |
| Speedup concurrency 32 | ~1.01x-1.45x |
| SGLang install | GitHub source |
| vLLM install | PR #52816 branch |

## Why this source matters for the RAG

It provides exact, copy-ready deployment commands for a cutting-edge speculative decoding setup in both major serving engines. This supports procedural and troubleshooting questions about running DFlash 2 with vLLM or SGLang.
