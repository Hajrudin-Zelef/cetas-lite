---
id: collect-261001-general-networking/general-networking/welcome-inkling-by-thinking-machines-2
title: "model_id = \"thinkingmachines/Inkling-NVFP4\""
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: []
keywords: ["nvfp4", "agent", "agentic", "context window", "cost", "decode", "distribution", "gguf", "gpu", "gpus", "inference", "inference engine"]
source: docs/RAG/collect-261001-general-networking/welcome-inkling-by-thinking-machines.md
source_anchor: ""
source_lines: [104, 286]
sha256: e5b74cd742b0e3e0f20c9683986eba5343cc1441c8a2e0e135336957221a433b
---

# model_id = "thinkingmachines/Inkling-NVFP4"

```
from transformers import AutoModelForMultimodalLM, AutoProcessor
model_id = "thinkingmachines/Inkling"
processor = AutoProcessor.from_pretrained(model_id)
model = AutoModelForMultimodalLM.from_pretrained(
    model_id,
    dtype="auto",
    device_map="auto",
)
image_url = (
    "https://huggingface.co/datasets/merve/vl-test-suite/"
    "resolve/main/pills.jpg"
)
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "image",
                "image": image_url,
            },
            {
                "type": "text",
                "text": "Do any of the components in this supplement interact?",
            },
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    add_generation_prompt=True,
    reasoning_effort="medium",
    return_dict=True,
    return_tensors="pt",
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
outputs = model.generate(**inputs, max_new_tokens=2000)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
processor.parse_response(response)
```
Inkling also takes in audio input. Below is an example inference snippet, which still uses the same `AutoModelForMultimodalLM` class.

## Text with audio inference

```
from transformers import AutoModelForMultimodalLM, AutoProcessor
model_id = "thinkingmachines/Inkling"
processor = AutoProcessor.from_pretrained(model_id)
model = AutoModelForMultimodalLM.from_pretrained(
    model_id,
    dtype="auto",
    device_map="auto",
)
audio_url = (
    "https://huggingface.co/datasets/merve/vl-test-suite/"
    "resolve/main/example_audio.mp3"
)
messages = [
    {
        "role": "user",
        "content": [
            {"type": "text", "text": "Transcribe the following speech to text."},
            {
                "type": "audio",
                "audio": audio_url,
            },
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
outputs = model.generate(**inputs, max_new_tokens=512)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
processor.parse_response(response)
```
For more realistic parallel deployment in a cluster of several nodes, please refer to the Slurm section below.

SGLang is one of the fastest deployment frameworks for Inkling at the time of release, as it includes a custom model implementation. The launch command below shards the model across 8 GPUs and serves an OpenAI-compatible API on port 30000.

```
pip install sglang
python3 -m sglang.launch_server \
 --model-path thinkingmachine/Inkling \
 --tp-size 8 \
 --served-model-name inkling \
 --host 0.0.0.0 \
 --port 30000
```
Match `--tp-size` to your GPU count. Add `--mem-fraction-static` (e.g. `0.85`) if you need to leave more headroom for the KV cache.

vLLM is strong for production serving. A single `vllm serve` command downloads the weights from the Hub, shards the model across your GPUs with tensor parallelism, and starts an OpenAI-compatible server on port 8000. You can view the vLLM Recipe for the model here to customize to your hardware.

```
# Requires nightly or vllm>=0.26 (once released)
uv pip install -U vllm --pre \
  --extra-index-url https://wheels.vllm.ai/nightly/cu130 \
  --extra-index-url https://download.pytorch.org/whl/cu130 \
  --index-strategy unsafe-best-match
vllm serve thinkingmachines/Inkling-NVFP4 \
  --trust-remote-code \
  --tokenizer-mode inkling \
  --tensor-parallel-size 8 \
  --enable-auto-tool-choice \
  --tool-call-parser inkling \
  --reasoning-parser inkling \
  --served-model-name inkling
```
In practice, you might need multiple nodes and a distribution tool like SLURM (see below). Key parameters are `--tensor-parallel-size` to the number of GPUs on your node, and use `--max-model-len` to cap the context window if you hit KV-cache memory limits.

```
curl http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "inkling",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```
You can infer with this model using several inference providers through Hugging Face. You can see all the code snippets to consume here. Below you can see how to use with the OpenAI client.

```
import os
from openai import OpenAI
client = OpenAI(
    base_url="https://router.huggingface.co/v1",
    api_key=os.environ["HF_TOKEN"],
)
completion = client.chat.completions.create(
    model="thinkingmachines/Inkling:auto",
    messages=[
        {
            "role": "user",
            "content": "What is the capital of France?",
        },
    ],
)
print(completion.choices[0].message)
```
Using the `“:auto”` suffix routes to your preferred provider in your settings; you can also use `“cheapest”` or `“:fastest”` as well. For this release, we cover the inference costs for 2 hours within the release for everyone.

Note: audio support in Inference Providers is work in progress and will be added shortly.

You can use `llama.cpp` to run quantized versions of the model on limited hardware. Unsloth have quantized the model down to 1-bit precision, reducing VRAM consumption by 95% over the original model.

```
llama serve -hf unsloth/inkling-GGUF:UD-IQ1_S
```
This starts an OpenAI-compatible server running at `http://localhost:8080``/v1` that you connect to in your preferred tool or clients. Heading there, you can start chatting with the model, and set it up with your favorite MCPs, pass in images or files conveniently and more!

Llama cpp also ships with a built-in UI that supports tools, mcp, and agentic workloads. Checkout Inkling running at 1-bit precision in the llama app:

Inkling GGUFs are also runnable in Unsloth Studio with dynamic 1-bit GGUFs which retain ~74.2% of top-1% accuracy whilst being 86% smaller.

Pi is a minimal coding agent harness you can use with different language models. You can use Pi with either an inference engine server endpoint, such as llama.cpp, or with Inference Providers on Hugging Face by adding this to your `~/.pi/agent/models.json` after installation.

```
{
  "providers": {
    "inference-providers": {
      "baseUrl": "https://router.huggingface.co/v1",
      "api": "openai-completions",
      "apiKey": "hf_...",
      "models": [
        {
          "id": "thinkingmachines/Inkling"
        }
      ]
    }
  }
}
```
Then you can start Pi in your project directory by calling `pi` and you’re good to go! In this demo, we give the model a hard math reasoning problem and it uses tools in pi to solve it.

Inkling is focused on broad multimodality reasoning and low token consumption, so try it out with document processing or audio tasks.

MTP adds extra layers to the model that predict several tokens at once, not just the next one. During inference, the extra layers act as “drafters” for speculative decoding, speeding up generation without compromising performance. With MTP, you get the exact same generated outputs, multipliers in generation speed-up at small memory cost in VRAM (due to serving the drafter). Thinking Machines also provides an MTP drafter with this release.

