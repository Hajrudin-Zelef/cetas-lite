---
id: collect-240926-huggingface/huggingface/realrebelai-qwen-image-2-1-ggufs-hugging-face
title: "realrebelai-qwen-image-2-1-ggufs-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["gguf", "qwen", "attention", "diffusion", "llama", "llama.cpp", "memory", "packaging", "quantization", "safetensors"]
source: docs/RAG/clean_en/huggingface/realrebelai-qwen-image-2-1-ggufs-hugging-face.md
source_anchor: ""
source_lines: [1, 134]
sha256: 64d32fd9e0c8081988e0ef16dbf8752f74efac642775689811569d798ff51252
---

# realrebelai-qwen-image-2-1-ggufs-hugging-face

<!-- source: https://huggingface.co/realrebelai/Qwen-Image-2.1_GGUFs -->

High-quality mixed-precision GGUF quantizations of **Qwen-Image 2.1** for use in **ComfyUI** with **ComfyUI-GGUF**.

These files were converted from the **Comfy-Org Qwen-Image 2.1 BF16 diffusion model** and preserve the native Comfy tensor naming/layout. No additional tensor-name remapping is required after conversion.

This repository contains the **diffusion model only**. You still need the normal Qwen-Image 2.1 supporting models required by your ComfyUI workflow, including the text encoder and VAE.


Rebels Q4

Comfy INT8

A standard generic Q4_K_M conversion showed visible quality loss in fine structure and anatomy compared with the reference INT8 model.

To preserve quality, these GGUFs use a **Qwen-Image 2.1-specific mixed-precision policy**. The requested quant level is still the base quant for the model, but precision-sensitive transformer projections are kept at higher precision.

The corrected Q4_K_M build was A/B tested against the INT8 reference with identical generation settings and showed a substantial restoration of anatomy, face detail, and overall structural consistency.

| File | Base quant | Attention Q/K/V/Out | `img_mlp.out` | `img_mlp.gate_up` | 
|---|---|---|---|---|
| Q8_0-HQv3 | Q8_0 | Q8_0 | Q8_0 | Q8_0 | 
| Q6_K-HQv3 | Q6_K | Q8_0 | Q8_0 | Q8_0 | 
| Q5_K_M-HQv3 | Q5_K_M | Q8_0 | Q8_0 | Q6_K | 
| Q4_K_M-HQv3 | Q4_K_M | Q8_0 | Q8_0 | Q5_K | 
| Q3_K_M-HQv3 | Q3_K_M | Q6_K | Q6_K | Q4_K | 
| Q2_K-HQv3 | Q2_K | Q5_K | Q5_K | Q3_K | 

The following top-level Qwen modules remain at their original higher precision where applicable:

```
img_in
txt_in
time_text_embed
modulation
norm_out
proj_out
```
Norm and other small tensors remain high precision where required by the GGUF conversion path.

BF16 diffusion model source:

**Comfy-Org/Qwen-Image-2.1**

```
diffusion_models/qwen_image_2.1_bf16.safetensors
```
The GGUF files use:

```
general.architecture = qwen_image
```
The source transformer uses:

- 32 transformer blocks
- 4096 hidden dimension
- Qwen-Image 2.1 native Comfy tensor naming
- `attn.to_q`
- `attn.to_k`
- `attn.to_v`
- `attn.to_out.0`
- `img_mlp.gate_up`
- `img_mlp.out`

The Comfy-Org BF16 checkpoint already uses the tensor layout expected by ComfyUI, so these conversions preserve those names rather than remapping them to a separate llama.cpp-style naming scheme.

Where a tensor must be physically reshaped to satisfy GGUF quantization requirements, the original logical shape is stored using:

```
comfy.gguf.orig_shape.*
```
ComfyUI-GGUF uses this metadata when loading the model.

For most users, start with:

```
Qwen-Image-2.1-Q4_K_M-HQv3.gguf
```
This is the main balanced release: substantially smaller than higher-precision variants while retaining much better structural quality than a generic Q4 conversion.

If you have more memory available, Q5_K_M, Q6_K, or Q8_0 provide additional precision.

Q3_K_M and Q2_K are intended for more memory-constrained systems and may still show increasing quality loss despite the protected high-precision layers.

Install **ComfyUI-GGUF** in your ComfyUI custom nodes directory.

Portable Windows example:

```
cd /d D:\AI_Tools\ComfyUI_windows_portable\ComfyUI\custom_nodes && git clone https://github.com/city96/ComfyUI-GGUF
```
Install its requirements:

```
D:\AI_Tools\ComfyUI_windows_portable\python_embeded\python.exe -m pip install -r D:\AI_Tools\ComfyUI_windows_portable\ComfyUI\custom_nodes\ComfyUI-GGUF\requirements.txt
```
If ComfyUI-GGUF is already installed, update your existing installation instead of cloning a second copy.

Place the GGUF diffusion model in:

```
ComfyUI/models/diffusion_models/
```
Load it using the GGUF diffusion-model / UNet loader supplied by ComfyUI-GGUF.

Continue using the standard Qwen-Image 2.1 text encoder, VAE, conditioning, sampler, and workflow components around it.

These are **unofficial community quantizations**.

The `HQv3` suffix identifies the Qwen-specific high-quality mixed-precision policy used for these releases. It is intentionally different from a generic llama.cpp Q2/Q3/Q4/Q5/Q6/Q8 conversion.

Lower-bit quantization can still affect:

- fine detail
- hands and anatomy
- typography
- prompt adherence
- edit fidelity
- texture consistency

If a lower quant shows noticeable degradation, move up one quant level.

- **Qwen** — Qwen-Image / Qwen-Image 2.1
- **Comfy-Org** — ComfyUI-compatible Qwen-Image 2.1 model packaging
- **City96 / ComfyUI-GGUF** — GGUF loading and image-model quantization infrastructure
- **llama.cpp / ggml** — GGUF quantization infrastructure

Quantizations and Qwen-Image 2.1 mixed-precision work by **RealRebelAI**.

- Qwen-Image 2.1: https://huggingface.co/Comfy-Org/Qwen-Image-2.1
- ComfyUI-GGUF: https://github.com/city96/ComfyUI-GGUF

- Downloads last month
- 11,389
