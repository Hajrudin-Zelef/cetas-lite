---
id: collect-240926-huggingface/huggingface/unsloth-qwen3-8-2-4t-a95b-gguf-hugging-face-2
title: "Read our How to Run Qwen3.8 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI"]
dates: []
keywords: ["qwen", "agent", "agentic", "benchmark", "claude", "cost", "fable 5", "gemini", "gpt-5.6", "inference", "multimodal", "opus 4"]
source: docs/RAG/clean_en/huggingface/unsloth-qwen3-8-2-4t-a95b-gguf-hugging-face.md
source_anchor: ""
source_lines: [121, 217]
sha256: 0d583d3f5ce9cf41e6e79e7e4cb1476c9a73e23cd75a527bd30a84de2aa4e976
---

# Read our How to Run Qwen3.8 Guide!

15. SkillsBench: Evaluated on the public SkillsBench v1.1 benchmark across 87 tasks, reporting the average score over three runs per task. Opus 4.8 and Fable 5 are evaluated on Claude Code; GPT-5.6 Sol is evaluated on Codex; the Qwen-series are evaluated on OpenCode. All results are from our own testing.

16. Automation-Bench: Evaluated on the 600-task public subset.

17. WideSearch: Evaluated with the Claude Code harness for external models and the Qwen-Agent harness for ours, reporting the average item-F1 over four runs.

18. $OneMillion-Bench: Evaluated using gemini-3.1-pro-preview.

19. PLawBench: Evaluated using gemini-3.1-pro-preview.

20. Empty cells (--): Scores are not yet available or are not applicable.

Qwen3.8-2.4T-A95B is a text-only model that requires thinking mode for all interactions. Multimodal inputs are not supported, and thinking cannot be disabled. Every response will automatically begin with reasoning enclosed in `<think>\n...</think>\n\n` before the final output.


We recommend using the following set of sampling parameters for generation:


Please note that the support for sampling parameters varies according to inference frameworks.


`reasoning_effort`, which can be used to adjust reasoning depth and control cost:  

In addition, `preserve_thinking` is enabled by default for all workloads for the best out-of-the-box experience.

```
pip install -U openai
# Set the following accordingly
export OPENAI_BASE_URL='your-base-url'
export OPENAI_API_KEY='your-api-key'
```
```
from openai import OpenAI
# Configured by environment variables
client = OpenAI()
messages = [{"role": "user", "content": "Write a Python function to merge two sorted linked lists."}]
completion = client.chat.completions.create(
    model="Qwen/Qwen3.8-2.4T-A95B",
    messages=messages,
    extra_body={
        "chat_template_kwargs": {
            "enable_thinking": True,  # on by default; should not be turned off
            "preserve_thinking": True, # on by default
        },
    },
    reasoning_effort="xhigh",  # xhigh by default; supported levels are xhigh, medium, and low
    stream=True,
    stream_options={"include_usage": True},
)
reasoning_content = ""
answer_content = ""
is_answering = False
print("\n" + "=" * 20 + "Reasoning" + "=" * 20 + "\n")
for chunk in completion:
    if not chunk.choices:
        print("\nUsage:")
        print(chunk.usage)
        continue
    delta = chunk.choices[0].delta
    if hasattr(delta, "reasoning_content") and delta.reasoning_content is not None:
        if not is_answering:
            print(delta.reasoning_content, end="", flush=True)
        reasoning_content += delta.reasoning_content
    if hasattr(delta, "content") and delta.content:
        if not is_answering:
            print("\n" + "=" * 20 + "Answer" + "=" * 20 + "\n")
            is_answering = True
        print(delta.content, end="", flush=True)
        answer_content += delta.content
```
If you are using APIs from Qwen Cloud, in addition to changing `model`, please pass `extra_body={"enable_thinking": True, "preserve_thinking": True}` instead of `extra_body={"chat_template_kwargs": {"enable_thinking": True, "preserve_thinking": True}}`.


To achieve optimal performance, we recommend the following settings:

1. **Sampling Parameters** :
  - We suggest using the following set of sampling parameters:  
    - `temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
  - For supported frameworks, you can adjust the `presence_penalty` parameter between 0 and 2 to reduce endless repetition. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. We suggest using the following set of sampling parameters:  
3. **Adequate Output Length** : To optimize performance on agentic tasks, we recommend allocating sufficient output length to allow the model to generate detailed and comprehensive responses. For frameworks that support separate token limits for internal reasoning and final outputs, we suggest the following configuration within the 1M context length:
  - **Reasoning Content:** Set the maximum output length to 262,144 tokens.
  - **Final Response:** Set the maximum output length to 131,072 tokens.

If you find our work helpful, feel free to give us a cite.

```
@misc{qwen38,
    title = {{Qwen3.8-Max}: A New Bar for Coding and Cowork},
    url = {https://qwen.ai/blog?id=qwen3.8},
    author = {{Qwen Team}},
    month = {August},
    year = {2026}
}
```
- Downloads last month
- 10,594
