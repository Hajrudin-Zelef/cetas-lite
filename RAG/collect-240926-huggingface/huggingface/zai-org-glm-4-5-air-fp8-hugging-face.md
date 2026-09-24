---
id: collect-240926-huggingface/huggingface/zai-org-glm-4-5-air-fp8-hugging-face
title: "zai-org-glm-4-5-air-fp8-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["China", "Hugging Face", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["fp8", "glm", "agent", "agents", "attention", "benchmarks", "gpu", "gpus", "inference", "license", "llama", "lora"]
source: docs/RAG/clean_en/huggingface/zai-org-glm-4-5-air-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 153]
sha256: b1cd51b52066046a9e8a02d2f657030975f4be7fe479961458eeb2fe0e2a3a72
---

# zai-org-glm-4-5-air-fp8-hugging-face

<!-- source: https://huggingface.co/zai-org/GLM-4.5-Air-FP8 -->

👋 Join our Discord community.
    

    📖 Check out the GLM-4.5 technical blog, technical report, and Zhipu AI technical documentation.
    

     Code on GitHub
    

    📍 Use GLM-4.5 API services on Z.ai API Platform (Global) or 

 Zhipu AI Open Platform (Mainland China).
    

    👉 One click to GLM-4.5.

The **GLM-4.5** series models are foundation models designed for intelligent agents. GLM-4.5 has **355** billion total parameters with **32** billion active parameters, while GLM-4.5-Air adopts a more compact design with **106** billion total parameters and **12** billion active parameters. GLM-4.5 models unify reasoning, coding, and intelligent agent capabilities to meet the complex demands of intelligent agent applications.

Both GLM-4.5 and GLM-4.5-Air are hybrid reasoning models that provide two modes: thinking mode for complex reasoning and tool usage, and non-thinking mode for immediate responses.

We have open-sourced the base models, hybrid reasoning models, and FP8 versions of the hybrid reasoning models for both GLM-4.5 and GLM-4.5-Air. They are released under the MIT open-source license and can be used commercially and for secondary development.

As demonstrated in our comprehensive evaluation across 12 industry-standard benchmarks, GLM-4.5 achieves exceptional performance with a score of **63.2**, in the **3rd** place among all the proprietary and open-source models. Notably, GLM-4.5-Air delivers competitive results at **59.8** while maintaining superior efficiency.

For more eval results, show cases, and technical details, please visit our technical report or technical blog.

The model code, tool parser and reasoning parser can be found in the implementation of transformers, vLLM and SGLang.

You can directly experience the model on Hugging Face or ModelScope or download the model by following the links below.

| Model | Download Links | Model Size | Precision | 
|---|---|---|---|
| GLM-4.5 | 🤗 Hugging Face 🤖 ModelScope | 355B-A32B | BF16 | 
| GLM-4.5-Air | 🤗 Hugging Face 🤖 ModelScope | 106B-A12B | BF16 | 
| GLM-4.5-FP8 | 🤗 Hugging Face 🤖 ModelScope | 355B-A32B | FP8 | 
| GLM-4.5-Air-FP8 | 🤗 Hugging Face 🤖 ModelScope | 106B-A12B | FP8 | 
| GLM-4.5-Base | 🤗 Hugging Face 🤖 ModelScope | 355B-A32B | BF16 | 
| GLM-4.5-Air-Base | 🤗 Hugging Face 🤖 ModelScope | 106B-A12B | BF16 | 

We provide minimum and recommended configurations for "full-featured" model inference. The data in the table below is based on the following conditions:

1. All models use MTP layers and specify
`--speculative-num-steps 3 --speculative-eagle-topk 1 --speculative-num-draft-tokens 4` to ensure competitive
inference speed.
2. The `cpu-offload` parameter is not used.
3. Inference batch size does not exceed `8` .
4. All are executed on devices that natively support FP8 inference, ensuring both weights and cache are in FP8 format.
5. Server memory must exceed `1T` to ensure normal model loading and operation.

The models can run under the configurations in the table below:

| Model | Precision | GPU Type and Count | Test Framework | 
|---|---|---|---|
| GLM-4.5 | BF16 | H100 x 16 / H200 x 8 | sglang | 
| GLM-4.5 | FP8 | H100 x 8 / H200 x 4 | sglang | 
| GLM-4.5-Air | BF16 | H100 x 4 / H200 x 2 | sglang | 
| GLM-4.5-Air | FP8 | H100 x 2 / H200 x 1 | sglang | 

