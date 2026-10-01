---
id: collect-261001-ia-llm/ia-llm/glm-5-2-built-for-long-horizon-tasks-2
title: "glm-5-2-built-for-long-horizon-tasks"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "MiniMax", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["glm", "agent", "agentic", "agents", "benchmark", "claude", "deepseek", "fp8", "gemini", "inference", "mcp", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/glm-5-2-built-for-long-horizon-tasks.md
source_anchor: ""
source_lines: [37, 88]
sha256: 3c07902f1ff633ed5a29c11fdbca5acd97073be71c5aa8c5d489cdc824d66597
---

# glm-5-2-built-for-long-horizon-tasks

The agentic RL post-training of GLM-5.2 involves tasks at larger scale, across more domains, and with more complex execution patterns. Heterogeneous data and tasks need to be organized within a unified training process, while long-horizon interactions, tool use, sub-task decomposition, and multi-turn environment feedback all impose higher requirements on rollout and training orchestration. To support this process, slime serves as an integrated infrastructure layer from training to large-scale inference rollout. It supports multiple training and task organization modes, including white-box rollout, black-box rollout, compact trajectory, and sub-agent workflow, enabling the same system to scale to larger and more complex RL and OPD training workloads. In the post-training process of GLM-5.2, we used the slime framework to conduct parallel OPD training, efficiently merging more than ten expert models into the final model. The entire OPD training process took approximately two days, demonstrating high training efficiency.

Agentic RL also places higher demands on system resources and inference infrastructure. slime provides a highly open and flexible interface to inference systems: the training side can connect to inference services in different forms, and flexibly adapt to different parallelism strategies, routing policies, PD disaggregation setups, and deployment patterns. At the same time, the configuration experience, scheduling strategies, and optimization paths accumulated during RL rollout can be reused and further refined in the production serving stage, allowing the training side and the serving side to reinforce each other. This creates a more direct path from post-training to production deployment. Together with flexible training-inference resource organization and KV-cache FP8, slime provides critical infrastructure support for GLM-5.2’s large-scale agentic RL training, further improving system efficiency, rollout throughput, and large-scale inference concurrency.

**RL for Long-Horizon Tasks**. For GLM-5.2, long-horizon tasks produce substantially longer execution traces, and once a super-long trajectory is split by compaction into multiple sub-traces, different rollouts under the same prompt yield different numbers of trainable traces with highly variable lengths. We therefore move from group-wise optimization to a critic-based PPO formulation that learns from individual rollouts, relying on a critic to estimate token-level advantages rather than group-relative comparisons. This single-rollout formulation fits compaction naturally, as it places no constraint on how many traces a prompt produces or on their relative lengths: we bring compaction into training by including all compacted sub-traces as trainable trajectories, and apply a token-level loss to address their length imbalance.

**Anti-Hack in Coding agents**. Coding RL is especially vulnerable to reward hacking because the reward is typically a verifiable pass/fail signal. We find that GLM-5.2 shows more potential hacking behavior than GLM-5.1. This makes the verification signal easy to optimize, but fails to actually improve the fundamental capabilities of the model. An agent can read protected evaluation artifacts, copy answer content from references or upstream commits, or directly fetch the target source in GitHub-related tasks. For example, the agent may download solution via `curl https://raw.githubusercontent.com/<path-to-file>` or even chained leakage like

```
1. find /workspace -name "*hidden*"
2. cat /workspace/.eval/secret_cases.json
3. python solve.py --case "$(cat /workspace/.eval/secret_cases.json)"
```
These behaviors inflate rewards and corrupt the training signal, requiring a clear mechanism to separate real task-solving from shortcuts. To address this, we introduce an anti-hack module for both RL training and evaluation. The detection process has two stages: a rule-based filter first catches potential hacks to maximize recall, and then an LLM judge checks the intent of these flagged actions to keep precision high. We use an online strategy that monitors the tool calls at each step. If a hack is detected, the system blocks the call and returns dummy information as the result. Importantly, this online guard allows the model to continue the rollout even after a hacked action is caught. By handling the specific invalid behavior instead of rejecting the entire trajectory, this approach helps prevent the training instability and model collapse that can happen when rollouts are abruptly stopped.

| Benchmark | GLM-5.2 | GLM-5.1 | Qwen3.7-Max | MiniMax M3 | DeepSeek-V4-Pro | Claude Opus 4.8 | GPT-5.5 | Gemini 3.1 Pro | 
|---|---|---|---|---|---|---|---|---|
| Reasoning |  |  |  |  |  |  |  |  | 
| HLE | 40.5 | 31 | 41.4 | 37 | 37.7 | 49.8* | 41.4* | 45 | 
| HLE (w/ Tools) | 54.7 | 52.3 | 53.5 | - | 48.2 | 57.9* | 52.2* | 51.4* | 
| CritPt | 16.7 | 4.6 | 13.4 | 3.7 | 12.9 | 20.9 | 27.1 | 17.7 | 
| AIME 2026 | 99.2 | 95.3 | 97 | - | 94.6 | 95.7 | 98.3 | 98.2 | 
| HMMT Nov. 2025 | 94.4 | 94 | 95 | 84.4 | 94.4 | 96.5 | 96.5 | 94.8 | 
| HMMT Feb. 2026 | 92.5 | 82.6 | 97.1 | 84.4 | 95.2 | 96.7 | 96.7 | 87.3 | 
| IMOAnswerBench | 91.0 | 83.8 | 90 | - | 89.8 | 83.5 | - | 81 | 
| GPQA-Diamond | 91.2 | 86.2 | 90 | 93 | 90.1 | 93.6 | 93.6 | 94.3 | 
| Coding |  |  |  |  |  |  |  |  | 
| SWE-bench Pro | 62.1 | 58.4 | 60.6 | 59 | 55.4 | 69.2 | 58.6 | 54.2 | 
| NL2Repo | 48.9 | 42.7 | 47.2 | 42.1 | 35.5 | 69.7 | 50.7 | 33.4 | 
| DeepSWE | 46.2 | 18 | 18 | 20 | 8 | 58 | 70 | 10 | 
| ProgramBench | 63.7 | 50.9 | - | - | 47.8 | 71.9 | 70.8 | 39.5 | 
| Terminal Bench 2.1 (Terminus-2) | 81.0 | 63.5 | 75 | 65 | 64 | 85 | 84 | 74 | 
| Terminal Bench 2.1 (Best Reported Harness) | 82.7 | 69 | - | - | - | 78.9 | 83.4 | 70.7 | 
| FrontierSWE (Dominance) | 74.4 | 30.5 | - | - | 29.0 | 75.1 | 72.6 | 39.6 | 
| PostTrainBench | 34.3 | 20.1 | - | - | - | 37.2 | 28.4 | 21.6 | 
| SWE-Marathon | 13.0 | 1.0 | - | - | - | 26.0 | 12.0 | 4.0 | 
| Agentic |  |  |  |  |  |  |  |  | 
| MCP-Atlas (Public Set) | 76.8 | 71.8 | 76.4 | 74.2 | 73.6 | 77.8 | 75.3 | 69.2 | 
| Tool-Decathlon | 48.2 | 40.7 | - | - | 52.8 | 59.9 | 55.6 | 48.8 | 

Try **GLM-5.2** in your favorite coding agents—**ZCode, Claude Code, OpenCode**, and more. https://docs.z.ai/devpack/overview

**For GLM Coding Plan subscribers:** We already rolled out GLM-5.2 to all Coding Plan users. You can enable GLM-5.2 now by updating the model name to `"GLM-5.2"` (or GLM-5.2[1m] in Claude Code to enable 1M context length). You can also choose different thinking effort, High or Max, depending on the task. As our most capable model, GLM-5.2 consumes quota at 3× during peak hours and 2× during off-peak hours. As a limited-time promotion through the end of September, off-peak usage is billed at 1×. (Peak hours are 14:00–18:00 UTC+8 (Beijing Time) daily).

Prefer a GUI? We offer **ZCode** —a desktop agent powered by GLM-5.2, with /goal for long-horizon tasks, SSH remote development, and mobile control. **Special offer**: use GLM-5.2 through Coding Plan inside ZCode and get 1.5x effective quota until June 30.

**Start building now:** https://z.ai/subscribe

GLM-5.2 is now available on Z.ai.

The model weights of GLM-5.2 are publicly available on HuggingFace and ModelScope. For local deployment, GLM-5.2 supports inference frameworks including transformers, vLLM, SGLang, xLLM, ktransformers.

