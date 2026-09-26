---
id: collect-240926-huggingface/huggingface/unsloth-kimi-k2-6-gguf-hugging-face-1
title: "Read our How to Run Kimi K2.6 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Moonshot", "OpenAI", "Unsloth"]
dates: []
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "context window", "deepseek", "gemini", "gguf"]
source: docs/RAG/clean_en/huggingface/unsloth-kimi-k2-6-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 104]
sha256: 8e417b16fd66671d3ecc0461d23481564331bfcb3fe6616bc34889f8ff02bfbd
---

# Read our How to Run Kimi K2.6 Guide!

<!-- source: https://huggingface.co/unsloth/Kimi-K2.6-GGUF -->

# Read our How to Run Kimi K2.6 Guide!

    *See Unsloth Dynamic 2.0 GGUFs for our quantization benchmarks.*
  

- To run Kimi K2.6 in full precision lossless, run Q8 (UD-Q8_K_XL), which is 595GB and only 10GB bigger than Q4 (UD-Q4_K_XL).
- See our Kimi K2.6 guide for quantization analysis and instructions.

Kimi K2.6 is an open-source, native multimodal agentic model that advances practical capabilities in long-horizon coding, coding-driven design, proactive autonomous execution, and swarm-based task orchestration.

- **Long-Horizon Coding** : K2.6 achieves significant improvements on complex, end-to-end coding tasks, generalizing robustly across programming languages (Rust, Go, Python) and domains spanning front-end, DevOps, and performance optimization.
- **Coding-Driven Design** : K2.6 is capable of transforming simple prompts and visual inputs into production-ready interfaces and lightweight full-stack workflows, generating structured layouts, interactive elements, and rich animations with deliberate aesthetic precision.
- **Elevated Agent Swarm** : Scaling horizontally to 300 sub-agents executing 4,000 coordinated steps, K2.6 can dynamically decompose tasks into parallel, domain-specialized subtasks, delivering end-to-end outputs from documents to websites to spreadsheets in a single autonomous run.
- **Proactive & Open Orchestration** : For autonomous tasks, K2.6 demonstrates strong performance in powering persistent, 24/7 background agents that proactively manage schedules, execute code, and orchestrate cross-platform operations without human oversight.

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

| Benchmark | <sup>Kimi K2.6</sup> | <sup>GPT-5.4  <sup>(xhigh)</sup></sup> | <sup>Claude Opus 4.6  <sup>(max effort)</sup></sup> | <sup>Gemini 3.1 Pro <sup>(thinking high)</sup></sup> | <sup>Kimi K2.5</sup> | 
|---|---|---|---|---|---|
| **Agentic** |  |  |  |  |  | 
| HLE-Full (w/ tools) | 54.0 | 52.1 | 53.0 | 51.4 | 50.2 | 
| BrowseComp | 83.2 | 82.7 | 83.7 | 85.9 | 74.9 | 
| BrowseComp (Agent Swarm) | 86.3 |  |  |  | 78.4 | 
| DeepSearchQA (f1-score) | 92.5 | 78.6 | 91.3 | 81.9 | 89.0 | 
| DeepSearchQA (accuracy) | 83.0 | 63.7 | 80.6 | 60.2 | 77.1 | 
| WideSearch (item-f1) | 80.8 | - | - | - | 72.7 | 
| Toolathlon | 50.0 | 54.6 | 47.2 | 48.8 | 27.8 | 
| MCPMark | 55.9 | 62.5* | 56.7* | 55.9* | 29.5 | 
| Claw Eval (pass^3) | 62.3 | 60.3 | 70.4 | 57.8 | 52.3 | 
| Claw Eval (pass@3) | 80.9 | 78.4 | 82.4 | 82.9 | 75.4 | 
| APEX-Agents | 27.9 | 33.3 | 33.0 | 32.0 | 11.5 | 
| OSWorld-Verified | 73.1 | 75.0 | 72.7 | - | 63.3 | 
| **Coding** |  |  |  |  |  | 
| Terminal-Bench 2.0 (Terminus-2) | 66.7 | 65.4* | 65.4 | 68.5 | 50.8 | 
| SWE-Bench Pro | 58.6 | 57.7 | 53.4 | 54.2 | 50.7 | 
| SWE-Bench Multilingual | 76.7 | - | 77.8 | 76.9* | 73.0 | 
| SWE-Bench Verified | 80.2 | - | 80.8 | 80.6 | 76.8 | 
| SciCode | 52.2 | 56.6 | 51.9 | 58.9 | 48.7 | 
| OJBench (python) | 60.6 | - | 60.3 | 70.7 | 54.7 | 
| LiveCodeBench (v6) | 89.6 | - | 88.8 | 91.7 | 85.0 | 
| **Reasoning & Knowledge** |  |  |  |  |  | 
| HLE-Full | 34.7 | 39.8 | 40.0 | 44.4 | 30.1 | 
| AIME 2026 | 96.4 | 99.2 | 96.7 | 98.3 | 95.8 | 
| HMMT 2026 (Feb) | 92.7 | 97.7 | 96.2 | 94.7 | 87.1 | 
| IMO-AnswerBench | 86.0 | 91.4 | 75.3 | 91.0* | 81.8 | 
| GPQA-Diamond | 90.5 | 92.8 | 91.3 | 94.3 | 87.6 | 
| **Vision** |  |  |  |  |  | 
| MMMU-Pro | 79.4 | 81.2 | 73.9 | 83.0* | 78.5 | 
| MMMU-Pro (w/ python) | 80.1 | 82.1 | 77.3 | 85.3* | 77.7 | 
| CharXiv (RQ) | 80.4 | 82.8* | 69.1 | 80.2* | 77.5 | 
| CharXiv (RQ) (w/ python) | 86.7 | 90.0* | 84.7 | 89.9* | 78.7 | 
| MathVision | 87.4 | 92.0* | 71.2* | 89.8* | 84.2 | 
| MathVision (w/ python) | 93.2 | 96.1* | 84.6* | 95.7* | 85.0 | 
| BabyVision | 39.8 | 49.7 | 14.8 | 51.6 | 36.5 | 
| BabyVision (w/ python) | 68.5 | 80.2* | 38.4* | 68.3* | 40.5 | 
| V* (w/ python) | 96.9 | 98.4* | 86.4* | 96.9* | 86.9 | 

