---
id: collect-240926-huggingface/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face-2
title: "Log in once; the token is cached at ~/.cache/huggingface/token"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["context window", "cost", "decode", "fp8", "gpu", "gpus", "kv cache", "lpddr5x", "memory", "nvfp4", "nvidia", "omni"]
source: docs/RAG/clean_en/huggingface/nvidia-nemotron-3-nano-omni-30b-a3b-reasoning-fp8-hugging-face.md
source_anchor: ""
source_lines: [186, 421]
sha256: 353906c102d51b6876a932450c5dceafd3665a61b2369065b9654c1dc83f58a8
---

# Log in once; the token is cached at ~/.cache/huggingface/token

```
docker pull vllm/vllm-openai:v0.20.0
```
```
WEIGHTS=/path/to/nemotron-3-nano-omni-weights
# The image does not include audio packages so we need to install them with "pip install vllm[audio]" as done in the command below
docker run --rm -it \
  --gpus all \
  --ipc=host -p 8000:8000 \
  --shm-size=16g \
  --name vllm-nemotron-omni \
  -v "${WEIGHTS}:/model:ro" \
  --entrypoint /bin/bash \
  vllm/vllm-openai:v0.20.0 -c  \
  "pip install vllm[audio] && vllm serve /model \
  --served-model-name=nemotron_3_nano_omni \
  --max-num-seqs 8 \
  --max-model-len 131072 \
  --port 8000 \
  --trust-remote-code \
  --gpu-memory-utilization 0.8 \
  --limit-mm-per-prompt '{\"video\": 1, \"image\": 1, \"audio\": 1}' \
  --media-io-kwargs '{\"video\": {\"fps\": 2,  \"num_frames\": 256}}' \
  --allowed-local-media-path=/ \
  --enable-prefix-caching \
  --max-num-batched-tokens 32768 \
  --reasoning-parser nemotron_v3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder"
```
In another terminal, verify the server is ready:

```
curl -sS http://localhost:8000/v1/models | python3 -m json.tool
```
| Flag | Purpose | Spark Guidance | 
|---|---|---|
| `--gpus all` | Select GPU | Spark has one GB10 GPU; `all` is equivalent to`device=0` | 
| `--max-model-len` | Max context window | Start at 131072; reduce if you hit OOM (see Memory Tuning below) | 

Spark uses **unified LPDDR5X memory** (~128 GB shared between CPU and GPU), not separate system + VRAM pools. Two levers, in order of impact:

1. **Lower `--gpu-memory-utilization`** from 0.85 → 0.70 to free ~19 GB back to the OS and re-enable weight prefetch. Cost: smaller KV cache budget.
2. **Lower `--max-model-len`** to reduce KV cache allocation (e.g. halving context window halves KV cache at`--max-num-seqs=1` ).
Combined override:

```
  --gpu-memory-utilization=0.70 \
  --max-model-len=32768 \
```
This model can also be deployed with TensorRT-LLM - see relevant instructions here.

This model can also be deployed with TensorRT Edge-LLM on NVIDIA Jetson Thor - see the Jetson AI Lab model page and the TensorRT Edge-LLM Quick Start Guide.

The BF16 variant of this model is supported on SGLang, with the following images:

- **CUDA 13.0:**`lmsysorg/sglang:dev-cu13-nemotronh-nano-omni-reasoning-v3`
- **CUDA 12.9:**`lmsysorg/sglang:dev-nemotronh-nano-omni-reasoning-v3`

`librosa` must be installed first:
`pip install librosa --break-system-packages`

To serve:
`sglang serve --model-path nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-BF16 --trust-remote-code`

NVFP4 and FP8 support to come.


For everything not covered here (API examples, reasoning mode, video tuning), follow the general instructions.

Use the upstream multi-arch **CUDA 13.0** docker image linked above. Docker will automatically pull the `arm64` variant.

```
docker pull lmsysorg/sglang:dev-cu13-nemotronh-nano-omni-reasoning-v3
```
```
WEIGHTS=/path/to/nemotron-3-nano-omni-weights
# The image does not include audio packages so we need to install them with "pip install librosa" as done in the command below
docker run --gpus all -it --rm \
  -p 30000:30000 \
  -v "${WEIGHTS}:/model:ro" \
  --shm-size 16g \
  lmsysorg/sglang:dev-cu13-nemotronh-nano-omni-reasoning-v3 \
  bash -c "pip install librosa && python3 -m sglang.launch_server --model-path /model \
  --host 0.0.0.0 \
  --port 30000 \
  --trust-remote-code \
  --mem-fraction-static 0.8 \
  --max-running-requests 8 \
  --tool-call-parser qwen3_coder \
  --reasoning-parser nemotron_3"
```
In another terminal, verify the server is ready:

