---
id: collect-250926-servers-hardware/servers-hardware/nvidia-dgx-spark-performance-ollama-blog
title: "NVIDIA DGX Spark performance"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia", "OpenAI"]
dates: ["2025-10-23"]
keywords: ["nvidia", "attention", "blackwell", "decode", "deepseek", "distribution", "mxfp4", "prefill", "quantization", "tokens per second"]
source: docs/RAG/clean4/nvidia-dgx-spark-performance-ollama-blog.md
source_anchor: ""
source_lines: [1, 82]
sha256: 30aa51a2a4623f52a7cc591c61bd389ea66520a10269a7fb00f01263006c72d9
---

# NVIDIA DGX Spark performance

## October 23, 2025

## Performance

We ran performance tests on release day firmware and an updated Ollama version to see how Ollama performs.

The tests were run using the latest NVIDIA DGX Spark firmware (580.95.05) and Ollama v0.12.6.

Each test is performed:

- 10 times
- Temperature set to 0
- Constrained to 500 tokens output
- Prompt: “write an in-depth summary of this story: $(head -n200 pg98.txt)” (please see the test script for the book, “A Tale of Two Cities”)
- Caching is disabled so repeated tests will not be faster

The test script and its readme are made available and can be customized for your own testing.

| **Device** | **Model name** | **Model size** | **Quantization** | **Prefill (tokens per second)** | **Decode (tokens per second)** | 
|---|---|---|---|---|---|
| NVIDIA DGX Spark | gpt-oss | 20B | MXFP4 | 3.224k | 58.27 | 
| NVIDIA DGX Spark | gpt-oss | 120B | MXFP4 | 1.169k | 41.14 | 
| NVIDIA DGX Spark | gemma3 | 12B | q4_K_M | 1.894k | 24.25 | 
| NVIDIA DGX Spark | gemma3 | 12B | q8_0 | 1.406k | 15.46 | 
| NVIDIA DGX Spark | gemma3 | 27B | q4_K_M | 834.1 | 10.83 | 
| NVIDIA DGX Spark | gemma3 | 27B | q8_0 | 585.4 | 7.210 | 
| NVIDIA DGX Spark | llama3.1 | 8B | q4_K_M | 7.614k | 38.02 | 
| NVIDIA DGX Spark | llama3.1 | 8B | q8_0 | 6.110k | 25.23 | 
| NVIDIA DGX Spark | llama3.1 | 70B | q4_K_M | 1.911k | 4.423 | 
| NVIDIA DGX Spark | deepseek-r1 | 14B | q4_K_M | 5.919k | 19.99 | 
| NVIDIA DGX Spark | deepseek-r1 | 14B | q8_0 | 4.667k | 13.32 | 
| NVIDIA DGX Spark | qwen3 | 32B | q4_K_M | 705.0 | 9.411 | 
| NVIDIA DGX Spark | qwen3 | 32B | q8_0 | 487.2 | 6.240 | 

*OpenAI’s gpt-oss models are tested using models officially provided by OpenAI, distributed via Ollama. Some GGUFs distributed online labeled as MXFP4 are further quantized to q8_0 in the attention layers. The same layers are BF16 on Ollama as intended by OpenAI.

### NVIDIA Firmware update

If you are using a DGX Spark firmware version below 580.95.05, it is recommended to use the DGX Dashboard to perform updates.

If you want to upgrade via the CLI, you will need to upgrade both the Ubuntu distribution as well as the firmware. Use the following commands:

```
sudo apt update
sudo apt dist-upgrade
sudo fwupdmgr refresh
sudo fwupdmgr upgrade
sudo reboot
```
## Get started with Ollama

Install Ollama:

```
curl -fsSL https://ollama.com/install.sh | sh
```
Then run a model:

```
ollama run gpt-oss
```
## Coding with Codex & Ollama

OpenAI’s Codex and Ollama work seamlessly together.

Install OpenAI’s Codex:

```
npm install -g @openai/codex
```
Once Codex is installed, use:

```
codex --oss --model gpt-oss 
```
The DGX Spark also supports the larger gpt-oss-120b model, fitting the entire model into the 120GB of VRAM provided by the GB10 Grace Blackwell Superchip:

```
codex --oss --model gpt-oss:120b
```
