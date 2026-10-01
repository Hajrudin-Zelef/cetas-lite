---
id: collect-261001-ia-llm/ia-llm/glm-5-2-built-for-long-horizon-tasks-3
title: "glm-5-2-built-for-long-horizon-tasks"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "claude", "context window", "gemini", "mcp", "memory", "parameters", "reasoning", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/glm-5-2-built-for-long-horizon-tasks.md
source_anchor: ""
source_lines: [89, 100]
sha256: b25e7d8faab9a275ca080a21f1aeff8c4475ac90ff014cde0e65a82964f22b6b
---

# glm-5-2-built-for-long-horizon-tasks

- **Humanity’s Last Exam (HLE) & other reasoning tasks** : We use sampling parameters of`temperature=1.0` ,`top_p=0.95` for evaluation. We evaluate with a maximum generation length of`163,840` tokens. By default, we report the text-only subset; results marked with * are from the full set. For AIME, HMMT and IMOAnswerBench, we evaluate  each question using the following system prompt:`Your response should be in the following format:\nExplanation: {your explanation for your final answer}\nExact Answer: {your succinct, final answer}\nConfidence: {your confidence score between 0% and 100% for your answer}.` We use GPT-5.5 (medium) as the judge model. For HLE-with-tools, we use a maximum context length of 300,000 tokens, with no context management strategy.
- **SWE-Bench Pro** : We run the SWE-Bench Pro suite with OpenHands using a tailored instruction prompt. Settings:`temperature=1` ,`top_p=1` ,`max_new_tokens=32k` , with a 400K context window.
- **NL2Repo** : We evaluated NL2Repo with`temperature=1.0` ,`top_p=1.0` , and`max_new_tokens=48k` under 400k context. To prevent hacking, we use rule-based and a LLM-based judgement to prevent malicious behaviors (e.g., unauthorized pip or curl operations).
- **DeepSWE** : We run DeepSWE with the official pier evaluation framework and the mini-swe-agent harness (`temperature=1.0` ,`top_p=1.0` ,`timeout=2h` , 400K context). Each task is solved in an isolated container with 2 CPUs, 8 GB RAM, and no internet access.
- **ProgramBench** : We evaluate ProgramBench (200 instances) with Claude-Code 2.1.156 using`temperature=1.0, top_p=1.0, max_tokens=64000, max_turns=2000, sample_timeout=6h, reasoning_effort=max` , with a 400K context window. Each instance runs in a (4 CPUs, 8 GB RAM) sandbox with internet access disabled.
- **Terminal-Bench 2.1 (Terminus 2)** : We evaluate Terminal-Bench 2.1 with Terminus-2 framework using`parser=json` ,`timeout=4h` ,`temperature=1.0` ,`top_p=1.0` ,`max_new_tokens=48k` ,`max_episodes=500` , with a 256K context window. Resource limits are capped at 4 CPUs and 8 GB RAM.
- **Terminal-Bench 2.1 (Claude Code)** : We evaluate in Claude Code 2.1.167 with`temperature=1.0, top_p=0.95, max_new_tokens=131072` . We override max_new_tokens to 128k via a transparent proxy, bypassing the 64k CLI cap to restore the configurability of`CLAUDE_CODE_MAX_OUTPUT_TOKENS` . We remove wall-clock time limits, while preserving per-task CPU and memory constraints. Scores are averaged over 5 runs.
- **MCP-Atlas** : All models were evaluated in think mode on the 500-task public subset with a 10-minute timeout per task. We use Gemini-3.0-Pro as the judge model for evaluation.
- **Tool-Decathlon** : We use the official evaluation service and set max_token to 128K.
- **FrontierSWE** : The evaluation was conducted by Proximal with 1M context length, max effort level, and 128K maximum output tokens. Dominance score reported as of 2026/06/16.
- **PostTrainBench** : The evaluation was conducted by PostTrainBench with 1M context length, max effort level, and 128K maximum output tokens.
- **SWE-Marathon** : The evaluation was conducted by Abundant AI with 1M context length, max effort level, and 128K maximum output tokens.
