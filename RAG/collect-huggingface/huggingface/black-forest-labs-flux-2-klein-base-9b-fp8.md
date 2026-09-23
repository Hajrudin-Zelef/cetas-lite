---
id: collect-huggingface/huggingface/black-forest-labs-flux-2-klein-base-9b-fp8
title: "FLUX.2-klein-base-9b-fp8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Apple", "Hugging Face", "Microsoft", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["fp8", "apache", "diffusion", "fine-tuning", "governance", "inference", "license", "nvidia", "open-weight", "safetensors", "text-to-image", "training"]
source: docs/RAG/Collect RAG/03_huggingface/black-forest-labs-FLUX.2-klein-base-9b-fp8.md
source_anchor: ""
source_lines: [1, 49]
sha256: 817843f2a9058daaea0e35b2a01d127408426f3ac0395761999e1b505515db0c
---

# FLUX.2-klein-base-9b-fp8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/black-forest-labs/FLUX.2-klein-base-9b-fp8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

FLUX.2 [klein] 9B Base is a 9-billion-parameter rectified flow transformer from Black Forest Labs, capable of generating images from text descriptions and supporting multi-reference editing capabilities. This repository holds the FP8 version (flux-2-klein-base-9b-fp8.safetensors) of the model; the full BF16 weights live in the main repo (black-forest-labs/FLUX.2-klein-base-9B). It is positioned as a step "towards interactive visual intelligence" (per the BFL blog post on FLUX.2 klein). Task tags: Image-to-Image, image generation and image editing; Diffusers and Diffusion Single File libraries. The model fits in ~29GB VRAM and is accessible on NVIDIA RTX 4090 and above.

Usage via Diffusers: `DiffusionPipeline.from_pretrained("black-forest-labs/FLUX.2-klein-base-9b-fp8", dtype=torch.bfloat16, device_map="cuda")`, then pipe(image=input_image, prompt=prompt) for editing tasks; switch to "mps" for Apple devices. License: FLUX Non-Commercial License (gated access, contact-info form). Limitations: not intended to provide factual information; text rendering may be inaccurate/distorted; may amplify training-data biases; may fail to match prompts; prompt following heavily influenced by prompting style. Out-of-scope uses prohibited include unlawful content, harm to minors, deceptive content, PII harm, harassment, non-consensual intimate imagery, and fully automated high-risk decision making.

Responsible AI development section details: pre-training NSFW/CSAM filtering (partnership with IWF), multi-round post-training safety fine-tuning (T2I and I2I attacks), ongoing internal/external third-party evaluations, a final third-party release evaluation (9B released under non-commercial license; 4B variants under Apache 2.0), inference NSFW/protected-content filters (in-house + thehive.ai + Microsoft filters), content provenance (pixel-layer watermarking, C2PA metadata), policies, and safety monitoring with a dedicated contact email. Downloads/month: 34,594.

## Key points

- 9B rectified flow transformer for text-to-image and multi-reference image editing.
- This repo holds the FP8 checkpoint; BF16 weights in the main repo.
- Fits in ~29 GB VRAM; runs on NVIDIA RTX 4090 and above.
- Non-commercial license (FLUX Non-Commercial License), gated access.
- Extensive safety mitigations: NSFW/CSAM filtering, safety fine-tuning, third-party evals.
- Content provenance: pixel-layer watermarking and C2PA metadata.
- Diffusers + Diffusion Single File support; device_map="cuda" or "mps".
- 34,594 downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | black-forest-labs |
| Model name | FLUX.2-klein-base-9b-fp8 |
| Architecture | Rectified flow transformer |
| Total params | 9B |
| Precision | FP8 (checkpoint) |
| Task | Text-to-Image / Image-to-Image / multi-reference editing |
| Libraries | Diffusers, Diffusion Single File |
| VRAM | ~29 GB (RTX 4090+) |
| License | FLUX Non-Commercial License |
| Main repo | black-forest-labs/FLUX.2-klein-base-9B (BF16) |
| Downloads/month | 34,594 |

## Why this source matters for the RAG

This card documents a frontier open-weight image generation/editing model's FP8 release, including hardware requirements, licensing, and a detailed responsible-AI safety pipeline. It is essential reference material for retrieval on rectified flow transformers, image editing models, FP8 checkpoints, and open-weight model safety/governance practices.
