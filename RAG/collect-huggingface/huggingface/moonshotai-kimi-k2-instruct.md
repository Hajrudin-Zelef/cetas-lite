---
id: collect-huggingface/huggingface/moonshotai-kimi-k2-instruct
title: "Kimi K2 Instruct - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agentic", "attention", "benchmarks", "claude", "deepseek", "fp8", "gemini", "license", "mit license", "moe", "opus 4"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-K2-Instruct.md
source_anchor: ""
source_lines: [1, 51]
sha256: f211bc3fd7575fdb640cb52d0dcf87532eac336747b4b8fea0781d4868af3c4a
---

# Kimi K2 Instruct - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-K2-Instruct
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi K2 Instruct is Moonshot AI's state-of-the-art Mixture-of-Experts language model: 1 trillion total parameters with 32 billion activated parameters, trained with the Muon optimizer on 15.5T tokens with zero training instability. It is the post-trained "drop-in" instruction model of the Kimi K2 family (sibling of Kimi-K2-Base), described as a reflex-grade model without long thinking, designed for general-purpose chat and agentic experiences. Architecture: 61 layers (1 dense), MLA attention (hidden 7168, 64 heads), MoE hidden 2048 per expert, 384 experts with 8 selected per token plus 1 shared expert, SwiGLU activation, 160K vocabulary, 128K context. Checkpoints are stored in block-fp8 format. License: Modified MIT. Downloads ~239K/month. A newer version (Kimi-K2-Instruct-0905) exists.

Evaluation (instruction model) shows strong agentic and coding performance: LiveCodeBench v6 53.7, OJBench 27.1, MultiPL-E 85.7, SWE-bench Verified 65.8 (agentic single attempt) / 71.6 (multi-attempt), SWE-bench Multilingual 47.3, TerminalBench 25.0 (Terminus), Aider-Polyglot 60.0, Tau2 retail 70.6 / airline 56.5 / telecom 65.8, AceBench 76.5, AIME 2024 69.6, AIME 2025 49.5, MATH-500 97.4, GPQA-Diamond 75.1, SuperGPQA 57.2, MMLU 89.5, MMLU-Pro 81.1, IFEval 89.8, SimpleQA 31.0, Livebench 76.4, HLE 4.7 — with many scores above DeepSeek-V3-0324, Qwen3-235B-A22B, Claude Sonnet 4, Claude Opus 4, GPT-4.1 and Gemini 2.5 Flash. Base-model results are also tabulated (e.g., MMLU 87.8, GPQA-Diamond 48.1, EvalPlus 80.3).

Usage: chat via OpenAI-compatible API with recommended temperature 0.6 and the default system prompt "You are Kimi, an AI assistant created by Moonshot AI." Tool calling is enabled by passing a tools list and executing the native Kimi K2 tool-parsing loop. Deployment: vLLM, SGLang, KTransformers, TensorRT-LLM; the Anthropic-compatible API maps temperature as `real = request * 0.6`. Changelog: chat template robustness fixes (2025.7.15/7.18) and `name` field support with a standalone chat template file (2025.8.11).

## Key points

- Kimi K2 Instruct: 1T total / 32B active MoE, Muon optimizer, 15.5T tokens.
- Reflex-grade general-purpose + agentic chat model (no long thinking).
- 128K context; MLA; 384 experts (8/token); block-fp8 checkpoints.
- Modified MIT license; newer 0905 revision available.
- Strong coding/agentic results: SWE-bench Verified 65.8/71.6, LiveCodeBench 53.7, TerminalBench 25.0, AIME 2024 69.6, GPQA-D 75.1.
- Recommend temperature 0.6; native tool-calling loop.
- Deploy via vLLM / SGLang / KTransformers / TRT-LLM.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-K2-Instruct |
| Architecture | MoE, 61 layers (1 dense), MLA, hidden 7168, 64 heads |
| Total params | 1T |
| Active params | 32B |
| Experts | 384 (8/token) + 1 shared |
| Context length | 128K |
| Vocab | 160K |
| Optimizer | Muon (MuonClip) |
| Pre-training | 15.5T tokens |
| Quantization | block-fp8 |
| License | modified-mit |
| Key benchmarks | SWE-bench Verified 65.8/71.6, SWE-bench Multilingual 47.3, LiveCodeBench v6 53.7, AIME 2024 69.6, MATH-500 97.4, GPQA-D 75.1, MMLU 89.5, IFEval 89.8 |
| Downloads/month | ~239,492 |

## Why this source matters for the RAG

This card is a primary reference for Kimi K2 Instruct, a landmark open 1T-class agentic MoE (the first in the K2.x series), documenting Muon-scale training, agentic/coding benchmarks, tool-calling API and multi-engine deployment. It anchors historical and comparative knowledge of open frontier MoE models for the RAG.
