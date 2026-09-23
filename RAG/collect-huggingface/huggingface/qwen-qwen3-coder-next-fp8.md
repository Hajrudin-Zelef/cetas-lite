---
id: collect-huggingface/huggingface/qwen-qwen3-coder-next-fp8
title: "Qwen3-Coder-Next-FP8 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "Hugging Face", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["fp8", "qwen", "agent", "agentic", "agents", "apache", "attention", "claude", "cost", "embedding", "inference", "license"]
source: docs/RAG/Collect RAG/03_huggingface/Qwen-Qwen3-Coder-Next-FP8.md
source_anchor: ""
source_lines: [1, 48]
sha256: 4bac1bb46e3a1012c87e05fec34ac246191a72104dcab9febfaaf16429f15491
---

# Qwen3-Coder-Next-FP8 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/Qwen/Qwen3-Coder-Next-FP8
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Qwen3-Coder-Next-FP8 is an open-weight language model designed specifically for coding agents and local development, released under Apache-2.0. It is the FP8-quantized checkpoint of Qwen3-Coder-Next, using fine-grained FP8 quantization with a block size of 128. It has 80B total parameters with only 3B activated per token, achieving performance comparable to models with 10-20x more active parameters, making it highly cost-effective for agent deployment.

The model is a causal language model with 79B non-embedding parameters, hidden dimension 2048, 48 layers, and a hybrid layout of 12 × (3 × (Gated DeltaNet → MoE) → 1 × (Gated Attention → MoE)). The MoE has 512 experts with 10 activated plus 1 shared expert; expert intermediate dimension is 512. Gated Attention uses 16 Q heads and 2 KV heads; Gated DeltaNet uses 32 V heads and 16 QK heads. Context length is 262,144 tokens natively. It supports only non-thinking mode and does not generate thinking blocks.

Agentic capabilities: excels at long-horizon reasoning, complex tool usage, and recovery from execution failures. Its 256k context and adaptability to scaffold templates enable integration with CLI/IDE platforms such as Claude Code, Qwen Code, Qoder, Kilo, Trae, and Cline. Deployment: SGLang >=v0.5.8 or vLLM >=0.15.0, with `--tool-call-parser qwen3_coder`. Recommended sampling: temperature=1.0, top_p=0.95, top_k=40. Works with Ollama, LM Studio, MLX-LM, llama.cpp, and KTransformers for local use. ~1.16M downloads per month.

## Key points

- Coding-agent model: 80B total / 3B activated MoE.
- FP8-quantized (fine-grained, block 128), near BF16 performance.
- 512 experts (10 active + 1 shared); 262,144 native context.
- Non-thinking mode only; strong tool calling (`qwen3_coder` parser).
- Integrates with Claude Code, Qwen Code, Qoder, Trae, Cline, etc.
- Apache-2.0; served via SGLang, vLLM, Ollama, LM Studio, KTransformers.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 80B (79B non-embedding) |
| Active parameters | 3B |
| Architecture | Hybrid Gated DeltaNet + Gated Attention + MoE |
| Layers | 48 |
| Hidden dimension | 2048 |
| Experts | 512 (10 active + 1 shared) |
| Expert intermediate dimension | 512 |
| Context length | 262,144 (256K) |
| Quantization | FP8 (block size 128) |
| License | Apache-2.0 |
| Recommended sampling | temp=1.0, top_p=0.95, top_k=40 |
| Monthly downloads | ~1.16M |

## Why this source matters for the RAG

Qwen3-Coder-Next-FP8 is a key reference for open-weight coding agents with an extremely low activation footprint (3B active), directly relevant to cost-efficient agent deployment. Its FP8 quantization and scaffolding support details ground the knowledge base on running local/agentic coding models. It complements the coding-agent and local-inference coverage of the RAG.
