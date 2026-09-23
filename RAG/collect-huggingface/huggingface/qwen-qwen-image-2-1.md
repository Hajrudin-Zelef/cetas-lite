---
id: collect-huggingface/huggingface/qwen-qwen-image-2-1
title: "Qwen-Image-2.1 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face"]
dates: ["2026-09-23"]
keywords: ["qwen", "apache", "attention", "consumer", "inference", "kv cache", "license", "memory", "parameters", "research", "text-to-image"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen-Image-2.1.md
source_anchor: ""
source_lines: [1, 51]
sha256: 2b5983bc5e57f64957458d682dd1b0030e07e98452dc8c9444c8ada734997edc
---

# Qwen-Image-2.1 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen-Image-2.1
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen-Image-2.1 is a unified text-to-image generation and image-editing model from the Qwen team, with a 7B-parameter visual generation component (32 Single-Stream DiT layers) that balances generation quality, inference efficiency, and versatility. It is a Text-to-Image diffusers model using the `QwenImage21Pipeline`, released under the Qwen Research License, with BF16 weights and 28,407 monthly downloads. The card highlights four key improvements: a compact and efficient architecture with mixed-granularity attention and prefix KV cache reuse; native transparency for generating regular or RGBA (transparent) images, editing transparent layers, and extracting subjects from photographs in one model; versatile editing supporting up to 10 reference images with local edits via circles, painted annotations, or separate masks while preserving identity for people and products; and realistic textures with improved typography, portrait lighting, and fine detail.

Quick-start instructions are provided for the Diffusers library: `pip install torch>=2.4.0`, `transformers>=5.17`, diffusers from git, plus accelerate and pillow. Text-to-image uses `QwenImage21Pipeline.from_pretrained("Qwen/Qwen-Image-2.1", torch_dtype=torch.bfloat16).to("cuda")` and `pipe(prompt=..., width=2048, height=2048, num_inference_steps=40, generator=...)`. Image editing passes an `image=input_image` to the same pipeline, and transparent (RGBA) image generation uses a recommended prompt format ("This is an RGBA image with transparency. ..."). Supported aspect ratios are documented as a dict: 1:1 (2048x2048), 4:3 (2400x1792), 3:4 (1792x2400), 3:2 (2528x1696), 2:3 (1696x2528), 16:9 (2752x1536), and 9:16 (1536x2752). Memory optimization is available via `pipe.enable_model_cpu_offload()`.

The card links to ModelScope, Hugging Face, the Qwen blog, an online demo Space, Discord, and WeChat, and to the GitHub repo (QwenLM/Qwen-Image-2.1). It notes Draw Things and DiffusionBee as supported local apps. The model tree shows 8 adapters, 26 finetunes, 51 quantizations, and 77 Spaces using it. The license is the Qwen Research License Agreement (research use; not Apache-2.0).

## Key points

- Unified text-to-image generation and editing model; 7B-param DiT (32 Single-Stream layers).
- Mixed-granularity attention + prefix KV cache reuse for efficiency.
- Native RGBA/transparent image generation and editing; subject extraction from photos.
- Versatile editing: up to 10 reference images, circles/painted annotations/masks, identity preservation.
- Diffusers `QwenImage21Pipeline`; BF16; aspect ratios from 1:1 to 16:9.
- Memory optimization via CPU offload.
- Licensed under Qwen Research License (not open-source Apache).
- 28,407 downloads/month; 77 Spaces; Draw Things / DiffusionBee support.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Qwen |
| Model name | Qwen-Image-2.1 |
| Type | Text-to-Image / image editing (unified) |
| Pipeline | QwenImage21Pipeline (Diffusers) |
| Parameters | 7B (visual generation component) |
| Architecture | 32 Single-Stream DiT layers |
| Weight format | BF16 |
| License | qwen-research |
| Features | RGBA/transparent gen, editing, up to 10 refs, subject extraction |
| Aspect ratios | 1:1, 4:3, 3:4, 3:2, 2:3, 16:9, 9:16 (up to 2752x1536) |
| Inference steps | 40 (example) |
| Local apps | Draw Things, DiffusionBee |
| Downloads/month | 28,407 |

## Why this source matters for the RAG

This card documents a capable open-weights image generation and editing model with concrete Diffusers recipes and RGBA transparency support, valuable for retrieval on generative vision, image editing, and local text-to-image deployment. Its compact 7B size makes it a practical reference for consumer/workstation setups.
