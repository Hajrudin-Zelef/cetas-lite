---
id: collect-250926-servers-hardware/servers-hardware/qwen3-6-how-to-run-locally-unsloth-documentation-1
title: "qwen3-6-how-to-run-locally-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Apple", "Hugging Face", "Unsloth"]
dates: []
keywords: ["qwen", "agentic", "benchmarks", "context window", "fine-tuning", "gguf", "gpu", "gpus", "inference", "llama", "llama.cpp", "memory"]
source: docs/RAG/clean4/qwen3-6-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 201]
sha256: 93706d24f2b7ba1f85deb02dc8a7ff536a3c8bad5ec72dd0768cf1e601502a1b
---

# qwen3-6-how-to-run-locally-unsloth-documentation

Qwen3.6 is Alibaba’s new family of multimodal hybrid-thinking models, including: **Qwen3.6-27B** and **35B-A3B**. It delivers top performance for its size, supports 256K context across 201 languages. It excels in agentic coding, vision, chat tasks. Qwen3.6-27B runs on **18GB RAM** setups and 35B-A3B runs on **22GB**. You can now run and train the models in Unsloth Desktop.

**July 10:** We released new **NVFP4** quants to run Qwen3.6 2.5x faster on GPUs.

**Qwen3.6 MTP is here****!** MTP enables 1.4-2.2x faster inference without accuracy loss. Run MTP directly in Unsloth Studio. We conducted Qwen3.6 GGUF Benchmarks to help you pick the best quant.

Qwen3.6 GGUFs use Unsloth Dynamic 2.0 for SOTA quant performance - so quants are calibrated on real world use-case datasets and important layers are upcasted. *Thank you Qwen for day zero access.*

- **Developer Role Support** for Codex, OpenCode and more:
Our uploads now support the`developer role` for agentic coding tools.

**Table: Inference hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

**27B**

15 GB

18 GB

24 GB

30 GB

55 GB

**35B-A3B**

17 GB

23 GB

30 GB

38 GB

70 GB

For best performance, make sure your total available memory (VRAM + system RAM) exceeds the size of the quantized model file you’re downloading. If it doesn’t, you can still run via SSD/HDD offloading, but inference will be slower.

**To train Qwen3.6, you can refer to our previous** **Qwen3.5 fine-tuning guide****.**

- **Maximum context window:**`262,144` (can be extended to 1M via YaRN)
- `presence_penalty = 0.0 to 2.0` default this is off, but to reduce repetitions, you can use this, however using a higher value may result in**slight decrease in performance**
- **Adequate Output Length** :`32,768` tokens for most queries

As Qwen3.6 is hybrid reasoning, thinking and non-thinking mode have different settings:

temperature = 1.0

temperature = 0.6

top_p = 0.95

top_p = 0.95

top_k = 20

top_k = 20

min_p = 0.0

min_p = 0.0

presence_penalty = 0.0

presence_penalty = 0.0

repeat_penalty = disabled or 1.0

repeat_penalty = disabled or 1.0

Thinking mode for general tasks:

Thinking mode for precise coding tasks:

temperature = 0.7

top_p = 0.8

top_k = 20

min_p = 0.0

presence_penalty = 1.5

repeat_penalty = disabled or 1.0

To disable thinking / reasoning, use `--chat-template-kwargs '{"enable_thinking":false}'`

If you're on **Windows** Powershell, use: `--chat-template-kwargs "{\"enable_thinking\":false}"`

Use 'true' and 'false' interchangeably.

Instruct (non-thinking) for general tasks:

We'll be using Dynamic 4-bit `UD-Q4_K_XL` GGUF variants for inference workloads. Click below to navigate to designated model instructions:

Do NOT use CUDA 13.2 as you may get gibberish outputs. Use below CUDA 13.2 or CUDA 13.3.

Qwen3.6 and Qwen3.6 MTP can now be run in Unsloth Desktop, our new open-source UI for local AI. Unsloth Studio lets you run models locally on **MacOS, Windows**, Linux and:

- Fast CPU + GPU inference via llama.cpp

The easiest way to get started is by downloading the Unsloth Desktop app. Works on macOS, Windows, and Linux.

Or, if you prefer to install manually:

**MacOS, Linux, WSL:**

**Windows PowerShell:**

**Installation will be quick and take approx 20 sec - 1 mins.**

**MacOS, Linux, WSL and Windows:**

Then open `http://127.0.0.1:8888` (or your specific URL) in your browser.

**Launch Unsloth securely with HTTPS and Cloudflare**

**NEW!** Unsloth now provides a secure way to launch Unsloth over HTTPS through a free Cloudflare tunnel. Use the below (works in Windows, Mac & Linux):

On first launch you will need to create a password to secure your account and sign in again later. Then go to the Unsloth Chat tab and search for Qwen3.6 or Qwen3.6 MTP in the search bar and download your desired model and quant.

Inference parameters should be auto-set when using Unsloth Studio, however you can still change it manually. You can also edit the context length, chat template and other settings.

For more information, you can view our Unsloth Studio inference guide. Below, the 2-bit Qwen3.6 GGUF made 30+ tool calls, searched 20 sites and executed Python code:

MTP (Multi Token Prediction) speculative decoding enables models like Qwen3.6 to have **~1.4-2.2x faster generation with** **no change in accuracy**. This enables Qwen3.6 27B and 35B-A3B to have **>1.4x speed-up** over the original baseline which is especially useful for local models.

Unsloth Qwen3.6 MTP GGUFs are no longer in experimental mode, and llama.cpp has merged MTP support. Run directly in Unsloth Studio’s UI or via llama.cpp. **Qwen3.6 27B MTP now runs at 160 tokens/s generation and Qwen3.6 35B-A3B at 240 tokens/s on a RTX 6000 GPU.** See MTP Benchmarks.

Unsloth Studio automatically sets the ideal MTP settings optimized for your specific hardware (Mac, CPU, GPU etc.) - you can still change it later.

In practice, MTP predicts several future tokens, then the main model verifies those tokens in parallel. This reduces the number of forward passes needed during generation and make output faster. **We found** **--spec-draft-n-max 2** **to work best in most setups.** **However, do not assume** **2** **is optimal, as performance is hardware-dependent. Try values from** **1** **through** **6** **and use whichever is fastest for your system.**

We also uploaded MTP GGUFs for the **Qwen3.5** **model family** including: 0.8B, 2B, 4B, 9B, 27B, 35B-A3B, 122B-A10B and 397B-A17B. Llama.cpp is continually improving MTP performance, so expect it to get faster overtime!

**Table: MTP hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

**27B**

16 GB

19 GB

25 GB

31 GB

56 GB

**35B-A3B**

18 GB

24 GB

31 GB

39 GB

71 GB

Unsloth Studio automatically sets the ideal MTP settings optimized for your specific hardware (Mac, CPU, GPU etc.) - you can still change it later.

On first launch you will need to create a password to secure your account and sign in again later. Then go to the Unsloth Chat tab and search for Qwen3.6 MTP in the search bar and download your desired model and quant.

Inference parameters should be auto-set when using Unsloth Studio, however you can still change it manually. You can also edit the context length, chat template and other settings.

For more information, you can view our Unsloth Studio inference guide. Below, the 2-bit Qwen3.6 MTP GGUF made 10+ tool calls, searched 10 sites and executed Python code:

Install the latest version of `llama.cpp` on **GitHub here**. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

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

