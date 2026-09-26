---
id: collect-240926-huggingface/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face-1
title: "unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "Unsloth", "vLLM"]
dates: ["2025-09", "2025-12", "2026-05", "2026-08-11", "2026-11-08"]
keywords: ["gguf", "nvidia", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "blackwell", "data centre", "distillation", "gpu"]
source: docs/RAG/clean_en/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 116]
sha256: 1f1776ad43038feca6f9cbe8c274c7c84f275b7c14e92cc86565cc046110857c
---

# unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face

<!-- source: https://huggingface.co/unsloth/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF -->

## Read our How to Run Nemotron 3.5 Guide!

    *Unsloth Dynamic 2.0 achieves superior accuracy & outperforms other leading quants.*
  

- You can now run Nemotron-3.5 in Unsloth with toggles for thinking.
- See below for 2-bit Nemotron 3.5 call tools for 10mins straight inside of Unsloth:
- The pre-training data has a cutoff date of September 2025.
- The post-training data has a cutoff date of May 2026.

| **Total Parameters** | 30B (3B active) | 
| **Architecture** | MoE — Mamba-2 + MoE + Attention hybrid | 
| **Precision** | BF16 (full-precision reference weights) | 
| **Context Length** | Up to 1M tokens (for single H100 deployment, we use 256K) | 
| **Single-GPU Deployment** | 1× H100 80GB (or 1× A100 80GB) | 
| **Supported Hardware** | NVIDIA Blackwell (GB200, GeForce RTX 5090); NVIDIA Hopper (H100, H200); NVIDIA Ampere (A100) | 
| **Supported Languages** | English (and coding languages), Spanish, French, German, Italian, Japanese | 
| **Speculative Decoding** | DSpark for Low Concurrency Data Centre Deployments — Read more below | 
| **Reasoning Mode** | Configurable on/off via chat template ( `enable_thinking=True/False` ) | 
| **Recommended Sampling** | Temperature 1.0, Top_P 0.95 | 
| **Best For** | Customization — post-training (SFT, RL, distillation), domain adaptation, building quantized variants, and research/evaluation at full precision | 
| **Looking to Deploy?** | For optimized inference, see NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 | 
| **License** | OpenMDW License Agreement, version 1.1 | 
| **Release Date** | August 11, 2026 | 

**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - May 2026

**Data Freshness:**

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

**NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16** is a large language model (LLM) trained by NVIDIA. This is the full-precision (BF16) release of Nemotron 3.5 Lightning — the reference weights of the model, intended primarily as the starting point for customization: post-training (SFT, RL, distillation), domain adaptation, and producing your own quantized or GGUF variants. For latency- and throughput-optimized inference, use the NVFP4 release instead.

The model employs a hybrid **Mixture-of-Experts** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. The Lightning 3.5 model is released alongside a number of speculative decoding methods for faster text generation. The model has **3B active parameters** and **30B parameters in total**.

This model is ready for commercial use.

*For running Nemotron 3.5 Lightning fast — with NVFP4 quantization, W4A16 for broad hardware coverage, and the DSpark recipe for DGX Spark — please see: NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4*


To get quickly started on a single H100 you can use the following command.

Grab the model:

```
export MODEL_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16
```
Run it with vLLM! (vLLM Nightly: `vllm/vllm-openai:v0.27.1`)

```
vllm serve --model $MODEL_CKPT \
    --max-num-seqs 128 \
    --enable-prefix-caching \
    --async-scheduling \
    --mamba-backend flashinfer \
    --mamba-ssm-cache-dtype float16 \
    --enable-mamba-cache-stochastic-rounding \
    --mamba-cache-philox-rounds 5
```
For more details on how to deploy and use the model — see the Quick Start Guide below!

**Governing Download Terms:** Use of this model is governed by the OpenMDW-1.1 model license.

We evaluated our model on the following benchmarks:

| Task | Nemotron-3.5-Lightning-30B-A3B-BF16 | Qwen 3.6 35B A3B | Gemma 4 26B A4B | Nemotron 3 Nano | Nemotron 3 Super | GPT-OSS 20B | 
|---|---|---|---|---|---|---|
| **General Knowledge** |  |  |  |  |  |  | 
| MMLU Pro | 81.94 | **85.63** | 85.20 | 78.46 | 83.89 | 76.40 | 
| AA-Omniscience | 17.50 | 19.47 | 22.17 | 20.15 | **26.68** | 16.62 | 
| **Reasoning** |  |  |  |  |  |  | 
| GPQA Diamond (no tools) | 75.44 | **83.40** | 79.61 | 74.05 | 78.60 | 71.46 | 
| HLE (text-only, no tools) | 11.72 | 19.56 | 17.42 | 10.89 | **20.30** | 13.76 | 
| SciCode | 32.60 | 35.33 | **40.28** | 30.08 | 35.11 | 38.63 | 
| **Coding & Agentic** |  |  |  |  |  |  | 
| SWE-bench Verified | 51.56 | **70.12** | 57.40 | 34.08 | 63.08 | 52.44 | 
| SWE-bench Multilingual | 39.33 | **63.40** | 43.40 | 14.07 | 49.80 | 41.93 | 
| Terminal-Bench 2.1 | 24.58 | **44.38** | 37.22 | 8.29 | 39.61 | 15.17 | 
| PinchBench | 85.37 | **88.07** | 74.70 | 66.11 | 80.36 | 57.20 | 
| BrowseComp | 36.97 | **48.74** | 26.30 | 13.74 | 22.77 | – | 
| τ³-bench (Banking) | 9.28 | 10.52 | **14.02** | 7.01 | 12.37 | – | 
| GDPval-AA-V2 | 832 | **1015** | 807 | 473 | 746 | – | 
| **Instruction Following** |  |  |  |  |  |  | 
| IFBench (loose) | 71.88 | 63.71 | **77.25** | 72.17 | 71.92 | 68.50 | 
| **Long Context** |  |  |  |  |  |  | 
| AA-LCR | 52.00 | **61.06** | 57.56 | 32.75 | 58.44 | 32.88 | 

Accuracy numbers measured by NVIDIA under a consistent harness (NeMo Gym / Nemo Evaluator SDK); they may differ from vendors' self-reported numbers.

For reproducibility, the evaluation recipes, installation instructions, and commands for NVIDIA Nemotron 3.5 Lightning were collected and published in NeMo Gym. The reported results cover the release evaluation suite, including knowledge and reasoning, instruction following, coding, agentic, tool-use, and long-context. Most evaluations use NeMo Gym-native harnesses while a small subset, including SWE-Bench and Terminal-Bench, used NeMo Evaluator natively. The published recipes specify the benchmark-specific containers, prompts, inference parameters, parser configurations, and scoring settings used to produce the results.

Additional harness-level coding-agent results for SWE-Bench Verified and Terminal-Bench 2.1 are shown below.

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 is the full-precision release of a general purpose reasoning and chat model, and is intended primarily for **customization and post-training** rather than direct production inference. It is intended to be used by developers who want to: post-train the model on their own data (SFT, RL via NeMo RL and NeMo Gym, or distillation), adapt it to a domain or task, produce quantized variants (NVFP4, W4A16, GGUF) for their own deployment targets, or run full-precision research and evaluation. English and coding languages are the primary languages, with Spanish, French, German, Italian, and Japanese also supported.

For developers who want to *deploy* Lightning 3.5 directly — in AI agent systems, chatbots, RAG systems, and instruction-following applications — the NVFP4 release is the recommended path, with optimized recipes for data centre and DGX Spark deployments.

Hugging Face — 08/11/2026

- **Architecture Type:** Mixture-of-Experts Hybrid (Mamba + Transformer)
- **Network Architecture:** Nemotron-3-Lightning + Multi-Token Prediction (MTP)
- **Number of model parameters:** 30B Total / 3B Active

The model was pre-trained with over 20T tokens and supports up to 1M context length. The pre-training phase used an NVFP4 recipe. The model includes **Multi-Token Prediction (MTP)** layers, which predict multiple future tokens to provide richer training signals.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 model was pre-trained using an NVFP4 recipe with crawled and synthetic code, math, science, and general knowledge data.
- Software used for pre-training: Megatron-LM

Stage 2: Continued Pre-Training for Multi-Token Prediction (MTP)

