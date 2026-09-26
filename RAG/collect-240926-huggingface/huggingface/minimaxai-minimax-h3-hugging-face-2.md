---
id: collect-240926-huggingface/huggingface/minimaxai-minimax-h3-hugging-face-2
title: "Original checkpoint, both task families (SGLang, vLLM):"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "MiniMax", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "attention", "cost", "embeddings", "fine-tuning", "gpus", "inference", "multimodal", "omni", "parameters", "rotary"]
source: docs/RAG/clean_en/huggingface/minimaxai-minimax-h3-hugging-face.md
source_anchor: ""
source_lines: [70, 189]
sha256: 364d5578092b0bc307694b6034752523295f81cad3633dea3617f0ef0f3197b6
---

# Original checkpoint, both task families (SGLang, vLLM):

- For scalability and generalization, we adopt a relatively simple Transformer block design. H3-Omni-Transformer is a 33B-parameter dense, single-stream Transformer, with approximately 13B parameters residing in AdaLN-related branches. Because the AdaLN modulation outputs can be precomputed and cached, these parameters do not need to be loaded for inference-only deployment. We release the complete model weights to support further development, including fine-tuning.
- Neither the attention layers nor the FFN layers contain modality-specific structures. Modality-specific parameters are confined to the input/output layers and the AdaLN branches. In particular, modality-specific AdaLN improves generation quality with relatively low additional training and inference costs.
- The model uses three-dimensional Multimodal Rotary Position Embeddings (MM-RoPE) to represent positional relationships across the temporal and two spatial dimensions, `(t, h, w)` .
- During the final stage of training, we introduce native sparse attention to reduce the computational cost of long sequences. The sparse-attention implementation is not included in the initial open-source release and will be published separately in a future update.

- For H3's 2K-resolution output, instead of using a conventional dedicated super-resolution module, we use the H3 base model to regenerate its own low-resolution result through an in-context manner.
- This approach provides two advantages: (1) the regeneration process can reuse the generative capabilities of H3 base model to the greatest extent possible; and (2) the in-context format can reuse the original multimodal context when producing high-resolution output, allowing it to recover information that conventional super-resolution methods would otherwise have to “guess,” such as small text and fine details.
- In-context regeneration is also an example of task generalization.
- **Due to the complexity of the system, this module is not yet open-sourced. We will release it once it is ready.** We provide an API for validating the official results; see "Full 2K Workflow" below.

To help the community deploy MiniMax H3 correctly, we provide two validation methods.

Since the complete H3 system consists of three modules—H3-Context-IR, H3-Base, and H3-Regenerate-2K—the “Full 2K Workflow” provides an end-to-end validation pipeline for 2K output, combining the Open Platform API with a locally deployed H3-Base. The “Local Deployment of H3-Base” section provides a method for validating 768p output using only a locally deployed H3-Base.

In addition, the “Prompting Guidance” section provides a detailed tutorial to help the community develop their own prompting systems.

MiniMax H3 is released as two task-specific checkpoints. Each checkpoint contains a specialized Omni Transformer Model together with the required processor, tokenizer, text encoder, Visual VAE, and standalone Audio VAE components.

| Checkpoint | Supported Tasks | Input Conditions | Output | Precision | 
|---|---|---|---|---|
| MiniMax-H3 Base FL2VA | Text-to-Audio-Video ( `t2va` ), First/Last-Frame-to-Audio-Video (`fl2va` ) | Text; optional first frame, last frame, or both | Video and audio | BF16 | 
| MiniMax-H3 Base Ref2VA | Reference-to-Audio-Video ( `ref2va` ) | Text with reference images, videos, and/or audio | Video and audio | BF16 | 

The released checkpoints are CFG-distilled Omni Transformer model weights.

Each checkpoint is distributed as a self-contained Hugging Face-style repository with the following components:

```
<TASK>/
├── model_index.json
├── processor/
├── tokenizer/
├── text_encoder/
├── transformer/
├── visual_vae/
└── audio_vae/
```
Download the model. The repository hosts the original checkpoint (`FL2VA/`, `Ref2VA/`) and the diffusers format side by side, so scope the download to what your framework needs:

`model_index.json` is the repository-level public entry. The task-family-specific diffusers indexes remain under `FL2VA/model_index.json` and `Ref2VA/model_index.json`.

```
# Original checkpoint, both task families (SGLang, vLLM):
hf download MiniMaxAI/MiniMax-H3 --include "model_index.json" "FL2VA/*" "Ref2VA/*" --local-dir MiniMax-H3
# Or a single task family:
hf download MiniMaxAI/MiniMax-H3 --include "model_index.json" "FL2VA/*" --local-dir MiniMax-H3
```
diffusers users do not need a manual download: `ModularPipeline.from_pretrained("MiniMaxAI/MiniMax-H3")` fetches exactly the components it needs. See the diffusers documentation for loading recipes.

We recommend the following inference frameworks to serve the model:

- SGLang - see cookbook
- vLLM - see vllm recipes
- diffusers - see diffusers docs
- ComfyUI - see Comfy tutorial; use R2V template / T2V template

Here we use sglang as a deployment example. See the MiniMax-H3 deployment guide for additional deployment configurations.

FL2VA:

```
sglang serve \
  --model-path MiniMaxAI/MiniMax-H3 \
  --num-gpus 4 \
  --ulysses-degree 4 \
  --performance-mode speed \
  --host 0.0.0.0 \
  --port 30010 \
  --model-variant fl2va
```
Ref2VA:

```
sglang serve \
  --model-path MiniMaxAI/MiniMax-H3 \
  --num-gpus 4 \
  --ulysses-degree 4 \
  --performance-mode speed \
  --host 0.0.0.0 \
  --port 30011 \
  --model-variant ref2va
```
The following three use cases T2VA, FL2VA, and Ref2VA demonstrate how to reproduce MiniMax-H3 video-audio generation.

| Use case | Request | Result | 
|---|---|---|
| T2VA | View script | t2va.mp4 | 
| FL2VA | View script | fl2va.mp4 | 
| Ref2VA | View script | ref2va.mp4 | 

This section explains how to combine a locally deployed SGLang service with the official **H3-Context-IR** and **H3-Regenerate-2K** APIs to reproduce the quality of 2K videos generated directly by the MiniMax API.
Before you begin, configure the SGLang endpoint and your MiniMax API credentials:

```
# URL of your SGLang deployment
SGLANG_DEPLOYMENT_URL="<sglang-deployment-url>"
# MiniMax API endpoint (choose one)
# CN
MINIMAX_API_BASE="https://api.minimaxi.com"
# Global
# MINIMAX_API_BASE="https://api.minimax.io"
# API token obtained from the MiniMax platform
TOKEN="<token>"
```
MiniMax platform:

API docs:

- Create H3-2K: use /video-generation-v2-create EN-docs, CN-docs
- H3-Context-IR：use /video-generation-v2-h3-context-ir EN-docs, CN-docs
- H3-Regenerate-2K：use /video-generation-v2-regeneration EN-docs, CN-docs

The examples below encode local H3-Base output files as Base64 Data URLs. For production use, uploading the video to a publicly accessible URL and passing that URL as `base_video` is recommended.

For each case below, we provide reference outputs at both 2K and 768p generated directly through the Open Platform API, making it easier to validate the results.

- Type: Text-to-video
- Duration: 10 seconds
- Aspect ratio: 16:9

