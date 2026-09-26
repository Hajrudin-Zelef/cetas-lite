---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k2-5-hugging-face-3
title: "moonshotai-kimi-k2-5-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["kimi", "agent", "decode", "license", "mit license", "reasoning", "research", "sglang", "vllm"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k2-5-hugging-face.md
source_anchor: ""
source_lines: [183, 270]
sha256: 56814cdb127aad786d21acb7d0d189cfe012363b583d6e538a9894aeaeae016d
---

# moonshotai-kimi-k2-5-hugging-face

```
import openai
import base64
import requests
def chat_with_image(client: openai.OpenAI, model_name: str):
    url = 'https://huggingface.co/moonshotai/Kimi-K2.5/resolve/main/figures/kimi-logo.png'
    image_base64 = base64.b64encode(requests.get(url).content).decode()
    messages = [
        {
            'role': 'user',
            'content': [
                {
                    'type': 'image_url',
                    'image_url': {'url': f'data:image/png;base64, {image_base64}'},
                },
                {'type': 'text', 'text': 'Describe this image in detail.'},
            ],
        }
    ]
    response = client.chat.completions.create(
        model=model_name, messages=messages, stream=False, max_tokens=8192
    )
    print('====== Below is reasoning_content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning_content}')
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
The following example demonstrates how to call K2.5 API with video input:

```
import openai
import base64
import requests
def chat_with_video(client: openai.OpenAI, model_name:str):
    url = 'https://huggingface.co/moonshotai/Kimi-K2.5/resolve/main/figures/demo_video.mp4'
    video_base64 = base64.b64encode(requests.get(url).content).decode()
    messages = [
        {
            "role": "user",
            "content": [
                {
                    "type": "video_url",
                    "video_url": {"url": f"data:video/mp4;base64,{video_base64}"},
                },
                {"type": "text","text": "Describe the video in detail."},
            ],
        }
    ]
    response = client.chat.completions.create(model=model_name, messages=messages)
    print('====== Below is reasoning_content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning_content}')
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
K2.5 shares the same design of Interleaved Thinking and Multi-Step Tool Call as K2 Thinking. For usage example, please refer to the K2 Thinking documentation.

Kimi K2.5 works best with Kimi Code CLI as its agent framework — give it a try at https://www.kimi.com/code.

Both the code repository and the model weights are released under the Modified MIT License.

If you have any questions, please reach out at support@moonshot.cn.

If you find K2.5 useful for your research, please kindly cite K2.5 technical report as follows:

