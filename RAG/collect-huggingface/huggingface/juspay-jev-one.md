---
id: collect-huggingface/huggingface/juspay-jev-one
title: "jev-one - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Nvidia", "SGLang"]
dates: ["2026-09-23"]
keywords: ["apache", "blackwell", "distribution", "gpus", "inference", "license", "memory", "moe", "nvidia", "parameters", "prefill", "qwen"]
source: docs/RAG/Collect RAG/03_huggingface/juspay-jev-one.md
source_anchor: ""
source_lines: [1, 52]
sha256: d0bc205b197d6df7430e946eba809c061defbf7800a257a3fdca63c4242f59e7
---

# jev-one - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/juspay/jev-one
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

JevOne (`jev-one`) is a post-trained version of Qwen/Qwen3.6-35B-A3B developed by Juspay Technologies for typed decision tasks. It is served through a TypeSafe-compatible `/v1/systemone` API rather than as a general chat model. The base checkpoint is Qwen3.6-35B-A3B, a Mixture-of-Experts causal language model with 35 billion total parameters and approximately 3 billion activated parameters per token. It is released in BF16 precision with fully merged weights (no adapter loading or merging required), under the Apache License 2.0. The server interface accepts a state and a map of typed questions with three output modes: `noul` (binary probability), `choice` (categorical decision plus full probability distribution), and `score` (expected ordinal score plus full probability distribution). The serving layer performs deterministic single-token candidate readout, forward and reverse option-order evaluation, probability calibration, and schema conversion; Juspay emphasizes that this serving layer is part of the released inference configuration and must be used for reproducible results. The quick-start procedure pins revision `4de0db772d718134bd32fe5d4c8da77d97a36240`, downloads and SHA256-verifies a serving bundle, extracts it, and runs on two GPUs with tensor parallelism 2, exposing the API at `http://127.0.0.1:49001/v1/systemone`. The setup requires Linux x86-64, the Hugging Face CLI, Docker Engine with Docker Compose v2, the NVIDIA Container Toolkit, and approximately 120 GB of free disk space. JevOne was self-evaluated on the public tiers of JEVBench using harness commit `fd51755eb0c0b546ca206d764faf3302feca913e` and the `typesafe` adapter on 2× NVIDIA RTX PRO 6000 Blackwell Server Edition GPUs (96 GB each). Results: Easy tier 48/48 correct (accuracy 1.0000, Brier 0.0018, ECE 0.0257, p50 0.0766s); Original tier 70/72 correct (accuracy 0.9722, Brier 0.0897, ECE 0.1285); Hard public tier 86/111 correct (accuracy 0.7748, macro accuracy 0.8033, Brier 0.3460, ECE 0.0579, p50 0.1361s, p95 0.3887s). Operational success, coverage, schema validity, and strict schema validity were 1.0000 across all three tiers; the author notes these are self-run public-tier results, not an official JEVBench rank. The validated runtime pins a specific SGLang image (sha256:6bcaa47db52f78ce0d67863b8b2431221b79bc23204a80cad757fa819d00e921), tensor parallelism 2, maximum prefill of 250,000 tokens, and a static memory fraction of 0.85. Model files occupy ~66 GB. The service needs no API key when bound to localhost, but remote deployments must add authentication, TLS, rate limits, and request-size limits at the ingress layer. The hub reports 35B params (BF16) and 757 monthly downloads.

## Key points

- JevOne is a post-trained Qwen3.6-35B-A3B for typed decision tasks, served via a TypeSafe-compatible `/v1/systemone` API.
- MoE causal LM: 35B total parameters, ~3B activated per token, BF16, fully merged weights.
- Three typed output modes: `noul` (binary probability), `choice` (categorical + distribution), `score` (ordinal + distribution).
- Serving layer does deterministic single-token readout, forward/reverse option-order evaluation, calibration, and schema conversion.
- Apache 2.0 license; base checkpoint Qwen/Qwen3.6-35B-A3B.
- Self-run JEVBench public tiers: Easy 1.0000 accuracy, Original 0.9722, Hard public 0.7748 (macro 0.8033).
- Validated runtime: SGLang pinned image, TP=2, 2× RTX PRO 6000 Blackwell (96 GB), max prefill 250,000, static memory fraction 0.85.
- Requires Linux x86-64, Docker Compose v2, NVIDIA Container Toolkit, ~120 GB free disk; model files ~66 GB.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | juspay |
| Model name | jev-one |
| Base model | Qwen/Qwen3.6-35B-A3B |
| Architecture | Mixture-of-experts causal LM (qwen3_5_moe) |
| Total params | 35B |
| Activated params | ~3B per token |
| Precision | BF16 (fully merged weights) |
| License | Apache 2.0 |
| Task | Typed decision (noul / choice / score) |
| Serving API | `/v1/systemone` (TypeSafe-compatible) |
| Validated runtime | SGLang pinned image; TP=2; 2× RTX PRO 6000 Blackwell 96GB |
| Max prefill tokens | 250,000 |
| Static memory fraction | 0.85 |
| Disk (model files) | ~66 GB |
| JEVBench (Easy) | 48/48, acc 1.0000, Brier 0.0018, ECE 0.0257 |
| JEVBench (Original) | 70/72, acc 0.9722, Brier 0.0897, ECE 0.1285 |
| JEVBench (Hard public) | 86/111, acc 0.7748, macro 0.8033, Brier 0.3460 |
| Downloads/month | 757 |

## Why this source matters for the RAG

This card documents a specialized post-trained MoE model for calibrated, schema-typed decision tasks, distinct from general chat models, including its pinned serving stack and JEVBench results. It is a key reference for decision/classification use cases, probability calibration, and enterprise deployment patterns.