## **Footnotes**

1. **General Testing Details**
  - We report results for Kimi K2.6 and Kimi K2.5 with thinking mode enabled, Claude Opus 4.6 with max effort, GPT-5.4 with xhigh reasoning effort, and Gemini 3.1 Pro with a high thinking level.
  - Unless otherwise specified, all Kimi K2.6 experiments were conducted with temperature = 1.0, top-p = 1.0, and a context length of 262,144 tokens.
  - Benchmarks without publicly available scores were re-evaluated under the same conditions used for Kimi K2.6 and are marked with an asterisk (`*` ). Except where noted with an asterisk, all other results are cited from official reports.
2. **Reasoning Benchmarks**
  - IMO-AnswerBench scores for GPT-5.4 and Claude 4.6 were obtained from z.ai/blog/glm-5.1.
  - Humanity's Last Exam (HLE) and other reasoning tasks were evaluated with a maximum generation length of 98,304 tokens. By default, we report results on the HLE full set. For the text-only subset, Kimi K2.6 achieves 36.4% accuracy without tools and 55.5% with tools.
3. **Tool-Augmented / Agentic Tasks**
  - Kimi K2.6 was equipped with search, code-interpreter, and web-browsing tools for HLE with tools, BrowseComp, DeepSearchQA, and WideSearch.
  - For HLE-Full with tools, the maximum generation length is 262,144 tokens with a per-step limit of 49,152 tokens. We employ a simple context management strategy: once the context window exceeds the threshold, only the most recent round of tool-related messages is retained.
  - For BrowseComp, we report scores obtained with context management using the same discard-all strategy as Kimi K2.5 and DeepSeek-V3.2.
  - For DeepSearchQA, no context management was applied to Kimi K2.6 tests, and tasks exceeding the supported context length were directly counted as failed. Scores for Claude Opus 4.6, GPT-5.4, and Gemini 3.1 Pro on DeepSearchQA are cited from the Claude Opus 4.7 System Card.
  - For WideSearch, we report results under the "hide tool result" context management setting. Once the context window exceeds the threshold, only the most recent round of tool-related messages is retained.
  - The test system prompts are identical to those used in the Kimi K2.5 technical report.
  - Claw Eval was conducted using version 1.1 with max-tokens-per-step = 16384.
  - For APEX-Agents, we evaluate 452 tasks from the public 480-task release, as done by Artificial Analysis(excluding Investment Banking Worlds 244 and 246, which have external runtime dependencies)
4. **Coding Tasks**
  - Terminal-Bench 2.0 scores were obtained with the default agent framework (Terminus-2) and the provided JSON parser, operating in preserve thinking mode.
  - For the SWE-Bench series of evaluations (including Verified, Multilingual, and Pro), we used an in-house evaluation framework adapted from SWE-agent. This framework includes a minimal set of tools—bash tool, createfile tool, insert tool, view tool, strreplace tool, and submit tool.
  - All reported scores for coding tasks are averaged over 10 independent runs.
5. **Vision Benchmarks**
  - Max-tokens = 98,304, averaged over three runs (avg@3).
  - Settings with Python tool use max-tokens-per-step = 65,536 and max-steps = 50 for multi-step reasoning.
  - MMMU-Pro follows the official protocol, preserving input order and prepending images.

Kimi-K2.6 adopts the same native int4 quantization method as Kimi-K2-Thinking.

