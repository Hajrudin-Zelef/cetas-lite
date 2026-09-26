---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-1
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "China", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "Z.ai"]
dates: ["2025-09", "2025-12", "2026-04", "2026-04-06", "2026-05", "2026-06-04"]
keywords: ["agent", "agentic", "agents", "attention", "aws", "benchmark", "benchmarks", "compute", "glm", "gpu", "kimi", "license"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [1, 95]
sha256: 030702a9d38aac4c5e271709d8c75f97bcef8e08ff62d7205874e01700e3959a
---

# Set the IP for the head node in RAY_HEAD_IP

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 -->

| **Total Parameters** | 550B (55B active) | 
| **Architecture** | LatentMoE - Mamba-2 + MoE + Attention hybrid with Multi-Token Prediction (MTP) | 
| **Context Length** | Up to 1M tokens | 
| **Minimum GPU Requirement** | 8x GB200/B200/GB300/B300, 16x H100, 8x H200 | 
| **Supported Languages** | English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, and Chinese | 
| **Best For** | Frontier reasoning, complex agentic workflows, long-context analysis, tool use, multilingual reasoning, high-stakes RAG | 
| **Reasoning Mode** | Configurable on/off via chat template ( `enable_thinking=True/False` ) | 
| **License** | OpenMDW License Agreement, version 1.1 | 
| **Release Date** | June 4, 2026 | 

For more details on how to deploy and use the model - see the Quick Start Guide below!

*For running Nemotron 3 Ultra on a smaller footprint, please see: NVIDIA-Nemotron-3-Ultra-550B-A55B-NVFP4*


**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - April 2026

**Data Freshness:**

- The post-training data has a cutoff date of May 2026.
- The pre-training data has a cutoff date of September 2025.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

**Nemotron-3-Ultra-550B-A55B-BF16** is a frontier-scale large language model (LLM) trained by NVIDIA, designed to deliver strong agentic, reasoning, and conversational capabilities. It is optimized for the most demanding workloads, including complex multi-step agents, long-context analysis, and high-accuracy reasoning over code, math, and science. Like other models in the family, it responds to user queries and tasks by first generating a reasoning trace and then concluding with a final response. The model's reasoning capabilities can be configured through a flag in the chat template.

The model employs a hybrid **Latent Mixture-of-Experts (LatentMoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. Like the Super model, the Ultra model incorporates **Multi-Token Prediction (MTP)** layers for faster text generation and improved quality, and it is trained using an **NVFP4** pre-training recipe to maximize compute efficiency. The model has **55B active parameters** and **550B parameters in total**.

The supported languages include: English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, and Chinese.

This model is ready for commercial and non-commercial use.

**Governing Download Terms:** Use of this model is governed by the OpenMDW License Agreement, version 1.1 (OpenMDW-1.1).

| Benchmark | N-3-Ultra 550B-A55B | MiniMax-2.7 230B-A10B | GLM-5.1 744B-A40B | Kimi-K2.6 1T-A32B | Qwen-3.5 397B-17B | DS-v4-Pro 1.6T-A49B | DS-v4-Flash 284B-A13B | 
|---|---|---|---|---|---|---|---|
| **Agentic** |  |  |  |  |  |  |  | 
| Terminal Bench 2.1 | 56.4 | 55.5 | 59.3 | 67.2 | 49.9 | 49.2 | 54.2 | 
| GDPVal | 46.7 | 47.6 | 54.7 | 50.4 | 34.6 | 54.6 | 50.2 | 
| SWE-Bench Verified | 70.7 | 75.3 | 76.2 | 75.7 | 73.6 | 74.5 | 73.5 | 
| SWE-Bench Multilingual | 67.7 | 71.8 | 74.8 | 77.1 | 70.9 | 76.5 | 75.0 | 
| ProfBench (Search) | 56.0 | 52.0 | 46.0 | 56.0 | 53.0 | 59.9 | 57.0 | 
| PinchBench | 90.0 | 77.6 | 81.2 | 90.2 | 86.6 | 88.6 | 91.3 | 
| TauBench V3 |  |  |  |  |  |  |  | 
| Airline | 81.5 | 75.3 | 85.0 | 85.8 | 76.5 | 80.8 | 80.8 | 
| Retail | 86.4 | 84.9 | 84.1 | 82.9 | 88.5 | 88.9 | 89.1 | 
| Telecom | 92.9 | 89.6 | 96.9 | 97.8 | 98.0 | 96.3 | 98.3 | 
| Banking | 22.6 | 14.6 | 12.8 | 23.1 | 20.9 | 25.9 | 26.7 | 
| Average | 70.9 | 66.1 | 69.7 | 72.4 | 71.0 | 73.2 | 73.7 | 
| BrowseComp | 44.4 | 54.1 | 59.4 | 61.3 | 40.5 | 59.4 | 46.9 | 
| Vals.ai Financial Agent 1.1 |  |  |  |  |  |  |  | 
| without web search | 60.1 | 51.3 | 60.2 | 54.0 | 61.3 | 58.9 | 58.4 | 
| with web search | 53.7 | 50.5 | 60.7 | 58.8 | 59.0 | 62.3 | 60.1 | 
| **Reasoning and Knowledge** |  |  |  |  |  |  |  | 
| IOI 2025 | 570.0 | -- | 456.5 | 585.0 | 441.3 | 580.1 | -- | 
| LiveCodeBench (v6) | 89.0 | 77.2 | 85.7 | 90.2 | 79.3 | 92.5 | 90.9 | 
| IMOAnswerBench (no tools) | 88.6 | 68.3 | 86.8 | 91.1 | 83.1 | 93.0 | 91.1 | 
| IMOAnswerBench (with tools) | 92.3 | 75.1 | 91.1 | 93.71 | 84.51 | 85.4 | 89.6 | 
| Apex-Shortlist (no tools) | 74.9 | 28.9 | 71.1 | 77.4 | 61.4 | 85.8 | 82.4 | 
| Apex-Shortlist (with tools) | 84.8 | 51.9 | 79.0 | 73.2 | 60.4 | 86.5 | 82.0 | 
| GPQA (no tools) | 87.0 | 86.6 | 86.1 | 91.0 | 87.1 | 87.8 | 88.5 | 
| SciCode (subtask) | 44.6 | 38.3 | 47.7 | 52.0 | 48.0 | 50.5 | 48.2 | 
| HLE (no tools) | 26.7 | 23.1 | 27.2 | 34.8 | 28.5 | 37.7 | 32.2 | 
| HLE (with tools) | 37.4 | -- | 50.4 | 54.0 | 48.3 | 48.2 | 45.1 | 
| CritPt (no tools) | 3.1 | 0.6 | 3.7 | 9.1 | 2.4 | 14.0 | 10.6 | 
| MMLU-Pro | 86.8 | 81.9 | 85.9 | 88.1 | 88.3 | 87.5 | 86.4 | 
| OmniScience Accuracy | 24.1 | 20.5 | 31.3 | 35.5 | 35.9 | 46.8 | 39.9 | 
| OmniScience Non-Hallucination | 78.7 | 74.4 | 66.8 | 67.1 | 7.4 | 5.7 | 2.8 | 
| **Chat & Instruction Following** |  |  |  |  |  |  |  | 
| IFBench (prompt loose) | 81.7 | 74.6 | 76.6 | 73.7 | 78.2 | 79.1 | 82.0 | 
| Multi-Challenge | 63.8 | 42.5 | 63.0 | 63.1 | 63.9 | 64.1 | 63.5 | 
| **Long Context** |  |  |  |  |  |  |  | 
| AA-LCR | 65.4 | 69.8 | 66.9 | 70.2 | 68.3 | 67.3 | 62.7 | 
| RULER (1M) | 94.7 | -- | -- | -- | 90.1 | 94.2 | 87.7 | 
| Longbench v2 (≤ 1M) | 61.9 | -- | -- | -- | 68.9 | 62.1 | 57.0 | 
| **Multilingual** |  |  |  |  |  |  |  | 
| MMLU-ProX (avg en/de/fr/es/it/ja/zh/hi/pt/ko) | 83.0 | 78.4 | 85.8 | 85.0 | 86.4 | 85.6 | 84.3 | 
| WMT24++ (en→xx) | 83.7 | 82.8 | 84.4 | 84.5 | 86.8 | 85.9 | 85.9 | 

All evaluation results were collected via Nemo Evaluator SDK. We used three main evaluation harnesses: Nemo Gym, Nemo Skills, and Harbor with extended sandboxing support via AWS ECS on Nemo Evaluator. In addition, the evaluations also used dedicated open-source packaged containers for ScaleAI Multi Challenge Multi Turn Instruction Following and KernelBench. For reproducibility purposes, more details on the evaluation settings and pinned containers can be found in the Nemo Evaluator SDK examples folder and the reproducibility tutorial for Nemotron 3 Ultra.

The following benchmarks are not onboarded yet in our open source tools and for these we used either their official open source implementation or otherwise an internal scaffolding that we plan to open source in the future: BrowseComp with Search, Tau Bench 3, ProfBench with Search, PinchBench, Vals.ai, LongBench v2.

NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 is a frontier-scale general purpose reasoning and chat model intended to be used in English, Code, and supported multilingual contexts. This model is optimized for complex agentic workflows, long-context reasoning, and high-stakes analytical workloads. It is intended to be used by developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. This model is also suitable for complex instruction-following tasks and long-context reasoning over very large documents and codebases.

Hugging Face - 06/04/2026 via Hugging Face

- **Architecture Type:** Mamba2-Transformer Hybrid Latent Mixture of Experts (LatentMoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid LatentMoE
- **Number of model parameters:** 550B Total / 55B Active

