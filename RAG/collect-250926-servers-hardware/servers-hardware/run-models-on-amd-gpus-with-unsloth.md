---
id: collect-250926-servers-hardware/servers-hardware/run-models-on-amd-gpus-with-unsloth
title: "Train & run models on AMD GPUs with Unsloth"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Z.ai", "vLLM"]
dates: []
keywords: ["amd", "gpu", "gpus", "agent", "agents", "benchmarks", "claude", "consumer", "datacenter", "deepseek", "fine-tuning", "gguf"]
source: docs/RAG/clean4/run models on AMD GPUs with Unsloth.md
source_anchor: ""
source_lines: [1, 282]
sha256: ef8a1a12735aa4b4817f8aa7f560f9d0a7fb9eb3d1886b3428d378f2cd8a8f96
---

# Train & run models on AMD GPUs with Unsloth

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/basics/amd.md).
# Train & run models on AMD GPUs with Unsloth
Unsloth now officially supports AMD hardware, making it easy to run, train, fine-tune, RL, and deploy models locally on AMD GPUs. Fully open source, [Unsloth](/docs/desktop.md) runs on Windows, WSL, and Linux, supporting RDNA 1–4, Radeon RX 9000 and 7000 series GPUs, Instinct MI350 and MI300 data center GPUs, Vulkan, Strix Halo - powered Ryzen AI Max systems, and more.
We collaborated with [AMD](https://www.amd.com/en/developer/resources/technical-articles/2026/train-and-run-models-on-amd-gpus-with-unsloth.html) and our community to enable up to 2x faster and 70% less VRAM training for all models with no accuracy loss. If you don't have a GPU, Unsloth still supports native AMD inference for [Qwen3.6](/docs/models/qwen3.6.md), [Gemma 4](/docs/models/gemma-4.md), DeepSeek-V4, GLM 5.2, DiffusionGemma, Kimi K2.7 and other models.
QuickstartFeaturesGithub
{% hint style="success" %}
**Sept 18 Update:** Added **RDNA 1, 2, Gorgon Halo, Vulkan support** + fixed not detecting GPUs on AMD GPUs.\
Better RDNA4, HIP / ROCm failure auto fixing and catching.
{% endhint %}
To install Unsloth on AMD, the easiest way is to download our [Desktop app](/docs/desktop.md).
Download for WindowsDownload for Linux\
\
For manual installation:
**MacOS, Linux, WSL:**
{% code overflow="wrap" expandable="true" %}
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
{% endcode %}
**Windows Powershell:**

irm https://unsloth.ai/install.ps1 | iex

### ⭐ Features
Unsloth supports both inference and training for AMD, so if you don't have a GPU and only CPU, you can still run models with Unsloth. When using Unsloth, you get:
* [**Tool call healing**](/docs/new/studio/chat.md#auto-healing-tool-calling) **for inference to increase accuracy by 50%**, unlimited free [web search](/docs/new/studio/chat.md#advanced-web-search), HTML canvases, [**secure HTTPS deployment**](#unsloth-on-phones-and-remote-https-tunneling) via Cloudflare and [Code Execution](/docs/new/studio.md#execute-code--heal-tool-calling) all work.
* **Train Gemma 4 models in 8GB VRAM** or Qwen3.5 in 3GB VRAM. Triton kernels, math algorithms and memory tricks allow 70% less VRAM use for AMD with no accuracy loss.
* Up to **80% less VRAM usage for** [**reinforcement learning**](/docs/get-started/reinforcement-learning-rl-guide/grpo-long-context.md) such as GRPO. And weight sharing with vLLM to save 50% memory usage.
* Optimized support for **RDNA 3** and newer GPUs & data-center grade GPUs. Strix Halo is also optimized for Linux, WSL and Windows.
{% columns %}
{% column %}
* [**Connect**](/docs/integrations/unsloth-start.md) **your local models to any agent harness**: [Claude Code](/docs/basics/claude-code.md), [Codex](/docs/basics/codex.md), [Hermes](/docs/integrations/hermes-agent.md), [OpenClaw](/docs/integrations/openclaw.md) and more.
* Support inference and training for nearly **all models**, including the latest [DeepSeek-V4](/docs/models/deepseek-v4.md), [GLM-5.2](/docs/models/glm-5.2.md), [Kimi-K2.7](/docs/models/kimi-k2.7-code.md), [MiniMax M3](/docs/models/minimax-m3.md), [DiffusionGemma](/docs/models/diffusiongemma.md), [Inkling](/docs/models/inkling.md) and more!
* AMD multi-GPU support + VRAM/RAM tracking. `llama.cpp` ROCm prebuilts are provided daily, so you always get the latest model updates!
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
### ⚡Quickstart
Unsloth auto installs the best ROCm builds for PyTorch, llama.cpp prebuilts, bitsandbytes, ROCm optimized kernels & Triton. We also ship our AMD optimizations and fixes together. You can read our more [detailed guide here](#guide-to-using-unsloth-on-amd) for more details.
We made Unsloth installations seamless for Windows and WSL folks - but we generally recommend using Linux directly since it'll have the broadest support. Unsloth supports MacOS, Linux, [Windows](/docs/get-started/install/windows-installation.md), [NVIDIA](/docs/get-started/install/pip-install.md), Intel and CPU setups. See: [Unsloth Requirements](/docs/get-started/fine-tuning-for-beginners/unsloth-requirements.md). Updating are the same commands!
**MacOS, Linux, WSL:**
{% code overflow="wrap" %}
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
{% endcode %}
**Windows PowerShell:**
{% code overflow="wrap" %}
```bash
irm https://unsloth.ai/install.ps1 | iex
```
{% endcode %}
{% columns %}
{% column %}
Unsloth will auto install ROCm, PyTorch, custom llama.cpp prebuilts and more!
#### Accessing Unsloth remotely
We also enabled [Cloudflare HTTPS tunneling](#unsloth-on-phones-and-remote-https-tunneling) for free - so you can access Unsloth securely via HTTPS remotely.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### :inbox\_tray: Install Unsloth Notebooks & pip install
If you want plain Unsloth notebooks, see the general [AMD](/docs/get-started/install/amd.md) install guide. Be sure to read[AMD](/docs/get-started/install/amd.md) install guide! After you follow it for installing PyTorch and ROCm, then do:
{% code overflow="wrap" %}
```bash
uv pip install unsloth[amd]
```
{% endcode %}
#### :play\_pause:Unsloth Start
{% columns %}
{% column width="58.333333333333336%" %}
[Unsloth Start](/docs/integrations/unsloth-start.md) lets you connect [Claude Code](/docs/basics/claude-code.md), [Codex](/docs/basics/codex.md) and other agents to local models via the `unsloth start` command.
Start Unsloth, open your project folder, and run:
{% code overflow="wrap" %}
```bash
unsloth start claude --model \
unsloth/gemma-4-E2B-it-GGUF:UD-Q4_K_XL
```
{% endcode %}
Replace `claude` with any agent below:
{% endcolumn %}
{% column width="41.666666666666664%" %}
{% endcolumn %}
{% endcolumns %}
| Agent | Command |
| ------------------------------------------------------------------ | ------------------------ |
| :claude: Claude Code | `unsloth start claude` |
| :openai: OpenAI Codex | `unsloth start codex` |
| :caduceus: Hermes Agent | `unsloth start hermes` |
| :lobster: OpenClaw | `unsloth start openclaw` |
| :rectangle-vertical: OpenCode | `unsloth start opencode` |

### :bar\_chart: AMD Analysis
We brought many of our optimizations we made for non AMD GPUs over to make AMD GPUs shine! Getting AMD training to work on Windows devices and ensuring compatbility across nearly every AMD device was a lot of work but we're happy in the current state. We are still actively working on making AMD GPUs even better! Below are some analysis/benchmarks we conducted on a AMD MI300X.
#### Reinforcement learning: more accurate LoRA and QLoRA
During RL on AMD, Unsloth supports **weight sharing** with vLLM as introduced in [Memory Efficient RL](/docs/get-started/reinforcement-learning-rl-guide/memory-efficient-rl.md). We **halve memory usage** by directly accessing vLLM's allocation of the model.
During LoRA and QLoRA, we also do not do merging and demerging of the LoRA weights like below:
$$
W = W + sAB \\
\text{inference} = XW \\
W = W - sAB
$$
The above causes issues as **IEEE floats are not associative due to rounding**. When W is in BF16, this causes subtle differences when subtracting as done in other RL engines. This means in IEEE land:
$$
W \ne(W + sAB) - (sAB)
$$
And when W is in bfloat16 with few mantissa bits, it **rounds small numbers to zero**. Since A and B are float32 during LoRA, this means W can drift over time after merging & demerging. Unsloth uses vLLM's LoRA path directly, so we avoid this. The base weights are never edited and will never drift.
#### Faster Reinforcement Learning
{% columns %}
{% column %}
We benchmarked other setups with FA2 and vLLM co-located. Both use vLLM for generation, Unsloth pushes more of the step into generation while making training and roughly 2x faster.
All three benchmarks are reproducible: same seeds, same token budgets, and the
loss curves match between the two stacks.
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
#### Faster & more memory efficient training
On Llama-3.1-8B LoRA SFT (batch 2 x grad-accum 4 x 2048 = 16,384 tokens/step, packed), Unsloth trains at 2.07 s/step vs 2.87 s/step for TRL + FA2, and peaks at 18.3 GB versus 24.3 GB. That is 1.39x faster and 1.33x less memory usage, with no change in accuracy. We plan to make this even better!

The memory savings are not just a lower peak, they hold for the whole run.
Sampling allocated VRAM every 0.1s, Unsloth stays flat and low the entire time,
while other setups + FA2 spikes toward 22.8 GB on nearly every step and takes longer to
finish (75s vs 56s for the same 25 steps).
### :sparkles:Supported AMD Hardware
Unsloth’s AMD support focuses on popular consumer, workstation, and datacenter GPUs. Training, Unsloth Studio, and GGUF/llama.cpp inference support can vary by architecture, so the table below separates optimized support from compatibility paths.

Architecture

Series

gfx

Support

Platform

RDNA 4

Radeon™ RX 9000 Series

gfx1200, gfx1201

Full

Windows + WSL + Linux

RDNA 3.5

Ryzen AI 300 / Ryzen AI MAX (Strix Halo)

gfx1150, gfx1151, gfx1152

Full

Windows + WSL + Linux

RDNA 3

Radeon™ RX 7000 Series

gfx1100, gfx1101, gfx1102

Full

Windows + WSL + Linux

RDNA 2

Radeon™ RX 6000 Series

gfx1030, gfx1031, gfx1032, gfx1034

Full

Windows + WSL + Linux

RDNA 1

Radeon™ RX 5000 Series

gfx1010, gfx1012

Vulkan inference

Windows + WSL + Linux

CDNA 4

Instinct™ MI350 Series GPUs

gfx950

Full

Linux only

CDNA 3

Instinct™ MI300 Series GPUs

gfx940, gfx941, gfx942

Full

Linux only

CDNA 2

Instinct™ MI200 Series GPUs

gfx90a

Full

Linux only

CDNA 1

Instinct™ MI100 Series GPUs

gfx908

Full

Linux only

We are actively working on supporting more AMD GPUs, however older ones unfortunately do not have the required hardware support for us to work on.
### :desktop: Guide to using Unsloth on AMD
After installation, you can open Unsloth Studio in your browser. From there, you can run and fine-tune local LLMs on AMD hardware automatically. See below for detailed install instructions:
{% columns %}
{% column width="50%" %} AMD Install GuideInstall Unsloth Studio
Select your model, dataset, and training settings to begin. For example, you can run QLoRA 4-bit fine-tuning on Gemma 4 12B with a context length of \~16k, LoRA rank 16, learning rate 0.0002, and 30 max steps.
{% endcolumn %}
{% column width="50%" %}
{% endcolumn %}
{% endcolumns %}
{% columns %}
{% column %}
Unsloth Studio’s fine-tuning screen is fully customizable, letting you configure your model, dataset, and training parameters. During training, Unsloth tracks progress in real time, including loss, RAM and VRAM usage, as well as support for multi-gpu. Once training starts you will see the progress plotted in real time.
After training, you can visit your history to see past and active training sessions. Select from any of your fine-tuned models to view past training logs:
{% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}

To set it up, first in terminal after installing Unsloth, type `unsloth studio --secure` . Setup a password so no one unauthorized can access the HTTPS link!
Then in the logs - use the **green Cloudflare link** and enter this on our Phone or any remote device!
And you're in! You can always monitor / cancel the live HTTPS served link by CTRL+C ing the terminal where Unsloth Studio was launched!

### :question:How AMD support was built for Unsloth
Unsloth’s custom kernels were ported to AMD in close collaboration with the AMD team. This required changes across the installer, hardware detection, runtime routing, monitoring, and prebuilt asset selection.
* **ROCm kernels:** HIP/Triton ports tuned for RDNA and CDNA, with architecture-specific precision, performance, and stability fixes. ([#2520](https://github.com/unslothai/unsloth/pull/2520))
* **4-bit training:** bitsandbytes-based QLoRA support for Radeon and CDNA GPUs. ([#3748](https://github.com/unslothai/unsloth/pull/3748))
* **Automatic setup:** robust AMD GPU detection and installation of the correct ROCm PyTorch wheel, including repair and platform safeguards. ([#4770](https://github.com/unslothai/unsloth/pull/4770))
* **Windows and Strix Halo:** native Windows ROCm installation, runtime compatibility patches, unified-memory reporting, and WSL support. ([#5301](https://github.com/unslothai/unsloth/pull/5301), [#6227](https://github.com/unslothai/unsloth/pull/6227))
* **ROCm inference:** architecture-specific llama.cpp prebuilts, source-build fallback, AMD VRAM detection, and monitoring. ([#5172](https://github.com/unslothai/unsloth/pull/5172))
### :clock1: Future work
Thanks so much to the AMD team for collaborating with us on making AMD work well! Upcoming priorities include:
* Continued development and engagement with the AMD team.
* Expanded hardware-backend CI/CD using AMD Strix Halo and Radeon Lab hardware.
* Support for new AMD GPU cards and architectures as they launch.
* Clearer multi-GPU guidance: AMD distributed training is currently Linux-only. On Linux, ROCm uses RCCL, so distributed training works. On Windows, AMD ROCm does not yet provide a GPU-collective backend, with GLOO/CPU-only support as of [TheRock #5694](https://github.com/ROCm/TheRock/pull/5694), use Linux if you plan to do AMD multi-GPU training.
* Further speed optimizations with respect to the ROCm stack.
### :love\_letter: Thanks to AMD for collaborating!
Thanks so much to the AMD team and our amazing Unsloth community for collaborating with us on this release. We thank Strahinja Stamenkovic, Filip Jankovic, Iswarya Alex, Yue Yuan Bill He, Eda Zhou, Mahdi Ghodsi, Guruprasad and the entire AMD team for collaborating on making AMD support work great with Unsloth! Also huge thanks to community members: Nate, UBER6, AiwendilOfMirk..., Gene, Jimster480, VELICAN, Vintheboy, archmaker, BichonFriseMax, Calandracas, Chigoma333, Edd, electroglyph, jslowbell, mk 27B UD-, Satyam, snapcast3r, TotoMC for testing and debugging with us.
An easy way to use Unsloth on AMD is through [AMD Developer Cloud](https://www.amd.com/en/developer/resources/cloud-access/amd-developer-cloud.html), with one-click notebooks powered by MI300X GPUs with 192GB VRAM. AMD also offers free credits through the [AMD AI Developer Program](https://www.amd.com/en/developer/ai-dev-program.html). To run an Unsloth notebook, use any AMD notebook from [unslothai/notebooks](https://github.com/unslothai/notebooks) and swap the GitHub domain with the AMD Developer Cloud URL, since the notebook paths are the same.
Need help or want to share feedback? You can join our [Discord](https://discord.com/invite/unsloth), read the [AMD docs](https://unsloth.ai/docs/get-started/install/amd), open an issue or PR on [GitHub](https://github.com/unslothai/unsloth/issues), or post on our [Reddit r/unsloth](https://www.reddit.com/r/unsloth/).
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/amd.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
