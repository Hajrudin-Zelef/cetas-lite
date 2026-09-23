---
id: collect-huggingface/huggingface/black-forest-labs-flux-1-schnell
title: "FLUX.1 [schnell] - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["apache", "diffusion", "distillation", "gguf", "inference", "license", "open-weight", "parameters", "qwen", "safetensors", "text-to-image"]
source: docs/RAG/Collect RAG/03_huggingface/black-forest-labs-FLUX.1-schnell.md
source_anchor: ""
source_lines: [1, 46]
sha256: 922bbad67cd43c01602b5b885de8fb06457aa8f893f9f9318c14368366057ce3
---

# FLUX.1 [schnell] - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/black-forest-labs/FLUX.1-schnell
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

FLUX.1 [schnell] is a 12-billion-parameter rectified flow transformer developed by Black Forest Labs that generates images from text descriptions. It is the fast, distilled variant of the FLUX.1 family, trained using latent adversarial diffusion distillation, which allows it to produce high-quality images in only 1 to 4 inference steps rather than the many steps required by standard diffusion models. The model delivers cutting-edge output quality and competitive prompt following that matches closed-source alternatives, and is released under the permissive Apache 2.0 license, allowing personal, scientific, and commercial use. It supports English prompting and is distributed as BF16 Safetensors. The model can be run through the Diffusers library using `FluxPipeline`, where recommended settings include `guidance_scale=0.0`, `num_inference_steps=4`, and `max_sequence_length=256`. It is also available via ComfyUI for local node-based workflows and through multiple hosted API endpoints (bfl.ml, replicate.com, fal.ai, mystic.ai). The reference implementation and sampling code live in the dedicated Black Forest Labs GitHub repository. Known limitations are that the model is not intended to provide factual information, may amplify societal biases as a statistical model, may fail to match prompts, and its prompt following is heavily influenced by prompting style. Out-of-scope uses include illegal activity, exploitation of minors, disinformation, generation of personal identifiable information, harassment, non-consensual nudity, and fully automated decisions affecting legal rights. The model reports 525,415 downloads in the last month and has 283 adapters and 38 quantizations in the community ecosystem.

## Key points

- 12B-parameter rectified flow transformer for text-to-image generation by Black Forest Labs.
- Latent adversarial diffusion distillation enables high-quality images in only 1–4 steps.
- Apache 2.0 license permits personal, scientific, and commercial use.
- Recommended Diffusers settings: guidance_scale=0.0, num_inference_steps=4, max_sequence_length=256.
- Run via Diffusers (`FluxPipeline`), ComfyUI, or hosted APIs (bfl.ml, replicate, fal.ai, mystic.ai).
- Gated repository: users must accept conditions and share contact info to access files.
- Limitations: not factual, can amplify biases, prompt following is prompt-style sensitive.

## Technical data / figures

| Attribute | Value |
|---|---|
| Parameters | 12B |
| Architecture | Rectified flow transformer (diffusion) |
| Task | Text-to-Image |
| Language | English |
| Tensor type | BF16 |
| Library | Diffusers (FluxPipeline), Safetensors |
| License | Apache 2.0 |
| Recommended steps | 1–4 |
| Recommended guidance_scale | 0.0 |
| Recommended max_sequence_length | 256 |
| Access | Gated (accept conditions) |
| Downloads last month | 525,415 |
| Community | 283 adapters, 70 finetunes, 38 quantizations |

## Why this source matters for the RAG

This card documents a widely used open-weight image generation model and its distillation-based speed advantage, providing a canonical reference for text-to-image architectures, licensing and recommended inference parameters. It complements the Qwen-Image 2.1 GGUF entry by covering the other major open image-generation family.
