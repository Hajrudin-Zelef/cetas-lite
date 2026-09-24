---
id: collect-240926-huggingface/huggingface/comfy-org-qwen-image-comfyui-hugging-face
title: "comfy-org-qwen-image-comfyui-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["qwen", "diffusion", "fp8", "nvfp4", "safetensors"]
source: docs/RAG/clean_en/huggingface/comfy-org-qwen-image-comfyui-hugging-face.md
source_anchor: ""
source_lines: [1, 45]
sha256: 6081c831b21ed8f0190c74f39b2a3c2a832047a3ed89370949e205cb612cfde0
---

# comfy-org-qwen-image-comfyui-hugging-face

<!-- source: https://huggingface.co/Comfy-Org/Qwen-Image_ComfyUI -->

Repackaged model files for ComfyUI.

Original model repository:

- https://huggingface.co/Qwen/Qwen-Image
- https://huggingface.co/Qwen/Qwen-Image-2512
- https://huggingface.co/DiffSynth-Studio/Qwen-Image-Distill-Full

Place the files in the following folders:

```
📂 ComfyUI/
├── 📂 models/
│   ├── 📂 diffusion_models/
│   │   ├── qwen_image_distill_full_bf16.safetensors
│   │   ├── qwen_image_distill_full_fp8_e4m3fn.safetensors
│   │   ├── qwen_image_2512_bf16.safetensors
│   │   ├── qwen_image_2512_fp8_e4m3fn.safetensors
│   │   ├── qwen_image_bf16.safetensors
│   │   ├── qwen_image_fp8_e4m3fn.safetensors
│   │   ├── qwen_image_fp8_hq.safetensors
│   │   ├── qwen_image_fp8mixed.safetensors
│   │   └── qwen_image_nvfp4.safetensors
│   ├── 📂 text_encoders/
│   │   ├── qwen_2.5_vl_7b.safetensors
│   │   ├── qwen_2.5_vl_7b_fp8_scaled.safetensors
│   │   └── qwen_2.5_vl_7b_nvfp4.safetensors
│   ├── 📂 vae/
│   │   └── qwen_image_vae.safetensors
```
| Workflow | Thumb |  | 
|---|---|---|
| Qwen Image 2512 |  |  | 
| Qwen-Image InstantX Inpainting ControlNet | Before | After | 
| Qwen-Image 2512 Turbo |  |  | 
| Qwen-Image ControlNet Model Patch | Before | After | 
| Qwen-Image InstantX Union ControlNet | Before | After | 
| Qwen-Image 2512: Fun Union ControlNet | Before | After | 
| Qwen-Image: Text to Image |  |  | 
| Qwen-Image Union Control | Before | After | 

- Downloads last month
- 2,689,896
