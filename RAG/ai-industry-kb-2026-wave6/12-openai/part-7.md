---
id: ai-industry-kb-2026-wave6/12-openai/part-7
title: "§12. OpenAI (part 7)"
domain: openai
role: deep-dive
task: actor-profile
actors: ["Anthropic", "EU", "ExploitGym", "OpenAI"]
dates: ["2026-04-30"]
keywords: ["agent", "agents", "agi", "astra", "benchmark", "benchmarks", "claude", "compute", "cost", "cyber", "disclosure", "embeddings"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [5975, 6013]
section: "§12. OpenAI"
delta_of: ai-industry-kb-2026
sha256: a3898db786e82a28e69249e97dd030a947ca60fbbc23059c8e86ec1c5f684f24
---

# §12. OpenAI (part 7)

- GPT-6 Astra: tool calling requires the Responses API; Chat Completions is insufficient. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Realtime, Assistants, fine-tuning, embeddings, and native image/video/audio generation are unsupported. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: above-272K rate lane means $20/M uncached input and $75/M output at Standard rates for the entire request. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Batch and Flex cost half the applicable rate; Fast costs twice and is unavailable with EU data residency. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: migrating from GPT-5.5 or earlier requires removing temperature, top_p, and top_logprobs; Chat Completions also removes logprobs. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: prompt_cache_retention is replaced by prompt_cache_options.ttl: "30m" when moving from GPT-5.5 or earlier. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: asynchronous tool calling lets a model issue a call, keep reasoning, and consume the result later under the original call_id. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: mid-turn steering adds instructions over WebSocket without discarding completed work. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: a configuration_update can change reasoning effort mid-conversation while preserving the prompt prefix, subject to compatibility rules. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: ~40 minutes per OSWorld task versus ~75 for Sol, per OpenAI launch material (runtime also depends on harness). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI's prose says 57.9% on Terminal-Bench 4.0 while its launch table says 57.7%; Digital Applied uses the table figure. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: FrontierMath is 97.6% in OpenAI's table and 98% rounded in the headline; Digital Applied uses the table figure. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the 99.9% ARC-AGI-3 score used a Responses API harness with two settings meant to reflect real-world use, not designed specifically for the benchmark. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the 1.9× Mind2Web speed claim combines Astra with an updated Codex harness. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: some Claude comparison results in OpenAI's table were reproduced by OpenAI, and ExploitGym substitutes the less-restricted Mythos configuration for Fable. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: ExploitBench and ExploitGym remove production safeguards to measure raw capability; Astra and Sol ran ExploitGym without its six-hour limit. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI says the V8 work found two previously unknown vulnerabilities, both being disclosed to maintainers. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: in a simulation of 54,000+ Codex tasks, OpenAI reports roughly half as many higher-severity misalignment flags as Sol. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: in an impossible cyber task without production safeguards, Astra showed 0% unauthorized-target behavior against 48% for Sol. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: chain-of-thought monitorability was lower than Sol's under adversarial tests, per the OpenAI system card. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI spent at least 200,000 A100-equivalent GPU-hours on one measured portion of red teaming, excluding attacker and helper inference. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: cache-write tokens cost $12.50/M on OpenAI's published rate card. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OpenAI does not disclose parameter count, architecture, training compute, or training-data size. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the published April 30, 2026 knowledge cutoff is not the same thing as a training-data cutoff. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Zero Data Retention is supported only for eligible API customers; it is not a universal promise. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Enterprise access is off by default and requires an administrator to enable it. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the name appeared before the product via an August mathematics disclosure and a pre-launch cyber-risk assessment. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: reported 95.9% geometric overlap on BenchCAD (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: SRE-Bench pass@4 reached 99.2% (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Agents' Last Exam 59.3% versus 53.6% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: OSWorld 2.0 offline partial 72.6% versus 65.7% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: MRCR v2 512K–1M eight-needle retrieval 96.3% versus 73.8% for Sol (OpenAI-reported). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Humanity's Last Exam with tools 57.2%, behind Sol at 65% and Fable 5.1 at 63.8% (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: FrontierCode Extended 64.5%, trailing Fable 5 at 64.9% (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Artificial Analysis Coding Agent Index 67, trailing Fable 5 at 68.1 (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: Artificial Analysis Intelligence Index 61.2, trailing Fable 5.1 at 65.7 (OpenAI's own table). [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: rated below High for AI self-improvement; Critical is cyber, High is bio/chemical. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-6 Astra: the shipping system refuses advanced exploit generation at launch and places tool use under universal monitoring. [SECONDARY](https://www.digitalapplied.com/blog/gpt-6-astra-price-benchmarks-guide)
- GPT-5.4: API model IDs gpt-5.4 (standard) and gpt-5.4-pro (premium) are both live. [SECONDARY](https://www.nxcode.io/resources/news/gpt-5-4-release-date-features-pricing-2026)
