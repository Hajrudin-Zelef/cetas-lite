---
id: collect-240926-huggingface/huggingface/unsloth-glm-5-3-gguf-hugging-face-2
title: "Read our How to Run GLM-5.3 Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "ExploitGym", "Moonshot", "Nvidia", "OpenAI", "Z.ai", "vLLM"]
dates: []
keywords: ["glm", "agent", "claude", "compute", "context window", "gpt-5.6", "inference", "kimi", "luna", "nvidia", "parameters", "reasoning"]
source: docs/RAG/clean_en/huggingface/unsloth-glm-5-3-gguf-hugging-face.md
source_anchor: ""
source_lines: [38, 55]
sha256: c67b6910ff2c4500e2787eee3377bc174ef9e5da338b60f3caaaffc10d1da2db
---

# Read our How to Run GLM-5.3 Guide!

- **HLE w/ tools** : We use sampling parameters of`temperature=1.0` and`top_p=0.95` for evaluation, with a maximum generation length of`163,840` tokens. The evaluation is conducted with a maximum context length of`300,000` tokens, using a context management strategy. We use GPT-5.6-luna (medium) as the judge model.
- **NL2Repo** : We evaluated NL2Repo with`temperature=1.0` ,`top_p=1.0` , and`max_new_tokens=64k` under 1M context. To prevent hacking, we use rule-based and a LLM-based judgement to prevent malicious behaviors (e.g., unauthorized pip or curl operations).
- **DeepSWE** : We run DeepSWE using the mini-swe-agent harness with`temperature=0.95` ,`top_p=1.0` ,`timeout=6h` and 400K context.
- **Terminal-Bench 2.1** : We evaluate in Claude Code 2.1.207 with`temperature=1.0` ,`top_p=1` ,`max_new_tokens=65536` with 6h timeout.
- **Terminal-Bench 3.0** : We evaluate Terminal-Bench-3 tasks with the Claude Code 2.1.207 harness (reasoning effort=max, 400K context, and 128K maximum output), reporting avg@3 over three rollouts per task. Each rollout runs in an isolated container built from the task's official image, and is capped at 600 agent turns with a 10-hour timeout. Tool Search is disabled, and the artifacts each agent produces are scored by the task's official separate verifier.
- **Agent's Last Exam (CLI)** : We evaluate ALE using the official evaluation protocol with the Claude Code harness (reasoning effort=max, 1M context, and 64K maximum output). Each of the 105 tasks runs in an isolated Docker container using the resources declared in its Task Card. The default timeout is 4 hours, with task-specific limits taking precedence (up to 8 hours). Tool Search is disabled, and results are scored by the official ALE evaluators.
- **Toolathlon Verified** : We obtain all results via the official evaluation service and report pass@1 averaged over 3 independent runs.
- **AutomationBench** : We evaluate on AutomationBench**v1.0.6** , incorporating the fix for the`null` -type handling issue introduced in PR #13.
- **GDPval-AA v2** : Models are evaluated by Artificial Analysis.
- **CyberGym** : We evaluate GLM-5.3 in Claude Code 2.1.207 (max reasoning effort, no web tools with`temperature=1.0` ,`top_p=1.0` ,`max_new_tokens=128000` ). All evaluations are under unlimited timeout per task and results are single-run Pass@1 over 1,507 tasks. To simulate real-world usage scenarios, we place the agent inside the task container. We also remove all Git-related information and apply a domain whitelist (allowing only essential domains such as pypi.org and deb.debian.org for basic tool installation) to prevent the agent from cheating.
- **ExploitGym** : We evaluate GLM-5.3, Kimi-K3 and Qwen3.8 Max in Claude Code 2.1.207 (max reasoning effort, no web tools with`temperature=1.0` ,`top_p=1.0` ,`max_new_tokens=128000` ). The reported results are single-run Pass@1 on 869 tasks under two timeout budgets: 2 hours and 6 hours, which are calculated as the API inference time rescaled by per-model tokens per second rate (per-model TPS sourced from Artificial Analysis; that is, we rescale GLM-5.3's results by 115 TPS, Kimi K3's results by 40 TPS and Qwen3.8 Max's results by 47 TPS), plus the non-API overhead. We also apply a domain whitelist (allowing only essential domains such as pypi.org and deb.debian.org for basic tool installation) to prevent the agent from cheating.
- **ExploitBench** : We evaluate GLM-5.3 in Claude Code 2.1.207 (max reasoning effort, no web tools with`temperature=1.0` ,`top_p=1.0` ,`max_new_tokens=128000` ). Following the official evaluation settings, we limit the maximum number of interaction rounds between the agent and the environment to 300, and compute the average coverage score over all 41 tasks across 3 revisions. The coverage result of a task is determined by taking the union of capabilities achieved across all revisions, and the average score is obtained by averaging the results. We also apply a domain whitelist (allowing only essential domains such as pypi.org and deb.debian.org for basic tool installation) to prevent the agent from cheating.
- **FrontierSWE** : The evaluation was conducted by Proximal with 1M context length, max effort level, and 128K maximum output tokens. Dominance score reported as of 2026/08/14.
- **PostTrainBench** : We evaluate GLM-5.3 using Claude Code 2.1.207 with max effort level,`temperature = 1.0` ,`top_p = 1.0` ,`max_new_tokens = 128000` , and a 1M-token context window. We report the weighted average over 3 runs. Runs that fail to produce a score fall back to the official zero-shot base-model baseline score. For checks intended to prevent the use of third-party APIs, we removed the original pattern-matching-based checks, as they produced false positives when a local vLLM endpoint was accessed through the OpenAI SDK. Instead, we use an LLM agent to inspect solutions for external API usage.
- **SWE-Marathon** : We evaluate GLM-5.3 using Claude Code 2.1.207 with maximum effort level,`temperature = 1.0` ,`top_p = 0.95` ,`max_new_tokens = 128000` , and a 1M-token context window. For`strip-clone` , the original anti-cheat checks used overly broad import detection that could reject valid implementations. We removed the affected checks and performed llm-based inspection instead to avoid false positives. For`parameter-golf` and`trimul-cuda` , changes to the NVIDIA wheels caused the Docker image builds to fail, so we added`--extra-index-url https://pypi.org/simple` to restore successful builds.

If you find GLM-5.3 useful in your research, please cite our technical report:

