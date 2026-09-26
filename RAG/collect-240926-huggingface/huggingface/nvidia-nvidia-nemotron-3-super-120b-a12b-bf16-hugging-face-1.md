---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face-1
title: "with uv: uv pip install vllm==0.18.1 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Hugging Face", "Meta", "Nvidia"]
dates: ["2025-06", "2025-12", "2026-02", "2026-03", "2026-03-11", "2026-11-03"]
keywords: ["agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "compute", "embeddings", "fine-tuning", "gpu", "inference", "license"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face.md
source_anchor: ""
source_lines: [1, 106]
sha256: 7a3f84351febde4ffbb23cb3e794cacbb72129c592811f96ded9a453b37b786d
---

# with uv: uv pip install vllm==0.18.1 --torch-backend=auto

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16 -->

| **Total Parameters** | 120B (12B active) | 
| **Architecture** | LatentMoE - Mamba-2 + MoE + Attention hybrid with Multi-Token Prediction (MTP) | 
| **Context Length** | Up to 1M tokens | 
| **Minimum GPU Requirement** | 8× H100-80GB | 
| **Supported Languages** | English, French, German, Italian, Japanese, Spanish, Chinese | 
| **Best For** | Agentic workflows, long-context reasoning, high-volume workloads (e.g. IT ticket automation), tool use, RAG | 
| **Reasoning Mode** | Configurable on/off via chat template ( `enable_thinking=True/False` ) | 
| **Speculative Decoding** | Includes a built-in MTP head, with an updated MTPv2 head available as a separate checkpoint | 
| **License** | NVIDIA Nemotron Open Model License | 
| **Release Date** | March 11, 2026 | 

Use `temperature=1.0` and `top_p=0.95` across **all tasks and serving backends** — reasoning, tool calling, and general chat alike.


For more details on how to deploy and use the model - see the Quick Start Guide below!

*For running Nemotron 3 Super on a single B200 or DGX Spark - please see: NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4*


**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - March 2026

**Data Freshness:**

- The post-training data has a cutoff date of February 2026.
- The pre-training data has a cutoff date of June 2025.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

**Nemotron-3-Super-120B-A12B-BF16** is a large language model (LLM) trained by NVIDIA, designed to deliver strong agentic, reasoning, and conversational capabilities. It is optimized for collaborative agents and high-volume workloads such as IT ticket automation. Like other models in the family, it responds to user queries and tasks by first generating a reasoning trace and then concluding with a final response. The model's reasoning capabilities can be configured through a flag in the chat template.

The model employs a hybrid **Latent Mixture-of-Experts (LatentMoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. Distinct from the Nano model, the Super model incorporates **Multi-Token Prediction (MTP)** layers for faster text generation and improved quality, and it is trained using **NVFP4** quantization to maximize compute efficiency. The model has **12B active parameters** and **120B parameters in total**.

The supported languages include: English, French, German, Italian, Japanese, Spanish, and Chinese

This model is ready for commercial use.

**Governing Download Terms:** Use of this model is governed by the NVIDIA Nemotron Open Model License.

**Governing Download Terms with NIM:** The NIM container is governed by the NVIDIA Software License Agreement and Product-Specific Terms for AI Products. Use of this model is governed by the NVIDIA Nemotron Open Model License.

| **Benchmark** | **Nemotron 3 Super** | **Qwen3.5-122B-A10B** | **GPT-OSS-120B** | 
|---|---|---|---|
| **General Knowledge** |  |  |  | 
| MMLU-Pro | 83.73 | 86.70 | 81.00 | 
| **Reasoning** |  |  |  | 
| AIME25 (no tools) | 90.21 | 90.36 | 92.50 | 
| HMMT Feb25 (no tools) | 93.67 | 91.40 | 90.00 | 
| HMMT Feb25 (with tools) | 94.73 | 89.55 | — | 
| GPQA (no tools) | 79.23 | 86.60 | 80.10 | 
| GPQA (with tools) | 82.70 | — | 80.09 | 
| LiveCodeBench (v5 2024-07↔2024-12) | 81.19 | 78.93 | 88.00 | 
| SciCode (subtask) | 42.05 | 42.00 | 39.00 | 
| HLE (no tools) | 18.26 | 25.30 | 14.90 | 
| HLE (with tools) | 22.82 | — | 19.0 | 
| **Agentic** |  |  |  | 
| Terminal Bench (hard subset) | 25.78 | 26.80 | 24.00 | 
| Terminal Bench Core 2.0 | 31.00 | 37.50 | 18.70 | 
| SWE-Bench (OpenHands) | 60.47 | 66.40 | 41.9 | 
| SWE-Bench (OpenCode) | 59.20 | 67.40 | — | 
| SWE-Bench (Codex) | 53.73 | 61.20 | — | 
| SWE-Bench Multilingual (OpenHands) | 45.78 | — | 30.80 | 
| **TauBench V2** |  |  |  | 
| Airline | 56.25 | 66.0 | 49.2 | 
| Retail | 62.83 | 62.6 | 67.80 | 
| Telecom | 64.36 | 95.00 | 66.00 | 
| Average | 61.15 | 74.53 | 61.0 | 
| BrowseComp with Search | 31.28 | — | 33.89 | 
| BIRD Bench | 41.80 | — | 38.25 | 
| **Chat & Instruction Following** |  |  |  | 
| IFBench (prompt) | 72.56 | 73.77 | 68.32 | 
| Scale AI Multi-Challenge | 55.23 | 61.50 | 58.29 | 
| Arena-Hard-V2 | 73.88 | 75.15 | 90.26 | 
| **Long Context** |  |  |  | 
| AA-LCR | 58.31 | 66.90 | 51.00 | 
| RULER @ 256k | 96.30 | 96.74 | 52.30 | 
| RULER @ 512k | 95.67 | 95.95 | 46.70 | 
| RULER @ 1M | 91.75 | 91.33 | 22.30 | 
| **Multilingual** |  |  |  | 
| MMLU-ProX (avg over langs) | 79.36 | 85.06 | 76.59 | 
| WMT24++ (en→xx) | 86.67 | 87.84 | 88.89 | 

All evaluation results were collected via Nemo Evaluator SDK and for most benchmarks, the Nemo Skills Harness. For reproducibility purposes, more details on the evaluation settings can be found in the Nemo Evaluator SDK configs folder and the reproducibility tutorial for Nemotron 3 Super. The open source container on Nemo Skills packaged via NVIDIA's Nemo Evaluator SDK used for evaluations can be found here. In addition to Nemo Skills, the evaluations also used dedicated open-source packaged containers for Tau-2 Bench (default prompt), Terminal Bench Hard (48 tasks), ScaleAI Multi Challenge Multi-turn Instruction Following, and Ruler.

The following benchmarks are not onboarded yet in our open source tools and for these we used either their official open source implementation or otherwise an internal scaffolding that we plan to open source in the future: SWE Bench Verified (OpenHands), SWE Bench Multilingual (OpenHands), BrowseComp with Search (internal implementation with Serp API), Terminal Bench Core 2.0 (Harbor).

NVIDIA-Nemotron-3-Super-120B-A12B-BF16 is a general purpose reasoning and chat model intended to be used in English, Code, and supported multilingual contexts. This model is optimized for collaborative agents and high-volume workloads. It is intended to be used by developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. This model is also suitable for complex instruction-following tasks and long-context reasoning.

Hugging Face - 03/11/2026 via Hugging Face

- **Architecture Type:** Mamba2-Transformer Hybrid Latent Mixture of Experts (LatentMoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid LatentMoE
- **Number of model parameters:** 120B Total / 12B Active

The model utilizes the **LatentMoE** architecture, where tokens are projected into a smaller latent dimension for expert routing and computation, improving accuracy per byte. The Super model is pre-trained using NVFP4 quantization — the first model in the Nemotron 3 family trained at this precision. The majority of linear layers use NVFP4 for weights, activations, and gradients, while select layers (including latent projections, MTP layers, QKV/attention projections, and embeddings) are maintained in BF16 or MXFP8 for training stability. The model includes **Multi-Token Prediction (MTP)** layers using a shared-weight design across prediction heads. This improves training signal quality, enables faster inference via native speculative decoding, and supports more stable autoregressive drafting at longer draft lengths compared to independently trained offset heads.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3-Super-120B-A12B-Base-BF16 model was pre-trained for over 25T tokens using crawled and synthetic code, math, science, and general knowledge data. Training leveraged NVFP4 quantization for efficiency. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the pre-training corpus are released in the Nemotron-Pre-Training-Datasets collection.
- Software used for pre-training: Megatron-LM

Stage 2: Supervised Fine-Tuning

