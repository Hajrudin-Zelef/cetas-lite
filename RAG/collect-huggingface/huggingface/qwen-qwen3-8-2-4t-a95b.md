---
id: collect-huggingface/huggingface/qwen-qwen3-8-2-4t-a95b
title: "Qwen3.8-2.4T-A95B - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["qwen", "agentic", "agents", "attention", "benchmark", "benchmarks", "embedding", "fable 5", "fp8", "license", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3.8-2.4T-A95B.md
source_anchor: ""
source_lines: [1, 54]
sha256: 47cc51cdbe9800aaa312c4df521d5f985684b97350eefb9f86619fbfb6b114c8
---

# Qwen3.8-2.4T-A95B - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3.8-2.4T-A95B
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3.8-2.4T-A95B is the full-precision open release of the Qwen3.8-Max-class model — the first time Qwen brings a Qwen-Max-class model to open release. It is a text-only Causal Language Model (no multimodal input) with 2.4T total parameters and 95B activated, licensed under the qwen3.8-max license, tagged qwen3_5_moe_text/conversational, with 54,427 monthly downloads. Architecture: hidden dim 8192, 92 layers, token embedding 248,320 (padded); hidden layout 23 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)). Gated DeltaNet: 128 V heads / 16 QK heads, dim 128. Gated Attention: 64 Q / 4 KV heads, dim 256, RoPE dim 64. MoE: 512 experts, 10 routed + 1 shared activated, expert intermediate dim 2048. MTP (Multi-Token Prediction) is trained with multiple steps. Context length 262,144 natively, extensible up to 1,010,000 tokens. Hub weights are BF16.

This is a text-only model requiring thinking mode for all interactions: multimodal inputs are not supported and thinking cannot be disabled; every response begins with reasoning in ` thinking\n... response\n\n`. `reasoning_effort` supports `xhigh` (default), `medium`, `low`; `preserve_thinking` is enabled by default. Recommended sampling: temperature 1.0, top_p 0.95, top_k 20, min_p 0.0, presence_penalty 0.0, repetition_penalty 1.0. Within the 1M context, it is recommended to set reasoning max output to 262,144 tokens and final response max to 131,072 tokens.

Benchmarks (as Qwen3.8-Max vs Opus 4.8, Fable 5, GPT 5.6 Sol max, Qwen3.7-Max): Terminal Bench 2.1 86.6, SWE-bench Pro 67.7, DeepSWE 1.1 56.6, NL2Repo-Bench 55.9, FrontierSWE 73.5, MLS-Bench-Lite 41.0, PaperBench 93.0, AndroidBench 75.1, QwenSWEBench 80.7, QwenQoderBench 58.4, QwenReactBench 1724, QwenSVGBench 1713, CoWorkBench 74.8, WorkSpaceBench 67.7, JobBench 53.4, SkillsBench 70.2, Agents' Last Exam 27.0/52.4, Automation-Bench 27.3, Toolathlon Verified 72.5, WideSearch 81.9, HLE w/ tools 56.2, GPQA Diamond 92.6, HLE 43.6, IFBench 82.8, $OneMillion-Bench 52.5, HealthBench 60.2, PLawBench 73.2, PRBench-Legal 57.6, PRBench-Finance 58.3, MRCR v2 256K 92.9, LongBench v2 66.3. The official Qwen3.8-Max API (vision input, non-thinking support, 1M context by default, built-in tools) is based on this model. Serving via Transformers, vLLM, SGLang, TokenSpeed.

## Key points

- First Qwen-Max-class open release: 2.4T total / 95B activated, BF16.
- Text-only; thinking mode required (cannot be disabled); no multimodal input.
- 92 layers, hidden 8192, 512-expert MoE (10 routed + 1 shared), Gated DeltaNet + Gated Attention, MTP.
- Native context 262,144, extensible to 1,010,000; qwen3.8-max license.
- reasoning_effort xhigh/medium/low; preserve_thinking default on.
- Benchmarks: Terminal Bench 2.1 86.6, SWE Pro 67.7, PaperBench 93.0, GPQA-D 92.6, IFBench 82.8.
- Recommended sampling temp 1.0, top_p 0.95; serve via Transformers/vLLM/SGLang/TokenSpeed.
- 54,427 downloads/month; base for the Qwen3.8-Max API; FP8 variant available separately.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | Qwen |
| Model name | Qwen3.8-2.4T-A95B |
| Type | Text-only Causal LM |
| Total params | 2.4T |
| Activated params | 95B |
| Layers | 92 (hidden 8192) |
| MoE | 512 experts; 10 routed + 1 shared; expert dim 2048 |
| Attention | Gated DeltaNet + Gated Attention (hybrid) |
| MTP | Multi-Token Prediction |
| Context length | 262,144 native; up to 1,010,000 |
| Weight format | BF16 |
| License | qwen3.8-max |
| Thinking | Required (cannot be disabled); reasoning_effort xhigh/medium/low |
| Key benchmarks | Terminal Bench 2.1 86.6; SWE Pro 67.7; PaperBench 93.0; GPQA-D 92.6; MRCR v2 92.9 |
| Serving | Transformers, vLLM, SGLang, TokenSpeed |
| Downloads/month | 54,427 |

## Why this source matters for the RAG

This card is the canonical reference for Qwen's largest open-weight release, providing full architecture specs, licensing, and an exhaustive benchmark table across coding, agentic, and reasoning domains. It is essential for retrieval on frontier open MoE models, million-token context, and thinking-mode-only deployment.