```
curl -sS http://localhost:30000/v1/models | python3 -m json.tool
```
| Flag | Purpose | Spark Guidance | 
|---|---|---|
| `--gpus all` | Select GPU | Spark has one GB10 GPU; `all` is equivalent to`device=0` | 
| `--context-length` | Max context window | Start with default; reduce if you hit OOM (see Memory Tuning below) | 

Spark uses **unified LPDDR5X memory** (~128 GB shared between CPU and GPU), not separate system + VRAM pools. Two levers, in order of impact:

1. **Lower `--mem-fraction-static`** from 0.80 → 0.70 to free ~13 GB back to the OS and re-enable weight prefetch. Cost: smaller KV cache budget.
2. **Lower `--context-length`** to reduce KV cache allocation (e.g. halving context window halves KV cache at`--max-running-requests=1` ).
Combined override:

```
  --mem-fraction-static=0.70 \
  --context-length=32768 \
```
```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="")
MODEL = "nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4"
```
**Image Example**

```
import base64
 
def image_to_data_url(path: str) -> str:
    with open(path, "rb") as f:
        b64 = base64.b64encode(f.read()).decode("utf-8")
    return f"data:image/jpeg;base64,{b64}"
 
image_url = image_to_data_url("media/example1a.jpeg")
 
response = client.chat.completions.create(
    model=MODEL,
    messages=[
        {
            "role": "user",
            "content": [
                {"type": "text", "text": "Describe this image in detail."},
                {"type": "image_url", "image_url": {"url": image_url}},
            ],
        }
    ],
    max_tokens=1024,
    temperature=0.2,
    extra_body={"top_k": 1, "chat_template_kwargs": {"enable_thinking": False}},
)
print(response.choices[0].message.content)
```
**Audio Example**

```
from pathlib import Path
 
audio_url = Path("media/2414-165385-0000.wav").resolve().as_uri()
 
response = client.chat.completions.create(
    model=MODEL,
    messages=[
        {
            "role": "user",
            "content": [
                {"type": "audio_url", "audio_url": {"url": audio_url}},
                {"type": "text", "text": "Transcribe this audio."},
            ],
        }
    ],
    max_tokens=1024,
    temperature=0.2,
    extra_body={"top_k": 1, "chat_template_kwargs": {"enable_thinking": False}},
)
print(response.choices[0].message.content)
```
**Video Example**

```
from pathlib import Path
 
video_url = Path("media/demo.mp4").resolve().as_uri()
reasoning_budget = 16384
grace_period = 1024
 
response = client.chat.completions.create(
    model=MODEL,
    messages=[
        {
            "role": "user",
            "content": [
                {"type": "video_url", "video_url": {"url": video_url}},
                {"type": "text", "text": "Describe this video."},
            ],
        }
    ],
    max_tokens=20480,
    temperature=0.6,
    top_p=0.95,
    extra_body={
        "thinking_token_budget": reasoning_budget + grace_period,
        "chat_template_kwargs": {
            "enable_thinking": True,
            "reasoning_budget": reasoning_budget,
        },
        "mm_processor_kwargs": {"use_audio_in_video": False},
    },
)
print(response.choices[0].message.content)
```
**Text Example (curl)**

```
curl -sS http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model":"nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4","messages":[{"role":"user","content":"Hello, what can you do?"}],"temperature":0.2,"top_k":1}' \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['choices'][0]['message']['content'])"
```
**PDF Example (page-by-page via Python)**

The API accepts **images**, not raw PDF files. The script below renders each page to PNG and sends it as base64. Save as **`pdf_vlm_chat.py`** and install dependencies: `pip install pymupdf pillow requests`.

## pdf_vlm_chat.py (click to expand)

```
#!/usr/bin/env python3
"""Send PDF page(s) as images to a vLLM /v1/chat/completions endpoint."""
from __future__ import annotations
 
import argparse, base64, sys
from io import BytesIO
from pathlib import Path
 
import requests
 
try:
    import fitz
    from PIL import Image
except ImportError:
    print("Install: pip install pymupdf pillow requests", file=sys.stderr)
    sys.exit(1)
 
