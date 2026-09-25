---
id: collect-250926-servers-hardware/servers-hardware/docker-ollama
title: "docker-ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "gpus", "nvidia"]
source: docs/RAG/clean4/docker-ollama.md
source_anchor: ""
source_lines: [1, 36]
sha256: f250f01f0907acd294fcac8d28cbd6619c1a20d88cc02294392e3800ae0dea22
---

# docker-ollama

## CPU only

## Nvidia GPU

Install the NVIDIA Container Toolkit.
### Install with Apt

1. 
Configure the repository
2. 
Install the NVIDIA Container Toolkit packages

### Install with Yum or Dnf

1. 
Configure the repository
2. 
Install the NVIDIA Container Toolkit packages

### Configure Docker to use Nvidia driver

### Start the container

If you’re running on an NVIDIA JetPack system, Ollama can’t automatically discover the correct JetPack version.
Pass the environment variable 

`JETSON_JETPACK=5` or `JETSON_JETPACK=6` to the container to select version 5 or 6.
## AMD GPU

To run Ollama using Docker with AMD GPUs, use the`rocm` tag and the following command:
## Vulkan Support

Vulkan is bundled into the`ollama/ollama` image and is enabled by default when
the container can access the GPU devices.
`OLLAMA_VULKAN=0` to disable Vulkan, or `GGML_VK_VISIBLE_DEVICES=<ids>` to
select specific Vulkan devices.
