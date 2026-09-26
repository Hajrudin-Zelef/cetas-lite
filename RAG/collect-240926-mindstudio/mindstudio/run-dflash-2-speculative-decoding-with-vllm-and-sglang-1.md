---
id: collect-240926-mindstudio/mindstudio/run-dflash-2-speculative-decoding-with-vllm-and-sglang-1
title: "run-dflash-2-speculative-decoding-with-vllm-and-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "agents", "attention", "benchmark", "benchmarks", "cohere", "compute", "diffusion", "distribution", "exploit", "gpu"]
source: docs/RAG/clean_en/mindstudio/run-dflash-2-speculative-decoding-with-vllm-and-sglang.md
source_anchor: ""
source_lines: [1, 87]
sha256: 6985310508c85958f23ddb5cbf26c0dddbc95f74aff8667e917a027e770af771
---

# run-dflash-2-speculative-decoding-with-vllm-and-sglang

<!-- source: https://www.mindstudio.ai/blog/run-dflash2-vllm-sglang-locally -->

## What is DFlash 2 and why does it speed up inference?

DFlash 2 is a draft model built for speculative decoding with Qwen/Qwen3.8-27B. Instead of guessing one token at a time, it predicts a whole block of tokens in a single forward pass, keeps multiple candidates at each position, and uses a lightweight selector to pick one coherent path through them. The target model then verifies that path in one step. On a single NVIDIA H200, this setup pushed throughput up to 3.43x over plain autoregressive decoding at low concurrency, according to the model’s published benchmarks. Decoding stays lossless: greedy output matches the target model exactly, and sampled output preserves the same distribution.

## TL;DR

- **DFlash 2** is a block-diffusion draft model that proposes several tokens per step instead of one, then lets the target model verify them together.
- It’s designed specifically for **Qwen/Qwen3.8-27B** and ships as a companion checkpoint, not a standalone LLM you can query on its own.
- Decoding is **lossless** : the final output matches what the target model would have produced running alone, whether you use greedy decoding or sampling.
- Benchmarks on a single **H200 GPU** show throughput gains up to 3.43x at concurrency 1, dropping to roughly 1.0x to 1.45x at concurrency 32, where GPU compute is already saturated.
- It outperforms both Qwen3.8’s built-in multi-token prediction (MTP) head and the community **DSpark** drafter on acceptance length and throughput across every tested benchmark.
- You can serve it today with a **SGLang** build from source or a**vLLM** pull-request branch; neither integration has landed in a stable release yet.
- The two-tap dynamic convolution in its backbone is what keeps draft quality from degrading toward the end of each token block, a common failure mode in earlier speculative drafters.

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

## How does DFlash 2 actually work?

Speculative decoding normally works by having a small, fast draft model guess the next few tokens, then having the large target model check those guesses in parallel. If the target model accepts a guess, you save a full forward pass. If it rejects one, you fall back to normal generation from that point.

DFlash 2 changes how the draft itself is generated. Rather than drafting tokens sequentially, it uses a block-diffusion approach: it predicts an entire block of tokens at once and retains the top candidates at every position within that block. A small selector module then traces a single, consistent path through those candidates rather than just taking the greedy pick at each slot independently. This matters because naive parallel drafting tends to produce blocks where later tokens don’t cohere with earlier ones. The selector’s job is to fix that before the target model ever sees the draft.

The backbone also uses “two-tap dynamic convolutions,” which the model card credits with preventing draft quality from decaying toward the tail end of each block, a known weakness in earlier drafters where confidence and accuracy fall off after the first token or two.

The practical effect shows up in “acceptance length,” the average number of tokens the target model accepts per verification step. Longer acceptance length means fewer verification passes for the same output, which is where the speedup comes from. On GSM8K, DFlash 2 hit an acceptance length of 5.46 tokens, compared to 5.02 for Qwen3.8’s built-in MTP head and 4.36 for the DSpark drafter. Similar gaps show up on MATH-500, HumanEval, MBPP, and MT-Bench.

