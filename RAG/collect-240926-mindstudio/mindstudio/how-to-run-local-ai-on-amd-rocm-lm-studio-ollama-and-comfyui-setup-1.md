---
id: collect-240926-mindstudio/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup-1
title: "Download the amdgpu-install package (check AMD's site for the latest version)"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "Nvidia", "OpenAI"]
dates: []
keywords: ["amd", "gpu", "agents", "attention", "compute", "consumer", "gpus", "inference", "llama", "llama.cpp", "mistral", "nvidia"]
source: docs/RAG/clean_en/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup.md
source_anchor: ""
source_lines: [1, 185]
sha256: 1f58bd6197d902728d0435fcf2c5ad2650dc71b40f208fd8959b902ea373aae5
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

