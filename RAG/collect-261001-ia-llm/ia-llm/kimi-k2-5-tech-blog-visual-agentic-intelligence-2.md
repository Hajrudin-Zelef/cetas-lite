---
id: collect-261001-ia-llm/ia-llm/kimi-k2-5-tech-blog-visual-agentic-intelligence-2
title: "Kimi K2.5: Visual Agentic Intelligence"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "kimi", "agents", "agi", "benchmark", "benchmarks", "claude", "deepseek", "gemini", "opus 4", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/kimi-k2-5-tech-blog-visual-agentic-intelligence.md
source_anchor: ""
source_lines: [78, 138]
sha256: bbb2ba70cbb3a2dec0f55a93e1c747e507a3413be0b77059b4d2c39588f725c6
---

# Kimi K2.5: Visual Agentic Intelligence

K2.5 agent supports advanced tasks such as **adding annotations in Word, constructing financial models with Pivot Tables, and writing LaTeX equations in PDFs**, while scaling to long-form outputs like **10,000-word papers or 100-page documents**.

Tasks that once took hours or days now complete in minutes. Here are some examples:

## 4. Conclusion

Grounded in advances in coding with vision, agent swarms, and office productivity, Kimi K2.5 represents a meaningful step toward AGI for the open-source community, demonstrating strong capability on real-world tasks under real-world constraints. Looking ahead, we will push further into the frontier of agentic intelligence, redefining the boundaries of AI in knowledge work.

## Appendix

### Benchmark table

To reproduce official **Kimi-K2.5** benchmark results, we recommend using the official API. For third-party providers, refer to **Kimi Vendor Verifier (KVV)** to choose high-accuracy services. Details: https://www.kimi.ai/blog/kimi-vendor-verifier

## Footnotes

**1. General Testing Details**

- We report results for Kimi K2.5 and DeepSeek-V3.2 with thinking mode enabled, Claude Opus 4.5 with extended thinking mode, GPT-5.2 with xhigh reasoning effort, and Gemini 3 Pro with a high thinking level. For vision benchmarks, we additionally report results for Qwen3-VL-235B-A22B-Thinking.
- Unless otherwise specified, all Kimi K2.5 experiments were conducted with temperature = 1.0, top-p = 0.95, and a context length of 256k tokens.
- Benchmarks without publicly available scores were re-evaluated under the same conditions used for Kimi K2.5 and are marked with an asterisk (*).
- We could not evaluate GPT-5.2 xhigh on all benchmarks due to service stability issues. For benchmarks that were not tested, we mark them as "-".

**2.** **Text and Reasoning**

- HLE, AIME 2025, HMMT 2025 (Feb), GPQA-Diamond and IMO-AnswerBench were evaluated with a maximum completion budget of 96k tokens.
- Results for AIME and HMMT are averaged over 32 runs (avg@32); GPQA-Diamond over 8 runs (avg@8).
- For HLE, we report scores on the full set (text & image). Kimi K2.5 scores 31.5 (text) and 21.3 (image) without tools, and 51.8 (text) and 39.8 (image) with tools. The DeepSeek-V3.2 score corresponds to its text-only subset (marked with †) . Hugging Face access was blocked to prevent potential data leakage. HLE with tools uses simple context management: once the context exceeds a threshold, only the latest round of tool messages is retained.

**3.** **Tool-Augmented / Agentic Search**

- Kimi K2.5 was equipped with search, code-interpreter, and web-browsing tools for HLE with tools and all agentic search benchmarks.
- Except for BrowseComp (where K2.5 and DeepSeek-V3.2 used the discard-all strategy), no context management was applied, and tasks exceeding the supported context length were directly counted as failed.
- The test system prompts emphasize deep and proactive tool use, instructing models to reason carefully, leverage tools, and verify uncertain information. Full prompts will be provided in the technical report.
- Results for Seal-0 and WideSearch are averaged over four runs (avg@4).

**4.** **Vision Benchmarks**

- Max-tokens = 64k, averaged over three runs (avg@3).
- ZeroBench (w/ tools) uses max-tokens-per-step = 24k and max-steps = 30 for multi-step reasoning.
- MMMU-Pro follows the official protocol, preserving input order and prepending images.
- GPT-5.2-xhigh had ~10% failure rate (no output despite 3 retries), treated as incorrect; reported scores likely underestimate true performance.
- WorldVQA, a benchmark designed to evaluate atomic vision-centric world knowledge. Access WorldVQA at https://github.com/MoonshotAI/WorldVQA.
- OmniDocBench Score is computed as (1 − normalized Levenshtein distance) × 100, where a higher score denotes superior accuracy.

**5.** **Coding Tasks**

- Terminal-Bench 2.0 scores were obtained with the default agent framework (Terminus-2) and the provided JSON parser. In our implementation, we evaluated Terminal-Bench 2.0 under non-thinking mode. This choice was made because our current context management strategy for the thinking mode is incompatible with Terminus-2.
- For the SWE-Bench series of evaluations (including verified, multilingual, and pro), we used an internally developed evaluation framework. This framework includes a minimal set of tools—bash tool, createfile tool, insert tool, view tool, strreplace tool, and submit tool—along with tailored system prompts designed for the tasks. The highest scores were achieved under non-thinking mode.
- The score of Claude Opus 4.5 on CyberGym is reported under the non-thinking setting.
- All reported scores of coding tasks are averaged over 5 independent runs.

**6.** **Long-Context Benchmarks**

- AA-LCR: scores averaged over three runs (avg@3).
- LongBench-V2: identical prompts and input contexts standardized to ~128k tokens.

**7.** **Agent Swarm**

- BrowseComp (Swarm Mode): main agent max 15 steps; sub-agents max 100 steps.
- WideSearch (Swarm Mode): main and sub-agents max 100 steps.
