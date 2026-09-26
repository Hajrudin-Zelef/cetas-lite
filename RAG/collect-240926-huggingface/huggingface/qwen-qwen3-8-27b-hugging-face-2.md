---
id: collect-240926-huggingface/huggingface/qwen-qwen3-8-27b-hugging-face-2
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-03-05"]
keywords: ["agent", "agentic", "benchmark", "benchmarks", "claude", "cost", "inference", "latency", "multimodal", "opus 4", "parameters", "qwen"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-8-27b-hugging-face.md
source_anchor: ""
source_lines: [90, 131]
sha256: 76c305ccc4a03ccf048a441523fa02a18a42576f867a7c591215702f0115418c
---

# Set the following accordingly

1. MathVision, BabyVision, and CharXiv (RQ): Where both settings are available, cells report “Without CI” and “With CI” separately; otherwise, only the available setting is shown. A small number of incorrect ground-truth annotations in MathVision and CharXiv (RQ) were corrected following manual verification, and all reported scores on those benchmarks were computed using the corrected annotations.
2. MathVision: Qwen3.8-27B is evaluated using the fixed prompt: “Please reason step by step, and put your final answer within `\boxed{}` .” For the remaining models, we report the higher score from two prompt variants—one with and one without the`\boxed{}` formatting requirement.
3. WebArena-Verified: Scores are computed with the official WebArena-Verified grader under the OSWorld scaffold.
4. RecreationBench: An in-house, long-horizon application-recreation benchmark designed to evaluate hybrid-agent capabilities across five platforms: desktop (Ubuntu, macOS, and Windows), mobile (Android), and the web.
5. ClawEval-MM: Scores are reported as “Pass@3 / average score.” Pass@3 is the percentage of tasks passed in at least one of three trials; the average score is the mean benchmark score across the three trials.
6. Vision2Web: Scores are averaged across the frontend, webpage, and website categories. Evaluations use the Claude Code harness and are judged by `gpt-5.4-2026-03-05` .
7. SWE-MM: Scores are evaluated on the Claude Code harness using the public dev split of SWE-bench Multimodal, with the modifications described in Appendix 8.3 of the Claude Opus 4.7 system card.
8. Empty cells (--) indicate that results are not yet available or not applicable.

For streamlined integration, we recommend using Qwen3.8 via APIs.

Inference efficiency and throughput vary significantly across frameworks. We recommend using the latest framework versions to ensure optimal performance and compatibility. For production workloads or high-throughput scenarios, dedicated serving engines such as SGLang, vLLM, or TokenSpeed are recommended.


Qwen3.8 can be deployed with popular inference frameworks, e.g.:

Qwen3.8 models operate in thinking mode by default, generating thinking content signified by `<think>\n...</think>\n\n` before producing the final response.
To disable thinking content and obtain a direct response, refer to the examples here.


We recommend using the following sets of sampling parameters for generation:


- Thinking Mode:
`temperature=1.0`, `top_p=0.95`, `top_k=20`, `min_p=0.0`, `presence_penalty=0.0`, `repetition_penalty=1.0`- Instruct (or non-thinking) mode:
`temperature=0.7`, `top_p=0.80`, `top_k=20`, `min_p=0.0`, `presence_penalty=1.5`, `repetition_penalty=1.0`
Please note that the support for sampling parameters varies according to inference frameworks.


Qwen3.8 comes with official support for `reasoning_effort`, which can be used to adjust reasoning depth and control cost:  

- `xhigh` (default): for complex tasks demanding thorough analysis
- `medium` : balancing accuracy and speed
- `low` : efficient reasoning optimizing for speed and cost

In addition, `preserve_thinking` is enabled by default for all workloads for the best out-of-the-box experience. To disable preserved thinking, refer to the examples here.

In multi-turn agentic tasks, lower reasoning effort does not always reduce overall task completion time. Although it may produce faster per-turn responses, it can also lead to insufficient analysis, more failures, and repeated retries, which may increase total latency and token consumption.


The Chat Completions API can be used with most inference frameworks, as well as Qwen Cloud. Before starting, make sure the OpenAI Python SDK is installed and the API key and the API base URL are configured, e.g.:

