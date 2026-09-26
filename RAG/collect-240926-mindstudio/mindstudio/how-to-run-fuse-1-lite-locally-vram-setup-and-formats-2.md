---
id: collect-240926-mindstudio/mindstudio/how-to-run-fuse-1-lite-locally-vram-setup-and-formats-2
title: "how-to-run-fuse-1-lite-locally-vram-setup-and-formats"
domain: mindstudio
role: reference
task: reference
actors: ["Apple", "vLLM"]
dates: []
keywords: ["agents", "alignment", "attention", "gguf", "gpu", "llama", "llama.cpp", "memory", "quantization", "throughput", "vllm"]
source: docs/RAG/clean_en/mindstudio/how-to-run-fuse-1-lite-locally-vram-setup-and-formats.md
source_anchor: ""
source_lines: [101, 133]
sha256: b2d52a6fb4ee639a8d2109cb476b84d279e85473768090348f2086a1e9cbf966
---

# how-to-run-fuse-1-lite-locally-vram-setup-and-formats

```
pip install git+https://huggingface.co/Akahsizrr/fuse-1-Lite-vLLM
vllm serve Akahsizrr/fuse-1-Lite \
  --mamba-cache-mode align \
  --max-model-len 4096
```
The `--mamba-cache-mode align` flag matters here because LFM2's short-convolution layers need cache alignment that differs from standard transformer attention caching. This path needs roughly 12 GB of VRAM and is best suited to A10G, A100, or H100 class hardware, given vLLM's throughput-oriented design.

## Frequently Asked Questions

### Can you run fuse-1 Lite on a laptop GPU or Apple Silicon?

Yes. The 4-bit bitsandbytes version needs only 3.36 GB, which fits an RTX 3060 or an Apple M2 Pro. For Mac users specifically, the MLX build runs natively on M1 Pro and later chips, though the unquantized MLX version needs around 12 GB of unified memory.

### Do you need a special version of llama.cpp to run the GGUF file?

Yes. fuse-1 Lite's GGUF uses a custom `fuse3` architecture that stock llama.cpp cannot load. You need to build a fork with Fuse3 support, using the C++ graph builder and integration guide provided in the GGUF repository.

### What's the difference between the 4-bit and 8-bit versions in practice?

4-bit NF4 quantization cuts VRAM to 3.36 GB but introduces more precision loss, while 8-bit quantization uses 6.00 GB and stays closer to the original bfloat16 weights. Both are applied at runtime through `BitsAndBytesConfig` rather than downloaded as separately pre-quantized checkpoints.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

### Can you turn off the coding experts and use it as a general chat model?

Yes. The model exposes a `set_coding_enabled(False)` method that disables the transplanted expert layers, effectively reverting to the underlying LFM2.5-2.6B host model's behavior. Calling `set_coding_enabled(True)` re-enables the coding experts.

### Why does fuse-1 Lite need `trust_remote_code=True`?

Because it isn't a standard transformers architecture. It uses a custom `Fuse3ForCausalLM` class that implements the expert augmentation logic, router, and scaling mechanism, none of which exist in the base transformers library.
