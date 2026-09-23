---
id: collect-huggingface/huggingface/comfy-org-qwen-image-comfyui
title: "Qwen-Image_ComfyUI - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["qwen", "apache", "diffusion", "fp8", "license", "lora", "nvfp4", "quantization", "safetensors", "text-to-image", "video generation"]
source: docs/RAG/Collect RAG/03_huggingface/Comfy-Org-Qwen-Image_ComfyUI.md
source_anchor: ""
source_lines: [1, 48]
sha256: 497be45bff938cbb0b47ff5de7c26440c1d2b600cbbed5c4f6d76c01c405915b
---

# Qwen-Image_ComfyUI - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Comfy-Org/Qwen-Image_ComfyUI
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Comfy-Org/Qwen-Image_ComfyUI is a repackaging of the Qwen-Image text-to-image diffusion model files for use with ComfyUI. It aggregates weights from three upstream sources: Qwen/Qwen-Image, Qwen/Qwen-Image-2512, and DiffSynth-Studio/Qwen-Image-Distill-Full. The repository provides single-file diffusion model checkpoints, text encoders, and a VAE, organized for drop-in placement in ComfyUI's folder structure.

The repository includes multiple diffusion model variants: `qwen_image_distill_full_bf16`, `qwen_image_distill_full_fp8_e4m3fn`, `qwen_image_2512_bf16`, `qwen_image_2512_fp8_e4m3fn`, `qwen_image_bf16`, `qwen_image_fp8_e4m3fn`, `qwen_image_fp8_hq`, `qwen_image_fp8mixed`, and `qwen_image_nvfp4`. Text encoders include `qwen_2.5_vl_7b.safetensors`, an FP8 scaled version, and an NVFP4 version. The VAE is `qwen_image_vae.safetensors`. Files are placed under `ComfyUI/models/diffusion_models/`, `ComfyUI/models/text_encoders/`, and `ComfyUI/models/vae/`.

The model uses Qwen2.5-VL-7B as its text encoder, reflecting Qwen-Image's strong text rendering and prompt comprehension. The card links to official ComfyUI example workflows, including Qwen Image 2512 text-to-image, Qwen-Image 2512 Turbo (2-step LoRA), InstantX Inpainting ControlNet, InstantX Union ControlNet, Fun Union ControlNet, Qwen-Image ControlNet Model Patch, and Qwen-Image Union Control LoRA. This makes it a practical hub for advanced image generation, inpainting, and structural control.

License is Apache-2.0. The repository is tagged as a Diffusion Single File and comfyui library. It has over 2.6M monthly downloads, indicating very high community usage for local image generation. No code snippets are provided; usage is through ComfyUI's node-based interface and the referenced workflow JSON templates.

## Key points

- ComfyUI repackaging of Qwen-Image (plus Qwen-Image-2512 and Distill-Full variants).
- Multiple formats: BF16, FP8 (e4m3fn, hq, mixed), and NVFP4 diffusion models.
- Uses Qwen2.5-VL-7B as text encoder; dedicated Qwen-Image VAE.
- Includes ControlNet, inpainting, union control, and 2-step Turbo LoRA workflows.
- Apache-2.0 license; Diffusion Single File / comfyui.
- ~2.6M monthly downloads; widely used for local image generation.

## Technical data / figures

| Attribute | Value |
|---|---|
| Model type | Text-to-image diffusion (single-file) |
| Text encoder | Qwen2.5-VL-7B (BF16 / FP8 / NVFP4) |
| VAE | qwen_image_vae.safetensors |
| Quantization variants | BF16, FP8 e4m3fn, FP8 hq, FP8 mixed, NVFP4 |
| Distill variants | qwen_image_distill_full (bf16/fp8) |
| 2512 variants | qwen_image_2512 (bf16/fp8) |
| Upstream sources | Qwen/Qwen-Image, Qwen/Qwen-Image-2512, DiffSynth-Studio/Qwen-Image-Distill-Full |
| License | Apache-2.0 |
| Library | diffusion-single-file, comfyui |
| Monthly downloads | ~2.7M |

## Why this source matters for the RAG

This card documents how Qwen-Image is packaged and run in ComfyUI, covering the practical file layout and quantization options needed for local image generation. It connects the Qwen LLM/VLM ecosystem to diffusion-based creative tooling. It supports questions about local AI image and video generation pipelines.
