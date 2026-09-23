---
id: collect-mindstudio/mindstudio/nvidia-switchyard-local-model-router
title: "What Is NVIDIA SwitchYard? The Open-Source Local AI Model Router"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Cerebras", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["nvidia", "agent", "agentic", "benchmark", "compute", "cost", "gpu", "latency", "memory", "moe", "prefill", "throughput"]
source: docs/RAG/Collect RAG/02_mindstudio/nvidia-switchyard-local-model-router.md
source_anchor: ""
source_lines: [1, 53]
sha256: 5d7fd394e14cd42fcd2142d964ecbbb33738881fd1cd85e4bf30fa8e6c168192
---

# What Is NVIDIA SwitchYard? The Open-Source Local AI Model Router

## Metadata

- **Source** : https://www.mindstudio.ai/blog/nvidia-switchyard-local-model-router
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

**NVIDIA SwitchYard** is an open-source routing library that sits between an AI agent and the models it calls, deciding per request which model should handle the job. It is not a model itself but infrastructure: a classifier, a proxy for translating between different API formats, and a set of routing strategies that let a team mix local models with frontier APIs based on task difficulty. It builds on an earlier project called **RouteLLM** and adds session state tracking, action traces, and multiple routing algorithms out of the box.

The article explains why agent systems need a router: most builders default to the most capable model for every task, which breaks down when an agent takes hundreds of steps (planning, tool calls, retrieval, summarization, error handling, output generation). There are four reasons teams adopt routing: **cost** (simple tasks like classification or entity extraction go to cheaper models), **speed** (faster models keep long sequential pipelines moving — e.g., OpenAI's planned deployment on Cerebras hardware), **specialization** (fine-tuned/purpose-built models for narrow jobs), and **privacy** (keeping sensitive data such as hostnames, IPs, or account numbers on local infrastructure, sending only sanitized data to external APIs).

SwitchYard's core building blocks are **profiles** (which router mode a route uses) and **targets** (available models, labeled e.g. "capable" for a frontier API and "efficient" for a local model, plus the serving provider). It offers four routing strategies:

- **Random routing** — splits traffic at a fixed ratio; useful for A/B tests or cost/quality baselines.
- **LLM classifier routing** — a secondary model reads the incoming request and picks the tier before the task runs; good for domain-based routing (coding vs. general writing).
- **Stage routing** — skips extra classification and reads signals already in the conversation (failing tool calls, successful code edits) to decide escalation; nearly free to run.
- **Escalation routing** — starts on the cheaper model with a judge model watching; on real trouble it bumps the session to a stronger model and keeps it there, rather than re-evaluating every turn.

A key operational point: re-routing on every conversational turn is expensive and breaks prompt caching, so the general pattern is to decide once near the start of a task and stick with it. Costs differ per strategy: random is free, LLM classifier adds one call per query, stage routing is near-free, and escalation costs one judge call per session. SwitchYard also supports a "tunable" prefill router — a trained model that reads early signals and predicts success likelihood, blending prediction with cost/latency targets (a learned policy).

NVIDIA claims routing open and proprietary models together can deliver **~50% faster responses and 25% better token efficiency**, though this is the company's framing rather than an independently verified benchmark. Most routing decision logic runs on **CPU**, so the router does not compete with models for GPU memory. A demonstrated setup processed sensitive network-alert data (hostnames, IPs) locally with a small model, forwarding only sanitized summaries to a larger API model for root-cause analysis — solving privacy and cost at once. The tradeoff: SwitchYard provides plumbing, not policy — teams must still define which tasks are simple enough for local models and where the cost/quality line sits. It pairs well with fast local models like **Nemotron 3.5 Lightning**, a 30B MoE model NVIDIA released alongside it.

## Key points

- SwitchYard is a routing layer (not a model) built on RouteLLM, deciding which model handles each agent step.
- Four strategies ship out of the box: random, LLM classifier, stage router, and escalation router.
- An SDK translates between OpenAI-style, Anthropic-style, and the newer responses API, easing provider switches.
- NVIDIA claims ~50% faster responses and 25% better token efficiency from hybrid routing (company framing, not independently verified).
- Routing addresses four needs: cost, speed, specialization, and privacy.
- Most routing logic runs on CPU, so it does not compete for GPU memory.
- Nemotron 3.5 Lightning (30B MoE) is the companion fast local model example.

## Technical data / figures

| Item | Detail |
|---|---|
| Base project | RouteLLM |
| Routing strategies | random, LLM classifier, stage, escalation (plus tunable prefill router) |
| API formats | OpenAI-style, Anthropic-style, responses API |
| Claimed benefits | ~50% faster responses, 25% better token efficiency |
| Compute for routing | CPU (most configurations) |
| Companion model | Nemotron 3.5 Lightning, 30B MoE, high-throughput/low-latency |
| Cost per strategy | random: free; LLM classifier: +1 call/query; stage: near-free; escalation: 1 judge call/session |

## Why this source matters for the RAG

It documents NVIDIA's open-source hybrid routing infrastructure, a key pattern for combining local models with frontier APIs to cut cost and preserve privacy in agentic workflows. It provides concrete routing strategies, their cost profiles, and the operational caveat about prompt caching, which is central to designing local/cloud hybrid systems.
