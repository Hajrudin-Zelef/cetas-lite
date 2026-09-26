---
id: collect-240926-huggingface/huggingface/unsloth-kimi-k2-6-gguf-hugging-face-2
title: "Read our How to Run Kimi K2.6 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["kimi", "agent", "decode", "inference", "license", "mit license", "reasoning", "sglang", "vllm"]
source: docs/RAG/clean_en/huggingface/unsloth-kimi-k2-6-gguf-hugging-face.md
source_anchor: ""
source_lines: [105, 290]
sha256: e4c9034d4534ae6998bf27ddcaa204c8f9f053dd844b18765bcb221536057758
---

# Read our How to Run Kimi K2.6 Guide!

You can access Kimi-K2.6's API on https://platform.moonshot.ai and we provide OpenAI/Anthropic-compatible API for you. To verify the deployment is correct, we also provide the Kimi Vendor Verifier. Currently, Kimi-K2.6 is recommended to run on the following inference engines:


- vLLM
- SGLang
- KTransformers

Kimi-K2.6 has the same architecture as Kimi-K2.5, and the deployment method can be directly reused.

The version requirement for `transformers` is `>=4.57.1, <5.0.0`.

Deployment examples can be found in the Model Deployment Guide.

The usage demos below demonstrate how to call our official API.

For third-party APIs deployed with vLLM or SGLang, please note that:


The recommended `temperature` will be `1.0` for Thinking mode and `0.6` for Instant mode.

The recommended `top_p` is `0.95`.

To use instant mode, you need to pass `{'chat_template_kwargs': {"thinking": False}}` in `extra_body`.


This is a simple chat completion script which shows how to call K2.6 API in Thinking and Instant modes.

```
import openai
import base64
import requests
def simple_chat(client: openai.OpenAI, model_name: str):
    messages = [
        {'role': 'system', 'content': 'You are Kimi, an AI assistant created by Moonshot AI.'},
        {
            'role': 'user',
            'content': [
                {'type': 'text', 'text': 'which one is bigger, 9.11 or 9.9? think carefully.'}
            ],
        },
    ]
    response = client.chat.completions.create(
        model=model_name, messages=messages, stream=False, max_tokens=4096
    )
    print('====== Below is reasoning content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning}')
    print('====== Below is response in Thinking Mode ======')
    print(f'response: {response.choices[0].message.content}')
    # To use instant mode, pass {"thinking" = {"type":"disabled"}}
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        max_tokens=4096,
        extra_body={'thinking': {'type': 'disabled'}},  # this is for official API
        # extra_body= {'chat_template_kwargs': {"thinking": False}}  # this is for vLLM/SGLang
    )
    print('====== Below is response in Instant Mode ======')
    print(f'response: {response.choices[0].message.content}')
```
K2.6 supports Image and Video input.

The following example demonstrates how to call K2.6 API with image input:

```
import openai
import base64
import requests
def chat_with_image(client: openai.OpenAI, model_name: str):
    url = 'https://huggingface.co/moonshotai/Kimi-K2.6/resolve/main/figures/kimi-logo.png'
    image_base64 = base64.b64encode(requests.get(url).content).decode()
    messages = [
        {
            'role': 'user',
            'content': [
                {'type': 'text', 'text': 'Describe this image in detail.'},
                {
                    'type': 'image_url',
                    'image_url': {'url': f'data:image/png;base64, {image_base64}'},
                },
            ],
        }
    ]
    response = client.chat.completions.create(
        model=model_name, messages=messages, stream=False, max_tokens=8192
    )
    print('====== Below is reasoning content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning}')
    print('====== Below is response in Thinking Mode ======')
    print(f'response: {response.choices[0].message.content}')
    # Also support instant mode if you pass {"thinking" = {"type":"disabled"}}
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        max_tokens=4096,
        extra_body={'thinking': {'type': 'disabled'}},  # this is for official API
        # extra_body= {'chat_template_kwargs': {"thinking": False}}  # this is for vLLM/SGLang
    )
    print('====== Below is response in Instant Mode ======')
    print(f'response: {response.choices[0].message.content}')
    return response.choices[0].message.content
```
The following example demonstrates how to call K2.6 API with video input:

```
import openai
import base64
import requests
def chat_with_video(client: openai.OpenAI, model_name:str):
    url = 'https://huggingface.co/moonshotai/Kimi-K2.6/resolve/main/figures/demo_video.mp4'
    video_base64 = base64.b64encode(requests.get(url).content).decode()
    messages = [
        {
            "role": "user",
            "content": [
                {"type": "text","text": "Describe the video in detail."},
                {
                    "type": "video_url",
                    "video_url": {"url": f"data:video/mp4;base64,{video_base64}"},
                },
            ],
        }
    ]
    response = client.chat.completions.create(model=model_name, messages=messages)
    print('====== Below is reasoning content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning}')
    print('====== Below is response in Thinking Mode ======')
    print(f'response: {response.choices[0].message.content}')
    # Also support instant mode if pass {"thinking" = {"type":"disabled"}}
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        max_tokens=4096,
        extra_body={'thinking': {'type': 'disabled'}},  # this is for official API
        # extra_body= {'chat_template_kwargs': {"thinking": False}}  # this is for vLLM/SGLang
    )
    print('====== Below is response in Instant Mode ======')
    print(f'response: {response.choices[0].message.content}')
    return response.choices[0].message.content
```
Kimi K2.6 supports `preserve_thinking` mode, which retains full reasoning content across multi-turn interactions and enhances performance in coding agent scenarios.

This feature is disabled by default. The following example demonstrates how to call K2.6 API in `preserve_thinking` mode:

```
def chat_with_preserve_thinking(client: openai.OpenAI, model_name: str):
    messages = [
        {
            "role": "user",
            "content": "Tell me three random numbers."
        },
        {
            "role": "assistant",
            "reasoning_content": "I'll start by listing five numbers: 473, 921, 235, 215, 222, and I'll tell you the first three.",
            "content": "473, 921, 235"
        },
        {
            "role": "user",
            "content": "What are the other two numbers you have in mind?"
        }
    ]
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        max_tokens=4096,
        extra_body={'thinking': {'type': 'enabled', 'keep': 'all'}},  # this is for official API
        # extra_body={"chat_template_kwargs": {"thinking":True, "preserve_thinking": True}},  # this is for vLLM/SGLang
        # We recommend enabling preserve_thinking only in think mode.
    )
    # the assistant should mention 215 and 222 that appear in the prior reasoning content
    print(f"response: {response.choices[0].message.reasoning}")
    return response.choices[0].message.content
```
K2.6 shares the same design of Interleaved Thinking and Multi-Step Tool Call as K2 Thinking. For usage example, please refer to the K2 Thinking documentation.

Kimi K2.6 works best with Kimi Code CLI as its agent framework — give it a try at https://www.kimi.com/code.

Both the code repository and the model weights are released under the Modified MIT License.

If you have any questions, please reach out at support@moonshot.ai.

- Downloads last month
- 264,449
