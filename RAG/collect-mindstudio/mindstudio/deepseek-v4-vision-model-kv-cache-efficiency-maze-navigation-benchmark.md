---
id: collect-mindstudio/mindstudio/deepseek-v4-vision-model-kv-cache-efficiency-maze-navigation-benchmark
title: "DeepSeek V4 Vision Model: 10x KV-Cache Efficiency and 67% Maze Navigation vs GPT-5.4's 50%"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI", "OpenRouter"]
dates: ["2024-03", "2026-09-23"]
keywords: ["deepseek", "attention", "benchmark", "claude", "cost", "distillation", "gemini", "gpus", "inference", "kv cache", "memory", "moe"]
source: docs/RAG/Collect RAG/02_mindstudio/deepseek-v4-vision-model-kv-cache-efficiency-maze-navigation-benchmark.md
source_anchor: ""
source_lines: [1, 50]
sha256: 768c5ba40e4b0aa96e356dc05a91735adeb6b288876da0433b9e13fb6fcde43d
---

# DeepSeek V4 Vision Model: 10x KV-Cache Efficiency and 67% Maze Navigation vs GPT-5.4's 50%

## Metadata

- **Source** : https://www.mindstudio.ai/blog/deepseek-v4-vision-model-kv-cache-efficiency-maze-navigation-benchmark
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article details DeepSeek's vision model (built on the V4 Flash backbone): ~90 KV-cache entries per 80×80 image vs ~870 for Claude Sonnet 4.6 (~10x memory reduction), while scoring 67% on maze navigation vs GPT-5.4's 50% and Gemini Flash 3's 49%.

The compression is a multi-stage pipeline, not a single trick: a 756×756 image (571,000 pixels) is processed by the "DeepSeek Vision Transformer" (built from scratch for arbitrary resolution) using 14×4 patches → ~2,916 patch tokens → a 3×3 spatial compression along the channel dimension collapses nine adjacent patches into one → ~324 tokens → compressed sparse attention from the V4 paper applies another 4x compression to the KV cache → ~81 entries (~7,000x total compression from raw pixels). The language backbone is DeepSeek V4 Flash: a 284B-parameter MoE with 13B active parameters at inference — frontier-quality reasoning at a fraction of activation cost.

The "Thinking with Visual Primitives" paper argues multimodal models have two gaps: the perception gap (can't always see fine detail — what 2024 high-res cropping/dynamic patching targeted) and the reference gap (language is too imprecise to point at things reliably as reasoning chains lengthen). DeepSeek's solution: make spatial coordinates first-class tokens in the chain of thought — the model emits inline `<ref>person_3</ref><box>(x1,y1),(x2,y2)</box>` tokens mid-thought, anchoring reasoning like a human pointing a finger. This explains the maze navigation result: path tracing and topological reasoning are exactly where language is worst at trajectory description; a model that can mark cells and reason forward from marks doesn't lose its place.

Two years of the same story (cheapest representation that still works): DeepSeek VL (March 2024, hybrid SigLIP+SAM encoder); Janus (Oct 2024, decoupled encoders for understanding vs generation); VL2 (Dec 2024, MoE + multi-head latent attention; 1B-activated version scored 80.9 on OCR Bench); DeepSeek OCR (Oct 2025 — 1,000 text tokens rendered as image → 100 vision tokens reconstruct text at 97% accuracy, 10x long-context compression; Karpathy: "the tokenizer must go, pixels may be better inputs to language models than text").

Benchmark honesty: the paper's footnote scopes results to evaluation dimensions directly relevant to its research focus — DeepSeek claims to beat GPT-5.4 on visually grounded reasoning (maze navigation, path tracing, counting in dense scenes), not across the board. Gemini Flash 3 is still ahead on raw count QA. Three explicit limitations: resolution-bound, visual primitives mode must be triggered explicitly, and point-based topological reasoning doesn't generalize across all scenarios.

Practical consequences: the 10x KV-cache reduction means ~9x more concurrent image requests in the same memory budget or ~9x fewer GPUs for the same throughput. Pricing context: DeepSeek V4 Flash at $1.74/M input and $3.48/M output vs Claude Opus 4.7 at $5/M input and $25/M output, or GPT-5.5 at $5/$30 — the KV-cache efficiency multiplies the cost advantage when serving images at volume.

Training architecture: two separate specialist models first — one for "thinking with grounding" (bounding boxes), one for "thinking with pointing" (point coordinates) — each with its own SFT pass and GRPO RL with three reward heads (format, quality, accuracy), then a unified reward-free distillation merges them into a single student model. Rollout status: vision mode began rolling out in limited form in the app/web alongside fast and expert modes; the OpenRouter ID for the text backbone is deepseek/deepseek-v4-flash.

## Key points

- ~90 vs ~870 vs ~1,000 KV-cache entries per 80×80 image for DeepSeek / Sonnet 4.6 / Flash 3 — a structural ~10x efficiency difference.
- 67% maze navigation vs GPT-5.4's 50% and Gemini Flash 3's 49% — visual primitives (inline bounding-box tokens) solve the reference gap.
- Compression pipeline: 2,916 patch tokens → 324 (3×3 spatial) → ~81 KV entries (4x sparse attention) ≈ 7,000x total.
- Backbone: DeepSeek V4 Flash, 284B MoE / 13B active.
- Training: separate grounding/pointing specialists with GRPO + three reward heads, merged via distillation.
- Honest scoping and three admitted limitations (resolution-bound, explicit triggering, limited topological generalization).

## Technical data / figures

| Metric | Value |
|---|---|
| KV entries (80×80) | DeepSeek ~90, Sonnet 4.6 ~870, Flash 3 ~1,000 |
| Maze navigation | DeepSeek 67%, GPT-5.4 50%, Flash 3 49%, Sonnet 4.6 49% |
| Patch pipeline | 571k px → 2,916 → 324 → ~81 KV |
| Total compression | ~7,000x |
| DeepSeek V4 Flash pricing | $1.74/M in, $3.48/M out |
| Claude Opus 4.7 / GPT-5.5 pricing | $5/$25 and $5/$30 |

## Why this source matters for the RAG

Provides the benchmark and efficiency numbers (KV cache, maze navigation, compression pipeline) that justify using DeepSeek V4 Vision in spatial/document-reasoning RAG pipelines at a fraction of the cost. Directly useful for model selection and cost forecasting in vision-heavy retrieval systems.
