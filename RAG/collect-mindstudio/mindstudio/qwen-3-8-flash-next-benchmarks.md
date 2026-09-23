---
id: collect-mindstudio/mindstudio/qwen-3-8-flash-next-benchmarks
title: "Qwen 3.8 Flash Next Benchmarks: Coding and Agentic Test Results"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Anthropic", "DeepSeek", "OpenAI", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: ["2026-09-23"]
keywords: ["agentic", "benchmark", "benchmarks", "qwen", "attention", "claude", "cost", "deepseek", "embedding", "embeddings", "fp8", "gguf"]
source: docs/RAG/Collect RAG/02_mindstudio/qwen-3-8-flash-next-benchmarks.md
source_anchor: ""
source_lines: [1, 57]
sha256: 2fa488da3cb65bf60cbe1766055fb6012a408619f148e242d546cc290c6110a6
---

# Qwen 3.8 Flash Next Benchmarks: Coding and Agentic Test Results

## Metadata

- **Source** : https://www.mindstudio.ai/blog/qwen-3-8-flash-next-benchmarks
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article reports independent KingBench test results for Qwen 3.8 Flash Next, Alibaba's experimental preview of the architecture it plans to use for Qwen 4. It is not a flagship release but a small, cheap, fast model built to validate new design choices before they are scaled up. On KingBench's fixed eight-question coding and agentic suite (each task scored out of 10, maximum 80 points), Qwen 3.8 Flash Next scored 56/80 (70%) against GLM 5.3 Flash's 63/80 (78.75%). The gap comes almost entirely from front-end and 3D rendering tasks, while the two models tied with perfect 10s on math and agentic pipeline tests.

Architecturally, the model is a 125 billion parameter mixture-of-experts (MoE) model that activates only 6 billion parameters per token. It uses a hybrid attention system across 48 layers: three gated delta blocks (a form of linear attention) followed by one block of Qwen Sparse Attention (QSA), which operates at the microblock level to cut latency for long-context and agentic workloads. A notable choice is an n-gram embedding layer: roughly 51 billion of the 125 billion parameters are dedicated to bigram and trigram embeddings stored as fixed phrase lookups that can live in ordinary system RAM instead of GPU memory. MoE routing uses 512 experts with 11 activated per token, plus multi-token prediction and gated residuals. The model ships as a thinking model by default (can be turned off), with a native 262,000-token context extendable to 1 million via YaRN.

Alibaba's own benchmarks claim it beats Claude Opus 4.5 on SWE-bench Pro (62.5 vs 53.4), scores 91.7 on GPQA Diamond, and outperforms DeepSeek V4 Flash on coding. Vendor benchmarks are treated skeptically; independent testing is preferred. On the broader KingBench leaderboard, Flash Next outperforms GPT-5.6, Sonnet 5, and Grok 4.5 — strong for a model activating only 6B parameters per token. Full-size GLM 5.3 leads at 91.25%, while Qwen 3.8 Max scores 81.25%.

API pricing is aggressive: $0.16 per million input tokens and $0.47 per million output tokens, about 12x cheaper than Qwen 3.8 Max. The model is open weight under the Qwen community license, with GGUF quantizations available (roughly 72–74GB for 1-bit dynamic quants up to 94–111GB for 4-bit; official FP8 exists). Comfortable local operation requires 96–128GB of RAM; Ollama is the simplest path, Llama.cpp for an OpenAI-compatible local server, and vLLM/SGLang for production. Given the low API price, local running mostly makes sense for privacy or offline use.

## Key points

- Qwen 3.8 Flash Next is a 125B-parameter MoE preview model activating only 6B parameters per token.
- Scored 56/80 (70%) on KingBench's eight-task coding/agentic suite vs GLM 5.3 Flash's 63/80 (78.75%).
- Tied GLM 5.3 Flash on math and agentic tasks (perfect 10s); lost on 3D rendering and front-end polish, with a near-total failure on the folding table animation test.
- Uses hybrid attention (gated delta blocks + Qwen Sparse Attention) and an n-gram embedding layer storing ~51B parameters as phrase lookups in RAM.
- API pricing: $0.16/M input and $0.47/M output tokens, about 12x cheaper than Qwen 3.8 Max.
- Open weight (Qwen community license); GGUF quants available; 96–128GB RAM recommended for local use.
- Outperforms GPT-5.6, Sonnet 5, and Grok 4.5 on the broader KingBench leaderboard.
- Best suited as a cheap, fast backend for agentic pipelines, tool-calling, and long-context processing — not front-end/3D polish.

## Technical data / figures

| Item | Value |
|---|---|
| Total parameters | 125B (MoE) |
| Active parameters per token | 6B |
| Experts / active per token | 512 / 11 |
| Layers | 48 (3 gated delta blocks : 1 QSA block) |
| N-gram embedding parameters | ~51B (bigram/trigram) |
| Native context | 262,000 tokens (1M via YaRN) |
| KingBench score | 56/80 (70%) |
| GLM 5.3 Flash KingBench | 63/80 (78.75%) |
| SWE-bench Pro (vendor) | 62.5 vs Claude Opus 4.5's 53.4 |
| GPQA Diamond (vendor) | 91.7 |
| API input price | $0.16 / M tokens |
| API output price | $0.47 / M tokens |
| GGUF quant sizes | ~72–74GB (1-bit) to 94–111GB (4-bit) |
| Recommended local RAM | 96–128GB |

KingBench task breakdown: elevator simulation 5 vs 6; Three.js contact lens case 8 vs 7 (win); Three.js folding table 2 vs 7 (biggest gap); SVG panda 6 vs 8; bow and arrow 8 vs 8 (tie); hard math (answer 2460) 10 vs 10; agentic pipeline (Pandemma 2B) 10 vs 10; 3D wristwatch 7 vs 7.

## Why this source matters for the RAG

It provides concrete, independently verified benchmark and architecture data for a low-cost open-weight MoE model central to 2026 local/agentic deployment discussions. The detailed pricing, quantization sizes, and RAM requirements make it directly actionable for hardware and routing decisions in a local-AI knowledge base.

