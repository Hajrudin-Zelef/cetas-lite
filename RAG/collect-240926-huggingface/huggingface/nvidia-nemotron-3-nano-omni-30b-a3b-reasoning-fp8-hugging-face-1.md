---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-1
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-28-04"]
keywords: ["agentic", "agents", "blackwell", "fp8", "gpu", "incident", "inference", "llama", "llama.cpp", "mixture of experts", "moe", "multimodal"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 185]
sha256: 2f5f282623e7a9f2298c13717c9e89008c0eee477d0b121744291966169611a9
---

# Log in once; the token is cached at ~/.cache/huggingface/token

<!-- source: https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8 -->

| **Total parameters** | 31B (Mamba2-Transformer hybrid MoE) | 
| **Active parameters** | ~3B per token | 
| **Max context** | 256k tokens | 
| **Modalities (in)** | Video, Audio, Image, Text | 
| **Modality (out)** | Text | 
| **Reasoning mode** | On by default; toggle via `enable_thinking` | 
| **Best for** | Video+speech analysis, document intelligence (OCR/charts/long docs), GUI/agentic workflows, ASR | 
| **Minimum GPU (BF16)** | 1× H100 80GB (single-GPU); 1× B200 / 1× H200 recommended | 
| **Minimum GPU (FP8)** | 1× L40S 48GB; 1× RTX Pro 6000 / 1× B200 recommended | 
| **Minimum GPU (NVFP4)** | 1× RTX 5090 32GB; 1× DGX Spark / 1× Jetson Thor also supported | 
| **Precisions** | BF16 (62 GB) · FP8 (33 GB) · NVFP4 (21 GB) | 

| Mode | temperature | top_p | top_k | max_tokens | reasoning_budget | grace_period | 
|---|---|---|---|---|---|---|
| **Thinking mode** | 0.6 | 0.95 | — | 20480 | 16384 | 1024 | 
| **Instruct mode** | 0.2 | — | 1 | 1024 | — | — | 

NVIDIA Nemotron 3 Nano Omni is a multimodal large language model that unifies video, audio, image, and text understanding to support enterprise-grade Q&A, summarization, transcription, and document intelligence workflows. It extends the Nemotron Nano family with integrated video+speech comprehension, Graphical User Interface (GUI), Optical Character Recognition (OCR), and speech transcription capabilities, enabling end-to-end processing of rich enterprise content such as meeting recordings, M&E assets, training videos, and complex business documents. NVIDIA Nemotron 3 Nano Omni was developed by NVIDIA as part of the Nemotron model family. 

This model is available for commercial use. 

This model was improved using Qwen3-VL-30B-A3B-Instruct, Qwen3.5-122B-A10B, Qwen3.5-397B-A17B, Qwen2.5-VL-72B-Instruct, and gpt-oss-120b. For more information, please see the Training Dataset section below. 

Governing Terms: Use of this model is governed by the  NVIDIA Open Model Agreement

Global 

This model is designed for enterprise customers requiring multimodal understanding capabilities. Expected users include:

- Customer service applications (e.g., Doordash video of drop-off at a given address via OCR, drive-thru order verification)
- Media and Entertainment (M&E) — video and speech analysis, dense captions, video search and summarization
- Document intelligence for AI assistants (contracts, SOW/MSA, scientific discovery, financial documents)
- GUI automation for AI agentic applications (incident management, agentic search, browser agents, email agents)

Build.Nvidia.com 04/28/2026 via URL 

 
Hugging Face 04/28/2026 via:

NGC 04/28/2026  via URL 

**Architecture Type:** Mamba2-Transformer Hybrid Mixture of Experts (MoE) 

**Network Architecture:**

- Nemotron 3 Nano LLM (30B A3B) — 31B-parameter Mamba2-Transformer hybrid MoE backbone with ~3B active parameters per token.
- CRADIO v4-H — vision encoder for image and video frames.
- Parakeet — speech encoder for audio inputs.

**Number of model parameters:** 3.1 x 10^10 (31B A3B) 

**Input Type(s):** Video, Audio, Image, Text 

**Input Format(s):** 

 

- Video: mp4, up to 2 minutes. For 1080p videos, sample up to 1 FPS / 128 frames. For lower-resolution videos such as 720p, higher temporal sampling such as 2 FPS / 256 frames may be used.
- Audio: wav, mp3 files (up to 1 hour), 8kHz and higher sampling rates
- Image: Red, Green, Blue (RGB) (jpeg, png)
- Text: String 

**Input Parameters:** 

- Video: Three-Dimensional (3D)
- Audio: One-Dimensional (1D)
- Image: Two-Dimensional (2D)
- Text: One-Dimensional (1D) 

**Other Properties Related to Input:** 

- Maximum context length up to 256k tokens
- Language support: English only  

**Output Type(s):** Text 

**Output Format(s):** 

- Text: String 

**Output Parameters:** 

- Text: One-Dimensional (1D) 

**Other Properties Related to Output:** 

- Maximum context length up to 256k tokens.
- Supports JSON output format
- Supports reasoning output with chain-of-thought
- Supports tool calling
- Supports word-level timestamps for transcription 

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions. 

**Runtime Engine(s):**

- vLLM 
- NeMo 
- Megatron 
- NeMo-RL 

**Supported Hardware Microarchitecture Compatibility:** 

- NVIDIA Ampere (A100 80GB SXM/NVLink) 
- NVIDIA Blackwell (B200 SXM/NVLink, RTX Pro 6000 SE, DGX Spark, Jetson Thor, RTX 5090) 
- NVIDIA Hopper (H100 SXM/NVLink, H200 SXM/NVLink) 
- NVIDIA Lovelace (L40S) 

**Preferred/Supported Operating System(s):**

- Linux 

**Inference Runtimes:** 

- vLLM 
- TensorRT LLM 
- TensorRT Edge-LLM 
- llama.cpp 
- Ollama 
- SGLang 

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment. 

This AI model can be embedded as an Application Programming Interface (API) call into the software environment described above. 

Nemotron-3-Nano-Omni-30B-A3B-Reasoning 

| Precision | Technical Name | HuggingFace URL | 
|---|---|---|
| BF16 | `Nemotron-3-Nano-Omni-30B-A3B-Reasoning-BF16` | https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-BF16 | 
| FP8 | `Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8` | https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8 | 
| NVFP4 | `Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4` | https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4 | 

```
pip install -U "huggingface_hub[hf_xet]"
 
# Log in once; the token is cached at ~/.cache/huggingface/token
hf auth login
 
# Sanity check: should print your username and orgs
hf auth whoami
```
**Required version:** vLLM **0.20.0** is needed. This means one of these containers:


**CUDA 13.0:** 'vllm/vllm-openai:v0.20.0'
**CUDA 12.9:** 'vllm/vllm-openai:v0.20.0-cu129'

```
docker pull vllm/vllm-openai:v0.20.0
```
**Audio support:** Within the vLLM container, before running `vllm serve`, if *any* audio will be used (including passing `use_audio_in_video: true`):

```
python3 -m pip install "vllm[audio]"
```

```
# vllm serve nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-BF16 \
# vllm serve nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8 \
vllm serve nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4 \
  --host 0.0.0.0 \
  --max-model-len 131072 \
  --tensor-parallel-size 1 \
  --trust-remote-code \
  --video-pruning-rate 0.5 \
  --max-num-seqs 384 \
  --allowed-local-media-path / \
  --media-io-kwargs '{"video": {"fps": 2, "num_frames": 256}}' \
  --reasoning-parser nemotron_v3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \
  --kv-cache-dtype fp8 # Omit this for BF16
```
Efficient Video Sampling: video-pruning-rate=0.5 drops 50% of redundant video tokens; halves video-prefill VRAM/TTFT.

**RTX Pro:** Due to a current bug with FlashInfer + RTX Pro, append: `--moe-backend triton`


For everything not covered here (API examples, reasoning mode, video tuning), follow the general instructions.

Use the upstream multi-arch **vLLM v0.20.0** docker image. Docker will automatically pull the `arm64` variant.

