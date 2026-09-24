---
id: collect-240926-misc/misc/qwen3-8-flash-next-how-to-run-locally-unsloth-documentation
title: "qwen3-8-flash-next-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Nvidia", "Unsloth"]
dates: []
keywords: ["qwen", "benchmarks", "claude", "context window", "cost", "embeddings", "gguf", "gpu", "gpus", "inference", "llama", "llama.cpp"]
source: docs/RAG/clean_en/misc/qwen3-8-flash-next-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 266]
sha256: b8327678d1bddebea774267d8ae0c56acd3dab5c24ee1fb5556f149531c56c73
---

# qwen3-8-flash-next-how-to-run-locally-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/qwen3.8-next -->

Qwen3.8-Flash-Next is a new open-weight, **125B parameter** MoE multimodal model from Qwen. Built on the new Qwen4 architecture, it supports a 262K context window and advanced reasoning. Qwen3.8-Flash-Next outperforms Claude-4.6-Opus (Max) and can run locally on devices with **75GB RAM**/unified memory with no GPU VRAM required. To run the model, use our GGUFs via llama.cpp or Unsloth Desktop. Thank you Qwen for day zero access.

Whether you run **Qwen3.8-Flash-Next** on a CPU with system RAM or on a GPU with VRAM may make relatively little difference. Its unique architecture allows inference using RAM or unified memory to achieve performance closer to that of GPU VRAM than is typical for other models. This makes it particularly well suited to Macs, NVIDIA DGX Spark systems, and other devices with large memory capacities.

You will need at least **75 GB of RAM or unified memory** to run the model. Its smallest 1-bit quantized version is larger than usual because of new Ngram layers or per layer embeddings which is like a lookup table. However, this also means the quantization is less aggressive, allowing the model to retain more of its original accuracy than more heavily quantized models. You can also offload the PLE / Ngram layer to SSD and use mmap which allows less usage of CPU and GPU VRAM.

The smallest quant works on 75GB RAM so it's best to have a 96GB RAM/unified memory device.
**Table: Hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

75 GB

79 GB

90 GB

96-114 GB

163 GB

200 GB

355 GB

Qwen3.8-Flash-Next is a **hybrid thinking** model with different default settings for thinking and non-thinking modes. Extra high is enabled by default so if you want shorter thinking traces, you can adjust the thinking effort:

`temperature`

1.0

0.7

`top_p`

0.95

0.80

`top_k`

20

20

`min_p`

0.0

0.0

`presence_penalty`

0.0

1.5

`repetition_penalty`

1.0

1.0

- Context length = up to `262,144`
- Thinking Mode: `temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
- Instruct (or non-thinking) mode: `temperature=0.7` ,`top_p=0.80` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`

Qwen3.8-Flash-Next has **Preserve Thinking** which leaves the thinking trace from the previous conversation. This increases the number of tokens you use, but could increase accuracy in continued conversations. Unsloth has 'Think' and Preserved Thinking toggles for Qwen3.8 (see right):

Qwen3.8-Flash-Next comes with support for `reasoning_effort`, which can be used to adjust reasoning depth and control cost. These toggles are automatically enabled in Unsloth:

- `xhigh` (default): for complex tasks demanding thorough analysis
- `medium` : balancing accuracy and speed
- `low` : efficient reasoning optimizing for speed and cost
- none

To change thinking / reasoning effort in `unsloth run` or `llama-server`, use `--chat-template-kwargs '{"reasoning_effort":"medium"}'`

If you're on **Windows** Powershell, use: `--chat-template-kwargs "{\"reasoning_effort\":\"medium\"}"`

Change `medium` to your desired reasoning level.

We ran KLD for Qwen3.8-Flash quants, and show that 80% top-1 accuracy recovery is possible with 79% less disk space usage. The new architecture uses PLE / Ngrams, and these are not quantized that heavily (4-bit minimum) since they have random access pattern, and quantizing them heavily will damage the model.

UD-IQ1_S

72.5

77.325

0.396070

7.2126

UD-IQ1_M

74.5

79.691

0.314739

6.1965

UD-Q2_K_XL

78.9

82.715

0.224607

4.9121

UD-IQ3_XXS

82.0

85.414

0.165120

4.0375

UD-Q3_K_XL

90.0

88.315

0.106504

3.0538

UD-IQ4_XS

93.7

89.554

0.083630

2.3677

UD-Q4_K_XL

111.3

92.255

0.046893

1.5468

UD-Q5_K_XL

158.3

93.680

0.030415

1.0036

UD-Q6_K_XL

169.2

94.089

0.027091

0.8416

Q8_0

188.2

94.122

0.026574

0.8118

You can now run Qwen3.8-Flash-Next in Unsloth Desktop and llama.cpp. Feel free to change quantization type.

Qwen3.8-Flash-Next now is able to run in Unsloth Desktop, an open-source UI app for local AI. **Unsloth automatically offloads to RAM and detects multiGPU setups**. With Unsloth Desktop, you can run models locally on **MacOS, Windows**, Linux and:

- Fast CPU + GPU inference via MLX and llama.cpp

MTP is automatically enabled bit you can disable it. Inference parameters should be auto-set when using Unsloth, however you can still change it manually. You can also edit the context length, chat template and other settings.

For more information, you can view our Unsloth inference guide.

For example using Unsloth Desktop with the 397GB Qwen3.8 (-91% smaller) allows you to toggle thinking modes, allow inline canvas, web search and code execution and much more.

You can use `unsloth run` command and serve Qwen3.8 via an API using `llama-server` runtime flags, including context sizing, GPU layers, threading, sampling, networking, and tool configuration. For more info see our API docs or unsloth start.

Install the latest version of llama.cpp. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

To run the model, you can do:

Then to run it:

Qwen3.8-Flash can run with 1.3 to **1.7x faster inference** via MTP (Multi-Token Prediction) with no accuracy degradation! MTP enables Qwen3.8-Flash to reach **170 tokens/s** on 1x RTX 6000 PRO GPU compared to the 100 token baseline. MTP speeds up inference by letting a model predict multiple upcoming tokens at once instead of generating one token per step, and is especially effective on GPUs.

To run Qwen3.8-Flash with MTP, MTP is enabled by default in Unsloth Desktop or you can use our custom llama.cpp PR.

Gains are smaller on devices with lower memory bandwidth, such as older Macs. We created both shared MTP modules (excludes the embed_tokens and shares it with the main model) to save disk space, RAM and VRAM usage by around 1 to 2GB.

BF16

7.77 GB

5.23 GB

2.54 GB

Q8_0

4.14 GB

2.79 GB

1.35 GB

Q4_K_M

2.79 GB

1.91 GB

880 MB

The 3-bit MTP quant works on 91GB RAM so it's best to have a 96GB RAM/unified memory device.
**Table: MTP Hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

76 GB

80 GB

91 GB

97-115 GB

164 GB

200 GB

355 GB

To run Qwen3.8-Flash with MTP, all you need to do is **install Unsloth Desktop** or update to the latest verison of Unsloth then re-download the model or download the MTP file. See further below for llama.cpp instructions.

Unsloth Desktop Works on macOS, Windows, and Linux.

In Unsloth Desktop, you can also change the number of draft tokens or customize MTP. Use the advanced settings in the right sidebar, and enable "Advanced settings" and you can select MTP / Ngram speculative decoding, the number of draft tokens and more:

Then to download the shared MTP module:

And to use llama-server with it:

For GGUF quantization benchmarks you can see above for our quantization analysis or Dynamic V3.0 article.

Last updated

Was this helpful?
