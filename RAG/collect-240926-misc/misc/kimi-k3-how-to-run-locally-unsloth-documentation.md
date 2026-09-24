---
id: collect-240926-misc/misc/kimi-k3-how-to-run-locally-unsloth-documentation
title: "kimi-k3-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Anthropic", "Apple", "Moonshot", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["kimi", "agentic", "benchmarks", "claude", "compute", "context window", "gguf", "gpt-5.6", "gpu", "inference", "llama", "llama.cpp"]
source: docs/RAG/clean_en/misc/kimi-k3-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 305]
sha256: 6294863aa7a03ac868395a2c5836cc141d4e929a69307b83d2319f2b5384566a
---

# kimi-k3-how-to-run-locally-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/kimi-k3 -->

Kimi K3 by Moonshot AI is a 2.8T parameter open-weight model (104B active) built for SOTA coding, agentic, long-context, and chat workloads. It is the **strongest open model** to date, rivaling Claude 4.8 Opus and GPT-5.6. Kimi K3 has native vision, a 1M-token context window and uses MXFP4. Full-precision inference requires 1.56 TB of storage and 1-bit Kimi K3 Unsloth Dynamic GGUF requires **594 GB (62% less)**.

Dynamic 1-bit (see right) reaches **~78.9%** top-1 accuracy while being **62% smaller**. Dynamic 2-bit 861.3GB reaches **~90%** accuracy while being **45% smaller**. Run **Kimi-K3-GGUF** via Unsloth Studio or llama.cpp. Kimi K3 can run on a NVIDIA DGX Station, or Mac Studio connected to a 128GB RAM device. 

For **lossless** Kimi K3, use Q8 (`UD-Q8_K_XL`), which is **50GB larger** than Q4 (`UD-Q4_K_XL`). We are still investigating if we can push it under 512GiB (dynamic 1-bit is 553.2 GiB) without damaging the model.

**Table: Hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

610 GB

665 GB

726 GB

880 GB

1.6 TB

We built on top of llama.cpp PR with our fork which includes vision support and some bug fixes.

1. The mmproj / vision tower is similar to the Kimi-K2.5 tower, but with RMSNorm, no biases, a non-square fused QKV (qkv width != n_embd) and a post-norm projector.
2. We found when running llama.cpp, the `n_tokens * 40` budget failed at large batch sizes, and so we had to up to`n_tokens * 160`
3. We also had to convert the Kimi chat template to a jinja format.
4. We tested as much as we could to cover all cases. Kimi by default has been trained with **preserved thinking on** , so all thinking traces are not deleted, but kept.

Like Kimi K2.6 and K2.7, K3's `UD-Q8_K_XL` is lossless because Kimi uses MXFP4 for MoE weights and BF16 for everything else, and `Q8_K_XL` follows that exactly. `UD-Q4_K_XL` is similar except some of remaining tensors (except norms etc) are `Q8_0`, so it is near full precision and requires 1.56 TB RAM/VRAM. `UD-Q8_K_XL` is 'truly lossless' vs the MXFP4 full safetensors version.

`UD-IQ1_S`

594.0

0.5645

2.5789

78.875 +/- 0.107

36.495

`UD-IQ1_M`

648.9

0.4789

2.3639

81.219 +/- 0.103

33.629

`UD-IQ2_XXS`

711.1

0.3784

2.1266

84.127 +/- 0.096

29.826

`UD-Q2_K_XL`

861.3

0.1779

1.7359

90.390 +/- 0.077

19.862

`UD-Q4_K_XL`

1,510

1.4579

`UD-Q8_K_XL`

1,560

1.4581

For imatrix generation and quantization, we used the 1.56 TB lossless `UD-Q8_K_XL` throughout calibration; its perplexity is 1.4581. Our Dynamic-1bit quant reaches 2.58 perplexity with 79% top-1 accuracy, making it surprisingly usable.

Other community quants are larger yet degrade far more. For example, one 618.9 GB quant `IQ1_M` exceeds our 594 GB 1-bit quant, but its perplexity jumps to 54.56 - 21× worse. The same pattern holds for `IQ2_XXS`: 725 GB at 96 PPL vs. our 711 GB at 2.12 PPL - 45× worse, meaning their 2-bit performs even worse than their 1-bit. This highlights importance of dynamic quantization + proper calibration.

**We also provide Top-1% Accuracy, KLD plots:**

