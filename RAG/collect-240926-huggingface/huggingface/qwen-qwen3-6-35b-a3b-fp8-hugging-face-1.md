---
id: collect-240926-huggingface/huggingface/qwen-qwen3-6-35b-a3b-fp8-hugging-face-1
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "attention", "benchmark", "claude", "context window", "distribution", "embedding", "fp8", "gemini", "inference", "mcp"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-6-35b-a3b-fp8-hugging-face.md
source_anchor: ""
source_lines: [1, 142]
sha256: 5735ba0013a464c59c85d67793655e916218562e25c244b7e3acf610fb1525fe
---

# Set the following accordingly

<!-- source: https://huggingface.co/Qwen/Qwen3.6-35B-A3B-FP8 -->

This repository contains FP8-quantized model weights and configuration files for the post-trained model in the Hugging Face Transformers format.

These artifacts are compatible with Hugging Face Transformers, vLLM, SGLang, KTransformers, etc.

The quantization method is fine-grained fp8 quantization with block size of 128, and its performance metrics are nearly identical to those of the original model.


Following the February release of the Qwen3.5 series, we're pleased to share the first open-weight variant of Qwen3.6. Built on direct feedback from the community, Qwen3.6 prioritizes stability and real-world utility, offering developers a more intuitive, responsive, and genuinely productive coding experience.

This release delivers substantial upgrades, particularly in

- **Agentic Coding:** the model now handles frontend workflows and repository-level reasoning with greater fluency and precision.
- **Thinking Preservation:** we've introduced a new option to retain reasoning context from historical messages, streamlining iterative development and reducing overhead.

For more details, please refer to our blog post Qwen3.6-35B-A3B.

- Type: Causal Language Model with Vision Encoder
- Training Stage: Pre-training & Post-training
- Language Model
  - Number of Parameters: 35B in total and 3B activated
  - Hidden Dimension: 2048
  - Token Embedding: 248320 (Padded)
  - Number of Layers: 40
  - Hidden Layout: 10 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE))
  - Gated DeltaNet:
    - Number of Linear Attention Heads: 32 for V and 16 for QK
    - Head Dimension: 128
  - Gated Attention:
    - Number of Attention Heads: 16 for Q and 2 for KV
    - Head Dimension: 256
    - Rotary Position Embedding Dimension: 64
  - Mixture Of Experts
    - Number of Experts: 256
    - Number of Activated Experts: 8 Routed + 1 Shared
    - Expert Intermediate Dimension: 512
  - LM Output: 248320 (Padded)
  - MTP: trained with multi-steps
- Context Length: 262,144 natively and extensible up to 1,010,000 tokens.

