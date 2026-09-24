---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-nano-12b-v2-vl-fp8-hugging-face
title: "nvidia-nvidia-nemotron-nano-12b-v2-vl-fp8-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "Hugging Face", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["fp8", "nvidia", "alignment", "benchmark", "benchmarks", "foundry", "fp4", "gpu", "guardrails", "inference", "license", "nvfp4"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-nano-12b-v2-vl-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 232]
sha256: 08d470af5a6e59adafb303bfc53879cb763379a406460eec0dcf64e2c62f6e8c
---

# nvidia-nvidia-nemotron-nano-12b-v2-vl-fp8-hugging-face

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-Nano-12B-v2-VL-FP8 -->

**Configuration Parsing
			Warning:**Invalid JSON for config file config.json

NVIDIA-Nemotron-Nano-VL-12B-V2-FP8 is the quantized version of the NVIDIA Nemotron Nano VL V2 model, which is an auto-regressive vision language model that uses an optimized transformer architecture. For more information, please check here. The NVIDIA Nemotron Nano VL FP4 QAD model is quantized with TensorRT Model Optimizer.

This model was trained on commercial images for all three stages of training and supports single image inference.

**Governing Terms:**

Your use of the model is governed by the NVIDIA Open License Agreement.

**Additional Information:**

Backbone LLM: NVIDIA-Nemotron-Nano-12B-v2.

Global

Customers: AI foundry enterprise customers

Use Cases: Image summarization. Text-image analysis, Optical Character Recognition, Interactive Q&A on images, Text Chain-of-Thought reasoning

- Build.Nvidia.com [October 28th, 2025] via nvidia/NVIDIA-Nemotron-Nano-VL-12B-V2
- Hugging Face [October 28th, 2025] via nvidia/NVIDIA-Nemotron-Nano-VL-12B-V2-BF16
- Hugging Face [October 28th, 2025] via nvidia/NVIDIA-Nemotron-Nano-VL-12B-V2-FP8
- Hugging Face [October 28th, 2025] via nvidia/NVIDIA-Nemotron-Nano-VL-12B-V2-NVFP4

**Network Type:** Transformer

**Network Architecture:**

Vision Encoder: C-RADIOv2-H

Language Encoder: NVIDIA-Nemotron-Nano-12B-v2

Input Type(s): Image, Text

- Input Images
- Language Supported: German, Spanish, French, Italian, Korean, Portuguese, Russian, Japanese, Chinese, English

Input Format(s): Image (Red, Green, Blue (RGB)), and Text (String)

Input Parameters: Image (2D), Text (1D)

Other Properties Related to Input:

- Context length up to 128K
- Maximum Resolution: Determined by a 12-tile layout constraint, with each tile being 512 × 512 pixels. This supports aspect ratios such as:
  - 4 × 3 layout: up to 2048 × 1536 pixels
  - 3 × 4 layout: up to 1536 × 2048 pixels
  - 2 × 6 layout: up to 1024 × 3072 pixels
  - 6 × 2 layout: up to 3072 × 1024 pixels
  - Other configurations allowed, provided total tiles ≤ 12
- Channel Count: 3 channels (RGB)
- Alpha Channel: Not supported (no transparency)

Output Type(s): Text

Output Formats: String

Output Parameters: One-Dimensional (1D): Sequences up to 128K

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA’s hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

Runtime Engine(s): vLLM

Supported Hardware Microarchitecture Compatibility: H100 SXM 80GB

Supported Operating System(s): Linux

Nemotron-Nano-VL-12B-V2-FP8

```
pip install causal_conv1d "transformers>4.53,<4.54" torch timm "mamba-ssm==2.2.5" accelerate open_clip_torch numpy pillow
```
To serve this checkpoint with vLLM, you can start the docker `vllm/vllm-openai:nightly` and run the sample command below:

```
python3 -m vllm.entrypoints.openai.api_server --model nvidia/Nemotron-Nano-VL-12B-V2-FP8 --trust-remote-code --quantization modelopt
```
**Data Modalities** 

** Total Size: 39'486'703 samples 

** Total Number of Datasets: 270 

** Text-only datasets: 33 

** Text-and-image datasets: 176 

** Video-and-text datasets: 61 

** Total size: 27.7 TB 

** Data modalities: Text, Image, Video 

** Data Collection Method by dataset: Hybrid: Automated, Human, Synthetic 

** Labeling Method by dataset: Hybrid: Automated, Human, Synthetic 

** Dataset partition: Training [100%], Testing [0%], Validation [0%] 

 
** Time period for training data collection: 2023-2025 

 
** Time period for testing data collection: N/A 

 
** Time period for validation data collection: N/A 

The post-training datasets consist of a mix of internal and public datasets designed for training vision language models across various tasks. It includes:

- Public datasets sourced from publicly available images and annotations, supporting tasks like classification, captioning, visual question answering, conversation modeling, document analysis and text/image reasoning.
- Internal text and image datasets built with public commercial images and internal labels, adapted for the same tasks as listed above.
- Synthetic image datasets generated programmatically for specific tasks like tabular data understanding and optical character recognition (OCR), for English, Chinese as well as other languages.
- Video datasets supporting video question answering and reasoning tasks from publicly available video sources, with either publicly available or internally generated annotations.
- Specialized datasets for safety alignment, function calling, and domain-specific tasks (e.g., science diagrams, financial question answering).
- NVIDIA-Sourced Synthetic Datasets for text reasoning.
- Private datasets for safety alignment or VQA on invoices.
- Crawled or scraped captioning, VQA, and video datasets.
- Some datasets were improved with Qwen2.5-72B-Instruct annotations

For around ~30% of our total training corpus and several of the domains listed above, we used commercially permissive models to perform:

- Language translation
- Re-labeling of annotations for text, image and video datasets
- Synthetic data generation
- Generating chain-of-thought (CoT) traces

Additional processing for several datasets included rule-based QA generation (e.g., with templates), expanding short answers into longer responses, as well as proper reformatting. More details can be found here.

** Image based datasets were all scanned against known CSAM to make sure no such content was included in training.

| Type | Data Type | Total Samples | Total Size (GB) | 
|---|---|---|---|
| Function call | text | 8,000 | 0.02 | 
| Image Captioning | image, text | 1,422,102 | 1,051.04 | 
| Image Reasoning | image, text | 1,888,217 | 286.95 | 
| OCR | image, text | 9,830,570 | 5,317.60 | 
| Referring Expression Grounding | image, text | 14,694 | 2.39 | 
| Safety | image, text | 34,187 | 9.21 | 
| Safety | text | 57,223 | 0.52 | 
| Safety | video, text | 12,988 | 11.78 | 
| Text Instruction Tuning | text | 245,056 | 1.13 | 
| Text Reasoning | text | 225,408 | 4.55 | 
| VQA | image, text | 8,174,136 | 2,207.52 | 
| VQA | video, text | 40,000 | 46.05 | 
| Video Captioning | video, text | 3,289 | 6.31 | 
| Video Reasoning | video, text | 42,620 | 49.10 | 
| VideoQA | video, text | 1,371,923 | 17,641.79 | 
| Visual Instruction Tuning | image, text | 1,173,877 | 167.79 | 
| **TOTAL** |  | **24,544,290** | **26,803.75** | 

| Type | Modalities | Total Samples | Total Size (GB) | 
|---|---|---|---|
| Image Reasoning | image, text | 17,729 | 15.41 | 
| Text Reasoning | text | 445,958 | 9.01 | 
| **TOTAL** |  | **463,687** | **24.42** | 

| Type | Modalities | Total Samples | Total Size (GB) | 
|---|---|---|---|
| Image Captioning | image, text | 39,870 | 10.24 | 
| VQA | image, text | 40,348 | 3.94 | 
| VideoQA | video, text | 288,728 | 393.30 | 
| **TOTAL** |  | **368,946** | **407.48** | 

| Type | Data Type | Total Samples | Total Size (GB) | 
|---|---|---|---|
| Code | text | 1,165,591 | 54.15 | 
| OCR | image, text | 216,332 | 83.53 | 
| Text Reasoning | text | 12,727,857 | 295.80 | 
| **TOTAL** |  | **14,109,780** | **433.48** | 

**Properties**

- Additionally, the dataset collection (for training and evaluation) consists of a mix of internal and public datasets designed for training and evaluation across various tasks. It includes: 
  - Internal datasets built with public commercial images and internal labels, supporting tasks like conversation modeling and document analysis.
  - Public datasets sourced from publicly available images and annotations, adapted for tasks such as image captioning and visual question answering.
  - Synthetic datasets generated programmatically for specific tasks like tabular data understanding.
  - Specialized datasets for safety alignment, function calling, and domain-specific tasks (e.g., science diagrams, financial question answering).

For more information about the datasets used to train this model, please see the Public Summary of Training Content

The following external benchmarks are used for evaluating the model: 

Data Collection Method by dataset:  

- Hybrid: Human, Automated 

Labeling Method by dataset:  

- Hybrid: Human, Automated  

**Properties (Quantity, Dataset Descriptions, Sensor(s)):** N/A 

**Dataset License(s):** N/A 

| Benchmark | Score (FP8) | Score (BF16) | 
|---|---|---|
| AI2D | 87.6% | 87.1% | 
| OCRBenchV2 | 61.8% | 62.0% | 
| OCRBench | 85.4% | 85.6% | 
| ChartQA | 89.4% | 89.7% | 
| DocVQA val | 94.3% | 94.4% | 

**Engine:** vLLM 

**Test Hardware:** 

- 1x NVIDIA H100 SXM 80GB

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse. For more detailed information on ethical considerations for this model, please see the Model Card++ Explainability, Bias, Safety & Security, and Privacy Subcards. Please report security vulnerabilities or NVIDIA AI Concerns here.

Users are responsible for model inputs and outputs. Users are responsible for ensuring safe integration of this model, including implementing guardrails as well as other safety mechanisms, prior to deployment.

Outputs generated by these models may contain political content or other potentially misleading information, issues with content security and safety, or unwanted bias that is independent of our oversight.

```
@misc{nvidia2025nvidianemotronnanov2,
      title={NVIDIA Nemotron Nano V2 VL},
      author={NVIDIA},
      year={2025},
      eprint={2511.03929},
      archivePrefix={arXiv},
      primaryClass={cs.LG},
      url={https://arxiv.org/abs/2511.03929},
}
```
- Downloads last month
- 117,348
