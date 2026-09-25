---
id: collect-250926-servers-hardware/servers-hardware/nvidia-nemotron-3-5-lightning-ollama-blog
title: "NVIDIA Nemotron 3.5 Lightning"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Apple", "Nvidia"]
dates: ["2026-08-11"]
keywords: ["nvidia", "agent", "agentic", "agents", "benchmarks", "claude", "datacenter", "inference", "open weights", "parameters", "reasoning", "speculative decoding"]
source: docs/RAG/clean4/nvidia-nemotron-3-5-lightning-ollama-blog.md
source_anchor: ""
source_lines: [1, 64]
sha256: edb9f2423dd7fe9aaf6bdded564706f853818fe28ca88512acf567b4bb627595
---

# NVIDIA Nemotron 3.5 Lightning

## August 11, 2026

NVIDIA Nemotron 3.5 Lightning is now available on Ollama, and it runs completely on your own device. It’s a 30 billion parameter (3B active) open model from NVIDIA built for agents that stay running: gathering context, calling tools, and working through multi-step tasks.

Nemotron 3.5 Lightning is made for agentic tasks such as reading a file, calling a tool, sorting a result, and retrying something that failed. Most of these steps don’t need a large model. At 3B active parameters per token, it’s built for local systems rather than the datacenter, and running locally means your data stays on your device.

### Model highlights

- **Runs where you work:** 30B total parameters with only 3B active per token, on a hybrid Mixture-of-Experts architecture. It runs locally on NVIDIA RTX PCs, NVIDIA RTX PRO workstations, NVIDIA DGX Spark and DGX Station, and in the datacenter and cloud.
- **Built for agent harnesses:** developed with the Nemotron Coalition and trained for the tools developers already use, across coding, tool calling, instruction following and multi-turn work.
- **1M token context:** a context length up to 1M, which leaves room for long tool histories across multi-turn workflows.
- **Optimized inference:** speculative decoding using multi-token prediction (MTP), DFlash or DSpark, offering up to 4x higher throughput than comparable open models.
- **Yours to customize:** an open model trained on open datasets. Post-train it for a specific task and run the result anywhere, from edge to datacenter.

### What you can build

Nemotron 3.5 Lightning excels on the following workloads:

- **Long-running personal assistants.** Email, calendar, projects and bookings. Running locally, the agent can use local context and none of it is sent elsewhere.
- **Coding sub-agents.** Running tests, searching the codebase and applying refactors, inside the harnesses you already use.
- **Security operations.** Enriching alerts, classifying incidents, querying logs, correlating indicators and preparing structured findings for analysts.
- **A local tier alongside the cloud.** Nemotron 3.5 Lightning handles the high-volume steps locally, and a larger hosted model picks up the few that need one. Same CLI, same API.
- **A specialist you train yourself.** Open weights and open datasets, so you can post-train Nemotron 3.5 Lightning for one narrow job and run the result locally.

### Get started

Download Ollama, then run Nemotron 3.5 Lightning with your tool of choice.

**General chat**

```
ollama run nemotron-3.5-lightning
```
**Claude Code**

```
ollama launch claude --model nemotron-3.5-lightning
```
**OpenClaw**

```
ollama launch openclaw --model nemotron-3.5-lightning
```
**Hermes Agent**

```
ollama launch hermes --model nemotron-3.5-lightning
```
**OpenCode**

```
ollama launch opencode --model nemotron-3.5-lightning
```
For users on Apple silicon, Ollama offers the model with state-of-the-art performance: `nemotron-3.5-lightning:30b-mlx`.

See more integrations on the model page.

The same pattern works for models running in Ollama’s cloud, so an agent can send an individual step to a larger model without changing anything else.

### Benchmarks

Nemotron 3.5 Lightning offers 4x higher throughput and 30% faster task completion time compared to other leading open models of similar size and offers leading accuracy across agentic, coding and reasoning tasks. For agents that stay running, throughput is the number that matters most: more steps per minute means long tasks finish. Full results and test configurations are in NVIDIA’s launch blog.
