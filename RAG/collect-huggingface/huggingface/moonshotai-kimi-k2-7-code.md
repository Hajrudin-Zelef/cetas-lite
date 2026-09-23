---
id: collect-huggingface/huggingface/moonshotai-kimi-k2-7-code
title: "Kimi K2.7 Code - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agent", "agentic", "attention", "benchmarks", "claude", "int4", "license", "mcp", "mit license", "moe", "multimodal"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-K2.7-Code.md
source_anchor: ""
source_lines: [1, 57]
sha256: 8b4f028a9267a1bc177112590f9100938017a9612562ba55ba66b7f0afe45520
---

# Kimi K2.7 Code - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-K2.7-Code
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi K2.7 Code is a coding-focused agentic model from Moonshot AI, built on top of Kimi K2.6, designed to excel at real-world long-horizon software engineering tasks while improving token efficiency — reducing thinking-token usage by approximately 30% compared to K2.6. It is a Mixture-of-Experts (MoE) model with 1T total parameters and 32B activated parameters, 61 layers (1 dense), attention hidden dimension 7168, MoE hidden dimension 2048 per expert, 64 attention heads, 384 experts with 8 selected per token plus 1 shared expert, a 160K vocabulary, 256K context length, MLA attention, SwiGLU activation, and a 400M-parameter MoonViT vision encoder — making it an Image-Text-to-Text model supporting image and (experimental) video input.

Benchmarks: Kimi Code Bench v2 62.0 (vs K2.6 50.9, GPT-5.5 69.0, Claude Opus 4.8 67.4), Program Bench 53.6, MLS-Bench Lite 35.1, Kimi Claw 24/7 Bench 46.9, MCP Atlas 76.0, MCP Mark Verified 81.1. It is a thinking model that forces thinking and preserve_thinking modes (retaining full reasoning content across multi-turn interactions, enabled by default), uses interleaved thinking and multi-step tool calling, and is intended to be run with the Kimi Code CLI agent framework (temp 1.0, top_p 0.95, 262,144-token context). Native INT4 quantization is supported using the same method as Kimi-K2-Thinking.

The model is released under a Modified MIT License (both weights and code), with third-party notices provided. Deployment is supported on vLLM, SGLang, and KTransformers with transformers >= 4.57.1, < 5.0.0; it shares the K2.5/K2.6 architecture so existing deployment guides apply. The official API is available at platform.moonshot.ai with OpenAI/Anthropic-compatible endpoints. Hub reports 1T params (BF16/F32/I32), about 104,772 monthly downloads, 28 community quantizations, and 10 finetunes.

## Key points

- Coding-focused agentic MoE model built on Kimi K2.6, cutting thinking-token usage ~30%.
- 1T total / 32B activated parameters, 384 experts (8 active + 1 shared), 256K context.
- Multimodal: 400M-param MoonViT vision encoder supports image and experimental video input.
- Forces thinking + preserve_thinking; interleaved thinking and multi-step tool calling.
- Strong agentic/coding benchmarks: Kimi Code Bench v2 62.0, MCP Mark Verified 81.1.
- Native INT4 quantization (same scheme as Kimi-K2-Thinking).
- Modified MIT license for weights and code.
- Runs on vLLM, SGLang, KTransformers; transformers>=4.57.1,<5.0.0.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-K2.7-Code |
| Architecture | MoE (MLA + MoE FFN) |
| Total params | 1T |
| Activated params | 32B |
| Layers | 61 (1 dense) |
| Experts | 384 (8 selected/token) + 1 shared |
| Hidden dimension | 768 (attention), 2048 per expert MoE |
| Attention heads | 64 |
| Vocabulary size | 160K |
| Context length | 256K |
| Activation | SwiGLU |
| Vision encoder | MoonViT, 400M params |
| Modalities | Text, Image, Video (experimental) |
| License | Modified MIT |
| Weight formats | BF16, F32, I32 |
| Benchmarks | Kimi Code Bench v2 62.0, Program Bench 53.6, MCP Mark Verified 81.1 |
| Deployment | vLLM, SGLang, KTransformers |
| Downloads/month | 104,772 |

## Why this source matters for the RAG

This card documents a flagship open-weight coding agent model with detailed architecture specs, comparison benchmarks against GPT-5.5 and Claude Opus 4.8, and its distinctive forced-thinking/preserve-thinking design. It is key reference material for retrieval on agentic coding models, MoE architectures, thinking-token efficiency, and open-weight model licensing.
