---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face-7
title: "with uv: uv pip install vllm==0.18.1 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "MiniMax", "Moonshot", "Nvidia"]
dates: []
keywords: ["agentic", "blackwell", "deepseek", "distribution", "guardrails", "kimi", "leaderboard", "nvidia", "reasoning", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face.md
source_anchor: ""
source_lines: [637, 713]
sha256: de535887cafbf6b63543786c072255fbed84715decce917cc16f443e5dbe3ad3
---

# with uv: uv pip install vllm==0.18.1 --torch-backend=auto

| Synthetic GPQA from Deepseek-V3 | Text | Undisclosed | Stack Exchange | DeepSeek-V3-0324 | 
| Synthetic Vedantu from Deepseek-V3 | Text | Undisclosed | Vedantu | DeepSeek-V3-0324 | 
| Synthetic Tool Call Schema for RL | Text | Undisclosed | ToolBench; glaive-function-calling-v2; APIGen Function-Calling; Nemotron-Personas-USA | Qwen3-235B-A22B-Thinking-2507; Qwen3-Next-80B-A3B-Thinking | 
| Synthetic Data for Search | Text | Undisclosed | Wikimedia | MiniMax-M2 | 
| Synthetic Instruction Following for RL | Text | Undisclosed | - | NVIDIA-Nemotron-Nano-9B-v2; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Conversational Agentic Tool-Use RL | Text | Undisclosed | - | DeepSeek-V3.2; DeepSeek-R1-0528; Qwen3-235B-A22B-Thinking-2507; Qwen3-32B; gpt-oss-120b; Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic Terminal Pivot RL | Text | Undisclosed | SWE-smith; Nemotron-Cascade-RL-SWE; Vendor supplied | DeepSeek-V3.2; Qwen3-Coder-480B-A35B-Instruct; Kimi-K2.5; Qwen3-235B-A22B-Instruct-2507 | 

For our post-training recipe, we focused on 9 main languages in addition to English: French, German, Italian, Japanese, Spanish, and Chinese

Those languages were represented in the form of multilingual reasoning and translation tasks.

The following table depicts our sample distribution for the 6 languages and 5 translation pairs.

| Language | Size | 
|---|---|
| English | 13.48M | 
| Italian | 53k | 
| German | 53k | 
| Spanish | 53k | 
| French | 53k | 
| Japanese | 53k | 
| Chinese | 53k | 
| English <-> Italian | 43.2k | 
| English <-> German | 43.2k | 
| English <-> Spanish | 43.2k | 
| English <-> French | 43.2k | 
| English <-> Japanese | 43.2k | 

- **Data Collection Method by dataset** : Hybrid: Human, Synthetic
- **Labeling Method by dataset** : Hybrid: Automated, Human, Synthetic

- **Acceleration Engine:** PyTorch
- **Test Hardware:**
  - NVIDIA Hopper
    - 1-8x H100
    - 1-8x H200
  - NVIDIA Grace Blackwell
    - GB200
- NVIDIA Hopper

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case. For more details: Safety and Explainability Subcards.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, and Privacy Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

```
@misc{nvidia_nemotron_3_2025,
  title  = {NVIDIA Nemotron 3: Efficient and Open Intelligence},
  author = {{NVIDIA}},
  year   = {2025},
  url    = {https://arxiv.org/abs/2512.20856},
  note   = {White Paper}
}
```
- Downloads last month
- 1,229,837

## Spaces using nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16 10

## Collection including nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16

## Papers for nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16

## Article mentioning nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16

- Idavidrein/gpqa · Diamond leaderboard
- default View evaluation results79.23
- With tools View evaluation results82.7<sup>*</sup>
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved leaderboard
- OpenHands harness View evaluation results60.47<sup>*</sup>
- OpenCode harness View evaluation results59.2<sup>*</sup>
- Codex harness View evaluation results53.73<sup>*</sup>
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  83.73
