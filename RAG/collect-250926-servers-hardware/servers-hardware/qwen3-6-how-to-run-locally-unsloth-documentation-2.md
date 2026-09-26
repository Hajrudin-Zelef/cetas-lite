---
id: collect-250926-servers-hardware/servers-hardware/qwen3-6-how-to-run-locally-unsloth-documentation-2
title: "qwen3-6-how-to-run-locally-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Hugging Face", "Nvidia", "Perplexity", "SGLang", "Unsloth", "vLLM"]
dates: []
keywords: ["agentic", "benchmarks", "blackwell", "claude", "decode", "distribution", "fp4", "fp8", "gguf", "gpu", "gpus", "inference"]
source: docs/RAG/clean4/qwen3-6-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [202, 478]
sha256: afc251d3e89456ebbcec4bc98431b745c0fa59121488c50d570407bf16148c9b
---

# qwen3-6-how-to-run-locally-unsloth-documentation

We also uploaded dynamic Qwen3.6 4bit and 8bit quants for MacOS devices! Our MLX quant algorithm is still evolving, and we’re actively refining it wherever improvements can be made.

You can run all MLX models in Unsloth Studio!

**Qwen3.6-27B MLX:**

**Qwen3.6-35B-A3B MLX:**

To try them out use:

See below for Qwen3.6-27B KL Divergence (KLD) and Perplexity (PPL) scores (lower is better):

**July 10 2026:** We’re releasing new dynamic NVFP4 Qwen3.6 quants that run ~**2.5× faster** than other NVFP4 quants, with **better performance** and comparable file sizes. Run Qwen3.6-27B NVFP4 **2.5x faster** on **24GB VRAM** and Qwen3.6-35B-A3B **1.7x faster** on **32GB VRAM**. We also added **FP8 KV cache calibration** for 2x longer context lengths! NVFP4 requires NVIDIA's Blackwell GPUs like RTX 50X, DGX Spark (see DGX Spark with NVFP4 quants), B200, B300 GPUs. For older GPUs, our GGUFs work well!

All benchmarks use 1x B200 128 concurrency. Higher concurrency can boost 35B to 17,561 tokens / s. We're also releasing two 35B-A3B NVFP4 versions:

For accuracy benchmarks, we conducted MMLU-Pro, AIME 2025, GPQA for FP8, BF16, NVIDIA's NVFP4 and our NVFP4s - we show our faster quants do similarly on all:

**MTP tensors are also built directly into the quants for additional speedups.** Accuracy gains come from improvements to Qwen3.6’s chat template and dataset calibration. We use our previous chat template updates to help improve coding and tool-calling consistency while reducing looping and other reported issues. Our calibration uses a mix of our dataset optimized for coding, tool-calling and chat alongside UltraChat.

For Decode speed (tokens per person), ours is 1.03x faster for 27B and 1.17x and 1.22x faster for 35B.

NVFP4 runs 4-bit weights and matrix multiplications directly on Blackwell Tensor Cores. Our Qwen3.6 NVFP4 quants use W4A4 so they actually use the FP4 tensor cores, so they decode faster than NVIDIA's which use W4A16. We also dynamically quantize layers to retain accuracy, and we conducted MMLU-Pro, AIME 2025, GPQA for all quants including comparing to FP8 and BF16.

**Qwen3.6-27B NVFP4 Accuracy Benchmarks**

Unsloth

86.25

86.34

93.12

NVIDIA

85.96

86.87

93.12

FP8

86.11

86.87

93.75

BF16

85.96

88.13

93.33

**Qwen3.6-35B-A3B NVFP4 Accuracy Benchmarks**

Unsloth

85.85

86.74

92.29

**Unsloth Fast**

85.58

87.75

91.67

NVIDIA

85.60

87.12

91.88

FP8

85.75

86.74

93.12

BF16

85.75

86.36

92.50

We also checked the output length of all benchmarks, and they are comparable, so the new NVFP4 quants do not think for longer which defeats the purpose of quantizing them! (Ie if it's 2x faster, but thinks 2x more, then that's useless)

We also found Marlin kernels to not support W4A4 well - enabling it will cause a 2.5x performance degradation - so use CUTLASS, Flashinfer-TRTLLM or Cute-DSL (auto enabled in vLLM)! Also if you have a DGX Spark, see Qwen3.6 you must use `--moe-backend flashinfer_b12x` or you will get 2.5x slower inference.

**So don't set any backend - vLLM auto selects the best.**

nvidia 27B

W4A16

marlin (auto)

115.6

2,403

unsloth 27B

W4A4

marlin

105.6

2,127

unsloth 27B

W4A4

cutlass

113.5

6,681

unsloth 27B

W4A4

flashinfer_trtllm

112.6

6,158

unsloth 27B

W4A4

**cute-DSL (auto)**

125.9

**6,863**

nvidia 35B-A3B

W4A4

marlin (auto)

240.8

8,721

unsloth 35B-A3B

W4A4

marlin

215.8

8,619

unsloth 35B-A3B

W4A4

cutlass

158.3

11,017

unsloth 35B-A3B

W4A4

**cute-DSL (auto)**

295.2

**15,636**

To run NVFP4 quants, see below for commands to run Qwen3.6-27B in vLLM and SGLang (you can change model name to `Qwen3.6-35-A3B-NVFP4`). Also do NOT select any MoE backend - leave vLLM to select it - for eg Marlin is 2.5x slower! See Marlin vs Flashinfer vs cutlass vs cute-DSLIf you have a DGX Spark, see Qwen3.6 you must use `--moe-backend flashinfer_b12x` or you will get much slower inference.

To install vLLM in a separate venv:

Then to serve the 35B Fast variant:

Change `unsloth/Qwen3.6-35B-A3B-NVFP4-Fast` to the NVFP4 quant names!

To enable MTP / speculative decoding (faster decode but somewhat less throughput), use:

If you get Torchcodec issues, be sure to do the below then relaunch vllm.

To ensure DGX Spark has the correct kernels (or you will get **2x SLOWER inference**), first check:

which should NOT error out - if it did, please update vllm or reinstall via:

Then to serve in vLLM for DGX Spark:

If you get Torchcodec issues, be sure to do the below then relaunch vllm.

For this guide we will be utilizing Dynamic 4-bit which works great on a 24GB RAM / Mac device for fast inference on llama.cpp. Because the model is only around 72GB at full F16 precision, we won't need to worry much about performance. See our GGUF collection.

Obtain the latest `llama.cpp` **on** **GitHub here**. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

If you want to use `llama.cpp` directly to load models, you can do the below: (:`Q4_K_XL`) is the quantization type. You can also download via Hugging Face (point 3). This is similar to `ollama run` . Use `export LLAMA_CACHE="folder"` to force `llama.cpp` to save to a specific location. The model has a maximum of 256K context length.

Follow one of the commands for the specific models:

**Thinking mode:**

General tasks:

For precise coding tasks, change: `temperature=0.6`

**Non-thinking mode:**

General tasks:

**Thinking mode:**

General tasks:

For precise coding tasks, change: `temperature=0.6`

**Non-thinking mode:**

General tasks:

You can also download the model manually as well via the code below (after installing `pip install huggingface_hub`). You can choose Q4_K_M or other quantized versions like `UD-Q4_K_XL` . We recommend using at least 2-bit dynamic quant `UD-Q2_K_XL` to balance size and accuracy. If downloads get stuck, see: Hugging Face Hub, XET debugging

Then run the model in conversation mode:

Qwen3.6 also has **Preserve Thinking** which leaves the thinking trace from the previous conversation. This increases the number of tokens you use, but could increase accuracy in continued conversations. Unsloth Studio has 'Think' and Preserved Thinking toggles for Qwen3.6:

To enable **preserve thinking** in llama.cpp use (change to 'true' or 'false') '`preserve_thinking`' instead of '`enable_thinking`' or '`disable_thinking`'.

For normal thinking, you can enable / disable thinking in llama.cpp by following the below commands. Use '`true`' and '`false`' interchangeably.

Linux, MacOS, WSL:

Windows / Powershell:

As an example for Qwen3.6-35B-A3B to enable preserve thinking (default is enabled):

And then in Python:

To run the model via local coding agentic workloads, you can follow our guide. Use the `llama-server` we just set up just then, and set the model name to the exact id it reports at `GET /v1/models` (the `--alias` value above, e.g. `unsloth/Qwen3.6-35B-A3B-GGUF`). Follow the correct Qwen3.6 parameters and usage instructions.

After following the instructions for Claude Code for example you will see:

We can then ask say `Create a Python game for Chess` :

We conducted Mean KL Divergence benchmarks for Qwen3.6-35-A3B GGUFs across providers to help you pick the best quant.

- KL Divergence puts nearly all Unsloth GGUFs on the SOTA Pareto frontier
- KLD shows how well a quantized model matches the original BF16 output distribution, indicating retained accuracy.
- This makes Unsloth the top-performing in 21 of 22 sizes
- Only Q6_K was updated for more Dynamic layers and we introduced a new `UD-IQ4_NL_XL` quant

We benchmarked the new quants we made for 27B and 35B MoE. In general, dense models are much more accelerated with MTP (1.4-2x) vs MoE models (1.15-1.25x).

