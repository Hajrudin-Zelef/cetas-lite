---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face-2
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "TensorRT-LLM", "vLLM"]
dates: ["2026-11-08"]
keywords: ["fp4", "nvfp4", "nvidia", "blackwell", "context window", "data centre", "diffusion", "fp8", "gpu", "gpus", "inference", "kv cache"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [135, 307]
sha256: f38e76817a26917cc42c94a63e59206d9551101b45f632e78b75a4ca22935d06
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face

- We performed post-training quantization (PTQ) with Nvidia Model Optimizer using the following recipe: Four Over Six NVFP4 (a variant of static MSE calibration) W4A16 on routed and shared experts, FP8 per-tensor dynamic scales on mamba in_proj/out_proj and KV cache. We used a subset of the Nemotron Ultra validation set for calibration with 1000 samples at 32k token length.

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 is a result of the above work.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. Supported languages include English, Spanish, French, German, Italian, and Japanese.

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- **Runtime Engine(s):** PyTorch
- **Supported Hardware Microarchitecture Compatibility:** NVIDIA Blackwell; NVIDIA Hopper (NVFP4 / W4A16); NVIDIA Ampere (W4A16)
- **Preferred/Supported Operating System(s):** Linux

- GA (08/11/2026)

| Setup & goal | Recommended path | Speculative decoding | Validated context | 
|---|---|---|---|
| **DGX Spark (GB10)** | vLLM + DSpark | DSpark ( `num_speculative_tokens 3` ) | 1M | 
| **1× H100 — max throughput (batch)** | vLLM, no spec decoding ( `humming` backend) | None | 1M (default) | 
| **1× H100 — interactive (40+ TPS/user)** | vLLM + DSpark, concurrency ≤ 128 | DSpark ( `num_speculative_tokens 3` ) | 1M (default) | 
| **8× H100 — long context** | vLLM, TP8 + expert parallel | None | 1M | 
| **1× GB200** | vLLM + DSpark | DSpark ( `num_speculative_tokens 5` ) | 1M (default) | 
| **Ampere (A100, etc.)** | vLLM, W4A16 kernels | None | 1M (default) | 
| **Local — llama.cpp** | `llama-server ...` | — | ~40K example, VRAM-bound | 
| **Local — Ollama / LM Studio / Jan…** | `ollama run nemotron-3.5-lightning` | built-in | varies | 

All deployment snippets below assume:

```
export MODEL_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4
```
And for DSpark:

```
export DSPARK_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark
```
Lightning 3.5 ships with two external draft models for speculative decoding as well as MTP (Multi-Token Prediction). While DSpark is our recommended default (see Choose your deployment), some workloads are better served without speculative decoding, or by DFlash or MTP:

- **DSpark:** A semi-autoregressive speculative-decoding drafter that proposes a whole block of candidate tokens in a single forward pass from a parallel backbone. This is recommended for DGX Spark, as well as low-concurrency data centre deployments.
- **DFlash:** A speculative-decoding drafter that uses a lightweight block-diffusion model to generate an entire draft block in one forward pass.
- **MTP:** A modeling technique that trains the network to predict several future tokens at each position instead of only the next one.

**DSpark is the recommended default** for DGX Spark and latency-sensitive, low-concurrency serving. For maximum-throughput batch serving on H100, **no** speculative decoding is fastest.

For more indepth instructions on how to deploy through vLLM, head here


- vLLM version: `vllm/vllm-openai:v0.27.1`

```
vllm serve --model $MODEL_CKPT \
  --moe-backend marlin \
  --kv-cache-dtype fp8 \
  --enable-prefix-caching \
  --gpu-memory-utilization 0.85 \
  --speculative_config.num_speculative_tokens 3 \
  --mamba-backend flashinfer \
  --mamba-cache-mode align \
  --reasoning-parser nemotron_v3 \
  --speculative_config.model $DSPARK_CKPT \
  --tool-call-parser qwen3_coder \
  --enable-auto-tool-choice
```
**Validated context:** 1M tokens (default).

For max throughput deployments, use the following configuration, no speculative decoding strategy is best for this serving configuration, and due to memory constraints the Mamba cache `dtype` is set as FP16:

```
VLLM_HUMMING_MOE_GEMM_TYPE=indexed \
vllm serve --model $MODEL_CKPT \
    --max-num-seqs 256 \
    --max-num-batched-tokens 16384 \
    --enable-prefix-caching \
    --async-scheduling \
    --mamba-backend flashinfer \
    --moe-backend humming \
    --linear-backend humming \
    --mamba-ssu-algorithm horizontal \
    --mamba-cache-mode align \
    --mamba-ssm-cache-dtype float16 \
    --enable-mamba-cache-stochastic-rounding \
    --mamba-cache-philox-rounds 5 \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
**Validated context:** 1M tokens (default).

For interactive usage scenarios (achieving 40+ TPS/User) use a lower concurrency (<=128) with DSpark:

```
vllm serve --model $MODEL_CKPT \
    --max-num-seqs 128 \
    --enable-prefix-caching \
    --async-scheduling \
    --speculative_config.model $DSPARK_CKPT \
    --speculative_config.num_speculative_tokens 3 \
    --mamba-backend flashinfer \
    --mamba-ssm-cache-dtype float16 \
    --mamba-cache-mode align \
    --enable-mamba-cache-stochastic-rounding \
    --mamba-cache-philox-rounds 5 \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
**Validated context:** 1M tokens, served by default. Lower `--max-model-len` for more KV-cache headroom at high concurrency.

For long-context, multi-GPU serving (TP8 with expert parallelism):

```
vllm serve --model $MODEL_CKPT \
    --mamba-backend flashinfer \
    --async-scheduling \
    --enable-prefix-caching \
    --mamba-cache-mode align \
    --enable-expert-parallel \
    --tensor-parallel-size 8 \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
**Validated context:** 1M tokens (TP8 + expert parallel).

```
vllm serve --model $MODEL_CKPT \
    --max-num-batched-tokens 10240 \
    --no-enable-prefix-caching \
    --async-scheduling \
    --speculative_config.model $DSPARK_CKPT \
    --speculative_config.num_speculative_tokens 5 \
    --mamba-backend flashinfer \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
**Validated context:** 1M tokens (default).

The same checkpoint also serves via W4A16 kernels, extending coverage to Ampere-class GPUs, like the A100 80GB:

```
VLLM_HUMMING_MOE_GEMM_TYPE=indexed \
vllm serve --model nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 \
    --moe-backend humming \
    --linear-backend humming \
    --max-num-seqs 256 \
    --max-num-batched-tokens 32768 \
    --enable-prefix-caching \
    --async-scheduling \
    --quantization modelopt_fp4 \
    --mamba-backend flashinfer \
    --mamba-cache-mode align \
    --mamba-ssu-algorithm simple \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
**Validated context:** 1M tokens (default)

- **Context Length:** The H100 and GB200 snippets above serve the model's full 1M-token context window by default. If you're memory-constrained — or want more KV-cache headroom at high concurrency — lower`--max-model-len` to match your workload.

For more indepth instructions on how to deploy through TensorRT-LLM, head here


Container: `nvcr.io/nvidia/tensorrt-llm/release:1.3.0rc24`

