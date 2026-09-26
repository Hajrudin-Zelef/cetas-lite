---
id: collect-240926-huggingface/huggingface/moonshotai-kimi-k2-5-hugging-face-1
title: "moonshotai-kimi-k2-5-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI"]
dates: []
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "claude", "deepseek", "gemini", "moe", "multimodal", "parameters"]
source: docs/RAG/clean_en/huggingface/moonshotai-kimi-k2-5-hugging-face.md
source_anchor: ""
source_lines: [1, 83]
sha256: db43719b3af051acdfcc89d3266788f04c27536f9e64510f179a1440e07c8055
---

# moonshotai-kimi-k2-5-hugging-face

<!-- source: https://huggingface.co/moonshotai/Kimi-K2.5 -->

- 2026.1.29: 
  - The default system prompt might cause confusion to users and unexpected behaviours, so we remove it.
  - The token `<|media_start|>` is incorrect; it has been replaced with`<|media_begin|>` in the chat template.

Kimi K2.5 is an open-source, native multimodal agentic model built through continual pretraining on approximately 15 trillion mixed visual and text tokens atop Kimi-K2-Base. It seamlessly integrates vision and language understanding with advanced agentic capabilities, instant and thinking modes, as well as conversational and agentic paradigms.

- **Native Multimodality** : Pre-trained on vision–language tokens, K2.5 excels in visual knowledge, cross-modal reasoning, and agentic tool use grounded in visual inputs.
- **Coding with Vision** : K2.5 generates code from visual specifications (UI designs, video workflows) and autonomously orchestrates tools for visual data processing.
- **Agent Swarm** : K2.5 transitions from single-agent scaling to a self-directed, coordinated swarm-like execution scheme. It decomposes complex tasks into parallel sub-tasks executed by dynamically instantiated, domain-specific agents.

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
| **Context Length** | 256K | 
| **Attention Mechanism** | MLA | 
| **Activation Function** | SwiGLU | 
| **Vision Encoder** | MoonViT | 
| **Parameters of Vision Encoder** | 400M | 

| Benchmark | <sup>Kimi K2.5 <sup>(Thinking)</sup></sup> | <sup>GPT-5.2  <sup>(xhigh)</sup></sup> | <sup>Claude 4.5 Opus  <sup>(Extended Thinking)</sup></sup> | <sup>Gemini 3 Pro  <sup>(High Thinking Level)</sup></sup> | <sup>DeepSeek V3.2  <sup>(Thinking)</sup></sup> | <sup>Qwen3-VL- 235B-A22B- Thinking</sup> |  | 
|---|---|---|---|---|---|---|---|
| **Reasoning & Knowledge** |  |  |  |  |  |  |  | 
| HLE-Full | 30.1 | 34.5 | 30.8 | 37.5 | 25.1 <sup>†</sup> | - |  | 
| HLE-Full (w/ tools) | 50.2 | 45.5 | 43.2 | 45.8 | 40.8 <sup>†</sup> | - |  | 
| AIME 2025 | 96.1 | 100 | 92.8 | 95.0 | 93.1 | - |  | 
| HMMT 2025 (Feb) | 95.4 | 99.4 | 92.9* | 97.3* | 92.5 | - |  | 
| IMO-AnswerBench | 81.8 | 86.3 | 78.5* | 83.1* | 78.3 | - |  | 
| GPQA-Diamond | 87.6 | 92.4 | 87.0 | 91.9 | 82.4 | - |  | 
| MMLU-Pro | 87.1 | 86.7* | 89.3* | 90.1 | 85.0 | - |  | 
| **Image & Video** |  |  |  |  |  |  |  | 
| MMMU-Pro | 78.5 | 79.5* | 74.0 | 81.0 | - | 69.3 |  | 
| CharXiv (RQ) | 77.5 | 82.1 | 67.2* | 81.4 | - | 66.1 |  | 
| MathVision | 84.2 | 83.0 | 77.1* | 86.1* | - | 74.6 |  | 
| MathVista (mini) | 90.1 | 82.8* | 80.2* | 89.8* | - | 85.8 |  | 
| ZeroBench | 9 | 9* | 3* | 8* | - | 4* |  | 
| ZeroBench (w/ tools) | 11 | 7* | 9* | 12* | - | 3* |  | 
| OCRBench | 92.3 | 80.7* | 86.5* | 90.3* | - | 87.5 |  | 
| OmniDocBench 1.5 | 88.8 | 85.7 | 87.7* | 88.5 | - | 82.0* |  | 
| InfoVQA (val) | 92.6 | 84* | 76.9* | 57.2* | - | 89.5 |  | 
| SimpleVQA | 71.2 | 55.8* | 69.7* | 69.7* | - | 56.8* |  | 
| WorldVQA | 46.3 | 28.0 | 36.8 | 47.4 | - | 23.5 |  | 
| VideoMMMU | 86.6 | 85.9 | 84.4* | 87.6 | - | 80.0 |  | 
| MMVU | 80.4 | 80.8* | 77.3 | 77.5 | - | 71.1 |  | 
| MotionBench | 70.4 | 64.8 | 60.3 | 70.3 | - | - |  | 
| VideoMME | 87.4 | 86.0* | - | 88.4* | - | 79.0 |  | 
| LongVideoBench | 79.8 | 76.5* | 67.2* | 77.7* | - | 65.6* |  | 
| LVBench | 75.9 | - | - | 73.5* | - | 63.6 |  | 
| **Coding** |  |  |  |  |  |  |  | 
| SWE-Bench Verified | 76.8 | 80.0 | 80.9 | 76.2 | 73.1 | - |  | 
| SWE-Bench Pro | 50.7 | 55.6 | 55.4* | - | - | - |  | 
| SWE-Bench Multilingual | 73.0 | 72.0 | 77.5 | 65.0 | 70.2 | - |  | 
| Terminal Bench 2.0 | 50.8 | 54.0 | 59.3 | 54.2 | 46.4 | - |  | 
| PaperBench | 63.5 | 63.7* | 72.9* | - | 47.1 | - |  | 
| CyberGym | 41.3 | - | 50.6 | 39.9* | 17.3* | - |  | 
| SciCode | 48.7 | 52.1 | 49.5 | 56.1 | 38.9 | - |  | 
| OJBench (cpp) | 57.4 | - | 54.6* | 68.5* | 54.7* | - |  | 
| LiveCodeBench (v6) | 85.0 | - | 82.2* | 87.4* | 83.3 | - |  | 
| **Long Context** |  |  |  |  |  |  |  | 
| Longbench v2 | 61.0 | 54.5* | 64.4* | 68.2* | 59.8* | - |  | 
| AA-LCR | 70.0 | 72.3* | 71.3* | 65.3* | 64.3* | - |  | 
| **Agentic Search** |  |  |  |  |  |  |  | 
| BrowseComp | 60.6 | 65.8 | 37.0 | 37.8 | 51.4 | - |  | 
| BrowseComp (w/ctx manage) | 74.9 |  | 57.8 | 59.2 | 67.6 | - |  | 
| BrowseComp (Agent Swarm) | 78.4 | - | - | - | - | - |  | 
| WideSearch (item-f1) | 72.7 | - | 76.2* | 57.0 | 32.5* | - |  | 
| WideSearch (item-f1 Agent Swarm) | 79.0 | - | - | - | - | - |  | 
| DeepSearchQA | 77.1 | 71.3* | 76.1* | 63.2* | 60.9* | - |  | 
| FinSearchCompT2&T3 | 67.8 | - | 66.2* | 49.9 | 59.1* | - |  | 
| Seal-0 | 57.4 | 45.0 | 47.7* | 45.5* | 49.5* | - |  | 

## **Footnotes**

