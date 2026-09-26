---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face-1
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "vLLM"]
dates: ["2025-09", "2025-12", "2026-05", "2026-08-11", "2026-11-08"]
keywords: ["fp4", "nvfp4", "nvidia", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "blackwell", "compute", "data centre"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 134]
sha256: 06d5aaeb6d896bf7788e288fb7ba0277e4e74c40e6bf67bed1301483ba9e910b
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 -->

| **Total Parameters** | 30B (3B active) | 
| **Architecture** | MoE - Mamba-2 + MoE + Attention hybrid | 
| **Context Length** | Up to 1M tokens | 
| **Single-GPU Deployment** | 1× DGX Spark (GB10) or 1× H100 | 
| **Supported Hardware** | NVIDIA Blackwell (DGX Spark / GB10, GB200, GeForce RTX 5090); NVIDIA Hopper (H100, H200); NVIDIA Ampere via W4A16 | 
| **Supported Languages** | English (and coding languages), Spanish, French, German, Italian, Japanese | 
| **Speculative Decoding** | DSpark for low-concurrency Data Centre and DGX Spark Workflows — Read more below, also provided are MTP (Multi-Token Prediction) and DFlash | 
| **Recommended Sampling** | Temperature 1.0, Top_P 0.95 | 
| **Best For** | Long-running autonomous agents, sub-agent workhorse deployments, and efficient local inference on personal hardware | 
| **License** | OpenMDW License Agreement, version 1.1 | 
| **Release Date** | August 11, 2026 | 

## Hardware Matrix

| Hardware | Stored precision | Compute path | MoE backend (per recipe) | Native FP4 tensor-core path | Validated context | 
|---|---|---|---|---|---|
| **Blackwell — GB200** | NVFP4 | FP4 | default | Yes | 1M (default) | 
| **Blackwell — DGX Spark (GB10)** | NVFP4 | W4A16 | `marlin` | No — runs via Marlin | 1M (default) | 
| **Blackwell — GeForce RTX 5090** | NVFP4 | *confirm* | *no recipe published* | Hardware: yes | *confirm* | 
| **Hopper — H100 / H200** | NVFP4 | W4A16 | `humming` (max-tput) / default | No — Hopper has no FP4 tensor cores | 1M (default) | 
| **Ampere — A100, etc.** | NVFP4 | W4A16 | `humming` | No | 1M (default) | 

**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - May 2026

**Data Freshness:**

- The pre-training data has a cutoff date of September 2025.
- The post-training data has a cutoff date of May 2026.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

**NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4** is a large language model (LLM) trained by NVIDIA.

The model employs a hybrid **Mixture-of-Experts** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. The Lightning 3.5 model is released alongside a number of speculative decoding methods for faster text generation. The model has **3B active parameters** and **30B parameters in total**.

This model is ready for commercial use.

To get quickly started on DGX Spark (GB10) you can use the following command.

Grab the model:

```
export MODEL_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4
export DSPARK_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark
```
Run it with vLLM — this recipe uses DSpark speculative decoding, tuned for DGX Spark. (vLLM version: `vllm/vllm-openai:v0.27.1`)

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

For more details on how to deploy and use the model — see the Quick Start Guide below!

**Governing Download Terms:** Use of this model is governed by the OpenMDW-1.1 model license.

We evaluated our model on the following benchmarks:

| Task | Nemotron-3.5-Lightning-30B-A3B-BF16 | Nemotron-3.5-Lightning-30B-A3B-NVFP4 | 
|---|---|---|
| **General Knowledge** |  |  | 
| MMLU Pro | 81.94 | 81.62 | 
| AA-Omniscience | 17.50 | 16.63 | 
| **Reasoning** |  |  | 
| GPQA Diamond (no tools) | 75.44 | 75.57 | 
| HLE (text-only, no tools) | 11.72 | 10.47 | 
| SciCode | 32.60 | 31.38 | 
| **Coding & Agentic** |  |  | 
| SWE-bench Verified | 51.56 | 52.80 | 
| SWE-bench Multilingual | 39.33 | 36.47 | 
| Terminal-Bench 2.1 | 24.58 | 23.46 | 
| PinchBench | 85.37 | 83.43 | 
| BrowseComp | 36.97 | 36.81 | 
| τ³-bench (Banking) | 9.28 | 9.48 | 
| GDPval-AA-V2 | 832 | 865 | 
| **Instruction Following** |  |  | 
| IFBench (loose) | 71.88 | 72.88 | 
| **Long Context** |  |  | 
| AA-LCR | 52.00 | 49.19 | 

Accuracy numbers measured by NVIDIA under a consistent harness (NeMo Gym / Nemo Evaluator SDK); they may differ from vendors' self-reported numbers.

For reproducibility, the evaluation recipes, installation instructions, and commands for NVIDIA Nemotron 3.5 Lightning were collected and published in NeMo Gym. The reported results cover the release evaluation suite, including knowledge and reasoning, instruction following, coding, agentic, tool-use, and long-context. Most evaluations use NeMo Gym-native harnesses while a small subset, including SWE-Bench and Terminal-Bench, used NeMo Evaluator natively. The published recipes specify the benchmark-specific containers, prompts, inference parameters, parser configurations, and scoring settings used to produce the results.

These numbers were measured with and apply to the official NVFP4 checkpoint


Additional harness-level coding-agent results for SWE-Bench Verified and Terminal-Bench 2.1 are shown below.

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 is a general purpose reasoning and chat model intended to be used in English and coding languages. Other non-English languages (Spanish, French, German, Italian, Japanese) are also supported. Intended for developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. Also suitable for typical instruction-following tasks.

Hugging Face — 08/11/2026

- **Architecture Type:** Mixture-of-Experts Hybrid (Mamba + Transformer)
- **Network Architecture:** Nemotron-3-Lightning + Multi-Token Prediction (MTP)
- **Number of model parameters:** 30B Total / 3B Active

The model was pre-trained with over 20T tokens and supports up to 1M context length. The pre-training phase used an NVFP4 recipe. The model includes **Multi-Token Prediction (MTP)** layers, which predict multiple future tokens to provide richer training signals.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 model was pre-trained using an NVFP4 recipe with crawled and synthetic code, math, science, and general knowledge data.
- Software used for pre-training: Megatron-LM

Stage 2: Continued Pre-Training for Multi-Token Prediction (MTP)

- The model underwent a continued pre-training phase to train its **Multi-Token Prediction (MTP)** layers. In this stage, MTP heads learn to predict multiple future tokens, providing richer training signals to the base model. This phase aligns the MTP layers with the base model's distribution.

Stage 3: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation.

Stage 4: Reinforcement Learning

- The model underwent multi-environment reinforcement learning using GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. It utilized an asynchronous RL architecture that decouples training from inference and leverages MTP to accelerate rollout generation.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

Stage 5: Post-training Quantization (PTQ)

