---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k2-5-hugging-face-2
title: "moonshotai-kimi-k2-5-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["kimi", "agent", "agentic", "agents", "benchmark", "benchmarks", "claude", "deepseek", "gemini", "inference", "int4", "opus 4"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k2-5-hugging-face.md
source_anchor: ""
source_lines: [84, 182]
sha256: 07f495a1cc83e53a98f1456312c0a47ebb1e77ad9c854f92c8b0b7bd0e14bec9
---

# moonshotai-kimi-k2-5-hugging-face

1. General Testing Details
  - We report results for Kimi K2.5 and DeepSeek-V3.2 with thinking mode enabled, Claude Opus 4.5 with extended thinking mode, GPT-5.2 with xhigh reasoning effort, and Gemini 3 Pro with a high thinking level. For vision benchmarks, we additionally report results for Qwen3-VL-235B-A22B-Thinking.
  - Unless otherwise specified, all Kimi K2.5 experiments were conducted with temperature = 1.0, top-p = 0.95, and a context length of 256k tokens.
  - Benchmarks without publicly available scores were re-evaluated under the same conditions used for Kimi K2.5 and are marked with an asterisk (*).
  - We could not evaluate GPT-5.2 xhigh on all benchmarks due to service stability issues. For benchmarks that were not tested, we mark them as "-".
2. Text and Reasoning
  - HLE, AIME 2025, HMMT 2025 (Feb), and GPQA-Diamond were evaluated with a maximum completion budget of 96k tokens.
  - Results for AIME and HMMT are averaged over 32 runs (avg@32); GPQA-Diamond over 8 runs (avg@8).
  - For HLE, we report scores on the full set (text & image). Kimi K2.5 scores 31.5 (text) and 21.3 (image) without tools, and 51.8 (text) and 39.8 (image) with tools. The DeepSeek-V3.2 score corresponds to its text-only subset (marked with †) . Hugging Face access was blocked to prevent potential data leakage. HLE with tools uses simple context management: once the context exceeds a threshold, only the latest round of tool messages is retained.
3. Tool-Augmented / Agentic Search
  - Kimi K2.5 was equipped with search, code-interpreter, and web-browsing tools for HLE with tools and all agentic search benchmarks.
  - Except for BrowseComp (where K2.5 and DeepSeek-V3.2 used the discard-all strategy), no context management was applied, and tasks exceeding the supported context length were directly counted as failed.
  - The test system prompts emphasize deep and proactive tool use, instructing models to reason carefully, leverage tools, and verify uncertain information. Full prompts will be provided in the technical report.
  - Results for Seal-0 and WideSearch are averaged over four runs (avg@4).
4. Vision Benchmarks
  - Max-tokens = 64k, averaged over three runs (avg@3).
  - ZeroBench (w/ tools) uses max-tokens-per-step = 24k and max-steps = 30 for multi-step reasoning.
  - MMMU-Pro follows the official protocol, preserving input order and prepending images.
  - GPT-5.2-xhigh had ~10% failure rate (no output despite 3 retries), treated as incorrect; reported scores likely underestimate true performance.
  - WorldVQA, a benchmark designed to evaluate atomic vision-centric world knowledge. Access WorldVQA at https://github.com/MoonshotAI/WorldVQA.
  - OmniDocBench Score is computed as (1 − normalized Levenshtein distance) × 100, where a higher score denotes superior accuracy.
5. Coding Tasks
  - Terminal-Bench 2.0 scores were obtained with the default agent framework (Terminus-2) and the provided JSON parser. In our implementation, we evaluated Terminal-Bench 2.0 under non-thinking mode. This choice was made because our current context management strategy for the thinking mode is incompatible with Terminus-2.
  - For the SWE-Bench series of evaluations (including verified, multilingual, and pro), we used an internally developed evaluation framework. This framework includes a minimal set of tools—bash tool, createfile tool, insert tool, view tool, strreplace tool, and submit tool—along with tailored system prompts designed for the tasks. The highest scores were achieved under non-thinking mode.
  - The score of Claude Opus 4.5 on CyberGym is reported under the non-thinking setting.
  - All reported scores of coding tasks are averaged over 5 independent runs.
6. Long-Context Benchmarks
  - AA-LCR: scores averaged over three runs (avg@3).
  - LongBench-V2: identical prompts and input contexts standardized to ~128k tokens.
7. Agent Swarm
  - BrowseComp (Swarm Mode): main agent max 15 steps; sub-agents max 100 steps.
  - WideSearch (Swarm Mode): main and sub-agents max 100 steps.

Kimi-K2.5 adopts the same native int4 quantization method as Kimi-K2-Thinking.

You can access Kimi-K2.5's API on https://platform.moonshot.ai and we provide OpenAI/Anthropic-compatible API for you. To verify the deployment is correct, we also provide the Kimi Vendor Verifier. Currently, Kimi-K2.5 is recommended to run on the following inference engines:


- vLLM
- SGLang
- KTransformers

The minimum version requirement for `transformers` is `4.57.1`.

Deployment examples can be found in the Model Deployment Guide.

The usage demos below demonstrate how to call our official API.

For third-party APIs deployed with vLLM or SGLang, please note that:


Chat with video content is an experimental feature and is only supported in our official API for now.

The recommended `temperature` will be `1.0` for Thinking mode and `0.6` for Instant mode.

The recommended `top_p` is `0.95`.

To use instant mode, you need to pass `{'chat_template_kwargs': {"thinking": False}}` in `extra_body`.


This is a simple chat completion script which shows how to call K2.5 API in Thinking and Instant modes.

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
    print('====== Below is reasoning_content in Thinking Mode ======')
    print(f'reasoning content: {response.choices[0].message.reasoning_content}')
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
K2.5 supports Image and Video input.

The following example demonstrates how to call K2.5 API with image input:

