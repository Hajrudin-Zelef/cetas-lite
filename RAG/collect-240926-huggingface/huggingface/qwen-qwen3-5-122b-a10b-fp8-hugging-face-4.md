---
id: collect-240926-huggingface/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face-4
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agents", "embeddings", "fp8", "inference", "mcp", "multimodal", "parameters", "qwen", "reasoning", "rotary", "sglang"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face.md
source_anchor: ""
source_lines: [410, 537]
sha256: 0c7e7a97066a37fe895727399c0cbdae6b458bc673965e0bebe40ea21f88c40f
---

# Set the following accordingly

```
import os
from qwen_agent.agents import Assistant
# Define LLM
# Using Alibaba Cloud Model Studio
llm_cfg = {
    # Use the OpenAI-compatible model service provided by DashScope:
    'model': 'Qwen3.5-122B-A10B-FP8',
    'model_type': 'qwenvl_oai',
    'model_server': 'https://dashscope.aliyuncs.com/compatible-mode/v1',
    'api_key': os.getenv('DASHSCOPE_API_KEY'),
    'generate_cfg': {
        'use_raw_api': True,
        # When using Dash Scope OAI API, pass the parameter of whether to enable thinking mode in this way
        'extra_body': {
            'enable_thinking': True
        },
    },
}
# Using OpenAI-compatible API endpoint.
# functionality of the deployment frameworks and let Qwen-Agent automate the related operations.
#
# llm_cfg = {
#     # Use your own model service compatible with OpenAI API by vLLM/SGLang:
#     'model': 'Qwen/Qwen3.5-122B-A10B-FP8',
#     'model_type': 'qwenvl_oai',
#     'model_server': 'http://localhost:8000/v1',  # api_base
#     'api_key': 'EMPTY',
#
#     'generate_cfg': {
#         'use_raw_api': True,
#         # When using vLLM/SGLang OAI API, pass the parameter of whether to enable thinking mode in this way
#         'extra_body': {
#             'chat_template_kwargs': {'enable_thinking': True}
#         },
#     },
# }
# Define Tools
tools = [
    {'mcpServers': {  # You can specify the MCP configuration file
            "filesystem": {
                "command": "npx",
                "args": ["-y", "@modelcontextprotocol/server-filesystem", "/Users/xxxx/Desktop"]
            }
        }
    }
]
# Define Agent
bot = Assistant(llm=llm_cfg, function_list=tools)
# Streaming generation
messages = [{'role': 'user', 'content': 'Help me organize my desktop.'}]
for responses in bot.run(messages=messages):
    pass
print(responses)
# Streaming generation
messages = [{'role': 'user', 'content': 'Develop a dog website and save it on the desktop'}]
for responses in bot.run(messages=messages):
    pass
print(responses)
```
Qwen Code is an open-source AI agent for the terminal, optimized for Qwen models. It helps you understand large codebases, automate tedious work, and ship faster.

For more information, please refer to Qwen Code.

Qwen3.5 natively supports context lengths of up to 262,144 tokens. For long-horizon tasks where the total length (including both input and output) exceeds this limit, we recommend using RoPE scaling techniques to handle long texts effectively., e.g., YaRN.

YaRN is currently supported by several inference frameworks, e.g., `transformers`, `vllm`, `ktransformers` and `sglang`. 
In general, there are two approaches to enabling YaRN for supported frameworks:

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
    "original_max_position_embeddings": 262144,
}
```
- Passing command line arguments: For `vllm` , you can use```
VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 vllm serve ... --hf-overrides '{"text_config": {"rope_parameters": {"mrope_interleaved": true, "mrope_section": [11, 11, 10], "rope_type": "yarn", "rope_theta": 10000000, "partial_rotary_factor": 0.25, "factor": 4.0, "original_max_position_embeddings": 262144}}}' --max-model-len 1010000  
```
For `sglang` and`ktransformers` , you can use```
SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1 python -m sglang.launch_server ... --json-model-override-args '{"text_config": {"rope_parameters": {"mrope_interleaved": true, "mrope_section": [11, 11, 10], "rope_type": "yarn", "rope_theta": 10000000, "partial_rotary_factor": 0.25, "factor": 4.0, "original_max_position_embeddings": 262144}}}' --context-length 1010000
```

**potentially impacting performance on shorter texts.**
We advise modifying the `rope_parameters` configuration only when processing long contexts is required. 
It is also recommended to modify the `factor` as needed. For example, if the typical context length for your application is 524,288 tokens, it would be better to set `factor` as 2.0. 


To achieve optimal performance, we recommend the following settings:

1. **Sampling Parameters** :
  - We suggest using the following sets of sampling parameters depending on the mode and task type:  
    - **Thinking mode for general tasks** :`temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
    - **Thinking mode for precise coding tasks (e.g., WebDev)** :`temperature=0.6` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
    - **Instruct (or non-thinking) mode for general tasks** :`temperature=0.7` ,`top_p=0.8` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
    - **Instruct (or non-thinking) mode for reasoning tasks** :`temperature=1.0` ,`top_p=1.0` ,`top_k=40` ,`min_p=0.0` ,`presence_penalty=2.0` ,`repetition_penalty=1.0`
  - For supported frameworks, you can adjust the `presence_penalty` parameter between 0 and 2 to reduce endless repetitions. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. We suggest using the following sets of sampling parameters depending on the mode and task type:  
3. **Adequate Output Length** : We recommend using an output length of 32,768 tokens for most queries. For benchmarking on highly complex problems, such as those found in math and programming competitions, we suggest setting the max output length to 81,920 tokens. This provides the model with sufficient space to generate detailed and comprehensive responses, thereby enhancing its overall performance.
4. **Standardize Output Format** : We recommend using prompts to standardize model outputs when benchmarking.
  - **Math Problems** : Include "Please reason step by step, and put your final answer within \boxed{}." in the prompt.
  - **Multiple-Choice Questions** : Add the following JSON structure to the prompt to standardize responses: "Please show your choice in the`answer` field with only the choice letter, e.g.,`"answer": "C"` ."
5. **No Thinking Content in History** : In multi-turn conversations, the historical model output should only include the final output part and does not need to include the thinking content. It is implemented in the provided chat template in Jinja2. However, for frameworks that do not directly use the Jinja2 chat template, it is up to the developers to ensure that the best practice is followed.
6. **Long Video Understanding** : To optimize inference efficiency for plain text and images, the`size` parameter in the released`video_preprocessor_config.json` is conservatively configured. It is recommended to set the`longest_edge` parameter in the video_preprocessor_config file to 469,762,048 (corresponding to 224k video tokens) to enable higher frame-rate sampling for hour-scale videos and thereby achieve superior performance. For example,```
{"longest_edge": 469762048, "shortest_edge": 4096}
```

If you find our work helpful, feel free to give us a cite.

```
@misc{qwen3.5,
    title  = {{Qwen3.5}: Towards Native Multimodal Agents},
    author = {{Qwen Team}},
    month  = {February},
    year   = {2026},
    url    = {https://qwen.ai/blog?id=qwen3.5}
}
```
- Downloads last month
- 886,121
