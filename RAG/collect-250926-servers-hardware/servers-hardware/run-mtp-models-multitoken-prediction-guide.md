---
id: collect-250926-servers-hardware/servers-hardware/run-mtp-models-multitoken-prediction-guide
title: "How to Run MTP Models: Multi-Token Prediction Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Google", "Hugging Face", "Unsloth"]
dates: []
keywords: ["agent", "agents", "gguf", "gpu", "gpus", "inference", "llama", "llama.cpp", "memory", "qwen", "speculative decoding", "throughput"]
source: docs/RAG/clean4/Run MTP Models MultiToken Prediction Guide.md
source_anchor: ""
source_lines: [1, 244]
sha256: 74295ae128361307035d1981692052a07d8e1811b50a1012eb30d635241d73fd
---

# How to Run MTP Models: Multi-Token Prediction Guide

> For the complete documentation index, see [llms.txt](https://unsloth.ai/docs/llms.txt). Markdown versions of documentation pages are available by appending `.md` to page URLs; this page is available as [Markdown](https://unsloth.ai/docs/models/mtp.md).
# How to Run MTP Models: Multi-Token Prediction Guide
MTP, or Multi-Token Prediction, speeds up inference by letting a model predict multiple upcoming tokens at once instead of generating one token per step. It enables faster inference without accuracy loss and is especially effective on GPUs. In this guide, you’ll learn how to use MTP models like [Gemma 4](/docs/models/gemma-4.md) or [Qwen3.8](/docs/models/qwen3.6.md) on your local device.
MTP predicts multiple future tokens, which the main model verifies in parallel. This reduces generation forward passes, speeding output while preserving quality because only verified tokens are kept.
When running [GGUFs](/docs/basics/dynamic-3.0-ggufs.md), MTP can make generation **\~1.4× to 2.2× faster**. Dense models like Gemma-4-31B benefit most, reaching **>1.4× speedup** over the original. Gains are smaller on devices with lower memory bandwidth, such as older Macs. You can run MTP models directly in [Unsloth Studio’s UI](/docs/new/studio.md) or llama.cpp.
{% hint style="info" %}
**MTP uses more memory than standard**, so plan for \~2 GB additional RAM/VRAM headroom.
{% endhint %}
Gemma 4 MTPQwen3.6 MTP
We found `--spec-draft-n-max 2` is the best starting point however, **do not assume `2` is optimal**, as performance is hardware-dependent. Try any value from `1` through `6` and use whichever is fastest for your system. Unsloth Studio automatically sets the ideal MTP settings optimized for your specific hardware (Mac, CPU, GPU etc.) - you can still change it later.
### Gemma 4 MTP
Google DeepMind trained MTP separately from the original [Gemma 4](/docs/models/gemma-4/qat.md) models, including for [QAT variants](/docs/models/gemma-4/qat.md). Unlike Qwen, Google released specific MTP variants under the `assistant` name. For best results, we only upload 3 precision options: **8-bit** and **16-bit** (BF16, F16). For QAT - we applied the [smart 4-bit recovery process](/docs/models/gemma-4/qat.md#qat-analysis) like we did for Gemma 4 QAT quants, and so the MTP quants are also smart 4-bit derived.
We uploaded `mtp-` prefixed GGUFs to each repo, so you only need to use the **regular original Gemma 4 GGUFs**, no separate repo is needed. You can access Gemma [MTP models here](https://huggingface.co/collections/unsloth/gemma-4) and they can now run in [Unsloth](#unsloth-studio-mtp-guide). We benchmarked Gemma 4 QAT with MTP, and it runs 1.5x - 2.2x faster:

🦥 Run in Unsloth Studio🦙 Run in llama.cpp
so the below just works (this uses the 8-bit one)
{% code overflow="wrap" %}
```bash
llama-server \
-hf unsloth/gemma-4-31B-it-GGUF \
--spec-type draft-mtp \
--spec-draft-n-max 4
```
{% endcode %}
### Qwen3.6 MTP
Qwen directly trained MTP inside of the [Qwen3.6](/docs/models/qwen3.6.md) and [Qwen3.5](/docs/models/qwen3.5.md) models. This enables Qwen3.6 27B MTP to reach 160 tokens/s and Qwen3.6 35B-A3B reach 240 tokens/s on an RTX 6000 GPU. GGUF uploads:
| [Qwen3.6-27B-MTP-GGUF](https://huggingface.co/unsloth/Qwen3.6-27B-MTP-GGUF) | [Qwen3.6-35B-A3B-MTP-GGUF](https://huggingface.co/unsloth/Qwen3.6-35B-A3B-MTP-GGUF) |
| --------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
**Table: MTP hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

Qwen3.6

3-bit

4-bit

6-bit

8-bit

BF16

27B

16 GB

19 GB

25 GB

31 GB

56 GB

35B-A3B

18 GB

24 GB

31 GB

39 GB

71 GB

Below are graphs of inference throughput for MTP vs. no MTP:

We also [uploaded MTP GGUFs](https://huggingface.co/unsloth/models?search=mtp) for the [Qwen3.5](/docs/models/qwen3.5.md) **model family** including: 0.8B, 2B, 4B, 9B, 27B, 35B-A3B, 122B-A10B and 397B-A17B. Llama.cpp is continually improving MTP performance, so expect it to get faster overtime!
To run the Qwen MTP models, follow the steps either for [Unsloth Studio](#unsloth-studio-mtp-guide) or [llama.cpp](#llama.cpp-mtp-guide).
### 🦥 Unsloth Studio MTP Guide
Unsloth Studio automatically sets the ideal MTP settings optimized for your specific hardware (Mac, CPU, GPU etc.) - you can still change it later.
{% stepper %}
{% step %}
#### Install Unsloth
Run in your terminal:
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
#### Launch Unsloth
**MacOS, Linux, WSL and Windows:**
```bash
unsloth studio -H 127.0.0.1 -p 8888
```
Then open `http://127.0.0.1:8888` (or your specific URL) in your browser.
{% endstep %}
{% step %}
#### Search and download your desired model
On first launch you will need to create a password to secure your account and sign in again later. Then go to the [Unsloth Chat](/docs/new/studio/chat.md) tab and search for Qwen3.6 MTP or Gemma 4 in the search bar and download your desired model and quant.
{% hint style="warning" %}
**Gemma 4 MTP is automatically enabled in Unsloth. You only need to download the regular original Gemma 4 GGUF.** We updated the Gemma 4 GGUF files to include an additional MTP file inside a separate folder within the GGUF package, so there is no need to download a separate Gemma 4 assistant GGUF.
The only model that still requires a separate MTP GGUF is Qwen3.6.
{% endhint %}

{% endstep %}
{% step %}
#### Run your MTP model
Inference, MTP and speculative **decoding settings** should be auto-set when using Unsloth Studio, however you can still change it manually. You can also edit speculative decoding, the context length, chat template and other settings in the right side bar.
Gemma 4Qwen3.6
#### Gemma 4 MTP:
Don't forget to **change the model name** to your desired Gemma 4 model size like Gemma-4-26B-A4B etc. as the instructions below are for Gemma-4-12B. Notice we provided a `mtp-` prefixed GGUF, so the below `-hf` command should auto download and use MTP.
**Thinking mode:**
```bash
export LLAMA_CACHE="unsloth/gemma-4-12b-it-GGUF"
./llama.cpp/llama-cli \
-hf unsloth/gemma-4-12b-it-GGUF:UD-Q4_K_XL \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--spec-type draft-mtp --spec-draft-n-max 2
```
{% hint style="info" %}
Please see Gemma 4's new [Preserved Thinking](#thinking-enable-disable--preserve-thinking).
{% endhint %}
**Non-thinking mode**:
```bash
export LLAMA_CACHE="unsloth/gemma-4-12b-it-GGUF"
./llama.cpp/llama-cli \
-hf unsloth/gemma-4-12b-it-GGUF:UD-Q4_K_XL \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--spec-type draft-mtp --spec-draft-n-max 2 \
--chat-template-kwargs '{"enable_thinking":false}'
```
#### Qwen3.6 MTP:
Don't forget to **change the model name** to your desired Qwen3.6 variant like Qwen3.6-35B-A3B or Qwen3.5 etc. as the instructions below are for Qwen3.6-27B:
**Thinking mode** (General tasks)**:**
```bash
export LLAMA_CACHE="unsloth/Qwen3.6-27B-MTP-GGUF"
./llama.cpp/llama-cli \
-hf unsloth/Qwen3.6-27B-MTP-GGUF:UD-Q4_K_XL \
--temp 1.0 \
--top-p 0.95 \
--top-k 20 \
--min-p 0.00 \
--spec-type draft-mtp --spec-draft-n-max 2
```
For precise coding tasks, change: `temperature=0.6`
{% hint style="info" %}
Please see Qwen3.6's new [Preserved Thinking](#thinking-enable-disable--preserve-thinking).
{% endhint %}
**Non-thinking mode** (General tasks):
```bash
export LLAMA_CACHE="unsloth/Qwen3.6-27B-MTP-GGUF"
./llama.cpp/llama-server \
-hf unsloth/Qwen3.6-27B-MTP-GGUF:UD-Q4_K_XL \
--temp 0.7 \
--top-p 0.8 \
--top-k 20 \
--presence-penalty 1.5 \
--min-p 0.00 \
--spec-type draft-mtp --spec-draft-n-max 2 \
--chat-template-kwargs '{"enable_thinking":false}'
```
{% endstep %}
{% step %}
#### Manually downloading quants
If you want to manually download the quants and the MTP quants, you can also do that! Download the model via the code below (after installing `pip install huggingface_hub hf_transfer`). You can choose Q4\_K\_M or other quantized versions like `UD-Q4_K_XL` . We recommend using at least 2-bit dynamic quant `UD-Q2_K_XL` to balance size and accuracy. If downloads get stuck, see: [Hugging Face Hub, XET debugging](/docs/basics/troubleshooting-and-faqs/hugging-face-hub-xet-debugging.md)
#### Gemma 4 MTP:
```bash
hf download unsloth/gemma-4-12B-it-qat-GGUF \
--local-dir unsloth/gemma-4-12B-it-qat-GGUF \
--include "*mmproj-F16*" \
--include "mtp-*" \
--include "*UD-Q4_K_XL*" # Use "*UD-Q2_K_XL*" for Dynamic 2bit
```
#### Qwen3.6 MTP:
```bash
hf download unsloth/Qwen3.6-27B-MTP-GGUF \
--local-dir unsloth/Qwen3.6-27B-MTP-GGUF \
--include "*mmproj-F16*" \
--include "*UD-Q4_K_XL*" # Use "*UD-Q2_K_XL*" for Dynamic 2bit
```
{% endstep %}
{% step %}
Then run the model in conversation mode:
#### Gemma 4 MTP:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
--model unsloth/gemma-4-12B-it-qat-GGUF/gemma-4-12B-it-qat-UD-Q4_K_XL.gguf \
--mmproj unsloth/gemma-4-12B-it-qat-GGUF/mmproj-F16.gguf \
--model-draft unsloth/gemma-4-12B-it-qat-GGUF/mtp-gemma-4-12B-it.gguf \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--spec-type draft-mtp --spec-draft-n-max 2
```
{% endcode %}
And you will see the below - ignore the error messages as well
#### Qwen3.6 MTP:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
--model unsloth/Qwen3.6-27B-MTP-GGUF/Qwen3.6-27B-UD-Q4_K_XL.gguf \
--mmproj unsloth/Qwen3.6-27B-MTP-GGUF/mmproj-F16.gguf \
--temp 1.0 \
--top-p 0.95 \
--min-p 0.00 \
--top-k 20 \
--spec-type draft-mtp --spec-draft-n-max 2
```
{% endcode %}
{% endstep %}
{% step %}
#### Llama-server deployment
To deploy Gemma-4 on llama-server, use:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-server \
--model unsloth/gemma-4-12B-it-qat-GGUF/gemma-4-12B-it-qat-UD-Q4_K_XL.gguf \
--mmproj unsloth/gemma-4-12B-it-qat-GGUF/mmproj-F16.gguf \
--model-draft unsloth/gemma-4-12B-it-qat-GGUF/mtp-gemma-4-12B-it.gguf \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--alias "unsloth/gemma-4-12b-it-qat-GGUF" \
--port 8001 \
--chat-template-kwargs '{"enable_thinking":true}'
```
{% endcode %}
{% endstep %}
{% endstepper %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/models/mtp.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
