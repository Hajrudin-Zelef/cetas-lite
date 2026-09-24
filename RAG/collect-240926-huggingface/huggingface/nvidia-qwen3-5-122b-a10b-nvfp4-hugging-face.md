---
id: collect-240926-huggingface/huggingface/nvidia-qwen3-5-122b-a10b-nvfp4-hugging-face
title: "nvidia-qwen3-5-122b-a10b-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Nvidia", "vLLM"]
dates: ["2026-01-06"]
keywords: ["fp4", "nvfp4", "nvidia", "agent", "agents", "benchmark", "benchmarks", "blackwell", "fp8", "gpu", "inference", "memory"]
source: docs/RAG/clean_en/huggingface/nvidia-qwen3-5-122b-a10b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 120]
sha256: b2b654ace95274ebbf35aec18067fcc6820d388b3defc9e569a7e448ede52b0c
---

# nvidia-qwen3-5-122b-a10b-nvfp4-hugging-face

<!-- source: https://huggingface.co/nvidia/Qwen3.5-122B-A10B-NVFP4 -->

The NVIDIA Qwen3.5-122B-A10B-NVFP4 model is the quantized version of Alibaba's Qwen3.5-122B-A10B model, which is an auto-regressive language model that uses an optimized transformer architecture. For more information, please check here. The NVIDIA Qwen3.5-122B-A10B NVFP4 model is quantized with Model Optimizer.

This model is ready for commercial/non-commercial use.  

This model is not owned or developed by NVIDIA. This model has been developed and built to a third-party’s requirements for this application and use case; see link to Non-NVIDIA (Qwen3.5-122B-A10B) Model Card from Alibaba.

Nvidia Model Optimizer: https://github.com/NVIDIA/Model-Optimizer

Global 

Developers looking to take off-the-shelf, pre-quantized models for deployment in AI Agent systems, chatbots, RAG systems, and other AI-powered applications. 

Huggingface 06/01/2026 via https://huggingface.co/nvidia/Qwen3.5-122B-A10B-NVFP4 

**Architecture Type:** Transformers  

**Network Architecture:** Qwen3.5-122B-A10B Mixture of Experts 

**Number of Model Parameters:** 122B in total and 10B activated 

**Input Type(s):** Text, Image, Video 

**Input Format(s):** String, Red, Green, Blue (RGB), Video (MP4/WebM) 

**Input Parameters:** One-Dimensional (1D), Two-Dimensional (2D), Three-Dimensional (3D) 

**Other Properties Related to Input:** Context length up to 262K 

**Output Type(s):** Text 

**Output Format:** String 

**Output Parameters:** One-Dimensionali (1D): Sequences 

**Other Properties Related to Output:** None  

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA’s hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions. 

**Supported Runtime Engine(s):** 

- **vLLM**

**Supported Hardware Microarchitecture Compatibility:** 

- NVIDIA Blackwell 

**Preferred Operating System(s):** 

- Linux 

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

The model version is NVFP4 version and is quantized with nvidia-modelopt **v0.44.0**  

**Link:** cnn_dailymail, Nemotron-Post-Training-Dataset-v2 

**Data Collection Method by dataset:** Automated. 

**Labeling Method by dataset:** Automated. 

**Properties:** The cnn_dailymail dataset is an English-language dataset containing just over 300k unique news articles as written by journalists at CNN and the Daily Mail. The Nemotron-Post-Training-Dataset-v2 is a post-training dataset curated by NVIDIA containing multi-turn conversations across diverse topics. 

**Data Modality:** Undisclosed 

**Data Collection Method by dataset:** Undisclosed 

**Labeling Method by dataset:** Undisclosed 

**Data Size:** Undisclosed 

**Properties:** Undisclosed

**Datasets:** MMMU Pro, GPQA Diamond, tau2_bench_telecom, SciCode, AA-LCR, IFBench 

**Data Collection Method by dataset:** Hybrid: Automated, Human 

**Labeling Method by dataset:** Hybrid: Human, Automated 

**Properties:** We evaluated the model on text-based reasoning and coding benchmarks: MMMU Pro is an advanced benchmark designed to rigorously evaluate multimodal AI models by preventing text-based shortcuts; GPQA Diamond contains 448 graduate-level multiple-choice questions written by domain experts in biology, physics, and chemistry; tau2_bench_telecom is a simulation framework by Sierra Research designed to evaluate conversational AI agents in complex, multi-agent, dynamic environments; SciCode is a scientist-curated benchmark designed to evaluate language models on authentic scientific research programming tasks;  AA-LCR (Artificial Analysis Long Context Recall) evaluates a model's ability to accurately retrieve and recall information from long input contexts; IFBench is a benchmark for evaluating instruction-following capabilities across diverse and structured task constraints. 

**Acceleration Engine:**vLLM** 

**Test Hardware:** NVIDIA B200 

This model was obtained by quantizing the weights and activations of Qwen3.5-122B-A10B to NVFP4 data type, ready for inference with vLLM. Only the weights and activations of the linear operators within transformer blocks in MoE are quantized. This model was obtained by quantizing the weights and activations of Qwen3.5-122B-A10B to NVFP4 data type the number of bits per parameter from 16 to 4 reducing the disk size and GPU memory requirements by approximately 4x.

To serve this checkpoint with vLLM, you can start the docker `docker: nvcr.io/nvidia/vllm:26.04-py3` and run the sample command below:

```
vllm serve nvidia/Qwen3.5-122B-A10B-NVFP4 \
  --trust-remote-code \
  --quantization modelopt_fp4 \
  --kv-cache-dtype fp8 \
  --tensor-parallel-size 1 \
  --reasoning-parser qwen3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder
```
The accuracy benchmark results are presented in the table below:

| Precision | MMMU Pro | GPQA Diamond | SciCode | AA-LCR | IFBench | 
|---|---|---|---|---|---|
| FP8 | 75.90 | 87.37 | 42.16 | 65.5 | 70.91 | 
| NVFP4 | 75.55 | 86.77 | 41.79 | 67.13 | 70.80 | 

Baseline: Qwen3.5-122B-A10B-FP8. Benchmarked with temperature=0.6, top_p=0.95, max num tokens 64000


The base model was trained on data that contains toxic language and societal biases originally crawled from the internet. Therefore, the model may amplify those biases and return toxic responses especially when prompted with toxic prompts. The model may generate answers that may be inaccurate, omit key information, or include irrelevant or redundant text producing socially unacceptable or undesirable text, even if the prompt itself does not include anything explicitly offensive.

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. Developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

Please make sure you have proper rights and permissions for all input image and video content; if image or video includes people, personal health information, or intellectual property, the image or video generated will not blur or maintain proportions of image subjects included.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 1,337,015