## What hardware and models do you need?

The published benchmarks run on a single NVIDIA H200 with FlashAttention 3 handling attention for both the target and draft models. The target model is Qwen/Qwen3.8-27B, a 27-billion-parameter model, and the draft checkpoint (Qwen3.8-27B-DFlash2) is mirrored on Hugging Face under both `z-lab` and `incoai` namespaces. It is not a general-purpose model you can prompt directly; it only functions inside a speculative decoding server alongside its matching target model.

The evaluation setup used a speculation block size of 8, meaning 7 draft tokens get proposed per verification step, sampling temperature of 1.0, top-p of 0.95, top-k of 20 (Qwen3.8’s official recommended settings), and “xhigh” reasoning effort, with a cap of 4096 new tokens per generation.

## How do you set it up with SGLang?

SGLang support requires building from the project’s GitHub source rather than a tagged release:

```
pip install "sglang[all] @ git+https://github.com/sgl-project/sglang.git#subdirectory=python"
python -m sglang.launch_server \
  --model-path Qwen/Qwen3.8-27B \
  --speculative-algorithm DFLASH \
  --speculative-draft-model-path incoai/Qwen3.8-27B-DFlash2 \
  --speculative-num-draft-tokens 8
```
The `--speculative-algorithm DFLASH` flag tells SGLang to use the block-diffusion drafting path instead of a standard speculative decoding scheme. The draft token count here is set to 8, matching the block size used in the official evaluation.

## How do you set it up with vLLM?

vLLM support for DFlash currently lives on an open pull request, not in the main release, so you install from that PR branch directly:

```
pip install -U "vllm @ git+https://github.com/vllm-project/vllm.git@refs/pull/52816/head"
vllm serve Qwen/Qwen3.8-27B \
  --speculative-config '{
    "method": "dflash",
    "model": "incoai/Qwen3.8-27B-DFlash2",
    "num_speculative_tokens": 7
  }'
```
### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

Note the vLLM config uses `num_speculative_tokens: 7` rather than 8, reflecting the 7 draft tokens generated per verification step within the block-size-8 scheme. Since this integration depends on an unmerged PR, expect it to shift or break as the branch gets updated or eventually merged.

## Is DFlash 2 worth using over plain decoding or other drafters?

The gains depend heavily on concurrency, which is the number of simultaneous requests hitting the server. At concurrency 1, meaning a single request with no competing load, DFlash 2 delivered its biggest wins: 3.43x speedup on GSM8K, 3.34x on MATH-500, and 3.11x on HumanEval versus autoregressive decoding. At concurrency 8, gains stay strong, ranging from about 2.27x to 2.85x depending on the task.

At concurrency 32, the picture changes. GPU compute is largely saturated by then, so speculative decoding has less idle capacity to exploit. DFlash 2 still comes out ahead of plain autoregressive decoding on every benchmark (1.01x to 1.45x), but the built-in MTP head and DSpark drafter actually fall below 1.0x on several tasks at this concurrency, meaning they’re slower than just running the target model alone. That’s a meaningful data point: some speculative decoding setups can hurt throughput under heavy concurrent load, and DFlash 2’s advantage over autoregressive decoding narrows but doesn’t disappear.

Compared directly to the alternatives, DFlash 2 wins across the board. It beats Qwen3.8’s built-in MTP head and the community DSpark drafter on acceptance length and throughput at every concurrency level and every benchmark task tested (GSM8K, MATH-500, HumanEval, MBPP, MT-Bench). For anyone already serving Qwen3.8-27B and looking for a lossless speedup, it’s the strongest published option among the three.

The catch is maturity. Both integrations require building from source or an unmerged pull request rather than a stable pip package, so expect some friction and the likelihood that commands or flags change before this lands in official releases.

## Frequently Asked Questions

### What GPU do I need to run DFlash 2?

