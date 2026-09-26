---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-6
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: ["EU", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["attention", "benchmark", "benchmarks", "blackwell", "fp4", "fp8", "kv cache", "llama", "llama.cpp", "moe", "multimodal", "nvfp4"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [788, 894]
sha256: e5052efecc99866b6b6fd6c80accb2716c27305356779aef18f07053f12ad6cf
---

# Log in once; the token is cached at ~/.cache/huggingface/token

**Properties (Quantity, Dataset Descriptions, Sensor(s)):** 354,587,705 total data items across 1395 datasets. The training data spans five modality combinations: text+audio (259,178,821 samples), text+image (70,143,901 samples), text+video (15,837,673 samples), text+video+audio (8,720,044 samples), and text-only (707,187 samples). Content includes publicly available academic datasets, licensed third-party data, NVIDIA-internal collections, and synthetically generated annotations. The data is primarily in English. No sensor-derived data was used. 

**Benchmark Scores:** 

| Task | Multimodal Benchmarks | Nemotron 3 Nano Omni | Nemotron Nano VL V2 | % Improvement | 
|---|---|---|---|---|
| Grounding | CVBench2D | 83.95 | 78.3 | 6.73 | 
| Document | OCRBenchV2 (EN) | 67.04 | 54.8 | 18.26 | 
| Computer Use | OSWorld | 47.4 | 11.1 | 76.58 | 
| Chart Reasoning | Charxiv Reasoning | 63.6 | 41.3 | 35.06 | 
| Multi-Image Reasoning | MMlongBench Doc | 57.5 | 38 | 33.91 | 
| Math Reasoning | MathVista_MINI | 82.8 | 75.5 | 8.82 | 
| OCR Reasoning | OCR_Reasoning | 54.14 | 33.9 | 33.87 | 
| Video Q/A | Video MME | 72.2 | - | - | 
| Video + Audio Q/A | World Sense | 55.4 | - | - | 
| Video + Audio Q/A | Daily Omni | 74.52 | - | - | 
| Speech Instruction Following | Voice interaction | 89.39 | - | - | 

**Quantization Benchmark Scores:** 

We release FP8 and NVFP4 quantized variants alongside the BF16 model. The FP8 variant quantizes every linear layer in the language model to per-tensor E4M3 (with the exception of the MoE router and `lm_head`) and pairs it with an FP8 KV cache, yielding 8.5 effective bits per weight (32.8 GB). The NVFP4 variant uses a mixed-precision recipe inspired by Nemotron 3 Super: routed MoE experts are quantized to NVFP4 (FP4 E2M1 values with per-block FP8 E4M3 scales over groups of 16 elements and an additional per-tensor FP32 global scale), while the Mamba `in_proj` / `out_proj`, shared experts, and attention `o_proj` are quantized to FP8, yielding 4.98 effective bits per weight (20.9 GB). In both variants the vision and audio encoders and their MLP projectors are kept in BF16. 

The table below reports FP8 & NVFP4 accuracy against a BF16 baseline using non-reasoning mode. Across 9 multimodal benchmarks, both quantized variants stay within 1 point of BF16 on average.

| Footprint | BF16 | FP8 | NVFP4 | 
|---|---|---|---|
| Size (GB) | 61.5 | 32.8 | 20.9 | 
| Effective bpw | 16.00 | 8.5 | 4.98 | 

| Benchmark | BF16 | FP8 | NVFP4 | 
|---|---|---|---|
| MathVista_MINI | 71.90 | 71.05 | 71.30 | 
| Charxiv Reasoning | 49.10 | 48.05 | 47.95 | 
| MMlongBench Doc | 46.10 | 45.84 | 45.78 | 
| OCRBenchV2 (EN) | 65.80 | 65.63 | 65.77 | 
| CVBench2D | 84.20 | 85.62 | 85.27 | 
| Video MME | 70.80 | 69.40 | 69.60 | 
| Daily Omni | 74.50 | 74.06 | 74.23 | 
| World Sense | 55.20 | 54.40 | 54.60 | 
| MMAU | 74.62 | 74.56 | 74.34 | 
| Tedium Long (WER↓) | 3.11 | 3.12 | 3.04 | 
| HF-ASR (WER↓) | 5.95 | 5.97 | 5.95 | 
| **Mean (9 non-ASR)** | **65.80** | **65.40** | **65.43** | 
| **Median (9 non-ASR)** | **70.80** | **69.40** | **69.60** | 
| **Δ vs BF16 (mean)** | --- | −0.40 | −0.38 | 

Data Collection Method by dataset: 

- Hybrid: Human, Automated — Evaluation benchmarks are primarily human-curated public academic datasets with automated scoring. 

Labeling Method by dataset: 

- Human 

**Properties (Quantity, Dataset Descriptions, Sensor(s)):** 14 evaluation benchmarks spanning image understanding (MathVistaMini, Charxiv Reasoning, MMLongBench-Doc, OCR Reasoning, OCRBenchV2 English, CVBench2D, OSWorld), video understanding (Video MME), audio/speech understanding (VoiceBench, Tedium Long, HF-ASR, MMAU, World Sense), and multimodal omni-understanding (Daily Omni). All benchmarks are publicly available academic datasets in English. 

Prior to training this model, NVIDIA implemented measures to respect EU text and data mining opt-outs by (1) respecting robots.txt instructions to the extent such signals reflect valid rights reservations, and (2) filtering datasets on any actionable metadata identifiers provided by rightsholders.

**Acceleration Engine:** TensorRT-LLM, vLLM, TensorRT Edge-LLM, llama.cpp, ollama, SGlang 

**Test Hardware:** 

- NVIDIA H100 SXM 
- NVIDIA H200 SXM 
- NVIDIA B200 SXM 
- NVIDIA A100 80GB SXM 
- NVIDIA GB200 NVL72 
- NVIDIA RTX PRO 6000 SE Blackwell 
- NVIDIA L40S PCIe 48GB 
- NVIDIA DGX Spark 
- NVIDIA Jetson Thor 
- NVIDIA RTX 5090

We recommend following settings for reaching the optimal performance.

We suggest the following sampling parameters based on the mode and tasks.

- Thinking mode for long document analysis and multimodal reasoning tasks: 
 `temperature=0.6` ,`top_p=0.95` ,`grace_period=1024` ,`reasoning_budget=16384` ,`max_token=20480` , and`max_model_len=210000`
- Instruct mode (non-thinking) for general tasks:
 `temperature=0.2` ,`top_k=1`
- For ASR tasks, we recommend non-thinking mode with 
 `temperature=1.0` ,`top_k=1`

For most multimodel reasoning tasks, we recommend using output length of at least 20480. For complex reasoning questions especially in math and programing increasing the maximum output length to 210000 tokens can give the model enough room to produce more detailed and correct answers. We also found the proposed Budget-Controlled Reasoning effectiveness in answering complex reasoning questions.

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse. 

Please make sure you have proper rights and permissions for all input image and video content; if image or video includes people, personal health information, or intellectual property, the image or video generated will not blur or maintain proportions of image subjects included. 

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, Explainability, Safety & Security, and Privacy Subcards. 

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here. 

```
@misc{nvidia2026nemotron3nanoomni,
      title={Nemotron 3 Nano Omni: Efficient and Open Multimodal Intelligence},
      author={NVIDIA},
      year={2026},
      eprint={2604.24954},
      archivePrefix={arXiv},
      primaryClass={cs.LG},
      url={https://arxiv.org/abs/2604.24954},
}
```
- Downloads last month
- 1,079,861
