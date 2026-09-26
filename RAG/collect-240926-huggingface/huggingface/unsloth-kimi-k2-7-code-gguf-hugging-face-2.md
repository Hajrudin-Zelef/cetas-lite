---
id: collect-240926-huggingface/huggingface/unsloth-kimi-k2-7-code-gguf-hugging-face-2
title: "Read our How to Run Kimi K2.7 Code Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Moonshot", "OpenAI", "vLLM"]
dates: []
keywords: ["kimi", "agent", "decode", "license", "mit license", "reasoning", "vllm"]
source: docs/RAG/clean_en/huggingface/unsloth-kimi-k2-7-code-gguf-hugging-face.md
source_anchor: ""
source_lines: [128, 194]
sha256: ccc40b109523dd9cd51029da237b040db47bbbfef3186e95c3fc1a9efbc2f8c1
---

# Read our How to Run Kimi K2.7 Code Guide!

```
import openai
import base64
import requests
def chat_with_video(client: openai.OpenAI, model_name:str):
    url = 'https://huggingface.co/moonshotai/Kimi-K2.7-Code/resolve/main/figures/demo_video.mp4'
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
```
Kimi K2.7 Code forces `preserve_thinking` mode, which retains full reasoning content across multi-turn interactions and enhances performance in coding agent scenarios.

This feature is enabled by default and can't be disabled. The following example demonstrates how to call K2.7-Code API in `preserve_thinking` mode:

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
            # Some API (e.g. vLLM) may not support reasoning_content, you can try reasoning instead
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
    )
    # the assistant should mention 215 and 222 that appear in the prior reasoning content
    print(f"response: {response.choices[0].message.reasoning}")
    return response.choices[0].message.content
```
K2.7-Code shares the same design of Interleaved Thinking and Multi-Step Tool Call as K2 Thinking. For usage example, please refer to the K2 Thinking documentation.

Kimi K2.7-Code works best with Kimi Code CLI as its agent framework — give it a try at https://www.kimi.com/code.

Both the code repository and the model weights are released under the Modified MIT License.

If you have any questions, please reach out at support@moonshot.ai.

- Downloads last month
- 360,598
