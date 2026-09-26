---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face-3
title: "with uv: uv pip install vllm==0.18.1 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["China", "Hugging Face", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-02-24"]
keywords: ["vllm", "agent", "alignment", "decode", "gpus", "nvidia", "reasoning", "sglang", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face.md
source_anchor: ""
source_lines: [297, 485]
sha256: a65ac18f03e4b6da7774a3a3e3048ab269f0afe9751caf7d7fad1219c6c2f7f8
---

# with uv: uv pip install vllm==0.18.1 --torch-backend=auto

```
{
    "$schema": "https://opencode.ai/config.json",
    "model": "local/nvidia-nemotron-3-super",
    "provider": {
        "local": {
            "npm": "@ai-sdk/openai-compatible",
            "name": "local_backend",
            "options": {
                "baseURL": "http://localhost:8000/v1",
                "apiKey": "EMPTY"
            },
            "models": {
                "nvidia-nemotron-3-super": {
                    "name": "nvidia/nemotron-3-super",
                    "limit": {
                        "context": 1000000,
                        "output": 32768
                    }
                }
            }
        }
    },
    "agent": {
        "build": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        },
        "plan": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        }
    }
}
```
Update `baseURL` to match whichever backend you are running. The default port above (`8000`) matches the vLLM example; SGLang and TRT-LLM use `30000` and `8123` respectively.


To learn more about other supported agent scaffolds - check out this resource

##  **Advanced: Budget-Controlled Reasoning**

Set a hard token ceiling on the reasoning trace using `reasoning_budget`. The model will attempt to close the trace at the next newline before the budget is hit; if none is found within 500 tokens it closes abruptly at `reasoning_budget + 500`.

```
from typing import Any, Dict, List
import openai
from transformers import AutoTokenizer
class ThinkingBudgetClient:
    def __init__(self, base_url: str, api_key: str, tokenizer_name_or_path: str):
        self.tokenizer = AutoTokenizer.from_pretrained(tokenizer_name_or_path)
        self.client = openai.OpenAI(base_url=base_url, api_key=api_key)
    def chat_completion(        self,
        model: str,
        messages: List[Dict[str, Any]],
        reasoning_budget: int = 512,
        max_tokens: int = 1024,
        **kwargs,
    ) -> Dict[str, Any]:
        assert max_tokens > reasoning_budget, (
            f"reasoning_budget must be less than max_tokens. "
            f"Got {max_tokens=} and {reasoning_budget=}"
        )
        # Step 1: generate the reasoning trace up to the budget
        response = self.client.chat.completions.create(
            model=model, messages=messages, max_tokens=reasoning_budget, **kwargs
        )
        reasoning_content = response.choices[0].message.content
        if "" not in reasoning_content:
            reasoning_content = f"{reasoning_content}.\n\n\n"
        reasoning_tokens_len = len(
            self.tokenizer.encode(reasoning_content, add_special_tokens=False)
        )
        remaining_tokens = max_tokens - reasoning_tokens_len
        assert remaining_tokens > 0, (
            f"No tokens remaining for response ({remaining_tokens=}). "
            "Increase max_tokens or lower reasoning_budget."
        )
        # Step 2: continue from the closed reasoning trace
        messages.append({"role": "assistant", "content": reasoning_content})
        prompt = self.tokenizer.apply_chat_template(
            messages, tokenize=False, continue_final_message=True
        )
        response = self.client.completions.create(
            model=model, prompt=prompt, max_tokens=remaining_tokens, **kwargs
        )
        return {
            "reasoning_content": reasoning_content.strip().strip("").strip(),
            "content": response.choices[0].text,
            "finish_reason": response.choices[0].finish_reason,
        }
```
**Example usage** (32-token reasoning budget):

```
client = ThinkingBudgetClient(
    base_url="http://localhost:8000/v1",
    api_key="EMPTY",
    tokenizer_name_or_path="nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16",
)
result = client.chat_completion(
    model="nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16",
    messages=[
        {"role": "system", "content": "You are a helpful assistant. /think"},
        {"role": "user", "content": "What is 2+2?"},
    ],
    reasoning_budget=32,
    max_tokens=512,
    temperature=1.0,
    top_p=0.95,
)
print(result)
```
The model has been integrated into 🤗 Transformers since v5.3.0. We recommend using the Nemotron 3 Super container from the NeMo Framework to ensure all required libraries are available.

```
import torch
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16")
model = AutoModelForCausalLM.from_pretrained(
    "nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16",
    torch_dtype=torch.bfloat16,
    device_map="auto"
)
```
If your Transformers version is lower than v5.3.0, please add `trust_remote_code=True` when loading the model:

```
model = AutoModelForCausalLM.from_pretrained(
    "nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16",
    torch_dtype=torch.bfloat16,
    device_map="auto",
    trust_remote_code=True
)
```
Please note that the model supports up to a 1M context size, although the default context size in the Hugging Face configuration is 256k due to higher VRAM requirements.

Here is an example of generating outputs with reasoning enabled (the default):

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
if not isinstance(tokenized_chat, torch.Tensor):
    input_ids = tokenized_chat["input_ids"]
else:
    input_ids = tokenized_chat
outputs = model.generate(
    input_ids,
    max_new_tokens=50,
    temperature=1.0,
    top_p=0.95,
    eos_token_id=tokenizer.eos_token_id
)
print(tokenizer.decode(outputs[0]))
```
To disable reasoning, add `enable_thinking=False` to `apply_chat_template()`. By default, `enable_thinking` is set to `True`.

```
tokenized_chat = tokenizer.apply_chat_template(
    messages,
    tokenize=True,
    enable_thinking=False,
    add_generation_prompt=True,
    return_tensors="pt"
).to(model.device)
```
**Data Modality:** Text
**The total size:** 15,573,172,908,990 Tokens
**Total number of datasets:** 153
**Dataset partition:** *Training [100%], testing [0%], validation [0%]*
**Time period for training data collection:** 2013 to February 24, 2026
**Time period for testing data collection:** 2013 to February 24, 2026
**Time period for validation data collection:** 2013 to February 24, 2026
**Data Collection Method by dataset:** Hybrid: Automated, Human, Synthetic
**Labeling Method by dataset:** Hybrid: Automated, Human, Synthetic

NVIDIA-Nemotron-3-Super-120B-A12B-BF16 is pre-trained on a large corpus of high-quality curated and synthetically-generated data. It is trained in the English language, as well as 19 other languages and 43 programming languages. Our sources cover a variety of document types such as: webpages, dialogue, articles, and other written materials. The corpus spans domains including legal, math, science, finance, and more. We also include a small portion of question-answering, and alignment style data to improve model accuracy. The model was trained for approximately 25 trillion tokens.

The post-training corpus for NVIDIA-Nemotron-3-Super-120B-A12B-BF16 of high-quality curated and synthetically-generated data. Primary languages used for post-training include English, French, German, Italian, Japanese, Spanish, and Chinese.

