---
id: collect-240926-huggingface/huggingface/qwen-qwen3-8-flash-next-hugging-face-2
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "SGLang", "vLLM"]
dates: ["2026-03-05"]
keywords: ["agent", "agentic", "benchmark", "claude", "inference", "latency", "multimodal", "parameters", "qwen", "reasoning", "sglang", "throughput"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-8-flash-next-hugging-face.md
source_anchor: ""
source_lines: [95, 150]
sha256: fb5a7c477c2efa54410129466929650808636555c900bc34ae18d7b7264fad73
---

# Set the following accordingly

|  | Qwen3.8-Flash-Next | Qwen3.8-27B | Qwen3.7-Plus | Claude-Opus-4.6 (Max) | 
|---|---|---|---|---|
| Agentic Multimodal Intelligence |  |  |  |  | 
| Multimodal tool use ClawEval-MM | Pass@3 **64.4**Average **60.4** | Pass@3 57.4 Average 56.9 | Pass@3 57.4 Average 60.1 | Pass@3 52.5 Average 54.7 | 
| Application recreation RecreationBench | **49.9** | 47.1 | 30.2 | -- | 
| Mobile use AndroidWorld | **84.5** | 81.9 | 81.0 | 62.0 | 
| Computer use OSWorld 2.0 | Binary **19.4**Partial **52.3** | Binary 19.4 Partial 48.0 | Binary 2.8 Partial 21.5 | -- | 
| Visual web development Vision2Web | **64.0** | 62.9 | 42.1 | -- | 
| General Multimodal Intelligence |  |  |  |  | 
| Embodied intelligence ERQA | **72.3** | 65.5 | 69.8 | 40.8 | 
| Long video understanding LVBench | **76.6** | 72.4 | 76.2 | 63.0 | 
| Real-world perception RealWorldQA | **88.5** | 85.9 | 86.9 | 73.9 | 
| Visual math problem solving MathVision | Without CI **90.6**With CI **95.7** | Without CI 90.0 With CI 94.6 | Without CI 90.3 With CI 88.7 | Without CI 65.5 | 
| Scientific chart analysis CharXiv (RQ) | Without CI 84.6 With CI **90.6** | Without CI 83.7 With CI 90.2 | Without CI **85.8**With CI 85.9 | Without CI 66.0 | 

1. ClawEval-MM: scores are reported as "pass@3 / average score". Pass@3 measures the percentage passed in at least one of three trials, and the average score is the mean score across the three trials.

2. RecreationBench: an in-house long-horizon application-recreation benchmark for evaluating hybrid-agent abilities spanning five platforms — desktop (Ubuntu, macOS, Windows), mobile (Android) and web.

3. OSWorld 2.0: scores are reported as "binary / partial". The binary score is the percentage of tasks that receive the full task reward, while the partial score aggregates the partial rewards obtained across all tasks.

4. Vision2Web: scores are reported as the average over the frontend, webpage and website categories, using the Claude Code harness and judged by gpt-5.4-2026-03-05.

5. MathVision, CharXiv (RQ): scores are reported as "without CI / with CI". A small number of incorrect ground-truth annotations in MathVision were corrected after manual verification. Our model's score is evaluated using a fixed prompt, e.g. "Please reason step by step, and put your final answer within \boxed{}." For other models, we report the higher score between runs with and without the \boxed{} formatting.

6. The best result in each row is shown in bold.

7. Empty cells (--) indicate scores not yet available or not applicable.

For streamlined integration, we recommend using Qwen3.8-Flash-Next via APIs.

Inference efficiency and throughput vary significantly across frameworks. We recommend using the latest framework versions to ensure optimal performance and compatibility. For production workloads or high-throughput scenarios, dedicated serving engines such as SGLang, KTransformers or vLLM are strongly recommended.


Qwen3.8-Flash-Next can be deployed with popular inference frameworks, e.g.:

Qwen3.8-Flash-Next models operate in thinking mode by default, generating thinking content signified by `<think>\n...</think>\n\n` before producing the final responses.
To disable thinking content and obtain direct response, refer to the examples here.


We recommend using the following sets of sampling parameters for generation:


- Thinking Mode:
`temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0`- Instruct (or non-thinking) mode:
`temperature=0.7`, `top_p=0.80`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0`
Please note that the support for sampling parameters varies according to inference frameworks.


In multi-turn agentic tasks, lower reasoning effort does not always reduce overall task completion time. Although it may produce faster per-turn responses, it can also lead to insufficient analysis, more failures, and repeated retries, which may increase total latency and token consumption.


Qwen3.8-Flash-Next supports controlling thinking behavior via `enable_thinking`, `preserve_thinking`, and `reasoning_effort`.

The Chat Completions API can be used with most inference frameworks, as well as Qwen Cloud. Before starting, make sure it is installed and the API key and the API base URL is configured, e.g.:

