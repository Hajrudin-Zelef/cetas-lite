---
id: collect-240926-huggingface/huggingface/unsloth-glm-5-3-gguf-hugging-face-1
title: "Read our How to Run GLM-5.3 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "ExploitGym", "Moonshot", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["glm", "agents", "benchmark", "benchmarks", "cyber", "deepseek", "fable 5", "gguf", "gpt-5.6", "kimi", "leaderboard", "opus 4"]
source: docs/RAG/clean_en/huggingface/unsloth-glm-5-3-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 37]
sha256: dc0dae56abd1a3f1f3fa63d3f0d3e4be3bab4d9a8e63d8bcda697b9a52a4dc26
---

# Read our How to Run GLM-5.3 Guide!

<!-- source: https://huggingface.co/unsloth/GLM-5.3-GGUF -->

# Read our How to Run GLM-5.3 Guide!

    *See Unsloth Dynamic 3.0 GGUFs for our quantization benchmarks.*
  

- You can now run GLM-5.3 in Unsloth Desktop with toggles for Low, High and Max thinking.
- Read our GLM-5.3 guide for analysis and instructions.

GLM-5.3 uses the same base model as GLM-5.2 — every gain comes from post-training. Compared with GLM-5.2, it is much better at complex coding and long-horizon tasks:

- Stronger Coding: GLM-5.3 is the most capable open-weights model for coding, with a 50% improvement over GLM-5.2 on our in-house Z.ai Code Bench. It also achieve open-source SOTA on public benchmarks including Terminal Bench 3.0 and Agents' Last Exam.
- Emergent Cyber Capability: As we scaled post-training, cyber capability developed faster than we expected. GLM-5.3 is state of the art on CyberGym for vulnerability discovery, and its gains are largest further up the exploitation chain, where it more than doubles GLM-5.2 on exploitation benchmarks.

| Benchmark | GLM-5.3 | GLM-5.2 | Kimi K3 | DeepSeek-V4 Pro-0813 | Qwen3.8-Max | Opus 4.8 | Fable 5 (w/ fallback) | GPT-5.6 Sol | 
|---|---|---|---|---|---|---|---|---|
| Terminal Bench 2.1 | 88.2 | 81.0 | 88.3 | 87.9 | 86.6 | 85.0 | 88.0 | **88.8** | 
| Terminal Bench 3.0 | 28.3 | 4.6 | 17.4 | – | – | 21.1 | 33.7 | **34.6** | 
| DeepSWE (v1.1) | 66.9 | 46.2 | 67.5 | 62.7 | 56.6 | 58.0 | 69.7 | **72.7** | 
| NL2Repo | 58.0 | 48.9 | 58.0 | 61.1 | 55.9 | **69.7** | – | – | 
| ProgramBench (Almost Solved) | 19.0 | 9.5 | 17.5 | – | 10.5 | 15.5 | **33.0** | 23.0 | 
| FrontierSWE | 78.1 | 67.5 | – | – | – | 66.5 | **88.2** | – | 
| SWE-Marathon (v1.1) | 42.5 | 19.4 | 48.1 | – | – | **48.8** | 33.1 | 42.5 | 
| PostTrainBench | 39.8 | 31.7 | 32.0 | – | – | 32.9 | **41.8** | 36.2 | 
| CyberGym | **84.5** | 77.2 | 80.0 | 83.3 | 78.5 | 78.1 | 83.8 | 83.6 | 
| ExploitGym (2h / 6h) | 105 / 130 | 29 / 39 | 36 / 70 | – | 14 / 26 | 80 / 120 | 181 / 247 | **216 / 293** | 
| ExploitBench | 54.4 | 24.4 | 32.2 | – | 28.8 | 40.0 | **78.0** | 76.5 | 
| Toolathlon Verified | 73.0 | 59.9 | **76.5** | 74.1 | 72.5 | 76.2 | 74.7 | 74.9 | 
| AutomationBench (v1.0.6) | **48.2** | 26.2 | 46.7 | 43.2 | 39.8 | 41.0 | 46.2 | 45.8 | 
| Agents' Last Exam (ALE-CLI) | 28.5 | 23.8 | 27.6 | 25.7 | 27.0 | 25.7 | 23.8 | **28.6** | 
| HLE w/ Tools | 62.5 | 54.7 | 59.8 | 60.0 | 56.2 | 57.9 | 63.9 | **64.5** | 
| GDPval-AA v2 | **1769** | 1508 | 1682 | 1590 | 1739 | 1588 | 1743 | 1730 | 

- GLM-5.3 supports controlling the thinking budget through the `reasoning_effort` parameter, which accepts three levels:`low` ,`high` , and`max` . It defaults to`max` if not passed (or if set to any other value). To use`low` or`high` , pass them explicitly. For benchmark and leaderboard reproduction, keep the default`max` .
- In the chat template for GLM-5.3, `clear_thinking` defaults to`false` if not passed. For chat scenarios, explicitly pass`clear_thinking=true` .

