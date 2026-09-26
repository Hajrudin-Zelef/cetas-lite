---
id: collect-240926-huggingface/huggingface/unsloth-qwen3-8-flash-next-gguf-hugging-face-2
title: "Read our How to Run Qwen3.8-Flash-Next Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-03-05"]
keywords: ["qwen", "agent", "agentic", "benchmark", "claude", "embeddings", "inference", "kv cache", "latency", "parameters", "reasoning", "rotary"]
source: docs/RAG/clean_en/huggingface/unsloth-qwen3-8-flash-next-gguf-hugging-face.md
source_anchor: ""
source_lines: [109, 235]
sha256: 1420fd351101d18cdd2811808d54a85d28abec425c07c178aa08e5818de8c5f9
---

# Read our How to Run Qwen3.8-Flash-Next Guide!

1. ClawEval-MM: scores are reported as "pass@3 / average score". Pass@3 measures the percentage passed in at least one of three trials, and the average score is the mean score across the three trials.

2. RecreationBench: an in-house long-horizon application-recreation benchmark for evaluating hybrid-agent abilities spanning five platforms — desktop (Ubuntu, macOS, Windows), mobile (Android) and web.

3. OSWorld 2.0: scores are reported as "binary / partial". The binary score is the percentage of tasks that receive the full task reward, while the partial score aggregates the partial rewards obtained across all tasks.

4. Vision2Web: scores are reported as the average over the frontend, webpage and website categories, using the Claude Code harness and judged by gpt-5.4-2026-03-05.

5. MathVision, CharXiv (RQ): scores are reported as "without CI / with CI". A small number of incorrect ground-truth annotations in MathVision were corrected after manual verification. Our model's score is evaluated using a fixed prompt, e.g. "Please reason step by step, and put your final answer within \boxed{}." For other models, we report the higher score between runs with and without the \boxed{} formatting.

6. The best result in each row is shown in bold.

7. Empty cells (--) indicate scores not yet available or not applicable.

Qwen3.8-Flash-Next models operate in thinking mode by default, generating thinking content signified by `<think>\n...</think>\n\n` before producing the final responses.
To disable thinking content and obtain direct response, refer to the examples here.


We recommend using the following sets of sampling parameters for generation:


- Thinking Mode:
`temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0`- Instruct (or non-thinking) mode:
`temperature=0.7`, `top_p=0.80`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0`
Please note that the support for sampling parameters varies according to inference frameworks.


In multi-turn agentic tasks, lower reasoning effort does not always reduce overall task completion time. Although it may produce faster per-turn responses, it can also lead to insufficient analysis, more failures, and repeated retries, which may increase total latency and token consumption.


Qwen3.8-Flash-Next supports controlling thinking behavior via `enable_thinking`, `preserve_thinking`, and `reasoning_effort`.

t; supported levels are xhigh, medium, and low stream=True, stream_options={"include_usage": True}, )

Qwen3.8-Flash-Next will think by default before responding. You can obtain a direct response from the model without thinking by configuring the API parameters. For example,

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
    model="Qwen/Qwen3.8-Flash-Next",
    messages=messages,
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
If you are using APIs from Qwen Cloud, in addition to changing `model`, please use `"enable_thinking": False` instead of `"chat_template_kwargs": {"enable_thinking": False}`.


By default, Qwen3.8-Flash-Next retains thinking blocks from all historical messages, maintaining a complete reasoning trace across the conversation. This behavior, known as preserved thinking, ensures full context continuity and is especially beneficial for agent scenarios where decision consistency and reduced redundant reasoning are critical. It also improves KV cache utilization, optimizing inference efficiency in both thinking and non-thinking modes.

If you prefer to retain only the thinking blocks from the latest user message, you can disable this behavior by setting `preserve_thinking` to `False`:

```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [...]
chat_response = client.chat.completions.create(
    model="Qwen/Qwen3.8-Flash-Next",
    messages=messages,
    extra_body={
        "chat_template_kwargs": {"preserve_thinking": False},
    },
)
print("Chat response:", chat_response)
```
If you are using APIs from Qwen Cloud, in addition to changing `model`, please use `"preserve_thinking": False` directly instead of wrapping it in `chat_template_kwargs`.


To achieve optimal performance, we recommend the following settings:

1. **Sampling Parameters** : We suggest using the following sets of sampling parameters:
  - Thinking Mode: `temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
  - Instruct (or non-thinking) mode: `temperature=0.7` ,`top_p=0.80` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
 `presence_penalty` parameter between 0 and 2 to reduce endless repetition. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. Thinking Mode: 
3. **Adequate Output Length** : To optimize performance on agentic tasks, we recommend allocating sufficient output length to allow the model to generate detailed and comprehensive responses. For frameworks that support separate token limits for internal reasoning and final outputs, we suggest the following configuration within the 1M context length:
  - Reasoning Content: Set the maximum output length to 262,144 tokens.
  - Final Response: Set the maximum output length to 131,072 tokens.
4. **Processing Ultra-Long Texts** : Qwen3.8-Flash-Next natively supports context lengths of up to 262,144 tokens. For long-horizon tasks where the total length (including both input and output) exceeds this limit, we recommend using RoPE scaling techniques to handle long texts effectively, e.g., YaRN.YaRN is currently supported by several inference frameworks, e.g., vLLM, SGLang, and TokenSpeed. In general, there are two approaches to enabling YaRN for supported frameworks: 
  - Modifying the model configuration file: In the `config.json` file, change the`rope_parameters` fields in`text_config` to:```
{
    "mrope_interleaved": true,
    "mrope_section": [
        11,
        11,
        10
    ],
    "rope_type": "yarn",
    "rope_theta": 10000000,
    "partial_rotary_factor": 0.25,
    "factor": 4.0,
    "original_max_position_embeddings": 262144
}
```
 All the notable open-source frameworks implement static YaRN, which means the scaling factor remains constant regardless of input length, **potentially impacting performance on shorter texts.** We advise modifying the`rope_parameters` configuration only when processing long contexts is required. 
It is also recommended to modify the`factor` as needed. For example, if the typical context length for your application is 524,288 tokens, it would be better to set`factor` as 2.0.
5. **Long Video Understanding** : To optimize inference efficiency for plain text and images, the`size` parameter in the released`video_preprocessor_config.json` is conservatively configured. It is recommended to set the`longest_edge` parameter in the video_preprocessor_config file to 469,762,048 (corresponding to 224k video tokens) to enable higher frame-rate sampling for hour-scale videos and thereby achieve superior performance. For example,```
{"longest_edge": 469762048, "shortest_edge": 4096}
```
Alternatively, override the default values via engine startup parameters. For implementation details, refer to: vLLM / SGLang.