Under the configurations in the table below, the models can utilize their full 128K context length:

| Model | Precision | GPU Type and Count | Test Framework | 
|---|---|---|---|
| GLM-4.5 | BF16 | H100 x 32 / H200 x 16 | sglang | 
| GLM-4.5 | FP8 | H100 x 16 / H200 x 8 | sglang | 
| GLM-4.5-Air | BF16 | H100 x 8 / H200 x 4 | sglang | 
| GLM-4.5-Air | FP8 | H100 x 4 / H200 x 2 | sglang | 

The code can run under the configurations in the table below using Llama Factory:

| Model | GPU Type and Count | Strategy | Batch Size (per GPU) | 
|---|---|---|---|
| GLM-4.5 | H100 x 16 | Lora | 1 | 
| GLM-4.5-Air | H100 x 4 | Lora | 1 | 

The code can run under the configurations in the table below using Swift:

| Model | GPU Type and Count | Strategy | Batch Size (per GPU) | 
|---|---|---|---|
| GLM-4.5 | H20 (96GiB) x 16 | Lora | 1 | 
| GLM-4.5-Air | H20 (96GiB) x 4 | Lora | 1 | 
| GLM-4.5 | H20 (96GiB) x 128 | SFT | 1 | 
| GLM-4.5-Air | H20 (96GiB) x 32 | SFT | 1 | 
| GLM-4.5 | H20 (96GiB) x 128 | RL | 1 | 
| GLM-4.5-Air | H20 (96GiB) x 32 | RL | 1 | 

Please install the required packages according to `requirements.txt`.

```
pip install -r requirements.txt
```
Please refer to the `trans_infer_cli.py` code in the `inference` folder.

- Both BF16 and FP8 can be started with the following code:

```
vllm serve zai-org/GLM-4.5-Air \
    --tensor-parallel-size 8 \
    --tool-call-parser glm45 \
    --reasoning-parser glm45 \
    --enable-auto-tool-choice \
    --served-model-name glm-4.5-air
```
If you're using 8x H100 GPUs and encounter insufficient memory when running the GLM-4.5 model, you'll need
`--cpu-offload-gb 16` (only applicable to vLLM).

If you encounter `flash infer` issues, use `VLLM_ATTENTION_BACKEND=XFORMERS` as a temporary replacement. You can also
specify `TORCH_CUDA_ARCH_LIST='9.0+PTX'` to use `flash infer` (different GPUs have different TORCH_CUDA_ARCH_LIST
values, please check accordingly).

- BF16

```
python3 -m sglang.launch_server \
  --model-path zai-org/GLM-4.5-Air \
  --tp-size 8 \
  --tool-call-parser glm45  \
  --reasoning-parser glm45 \
  --speculative-algorithm EAGLE \
  --speculative-num-steps 3 \
  --speculative-eagle-topk 1 \
  --speculative-num-draft-tokens 4 \
  --mem-fraction-static 0.7 \
  --served-model-name glm-4.5-air \
  --host 0.0.0.0 \
  --port 8000
```
- FP8

```
python3 -m sglang.launch_server \
  --model-path zai-org/GLM-4.5-Air-FP8 \
  --tp-size 4 \
  --tool-call-parser glm45  \
  --reasoning-parser glm45  \
  --speculative-algorithm EAGLE \
  --speculative-num-steps 3  \
  --speculative-eagle-topk 1  \
  --speculative-num-draft-tokens 4 \
  --mem-fraction-static 0.7 \
  --disable-shared-experts-fusion \
  --served-model-name glm-4.5-air-fp8 \
  --host 0.0.0.0 \
  --port 8000
```
- When using `vLLM` and`SGLang` , thinking mode is enabled by default when sending requests. If you want to disable the
thinking switch, you need to add the`extra_body={"chat_template_kwargs": {"enable_thinking": False}}` parameter.
- Both support tool calling. Please use OpenAI-style tool description format for calls.
- For specific code, please refer to `api_request.py` in the`inference` folder.

- Downloads last month
- 52,793
