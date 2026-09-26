---
id: collect-240926-huggingface/huggingface/qwen-qwen3-6-35b-a3b-hugging-face-2
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["gpu", "gpus", "inference", "kv cache", "memory", "multimodal", "parameters", "qwen", "reasoning", "sglang", "throughput", "tool use"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-6-35b-a3b-hugging-face.md
source_anchor: ""
source_lines: [148, 323]
sha256: 6cdb381ed3414f852dc7ccd00103e1a3920eaca53ff917784e607f7317935ee0
---

# Set the following accordingly

- **Standard Version** : The following command can be used to create an API endpoint with maximum context length 262,144 tokens using tensor parallel on 8 GPUs.```
python -m sglang.launch_server --model-path Qwen/Qwen3.6-35B-A3B --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3
```
- **Tool Use** : To support tool use, you can use the following command.```
python -m sglang.launch_server --model-path Qwen/Qwen3.6-35B-A3B --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3 --tool-call-parser qwen3_coder
```
- **Multi-Token Prediction (MTP)** : The following command is recommended for MTP:```
python -m sglang.launch_server --model-path Qwen/Qwen3.6-35B-A3B --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3 --speculative-algo NEXTN --speculative-num-steps 3 --speculative-eagle-topk 1 --speculative-num-draft-tokens 4
```

For detailed deployment guide, see the SGLang Qwen3.5 Cookbook.

vLLM is a high-throughput and memory-efficient inference and serving engine for LLMs.
`vllm>=0.19.0` is recommended for Qwen3.6, which can be installed using the following command in a fresh environment:

```
uv pip install vllm --torch-backend=auto
```
See its documentation for more details.

The following will create API endpoints at `http://localhost:8000/v1`:

- **Standard Version** : The following command can be used to create an API endpoint with maximum context length 262,144 tokens using tensor parallel on 8 GPUs.```
vllm serve Qwen/Qwen3.6-35B-A3B --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 
```
- **Tool Call** : To support tool use, you can use the following command.```
vllm serve Qwen/Qwen3.6-35B-A3B --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --enable-auto-tool-choice --tool-call-parser qwen3_coder 
```
- **Multi-Token Prediction (MTP)** : The following command is recommended for MTP:```
vllm serve Qwen/Qwen3.6-35B-A3B --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --speculative-config '{"method":"qwen3_next_mtp","num_speculative_tokens":2}'
```
- **Text-Only** : The following command skips the vision encoder and multimodal profiling to free up memory for additional KV cache:```
vllm serve Qwen/Qwen3.6-35B-A3B --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --language-model-only
```

For detailed deployment guide, see the vLLM Qwen3.5 Recipe.

KTransformers is a flexible framework for experiencing cutting-edge LLM inference optimizations with CPU-GPU heterogeneous computing. For running Qwen3.6 with KTransformers, see the KTransformers Deployment Guide.

Hugging Face Transformers contains a *lightweight* server which can be used for quick testing and moderate load deployment.
The latest `transformers` is required for Qwen3.6:

```
pip install "transformers[serving]"
```
See its documentation for more details. Please also make sure torchvision and pillow are installed.

Then, run `transformers serve` to launch a server with API endpoints at `http://localhost:8000/v1`; it will place the model on accelerators if available:

```
transformers serve Qwen/Qwen3.6-35B-A3B --port 8000 --continuous-batching
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
`temperature=0.6, top_p=0.95, top_k=20, min_p=0.0, presence_penalty=0.0, repetition_penalty=1.0`- Instruct (or non-thinking) mode:
`temperature=0.7, top_p=0.80, top_k=20, min_p=0.0, presence_penalty=1.5, repetition_penalty=1.0`
Please note that the support for sampling parameters varies according to inference frameworks.


Qwen3.6 models operate in thinking mode by default, generating thinking content signified by `<think>\n...</think>\n\n` before producing the final responses.
To disable thinking content and obtain direct response, refer to the examples here.


```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [
    {"role": "user", "content": "Type \"I love Qwen3.6\" backwards"},
]
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.6-35B-A3B",
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
    model="Qwen/Qwen3.6-35B-A3B",
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
    model="Qwen/Qwen3.6-35B-A3B",
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
Qwen3.6 does not officially support the soft switch of Qwen3, i.e., `/think` and `/nothink`.


Qwen3.6 will think by default before response. You can obtain direct response from the model without thinking by configuring the API parameters. For example,

