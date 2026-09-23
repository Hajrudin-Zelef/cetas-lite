---
id: collect-mindstudio/mindstudio/run-qwen3-8-27b-obliterated-locally
title: "Run Qwen3.8-27B-OBLITERATED Locally: GGUF Sizes, VRAM, and Settings"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Apple"]
dates: ["2026-09-23"]
keywords: ["gguf", "agentic", "alignment", "consumer", "cost", "cyber", "datacenter", "fine-tuning", "gpu", "llama", "llama.cpp", "parameters"]
source: docs/RAG/Collect RAG/02_mindstudio/run-qwen3-8-27b-obliterated-locally.md
source_anchor: ""
source_lines: [1, 63]
sha256: 6d154bef496d82df23868440d74431b9947d1ac597b54b0969c6c40f70b7dffc
---

# Run Qwen3.8-27B-OBLITERATED Locally: GGUF Sizes, VRAM, and Settings

## Metadata

- **Source** : https://www.mindstudio.ai/blog/run-qwen3-8-27b-obliterated-locally
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This guide covers how to run the uncensored **Qwen3.8-27B-OBLITERATED** model locally: **GGUF** quant sizes, **VRAM** needs, and the exact settings that keep it from looping. The model is a modified version of Alibaba's **Qwen3.8-27B** with refusal behavior surgically removed through **abliteration**. The current release, **V3**, targets both hard refusals ("I cannot help with that") and softer safety-lecture deflections, aiming to answer restricted queries with substance rather than a warning label. It's distributed as GGUF files for llama.cpp-based tools like **Ollama** and **LM Studio**, as full **bfloat16 safetensors** for Transformers, and **MLX support is pending** upstream architecture compatibility.

Abliteration is a weight-space editing technique, not fine-tuning. It identifies **"refusal directions"** — specific vectors in the model's internal representations correlating with refusal behavior — and projects them out of the weight matrices, removing the mechanism that triggers refusals while leaving knowledge and reasoning intact. The project went through three iterations. **V1** used a single aggressive **SVD** pass removing five directions: it killed refusals completely but cost **6 percentage points of MMLU**. **V2** introduced **complementary abliteration blending**: running two methods (**aggressive SVD** and **LEACE**, which minimizes mutual information) that fail differently, then blending their weights at a **60/40 ratio**. SVD finds refusal directions well but damages capability; LEACE preserves capability but removes refusals less thoroughly. Blending canceled weaknesses, cutting the MMLU cost to **0.3 points**, but V2 still produced safety-lecture deflections. **V3** added **iterative stacking** (always refine the current best model rather than restarting from stock) and a **targeted corpus** (focused prompts aimed at specific deflection categories). V3 applies a gentle refinement pass, then a targeted surgery pass, then blends the two **50/50**, eliminating soft deflections as well as hard refusals at a cost of **2.1 MMLU points** versus stock.

The GGUF quantizations span a wide range, the main lever for fitting a 27B model onto consumer hardware: **Q8_0 ~27 GB** (maximum quality, high-VRAM/multi-GPU), **Q6_K ~21 GB**, **Q5_K_M ~18 GB** (solid for 24 GB cards), **Q4_K_M ~16 GB** (the commonly recommended sweet spot), **IQ4_XS ~14 GB** (compact, experimental), **Q3_K_M ~13 GB** (lower-VRAM), and **Q2_K ~11 GB** (minimum viable). As a rule, VRAM should be at least the GGUF file size plus headroom for context. A 16 GB card handles Q4_K_M with a few thousand tokens of context; a 24 GB card runs Q5_K_M or Q6_K with more headroom. The full bfloat16 release is ~**54 GB across 29 shards**, aimed at multi-GPU/datacenter Transformers workflows. Apple Silicon MLX is waiting on upstream **mlx_lm** support for the Qwen3.5 architecture.

Settings matter more than usual. **Temperature 0** (greedy decoding) produces the most complete, code-rich outputs; anything above 0.5 degrades noticeably. **Repetition penalty 1.15** is described as essential — without it, greedy decoding loops on imports or boilerplate (1.10-1.12 gives tighter output). **Max new tokens: 2048+**. **No system prompt** (A/B tested; adding one can reintroduce refusals). **Thinking mode off** (recommended; V3's chat template prefills an empty thinking block). Sampling parameters (top_p, top_k, min_p) are not needed. GGUF users must load the bundled chat template (`--jinja` in llama.cpp, or the built-in template in Ollama/LM Studio).

For **agentic use**, the model card recommends separate adjustments: keep repetition_penalty at 1.15 (critical for breaking loops), raise temperature to **0.1-0.3** to escape deterministic loops, cap max_tokens per turn at **1024-2048**, and summarize context after roughly **10 turns**. On capability tradeoff: stock Qwen3.8-27B scores **84.46%** on MMLU (0-shot, lm-eval-harness, ~5,700 questions); V3 lands at **82.33%** (-2.12 points). The hit is uneven: **STEM dropped most (-3.3pp)**, humanities barely moved (-1.0pp), and some subjects (philosophy, European history) scored higher, while abstract algebra and formal logic saw larger drops — suggesting refusal directions partially overlap with structured reasoning pathways. The card reports **20/20** successful completions on code/cyber prompts and **7/8** on advanced real-world tasks (async refactoring, security code review, Kubernetes debugging), matching stock's 7/8; multi-tool chaining failed for both.

## Key points

- V3 removes both hard refusals and soft safety-lecture deflections from stock Qwen3.8-27B, scoring 20/20 on code/cyber prompts.
- Capability cost is modest: MMLU 82.33% vs 84.46% (-2.1pp), concentrated in STEM (-3.3pp).
- GGUF ranges from Q2_K (~11 GB) to Q8_0 (~27 GB); Q4_K_M (~16 GB) is the recommended sweet spot.
- Essential settings: temperature 0, repetition_penalty 1.15, no system prompt, thinking off, max new tokens 2048+.
- GGUF users must load the bundled chat template (`--jinja` / built-in) to skip the thinking block.
- Agentic use: temperature 0.1-0.3, cap per-turn tokens, summarize context every ~10 turns.
- It's framed as a research tool for alignment researchers, red-teamers, and safety evaluators.

## Technical data / figures

| Quant | Size | Notes |
|---|---|---|
| Q8_0 | ~27 GB | Maximum quality; high-VRAM/multi-GPU |
| Q6_K | ~21 GB | Strong quality/size balance |
| Q5_K_M | ~18 GB | Solid for 24 GB cards |
| Q4_K_M | ~16 GB | Recommended sweet spot |
| IQ4_XS | ~14 GB | Compact, experimental |
| Q3_K_M | ~13 GB | Lower-VRAM setups |
| Q2_K | ~11 GB | Minimum viable quality |

| Other figure | Value |
|---|---|
| Stock MMLU | 84.46% |
| V3 MMLU | 82.33% (-2.12pp) |
| V1 MMLU cost | -6pp |
| V2 MMLU cost | -0.3pp |
| STEM drop (V3) | -3.3pp |
| Humanities drop (V3) | -1.0pp |
| V2 blend | 60/40 (SVD + LEACE) |
| V3 blend | 50/50 |
| Recommended decoding | temp 0, rep. penalty 1.15, no system prompt |
| Agentic decoding | temp 0.1-0.3, max_tokens 1024-2048 |
| Full bfloat16 | ~54 GB, 29 shards |
| MLX | Pending |

## Why this source matters for the RAG

It is the operational companion to the OBLITERATED V3 explainer, giving exact quantization sizes, VRAM guidance, and decoding settings. It is highly useful for procedural questions about running uncensored models locally and avoiding common looping failures.
