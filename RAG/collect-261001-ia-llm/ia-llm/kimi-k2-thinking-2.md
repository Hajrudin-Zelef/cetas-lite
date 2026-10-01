---
id: collect-261001-ia-llm/ia-llm/kimi-k2-thinking-2
title: "Introducing Kimi K2 Thinking"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Hugging Face", "OpenAI", "xAI"]
dates: []
keywords: ["kimi", "agent", "agentic", "benchmark", "benchmarks", "claude", "deepseek", "grok", "leaderboard", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/kimi-k2-thinking.md
source_anchor: ""
source_lines: [59, 75]
sha256: 9837f9c9273dab196fb87397e276eec2c32269bcd4c14830ebafc79198de7b52
---

# Introducing Kimi K2 Thinking

1. To ensure a fast, lightweight experience, we selectively employ a subset of tools and reduce the number of tool call turns under the chat mode on kimi.ai. As a result, chatting on kimi.ai may not reproduce our benchmark scores. Our agentic mode will be updated soon to reflect the full capabilities of K2 Thinking.
2. **Testing Details:** a. All benchmarks were evaluated at temperature = 1.0 and 256 k context length for K2 Thinking, except for SciCode, for which we followed the official temperature setting of 0.0.
b. HLE (no tools), AIME25, HMMT25, and GPQA were capped at a 96k thinking-token budget, while IMO-Answer Bench, LiveCodeBench and OJ-Bench were capped at a 128k thinking-token budget. Longform Writing was capped at a 32k completion-token budget.
c. For AIME and HMMT (no tools), we report the average of 32 runs (avg@32). For AIME and HMMT (with Python), we report the average of 16 runs (avg@16). For IMO-AnswerBench, we report the average of 8 runs (avg@8).
3. **Baselines:** a. GPT-5, Claude-4.5-sonnet, Grok-4 results and DeepSeek-V3.2 results are quoted from the GPT-5 post, GPT-5 for Developers post, GPT-5 system card, claude-sonnet-4-5, grok-4, deepseek-v3.2, the public Terminal-Bench leaderboard (Terminus-2), the public Vals AI leaderboard and the artificialanalysis. Benchmarks for which no available public scores were re-tested under the same conditions used for k2 thinking and are marked with an asterisk(*). For the GPT-5 test, we set the reasoning effort to high. 
b. The GPT-5 and Grok-4 on the HLE full set with tools are 35.2 and 38.6 from their official posts. In our internal evaluation on the HLE text-only subset, GPT-5 scores**41.7** and Grok-4 scores **38.6 ** (Grok-4’s launch cited**41.0** on the text-only subset). For GPT-5's HLE text-only w/o tool, we use score from Scale.ai, and the official GPT-5 score on the HLE full set (no tools) is 24.8. 
c. For IMO-AnswerBench: GPT-5 scored 65.6 in the benchmark paper. We re-evaluated GPT-5 with official API and obtained a score of 76.
4. **For HLE (w/ tools) and the agentic-search benchmarks:** a. K2 Thinking was equipped with search, code-interpreter, and web-browsing tools.
b. BrowseComp-ZH, Seal-0, FinSearchComp-T3 were run 4 times independently and the average is reported (avg@4).
c. The evaluation used o3-mini as judge, configured identically to the official HLE setting; judge prompts were taken verbatim from the official repository.
d. On HLE, the maximum step limit was 120, with a 48 k-token reasoning budget per step; on agentic-search tasks, the limit was 300 steps with a 24 k-token reasoning budget per step.
e. When tool execution results cause the accumulated input to exceed the model's context limit (256k), we employ a simple context management strategy that hides all previous tool outputs.
f. The web access to Hugging Face may lead to data leakage in certain benchmark tests, such as HLE. K2 Thinking can achieve a score of 51.3 on HLE without blocking Hugging Face. To ensure a fair and rigorous comparison, we blocked access to Hugging Face during testing.
5. **For Coding Tasks:** a. Terminal-Bench scores were obtained with the default agent framework (Terminus-2) and the provided JSON parser.
b. For other coding tasks, the result was produced with our in-house evaluation harness. The harness is derived from SWE-agent, but we clamp the context windows of the Bash and Edit tools and rewrite the system prompt to match the task semantics.
c. All reported scores of coding tasks are averaged over 5 independent runs.
6. **Heavy Mode** : K2 Thinking Heavy Mode employs an efficient parallel strategy: it first rolls out eight trajectories simultaneously, then reflectively aggregates all outputs to generate the final result. Heavy mode for GPT-5 denotes the official GPT-5 Pro score.
