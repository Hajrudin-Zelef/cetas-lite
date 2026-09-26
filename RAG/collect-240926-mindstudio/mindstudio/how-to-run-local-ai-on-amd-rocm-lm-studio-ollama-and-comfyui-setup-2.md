---
id: collect-240926-mindstudio/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup-2
title: "Download the amdgpu-install package (check AMD's site for the latest version)"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Mistral", "Nvidia", "TensorRT-LLM"]
dates: []
keywords: ["amd", "gpu", "agent", "agents", "benchmarks", "consumer", "diffusion", "fp8", "gpus", "inference", "llama", "memory"]
source: docs/RAG/clean_en/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup.md
source_anchor: ""
source_lines: [186, 348]
sha256: a8d59bedd7632ffd8ddce9f260ba18c1d027218645b80824abc70367acd9bad8
---

# Download the amdgpu-install package (check AMD's site for the latest version)

```
curl http://localhost:1234/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "local-model",
    "messages": [{"role": "user", "content": "Hello"}]
  }'
```
### Windows Users: DirectML in LM Studio

On Windows, LM Studio supports AMD GPU acceleration via DirectML rather than ROCm. In the model settings, select **Vulkan** or **DirectML** as the backend. Performance is generally lower than native ROCm on Linux, but it works without requiring WSL2.

## Step 4: Image Generation with ComfyUI on AMD

ComfyUI is a node-based interface for Stable Diffusion and other diffusion models. It supports AMD GPUs via ROCm on Linux.

### Install ComfyUI

Clone the repository and set up a Python virtual environment:

```
git clone https://github.com/comfyanonymous/ComfyUI.git
cd ComfyUI
python3 -m venv venv
source venv/bin/activate
```
### Install PyTorch with ROCm Support

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

This is the critical step. You need a PyTorch build compiled against ROCm, not the standard CUDA build:

`pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/rocm6.1`
Check PyTorch’s official install page to confirm the latest ROCm version available, as this URL changes with new releases.

Verify the installation:

```
import torch
print(torch.cuda.is_available())    # Should print True (ROCm uses CUDA-compatible API)
print(torch.cuda.get_device_name(0))  # Should show your AMD GPU
```
### Install ComfyUI Dependencies and Launch

```
pip install -r requirements.txt
python main.py
```
ComfyUI will start and print a local URL (default: `http://127.0.0.1:8188`). Open it in your browser.

### Download Models

Place model files in the appropriate directories:

- **Checkpoints** (SD 1.5, SDXL, SD3, Flux):`models/checkpoints/`
- **VAE** :`models/vae/`
- **LoRAs** :`models/loras/`
- **ControlNet** :`models/controlnet/`

For SDXL on 8GB VRAM, add `--lowvram` flag to the launch command:

`python main.py --lowvram`
For very limited VRAM, `--novram` forces CPU offloading (much slower but avoids OOM errors).

### Performance Expectations on AMD

ComfyUI on AMD with ROCm is fast but typically 10–30% slower than an equivalent NVIDIA card due to software optimization differences. An RX 7900 XTX will outperform an RTX 3080 but trail an RTX 4080 in most benchmarks.

SDXL generation times at 1024×1024 (20 steps, Euler a sampler):

- **RX 7900 XTX** : ~4–6 seconds
- **RX 6800 XT** : ~7–10 seconds
- **RX 7800 XT** : ~8–12 seconds

Flux models require significantly more VRAM (10GB+ for FP8 quantized versions) and are slower overall.

## Troubleshooting Common AMD Setup Issues

### GPU Not Detected by Ollama or PyTorch

The most common cause is missing group permissions. Check:

`groups $USER`
You need `render` and `video` in the output. If not:

```
sudo usermod -aG render,video $USER
# Then log out and back in — newgrp won't fix this fully
```
### Unsupported GPU Architecture (gfx error)

If you have an unsupported GPU (like RX 6600 with gfx1032 architecture), you can override the architecture detection:

