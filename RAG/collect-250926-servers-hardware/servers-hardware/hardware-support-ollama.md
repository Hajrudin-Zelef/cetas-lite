---
id: collect-250926-servers-hardware/servers-hardware/hardware-support-ollama
title: "hardware-support-ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Apple", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "compute", "gpu", "gpus", "intel", "nvidia"]
source: docs/RAG/clean4/hardware-support-ollama.md
source_anchor: ""
source_lines: [1, 78]
sha256: cd1b3544e298c7acd53c9d9f0864e3632cd500c63d6405a0c125ab0e68cd54c8
---

# hardware-support-ollama

## Nvidia

Ollama supports Nvidia GPUs with compute capability 5.0+ and driver version 550 and newer. Nvidia GPUs with compute capability 5.0 through 6.2 require driver version 570 or newer. Check your compute compatibility to see if your card is supported: https://developer.nvidia.com/cuda-gpus
For building locally to support older GPUs, see development.

### GPU Selection

If you have multiple NVIDIA GPUs in your system and want to limit Ollama to use a subset, you can set`CUDA_VISIBLE_DEVICES` to a comma separated list of GPUs.
Numeric IDs may be used, however ordering may vary, so UUIDs are more reliable.
You can discover the UUID of your GPUs by running `nvidia-smi -L` If you want to
ignore the GPUs and force CPU usage, use an invalid GPU ID (e.g., “-1”)
### Linux Suspend Resume

On linux, after a suspend/resume cycle, sometimes Ollama will fail to discover your NVIDIA GPU, and fallback to running on the CPU. You can workaround this driver bug by reloading the NVIDIA UVM driver with`sudo rmmod nvidia_uvm && sudo modprobe nvidia_uvm`
## AMD Radeon

Ollama supports the following AMD GPUs via the ROCm library:
**NOTE:**
Additional AMD GPU support is provided by the Vulkan Library - see below.

### Linux Support

Ollama requires the AMD ROCm v7 driver on Linux. You can install or upgrade using the`amdgpu-install` utility from
AMD’s ROCm documentation.
### Windows Support

Ollama requires an AMD ROCm v7 / HIP7-capable driver stack on Windows.
### Overrides on Linux

Ollama leverages the AMD ROCm library, which does not support all AMD GPUs. In some cases you can force the system to try to use a similar LLVM target that is close. For example The Radeon RX 5400 is`gfx1034` (also known as 10.3.4)
however, ROCm does not currently support this target. The closest support is
`gfx1030`. You can use the environment variable `HSA_OVERRIDE_GFX_VERSION` with
`x.y.z` syntax. So for example, to force the system to run on the RX 5400, you
would set `HSA_OVERRIDE_GFX_VERSION="10.3.0"` as an environment variable for the
server. If you have an unsupported AMD GPU you can experiment using the list of
supported types below.
If you have multiple GPUs with different GFX versions, append the numeric device
number to the environment variable to set them individually. For example,
`HSA_OVERRIDE_GFX_VERSION_0=10.3.0` and `HSA_OVERRIDE_GFX_VERSION_1=11.0.0`
At this time, the known supported GPU types on linux are the following LLVM Targets.
This table shows some example GPUs that map to these LLVM targets:
Reach out on Discord or file an
issue for additional help.

### GPU Selection

If you have multiple AMD GPUs in your system and want to limit Ollama to use a subset, you can set`ROCR_VISIBLE_DEVICES` to a comma separated list of GPUs.
You can see the list of devices with `rocminfo`. If you want to ignore the GPUs
and force CPU usage, use an invalid GPU ID (e.g., “-1”). When available, use the
`Uuid` to uniquely identify the device instead of numeric value.
### Container Permission

In some Linux distributions, SELinux can prevent containers from accessing the AMD GPU devices. On the host system you can run`sudo setsebool container_use_devices=1` to allow containers to use devices.
## Metal (Apple GPUs)

Ollama supports GPU acceleration on Apple devices via the Metal API.
## Vulkan GPU Support

Additional GPU support on Windows and Linux is provided via Vulkan. Vulkan is enabled by default when the backend is installed. On Windows most GPU vendors drivers come bundled with Vulkan support and require no additional setup steps. Most Linux distributions require installing additional components, and you may have multiple options for Vulkan drivers between Mesa and GPU Vendor specific packages
- Linux Intel GPU Instructions - https://dgpu-docs.intel.com/driver/client/overview.html
- Linux AMD GPU Instructions - https://amdgpu-install.readthedocs.io/en/latest/install-script.html#specifying-a-vulkan-implementation

`ollama` user to the `render` group.
The Ollama scheduler leverages available VRAM data reported by the GPU libraries to
make optimal scheduling decisions.  Vulkan requires additional capabilities or
running as root to expose this available VRAM data.  If neither root access or this
capability are granted, Ollama will use approximate sizes of the models
to make best effort scheduling decisions.
### GPU Selection

To select specific Vulkan GPU(s), you can set the environment variable`GGML_VK_VISIBLE_DEVICES` to one or more numeric IDs on the Ollama server as
described in the FAQ. If you
encounter any problems with Vulkan based GPUs, you can disable all Vulkan GPUs
by setting `OLLAMA_VULKAN=0` or `GGML_VK_VISIBLE_DEVICES=-1`.
On mixed iGPU/dGPU systems where the Vulkan iGPU is unstable, keep Vulkan
enabled and set `GGML_VK_VISIBLE_DEVICES` to the discrete GPU index. For
example, use `GGML_VK_VISIBLE_DEVICES=1` when `Vulkan1` is the discrete
GPU.