Kimi K3 is **thinking-only**, with **preserve_thinking** **always enabled** and **max** thinking on by default. Instant mode is not supported. Thinking effort is configured with the `reasoning_effort` request field, and K3 supports `"low"`, `"high"`, and `"max"` thinking efforts.

temperature = 1.0

temperature = 1.0

top_p = 0.95

top_p = 1.0

- Context length = up to `1,048,576`
- Low, High, Max Thinking is toggagle in Unsloth

If the model fits, you will get ~20 tokens/s generation when using B200s and >120 tokens / s throughput. We recommend `UD-IQ1_S` (594GB) as a good size/quality balance. Best rule of thumb: RAM+VRAM ≈ the quant size; otherwise it’ll still work, just much slower due to disk offloading.

You can now run Kimi K3 in llama.cpp and Unsloth Desktop. We will be utilizing the 594GB `UD-IQ1_S` quant for best results in terms of accessibility and accuracy and it will require at least 610GB RAM. Feel free to change quantization type. GGUF: **Kimi-K3-GGUF**

Kimi K3 can run in Unsloth Desktop, an open-source desktop UI for local AI. **Unsloth Desktop automatically offloads to RAM and detects multiGPU setups**. With Unsloth Studio, you can run models locally on **MacOS, Windows**, Linux and:

- Fast CPU + GPU inference via llama.cpp

**Install and Launch Unsloth**

The easiest way to get started is by downloading the Unsloth Desktop app. Works on macOS, Windows, and Linux.

Or, if you prefer to install manually:

MacOS, Linux, WSL:

Windows PowerShell:

**Launch Unsloth**

MacOS, Linux, WSL and Windows:

Then open `http://127.0.0.1:8888` (or your specific URL) in your browser.

**Launch Unsloth securely with HTTPS and Cloudflare**

**NEW!** Unsloth now provides a secure way to launch Unsloth over HTTPS through a free Cloudflare tunnel. Use the below (works in Windows, Mac & Linux):

**Search and download Kimi K3**

Unsloth Studio automatically offloads to RAM and detects multiGPU setups. On first launch you will need to create a password to secure your account and sign in again later.

Then go to the Model hub tab and search for **Kimi K3** in the search bar and download your desired model and quant. Ensure you have enough compute the run the model.

**Run Kimi K3**

Inference parameters should be auto-set when using Unsloth Studio, however you can still change it manually. You can also toggle **low, high or max thinking**, edit the context length, chat template and other settings.

For more information, you can view our Unsloth Studio inference guide.

For these tutorials, we will use llama.cpp for fast local inference, especially if you have a CPU. We created a fork specifically to support Kimi K3 vision thus. This builds on top of another llama.cpp PR.

Obtain the SPECIFIC Unsloth fork of `llama.cpp` on **GitHub here** to enable vision support. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

You can now use `llama.cpp` directly to load and download models, just like `ollama run`. First, select the quantization type you want like `IQ1_S`. Also use `export LLAMA_CACHE="folder"` to force `llama.cpp` to save to a specific location. **Note this download process might be very slow**, so it's probably best to use the manual download process in the next section.

Then run the model in conversation mode:

You will then see:

And I asked "What is the sqrt of -1":

Kimi K3 also supports images for eg loading the Unsloth image:

And then we use the sloth image and ask how it's related:

You can view further below for benchmarks in table format:

**Reasoning & Knowledge**

GPQA Diamond

93.5

92.6

**94.1**

91.0

93.5

91.2

HLE-Full

43.5 / 56.0

**53.3 / 63.0**

44.5 / 58.0

49.8 / 57.9

41.4 / 52.2

—

**Coding**

DeepSWE

67.5

70.0

**73.0**

59.0

67.0

46.2

Terminal-Bench 2.1

88.3

88.0

**88.8**

84.6

83.4

82.7

**Agentic**

BrowseComp

**91.2**

88.0

90.4

84.3

84.4

—

GDPval-AA v2 (Elo)

1686

**1747**

1736

1593

1491

1510

OSWorld 2.0

58.3

**66.1**

62.6

55.7

49.5

—

**Vision**

MMMU-Pro

81.6 / 83.4

81.2 / **86.5**

**83.0** / 84.6

78.9 / 82.7

81.2 / 83.2

—

MathVision

94.3 / 97.8

94.8 / **98.6**

**95.8** / 97.8

86.7 / 97.1

92.2 / 96.8

—

DeepSWE Benchmarks show Kimi-K3 doing very efficiently!

Last updated

Was this helpful?
