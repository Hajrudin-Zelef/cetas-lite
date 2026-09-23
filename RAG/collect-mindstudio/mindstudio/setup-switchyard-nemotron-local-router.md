---
id: collect-mindstudio/mindstudio/setup-switchyard-nemotron-local-router
title: "Set Up a Local AI Router With SwitchYard and Nemotron Lightning"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Moonshot", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["benchmark", "claude", "context window", "cost", "kimi", "latency", "moe", "nvidia", "reasoning", "speculative decoding", "throughput", "tokens per second"]
source: docs/RAG/Collect RAG/02_mindstudio/setup-switchyard-nemotron-local-router.md
source_anchor: ""
source_lines: [1, 50]
sha256: 02385d366f3fe1eda8a7836c30d726c4accc53f341eb6ac813d570346c3e4d80
---

# Set Up a Local AI Router With SwitchYard and Nemotron Lightning

## Metadata

- **Source** : https://www.mindstudio.ai/blog/setup-switchyard-nemotron-local-router
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how to configure **NVIDIA's SwitchYard** router with a locally hosted **Nemotron 3.5 Lightning** model to cut API costs without losing task quality. A local AI router sits between an application and multiple language models, deciding in real time which model should handle a request — matching task complexity to model capability rather than sending everything to a frontier model. SwitchYard is an open-source router built on **RouteLLM** that adds a classifier, a proxy layer, and several routing strategies, and it can run entirely on your own hardware alongside a local model.

**Nemotron 3.5 Lightning** is a **30 billion parameter MoE model** from NVIDIA built for high throughput rather than raw intelligence. It has **30 total experts** and is a **hybrid of Mamba and transformer components**, preserving a **1 million token context window** more efficiently than a pure transformer. It uses a "latent MoE" approach: tokens are projected into a smaller latent space before being routed to active experts. It supports speculative decoding techniques including MTP, Dlash, and DSpark. On an AI capability index it lands **on par with GPT-OSS** but at meaningfully higher throughput. Locally on a **DGX Spark**, it produced ~**71 tokens per second** on a single stream, with multi-token prediction accepting ~70% of speculative guesses. In a head-to-head on a 1,400-token reasoning task, local Lightning finished in ~**20 seconds** vs. ~**50 seconds** for **Kimi K3** over an API.

Routing exists for four reasons: **cost, speed, specialization, and privacy**, each with different tradeoffs. SwitchYard's core primitives are **profiles** (overall routing behavior) and **targets** (available models, labeled e.g. "capable" and "efficient", pointing to API providers). A local Lightning on DGX Spark can sit alongside a hosted model like Kimi K3 as separate targets.

The four strategies: **Random** (fixed traffic split for A/B testing), **LLM classifier** (a second model reads the request and picks the tier — adds an extra model call per request), **stage router** (reads signals already present in the conversation, like failing tool calls — essentially free), and **escalation** (tasks start on the cheaper model, a judge model escalates on trouble, and the session stays on the stronger model). Cost matters: random and stage routing are essentially free; LLM classification adds a call per request; escalation costs one judge call until a session locks into its tier.

Economics: frontier-only setups pay a steep premium for the last few percentage points of completion. In one referenced internal comparison, **Claude Opus** completed ~80% of tasks in a benchmark at a cost of ~$180 for that run; pairing it with a workhorse model through a router achieved similar completion at substantially lower cost. A practical privacy workflow: Nemotron Lightning runs locally and generates a compact digest of each network alert (stripping hostnames, IPs, account numbers); only the sanitized digest goes to a hosted model like Kimi K3 for root-cause analysis. Routing decisions should happen **once per task, not per turn**, because per-turn routing breaks caching and adds classifier overhead.

## Key points

- SwitchYard (built on RouteLLM) adds a classifier, request proxy, and state tracking; it runs on your own hardware.
- Nemotron 3.5 Lightning: 30B MoE, 30 experts, hybrid Mamba+transformer, 1M-token context, latent MoE routing.
- Lightning is on par with GPT-OSS on capability but at higher throughput; ~71 tok/s on DGX Spark, ~20s vs Kimi K3's ~50s on a 1,400-token reasoning task.
- Four strategies: random, LLM classifier, stage router, escalation — with different cost profiles.
- Hybrid routing can approximate frontier-only completion quality at a fraction of cost (e.g., Opus ~80% completion at ~$180 vs router-based similar quality).
- Privacy pattern: local model sanitizes sensitive data; only digests go to the hosted API.
- Route once per task, not per turn, to preserve caching and avoid classifier overhead.

## Technical data / figures

| Item | Detail |
|---|---|
| Router | NVIDIA SwitchYard (open-source, built on RouteLLM) |
| Workhorse model | Nemotron 3.5 Lightning, 30B MoE, 30 experts |
| Architecture | Hybrid Mamba + transformer; "latent MoE"; speculative decoding MTP/Dlash/DSpark |
| Context window | 1M tokens |
| Capability index | on par with GPT-OSS, higher throughput |
| Throughput | ~71 tok/s single stream on DGX Spark; ~70% MTP acceptance |
| Latency test | 1,400-token reasoning: Lightning ~20s local vs Kimi K3 ~50s API |
| Strategies | random, LLM classifier, stage, escalation |
| Cost reference | Claude Opus ~80% completion ~$180; router pairing similar quality lower cost |

## Why this source matters for the RAG

It gives a concrete, config-level guide to building a local/cloud hybrid routing setup with an open-source router and a fast local MoE model, including throughput and latency figures. It quantifies the cost/quality economics of escalation routing and the privacy pattern of sanitizing sensitive data locally, both key concepts for hybrid AI architectures.
