---
id: collect-260926-rattrapage/rattrapage/run-models-on-amd-gpus-with-unsloth-1
title: "Train & run models on AMD GPUs with Unsloth"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Z.ai", "vLLM"]
dates: []
keywords: ["amd", "gpu", "gpus", "agent", "agents", "claude", "deepseek", "fine-tuning", "gguf", "glm", "inference", "intel"]
source: docs/RAG/lot-rattrapage/fine-tuning/run models on AMD GPUs with Unsloth.md
source_anchor: ""
source_lines: [1, 93]
sha256: a71fca87b420bcf9eb6b57c5bc50bc0488c15547713608cbb27bcb86b50e402a
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

