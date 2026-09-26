---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-1
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia"]
dates: ["2025-06-25", "2025-09", "2025-11-28", "2025-12", "2026-01-28"]
keywords: ["agent", "agentic", "agents", "attention", "benchmarks", "distillation", "fine-tuning", "fp8", "gpu", "gqa", "inference", "kv cache"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 106]
sha256: c048aeed1c4552bec44bece333bc0cac7494eb071460c2196c8fe93bfe484a35
---

# Load tokenizer and model

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 -->

**Model Developer:** NVIDIA Corporation

**Model Dates:**

September 2025 - December 2025

**Data Freshness:**

- The post-training data has a cutoff date of November 28, 2025.
- The pre-training data has a cutoff date of June 25, 2025.

Nemotron-Nano-3-30B-A3B-NVFP4 is a quantized version of Nemotron-Nano-3-30B-A3B and is a large language model (LLM) trained from scratch by NVIDIA, and designed as a unified model for both reasoning and non-reasoning tasks. It responds to user queries and tasks by first generating a reasoning trace and then concluding with a final response. The model's reasoning capabilities can be configured through a flag in the chat template. If the user prefers the model to provide its final answer without intermediate reasoning traces, it can be configured to do so, albeit with a slight decrease in accuracy for harder prompts that require reasoning. Conversely, allowing the model to generate reasoning traces first generally results in higher-quality final solutions to queries and tasks.

The model employs a hybrid Mixture-of-Experts (MoE) architecture, consisting of 23 Mamba-2 and MoE layers, along with 6 Attention layers. Each MoE layer includes 128 experts plus 1 shared expert, with 6 experts activated per token. The model has 3.5B active parameters and 30B parameters in total.

The supported languages include: English, German, Spanish, French, Italian, and Japanese. Improved using Qwen.

This model is ready for commercial use.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

To get started, you can use our quickstart guide below.

Governing Terms: Use of this model is governed by the NVIDIA Nemotron Open Model License.

We evaluated our model on the following benchmarks:

| Task | NVIDIA-Nemotron-3-Nano-30B-A3B-BF16 | NVIDIA-Nemotron-3-Nano-30B-A3B-FP8 | NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 | 
|---|---|---|---|
| **General Knowledge** |  |  |  | 
| MMLU-Pro | 78.3 | 78.1 | 77.4 | 
| **Reasoning** |  |  |  | 
| AIME25 (no tools) | 89.1 | 87.7 | 86.7 | 
| GPQA (no tools) | 73.0 | 72.5 | 71.9 | 
| LiveCodeBench (v6 2025-08–2025-05) | 68.3 | 67.6 | 65.4 | 
| SciCode (subtask) | 33.0 | 31.9 | 30.7 | 
| HLE (no tools) | 10.2 | 10.3 | 9.4 | 
| **Agentic** |  |  |  | 
| TauBench V2 (Airline) | 48.0 | 44.8 | 41.5 | 
| TauBench V2 (Retail) | 56.9 | 55.6 | 54.1 | 
| TauBench V2 (Telecom) | 42.2 | 40.8 | 41.2 | 
| TauBench V2 (Average) | 49.0 | 47.0 | 45.6 | 
| **Chat & Instruction Following** |  |  |  | 
| IFBench (prompt) | 71.5 | 72.2 | 70.7 | 
| **Long Context** |  |  |  | 
| AA-LCR | 35.9 | 36.1 | 33.3 | 
| **Multilingual** |  |  |  | 
| MMLU-ProX (avg over langs) | 59.50 | 59.6 | 57.8 | 

All evaluation results were collected via Nemo Evaluator SDK and Nemo Skills. The open source container on Nemo Skills packaged via NVIDIA’s Nemo Evaluator SDK used for evaluations can be found here. In addition to Nemo Skills, the evaluations also used dedicated packaged containers for Tau-2 Bench, ArenaHard v2, AA_LCR. A reproducibility tutorial along with all configs can be found in Nemo Evaluator SDK examples. The configs are also available in this HF repo here.

NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 is a general purpose reasoning and chat model intended to be used in English and coding languages. Other non-English languages (English, Spanish, French, German, Japanese, Italian) are also supported. This model is intended to be used by developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. This model is also suitable for typical instruction-following tasks.

January 28, 2026 via Hugging Face

- **Architecture Type:** Mamba2-Transformer Hybrid Mixture of Experts (MoE)
- **Network Architecture:** Nemotron Hybrid MoE
- **Number of model parameters:** 30B

The model was trained with 25T tokens, with a batch size of 3072, and used the Warmup-Stable-Decay (WSD) learning rate schedule with 8B tokens of learning rate warm up, peak learning rate of 1e-3 and minimum learning rate of 1e-5. There are a total of 52 layers, of which there are 23 of each MoE and Mamba-2 and the remaining 6 layers use grouped query attention (GQA) with 2 groups. Each MoE layer includes 128 routed experts plus 1 shared expert, with 6 experts activated per token.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3-Nano-30B-A3B-Base-BF16 model was pre-trained using crawled and synthetic code, math, science, and general knowledge data. All datasets are disclosed in the Training, Testing, and Evaluation Datasets section of this document. Major portions of the pre-training corpus are released in the Nemotron-Pre-Training-Datasets collection.
- Software used for pre-training: Megatron-LM

Stage 2: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. All datasets are disclosed in the Training, Testing, and Evaluation Datasets section of this document. Major portions of the fine-tuning corpus are released in the Nemotron-Post-Training-v3 collection. Data Designer is one of the libraries used to prepare these corpora.
- Software used for supervised fine-tuning: Megatron-LM

Stage 3: Reinforcement Learning

- The model underwent multi-environment reinforcement learning using synchronous GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. Conversational quality was further refined through RLHF using a generative reward model. All datasets are disclosed in the *Training, Testing, and Evaluation Datasets* section of this document. The RL environments and datasets are released as part of NeMo Gym.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

Stage 4: Post-Training Quantization

- The model was quantized to NVFP4 with KV Cache quantized to FP8 using Post-Training Quantization (PTQ). To preserve accuracy while improving efficiency, we applied a selective quantization strategy, keeping the attention layers and the Mamba layers that feed into those attention layers in BF16. After PTQ, Quantization-Aware Distillation (QAD) was applied for further accuracy recovery.
- Software used for quantization: Model Optimizer

NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 model is a result of the above work.

The end-to-end training recipe is available in the NVIDIA Nemotron Developer Repository. Evaluation results can be replicated using the NeMo Evaluator SDK. More details on the datasets and synthetic data generation methods can be found in the technical report NVIDIA Nemotron 3 Nano.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Maximum input size:** 1M tokens
- **Other Properties Related to Input:** Supported languages include: English, Spanish, French, German, Japanese, Italian

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Maximum output size:** 1M tokens

Our AI models are designed and optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- Runtime Engine(s): NeMo 25.11.01
- Supported Hardware Microarchitecture Compatibility: NVIDIA H100-80GB, NVIDIA A100
- Operating System(s): Linux

The snippet below shows how to use this model with Huggingface Transformers (tested on version 4.57.3). We recommend using NeMo Framework 25.11.01 to ensure all required libraries are available.

