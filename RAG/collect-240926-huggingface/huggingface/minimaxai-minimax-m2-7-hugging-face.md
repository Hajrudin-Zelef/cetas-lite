---
id: collect-240926-huggingface/huggingface/minimaxai-minimax-m2-7-hugging-face
title: "minimaxai-minimax-m2-7-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "MiniMax", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "benchmark", "benchmarks", "cost", "incident", "inference", "leaderboard", "memory", "nvidia", "open-weight", "opus 4", "parameters"]
source: docs/RAG/clean_en/huggingface/minimaxai-minimax-m2-7-hugging-face.md
source_anchor: ""
source_lines: [1, 59]
sha256: 500924cfbea8e84010f9b69b35a14c29464830b3369a89b5743f18711af2cac8
---

# minimaxai-minimax-m2-7-hugging-face

<!-- source: https://huggingface.co/MiniMaxAI/MiniMax-M2.7 -->

**MiniMax-M2.7** is our first model deeply participating in its own evolution. M2.7 is capable of building complex agent harnesses and completing highly elaborate productivity tasks, leveraging Agent Teams, complex Skills, and dynamic tool search. For more details, see our blog post.


M2.7 initiates a cycle of model self-evolution: during development, we let the model update its own memory, build dozens of complex skills for RL experiments, and improve its own learning process based on experiment results. An internal version of M2.7 autonomously optimized a programming scaffold over 100+ rounds — analyzing failure trajectories, modifying code, running evaluations, and deciding to keep or revert — achieving a **30% performance improvement**. On MLE Bench Lite (22 ML competitions), M2.7 achieved a **66.6% medal rate**, second only to Opus-4.6 and GPT-5.4.



M2.7 delivers outstanding real-world programming capabilities spanning log analysis, bug troubleshooting, refactoring, code security, and machine learning. Beyond code generation, M2.7 demonstrates strong system-level reasoning — correlating monitoring metrics, conducting trace analysis, verifying root causes in databases, and making SRE-level decisions. Using M2.7, we have reduced live production incident recovery time to **under three minutes** on multiple occasions.

On SWE-Pro, M2.7 achieved **56.22%**, matching GPT-5.3-Codex, with even stronger performance on real-world engineering benchmarks: **SWE Multilingual (76.5)** and **Multi SWE Bench (52.7)**. On **VIBE-Pro (55.6%)**, M2.7 is nearly on par with Opus 4.6. On **Terminal Bench 2 (57.0%)** and **NL2Repo (39.8%)**, M2.7 demonstrates deep understanding of complex engineering systems. M2.7 also supports native **Agent Teams** for multi-agent collaboration with stable role identity and autonomous decision-making.


M2.7 achieved an **ELO score of 1495** on GDPval-AA (highest among open-weight models), surpassing GPT5.3. It handles Word, Excel, and PPT with high-fidelity multi-round editing, producing editable deliverables. On Toolathon, M2.7 reached **46.3%** accuracy (global top tier), and maintains **97% skill compliance** across 40+ complex skills on MM Claw. On the MM Claw end-to-end benchmark, M2.7 achieved **62.7%**, close to Sonnet 4.6.

M2.7 features strengthened character consistency and emotional intelligence. We open-sourced OpenRoom, an interactive demo that places AI interaction within a Web GUI space with real-time visual feedback and scene interactions. Try it at openroom.ai.

- MiniMax Agent: https://agent.minimax.io/
- MiniMax API: https://platform.minimax.io/
- Token Plan: https://platform.minimax.io/subscribe/token-plan

Download the model from HuggingFace repository: https://huggingface.co/MiniMaxAI/MiniMax-M2.7

We recommend using the following inference frameworks (listed alphabetically) to serve the model:

We recommend using SGLang to serve MiniMax-M2.7. Please refer to our SGLang Deployment Guide.

We recommend using vLLM to serve MiniMax-M2.7. Please refer to our vLLM Deployment Guide.

We recommend using Transformers to serve MiniMax-M2.7. Please refer to our Transformers Deployment Guide.

You also can get model weights from modelscope.

MiniMax M2.7 is also available on NVIDIA NIM Endpoint.

We recommend using the following parameters for best performance: `temperature=1.0`, `top_p = 0.95`, `top_k = 40`. Default system prompt:

```
You are a helpful assistant. Your name is MiniMax-M2.7 and is built by MiniMax.
```
Please refer to our Tool Calling Guide.

Contact us at model@minimax.io.

- Downloads last month
- 1,326,813

## Spaces using MiniMaxAI/MiniMax-M2.7 100

## Collection including MiniMaxAI/MiniMax-M2.7

- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results    leaderboard  56.2
- internlm/WildClawBench leaderboard
- Overall View evaluation resultssource33.8
- Avg Time View evaluation resultssource
- Avg Cost View evaluation resultssource7.2
- SWE-bench/SWE-bench_Multilingual · Swe Bench Multilingual Resolved View evaluation results    leaderboard  76.5
- benchflow/skillsbench · Skillsbench V1 1 View evaluation results  source leaderboard34.9<sup>*</sup>
