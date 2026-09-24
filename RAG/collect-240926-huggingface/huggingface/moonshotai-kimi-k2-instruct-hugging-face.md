---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k2-instruct-hugging-face
title: "Your tool implementation"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Moonshot", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agentic", "agents", "attention", "benchmark", "claude", "compute", "deepseek", "fine-tuning", "fp8", "gemini", "inference", "inference engine"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k2-instruct-hugging-face.md
source_anchor: ""
source_lines: [1, 230]
sha256: 07f27dd98bbb09d8ca5fe43458e549dc34e062d182b21e369076e129149427ce
---

# Your tool implementation

<!-- source: https://huggingface.co/moonshotai/Kimi-K2-Instruct -->

- Messages with `name` field are now supported. We’ve also moved the chat template to a standalone file for easier viewing.

- We further modified our chat template to improve its robustness. The default system prompt has also been updated.

- We have updated our tokenizer implementation. Now special tokens like `[EOS]` can be encoded to their token ids.
- We fixed a bug in the chat template that was breaking multi-turn tool calls.

Kimi K2 is a state-of-the-art mixture-of-experts (MoE) language model with 32 billion activated parameters and 1 trillion total parameters. Trained with the Muon optimizer, Kimi K2 achieves exceptional performance across frontier knowledge, reasoning, and coding tasks while being meticulously optimized for agentic capabilities.

- Large-Scale Training: Pre-trained a 1T parameter MoE model on 15.5T tokens with zero training instability.
- MuonClip Optimizer: We apply the Muon optimizer to an unprecedented scale, and develop novel optimization techniques to resolve instabilities while scaling up.
- Agentic Intelligence: Specifically designed for tool use, reasoning, and autonomous problem-solving.

- **Kimi-K2-Base** : The foundation model, a strong start for researchers and builders who want full control for fine-tuning and custom solutions.
- **Kimi-K2-Instruct** : The post-trained model best for drop-in, general-purpose chat and agentic experiences. It is a reflex-grade model without long thinking.

| **Architecture** | Mixture-of-Experts (MoE) | 
| **Total Parameters** | 1T | 
| **Activated Parameters** | 32B | 
| **Number of Layers** (Dense layer included) | 61 | 
| **Number of Dense Layers** | 1 | 
| **Attention Hidden Dimension** | 7168 | 
| **MoE Hidden Dimension** (per Expert) | 2048 | 
| **Number of Attention Heads** | 64 | 
| **Number of Experts** | 384 | 
| **Selected Experts per Token** | 8 | 
| **Number of Shared Experts** | 1 | 
| **Vocabulary Size** | 160K | 
| **Context Length** | 128K | 
| **Attention Mechanism** | MLA | 
| **Activation Function** | SwiGLU | 

