---
id: collect-benchmarks/benchmarks/mimo-v2-6-flash
title: "MiMo V2.6 Flash — pricing, benchmarks & speed"
domain: benchmarks
role: reference
task: benchmark
actors: ["Anthropic", "DeepSeek", "Meta", "Xiaomi"]
dates: ["2026-09-23"]
keywords: ["benchmark", "benchmarks", "pricing", "agent", "agentic", "claude", "context window", "cost", "deepseek", "latency", "multimodal", "muse"]
source: docs/RAG/Collect RAG/06_benchmarks/mimo-v2-6-flash.md
source_anchor: ""
source_lines: [1, 55]
sha256: eb3bb86aff491a4af1cf1aa15d1cdeefc055d31a003845eb380b87d28b59734d
---

# MiMo V2.6 Flash — pricing, benchmarks & speed

## Metadata

- **Source** : https://commandcode.ai/models/mimo-v2-6-flash
- **Site** : CommandCode
- **Type** : Benchmark
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Command Code page profiles Xiaomi's MiMo V2.6 Flash (`xiaomi/mimo-v2.6-flash`), an efficient multimodal agentic coding model with a 1.05M-token context window. The model is not yet scored: both its Intelligence Index and Coding Index are listed as "not yet scored." The page positions it as the budget tier of the MiMo V2.6 line and compares it against sibling and rival models on price, context, and scored intelligence/coding metrics. Benchmarks are sourced from Artificial Analysis (v4.3).

Pricing is $0.14 per million input tokens and $0.28 per million output tokens, with cache reads at $0.0028 per million and a blended rate of $0.18 per million. Command Code emphasizes that in an agent loop most input is cache-read, yielding an effective input rate of roughly $0.04 per million. It sits alongside MiMo V2.6 Pro ($0.43/$0.87) and MiMo V2.6 Pro UltraSpeed ($4.35/$8.70, a low-latency serving tier), all with ~1.05M context. In the comparison table, only rivals carry scores: Muse Spark 1.3 Contributor leads the shown set with Intelligence 48.2 and Coding 75.8 at $0.10/$0.20; DeepSeek V4 Pro scores 36.3/68.8 at $0.66/$1.98; DeepSeek V4 Flash 35/69.1 at $0.15/$0.60; MiMo V2.5 Pro 26.4/60.2 at $0.43/$0.87; MiMo V2.5 22.3/56.8 at $0.14/$0.28.

The page provides a usage calculator and real-task cost estimates. A typical request (800 fresh input, 180 output, 50K cache-read tokens) costs ~$0.0003, splitting ~37% fresh input, 17% output, 46% cache reads; $10 in credits supports ~33K such requests. End-to-end coding task costs: quick lookup $0.0004; review a 500-line PR $0.0038; fix a bug via agent loop $0.01; refactor a module $0.02; full-repo agent run $0.04. The same tasks on Muse Spark 1.3 Contributor cost $0.0003, $0.0027, $0.0072, $0.01, $0.03, and on Claude Haiku 4.5 cost $0.0065, $0.04, $0.12, $0.21, $0.49.

Availability: MiMo V2.6 Flash is offered on the Go plan and above. Users switch via `cmd --model xiaomi/mimo-v2.6-flash` or `/model` mid-session without losing context. The page markets Command Code as an AI coding agent that "learns your taste," starting at $1.

## Key points

- MiMo V2.6 Flash is Xiaomi's efficient multimodal agentic coding model with a 1.05M-token context window.
- Intelligence and Coding indices are "not yet scored" on Command Code; benchmark data sourced from Artificial Analysis v4.3.
- Pricing $0.14/M input, $0.28/M output, $0.0028/M cache read; effective agent-loop input ≈ $0.04/M.
- Comparison set: Muse Spark 1.3 Contributor (48.2/75.8), DeepSeek V4 Pro (36.3/68.8), DeepSeek V4 Flash (35/69.1), MiMo V2.5 Pro (26.4/60.2), MiMo V2.5 (22.3/56.8).
- Real task costs from $0.0004 (lookup) to $0.04 (full-repo agent run), cheaper than Claude Haiku 4.5.
- Available on Go plan and above; switchable mid-session via `/model`.

## Technical data / figures

| Model | Intelligence | Coding | Speed | Input $/M | Output $/M | Blended $/M | Context |
|-------|-------------|--------|-------|-----------|------------|-------------|---------|
| Muse Spark 1.3 Contributor | 48.2 | 75.8 | 418.4 | $0.10 | $0.20 | $0.13 | 1.05M |
| MiMo V2.5 Pro | 26.4 | 60.2 | 43.4 | $0.43 | $0.87 | $0.54 | 1M |
| MiMo V2.5 | 22.3 | 56.8 | 46.2 | $0.14 | $0.28 | $0.18 | 1M |
| MiMo V2.6 Flash | — | — | — | $0.14 | $0.28 | $0.18 | 1.05M |
| MiMo V2.6 Pro | — | — | — | $0.43 | $0.87 | $0.54 | 1.05M |
| MiMo V2.6 Pro UltraSpeed | — | — | — | $4.35 | $8.70 | $5.44 | 1.05M |
| DeepSeek V4 Pro | 36.3 | 68.8 | 65.3 | $0.66 | $1.98 | $0.99 | 1M |
| DeepSeek V4 Flash | 35 | 69.1 | 236.4 | $0.15 | $0.60 | $0.26 | 1M |

| Task (tokens in · out) | MiMo V2.6 Flash | Muse Spark 1.3 | Claude Haiku 4.5 |
|------------------------|-----------------|----------------|------------------|
| Quick lookup (8K · 1K) | $0.0004 | $0.0003 | $0.0065 |
| Review 500-line PR (60K · 4K) | $0.0038 | $0.0027 | $0.04 |
| Fix a bug, agent loop (180K · 12K) | $0.01 | $0.0072 | $0.12 |
| Refactor a module (320K · 20K) | $0.02 | $0.01 | $0.21 |
| Full-repo agent run (900K · 45K) | $0.04 | $0.03 | $0.49 |

## Why this source matters for the RAG

It supplies concrete pricing, cache economics, and real-task cost data for a budget agentic coding model, complementing score-only benchmark pages. The comparison table and per-task cost breakdowns enable cost-efficiency retrieval and head-to-head positioning against DeepSeek, Muse Spark, and Claude Haiku.

