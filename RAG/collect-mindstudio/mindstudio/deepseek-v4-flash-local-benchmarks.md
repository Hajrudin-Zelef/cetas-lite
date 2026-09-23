---
id: collect-mindstudio/mindstudio/deepseek-v4-flash-local-benchmarks
title: "DeepSeek V4 Flash: The Cheapest Frontier-Level Open Model Yet"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agentic", "agi", "benchmark", "benchmarks", "cost", "gemini", "glm", "kv cache", "nvidia", "opus 4", "parameters"]
source: docs/RAG/Collect RAG/02_mindstudio/deepseek-v4-flash-local-benchmarks.md
source_anchor: ""
source_lines: [1, 51]
sha256: 7c29b522bf722afa7d88b3c7037e8c615ccc67022a650139011b1e459e2b635d
---

# DeepSeek V4 Flash: The Cheapest Frontier-Level Open Model Yet

## Metadata

- **Source** : https://www.mindstudio.ai/blog/deepseek-v4-flash-local-benchmarks
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article analyzes **DeepSeek V4 Flash**, a post-trained update to DeepSeek's existing V4 Flash preview model — not a new architecture or a larger parameter count. At **284 billion parameters**, it now scores close to much bigger frontier models on several benchmarks while staying priced near the bottom of the market: around **2 cents per million input tokens and ~30 cents per million output tokens**, placing it near the front of the cost-versus-intelligence curve for teams running agentic coding workloads at scale.

The headline number is the jump on the **DeepSweep agentic coding benchmark**: from **7% under the previous V4 Flash preview to 54%** under the new release — a categorical shift driven entirely by post-training, not a new model. DeepSeek reports the new model **beats GLM 5.2 on nearly every published benchmark**, despite GLM 5.2 running at nearly three times the parameter count. On the **Artificial Intelligence Index** (a composite measure), the jump puts V4 Flash right next to **Gemini 2.6 Flash**, a model released only days earlier.

The caveat: DeepSeek ran its own model under its **own optimized harness** (the scaffolding managing tool calls, retries, context handling, and task execution). Other models in the comparison table likely didn't have access to that same harness, so results aren't a clean head-to-head. As reference, OpenAI's own ARC-AGI study showed switching from the standard public harness to an internally optimized harness moved the score from ~13% to ~40% — roughly triple, with no model change. So some of the reported gain is model, and some is harness.

On DeepSeek's internal benchmarks — **DS-Bench Full Stack** and **DS-Bench Hard** — which are harder to game because they're not public, V4 Flash improved but **still lagged behind Opus 4.8**, giving a more honest read on its ceiling. On **Terminal Bench 2.1**, widely referenced for agentic coding, it scores comparably to some of the best models at a fraction of their per-token cost. It supports the **responses API** (usable directly in tools like Codex) and open-source harnesses like **Open Code**, positioning it as a practical "executor" model paired with a larger planning model rather than a single frontier system.

Local deployment: at **4-bit precision** it needs ~**168 GB of VRAM** (e.g., two networked Nvidia DGX Spark units, each ~115 GB usable VRAM); at **3-bit** ~**110 GB**. These cover model loading only; usable context demands more. Alternative engines using SSD offloading, KV cache manipulation, and mixed-precision weight loading can run the architecture on ~**128 GB of VRAM**. Local testing on a two-DGX-Spark setup produced **~20–30 tokens per second**, usable for agentic coding. The model is **text-only** with no native image or vision support.

## Key points

- V4 Flash is a post-training update of the V4 Flash preview; 284B parameters, unchanged architecture.
- DeepSweep agentic coding jumped from 7% (preview) to 54%; beats GLM 5.2 on nearly every reported benchmark despite GLM 5.2 being ~3x larger.
- Pricing ~2¢/M input, ~30¢/M output tokens; Pareto frontier of cost vs intelligence near Gemini 2.6 Flash.
- Comparisons are harness-dependent; OpenAI's ARC-AGI harness study shows scores can triple from harness alone.
- On internal DS-Bench Full Stack / DS-Bench Hard, V4 Flash still trails Opus 4.8.
- Runs locally: ~168 GB VRAM at 4-bit, ~110 GB at 3-bit; ~20–30 tok/s on two DGX Sparks.
- Text-only; no native vision/image support.

## Technical data / figures

| Item | Detail |
|---|---|
| Parameters | 284B (post-trained update of V4 Flash preview) |
| DeepSweep | 7% (preview) → 54% |
| vs GLM 5.2 | beats on nearly every published benchmark (GLM 5.2 ~3x parameters) |
| AI Index | next to Gemini 2.6 Flash |
| API pricing | ~$0.02/M input, ~$0.30/M output |
| VRAM (4-bit) | ~168 GB |
| VRAM (3-bit) | ~110 GB |
| Local throughput | ~20–30 tok/s (two DGX Sparks) |
| Internal evals | DS-Bench Full Stack, DS-Bench Hard (still trails Opus 4.8) |
| Modality | text-only |

## Why this source matters for the RAG

It quantifies the cost-versus-intelligence positioning of a frontier-adjacent open model and its real local hardware footprint, key data for comparing cloud vs local economics. It also gives a rigorous treatment of harness effects on benchmark scores — an important analytical caveat when evaluating any model's published results.
