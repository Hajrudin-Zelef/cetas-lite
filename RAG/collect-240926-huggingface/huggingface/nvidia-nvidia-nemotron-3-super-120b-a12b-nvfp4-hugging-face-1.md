---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face-1
title: "with uv: uv pip install vllm==0.20.0 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["China", "Hugging Face", "Meta", "Nvidia"]
dates: ["2025-06", "2025-12", "2026-02", "2026-03", "2026-03-11", "2026-11-03"]
keywords: ["agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "compute", "embeddings", "fine-tuning", "fp8", "gpu", "inference"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 94]
sha256: 6b4ef7ba6567a138f4b04023c389bdd57668f66d58f48d824c1ce41cbcc9332e
---

# with uv: uv pip install vllm==0.20.0 --torch-backend=auto

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 -->

| **Total Parameters** | 120B (12B active) | 
| **Architecture** | LatentMoE - Mamba-2 + MoE + Attention hybrid with Multi-Token Prediction (MTP) | 
| **Context Length** | Up to 1M tokens | 
| **Minimum GPU Requirement** | 1× B200 OR 1× DGX Spark | 
| **Supported Languages** | English, French, German, Italian, Japanese, Spanish, Chinese | 
| **Best For** | Agentic workflows, long-context reasoning, high-volume workloads (e.g. IT ticket automation), tool use, RAG | 
| **Reasoning Mode** | Configurable on/off via chat template ( `enable_thinking=True/False` ) | 
| **Speculative Decoding** | Includes a built-in MTP head, with an updated MTPv2 head available as a separate checkpoint | 
| **License** | NVIDIA Nemotron Open Model License | 
| **Release Date** | March 11, 2026 | 

Use `temperature=1.0` and `top_p=0.95` across **all tasks and serving backends** — reasoning, tool calling, and general chat alike.


For more details on how to deploy and use the model - see the Quick Start Guide below!

**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - March 2026

**Data Freshness:**

- The post-training data has a cutoff date of February 2026.
- The pre-training data has a cutoff date of June 2025.

**Nemotron-3-Super-120B-A12B-NVFP4** is a large language model (LLM) trained by NVIDIA, designed to deliver strong agentic, reasoning, and conversational capabilities. It is optimized for collaborative agents and high-volume workloads such as IT ticket automation. Like other models in the family, it responds to user queries and tasks by first generating a reasoning trace and then concluding with a final response. The model's reasoning capabilities can be configured through a flag in the chat template.

The model employs a hybrid **Latent Mixture-of-Experts (LatentMoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. Distinct from the Nano model, the Super model incorporates **Multi-Token Prediction (MTP)** layers for faster text generation and improved quality, and it is trained using **NVFP4** quantization to maximize compute efficiency. The model has **12B active parameters** and **120B parameters in total**.

The supported languages include: English, French, German, Italian, Japanese, Spanish, and Chinese

This model is ready for commercial use.

**Governing Download Terms:** Use of this model is governed by the NVIDIA Nemotron Open Model License.

**Governing Download Terms with NIM:** The NIM container is governed by the NVIDIA Software License Agreement and Product-Specific Terms for AI Products. Use of this model is governed by the NVIDIA Nemotron Open Model License.

| Benchmark | Nemotron-3-Super | Nemotron-3-Super FP8 | Nemotron-3-Super NVFP4 | 
|---|---|---|---|
| **General Knowledge** |  |  |  | 
| MMLU-Pro | 83.73 | 83.63 | 83.33 | 
| **Reasoning** |  |  |  | 
| HMMT Feb25 (with tools) | 94.73 | 94.38 | 95.36 | 
| GPQA (no tools) | 79.23 | 79.36 | 79.42 | 
| LiveCodeBench (v6 2024-08↔2025-05) | 78.69 | 78.44 | 78.44 | 
| LiveCodeBench (v5 2024-07↔2024-12) | 81.19 | 80.99 | 80.56 | 
| SciCode (subtask) | 42.05 | 41.38 | 40.83 | 
| HLE (no tools) | 18.26 | 17.42 | 17.42 | 
| **Agentic** |  |  |  | 
| Terminal Bench (hard subset) | 25.78 | 26.04 | 24.48 | 
| **TauBench V2** |  |  |  | 
| Airline | 56.25 | 56.25 | 54.75 | 
| Retail | 62.83 | 63.05 | 63.38 | 
| Telecom | 64.36 | 63.93 | 63.27 | 
| Average | 61.15 | 61.07 | 60.46 | 
| **Chat & Instruction Following** |  |  |  | 
| IFBench (prompt) | 72.58 | 72.32 | 73.30 | 
| Scale AI Multi-Challenge | 55.23 | 54.35 | 52.8 | 
| Arena-Hard-V2 (Hard Prompt) | 73.88 | 76.06 | 76.00 | 
| **Long Context** |  |  |  | 
| AA-LCR | 58.31 | 57.69 | 58.06 | 
| RULER-500 @ 128k (500 samples per task) | 96.79 | 96.85 | 95.99 | 
| RULER-500 @ 256k (500 samples per task) | 96.60 | 96.33 | 96.52 | 
| RULER-500 @ 512k (500 samples per task) | 96.09 | 95.66 | 96.23 | 
| **Multilingual** |  |  |  | 
| MMLU-ProX (avg over languages) | 79.35 | 79.21 | 79.37 | 

All evaluation results were collected via Nemo Evaluator SDK and for most benchmarks, the Nemo Skills Harness. For reproducibility purposes, more details on the evaluation settings can be found in the Nemo Evaluator SDK configs folder and the reproducibility tutorial for Nemotron 3 Super. The open source container on Nemo Skills packaged via NVIDIA's Nemo Evaluator SDK used for evaluations can be found here. In addition to Nemo Skills, the evaluations also used dedicated open-source packaged containers for Tau-2 Bench (default prompt), Terminal Bench Hard (48 tasks), ScaleAI Multi Challenge Multi-turn Instruction Following, and Ruler.

The following benchmarks are not onboarded yet in our open source tools and for these we used either their official open source implementation or otherwise an internal scaffolding that we plan to open source in the future: SWE Bench Verified (OpenHands), SWE Bench Multilingual (OpenHands), BrowseComp with Search (internal implementation with Serp API), Terminal Bench Core 2.0 (Harbor).

NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 is a general purpose reasoning and chat model intended to be used in English, Code, and supported multilingual contexts. This model is optimized for collaborative agents and high-volume workloads. It is intended to be used by developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. This model is also suitable for complex instruction-following tasks and long-context reasoning.

Hugging Face - 03/11/2026 via Hugging Face

- **Architecture Type:** Mamba2-Transformer Hybrid Latent Mixture of Experts (LatentMoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid LatentMoE
- **Number of model parameters:** 120B Total / 12B Active

The model utilizes the **LatentMoE** architecture, where tokens are projected into a smaller latent dimension for expert routing and computation, improving accuracy per byte. The Super model is pre-trained using NVFP4 quantization — the first model in the Nemotron 3 family trained at this precision. The majority of linear layers use NVFP4 for weights, activations, and gradients, while select layers (including latent projections, MTP layers, QKV/attention projections, and embeddings) are maintained in BF16 or MXFP8 for training stability. The model includes **Multi-Token Prediction (MTP)** layers using a shared-weight design across prediction heads. This improves training signal quality, enables faster inference via native speculative decoding, and supports more stable autoregressive drafting at longer draft lengths compared to independently trained offset heads.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3-Super-120B-A12B-Base-BF16 model was pre-trained for over 25T tokens using crawled and synthetic code, math, science, and general knowledge data. Training leveraged NVFP4 quantization for efficiency. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the pre-training corpus are released in the Nemotron-Pre-Training-Datasets collection.
- Software used for pre-training: Megatron-LM

Stage 2: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the fine-tuning corpus are released in the Nemotron-Post-Training-v3 collection. Data Designer is one of the libraries used to prepare these corpora.

Stage 3: Reinforcement Learning

