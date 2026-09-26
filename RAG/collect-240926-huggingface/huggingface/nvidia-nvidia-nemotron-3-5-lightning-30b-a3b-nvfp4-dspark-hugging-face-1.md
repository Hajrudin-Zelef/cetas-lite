---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face-1
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-11-08"]
keywords: ["nvfp4", "nvidia", "agentic", "attention", "benchmarks", "blackwell", "data centre", "embedding", "gpu", "gpus", "gqa", "inference"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face.md
source_anchor: ""
source_lines: [1, 90]
sha256: 19d9ae8d7e47f7a47399d38507d9c4421498ecb82d67b94257add1349806fe70
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-dspark-hugging-face

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark -->

The NVIDIA Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark model is the DSpark speculative decoding checkpoint for NVIDIA's Nemotron-3.5-Lightning-30B-A3B model family, which is a hybrid LatentMoE language model designed for reasoning, chat, and agentic workflows. For more information, please check BF16, NVFP4. The NVIDIA Nemotron-3.5-Lightning-30B-A3B-DSpark-NVFP4 model is intended for DSpark speculative decoding deployments tuned for DGX Spark and low-concurrency data centre workflows.

This model is ready for commercial or non-commercial use.

**GOVERNING DOWNLOAD TERMS:** Use of this model is governed by the OpenMDW-1.1 model license.

Global

Developers deploying Nemotron-3.5-Lightning-30B-A3B for reasoning, chat, RAG, and agentic workflows that benefit from lower-latency speculative decoding on DGX Spark and data centre GPUs. This release is intended for DSpark-assisted serving of Nemotron-3.5-Lightning-30B-A3B rather than as a standalone target model checkpoint.

Hugging Face 08/11/2026 via https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark

The DSpark model architecture is as follows:

**Architecture Type:** Dense GQA (Dense MLP + GQA Attention)**Network Architecture:** Dense FFN MLP, and GQA Attention layers; DSpark speculative decoding attention uses causal grouped-query attention (GQA) with a sliding window of size 1024 on all layers, and per-head attention sink bias.**Number of Model Parameters:** 967M total parameters, of which 615M are non-embedding parameters.  

For more information about the underlying model's architecture, please see this Nemotron-3.5-Lightning-30B-A3B-BF16, Nemotron-3.5-Lightning-30B-A3B-NVFP4.

**Input Type(s):** Text**Input Format(s):** String**Input Parameters:** One-Dimensional (1D): Sequences**Other Properties Related to Input:** Maximum context length up to 1M tokens. Supported languages include English, Spanish, French, German, Italian, and Japanese.  

**Output Type(s):** Text**Output Format:** String**Output Parameters:** One-Dimensional (1D): Sequences**Other Properties Related to Output:** Outputs may include natural-language responses, reasoning traces, tool-use content, and structured outputs depending on chat-template configuration and application-level tooling.  

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

**Supported Runtime Engine(s):**

- vLLM

**Supported Hardware Microarchitecture Compatibility:**

- NVIDIA Blackwell - including DGX Spark (GB10)
- NVIDIA Hopper

**Preferred Operating System(s):**

- Linux

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

The model is a Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark speculative decoding version quantized with Model Optimizer **0.45.0 version**.

For more information about the underlying Nemotron-3.5-Lightning-30B-A3B model, please visit the model card - Nemotron-3.5-Lightning-30B-A3B-NVFP4, Nemotron-3.5-Lightning-30B-A3B-BF16.

**Link**: Nemotron-Post-Training-Dataset-v2, and Nemotron-Post-Training-Dataset-v3 only prompts from the datasets were used for data synthesis, (the original responses from GPT were not used), which is then used to train the DSpark modules.**Data Modality:** Text**Text Training Data Size:** 66 Billion Tokens, repeated for 2 epochs**Data Collection Method by dataset:** Hybrid: Automated, manually-collected, Synthetic**Labeling Method by dataset:** Hybrid: Automated, manually-labelled, Synthetic**Properties:** The Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark model was trained exclusively on a post-training corpus includes synthetic data for reasoning, code, science, tool use, instruction following, structured outputs, and multilingual tasks. Dataset disclosures are published in the `nvidia/nemotron-post-training-dataset-v2` and `nvidia/nemotron-post-training-v3` Hugging Face collections. 

For more information about the datasets used to train this model, please see the Public Summary of Training Content

The model is evaluated on the SPEED-Bench dataset for speculative decoding benchmarking.

**Data Collection Method by dataset:** Hybrid: Automated, manually-collected, Synthetic**Labeling Method by dataset:** Hybrid: Automated, manually-labelled, Synthetic**Properties:** This corpus comprises a mix of high-quality standard benchmarks and test suites for LLMs. These benchmarks prompt the underlying model on a diverse set of tasks such as coding, math, writing, translation, etc., and the acceptance rates achieved from speculative decoding over the responses are measured in order to evaluate the quality of the speculation model.

**Acceleration Engine:** vLLM, llama.cpp**Test Hardware:** NVIDIA Hopper - H100; NVIDIA Blackwell - GB200; NVIDIA Blackwell - GeForce RTX 5090; NVIDIA Blackwell - DGX Spark (GB10)

Nemotron-3.5-Lightning-30B-A3B supports multiple speculative decoding paths, including native MTP, DFlash, and DSpark for DGX Spark and low-concurrency data centre workflows. This release is the DSpark speculative decoding checkpoint intended to be paired with the Nemotron-3.5-Lightning-30B-A3B target model in vLLM. DSpark is designed to improve accepted length while preserving the latency benefits of block drafting on compact Blackwell systems such as DGX Spark.

To serve the checkpoint with vLLM, refer to the DSpark recipes in the NVIDIA Nemotron-3.5-Lightning-30B-A3B Model Card

Acceptance rate on SPEED-Bench with draft length 7:

| Category | SPEED-Bench Acceptance Length | 
|---|---|
| coding | **4.38** | 
| humanities | **3.18** | 
| math | **4.17** | 
| multilingual | **4.55** | 
| qa | **3.36** | 
| rag | **4.25** | 
| reasoning | **3.90** | 
| roleplay | **3.06** | 
| stem | **3.40** | 
| summarization | **4.15** | 
| writing | **2.83** | 
| **Overall Average** | **3.75** | 

Baseline: NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16
Benchmarked with `temperature=1.0`, `top_p=0.95`.


The base model may generate inaccurate, incomplete, or otherwise undesirable responses, even when prompts are benign. It may reflect biases or content artifacts present in its training data, and output quality can vary by domain, language, reasoning depth, and context length. Developers should add application-specific safeguards and evaluate the system in the intended deployment environment before production use.

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. Developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, Explainability, Safety & Security, and Privacy Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

SUBCARDS:

