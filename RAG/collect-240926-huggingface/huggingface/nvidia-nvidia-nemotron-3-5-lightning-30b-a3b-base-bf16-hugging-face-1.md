---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face-1
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Hugging Face", "Nvidia"]
dates: ["2025-09", "2025-12", "2026-05", "2026-11-08"]
keywords: ["nvidia", "agents", "attention", "benchmark", "blackwell", "compute", "distillation", "fine-tuning", "gpu", "inference", "license", "mixture of experts"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face.md
source_anchor: ""
source_lines: [1, 110]
sha256: d74f9a566fe0605a745262d67150869267e23a533ae8cb1a32c4181a929f03e1
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 -->

*Looking for the post-trained model? See NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 for the full-precision release, or NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 for optimized inference.*


**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - May 2026

**Data Freshness:**

- The pre-training data has a cutoff date of September 2025.

**NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16** is a large language model (LLM) trained by NVIDIA. This is the base (pre-trained) checkpoint of the Nemotron 3.5 Lightning family — no supervised fine-tuning, reinforcement learning, or distillation has been applied — making it the natural starting point for developers and researchers building their own post-trained models.

The model employs a hybrid **Mixture-of-Experts (MoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. The model incorporates **Multi-Token Prediction (MTP)** layers — trained via a dedicated continued pre-training phase — for richer training signals and native speculative decoding, and it is pre-trained using an **NVFP4** recipe to maximize compute efficiency. The model has **3B active parameters** and **30B parameters in total**.

The pre-training corpus spans English, 19 other spoken languages, and 43 programming languages.

This model is ready for commercial and non-commercial use.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

Use of this model is governed by the OpenMDW License Agreement, version 1.1 (OpenMDW-1.1).

| Benchmark | Qwen3.5-35B-A3B | Gemma-4-26B-A4B | Nemotron-3 Nano 30B-A3B | Nemotron-3.5 Lightning | Nemotron-3 Super 120B-A12B | 
|---|---|---|---|---|---|
| **General** |  |  |  |  |  | 
| MMLU | 81.07 | 77.81 | 78.48 | 78.59 | 86.01 | 
| MMLU-Pro (5-shot) | 64.49 | 50.02 | 64.20 | 67.94 | 74.43 | 
| AGIEval-EN (CoT) | 70.51 | 55.28 | 68.45 | 70.02 | 77.92 | 
| **Math** |  |  |  |  |  | 
| GSM8K (8-shot, CoT) | 90.07 | 77.03 | 91.43 | 91.28 | 90.67 | 
| Minerva Math (4-shot) | 59.66 | 43.74 | 82.64 | 82.78 | 84.84 | 
| **Code** |  |  |  |  |  | 
| MBPP (3-shot) | 70.76 | 68.39 | 73.82 | 78.59 | 81.71 | 
| HumanEval | 66.46 | 50.00 | 75.00 | 77.44 | 80.49 | 
| **Commonsense understanding** |  |  |  |  |  | 
| ARC-Challenge (25-shot) | 95.39 | 92.83 | 91.98 | 92.66 | 96.08 | 
| HellaSwag | 85.61 | 85.26 | 85.55 | 85.55 | 88.97 | 
| OpenBookQA | 44.20 | 48.80 | 46.80 | 47.60 | 48.60 | 
| PIQA (acc) | 82.32 | 82.21 | 82.64 | 83.35 | 83.90 | 
| PIQA (acc-norm) | 82.54 | 83.84 | 84.33 | 85.20 | 85.47 | 
| WinoGrande (5-shot) | 79.24 | 79.08 | 79.16 | 79.95 | 78.93 | 
| **Global-MMLU-Lite (5-shot)** |  |  |  |  |  | 
| Average | 80.94 | 74.78 | 74.62 | 75.53 | 85.72 | 
| German (de) | 81.25 | 75.50 | 75.75 | 76.00 | 87.25 | 
| Spanish (es) | 83.25 | 76.50 | 79.25 | 78.00 | 87.25 | 
| French (fr) | 82.00 | 74.50 | 74.75 | 76.25 | 85.75 | 
| Italian (it) | 85.25 | 76.50 | 77.00 | 77.75 | 86.75 | 
| Japanese (ja) | 77.25 | 73.75 | 70.75 | 73.25 | 84.25 | 
| Korean (ko) | 78.25 | 73.50 | 70.50 | 71.25 | 82.50 | 
| Portuguese (pt) | 82.25 | 76.25 | 75.00 | 77.00 | 87.50 | 
| Chinese (zh) | 78.00 | 71.75 | 74.00 | 74.75 | 84.50 | 
| **Multilingual Math — MGSM (8-shot)** |  |  |  |  |  | 
| German (de) | 84.80 | 69.60 | 85.60 | 84.80 | 90.40 | 
| Spanish (es) | 89.20 | 78.80 | 84.40 | 88.00 | 88.00 | 
| French (fr) | 84.40 | 67.20 | 81.60 | 82.40 | 85.60 | 
| Japanese (ja) | 72.80 | 54.80 | 70.40 | 69.60 | 81.60 | 
| Russian (ru) | 90.00 | 76.40 | 86.40 | 87.60 | 91.20 | 
| Chinese (zh) | 85.60 | 68.00 | 82.80 | 78.40 | 85.60 | 
| **Long context — RULER** |  |  |  |  |  | 
| RULER 256K | 82.36 | 85.73 | 71.71 | 76.88 | 83.03 | 
| RULER 1M | 56.43 | 72.93 | 51.23 | 69.62 | 66.98 | 

Accuracy numbers measured by NVIDIA under a consistent harness (NeMo Gym / Nemo Evaluator SDK); they may differ from vendors' self-reported numbers.

All evaluation results were collected via Nemo Evaluator SDK and NVIDIA's open source container of LM Evaluation Harness except for RULER which uses the NeMo Skills harness. The open source container on LM Evaluation Harness packaged via NVIDIA's Nemo Evaluator SDK used for evaluations can be found here. OSS Evaluation Recipes for Nemotron-3.5 Lightning are available in NeMo Gym.

This model is intended for developers and researchers building LLMs. As the base checkpoint of the Nemotron 3.5 Lightning family, it is the recommended starting point for pre-training research, continued pre-training on domain corpora, and building custom post-trained variants (SFT, RL, and distillation) via NeMo RL, NeMo Gym, and Megatron-LM.

Hugging Face - 08/11/2026

- **Architecture Type:** Mamba2-Transformer Hybrid Mixture of Experts (MoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid MoE
- **Number of model parameters:** 30B Total / 3B Active

The model was pre-trained with over 20T tokens and supports up to 1M context length. The pre-training phase used an NVFP4 recipe. The model includes **Multi-Token Prediction (MTP)** layers, which predict multiple future tokens to provide richer training signals and enable faster inference via speculative decoding.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 model was pre-trained using an NVFP4 recipe with crawled and synthetic code, math, science, and general knowledge data.
- Software used for pre-training: Megatron-LM

Stage 2: Continued Pre-Training for Multi-Token Prediction (MTP)

- The model underwent a continued pre-training phase to train its **Multi-Token Prediction (MTP)** layers. In this stage, MTP heads learn to predict multiple future tokens, providing richer training signals to the base model and enabling native speculative decoding at inference time.

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 model is a result of the above work.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. The pre-training corpus spans English, 19 other spoken languages, and 43 programming languages.

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- **Runtime Engine(s):** PyTorch; Megatron-LM; NeMo
- **Supported Hardware Microarchitecture Compatibility:** NVIDIA Ampere - A100; NVIDIA Blackwell; NVIDIA Hopper - H100-80GB
- **Operating System(s):** Linux

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

- v1.0 - GA (08/11/2026)

