---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-2
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agent", "attention", "decode", "fp4", "fp8", "gpus", "moe", "nvfp4", "nvidia", "reasoning", "sglang", "tensorrt"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [107, 270]
sha256: bd6d891abb389c18dec77331a74b115d21b41dc6cee433cc73fd552eef4cd54d
---

# Load tokenizer and model

Please note that the model supports up to a 1M context size, although the default context size in the Hugging Face configuration is 256k due to higher VRAM requirements.

```
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
# Load tokenizer and model
tokenizer = AutoTokenizer.from_pretrained("nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4")
model = AutoModelForCausalLM.from_pretrained(
    "nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4",
    torch_dtype=torch.bfloat16,
    trust_remote_code=True,
    device_map="auto"
)
```
```
messages = [
    {"role": "user", "content": "Write a haiku about GPUs"},
]
tokenized_chat = tokenizer.apply_chat_template(
    messages,
    tokenize=True,
    add_generation_prompt=True,
    return_tensors="pt"
).to(model.device)
outputs = model.generate(
    tokenized_chat,
    max_new_tokens=1024,
    temperature=1.0,
    top_p=1.0,
    eos_token_id=tokenizer.eos_token_id
)
print(tokenizer.decode(outputs[0]))
```
`temperature=1.0` and `top_p=1.0` are recommended for reasoning tasks, while `temperature=0.6` and `top_p=0.95` are recommended for tool calling. 

If you’d like to use reasoning off, add `enable_thinking=False` to `apply_chat_template()`. By default, `enable_thinking` is set to be `True`.

```
tokenized_chat = tokenizer.apply_chat_template(
    messages,
    tokenize=True,
    enable_thinking=False,
    add_generation_prompt=True,
    return_tensors="pt"
).to(model.device)
# Use Greedy Search for reasoning off
outputs = model.generate(
    tokenized_chat,
    max_new_tokens=32,
    do_sample=False,
    num_beams=1,
    eos_token_id=tokenizer.eos_token_id
)
print(tokenizer.decode(outputs[0]))
```
For more detailed information on how to use the model with vLLM, please see this cookbook. If you are on Jetson Thor or DGX Spark, please use this vllm container.

```
pip install -U "vllm>=0.12.0"
```
Download the custom parser from the Hugging Face repository.

```
wget https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4/resolve/main/nano_v3_reasoning_parser.py
```
Launch a vLLM server using the custom parser.

```
VLLM_USE_FLASHINFER_MOE_FP4=1 \
VLLM_FLASHINFER_MOE_BACKEND=throughput \
vllm serve nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 \
  --served-model-name model \
  --max-num-seqs 8 \
  --tensor-parallel-size 1 \
  --max-model-len 262144 \
  --port 8000 \
  --trust-remote-code \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \
  --reasoning-parser-plugin nano_v3_reasoning_parser.py \
  --reasoning-parser nano_v3 \
  --kv-cache-dtype fp8
```
In the example above, we use a context length of 256k. You can increase the context size up to 1M to support longer contexts.
To enable this, set the `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` environment variable as shown below:

```
VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 \
VLLM_USE_FLASHINFER_MOE_FP4=1 \
VLLM_FLASHINFER_MOE_BACKEND=throughput \
vllm serve nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 \
  --served-model-name model \
  --max-num-seqs 8 \
  --tensor-parallel-size 1 \
  --max-model-len 262144 \
  --port 8000 \
  --trust-remote-code \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \
  --reasoning-parser-plugin nano_v3_reasoning_parser.py \
  --reasoning-parser nano_v3 \
  --kv-cache-dtype fp8
```
Here is an example client code for vLLM. By default, the endpoint has reasoning enabled. We recommend setting a high value (e.g., 10,000) for `max_tokens`.

```
curl http://localhost:8000/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d '{
        "model": "model",
        "messages":[{"role": "user", "content": "Write a haiku about GPUs"}],
        "max_tokens": 10000
    }'
```
If you’d like to use reasoning off with vLLM, you can do the following:

vLLM OpenAI curl request:

```
curl http://localhost:8000/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d '{
        "model": "model",
        "messages":[{"role": "user", "content": "Write a haiku about GPUs"}],
        "chat_template_kwargs": {"enable_thinking": false}
    }'
```
vLLM OpenAI client:

```
response = client.chat.completions.create(model=model, messages=messages, extra_body={"chat_template_kwargs": {"enable_thinking": False}})
```
For more detailed information on how to use the model with TRT-LLM, please see this cookbook.

```
# nano_v3 example yaml is https://github.com/NVIDIA/TensorRT-LLM/blob/main/examples/auto_deploy/nano_v3.yaml
trtllm-serve <model_path> \
--backend _autodeploy \
--trust_remote_code \
--reasoning_parser nano-v3 \
--tool_parser qwen3_coder \
--extra_llm_api_options nano_v3.yaml
```
### For more detailed information on how to use the model with SGLang, please see this cookbook.

```
python3 -m sglang.launch_server --model-path nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4 \
  --trust-remote-code \
  --tp 1 \
  --attention-backend flashinfer \
  --tool-call-parser qwen3_coder \
  --reasoning-parser nano_v3
```
The thinking budget allows developers to keep accuracy high and meet response‑time targets - which is especially crucial for customer support, autonomous agent steps, and edge devices where every millisecond counts.

With budget control, you can set a limit for internal reasoning:

- `reasoning_budget` : This is a threshold that will attempt to end the reasoning trace at the next newline encountered in the reasoning trace. If no newline is encountered within 500 tokens, it will abruptly end the reasoning trace at`reasoning_budget + 500` .

NOTE: This client will work with any OpenAI API compatible endpoint.


Client for supporting budget control:

