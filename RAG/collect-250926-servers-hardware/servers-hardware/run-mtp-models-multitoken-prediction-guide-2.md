---
id: collect-250926-servers-hardware/servers-hardware/run-mtp-models-multitoken-prediction-guide-2
title: "How to Run MTP Models: Multi-Token Prediction Guide"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Unsloth"]
dates: []
keywords: ["agent", "agents", "gguf", "inference", "llama", "llama.cpp", "speculative decoding"]
source: docs/RAG/clean4/Run MTP Models MultiToken Prediction Guide.md
source_anchor: ""
source_lines: [102, 244]
sha256: 595211308922fdd408018786cd11b4638c76fbca6a55196f24dda261858d8b45
---

# How to Run MTP Models: Multi-Token Prediction Guide

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
