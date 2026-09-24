---
id: collect-240926-misc/misc/llama-4-how-to-run-fine-tune-unsloth-documentation
title: "llama-4-how-to-run-fine-tune-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Apple", "Hugging Face", "Meta", "Microsoft", "Unsloth"]
dates: []
keywords: ["llama", "attention", "gguf", "gpu", "gpus", "inference", "llama.cpp", "memory", "moe", "parameters", "quantization", "scout"]
source: docs/RAG/clean_en/misc/llama-4-how-to-run-fine-tune-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 47]
sha256: 7ba71d30c5587242b798d3e9374d6f1366eb75f0e91fb799c9115955e3ff13c3
---

# llama-4-how-to-run-fine-tune-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/tutorials/llama-4-how-to-run-and-fine-tune#tutorial-how-to-run-llama-4-scout-in-llama.cpp -->

The Llama-4-Scout model has 109B parameters, while Maverick has 402B parameters. The full unquantized version requires 113GB of disk space whilst the 1.78-bit version uses 33.8GB (-75% reduction in size). **Maverick** (402Bs) went from 422GB to just 122GB (-70%).

Both text AND **vision** is now supported! Plus multiple improvements to tool calling.

Scout 1.78-bit fits in a 24GB VRAM GPU for fast inference at ~20 tokens/sec. Maverick 1.78-bit fits in 2x48GB VRAM GPUs for fast inference at ~40 tokens/sec.

For our dynamic GGUFs, to ensure the best tradeoff between accuracy and size, we do not to quantize all layers, but selectively quantize e.g. the MoE layers to lower bit, and leave attention and other layers in 4 or 6bit.

**Scout - Unsloth Dynamic GGUFs with optimal configs:**

**Maverick - Unsloth Dynamic GGUFs with optimal configs:**

According to Meta, these are the recommended settings for inference:

- **Temperature of 0.6**
- Min_P of 0.01 (optional, but 0.01 works well, llama.cpp default is 0.1)
- Top_P of 0.9
- Chat template/prompt format:

- A BOS token of `<|begin_of_text|>` is auto added during tokenization (do NOT add it manually!)

1. Obtain the latest `llama.cpp` on GitHub here. You can follow the build instructions below as well. Change`-DGGML_CUDA=ON` to`-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference.**For Apple Mac / Metal devices** , set`-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

1. Download the model via (after installing `pip install huggingface_hub hf_transfer` ). You can choose Q4_K_M, or other quantized versions (like BF16 full precision). More versions at: https://huggingface.co/unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF

1. Run the model and try any prompt.
2. Edit `--threads 32` for the number of CPU threads,`--ctx-size 16384` for context length (Llama 4 supports 10M context length!),`--n-gpu-layers 99` for GPU offloading on how many layers. Try adjusting it if your GPU goes out of memory. Also remove it if you have CPU only inference.

Use `-ot ".ffn_.*_exps.=CPU"` to offload all MoE layers to the CPU! This effectively allows you to fit all non MoE layers on 1 GPU, improving generation speeds. You can customize the regex expression to fit more layers if you have more GPU capacity.

For Llama 4 Maverick - it's best to have 2 RTX 4090s (2 x 24GB)

During quantization of Llama 4 Maverick (the large model), we found the 1st, 3rd and 45th MoE layers could not be calibrated correctly. Maverick uses interleaving MoE layers for every odd layer, so Dense->MoE->Dense and so on.

We tried adding more uncommon languages to our calibration dataset, and tried using more tokens (1 million) vs Scout's 250K tokens for calibration, but we still found issues. We decided to leave these MoE layers as 3bit and 4bit.

For Llama 4 Scout, we found we should not quantize the vision layers, and leave the MoE router and some other layers as unquantized - we upload these to https://huggingface.co/unsloth/Llama-4-Scout-17B-16E-Instruct-unsloth-dynamic-bnb-4bit

We also had to convert `torch.nn.Parameter` to `torch.nn.Linear` for the MoE layers to allow 4bit quantization to occur. This also means we had to rewrite and patch over the generic Hugging Face implementation. We upload our quantized versions to https://huggingface.co/unsloth/Llama-4-Scout-17B-16E-Instruct-unsloth-bnb-4bit and https://huggingface.co/unsloth/Llama-4-Scout-17B-16E-Instruct-unsloth-bnb-8bit for 8bit.

Llama 4 also now uses chunked attention - it's essentially sliding window attention, but slightly more efficient by not attending to previous tokens over the 8192 boundary.

Last updated

Was this helpful?
