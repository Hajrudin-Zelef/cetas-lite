---
id: collect-huggingface/huggingface/qwen-qwen3-8-2-4t-a95b-fp8
title: "Qwen3.8-2.4T-A95B-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "qwen", "agentic", "agents", "attention", "benchmark", "benchmarks", "embedding", "fable 5", "license", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.8-2.4T-A95B-FP8.md
source_anchor: ""
source_lines: [1, 54]
sha256: 66c5e06c6f2e0365b59a1d508c9b2465b512fdc9b28072737098b62255922b73
---

# Qwen3.8-2.4T-A95B-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.8-2.4T-A95B-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.8-2.4T-A95B-FP8 is the FP8-quantized release of the Qwen3.8-Max-class open model, the first time Qwen brings a Qwen-Max-class model to open release. It is a text-only Causal Language Model (no vision) with 2.4T total parameters and 95B activated, licensed under the qwen3.8-max license, tagged qwen3_5_moe_text/conversational/fp8, with 8,319 monthly downloads. Hidden dim 8192, 92 layers, token embedding 248,320 (padded). Hidden layout: 23 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)). Gated DeltaNet: 128 V heads / 16 QK heads, dim 128. Gated Attention: 64 Q / 4 KV heads, dim 256, RoPE dim 64. MoE: 512 experts, 10 routed + 1 shared activated, expert intermediate dim 2048. MTP (Multi-Token Prediction) trained with multiple steps. Context length 262,144 natively, extensible up to 1,010,000 tokens.

The quantization is fine-grained FP8 with block size 128, claimed nearly identical to the original model's performance. It is a text-only model that requires thinking mode for all interactions: multimodal inputs are not supported and thinking cannot be disabled; every response begins with reasoning in ` thinking\n... response\n\n`. `reasoning_effort` supports `xhigh` (default), `medium`, `low`; `preserve_thinking` is on by default. Recommended sampling: temperature 1.0, top_p 0.95, top_k 20, min_p 0.0, presence_penalty 0.0, repetition_penalty 1.0. Best practices: within 1M context, set reasoning output max to 262,144 tokens and final response max to 131,072 tokens.

Benchmarks (as Qwen3.8-Max vs Opus 4.8, Fable 5, GPT 5.6 Sol max, Qwen3.7-Max): Terminal Bench 2.1 86.6, SWE-bench Pro 67.7, DeepSWE 1.1 56.6, NL2Repo-Bench 55.9, FrontierSWE 73.5, PaperBench 93.0, AndroidBench 75.1, QwenSWEBench 80.7, QwenReactBench 1724, CoWorkBench 74.8, WorkSpaceBench 67.7, JobBench 53.4, SkillsBench 70.2, Agents' Last Exam 27.0/52.4, Automation-Bench 27.3, Toolathlon Verified 72.5, WideSearch 81.9, HLE w/ tools 56.2, GPQA Diamond 92.6, HLE 43.6, IFBench 82.8, $OneMillion-Bench 52.5, HealthBench 60.2, PLawBench 73.2, PRBench-Legal 57.6, PRBench-Finance 58.3, MRCR v2 256K 92.9, LongBench v2 66.3. The official Qwen3.8-Max API (vision, non-thinking, 1M context, built-in tools) is based on this model. Serving via vLLM, SGLang, TokenSpeed; hub tensors BF16 + F8_E4M3.

## Key points

- FP8 (block 128) quantization of the 2.4T-param Qwen3.8-Max-class model; 95B activated.
- Text-only; thinking mode required and cannot be disabled; no multimodal input.
- Architecture: 92 layers, 512-expert MoE (10 routed + 1 shared), Gated DeltaNet + Gated Attention, MTP.
- Native context 262,144, extensible to 1,010,000; qwen3.8-max license.
- reasoning_effort xhigh/medium/low; preserve_thinking on by default.
- Benchmarks: Terminal Bench 2.1 86.6, SWE Pro 67.7, PaperBench 93.0, GPQA-D 92.6, IFBench 82.8, MRCR v2 92.9.
- Recommended sampling temp 1.0, top_p 0.95; serve via vLLM/SGLang/TokenSpeed.
- 8,319 downloads/month; base for Qwen3.8-Max API.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Qwen |
| Model name | Qwen3.8-2.4T-A95B-FP8 |
| Type | Text-only Causal LM (FP8 quantized) |
| Total params | 2.4T |
| Activated params | 95B |
| Layers | 92 (hidden 8192) |
| MoE | 512 experts; 10 routed + 1 shared; expert dim 2048 |
| Attention | Gated DeltaNet + Gated Attention (hybrid) |
| MTP | Multi-Token Prediction |
| Context length | 262,144 native; up to 1,010,000 |
| Quantization | Fine-grained FP8, block 128 (F8_E4M3) |
| License | qwen3.8-max |
| Thinking | Required (cannot be disabled); reasoning_effort xhigh/medium/low |
| Key benchmarks | Terminal Bench 2.1 86.6; SWE Pro 67.7; PaperBench 93.0; GPQA-D 92.6; MRCR v2 92.9 |
| Serving | vLLM, SGLang, TokenSpeed |
| Downloads/month | 8,319 |

## Why this source matters for the RAG

This card documents the FP8 quantization and serving of the largest open-weight Qwen model (2.4T total / 95B active), with an exhaustive agentic and reasoning benchmark table. It is essential for retrieval on extreme-scale MoE deployment, FP8 quantization fidelity, and frontier open-model capabilities.
