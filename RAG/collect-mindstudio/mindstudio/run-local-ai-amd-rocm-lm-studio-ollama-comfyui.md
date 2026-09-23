---
id: collect-mindstudio/mindstudio/run-local-ai-amd-rocm-lm-studio-ollama-comfyui
title: "How to Run Local AI on AMD: ROCm, LM Studio, Ollama, and ComfyUI Setup"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["amd", "llama", "compute", "consumer", "cost", "embeddings", "fp8", "gpu", "gpus", "inference", "llama.cpp", "mistral"]
source: docs/RAG/Collect RAG/02_mindstudio/run-local-ai-amd-rocm-lm-studio-ollama-comfyui.md
source_anchor: ""
source_lines: [1, 63]
sha256: 8d1ef8a86818cf945972ee3203a49ddda0968d39092a4244f749c9b89a40a005
---

# How to Run Local AI on AMD: ROCm, LM Studio, Ollama, and ComfyUI Setup

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-local-ai-amd-rocm-lm-studio-ollama-comfyui
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This practical guide explains how to run a full local AI stack on AMD hardware — ROCm installation, Ollama for LLMs, LM Studio as a desktop chat interface, and ComfyUI for image generation — targeting RX 6000/RX 7000 and supported professional cards.

ROCm (Radeon Open Compute) is AMD's open-source GPU compute platform, the equivalent of NVIDIA's CUDA. It supports PyTorch (official builds), TensorFlow (rocTF), HIP (CUDA-compatible programming model), and MIOpen (deep learning primitives). Without ROCm, AI tools default to slow CPU-only mode. Supported GPUs: RX 7900 XTX/XT/GRE (top consumer), RX 7800 XT/7700 XT, RX 6900 XT/6800 XT/6800 (RDNA2), RX 6700 XT (caveats), Radeon Pro and Instinct. Unsupported cards (RX 6600/6500 XT) may work via HSA_OVERRIDE_GFX_VERSION but with varying results. ROCm is primarily Linux technology; on Windows, DirectML is the alternative (LM Studio supports it; Ollama/ComfyUI need WSL2).

Setup: add user to render/video groups, install via AMD's amdgpu-install script (repo.radeon.com, usecase=rocm or graphics,rocm), verify with `rocminfo` and `rocm-smi` (ROCm equivalent of nvidia-smi).

Ollama on AMD: `curl -fsSL https://ollama.com/install.sh | sh` auto-detects the GPU with ROCm installed. Pull 7B models (8GB VRAM Q4) like llama3.2, or 13B (12GB+ VRAM) like mistral-nemo. VRAM sizing table: 8GB→7B Q4; 12GB→7B Q8 or 13B Q4; 16GB→13B Q8 or 30B Q4; 24GB→70B Q4 or 30B Q8. Q4_K_M is the best quality/speed/VRAM balance. Confirm GPU usage by watching `rocm-smi` during inference.

LM Studio on AMD: uses llama.cpp with ROCm via the hipBLAS backend. Enable GPU acceleration in Settings → My Models, set GPU Layers high (e.g., 99). Its local API server exposes OpenAI-compatible endpoint at http://localhost:1234/v1 (drop-in for curl/Python/other clients). On Windows, select Vulkan or DirectML backend.

ComfyUI on AMD: clone the repo, create a venv, install PyTorch built against ROCm (`pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/rocm6.1`), verify with `torch.cuda.is_available()`. Launch at http://127.0.0.1:8188. Model dirs: checkpoints/, vae/, loras/, controlnet/. Use `--lowvram` for SDXL on 8GB, `--novram` for CPU offloading. Performance: 10–30% slower than equivalent NVIDIA; SDXL 1024×1024 (20 steps, Euler a): RX 7900 XTX ~4–6s, RX 6800 XT ~7–10s, RX 7800 XT ~8–12s. Flux needs 10GB+ VRAM (FP8).

Troubleshooting: missing group permissions (render/video), gfx architecture override (HSA_OVERRIDE_GFX_VERSION=10.3.0 for RX 6600), OOM (lower quantization or --lowvram/--novram), CPU fallback (verify GPU utilization), and ROCm version mismatches between system libraries and PyTorch builds.

FAQ: Ollama has official ROCm support on Linux since 2024; Windows works via DirectML with limitations; RX 7900 XTX (24GB) is the best consumer AMD option; NVIDIA is typically 15–30% faster at equivalent price points but the gap has narrowed for LLM inference; 8GB VRAM is the practical minimum for useful LLM inference; ComfyUI is reasonably stable on ROCm (CUDA-specific custom nodes may fail).

## Key points

- ROCm is AMD's CUDA equivalent — correct installation unlocks GPU acceleration for everything else.
- Linux is the recommended environment for full AMD acceleration; Windows uses DirectML with more limitations.
- Ollama is the fastest path to LLMs on AMD (one install script, automatic GPU detection).
- LM Studio adds a GUI and an OpenAI-compatible local API server (localhost:1234/v1).
- ComfyUI works on ROCm with the ROCm PyTorch wheel; 10–30% slower than equivalent NVIDIA cards.
- VRAM is the binding constraint — size models around it.

## Technical data / figures

| VRAM | Recommended model size |
|---|---|
| 8GB | 7B (Q4) |
| 12GB | 7B (Q8) or 13B (Q4) |
| 16GB | 13B (Q8) or 30B (Q4) |
| 24GB | 70B (Q4) or 30B (Q8) |

| SDXL 1024×1024 (20 steps, Euler a) | Time |
|---|---|
| RX 7900 XTX | ~4–6 s |
| RX 6800 XT | ~7–10 s |
| RX 7800 XT | ~8–12 s |

| Key endpoints | Value |
|---|---|
| Ollama server | localhost:11434 |
| LM Studio API | localhost:1234/v1 |
| ComfyUI | 127.0.0.1:8188 |
| ROCm PyTorch index | https://download.pytorch.org/whl/rocm6.1 |

## Why this source matters for the RAG

Provides a complete, actionable setup path for running local LLMs, embeddings, and image generation on AMD hardware — relevant for building cost-effective local RAG stacks outside the NVIDIA ecosystem. Includes concrete commands, VRAM sizing rules, and troubleshooting that apply directly to on-prem deployment.
