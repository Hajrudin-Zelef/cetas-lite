---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face-10
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Moonshot", "Nvidia"]
dates: []
keywords: ["nvfp4", "nvidia", "agent", "agentic", "alignment", "benchmark", "benchmarks", "blackwell", "deepseek", "guardrails", "kimi", "leaderboard"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [769, 828]
sha256: 90c986e668830480efc21370b823cb3fc680a8af288433ac2d719193d6f28e16
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face

| Synthetic Safety from gemma-3-4b-it, Nemotron-Nano-9B-v2, and gpt-oss-120b | Text | 44,091 | [Safety SFT Data] | [google/gemma-3-4b-it]; [Nemotron-Nano-9B-v2]; [gpt-oss-120b] | 
| Synthetic Safety from DeepSeek-R1-0528, gpt-oss-120b, DeepSeek-R1-Distill-Qwen-7B, and Mixtral-8x7B-v0.1 | Text | Undisclosed | [Nemotron Content Safety Dataset V2]; [Gretel Synthetic Safety Alignment Dataset]; [RedTeam-2K]; [Malicious Tasks]; [Nemotron-Personas-USA] | [DeepSeek-R1-0528]; [gpt-oss-120b]; [DeepSeek-R1-Distill-Qwen-7B]; [Qwen3-30B-A3B-Thinking-2507]; [Qwen3-235B-A22B-Instruct-2507]; [Mixtral-8x7B-v0.1] | 
| Synthetic Tool Calling from Qwen3-235B-A22B-Thinking-2507 and Qwen3-Next-80B-A3B-Thinking | Text | Undisclosed | [ToolBench]; [glaive-function-calling-v2]; [APIGen Function-Calling]; [Nemotron-Personas-USA] | [Qwen3-235B-A22B-Thinking-2507]; [Qwen3-Next-80B-A3B-Thinking] | 
| Synthetic Chat from gpt-oss-120b, Mixtral-8x22B-Instruct-v0.1, Qwen3-235B-A22B-Instruct-2507, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | [C4]; [LMSYS-Chat-1M]; [ShareGPT]; [GSM8K]; [PRM800K]; [FinQA]; [WikiTableQuestions]; [Riddles]; [glaive-function-calling-v2]; [SciBench]; [tigerbot-kaggle-leetcodesolutions-en-2k]; [OpenBookQA]; [Advanced Reasoning Benchmark]; Software Heritage; [Khan Academy Math Keywords]; [WildChat-1M]; [Nemotron-Personas-USA] | [gpt-oss-120b]; [Mixtral-8x22B-Instruct-v0.1]; [Qwen3-235B-A22B-Instruct-2507]; [Qwen3-235B-A22B-Thinking-2507] | 
| Synthetic Tool Use Interactive Agent from gpt-oss-120b, DeepSeek-R1-0528, Qwen3-32B, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | NVIDIA Internal | [gpt-oss-120b]; [DeepSeek-R1-0528]; [Qwen3-32B]; [Qwen3-235B-A22B-Thinking-2507] | 
| Synthetic DocFinQA and SWE-smith from Qwen3-Coder-480B-A35B-Instruct and Kimi-K2-Thinking | Text | Undisclosed | [DocFinQA]; [SWE-smith] | [Qwen3-Coder-480B-A35B-Instruct]; [Kimi-K2-Thinking] | 
| Synthetic SWE-Gym from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | [SWE-Gym] | [Qwen3-Coder-480B-A35B-Instruct] | 
| Synthetic SWE-Gym and R2E-Gym-Subset from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | [SWE-Gym]; [R2E-Gym-Subset] | [Qwen3-Coder-480B-A35B-Instruct] | 
| Synthetic SWE-Gym and R2E-Gym-Subset from DeepSeek-R1-0528 | Text | Undisclosed | [SWE-Gym]; [R2E-Gym-Subset] | [DeepSeek-R1-0528] | 
| Synthetic HelpSteer, LMSYS-Chat-1M, and Nemotron-Personas-USA from gpt-oss-120b, Qwen3-235B-A22B-Instruct-2507, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | [HelpSteer2]; [HelpSteer3]; [LMSYS-Chat-1M]; [Nemotron-Personas-USA] | [gpt-oss-120b]; [Qwen3-235B-A22B-Instruct-2507]; [Qwen3-235B-A22B-Thinking-2507] | 
| Synthetic Nemotron-Personas-USA from gpt-oss-120b and Qwen3-8B | Text | Undisclosed | [Nemotron-Personas-USA] | [gpt-oss-120b]; [Qwen3-8B] | 
| Vendor Terminal Bench-like Tasks (Droid) | Text | Undisclosed | [Droid Harness Pivot vendor data] | Undisclosed | 

For our post-training recipe, we focused on the following languages in addition to English: French, German, Italian, Japanese, Spanish, and Chinese. Those languages were represented in the form of multilingual reasoning and translation tasks.

**Data Collection Method by dataset** 

- Hybrid: Automated, Manually-Collected, Synthetic
**Labeling Method by dataset**
- Hybrid: Automated, Manually-Labeled, Synthetic
**Properties:** This corpus comprises a mix of high-quality standard benchmarks and test suites for modern agentic AI. These benchmarks test model capabilities on tasks such as tool-calling and instruction following.

**Data Collection Method by dataset** 

- Hybrid: Automated, Manually-Collected, Synthetic
**Labeling Method by dataset**
- Hybrid: Automated, Manually-Labeled, Synthetic
**Properties:** This corpus comprises a mix of high-quality standard benchmarks and test suites for modern agentic AI. These benchmarks test model capabilities on tasks such as tool-calling and instruction following.

- **Acceleration Engine:** PyTorch
- **Test Hardware:**
  - NVIDIA Hopper
    - 1-8x H100
    - 1-8x H200
  - NVIDIA Blackwell
    - GB200
    - DGX Spark (GB10)
    - GeForce RTX 5090
- NVIDIA Hopper

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. Developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case. For more details: Safety and Explainability Subcards.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, and Privacy Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 951,877

## Spaces using nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 3

## Collections including nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  75.57
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results    leaderboard  52.8
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  81.62
- cais/hle · Hle View evaluation results     10.47<sup>*</sup>
- SWE-bench/SWE-bench_Multilingual · Swe Bench Multilingual Resolved View evaluation results    leaderboard  36.47
