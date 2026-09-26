---
id: collect-250926-servers-hardware/servers-hardware/unsloth-dynamic-3-0-ggufs-3
title: "Unsloth Dynamic 3.0 GGUFs"
domain: servers-hardware
role: reference
task: reference
actors: ["Google", "Microsoft", "Unsloth", "vLLM"]
dates: []
keywords: ["gguf", "agent", "agents", "benchmarks", "gpu", "inference", "llama", "llama.cpp", "quantization", "scout", "vllm"]
source: docs/RAG/clean4/Unsloth Dynamic 3.0 GGUFs.md
source_anchor: ""
source_lines: [132, 243]
sha256: ccd7d61aa15876087f8d0a2d536ec8470015d9f8c032f69d5fe731adb14d2eb4
---

# Unsloth Dynamic 3.0 GGUFs

We designed a new **Efficiency metric** which calculates the usefulness of the model whilst also taking into account its disk size and MMLU 5 shot score:
$$
\text{Efficiency} = \frac{\text{MMLU 5 shot score} - 25}{\text{Disk Space GB}}
$$
{% hint style="warning" %}
We have to **minus 25** since MMLU has 4 multiple choices - A, B, C or D. Assume we make a model that simply randomly chooses answers - it'll get 25% accuracy, and have a disk space of a few bytes. But clearly this is not a useful model.
{% endhint %}
On KL Divergence vs the base model, below is a table showcasing the improvements. Reminder the closer the KL Divergence is to 0, the better (ie 0 means identical to the full precision model)
| Quant | Baseline KLD | GB | New KLD | GB |
| --------- | ------------ | ----- | -------- | ----- |
| IQ1\_S | 1.035688 | 5.83 | 0.972932 | 6.06 |
| IQ1\_M | 0.832252 | 6.33 | 0.800049 | 6.51 |
| IQ2\_XXS | 0.535764 | 7.16 | 0.521039 | 7.31 |
| IQ2\_M | 0.26554 | 8.84 | 0.258192 | 8.96 |
| Q2\_K\_XL | 0.229671 | 9.78 | 0.220937 | 9.95 |
| Q3\_K\_XL | 0.087845 | 12.51 | 0.080617 | 12.76 |
| Q4\_K\_XL | 0.024916 | 15.41 | 0.023701 | 15.64 |
If we plot the ratio of the disk space increase and the KL Divergence ratio change, we can see a much clearer benefit! Our dynamic 2bit Q2\_K\_XL reduces KLD quite a bit (around 7.5%).
Truncated table of results for MMLU for Gemma 3 (27B). See below.
1. **Our dynamic 4bit version is 2GB smaller whilst having +1% extra accuracy vs the QAT version!**
2. Efficiency wise, 2bit Q2\_K\_XL and others seem to do very well!
| Quant | Unsloth | Unsloth + QAT | Disk Size | Efficiency |
| -------------- | --------- | ------------- | --------- | ---------- |
| IQ1\_M | 48.10 | 47.23 | 6.51 | 3.42 |
| IQ2\_XXS | 59.20 | 56.57 | 7.31 | 4.32 |
| IQ2\_M | 66.47 | 64.47 | 8.96 | 4.40 |
| Q2\_K\_XL | 68.70 | 67.77 | 9.95 | 4.30 |
| Q3\_K\_XL | 70.87 | 69.50 | 12.76 | 3.49 |
| **Q4\_K\_XL** | **71.47** | **71.07** | **15.64** | **2.94** |
| **Google QAT** | | **70.64** | **17.2** | **2.65** |
Click here for Full Google's Gemma 3 (27B) QAT Benchmarks:
| Model | Unsloth | Unsloth + QAT | Disk Size | Efficiency |
| -------------- | --------- | ------------- | --------- | ---------- |
| IQ1\_S | 41.87 | 43.37 | 6.06 | 3.03 |
| IQ1\_M | 48.10 | 47.23 | 6.51 | 3.42 |
| IQ2\_XXS | 59.20 | 56.57 | 7.31 | 4.32 |
| IQ2\_M | 66.47 | 64.47 | 8.96 | 4.40 |
| Q2\_K | 68.50 | 67.60 | 9.78 | 4.35 |
| Q2\_K\_XL | 68.70 | 67.77 | 9.95 | 4.30 |
| IQ3\_XXS | 68.27 | 67.07 | 10.07 | 4.18 |
| Q3\_K\_M | 70.70 | 69.77 | 12.51 | 3.58 |
| Q3\_K\_XL | 70.87 | 69.50 | 12.76 | 3.49 |
| Q4\_K\_M | 71.23 | 71.00 | 15.41 | 2.98 |
| **Q4\_K\_XL** | **71.47** | **71.07** | **15.64** | **2.94** |
| Q5\_K\_M | 71.77 | 71.23 | 17.95 | 2.58 |
| Q6\_K | 71.87 | 71.60 | 20.64 | 2.26 |
| Q8\_0 | 71.60 | 71.53 | 26.74 | 1.74 |
| **Google QAT** | | **70.64** | **17.2** | **2.65** |
## :llama: Llama 4 Bug Fixes + Run
We also helped and fixed a few Llama 4 bugs:
* Llama 4 Scout changed the RoPE Scaling configuration in their official repo. We helped resolve issues in llama.cpp to enable this [change here](https://github.com/ggml-org/llama.cpp/pull/12889)
* Llama 4's QK Norm's epsilon for both Scout and Maverick should be from the config file - this means using 1e-05 and not 1e-06. We helped resolve these in [llama.cpp](https://github.com/ggml-org/llama.cpp/pull/12889) and [transformers](https://github.com/huggingface/transformers/pull/37418)
* The Llama 4 team and vLLM also independently fixed an issue with QK Norm being shared across all heads (should not be so) [here](https://github.com/vllm-project/vllm/pull/16311). MMLU Pro increased from 68.58% to 71.53% accuracy.
* [Wolfram Ravenwolf](https://x.com/WolframRvnwlf/status/1909735579564331016) showcased how our GGUFs via llama.cpp attain much higher accuracy than third party inference providers - this was most likely a combination of the issues explained above, and also probably due to quantization issues.
As shown in our graph, our 4-bit Dynamic QAT quantization deliver better performance on 5-shot MMLU while also being smaller in size.
### Running Llama 4 Scout:
To run Llama 4 Scout for example, first clone llama.cpp:
```bash
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
-DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON -DLLAMA_CURL=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
Then download out new dynamic v 2.0 quant for Scout:
```python
# !pip install huggingface_hub hf_transfer
import os
os.environ["HF_HUB_ENABLE_HF_TRANSFER"] = "1"
from huggingface_hub import snapshot_download
snapshot_download(
repo_id = "unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF",
local_dir = "unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF",
allow_patterns = ["*IQ2_XXS*"],
)
```
And and let's do inference!
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
--model unsloth/Llama-4-Scout-17B-16E-Instruct-GGUF/Llama-4-Scout-17B-16E-Instruct-UD-IQ2_XXS.gguf \
--threads 32 \
--ctx-size 16384 \
--n-gpu-layers 99 \
-ot ".ffn_.*_exps.=CPU" \
--seed 3407 \
--prio 3 \
--temp 0.6 \
--min-p 0.01 \
--top-p 0.9 \
-no-cnv \
--prompt "<|header_start|>user<|header_end|>\n\nCreate a Flappy Bird game.<|eot|><|header_start|>assistant<|header_end|>\n\n"
```
{% endcode %}
{% hint style="success" %}
Read more on running Llama 4 here:
{% endhint %}
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/basics/dynamic-3.0-ggufs.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
