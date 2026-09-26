---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k2-instruct-hugging-face-2
title: "Your tool implementation"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "Moonshot", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agents", "fp8", "inference", "inference engine", "kimi", "leaderboard", "license", "mit license", "parameters", "sglang", "tensorrt", "tensorrt-llm"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k2-instruct-hugging-face.md
source_anchor: ""
source_lines: [106, 230]
sha256: fa773296d0d3ca2af55eecf5b3d67b72a09c98714acc235431c4313c4999d360
---

# Your tool implementation

<sup>• We only evaluate open-source pretrained models in this work. We report results for Qwen2.5-72B because the base checkpoint for Qwen3-235B-A22B was not open-sourced at the time of our study.</sup>

<sup>• All models are evaluated using the same evaluation protocol.</sup>

You can access Kimi K2's API on https://platform.moonshot.ai , we provide OpenAI/Anthropic-compatible API for you.

The Anthropic-compatible API maps temperature by `real_temperature = request_temperature * 0.6` for better compatible with existing applications.


Our model checkpoints are stored in the block-fp8 format, you can find it on Huggingface.

Currently, Kimi-K2 is recommended to run on the following inference engines:

- vLLM
- SGLang
- KTransformers
- TensorRT-LLM

Deployment examples for vLLM and SGLang can be found in the Model Deployment Guide.

Once the local inference service is up, you can interact with it through the chat endpoint:

```
def simple_chat(client: OpenAI, model_name: str):
    messages = [
        {"role": "system", "content": "You are Kimi, an AI assistant created by Moonshot AI."},
        {"role": "user", "content": [{"type": "text", "text": "Please give a brief self-introduction."}]},
    ]
    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False,
        temperature=0.6,
        max_tokens=256
    )
    print(response.choices[0].message.content)
```
The recommended temperature for Kimi-K2-Instruct is `temperature = 0.6`.
If no special instructions are required, the system prompt above is a good default.


Kimi-K2-Instruct has strong tool-calling capabilities. To enable them, you need to pass the list of available tools in each request, then the model will autonomously decide when and how to invoke them.

The following example demonstrates calling a weather tool end-to-end:

```
# Your tool implementation
def get_weather(city: str) -> dict:
    return {"weather": "Sunny"}
# Tool schema definition
tools = [{
    "type": "function",
    "function": {
        "name": "get_weather",
        "description": "Retrieve current weather information. Call this when the user asks about the weather.",
        "parameters": {
            "type": "object",
            "required": ["city"],
            "properties": {
                "city": {
                    "type": "string",
                    "description": "Name of the city"
                }
            }
        }
    }
}]
# Map tool names to their implementations
tool_map = {
    "get_weather": get_weather
}
def tool_call_with_client(client: OpenAI, model_name: str):
    messages = [
        {"role": "system", "content": "You are Kimi, an AI assistant created by Moonshot AI."},
        {"role": "user", "content": "What's the weather like in Beijing today? Use the tool to check."}
    ]
    finish_reason = None
    while finish_reason is None or finish_reason == "tool_calls":
        completion = client.chat.completions.create(
            model=model_name,
            messages=messages,
            temperature=0.6,
            tools=tools,          # tool list defined above
            tool_choice="auto"
        )
        choice = completion.choices[0]
        finish_reason = choice.finish_reason
        if finish_reason == "tool_calls":
            messages.append(choice.message)
            for tool_call in choice.message.tool_calls:
                tool_call_name = tool_call.function.name
                tool_call_arguments = json.loads(tool_call.function.arguments)
                tool_function = tool_map[tool_call_name]
                tool_result = tool_function(**tool_call_arguments)
                print("tool_result:", tool_result)
                messages.append({
                    "role": "tool",
                    "tool_call_id": tool_call.id,
                    "name": tool_call_name,
                    "content": json.dumps(tool_result)
                })
    print("-" * 100)
    print(choice.message.content)
```
The `tool_call_with_client` function implements the pipeline from user query to tool execution.
This pipeline requires the inference engine to support Kimi-K2’s native tool-parsing logic.
For streaming output and manual tool-parsing, see the Tool Calling Guide.

Both the code repository and the model weights are released under the Modified MIT License.

If you have any questions, please reach out at support@moonshot.cn.

- Downloads last month
- 252,250

## Spaces using moonshotai/Kimi-K2-Instruct 100

## Collection including moonshotai/Kimi-K2-Instruct

- mercor/apex-agents · Apex Agents View evaluation results  source leaderboard4.1<sup>*</sup>
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results  source leaderboard27.67
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results  source  leaderboard  81
- mercor/APEX-SWE · Apex Swe View evaluation results  source
- harborframework/terminal-bench-2.0 · Terminalbench 2 View evaluation results  source leaderboard27.8<sup>*</sup>
- SWE-bench/SWE-bench_Multilingual · Swe Bench Multilingual Resolved View evaluation results    leaderboard  47.3