| Benchmark | Metric | <sup>Kimi K2 Instruct</sup> | <sup>DeepSeek-V3-0324</sup> | <sup>Qwen3-235B-A22B  <sup>(non-thinking)</sup></sup> | <sup>Claude Sonnet 4  <sup>(w/o extended thinking)</sup></sup> | <sup>Claude Opus 4  <sup>(w/o extended thinking)</sup></sup> | <sup>GPT-4.1</sup> | <sup>Gemini 2.5 Flash   Preview (05-20)</sup> | 
|---|---|---|---|---|---|---|---|---|
| **Coding Tasks** |  |  |  |  |  |  |  |  | 
| LiveCodeBench v6 <sup>(Aug 24 - May 25)</sup> | Pass@1 | **53.7** | 46.9 | 37.0 | 48.5 | 47.4 | 44.7 | 44.7 | 
| OJBench | Pass@1 | **27.1** | 24.0 | 11.3 | 15.3 | 19.6 | 19.5 | 19.5 | 
| MultiPL-E | Pass@1 | **85.7** | 83.1 | 78.2 | 88.6 | **89.6** | 86.7 | 85.6 | 
| SWE-bench Verified <sup>(Agentless Coding)</sup> | Single Patch w/o Test (Acc) | **51.8** | 36.6 | 39.4 | 50.2 | **53.0** | 40.8 | 32.6 | 
| SWE-bench Verified <sup>(Agentic Coding)</sup> | Single Attempt (Acc) | **65.8** | 38.8 | 34.4 | **72.7**<sup>*</sup> | 72.5 <sup>*</sup> | 54.6 | — | 
|  | Multiple Attempts (Acc) | **71.6** | — | — | **80.2** | 79.4 <sup>*</sup> | — | — | 
| SWE-bench Multilingual <sup>(Agentic Coding)</sup> | Single Attempt (Acc) | **47.3** | 25.8 | 20.9 | **51.0** | — | 31.5 | — | 
| TerminalBench | Inhouse Framework (Acc) | **30.0** | — | — | 35.5 | **43.2** | 8.3 | — | 
|  | Terminus (Acc) | **25.0** | 16.3 | 6.6 | — | — | **30.3** | 16.8 | 
| Aider-Polyglot | Acc | 60.0 | 55.1 | **61.8** | 56.4 | **70.7** | 52.4 | 44.0 | 
| **Tool Use Tasks** |  |  |  |  |  |  |  |  | 
| Tau2 retail | Avg@4 | **70.6** | 69.1 | 57.0 | 75.0 | **81.8** | 74.8 | 64.3 | 
| Tau2 airline | Avg@4 | **56.5** | 39.0 | 26.5 | 55.5 | **60.0** | 54.5 | 42.5 | 
| Tau2 telecom | Avg@4 | **65.8** | 32.5 | 22.1 | 45.2 | 57.0 | 38.6 | 16.9 | 
| AceBench | Acc | **76.5** | 72.7 | 70.5 | 76.2 | 75.6 | **80.1** | 74.5 | 
| **Math & STEM Tasks** |  |  |  |  |  |  |  |  | 
| AIME 2024 | Avg@64 | **69.6** | 59.4 <sup>*</sup> | 40.1 <sup>*</sup> | 43.4 | 48.2 | 46.5 | 61.3 | 
| AIME 2025 | Avg@64 | **49.5** | 46.7 | 24.7 <sup>*</sup> | 33.1 <sup>*</sup> | 33.9 <sup>*</sup> | 37.0 | 46.6 | 
| MATH-500 | Acc | **97.4** | 94.0 <sup>*</sup> | 91.2 <sup>*</sup> | 94.0 | 94.4 | 92.4 | 95.4 | 
| HMMT 2025 | Avg@32 | **38.8** | 27.5 | 11.9 | 15.9 | 15.9 | 19.4 | 34.7 | 
| CNMO 2024 | Avg@16 | 74.3 | **74.7** | 48.6 | 60.4 | 57.6 | 56.6 | **75.0** | 
| PolyMath-en | Avg@4 | **65.1** | 59.5 | 51.9 | 52.8 | 49.8 | 54.0 | 49.9 | 
| ZebraLogic | Acc | **89.0** | 84.0 | 37.7 <sup>*</sup> | 73.7 | 59.3 | 58.5 | 57.9 | 
| AutoLogi | Acc | **89.5** | 88.9 | 83.3 | **89.8** | 86.1 | 88.2 | 84.1 | 
| GPQA-Diamond | Avg@8 | **75.1** | 68.4 <sup>*</sup> | 62.9 <sup>*</sup> | 70.0 <sup>*</sup> | 74.9 <sup>*</sup> | 66.3 | 68.2 | 
| SuperGPQA | Acc | **57.2** | 53.7 | 50.2 | 55.7 | 56.5 | 50.8 | 49.6 | 
| Humanity's Last Exam <sup>(Text Only)</sup> | - | 4.7 | 5.2 | **5.7** | 5.8 | **7.1** | 3.7 | 5.6 | 
| **General Tasks** |  |  |  |  |  |  |  |  | 
| MMLU | EM | **89.5** | 89.4 | 87.0 | 91.5 | **92.9** | 90.4 | 90.1 | 
| MMLU-Redux | EM | **92.7** | 90.5 | 89.2 | 93.6 | **94.2** | 92.4 | 90.6 | 
| MMLU-Pro | EM | 81.1 | **81.2**<sup>*</sup> | 77.3 | 83.7 | **86.6** | 81.8 | 79.4 | 
| IFEval | Prompt Strict | **89.8** | 81.1 | 83.2 <sup>*</sup> | 87.6 | 87.4 | 88.0 | 84.3 | 
| Multi-Challenge | Acc | **54.1** | 31.4 | 34.0 | 46.8 | 49.0 | 36.4 | 39.5 | 
| SimpleQA | Correct | **31.0** | 27.7 | 13.2 | 15.9 | 22.8 | **42.3** | 23.3 | 
| Livebench | Pass@1 | **76.4** | 72.4 | 67.6 | 74.8 | 74.6 | 69.8 | 67.8 | 

<sup>• Bold denotes global SOTA, and underlined denotes open-source SOTA.</sup>

<sup>• Data points marked with * are taken directly from the model's tech report or blog.</sup>

<sup>• All metrics, except for SWE-bench Verified (Agentless), are evaluated with an 8k output token length. SWE-bench Verified (Agentless) is limited to a 16k output token length.</sup>

<sup>• Kimi K2 achieves 65.8% pass@1 on the SWE-bench Verified tests with bash/editor tools (single-attempt patches, no test-time compute). It also achieves a 47.3% pass@1 on the SWE-bench Multilingual tests under the same conditions. Additionally, we report results on SWE-bench Verified tests (71.6%) that leverage parallel test-time compute by sampling multiple sequences and selecting the single best via an internal scoring model.</sup>

<sup>• To ensure the stability of the evaluation, we employed avg@k on the AIME, HMMT, CNMO, PolyMath-en, GPQA-Diamond, EvalPlus, Tau2.</sup>

<sup>• Some data points have been omitted due to prohibitively expensive evaluation costs.</sup>

| Benchmark | Metric | Shot | Kimi K2 Base | Deepseek-V3-Base | Qwen2.5-72B | Llama 4 Maverick | 
|---|---|---|---|---|---|---|
| **General Tasks** |  |  |  |  |  |  | 
| MMLU | EM | 5-shot | **87.8** | 87.1 | 86.1 | 84.9 | 
| MMLU-pro | EM | 5-shot | **69.2** | 60.6 | 62.8 | 63.5 | 
| MMLU-redux-2.0 | EM | 5-shot | **90.2** | 89.5 | 87.8 | 88.2 | 
| SimpleQA | Correct | 5-shot | **35.3** | 26.5 | 10.3 | 23.7 | 
| TriviaQA | EM | 5-shot | **85.1** | 84.1 | 76.0 | 79.3 | 
| GPQA-Diamond | Avg@8 | 5-shot | 48.1 | **50.5** | 40.8 | 49.4 | 
| SuperGPQA | EM | 5-shot | **44.7** | 39.2 | 34.2 | 38.8 | 
| **Coding Tasks** |  |  |  |  |  |  | 
| LiveCodeBench v6 | Pass@1 | 1-shot | **26.3** | 22.9 | 21.1 | 25.1 | 
| EvalPlus | Pass@1 | - | **80.3** | 65.6 | 66.0 | 65.5 | 
| **Mathematics Tasks** |  |  |  |  |  |  | 
| MATH | EM | 4-shot | **70.2** | 60.1 | 61.0 | 63.0 | 
| GSM8k | EM | 8-shot | **92.1** | 91.7 | 90.4 | 86.3 | 
| **Chinese Tasks** |  |  |  |  |  |  | 
| C-Eval | EM | 5-shot | **92.5** | 90.0 | 90.9 | 80.9 | 
| CSimpleQA | Correct | 5-shot | **77.6** | 72.1 | 50.5 | 53.5 | 

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
