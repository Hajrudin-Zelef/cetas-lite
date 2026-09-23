---
id: vague2-nerdykings/nerdykings/deepseek-dual-path-gpu
title: "DeepSeek Dual Path : Doubler Les Perfs IA Sans GPU"
domain: nerdykings
role: reference
task: article
actors: ["DeepSeek", "United States"]
dates: ["2026-09-23"]
keywords: ["deepseek", "gpu", "agent", "agentic", "compute", "decode", "gpus", "inference", "kv cache", "memory", "prefill", "throughput"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/deepseek-dual-path-gpu.md
source_anchor: ""
source_lines: [1, 50]
sha256: bb04f7e72d4e8636a23f244d51f4bd4da9960ae09c21c3ad5b1d8f7eb67c300f
---

# DeepSeek Dual Path : Doubler Les Perfs IA Sans GPU

## Metadata

- **Source** : https://www.nerdykings.com/blog/deepseek-dual-path-gpu.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article explains DeepSeek's "Dual Path" paper, which addresses the fact that AI servers' GPUs run at only 20-30% capacity much of the time — not because they are weak, but because they wait for data. It first distinguishes two LLM phases: prefill (the model reads the prompt at once, analyzes all tokens in parallel, computes each word's state into the KV cache; GPU usage 90-95%) and decode (generates one token at a time, re-reading the entire KV cache from memory repeatedly; low compute but huge memory bandwidth demand; GPU usage drops to 20-30%). Decode represents 95% of a request's lifetime: for a 300-token response, 20 ms of prefill versus 9 seconds of decode.

The industry first addressed this with disaggregation — separating prefill machines from decode machines — yielding 2-7x throughput gains. But this created a new problem in the agentic era. A classic chatbot has one question and one answer; an autonomous agent loops, executes code, calls tools, and reasons over dozens to hundreds of iterations, accumulating context. DeepSeek analyzed its own production traces: an average agentic session has 157 iterations with 32,000 tokens of cumulative context, but each turn adds only ~409 new tokens on average. That means 98.7% of the context already exists from previous turns. The KV cache is almost always on disk, so prefill spends its time fetching cache from storage rather than computing — and the prefill machine's storage network card saturates while decode machines' storage cards sit idle. The pipes are empty and global throughput collapses.

Dual Path's insight is simple: why should only the prefill machine fetch data from storage? It creates two simultaneous paths. Path A (classic): the prefill machine fetches data and sends results to the decode machine via the inter-GPU network. Path B (the detour): in parallel, the decode machine uses its own idle storage network cards to fetch the cache history directly from storage. Since the decode machine already has most of the historical cache after prefill, prefill only needs to send the small incremental cache of new tokens; decode merges and proceeds. To avoid contention with model generation traffic, model traffic always has priority (like an emergency lane on a highway), and Dual Path uses only remaining bandwidth. The system monitors machine load in real time and switches paths accordingly.

Production results: machine utilization rises from 40% to 80%, inference throughput doubles, time-to-first-token is 56% faster, with zero new GPUs purchased. The article situates this in a broader movement — HiCache (3-tier KV cache), Mooncake (Best Paper at FAST 2025), and 3FS (DeepSeek's open-source distributed file system) — and frames it as a software response to US GPU sanctions since 2022.

## Key points

- GPUs idle at 20-30% during decode because of memory-bandwidth bottleneck, not weak compute.
- Decode is 95% of a request's lifetime (300-token response: 20 ms prefill vs 9 s decode).
- Agentic sessions: 157 iterations, 32K cumulative context, but only ~409 new tokens/turn (98.7% already cached).
- Disaggregation (prefill/decode split) gave 2-7x throughput but saturated prefill storage network cards.
- Dual Path uses both prefill and decode storage network cards simultaneously (two paths).
- Model traffic has priority; the system dynamically switches paths based on load.
- Results: utilization 40% → 80%, throughput x2, first token 56% faster, no new GPU.
- Related systems: HiCache, Mooncake, 3FS.

## Technical data / figures

| Metric | Before | After |
|---|---|---|
| Machine utilization | 40% | 80% |
| Inference throughput | 1x | 2x |
| Time to first token | baseline | 56% faster |
| GPUs purchased | — | 0 |
| Prefill GPU usage | 90-95% | — |
| Decode GPU usage | 20-30% | — |
| Agentic session length | — | 157 iterations |
| Cumulative context | — | 32,000 tokens |
| New tokens per turn | — | ~409 (98.7% cached) |
| Prefill/decode time (300 tok) | 20 ms / 9 s | — |

## Why this source matters for the RAG

This source provides a rigorous infrastructure-level explainer of KV cache management and GPU utilization optimization, with concrete production metrics. It is valuable for RAG corpora on AI infrastructure, inference optimization, and the geopolitics of GPU access.
