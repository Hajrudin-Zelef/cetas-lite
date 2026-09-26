---
id: collect-240926-mindstudio/mindstudio/splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan-2
title: "splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu", "moe", "gpus", "llama", "llama.cpp", "memory", "parameters", "quantization", "throughput"]
source: docs/RAG/clean_en/mindstudio/splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan.md
source_anchor: ""
source_lines: [59, 75]
sha256: 21a39ed5d585ac585b6fa76dff0ea1873f7cd7676cba655f394de37fcd4e2066
---

# splitting-a-122b-moe-model-across-an-nvidia-and-amd-gpu-with-vulkan

Splitting layers across multiple GPUs works for dense and mixture-of-experts models alike in llama.cpp. MoE models aren’t required for cross-GPU splitting, but they make very large parameter counts more practical to run quickly once split, since only a fraction of parameters activate per token.

### What is an Oculink port and why does it matter here?

Oculink is an external PCIe connection running at roughly 63GB/s. It functions similarly to the internal slot a desktop graphics card uses, but it’s exposed outside the case, which lets a mini PC or small-form-factor system connect to a full external GPU (an eGPU) that it otherwise couldn’t fit or power internally.

### Does splitting a model across two GPUs reduce quality or accuracy?

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

No. Splitting only changes where the model’s layers physically execute, not the math being performed. Output quality depends on the model itself and its quantization level, not on how its layers are distributed across hardware.

### Is this setup practical for everyday use, or just a proof of concept?

It sits closer to an advanced workaround than a mainstream setup. It requires specific hardware (an APU with an Oculink port, a compatible eGPU, manual BIOS memory allocation) and yields lower throughput than running a model that fits on one GPU. It’s most useful for people trying to run models that exceed any single GPU’s memory without buying additional matching hardware.
