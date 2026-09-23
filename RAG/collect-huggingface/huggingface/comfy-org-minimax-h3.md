---
id: collect-huggingface/huggingface/comfy-org-minimax-h3
title: "MiniMax H3 (ComfyUI) - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "MiniMax", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["awq", "blackwell", "diffusion", "distribution", "embedding", "embeddings", "fp8", "gpu", "license", "lora", "multimodal", "nvfp4"]
source: docs/RAG/Collect RAG/03_huggingface/Comfy-Org-MiniMax-H3.md
source_anchor: ""
source_lines: [1, 48]
sha256: ccc6c8b7b7e6b46267df0a70c850fdc9dfb3110c2093d4833a925a289eb7d616
---

# MiniMax H3 (ComfyUI) - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Comfy-Org/MiniMax-H3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository (Comfy-Org/MiniMax-H3) is the ComfyUI repackaged distribution of the MiniMax H3 family of video-generation models, rather than a standard LLM card. It consolidates model files from several original repositories — MiniMaxAI/MiniMax-H3, lightx2v/Minimax-h3-Turbo, alibaba-pai/MiniMax-H3-Fun-Controlnet-Union, and Kijai/MiniMax-H3-experimental — into a "Diffusion Single File" format ready for ComfyUI, licensed under the miniMax-H3 community license agreement. It is very popular, with ~21.5 million downloads/month.

The card provides an exact file-layout guide for ComfyUI: `diffusion_models/` holds the H3 diffusion backbones in multiple formats (fl2va = first-last-two-VAE?, ref2va = reference-to-VAE variants; bf16, int8_convrot, fp8_scaled, pruned variants); `text_encoders/` holds the Qwen3-VL-32B text encoder quantized as bf16, int8_convrot, or nvfp4_awq (the NVFP4 quant converted from cybermotaz/Qwen3-VL-32B-Instruct-NVFP4, which does not require a Blackwell GPU); `loras/` holds Turbo LoRAs (fl2v 4-step 768p, fl2v 8-step, ref2v 4-step); `vae/` holds audio and video VAEs (audio VAE fp32, video VAE int8_convrot); `model_patches/` holds Fun ControlNet Union patches; and `embeddings/` holds 10 style embeddings (e.g., minimaxh3_art_is_explosion, blooming_flowers, bullet_time, dark_magic, fire_breath, four_seasons, kiss_camera, spiral_ascent, storm_magic, truman_show) that are invoked in the CLIPTextEncode node as `embedding:<filename>`.

The card recommends preferring `int8_convrot` for diffusion models when running PyTorch with cu130 (CUDA 13), and using `fp8_scaled` only if int8_convrot is unavailable. It also links six official ComfyUI workflow templates: Image-to-Video, Reference-to-Video, Text-to-Video, Image-to-Video Continuation, Multiframe Reference, and Fun ControlNet Union, plus a docs tutorial (docs.comfy.org/tutorials/video/minimax/minimax-h3). This makes the card a practical installation/format reference for running MiniMax H3 video generation inside ComfyUI, including quantization choices, text-encoder variants, Turbo LoRA steps, and embeddings.

## Key points

- ComfyUI repackaging of MiniMax H3 video models (fl2va/ref2va) + Turbo LoRAs + VAEs.
- Qwen3-VL-32B text encoder with bf16 / int8_convrot / nvfp4_awq variants (nvfp4 works without Blackwell).
- Preference order for diffusion models: int8_convrot (cu130) > fp8_scaled.
- 10 style embeddings invokable as `embedding:<name>` in CLIPTextEncode.
- 6 official workflows: T2V, I2V, R2V, I2V continuation, multiframe reference, Fun ControlNet Union.
- miniMax-H3 community license; ~21.5M downloads/month.
- Comprehensive ComfyUI folder-layout guide provided.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Comfy-Org |
| Model name | MiniMax-H3 |
| Type | Diffusion Single File (video generation) for ComfyUI |
| Original sources | MiniMaxAI/MiniMax-H3, lightx2v/Minimax-h3-Turbo, alibaba-pai MiniMax-H3-Fun-Controlnet-Union, Kijai |
| Text encoder | Qwen3-VL-32B (bf16, int8_convrot, nvfp4_awq) |
| Diffusion formats | bf16, int8_convrot, fp8_scaled, pruned |
| LoRAs | fl2v turbo 4-step (768p), fl2v turbo 8-step, ref2v turbo 4-step |
| VAEs | audio VAE fp32, video VAE int8_convrot |
| License | miniMax-h3-community-license-agreement |
| Downloads/month | ~21,517,501 |
| Workflows | 6 official templates + docs.comfy.org tutorial |

## Why this source matters for the RAG

This card is the authoritative ComfyUI-side reference for running the MiniMax H3 open video-generation model locally, documenting exact file formats, quantization variants (bf16/int8/fp8/NVFP4 text encoder), folder layout, LoRAs, VAEs and embeddings. It provides practical, citable data on open image/video generation deployment and ComfyUI integration for multimodal RAG knowledge.
