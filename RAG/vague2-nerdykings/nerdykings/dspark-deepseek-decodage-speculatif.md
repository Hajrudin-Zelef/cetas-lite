---
id: vague2-nerdykings/nerdykings/dspark-deepseek-decodage-speculatif
title: "DSpark : DeepSeek Accélère Son Modèle De 78% Sans Le Toucher"
domain: nerdykings
role: reference
task: article
actors: ["Alibaba", "China", "DeepSeek"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agentic", "cost", "gpu", "gpus", "inference", "kv cache", "license", "memory", "qwen", "speculative decoding"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/dspark-deepseek-decodage-speculatif.md
source_anchor: ""
source_lines: [1, 52]
sha256: 5630abadc3c08503ff13f0b71dfb68b28e6455f8225faf77ba35dac44987d011
---

# DSpark : DeepSeek Accélère Son Modèle De 78% Sans Le Toucher

## Metadata

- **Source** : https://www.nerdykings.com/blog/dspark-deepseek-decodage-speculatif.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article explains DSpark, a DeepSeek optimization that speeds up their V4 Pro model by up to 78% without a new model, retraining, or touching weights — the fifth "optimize instead of scale" move in six months. It first explains why LLMs are slow: autoregressive generation produces one word at a time, and each new word requires re-reading the entire KV cache, so longer answers cost more per word. During decoding, GPUs run at only 20-30% capacity, waiting on memory.

Speculative decoding is the base idea: a small fast "draft" model proposes a burst of 8 words, and the big "target" model verifies all 8 in a single parallel pass. If validated, 8 words are generated for the price of one; if rejected from word 5, the first 4 are kept and the process restarts. Existing systems had a structural flaw: autoregressive drafts (Eagle 3) are accurate but slow to generate; parallel drafts (10 Flash) are fast but incoherent, collapsing acceptance rates on later words.

DSpark's key innovation is a "semi-autoregressive" architecture with a Markov head — a mini-layer just before final word selection that looks only at the immediately preceding word (not the whole context) to correct local errors. Technically, this Markov head uses a rank-256 factorization, compressing the operation so it costs almost nothing. The draft becomes much cleaner and is accepted far more often. The second innovation is confidence-planned verification: a confidence head scores each draft word, and an algorithm called Sequential Temperature Scaling recalibrates those scores (neural networks are overconfident), reducing calibration error from 3-8% to about 1%. The system monitors GPU load in real time: when free, it verifies large blocks; when saturated, it shortens the block and sends only the most likely accepted words.

Production results: +57% to +78% generation speed per user on DeepSeek V4 Pro, +60% to +85% on V4 Flash, zero quality degradation. On Qwen 3 4B, DSpark accepts 26-31% more words than Eagle 3 and 16-18% more than 10 Flash. DeepSeek open-sourced everything under MIT as "DeepSpec." The article notes DeepSeek's pattern of six releases in six months (MHC, Engram, DeepSeek V4, Dual Path, DSpark), and two remaining blind spots: weak long agentic tasks (Terminal Bench 2.0: 67.9% vs GPT 5.5's 82.7%) and a 94% hallucination rate on unknown factual questions.

## Key points

- DSpark speeds up DeepSeek V4 Pro by +57% to +78% with no new model or retraining.
- Autoregressive decoding is memory-bandwidth-bound; GPUs idle at 20-30% capacity.
- Semi-autoregressive architecture with a rank-256 factorized Markov head cleans up parallel drafts.
- Confidence-planned verification uses a confidence head + Sequential Temperature Scaling (error 3-8% → ~1%).
- Production: +60-85% on V4 Flash, zero quality degradation.
- Qwen 3 4B: +26-31% accepted words vs Eagle 3, +16-18% vs 10 Flash.
- Open-sourced under MIT as "DeepSpec."
- Weaknesses remain: long agentic tasks (67.9% Terminal Bench 2.0) and 94% hallucination on unknown facts.

## Technical data / figures

| Item | Value |
|---|---|
| Speed gain V4 Pro | +57% to +78% |
| Speed gain V4 Flash | +60% to +85% |
| Quality degradation | 0% |
| Qwen 3 4B vs Eagle 3 | +26-31% accepted words |
| Qwen 3 4B vs 10 Flash | +16-18% accepted words |
| Markov head factorization rank | 256 |
| Calibration error (before → after) | 3-8% → ~1% |
| Draft burst size | 8 words |
| Terminal Bench 2.0 V4 Pro | 67.9% |
| Terminal Bench 2.0 GPT 5.5 | 82.7% |
| Hallucination rate (unknown facts) | 94% |
| License | MIT (DeepSpec) |

## Why this source matters for the RAG

This source is a detailed technical explainer of speculative decoding and DeepSeek's production optimization, offering concrete numbers and architecture details. It is highly relevant for RAG content on LLM inference efficiency, KV cache optimization, and the open-source Chinese AI ecosystem.
