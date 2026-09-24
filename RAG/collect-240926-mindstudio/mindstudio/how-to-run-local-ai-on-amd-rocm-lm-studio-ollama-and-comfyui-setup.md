---
id: collect-240926-mindstudio/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup
title: "Download the amdgpu-install package (check AMD's site for the latest version)"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Mistral", "Nvidia", "OpenAI", "TensorRT-LLM"]
dates: []
keywords: ["amd", "gpu", "agent", "agents", "attention", "benchmarks", "compute", "consumer", "diffusion", "fp8", "gpus", "inference"]
source: docs/RAG/clean_en/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup.md
source_anchor: ""
source_lines: [1, 355]
sha256: 49ee6749fedbf54008179adbe59fb2a12d5203b117727a584939754634a61a09
---

# Download the amdgpu-install package (check AMD's site for the latest version)

<!-- source: https://www.mindstudio.ai/blog/run-local-ai-amd-rocm-lm-studio-ollama-comfyui -->

## AMD GPUs Are Ready for Local AI — Here’s How to Set Them Up

For most of the past few years, running local AI on AMD hardware meant fighting through compatibility issues, hacky workarounds, and half-broken setups. NVIDIA’s CUDA ecosystem had all the tooling, and AMD was an afterthought.

That’s changed. AMD’s ROCm platform has matured significantly, and tools like Ollama, LM Studio, and ComfyUI now ship with first-class AMD support. If you have an RX 6000, RX 7000, or a supported professional card, you can run local large language models and image generation with a legitimate, maintained setup — not a workaround.

This guide walks through the full stack: ROCm installation, Ollama for running LLMs on AMD, LM Studio for a desktop chat interface, and ComfyUI for image generation. By the end, you’ll have a working local AI environment running entirely on your AMD GPU.

## What ROCm Is and Why It Matters

ROCm (Radeon Open Compute) is AMD’s open-source GPU compute platform. It’s the AMD equivalent of NVIDIA’s CUDA — it provides the underlying libraries that let software frameworks like PyTorch offload computation to AMD GPUs.

Without ROCm, every AI tool defaults to CPU-only mode, which is dramatically slower. With ROCm, your GPU handles the heavy lifting: matrix multiplications, attention computations, and all the operations that make running a 7B or 13B parameter model feel responsive.

ROCm supports:

- **PyTorch** (via AMD’s official PyTorch builds)
- **TensorFlow** (via rocTF)
- **HIP** (AMD’s CUDA-compatible programming model)
- **MIOpen** (AMD’s deep learning primitives library)

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

Most AI tools that support AMD GPU acceleration are built on top of these libraries. So getting ROCm right is the foundation everything else depends on.

### Which AMD GPUs Are Supported

ROCm’s official support list includes:

- **RX 7900 XTX, 7900 XT, 7900 GRE** — top-tier consumer support
- **RX 7800 XT, 7700 XT** — good support, slightly lower VRAM
- **RX 6900 XT, 6800 XT, 6800** — well-supported RDNA2 cards
- **RX 6700 XT** — supported with some caveats
- **Radeon Pro and Instinct series** — professionally supported

Cards not on AMD’s official list (like the RX 6600 or 6500 XT) may work with unofficial patches (via `HSA_OVERRIDE_GFX_VERSION`), but results vary. Budget more time for troubleshooting if you’re on an unsupported card.

### Linux vs. Windows

ROCm is primarily a Linux technology. On Linux, you get:

- Native ROCm support for Ollama, PyTorch, and ComfyUI
- Full GPU acceleration for all major AI frameworks
- The most stable and tested path

On Windows, AMD offers **DirectML** as an alternative compute backend. Some tools (like LM Studio) support DirectML-accelerated inference on Windows. Others (like Ollama and ComfyUI) have limited or no native Windows+AMD GPU support — you’d need WSL2 with ROCm passthrough, which works but adds complexity.

If you’re serious about local AI on AMD, Linux (Ubuntu 22.04 or 24.04) is the most reliable environment.

## Step 1: Install ROCm on Linux

### System Prerequisites

Before installing ROCm, confirm your environment:

```
uname -r          # Check kernel version (5.15+ recommended)
lspci | grep -i amd   # Confirm AMD GPU is detected
```
You’ll also need your user in the `render` and `video` groups to access the GPU without root:

`sudo usermod -aG render,video $USER`
Log out and back in after running this.

### Install ROCm via AMD’s Official Repository

AMD maintains package repositories for Ubuntu. The cleanest installation method is the `amdgpu-install` script:

```
# Download the amdgpu-install package (check AMD's site for the latest version)
wget https://repo.radeon.com/amdgpu-install/6.1.3/ubuntu/jammy/amdgpu-install_6.1.60103-1_all.deb
# Install it
sudo apt install ./amdgpu-install_6.1.60103-1_all.deb
# Install ROCm (usecase=rocm installs compute libraries without display drivers)
sudo amdgpu-install --usecase=rocm
```
If you want GPU display drivers alongside ROCm:

`sudo amdgpu-install --usecase=graphics,rocm`
Reboot after installation.

### Verify ROCm Is Working

After rebooting, check that ROCm can see your GPU:

`rocminfo`
You should see your GPU listed with its architecture details. Also run:

`rocm-smi`
This shows GPU utilization, temperature, VRAM usage — the ROCm equivalent of `nvidia-smi`. If both commands return GPU info, your ROCm installation is good.

## Step 2: Run Local LLMs with Ollama on AMD

Ollama is the fastest way to get language models running locally. It handles model downloads, quantization selection, and inference — and it has native ROCm support on Linux.

### Install Ollama

`curl -fsSL https://ollama.com/install.sh | sh`
The installer automatically detects your GPU environment. On a system with ROCm installed, it will configure Ollama to use your AMD GPU.

After installation, start the Ollama service:

`ollama serve`
### Pull and Run a Model

```
# Pull a 7B model (fits in 8GB VRAM with 4-bit quantization)
ollama pull llama3.2
# Run it
ollama run llama3.2
```
For a 13B model, you’ll want at least 12GB VRAM:

```
ollama pull mistral-nemo
ollama run mistral-nemo
```
### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

Ollama will automatically select the right quantization level based on available VRAM. You can also specify explicitly:

`ollama pull llama3.2:8b-instruct-q4_K_M`
### Confirm GPU Is Being Used

Watch GPU utilization while a model is loading or running inference:

`watch -n 1 rocm-smi`
You should see VRAM consumption jump when a model loads and GPU utilization spike during inference. If you see 0% utilization and the CPU is spiking instead, ROCm isn’t being picked up — check that `rocminfo` works and that the Ollama service was started after ROCm was installed.

### Choosing the Right Model Size for Your VRAM

| VRAM | Recommended Model Size |
|---|---|
| 8 GB | 7B models (Q4 quantization) |
| 12 GB | 7B (Q8) or 13B (Q4) |
| 16 GB | 13B (Q8) or 30B (Q4) |
| 24 GB | 70B (Q4) or 30B (Q8) |

The `Q4_K_M` quantization format offers the best balance of speed, quality, and VRAM efficiency for most use cases.

## Step 3: Set Up LM Studio on AMD

LM Studio provides a desktop GUI for running local models — useful if you want a chat interface without using the terminal, or if you want to run an OpenAI-compatible local API server.

### Installation

Download LM Studio from lmstudio.ai. The Linux AppImage works on most distributions.

```
chmod +x LM_Studio-*.AppImage
./LM_Studio-*.AppImage
```
### Enabling AMD GPU Acceleration in LM Studio

LM Studio uses `llama.cpp` under the hood. For AMD GPU support on Linux, it relies on ROCm through the `hipBLAS` backend.

When you load a model in LM Studio:

1. Open **Settings → My Models**
2. Click the gear icon next to a loaded model
3. Under **GPU Acceleration** , select your AMD GPU from the device list
4. Set **GPU Layers** to a high number (e.g., 99) to offload all layers to the GPU

If your AMD GPU doesn’t appear in the device list, the issue is almost always ROCm not being detected. Ensure ROCm is installed, `rocminfo` works, and you’ve added your user to the `render` and `video` groups.

### Using LM Studio as a Local API Server

This is one of LM Studio’s most useful features. It exposes an OpenAI-compatible API endpoint locally:

1. Load a model
2. Navigate to the **Local Server** tab
3. Click **Start Server**

The server runs at `http://localhost:1234/v1` and accepts standard OpenAI API calls. Any tool that supports a custom API base URL — including curl, Python scripts, and most AI clients — can use this as a drop-in local replacement.

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

- **ROCm is AMD’s answer to CUDA** — get it installed correctly and everything else becomes much easier.
- **Linux is the recommended environment** for full AMD GPU acceleration. Windows works for some tools via DirectML, but has more limitations.
- **Ollama is the fastest path** to running LLMs on AMD — one install script, automatic GPU detection, and a clean CLI.
- **LM Studio adds a GUI and local API server** useful for chat interfaces and connecting local models to other tools.
- **ComfyUI on ROCm works well** for image generation — install PyTorch with the ROCm wheel, and most standard workflows run without issues.
- **VRAM is the binding constraint** — plan your model choices around how much you have, not the other way around.
- For those who want to go further than raw inference — building agents, automating workflows, or combining local and cloud models — platforms like MindStudio can connect to local backends and handle the orchestration layer.
