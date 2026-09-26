---
id: collect-260926-rattrapage/rattrapage/nvidia-nemotron-3-ultra-how-to-run-locally-1
title: "NVIDIA Nemotron 3 Ultra - How To Run Locally"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Hugging Face", "Nvidia", "Unsloth"]
dates: []
keywords: ["nvidia", "agent", "agents", "benchmarks", "gguf", "gpu", "inference", "license", "llama", "llama.cpp", "moe", "nvfp4"]
source: docs/RAG/lot-rattrapage/servers-reviews/NVIDIA Nemotron 3 Ultra - How To Run Locally.md
source_anchor: ""
source_lines: [1, 144]
sha256: e8a21d20d15388be64622e2b8e4a78a4087563378540ee7d991df9ff15852082
---

# NVIDIA Nemotron 3 Ultra - How To Run Locally

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/nemotron-3-ultra.md).
# NVIDIA Nemotron 3 Ultra - How To Run Locally
Run Nemotron-3-Ultra-550B-A55B locally on your device!
NVIDIA Nemotron 3 Ultra is an open **550B parameter, 55B active** frontier-reasoning model and is NVIDIA's **largest model** released so far. Nemotron-3-Ultra-550B-A55B is built for long-running autonomous agents and reasoning across coding, deep research workflows. It is the **strongest Western open model**, and adopts the new Open Model, Weights & Data License.
With up to **1M context**, Nemotron 3 Ultra uses a Hybrid Transformer-Mamba MoE architecture and can preserve long agent state, logs, and plans across sustained sessions. GGUFs are at [Nemotron-3-Ultra-550B-A55B](https://huggingface.co/unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF) with dynamic 1bit taking 189GB of disk space. It's also pretrained using NVFP4. We als did [GGUF KLD Benchmarks](#kld-benchmarks).
### ⚙️ Usage Guide
NVIDIA recommends these settings for inference:
* `temperature = 1.0`
* `top_p = 0.95`
| Detail         | Nemotron 3 Ultra                                                                                             |
| -------------- | ------------------------------------------------------------------------------------------------------------ |
| Model size     | 550B total parameters / 55B active parameters                                                                |
| Context length | Up to 1M tokens                                                                                              |
| Architecture   | Hybrid Transformer-Mamba MoE with Latent MoE, Multi-Token Prediction (MTP currently not supported for GGUFs) |
| Model I/O      | Text input, text output                                                                                      |
The chat template is like below:
{% code overflow="wrap" %}
```
<|im_start|>system\n<|im_end|>\n<|im_start|>user\nWhat is 1+1?<|im_end|>\n<|im_start|>assistant\n2<|im_end|>\n<|im_start|>assistant\n\n
```
{% endcode %}
### Run Nemotron-3-Ultra
The 3-bit versions of the model requires \~256GB RAM, 4-bit needs \~300GB and 8-bit requires 600GB. For these guides, we will be using 3-bit `UD-IQ3_XXS` which fits on a 256GB device and is a good balance between size and accuracy. Depending on your use-case you will need to use [different settings](#usage-guide). **GGUF:** [Nemotron-3-Ultra-550B-A55B](https://huggingface.co/unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF)
Run in Unsloth StudioRun in llama.cpp
### 🦥 Unsloth Studio Guide
For this tutorial, we will be using [Unsloth Studio](/docs/new/studio.md), which is our UI for running and training LLMs. With Unsloth Studio, you can run models and input image and text locally on **Mac, Windows**, and Linux and:
{% columns %}
{% column %}
* Search, download, [run GGUFs](/docs/new/studio.md#run-models-locally) and safetensor models
* **Compare** models **side-by-side**
* [**Self-healing** tool calling](/docs/new/studio.md#execute-code--heal-tool-calling) + **web search**
* [**Code execution**](/docs/new/studio.md#run-models-locally) (Python, Bash)
* [Automatic inference](https://unsloth.ai/docs/desktop#feature-deep-dive) parameter tuning (temp, top-p, etc.)
* [Train LLMs](/docs/new/studio.md#no-code-training) 2x faster with 70% less VRAM
  {% endcolumn %}
{% column %}
{% endcolumn %}
{% endcolumns %}
{% stepper %}
{% step %}
#### Install Unsloth
**MacOS, Linux, WSL:**
```bash
curl -fsSL https://unsloth.ai/install.sh | sh
```
**Windows PowerShell:**
```bash
irm https://unsloth.ai/install.ps1 | iex
```
{% endstep %}
{% step %}
#### Setup Unsloth Studio (one time)
Setup automatically installs Node.js (via nvm), builds the frontend, installs all Python dependencies, and builds llama.cpp with CUDA support.
{% hint style="info" %}
**WSL users:** you will be prompted for your `sudo` password to install build dependencies (`cmake`, `git`, `libcurl4-openssl-dev`).
{% endhint %}
{% endstep %}
{% step %}
#### Launch Unsloth
**MacOS, Linux, WSL:**
```bash
source unsloth_studio/bin/activate
unsloth studio -H 0.0.0.0 -p 8888
```
**Windows Powershell:**
```bash
unsloth studio
```
Then open `http://127.0.0.1:8888` in your browser.
{% endstep %}
{% step %}
#### Search and download Nemotron-3-Ultra
On first launch you will need to create a password to secure your account and sign in again later. Then go to the [Unsloth Chat](/docs/new/studio/chat.md) tab and search for Nemotron-3-Ultra in the search bar and download your desired model and quant.
{% endstep %}
{% step %}
#### Run Nemotron-3-Ultra
Inference parameters should be auto-set when using Unsloth Studio, however you can still change it manually. You can also edit the context length, chat template and other settings.
For more information, you can view our [Unsloth Studio inference guide](/docs/new/studio/chat.md).
{% endstep %}
{% step %}
#### Serving Nemotron-3-Ultra
You can also use `unsloth studio run` to serve the model via llama-server like so:
{% code overflow="wrap" %}
```bash
unsloth studio run --model unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF:UD-Q4_K_XL
```
{% endcode %}
{% endstep %}
{% endstepper %}
### 🦙 Llama.cpp Tutorial:
Instructions to run in llama.cpp (note we will be using 4-bit to fit most devices):
{% stepper %}
{% step %}
Obtain the latest `llama.cpp` on [GitHub here](https://github.com/ggml-org/llama.cpp). You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.
{% code overflow="wrap" %}
```bash
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
    -DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON -DLLAMA_CURL=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-mtmd-cli llama-server llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
{% endcode %}
{% endstep %}
{% step %}
Download the model via the code below (after installing `pip install huggingface_hub`). You can choose Q4\_K\_M or other quantized versions like `UD-Q4_K_XL` . We recommend using at least 2-bit dynamic quant `UD-Q2_K_XL` to balance size and accuracy. If downloads get stuck, see: [Hugging Face Hub, XET debugging](/docs/basics/troubleshooting-and-faqs/hugging-face-hub-xet-debugging.md)
{% code overflow="wrap" %}
```bash
pip install huggingface_hub
hf download unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF \
    --local-dir unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF \
    --include "*UD-IQ3_XXS*" # Use "*UD-Q2_K_XL*" for Dynamic 2bit
```
{% endcode %}
{% endstep %}
{% step %}
Then run the model in conversation mode:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
    --model unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF/UD-IQ3_XXS/NVIDIA-Nemotron-3-Ultra-550B-A55B-UD-IQ3_XXS-00001-of-00006.gguf \
    --temp 1.0 \
    --top-p 0.95 \
    --min-p 0.01
```
{% endcode %}
{% endstep %}
{% endstepper %}
#### Llama-server serving & deployment
To deploy Nemotron-3-Ultra locally, use `llama-server`. In a new terminal, for example via `tmux`, deploy the model:
```bash
./llama.cpp/llama-server \
    -hf unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF:UD-IQ3_XXS \
    --alias "unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B" \
    --temp 1.0 \
    --top-p 0.95 \
    --port 8001
```
If you downloaded the model manually, use:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-server \
