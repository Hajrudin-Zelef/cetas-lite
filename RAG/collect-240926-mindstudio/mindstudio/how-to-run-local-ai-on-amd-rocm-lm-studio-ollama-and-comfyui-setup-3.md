---
id: collect-240926-mindstudio/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup-3
title: "Download the amdgpu-install package (check AMD's site for the latest version)"
domain: mindstudio
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "gpu", "agents", "inference"]
source: docs/RAG/clean_en/mindstudio/how-to-run-local-ai-on-amd-rocm-lm-studio-ollama-and-comfyui-setup.md
source_anchor: ""
source_lines: [349, 355]
sha256: 153f3689873fc2bc5631e1a36b9b122513940a36c8cbb45af3387704f1524b5f
---

# Download the amdgpu-install package (check AMD's site for the latest version)

- **ROCm is AMD’s answer to CUDA** — get it installed correctly and everything else becomes much easier.
- **Linux is the recommended environment** for full AMD GPU acceleration. Windows works for some tools via DirectML, but has more limitations.
- **Ollama is the fastest path** to running LLMs on AMD — one install script, automatic GPU detection, and a clean CLI.
- **LM Studio adds a GUI and local API server** useful for chat interfaces and connecting local models to other tools.
- **ComfyUI on ROCm works well** for image generation — install PyTorch with the ROCm wheel, and most standard workflows run without issues.
- **VRAM is the binding constraint** — plan your model choices around how much you have, not the other way around.
- For those who want to go further than raw inference — building agents, automating workflows, or combining local and cloud models — platforms like MindStudio can connect to local backends and handle the orchestration layer.
