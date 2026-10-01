---
id: collect-261001-ia-llm/ia-llm/granite-4-2-llms-how-they-re-built-3
title: "Default: previous thinking is stripped to save context"
domain: ia-llm
role: reference
task: reference
actors: ["China", "CoreWeave", "Hugging Face", "Nvidia", "vLLM"]
dates: []
keywords: ["agent", "agentic", "alignment", "benchmark", "benchmarks", "fp8", "gguf", "gptq", "gpu", "gpus", "inference", "jailbreak"]
source: docs/RAG/collect-261001-ia-llm/granite-4-2-llms-how-they-re-built.md
source_anchor: ""
source_lines: [146, 252]
sha256: 0b58b2e3b3f7ee80f409288bbbefd3e7a74940e9582da11920ab029d09eb066c
---

# Default: previous thinking is stripped to save context

- **SWE agent (software engineering).** Each task is a real repository in its own sandbox. Driven by the OpenHands harness, the model reads code, edits files, and runs the test suite over many internal turns. The reward is verifiable: do the hidden tests pass? Tasks are drawn from open-source SWE datasets, each instance backed by a per-repo container image.
- **Terminal agent (terminal / OS operation).** Multi-step tasks in a live shell, run through the Harbor / Terminus-2 agent harness. The model plans a sequence of commands, observes their output, and recovers from errors. Reward is assigned when the task completed successfully. This is the one stage that drives its multi-turn agent loop at the GRPO level, with rollouts spanning up to 64 environment turns.
- **Search agent (deep research).** The model answers hard, multi-hop questions using live web-search tool calls inside a browsing agent loop: gather evidence across hops, reason over it, and produce an answer. Because correctness here is open-ended, the reward is an LLM judge on the final answer.

The final stage of every model is **RLHF for human preference and safety.** It optimizes against a generative reward model (GenRM) for preference, plus a safety reward covering jailbreak resistance and appropriate refusals. This stage uses the highest KL penalty in the pipeline, aligning tone and safety without eroding the capabilities the earlier stages built. In addition to human preference and safety alignment, this stage also applies a reasoning-length penalty to discourage overly verbose reasoning behavior acquired during earlier stages.

Same method and infrastructure; the difference is how far up the ladder each model goes.

| Stage | 3B | 8B | 30B | 
|---|---|---|---|
| **RLVR** (verifiable) | ×2 | ×2 | ×3 | 
| **Skill boosters** | code | IF · GPQA · code | IF · code | 
| **SWE agent** | — | ✓ | ✓ | 
| **Terminal agent** | — | ✓ | ✓ | 
| **Search agent** | — | ✓ | ✓ | 
| **RLHF** (preference + safety) | ✓ | ✓ | ✓ | 

3B is a strong foundational-RL model; 8B and 30B add the agentic-RL block on top, learning to act with tools in real environments.

Reinforcement learning at this scale needs infrastructure that can drive a training loop and a fleet of live environments at the same time. This matters most in the agentic stages, where every training example is a multi-turn rollout that edits code, runs commands, or browses the web. Granite 4.2's RL runs on two open components: **NeMo-RL** on the training side and **NeMo-Gym** on the rollout side.

*Figure 3. The RL system. NeMo-RL drives the GRPO loop (Megatron-Core training backend, vLLM generation, and Megatron-Bridge for HF⇄Megatron weight conversion). NeMo-Gym orchestrates rollouts and hosts the tools, sandboxes, and reward/verifier calls as pluggable **Resources**.*

The division of labor:

- **NeMo-RL (training side).** Megatron-Core is the training backend; vLLM generates rollouts;**Megatron-Bridge** converts weights between Megatron and Hugging Face formats, so each stage can export a clean HF checkpoint for the next one.
- **NeMo-Gym (rollout side).** It exposes each environment as a set of**Resources** (verifiers, tools, sandboxes, and reward models) behind a uniform interface. This is the plug point for the agentic stages: the SWE repo sandboxes, the terminal harness, and the web-search tools all attach here, and to the training loop they look the same as a simple math verifier.

That uniformity is what makes the staged curriculum above practical: a booster's rule-based checker and a full SWE sandbox present the same interface to GRPO.

This split is also what makes the **asynchronous** training loop described above physically possible: generation and policy updates live on separate GPU pools, so the expensive generation fleet — including the live agentic environments — stays busy instead of idling through optimizer steps.

Granite 4.2 was evaluated across agentic coding, general agentic and tool use, reasoning, chat and instruction following, and long context. The full benchmark table is below, followed by charts that break out the headline results by model size.

| Task | 3B Dense | 8B Dense | 30B Dense | 
|---|---|---|---|
| Agentic (Coding) |  |  |  | 
| SWE Bench Multilingual | NA | 30.78 | 41.89 | 
| SWE Bench Pro | NA | 19.11 | 33.29 | 
| SWE Bench Verified | NA | 47.67 | 57.00 | 
| Terminal-Bench 2.1 | NA | 20.56 | 29.24 | 
| Agentic (General) |  |  |  | 
| τ³-bench | 45.78 | 58.06 | 62.00 | 
| BFCL (v4) | 52.41 | 50.29 | 61.39 | 
| ProfBench | 32.10 | 41.20 | 42.90 | 
| BirdBench | NA | 41.07 | 41.85 | 
| GDPval | NA | 1189.00 | 1225.00 | 
| Reasoning |  |  |  | 
| AIME25 | 78.33 | 86.67 | 89.17 | 
| HMMT Feb25 | 66.67 | 78.33 | 89.17 | 
| GPQA | 54.80 | 64.14 | 66.41 | 
| LiveCodeBench v6 | 69.71 | 73.24 | 75.77 | 
| SciCode | 24.11 | 36.09 | 38.76 | 
| Chat & Instruction Following |  |  |  | 
| MMLU-Pro | 67.84 | 74.04 | 77.60 | 
| MMLU-ProX lite (IBM) | 27.78 | 61.06 | 66.64 | 
| Arena-Hard-V2 | 34.96 | 65.19 | 67.93 | 
| IFBench (prompt) | 74.33 | 79.33 | 77.17 | 
| Long Context |  |  |  | 
| RULER 64K | 67.52 | 80.99 | 89.96 | 
| RULER 128K | 55.30 | 71.41 | 81.38 | 

**Supported languages:** English, German, Spanish, French, Japanese, Portuguese, Arabic, Czech, Italian, Korean, Dutch, and Chinese.

The charts below break these results out by capability area.

*Figure 4. Reasoning (pass@1). Scores rise consistently with model size across math (AIME25, HMMT), science (GPQA), and code reasoning (LiveCodeBench, SciCode).*

*Figure 5. Agentic coding resolve rates. The agentic-RL block is trained only for 8B and 30B; the 30B model leads across SWE-Bench variants and Terminal-Bench.*

*Figure 6. General agentic and tool-use benchmarks, reported for all three sizes.*

We also released four quantized variants of the Granite 4.2 models for inference with vLLM. The models are converted to FP8, NVFP4, and MXFP4 using LLM Compressor, and to the GGUF format using the llama.cpp framework for reduced-memory deployment.

The FP8 version is quantized with dynamic per-channel weights and per-token activations. No calibration is used.

The NVFP4 and MXFP4 versions are quantized using GPTQ calibrated on 2K samples drawn from the SFT dataset. Max context length is 2K during calibration.

Conversion of the Granite 4.2 models to GGUF is done with the canonical llama.cpp tool as described in https://github.com/IBM/gguf#gguf-conversion--quantization.

Several GGUF formats are provided:

- Q8_0
- Q6_K
- Q5_K_S
- Q5_K_M
- Q5_1
- Q5_0
- Q4_K_S
- Q4_K_M
- Q4_1
- Q4_0
- Q3_K_S
- Q3_K_M
- Q3_K_L
- Q2_K

We trained the Granite 4.2 language models on an NVIDIA GB200 NVL72 cluster hosted by CoreWeave, featuring:

- A 72-GPU NVLink domain for high-speed intra-rack communication
- A non-blocking Fat-Tree NDR 400 Gb/s InfiniBand fabric for full-bandwidth inter-rack connectivity
- Thousands of GPUs operating at cluster scale

This infrastructure delivers the high-bandwidth, low-latency communication required for efficient large-scale distributed training.

The training software stack is packaged into `.sqsh` container images, each giving a run a reproducible, portable environment with its SBSA-compatible CUDA targets, Linux aarch64 Python wheels, and GPU-specific binaries pinned. The large-scale SFT runs build on an NGC PyTorch base image (Ubuntu 22.04, CUDA 12.8, Python 3.12); the RL stack runs in its own NeMo-RL container.

