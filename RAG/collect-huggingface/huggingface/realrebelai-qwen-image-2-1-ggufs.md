---
id: collect-huggingface/huggingface/realrebelai-qwen-image-2-1-ggufs
title: "Qwen-Image 2.1 GGUFs - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["gguf", "qwen", "attention", "diffusion", "license", "memory", "packaging", "quantization", "text-to-image"]
source: docs/RAG/Collect RAG/03_huggingface/realrebelai-Qwen-Image-2.1_GGUFs.md
source_anchor: ""
source_lines: [1, 45]
sha256: 26baddd8bd8d8662b6075a0df3f2d62b2a6e5329e7e78fe3dbc467f19867fcab
---

# Qwen-Image 2.1 GGUFs - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/realrebelai/Qwen-Image-2.1_GGUFs
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This repository by RealRebelAI provides high-quality mixed-precision GGUF quantizations of the Qwen-Image 2.1 text-to-image diffusion model for use in ComfyUI via the ComfyUI-GGUF custom node. The files are converted from the Comfy-Org Qwen-Image 2.1 BF16 diffusion model and preserve the native Comfy tensor naming and layout, so no additional tensor-name remapping is required. The repository contains only the diffusion model; users still need the standard Qwen-Image 2.1 text encoder and VAE required by their workflow. The key innovation is a Qwen-Image 2.1-specific mixed-precision policy: because a generic Q4_K_M conversion showed visible quality loss in fine structure and anatomy versus the reference INT8 model, precision-sensitive transformer projections (attention Q/K/V/Out, `img_mlp.out`, `img_mlp.gate_up`) are kept at higher precision while the requested quant level remains the base. The corrected Q4_K_M build was A/B tested against the INT8 reference and showed substantial restoration of anatomy, face detail and structural consistency. The quant ladder ranges from Q2_K-HQv3 to Q8_0-HQv3, with recommended starting point Q4_K_M-HQv3 for the best size/quality balance; Q5_K_M, Q6_K and Q8_0 offer more precision, while Q3_K_M and Q2_K target memory-constrained systems. The source transformer uses 32 blocks with 4096 hidden dimension, architecture string `qwen_image`, and standard Comfy attention/MLP tensor names. Top-level modules (`img_in`, `txt_in`, `time_text_embed`, `modulation`, `norm_out`, `proj_out`) remain at higher precision, and reshaped tensors store original shapes under `comfy.gguf.orig_shape.*`. It is an unofficial community quantization; model size is reported as 7B params. Installation requires cloning ComfyUI-GGUF into custom_nodes and placing the GGUF in `ComfyUI/models/diffusion_models/`.

## Key points

- Mixed-precision GGUF quantizations of Qwen-Image 2.1 for ComfyUI + ComfyUI-GGUF.
- Diffusion model only; text encoder and VAE are still required separately.
- Custom HQv3 precision policy protects attention and MLP projections that degrade under generic quants.
- Quant ladder: Q2_K, Q3_K_M, Q4_K_M (recommended), Q5_K_M, Q6_K, Q8_0, all -HQv3.
- Source: Comfy-Org/Qwen-Image-2.1 BF16 diffusion model; architecture `qwen_image`, 32 blocks, 4096 hidden dim.
- Unofficial community quantization by RealRebelAI; reported model size 7B params.
- 10,135 downloads last month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Base model | Qwen/Qwen-Image-2.1 (Comfy-Org packaging) |
| Pipeline | Text-to-Image |
| Architecture | qwen_image |
| Transformer blocks | 32 |
| Hidden dimension | 4096 |
| Reported model size | 7B params |
| Format | GGUF (mixed precision, HQv3 policy) |
| Quant variants | Q2_K, Q3_K_M, Q4_K_M, Q5_K_M, Q6_K, Q8_0 (all HQv3) |
| Recommended quant | Q4_K_M-HQv3 |
| License | Not specified on card (base model governs) |
| Runtime | ComfyUI + ComfyUI-GGUF |
| Downloads last month | 10,135 |

## Why this source matters for the RAG

This card documents practical, community-driven quantization of an image-generation model, capturing the trade-offs between file size and visual quality that matter for local ComfyUI deployment. It adds image-generation and GGUF quantization coverage to a knowledge base otherwise dominated by language models.