|  | Qwen3.5-27B | Gemma4-31B | Qwen3.5-35BA3B | Gemma4-26BA4B | Qwen3.6-35BA3B | 
|---|---|---|---|---|---|
| Coding Agent |  |  |  |  |  | 
| SWE-bench Verified | 75.0 | 52.0 | 70.0 | 17.4 | 73.4 | 
| SWE-bench Multilingual | 69.3 | 51.7 | 60.3 | 17.3 | 67.2 | 
| SWE-bench Pro | 51.2 | 35.7 | 44.6 | 13.8 | 49.5 | 
| Terminal-Bench 2.0 | 41.6 | 42.9 | 40.5 | 34.2 | 51.5 | 
| Claw-Eval <sub>Avg</sub> | 64.3 | 48.5 | 65.4 | 58.8 | 68.7 | 
| Claw-Eval <sub>Pass^3</sub> | 46.2 | 25.0 | 51.0 | 28.0 | 50.0 | 
| SkillsBench <sub>Avg5</sub> | 27.2 | 23.6 | 4.4 | 12.3 | 28.7 | 
| QwenClawBench | 52.2 | 41.7 | 47.7 | 38.7 | 52.6 | 
| NL2Repo | 27.3 | 15.5 | 20.5 | 11.6 | 29.4 | 
| QwenWebBench | 1068 | 1197 | 978 | 1178 | 1397 | 
| General Agent |  |  |  |  |  | 
| TAU3-Bench | 68.4 | 67.5 | 68.9 | 59.0 | 67.2 | 
| VITA-Bench | 41.8 | 43.0 | 29.1 | 36.9 | 35.6 | 
| DeepPlanning | 22.6 | 24.0 | 22.8 | 16.2 | 25.9 | 
| Tool Decathlon | 31.5 | 21.2 | 28.7 | 12.0 | 26.9 | 
| MCPMark | 36.3 | 18.1 | 27.0 | 14.2 | 37.0 | 
| MCP-Atlas | 68.4 | 57.2 | 62.4 | 50.0 | 62.8 | 
| WideSearch | 66.4 | 35.2 | 59.1 | 38.3 | 60.1 | 
| Knowledge |  |  |  |  |  | 
| MMLU-Pro | 86.1 | 85.2 | 85.3 | 82.6 | 85.2 | 
| MMLU-Redux | 93.2 | 93.7 | 93.3 | 92.7 | 93.3 | 
| SuperGPQA | 65.6 | 65.7 | 63.4 | 61.4 | 64.7 | 
| C-Eval | 90.5 | 82.6 | 90.2 | 82.5 | 90.0 | 
| STEM & Reasoning |  |  |  |  |  | 
| GPQA | 85.5 | 84.3 | 84.2 | 82.3 | 86.0 | 
| HLE | 24.3 | 19.5 | 22.4 | 8.7 | 21.4 | 
| LiveCodeBench v6 | 80.7 | 80.0 | 74.6 | 77.1 | 80.4 | 
| HMMT Feb 25 | 92.0 | 88.7 | 89.0 | 91.7 | 90.7 | 
| HMMT Nov 25 | 89.8 | 87.5 | 89.2 | 87.5 | 89.1 | 
| HMMT Feb 26 | 84.3 | 77.2 | 78.7 | 79.0 | 83.6 | 
| IMOAnswerBench | 79.9 | 74.5 | 76.8 | 74.3 | 78.9 | 
| AIME26 | 92.6 | 89.2 | 91.0 | 88.3 | 92.7 | 

* SWE-Bench Series: Internal agent scaffold (bash + file-edit tools); temp=1.0, top_p=0.95, 200K context window. We correct some problematic tasks in the public set of SWE-bench Pro and evaluate all baselines on the refined benchmark.

* Terminal-Bench 2.0: Harbor/Terminus-2 harness; 3h timeout, 32 CPU/48 GB RAM; temp=1.0, top_p=0.95, top_k=20, max_tokens=80K, 256K ctx; avg of 5 runs.

* SkillsBench: Evaluated via OpenCode on 78 tasks (self-contained subset, excluding API-dependent tasks); avg of 5 runs.

* NL2Repo: Others are evaluated via Claude Code (temp=1.0, top_p=0.95, max_turns=900).

* QwenClawBench: An internal real-user-distribution Claw agent benchmark (open-sourcing soon); temp=0.6, 256K ctx.

* QwenWebBench: An internal front-end code generation benchmark; bilingual (EN/CN), 7 categories (Web Design, Web Apps, Games, SVG, Data Visualization, Animation, and 3D); auto-render + multimodal judge (code/visual correctness); BT/Elo rating system.

* TAU3-Bench: We use the official user model (gpt-5.2, low reasoning effort) + default BM25 retrieval.

* VITA-Bench: Avg subdomain scores; using claude-4-sonnet as judger, as the official judger (claude-3.7-sonnet) is no longer available.

* MCPMark: GitHub MCP v0.30.3; Playwright responses truncated at 32K tokens.

* MCP-Atlas: Public set score; gemini-2.5-pro judger.

* AIME 26: We use the full AIME 2026 (I & II), where the scores may differ from Qwen 3.5 notes.


|  | Qwen3.5-27B | Claude-Sonnet-4.5 | Gemma4-31B | Gemma4-26BA4B | Qwen3.5-35B-A3B | Qwen3.6-35B-A3B | 
|---|---|---|---|---|---|---|
| STEM and Puzzle |  |  |  |  |  |  | 
| MMMU | 82.3 | 79.6 | 80.4 | 78.4 | 81.4 | 81.7 | 
| MMMU-Pro | 75.0 | 68.4 | 76.9* | 73.8* | 75.1 | 75.3 | 
| Mathvista(mini) | 87.8 | 79.8 | 79.3 | 79.4 | 86.2 | 86.4 | 
| ZEROBench_sub | 36.2 | 26.3 | 26.0 | 26.3 | 34.1 | 34.4 | 
| General VQA |  |  |  |  |  |  | 
| RealWorldQA | 83.7 | 70.3 | 72.3 | 72.2 | 84.1 | 85.3 | 
| MMBench <sub>EN-DEV-v1.1</sub> | 92.6 | 88.3 | 90.9 | 89.0 | 91.5 | 92.8 | 
| SimpleVQA | 56.0 | 57.6 | 52.9 | 52.2 | 58.3 | 58.9 | 
| HallusionBench | 70.0 | 59.9 | 67.4 | 66.1 | 67.9 | 69.8 | 
| Text Recognition and Document Understanding |  |  |  |  |  |  | 
| OmniDocBench1.5 | 88.9 | 85.8 | 80.1 | 74.4 | 89.3 | 89.9 | 
| CharXiv(RQ) | 79.5 | 67.2 | 67.9 | 69.0 | 77.5 | 78.0 | 
| CC-OCR | 81.0 | 68.1 | 75.7 | 74.5 | 80.7 | 81.9 | 
| AI2D_TEST | 92.9 | 87.0 | 89.0 | 88.3 | 92.6 | 92.7 | 
| Spatial Intelligence |  |  |  |  |  |  | 
| RefCOCO(avg) | 90.9 | -- | -- | -- | 89.2 | 92.0 | 
| ODInW13 | 41.1 | -- | -- | -- | 42.6 | 50.8 | 
| EmbSpatialBench | 84.5 | 71.8 | -- | -- | 83.1 | 84.3 | 
| RefSpatialBench | 67.7 | -- | -- | -- | 63.5 | 64.3 | 
| Video Understanding |  |  |  |  |  |  | 
| VideoMME <sub>(w sub.)</sub> | 87.0 | 81.1 | -- | -- | 86.6 | 86.6 | 
| VideoMME <sub>(w/o sub.)</sub> | 82.8 | 75.3 | -- | -- | 82.5 | 82.5 | 
| VideoMMMU | 82.3 | 77.6 | 81.6 | 76.0 | 80.4 | 83.7 | 
| MLVU | 85.9 | 72.8 | -- | -- | 85.6 | 86.2 | 
| MVBench | 74.6 | -- | -- | -- | 74.8 | 74.6 | 
| LVBench | 73.6 | -- | -- | -- | 71.4 | 71.4 | 

* Empty cells (--) indicate scores not available or not applicable.

For streamlined integration, we recommend using Qwen3.6 via APIs. Below is a guide to use Qwen3.6 via OpenAI-compatible API.

Qwen3.6 can be served via APIs with popular inference frameworks. In the following, we show example commands to launch OpenAI-Compatible API servers for Qwen3.6 models.

Inference efficiency and throughput vary significantly across frameworks. We recommend using the latest framework versions to ensure optimal performance and compatibility. For production workloads or high-throughput scenarios, dedicated serving engines such as SGLang, KTransformers or vLLM are strongly recommended.


The model has a default context length of 262,144 tokens. If you encounter out-of-memory (OOM) errors, consider reducing the context window. However, because Qwen3.6 leverages extended context for complex tasks, we advise maintaining a context length of at least 128K tokens to preserve thinking capabilities.


