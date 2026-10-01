---
id: collect-261001-ia-llm/ia-llm/granite-4-2-llms-how-they-re-built-2
title: "Default: previous thinking is stripped to save context"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "vLLM"]
dates: []
keywords: ["agent", "agentic", "alignment", "kv cache", "lean", "parameters", "reasoning", "rlhf", "sandbox", "tool use", "training", "vllm"]
source: docs/RAG/collect-261001-ia-llm/granite-4-2-llms-how-they-re-built.md
source_anchor: ""
source_lines: [71, 145]
sha256: a6dac8bd7c57080803991faea721ca05b7d8aa9c4f713dca06cf46a069e42b52
---

# Default: previous thinking is stripped to save context

After SFT, we apply a **multi-stage, multi-environment reinforcement learning pipeline**. Rather than a single RL pass, we run a *chain* of focused stages spanning many environments: math, code, science, instruction following, tool use, and structured output, then software engineering, terminal use, and web search. Each stage is an independent RL run that targets one capability and warm-starts from the previous stage's checkpoint.

*Figure 1. The staged RL curriculum. Foundational RL (verifiable rewards + skill boosters) runs for all sizes; the agentic RL block (SWE → Terminal → Search) runs for 8B and 30B only. Every model finishes with RLHF. Each stage is a separate GRPO run that warm-starts from the previous checkpoint.*

Every stage trains with **asynchronous GRPO** (Group Relative Policy Optimization), so the generator and trainer halves of the loop never block on each other. A pool of generation workers keeps sampling responses and dropping the finished trajectories into a shared buffer; once the buffer holds a full step's worth, the trainer pulls that batch, takes an optimizer step, and streams the updated parameters back to the generator workers without pausing them. A refresh can land partway through a rollout, leaving a single trajectory stitched together from two adjacent policy versions. We allow this instead of paying to prevent it: the workers reuse their existing KV cache rather than rebuilding it after each refresh, and the one guardrail is a limit that keeps them from drifting more than a single update behind the trainer, which bounds how off-policy any sample can get. Whatever mismatch survives that limit is handled in the objective by *truncated importance sampling*, which clamps the train-versus-generation log-probability ratio to a fixed ceiling so a handful of stale tokens cannot dominate an update.

Advantages are group-relative with a **leave-one-out baseline**: each response is judged against the mean reward of the *other* samples drawn for the same prompt, which removes the need for a separate value network. To make this concrete, take **RLVR**, the first and longest-running stage: each step pairs **256 prompts** with **16 sampled responses apiece** for a **4,096**-example batch, which the trainer consumes in a single optimizer step before the next rollout begins. Later stages keep this machinery unchanged and adjust only the per-stage shape, shown next.

The pipeline keeps a common backbone of hyperparameters across every stage, which makes the curriculum easier to run and compare. A handful of knobs are fixed everywhere:

| Parameter | Value (shared across stages) | 
|---|---|
| Algorithm | GRPO (no value network; group-relative advantages) | 
| Training stack | NeMo-RL (Megatron-Core + vLLM) with NeMo-Gym environments | 
| Ratio clip (min / max) | 0.2 / 0.28 | 
| Micro-batch size | 1 | 
| Parallelism | tensor-parallel 2–4; no pipeline- or context-parallelism | 

What changes from stage to stage is the *shape* of each run: how many prompts and generations per step, how long the context is, whether the agent loop runs, and how hard we pull back toward the reference policy. The table below gives the exact settings for the **30B** chain, stage by stage:

| Stage | Prompts/step | Gens/prompt | Max seq len | Rollout turns | KL | LR | 
|---|---|---|---|---|---|---|
| RLVR (×3) | 256 | 16 | 64K | 1 | 0 | 5e-7 | 
| IF booster | 256 | 16 | 64K | 1 | 0 | 5e-7 | 
| Code booster | 64 | 16 | 64K | 1 | 0.05 | 5e-7 | 
| SWE 1 | 64 | 16 | 128K | 1 | 0.01 | 5e-7 | 
| SWE 2 | 32 | 16 | 128K | **128** | 0 | 5e-7 | 
| Terminal | 8 | 32 | 64K | **64** | 0.01 | **1e-6** | 
| Search | 32 | 16 | 128K | **64** | 0.01 | 5e-7 | 
| RLHF | 128 | 16 | 48K | 1 | 0.05 | 5e-7 | 

*Parameters shown for the 30B model. Global batch size = prompts/step × generations/prompt (e.g. 256 × 16 = 4096 for RLVR). The 3B and 8B models use the same recipe and hyperparameters with **fewer stages** (see How the Three Sizes Differ); the stage list is what changes, not the knobs.*

The **KL schedule** follows the reward type: explore freely where the reward is objective and verifiable (RLVR and SWE 2 run at KL 0), and stay close to the reference where the objective is preference, safety, or a narrow skill graft (RLHF and the code booster use KL 0.05). The **rollout-turns** column counts the environment interactions *GRPO itself* sees per rollout. In every case the model is trained on complete, real-environment trajectories.

Each stage is a separate RL run with a single objective and its own reward signal. When it finishes, its policy is exported to Hugging Face format and becomes the base model for the next stage, so the pipeline is a sequence of warm-starts:

```
SFT ─▶ RLVR ─▶ Skill boosters ─▶ SWE agent ─▶ Terminal ─▶ Search ─▶ RLHF
      └──────── foundational RL ────────┘   └──────── agentic RL (8B / 30B) ────────┘
```
The **8B and 30B** models follow the full ladder. The **3B** model takes a shortened path: foundational RL and alignment, without the agentic block.

A stage is defined mostly by *how it is rewarded*. Across the pipeline there are three reward types, and a single stage can use more than one:

| Reward type | What it measures | Used in | 
|---|---|---|
| **Verifiable** | Exact-match, unit tests, format checkers, rule-based checkers on ground truth | RLVR · boosters · SWE | 
| **Reward model / LLM judge** | Open-ended quality, preference, safety, answer correctness | RLVR · Search · RLHF | 
| **Agentic outcome** | Did the model actually solve a task in a real environment? | SWE · Terminal · Search | 

Verifiable rewards are objective and hard to game, so the pipeline front-loads them. Judge- and preference-based rewards handle open-ended qualities that no checker can express. Agentic-outcome rewards are the sparsest: often a single bit at the end of a long tool-use trajectory.

RLVR is the foundational stage and the broadest data mix in the pipeline: a single blended dataset spanning many verifiable domains.

- **Math:** chain-of-thought with boxed-answer checking, plus formal proving in**Lean**
- **Competitive coding:** solutions checked against hidden tests in a sandbox
- **STEM / graduate-level science MCQA** and general knowledge
- **Instruction following:** structured-output and inverse-instruction tasks
- **Tool / function calling:** single-step tool use
- **Reasoning puzzles** and**abstention** (knowing when to refuse)

Each task type carries its own verifier, so the reward is grounded per example. RLVR runs for **two rounds on 3B and 8B, and three on 30B**. Each round is a fresh warm-started run on a re-weighted mix of public and internally curated RL data.

After RLVR, a few short **booster** stages sharpen specific capabilities that benefit from concentrated training focusing on the following domains:

- **Instruction following (IF):** multi-turn chat, inverse-IFEval, structured outputs
- **Code:** competitive coding only

Boosters are small, focused runs. A light KL penalty keeps the model close to its current behavior while nudging one skill.

In the agentic stages the model learns to **act**: call tools, observe results, and iterate inside a real environment, rewarded on whether the task was actually solved. These stages share the same shape: multi-turn tool use, real (not simulated) environments, sparse outcome rewards, and GRPO, warm-started from the coding-boosted checkpoint. They run in order: **SWE → Terminal → Search**.

*Figure 2. The three agentic-RL environments. Each pairs a real harness with a real environment and a sparse, outcome-based reward. The 3B model runs none of these.*

