---
id: collect-240926-huggingface/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face-3
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "OpenAI", "vLLM"]
dates: []
keywords: ["agent", "fp8", "gpu", "inference", "mcp", "parameters", "qwen", "reasoning", "tool calling", "vllm"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face.md
source_anchor: ""
source_lines: [230, 409]
sha256: 491974a2e242c827c376dfc99f25964f59750307ea2ac24d62ecf2b0530768d4
---

# Set the following accordingly

KTransformers is a flexible framework for experiencing cutting-edge LLM inference optimizations with CPU-GPU heterogeneous computing. For running Qwen3.5 with KTransformers, see the KTransformers Deployment Guide.

Hugging Face Transformers contains a *lightweight* server which can be used for quick testing and moderate load deployment.
The latest `transformers` is required for Qwen3.5:

```
pip install "transformers[serving] @ git+https://github.com/huggingface/transformers.git@main"
```
See its documentation for more details. Please also make sure torchvision and pillow are installed.

Then, run `transformers serve` to launch a server with API endpoints at `http://localhost:8000/v1`; it will place the model on accelerators if available:

```
transformers serve --force-model Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --continuous-batching
```
The chat completions API is accessible via standard HTTP requests or OpenAI SDKs. Here, we show examples using the OpenAI Python SDK.

Before starting, make sure it is installed and the API key and the API base URL is configured, e.g.:

```
pip install -U openai
# Set the following accordingly
export OPENAI_BASE_URL="http://localhost:8000/v1"
export OPENAI_API_KEY="EMPTY"
```
We recommend using the following set of sampling parameters for generation


- Thinking mode for general tasks:
`temperature=1.0, top_p=0.95, top_k=20, min_p=0.0, presence_penalty=1.5, repetition_penalty=1.0`- Thinking mode for precise coding tasks (e.g. WebDev):
`temperature=0.6, top_p=0.95, top_k=20, min_p=0.0, presence_penalty=0.0, repetition_penalty=1.0`- Instruct (or non-thinking) mode for general tasks:
`temperature=0.7, top_p=0.8, top_k=20, min_p=0.0, presence_penalty=1.5, repetition_penalty=1.0`- Instruct (or non-thinking) mode for reasoning tasks:
`temperature=1.0, top_p=1.0, top_k=40, min_p=0.0, presence_penalty=2.0, repetition_penalty=1.0`
Please note that the support for sampling parameters varies according to inference frameworks.


```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [
    {"role": "user", "content": "Type \"I love Qwen3.5\" backwards"},
]
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.5-122B-A10B-FP8",
    messages=messages,
    max_tokens=81920,
    temperature=1.0,
    top_p=0.95,
    presence_penalty=1.5,
    extra_body={
        "top_k": 20,
    }, 
)
print("Chat response:", chat_response)
```
```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "image_url",
                "image_url": {
                    "url": "https://qianwen-res.oss-accelerate.aliyuncs.com/Qwen3.5/demo/CI_Demo/mathv-1327.jpg"
                }
            },
            {
                "type": "text",
                "text": "The centres of the four illustrated circles are in the corners of the square. The two big circles touch each other and also the two little circles. With which factor do you have to multiply the radii of the little circles to obtain the radius of the big circles?\nChoices:\n(A) $\\frac{2}{9}$\n(B) $\\sqrt{5}$\n(C) $0.8 \\cdot \\pi$\n(D) 2.5\n(E) $1+\\sqrt{2}$"
            }
        ]
    }
]
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.5-122B-A10B-FP8",
    messages=messages,
    max_tokens=81920,
    temperature=1.0,
    top_p=0.95,
    presence_penalty=1.5,
    extra_body={
        "top_k": 20,
    }, 
)
print("Chat response:", chat_response)
```
```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "video_url",
                "video_url": {
                    "url": "https://qianwen-res.oss-accelerate.aliyuncs.com/Qwen3.5/demo/video/N1cdUjctpG8.mp4"
                }
            },
            {
                "type": "text",
                "text": "How many porcelain jars were discovered in the niches located in the primary chamber of the tomb?"
            }
        ]
    }
]
# When vLLM is launched with `--media-io-kwargs '{"video": {"num_frames": -1}}'`,
# video frame sampling can be configured via `extra_body` (e.g., by setting `fps`).
# This feature is currently supported only in vLLM.
#
# By default, `fps=2` and `do_sample_frames=True`.
# With `do_sample_frames=True`, you can customize the `fps` value to set your desired video sampling rate.
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.5-122B-A10B-FP8",
    messages=messages,
    max_tokens=81920,
    temperature=1.0,
    top_p=0.95,
    presence_penalty=1.5,
    extra_body={
        "top_k": 20,
        "mm_processor_kwargs": {"fps": 2, "do_sample_frames": True},
    }, 
)
print("Chat response:", chat_response)
```
Qwen3.5 does not officially support the soft switch of Qwen3, i.e., `/think` and `/nothink`.


Qwen3.5 will think by default before response. You can obtain direct response from the model without thinking by configuring the API parameters. For example,

```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [
    {
        "role": "user",
        "content": [
            {
                "type": "image_url",
                "image_url": {
                    "url": "https://qianwen-res.oss-accelerate.aliyuncs.com/Qwen3.5/demo/RealWorld/RealWorld-04.png"
                }
            },
            {
                "type": "text",
                "text": "Where is this?"
            }
        ]
    }
]
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.5-122B-A10B-FP8",
    messages=messages,
    max_tokens=32768,
    temperature=0.7,
    top_p=0.8,
    presence_penalty=1.5,
    extra_body={
        "top_k": 20,
        "chat_template_kwargs": {"enable_thinking": False},
    }, 
)
print("Chat response:", chat_response)
```
If you are using APIs from Alibaba Cloud Model Studio, in addition to changing `model`, please use `"enable_thinking": False` instead of `"chat_template_kwargs": {"enable_thinking": False}`.


Qwen3.5 excels in tool calling capabilities.

We recommend using Qwen-Agent to quickly build Agent applications with Qwen3.5.

To define the available tools, you can use the MCP configuration file, use the integrated tool of Qwen-Agent, or integrate other tools by yourself.

