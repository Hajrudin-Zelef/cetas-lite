---
id: collect-240926-huggingface/huggingface/minimaxai-minimax-m2-5-hugging-face-2
title: "minimaxai-minimax-m2-5-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Google", "Hugging Face", "MiniMax", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "benchmark", "claude", "gemini", "inference", "inference engine", "moe", "opus 4", "parameters", "research"]
source: docs/RAG/clean_en/huggingface/minimaxai-minimax-m2-5-hugging-face.md
source_anchor: ""
source_lines: [47, 111]
sha256: 42eb61e09a28bf55745da0d6844beef114d09c11813c7b5712660c38821678d1
---

# minimaxai-minimax-m2-5-hugging-face

One of the key drivers of the aforementioned developments is the scaling of reinforcement learning. As we train our models, we also benefit from their abilities. Most of the tasks and workspaces that we perform in our company have been made into training environments for RL. To date, there are already hundreds of thousands of such environments. At the same time, we did plenty of work on our agentic RL framework, algorithms, reward signals, and infrastructure engineering to support the continued scaling of our RL training.

We designed an agent-native RL framework in-house, called Forge, which introduces an intermediary layer that fully decouples the underlying training-inference engine from the agent, supporting the integration of arbitrary agents and enabling us to optimize the model's generalization across agent scaffolds and tools. To improve system throughput, we optimized asynchronous scheduling strategies to balance system throughput against sample off-policyness, and designed a tree-structured merging strategy for training samples, achieving approximately 40x training speedup.


On the algorithm side, we continued using the CISPO algorithm we proposed at the beginning of last year to ensure the stability of MoE models during large-scale training. To address the credit assignment challenge posed by long contexts in agent rollouts, we introduced a process reward mechanism for end-to-end monitoring of generation quality. Furthermore, to deeply align with user experience, we evaluated task completion time through agent trajectories, achieving an optimal trade-off between model intelligence and response speed.


We will release a more comprehensive introduction to RL scaling soon in a separate technical blogpost.

M2.5 has been fully deployed in MiniMax Agent, delivering the best agentic experience.

We have distilled core information-processing capabilities into standardized Office Skills deeply integrated within MiniMax Agent. In MAX mode, when handling tasks such as Word formatting, PowerPoint editing, and Excel calculations, MiniMax Agent automatically loads the corresponding Office Skills based on file type, improving the quality of task outputs.

Furthermore, users can combine Office Skills with domain-specific industry expertise to create reusable Experts tailored to specific task scenarios.

Take industry research as an example: by merging a mature research framework SOP (standard operating procedure) with Word Skills, the Agent can strictly follow the established framework to automatically fetch data, organize analytical logic, and output properly formatted research reports — rather than merely generating a raw block of text. In financial modeling scenarios, by combining an organization's proprietary modeling standards with Excel Skills, the Agent can follow specific risk control logic and calculation standards to automatically generate and validate complex financial models, rather than simply outputting a basic spreadsheet.

To date, users have built over 10,000 Experts on MiniMax Agent, and this number is still growing rapidly. MiniMax has also built multiple sets of deeply optimized, ready-to-use Expert suites on MiniMax Agent for high-frequency scenarios such as office work, finance, and programming.

MiniMax itself has been among the first to benefit from M2.5's capabilities. Throughout the company's daily operations, 30% of overall tasks are autonomously completed by M2.5, spanning functions including R&D, product, sales, HR, and finance — and the penetration rate continues to rise. Performance in coding scenarios has been particularly notable, with M2.5-generated code accounting for 80% of newly committed code.

MiniMax Agent: https://agent.minimax.io/

MiniMax API Platform: https://platform.minimax.io/

MiniMax Coding Plan: https://platform.minimax.io/subscribe/coding-plan

Download the model from HuggingFace repository: https://huggingface.co/MiniMaxAI/MiniMax-M2.5

We recommend using the following inference frameworks (listed alphabetically) to serve the model:

We recommend using SGLang to serve MiniMax-M2.5. Please refer to our SGLang Deployment Guide.

We recommend using vLLM to serve MiniMax-M2.5. Please refer to our vLLM Deployment Guide.

We recommend using Transformers to serve MiniMax-M2.5. Please refer to our Transformers Deployment Guide.

We recommend using KTransformers to serve MiniMax-M2.5. Please refer to KTransformers Deployment Guide

You also can get model weights from modelscope.

We recommend using the following parameters for best performance: `temperature=1.0`, `top_p = 0.95`, `top_k = 40`. Default system prompt:

```
You are a helpful assistant. Your name is MiniMax-M2.5 and is built by MiniMax.
```
Please refer to our Tool Calling Guide.

Contact us at model@minimax.io.

Further benchmark results of M2.5:

| Benchmark | MiniMax-M2.5 | MiniMax-M2.1 | Claude Sonnet 4.5 | Claude Opus 4.5 | Claude Opus 4.6 | Gemini 3 Pro | GPT-5.2 (thinking) | 
|---|---|---|---|---|---|---|---|
| AIME25 | 86.3 | 83.0 | 88.0 | 91.0 | 95.6 | 96.0 | 98.0 | 
| GPQA-D | 85.2 | 83.0 | 83.0 | 87.0 | 90.0 | 91.0 | 90.0 | 
| HLE w/o tools | 19.4 | 22.2 | 17.3 | 28.4 | 30.7 | 37.2 | 31.4 | 
| SciCode | 44.4 | 41.0 | 45.0 | 50.0 | 52.0 | 56.0 | 52.0 | 
| IFBench | 70.0 | 70.0 | 57.0 | 58.0 | 53.0 | 70.0 | 75.0 | 
| AA-LCR | 69.5 | 62.0 | 66.0 | 74.0 | 71.0 | 71.0 | 73.0 | 

Evaluation methods:


