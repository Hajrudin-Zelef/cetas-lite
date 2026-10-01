---
id: collect-261001-ia-llm/ia-llm/granite-4-2-llms-how-they-re-built-4
title: "Default: previous thinking is stripped to save context"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI", "vLLM"]
dates: []
keywords: ["agent", "agentic", "decode", "parameters", "reasoning", "tool calling", "vllm"]
source: docs/RAG/collect-261001-ia-llm/granite-4-2-llms-how-they-re-built.md
source_anchor: ""
source_lines: [253, 473]
sha256: b8d25ea6e74f1f453f62f0f151c105c3273770bff9b908d535bb7040d5583595
---

# Default: previous thinking is stripped to save context

```
pip install torch
pip install accelerate transformers
```
```
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
model_path = "ibm-granite/granite-4.2-30b"
tokenizer = AutoTokenizer.from_pretrained(model_path)
model = AutoModelForCausalLM.from_pretrained(model_path, device_map="cuda", torch_dtype=torch.bfloat16)
model.eval()
messages = [
    {"role": "user", "content": "How many r's are in the word 'strawberry'?"},
]
text = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True, enable_thinking=True)
inputs = tokenizer(text, return_tensors="pt").to(model.device)
with torch.no_grad():
    output = model.generate(**inputs, max_new_tokens=8192, temperature=1.0, top_p=0.95, do_sample=True)
print(tokenizer.decode(output[0][inputs.input_ids.shape[-1]:], skip_special_tokens=False))
```
## **Example Output**

```
<think>
Okay, let's see. The problem is to find how many 'r's are in the word 'strawberry'.
First, I need to write out the word: s t r a w b e r r y.
Now, I need to count the number of 'r' letters. Let's list each letter and check for 'r'.
1. s – not r
2. t – not r
3. r – yes, that's one
4. a – no
5. w – no
6. b – no
7. e – no
8. r – yes, that's two
9. r – yes, that's three
10. y – no
Total r's = 3.
</think>
There are **3** r's in the word "strawberry".<|im_end|>
```
```
messages = [
    {"role": "user", "content": "What is the capital of France?"},
]
text = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True, enable_thinking=False)
inputs = tokenizer(text, return_tensors="pt").to(model.device)
output = model.generate(**inputs, max_new_tokens=2048, temperature=1.0, top_p=0.95, do_sample=True)
print(tokenizer.decode(output[0][inputs.input_ids.shape[-1]:], skip_special_tokens=False))
```
## **Example Output**

```
<think></think>The capital of France is Paris.<|im_end|>
```
```
messages = [
    {"role": "user", "content": "What is 2 + 2?"},
]
text = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True,
                                     enable_thinking=True, low_effort=True)
inputs = tokenizer(text, return_tensors="pt").to(model.device)
output = model.generate(**inputs, max_new_tokens=4096, temperature=1.0, top_p=0.95, do_sample=True)
print(tokenizer.decode(output[0][inputs.input_ids.shape[-1]:], skip_special_tokens=False))
```
## **Example Output**

```
<think>
Simple answer.
</think>
2 + 2 = 4.<|im_end|>
```
Granite models support tool calling with integrated reasoning: the model reasons about which tool to call and why before calling it. Tools are defined with the OpenAI function definition schema.

```
tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "Get the current weather for a specified city.",
            "parameters": {
                "type": "object",
                "properties": {
                    "city": {"type": "string", "description": "Name of the city"}
                },
                "required": ["city"]
            }
        }
    }
]
messages = [
    {"role": "user", "content": "What's the weather like in Boston right now?"},
]
text = tokenizer.apply_chat_template(messages, tokenize=False, tools=tools,
                                     add_generation_prompt=True, enable_thinking=True)
inputs = tokenizer(text, return_tensors="pt").to(model.device)
output = model.generate(**inputs, max_new_tokens=4096, temperature=1.0, top_p=0.95, do_sample=True)
print(tokenizer.decode(output[0][inputs.input_ids.shape[-1]:], skip_special_tokens=False))
```
## **Example Output**

