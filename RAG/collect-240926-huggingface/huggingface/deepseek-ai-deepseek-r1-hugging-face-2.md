---
id: collect-240926-huggingface/huggingface/deepseek-ai-deepseek-r1-hugging-face-2
title: "deepseek-ai-deepseek-r1-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "DeepSeek"]
dates: []
keywords: ["deepseek", "apache", "benchmark", "distillation", "leaderboard", "license", "llama", "mit license", "qwen", "reasoning", "training"]
source: docs/RAG/clean_en/huggingface/deepseek-ai-deepseek-r1-hugging-face.md
source_anchor: ""
source_lines: [95, 135]
sha256: ce7dbf6904c8d3b13b77823beb5f76e10a69f6506d760504cafbdd6074a7acd2
---

# deepseek-ai-deepseek-r1-hugging-face

1. Set the temperature within the range of 0.5-0.7 (0.6 is recommended) to prevent endless repetitions or incoherent outputs.
2. **Avoid adding a system prompt; all instructions should be contained within the user prompt.**
3. For mathematical problems, it is advisable to include a directive in your prompt such as: "Please reason step by step, and put your final answer within \boxed{}."
4. When evaluating model performance, it is recommended to conduct multiple tests and average the results.

Additionally, we have observed that the DeepSeek-R1 series models tend to bypass thinking pattern (i.e., outputting "<think>\n\n</think>") when responding to certain queries, which can adversely affect the model's performance.
**To ensure that the model engages in thorough reasoning, we recommend enforcing the model to initiate its response with "<think>\n" at the beginning of every output.**

This code repository and the model weights are licensed under the MIT License. DeepSeek-R1 series support commercial use, allow for any modifications and derivative works, including, but not limited to, distillation for training other LLMs. Please note that:

- DeepSeek-R1-Distill-Qwen-1.5B, DeepSeek-R1-Distill-Qwen-7B, DeepSeek-R1-Distill-Qwen-14B and DeepSeek-R1-Distill-Qwen-32B are derived from Qwen-2.5 series, which are originally licensed under Apache 2.0 License, and now finetuned with 800k samples curated with DeepSeek-R1.
- DeepSeek-R1-Distill-Llama-8B is derived from Llama3.1-8B-Base and is originally licensed under llama3.1 license.
- DeepSeek-R1-Distill-Llama-70B is derived from Llama3.3-70B-Instruct and is originally licensed under llama3.3 license.

```
@misc{deepseekai2025deepseekr1incentivizingreasoningcapability,
      title={DeepSeek-R1: Incentivizing Reasoning Capability in LLMs via Reinforcement Learning}, 
      author={DeepSeek-AI},
      year={2025},
      eprint={2501.12948},
      archivePrefix={arXiv},
      primaryClass={cs.CL},
      url={https://arxiv.org/abs/2501.12948}, 
}
```
If you have any questions, please raise an issue or contact us at service@deepseek.com.

- Downloads last month
- 822,578

## Spaces using deepseek-ai/DeepSeek-R1 100

## Collection including deepseek-ai/DeepSeek-R1

## Paper for deepseek-ai/DeepSeek-R1

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  71.5
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  84
- LEXam-Benchmark/LEXam leaderboard
- Open Question View evaluation results source
- Mcq 4 Choices View evaluation results source