`export HSA_OVERRIDE_GFX_VERSION=10.3.0`
Add this to your `.bashrc` or `.profile` for persistence. This tells ROCm to treat your GPU as a supported variant. It works for many RX 6000 series cards that aren’t on the official list, though not all features are guaranteed.

### Out of Memory Errors During Inference

If you’re hitting VRAM limits:

- Switch to a smaller model or lower quantization (Q4 instead of Q8)
- In ComfyUI, add `--lowvram` or`--novram`
- In Ollama, explicitly pull a smaller quantization: `ollama pull llama3.2:8b-instruct-q4_0`
- Close other GPU-using applications (browsers, games, etc.)

### Slow Performance (CPU Fallback)

If inference is unusually slow (CPU-level speed), the GPU isn’t being used. Check:

`rocm-smi  # Watch during inference — VRAM should fill, GPU utilization should spike`
For PyTorch-based tools, verify the tensor is on the GPU:

```
import torch
x = torch.tensor([1.0]).cuda()
print(x.device)  # Should print cuda:0
```
### ROCm Version Mismatches

Mismatched ROCm versions between system libraries and PyTorch builds cause cryptic errors. Always match:

- The ROCm version installed on your system (`rocminfo | grep "ROCm version"` )
- The ROCm version in your PyTorch install URL

If they don’t match, either upgrade your system ROCm or install the PyTorch build that matches your system version.

## FAQ: Running Local AI on AMD

### Does Ollama officially support AMD GPUs?

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Yes. Ollama has had official ROCm support on Linux since 2024. When you install Ollama on a Linux system with ROCm properly configured, it automatically detects and uses your AMD GPU. No additional configuration is required in most cases.

### Can I run local AI on AMD on Windows without Linux?

Yes, but with limitations. LM Studio supports AMD GPU acceleration on Windows via DirectML, which requires no additional setup beyond installing the application. ComfyUI and Ollama have limited Windows+AMD support — ComfyUI can use DirectML as a backend, but ROCm itself doesn’t run natively on Windows (it requires WSL2). For full GPU acceleration across all tools, Linux is significantly more capable.

### Which AMD GPU is best for local AI?

The **RX 7900 XTX** (24GB VRAM) is the best consumer AMD option for local AI. More VRAM lets you run larger models without quantization compromises. The **RX 7900 GRE** (16GB) offers a good balance of price and capability. For image generation, even an **RX 6800 XT** (16GB) performs well. Avoid low-VRAM cards (8GB or less) if you plan to run 13B+ models.

### How does AMD ROCm performance compare to NVIDIA CUDA for AI?

In most local AI workloads, NVIDIA still has an edge — typically 15–30% faster inference at equivalent price points, largely due to more mature software optimization and libraries like cuDNN and TensorRT being CUDA-exclusive. However, the gap has narrowed significantly. For LLM inference (not training), AMD GPUs with ROCm perform well and the real-world difference is often smaller than benchmarks suggest.

### What’s the minimum VRAM to run local LLMs on AMD?

8GB is the practical minimum for useful LLM inference. With 8GB, you can run 7B parameter models at Q4 quantization (like Llama 3.2 7B or Mistral 7B). Responses will be quick and quality is reasonable. With 16GB you can run 13B models comfortably, or 7B models at higher quality quantization. Below 8GB, you’re limited to very small models (1B–3B) or CPU-heavy offloading.

### Is ComfyUI stable on AMD ROCm?

Reasonably stable for most standard workflows. ComfyUI’s core functionality works well on ROCm — SD 1.5, SDXL, and most ControlNet workflows run without issues. Some custom nodes that use CUDA-specific operations may fail or require modifications. Flux models work but need careful VRAM management. The community around AMD ComfyUI setups is smaller than NVIDIA, but the GitHub issues tracker and relevant subreddits are good resources when something breaks.

## Key Takeaways