```
<think>
Okay, the user is asking for the weather in Boston right now. Let me check the tools
available. There's a function called get_current_weather that takes a city parameter.
I need to call that with the city set to Boston.
</think>
<tool_call>
<function=get_current_weather>
<parameter=city>
Boston
</parameter>
</function>
</tool_call>
<|im_end|>
```
```
messages = [
    {"role": "user", "content": "What's the weather like in Boston right now?"},
    {"role": "assistant", "content": "<think>\nThe user wants to know the current weather in Boston. I should call get_current_weather.\n</think>",
     "tool_calls": [{"function": {"name": "get_current_weather", "arguments": {"city": "Boston"}}}]},
    {"role": "tool", "content": '{"temperature": "72°F", "condition": "Partly cloudy", "humidity": "65%"}'},
]
text = tokenizer.apply_chat_template(messages, tokenize=False, tools=tools,
                                     add_generation_prompt=True, enable_thinking=True)
inputs = tokenizer(text, return_tensors="pt").to(model.device)
output = model.generate(**inputs, max_new_tokens=4096, temperature=1.0, top_p=0.95, do_sample=True)
print(tokenizer.decode(output[0][inputs.input_ids.shape[-1]:], skip_special_tokens=False))
```
## **Example Output**

```
<think>
The tool returned the weather data for Boston: temperature 72°F, partly cloudy, humidity 65%.
I need to present this information clearly to the user.
</think>
The current weather in Boston is 72°F, partly cloudy, with 65% humidity.<|im_end|>
```
```
messages = [
    {"role": "user", "content": "What is 15 * 37?"},
    {"role": "assistant", "content": "<think>\nLet me calculate 15 * 37.\n15 * 37 = 15 * 30 + 15 * 7 = 450 + 105 = 555\n</think>\n15 * 37 = 555"},
    {"role": "user", "content": "Now divide that by 5"},
]
# Default: previous thinking is stripped to save context
text = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True,
                                     enable_thinking=True, truncate_history_thinking=True)
# To preserve full history:
text_full = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True,
                                          enable_thinking=True, truncate_history_thinking=False)
```
```
import re
def parse_model_output(text):
    """Separate thinking content from final answer."""
    think_match = re.search(r'<think>(.*?)</think>', text, re.DOTALL)
    if think_match:
        thinking = think_match.group(1).strip()
        answer_start = text.find('</think>') + len('</think>')
        answer_end = text.find('<|im_end|>', answer_start)
        answer = text[answer_start:answer_end].strip() if answer_end != -1 else text[answer_start:].strip()
    else:
        thinking, answer = "", text.strip()
    return thinking, answer
thinking, answer = parse_model_output(output_text)
```
Granite models can serve as the backbone for agentic coding tools. Because they support reasoning and tool calling through the OpenAI-compatible API, they integrate with popular agentic harnesses without extra adapters. Start the vLLM server, then follow the harness-specific instructions below.

OpenCode is an AI coding agent that runs in your terminal.

**Install:**

```
curl -fsSL https://opencode.ai/install | bash
```
**Configure** `~/.config/opencode/opencode.json`:

```
{
  "$schema": "https://opencode.ai/config.json",
  "model": "local/granite-4.2-30b",
  "provider": {
    "local": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "vLLM (local)",
      "options": {
        "baseURL": "http://localhost:8000/v1",
        "apiKey": "EMPTY"
      },
      "models": {
        "granite-4.2-30b": {
          "name": "Granite 4.2 30B",
          "limit": {
            "context": 131072,
            "output": 8192
          }
        }
      }
    }
  }
}
```
**Run:**

```
opencode
opencode run "your task description"
```
For full documentation, see opencode.ai/docs.

Pi is a minimal agent harness for AI-powered coding that runs in your terminal. It supports custom providers via a `models.json` configuration file.

**Install:**

```
curl -fsSL https://pi.dev/install.sh | sh
```
**Configure** `~/.pi/agent/models.json`:

